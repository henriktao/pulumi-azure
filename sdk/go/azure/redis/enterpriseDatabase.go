// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Redis Enterprise Database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/redis"
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
// 		exampleEnterpriseCluster, err := redis.NewEnterpriseCluster(ctx, "exampleEnterpriseCluster", &redis.EnterpriseClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("Enterprise_E20-4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redis.NewEnterpriseDatabase(ctx, "exampleEnterpriseDatabase", &redis.EnterpriseDatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ClusterId:         exampleEnterpriseCluster.ID(),
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
// Redis Enterprise Databases can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:redis/enterpriseDatabase:EnterpriseDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1/databases/database1
// ```
type EnterpriseDatabase struct {
	pulumi.CustomResourceState

	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol pulumi.StringPtrOutput `pulumi:"clientProtocol"`
	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy pulumi.StringPtrOutput `pulumi:"clusteringPolicy"`
	// Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// A `module` block as defined below.
	Modules EnterpriseDatabaseModuleArrayOutput `pulumi:"modules"`
	// The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The Primary Access Key for the Redis Enterprise Database Instance.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Redis Enterprise Database Instance.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
}

// NewEnterpriseDatabase registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseDatabase(ctx *pulumi.Context,
	name string, args *EnterpriseDatabaseArgs, opts ...pulumi.ResourceOption) (*EnterpriseDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EnterpriseDatabase
	err := ctx.RegisterResource("azure:redis/enterpriseDatabase:EnterpriseDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseDatabase gets an existing EnterpriseDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseDatabaseState, opts ...pulumi.ResourceOption) (*EnterpriseDatabase, error) {
	var resource EnterpriseDatabase
	err := ctx.ReadResource("azure:redis/enterpriseDatabase:EnterpriseDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseDatabase resources.
type enterpriseDatabaseState struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol *string `pulumi:"clientProtocol"`
	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterId *string `pulumi:"clusterId"`
	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy *string `pulumi:"clusteringPolicy"`
	// Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// A `module` block as defined below.
	Modules []EnterpriseDatabaseModule `pulumi:"modules"`
	// The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
	Name *string `pulumi:"name"`
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
	Port *int `pulumi:"port"`
	// The Primary Access Key for the Redis Enterprise Database Instance.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Redis Enterprise Database Instance.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
}

type EnterpriseDatabaseState struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol pulumi.StringPtrInput
	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterId pulumi.StringPtrInput
	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy pulumi.StringPtrInput
	// Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy pulumi.StringPtrInput
	// A `module` block as defined below.
	Modules EnterpriseDatabaseModuleArrayInput
	// The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
	Name pulumi.StringPtrInput
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
	Port pulumi.IntPtrInput
	// The Primary Access Key for the Redis Enterprise Database Instance.
	PrimaryAccessKey pulumi.StringPtrInput
	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Access Key for the Redis Enterprise Database Instance.
	SecondaryAccessKey pulumi.StringPtrInput
}

func (EnterpriseDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseDatabaseState)(nil)).Elem()
}

type enterpriseDatabaseArgs struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol *string `pulumi:"clientProtocol"`
	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterId string `pulumi:"clusterId"`
	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy *string `pulumi:"clusteringPolicy"`
	// Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// A `module` block as defined below.
	Modules []EnterpriseDatabaseModule `pulumi:"modules"`
	// The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
	Name *string `pulumi:"name"`
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
	Port *int `pulumi:"port"`
	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EnterpriseDatabase resource.
type EnterpriseDatabaseArgs struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are `Encrypted` and `Plaintext`. Defaults to `Encrypted`. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol pulumi.StringPtrInput
	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterId pulumi.StringInput
	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are `EnterpriseCluster` and `OSSCluster`. Defaults to `OSSCluster`. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy pulumi.StringPtrInput
	// Redis eviction policy - default is VolatileLRU. Possible values are `AllKeysLFU`, `AllKeysLRU`, `AllKeysRandom`, `VolatileLRU`, `VolatileLFU`, `VolatileTTL`, `VolatileRandom` and `NoEviction`. Defaults to `VolatileLRU`. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy pulumi.StringPtrInput
	// A `module` block as defined below.
	Modules EnterpriseDatabaseModuleArrayInput
	// The name which should be used for this Redis Enterprise Database. Currently the acceptable value for this argument is `default`. Defaults to `default`. Changing this forces a new Redis Enterprise Database to be created.
	Name pulumi.StringPtrInput
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created.
	Port pulumi.IntPtrInput
	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName pulumi.StringInput
}

func (EnterpriseDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseDatabaseArgs)(nil)).Elem()
}

type EnterpriseDatabaseInput interface {
	pulumi.Input

	ToEnterpriseDatabaseOutput() EnterpriseDatabaseOutput
	ToEnterpriseDatabaseOutputWithContext(ctx context.Context) EnterpriseDatabaseOutput
}

func (*EnterpriseDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseDatabase)(nil)).Elem()
}

