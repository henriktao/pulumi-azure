// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Traffic Manager Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		server, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.AnyMap{
// 				"azi_id": pulumi.Any(1),
// 			},
// 			ByteLength: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTrafficManagerProfile, err := network.NewTrafficManagerProfile(ctx, "exampleTrafficManagerProfile", &network.TrafficManagerProfileArgs{
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			TrafficRoutingMethod: pulumi.String("Weighted"),
// 			DnsConfig: &network.TrafficManagerProfileDnsConfigArgs{
// 				RelativeName: server.Hex,
// 				Ttl:          pulumi.Int(100),
// 			},
// 			MonitorConfig: &network.TrafficManagerProfileMonitorConfigArgs{
// 				Protocol:                  pulumi.String("http"),
// 				Port:                      pulumi.Int(80),
// 				Path:                      pulumi.String("/"),
// 				IntervalInSeconds:         pulumi.Int(30),
// 				TimeoutInSeconds:          pulumi.Int(9),
// 				ToleratedNumberOfFailures: pulumi.Int(3),
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewTrafficManagerEndpoint(ctx, "exampleTrafficManagerEndpoint", &network.TrafficManagerEndpointArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ProfileName:       exampleTrafficManagerProfile.Name,
// 			Type:              pulumi.String("externalEndpoints"),
// 			Weight:            pulumi.Int(100),
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
// Traffic Manager Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:trafficmanager/endpoint:Endpoint exampleEndpoints /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1/azureEndpoints/mytrafficmanagerendpoint
// ```
//
// Deprecated: azure.trafficmanager.Endpoint has been deprecated in favor of azure.network.TrafficManagerEndpoint
type Endpoint struct {
	pulumi.CustomResourceState

	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayOutput `pulumi:"customHeaders"`
	// Specifies the Azure location of the Endpoint,
	// this must be specified for Profiles using the `Performance` routing method
	// if the Endpoint is of either type `nestedEndpoints` or `externalEndpoints`.
	// For Endpoints of type `azureEndpoints` the value will be taken from the
	// location of the Azure target resource.
	EndpointLocation      pulumi.StringOutput `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringOutput `pulumi:"endpointMonitorStatus"`
	// The status of the Endpoint, can be set to
	// either `Enabled` or `Disabled`. Defaults to `Enabled`.
	EndpointStatus pulumi.StringPtrOutput `pulumi:"endpointStatus"`
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
	Subnets EndpointSubnetArrayOutput `pulumi:"subnets"`
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

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
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
	var resource Endpoint
	err := ctx.RegisterResource("azure:trafficmanager/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure:trafficmanager/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []EndpointCustomHeader `pulumi:"customHeaders"`
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
	Subnets []EndpointSubnet `pulumi:"subnets"`
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

type EndpointState struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayInput
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
	Subnets EndpointSubnetArrayInput
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

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders []EndpointCustomHeader `pulumi:"customHeaders"`
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
	Subnets []EndpointSubnet `pulumi:"subnets"`
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

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// One or more `customHeader` blocks as defined below
	CustomHeaders EndpointCustomHeaderArrayInput
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
	Subnets EndpointSubnetArrayInput
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

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

func (i *Endpoint) ToEndpointPtrOutput() EndpointPtrOutput {
	return i.ToEndpointPtrOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPtrOutput)
}

type EndpointPtrInput interface {
	pulumi.Input

	ToEndpointPtrOutput() EndpointPtrOutput
	ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput
}

type endpointPtrType EndpointArgs

func (*endpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil))
}

func (i *endpointPtrType) ToEndpointPtrOutput() EndpointPtrOutput {
	return i.ToEndpointPtrOutputWithContext(context.Background())
}

func (i *endpointPtrType) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPtrOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//          EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//          EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointPtrOutput() EndpointPtrOutput {
	return o.ToEndpointPtrOutputWithContext(context.Background())
}

func (o EndpointOutput) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Endpoint) *Endpoint {
		return &v
	}).(EndpointPtrOutput)
}

type EndpointPtrOutput struct{ *pulumi.OutputState }

func (EndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil))
}

func (o EndpointPtrOutput) ToEndpointPtrOutput() EndpointPtrOutput {
	return o
}

func (o EndpointPtrOutput) ToEndpointPtrOutputWithContext(ctx context.Context) EndpointPtrOutput {
	return o
}

func (o EndpointPtrOutput) Elem() EndpointOutput {
	return o.ApplyT(func(v *Endpoint) Endpoint {
		if v != nil {
			return *v
		}
		var ret Endpoint
		return ret
	}).(EndpointOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoint)(nil))
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Endpoint {
		return vs[0].([]Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Endpoint)(nil))
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Endpoint {
		return vs[0].(map[string]Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPtrInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointPtrOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
