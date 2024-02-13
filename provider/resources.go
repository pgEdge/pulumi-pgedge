package pgedge

import (
	"fmt"
	"path"

	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	// shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	// "github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pgEdge/pulumi-pgedge/provider/pkg/version"
	shim "github.com/pgEdge/pulumi-pgedge/provider/shim"
)

// all of the token components used below.
const (
	mainPkg = "pgedge"
	mainMod = "index"
)

// func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
// 	return nil
// }

//go:embed cmd/pulumi-resource-pgedge/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(shim.NewProvider()),
		Name:              "pgedge",
		DisplayName:       "pgEdge",
		Publisher:         "pgEdge",
		LogoURL:           "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing pgedge cloud resources.",
		Keywords:          []string{"pulumi", "pgedge", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pgedge.com",
		Repository:        "https://github.com/pgEdge/pulumi-pgedge",
		Version:           version.Version,
		GitHubOrg:         "pgEdge",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		// PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"pgedge_database": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Database"),
				// Fields: map[string]*tfbridge.SchemaInfo{
				// 	"database_name": {
				// 		Name:           "",
				// 		Secret:         tfbridge.True(),
				// 		CSharpName:     "",
				// 		MarkAsOptional: false,
				// 	},
				// },
				// Docs: &tfbridge.DocInfo{},
			},
			"pgedge_cluster": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cluster"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"pgedge_databases": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDatabases"),
			},
			"pgedge_clusters": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getClusters"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pgEdge/pulumi-pgedge",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pgEdge_pulumi_pgedge",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pgEdge/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pgedge",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("pgedge_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
