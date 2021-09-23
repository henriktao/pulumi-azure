// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Sql Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/synapse"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("BlobStorage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkspace, err := synapse.NewWorkspace(ctx, "exampleWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewSqlPool(ctx, "exampleSqlPool", &synapse.SqlPoolArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			SkuName:            pulumi.String("DW100c"),
// 			CreateMode:         pulumi.String("Default"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Synapse Sql Pool can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/sqlPool:SqlPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1
// ```
type SqlPool struct {
	pulumi.CustomResourceState

	// The name of the collation to use with this pool, only applicable when `createMode` is set to `Default`. Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// Specifies how to create the Sql Pool. Valid values are: `Default`, `Recovery` or `PointInTimeRestore`. Must be `Default` to create a new database. Defaults to `Default`.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// Is transparent data encryption enabled? Defaults to `false`.
	DataEncrypted pulumi.BoolPtrOutput `pulumi:"dataEncrypted"`
	// The name which should be used for this Synapse Sql Pool. Changing this forces a new synapse SqlPool to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Synapse Sql Pool or Sql Database which is to back up, only applicable when `createMode` is set to `Recovery`. Changing this forces a new Synapse Sql Pool to be created.
	RecoveryDatabaseId pulumi.StringPtrOutput `pulumi:"recoveryDatabaseId"`
	// A `restore` block as defined below. only applicable when `createMode` is set to `PointInTimeRestore`.
	Restore SqlPoolRestorePtrOutput `pulumi:"restore"`
	// Specifies the SKU Name for this Synapse Sql Pool. Possible values are `DW100c`, `DW200c`, `DW300c`, `DW400c`, `DW500c`, `DW1000c`, `DW1500c`, `DW2000c`, `DW2500c`, `DW3000c`, `DW5000c`, `DW6000c`, `DW7500c`, `DW10000c`, `DW15000c` or `DW30000c`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// The ID of Synapse Workspace within which this Sql Pool should be created. Changing this forces a new Synapse Sql Pool to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Sql Pool.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSqlPool registers a new resource with the given unique name, arguments, and options.
func NewSqlPool(ctx *pulumi.Context,
	name string, args *SqlPoolArgs, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	var resource SqlPool
	err := ctx.RegisterResource("azure:synapse/sqlPool:SqlPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlPool gets an existing SqlPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolState, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	var resource SqlPool
	err := ctx.ReadResource("azure:synapse/sqlPool:SqlPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlPool resources.
type sqlPoolState struct {
	// The name of the collation to use with this pool, only applicable when `createMode` is set to `Default`. Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the Sql Pool. Valid values are: `Default`, `Recovery` or `PointInTimeRestore`. Must be `Default` to create a new database. Defaults to `Default`.
	CreateMode *string `pulumi:"createMode"`
	// Is transparent data encryption enabled? Defaults to `false`.
	DataEncrypted *bool `pulumi:"dataEncrypted"`
	// The name which should be used for this Synapse Sql Pool. Changing this forces a new synapse SqlPool to be created.
	Name *string `pulumi:"name"`
	// The ID of the Synapse Sql Pool or Sql Database which is to back up, only applicable when `createMode` is set to `Recovery`. Changing this forces a new Synapse Sql Pool to be created.
	RecoveryDatabaseId *string `pulumi:"recoveryDatabaseId"`
	// A `restore` block as defined below. only applicable when `createMode` is set to `PointInTimeRestore`.
	Restore *SqlPoolRestore `pulumi:"restore"`
	// Specifies the SKU Name for this Synapse Sql Pool. Possible values are `DW100c`, `DW200c`, `DW300c`, `DW400c`, `DW500c`, `DW1000c`, `DW1500c`, `DW2000c`, `DW2500c`, `DW3000c`, `DW5000c`, `DW6000c`, `DW7500c`, `DW10000c`, `DW15000c` or `DW30000c`.
	SkuName *string `pulumi:"skuName"`
	// The ID of Synapse Workspace within which this Sql Pool should be created. Changing this forces a new Synapse Sql Pool to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Sql Pool.
	Tags map[string]string `pulumi:"tags"`
}

type SqlPoolState struct {
	// The name of the collation to use with this pool, only applicable when `createMode` is set to `Default`. Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the Sql Pool. Valid values are: `Default`, `Recovery` or `PointInTimeRestore`. Must be `Default` to create a new database. Defaults to `Default`.
	CreateMode pulumi.StringPtrInput
	// Is transparent data encryption enabled? Defaults to `false`.
	DataEncrypted pulumi.BoolPtrInput
	// The name which should be used for this Synapse Sql Pool. Changing this forces a new synapse SqlPool to be created.
	Name pulumi.StringPtrInput
	// The ID of the Synapse Sql Pool or Sql Database which is to back up, only applicable when `createMode` is set to `Recovery`. Changing this forces a new Synapse Sql Pool to be created.
	RecoveryDatabaseId pulumi.StringPtrInput
	// A `restore` block as defined below. only applicable when `createMode` is set to `PointInTimeRestore`.
	Restore SqlPoolRestorePtrInput
	// Specifies the SKU Name for this Synapse Sql Pool. Possible values are `DW100c`, `DW200c`, `DW300c`, `DW400c`, `DW500c`, `DW1000c`, `DW1500c`, `DW2000c`, `DW2500c`, `DW3000c`, `DW5000c`, `DW6000c`, `DW7500c`, `DW10000c`, `DW15000c` or `DW30000c`.
	SkuName pulumi.StringPtrInput
	// The ID of Synapse Workspace within which this Sql Pool should be created. Changing this forces a new Synapse Sql Pool to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Synapse Sql Pool.
	Tags pulumi.StringMapInput
}

func (SqlPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolState)(nil)).Elem()
}

type sqlPoolArgs struct {
	// The name of the collation to use with this pool, only applicable when `createMode` is set to `Default`. Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the Sql Pool. Valid values are: `Default`, `Recovery` or `PointInTimeRestore`. Must be `Default` to create a new database. Defaults to `Default`.
	CreateMode *string `pulumi:"createMode"`
	// Is transparent data encryption enabled? Defaults to `false`.
	DataEncrypted *bool `pulumi:"dataEncrypted"`
	// The name which should be used for this Synapse Sql Pool. Changing this forces a new synapse SqlPool to be created.
	Name *string `pulumi:"name"`
	// The ID of the Synapse Sql Pool or Sql Database which is to back up, only applicable when `createMode` is set to `Recovery`. Changing this forces a new Synapse Sql Pool to be created.
	RecoveryDatabaseId *string `pulumi:"recoveryDatabaseId"`
	// A `restore` block as defined below. only applicable when `createMode` is set to `PointInTimeRestore`.
	Restore *SqlPoolRestore `pulumi:"restore"`
	// Specifies the SKU Name for this Synapse Sql Pool. Possible values are `DW100c`, `DW200c`, `DW300c`, `DW400c`, `DW500c`, `DW1000c`, `DW1500c`, `DW2000c`, `DW2500c`, `DW3000c`, `DW5000c`, `DW6000c`, `DW7500c`, `DW10000c`, `DW15000c` or `DW30000c`.
	SkuName string `pulumi:"skuName"`
	// The ID of Synapse Workspace within which this Sql Pool should be created. Changing this forces a new Synapse Sql Pool to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Sql Pool.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqlPool resource.
type SqlPoolArgs struct {
	// The name of the collation to use with this pool, only applicable when `createMode` is set to `Default`. Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the Sql Pool. Valid values are: `Default`, `Recovery` or `PointInTimeRestore`. Must be `Default` to create a new database. Defaults to `Default`.
	CreateMode pulumi.StringPtrInput
	// Is transparent data encryption enabled? Defaults to `false`.
	DataEncrypted pulumi.BoolPtrInput
	// The name which should be used for this Synapse Sql Pool. Changing this forces a new synapse SqlPool to be created.
	Name pulumi.StringPtrInput
	// The ID of the Synapse Sql Pool or Sql Database which is to back up, only applicable when `createMode` is set to `Recovery`. Changing this forces a new Synapse Sql Pool to be created.
	RecoveryDatabaseId pulumi.StringPtrInput
	// A `restore` block as defined below. only applicable when `createMode` is set to `PointInTimeRestore`.
	Restore SqlPoolRestorePtrInput
	// Specifies the SKU Name for this Synapse Sql Pool. Possible values are `DW100c`, `DW200c`, `DW300c`, `DW400c`, `DW500c`, `DW1000c`, `DW1500c`, `DW2000c`, `DW2500c`, `DW3000c`, `DW5000c`, `DW6000c`, `DW7500c`, `DW10000c`, `DW15000c` or `DW30000c`.
	SkuName pulumi.StringInput
	// The ID of Synapse Workspace within which this Sql Pool should be created. Changing this forces a new Synapse Sql Pool to be created.
	SynapseWorkspaceId pulumi.StringInput
	// A mapping of tags which should be assigned to the Synapse Sql Pool.
	Tags pulumi.StringMapInput
}

func (SqlPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolArgs)(nil)).Elem()
}

type SqlPoolInput interface {
	pulumi.Input

	ToSqlPoolOutput() SqlPoolOutput
	ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput
}

func (*SqlPool) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPool)(nil))
}

