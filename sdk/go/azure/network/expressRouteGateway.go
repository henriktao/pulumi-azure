// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ExpressRoute gateway.
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
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewExpressRouteGateway(ctx, "exampleExpressRouteGateway", &network.ExpressRouteGatewayArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualHubId:      exampleVirtualHub.ID(),
// 			ScaleUnits:        pulumi.Int(1),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// ExpressRoute Gateways can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/expressRouteGateway:ExpressRouteGateway example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteGateways/myExpressRouteGateway
// ```
type ExpressRouteGateway struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the ExpressRoute gateway. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	ScaleUnits pulumi.IntOutput `pulumi:"scaleUnits"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
}

// NewExpressRouteGateway registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteGateway(ctx *pulumi.Context,
	name string, args *ExpressRouteGatewayArgs, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScaleUnits == nil {
		return nil, errors.New("invalid value for required argument 'ScaleUnits'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource ExpressRouteGateway
	err := ctx.RegisterResource("azure:network/expressRouteGateway:ExpressRouteGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteGateway gets an existing ExpressRouteGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteGatewayState, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	var resource ExpressRouteGateway
	err := ctx.ReadResource("azure:network/expressRouteGateway:ExpressRouteGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteGateway resources.
type expressRouteGatewayState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	ScaleUnits *int `pulumi:"scaleUnits"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
}

type ExpressRouteGatewayState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	ScaleUnits pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	VirtualHubId pulumi.StringPtrInput
}

func (ExpressRouteGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayState)(nil)).Elem()
}

type expressRouteGatewayArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	ScaleUnits int `pulumi:"scaleUnits"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	VirtualHubId string `pulumi:"virtualHubId"`
}

// The set of arguments for constructing a ExpressRouteGateway resource.
type ExpressRouteGatewayArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the ExpressRoute gateway. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The number of scale units with which to provision the ExpressRoute gateway. Each scale unit is equal to 2Gbps, with support for up to 10 scale units (20Gbps).
	ScaleUnits pulumi.IntInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of a Virtual HUB within which the ExpressRoute gateway should be created.
	VirtualHubId pulumi.StringInput
}

func (ExpressRouteGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayArgs)(nil)).Elem()
}

type ExpressRouteGatewayInput interface {
	pulumi.Input

	ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput
	ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput
}

func (*ExpressRouteGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteGateway)(nil))
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return i.ToExpressRouteGatewayOutputWithContext(context.Background())
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayOutput)
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayPtrOutput() ExpressRouteGatewayPtrOutput {
	return i.ToExpressRouteGatewayPtrOutputWithContext(context.Background())
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayPtrOutputWithContext(ctx context.Context) ExpressRouteGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayPtrOutput)
}

type ExpressRouteGatewayPtrInput interface {
	pulumi.Input

	ToExpressRouteGatewayPtrOutput() ExpressRouteGatewayPtrOutput
	ToExpressRouteGatewayPtrOutputWithContext(ctx context.Context) ExpressRouteGatewayPtrOutput
}

type expressRouteGatewayPtrType ExpressRouteGatewayArgs

func (*expressRouteGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil))
}

func (i *expressRouteGatewayPtrType) ToExpressRouteGatewayPtrOutput() ExpressRouteGatewayPtrOutput {
	return i.ToExpressRouteGatewayPtrOutputWithContext(context.Background())
}

func (i *expressRouteGatewayPtrType) ToExpressRouteGatewayPtrOutputWithContext(ctx context.Context) ExpressRouteGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayPtrOutput)
}

// ExpressRouteGatewayArrayInput is an input type that accepts ExpressRouteGatewayArray and ExpressRouteGatewayArrayOutput values.
// You can construct a concrete instance of `ExpressRouteGatewayArrayInput` via:
//
//          ExpressRouteGatewayArray{ ExpressRouteGatewayArgs{...} }
type ExpressRouteGatewayArrayInput interface {
	pulumi.Input

	ToExpressRouteGatewayArrayOutput() ExpressRouteGatewayArrayOutput
	ToExpressRouteGatewayArrayOutputWithContext(context.Context) ExpressRouteGatewayArrayOutput
}

type ExpressRouteGatewayArray []ExpressRouteGatewayInput

func (ExpressRouteGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExpressRouteGateway)(nil)).Elem()
}

func (i ExpressRouteGatewayArray) ToExpressRouteGatewayArrayOutput() ExpressRouteGatewayArrayOutput {
	return i.ToExpressRouteGatewayArrayOutputWithContext(context.Background())
}

func (i ExpressRouteGatewayArray) ToExpressRouteGatewayArrayOutputWithContext(ctx context.Context) ExpressRouteGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayArrayOutput)
}

// ExpressRouteGatewayMapInput is an input type that accepts ExpressRouteGatewayMap and ExpressRouteGatewayMapOutput values.
// You can construct a concrete instance of `ExpressRouteGatewayMapInput` via:
//
//          ExpressRouteGatewayMap{ "key": ExpressRouteGatewayArgs{...} }
type ExpressRouteGatewayMapInput interface {
	pulumi.Input

	ToExpressRouteGatewayMapOutput() ExpressRouteGatewayMapOutput
	ToExpressRouteGatewayMapOutputWithContext(context.Context) ExpressRouteGatewayMapOutput
}

type ExpressRouteGatewayMap map[string]ExpressRouteGatewayInput

func (ExpressRouteGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExpressRouteGateway)(nil)).Elem()
}

func (i ExpressRouteGatewayMap) ToExpressRouteGatewayMapOutput() ExpressRouteGatewayMapOutput {
	return i.ToExpressRouteGatewayMapOutputWithContext(context.Background())
}

func (i ExpressRouteGatewayMap) ToExpressRouteGatewayMapOutputWithContext(ctx context.Context) ExpressRouteGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayMapOutput)
}

type ExpressRouteGatewayOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteGateway)(nil))
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return o
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return o
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayPtrOutput() ExpressRouteGatewayPtrOutput {
	return o.ToExpressRouteGatewayPtrOutputWithContext(context.Background())
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayPtrOutputWithContext(ctx context.Context) ExpressRouteGatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteGateway) *ExpressRouteGateway {
		return &v
	}).(ExpressRouteGatewayPtrOutput)
}

type ExpressRouteGatewayPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil))
}

func (o ExpressRouteGatewayPtrOutput) ToExpressRouteGatewayPtrOutput() ExpressRouteGatewayPtrOutput {
	return o
}

func (o ExpressRouteGatewayPtrOutput) ToExpressRouteGatewayPtrOutputWithContext(ctx context.Context) ExpressRouteGatewayPtrOutput {
	return o
}

func (o ExpressRouteGatewayPtrOutput) Elem() ExpressRouteGatewayOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) ExpressRouteGateway {
		if v != nil {
			return *v
		}
		var ret ExpressRouteGateway
		return ret
	}).(ExpressRouteGatewayOutput)
}

type ExpressRouteGatewayArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteGateway)(nil))
}

func (o ExpressRouteGatewayArrayOutput) ToExpressRouteGatewayArrayOutput() ExpressRouteGatewayArrayOutput {
	return o
}

func (o ExpressRouteGatewayArrayOutput) ToExpressRouteGatewayArrayOutputWithContext(ctx context.Context) ExpressRouteGatewayArrayOutput {
	return o
}

func (o ExpressRouteGatewayArrayOutput) Index(i pulumi.IntInput) ExpressRouteGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteGateway {
		return vs[0].([]ExpressRouteGateway)[vs[1].(int)]
	}).(ExpressRouteGatewayOutput)
}

type ExpressRouteGatewayMapOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExpressRouteGateway)(nil))
}

func (o ExpressRouteGatewayMapOutput) ToExpressRouteGatewayMapOutput() ExpressRouteGatewayMapOutput {
	return o
}

func (o ExpressRouteGatewayMapOutput) ToExpressRouteGatewayMapOutputWithContext(ctx context.Context) ExpressRouteGatewayMapOutput {
	return o
}

func (o ExpressRouteGatewayMapOutput) MapIndex(k pulumi.StringInput) ExpressRouteGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExpressRouteGateway {
		return vs[0].(map[string]ExpressRouteGateway)[vs[1].(string)]
	}).(ExpressRouteGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteGatewayOutput{})
	pulumi.RegisterOutputType(ExpressRouteGatewayPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteGatewayArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteGatewayMapOutput{})
}
