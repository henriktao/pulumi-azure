// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage Private DNS zone Virtual Network Links. These Links enable DNS resolution and registration inside Azure Virtual Networks using Azure Private DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/privatedns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleZone, err := privatedns.NewZone(ctx, "exampleZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewZoneVirtualNetworkLink(ctx, "exampleZoneVirtualNetworkLink", &privatedns.ZoneVirtualNetworkLinkArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			PrivateDnsZoneName: exampleZone.Name,
// 			VirtualNetworkId:   pulumi.Any(azurerm_virtual_network.Example.Id),
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
// Private DNS Zone Virtual Network Links can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink link1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1.com/virtualNetworkLinks/myVnetLink1
// ```
type ZoneVirtualNetworkLink struct {
	pulumi.CustomResourceState

	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringOutput `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrOutput `pulumi:"registrationEnabled"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringOutput `pulumi:"virtualNetworkId"`
}

// NewZoneVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *ZoneVirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*ZoneVirtualNetworkLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateDnsZoneName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateDnsZoneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkId'")
	}
	var resource ZoneVirtualNetworkLink
	err := ctx.RegisterResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneVirtualNetworkLink gets an existing ZoneVirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneVirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*ZoneVirtualNetworkLink, error) {
	var resource ZoneVirtualNetworkLink
	err := ctx.ReadResource("azure:privatedns/zoneVirtualNetworkLink:ZoneVirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneVirtualNetworkLink resources.
type zoneVirtualNetworkLinkState struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName *string `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled *bool `pulumi:"registrationEnabled"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
}

type ZoneVirtualNetworkLinkState struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringPtrInput
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringPtrInput
}

func (ZoneVirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVirtualNetworkLinkState)(nil)).Elem()
}

type zoneVirtualNetworkLinkArgs struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName string `pulumi:"privateDnsZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled *bool `pulumi:"registrationEnabled"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId string `pulumi:"virtualNetworkId"`
}

// The set of arguments for constructing a ZoneVirtualNetworkLink resource.
type ZoneVirtualNetworkLinkArgs struct {
	// The name of the Private DNS Zone Virtual Network Link. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Private DNS zone (without a terminating dot). Changing this forces a new resource to be created.
	PrivateDnsZoneName pulumi.StringInput
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled? Defaults to `false`.
	RegistrationEnabled pulumi.BoolPtrInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringInput
}

func (ZoneVirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVirtualNetworkLinkArgs)(nil)).Elem()
}

type ZoneVirtualNetworkLinkInput interface {
	pulumi.Input

	ToZoneVirtualNetworkLinkOutput() ZoneVirtualNetworkLinkOutput
	ToZoneVirtualNetworkLinkOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkOutput
}

func (*ZoneVirtualNetworkLink) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneVirtualNetworkLink)(nil))
}

func (i *ZoneVirtualNetworkLink) ToZoneVirtualNetworkLinkOutput() ZoneVirtualNetworkLinkOutput {
	return i.ToZoneVirtualNetworkLinkOutputWithContext(context.Background())
}

func (i *ZoneVirtualNetworkLink) ToZoneVirtualNetworkLinkOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVirtualNetworkLinkOutput)
}

func (i *ZoneVirtualNetworkLink) ToZoneVirtualNetworkLinkPtrOutput() ZoneVirtualNetworkLinkPtrOutput {
	return i.ToZoneVirtualNetworkLinkPtrOutputWithContext(context.Background())
}

func (i *ZoneVirtualNetworkLink) ToZoneVirtualNetworkLinkPtrOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVirtualNetworkLinkPtrOutput)
}

type ZoneVirtualNetworkLinkPtrInput interface {
	pulumi.Input

	ToZoneVirtualNetworkLinkPtrOutput() ZoneVirtualNetworkLinkPtrOutput
	ToZoneVirtualNetworkLinkPtrOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkPtrOutput
}

type zoneVirtualNetworkLinkPtrType ZoneVirtualNetworkLinkArgs

func (*zoneVirtualNetworkLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneVirtualNetworkLink)(nil))
}

func (i *zoneVirtualNetworkLinkPtrType) ToZoneVirtualNetworkLinkPtrOutput() ZoneVirtualNetworkLinkPtrOutput {
	return i.ToZoneVirtualNetworkLinkPtrOutputWithContext(context.Background())
}

func (i *zoneVirtualNetworkLinkPtrType) ToZoneVirtualNetworkLinkPtrOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVirtualNetworkLinkPtrOutput)
}

// ZoneVirtualNetworkLinkArrayInput is an input type that accepts ZoneVirtualNetworkLinkArray and ZoneVirtualNetworkLinkArrayOutput values.
// You can construct a concrete instance of `ZoneVirtualNetworkLinkArrayInput` via:
//
//          ZoneVirtualNetworkLinkArray{ ZoneVirtualNetworkLinkArgs{...} }
type ZoneVirtualNetworkLinkArrayInput interface {
	pulumi.Input

	ToZoneVirtualNetworkLinkArrayOutput() ZoneVirtualNetworkLinkArrayOutput
	ToZoneVirtualNetworkLinkArrayOutputWithContext(context.Context) ZoneVirtualNetworkLinkArrayOutput
}

type ZoneVirtualNetworkLinkArray []ZoneVirtualNetworkLinkInput

func (ZoneVirtualNetworkLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ZoneVirtualNetworkLink)(nil))
}

