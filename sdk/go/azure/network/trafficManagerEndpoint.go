// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Traffic Manager Endpoint.
//
// ## Import
//
// Traffic Manager Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/trafficManagerEndpoint:TrafficManagerEndpoint exampleEndpoints /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1/azureEndpoints/mytrafficmanagerendpoint
// ```
type TrafficManagerEndpoint struct {
	pulumi.CustomResourceState

	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeaderArrayOutput `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      pulumi.StringOutput `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringOutput `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringOutput `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayOutput `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and has to be larger than `0`.
	MinChildEndpoints pulumi.IntPtrOutput `pulumi:"minChildEndpoints"`
	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv4 pulumi.IntPtrOutput `pulumi:"minimumRequiredChildEndpointsIpv4"`
	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv6 pulumi.IntPtrOutput `pulumi:"minimumRequiredChildEndpointsIpv6"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetArrayOutput `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringOutput `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewTrafficManagerEndpoint registers a new resource with the given unique name, arguments, and options.
func NewTrafficManagerEndpoint(ctx *pulumi.Context,
	name string, args *TrafficManagerEndpointArgs, opts ...pulumi.ResourceOption) (*TrafficManagerEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:trafficmanager/endpoint:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource TrafficManagerEndpoint
	err := ctx.RegisterResource("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficManagerEndpoint gets an existing TrafficManagerEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficManagerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerEndpointState, opts ...pulumi.ResourceOption) (*TrafficManagerEndpoint, error) {
	var resource TrafficManagerEndpoint
	err := ctx.ReadResource("azure:network/trafficManagerEndpoint:TrafficManagerEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficManagerEndpoint resources.
type trafficManagerEndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []TrafficManagerEndpointCustomHeader `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      *string `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings []string `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and has to be larger than `0`.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv4 *int `pulumi:"minimumRequiredChildEndpointsIpv4"`
	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv6 *int `pulumi:"minimumRequiredChildEndpointsIpv6"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority *int `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName *string `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets []TrafficManagerEndpointSubnet `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target *string `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type *string `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight *int `pulumi:"weight"`
}

type TrafficManagerEndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeaderArrayInput
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      pulumi.StringPtrInput
	EndpointMonitorStatus pulumi.StringPtrInput
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringPtrInput
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and has to be larger than `0`.
	MinChildEndpoints pulumi.IntPtrInput
	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv4 pulumi.IntPtrInput
	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv6 pulumi.IntPtrInput
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntPtrInput
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringPtrInput
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetArrayInput
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringPtrInput
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrInput
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringPtrInput
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntPtrInput
}

func (TrafficManagerEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerEndpointState)(nil)).Elem()
}

type trafficManagerEndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []TrafficManagerEndpointCustomHeader `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings []string `pulumi:"geoMappings"`
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and has to be larger than `0`.
	MinChildEndpoints *int `pulumi:"minChildEndpoints"`
	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv4 *int `pulumi:"minimumRequiredChildEndpointsIpv4"`
	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv6 *int `pulumi:"minimumRequiredChildEndpointsIpv6"`
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority *int `pulumi:"priority"`
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `subnet` blocks as defined below
	Subnets []TrafficManagerEndpointSubnet `pulumi:"subnets"`
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target *string `pulumi:"target"`
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type string `pulumi:"type"`
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a TrafficManagerEndpoint resource.
type TrafficManagerEndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders TrafficManagerEndpointCustomHeaderArrayInput
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation pulumi.StringPtrInput
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringPtrInput
	// A list of Geographic Regions used to distribute traffic, such as `WORLD`, `UK` or `DE`. The same location can't be specified in two endpoints. [See the Geographic Hierarchies documentation for more information](https://docs.microsoft.com/en-us/rest/api/trafficmanager/geographichierarchies/getdefault).
	GeoMappings pulumi.StringArrayInput
	// This argument specifies the minimum number
	// of endpoints that must be ‘online’ in the child profile in order for the
	// parent profile to direct traffic to any of the endpoints in that child
	// profile. This argument only applies to Endpoints of type `nestedEndpoints`
	// and has to be larger than `0`.
	MinChildEndpoints pulumi.IntPtrInput
	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv4 pulumi.IntPtrInput
	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type `nestedEndpoints` and defaults to `1`.
	MinimumRequiredChildEndpointsIpv6 pulumi.IntPtrInput
	// The name of the Traffic Manager endpoint. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of this Endpoint, this must be
	// specified for Profiles using the `Priority` traffic routing method. Supports
	// values between 1 and 1000, with no Endpoints sharing the same value. If
	// omitted the value will be computed in order of creation.
	Priority pulumi.IntPtrInput
	// The name of the Traffic Manager Profile to attach
	// create the Traffic Manager endpoint.
	ProfileName pulumi.StringInput
	// The name of the resource group where the Traffic Manager Profile exists.
	ResourceGroupName pulumi.StringInput
	// One or more `subnet` blocks as defined below
	Subnets TrafficManagerEndpointSubnetArrayInput
	// The FQDN DNS name of the target. This argument must be
	// provided for an endpoint of type `externalEndpoints`, for other types it
	// will be computed.
	Target pulumi.StringPtrInput
	// The resource id of an Azure resource to
	// target. This argument must be provided for an endpoint of type
	// `azureEndpoints` or `nestedEndpoints`.
	TargetResourceId pulumi.StringPtrInput
	// The Endpoint type, must be one of:
	// - `azureEndpoints`
	// - `externalEndpoints`
	// - `nestedEndpoints`
	Type pulumi.StringInput
	// Specifies how much traffic should be distributed to this
	// endpoint, this must be specified for Profiles using the  `Weighted` traffic
	// routing method. Supports values between 1 and 1000.
	Weight pulumi.IntPtrInput
}

func (TrafficManagerEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerEndpointArgs)(nil)).Elem()
}

type TrafficManagerEndpointInput interface {
	pulumi.Input

	ToTrafficManagerEndpointOutput() TrafficManagerEndpointOutput
	ToTrafficManagerEndpointOutputWithContext(ctx context.Context) TrafficManagerEndpointOutput
}

func (*TrafficManagerEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerEndpoint)(nil))
}

