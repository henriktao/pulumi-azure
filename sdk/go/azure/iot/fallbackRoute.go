// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IotHub Fallback Route
//
// ## Disclaimers
//
// > **Note:** Fallback route can be defined either directly on the `iot.IoTHub` resource, or using the `iot.FallbackRoute` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// > **Note:** Since this resource is provisioned by default, the Azure Provider will not check for the presence of an existing resource prior to attempting to create it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 			Tags: pulumi.StringMap{
// 				"purpose": pulumi.String("testing"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEndpointStorageContainer, err := iot.NewEndpointStorageContainer(ctx, "exampleEndpointStorageContainer", &iot.EndpointStorageContainerArgs{
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			IothubName:              exampleIoTHub.Name,
// 			ConnectionString:        exampleAccount.PrimaryBlobConnectionString,
// 			BatchFrequencyInSeconds: pulumi.Int(60),
// 			MaxChunkSizeInBytes:     pulumi.Int(10485760),
// 			ContainerName:           exampleContainer.Name,
// 			Encoding:                pulumi.String("Avro"),
// 			FileNameFormat:          pulumi.String("{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewFallbackRoute(ctx, "exampleFallbackRoute", &iot.FallbackRouteArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			Condition:         pulumi.String("true"),
// 			EndpointNames: pulumi.String{
// 				exampleEndpointStorageContainer.Name,
// 			},
// 			Enabled: pulumi.Bool(true),
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
// IoTHub Fallback Route can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/fallbackRoute:FallbackRoute route1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/FallbackRoute/default
// ```
type FallbackRoute struct {
	pulumi.CustomResourceState

	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrOutput `pulumi:"condition"`
	// Used to specify whether the fallback route is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames pulumi.StringOutput `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewFallbackRoute registers a new resource with the given unique name, arguments, and options.
func NewFallbackRoute(ctx *pulumi.Context,
	name string, args *FallbackRouteArgs, opts ...pulumi.ResourceOption) (*FallbackRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.EndpointNames == nil {
		return nil, errors.New("invalid value for required argument 'EndpointNames'")
	}
	if args.IothubName == nil {
		return nil, errors.New("invalid value for required argument 'IothubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource FallbackRoute
	err := ctx.RegisterResource("azure:iot/fallbackRoute:FallbackRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFallbackRoute gets an existing FallbackRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFallbackRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FallbackRouteState, opts ...pulumi.ResourceOption) (*FallbackRoute, error) {
	var resource FallbackRoute
	err := ctx.ReadResource("azure:iot/fallbackRoute:FallbackRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FallbackRoute resources.
type fallbackRouteState struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `pulumi:"condition"`
	// Used to specify whether the fallback route is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames *string `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IothubName *string `pulumi:"iothubName"`
	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type FallbackRouteState struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrInput
	// Used to specify whether the fallback route is enabled.
	Enabled pulumi.BoolPtrInput
	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames pulumi.StringPtrInput
	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringPtrInput
	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (FallbackRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*fallbackRouteState)(nil)).Elem()
}

type fallbackRouteArgs struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `pulumi:"condition"`
	// Used to specify whether the fallback route is enabled.
	Enabled bool `pulumi:"enabled"`
	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames string `pulumi:"endpointNames"`
	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IothubName string `pulumi:"iothubName"`
	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a FallbackRoute resource.
type FallbackRouteArgs struct {
	// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to `true` by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition pulumi.StringPtrInput
	// Used to specify whether the fallback route is enabled.
	Enabled pulumi.BoolInput
	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames pulumi.StringInput
	// The name of the IoTHub to which this Fallback Route belongs. Changing this forces a new resource to be created.
	IothubName pulumi.StringInput
	// The name of the resource group under which the IotHub Storage Container Endpoint resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (FallbackRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fallbackRouteArgs)(nil)).Elem()
}

type FallbackRouteInput interface {
	pulumi.Input

	ToFallbackRouteOutput() FallbackRouteOutput
	ToFallbackRouteOutputWithContext(ctx context.Context) FallbackRouteOutput
}

func (*FallbackRoute) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRoute)(nil))
}

func (i *FallbackRoute) ToFallbackRouteOutput() FallbackRouteOutput {
	return i.ToFallbackRouteOutputWithContext(context.Background())
}

func (i *FallbackRoute) ToFallbackRouteOutputWithContext(ctx context.Context) FallbackRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRouteOutput)
}

func (i *FallbackRoute) ToFallbackRoutePtrOutput() FallbackRoutePtrOutput {
	return i.ToFallbackRoutePtrOutputWithContext(context.Background())
}

func (i *FallbackRoute) ToFallbackRoutePtrOutputWithContext(ctx context.Context) FallbackRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePtrOutput)
}

type FallbackRoutePtrInput interface {
	pulumi.Input

	ToFallbackRoutePtrOutput() FallbackRoutePtrOutput
	ToFallbackRoutePtrOutputWithContext(ctx context.Context) FallbackRoutePtrOutput
}

type fallbackRoutePtrType FallbackRouteArgs

func (*fallbackRoutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRoute)(nil))
}

func (i *fallbackRoutePtrType) ToFallbackRoutePtrOutput() FallbackRoutePtrOutput {
	return i.ToFallbackRoutePtrOutputWithContext(context.Background())
}

