// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a 3rd Generation (v3) App Service Environment.
//
// > **NOTE:** App Service Environment V3 is currently in Preview.
//
// ## Import
//
// A 3rd Generation (v3) App Service Environment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/environmentV3:EnvironmentV3 myAppServiceEnv /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Web/hostingEnvironments/myAppServiceEnv
// ```
type EnvironmentV3 struct {
	pulumi.CustomResourceState

	// Should new Private Endpoint Connections be allowed. Defaults to `true`.
	AllowNewPrivateEndpointConnections pulumi.BoolPtrOutput `pulumi:"allowNewPrivateEndpointConnections"`
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings EnvironmentV3ClusterSettingArrayOutput `pulumi:"clusterSettings"`
	// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
	DedicatedHostCount pulumi.IntPtrOutput `pulumi:"dedicatedHostCount"`
	// the DNS suffix for this App Service Environment V3.
	DnsSuffix pulumi.StringOutput `pulumi:"dnsSuffix"`
	// The external outbound IP addresses of the App Service Environment V3.
	ExternalInboundIpAddresses pulumi.StringArrayOutput `pulumi:"externalInboundIpAddresses"`
	// An Inbound Network Dependencies block as defined below.
	InboundNetworkDependencies EnvironmentV3InboundNetworkDependencyArrayOutput `pulumi:"inboundNetworkDependencies"`
	// The internal outbound IP addresses of the App Service Environment V3.
	InternalInboundIpAddresses pulumi.StringArrayOutput `pulumi:"internalInboundIpAddresses"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrOutput `pulumi:"internalLoadBalancingMode"`
	// The number of IP SSL addresses reserved for the App Service Environment V3.
	IpSslAddressCount pulumi.IntOutput `pulumi:"ipSslAddressCount"`
	// Outbound addresses of Linux based Apps in this App Service Environment V3
	LinuxOutboundIpAddresses pulumi.StringArrayOutput `pulumi:"linuxOutboundIpAddresses"`
	// The location where the App Service Environment exists.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pricing tier for the front end instances.
	PricingTier pulumi.StringOutput `pulumi:"pricingTier"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Outbound addresses of Windows based Apps in this App Service Environment V3.
	WindowsOutboundIpAddresses pulumi.StringArrayOutput `pulumi:"windowsOutboundIpAddresses"`
	ZoneRedundant              pulumi.BoolPtrOutput     `pulumi:"zoneRedundant"`
}

