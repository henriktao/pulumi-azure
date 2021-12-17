// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VPN Gateway Nat Rule.
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
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpnGateway, err := network.NewVpnGateway(ctx, "exampleVpnGateway", &network.VpnGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			VirtualHubId:      exampleVirtualHub.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVnpGatewayNatRule(ctx, "exampleVnpGatewayNatRule", &network.VnpGatewayNatRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			VpnGatewayId:      exampleVpnGateway.ID(),
// 			ExternalAddressSpaceMappings: pulumi.StringArray{
// 				pulumi.String("192.168.21.0/26"),
// 			},
// 			InternalAddressSpaceMappings: pulumi.StringArray{
// 				pulumi.String("10.4.0.0/26"),
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
// VPN Gateway Nat Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/vnpGatewayNatRule:VnpGatewayNatRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnGateways/vpnGateway1/natRules/natRule1
// ```
type VnpGatewayNatRule struct {
	pulumi.CustomResourceState

	// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
	ExternalAddressSpaceMappings pulumi.StringArrayOutput `pulumi:"externalAddressSpaceMappings"`
	// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
	InternalAddressSpaceMappings pulumi.StringArrayOutput `pulumi:"internalAddressSpaceMappings"`
	// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
	IpConfigurationId pulumi.StringPtrOutput `pulumi:"ipConfigurationId"`
	// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewVnpGatewayNatRule registers a new resource with the given unique name, arguments, and options.
func NewVnpGatewayNatRule(ctx *pulumi.Context,
	name string, args *VnpGatewayNatRuleArgs, opts ...pulumi.ResourceOption) (*VnpGatewayNatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalAddressSpaceMappings == nil {
		return nil, errors.New("invalid value for required argument 'ExternalAddressSpaceMappings'")
	}
	if args.InternalAddressSpaceMappings == nil {
		return nil, errors.New("invalid value for required argument 'InternalAddressSpaceMappings'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VpnGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VpnGatewayId'")
	}
	var resource VnpGatewayNatRule
	err := ctx.RegisterResource("azure:network/vnpGatewayNatRule:VnpGatewayNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVnpGatewayNatRule gets an existing VnpGatewayNatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVnpGatewayNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VnpGatewayNatRuleState, opts ...pulumi.ResourceOption) (*VnpGatewayNatRule, error) {
	var resource VnpGatewayNatRule
	err := ctx.ReadResource("azure:network/vnpGatewayNatRule:VnpGatewayNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VnpGatewayNatRule resources.
type vnpGatewayNatRuleState struct {
	// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
	ExternalAddressSpaceMappings []string `pulumi:"externalAddressSpaceMappings"`
	// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
	InternalAddressSpaceMappings []string `pulumi:"internalAddressSpaceMappings"`
	// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
	IpConfigurationId *string `pulumi:"ipConfigurationId"`
	// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
	Mode *string `pulumi:"mode"`
	// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type VnpGatewayNatRuleState struct {
	// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
	ExternalAddressSpaceMappings pulumi.StringArrayInput
	// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
	InternalAddressSpaceMappings pulumi.StringArrayInput
	// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
	IpConfigurationId pulumi.StringPtrInput
	// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
	Mode pulumi.StringPtrInput
	// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
	VpnGatewayId pulumi.StringPtrInput
}

func (VnpGatewayNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vnpGatewayNatRuleState)(nil)).Elem()
}

type vnpGatewayNatRuleArgs struct {
	// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
	ExternalAddressSpaceMappings []string `pulumi:"externalAddressSpaceMappings"`
	// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
	InternalAddressSpaceMappings []string `pulumi:"internalAddressSpaceMappings"`
	// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
	IpConfigurationId *string `pulumi:"ipConfigurationId"`
	// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
	Mode *string `pulumi:"mode"`
	// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VnpGatewayNatRule resource.
type VnpGatewayNatRuleArgs struct {
	// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
	ExternalAddressSpaceMappings pulumi.StringArrayInput
	// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
	InternalAddressSpaceMappings pulumi.StringArrayInput
	// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
	IpConfigurationId pulumi.StringPtrInput
	// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
	Mode pulumi.StringPtrInput
	// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
	VpnGatewayId pulumi.StringInput
}

func (VnpGatewayNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vnpGatewayNatRuleArgs)(nil)).Elem()
}

type VnpGatewayNatRuleInput interface {
	pulumi.Input

	ToVnpGatewayNatRuleOutput() VnpGatewayNatRuleOutput
	ToVnpGatewayNatRuleOutputWithContext(ctx context.Context) VnpGatewayNatRuleOutput
}

func (*VnpGatewayNatRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VnpGatewayNatRule)(nil))
}

func (i *VnpGatewayNatRule) ToVnpGatewayNatRuleOutput() VnpGatewayNatRuleOutput {
	return i.ToVnpGatewayNatRuleOutputWithContext(context.Background())
}

func (i *VnpGatewayNatRule) ToVnpGatewayNatRuleOutputWithContext(ctx context.Context) VnpGatewayNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnpGatewayNatRuleOutput)
}

