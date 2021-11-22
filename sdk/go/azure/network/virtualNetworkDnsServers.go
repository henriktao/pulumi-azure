// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Virtual Network DNS Servers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualNetworkDnsServers:VirtualNetworkDnsServers exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/dnsServers/default
// ```
type VirtualNetworkDnsServers struct {
	pulumi.CustomResourceState

	// List of IP addresses of DNS servers
	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringOutput `pulumi:"virtualNetworkId"`
}

// NewVirtualNetworkDnsServers registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkDnsServers(ctx *pulumi.Context,
	name string, args *VirtualNetworkDnsServersArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkDnsServers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkId'")
	}
	var resource VirtualNetworkDnsServers
	err := ctx.RegisterResource("azure:network/virtualNetworkDnsServers:VirtualNetworkDnsServers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkDnsServers gets an existing VirtualNetworkDnsServers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkDnsServers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkDnsServersState, opts ...pulumi.ResourceOption) (*VirtualNetworkDnsServers, error) {
	var resource VirtualNetworkDnsServers
	err := ctx.ReadResource("azure:network/virtualNetworkDnsServers:VirtualNetworkDnsServers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkDnsServers resources.
type virtualNetworkDnsServersState struct {
	// List of IP addresses of DNS servers
	DnsServers []string `pulumi:"dnsServers"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
}

type VirtualNetworkDnsServersState struct {
	// List of IP addresses of DNS servers
	DnsServers pulumi.StringArrayInput
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringPtrInput
}

func (VirtualNetworkDnsServersState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkDnsServersState)(nil)).Elem()
}

type virtualNetworkDnsServersArgs struct {
	// List of IP addresses of DNS servers
	DnsServers []string `pulumi:"dnsServers"`
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId string `pulumi:"virtualNetworkId"`
}

// The set of arguments for constructing a VirtualNetworkDnsServers resource.
type VirtualNetworkDnsServersArgs struct {
	// List of IP addresses of DNS servers
	DnsServers pulumi.StringArrayInput
	// The ID of the Virtual Network that should be linked to the DNS Zone. Changing this forces a new resource to be created.
	VirtualNetworkId pulumi.StringInput
}

func (VirtualNetworkDnsServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkDnsServersArgs)(nil)).Elem()
}

type VirtualNetworkDnsServersInput interface {
	pulumi.Input

	ToVirtualNetworkDnsServersOutput() VirtualNetworkDnsServersOutput
	ToVirtualNetworkDnsServersOutputWithContext(ctx context.Context) VirtualNetworkDnsServersOutput
}

func (*VirtualNetworkDnsServers) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkDnsServers)(nil))
}

func (i *VirtualNetworkDnsServers) ToVirtualNetworkDnsServersOutput() VirtualNetworkDnsServersOutput {
	return i.ToVirtualNetworkDnsServersOutputWithContext(context.Background())
}

func (i *VirtualNetworkDnsServers) ToVirtualNetworkDnsServersOutputWithContext(ctx context.Context) VirtualNetworkDnsServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkDnsServersOutput)
}

func (i *VirtualNetworkDnsServers) ToVirtualNetworkDnsServersPtrOutput() VirtualNetworkDnsServersPtrOutput {
	return i.ToVirtualNetworkDnsServersPtrOutputWithContext(context.Background())
}

func (i *VirtualNetworkDnsServers) ToVirtualNetworkDnsServersPtrOutputWithContext(ctx context.Context) VirtualNetworkDnsServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkDnsServersPtrOutput)
}

type VirtualNetworkDnsServersPtrInput interface {
	pulumi.Input

	ToVirtualNetworkDnsServersPtrOutput() VirtualNetworkDnsServersPtrOutput
	ToVirtualNetworkDnsServersPtrOutputWithContext(ctx context.Context) VirtualNetworkDnsServersPtrOutput
}

type virtualNetworkDnsServersPtrType VirtualNetworkDnsServersArgs

func (*virtualNetworkDnsServersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkDnsServers)(nil))
}

func (i *virtualNetworkDnsServersPtrType) ToVirtualNetworkDnsServersPtrOutput() VirtualNetworkDnsServersPtrOutput {
	return i.ToVirtualNetworkDnsServersPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkDnsServersPtrType) ToVirtualNetworkDnsServersPtrOutputWithContext(ctx context.Context) VirtualNetworkDnsServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkDnsServersPtrOutput)
}

// VirtualNetworkDnsServersArrayInput is an input type that accepts VirtualNetworkDnsServersArray and VirtualNetworkDnsServersArrayOutput values.
// You can construct a concrete instance of `VirtualNetworkDnsServersArrayInput` via:
//
//          VirtualNetworkDnsServersArray{ VirtualNetworkDnsServersArgs{...} }
type VirtualNetworkDnsServersArrayInput interface {
	pulumi.Input

	ToVirtualNetworkDnsServersArrayOutput() VirtualNetworkDnsServersArrayOutput
	ToVirtualNetworkDnsServersArrayOutputWithContext(context.Context) VirtualNetworkDnsServersArrayOutput
}

type VirtualNetworkDnsServersArray []VirtualNetworkDnsServersInput

func (VirtualNetworkDnsServersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNetworkDnsServers)(nil)).Elem()
}

func (i VirtualNetworkDnsServersArray) ToVirtualNetworkDnsServersArrayOutput() VirtualNetworkDnsServersArrayOutput {
	return i.ToVirtualNetworkDnsServersArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkDnsServersArray) ToVirtualNetworkDnsServersArrayOutputWithContext(ctx context.Context) VirtualNetworkDnsServersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkDnsServersArrayOutput)
}

// VirtualNetworkDnsServersMapInput is an input type that accepts VirtualNetworkDnsServersMap and VirtualNetworkDnsServersMapOutput values.
// You can construct a concrete instance of `VirtualNetworkDnsServersMapInput` via:
//
//          VirtualNetworkDnsServersMap{ "key": VirtualNetworkDnsServersArgs{...} }
type VirtualNetworkDnsServersMapInput interface {
	pulumi.Input

	ToVirtualNetworkDnsServersMapOutput() VirtualNetworkDnsServersMapOutput
	ToVirtualNetworkDnsServersMapOutputWithContext(context.Context) VirtualNetworkDnsServersMapOutput
}

type VirtualNetworkDnsServersMap map[string]VirtualNetworkDnsServersInput

func (VirtualNetworkDnsServersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNetworkDnsServers)(nil)).Elem()
}

func (i VirtualNetworkDnsServersMap) ToVirtualNetworkDnsServersMapOutput() VirtualNetworkDnsServersMapOutput {
	return i.ToVirtualNetworkDnsServersMapOutputWithContext(context.Background())
}

func (i VirtualNetworkDnsServersMap) ToVirtualNetworkDnsServersMapOutputWithContext(ctx context.Context) VirtualNetworkDnsServersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkDnsServersMapOutput)
}

type VirtualNetworkDnsServersOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkDnsServers)(nil))
}

func (o VirtualNetworkDnsServersOutput) ToVirtualNetworkDnsServersOutput() VirtualNetworkDnsServersOutput {
	return o
}

func (o VirtualNetworkDnsServersOutput) ToVirtualNetworkDnsServersOutputWithContext(ctx context.Context) VirtualNetworkDnsServersOutput {
	return o
}

func (o VirtualNetworkDnsServersOutput) ToVirtualNetworkDnsServersPtrOutput() VirtualNetworkDnsServersPtrOutput {
	return o.ToVirtualNetworkDnsServersPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkDnsServersOutput) ToVirtualNetworkDnsServersPtrOutputWithContext(ctx context.Context) VirtualNetworkDnsServersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkDnsServers) *VirtualNetworkDnsServers {
		return &v
	}).(VirtualNetworkDnsServersPtrOutput)
}

type VirtualNetworkDnsServersPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsServersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkDnsServers)(nil))
}

func (o VirtualNetworkDnsServersPtrOutput) ToVirtualNetworkDnsServersPtrOutput() VirtualNetworkDnsServersPtrOutput {
	return o
}

func (o VirtualNetworkDnsServersPtrOutput) ToVirtualNetworkDnsServersPtrOutputWithContext(ctx context.Context) VirtualNetworkDnsServersPtrOutput {
	return o
}

func (o VirtualNetworkDnsServersPtrOutput) Elem() VirtualNetworkDnsServersOutput {
	return o.ApplyT(func(v *VirtualNetworkDnsServers) VirtualNetworkDnsServers {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkDnsServers
		return ret
	}).(VirtualNetworkDnsServersOutput)
}

type VirtualNetworkDnsServersArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsServersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkDnsServers)(nil))
}

func (o VirtualNetworkDnsServersArrayOutput) ToVirtualNetworkDnsServersArrayOutput() VirtualNetworkDnsServersArrayOutput {
	return o
}

func (o VirtualNetworkDnsServersArrayOutput) ToVirtualNetworkDnsServersArrayOutputWithContext(ctx context.Context) VirtualNetworkDnsServersArrayOutput {
	return o
}

func (o VirtualNetworkDnsServersArrayOutput) Index(i pulumi.IntInput) VirtualNetworkDnsServersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkDnsServers {
		return vs[0].([]VirtualNetworkDnsServers)[vs[1].(int)]
	}).(VirtualNetworkDnsServersOutput)
}

type VirtualNetworkDnsServersMapOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsServersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualNetworkDnsServers)(nil))
}

func (o VirtualNetworkDnsServersMapOutput) ToVirtualNetworkDnsServersMapOutput() VirtualNetworkDnsServersMapOutput {
	return o
}

func (o VirtualNetworkDnsServersMapOutput) ToVirtualNetworkDnsServersMapOutputWithContext(ctx context.Context) VirtualNetworkDnsServersMapOutput {
	return o
}

func (o VirtualNetworkDnsServersMapOutput) MapIndex(k pulumi.StringInput) VirtualNetworkDnsServersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualNetworkDnsServers {
		return vs[0].(map[string]VirtualNetworkDnsServers)[vs[1].(string)]
	}).(VirtualNetworkDnsServersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkDnsServersInput)(nil)).Elem(), &VirtualNetworkDnsServers{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkDnsServersPtrInput)(nil)).Elem(), &VirtualNetworkDnsServers{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkDnsServersArrayInput)(nil)).Elem(), VirtualNetworkDnsServersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkDnsServersMapInput)(nil)).Elem(), VirtualNetworkDnsServersMap{})
	pulumi.RegisterOutputType(VirtualNetworkDnsServersOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsServersPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsServersArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsServersMapOutput{})
}
