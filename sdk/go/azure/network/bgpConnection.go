// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Bgp Connection for a Virtual Hub.
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
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku:               pulumi.String("Standard"),
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.5.0.0/16"),
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
// 			AddressPrefix:      pulumi.String("10.5.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHubIp, err := network.NewVirtualHubIp(ctx, "exampleVirtualHubIp", &network.VirtualHubIpArgs{
// 			VirtualHubId:              exampleVirtualHub.ID(),
// 			PrivateIpAddress:          pulumi.String("10.5.1.18"),
// 			PrivateIpAllocationMethod: pulumi.String("Static"),
// 			PublicIpAddressId:         examplePublicIp.ID(),
// 			SubnetId:                  exampleSubnet.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewBgpConnection(ctx, "exampleBgpConnection", &network.BgpConnectionArgs{
// 			VirtualHubId: exampleVirtualHub.ID(),
// 			PeerAsn:      pulumi.Int(65514),
// 			PeerIp:       pulumi.String("169.254.21.5"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleVirtualHubIp,
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
// Virtual Hub Bgp Connections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/bgpConnection:BgpConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/connection1
// ```
type BgpConnection struct {
	pulumi.CustomResourceState

	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntOutput `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewBgpConnection registers a new resource with the given unique name, arguments, and options.
func NewBgpConnection(ctx *pulumi.Context,
	name string, args *BgpConnectionArgs, opts ...pulumi.ResourceOption) (*BgpConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerAsn == nil {
		return nil, errors.New("invalid value for required argument 'PeerAsn'")
	}
	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource BgpConnection
	err := ctx.RegisterResource("azure:network/bgpConnection:BgpConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpConnection gets an existing BgpConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpConnectionState, opts ...pulumi.ResourceOption) (*BgpConnection, error) {
	var resource BgpConnection
	err := ctx.ReadResource("azure:network/bgpConnection:BgpConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpConnection resources.
type bgpConnectionState struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn *int `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp *string `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type BgpConnectionState struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntPtrInput
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringPtrInput
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
}

func (BgpConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpConnectionState)(nil)).Elem()
}

type bgpConnectionArgs struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn int `pulumi:"peerAsn"`
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp string `pulumi:"peerIp"`
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a BgpConnection resource.
type BgpConnectionArgs struct {
	// The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn pulumi.IntInput
	// The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
	PeerIp pulumi.StringInput
	// The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
}

func (BgpConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpConnectionArgs)(nil)).Elem()
}

type BgpConnectionInput interface {
	pulumi.Input

	ToBgpConnectionOutput() BgpConnectionOutput
	ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput
}

func (*BgpConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpConnection)(nil))
}

func (i *BgpConnection) ToBgpConnectionOutput() BgpConnectionOutput {
	return i.ToBgpConnectionOutputWithContext(context.Background())
}

func (i *BgpConnection) ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionOutput)
}

func (i *BgpConnection) ToBgpConnectionPtrOutput() BgpConnectionPtrOutput {
	return i.ToBgpConnectionPtrOutputWithContext(context.Background())
}

func (i *BgpConnection) ToBgpConnectionPtrOutputWithContext(ctx context.Context) BgpConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionPtrOutput)
}

type BgpConnectionPtrInput interface {
	pulumi.Input

	ToBgpConnectionPtrOutput() BgpConnectionPtrOutput
	ToBgpConnectionPtrOutputWithContext(ctx context.Context) BgpConnectionPtrOutput
}

type bgpConnectionPtrType BgpConnectionArgs

func (*bgpConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpConnection)(nil))
}

func (i *bgpConnectionPtrType) ToBgpConnectionPtrOutput() BgpConnectionPtrOutput {
	return i.ToBgpConnectionPtrOutputWithContext(context.Background())
}

func (i *bgpConnectionPtrType) ToBgpConnectionPtrOutputWithContext(ctx context.Context) BgpConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionPtrOutput)
}

// BgpConnectionArrayInput is an input type that accepts BgpConnectionArray and BgpConnectionArrayOutput values.
// You can construct a concrete instance of `BgpConnectionArrayInput` via:
//
//          BgpConnectionArray{ BgpConnectionArgs{...} }
type BgpConnectionArrayInput interface {
	pulumi.Input

	ToBgpConnectionArrayOutput() BgpConnectionArrayOutput
	ToBgpConnectionArrayOutputWithContext(context.Context) BgpConnectionArrayOutput
}

type BgpConnectionArray []BgpConnectionInput

func (BgpConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpConnection)(nil)).Elem()
}

func (i BgpConnectionArray) ToBgpConnectionArrayOutput() BgpConnectionArrayOutput {
	return i.ToBgpConnectionArrayOutputWithContext(context.Background())
}

func (i BgpConnectionArray) ToBgpConnectionArrayOutputWithContext(ctx context.Context) BgpConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionArrayOutput)
}

// BgpConnectionMapInput is an input type that accepts BgpConnectionMap and BgpConnectionMapOutput values.
// You can construct a concrete instance of `BgpConnectionMapInput` via:
//
//          BgpConnectionMap{ "key": BgpConnectionArgs{...} }
type BgpConnectionMapInput interface {
	pulumi.Input

	ToBgpConnectionMapOutput() BgpConnectionMapOutput
	ToBgpConnectionMapOutputWithContext(context.Context) BgpConnectionMapOutput
}

type BgpConnectionMap map[string]BgpConnectionInput

func (BgpConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpConnection)(nil)).Elem()
}

func (i BgpConnectionMap) ToBgpConnectionMapOutput() BgpConnectionMapOutput {
	return i.ToBgpConnectionMapOutputWithContext(context.Background())
}

func (i BgpConnectionMap) ToBgpConnectionMapOutputWithContext(ctx context.Context) BgpConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpConnectionMapOutput)
}

type BgpConnectionOutput struct{ *pulumi.OutputState }

func (BgpConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpConnection)(nil))
}

func (o BgpConnectionOutput) ToBgpConnectionOutput() BgpConnectionOutput {
	return o
}

func (o BgpConnectionOutput) ToBgpConnectionOutputWithContext(ctx context.Context) BgpConnectionOutput {
	return o
}

func (o BgpConnectionOutput) ToBgpConnectionPtrOutput() BgpConnectionPtrOutput {
	return o.ToBgpConnectionPtrOutputWithContext(context.Background())
}

func (o BgpConnectionOutput) ToBgpConnectionPtrOutputWithContext(ctx context.Context) BgpConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpConnection) *BgpConnection {
		return &v
	}).(BgpConnectionPtrOutput)
}

type BgpConnectionPtrOutput struct{ *pulumi.OutputState }

func (BgpConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpConnection)(nil))
}

func (o BgpConnectionPtrOutput) ToBgpConnectionPtrOutput() BgpConnectionPtrOutput {
	return o
}

func (o BgpConnectionPtrOutput) ToBgpConnectionPtrOutputWithContext(ctx context.Context) BgpConnectionPtrOutput {
	return o
}

func (o BgpConnectionPtrOutput) Elem() BgpConnectionOutput {
	return o.ApplyT(func(v *BgpConnection) BgpConnection {
		if v != nil {
			return *v
		}
		var ret BgpConnection
		return ret
	}).(BgpConnectionOutput)
}

type BgpConnectionArrayOutput struct{ *pulumi.OutputState }

func (BgpConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpConnection)(nil))
}

func (o BgpConnectionArrayOutput) ToBgpConnectionArrayOutput() BgpConnectionArrayOutput {
	return o
}

func (o BgpConnectionArrayOutput) ToBgpConnectionArrayOutputWithContext(ctx context.Context) BgpConnectionArrayOutput {
	return o
}

func (o BgpConnectionArrayOutput) Index(i pulumi.IntInput) BgpConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BgpConnection {
		return vs[0].([]BgpConnection)[vs[1].(int)]
	}).(BgpConnectionOutput)
}

type BgpConnectionMapOutput struct{ *pulumi.OutputState }

func (BgpConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BgpConnection)(nil))
}

func (o BgpConnectionMapOutput) ToBgpConnectionMapOutput() BgpConnectionMapOutput {
	return o
}

func (o BgpConnectionMapOutput) ToBgpConnectionMapOutputWithContext(ctx context.Context) BgpConnectionMapOutput {
	return o
}

func (o BgpConnectionMapOutput) MapIndex(k pulumi.StringInput) BgpConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BgpConnection {
		return vs[0].(map[string]BgpConnection)[vs[1].(string)]
	}).(BgpConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(BgpConnectionOutput{})
	pulumi.RegisterOutputType(BgpConnectionPtrOutput{})
	pulumi.RegisterOutputType(BgpConnectionArrayOutput{})
	pulumi.RegisterOutputType(BgpConnectionMapOutput{})
}
