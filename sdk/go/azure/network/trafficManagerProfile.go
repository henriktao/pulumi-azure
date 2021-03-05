// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Traffic Manager Profile to which multiple endpoints can be attached.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		server, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.Float64Map{
// 				"azi_id": pulumi.Float64(1),
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
// 		_, err = network.NewTrafficManagerProfile(ctx, "exampleTrafficManagerProfile", &network.TrafficManagerProfileArgs{
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
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Traffic Manager Profiles can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/trafficManagerProfile:TrafficManagerProfile exampleProfile /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1
// ```
type TrafficManagerProfile struct {
	pulumi.CustomResourceState

	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigOutput `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
	MaxReturn pulumi.IntPtrOutput `pulumi:"maxReturn"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigOutput `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringOutput `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringOutput `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled pulumi.BoolPtrOutput `pulumi:"trafficViewEnabled"`
}

// NewTrafficManagerProfile registers a new resource with the given unique name, arguments, and options.
func NewTrafficManagerProfile(ctx *pulumi.Context,
	name string, args *TrafficManagerProfileArgs, opts ...pulumi.ResourceOption) (*TrafficManagerProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsConfig == nil {
		return nil, errors.New("invalid value for required argument 'DnsConfig'")
	}
	if args.MonitorConfig == nil {
		return nil, errors.New("invalid value for required argument 'MonitorConfig'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficRoutingMethod == nil {
		return nil, errors.New("invalid value for required argument 'TrafficRoutingMethod'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:trafficmanager/profile:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource TrafficManagerProfile
	err := ctx.RegisterResource("azure:network/trafficManagerProfile:TrafficManagerProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficManagerProfile gets an existing TrafficManagerProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficManagerProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerProfileState, opts ...pulumi.ResourceOption) (*TrafficManagerProfile, error) {
	var resource TrafficManagerProfile
	err := ctx.ReadResource("azure:network/trafficManagerProfile:TrafficManagerProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficManagerProfile resources.
type trafficManagerProfileState struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig *TrafficManagerProfileDnsConfig `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn *string `pulumi:"fqdn"`
	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
	MaxReturn *int `pulumi:"maxReturn"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig *TrafficManagerProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `pulumi:"trafficViewEnabled"`
}

type TrafficManagerProfileState struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigPtrInput
	// The FQDN of the created Profile.
	Fqdn pulumi.StringPtrInput
	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
	MaxReturn pulumi.IntPtrInput
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigPtrInput
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringPtrInput
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled pulumi.BoolPtrInput
}

func (TrafficManagerProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerProfileState)(nil)).Elem()
}

type trafficManagerProfileArgs struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfig `pulumi:"dnsConfig"`
	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
	MaxReturn *int `pulumi:"maxReturn"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `pulumi:"trafficViewEnabled"`
}

// The set of arguments for constructing a TrafficManagerProfile resource.
type TrafficManagerProfileArgs struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig TrafficManagerProfileDnsConfigInput
	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
	MaxReturn pulumi.IntPtrInput
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig TrafficManagerProfileMonitorConfigInput
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringInput
	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled pulumi.BoolPtrInput
}

func (TrafficManagerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerProfileArgs)(nil)).Elem()
}

type TrafficManagerProfileInput interface {
	pulumi.Input

	ToTrafficManagerProfileOutput() TrafficManagerProfileOutput
	ToTrafficManagerProfileOutputWithContext(ctx context.Context) TrafficManagerProfileOutput
}

func (*TrafficManagerProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerProfile)(nil))
}

func (i *TrafficManagerProfile) ToTrafficManagerProfileOutput() TrafficManagerProfileOutput {
	return i.ToTrafficManagerProfileOutputWithContext(context.Background())
}

func (i *TrafficManagerProfile) ToTrafficManagerProfileOutputWithContext(ctx context.Context) TrafficManagerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerProfileOutput)
}

func (i *TrafficManagerProfile) ToTrafficManagerProfilePtrOutput() TrafficManagerProfilePtrOutput {
	return i.ToTrafficManagerProfilePtrOutputWithContext(context.Background())
}

func (i *TrafficManagerProfile) ToTrafficManagerProfilePtrOutputWithContext(ctx context.Context) TrafficManagerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerProfilePtrOutput)
}

type TrafficManagerProfilePtrInput interface {
	pulumi.Input

	ToTrafficManagerProfilePtrOutput() TrafficManagerProfilePtrOutput
	ToTrafficManagerProfilePtrOutputWithContext(ctx context.Context) TrafficManagerProfilePtrOutput
}

type trafficManagerProfilePtrType TrafficManagerProfileArgs

func (*trafficManagerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerProfile)(nil))
}

func (i *trafficManagerProfilePtrType) ToTrafficManagerProfilePtrOutput() TrafficManagerProfilePtrOutput {
	return i.ToTrafficManagerProfilePtrOutputWithContext(context.Background())
}

func (i *trafficManagerProfilePtrType) ToTrafficManagerProfilePtrOutputWithContext(ctx context.Context) TrafficManagerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerProfilePtrOutput)
}