func (i *SqlPool) ToSqlPoolOutput() SqlPoolOutput {
	return i.ToSqlPoolOutputWithContext(context.Background())
}

func (i *SqlPool) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolOutput)
}

func (i *SqlPool) ToSqlPoolPtrOutput() SqlPoolPtrOutput {
	return i.ToSqlPoolPtrOutputWithContext(context.Background())
}

func (i *SqlPool) ToSqlPoolPtrOutputWithContext(ctx context.Context) SqlPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolPtrOutput)
}

type SqlPoolPtrInput interface {
	pulumi.Input

	ToSqlPoolPtrOutput() SqlPoolPtrOutput
	ToSqlPoolPtrOutputWithContext(ctx context.Context) SqlPoolPtrOutput
}

type sqlPoolPtrType SqlPoolArgs

func (*sqlPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPool)(nil))
}

func (i *sqlPoolPtrType) ToSqlPoolPtrOutput() SqlPoolPtrOutput {
	return i.ToSqlPoolPtrOutputWithContext(context.Background())
}

func (i *sqlPoolPtrType) ToSqlPoolPtrOutputWithContext(ctx context.Context) SqlPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolPtrOutput)
}

// SqlPoolArrayInput is an input type that accepts SqlPoolArray and SqlPoolArrayOutput values.
// You can construct a concrete instance of `SqlPoolArrayInput` via:
//
//          SqlPoolArray{ SqlPoolArgs{...} }
type SqlPoolArrayInput interface {
	pulumi.Input

	ToSqlPoolArrayOutput() SqlPoolArrayOutput
	ToSqlPoolArrayOutputWithContext(context.Context) SqlPoolArrayOutput
}