func (i *VnpGatewayNatRule) ToVnpGatewayNatRulePtrOutput() VnpGatewayNatRulePtrOutput {
	return i.ToVnpGatewayNatRulePtrOutputWithContext(context.Background())
}

func (i *VnpGatewayNatRule) ToVnpGatewayNatRulePtrOutputWithContext(ctx context.Context) VnpGatewayNatRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnpGatewayNatRulePtrOutput)
}

type VnpGatewayNatRulePtrInput interface {
	pulumi.Input

	ToVnpGatewayNatRulePtrOutput() VnpGatewayNatRulePtrOutput
	ToVnpGatewayNatRulePtrOutputWithContext(ctx context.Context) VnpGatewayNatRulePtrOutput
}

type vnpGatewayNatRulePtrType VnpGatewayNatRuleArgs

func (*vnpGatewayNatRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VnpGatewayNatRule)(nil))
}

func (i *vnpGatewayNatRulePtrType) ToVnpGatewayNatRulePtrOutput() VnpGatewayNatRulePtrOutput {
	return i.ToVnpGatewayNatRulePtrOutputWithContext(context.Background())
}

func (i *vnpGatewayNatRulePtrType) ToVnpGatewayNatRulePtrOutputWithContext(ctx context.Context) VnpGatewayNatRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnpGatewayNatRulePtrOutput)
}

// VnpGatewayNatRuleArrayInput is an input type that accepts VnpGatewayNatRuleArray and VnpGatewayNatRuleArrayOutput values.
// You can construct a concrete instance of `VnpGatewayNatRuleArrayInput` via:
//
//          VnpGatewayNatRuleArray{ VnpGatewayNatRuleArgs{...} }
type VnpGatewayNatRuleArrayInput interface {
	pulumi.Input

	ToVnpGatewayNatRuleArrayOutput() VnpGatewayNatRuleArrayOutput
	ToVnpGatewayNatRuleArrayOutputWithContext(context.Context) VnpGatewayNatRuleArrayOutput
}

type VnpGatewayNatRuleArray []VnpGatewayNatRuleInput

func (VnpGatewayNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VnpGatewayNatRule)(nil)).Elem()
}

func (i VnpGatewayNatRuleArray) ToVnpGatewayNatRuleArrayOutput() VnpGatewayNatRuleArrayOutput {
	return i.ToVnpGatewayNatRuleArrayOutputWithContext(context.Background())
}

func (i VnpGatewayNatRuleArray) ToVnpGatewayNatRuleArrayOutputWithContext(ctx context.Context) VnpGatewayNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnpGatewayNatRuleArrayOutput)
}

// VnpGatewayNatRuleMapInput is an input type that accepts VnpGatewayNatRuleMap and VnpGatewayNatRuleMapOutput values.
// You can construct a concrete instance of `VnpGatewayNatRuleMapInput` via:
//
//          VnpGatewayNatRuleMap{ "key": VnpGatewayNatRuleArgs{...} }
type VnpGatewayNatRuleMapInput interface {
	pulumi.Input

	ToVnpGatewayNatRuleMapOutput() VnpGatewayNatRuleMapOutput
	ToVnpGatewayNatRuleMapOutputWithContext(context.Context) VnpGatewayNatRuleMapOutput
}

type VnpGatewayNatRuleMap map[string]VnpGatewayNatRuleInput

func (VnpGatewayNatRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VnpGatewayNatRule)(nil)).Elem()
}

func (i VnpGatewayNatRuleMap) ToVnpGatewayNatRuleMapOutput() VnpGatewayNatRuleMapOutput {
	return i.ToVnpGatewayNatRuleMapOutputWithContext(context.Background())
}

