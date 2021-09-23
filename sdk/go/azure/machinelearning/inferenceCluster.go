// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Machine Learning Inference Cluster.
//
// > **NOTE:** The Machine Learning Inference Cluster resource is used to attach an existing AKS cluster to the Machine Learning Workspace, it doesn't create the AKS cluster itself. Therefore it can only be created and deleted, not updated. Any change to the configuration will recreate the resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/machinelearning"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("west europe"),
// 			Tags: pulumi.StringMap{
// 				"stage": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			TenantId:               pulumi.String(current.TenantId),
// 			SkuName:                pulumi.String("standard"),
// 			PurgeProtectionEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkspace, err := machinelearning.NewWorkspace(ctx, "exampleWorkspace", &machinelearning.WorkspaceArgs{
// 			Location:              exampleResourceGroup.Location,
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			KeyVaultId:            exampleKeyVault.ID(),
// 			StorageAccountId:      exampleAccount.ID(),
// 			Identity: &machinelearning.WorkspaceIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefix:      pulumi.String("10.1.0.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKubernetesCluster, err := containerservice.NewKubernetesCluster(ctx, "exampleKubernetesCluster", &containerservice.KubernetesClusterArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
// 				Name:         pulumi.String("default"),
// 				NodeCount:    pulumi.Int(3),
// 				VmSize:       pulumi.String("Standard_D3_v2"),
// 				VnetSubnetId: exampleSubnet.ID(),
// 			},
// 			Identity: &containerservice.KubernetesClusterIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = machinelearning.NewInferenceCluster(ctx, "exampleInferenceCluster", &machinelearning.InferenceClusterArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ClusterPurpose:             pulumi.String("FastProd"),
// 			KubernetesClusterId:        exampleKubernetesCluster.ID(),
// 			Description:                pulumi.String("This is an example cluster used with Terraform"),
// 			MachineLearningWorkspaceId: exampleWorkspace.ID(),
// 			Tags: pulumi.StringMap{
// 				"stage": pulumi.String("example"),
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
// Machine Learning Inference Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:machinelearning/inferenceCluster:InferenceCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
// ```
type InferenceCluster struct {
	pulumi.CustomResourceState

	// The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
	ClusterPurpose pulumi.StringPtrOutput `pulumi:"clusterPurpose"`
	// The description of the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Identity InferenceClusterIdentityPtrOutput `pulumi:"identity"`
	// The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	KubernetesClusterId pulumi.StringOutput `pulumi:"kubernetesClusterId"`
	// The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringOutput `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Ssl InferenceClusterSslPtrOutput `pulumi:"ssl"`
	// A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewInferenceCluster registers a new resource with the given unique name, arguments, and options.
func NewInferenceCluster(ctx *pulumi.Context,
	name string, args *InferenceClusterArgs, opts ...pulumi.ResourceOption) (*InferenceCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesClusterId == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesClusterId'")
	}
	if args.MachineLearningWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'MachineLearningWorkspaceId'")
	}
	var resource InferenceCluster
	err := ctx.RegisterResource("azure:machinelearning/inferenceCluster:InferenceCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInferenceCluster gets an existing InferenceCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInferenceCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InferenceClusterState, opts ...pulumi.ResourceOption) (*InferenceCluster, error) {
	var resource InferenceCluster
	err := ctx.ReadResource("azure:machinelearning/inferenceCluster:InferenceCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InferenceCluster resources.
type inferenceClusterState struct {
	// The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
	ClusterPurpose *string `pulumi:"clusterPurpose"`
	// The description of the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Description *string `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Identity *InferenceClusterIdentity `pulumi:"identity"`
	// The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	KubernetesClusterId *string `pulumi:"kubernetesClusterId"`
	// The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
	MachineLearningWorkspaceId *string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Name *string `pulumi:"name"`
	// A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Ssl *InferenceClusterSsl `pulumi:"ssl"`
	// A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
}

type InferenceClusterState struct {
	// The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
	ClusterPurpose pulumi.StringPtrInput
	// The description of the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Description pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Identity InferenceClusterIdentityPtrInput
	// The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	KubernetesClusterId pulumi.StringPtrInput
	// The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Name pulumi.StringPtrInput
	// A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Ssl InferenceClusterSslPtrInput
	// A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Tags pulumi.StringMapInput
}

func (InferenceClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceClusterState)(nil)).Elem()
}