// TrafficManagerProfileArrayInput is an input type that accepts TrafficManagerProfileArray and TrafficManagerProfileArrayOutput values.
// You can construct a concrete instance of `TrafficManagerProfileArrayInput` via:
//
//          TrafficManagerProfileArray{ TrafficManagerProfileArgs{...} }
type TrafficManagerProfileArrayInput interface {
	pulumi.Input

	ToTrafficManagerProfileArrayOutput() TrafficManagerProfileArrayOutput
	ToTrafficManagerProfileArrayOutputWithContext(context.Context) TrafficManagerProfileArrayOutput
}

type TrafficManagerProfileArray []TrafficManagerProfileInput

func (TrafficManagerProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TrafficManagerProfile)(nil))
}

func (i TrafficManagerProfileArray) ToTrafficManagerProfileArrayOutput() TrafficManagerProfileArrayOutput {
	return i.ToTrafficManagerProfileArrayOutputWithContext(context.Background())
}

func (i TrafficManagerProfileArray) ToTrafficManagerProfileArrayOutputWithContext(ctx context.Context) TrafficManagerProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerProfileArrayOutput)
}

// TrafficManagerProfileMapInput is an input type that accepts TrafficManagerProfileMap and TrafficManagerProfileMapOutput values.
// You can construct a concrete instance of `TrafficManagerProfileMapInput` via:
//
//          TrafficManagerProfileMap{ "key": TrafficManagerProfileArgs{...} }
type TrafficManagerProfileMapInput interface {
	pulumi.Input

	ToTrafficManagerProfileMapOutput() TrafficManagerProfileMapOutput
	ToTrafficManagerProfileMapOutputWithContext(context.Context) TrafficManagerProfileMapOutput
}

type TrafficManagerProfileMap map[string]TrafficManagerProfileInput

func (TrafficManagerProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TrafficManagerProfile)(nil))
}

func (i TrafficManagerProfileMap) ToTrafficManagerProfileMapOutput() TrafficManagerProfileMapOutput {
	return i.ToTrafficManagerProfileMapOutputWithContext(context.Background())
}

func (i TrafficManagerProfileMap) ToTrafficManagerProfileMapOutputWithContext(ctx context.Context) TrafficManagerProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerProfileMapOutput)
}

type TrafficManagerProfileOutput struct {
	*pulumi.OutputState
}

func (TrafficManagerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficManagerProfile)(nil))
}

func (o TrafficManagerProfileOutput) ToTrafficManagerProfileOutput() TrafficManagerProfileOutput {
	return o
}

func (o TrafficManagerProfileOutput) ToTrafficManagerProfileOutputWithContext(ctx context.Context) TrafficManagerProfileOutput {
	return o
}

func (o TrafficManagerProfileOutput) ToTrafficManagerProfilePtrOutput() TrafficManagerProfilePtrOutput {
	return o.ToTrafficManagerProfilePtrOutputWithContext(context.Background())
}

func (o TrafficManagerProfileOutput) ToTrafficManagerProfilePtrOutputWithContext(ctx context.Context) TrafficManagerProfilePtrOutput {
	return o.ApplyT(func(v TrafficManagerProfile) *TrafficManagerProfile {
		return &v
	}).(TrafficManagerProfilePtrOutput)
}

type TrafficManagerProfilePtrOutput struct {
	*pulumi.OutputState
}

func (TrafficManagerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerProfile)(nil))
}

func (o TrafficManagerProfilePtrOutput) ToTrafficManagerProfilePtrOutput() TrafficManagerProfilePtrOutput {
	return o
}

func (o TrafficManagerProfilePtrOutput) ToTrafficManagerProfilePtrOutputWithContext(ctx context.Context) TrafficManagerProfilePtrOutput {
	return o
}

type TrafficManagerProfileArrayOutput struct{ *pulumi.OutputState }

func (TrafficManagerProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficManagerProfile)(nil))
}

func (o TrafficManagerProfileArrayOutput) ToTrafficManagerProfileArrayOutput() TrafficManagerProfileArrayOutput {
	return o
}

func (o TrafficManagerProfileArrayOutput) ToTrafficManagerProfileArrayOutputWithContext(ctx context.Context) TrafficManagerProfileArrayOutput {
	return o
}

func (o TrafficManagerProfileArrayOutput) Index(i pulumi.IntInput) TrafficManagerProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficManagerProfile {
		return vs[0].([]TrafficManagerProfile)[vs[1].(int)]
	}).(TrafficManagerProfileOutput)
}

type TrafficManagerProfileMapOutput struct{ *pulumi.OutputState }

func (TrafficManagerProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TrafficManagerProfile)(nil))
}

func (o TrafficManagerProfileMapOutput) ToTrafficManagerProfileMapOutput() TrafficManagerProfileMapOutput {
	return o
}

func (o TrafficManagerProfileMapOutput) ToTrafficManagerProfileMapOutputWithContext(ctx context.Context) TrafficManagerProfileMapOutput {
	return o
}

func (o TrafficManagerProfileMapOutput) MapIndex(k pulumi.StringInput) TrafficManagerProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TrafficManagerProfile {
		return vs[0].(map[string]TrafficManagerProfile)[vs[1].(string)]
	}).(TrafficManagerProfileOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficManagerProfileOutput{})
	pulumi.RegisterOutputType(TrafficManagerProfilePtrOutput{})
	pulumi.RegisterOutputType(TrafficManagerProfileArrayOutput{})
	pulumi.RegisterOutputType(TrafficManagerProfileMapOutput{})
}
