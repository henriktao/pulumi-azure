// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Mongo Collection within a Cosmos DB Account.
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
// 		exampleMongoDatabase, err := cosmosdb.NewMongoDatabase(ctx, "exampleMongoDatabase", &cosmosdb.MongoDatabaseArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewMongoCollection(ctx, "exampleMongoCollection", &cosmosdb.MongoCollectionArgs{
// 			ResourceGroupName: pulumi.String(exampleAccount.ResourceGroupName),
// 			AccountName:       pulumi.String(exampleAccount.Name),
// 			DatabaseName:      exampleMongoDatabase.Name,
// 			DefaultTtlSeconds: pulumi.Int(777),
// 			ShardKey:          pulumi.String("uniqueKey"),
// 			Throughput:        pulumi.Int(400),
// 			Indices: cosmosdb.MongoCollectionIndexArray{
// 				&cosmosdb.MongoCollectionIndexArgs{
// 					Keys: pulumi.StringArray{
// 						pulumi.String("_id"),
// 					},
// 					Unique: pulumi.Bool(true),
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
// CosmosDB Mongo Collection can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/mongoCollection:MongoCollection collection1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1/collections/collection1
// ```
type MongoCollection struct {
	pulumi.CustomResourceState

	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrOutput `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
	AutoscaleSettings MongoCollectionAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
	DefaultTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultTtlSeconds"`
	// One or more `index` blocks as defined below.
	Indices MongoCollectionIndexArrayOutput `pulumi:"indices"`
	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the key to partition on for sharding. There must not be any other unique index keys.
	ShardKey pulumi.StringPtrOutput `pulumi:"shardKey"`
	// One or more `systemIndexes` blocks as defined below.
	SystemIndexes MongoCollectionSystemIndexArrayOutput `pulumi:"systemIndexes"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewMongoCollection registers a new resource with the given unique name, arguments, and options.
func NewMongoCollection(ctx *pulumi.Context,
	name string, args *MongoCollectionArgs, opts ...pulumi.ResourceOption) (*MongoCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource MongoCollection
	err := ctx.RegisterResource("azure:cosmosdb/mongoCollection:MongoCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoCollection gets an existing MongoCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoCollectionState, opts ...pulumi.ResourceOption) (*MongoCollection, error) {
	var resource MongoCollection
	err := ctx.ReadResource("azure:cosmosdb/mongoCollection:MongoCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoCollection resources.
type mongoCollectionState struct {
	AccountName *string `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl *int `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
	AutoscaleSettings *MongoCollectionAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
	DefaultTtlSeconds *int `pulumi:"defaultTtlSeconds"`
	// One or more `index` blocks as defined below.
	Indices []MongoCollectionIndex `pulumi:"indices"`
	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the key to partition on for sharding. There must not be any other unique index keys.
	ShardKey *string `pulumi:"shardKey"`
	// One or more `systemIndexes` blocks as defined below.
	SystemIndexes []MongoCollectionSystemIndex `pulumi:"systemIndexes"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

type MongoCollectionState struct {
	AccountName pulumi.StringPtrInput
	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
	AutoscaleSettings MongoCollectionAutoscaleSettingsPtrInput
	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
	DefaultTtlSeconds pulumi.IntPtrInput
	// One or more `index` blocks as defined below.
	Indices MongoCollectionIndexArrayInput
	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the key to partition on for sharding. There must not be any other unique index keys.
	ShardKey pulumi.StringPtrInput
	// One or more `systemIndexes` blocks as defined below.
	SystemIndexes MongoCollectionSystemIndexArrayInput
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (MongoCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoCollectionState)(nil)).Elem()
}

type mongoCollectionArgs struct {
	AccountName string `pulumi:"accountName"`
	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl *int `pulumi:"analyticalStorageTtl"`
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
	AutoscaleSettings *MongoCollectionAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
	DefaultTtlSeconds *int `pulumi:"defaultTtlSeconds"`
	// One or more `index` blocks as defined below.
	Indices []MongoCollectionIndex `pulumi:"indices"`
	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the key to partition on for sharding. There must not be any other unique index keys.
	ShardKey *string `pulumi:"shardKey"`
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput *int `pulumi:"throughput"`
}

// The set of arguments for constructing a MongoCollection resource.
type MongoCollectionArgs struct {
	AccountName pulumi.StringInput
	// The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	AnalyticalStorageTtl pulumi.IntPtrInput
	// An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
	AutoscaleSettings MongoCollectionAutoscaleSettingsPtrInput
	// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
	DefaultTtlSeconds pulumi.IntPtrInput
	// One or more `index` blocks as defined below.
	Indices MongoCollectionIndexArrayInput
	// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the key to partition on for sharding. There must not be any other unique index keys.
	ShardKey pulumi.StringPtrInput
	// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
	Throughput pulumi.IntPtrInput
}

