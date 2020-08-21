// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Synapse Spark Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/synapse"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 			AccountKind:            pulumi.String("StorageV2"),
// 			IsHnsEnabled:           pulumi.Bool(true),
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
// 		_, err = synapse.NewSparkPool(ctx, "exampleSparkPool", &synapse.SparkPoolArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			NodeSizeFamily:     pulumi.String("MemoryOptimized"),
// 			NodeSize:           pulumi.String("Small"),
// 			AutoScale: &synapse.SparkPoolAutoScaleArgs{
// 				MaxNodeCount: pulumi.Int(50),
// 				MinNodeCount: pulumi.Int(3),
// 			},
// 			AutoPause: &synapse.SparkPoolAutoPauseArgs{
// 				DelayInMinutes: pulumi.Int(15),
// 			},
// 			Tags: pulumi.StringMap{
// 				"ENV": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SparkPool struct {
	pulumi.CustomResourceState

	// An `autoPause` block as defined below.
	AutoPause SparkPoolAutoPausePtrOutput `pulumi:"autoPause"`
	// An `autoScale` block as defined below. Exactly one of `nodeCount` or `autoScale` must be specified.
	AutoScale SparkPoolAutoScalePtrOutput `pulumi:"autoScale"`
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
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrOutput `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrOutput `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
	SparkVersion pulumi.StringPtrOutput `pulumi:"sparkVersion"`
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSparkPool registers a new resource with the given unique name, arguments, and options.
func NewSparkPool(ctx *pulumi.Context,
	name string, args *SparkPoolArgs, opts ...pulumi.ResourceOption) (*SparkPool, error) {
	if args == nil || args.NodeSize == nil {
		return nil, errors.New("missing required argument 'NodeSize'")
	}
	if args == nil || args.NodeSizeFamily == nil {
		return nil, errors.New("missing required argument 'NodeSizeFamily'")
	}
	if args == nil || args.SynapseWorkspaceId == nil {
		return nil, errors.New("missing required argument 'SynapseWorkspaceId'")
	}
	if args == nil {
		args = &SparkPoolArgs{}
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
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder *string `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
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
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrInput
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrInput
	// The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
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
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder *string `pulumi:"sparkLogFolder"`
	// The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
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
	// The Spark events folder. Defaults to `/events`.
	SparkEventsFolder pulumi.StringPtrInput
	// The default folder where Spark logs will be written. Defaults to `/logs`.
	SparkLogFolder pulumi.StringPtrInput
	// The Apache Spark version. Possible value is `2.4`. Defaults to `2.4`.
	SparkVersion pulumi.StringPtrInput
	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	SynapseWorkspaceId pulumi.StringInput
	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	Tags pulumi.StringMapInput
}

func (SparkPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sparkPoolArgs)(nil)).Elem()
}
