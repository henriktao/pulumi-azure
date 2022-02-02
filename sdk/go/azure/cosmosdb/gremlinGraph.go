// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Gremlin Graph within a Cosmos DB Account.
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
// 		exampleAccount, err := cosmosdb.LookupAccount(ctx, &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleGremlinDatabase, err := cosmosdb.NewGremlinDatabase(ctx, "exampleGremlinDatabase", &cosmosdb.GremlinDatabaseArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewGremlinGraph(ctx, "exampleGremlinGraph", &cosmosdb.GremlinGraphArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_cosmosdb_account.Example.Resource_group_name),
// 			AccountName:       pulumi.Any(azurerm_cosmosdb_account.Example.Name),
// 			DatabaseName:      exampleGremlinDatabase.Name,
// 			PartitionKeyPath:  pulumi.String("/Example"),
// 			Throughput:        pulumi.Int(400),
// 			IndexPolicies: cosmosdb.GremlinGraphIndexPolicyArray{
// 				&cosmosdb.GremlinGraphIndexPolicyArgs{
// 					Automatic:    pulumi.Bool(true),
// 					IndexingMode: pulumi.String("Consistent"),
// 					IncludedPaths: pulumi.StringArray{
// 						pulumi.String("/*"),
// 					},
// 					ExcludedPaths: pulumi.StringArray{
// 						pulumi.String("/\"_etag\"/?"),
// 					},
// 				},
// 			},
// 			ConflictResolutionPolicies: cosmosdb.GremlinGraphConflictResolutionPolicyArray{
// 				&cosmosdb.GremlinGraphConflictResolutionPolicyArgs{
// 					Mode:                   pulumi.String("LastWriterWins"),
// 					ConflictResolutionPath: pulumi.String("/_ts"),
// 				},
// 			},
// 			UniqueKeys: cosmosdb.GremlinGraphUniqueKeyArray{
// 				&cosmosdb.GremlinGraphUniqueKeyArgs{
// 					Paths: pulumi.StringArray{
// 						pulumi.String("/definition/id1"),
// 						pulumi.String("/definition/id2"),
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
// > **NOTE:** The CosmosDB Account needs to have the `EnableGremlin` capability enabled to use this resource - which can be done by adding this to the `capabilities` list within the `cosmosdb.Account` resource.
//
// ## Import
//
// Cosmos Gremlin Graphs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/gremlinGraph:GremlinGraph example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1/graphs/graphs1
// ```
type GremlinGraph struct {
	pulumi.CustomResourceState

	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringOutput                    `pulumi:"accountName"`
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayOutput `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	DefaultTtl pulumi.IntOutput `pulumi:"defaultTtl"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below.
	IndexPolicies GremlinGraphIndexPolicyArrayOutput `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringOutput `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrOutput `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayOutput `pulumi:"uniqueKeys"`
}

// NewGremlinGraph registers a new resource with the given unique name, arguments, and options.
func NewGremlinGraph(ctx *pulumi.Context,
	name string, args *GremlinGraphArgs, opts ...pulumi.ResourceOption) (*GremlinGraph, error) {
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
	var resource GremlinGraph
	err := ctx.RegisterResource("azure:cosmosdb/gremlinGraph:GremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinGraph gets an existing GremlinGraph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinGraphState, opts ...pulumi.ResourceOption) (*GremlinGraph, error) {
	var resource GremlinGraph
	err := ctx.ReadResource("azure:cosmosdb/gremlinGraph:GremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinGraph resources.
type gremlinGraphState struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName       *string                        `pulumi:"accountName"`
	AutoscaleSettings *GremlinGraphAutoscaleSettings `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicies []GremlinGraphConflictResolutionPolicy `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below.
	IndexPolicies []GremlinGraphIndexPolicy `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion *int `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

type GremlinGraphState struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringPtrInput
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrInput
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayInput
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	DefaultTtl pulumi.IntPtrInput
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below.
	IndexPolicies GremlinGraphIndexPolicyArrayInput
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayInput
}

func (GremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinGraphState)(nil)).Elem()
}

type gremlinGraphArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName       string                         `pulumi:"accountName"`
	AutoscaleSettings *GremlinGraphAutoscaleSettings `pulumi:"autoscaleSettings"`
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicies []GremlinGraphConflictResolutionPolicy `pulumi:"conflictResolutionPolicies"`
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below.
	IndexPolicies []GremlinGraphIndexPolicy `pulumi:"indexPolicies"`
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath string `pulumi:"partitionKeyPath"`
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion *int `pulumi:"partitionKeyVersion"`
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []GremlinGraphUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a GremlinGraph resource.
type GremlinGraphArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	AccountName       pulumi.StringInput
	AutoscaleSettings GremlinGraphAutoscaleSettingsPtrInput
	// A `conflictResolutionPolicy` blocks as defined below.
	ConflictResolutionPolicies GremlinGraphConflictResolutionPolicyArrayInput
	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	DefaultTtl pulumi.IntPtrInput
	// The configuration of the indexing policy. One or more `indexPolicy` blocks as defined below.
	IndexPolicies GremlinGraphIndexPolicyArrayInput
	// Specifies the name of the Cosmos DB Gremlin Graph. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringInput
	// Define a partition key version. Changing this forces a new resource to be created. Possible values are ` 1  `and `2`. This should be set to `2` in order to use large partition keys.
	PartitionKeyVersion pulumi.IntPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of the Gremlin graph (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys GremlinGraphUniqueKeyArrayInput
}

func (GremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinGraphArgs)(nil)).Elem()
}

type GremlinGraphInput interface {
	pulumi.Input

	ToGremlinGraphOutput() GremlinGraphOutput
	ToGremlinGraphOutputWithContext(ctx context.Context) GremlinGraphOutput
}

func (*GremlinGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraph)(nil)).Elem()
}

func (i *GremlinGraph) ToGremlinGraphOutput() GremlinGraphOutput {
	return i.ToGremlinGraphOutputWithContext(context.Background())
}

func (i *GremlinGraph) ToGremlinGraphOutputWithContext(ctx context.Context) GremlinGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphOutput)
}

// GremlinGraphArrayInput is an input type that accepts GremlinGraphArray and GremlinGraphArrayOutput values.
// You can construct a concrete instance of `GremlinGraphArrayInput` via:
//
//          GremlinGraphArray{ GremlinGraphArgs{...} }
type GremlinGraphArrayInput interface {
	pulumi.Input

	ToGremlinGraphArrayOutput() GremlinGraphArrayOutput
	ToGremlinGraphArrayOutputWithContext(context.Context) GremlinGraphArrayOutput
}

type GremlinGraphArray []GremlinGraphInput

func (GremlinGraphArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GremlinGraph)(nil)).Elem()
}

func (i GremlinGraphArray) ToGremlinGraphArrayOutput() GremlinGraphArrayOutput {
	return i.ToGremlinGraphArrayOutputWithContext(context.Background())
}

func (i GremlinGraphArray) ToGremlinGraphArrayOutputWithContext(ctx context.Context) GremlinGraphArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphArrayOutput)
}

// GremlinGraphMapInput is an input type that accepts GremlinGraphMap and GremlinGraphMapOutput values.
// You can construct a concrete instance of `GremlinGraphMapInput` via:
//
//          GremlinGraphMap{ "key": GremlinGraphArgs{...} }
type GremlinGraphMapInput interface {
	pulumi.Input

	ToGremlinGraphMapOutput() GremlinGraphMapOutput
	ToGremlinGraphMapOutputWithContext(context.Context) GremlinGraphMapOutput
}

type GremlinGraphMap map[string]GremlinGraphInput

func (GremlinGraphMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GremlinGraph)(nil)).Elem()
}

func (i GremlinGraphMap) ToGremlinGraphMapOutput() GremlinGraphMapOutput {
	return i.ToGremlinGraphMapOutputWithContext(context.Background())
}

func (i GremlinGraphMap) ToGremlinGraphMapOutputWithContext(ctx context.Context) GremlinGraphMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphMapOutput)
}

type GremlinGraphOutput struct{ *pulumi.OutputState }

func (GremlinGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraph)(nil)).Elem()
}

func (o GremlinGraphOutput) ToGremlinGraphOutput() GremlinGraphOutput {
	return o
}

func (o GremlinGraphOutput) ToGremlinGraphOutputWithContext(ctx context.Context) GremlinGraphOutput {
	return o
}

type GremlinGraphArrayOutput struct{ *pulumi.OutputState }

func (GremlinGraphArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GremlinGraph)(nil)).Elem()
}

func (o GremlinGraphArrayOutput) ToGremlinGraphArrayOutput() GremlinGraphArrayOutput {
	return o
}

func (o GremlinGraphArrayOutput) ToGremlinGraphArrayOutputWithContext(ctx context.Context) GremlinGraphArrayOutput {
	return o
}

func (o GremlinGraphArrayOutput) Index(i pulumi.IntInput) GremlinGraphOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GremlinGraph {
		return vs[0].([]*GremlinGraph)[vs[1].(int)]
	}).(GremlinGraphOutput)
}

type GremlinGraphMapOutput struct{ *pulumi.OutputState }

func (GremlinGraphMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GremlinGraph)(nil)).Elem()
}

func (o GremlinGraphMapOutput) ToGremlinGraphMapOutput() GremlinGraphMapOutput {
	return o
}

func (o GremlinGraphMapOutput) ToGremlinGraphMapOutputWithContext(ctx context.Context) GremlinGraphMapOutput {
	return o
}

func (o GremlinGraphMapOutput) MapIndex(k pulumi.StringInput) GremlinGraphOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GremlinGraph {
		return vs[0].(map[string]*GremlinGraph)[vs[1].(string)]
	}).(GremlinGraphOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GremlinGraphInput)(nil)).Elem(), &GremlinGraph{})
	pulumi.RegisterInputType(reflect.TypeOf((*GremlinGraphArrayInput)(nil)).Elem(), GremlinGraphArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GremlinGraphMapInput)(nil)).Elem(), GremlinGraphMap{})
	pulumi.RegisterOutputType(GremlinGraphOutput{})
	pulumi.RegisterOutputType(GremlinGraphArrayOutput{})
	pulumi.RegisterOutputType(GremlinGraphMapOutput{})
}
