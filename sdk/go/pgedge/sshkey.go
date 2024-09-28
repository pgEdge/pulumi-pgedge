// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pgedge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pgEdge/pulumi-pgedge/sdk/go/pgedge/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SSHKey struct {
	pulumi.CustomResourceState

	// The timestamp when the SSH key was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name of the SSH key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The public key.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewSSHKey registers a new resource with the given unique name, arguments, and options.
func NewSSHKey(ctx *pulumi.Context,
	name string, args *SSHKeyArgs, opts ...pulumi.ResourceOption) (*SSHKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SSHKey
	err := ctx.RegisterResource("pgedge:index/sSHKey:SSHKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSHKey gets an existing SSHKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSHKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSHKeyState, opts ...pulumi.ResourceOption) (*SSHKey, error) {
	var resource SSHKey
	err := ctx.ReadResource("pgedge:index/sSHKey:SSHKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSHKey resources.
type sshkeyState struct {
	// The timestamp when the SSH key was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The name of the SSH key.
	Name *string `pulumi:"name"`
	// The public key.
	PublicKey *string `pulumi:"publicKey"`
}

type SSHKeyState struct {
	// The timestamp when the SSH key was created.
	CreatedAt pulumi.StringPtrInput
	// The name of the SSH key.
	Name pulumi.StringPtrInput
	// The public key.
	PublicKey pulumi.StringPtrInput
}

func (SSHKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyState)(nil)).Elem()
}

type sshkeyArgs struct {
	// The name of the SSH key.
	Name *string `pulumi:"name"`
	// The public key.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SSHKey resource.
type SSHKeyArgs struct {
	// The name of the SSH key.
	Name pulumi.StringPtrInput
	// The public key.
	PublicKey pulumi.StringInput
}

func (SSHKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyArgs)(nil)).Elem()
}

type SSHKeyInput interface {
	pulumi.Input

	ToSSHKeyOutput() SSHKeyOutput
	ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput
}

func (*SSHKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKey)(nil)).Elem()
}

func (i *SSHKey) ToSSHKeyOutput() SSHKeyOutput {
	return i.ToSSHKeyOutputWithContext(context.Background())
}

func (i *SSHKey) ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyOutput)
}

// SSHKeyArrayInput is an input type that accepts SSHKeyArray and SSHKeyArrayOutput values.
// You can construct a concrete instance of `SSHKeyArrayInput` via:
//
//	SSHKeyArray{ SSHKeyArgs{...} }
type SSHKeyArrayInput interface {
	pulumi.Input

	ToSSHKeyArrayOutput() SSHKeyArrayOutput
	ToSSHKeyArrayOutputWithContext(context.Context) SSHKeyArrayOutput
}

type SSHKeyArray []SSHKeyInput

func (SSHKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKey)(nil)).Elem()
}

func (i SSHKeyArray) ToSSHKeyArrayOutput() SSHKeyArrayOutput {
	return i.ToSSHKeyArrayOutputWithContext(context.Background())
}

func (i SSHKeyArray) ToSSHKeyArrayOutputWithContext(ctx context.Context) SSHKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyArrayOutput)
}

// SSHKeyMapInput is an input type that accepts SSHKeyMap and SSHKeyMapOutput values.
// You can construct a concrete instance of `SSHKeyMapInput` via:
//
//	SSHKeyMap{ "key": SSHKeyArgs{...} }
type SSHKeyMapInput interface {
	pulumi.Input

	ToSSHKeyMapOutput() SSHKeyMapOutput
	ToSSHKeyMapOutputWithContext(context.Context) SSHKeyMapOutput
}

type SSHKeyMap map[string]SSHKeyInput

func (SSHKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKey)(nil)).Elem()
}

func (i SSHKeyMap) ToSSHKeyMapOutput() SSHKeyMapOutput {
	return i.ToSSHKeyMapOutputWithContext(context.Background())
}

func (i SSHKeyMap) ToSSHKeyMapOutputWithContext(ctx context.Context) SSHKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyMapOutput)
}

type SSHKeyOutput struct{ *pulumi.OutputState }

func (SSHKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKey)(nil)).Elem()
}

func (o SSHKeyOutput) ToSSHKeyOutput() SSHKeyOutput {
	return o
}

func (o SSHKeyOutput) ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput {
	return o
}

// The timestamp when the SSH key was created.
func (o SSHKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The name of the SSH key.
func (o SSHKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The public key.
func (o SSHKeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

type SSHKeyArrayOutput struct{ *pulumi.OutputState }

func (SSHKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKey)(nil)).Elem()
}

func (o SSHKeyArrayOutput) ToSSHKeyArrayOutput() SSHKeyArrayOutput {
	return o
}

func (o SSHKeyArrayOutput) ToSSHKeyArrayOutputWithContext(ctx context.Context) SSHKeyArrayOutput {
	return o
}

func (o SSHKeyArrayOutput) Index(i pulumi.IntInput) SSHKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SSHKey {
		return vs[0].([]*SSHKey)[vs[1].(int)]
	}).(SSHKeyOutput)
}

type SSHKeyMapOutput struct{ *pulumi.OutputState }

func (SSHKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKey)(nil)).Elem()
}

func (o SSHKeyMapOutput) ToSSHKeyMapOutput() SSHKeyMapOutput {
	return o
}

func (o SSHKeyMapOutput) ToSSHKeyMapOutputWithContext(ctx context.Context) SSHKeyMapOutput {
	return o
}

func (o SSHKeyMapOutput) MapIndex(k pulumi.StringInput) SSHKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SSHKey {
		return vs[0].(map[string]*SSHKey)[vs[1].(string)]
	}).(SSHKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyInput)(nil)).Elem(), &SSHKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyArrayInput)(nil)).Elem(), SSHKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyMapInput)(nil)).Elem(), SSHKeyMap{})
	pulumi.RegisterOutputType(SSHKeyOutput{})
	pulumi.RegisterOutputType(SSHKeyArrayOutput{})
	pulumi.RegisterOutputType(SSHKeyMapOutput{})
}