func (i *TrafficManagerEndpoint) ToTrafficManagerEndpointOutput() TrafficManagerEndpointOutput {
	return i.ToTrafficManagerEndpointOutputWithContext(context.Background())
}

func (i *TrafficManagerEndpoint) ToTrafficManagerEndpointOutputWithContext(ctx context.Context) TrafficManagerEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerEndpointOutput)
}

func (i *TrafficManagerEndpoint) ToTrafficManagerEndpointPtrOutput() TrafficManagerEndpointPtrOutput {
	return i.ToTrafficManagerEndpointPtrOutputWithContext(context.Background())
}

func (i *TrafficManagerEndpoint) ToTrafficManagerEndpointPtrOutputWithContext(ctx context.Context) TrafficManagerEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerEndpointPtrOutput)
}

type TrafficManagerEndpointPtrInput interface {
	pulumi.Input

	ToTrafficManagerEndpointPtrOutput() TrafficManagerEndpointPtrOutput
	ToTrafficManagerEndpointPtrOutputWithContext(ctx context.Context) TrafficManagerEndpointPtrOutput
}

type trafficManagerEndpointPtrType TrafficManagerEndpointArgs

func (*trafficManagerEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerEndpoint)(nil))
}

func (i *trafficManagerEndpointPtrType) ToTrafficManagerEndpointPtrOutput() TrafficManagerEndpointPtrOutput {
	return i.ToTrafficManagerEndpointPtrOutputWithContext(context.Background())
}

func (i *trafficManagerEndpointPtrType) ToTrafficManagerEndpointPtrOutputWithContext(ctx context.Context) TrafficManagerEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerEndpointPtrOutput)
}

// TrafficManagerEndpointArrayInput is an input type that accepts TrafficManagerEndpointArray and TrafficManagerEndpointArrayOutput values.
// You can construct a concrete instance of `TrafficManagerEndpointArrayInput` via:
//
//          TrafficManagerEndpointArray{ TrafficManagerEndpointArgs{...} }
type TrafficManagerEndpointArrayInput interface {
	pulumi.Input

	ToTrafficManagerEndpointArrayOutput() TrafficManagerEndpointArrayOutput
	ToTrafficManagerEndpointArrayOutputWithContext(context.Context) TrafficManagerEndpointArrayOutput
}

type TrafficManagerEndpointArray []TrafficManagerEndpointInput

func (TrafficManagerEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficManagerEndpoint)(nil)).Elem()
}

func (i TrafficManagerEndpointArray) ToTrafficManagerEndpointArrayOutput() TrafficManagerEndpointArrayOutput {
	return i.ToTrafficManagerEndpointArrayOutputWithContext(context.Background())
}

func (i TrafficManagerEndpointArray) ToTrafficManagerEndpointArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerEndpointArrayOutput)
}

