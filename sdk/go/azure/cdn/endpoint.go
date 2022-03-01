// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cdn"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleProfile, err := cdn.NewProfile(ctx, "exampleProfile", &cdn.ProfileArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard_Verizon"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cdn.NewEndpoint(ctx, "exampleEndpoint", &cdn.EndpointArgs{
// 			ProfileName:       exampleProfile.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Origins: cdn.EndpointOriginArray{
// 				&cdn.EndpointOriginArgs{
// 					Name:     pulumi.String("example"),
// 					HostName: pulumi.String("www.contoso.com"),
// 				},
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
// CDN Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cdn/endpoint:Endpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses pulumi.StringArrayOutput `pulumi:"contentTypesToCompresses"`
	// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
	DeliveryRules EndpointDeliveryRuleArrayOutput `pulumi:"deliveryRules"`
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters EndpointGeoFilterArrayOutput `pulumi:"geoFilters"`
	// Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
	GlobalDeliveryRule EndpointGlobalDeliveryRulePtrOutput `pulumi:"globalDeliveryRule"`
	// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Indicates whether compression is to be enabled.
	IsCompressionEnabled pulumi.BoolPtrOutput `pulumi:"isCompressionEnabled"`
	// Defaults to `true`.
	IsHttpAllowed pulumi.BoolPtrOutput `pulumi:"isHttpAllowed"`
	// Defaults to `true`.
	IsHttpsAllowed pulumi.BoolPtrOutput `pulumi:"isHttpsAllowed"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType pulumi.StringPtrOutput `pulumi:"optimizationType"`
	// The host header CDN provider will send along with content requests to origins.
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// The path used at for origin requests.
	OriginPath pulumi.StringOutput `pulumi:"originPath"`
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins EndpointOriginArrayOutput `pulumi:"origins"`
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath pulumi.StringOutput `pulumi:"probePath"`
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour pulumi.StringPtrOutput `pulumi:"querystringCachingBehaviour"`
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Origins == nil {
		return nil, errors.New("invalid value for required argument 'Origins'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("azure:cdn/endpoint:Endpoint", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:cdn/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses []string `pulumi:"contentTypesToCompresses"`
	// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
	DeliveryRules []EndpointDeliveryRule `pulumi:"deliveryRules"`
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters []EndpointGeoFilter `pulumi:"geoFilters"`
	// Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
	GlobalDeliveryRule *EndpointGlobalDeliveryRule `pulumi:"globalDeliveryRule"`
	// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
	HostName *string `pulumi:"hostName"`
	// Indicates whether compression is to be enabled.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Defaults to `true`.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Defaults to `true`.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType *string `pulumi:"optimizationType"`
	// The host header CDN provider will send along with content requests to origins.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// The path used at for origin requests.
	OriginPath *string `pulumi:"originPath"`
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins []EndpointOrigin `pulumi:"origins"`
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath *string `pulumi:"probePath"`
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName *string `pulumi:"profileName"`
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour *string `pulumi:"querystringCachingBehaviour"`
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EndpointState struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses pulumi.StringArrayInput
	// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
	DeliveryRules EndpointDeliveryRuleArrayInput
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters EndpointGeoFilterArrayInput
	// Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
	GlobalDeliveryRule EndpointGlobalDeliveryRulePtrInput
	// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
	HostName pulumi.StringPtrInput
	// Indicates whether compression is to be enabled.
	IsCompressionEnabled pulumi.BoolPtrInput
	// Defaults to `true`.
	IsHttpAllowed pulumi.BoolPtrInput
	// Defaults to `true`.
	IsHttpsAllowed pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType pulumi.StringPtrInput
	// The host header CDN provider will send along with content requests to origins.
	OriginHostHeader pulumi.StringPtrInput
	// The path used at for origin requests.
	OriginPath pulumi.StringPtrInput
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins EndpointOriginArrayInput
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath pulumi.StringPtrInput
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName pulumi.StringPtrInput
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour pulumi.StringPtrInput
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses []string `pulumi:"contentTypesToCompresses"`
	// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
	DeliveryRules []EndpointDeliveryRule `pulumi:"deliveryRules"`
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters []EndpointGeoFilter `pulumi:"geoFilters"`
	// Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
	GlobalDeliveryRule *EndpointGlobalDeliveryRule `pulumi:"globalDeliveryRule"`
	// Indicates whether compression is to be enabled.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Defaults to `true`.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Defaults to `true`.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType *string `pulumi:"optimizationType"`
	// The host header CDN provider will send along with content requests to origins.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// The path used at for origin requests.
	OriginPath *string `pulumi:"originPath"`
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins []EndpointOrigin `pulumi:"origins"`
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath *string `pulumi:"probePath"`
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName string `pulumi:"profileName"`
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour *string `pulumi:"querystringCachingBehaviour"`
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// An array of strings that indicates a content types on which compression will be applied. The value for the elements should be MIME types.
	ContentTypesToCompresses pulumi.StringArrayInput
	// Rules for the rules engine. An endpoint can contain up until 4 of those rules that consist of conditions and actions. A `deliveryRule` blocks as defined below.
	DeliveryRules EndpointDeliveryRuleArrayInput
	// A set of Geo Filters for this CDN Endpoint. Each `geoFilter` block supports fields documented below.
	GeoFilters EndpointGeoFilterArrayInput
	// Actions that are valid for all resources regardless of any conditions. A `globalDeliveryRule` block as defined below.
	GlobalDeliveryRule EndpointGlobalDeliveryRulePtrInput
	// Indicates whether compression is to be enabled.
	IsCompressionEnabled pulumi.BoolPtrInput
	// Defaults to `true`.
	IsHttpAllowed pulumi.BoolPtrInput
	// Defaults to `true`.
	IsHttpsAllowed pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// What types of optimization should this CDN Endpoint optimize for? Possible values include `DynamicSiteAcceleration`, `GeneralMediaStreaming`, `GeneralWebDelivery`, `LargeFileDownload` and `VideoOnDemandMediaStreaming`.
	OptimizationType pulumi.StringPtrInput
	// The host header CDN provider will send along with content requests to origins.
	OriginHostHeader pulumi.StringPtrInput
	// The path used at for origin requests.
	OriginPath pulumi.StringPtrInput
	// The set of origins of the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options. Each `origin` block supports fields documented below.
	Origins EndpointOriginArrayInput
	// the path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the `originPath`.
	ProbePath pulumi.StringPtrInput
	// The CDN Profile to which to attach the CDN Endpoint.
	ProfileName pulumi.StringInput
	// Sets query string caching behavior. Allowed values are `IgnoreQueryString`, `BypassCaching` and `UseQueryString`. `NotSet` value can be used for `Premium Verizon` CDN profile. Defaults to `IgnoreQueryString`.
	QuerystringCachingBehaviour pulumi.StringPtrInput
	// The name of the resource group in which to create the CDN Endpoint.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
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
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
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
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