type inferenceClusterArgs struct {
	// The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
	ClusterPurpose *string `pulumi:"clusterPurpose"`
	// The description of the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Description *string `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Identity *InferenceClusterIdentity `pulumi:"identity"`
	// The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	KubernetesClusterId string `pulumi:"kubernetesClusterId"`
	// The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
	MachineLearningWorkspaceId string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Name *string `pulumi:"name"`
	// A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Ssl *InferenceClusterSsl `pulumi:"ssl"`
	// A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InferenceCluster resource.
type InferenceClusterArgs struct {
	// The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
	ClusterPurpose pulumi.StringPtrInput
	// The description of the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Description pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Identity InferenceClusterIdentityPtrInput
	// The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	KubernetesClusterId pulumi.StringInput
	// The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringInput
	// The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Name pulumi.StringPtrInput
	// A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
	Ssl InferenceClusterSslPtrInput
	// A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
	Tags pulumi.StringMapInput
}

func (InferenceClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceClusterArgs)(nil)).Elem()
}

type InferenceClusterInput interface {
	pulumi.Input

	ToInferenceClusterOutput() InferenceClusterOutput
	ToInferenceClusterOutputWithContext(ctx context.Context) InferenceClusterOutput
}

func (*InferenceCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceCluster)(nil))
}

func (i *InferenceCluster) ToInferenceClusterOutput() InferenceClusterOutput {
	return i.ToInferenceClusterOutputWithContext(context.Background())
}

func (i *InferenceCluster) ToInferenceClusterOutputWithContext(ctx context.Context) InferenceClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterOutput)
}

func (i *InferenceCluster) ToInferenceClusterPtrOutput() InferenceClusterPtrOutput {
	return i.ToInferenceClusterPtrOutputWithContext(context.Background())
}

func (i *InferenceCluster) ToInferenceClusterPtrOutputWithContext(ctx context.Context) InferenceClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterPtrOutput)
}

type InferenceClusterPtrInput interface {
	pulumi.Input

	ToInferenceClusterPtrOutput() InferenceClusterPtrOutput
	ToInferenceClusterPtrOutputWithContext(ctx context.Context) InferenceClusterPtrOutput
}

type inferenceClusterPtrType InferenceClusterArgs

func (*inferenceClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceCluster)(nil))
}

func (i *inferenceClusterPtrType) ToInferenceClusterPtrOutput() InferenceClusterPtrOutput {
	return i.ToInferenceClusterPtrOutputWithContext(context.Background())
}

func (i *inferenceClusterPtrType) ToInferenceClusterPtrOutputWithContext(ctx context.Context) InferenceClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterPtrOutput)
}

// InferenceClusterArrayInput is an input type that accepts InferenceClusterArray and InferenceClusterArrayOutput values.
// You can construct a concrete instance of `InferenceClusterArrayInput` via:
//
//          InferenceClusterArray{ InferenceClusterArgs{...} }
type InferenceClusterArrayInput interface {
	pulumi.Input

	ToInferenceClusterArrayOutput() InferenceClusterArrayOutput
	ToInferenceClusterArrayOutputWithContext(context.Context) InferenceClusterArrayOutput
}

type InferenceClusterArray []InferenceClusterInput

func (InferenceClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InferenceCluster)(nil)).Elem()
}

func (i InferenceClusterArray) ToInferenceClusterArrayOutput() InferenceClusterArrayOutput {
	return i.ToInferenceClusterArrayOutputWithContext(context.Background())
}

func (i InferenceClusterArray) ToInferenceClusterArrayOutputWithContext(ctx context.Context) InferenceClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterArrayOutput)
}

