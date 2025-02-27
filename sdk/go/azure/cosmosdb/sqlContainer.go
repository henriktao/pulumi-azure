// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SQL Container within a Cosmos DB Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cosmosdb.NewSqlContainer(ctx, "example", &cosmosdb.SqlContainerArgs{
// 			ResourceGroupName:   pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:         pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:        pulumi.Any(azurerm_cosmosdb_sql_database.Example.Name),
// 			PartitionKeyPath:    pulumi.String("/definition/id"),
// 			PartitionKeyVersion: pulumi.Int(1),
// 			Throughput:          pulumi.Int(400),
// 			IndexingPolicy: &cosmosdb.SqlContainerIndexingPolicyArgs{
// 				IndexingMode: pulumi.String("Consistent"),
// 				IncludedPaths: cosmosdb.SqlContainerIndexingPolicyIncludedPathArray{
// 					&cosmosdb.SqlContainerIndexingPolicyIncludedPathArgs{
// 						Path: pulumi.String("/*"),
// 					},
// 					&cosmosdb.SqlContainerIndexingPolicyIncludedPathArgs{
// 						Path: pulumi.String("/included/?"),
// 					},
// 				},
// 				ExcludedPaths: cosmosdb.SqlContainerIndexingPolicyExcludedPathArray{
// 					&cosmosdb.SqlContainerIndexingPolicyExcludedPathArgs{
// 						Path: pulumi.String("/excluded/?"),
// 					},
// 				},
// 			},
// 			UniqueKeys: cosmosdb.SqlContainerUniqueKeyArray{
// 				&cosmosdb.SqlContainerUniqueKeyArgs{
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/definition/idlong"),
// 						pulumi.String("/definition/idshort"),
// 					},
// 				},
// 			},
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
// Cosmos SQL Containers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/sqlContainer:SqlContainer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1
// ```
type SqlContainer struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrOutput `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicy SqlContainerConflictResolutionPolicyOutput `pulumi:"conflictResolutionPolicy"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntOutput `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyOutput `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringOutput `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrOutput `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayOutput `pulumi:"uniqueKeys"`
}

// NewSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewSqlContainer(ctx *pulumi.Context,
	name string, args *SqlContainerArgs, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.PartitionKeyPath == nil {
		return nil, errors.New("invalid value for required argument 'PartitionKeyPath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource SqlContainer
	err := ctx.RegisterResource("azure:cosmosdb/sqlContainer:SqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlContainer gets an existing SqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlContainerState, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	var resource SqlContainer
	err := ctx.ReadResource("azure:cosmosdb/sqlContainer:SqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlContainer resources.
type sqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl *int `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings *SqlContainerAutoscaleSettings `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicy *SqlContainerConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy *SqlContainerIndexingPolicy `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion *int `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

type SqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrInput
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicy SqlContainerConflictResolutionPolicyPtrInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyPtrInput
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrInput
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerState)(nil)).Elem()
}

type sqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl *int `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings *SqlContainerAutoscaleSettings `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicy *SqlContainerConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// An `indexingPolicy` block as defined below.
	IndexingPolicy *SqlContainerIndexingPolicy `pulumi:"indexingPolicy"`
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath string `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion *int `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a SqlContainer resource.
type SqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// The default time to live of Analytical Storage for this SQL container. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply. Requires `partitionKeyPath` to be set.
	AutoscaleSettings SqlContainerAutoscaleSettingsPtrInput
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicy SqlContainerConflictResolutionPolicyPtrInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// An `indexingPolicy` block as defined below.
	IndexingPolicy SqlContainerIndexingPolicyPtrInput
	// Specifies the name of the Cosmos DB SQL Container. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringInput
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrInput
	// The name of the resource group in which the Cosmos DB SQL Container is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerArgs)(nil)).Elem()
}

type SqlContainerInput interface {
	pulumi.Input

	ToSqlContainerOutput() SqlContainerOutput
	ToSqlContainerOutputWithContext(ctx context.Context) SqlContainerOutput
}

func (*SqlContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainer)(nil)).Elem()
}

