// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Machine Learning Compute Cluster.
// **NOTE:** At this point in time the resource cannot be updated (not supported by the backend Azure Go SDK). Therefore it can only be created and deleted, not updated. At the moment, there is also no possibility to specify ssh User Account Credentials to ssh into the compute cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
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
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.1.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = machinelearning.NewComputeCluster(ctx, "test", &machinelearning.ComputeClusterArgs{
// 			Location:                   pulumi.String("West Europe"),
// 			VmPriority:                 pulumi.String("LowPriority"),
// 			VmSize:                     pulumi.String("Standard_DS2_v2"),
// 			MachineLearningWorkspaceId: exampleWorkspace.ID(),
// 			SubnetResourceId:           exampleSubnet.ID(),
// 			ScaleSettings: &machinelearning.ComputeClusterScaleSettingsArgs{
// 				MinNodeCount:                    pulumi.Int(0),
// 				MaxNodeCount:                    pulumi.Int(1),
// 				ScaleDownNodesAfterIdleDuration: pulumi.String("PT30S"),
// 			},
// 			Identity: &machinelearning.ComputeClusterIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
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
// Machine Learning Compute Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:machinelearning/computeCluster:ComputeCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
// ```
type ComputeCluster struct {
	pulumi.CustomResourceState

	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity ComputeClusterIdentityPtrOutput `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled pulumi.BoolPtrOutput `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringOutput `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings ComputeClusterScaleSettingsOutput `pulumi:"scaleSettings"`
	// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Ssh ComputeClusterSshPtrOutput `pulumi:"ssh"`
	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SshPublicAccessEnabled pulumi.BoolOutput `pulumi:"sshPublicAccessEnabled"`
	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceId pulumi.StringPtrOutput `pulumi:"subnetResourceId"`
	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are `Dedicated` and `LowPriority`.
	VmPriority pulumi.StringOutput `pulumi:"vmPriority"`
	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
}

// NewComputeCluster registers a new resource with the given unique name, arguments, and options.
func NewComputeCluster(ctx *pulumi.Context,
	name string, args *ComputeClusterArgs, opts ...pulumi.ResourceOption) (*ComputeCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineLearningWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'MachineLearningWorkspaceId'")
	}
	if args.ScaleSettings == nil {
		return nil, errors.New("invalid value for required argument 'ScaleSettings'")
	}
	if args.VmPriority == nil {
		return nil, errors.New("invalid value for required argument 'VmPriority'")
	}
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	var resource ComputeCluster
	err := ctx.RegisterResource("azure:machinelearning/computeCluster:ComputeCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeCluster gets an existing ComputeCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterState, opts ...pulumi.ResourceOption) (*ComputeCluster, error) {
	var resource ComputeCluster
	err := ctx.ReadResource("azure:machinelearning/computeCluster:ComputeCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeCluster resources.
type computeClusterState struct {
	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description *string `pulumi:"description"`
	// An `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity *ComputeClusterIdentity `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceId *string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name *string `pulumi:"name"`
	// A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings *ComputeClusterScaleSettings `pulumi:"scaleSettings"`
	// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Ssh *ComputeClusterSsh `pulumi:"ssh"`
	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SshPublicAccessEnabled *bool `pulumi:"sshPublicAccessEnabled"`
	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are `Dedicated` and `LowPriority`.
	VmPriority *string `pulumi:"vmPriority"`
	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VmSize *string `pulumi:"vmSize"`
}

type ComputeClusterState struct {
	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description pulumi.StringPtrInput
	// An `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity ComputeClusterIdentityPtrInput
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled pulumi.BoolPtrInput
	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name pulumi.StringPtrInput
	// A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings ComputeClusterScaleSettingsPtrInput
	// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Ssh ComputeClusterSshPtrInput
	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SshPublicAccessEnabled pulumi.BoolPtrInput
	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Tags pulumi.StringMapInput
	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are `Dedicated` and `LowPriority`.
	VmPriority pulumi.StringPtrInput
	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VmSize pulumi.StringPtrInput
}

func (ComputeClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterState)(nil)).Elem()
}