// TrafficManagerEndpointMapInput is an input type that accepts TrafficManagerEndpointMap and TrafficManagerEndpointMapOutput values.
// You can construct a concrete instance of `TrafficManagerEndpointMapInput` via:
//
//          TrafficManagerEndpointMap{ "key": TrafficManagerEndpointArgs{...} }
type TrafficManagerEndpointMapInput interface {
	pulumi.Input

	ToTrafficManagerEndpointMapOutput() TrafficManagerEndpointMapOutput
	ToTrafficManagerEndpointMapOutputWithContext(context.Context) TrafficManagerEndpointMapOutput
}

type TrafficManagerEndpointMap map[string]TrafficManagerEndpointInput

func (TrafficManagerEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficManagerEndpoint)(nil)).Elem()
}

func (i TrafficManagerEndpointMap) ToTrafficManagerEndpointMapOutput() TrafficManagerEndpointMapOutput {
	return i.ToTrafficManagerEndpointMapOutputWithContext(context.Background())
}

func (i TrafficManagerEndpointMap) ToTrafficManagerEndpointMapOutputWithContext(ctx context.Context) TrafficManagerEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerEndpointMapOutput)
}

type TrafficManagerEndpointOutput struct{ *pulumi.OutputState }

func (TrafficManagerEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerEndpoint)(nil))
}

func (o TrafficManagerEndpointOutput) ToTrafficManagerEndpointOutput() TrafficManagerEndpointOutput {
	return o
}

func (o TrafficManagerEndpointOutput) ToTrafficManagerEndpointOutputWithContext(ctx context.Context) TrafficManagerEndpointOutput {
	return o
}

func (o TrafficManagerEndpointOutput) ToTrafficManagerEndpointPtrOutput() TrafficManagerEndpointPtrOutput {
	return o.ToTrafficManagerEndpointPtrOutputWithContext(context.Background())
}

func (o TrafficManagerEndpointOutput) ToTrafficManagerEndpointPtrOutputWithContext(ctx context.Context) TrafficManagerEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrafficManagerEndpoint) *TrafficManagerEndpoint {
		return &v
	}).(TrafficManagerEndpointPtrOutput)
}

type TrafficManagerEndpointPtrOutput struct{ *pulumi.OutputState }

func (TrafficManagerEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerEndpoint)(nil))
}

func (o TrafficManagerEndpointPtrOutput) ToTrafficManagerEndpointPtrOutput() TrafficManagerEndpointPtrOutput {
	return o
}

func (o TrafficManagerEndpointPtrOutput) ToTrafficManagerEndpointPtrOutputWithContext(ctx context.Context) TrafficManagerEndpointPtrOutput {
	return o
}

func (o TrafficManagerEndpointPtrOutput) Elem() TrafficManagerEndpointOutput {
	return o.ApplyT(func(v *TrafficManagerEndpoint) TrafficManagerEndpoint {
		if v != nil {
			return *v
		}
		var ret TrafficManagerEndpoint
		return ret
	}).(TrafficManagerEndpointOutput)
}

type TrafficManagerEndpointArrayOutput struct{ *pulumi.OutputState }

func (TrafficManagerEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficManagerEndpoint)(nil))
}

func (o TrafficManagerEndpointArrayOutput) ToTrafficManagerEndpointArrayOutput() TrafficManagerEndpointArrayOutput {
	return o
}

func (o TrafficManagerEndpointArrayOutput) ToTrafficManagerEndpointArrayOutputWithContext(ctx context.Context) TrafficManagerEndpointArrayOutput {
	return o
}

func (o TrafficManagerEndpointArrayOutput) Index(i pulumi.IntInput) TrafficManagerEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficManagerEndpoint {
		return vs[0].([]TrafficManagerEndpoint)[vs[1].(int)]
	}).(TrafficManagerEndpointOutput)
}

type TrafficManagerEndpointMapOutput struct{ *pulumi.OutputState }

func (TrafficManagerEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TrafficManagerEndpoint)(nil))
}

func (o TrafficManagerEndpointMapOutput) ToTrafficManagerEndpointMapOutput() TrafficManagerEndpointMapOutput {
	return o
}

func (o TrafficManagerEndpointMapOutput) ToTrafficManagerEndpointMapOutputWithContext(ctx context.Context) TrafficManagerEndpointMapOutput {
	return o
}

func (o TrafficManagerEndpointMapOutput) MapIndex(k pulumi.StringInput) TrafficManagerEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TrafficManagerEndpoint {
		return vs[0].(map[string]TrafficManagerEndpoint)[vs[1].(string)]
	}).(TrafficManagerEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficManagerEndpointOutput{})
	pulumi.RegisterOutputType(TrafficManagerEndpointPtrOutput{})
	pulumi.RegisterOutputType(TrafficManagerEndpointArrayOutput{})
	pulumi.RegisterOutputType(TrafficManagerEndpointMapOutput{})
}
