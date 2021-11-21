// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Public IP Prefix.
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
// 		_, err = network.NewPublicIpPrefix(ctx, "examplePublicIpPrefix", &network.PublicIpPrefixArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PrefixLength:      pulumi.Int(31),
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
// Public IP Prefixes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/publicIpPrefix:PublicIpPrefix myPublicIpPrefix /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPPrefixes/myPublicIpPrefix1
// ```
type PublicIpPrefix struct {
	pulumi.CustomResourceState

	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The IP address prefix value that was allocated.
	IpPrefix pulumi.StringOutput `pulumi:"ipPrefix"`
	// The IP Version to use, `IPv6` or `IPv4`. Changing this forces a new resource to be created. Default is `IPv4`.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringOutput `pulumi:"zones"`
}

// NewPublicIpPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicIpPrefix(ctx *pulumi.Context,
	name string, args *PublicIpPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIpPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PublicIpPrefix
	err := ctx.RegisterResource("azure:network/publicIpPrefix:PublicIpPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIpPrefix gets an existing PublicIpPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIpPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpPrefixState, opts ...pulumi.ResourceOption) (*PublicIpPrefix, error) {
	var resource PublicIpPrefix
	err := ctx.ReadResource("azure:network/publicIpPrefix:PublicIpPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIpPrefix resources.
type publicIpPrefixState struct {
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The IP address prefix value that was allocated.
	IpPrefix *string `pulumi:"ipPrefix"`
	// The IP Version to use, `IPv6` or `IPv4`. Changing this forces a new resource to be created. Default is `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *int `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones *string `pulumi:"zones"`
}

type PublicIpPrefixState struct {
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringPtrInput
	// The IP address prefix value that was allocated.
	IpPrefix pulumi.StringPtrInput
	// The IP Version to use, `IPv6` or `IPv4`. Changing this forces a new resource to be created. Default is `IPv4`.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrInput
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringPtrInput
}

func (PublicIpPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpPrefixState)(nil)).Elem()
}