type computeClusterArgs struct {
	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description *string `pulumi:"description"`
	// An `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity *ComputeClusterIdentity `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceId string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name *string `pulumi:"name"`
	// A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings ComputeClusterScaleSettings `pulumi:"scaleSettings"`
	// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Ssh *ComputeClusterSsh `pulumi:"ssh"`
	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SshPublicAccessEnabled *bool `pulumi:"sshPublicAccessEnabled"`
	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Tags map[string]string `pulumi:"tags"`
	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are `Dedicated` and `LowPriority`.
	VmPriority string `pulumi:"vmPriority"`
	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VmSize string `pulumi:"vmSize"`
}

// The set of arguments for constructing a ComputeCluster resource.
type ComputeClusterArgs struct {
	// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
	Description pulumi.StringPtrInput
	// An `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Identity ComputeClusterIdentityPtrInput
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Cluster to be created.
	LocalAuthEnabled pulumi.BoolPtrInput
	// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
	MachineLearningWorkspaceId pulumi.StringInput
	// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Name pulumi.StringPtrInput
	// A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	ScaleSettings ComputeClusterScaleSettingsInput
	// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
	Ssh ComputeClusterSshPtrInput
	// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
	SshPublicAccessEnabled pulumi.BoolPtrInput
	// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
	SubnetResourceId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
	Tags pulumi.StringMapInput
	// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created. Accepted values are `Dedicated` and `LowPriority`.
	VmPriority pulumi.StringInput
	// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
	VmSize pulumi.StringInput
}

func (ComputeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterArgs)(nil)).Elem()
}

type ComputeClusterInput interface {
	pulumi.Input

	ToComputeClusterOutput() ComputeClusterOutput
	ToComputeClusterOutputWithContext(ctx context.Context) ComputeClusterOutput
}

func (*ComputeCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeCluster)(nil))
}

func (i *ComputeCluster) ToComputeClusterOutput() ComputeClusterOutput {
	return i.ToComputeClusterOutputWithContext(context.Background())
}

func (i *ComputeCluster) ToComputeClusterOutputWithContext(ctx context.Context) ComputeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterOutput)
}

func (i *ComputeCluster) ToComputeClusterPtrOutput() ComputeClusterPtrOutput {
	return i.ToComputeClusterPtrOutputWithContext(context.Background())
}

func (i *ComputeCluster) ToComputeClusterPtrOutputWithContext(ctx context.Context) ComputeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterPtrOutput)
}

type ComputeClusterPtrInput interface {
	pulumi.Input

	ToComputeClusterPtrOutput() ComputeClusterPtrOutput
	ToComputeClusterPtrOutputWithContext(ctx context.Context) ComputeClusterPtrOutput
}

type computeClusterPtrType ComputeClusterArgs

func (*computeClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeCluster)(nil))
}

func (i *computeClusterPtrType) ToComputeClusterPtrOutput() ComputeClusterPtrOutput {
	return i.ToComputeClusterPtrOutputWithContext(context.Background())
}

func (i *computeClusterPtrType) ToComputeClusterPtrOutputWithContext(ctx context.Context) ComputeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterPtrOutput)
}

// ComputeClusterArrayInput is an input type that accepts ComputeClusterArray and ComputeClusterArrayOutput values.
// You can construct a concrete instance of `ComputeClusterArrayInput` via:
//
//          ComputeClusterArray{ ComputeClusterArgs{...} }
type ComputeClusterArrayInput interface {
	pulumi.Input

	ToComputeClusterArrayOutput() ComputeClusterArrayOutput
	ToComputeClusterArrayOutputWithContext(context.Context) ComputeClusterArrayOutput
}

type ComputeClusterArray []ComputeClusterInput

func (ComputeClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeCluster)(nil)).Elem()
}