func (i ZoneVirtualNetworkLinkArray) ToZoneVirtualNetworkLinkArrayOutput() ZoneVirtualNetworkLinkArrayOutput {
	return i.ToZoneVirtualNetworkLinkArrayOutputWithContext(context.Background())
}

func (i ZoneVirtualNetworkLinkArray) ToZoneVirtualNetworkLinkArrayOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVirtualNetworkLinkArrayOutput)
}

// ZoneVirtualNetworkLinkMapInput is an input type that accepts ZoneVirtualNetworkLinkMap and ZoneVirtualNetworkLinkMapOutput values.
// You can construct a concrete instance of `ZoneVirtualNetworkLinkMapInput` via:
//
//          ZoneVirtualNetworkLinkMap{ "key": ZoneVirtualNetworkLinkArgs{...} }
type ZoneVirtualNetworkLinkMapInput interface {
	pulumi.Input

	ToZoneVirtualNetworkLinkMapOutput() ZoneVirtualNetworkLinkMapOutput
	ToZoneVirtualNetworkLinkMapOutputWithContext(context.Context) ZoneVirtualNetworkLinkMapOutput
}

type ZoneVirtualNetworkLinkMap map[string]ZoneVirtualNetworkLinkInput

func (ZoneVirtualNetworkLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ZoneVirtualNetworkLink)(nil))
}

func (i ZoneVirtualNetworkLinkMap) ToZoneVirtualNetworkLinkMapOutput() ZoneVirtualNetworkLinkMapOutput {
	return i.ToZoneVirtualNetworkLinkMapOutputWithContext(context.Background())
}

func (i ZoneVirtualNetworkLinkMap) ToZoneVirtualNetworkLinkMapOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVirtualNetworkLinkMapOutput)
}

type ZoneVirtualNetworkLinkOutput struct {
	*pulumi.OutputState
}

func (ZoneVirtualNetworkLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneVirtualNetworkLink)(nil))
}

func (o ZoneVirtualNetworkLinkOutput) ToZoneVirtualNetworkLinkOutput() ZoneVirtualNetworkLinkOutput {
	return o
}

func (o ZoneVirtualNetworkLinkOutput) ToZoneVirtualNetworkLinkOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkOutput {
	return o
}

func (o ZoneVirtualNetworkLinkOutput) ToZoneVirtualNetworkLinkPtrOutput() ZoneVirtualNetworkLinkPtrOutput {
	return o.ToZoneVirtualNetworkLinkPtrOutputWithContext(context.Background())
}

func (o ZoneVirtualNetworkLinkOutput) ToZoneVirtualNetworkLinkPtrOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkPtrOutput {
	return o.ApplyT(func(v ZoneVirtualNetworkLink) *ZoneVirtualNetworkLink {
		return &v
	}).(ZoneVirtualNetworkLinkPtrOutput)
}

type ZoneVirtualNetworkLinkPtrOutput struct {
	*pulumi.OutputState
}

func (ZoneVirtualNetworkLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneVirtualNetworkLink)(nil))
}

func (o ZoneVirtualNetworkLinkPtrOutput) ToZoneVirtualNetworkLinkPtrOutput() ZoneVirtualNetworkLinkPtrOutput {
	return o
}

func (o ZoneVirtualNetworkLinkPtrOutput) ToZoneVirtualNetworkLinkPtrOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkPtrOutput {
	return o
}

type ZoneVirtualNetworkLinkArrayOutput struct{ *pulumi.OutputState }

func (ZoneVirtualNetworkLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneVirtualNetworkLink)(nil))
}

func (o ZoneVirtualNetworkLinkArrayOutput) ToZoneVirtualNetworkLinkArrayOutput() ZoneVirtualNetworkLinkArrayOutput {
	return o
}

func (o ZoneVirtualNetworkLinkArrayOutput) ToZoneVirtualNetworkLinkArrayOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkArrayOutput {
	return o
}

func (o ZoneVirtualNetworkLinkArrayOutput) Index(i pulumi.IntInput) ZoneVirtualNetworkLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneVirtualNetworkLink {
		return vs[0].([]ZoneVirtualNetworkLink)[vs[1].(int)]
	}).(ZoneVirtualNetworkLinkOutput)
}

type ZoneVirtualNetworkLinkMapOutput struct{ *pulumi.OutputState }

func (ZoneVirtualNetworkLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ZoneVirtualNetworkLink)(nil))
}

func (o ZoneVirtualNetworkLinkMapOutput) ToZoneVirtualNetworkLinkMapOutput() ZoneVirtualNetworkLinkMapOutput {
	return o
}

func (o ZoneVirtualNetworkLinkMapOutput) ToZoneVirtualNetworkLinkMapOutputWithContext(ctx context.Context) ZoneVirtualNetworkLinkMapOutput {
	return o
}

func (o ZoneVirtualNetworkLinkMapOutput) MapIndex(k pulumi.StringInput) ZoneVirtualNetworkLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ZoneVirtualNetworkLink {
		return vs[0].(map[string]ZoneVirtualNetworkLink)[vs[1].(string)]
	}).(ZoneVirtualNetworkLinkOutput)
}

func init() {
	pulumi.RegisterOutputType(ZoneVirtualNetworkLinkOutput{})
	pulumi.RegisterOutputType(ZoneVirtualNetworkLinkPtrOutput{})
	pulumi.RegisterOutputType(ZoneVirtualNetworkLinkArrayOutput{})
	pulumi.RegisterOutputType(ZoneVirtualNetworkLinkMapOutput{})
}