type SqlPoolArray []SqlPoolInput

func (SqlPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlPool)(nil)).Elem()
}

func (i SqlPoolArray) ToSqlPoolArrayOutput() SqlPoolArrayOutput {
	return i.ToSqlPoolArrayOutputWithContext(context.Background())
}

func (i SqlPoolArray) ToSqlPoolArrayOutputWithContext(ctx context.Context) SqlPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolArrayOutput)
}

// SqlPoolMapInput is an input type that accepts SqlPoolMap and SqlPoolMapOutput values.
// You can construct a concrete instance of `SqlPoolMapInput` via:
//
//          SqlPoolMap{ "key": SqlPoolArgs{...} }
type SqlPoolMapInput interface {
	pulumi.Input

	ToSqlPoolMapOutput() SqlPoolMapOutput
	ToSqlPoolMapOutputWithContext(context.Context) SqlPoolMapOutput
}

type SqlPoolMap map[string]SqlPoolInput

func (SqlPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlPool)(nil)).Elem()
}

func (i SqlPoolMap) ToSqlPoolMapOutput() SqlPoolMapOutput {
	return i.ToSqlPoolMapOutputWithContext(context.Background())
}

func (i SqlPoolMap) ToSqlPoolMapOutputWithContext(ctx context.Context) SqlPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolMapOutput)
}

type SqlPoolOutput struct{ *pulumi.OutputState }

func (SqlPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPool)(nil))
}

func (o SqlPoolOutput) ToSqlPoolOutput() SqlPoolOutput {
	return o
}

func (o SqlPoolOutput) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return o
}

func (o SqlPoolOutput) ToSqlPoolPtrOutput() SqlPoolPtrOutput {
	return o.ToSqlPoolPtrOutputWithContext(context.Background())
}

func (o SqlPoolOutput) ToSqlPoolPtrOutputWithContext(ctx context.Context) SqlPoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlPool) *SqlPool {
		return &v
	}).(SqlPoolPtrOutput)
}

type SqlPoolPtrOutput struct{ *pulumi.OutputState }

func (SqlPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPool)(nil))
}

func (o SqlPoolPtrOutput) ToSqlPoolPtrOutput() SqlPoolPtrOutput {
	return o
}

func (o SqlPoolPtrOutput) ToSqlPoolPtrOutputWithContext(ctx context.Context) SqlPoolPtrOutput {
	return o
}

func (o SqlPoolPtrOutput) Elem() SqlPoolOutput {
	return o.ApplyT(func(v *SqlPool) SqlPool {
		if v != nil {
			return *v
		}
		var ret SqlPool
		return ret
	}).(SqlPoolOutput)
}

type SqlPoolArrayOutput struct{ *pulumi.OutputState }

func (SqlPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlPool)(nil))
}

func (o SqlPoolArrayOutput) ToSqlPoolArrayOutput() SqlPoolArrayOutput {
	return o
}

func (o SqlPoolArrayOutput) ToSqlPoolArrayOutputWithContext(ctx context.Context) SqlPoolArrayOutput {
	return o
}

func (o SqlPoolArrayOutput) Index(i pulumi.IntInput) SqlPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlPool {
		return vs[0].([]SqlPool)[vs[1].(int)]
	}).(SqlPoolOutput)
}

type SqlPoolMapOutput struct{ *pulumi.OutputState }

func (SqlPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SqlPool)(nil))
}

func (o SqlPoolMapOutput) ToSqlPoolMapOutput() SqlPoolMapOutput {
	return o
}

func (o SqlPoolMapOutput) ToSqlPoolMapOutputWithContext(ctx context.Context) SqlPoolMapOutput {
	return o
}

func (o SqlPoolMapOutput) MapIndex(k pulumi.StringInput) SqlPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SqlPool {
		return vs[0].(map[string]SqlPool)[vs[1].(string)]
	}).(SqlPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlPoolOutput{})
	pulumi.RegisterOutputType(SqlPoolPtrOutput{})
	pulumi.RegisterOutputType(SqlPoolArrayOutput{})
	pulumi.RegisterOutputType(SqlPoolMapOutput{})
}
