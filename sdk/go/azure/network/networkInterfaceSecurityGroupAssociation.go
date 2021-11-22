// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages the association between a Network Interface and a Network Security Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// 		_, err = network.NewNetworkInterfaceSecurityGroupAssociation(ctx, "exampleNetworkInterfaceSecurityGroupAssociation", &network.NetworkInterfaceSecurityGroupAssociationArgs{
// 			NetworkInterfaceId:     exampleNetworkInterface.ID(),
// 			NetworkSecurityGroupId: exampleNetworkSecurityGroup.ID(),
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
// Associations between Network Interfaces and Network Security Group can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkSecurityGroups/group1"
// ```
type NetworkInterfaceSecurityGroupAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringOutput `pulumi:"networkSecurityGroupId"`
}

// NewNetworkInterfaceSecurityGroupAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceSecurityGroupAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceSecurityGroupAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.NetworkSecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupId'")
	}
	var resource NetworkInterfaceSecurityGroupAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceSecurityGroupAssociation gets an existing NetworkInterfaceSecurityGroupAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceSecurityGroupAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceSecurityGroupAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAssociation, error) {
	var resource NetworkInterfaceSecurityGroupAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAssociation resources.
type networkInterfaceSecurityGroupAssociationState struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId *string `pulumi:"networkSecurityGroupId"`
}

type NetworkInterfaceSecurityGroupAssociationState struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringPtrInput
}

func (NetworkInterfaceSecurityGroupAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAssociationState)(nil)).Elem()
}

type networkInterfaceSecurityGroupAssociationArgs struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
}

// The set of arguments for constructing a NetworkInterfaceSecurityGroupAssociation resource.
type NetworkInterfaceSecurityGroupAssociationArgs struct {
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
	// The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringInput
}

func (NetworkInterfaceSecurityGroupAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceSecurityGroupAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput
	ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput
}

func (*NetworkInterfaceSecurityGroupAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (i *NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(context.Background())
}

func (i *NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationOutput)
}

func (i *NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationPtrOutput() NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(context.Background())
}

func (i *NetworkInterfaceSecurityGroupAssociation) ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationPtrOutput)
}

type NetworkInterfaceSecurityGroupAssociationPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAssociationPtrOutput() NetworkInterfaceSecurityGroupAssociationPtrOutput
	ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationPtrOutput
}

type networkInterfaceSecurityGroupAssociationPtrType NetworkInterfaceSecurityGroupAssociationArgs

func (*networkInterfaceSecurityGroupAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (i *networkInterfaceSecurityGroupAssociationPtrType) ToNetworkInterfaceSecurityGroupAssociationPtrOutput() NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceSecurityGroupAssociationPtrType) ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationPtrOutput)
}

// NetworkInterfaceSecurityGroupAssociationArrayInput is an input type that accepts NetworkInterfaceSecurityGroupAssociationArray and NetworkInterfaceSecurityGroupAssociationArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceSecurityGroupAssociationArrayInput` via:
//
//          NetworkInterfaceSecurityGroupAssociationArray{ NetworkInterfaceSecurityGroupAssociationArgs{...} }
type NetworkInterfaceSecurityGroupAssociationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAssociationArrayOutput() NetworkInterfaceSecurityGroupAssociationArrayOutput
	ToNetworkInterfaceSecurityGroupAssociationArrayOutputWithContext(context.Context) NetworkInterfaceSecurityGroupAssociationArrayOutput
}

type NetworkInterfaceSecurityGroupAssociationArray []NetworkInterfaceSecurityGroupAssociationInput

func (NetworkInterfaceSecurityGroupAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterfaceSecurityGroupAssociation)(nil)).Elem()
}

func (i NetworkInterfaceSecurityGroupAssociationArray) ToNetworkInterfaceSecurityGroupAssociationArrayOutput() NetworkInterfaceSecurityGroupAssociationArrayOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceSecurityGroupAssociationArray) ToNetworkInterfaceSecurityGroupAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationArrayOutput)
}