func (MongoCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoCollectionArgs)(nil)).Elem()
}

type MongoCollectionInput interface {
	pulumi.Input

	ToMongoCollectionOutput() MongoCollectionOutput
	ToMongoCollectionOutputWithContext(ctx context.Context) MongoCollectionOutput
}

func (*MongoCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCollection)(nil)).Elem()
}

func (i *MongoCollection) ToMongoCollectionOutput() MongoCollectionOutput {
	return i.ToMongoCollectionOutputWithContext(context.Background())
}

func (i *MongoCollection) ToMongoCollectionOutputWithContext(ctx context.Context) MongoCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoCollectionOutput)
}

// MongoCollectionArrayInput is an input type that accepts MongoCollectionArray and MongoCollectionArrayOutput values.
// You can construct a concrete instance of `MongoCollectionArrayInput` via:
//
//          MongoCollectionArray{ MongoCollectionArgs{...} }
type MongoCollectionArrayInput interface {
	pulumi.Input

	ToMongoCollectionArrayOutput() MongoCollectionArrayOutput
	ToMongoCollectionArrayOutputWithContext(context.Context) MongoCollectionArrayOutput
}

type MongoCollectionArray []MongoCollectionInput

func (MongoCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoCollection)(nil)).Elem()
}

func (i MongoCollectionArray) ToMongoCollectionArrayOutput() MongoCollectionArrayOutput {
	return i.ToMongoCollectionArrayOutputWithContext(context.Background())
}

func (i MongoCollectionArray) ToMongoCollectionArrayOutputWithContext(ctx context.Context) MongoCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoCollectionArrayOutput)
}

// MongoCollectionMapInput is an input type that accepts MongoCollectionMap and MongoCollectionMapOutput values.
// You can construct a concrete instance of `MongoCollectionMapInput` via:
//
//          MongoCollectionMap{ "key": MongoCollectionArgs{...} }
type MongoCollectionMapInput interface {
	pulumi.Input

	ToMongoCollectionMapOutput() MongoCollectionMapOutput
	ToMongoCollectionMapOutputWithContext(context.Context) MongoCollectionMapOutput
}

type MongoCollectionMap map[string]MongoCollectionInput

func (MongoCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoCollection)(nil)).Elem()
}

func (i MongoCollectionMap) ToMongoCollectionMapOutput() MongoCollectionMapOutput {
	return i.ToMongoCollectionMapOutputWithContext(context.Background())
}

func (i MongoCollectionMap) ToMongoCollectionMapOutputWithContext(ctx context.Context) MongoCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoCollectionMapOutput)
}

type MongoCollectionOutput struct{ *pulumi.OutputState }

func (MongoCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCollection)(nil)).Elem()
}

func (o MongoCollectionOutput) ToMongoCollectionOutput() MongoCollectionOutput {
	return o
}

func (o MongoCollectionOutput) ToMongoCollectionOutputWithContext(ctx context.Context) MongoCollectionOutput {
	return o
}

type MongoCollectionArrayOutput struct{ *pulumi.OutputState }

func (MongoCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoCollection)(nil)).Elem()
}

func (o MongoCollectionArrayOutput) ToMongoCollectionArrayOutput() MongoCollectionArrayOutput {
	return o
}

func (o MongoCollectionArrayOutput) ToMongoCollectionArrayOutputWithContext(ctx context.Context) MongoCollectionArrayOutput {
	return o
}

func (o MongoCollectionArrayOutput) Index(i pulumi.IntInput) MongoCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MongoCollection {
		return vs[0].([]*MongoCollection)[vs[1].(int)]
	}).(MongoCollectionOutput)
}

type MongoCollectionMapOutput struct{ *pulumi.OutputState }

func (MongoCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoCollection)(nil)).Elem()
}

func (o MongoCollectionMapOutput) ToMongoCollectionMapOutput() MongoCollectionMapOutput {
	return o
}

func (o MongoCollectionMapOutput) ToMongoCollectionMapOutputWithContext(ctx context.Context) MongoCollectionMapOutput {
	return o
}

func (o MongoCollectionMapOutput) MapIndex(k pulumi.StringInput) MongoCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MongoCollection {
		return vs[0].(map[string]*MongoCollection)[vs[1].(string)]
	}).(MongoCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MongoCollectionInput)(nil)).Elem(), &MongoCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoCollectionArrayInput)(nil)).Elem(), MongoCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoCollectionMapInput)(nil)).Elem(), MongoCollectionMap{})
	pulumi.RegisterOutputType(MongoCollectionOutput{})
	pulumi.RegisterOutputType(MongoCollectionArrayOutput{})
	pulumi.RegisterOutputType(MongoCollectionMapOutput{})
}