func (i ComputeClusterArray) ToComputeClusterArrayOutput() ComputeClusterArrayOutput {
	return i.ToComputeClusterArrayOutputWithContext(context.Background())
}

func (i ComputeClusterArray) ToComputeClusterArrayOutputWithContext(ctx context.Context) ComputeClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterArrayOutput)
}

// ComputeClusterMapInput is an input type that accepts ComputeClusterMap and ComputeClusterMapOutput values.
// You can construct a concrete instance of `ComputeClusterMapInput` via:
//
//          ComputeClusterMap{ "key": ComputeClusterArgs{...} }
type ComputeClusterMapInput interface {
	pulumi.Input

	ToComputeClusterMapOutput() ComputeClusterMapOutput
	ToComputeClusterMapOutputWithContext(context.Context) ComputeClusterMapOutput
}

type ComputeClusterMap map[string]ComputeClusterInput

func (ComputeClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeCluster)(nil)).Elem()
}

func (i ComputeClusterMap) ToComputeClusterMapOutput() ComputeClusterMapOutput {
	return i.ToComputeClusterMapOutputWithContext(context.Background())
}

func (i ComputeClusterMap) ToComputeClusterMapOutputWithContext(ctx context.Context) ComputeClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterMapOutput)
}

type ComputeClusterOutput struct{ *pulumi.OutputState }

func (ComputeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeCluster)(nil))
}

func (o ComputeClusterOutput) ToComputeClusterOutput() ComputeClusterOutput {
	return o
}

func (o ComputeClusterOutput) ToComputeClusterOutputWithContext(ctx context.Context) ComputeClusterOutput {
	return o
}

func (o ComputeClusterOutput) ToComputeClusterPtrOutput() ComputeClusterPtrOutput {
	return o.ToComputeClusterPtrOutputWithContext(context.Background())
}

func (o ComputeClusterOutput) ToComputeClusterPtrOutputWithContext(ctx context.Context) ComputeClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeCluster) *ComputeCluster {
		return &v
	}).(ComputeClusterPtrOutput)
}

type ComputeClusterPtrOutput struct{ *pulumi.OutputState }

func (ComputeClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeCluster)(nil))
}

func (o ComputeClusterPtrOutput) ToComputeClusterPtrOutput() ComputeClusterPtrOutput {
	return o
}

func (o ComputeClusterPtrOutput) ToComputeClusterPtrOutputWithContext(ctx context.Context) ComputeClusterPtrOutput {
	return o
}

func (o ComputeClusterPtrOutput) Elem() ComputeClusterOutput {
	return o.ApplyT(func(v *ComputeCluster) ComputeCluster {
		if v != nil {
			return *v
		}
		var ret ComputeCluster
		return ret
	}).(ComputeClusterOutput)
}

type ComputeClusterArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeCluster)(nil))
}

func (o ComputeClusterArrayOutput) ToComputeClusterArrayOutput() ComputeClusterArrayOutput {
	return o
}

func (o ComputeClusterArrayOutput) ToComputeClusterArrayOutputWithContext(ctx context.Context) ComputeClusterArrayOutput {
	return o
}

func (o ComputeClusterArrayOutput) Index(i pulumi.IntInput) ComputeClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeCluster {
		return vs[0].([]ComputeCluster)[vs[1].(int)]
	}).(ComputeClusterOutput)
}

type ComputeClusterMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ComputeCluster)(nil))
}

func (o ComputeClusterMapOutput) ToComputeClusterMapOutput() ComputeClusterMapOutput {
	return o
}

func (o ComputeClusterMapOutput) ToComputeClusterMapOutputWithContext(ctx context.Context) ComputeClusterMapOutput {
	return o
}

func (o ComputeClusterMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ComputeCluster {
		return vs[0].(map[string]ComputeCluster)[vs[1].(string)]
	}).(ComputeClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeClusterOutput{})
	pulumi.RegisterOutputType(ComputeClusterPtrOutput{})
	pulumi.RegisterOutputType(ComputeClusterArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterMapOutput{})
}