// NetworkInterfaceSecurityGroupAssociationMapInput is an input type that accepts NetworkInterfaceSecurityGroupAssociationMap and NetworkInterfaceSecurityGroupAssociationMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceSecurityGroupAssociationMapInput` via:
//
//          NetworkInterfaceSecurityGroupAssociationMap{ "key": NetworkInterfaceSecurityGroupAssociationArgs{...} }
type NetworkInterfaceSecurityGroupAssociationMapInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAssociationMapOutput() NetworkInterfaceSecurityGroupAssociationMapOutput
	ToNetworkInterfaceSecurityGroupAssociationMapOutputWithContext(context.Context) NetworkInterfaceSecurityGroupAssociationMapOutput
}

type NetworkInterfaceSecurityGroupAssociationMap map[string]NetworkInterfaceSecurityGroupAssociationInput

func (NetworkInterfaceSecurityGroupAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterfaceSecurityGroupAssociation)(nil)).Elem()
}

func (i NetworkInterfaceSecurityGroupAssociationMap) ToNetworkInterfaceSecurityGroupAssociationMapOutput() NetworkInterfaceSecurityGroupAssociationMapOutput {
	return i.ToNetworkInterfaceSecurityGroupAssociationMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceSecurityGroupAssociationMap) ToNetworkInterfaceSecurityGroupAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAssociationMapOutput)
}

type NetworkInterfaceSecurityGroupAssociationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationOutput() NetworkInterfaceSecurityGroupAssociationOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationPtrOutput() NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return o.ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceSecurityGroupAssociationOutput) ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceSecurityGroupAssociation) *NetworkInterfaceSecurityGroupAssociation {
		return &v
	}).(NetworkInterfaceSecurityGroupAssociationPtrOutput)
}

type NetworkInterfaceSecurityGroupAssociationPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceSecurityGroupAssociationPtrOutput) ToNetworkInterfaceSecurityGroupAssociationPtrOutput() NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationPtrOutput) ToNetworkInterfaceSecurityGroupAssociationPtrOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationPtrOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationPtrOutput) Elem() NetworkInterfaceSecurityGroupAssociationOutput {
	return o.ApplyT(func(v *NetworkInterfaceSecurityGroupAssociation) NetworkInterfaceSecurityGroupAssociation {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceSecurityGroupAssociation
		return ret
	}).(NetworkInterfaceSecurityGroupAssociationOutput)
}

type NetworkInterfaceSecurityGroupAssociationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceSecurityGroupAssociationArrayOutput) ToNetworkInterfaceSecurityGroupAssociationArrayOutput() NetworkInterfaceSecurityGroupAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationArrayOutput) ToNetworkInterfaceSecurityGroupAssociationArrayOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationArrayOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceSecurityGroupAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceSecurityGroupAssociation {
		return vs[0].([]NetworkInterfaceSecurityGroupAssociation)[vs[1].(int)]
	}).(NetworkInterfaceSecurityGroupAssociationOutput)
}

type NetworkInterfaceSecurityGroupAssociationMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkInterfaceSecurityGroupAssociation)(nil))
}

func (o NetworkInterfaceSecurityGroupAssociationMapOutput) ToNetworkInterfaceSecurityGroupAssociationMapOutput() NetworkInterfaceSecurityGroupAssociationMapOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationMapOutput) ToNetworkInterfaceSecurityGroupAssociationMapOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAssociationMapOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAssociationMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceSecurityGroupAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkInterfaceSecurityGroupAssociation {
		return vs[0].(map[string]NetworkInterfaceSecurityGroupAssociation)[vs[1].(string)]
	}).(NetworkInterfaceSecurityGroupAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociationInput)(nil)).Elem(), &NetworkInterfaceSecurityGroupAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociationPtrInput)(nil)).Elem(), &NetworkInterfaceSecurityGroupAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociationArrayInput)(nil)).Elem(), NetworkInterfaceSecurityGroupAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAssociationMapInput)(nil)).Elem(), NetworkInterfaceSecurityGroupAssociationMap{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAssociationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAssociationPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAssociationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAssociationMapOutput{})
}