func (i VnpGatewayNatRuleMap) ToVnpGatewayNatRuleMapOutputWithContext(ctx context.Context) VnpGatewayNatRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnpGatewayNatRuleMapOutput)
}

type VnpGatewayNatRuleOutput struct{ *pulumi.OutputState }

func (VnpGatewayNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnpGatewayNatRule)(nil))
}

func (o VnpGatewayNatRuleOutput) ToVnpGatewayNatRuleOutput() VnpGatewayNatRuleOutput {
	return o
}

func (o VnpGatewayNatRuleOutput) ToVnpGatewayNatRuleOutputWithContext(ctx context.Context) VnpGatewayNatRuleOutput {
	return o
}

func (o VnpGatewayNatRuleOutput) ToVnpGatewayNatRulePtrOutput() VnpGatewayNatRulePtrOutput {
	return o.ToVnpGatewayNatRulePtrOutputWithContext(context.Background())
}

func (o VnpGatewayNatRuleOutput) ToVnpGatewayNatRulePtrOutputWithContext(ctx context.Context) VnpGatewayNatRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VnpGatewayNatRule) *VnpGatewayNatRule {
		return &v
	}).(VnpGatewayNatRulePtrOutput)
}

type VnpGatewayNatRulePtrOutput struct{ *pulumi.OutputState }

func (VnpGatewayNatRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnpGatewayNatRule)(nil))
}

func (o VnpGatewayNatRulePtrOutput) ToVnpGatewayNatRulePtrOutput() VnpGatewayNatRulePtrOutput {
	return o
}

func (o VnpGatewayNatRulePtrOutput) ToVnpGatewayNatRulePtrOutputWithContext(ctx context.Context) VnpGatewayNatRulePtrOutput {
	return o
}

func (o VnpGatewayNatRulePtrOutput) Elem() VnpGatewayNatRuleOutput {
	return o.ApplyT(func(v *VnpGatewayNatRule) VnpGatewayNatRule {
		if v != nil {
			return *v
		}
		var ret VnpGatewayNatRule
		return ret
	}).(VnpGatewayNatRuleOutput)
}

type VnpGatewayNatRuleArrayOutput struct{ *pulumi.OutputState }

func (VnpGatewayNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnpGatewayNatRule)(nil))
}

func (o VnpGatewayNatRuleArrayOutput) ToVnpGatewayNatRuleArrayOutput() VnpGatewayNatRuleArrayOutput {
	return o
}

func (o VnpGatewayNatRuleArrayOutput) ToVnpGatewayNatRuleArrayOutputWithContext(ctx context.Context) VnpGatewayNatRuleArrayOutput {
	return o
}

func (o VnpGatewayNatRuleArrayOutput) Index(i pulumi.IntInput) VnpGatewayNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VnpGatewayNatRule {
		return vs[0].([]VnpGatewayNatRule)[vs[1].(int)]
	}).(VnpGatewayNatRuleOutput)
}

type VnpGatewayNatRuleMapOutput struct{ *pulumi.OutputState }

func (VnpGatewayNatRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VnpGatewayNatRule)(nil))
}

func (o VnpGatewayNatRuleMapOutput) ToVnpGatewayNatRuleMapOutput() VnpGatewayNatRuleMapOutput {
	return o
}

func (o VnpGatewayNatRuleMapOutput) ToVnpGatewayNatRuleMapOutputWithContext(ctx context.Context) VnpGatewayNatRuleMapOutput {
	return o
}

func (o VnpGatewayNatRuleMapOutput) MapIndex(k pulumi.StringInput) VnpGatewayNatRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VnpGatewayNatRule {
		return vs[0].(map[string]VnpGatewayNatRule)[vs[1].(string)]
	}).(VnpGatewayNatRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VnpGatewayNatRuleInput)(nil)).Elem(), &VnpGatewayNatRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VnpGatewayNatRulePtrInput)(nil)).Elem(), &VnpGatewayNatRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VnpGatewayNatRuleArrayInput)(nil)).Elem(), VnpGatewayNatRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VnpGatewayNatRuleMapInput)(nil)).Elem(), VnpGatewayNatRuleMap{})
	pulumi.RegisterOutputType(VnpGatewayNatRuleOutput{})
	pulumi.RegisterOutputType(VnpGatewayNatRulePtrOutput{})
	pulumi.RegisterOutputType(VnpGatewayNatRuleArrayOutput{})
	pulumi.RegisterOutputType(VnpGatewayNatRuleMapOutput{})
}