func (i *EnterpriseDatabase) ToEnterpriseDatabaseOutput() EnterpriseDatabaseOutput {
	return i.ToEnterpriseDatabaseOutputWithContext(context.Background())
}

func (i *EnterpriseDatabase) ToEnterpriseDatabaseOutputWithContext(ctx context.Context) EnterpriseDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseDatabaseOutput)
}

// EnterpriseDatabaseArrayInput is an input type that accepts EnterpriseDatabaseArray and EnterpriseDatabaseArrayOutput values.
// You can construct a concrete instance of `EnterpriseDatabaseArrayInput` via:
//
//          EnterpriseDatabaseArray{ EnterpriseDatabaseArgs{...} }
type EnterpriseDatabaseArrayInput interface {
	pulumi.Input

	ToEnterpriseDatabaseArrayOutput() EnterpriseDatabaseArrayOutput
	ToEnterpriseDatabaseArrayOutputWithContext(context.Context) EnterpriseDatabaseArrayOutput
}

type EnterpriseDatabaseArray []EnterpriseDatabaseInput

func (EnterpriseDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseDatabase)(nil)).Elem()
}

func (i EnterpriseDatabaseArray) ToEnterpriseDatabaseArrayOutput() EnterpriseDatabaseArrayOutput {
	return i.ToEnterpriseDatabaseArrayOutputWithContext(context.Background())
}

func (i EnterpriseDatabaseArray) ToEnterpriseDatabaseArrayOutputWithContext(ctx context.Context) EnterpriseDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseDatabaseArrayOutput)
}

// EnterpriseDatabaseMapInput is an input type that accepts EnterpriseDatabaseMap and EnterpriseDatabaseMapOutput values.
// You can construct a concrete instance of `EnterpriseDatabaseMapInput` via:
//
//          EnterpriseDatabaseMap{ "key": EnterpriseDatabaseArgs{...} }
type EnterpriseDatabaseMapInput interface {
	pulumi.Input

	ToEnterpriseDatabaseMapOutput() EnterpriseDatabaseMapOutput
	ToEnterpriseDatabaseMapOutputWithContext(context.Context) EnterpriseDatabaseMapOutput
}

type EnterpriseDatabaseMap map[string]EnterpriseDatabaseInput

func (EnterpriseDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseDatabase)(nil)).Elem()
}

func (i EnterpriseDatabaseMap) ToEnterpriseDatabaseMapOutput() EnterpriseDatabaseMapOutput {
	return i.ToEnterpriseDatabaseMapOutputWithContext(context.Background())
}

func (i EnterpriseDatabaseMap) ToEnterpriseDatabaseMapOutputWithContext(ctx context.Context) EnterpriseDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseDatabaseMapOutput)
}

type EnterpriseDatabaseOutput struct{ *pulumi.OutputState }

func (EnterpriseDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseDatabase)(nil)).Elem()
}

func (o EnterpriseDatabaseOutput) ToEnterpriseDatabaseOutput() EnterpriseDatabaseOutput {
	return o
}

func (o EnterpriseDatabaseOutput) ToEnterpriseDatabaseOutputWithContext(ctx context.Context) EnterpriseDatabaseOutput {
	return o
}

type EnterpriseDatabaseArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseDatabase)(nil)).Elem()
}

func (o EnterpriseDatabaseArrayOutput) ToEnterpriseDatabaseArrayOutput() EnterpriseDatabaseArrayOutput {
	return o
}

func (o EnterpriseDatabaseArrayOutput) ToEnterpriseDatabaseArrayOutputWithContext(ctx context.Context) EnterpriseDatabaseArrayOutput {
	return o
}

func (o EnterpriseDatabaseArrayOutput) Index(i pulumi.IntInput) EnterpriseDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseDatabase {
		return vs[0].([]*EnterpriseDatabase)[vs[1].(int)]
	}).(EnterpriseDatabaseOutput)
}

type EnterpriseDatabaseMapOutput struct{ *pulumi.OutputState }

func (EnterpriseDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseDatabase)(nil)).Elem()
}

func (o EnterpriseDatabaseMapOutput) ToEnterpriseDatabaseMapOutput() EnterpriseDatabaseMapOutput {
	return o
}

func (o EnterpriseDatabaseMapOutput) ToEnterpriseDatabaseMapOutputWithContext(ctx context.Context) EnterpriseDatabaseMapOutput {
	return o
}

func (o EnterpriseDatabaseMapOutput) MapIndex(k pulumi.StringInput) EnterpriseDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseDatabase {
		return vs[0].(map[string]*EnterpriseDatabase)[vs[1].(string)]
	}).(EnterpriseDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseDatabaseInput)(nil)).Elem(), &EnterpriseDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseDatabaseArrayInput)(nil)).Elem(), EnterpriseDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseDatabaseMapInput)(nil)).Elem(), EnterpriseDatabaseMap{})
	pulumi.RegisterOutputType(EnterpriseDatabaseOutput{})
	pulumi.RegisterOutputType(EnterpriseDatabaseArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseDatabaseMapOutput{})
}
