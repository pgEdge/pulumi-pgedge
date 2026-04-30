package pgedge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// schemaDoc is a partial view of schema.json sufficient for shape assertions.
type schemaDoc struct {
	Resources map[string]resourceDoc `json:"resources"`
	Functions map[string]functionDoc `json:"functions"`
}

type resourceDoc struct {
	InputProperties  map[string]propertyDoc `json:"inputProperties"`
	RequiredInputs   []string               `json:"requiredInputs"`
	Properties       map[string]propertyDoc `json:"properties"`
	Required         []string               `json:"required"`
	ReplaceOnChanges []string               `json:"replaceOnChanges"`
}

type functionDoc struct {
	Outputs struct {
		Properties map[string]propertyDoc `json:"properties"`
		Required   []string               `json:"required"`
	} `json:"outputs"`
}

type propertyDoc struct {
	Type string `json:"type"`
}

func loadSchema(t *testing.T) schemaDoc {
	t.Helper()
	path := filepath.Join("cmd", "pulumi-resource-pgedge", "schema.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var doc schemaDoc
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return doc
}

// resourceShape locks in the user-facing input/output surface of each resource.
// Fields are sorted; the test compares as sets. Adding or removing a field —
// or flipping required-ness — breaks the test on purpose, since each is a
// user-visible API change.
type resourceShape struct {
	token           string
	inputProperties []string // exact set
	requiredInputs  []string // exact set; "name" is autonamed via SetAutonaming so usually absent
	properties      []string // exact set (inputs + computed outputs)
	required        []string // exact set
	// arrayProps names properties that must serialize as `"type": "array"` in
	// the generated schema. Catches regressions in the bridge-metadata
	// `maxItemsOne: false` overrides — without those, the tfbridge collapses
	// single-element TF lists to scalars.
	arrayProps []string
}

var resourceShapes = []resourceShape{
	{
		token:           "pgedge:index/sSHKey:SSHKey",
		inputProperties: []string{"name", "publicKey"},
		requiredInputs:  []string{"publicKey"},
		properties:      []string{"createdAt", "name", "publicKey"},
		required:        []string{"createdAt", "name", "publicKey"},
	},
	{
		token:           "pgedge:index/cloudAccount:CloudAccount",
		inputProperties: []string{"credentials", "description", "name", "type"},
		requiredInputs:  []string{"credentials", "type"},
		properties:      []string{"createdAt", "credentials", "description", "name", "type"},
		required:        []string{"createdAt", "credentials", "name", "type"},
	},
	{
		token:           "pgedge:index/backupStore:BackupStore",
		inputProperties: []string{"cloudAccountId", "name", "region"},
		requiredInputs:  []string{"cloudAccountId", "region"},
		properties: []string{
			"cloudAccountId", "cloudAccountType", "clusterIds", "createdAt",
			"name", "properties", "region", "status",
		},
		required: []string{
			"cloudAccountId", "cloudAccountType", "clusterIds", "createdAt",
			"name", "properties", "region", "status",
		},
		arrayProps: []string{"clusterIds"},
	},
	{
		token: "pgedge:index/cluster:Cluster",
		inputProperties: []string{
			"backupStoreIds", "capacity", "cloudAccountId", "firewallRules",
			"name", "networks", "nodeLocation", "nodes", "regions",
			"resourceTags", "sshKeyId",
		},
		requiredInputs: []string{
			"cloudAccountId", "networks", "nodeLocation", "nodes", "regions",
		},
		properties: []string{
			"backupStoreIds", "capacity", "cloudAccountId", "createdAt",
			"firewallRules", "name", "networks", "nodeLocation", "nodes",
			"regions", "resourceTags", "sshKeyId", "status",
		},
		required: []string{
			"backupStoreIds", "capacity", "cloudAccountId", "createdAt",
			"name", "networks", "nodeLocation", "nodes", "regions",
			"resourceTags", "status",
		},
		arrayProps: []string{"backupStoreIds", "firewallRules", "networks", "nodes", "regions"},
	},
	{
		token: "pgedge:index/database:Database",
		inputProperties: []string{
			"backups", "clusterId", "configVersion", "displayName",
			"extensions", "name", "nodes", "options", "roles",
		},
		requiredInputs: []string{"clusterId", "nodes"},
		properties: []string{
			"backups", "clusterId", "components", "configVersion", "createdAt",
			"displayName", "domain", "extensions", "name", "nodes", "options",
			"pgVersion", "roles", "status",
		},
		required: []string{
			"backups", "clusterId", "components", "configVersion", "createdAt",
			"domain", "extensions", "name", "nodes", "pgVersion", "roles",
			"status",
		},
		arrayProps: []string{"components", "options", "roles"},
	},
}

func TestResourceFieldShape(t *testing.T) {
	doc := loadSchema(t)

	for _, want := range resourceShapes {
		t.Run(want.token, func(t *testing.T) {
			r, ok := doc.Resources[want.token]
			if !ok {
				t.Fatalf("schema.json resources missing %q", want.token)
			}

			assertSameSet(t, "inputProperties", keys(r.InputProperties), want.inputProperties)
			assertSameSet(t, "requiredInputs", r.RequiredInputs, want.requiredInputs)
			assertSameSet(t, "properties", keys(r.Properties), want.properties)
			assertSameSet(t, "required", r.Required, want.required)

			for _, name := range want.arrayProps {
				prop, ok := r.Properties[name]
				if !ok {
					t.Errorf("properties missing %q", name)
					continue
				}
				if prop.Type != "array" {
					t.Errorf("properties[%q].type = %q, want array (bridge-metadata maxItemsOne override regressed?)", name, prop.Type)
				}
			}
		})
	}
}

// dataSourceShape locks in the top-level output surface of each data source.
// Each list endpoint should return one named array of objects plus the
// implicit `id` field that tfbridge adds.
type dataSourceShape struct {
	token       string
	listField   string // the named array field, e.g. "databases"
	allOutputs  []string
	allRequired []string
}

var dataSourceShapes = []dataSourceShape{
	{
		token:       "pgedge:index/getSSHKeys:getSSHKeys",
		listField:   "sshKeys",
		allOutputs:  []string{"id", "sshKeys"},
		allRequired: []string{"id", "sshKeys"},
	},
	{
		token:       "pgedge:index/getDatabases:getDatabases",
		listField:   "databases",
		allOutputs:  []string{"databases", "id"},
		allRequired: []string{"databases", "id"},
	},
	{
		token:       "pgedge:index/getClusters:getClusters",
		listField:   "clusters",
		allOutputs:  []string{"clusters", "id"},
		allRequired: []string{"clusters", "id"},
	},
	{
		token:       "pgedge:index/getCloudAccounts:getCloudAccounts",
		listField:   "cloudAccounts",
		allOutputs:  []string{"cloudAccounts", "id"},
		allRequired: []string{"cloudAccounts", "id"},
	},
	{
		token:       "pgedge:index/getBackupStores:getBackupStores",
		listField:   "backupStores",
		allOutputs:  []string{"backupStores", "id"},
		allRequired: []string{"backupStores", "id"},
	},
}

func TestDataSourceFieldShape(t *testing.T) {
	doc := loadSchema(t)

	for _, want := range dataSourceShapes {
		t.Run(want.token, func(t *testing.T) {
			fn, ok := doc.Functions[want.token]
			if !ok {
				t.Fatalf("schema.json functions missing %q", want.token)
			}

			assertSameSet(t, "outputs.properties", keys(fn.Outputs.Properties), want.allOutputs)
			assertSameSet(t, "outputs.required", fn.Outputs.Required, want.allRequired)

			prop, ok := fn.Outputs.Properties[want.listField]
			if !ok {
				t.Fatalf("outputs.properties missing %q", want.listField)
			}
			if prop.Type != "array" {
				t.Errorf("outputs.properties[%q].type = %q, want array", want.listField, prop.Type)
			}
		})
	}
}

func keys(m map[string]propertyDoc) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func assertSameSet(t *testing.T, label string, got, want []string) {
	t.Helper()
	g := slices.Clone(got)
	w := slices.Clone(want)
	slices.Sort(g)
	slices.Sort(w)
	if !slices.Equal(g, w) {
		t.Errorf("%s set mismatch\n  got:  %v\n  want: %v", label, g, w)
	}
}
