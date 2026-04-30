package pgedge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/pgEdge/pulumi-pgedge/provider/pkg/version"
)

func init() {
	// version.Version is normally injected via -ldflags during build. In `go test`
	// it is empty, which makes Provider() panic inside tfbridge.GetModuleMajorVersion.
	if version.Version == "" {
		version.Version = "0.0.0-test"
	}
}

// expectedResources maps each Terraform resource name to the Pulumi type-name
// suffix it must produce. Renaming or dropping any of these is a breaking
// change for users of the generated SDKs.
var expectedResources = map[string]string{
	"pgedge_database":      "Database",
	"pgedge_cluster":       "Cluster",
	"pgedge_cloud_account": "CloudAccount",
	"pgedge_ssh_key":       "SSHKey",
	"pgedge_backup_store":  "BackupStore",
}

// expectedDataSources maps each Terraform data source name to the Pulumi
// function-name suffix it must produce.
var expectedDataSources = map[string]string{
	"pgedge_databases":      "getDatabases",
	"pgedge_clusters":       "getClusters",
	"pgedge_cloud_accounts": "getCloudAccounts",
	"pgedge_ssh_keys":       "getSSHKeys",
	"pgedge_backup_stores":  "getBackupStores",
}

func TestProviderResourceTokens(t *testing.T) {
	prov := Provider()

	if got, want := len(prov.Resources), len(expectedResources); got != want {
		t.Errorf("Provider().Resources has %d entries, want %d", got, want)
	}
	for tfName, typeName := range expectedResources {
		info, ok := prov.Resources[tfName]
		if !ok {
			t.Errorf("Provider().Resources missing %q", tfName)
			continue
		}
		tok := string(info.Tok)
		if !strings.HasSuffix(tok, ":"+typeName) {
			t.Errorf("Provider().Resources[%q].Tok = %q, want suffix :%s", tfName, tok, typeName)
		}
	}
}

func TestProviderDataSourceTokens(t *testing.T) {
	prov := Provider()

	if got, want := len(prov.DataSources), len(expectedDataSources); got != want {
		t.Errorf("Provider().DataSources has %d entries, want %d", got, want)
	}
	for tfName, fnName := range expectedDataSources {
		info, ok := prov.DataSources[tfName]
		if !ok {
			t.Errorf("Provider().DataSources missing %q", tfName)
			continue
		}
		tok := string(info.Tok)
		if !strings.HasSuffix(tok, ":"+fnName) {
			t.Errorf("Provider().DataSources[%q].Tok = %q, want suffix :%s", tfName, tok, fnName)
		}
	}
}

func TestProviderConfigBaseURL(t *testing.T) {
	prov := Provider()

	cfg, ok := prov.Config["base_url"]
	if !ok {
		t.Fatal(`Provider().Config["base_url"] missing — users rely on PGEDGE_BASE_URL to configure the API endpoint`)
	}
	if cfg.Default == nil {
		t.Fatal(`Provider().Config["base_url"].Default is nil`)
	}
	if !slices.Contains(cfg.Default.EnvVars, "PGEDGE_BASE_URL") {
		t.Errorf(`Provider().Config["base_url"].Default.EnvVars = %v, want PGEDGE_BASE_URL`, cfg.Default.EnvVars)
	}
}

func TestProviderMetadataLoads(t *testing.T) {
	prov := Provider()
	if prov.MetadataInfo == nil {
		t.Fatal("Provider().MetadataInfo is nil — embedded bridge-metadata.json failed to load")
	}
}

// TestSchemaJSONInSync asserts that the committed schema.json (the artifact
// that drives SDK codegen) reflects the current resources.go. Catches the
// "forgot to run `make tfgen`" class of mistake.
func TestSchemaJSONInSync(t *testing.T) {
	path := filepath.Join("cmd", "pulumi-resource-pgedge", "schema.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	var schema struct {
		Name      string                     `json:"name"`
		Resources map[string]json.RawMessage `json:"resources"`
		Functions map[string]json.RawMessage `json:"functions"`
		Config    struct {
			Variables map[string]json.RawMessage `json:"variables"`
		} `json:"config"`
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}

	if schema.Name != "pgedge" {
		t.Errorf("schema.json name = %q, want pgedge", schema.Name)
	}

	if got, want := len(schema.Resources), len(expectedResources); got != want {
		t.Errorf("schema.json has %d resources, want %d (run `make tfgen`?)", got, want)
	}
	for _, typeName := range expectedResources {
		if !hasTokenSuffix(schema.Resources, ":"+typeName) {
			t.Errorf("schema.json resources missing token ending in :%s (run `make tfgen`?)", typeName)
		}
	}

	if got, want := len(schema.Functions), len(expectedDataSources); got != want {
		t.Errorf("schema.json has %d functions, want %d (run `make tfgen`?)", got, want)
	}
	for _, fnName := range expectedDataSources {
		if !hasTokenSuffix(schema.Functions, ":"+fnName) {
			t.Errorf("schema.json functions missing token ending in :%s (run `make tfgen`?)", fnName)
		}
	}

	if _, ok := schema.Config.Variables["baseUrl"]; !ok {
		t.Errorf("schema.json config.variables missing baseUrl")
	}
}

func hasTokenSuffix(m map[string]json.RawMessage, suffix string) bool {
	for k := range m {
		if strings.HasSuffix(k, suffix) {
			return true
		}
	}
	return false
}
