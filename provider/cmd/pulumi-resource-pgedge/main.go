//go:generate go run ./generate.go
package main

import (
	"context"

	_ "embed"

	pgedge "github.com/pgEdge/pulumi-pgedge/provider"
	tfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
)

//go:embed schema.json
var pulumiSchema []byte

//go:embed bridge-metadata.json
var bridgeMetadata []byte

func main() {
	meta := tfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}
	tfbridge.Main(context.Background(), "pgedge", pgedge.Provider(), meta)
}
