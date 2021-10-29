// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Virtual Hub within a Virtual WAN.
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
// 		_, err = network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.0.0/23"),
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
// Virtual Hub's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualHub:VirtualHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1
// ```
type VirtualHub struct {
	pulumi.CustomResourceState

	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. [The address prefix subnet cannot be smaller than a `/24`. Azure recommends using a `/23`](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-faq#what-is-the-recommended-hub-address-space-during-hub-creation).
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// The ID of the default Route Table in the Virtual Hub.
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayOutput `pulumi:"routes"`
	// The sku of the Virtual Hub. Possible values are `Basic` and `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringPtrOutput `pulumi:"virtualWanId"`
}

// NewVirtualHub registers a new resource with the given unique name, arguments, and options.
func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource VirtualHub
	err := ctx.RegisterResource("azure:network/virtualHub:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHub gets an existing VirtualHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure:network/virtualHub:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHub resources.
type virtualHubState struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. [The address prefix subnet cannot be smaller than a `/24`. Azure recommends using a `/23`](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-faq#what-is-the-recommended-hub-address-space-during-hub-creation).
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The ID of the default Route Table in the Virtual Hub.
	DefaultRouteTableId *string `pulumi:"defaultRouteTableId"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes []VirtualHubRoute `pulumi:"routes"`
	// The sku of the Virtual Hub. Possible values are `Basic` and `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId *string `pulumi:"virtualWanId"`
}

type VirtualHubState struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. [The address prefix subnet cannot be smaller than a `/24`. Azure recommends using a `/23`](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-faq#what-is-the-recommended-hub-address-space-during-hub-creation).
	AddressPrefix pulumi.StringPtrInput
	// The ID of the default Route Table in the Virtual Hub.
	DefaultRouteTableId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayInput
	// The sku of the Virtual Hub. Possible values are `Basic` and `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapInput
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringPtrInput
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. [The address prefix subnet cannot be smaller than a `/24`. Azure recommends using a `/23`](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-faq#what-is-the-recommended-hub-address-space-during-hub-creation).
	AddressPrefix *string `pulumi:"addressPrefix"`
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `route` blocks as defined below.
	Routes []VirtualHubRoute `pulumi:"routes"`
	// The sku of the Virtual Hub. Possible values are `Basic` and `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the Virtual Hub.
	Tags map[string]string `pulumi:"tags"`
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId *string `pulumi:"virtualWanId"`
}

// The set of arguments for constructing a VirtualHub resource.
type VirtualHubArgs struct {
	// The Address Prefix which should be used for this Virtual Hub. Changing this forces a new resource to be created. [The address prefix subnet cannot be smaller than a `/24`. Azure recommends using a `/23`](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-faq#what-is-the-recommended-hub-address-space-during-hub-creation).
	AddressPrefix pulumi.StringPtrInput
	// Specifies the supported Azure location where the Virtual Hub should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Virtual Hub should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `route` blocks as defined below.
	Routes VirtualHubRouteArrayInput
	// The sku of the Virtual Hub. Possible values are `Basic` and `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the Virtual Hub.
	Tags pulumi.StringMapInput
	// The ID of a Virtual WAN within which the Virtual Hub should be created. Changing this forces a new resource to be created.
	VirtualWanId pulumi.StringPtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}

type VirtualHubInput interface {
	pulumi.Input

	ToVirtualHubOutput() VirtualHubOutput
	ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput
}

func (*VirtualHub) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHub)(nil))
}

func (i *VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

func (i *VirtualHub) ToVirtualHubPtrOutput() VirtualHubPtrOutput {
	return i.ToVirtualHubPtrOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubPtrOutputWithContext(ctx context.Context) VirtualHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubPtrOutput)
}

type VirtualHubPtrInput interface {
	pulumi.Input

	ToVirtualHubPtrOutput() VirtualHubPtrOutput
	ToVirtualHubPtrOutputWithContext(ctx context.Context) VirtualHubPtrOutput
}

type virtualHubPtrType VirtualHubArgs

func (*virtualHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil))
}

func (i *virtualHubPtrType) ToVirtualHubPtrOutput() VirtualHubPtrOutput {
	return i.ToVirtualHubPtrOutputWithContext(context.Background())
}