func (i *fallbackRoutePtrType) ToFallbackRoutePtrOutputWithContext(ctx context.Context) FallbackRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePtrOutput)
}

// FallbackRouteArrayInput is an input type that accepts FallbackRouteArray and FallbackRouteArrayOutput values.
// You can construct a concrete instance of `FallbackRouteArrayInput` via:
//
//          FallbackRouteArray{ FallbackRouteArgs{...} }
type FallbackRouteArrayInput interface {
	pulumi.Input

	ToFallbackRouteArrayOutput() FallbackRouteArrayOutput
	ToFallbackRouteArrayOutputWithContext(context.Context) FallbackRouteArrayOutput
}

type FallbackRouteArray []FallbackRouteInput

func (FallbackRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FallbackRoute)(nil)).Elem()
}

func (i FallbackRouteArray) ToFallbackRouteArrayOutput() FallbackRouteArrayOutput {
	return i.ToFallbackRouteArrayOutputWithContext(context.Background())
}

func (i FallbackRouteArray) ToFallbackRouteArrayOutputWithContext(ctx context.Context) FallbackRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRouteArrayOutput)
}

// FallbackRouteMapInput is an input type that accepts FallbackRouteMap and FallbackRouteMapOutput values.
// You can construct a concrete instance of `FallbackRouteMapInput` via:
//
//          FallbackRouteMap{ "key": FallbackRouteArgs{...} }
type FallbackRouteMapInput interface {
	pulumi.Input

	ToFallbackRouteMapOutput() FallbackRouteMapOutput
	ToFallbackRouteMapOutputWithContext(context.Context) FallbackRouteMapOutput
}

type FallbackRouteMap map[string]FallbackRouteInput

func (FallbackRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FallbackRoute)(nil)).Elem()
}

func (i FallbackRouteMap) ToFallbackRouteMapOutput() FallbackRouteMapOutput {
	return i.ToFallbackRouteMapOutputWithContext(context.Background())
}

func (i FallbackRouteMap) ToFallbackRouteMapOutputWithContext(ctx context.Context) FallbackRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRouteMapOutput)
}

type FallbackRouteOutput struct{ *pulumi.OutputState }

func (FallbackRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRoute)(nil))
}

func (o FallbackRouteOutput) ToFallbackRouteOutput() FallbackRouteOutput {
	return o
}

func (o FallbackRouteOutput) ToFallbackRouteOutputWithContext(ctx context.Context) FallbackRouteOutput {
	return o
}

func (o FallbackRouteOutput) ToFallbackRoutePtrOutput() FallbackRoutePtrOutput {
	return o.ToFallbackRoutePtrOutputWithContext(context.Background())
}

func (o FallbackRouteOutput) ToFallbackRoutePtrOutputWithContext(ctx context.Context) FallbackRoutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FallbackRoute) *FallbackRoute {
		return &v
	}).(FallbackRoutePtrOutput)
}

type FallbackRoutePtrOutput struct{ *pulumi.OutputState }

func (FallbackRoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRoute)(nil))
}

func (o FallbackRoutePtrOutput) ToFallbackRoutePtrOutput() FallbackRoutePtrOutput {
	return o
}

func (o FallbackRoutePtrOutput) ToFallbackRoutePtrOutputWithContext(ctx context.Context) FallbackRoutePtrOutput {
	return o
}

func (o FallbackRoutePtrOutput) Elem() FallbackRouteOutput {
	return o.ApplyT(func(v *FallbackRoute) FallbackRoute {
		if v != nil {
			return *v
		}
		var ret FallbackRoute
		return ret
	}).(FallbackRouteOutput)
}

type FallbackRouteArrayOutput struct{ *pulumi.OutputState }

func (FallbackRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FallbackRoute)(nil))
}

func (o FallbackRouteArrayOutput) ToFallbackRouteArrayOutput() FallbackRouteArrayOutput {
	return o
}

func (o FallbackRouteArrayOutput) ToFallbackRouteArrayOutputWithContext(ctx context.Context) FallbackRouteArrayOutput {
	return o
}

func (o FallbackRouteArrayOutput) Index(i pulumi.IntInput) FallbackRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FallbackRoute {
		return vs[0].([]FallbackRoute)[vs[1].(int)]
	}).(FallbackRouteOutput)
}

type FallbackRouteMapOutput struct{ *pulumi.OutputState }

func (FallbackRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FallbackRoute)(nil))
}

func (o FallbackRouteMapOutput) ToFallbackRouteMapOutput() FallbackRouteMapOutput {
	return o
}

func (o FallbackRouteMapOutput) ToFallbackRouteMapOutputWithContext(ctx context.Context) FallbackRouteMapOutput {
	return o
}

func (o FallbackRouteMapOutput) MapIndex(k pulumi.StringInput) FallbackRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FallbackRoute {
		return vs[0].(map[string]FallbackRoute)[vs[1].(string)]
	}).(FallbackRouteOutput)
}

func init() {
	pulumi.RegisterOutputType(FallbackRouteOutput{})
	pulumi.RegisterOutputType(FallbackRoutePtrOutput{})
	pulumi.RegisterOutputType(FallbackRouteArrayOutput{})
	pulumi.RegisterOutputType(FallbackRouteMapOutput{})
}