// NewEnvironmentV3 registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentV3(ctx *pulumi.Context,
	name string, args *EnvironmentV3Args, opts ...pulumi.ResourceOption) (*EnvironmentV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource EnvironmentV3
	err := ctx.RegisterResource("azure:appservice/environmentV3:EnvironmentV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentV3 gets an existing EnvironmentV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentV3State, opts ...pulumi.ResourceOption) (*EnvironmentV3, error) {
	var resource EnvironmentV3
	err := ctx.ReadResource("azure:appservice/environmentV3:EnvironmentV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentV3 resources.
type environmentV3State struct {
	// Should new Private Endpoint Connections be allowed. Defaults to `true`.
	AllowNewPrivateEndpointConnections *bool `pulumi:"allowNewPrivateEndpointConnections"`
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings []EnvironmentV3ClusterSetting `pulumi:"clusterSettings"`
	// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
	DedicatedHostCount *int `pulumi:"dedicatedHostCount"`
	// the DNS suffix for this App Service Environment V3.
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// The external outbound IP addresses of the App Service Environment V3.
	ExternalInboundIpAddresses []string `pulumi:"externalInboundIpAddresses"`
	// An Inbound Network Dependencies block as defined below.
	InboundNetworkDependencies []EnvironmentV3InboundNetworkDependency `pulumi:"inboundNetworkDependencies"`
	// The internal outbound IP addresses of the App Service Environment V3.
	InternalInboundIpAddresses []string `pulumi:"internalInboundIpAddresses"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// The number of IP SSL addresses reserved for the App Service Environment V3.
	IpSslAddressCount *int `pulumi:"ipSslAddressCount"`
	// Outbound addresses of Linux based Apps in this App Service Environment V3
	LinuxOutboundIpAddresses []string `pulumi:"linuxOutboundIpAddresses"`
	// The location where the App Service Environment exists.
	Location *string `pulumi:"location"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Pricing tier for the front end instances.
	PricingTier *string `pulumi:"pricingTier"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
	// Outbound addresses of Windows based Apps in this App Service Environment V3.
	WindowsOutboundIpAddresses []string `pulumi:"windowsOutboundIpAddresses"`
	ZoneRedundant              *bool    `pulumi:"zoneRedundant"`
}

type EnvironmentV3State struct {
	// Should new Private Endpoint Connections be allowed. Defaults to `true`.
	AllowNewPrivateEndpointConnections pulumi.BoolPtrInput
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings EnvironmentV3ClusterSettingArrayInput
	// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
	DedicatedHostCount pulumi.IntPtrInput
	// the DNS suffix for this App Service Environment V3.
	DnsSuffix pulumi.StringPtrInput
	// The external outbound IP addresses of the App Service Environment V3.
	ExternalInboundIpAddresses pulumi.StringArrayInput
	// An Inbound Network Dependencies block as defined below.
	InboundNetworkDependencies EnvironmentV3InboundNetworkDependencyArrayInput
	// The internal outbound IP addresses of the App Service Environment V3.
	InternalInboundIpAddresses pulumi.StringArrayInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrInput
	// The number of IP SSL addresses reserved for the App Service Environment V3.
	IpSslAddressCount pulumi.IntPtrInput
	// Outbound addresses of Linux based Apps in this App Service Environment V3
	LinuxOutboundIpAddresses pulumi.StringArrayInput
	// The location where the App Service Environment exists.
	Location pulumi.StringPtrInput
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Pricing tier for the front end instances.
	PricingTier pulumi.StringPtrInput
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
	// Outbound addresses of Windows based Apps in this App Service Environment V3.
	WindowsOutboundIpAddresses pulumi.StringArrayInput
	ZoneRedundant              pulumi.BoolPtrInput
}

func (EnvironmentV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentV3State)(nil)).Elem()
}

type environmentV3Args struct {
	// Should new Private Endpoint Connections be allowed. Defaults to `true`.
	AllowNewPrivateEndpointConnections *bool `pulumi:"allowNewPrivateEndpointConnections"`
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings []EnvironmentV3ClusterSetting `pulumi:"clusterSettings"`
	// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
	DedicatedHostCount *int `pulumi:"dedicatedHostCount"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags          map[string]string `pulumi:"tags"`
	ZoneRedundant *bool             `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a EnvironmentV3 resource.
type EnvironmentV3Args struct {
	// Should new Private Endpoint Connections be allowed. Defaults to `true`.
	AllowNewPrivateEndpointConnections pulumi.BoolPtrInput
	// Zero or more `clusterSetting` blocks as defined below.
	ClusterSettings EnvironmentV3ClusterSettingArrayInput
	// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
	DedicatedHostCount pulumi.IntPtrInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
	InternalLoadBalancingMode pulumi.StringPtrInput
	// The name of the App Service Environment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnetId`).
	ResourceGroupName pulumi.StringInput
	// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags          pulumi.StringMapInput
	ZoneRedundant pulumi.BoolPtrInput
}

func (EnvironmentV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentV3Args)(nil)).Elem()
}

type EnvironmentV3Input interface {
	pulumi.Input

	ToEnvironmentV3Output() EnvironmentV3Output
	ToEnvironmentV3OutputWithContext(ctx context.Context) EnvironmentV3Output
}

func (*EnvironmentV3) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentV3)(nil))
}

func (i *EnvironmentV3) ToEnvironmentV3Output() EnvironmentV3Output {
	return i.ToEnvironmentV3OutputWithContext(context.Background())
}

func (i *EnvironmentV3) ToEnvironmentV3OutputWithContext(ctx context.Context) EnvironmentV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentV3Output)
}

func (i *EnvironmentV3) ToEnvironmentV3PtrOutput() EnvironmentV3PtrOutput {
	return i.ToEnvironmentV3PtrOutputWithContext(context.Background())
}

func (i *EnvironmentV3) ToEnvironmentV3PtrOutputWithContext(ctx context.Context) EnvironmentV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentV3PtrOutput)
}

type EnvironmentV3PtrInput interface {
	pulumi.Input

	ToEnvironmentV3PtrOutput() EnvironmentV3PtrOutput
	ToEnvironmentV3PtrOutputWithContext(ctx context.Context) EnvironmentV3PtrOutput
}

type environmentV3PtrType EnvironmentV3Args

func (*environmentV3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentV3)(nil))
}

func (i *environmentV3PtrType) ToEnvironmentV3PtrOutput() EnvironmentV3PtrOutput {
	return i.ToEnvironmentV3PtrOutputWithContext(context.Background())
}

func (i *environmentV3PtrType) ToEnvironmentV3PtrOutputWithContext(ctx context.Context) EnvironmentV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentV3PtrOutput)
}

