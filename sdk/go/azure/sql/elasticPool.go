// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to manage an Azure SQL Elastic Pool.
//
// > **NOTE:** -  This version of the `Elasticpool` resource is being **deprecated** and should no longer be used. Please use the mssql.ElasticPool version instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sql"
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
// 		exampleSqlServer, err := sql.NewSqlServer(ctx, "exampleSqlServer", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("4dm1n157r470r"),
// 			AdministratorLoginPassword: pulumi.String("4-v3ry-53cr37-p455w0rd"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewElasticPool(ctx, "exampleElasticPool", &sql.ElasticPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ServerName:        exampleSqlServer.Name,
// 			Edition:           pulumi.String("Basic"),
// 			Dtu:               pulumi.Int(50),
// 			DbDtuMin:          pulumi.Int(0),
// 			DbDtuMax:          pulumi.Int(5),
// 			PoolSize:          pulumi.Int(5000),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **NOTE on `sql.ElasticPool`:** -  The values of `edition`, `dtu`, and `poolSize` must be consistent with the [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). Any inconsistent argument configuration will be rejected.
//
// ## Import
//
// SQL Elastic Pool's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/elasticPool:ElasticPool pool1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/pool1
// ```
type ElasticPool struct {
	pulumi.CustomResourceState

	// The creation date of the SQL Elastic Pool.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax pulumi.IntOutput `pulumi:"dbDtuMax"`
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin pulumi.IntOutput `pulumi:"dbDtuMin"`
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu pulumi.IntOutput `pulumi:"dtu"`
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition pulumi.StringOutput `pulumi:"edition"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize pulumi.IntOutput `pulumi:"poolSize"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dtu == nil {
		return nil, errors.New("invalid value for required argument 'Dtu'")
	}
	if args.Edition == nil {
		return nil, errors.New("invalid value for required argument 'Edition'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource ElasticPool
	err := ctx.RegisterResource("azure:sql/elasticPool:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure:sql/elasticPool:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticPool resources.
type elasticPoolState struct {
	// The creation date of the SQL Elastic Pool.
	CreationDate *string `pulumi:"creationDate"`
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax *int `pulumi:"dbDtuMax"`
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin *int `pulumi:"dbDtuMin"`
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu *int `pulumi:"dtu"`
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition *string `pulumi:"edition"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize *int `pulumi:"poolSize"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ElasticPoolState struct {
	// The creation date of the SQL Elastic Pool.
	CreationDate pulumi.StringPtrInput
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax pulumi.IntPtrInput
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin pulumi.IntPtrInput
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu pulumi.IntPtrInput
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize pulumi.IntPtrInput
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax *int `pulumi:"dbDtuMax"`
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin *int `pulumi:"dbDtuMin"`
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu int `pulumi:"dtu"`
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition string `pulumi:"edition"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize *int `pulumi:"poolSize"`
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax pulumi.IntPtrInput
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin pulumi.IntPtrInput
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu pulumi.IntInput
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize pulumi.IntPtrInput
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName pulumi.StringInput
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

func (i *ElasticPool) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return i.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPtrOutput)
}

type ElasticPoolPtrInput interface {
	pulumi.Input

	ToElasticPoolPtrOutput() ElasticPoolPtrOutput
	ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput
}

type elasticPoolPtrType ElasticPoolArgs

func (*elasticPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil))
}

func (i *elasticPoolPtrType) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return i.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (i *elasticPoolPtrType) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPtrOutput)
}

// ElasticPoolArrayInput is an input type that accepts ElasticPoolArray and ElasticPoolArrayOutput values.
// You can construct a concrete instance of `ElasticPoolArrayInput` via:
//
//          ElasticPoolArray{ ElasticPoolArgs{...} }
type ElasticPoolArrayInput interface {
	pulumi.Input

	ToElasticPoolArrayOutput() ElasticPoolArrayOutput
	ToElasticPoolArrayOutputWithContext(context.Context) ElasticPoolArrayOutput
}

type ElasticPoolArray []ElasticPoolInput

func (ElasticPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticPool)(nil)).Elem()
}

func (i ElasticPoolArray) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return i.ToElasticPoolArrayOutputWithContext(context.Background())
}

func (i ElasticPoolArray) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolArrayOutput)
}

// ElasticPoolMapInput is an input type that accepts ElasticPoolMap and ElasticPoolMapOutput values.
// You can construct a concrete instance of `ElasticPoolMapInput` via:
//
//          ElasticPoolMap{ "key": ElasticPoolArgs{...} }
type ElasticPoolMapInput interface {
	pulumi.Input

	ToElasticPoolMapOutput() ElasticPoolMapOutput
	ToElasticPoolMapOutputWithContext(context.Context) ElasticPoolMapOutput
}

type ElasticPoolMap map[string]ElasticPoolInput

func (ElasticPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticPool)(nil)).Elem()
}

func (i ElasticPoolMap) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return i.ToElasticPoolMapOutputWithContext(context.Background())
}

func (i ElasticPoolMap) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolMapOutput)
}

type ElasticPoolOutput struct{ *pulumi.OutputState }

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return o.ToElasticPoolPtrOutputWithContext(context.Background())
}

func (o ElasticPoolOutput) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ElasticPool) *ElasticPool {
		return &v
	}).(ElasticPoolPtrOutput)
}

type ElasticPoolPtrOutput struct{ *pulumi.OutputState }

func (ElasticPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil))
}

func (o ElasticPoolPtrOutput) ToElasticPoolPtrOutput() ElasticPoolPtrOutput {
	return o
}

func (o ElasticPoolPtrOutput) ToElasticPoolPtrOutputWithContext(ctx context.Context) ElasticPoolPtrOutput {
	return o
}

func (o ElasticPoolPtrOutput) Elem() ElasticPoolOutput {
	return o.ApplyT(func(v *ElasticPool) ElasticPool {
		if v != nil {
			return *v
		}
		var ret ElasticPool
		return ret
	}).(ElasticPoolOutput)
}

type ElasticPoolArrayOutput struct{ *pulumi.OutputState }

func (ElasticPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ElasticPool)(nil))
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutput() ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) ToElasticPoolArrayOutputWithContext(ctx context.Context) ElasticPoolArrayOutput {
	return o
}

func (o ElasticPoolArrayOutput) Index(i pulumi.IntInput) ElasticPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ElasticPool {
		return vs[0].([]ElasticPool)[vs[1].(int)]
	}).(ElasticPoolOutput)
}

type ElasticPoolMapOutput struct{ *pulumi.OutputState }

func (ElasticPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ElasticPool)(nil))
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutput() ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) ToElasticPoolMapOutputWithContext(ctx context.Context) ElasticPoolMapOutput {
	return o
}

func (o ElasticPoolMapOutput) MapIndex(k pulumi.StringInput) ElasticPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ElasticPool {
		return vs[0].(map[string]ElasticPool)[vs[1].(string)]
	}).(ElasticPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolInput)(nil)).Elem(), &ElasticPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolPtrInput)(nil)).Elem(), &ElasticPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolArrayInput)(nil)).Elem(), ElasticPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolMapInput)(nil)).Elem(), ElasticPoolMap{})
	pulumi.RegisterOutputType(ElasticPoolOutput{})
	pulumi.RegisterOutputType(ElasticPoolPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolMapOutput{})
}
