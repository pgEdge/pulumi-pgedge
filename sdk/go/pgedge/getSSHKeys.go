// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pgedge

import (
	"context"
	"reflect"

	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSSHKeys(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetSSHKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSSHKeysResult
	err := ctx.Invoke("pgedge:index/getSSHKeys:getSSHKeys", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getSSHKeys.
type GetSSHKeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string             `pulumi:"id"`
	SshKeys []GetSSHKeysSshKey `pulumi:"sshKeys"`
}

func GetSSHKeysOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetSSHKeysResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetSSHKeysResult, error) {
		r, err := GetSSHKeys(ctx, opts...)
		var s GetSSHKeysResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetSSHKeysResultOutput)
}

// A collection of values returned by getSSHKeys.
type GetSSHKeysResultOutput struct{ *pulumi.OutputState }

func (GetSSHKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSSHKeysResult)(nil)).Elem()
}

func (o GetSSHKeysResultOutput) ToGetSSHKeysResultOutput() GetSSHKeysResultOutput {
	return o
}

func (o GetSSHKeysResultOutput) ToGetSSHKeysResultOutputWithContext(ctx context.Context) GetSSHKeysResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSSHKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSSHKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSSHKeysResultOutput) SshKeys() GetSSHKeysSshKeyArrayOutput {
	return o.ApplyT(func(v GetSSHKeysResult) []GetSSHKeysSshKey { return v.SshKeys }).(GetSSHKeysSshKeyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSSHKeysResultOutput{})
}