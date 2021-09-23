// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages the association between a Network Interface and a Load Balancer's NAT Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/lb"
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
// 				pulumi.String("10.0.0.0/16"),
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
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              pulumi.String("primary"),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNatRule, err := lb.NewNatRule(ctx, "exampleNatRule", &lb.NatRuleArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			LoadbalancerId:              exampleLoadBalancer.ID(),
// 			Protocol:                    pulumi.String("Tcp"),
// 			FrontendPort:                pulumi.Int(3389),
// 			BackendPort:                 pulumi.Int(3389),
// 			FrontendIpConfigurationName: pulumi.String("primary"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("testconfiguration1"),
// 					SubnetId:                   exampleSubnet.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkInterfaceNatRuleAssociation(ctx, "exampleNetworkInterfaceNatRuleAssociation", &network.NetworkInterfaceNatRuleAssociationArgs{
// 			NetworkInterfaceId:  exampleNetworkInterface.ID(),
// 			IpConfigurationName: pulumi.String("testconfiguration1"),
// 			NatRuleId:           exampleNatRule.ID(),
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
// Associations between Network Interfaces and Load Balancer NAT Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
// ```
type NetworkInterfaceNatRuleAssociation struct {
	pulumi.CustomResourceState

	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringOutput `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringOutput `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceNatRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceNatRuleAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceNatRuleAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceNatRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'IpConfigurationName'")
	}
	if args.NatRuleId == nil {
		return nil, errors.New("invalid value for required argument 'NatRuleId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceNatRuleAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceNatRuleAssociation gets an existing NetworkInterfaceNatRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceNatRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceNatRuleAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceNatRuleAssociation, error) {
	var resource NetworkInterfaceNatRuleAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceNatRuleAssociation resources.
type networkInterfaceNatRuleAssociationState struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName *string `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId *string `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceNatRuleAssociationState struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringPtrInput
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceNatRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceNatRuleAssociationState)(nil)).Elem()
}

type networkInterfaceNatRuleAssociationArgs struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName string `pulumi:"ipConfigurationName"`
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId string `pulumi:"natRuleId"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceNatRuleAssociation resource.
type NetworkInterfaceNatRuleAssociationArgs struct {
	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringInput
	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NatRuleId pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceNatRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceNatRuleAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceNatRuleAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput
	ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput
}

func (*NetworkInterfaceNatRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNatRuleAssociation)(nil))
}

func (i *NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput {
	return i.ToNetworkInterfaceNatRuleAssociationOutputWithContext(context.Background())
}

func (i *NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationOutput)
}

func (i *NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationPtrOutput() NetworkInterfaceNatRuleAssociationPtrOutput {
	return i.ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(context.Background())
}

func (i *NetworkInterfaceNatRuleAssociation) ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationPtrOutput)
}

type NetworkInterfaceNatRuleAssociationPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceNatRuleAssociationPtrOutput() NetworkInterfaceNatRuleAssociationPtrOutput
	ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationPtrOutput
}

type networkInterfaceNatRuleAssociationPtrType NetworkInterfaceNatRuleAssociationArgs

func (*networkInterfaceNatRuleAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceNatRuleAssociation)(nil))
}

func (i *networkInterfaceNatRuleAssociationPtrType) ToNetworkInterfaceNatRuleAssociationPtrOutput() NetworkInterfaceNatRuleAssociationPtrOutput {
	return i.ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceNatRuleAssociationPtrType) ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationPtrOutput)
}

// NetworkInterfaceNatRuleAssociationArrayInput is an input type that accepts NetworkInterfaceNatRuleAssociationArray and NetworkInterfaceNatRuleAssociationArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceNatRuleAssociationArrayInput` via:
//
//          NetworkInterfaceNatRuleAssociationArray{ NetworkInterfaceNatRuleAssociationArgs{...} }
type NetworkInterfaceNatRuleAssociationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceNatRuleAssociationArrayOutput() NetworkInterfaceNatRuleAssociationArrayOutput
	ToNetworkInterfaceNatRuleAssociationArrayOutputWithContext(context.Context) NetworkInterfaceNatRuleAssociationArrayOutput
}

type NetworkInterfaceNatRuleAssociationArray []NetworkInterfaceNatRuleAssociationInput

func (NetworkInterfaceNatRuleAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterfaceNatRuleAssociation)(nil)).Elem()
}

