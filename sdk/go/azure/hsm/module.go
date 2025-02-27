// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hsm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Dedicated Hardware Security Module.
//
// > **Note:** Before using this resource, it's required to submit the request of registering the providers and features with Azure CLI `az provider register --namespace Microsoft.HardwareSecurityModules && az feature register --namespace Microsoft.HardwareSecurityModules --name AzureDedicatedHSM && az provider register --namespace Microsoft.Network && az feature register --namespace Microsoft.Network --name AllowBaremetalServers` and ask service team (hsmrequest@microsoft.com) to approve. See more details from https://docs.microsoft.com/en-us/azure/dedicated-hsm/tutorial-deploy-hsm-cli#prerequisites.
//
// > **Note:** If the quota is not enough in some region, please submit the quota request to service team.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/hsm"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.2.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.2.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example2, err := network.NewSubnet(ctx, "example2", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.2.1.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("first"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.HardwareSecurityModules/dedicatedHSMs"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/networkinterfaces/*"),
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example3, err := network.NewSubnet(ctx, "example3", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.2.255.0/26"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetworkGateway, err := network.NewVirtualNetworkGateway(ctx, "exampleVirtualNetworkGateway", &network.VirtualNetworkGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Type:              pulumi.String("ExpressRoute"),
// 			VpnType:           pulumi.String("PolicyBased"),
// 			Sku:               pulumi.String("Standard"),
// 			IpConfigurations: network.VirtualNetworkGatewayIpConfigurationArray{
// 				&network.VirtualNetworkGatewayIpConfigurationArgs{
// 					PublicIpAddressId:          examplePublicIp.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 					SubnetId:                   example3.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hsm.NewModule(ctx, "exampleModule", &hsm.ModuleArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("SafeNet Luna Network HSM A790"),
// 			NetworkProfile: &hsm.ModuleNetworkProfileArgs{
// 				NetworkInterfacePrivateIpAddresses: pulumi.StringArray{
// 					pulumi.String("10.2.1.8"),
// 				},
// 				SubnetId: example2.ID(),
// 			},
// 			StampId: pulumi.String("stamp2"),
// 			Tags: pulumi.StringMap{
// 				"env": pulumi.String("Test"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleVirtualNetworkGateway,
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
// Dedicated Hardware Security Module can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:hsm/module:Module example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
// ```
type Module struct {
	pulumi.CustomResourceState

	// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile ModuleNetworkProfileOutput `pulumi:"networkProfile"`
	// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
	StampId pulumi.StringPtrOutput `pulumi:"stampId"`
	// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewModule registers a new resource with the given unique name, arguments, and options.
func NewModule(ctx *pulumi.Context,
	name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*Module, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkProfile == nil {
		return nil, errors.New("invalid value for required argument 'NetworkProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource Module
	err := ctx.RegisterResource("azure:hsm/module:Module", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModule gets an existing Module resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleState, opts ...pulumi.ResourceOption) (*Module, error) {
	var resource Module
	err := ctx.ReadResource("azure:hsm/module:Module", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Module resources.
type moduleState struct {
	// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *ModuleNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
	SkuName *string `pulumi:"skuName"`
	// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
	StampId *string `pulumi:"stampId"`
	// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
	Tags map[string]string `pulumi:"tags"`
	// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
	Zones []string `pulumi:"zones"`
}

type ModuleState struct {
	// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile ModuleNetworkProfilePtrInput
	// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
	SkuName pulumi.StringPtrInput
	// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
	StampId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
	Tags pulumi.StringMapInput
	// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
	Zones pulumi.StringArrayInput
}

func (ModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleState)(nil)).Elem()
}

type moduleArgs struct {
	// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile ModuleNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
	SkuName string `pulumi:"skuName"`
	// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
	StampId *string `pulumi:"stampId"`
	// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
	Tags map[string]string `pulumi:"tags"`
	// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Module resource.
type ModuleArgs struct {
	// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile ModuleNetworkProfileInput
	// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
	ResourceGroupName pulumi.StringInput
	// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
	SkuName pulumi.StringInput
	// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
	StampId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
	Tags pulumi.StringMapInput
	// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
	Zones pulumi.StringArrayInput
}

func (ModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleArgs)(nil)).Elem()
}

type ModuleInput interface {
	pulumi.Input

	ToModuleOutput() ModuleOutput
	ToModuleOutputWithContext(ctx context.Context) ModuleOutput
}

func (*Module) ElementType() reflect.Type {
	return reflect.TypeOf((**Module)(nil)).Elem()
}

func (i *Module) ToModuleOutput() ModuleOutput {
	return i.ToModuleOutputWithContext(context.Background())
}

func (i *Module) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleOutput)
}

// ModuleArrayInput is an input type that accepts ModuleArray and ModuleArrayOutput values.
// You can construct a concrete instance of `ModuleArrayInput` via:
//
//          ModuleArray{ ModuleArgs{...} }
type ModuleArrayInput interface {
	pulumi.Input

	ToModuleArrayOutput() ModuleArrayOutput
	ToModuleArrayOutputWithContext(context.Context) ModuleArrayOutput
}

type ModuleArray []ModuleInput

func (ModuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Module)(nil)).Elem()
}