// InferenceClusterMapInput is an input type that accepts InferenceClusterMap and InferenceClusterMapOutput values.
// You can construct a concrete instance of `InferenceClusterMapInput` via:
//
//          InferenceClusterMap{ "key": InferenceClusterArgs{...} }
type InferenceClusterMapInput interface {
	pulumi.Input

	ToInferenceClusterMapOutput() InferenceClusterMapOutput
	ToInferenceClusterMapOutputWithContext(context.Context) InferenceClusterMapOutput
}

type InferenceClusterMap map[string]InferenceClusterInput

func (InferenceClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InferenceCluster)(nil)).Elem()
}

func (i InferenceClusterMap) ToInferenceClusterMapOutput() InferenceClusterMapOutput {
	return i.ToInferenceClusterMapOutputWithContext(context.Background())
}

func (i InferenceClusterMap) ToInferenceClusterMapOutputWithContext(ctx context.Context) InferenceClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterMapOutput)
}

type InferenceClusterOutput struct{ *pulumi.OutputState }

func (InferenceClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceCluster)(nil))
}

func (o InferenceClusterOutput) ToInferenceClusterOutput() InferenceClusterOutput {
	return o
}

func (o InferenceClusterOutput) ToInferenceClusterOutputWithContext(ctx context.Context) InferenceClusterOutput {
	return o
}

func (o InferenceClusterOutput) ToInferenceClusterPtrOutput() InferenceClusterPtrOutput {
	return o.ToInferenceClusterPtrOutputWithContext(context.Background())
}

func (o InferenceClusterOutput) ToInferenceClusterPtrOutputWithContext(ctx context.Context) InferenceClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InferenceCluster) *InferenceCluster {
		return &v
	}).(InferenceClusterPtrOutput)
}

type InferenceClusterPtrOutput struct{ *pulumi.OutputState }

func (InferenceClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceCluster)(nil))
}

func (o InferenceClusterPtrOutput) ToInferenceClusterPtrOutput() InferenceClusterPtrOutput {
	return o
}

func (o InferenceClusterPtrOutput) ToInferenceClusterPtrOutputWithContext(ctx context.Context) InferenceClusterPtrOutput {
	return o
}

func (o InferenceClusterPtrOutput) Elem() InferenceClusterOutput {
	return o.ApplyT(func(v *InferenceCluster) InferenceCluster {
		if v != nil {
			return *v
		}
		var ret InferenceCluster
		return ret
	}).(InferenceClusterOutput)
}

type InferenceClusterArrayOutput struct{ *pulumi.OutputState }

func (InferenceClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InferenceCluster)(nil))
}

func (o InferenceClusterArrayOutput) ToInferenceClusterArrayOutput() InferenceClusterArrayOutput {
	return o
}

func (o InferenceClusterArrayOutput) ToInferenceClusterArrayOutputWithContext(ctx context.Context) InferenceClusterArrayOutput {
	return o
}

func (o InferenceClusterArrayOutput) Index(i pulumi.IntInput) InferenceClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InferenceCluster {
		return vs[0].([]InferenceCluster)[vs[1].(int)]
	}).(InferenceClusterOutput)
}

type InferenceClusterMapOutput struct{ *pulumi.OutputState }

func (InferenceClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InferenceCluster)(nil))
}

func (o InferenceClusterMapOutput) ToInferenceClusterMapOutput() InferenceClusterMapOutput {
	return o
}

func (o InferenceClusterMapOutput) ToInferenceClusterMapOutputWithContext(ctx context.Context) InferenceClusterMapOutput {
	return o
}

func (o InferenceClusterMapOutput) MapIndex(k pulumi.StringInput) InferenceClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InferenceCluster {
		return vs[0].(map[string]InferenceCluster)[vs[1].(string)]
	}).(InferenceClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(InferenceClusterOutput{})
	pulumi.RegisterOutputType(InferenceClusterPtrOutput{})
	pulumi.RegisterOutputType(InferenceClusterArrayOutput{})
	pulumi.RegisterOutputType(InferenceClusterMapOutput{})
}