// EnvironmentV3ArrayInput is an input type that accepts EnvironmentV3Array and EnvironmentV3ArrayOutput values.
// You can construct a concrete instance of `EnvironmentV3ArrayInput` via:
//
//          EnvironmentV3Array{ EnvironmentV3Args{...} }
type EnvironmentV3ArrayInput interface {
	pulumi.Input

	ToEnvironmentV3ArrayOutput() EnvironmentV3ArrayOutput
	ToEnvironmentV3ArrayOutputWithContext(context.Context) EnvironmentV3ArrayOutput
}

type EnvironmentV3Array []EnvironmentV3Input

func (EnvironmentV3Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EnvironmentV3)(nil))
}

func (i EnvironmentV3Array) ToEnvironmentV3ArrayOutput() EnvironmentV3ArrayOutput {
	return i.ToEnvironmentV3ArrayOutputWithContext(context.Background())
}

func (i EnvironmentV3Array) ToEnvironmentV3ArrayOutputWithContext(ctx context.Context) EnvironmentV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentV3ArrayOutput)
}

// EnvironmentV3MapInput is an input type that accepts EnvironmentV3Map and EnvironmentV3MapOutput values.
// You can construct a concrete instance of `EnvironmentV3MapInput` via:
//
//          EnvironmentV3Map{ "key": EnvironmentV3Args{...} }
type EnvironmentV3MapInput interface {
	pulumi.Input

	ToEnvironmentV3MapOutput() EnvironmentV3MapOutput
	ToEnvironmentV3MapOutputWithContext(context.Context) EnvironmentV3MapOutput
}

type EnvironmentV3Map map[string]EnvironmentV3Input

func (EnvironmentV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EnvironmentV3)(nil))
}

func (i EnvironmentV3Map) ToEnvironmentV3MapOutput() EnvironmentV3MapOutput {
	return i.ToEnvironmentV3MapOutputWithContext(context.Background())
}

func (i EnvironmentV3Map) ToEnvironmentV3MapOutputWithContext(ctx context.Context) EnvironmentV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentV3MapOutput)
}

type EnvironmentV3Output struct {
	*pulumi.OutputState
}

func (EnvironmentV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentV3)(nil))
}

func (o EnvironmentV3Output) ToEnvironmentV3Output() EnvironmentV3Output {
	return o
}

func (o EnvironmentV3Output) ToEnvironmentV3OutputWithContext(ctx context.Context) EnvironmentV3Output {
	return o
}

func (o EnvironmentV3Output) ToEnvironmentV3PtrOutput() EnvironmentV3PtrOutput {
	return o.ToEnvironmentV3PtrOutputWithContext(context.Background())
}

func (o EnvironmentV3Output) ToEnvironmentV3PtrOutputWithContext(ctx context.Context) EnvironmentV3PtrOutput {
	return o.ApplyT(func(v EnvironmentV3) *EnvironmentV3 {
		return &v
	}).(EnvironmentV3PtrOutput)
}

type EnvironmentV3PtrOutput struct {
	*pulumi.OutputState
}

func (EnvironmentV3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentV3)(nil))
}

func (o EnvironmentV3PtrOutput) ToEnvironmentV3PtrOutput() EnvironmentV3PtrOutput {
	return o
}

func (o EnvironmentV3PtrOutput) ToEnvironmentV3PtrOutputWithContext(ctx context.Context) EnvironmentV3PtrOutput {
	return o
}

type EnvironmentV3ArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentV3)(nil))
}

func (o EnvironmentV3ArrayOutput) ToEnvironmentV3ArrayOutput() EnvironmentV3ArrayOutput {
	return o
}

func (o EnvironmentV3ArrayOutput) ToEnvironmentV3ArrayOutputWithContext(ctx context.Context) EnvironmentV3ArrayOutput {
	return o
}

func (o EnvironmentV3ArrayOutput) Index(i pulumi.IntInput) EnvironmentV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentV3 {
		return vs[0].([]EnvironmentV3)[vs[1].(int)]
	}).(EnvironmentV3Output)
}

type EnvironmentV3MapOutput struct{ *pulumi.OutputState }

func (EnvironmentV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EnvironmentV3)(nil))
}

func (o EnvironmentV3MapOutput) ToEnvironmentV3MapOutput() EnvironmentV3MapOutput {
	return o
}

func (o EnvironmentV3MapOutput) ToEnvironmentV3MapOutputWithContext(ctx context.Context) EnvironmentV3MapOutput {
	return o
}

func (o EnvironmentV3MapOutput) MapIndex(k pulumi.StringInput) EnvironmentV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EnvironmentV3 {
		return vs[0].(map[string]EnvironmentV3)[vs[1].(string)]
	}).(EnvironmentV3Output)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentV3Output{})
	pulumi.RegisterOutputType(EnvironmentV3PtrOutput{})
	pulumi.RegisterOutputType(EnvironmentV3ArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentV3MapOutput{})
}
