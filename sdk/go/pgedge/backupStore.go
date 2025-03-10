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

type BackupStore struct {
	pulumi.CustomResourceState

	CloudAccountId   pulumi.StringOutput      `pulumi:"cloudAccountId"`
	CloudAccountType pulumi.StringOutput      `pulumi:"cloudAccountType"`
	ClusterIds       pulumi.StringArrayOutput `pulumi:"clusterIds"`
	CreatedAt        pulumi.StringOutput      `pulumi:"createdAt"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Properties       pulumi.StringMapOutput   `pulumi:"properties"`
	Region           pulumi.StringOutput      `pulumi:"region"`
	Status           pulumi.StringOutput      `pulumi:"status"`
}

// NewBackupStore registers a new resource with the given unique name, arguments, and options.
func NewBackupStore(ctx *pulumi.Context,
	name string, args *BackupStoreArgs, opts ...pulumi.ResourceOption) (*BackupStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudAccountId == nil {
		return nil, errors.New("invalid value for required argument 'CloudAccountId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupStore
	err := ctx.RegisterResource("pgedge:index/backupStore:BackupStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupStore gets an existing BackupStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupStoreState, opts ...pulumi.ResourceOption) (*BackupStore, error) {
	var resource BackupStore
	err := ctx.ReadResource("pgedge:index/backupStore:BackupStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupStore resources.
type backupStoreState struct {
	CloudAccountId   *string           `pulumi:"cloudAccountId"`
	CloudAccountType *string           `pulumi:"cloudAccountType"`
	ClusterIds       []string          `pulumi:"clusterIds"`
	CreatedAt        *string           `pulumi:"createdAt"`
	Name             *string           `pulumi:"name"`
	Properties       map[string]string `pulumi:"properties"`
	Region           *string           `pulumi:"region"`
	Status           *string           `pulumi:"status"`
}

type BackupStoreState struct {
	CloudAccountId   pulumi.StringPtrInput
	CloudAccountType pulumi.StringPtrInput
	ClusterIds       pulumi.StringArrayInput
	CreatedAt        pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Properties       pulumi.StringMapInput
	Region           pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
}

func (BackupStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupStoreState)(nil)).Elem()
}

type backupStoreArgs struct {
	CloudAccountId string  `pulumi:"cloudAccountId"`
	Name           *string `pulumi:"name"`
	Region         string  `pulumi:"region"`
}

// The set of arguments for constructing a BackupStore resource.
type BackupStoreArgs struct {
	CloudAccountId pulumi.StringInput
	Name           pulumi.StringPtrInput
	Region         pulumi.StringInput
}

func (BackupStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupStoreArgs)(nil)).Elem()
}

type BackupStoreInput interface {
	pulumi.Input

	ToBackupStoreOutput() BackupStoreOutput
	ToBackupStoreOutputWithContext(ctx context.Context) BackupStoreOutput
}

func (*BackupStore) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupStore)(nil)).Elem()
}

func (i *BackupStore) ToBackupStoreOutput() BackupStoreOutput {
	return i.ToBackupStoreOutputWithContext(context.Background())
}

func (i *BackupStore) ToBackupStoreOutputWithContext(ctx context.Context) BackupStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStoreOutput)
}

// BackupStoreArrayInput is an input type that accepts BackupStoreArray and BackupStoreArrayOutput values.
// You can construct a concrete instance of `BackupStoreArrayInput` via:
//
//	BackupStoreArray{ BackupStoreArgs{...} }
type BackupStoreArrayInput interface {
	pulumi.Input

	ToBackupStoreArrayOutput() BackupStoreArrayOutput
	ToBackupStoreArrayOutputWithContext(context.Context) BackupStoreArrayOutput
}

type BackupStoreArray []BackupStoreInput

func (BackupStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupStore)(nil)).Elem()
}

func (i BackupStoreArray) ToBackupStoreArrayOutput() BackupStoreArrayOutput {
	return i.ToBackupStoreArrayOutputWithContext(context.Background())
}

func (i BackupStoreArray) ToBackupStoreArrayOutputWithContext(ctx context.Context) BackupStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStoreArrayOutput)
}

// BackupStoreMapInput is an input type that accepts BackupStoreMap and BackupStoreMapOutput values.
// You can construct a concrete instance of `BackupStoreMapInput` via:
//
//	BackupStoreMap{ "key": BackupStoreArgs{...} }
type BackupStoreMapInput interface {
	pulumi.Input

	ToBackupStoreMapOutput() BackupStoreMapOutput
	ToBackupStoreMapOutputWithContext(context.Context) BackupStoreMapOutput
}

type BackupStoreMap map[string]BackupStoreInput

func (BackupStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupStore)(nil)).Elem()
}

func (i BackupStoreMap) ToBackupStoreMapOutput() BackupStoreMapOutput {
	return i.ToBackupStoreMapOutputWithContext(context.Background())
}

func (i BackupStoreMap) ToBackupStoreMapOutputWithContext(ctx context.Context) BackupStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupStoreMapOutput)
}

type BackupStoreOutput struct{ *pulumi.OutputState }

func (BackupStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupStore)(nil)).Elem()
}

func (o BackupStoreOutput) ToBackupStoreOutput() BackupStoreOutput {
	return o
}

func (o BackupStoreOutput) ToBackupStoreOutputWithContext(ctx context.Context) BackupStoreOutput {
	return o
}

func (o BackupStoreOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.CloudAccountId }).(pulumi.StringOutput)
}

func (o BackupStoreOutput) CloudAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.CloudAccountType }).(pulumi.StringOutput)
}

func (o BackupStoreOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringArrayOutput { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o BackupStoreOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o BackupStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupStoreOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BackupStoreOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o BackupStoreOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupStore) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BackupStoreArrayOutput struct{ *pulumi.OutputState }

func (BackupStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupStore)(nil)).Elem()
}

func (o BackupStoreArrayOutput) ToBackupStoreArrayOutput() BackupStoreArrayOutput {
	return o
}

func (o BackupStoreArrayOutput) ToBackupStoreArrayOutputWithContext(ctx context.Context) BackupStoreArrayOutput {
	return o
}

func (o BackupStoreArrayOutput) Index(i pulumi.IntInput) BackupStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupStore {
		return vs[0].([]*BackupStore)[vs[1].(int)]
	}).(BackupStoreOutput)
}

type BackupStoreMapOutput struct{ *pulumi.OutputState }

func (BackupStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupStore)(nil)).Elem()
}

func (o BackupStoreMapOutput) ToBackupStoreMapOutput() BackupStoreMapOutput {
	return o
}

func (o BackupStoreMapOutput) ToBackupStoreMapOutputWithContext(ctx context.Context) BackupStoreMapOutput {
	return o
}

func (o BackupStoreMapOutput) MapIndex(k pulumi.StringInput) BackupStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupStore {
		return vs[0].(map[string]*BackupStore)[vs[1].(string)]
	}).(BackupStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStoreInput)(nil)).Elem(), &BackupStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStoreArrayInput)(nil)).Elem(), BackupStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupStoreMapInput)(nil)).Elem(), BackupStoreMap{})
	pulumi.RegisterOutputType(BackupStoreOutput{})
	pulumi.RegisterOutputType(BackupStoreArrayOutput{})
	pulumi.RegisterOutputType(BackupStoreMapOutput{})
}
