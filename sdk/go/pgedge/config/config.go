// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Base Url to use when connecting to the PgEdge service.
func GetBaseUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "pgedge:baseUrl")
}