func (i *virtualHubPtrType) ToVirtualHubPtrOutputWithContext(ctx context.Context) VirtualHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubPtrOutput)
}

// VirtualHubArrayInput is an input type that accepts VirtualHubArray and VirtualHubArrayOutput values.
// You can construct a concrete instance of `VirtualHubArrayInput` via:
//
//          VirtualHubArray{ VirtualHubArgs{...} }
type VirtualHubArrayInput interface {
	pulumi.Input

	ToVirtualHubArrayOutput() VirtualHubArrayOutput
	ToVirtualHubArrayOutputWithContext(context.Context) VirtualHubArrayOutput
}

type VirtualHubArray []VirtualHubInput

func (VirtualHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHub)(nil)).Elem()
}

func (i VirtualHubArray) ToVirtualHubArrayOutput() VirtualHubArrayOutput {
	return i.ToVirtualHubArrayOutputWithContext(context.Background())
}

func (i VirtualHubArray) ToVirtualHubArrayOutputWithContext(ctx context.Context) VirtualHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubArrayOutput)
}

// VirtualHubMapInput is an input type that accepts VirtualHubMap and VirtualHubMapOutput values.
// You can construct a concrete instance of `VirtualHubMapInput` via:
//
//          VirtualHubMap{ "key": VirtualHubArgs{...} }
type VirtualHubMapInput interface {
	pulumi.Input

	ToVirtualHubMapOutput() VirtualHubMapOutput
	ToVirtualHubMapOutputWithContext(context.Context) VirtualHubMapOutput
}

type VirtualHubMap map[string]VirtualHubInput

func (VirtualHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHub)(nil)).Elem()
}

func (i VirtualHubMap) ToVirtualHubMapOutput() VirtualHubMapOutput {
	return i.ToVirtualHubMapOutputWithContext(context.Background())
}

func (i VirtualHubMap) ToVirtualHubMapOutputWithContext(ctx context.Context) VirtualHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubMapOutput)
}

type VirtualHubOutput struct{ *pulumi.OutputState }

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHub)(nil))
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubPtrOutput() VirtualHubPtrOutput {
	return o.ToVirtualHubPtrOutputWithContext(context.Background())
}

func (o VirtualHubOutput) ToVirtualHubPtrOutputWithContext(ctx context.Context) VirtualHubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHub) *VirtualHub {
		return &v
	}).(VirtualHubPtrOutput)
}

type VirtualHubPtrOutput struct{ *pulumi.OutputState }

func (VirtualHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil))
}

func (o VirtualHubPtrOutput) ToVirtualHubPtrOutput() VirtualHubPtrOutput {
	return o
}

func (o VirtualHubPtrOutput) ToVirtualHubPtrOutputWithContext(ctx context.Context) VirtualHubPtrOutput {
	return o
}

func (o VirtualHubPtrOutput) Elem() VirtualHubOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualHub {
		if v != nil {
			return *v
		}
		var ret VirtualHub
		return ret
	}).(VirtualHubOutput)
}

type VirtualHubArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHub)(nil))
}

func (o VirtualHubArrayOutput) ToVirtualHubArrayOutput() VirtualHubArrayOutput {
	return o
}

func (o VirtualHubArrayOutput) ToVirtualHubArrayOutputWithContext(ctx context.Context) VirtualHubArrayOutput {
	return o
}

func (o VirtualHubArrayOutput) Index(i pulumi.IntInput) VirtualHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHub {
		return vs[0].([]VirtualHub)[vs[1].(int)]
	}).(VirtualHubOutput)
}

type VirtualHubMapOutput struct{ *pulumi.OutputState }

func (VirtualHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualHub)(nil))
}

func (o VirtualHubMapOutput) ToVirtualHubMapOutput() VirtualHubMapOutput {
	return o
}

func (o VirtualHubMapOutput) ToVirtualHubMapOutputWithContext(ctx context.Context) VirtualHubMapOutput {
	return o
}

func (o VirtualHubMapOutput) MapIndex(k pulumi.StringInput) VirtualHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualHub {
		return vs[0].(map[string]VirtualHub)[vs[1].(string)]
	}).(VirtualHubOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubOutput{})
	pulumi.RegisterOutputType(VirtualHubPtrOutput{})
	pulumi.RegisterOutputType(VirtualHubArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubMapOutput{})
}