func (i NetworkInterfaceNatRuleAssociationArray) ToNetworkInterfaceNatRuleAssociationArrayOutput() NetworkInterfaceNatRuleAssociationArrayOutput {
	return i.ToNetworkInterfaceNatRuleAssociationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceNatRuleAssociationArray) ToNetworkInterfaceNatRuleAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationArrayOutput)
}

// NetworkInterfaceNatRuleAssociationMapInput is an input type that accepts NetworkInterfaceNatRuleAssociationMap and NetworkInterfaceNatRuleAssociationMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceNatRuleAssociationMapInput` via:
//
//          NetworkInterfaceNatRuleAssociationMap{ "key": NetworkInterfaceNatRuleAssociationArgs{...} }
type NetworkInterfaceNatRuleAssociationMapInput interface {
	pulumi.Input

	ToNetworkInterfaceNatRuleAssociationMapOutput() NetworkInterfaceNatRuleAssociationMapOutput
	ToNetworkInterfaceNatRuleAssociationMapOutputWithContext(context.Context) NetworkInterfaceNatRuleAssociationMapOutput
}

type NetworkInterfaceNatRuleAssociationMap map[string]NetworkInterfaceNatRuleAssociationInput

func (NetworkInterfaceNatRuleAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterfaceNatRuleAssociation)(nil)).Elem()
}

func (i NetworkInterfaceNatRuleAssociationMap) ToNetworkInterfaceNatRuleAssociationMapOutput() NetworkInterfaceNatRuleAssociationMapOutput {
	return i.ToNetworkInterfaceNatRuleAssociationMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceNatRuleAssociationMap) ToNetworkInterfaceNatRuleAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceNatRuleAssociationMapOutput)
}

type NetworkInterfaceNatRuleAssociationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNatRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNatRuleAssociation)(nil))
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationOutput() NetworkInterfaceNatRuleAssociationOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationPtrOutput() NetworkInterfaceNatRuleAssociationPtrOutput {
	return o.ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceNatRuleAssociationOutput) ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceNatRuleAssociation) *NetworkInterfaceNatRuleAssociation {
		return &v
	}).(NetworkInterfaceNatRuleAssociationPtrOutput)
}

type NetworkInterfaceNatRuleAssociationPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNatRuleAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceNatRuleAssociation)(nil))
}

func (o NetworkInterfaceNatRuleAssociationPtrOutput) ToNetworkInterfaceNatRuleAssociationPtrOutput() NetworkInterfaceNatRuleAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationPtrOutput) ToNetworkInterfaceNatRuleAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationPtrOutput) Elem() NetworkInterfaceNatRuleAssociationOutput {
	return o.ApplyT(func(v *NetworkInterfaceNatRuleAssociation) NetworkInterfaceNatRuleAssociation {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceNatRuleAssociation
		return ret
	}).(NetworkInterfaceNatRuleAssociationOutput)
}

type NetworkInterfaceNatRuleAssociationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNatRuleAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceNatRuleAssociation)(nil))
}

func (o NetworkInterfaceNatRuleAssociationArrayOutput) ToNetworkInterfaceNatRuleAssociationArrayOutput() NetworkInterfaceNatRuleAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationArrayOutput) ToNetworkInterfaceNatRuleAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceNatRuleAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceNatRuleAssociation {
		return vs[0].([]NetworkInterfaceNatRuleAssociation)[vs[1].(int)]
	}).(NetworkInterfaceNatRuleAssociationOutput)
}

type NetworkInterfaceNatRuleAssociationMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNatRuleAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkInterfaceNatRuleAssociation)(nil))
}

func (o NetworkInterfaceNatRuleAssociationMapOutput) ToNetworkInterfaceNatRuleAssociationMapOutput() NetworkInterfaceNatRuleAssociationMapOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationMapOutput) ToNetworkInterfaceNatRuleAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceNatRuleAssociationMapOutput {
	return o
}

func (o NetworkInterfaceNatRuleAssociationMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceNatRuleAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkInterfaceNatRuleAssociation {
		return vs[0].(map[string]NetworkInterfaceNatRuleAssociation)[vs[1].(string)]
	}).(NetworkInterfaceNatRuleAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceNatRuleAssociationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceNatRuleAssociationPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceNatRuleAssociationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceNatRuleAssociationMapOutput{})
}
