// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Public IP Address.
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
// 		_, err = network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			AllocationMethod:  pulumi.String("Static"),
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
// Public IPs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/publicIp:PublicIp myPublicIp /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1
// ```
type PublicIp struct {
	pulumi.CustomResourceState

	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringOutput `pulumi:"allocationMethod"`
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrOutput `pulumi:"domainNameLabel"`
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The IP address value that was allocated.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// A mapping of IP tags to assign to the public IP.
	IpTags pulumi.StringMapOutput `pulumi:"ipTags"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrOutput `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrOutput `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// The SKU Tier that should be used for the Public IP. Possible values are `Regional` and `Global`. Defaults to `Regional`.
	SkuTier pulumi.StringPtrOutput `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringOutput `pulumi:"zones"`
}

// NewPublicIp registers a new resource with the given unique name, arguments, and options.
func NewPublicIp(ctx *pulumi.Context,
	name string, args *PublicIpArgs, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationMethod == nil {
		return nil, errors.New("invalid value for required argument 'AllocationMethod'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PublicIp
	err := ctx.RegisterResource("azure:network/publicIp:PublicIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIp gets an existing PublicIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpState, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	var resource PublicIp
	err := ctx.ReadResource("azure:network/publicIp:PublicIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIp resources.
type publicIpState struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod *string `pulumi:"allocationMethod"`
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn *string `pulumi:"fqdn"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The IP address value that was allocated.
	IpAddress *string `pulumi:"ipAddress"`
	// A mapping of IP tags to assign to the public IP.
	IpTags map[string]string `pulumi:"ipTags"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId *string `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// The SKU Tier that should be used for the Public IP. Possible values are `Regional` and `Global`. Defaults to `Regional`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones *string `pulumi:"zones"`
}

type PublicIpState struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringPtrInput
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringPtrInput
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrInput
	// Fully qualified domain name of the A DNS record associated with the public IP. `domainNameLabel` must be specified to get the `fqdn`. This is the concatenation of the `domainNameLabel` and the regionalized DNS zone
	Fqdn pulumi.StringPtrInput
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The IP address value that was allocated.
	IpAddress pulumi.StringPtrInput
	// A mapping of IP tags to assign to the public IP.
	IpTags pulumi.StringMapInput
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringPtrInput
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrInput
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// The SKU Tier that should be used for the Public IP. Possible values are `Regional` and `Global`. Defaults to `Regional`.
	SkuTier pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringPtrInput
}

func (PublicIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpState)(nil)).Elem()
}

type publicIpArgs struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod string `pulumi:"allocationMethod"`
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// A mapping of IP tags to assign to the public IP.
	IpTags map[string]string `pulumi:"ipTags"`
	// The IP Version to use, IPv6 or IPv4.
	IpVersion *string `pulumi:"ipVersion"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId *string `pulumi:"publicIpPrefixId"`
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `pulumi:"reverseFqdn"`
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// The SKU Tier that should be used for the Public IP. Possible values are `Regional` and `Global`. Defaults to `Regional`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIp resource.
type PublicIpArgs struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod pulumi.StringInput
	// The availability zone to allocate the Public IP in. Possible values are `Zone-Redundant`, `1`, `2`, `3`, and `No-Zone`. Defaults to `Zone-Redundant`.
	AvailabilityZone pulumi.StringPtrInput
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel pulumi.StringPtrInput
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// A mapping of IP tags to assign to the public IP.
	IpTags pulumi.StringMapInput
	// The IP Version to use, IPv6 or IPv4.
	IpVersion pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName pulumi.StringInput
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn pulumi.StringPtrInput
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// The SKU Tier that should be used for the Public IP. Possible values are `Regional` and `Global`. Defaults to `Regional`.
	SkuTier pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: This property has been deprecated in favour of `availability_zone` due to a breaking behavioural change in Azure: https://azure.microsoft.com/en-us/updates/zone-behavior-change/
	Zones pulumi.StringPtrInput
}

func (PublicIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpArgs)(nil)).Elem()
}

type PublicIpInput interface {
	pulumi.Input

	ToPublicIpOutput() PublicIpOutput
	ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput
}

func (*PublicIp) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIp)(nil))
}