func (i ModuleArray) ToModuleArrayOutput() ModuleArrayOutput {
	return i.ToModuleArrayOutputWithContext(context.Background())
}

func (i ModuleArray) ToModuleArrayOutputWithContext(ctx context.Context) ModuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleArrayOutput)
}

// ModuleMapInput is an input type that accepts ModuleMap and ModuleMapOutput values.
// You can construct a concrete instance of `ModuleMapInput` via:
//
//          ModuleMap{ "key": ModuleArgs{...} }
type ModuleMapInput interface {
	pulumi.Input

	ToModuleMapOutput() ModuleMapOutput
	ToModuleMapOutputWithContext(context.Context) ModuleMapOutput
}

type ModuleMap map[string]ModuleInput

func (ModuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Module)(nil)).Elem()
}

func (i ModuleMap) ToModuleMapOutput() ModuleMapOutput {
	return i.ToModuleMapOutputWithContext(context.Background())
}

func (i ModuleMap) ToModuleMapOutputWithContext(ctx context.Context) ModuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleMapOutput)
}

type ModuleOutput struct{ *pulumi.OutputState }

func (ModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Module)(nil)).Elem()
}

func (o ModuleOutput) ToModuleOutput() ModuleOutput {
	return o
}

func (o ModuleOutput) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return o
}

type ModuleArrayOutput struct{ *pulumi.OutputState }

func (ModuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Module)(nil)).Elem()
}

func (o ModuleArrayOutput) ToModuleArrayOutput() ModuleArrayOutput {
	return o
}

func (o ModuleArrayOutput) ToModuleArrayOutputWithContext(ctx context.Context) ModuleArrayOutput {
	return o
}

func (o ModuleArrayOutput) Index(i pulumi.IntInput) ModuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Module {
		return vs[0].([]*Module)[vs[1].(int)]
	}).(ModuleOutput)
}

type ModuleMapOutput struct{ *pulumi.OutputState }

func (ModuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Module)(nil)).Elem()
}

func (o ModuleMapOutput) ToModuleMapOutput() ModuleMapOutput {
	return o
}

func (o ModuleMapOutput) ToModuleMapOutputWithContext(ctx context.Context) ModuleMapOutput {
	return o
}

func (o ModuleMapOutput) MapIndex(k pulumi.StringInput) ModuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Module {
		return vs[0].(map[string]*Module)[vs[1].(string)]
	}).(ModuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleInput)(nil)).Elem(), &Module{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleArrayInput)(nil)).Elem(), ModuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleMapInput)(nil)).Elem(), ModuleMap{})
	pulumi.RegisterOutputType(ModuleOutput{})
	pulumi.RegisterOutputType(ModuleArrayOutput{})
	pulumi.RegisterOutputType(ModuleMapOutput{})
}
