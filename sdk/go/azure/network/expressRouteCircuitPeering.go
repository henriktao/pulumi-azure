// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ExpressRoute Circuit Peering.
//
// ## Example Usage
// ### Creating A Microsoft Peering)
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
// 		exampleExpressRouteCircuit, err := network.NewExpressRouteCircuit(ctx, "exampleExpressRouteCircuit", &network.ExpressRouteCircuitArgs{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Location:            exampleResourceGroup.Location,
// 			ServiceProviderName: pulumi.String("Equinix"),
// 			PeeringLocation:     pulumi.String("Silicon Valley"),
// 			BandwidthInMbps:     pulumi.Int(50),
// 			Sku: &network.ExpressRouteCircuitSkuArgs{
// 				Tier:   pulumi.String("Standard"),
// 				Family: pulumi.String("MeteredData"),
// 			},
// 			AllowClassicOperations: pulumi.Bool(false),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewExpressRouteCircuitPeering(ctx, "exampleExpressRouteCircuitPeering", &network.ExpressRouteCircuitPeeringArgs{
// 			PeeringType:                pulumi.String("MicrosoftPeering"),
// 			ExpressRouteCircuitName:    exampleExpressRouteCircuit.Name,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			PeerAsn:                    pulumi.Int(100),
// 			PrimaryPeerAddressPrefix:   pulumi.String("123.0.0.0/30"),
// 			SecondaryPeerAddressPrefix: pulumi.String("123.0.0.4/30"),
// 			VlanId:                     pulumi.Int(300),
// 			MicrosoftPeeringConfig: &network.ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs{
// 				AdvertisedPublicPrefixes: pulumi.StringArray{
// 					pulumi.String("123.1.0.0/24"),
// 				},
// 			},
// 			Ipv6: &network.ExpressRouteCircuitPeeringIpv6Args{
// 				PrimaryPeerAddressPrefix:   pulumi.String("2002:db01::/126"),
// 				SecondaryPeerAddressPrefix: pulumi.String("2003:db01::/126"),
// 				MicrosoftPeering: &network.ExpressRouteCircuitPeeringIpv6MicrosoftPeeringArgs{
// 					AdvertisedPublicPrefixes: pulumi.StringArray{
// 						pulumi.String("2002:db01::/126"),
// 					},
// 				},
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
// ExpressRoute Circuit Peerings can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering peering1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute/peerings/peering1
// ```
type ExpressRouteCircuitPeering struct {
	pulumi.CustomResourceState

	// The ASN used by Azure.
	AzureAsn pulumi.IntOutput `pulumi:"azureAsn"`
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName pulumi.StringOutput `pulumi:"expressRouteCircuitName"`
	// A `ipv6` block as defined below.
	Ipv6 ExpressRouteCircuitPeeringIpv6PtrOutput `pulumi:"ipv6"`
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringMicrosoftPeeringConfigPtrOutput `pulumi:"microsoftPeeringConfig"`
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	PeerAsn pulumi.IntOutput `pulumi:"peerAsn"`
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType pulumi.StringOutput `pulumi:"peeringType"`
	// The Primary Port used by Azure for this Peering.
	PrimaryAzurePort pulumi.StringOutput `pulumi:"primaryAzurePort"`
	// A subnet for the primary link.
	PrimaryPeerAddressPrefix pulumi.StringOutput `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Route Filter. Only available when `peeringType` is set to `MicrosoftPeering`.
	RouteFilterId pulumi.StringPtrOutput `pulumi:"routeFilterId"`
	// The Secondary Port used by Azure for this Peering.
	SecondaryAzurePort pulumi.StringOutput `pulumi:"secondaryAzurePort"`
	// A subnet for the secondary link.
	SecondaryPeerAddressPrefix pulumi.StringOutput `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key. Can be a maximum of 25 characters.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// A valid VLAN ID to establish this peering on.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewExpressRouteCircuitPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressRouteCircuitName == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRouteCircuitName'")
	}
	if args.PeeringType == nil {
		return nil, errors.New("invalid value for required argument 'PeeringType'")
	}
	if args.PrimaryPeerAddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryPeerAddressPrefix'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecondaryPeerAddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'SecondaryPeerAddressPrefix'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	var resource ExpressRouteCircuitPeering
	err := ctx.RegisterResource("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitPeering gets an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	var resource ExpressRouteCircuitPeering
	err := ctx.ReadResource("azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitPeering resources.
type expressRouteCircuitPeeringState struct {
	// The ASN used by Azure.
	AzureAsn *int `pulumi:"azureAsn"`
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName *string `pulumi:"expressRouteCircuitName"`
	// A `ipv6` block as defined below.
	Ipv6 *ExpressRouteCircuitPeeringIpv6 `pulumi:"ipv6"`
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringMicrosoftPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	PeerAsn *int `pulumi:"peerAsn"`
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType *string `pulumi:"peeringType"`
	// The Primary Port used by Azure for this Peering.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// A subnet for the primary link.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Route Filter. Only available when `peeringType` is set to `MicrosoftPeering`.
	RouteFilterId *string `pulumi:"routeFilterId"`
	// The Secondary Port used by Azure for this Peering.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// A subnet for the secondary link.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key. Can be a maximum of 25 characters.
	SharedKey *string `pulumi:"sharedKey"`
	// A valid VLAN ID to establish this peering on.
	VlanId *int `pulumi:"vlanId"`
}

type ExpressRouteCircuitPeeringState struct {
	// The ASN used by Azure.
	AzureAsn pulumi.IntPtrInput
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName pulumi.StringPtrInput
	// A `ipv6` block as defined below.
	Ipv6 ExpressRouteCircuitPeeringIpv6PtrInput
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringMicrosoftPeeringConfigPtrInput
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	PeerAsn pulumi.IntPtrInput
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType pulumi.StringPtrInput
	// The Primary Port used by Azure for this Peering.
	PrimaryAzurePort pulumi.StringPtrInput
	// A subnet for the primary link.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Route Filter. Only available when `peeringType` is set to `MicrosoftPeering`.
	RouteFilterId pulumi.StringPtrInput
	// The Secondary Port used by Azure for this Peering.
	SecondaryAzurePort pulumi.StringPtrInput
	// A subnet for the secondary link.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key. Can be a maximum of 25 characters.
	SharedKey pulumi.StringPtrInput
	// A valid VLAN ID to establish this peering on.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringState)(nil)).Elem()
}

type expressRouteCircuitPeeringArgs struct {
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName string `pulumi:"expressRouteCircuitName"`
	// A `ipv6` block as defined below.
	Ipv6 *ExpressRouteCircuitPeeringIpv6 `pulumi:"ipv6"`
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringMicrosoftPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	PeerAsn *int `pulumi:"peerAsn"`
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType string `pulumi:"peeringType"`
	// A subnet for the primary link.
	PrimaryPeerAddressPrefix string `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Route Filter. Only available when `peeringType` is set to `MicrosoftPeering`.
	RouteFilterId *string `pulumi:"routeFilterId"`
	// A subnet for the secondary link.
	SecondaryPeerAddressPrefix string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key. Can be a maximum of 25 characters.
	SharedKey *string `pulumi:"sharedKey"`
	// A valid VLAN ID to establish this peering on.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitPeeringArgs struct {
	// The name of the ExpressRoute Circuit in which to create the Peering.
	ExpressRouteCircuitName pulumi.StringInput
	// A `ipv6` block as defined below.
	Ipv6 ExpressRouteCircuitPeeringIpv6PtrInput
	// A `microsoftPeeringConfig` block as defined below. Required when `peeringType` is set to `MicrosoftPeering`.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringMicrosoftPeeringConfigPtrInput
	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	PeerAsn pulumi.IntPtrInput
	// The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
	PeeringType pulumi.StringInput
	// A subnet for the primary link.
	PrimaryPeerAddressPrefix pulumi.StringInput
	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Route Filter. Only available when `peeringType` is set to `MicrosoftPeering`.
	RouteFilterId pulumi.StringPtrInput
	// A subnet for the secondary link.
	SecondaryPeerAddressPrefix pulumi.StringInput
	// The shared key. Can be a maximum of 25 characters.
	SharedKey pulumi.StringPtrInput
	// A valid VLAN ID to establish this peering on.
	VlanId pulumi.IntInput
}

func (ExpressRouteCircuitPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringArgs)(nil)).Elem()
}

type ExpressRouteCircuitPeeringInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput
	ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput
}

func (*ExpressRouteCircuitPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeering)(nil))
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return i.ToExpressRouteCircuitPeeringOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringOutput)
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringPtrOutput() ExpressRouteCircuitPeeringPtrOutput {
	return i.ToExpressRouteCircuitPeeringPtrOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringPtrOutput)
}

type ExpressRouteCircuitPeeringPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringPtrOutput() ExpressRouteCircuitPeeringPtrOutput
	ToExpressRouteCircuitPeeringPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringPtrOutput
}

type expressRouteCircuitPeeringPtrType ExpressRouteCircuitPeeringArgs

func (*expressRouteCircuitPeeringPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeering)(nil))
}

func (i *expressRouteCircuitPeeringPtrType) ToExpressRouteCircuitPeeringPtrOutput() ExpressRouteCircuitPeeringPtrOutput {
	return i.ToExpressRouteCircuitPeeringPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitPeeringPtrType) ToExpressRouteCircuitPeeringPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringPtrOutput)
}

// ExpressRouteCircuitPeeringArrayInput is an input type that accepts ExpressRouteCircuitPeeringArray and ExpressRouteCircuitPeeringArrayOutput values.
// You can construct a concrete instance of `ExpressRouteCircuitPeeringArrayInput` via:
//
//          ExpressRouteCircuitPeeringArray{ ExpressRouteCircuitPeeringArgs{...} }
type ExpressRouteCircuitPeeringArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringArrayOutput() ExpressRouteCircuitPeeringArrayOutput
	ToExpressRouteCircuitPeeringArrayOutputWithContext(context.Context) ExpressRouteCircuitPeeringArrayOutput
}

type ExpressRouteCircuitPeeringArray []ExpressRouteCircuitPeeringInput

func (ExpressRouteCircuitPeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExpressRouteCircuitPeering)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringArray) ToExpressRouteCircuitPeeringArrayOutput() ExpressRouteCircuitPeeringArrayOutput {
	return i.ToExpressRouteCircuitPeeringArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringArray) ToExpressRouteCircuitPeeringArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringArrayOutput)
}

// ExpressRouteCircuitPeeringMapInput is an input type that accepts ExpressRouteCircuitPeeringMap and ExpressRouteCircuitPeeringMapOutput values.
// You can construct a concrete instance of `ExpressRouteCircuitPeeringMapInput` via:
//
//          ExpressRouteCircuitPeeringMap{ "key": ExpressRouteCircuitPeeringArgs{...} }
type ExpressRouteCircuitPeeringMapInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringMapOutput() ExpressRouteCircuitPeeringMapOutput
	ToExpressRouteCircuitPeeringMapOutputWithContext(context.Context) ExpressRouteCircuitPeeringMapOutput
}

type ExpressRouteCircuitPeeringMap map[string]ExpressRouteCircuitPeeringInput

func (ExpressRouteCircuitPeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExpressRouteCircuitPeering)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringMap) ToExpressRouteCircuitPeeringMapOutput() ExpressRouteCircuitPeeringMapOutput {
	return i.ToExpressRouteCircuitPeeringMapOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringMap) ToExpressRouteCircuitPeeringMapOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringMapOutput)
}

type ExpressRouteCircuitPeeringOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeering)(nil))
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return o
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return o
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringPtrOutput() ExpressRouteCircuitPeeringPtrOutput {
	return o.ToExpressRouteCircuitPeeringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeering) *ExpressRouteCircuitPeering {
		return &v
	}).(ExpressRouteCircuitPeeringPtrOutput)
}

type ExpressRouteCircuitPeeringPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeering)(nil))
}

func (o ExpressRouteCircuitPeeringPtrOutput) ToExpressRouteCircuitPeeringPtrOutput() ExpressRouteCircuitPeeringPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringPtrOutput) ToExpressRouteCircuitPeeringPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringPtrOutput) Elem() ExpressRouteCircuitPeeringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeering) ExpressRouteCircuitPeering {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeering
		return ret
	}).(ExpressRouteCircuitPeeringOutput)
}

type ExpressRouteCircuitPeeringArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeering)(nil))
}

func (o ExpressRouteCircuitPeeringArrayOutput) ToExpressRouteCircuitPeeringArrayOutput() ExpressRouteCircuitPeeringArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringArrayOutput) ToExpressRouteCircuitPeeringArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitPeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeering {
		return vs[0].([]ExpressRouteCircuitPeering)[vs[1].(int)]
	}).(ExpressRouteCircuitPeeringOutput)
}

type ExpressRouteCircuitPeeringMapOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExpressRouteCircuitPeering)(nil))
}

func (o ExpressRouteCircuitPeeringMapOutput) ToExpressRouteCircuitPeeringMapOutput() ExpressRouteCircuitPeeringMapOutput {
	return o
}

func (o ExpressRouteCircuitPeeringMapOutput) ToExpressRouteCircuitPeeringMapOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringMapOutput {
	return o
}

func (o ExpressRouteCircuitPeeringMapOutput) MapIndex(k pulumi.StringInput) ExpressRouteCircuitPeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeering {
		return vs[0].(map[string]ExpressRouteCircuitPeering)[vs[1].(string)]
	}).(ExpressRouteCircuitPeeringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringMapOutput{})
}