func (i *PublicIp) ToPublicIpOutput() PublicIpOutput {
	return i.ToPublicIpOutputWithContext(context.Background())
}

func (i *PublicIp) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpOutput)
}

func (i *PublicIp) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return i.ToPublicIpPtrOutputWithContext(context.Background())
}

func (i *PublicIp) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPtrOutput)
}

type PublicIpPtrInput interface {
	pulumi.Input

	ToPublicIpPtrOutput() PublicIpPtrOutput
	ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput
}

type publicIpPtrType PublicIpArgs

func (*publicIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil))
}

func (i *publicIpPtrType) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return i.ToPublicIpPtrOutputWithContext(context.Background())
}

func (i *publicIpPtrType) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPtrOutput)
}

// PublicIpArrayInput is an input type that accepts PublicIpArray and PublicIpArrayOutput values.
// You can construct a concrete instance of `PublicIpArrayInput` via:
//
//          PublicIpArray{ PublicIpArgs{...} }
type PublicIpArrayInput interface {
	pulumi.Input

	ToPublicIpArrayOutput() PublicIpArrayOutput
	ToPublicIpArrayOutputWithContext(context.Context) PublicIpArrayOutput
}

type PublicIpArray []PublicIpInput

func (PublicIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIp)(nil)).Elem()
}

func (i PublicIpArray) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return i.ToPublicIpArrayOutputWithContext(context.Background())
}

func (i PublicIpArray) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpArrayOutput)
}

// PublicIpMapInput is an input type that accepts PublicIpMap and PublicIpMapOutput values.
// You can construct a concrete instance of `PublicIpMapInput` via:
//
//          PublicIpMap{ "key": PublicIpArgs{...} }
type PublicIpMapInput interface {
	pulumi.Input

	ToPublicIpMapOutput() PublicIpMapOutput
	ToPublicIpMapOutputWithContext(context.Context) PublicIpMapOutput
}

type PublicIpMap map[string]PublicIpInput

func (PublicIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIp)(nil)).Elem()
}

func (i PublicIpMap) ToPublicIpMapOutput() PublicIpMapOutput {
	return i.ToPublicIpMapOutputWithContext(context.Background())
}

func (i PublicIpMap) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpMapOutput)
}

type PublicIpOutput struct{ *pulumi.OutputState }

func (PublicIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIp)(nil))
}

func (o PublicIpOutput) ToPublicIpOutput() PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return o.ToPublicIpPtrOutputWithContext(context.Background())
}

func (o PublicIpOutput) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIp) *PublicIp {
		return &v
	}).(PublicIpPtrOutput)
}

type PublicIpPtrOutput struct{ *pulumi.OutputState }

func (PublicIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil))
}

func (o PublicIpPtrOutput) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return o
}

func (o PublicIpPtrOutput) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return o
}

func (o PublicIpPtrOutput) Elem() PublicIpOutput {
	return o.ApplyT(func(v *PublicIp) PublicIp {
		if v != nil {
			return *v
		}
		var ret PublicIp
		return ret
	}).(PublicIpOutput)
}

type PublicIpArrayOutput struct{ *pulumi.OutputState }

func (PublicIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PublicIp)(nil))
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) Index(i pulumi.IntInput) PublicIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PublicIp {
		return vs[0].([]PublicIp)[vs[1].(int)]
	}).(PublicIpOutput)
}

type PublicIpMapOutput struct{ *pulumi.OutputState }

func (PublicIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PublicIp)(nil))
}

func (o PublicIpMapOutput) ToPublicIpMapOutput() PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) MapIndex(k pulumi.StringInput) PublicIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PublicIp {
		return vs[0].(map[string]PublicIp)[vs[1].(string)]
	}).(PublicIpOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIpOutput{})
	pulumi.RegisterOutputType(PublicIpPtrOutput{})
	pulumi.RegisterOutputType(PublicIpArrayOutput{})
	pulumi.RegisterOutputType(PublicIpMapOutput{})
}