func (i *SqlContainer) ToSqlContainerOutput() SqlContainerOutput {
	return i.ToSqlContainerOutputWithContext(context.Background())
}

func (i *SqlContainer) ToSqlContainerOutputWithContext(ctx context.Context) SqlContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerOutput)
}

// SqlContainerArrayInput is an input type that accepts SqlContainerArray and SqlContainerArrayOutput values.
// You can construct a concrete instance of `SqlContainerArrayInput` via:
//
//          SqlContainerArray{ SqlContainerArgs{...} }
type SqlContainerArrayInput interface {
	pulumi.Input

	ToSqlContainerArrayOutput() SqlContainerArrayOutput
	ToSqlContainerArrayOutputWithContext(context.Context) SqlContainerArrayOutput
}

type SqlContainerArray []SqlContainerInput

func (SqlContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlContainer)(nil)).Elem()
}

func (i SqlContainerArray) ToSqlContainerArrayOutput() SqlContainerArrayOutput {
	return i.ToSqlContainerArrayOutputWithContext(context.Background())
}

func (i SqlContainerArray) ToSqlContainerArrayOutputWithContext(ctx context.Context) SqlContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerArrayOutput)
}

// SqlContainerMapInput is an input type that accepts SqlContainerMap and SqlContainerMapOutput values.
// You can construct a concrete instance of `SqlContainerMapInput` via:
//
//          SqlContainerMap{ "key": SqlContainerArgs{...} }
type SqlContainerMapInput interface {
	pulumi.Input

	ToSqlContainerMapOutput() SqlContainerMapOutput
	ToSqlContainerMapOutputWithContext(context.Context) SqlContainerMapOutput
}

type SqlContainerMap map[string]SqlContainerInput

func (SqlContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlContainer)(nil)).Elem()
}

func (i SqlContainerMap) ToSqlContainerMapOutput() SqlContainerMapOutput {
	return i.ToSqlContainerMapOutputWithContext(context.Background())
}

func (i SqlContainerMap) ToSqlContainerMapOutputWithContext(ctx context.Context) SqlContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerMapOutput)
}

type SqlContainerOutput struct{ *pulumi.OutputState }

func (SqlContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainer)(nil)).Elem()
}

func (o SqlContainerOutput) ToSqlContainerOutput() SqlContainerOutput {
	return o
}

func (o SqlContainerOutput) ToSqlContainerOutputWithContext(ctx context.Context) SqlContainerOutput {
	return o
}

type SqlContainerArrayOutput struct{ *pulumi.OutputState }

func (SqlContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlContainer)(nil)).Elem()
}

func (o SqlContainerArrayOutput) ToSqlContainerArrayOutput() SqlContainerArrayOutput {
	return o
}

func (o SqlContainerArrayOutput) ToSqlContainerArrayOutputWithContext(ctx context.Context) SqlContainerArrayOutput {
	return o
}

func (o SqlContainerArrayOutput) Index(i pulumi.IntInput) SqlContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlContainer {
		return vs[0].([]*SqlContainer)[vs[1].(int)]
	}).(SqlContainerOutput)
}

type SqlContainerMapOutput struct{ *pulumi.OutputState }

func (SqlContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlContainer)(nil)).Elem()
}

func (o SqlContainerMapOutput) ToSqlContainerMapOutput() SqlContainerMapOutput {
	return o
}

func (o SqlContainerMapOutput) ToSqlContainerMapOutputWithContext(ctx context.Context) SqlContainerMapOutput {
	return o
}

func (o SqlContainerMapOutput) MapIndex(k pulumi.StringInput) SqlContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlContainer {
		return vs[0].(map[string]*SqlContainer)[vs[1].(string)]
	}).(SqlContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlContainerInput)(nil)).Elem(), &SqlContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlContainerArrayInput)(nil)).Elem(), SqlContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlContainerMapInput)(nil)).Elem(), SqlContainerMap{})
	pulumi.RegisterOutputType(SqlContainerOutput{})
	pulumi.RegisterOutputType(SqlContainerArrayOutput{})
	pulumi.RegisterOutputType(SqlContainerMapOutput{})
}
