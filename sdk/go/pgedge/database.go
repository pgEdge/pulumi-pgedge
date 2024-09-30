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

// Manages a pgEdge database.
type Database struct {
	pulumi.CustomResourceState

	// Backup configuration for the database.
	Backups DatabaseBackupsOutput `pulumi:"backups"`
	// The ID of the cluster this database belongs to.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// List of components in the database.
	Components DatabaseComponentArrayOutput `pulumi:"components"`
	// The configuration version of the database.
	ConfigVersion pulumi.StringOutput `pulumi:"configVersion"`
	// The timestamp when the database was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The domain associated with the database.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Extensions configuration for the database.
	Extensions DatabaseExtensionsOutput `pulumi:"extensions"`
	// The name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of nodes in the database.
	Nodes DatabaseNodesMapOutput `pulumi:"nodes"`
	// A list of options for the database.
	Options pulumi.StringArrayOutput `pulumi:"options"`
	// The PostgreSQL version of the database.
	PgVersion pulumi.StringOutput `pulumi:"pgVersion"`
	// List of roles in the database.
	Roles DatabaseRoleArrayOutput `pulumi:"roles"`
	// The current status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// The amount of storage used by the database in bytes.
	StorageUsed pulumi.IntOutput `pulumi:"storageUsed"`
	// The timestamp when the database was last updated.
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
	if args.Nodes == nil {
		return nil, errors.New("invalid value for required argument 'Nodes'")
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
	// Backup configuration for the database.
	Backups *DatabaseBackups `pulumi:"backups"`
	// The ID of the cluster this database belongs to.
	ClusterId *string `pulumi:"clusterId"`
	// List of components in the database.
	Components []DatabaseComponent `pulumi:"components"`
	// The configuration version of the database.
	ConfigVersion *string `pulumi:"configVersion"`
	// The timestamp when the database was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The domain associated with the database.
	Domain *string `pulumi:"domain"`
	// Extensions configuration for the database.
	Extensions *DatabaseExtensions `pulumi:"extensions"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// List of nodes in the database.
	Nodes map[string]DatabaseNodes `pulumi:"nodes"`
	// A list of options for the database.
	Options []string `pulumi:"options"`
	// The PostgreSQL version of the database.
	PgVersion *string `pulumi:"pgVersion"`
	// List of roles in the database.
	Roles []DatabaseRole `pulumi:"roles"`
	// The current status of the database.
	Status *string `pulumi:"status"`
	// The amount of storage used by the database in bytes.
	StorageUsed *int `pulumi:"storageUsed"`
	// The timestamp when the database was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DatabaseState struct {
	// Backup configuration for the database.
	Backups DatabaseBackupsPtrInput
	// The ID of the cluster this database belongs to.
	ClusterId pulumi.StringPtrInput
	// List of components in the database.
	Components DatabaseComponentArrayInput
	// The configuration version of the database.
	ConfigVersion pulumi.StringPtrInput
	// The timestamp when the database was created.
	CreatedAt pulumi.StringPtrInput
	// The domain associated with the database.
	Domain pulumi.StringPtrInput
	// Extensions configuration for the database.
	Extensions DatabaseExtensionsPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// List of nodes in the database.
	Nodes DatabaseNodesMapInput
	// A list of options for the database.
	Options pulumi.StringArrayInput
	// The PostgreSQL version of the database.
	PgVersion pulumi.StringPtrInput
	// List of roles in the database.
	Roles DatabaseRoleArrayInput
	// The current status of the database.
	Status pulumi.StringPtrInput
	// The amount of storage used by the database in bytes.
	StorageUsed pulumi.IntPtrInput
	// The timestamp when the database was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Backup configuration for the database.
	Backups *DatabaseBackups `pulumi:"backups"`
	// The ID of the cluster this database belongs to.
	ClusterId string `pulumi:"clusterId"`
	// The configuration version of the database.
	ConfigVersion *string `pulumi:"configVersion"`
	// Extensions configuration for the database.
	Extensions *DatabaseExtensions `pulumi:"extensions"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// List of nodes in the database.
	Nodes map[string]DatabaseNodes `pulumi:"nodes"`
	// A list of options for the database.
	Options []string `pulumi:"options"`
	// List of roles in the database.
	Roles []DatabaseRole `pulumi:"roles"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Backup configuration for the database.
	Backups DatabaseBackupsPtrInput
	// The ID of the cluster this database belongs to.
	ClusterId pulumi.StringInput
	// The configuration version of the database.
	ConfigVersion pulumi.StringPtrInput
	// Extensions configuration for the database.
	Extensions DatabaseExtensionsPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// List of nodes in the database.
	Nodes DatabaseNodesMapInput
	// A list of options for the database.
	Options pulumi.StringArrayInput
	// List of roles in the database.
	Roles DatabaseRoleArrayInput
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

// Backup configuration for the database.
func (o DatabaseOutput) Backups() DatabaseBackupsOutput {
	return o.ApplyT(func(v *Database) DatabaseBackupsOutput { return v.Backups }).(DatabaseBackupsOutput)
}

// The ID of the cluster this database belongs to.
func (o DatabaseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// List of components in the database.
func (o DatabaseOutput) Components() DatabaseComponentArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseComponentArrayOutput { return v.Components }).(DatabaseComponentArrayOutput)
}

// The configuration version of the database.
func (o DatabaseOutput) ConfigVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ConfigVersion }).(pulumi.StringOutput)
}

// The timestamp when the database was created.
func (o DatabaseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The domain associated with the database.
func (o DatabaseOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Extensions configuration for the database.
func (o DatabaseOutput) Extensions() DatabaseExtensionsOutput {
	return o.ApplyT(func(v *Database) DatabaseExtensionsOutput { return v.Extensions }).(DatabaseExtensionsOutput)
}

// The name of the database.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of nodes in the database.
func (o DatabaseOutput) Nodes() DatabaseNodesMapOutput {
	return o.ApplyT(func(v *Database) DatabaseNodesMapOutput { return v.Nodes }).(DatabaseNodesMapOutput)
}

// A list of options for the database.
func (o DatabaseOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Database) pulumi.StringArrayOutput { return v.Options }).(pulumi.StringArrayOutput)
}

// The PostgreSQL version of the database.
func (o DatabaseOutput) PgVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.PgVersion }).(pulumi.StringOutput)
}

// List of roles in the database.
func (o DatabaseOutput) Roles() DatabaseRoleArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseRoleArrayOutput { return v.Roles }).(DatabaseRoleArrayOutput)
}

// The current status of the database.
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The amount of storage used by the database in bytes.
func (o DatabaseOutput) StorageUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.StorageUsed }).(pulumi.IntOutput)
}

// The timestamp when the database was last updated.
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
