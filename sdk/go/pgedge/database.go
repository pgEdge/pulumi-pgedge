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

// Interface with the pgEdge service API.
type Database struct {
	pulumi.CustomResourceState

	// Cluster Id of the database
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Created at of the database
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Domain of the database
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Name of the location
	Name  pulumi.StringOutput     `pulumi:"name"`
	Nodes DatabaseNodeArrayOutput `pulumi:"nodes"`
	// Options for creating the database
	Options pulumi.StringArrayOutput `pulumi:"options"`
	// Status of the database
	Status pulumi.StringOutput `pulumi:"status"`
	// Updated at of the database
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("pgedge:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("pgedge:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Cluster Id of the database
	ClusterId *string `pulumi:"clusterId"`
	// Created at of the database
	CreatedAt *string `pulumi:"createdAt"`
	// Domain of the database
	Domain *string `pulumi:"domain"`
	// Name of the location
	Name  *string        `pulumi:"name"`
	Nodes []DatabaseNode `pulumi:"nodes"`
	// Options for creating the database
	Options []string `pulumi:"options"`
	// Status of the database
	Status *string `pulumi:"status"`
	// Updated at of the database
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DatabaseState struct {
	// Cluster Id of the database
	ClusterId pulumi.StringPtrInput
	// Created at of the database
	CreatedAt pulumi.StringPtrInput
	// Domain of the database
	Domain pulumi.StringPtrInput
	// Name of the location
	Name  pulumi.StringPtrInput
	Nodes DatabaseNodeArrayInput
	// Options for creating the database
	Options pulumi.StringArrayInput
	// Status of the database
	Status pulumi.StringPtrInput
	// Updated at of the database
	UpdatedAt pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Cluster Id of the database
	ClusterId string `pulumi:"clusterId"`
	// Name of the location
	Name *string `pulumi:"name"`
	// Options for creating the database
	Options []string `pulumi:"options"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Cluster Id of the database
	ClusterId pulumi.StringInput
	// Name of the location
	Name pulumi.StringPtrInput
	// Options for creating the database
	Options pulumi.StringArrayInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// Cluster Id of the database
func (o DatabaseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Created at of the database
func (o DatabaseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Domain of the database
func (o DatabaseOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Name of the location
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Nodes() DatabaseNodeArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseNodeArrayOutput { return v.Nodes }).(DatabaseNodeArrayOutput)
}

// Options for creating the database
func (o DatabaseOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Database) pulumi.StringArrayOutput { return v.Options }).(pulumi.StringArrayOutput)
}

// Status of the database
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Updated at of the database
func (o DatabaseOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
