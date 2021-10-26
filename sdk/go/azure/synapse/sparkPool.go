// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Spark Pool.
//
// ## Import
//
// Synapse Spark Pool can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/sparkPool:SparkPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
// ```
type SparkPool struct {
	pulumi.CustomResourceState

	// An `autoPause` block as defined below.
	AutoPause SparkPoolAutoPausePtrOutput `pulumi:"autoPause"`
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale SparkPoolAutoScalePtrOutput `pulumi:"autoScale"`
	// The cache size in the Spark Pool.
	CacheSize pulumi.IntPtrOutput `pulumi:"cacheSize"`
	// Indicates whether compute isolation is enabled or not. Defaults to `false`.
	ComputeIsolationEnabled pulumi.BoolPtrOutput `pulumi:"computeIsolationEnabled"`
	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to `false`.
	DynamicExecutorAllocationEnabled pulumi.BoolPtrOutput `pulumi:"dynamicExecutorAllocationEnabled"`
	// A `libraryRequirement` block as defined below.
	LibraryRequirement SparkPoolLibraryRequirementPtrOutput `pulumi:"libraryRequirement"`
	// The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
	NodeSize pulumi.StringOutput `pulumi:"nodeSize"`
	// The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
	NodeSizeFamily pulumi.StringOutput `pulumi:"nodeSizeFamily"`
	// Indicates whether session level packages are enabled or not. Defaults to `false`.
	SessionLevelPackagesEnabled pulumi.BoolPtrOutput `pulumi:"sessionLevelPackagesEnabled"`
	// A `sparkConfig` block as defined below.
	SparkConfig SparkPoolSparkConfigPtrOutput `pulumi:"sparkConfig"`
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrOutput `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrOutput `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible values are `2.4` and `3.1`. Defaults to `2.4`.
	SparkVersion pulumi.StringPtrOutput `pulumi:"sparkVersion"`
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSparkPool registers a new resource with the given unique name, arguments, and options.
func NewSparkPool(ctx *pulumi.Context,
	name string, args *SparkPoolArgs, opts ...pulumi.ResourceOption) (*SparkPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeSize == nil {
		return nil, errors.New("invalid value for required argument 'NodeSize'")
	}
	if args.NodeSizeFamily == nil {
		return nil, errors.New("invalid value for required argument 'NodeSizeFamily'")
	}
	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	var resource SparkPool
	err := ctx.RegisterResource("azure:synapse/sparkPool:SparkPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSparkPool gets an existing SparkPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSparkPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SparkPoolState, opts ...pulumi.ResourceOption) (*SparkPool, error) {
	var resource SparkPool
	err := ctx.ReadResource("azure:synapse/sparkPool:SparkPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SparkPool resources.
type sparkPoolState struct {
	// An `autoPause` block as defined below.
	AutoPause *SparkPoolAutoPause `pulumi:"autoPause"`
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale *SparkPoolAutoScale `pulumi:"autoScale"`
	// The cache size in the Spark Pool.
	CacheSize *int `pulumi:"cacheSize"`
	// Indicates whether compute isolation is enabled or not. Defaults to `false`.
	ComputeIsolationEnabled *bool `pulumi:"computeIsolationEnabled"`
	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to `false`.
	DynamicExecutorAllocationEnabled *bool `pulumi:"dynamicExecutorAllocationEnabled"`
	// A `libraryRequirement` block as defined below.
	LibraryRequirement *SparkPoolLibraryRequirement `pulumi:"libraryRequirement"`
	// The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
	Name *string `pulumi:"name"`
	// The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
	NodeCount *int `pulumi:"nodeCount"`
	// The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
	NodeSize *string `pulumi:"nodeSize"`
	// The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
	NodeSizeFamily *string `pulumi:"nodeSizeFamily"`
	// Indicates whether session level packages are enabled or not. Defaults to `false`.
	SessionLevelPackagesEnabled *bool `pulumi:"sessionLevelPackagesEnabled"`
	// A `sparkConfig` block as defined below.
	SparkConfig *SparkPoolSparkConfig `pulumi:"sparkConfig"`
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder *string `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible values are `2.4` and `3.1`. Defaults to `2.4`.
	SparkVersion *string `pulumi:"sparkVersion"`
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags map[string]string `pulumi:"tags"`
}

type SparkPoolState struct {
	// An `autoPause` block as defined below.
	AutoPause SparkPoolAutoPausePtrInput
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale SparkPoolAutoScalePtrInput
	// The cache size in the Spark Pool.
	CacheSize pulumi.IntPtrInput
	// Indicates whether compute isolation is enabled or not. Defaults to `false`.
	ComputeIsolationEnabled pulumi.BoolPtrInput
	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to `false`.
	DynamicExecutorAllocationEnabled pulumi.BoolPtrInput
	// A `libraryRequirement` block as defined below.
	LibraryRequirement SparkPoolLibraryRequirementPtrInput
	// The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
	Name pulumi.StringPtrInput
	// The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
	NodeCount pulumi.IntPtrInput
	// The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
	NodeSize pulumi.StringPtrInput
	// The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
	NodeSizeFamily pulumi.StringPtrInput
	// Indicates whether session level packages are enabled or not. Defaults to `false`.
	SessionLevelPackagesEnabled pulumi.BoolPtrInput
	// A `sparkConfig` block as defined below.
	SparkConfig SparkPoolSparkConfigPtrInput
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrInput
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrInput
	// The Apache Spark version. Possible values are `2.4` and `3.1`. Defaults to `2.4`.
	SparkVersion pulumi.StringPtrInput
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags pulumi.StringMapInput
}

func (SparkPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*sparkPoolState)(nil)).Elem()
}

type sparkPoolArgs struct {
	// An `autoPause` block as defined below.
	AutoPause *SparkPoolAutoPause `pulumi:"autoPause"`
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale *SparkPoolAutoScale `pulumi:"autoScale"`
	// The cache size in the Spark Pool.
	CacheSize *int `pulumi:"cacheSize"`
	// Indicates whether compute isolation is enabled or not. Defaults to `false`.
	ComputeIsolationEnabled *bool `pulumi:"computeIsolationEnabled"`
	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to `false`.
	DynamicExecutorAllocationEnabled *bool `pulumi:"dynamicExecutorAllocationEnabled"`
	// A `libraryRequirement` block as defined below.
	LibraryRequirement *SparkPoolLibraryRequirement `pulumi:"libraryRequirement"`
	// The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
	Name *string `pulumi:"name"`
	// The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
	NodeCount *int `pulumi:"nodeCount"`
	// The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
	NodeSize string `pulumi:"nodeSize"`
	// The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
	NodeSizeFamily string `pulumi:"nodeSizeFamily"`
	// Indicates whether session level packages are enabled or not. Defaults to `false`.
	SessionLevelPackagesEnabled *bool `pulumi:"sessionLevelPackagesEnabled"`
	// A `sparkConfig` block as defined below.
	SparkConfig *SparkPoolSparkConfig `pulumi:"sparkConfig"`
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder *string `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible values are `2.4` and `3.1`. Defaults to `2.4`.
	SparkVersion *string `pulumi:"sparkVersion"`
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SparkPool resource.
type SparkPoolArgs struct {
	// An `autoPause` block as defined below.
	AutoPause SparkPoolAutoPausePtrInput
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale SparkPoolAutoScalePtrInput
	// The cache size in the Spark Pool.
	CacheSize pulumi.IntPtrInput
	// Indicates whether compute isolation is enabled or not. Defaults to `false`.
	ComputeIsolationEnabled pulumi.BoolPtrInput
	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to `false`.
	DynamicExecutorAllocationEnabled pulumi.BoolPtrInput
	// A `libraryRequirement` block as defined below.
	LibraryRequirement SparkPoolLibraryRequirementPtrInput
	// The name which should be used for this Synapse Spark Pool. Changing this forces a new Synapse Spark Pool to be created.
	Name pulumi.StringPtrInput
	// The number of nodes in the Spark Pool. Exactly one of `nodeCount` or `autoScale` must be specified.
	NodeCount pulumi.IntPtrInput
	// The level of node in the Spark Pool. Possible value is `Small`, `Medium` and `Large`.
	NodeSize pulumi.StringInput
	// The kind of nodes that the Spark Pool provides. Possible value is `MemoryOptimized`.
	NodeSizeFamily pulumi.StringInput
	// Indicates whether session level packages are enabled or not. Defaults to `false`.
	SessionLevelPackagesEnabled pulumi.BoolPtrInput
	// A `sparkConfig` block as defined below.
	SparkConfig SparkPoolSparkConfigPtrInput
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrInput
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrInput
	// The Apache Spark version. Possible values are `2.4` and `3.1`. Defaults to `2.4`.
	SparkVersion pulumi.StringPtrInput
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId pulumi.StringInput
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags pulumi.StringMapInput
}

func (SparkPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sparkPoolArgs)(nil)).Elem()
}

type SparkPoolInput interface {
	pulumi.Input

	ToSparkPoolOutput() SparkPoolOutput
	ToSparkPoolOutputWithContext(ctx context.Context) SparkPoolOutput
}

func (*SparkPool) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkPool)(nil))
}

