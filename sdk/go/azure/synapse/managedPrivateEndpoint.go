// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Managed Private Endpoint.
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
// 			ManagedVirtualNetworkEnabled:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFirewallRule, err := synapse.NewFirewallRule(ctx, "exampleFirewallRule", &synapse.FirewallRuleArgs{
// 			SynapseWorkspaceId: pulumi.Any(azurerm_synapse_workspace.Test.Id),
// 			StartIpAddress:     pulumi.String("0.0.0.0"),
// 			EndIpAddress:       pulumi.String("255.255.255.255"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConnect, err := storage.NewAccount(ctx, "exampleConnect", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("BlobStorage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewManagedPrivateEndpoint(ctx, "exampleManagedPrivateEndpoint", &synapse.ManagedPrivateEndpointArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			TargetResourceId:   exampleConnect.ID(),
// 			SubresourceName:    pulumi.String("blob"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleFirewallRule,
// 		}))
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
// Synapse Managed Private Endpoint can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
// ```
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringOutput `pulumi:"subresourceName"`
	// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
	// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubresourceName == nil {
		return nil, errors.New("invalid value for required argument 'SubresourceName'")
	}
	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `pulumi:"subresourceName"`
	// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
	// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type ManagedPrivateEndpointState struct {
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringPtrInput
	// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
	// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName string `pulumi:"subresourceName"`
	// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
	// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringInput
	// The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringInput
	// The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

// ManagedPrivateEndpointArrayInput is an input type that accepts ManagedPrivateEndpointArray and ManagedPrivateEndpointArrayOutput values.
// You can construct a concrete instance of `ManagedPrivateEndpointArrayInput` via:
//
//          ManagedPrivateEndpointArray{ ManagedPrivateEndpointArgs{...} }
type ManagedPrivateEndpointArrayInput interface {
	pulumi.Input

	ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput
	ToManagedPrivateEndpointArrayOutputWithContext(context.Context) ManagedPrivateEndpointArrayOutput
}

type ManagedPrivateEndpointArray []ManagedPrivateEndpointInput

func (ManagedPrivateEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (i ManagedPrivateEndpointArray) ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput {
	return i.ToManagedPrivateEndpointArrayOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointArray) ToManagedPrivateEndpointArrayOutputWithContext(ctx context.Context) ManagedPrivateEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointArrayOutput)
}

// ManagedPrivateEndpointMapInput is an input type that accepts ManagedPrivateEndpointMap and ManagedPrivateEndpointMapOutput values.
// You can construct a concrete instance of `ManagedPrivateEndpointMapInput` via:
//
//          ManagedPrivateEndpointMap{ "key": ManagedPrivateEndpointArgs{...} }
type ManagedPrivateEndpointMapInput interface {
	pulumi.Input

	ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput
	ToManagedPrivateEndpointMapOutputWithContext(context.Context) ManagedPrivateEndpointMapOutput
}

type ManagedPrivateEndpointMap map[string]ManagedPrivateEndpointInput

func (ManagedPrivateEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (i ManagedPrivateEndpointMap) ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput {
	return i.ToManagedPrivateEndpointMapOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointMap) ToManagedPrivateEndpointMapOutputWithContext(ctx context.Context) ManagedPrivateEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointMapOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

type ManagedPrivateEndpointArrayOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointArrayOutput) ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput {
	return o
}

func (o ManagedPrivateEndpointArrayOutput) ToManagedPrivateEndpointArrayOutputWithContext(ctx context.Context) ManagedPrivateEndpointArrayOutput {
	return o
}

func (o ManagedPrivateEndpointArrayOutput) Index(i pulumi.IntInput) ManagedPrivateEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedPrivateEndpoint {
		return vs[0].([]*ManagedPrivateEndpoint)[vs[1].(int)]
	}).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointMapOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointMapOutput) ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput {
	return o
}

func (o ManagedPrivateEndpointMapOutput) ToManagedPrivateEndpointMapOutputWithContext(ctx context.Context) ManagedPrivateEndpointMapOutput {
	return o
}

func (o ManagedPrivateEndpointMapOutput) MapIndex(k pulumi.StringInput) ManagedPrivateEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedPrivateEndpoint {
		return vs[0].(map[string]*ManagedPrivateEndpoint)[vs[1].(string)]
	}).(ManagedPrivateEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointInput)(nil)).Elem(), &ManagedPrivateEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointArrayInput)(nil)).Elem(), ManagedPrivateEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointMapInput)(nil)).Elem(), ManagedPrivateEndpointMap{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointArrayOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointMapOutput{})
}
