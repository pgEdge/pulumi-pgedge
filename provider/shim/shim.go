package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/pgEdge/pulumi-pgedge/provider/pkg/version"
	pgedge "github.com/pgEdge/terraform-provider-pgedge/internals/provider"
)

func NewProvider() tfpf.Provider {
	return pgedge.New(version.Version)()
}