func (i *SparkPool) ToSparkPoolOutput() SparkPoolOutput {
	return i.ToSparkPoolOutputWithContext(context.Background())
}

func (i *SparkPool) ToSparkPoolOutputWithContext(ctx context.Context) SparkPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkPoolOutput)
}

func (i *SparkPool) ToSparkPoolPtrOutput() SparkPoolPtrOutput {
	return i.ToSparkPoolPtrOutputWithContext(context.Background())
}

func (i *SparkPool) ToSparkPoolPtrOutputWithContext(ctx context.Context) SparkPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkPoolPtrOutput)
}

type SparkPoolPtrInput interface {
	pulumi.Input

	ToSparkPoolPtrOutput() SparkPoolPtrOutput
	ToSparkPoolPtrOutputWithContext(ctx context.Context) SparkPoolPtrOutput
}

type sparkPoolPtrType SparkPoolArgs

func (*sparkPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkPool)(nil))
}

func (i *sparkPoolPtrType) ToSparkPoolPtrOutput() SparkPoolPtrOutput {
	return i.ToSparkPoolPtrOutputWithContext(context.Background())
}

func (i *sparkPoolPtrType) ToSparkPoolPtrOutputWithContext(ctx context.Context) SparkPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkPoolPtrOutput)
}

// SparkPoolArrayInput is an input type that accepts SparkPoolArray and SparkPoolArrayOutput values.
// You can construct a concrete instance of `SparkPoolArrayInput` via:
//
//          SparkPoolArray{ SparkPoolArgs{...} }
type SparkPoolArrayInput interface {
	pulumi.Input

	ToSparkPoolArrayOutput() SparkPoolArrayOutput
	ToSparkPoolArrayOutputWithContext(context.Context) SparkPoolArrayOutput
}