type publicIpPrefixArgs struct {
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The IP Version to use, `IPv6` or `IPv4`. Changing this forces a new resource to be created. Default is `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *int `pulumi:"prefixLength"`
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIpPrefix resource.
type PublicIpPrefixArgs struct {
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringPtrInput
	// The IP Version to use, `IPv6` or `IPv4`. Changing this forces a new resource to be created. Default is `IPv4`.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP Prefix resource . Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to `28`(16 addresses). Changing this forces a new resource to be created.
	PrefixLength pulumi.IntPtrInput
	// The name of the resource group in which to create the Public IP Prefix.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringPtrInput
}

func (PublicIpPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpPrefixArgs)(nil)).Elem()
}

type PublicIpPrefixInput interface {
	pulumi.Input

	ToPublicIpPrefixOutput() PublicIpPrefixOutput
	ToPublicIpPrefixOutputWithContext(ctx context.Context) PublicIpPrefixOutput
}

func (*PublicIpPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpPrefix)(nil))
}

func (i *PublicIpPrefix) ToPublicIpPrefixOutput() PublicIpPrefixOutput {
	return i.ToPublicIpPrefixOutputWithContext(context.Background())
}

func (i *PublicIpPrefix) ToPublicIpPrefixOutputWithContext(ctx context.Context) PublicIpPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPrefixOutput)
}

func (i *PublicIpPrefix) ToPublicIpPrefixPtrOutput() PublicIpPrefixPtrOutput {
	return i.ToPublicIpPrefixPtrOutputWithContext(context.Background())
}

func (i *PublicIpPrefix) ToPublicIpPrefixPtrOutputWithContext(ctx context.Context) PublicIpPrefixPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPrefixPtrOutput)
}

type PublicIpPrefixPtrInput interface {
	pulumi.Input

	ToPublicIpPrefixPtrOutput() PublicIpPrefixPtrOutput
	ToPublicIpPrefixPtrOutputWithContext(ctx context.Context) PublicIpPrefixPtrOutput
}

type publicIpPrefixPtrType PublicIpPrefixArgs

func (*publicIpPrefixPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpPrefix)(nil))
}

func (i *publicIpPrefixPtrType) ToPublicIpPrefixPtrOutput() PublicIpPrefixPtrOutput {
	return i.ToPublicIpPrefixPtrOutputWithContext(context.Background())
}

func (i *publicIpPrefixPtrType) ToPublicIpPrefixPtrOutputWithContext(ctx context.Context) PublicIpPrefixPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPrefixPtrOutput)
}

// PublicIpPrefixArrayInput is an input type that accepts PublicIpPrefixArray and PublicIpPrefixArrayOutput values.
// You can construct a concrete instance of `PublicIpPrefixArrayInput` via:
//
//          PublicIpPrefixArray{ PublicIpPrefixArgs{...} }
type PublicIpPrefixArrayInput interface {
	pulumi.Input

	ToPublicIpPrefixArrayOutput() PublicIpPrefixArrayOutput
	ToPublicIpPrefixArrayOutputWithContext(context.Context) PublicIpPrefixArrayOutput
}

type PublicIpPrefixArray []PublicIpPrefixInput

func (PublicIpPrefixArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIpPrefix)(nil)).Elem()
}

func (i PublicIpPrefixArray) ToPublicIpPrefixArrayOutput() PublicIpPrefixArrayOutput {
	return i.ToPublicIpPrefixArrayOutputWithContext(context.Background())
}

func (i PublicIpPrefixArray) ToPublicIpPrefixArrayOutputWithContext(ctx context.Context) PublicIpPrefixArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPrefixArrayOutput)
}

// PublicIpPrefixMapInput is an input type that accepts PublicIpPrefixMap and PublicIpPrefixMapOutput values.
// You can construct a concrete instance of `PublicIpPrefixMapInput` via:
//
//          PublicIpPrefixMap{ "key": PublicIpPrefixArgs{...} }
type PublicIpPrefixMapInput interface {
	pulumi.Input

	ToPublicIpPrefixMapOutput() PublicIpPrefixMapOutput
	ToPublicIpPrefixMapOutputWithContext(context.Context) PublicIpPrefixMapOutput
}

type PublicIpPrefixMap map[string]PublicIpPrefixInput

func (PublicIpPrefixMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIpPrefix)(nil)).Elem()
}

func (i PublicIpPrefixMap) ToPublicIpPrefixMapOutput() PublicIpPrefixMapOutput {
	return i.ToPublicIpPrefixMapOutputWithContext(context.Background())
}

func (i PublicIpPrefixMap) ToPublicIpPrefixMapOutputWithContext(ctx context.Context) PublicIpPrefixMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPrefixMapOutput)
}

type PublicIpPrefixOutput struct{ *pulumi.OutputState }

func (PublicIpPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpPrefix)(nil))
}

func (o PublicIpPrefixOutput) ToPublicIpPrefixOutput() PublicIpPrefixOutput {
	return o
}

func (o PublicIpPrefixOutput) ToPublicIpPrefixOutputWithContext(ctx context.Context) PublicIpPrefixOutput {
	return o
}

func (o PublicIpPrefixOutput) ToPublicIpPrefixPtrOutput() PublicIpPrefixPtrOutput {
	return o.ToPublicIpPrefixPtrOutputWithContext(context.Background())
}

func (o PublicIpPrefixOutput) ToPublicIpPrefixPtrOutputWithContext(ctx context.Context) PublicIpPrefixPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpPrefix) *PublicIpPrefix {
		return &v
	}).(PublicIpPrefixPtrOutput)
}

type PublicIpPrefixPtrOutput struct{ *pulumi.OutputState }

func (PublicIpPrefixPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpPrefix)(nil))
}

func (o PublicIpPrefixPtrOutput) ToPublicIpPrefixPtrOutput() PublicIpPrefixPtrOutput {
	return o
}

func (o PublicIpPrefixPtrOutput) ToPublicIpPrefixPtrOutputWithContext(ctx context.Context) PublicIpPrefixPtrOutput {
	return o
}

func (o PublicIpPrefixPtrOutput) Elem() PublicIpPrefixOutput {
	return o.ApplyT(func(v *PublicIpPrefix) PublicIpPrefix {
		if v != nil {
			return *v
		}
		var ret PublicIpPrefix
		return ret
	}).(PublicIpPrefixOutput)
}

type PublicIpPrefixArrayOutput struct{ *pulumi.OutputState }

func (PublicIpPrefixArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PublicIpPrefix)(nil))
}

func (o PublicIpPrefixArrayOutput) ToPublicIpPrefixArrayOutput() PublicIpPrefixArrayOutput {
	return o
}

func (o PublicIpPrefixArrayOutput) ToPublicIpPrefixArrayOutputWithContext(ctx context.Context) PublicIpPrefixArrayOutput {
	return o
}

func (o PublicIpPrefixArrayOutput) Index(i pulumi.IntInput) PublicIpPrefixOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PublicIpPrefix {
		return vs[0].([]PublicIpPrefix)[vs[1].(int)]
	}).(PublicIpPrefixOutput)
}

type PublicIpPrefixMapOutput struct{ *pulumi.OutputState }

func (PublicIpPrefixMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PublicIpPrefix)(nil))
}

func (o PublicIpPrefixMapOutput) ToPublicIpPrefixMapOutput() PublicIpPrefixMapOutput {
	return o
}

func (o PublicIpPrefixMapOutput) ToPublicIpPrefixMapOutputWithContext(ctx context.Context) PublicIpPrefixMapOutput {
	return o
}

func (o PublicIpPrefixMapOutput) MapIndex(k pulumi.StringInput) PublicIpPrefixOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PublicIpPrefix {
		return vs[0].(map[string]PublicIpPrefix)[vs[1].(string)]
	}).(PublicIpPrefixOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIpPrefixOutput{})
	pulumi.RegisterOutputType(PublicIpPrefixPtrOutput{})
	pulumi.RegisterOutputType(PublicIpPrefixArrayOutput{})
	pulumi.RegisterOutputType(PublicIpPrefixMapOutput{})
}