type SparkPoolArray []SparkPoolInput

func (SparkPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SparkPool)(nil)).Elem()
}

func (i SparkPoolArray) ToSparkPoolArrayOutput() SparkPoolArrayOutput {
	return i.ToSparkPoolArrayOutputWithContext(context.Background())
}

func (i SparkPoolArray) ToSparkPoolArrayOutputWithContext(ctx context.Context) SparkPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkPoolArrayOutput)
}

// SparkPoolMapInput is an input type that accepts SparkPoolMap and SparkPoolMapOutput values.
// You can construct a concrete instance of `SparkPoolMapInput` via:
//
//          SparkPoolMap{ "key": SparkPoolArgs{...} }
type SparkPoolMapInput interface {
	pulumi.Input

	ToSparkPoolMapOutput() SparkPoolMapOutput
	ToSparkPoolMapOutputWithContext(context.Context) SparkPoolMapOutput
}

type SparkPoolMap map[string]SparkPoolInput

func (SparkPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SparkPool)(nil)).Elem()
}

func (i SparkPoolMap) ToSparkPoolMapOutput() SparkPoolMapOutput {
	return i.ToSparkPoolMapOutputWithContext(context.Background())
}

func (i SparkPoolMap) ToSparkPoolMapOutputWithContext(ctx context.Context) SparkPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkPoolMapOutput)
}

type SparkPoolOutput struct{ *pulumi.OutputState }

func (SparkPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkPool)(nil))
}

func (o SparkPoolOutput) ToSparkPoolOutput() SparkPoolOutput {
	return o
}

func (o SparkPoolOutput) ToSparkPoolOutputWithContext(ctx context.Context) SparkPoolOutput {
	return o
}

func (o SparkPoolOutput) ToSparkPoolPtrOutput() SparkPoolPtrOutput {
	return o.ToSparkPoolPtrOutputWithContext(context.Background())
}

func (o SparkPoolOutput) ToSparkPoolPtrOutputWithContext(ctx context.Context) SparkPoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SparkPool) *SparkPool {
		return &v
	}).(SparkPoolPtrOutput)
}

type SparkPoolPtrOutput struct{ *pulumi.OutputState }

func (SparkPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkPool)(nil))
}

func (o SparkPoolPtrOutput) ToSparkPoolPtrOutput() SparkPoolPtrOutput {
	return o
}

func (o SparkPoolPtrOutput) ToSparkPoolPtrOutputWithContext(ctx context.Context) SparkPoolPtrOutput {
	return o
}

func (o SparkPoolPtrOutput) Elem() SparkPoolOutput {
	return o.ApplyT(func(v *SparkPool) SparkPool {
		if v != nil {
			return *v
		}
		var ret SparkPool
		return ret
	}).(SparkPoolOutput)
}

type SparkPoolArrayOutput struct{ *pulumi.OutputState }

func (SparkPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkPool)(nil))
}

func (o SparkPoolArrayOutput) ToSparkPoolArrayOutput() SparkPoolArrayOutput {
	return o
}

func (o SparkPoolArrayOutput) ToSparkPoolArrayOutputWithContext(ctx context.Context) SparkPoolArrayOutput {
	return o
}

func (o SparkPoolArrayOutput) Index(i pulumi.IntInput) SparkPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SparkPool {
		return vs[0].([]SparkPool)[vs[1].(int)]
	}).(SparkPoolOutput)
}

type SparkPoolMapOutput struct{ *pulumi.OutputState }

func (SparkPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SparkPool)(nil))
}

func (o SparkPoolMapOutput) ToSparkPoolMapOutput() SparkPoolMapOutput {
	return o
}

func (o SparkPoolMapOutput) ToSparkPoolMapOutputWithContext(ctx context.Context) SparkPoolMapOutput {
	return o
}

func (o SparkPoolMapOutput) MapIndex(k pulumi.StringInput) SparkPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SparkPool {
		return vs[0].(map[string]SparkPool)[vs[1].(string)]
	}).(SparkPoolOutput)
}

func init() {
	pulumi.RegisterOutputType(SparkPoolOutput{})
	pulumi.RegisterOutputType(SparkPoolPtrOutput{})
	pulumi.RegisterOutputType(SparkPoolArrayOutput{})
	pulumi.RegisterOutputType(SparkPoolMapOutput{})
}
