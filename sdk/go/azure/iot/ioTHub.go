// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IotHub
//
// > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.
//
// > **NOTE:** Routes can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Route` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// > **NOTE:** Enrichments can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Enrichment` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// > **NOTE:** Fallback route can be defined either directly on the `iot.IoTHub` resource, or using the `iot.FallbackRoute` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku:               pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAuthorizationRule, err := eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
// 			Send:              pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 			Endpoints: iot.IoTHubEndpointArray{
// 				&iot.IoTHubEndpointArgs{
// 					Type:                    pulumi.String("AzureIotHub.StorageContainer"),
// 					ConnectionString:        exampleAccount.PrimaryBlobConnectionString,
// 					Name:                    pulumi.String("export"),
// 					BatchFrequencyInSeconds: pulumi.Int(60),
// 					MaxChunkSizeInBytes:     pulumi.Int(10485760),
// 					ContainerName:           exampleContainer.Name,
// 					Encoding:                pulumi.String("Avro"),
// 					FileNameFormat:          pulumi.String("{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}"),
// 				},
// 				&iot.IoTHubEndpointArgs{
// 					Type:             pulumi.String("AzureIotHub.EventHub"),
// 					ConnectionString: exampleAuthorizationRule.PrimaryConnectionString,
// 					Name:             pulumi.String("export2"),
// 				},
// 			},
// 			Routes: iot.IoTHubRouteArray{
// 				&iot.IoTHubRouteArgs{
// 					Name:      pulumi.String("export"),
// 					Source:    pulumi.String("DeviceMessages"),
// 					Condition: pulumi.String("true"),
// 					EndpointNames: pulumi.StringArray{
// 						pulumi.String("export"),
// 					},
// 					Enabled: pulumi.Bool(true),
// 				},
// 				&iot.IoTHubRouteArgs{
// 					Name:      pulumi.String("export2"),
// 					Source:    pulumi.String("DeviceMessages"),
// 					Condition: pulumi.String("true"),
// 					EndpointNames: pulumi.StringArray{
// 						pulumi.String("export2"),
// 					},
// 					Enabled: pulumi.Bool(true),
// 				},
// 			},
// 			Enrichments: iot.IoTHubEnrichmentArray{
// 				&iot.IoTHubEnrichmentArgs{
// 					Key:   pulumi.String("tenant"),
// 					Value: pulumi.String(fmt.Sprintf("%v%v", "$", "twin.tags.Tenant")),
// 					EndpointNames: pulumi.StringArray{
// 						pulumi.String("export"),
// 						pulumi.String("export2"),
// 					},
// 				},
// 			},
// 			CloudToDevice: &iot.IoTHubCloudToDeviceArgs{
// 				MaxDeliveryCount: pulumi.Int(30),
// 				DefaultTtl:       pulumi.String("PT1H"),
// 				Feedbacks: iot.IoTHubCloudToDeviceFeedbackArray{
// 					&iot.IoTHubCloudToDeviceFeedbackArgs{
// 						TimeToLive:       pulumi.String("PT1H10M"),
// 						MaxDeliveryCount: pulumi.Int(15),
// 						LockDuration:     pulumi.String("PT30S"),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"purpose": pulumi.String("testing"),
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
// IoTHubs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/ioTHub:IoTHub hub1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1
// ```
type IoTHub struct {
	pulumi.CustomResourceState

	// A `cloudToDevice` block as defined below.
	CloudToDevice IoTHubCloudToDeviceOutput `pulumi:"cloudToDevice"`
	// An `endpoint` block as defined below.
	Endpoints IoTHubEndpointArrayOutput `pulumi:"endpoints"`
	// A `enrichment` block as defined below.
	Enrichments IoTHubEnrichmentArrayOutput `pulumi:"enrichments"`
	// The EventHub compatible endpoint for events data
	EventHubEventsEndpoint pulumi.StringOutput `pulumi:"eventHubEventsEndpoint"`
	// The EventHub namespace for events data
	EventHubEventsNamespace pulumi.StringOutput `pulumi:"eventHubEventsNamespace"`
	// The EventHub compatible path for events data
	EventHubEventsPath pulumi.StringOutput `pulumi:"eventHubEventsPath"`
	// The EventHub compatible endpoint for operational data
	EventHubOperationsEndpoint pulumi.StringOutput `pulumi:"eventHubOperationsEndpoint"`
	// The EventHub compatible path for operational data
	EventHubOperationsPath pulumi.StringOutput `pulumi:"eventHubOperationsPath"`
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount pulumi.IntOutput `pulumi:"eventHubPartitionCount"`
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays pulumi.IntOutput `pulumi:"eventHubRetentionInDays"`
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute IoTHubFallbackRouteOutput `pulumi:"fallbackRoute"`
	// A `fileUpload` block as defined below.
	FileUpload IoTHubFileUploadPtrOutput `pulumi:"fileUpload"`
	// The hostname of the IotHub Resource.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// An `identity` block as defined below.
	Identity IoTHubIdentityPtrOutput `pulumi:"identity"`
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules IoTHubIpFilterRuleArrayOutput `pulumi:"ipFilterRules"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the minimum TLS version to support for this hub. The only valid value is `1.2`. Changing this forces a new resource to be created.
	MinTlsVersion pulumi.StringPtrOutput `pulumi:"minTlsVersion"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `route` block as defined below.
	Routes IoTHubRouteArrayOutput `pulumi:"routes"`
	// One or more `sharedAccessPolicy` blocks as defined below.
	SharedAccessPolicies IoTHubSharedAccessPolicyArrayOutput `pulumi:"sharedAccessPolicies"`
	// A `sku` block as defined below.
	Sku IoTHubSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIoTHub registers a new resource with the given unique name, arguments, and options.
func NewIoTHub(ctx *pulumi.Context,
	name string, args *IoTHubArgs, opts ...pulumi.ResourceOption) (*IoTHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource IoTHub
	err := ctx.RegisterResource("azure:iot/ioTHub:IoTHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIoTHub gets an existing IoTHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTHubState, opts ...pulumi.ResourceOption) (*IoTHub, error) {
	var resource IoTHub
	err := ctx.ReadResource("azure:iot/ioTHub:IoTHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IoTHub resources.
type ioTHubState struct {
	// A `cloudToDevice` block as defined below.
	CloudToDevice *IoTHubCloudToDevice `pulumi:"cloudToDevice"`
	// An `endpoint` block as defined below.
	Endpoints []IoTHubEndpoint `pulumi:"endpoints"`
	// A `enrichment` block as defined below.
	Enrichments []IoTHubEnrichment `pulumi:"enrichments"`
	// The EventHub compatible endpoint for events data
	EventHubEventsEndpoint *string `pulumi:"eventHubEventsEndpoint"`
	// The EventHub namespace for events data
	EventHubEventsNamespace *string `pulumi:"eventHubEventsNamespace"`
	// The EventHub compatible path for events data
	EventHubEventsPath *string `pulumi:"eventHubEventsPath"`
	// The EventHub compatible endpoint for operational data
	EventHubOperationsEndpoint *string `pulumi:"eventHubOperationsEndpoint"`
	// The EventHub compatible path for operational data
	EventHubOperationsPath *string `pulumi:"eventHubOperationsPath"`
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount *int `pulumi:"eventHubPartitionCount"`
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays *int `pulumi:"eventHubRetentionInDays"`
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute *IoTHubFallbackRoute `pulumi:"fallbackRoute"`
	// A `fileUpload` block as defined below.
	FileUpload *IoTHubFileUpload `pulumi:"fileUpload"`
	// The hostname of the IotHub Resource.
	Hostname *string `pulumi:"hostname"`
	// An `identity` block as defined below.
	Identity *IoTHubIdentity `pulumi:"identity"`
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules []IoTHubIpFilterRule `pulumi:"ipFilterRules"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the minimum TLS version to support for this hub. The only valid value is `1.2`. Changing this forces a new resource to be created.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `route` block as defined below.
	Routes []IoTHubRoute `pulumi:"routes"`
	// One or more `sharedAccessPolicy` blocks as defined below.
	SharedAccessPolicies []IoTHubSharedAccessPolicy `pulumi:"sharedAccessPolicies"`
	// A `sku` block as defined below.
	Sku *IoTHubSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
	Type *string `pulumi:"type"`
}

type IoTHubState struct {
	// A `cloudToDevice` block as defined below.
	CloudToDevice IoTHubCloudToDevicePtrInput
	// An `endpoint` block as defined below.
	Endpoints IoTHubEndpointArrayInput
	// A `enrichment` block as defined below.
	Enrichments IoTHubEnrichmentArrayInput
	// The EventHub compatible endpoint for events data
	EventHubEventsEndpoint pulumi.StringPtrInput
	// The EventHub namespace for events data
	EventHubEventsNamespace pulumi.StringPtrInput
	// The EventHub compatible path for events data
	EventHubEventsPath pulumi.StringPtrInput
	// The EventHub compatible endpoint for operational data
	EventHubOperationsEndpoint pulumi.StringPtrInput
	// The EventHub compatible path for operational data
	EventHubOperationsPath pulumi.StringPtrInput
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount pulumi.IntPtrInput
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays pulumi.IntPtrInput
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute IoTHubFallbackRoutePtrInput
	// A `fileUpload` block as defined below.
	FileUpload IoTHubFileUploadPtrInput
	// The hostname of the IotHub Resource.
	Hostname pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity IoTHubIdentityPtrInput
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules IoTHubIpFilterRuleArrayInput
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the minimum TLS version to support for this hub. The only valid value is `1.2`. Changing this forces a new resource to be created.
	MinTlsVersion pulumi.StringPtrInput
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `route` block as defined below.
	Routes IoTHubRouteArrayInput
	// One or more `sharedAccessPolicy` blocks as defined below.
	SharedAccessPolicies IoTHubSharedAccessPolicyArrayInput
	// A `sku` block as defined below.
	Sku IoTHubSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the endpoint. Possible values are `AzureIotHub.StorageContainer`, `AzureIotHub.ServiceBusQueue`, `AzureIotHub.ServiceBusTopic` or `AzureIotHub.EventHub`.
	Type pulumi.StringPtrInput
}

func (IoTHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubState)(nil)).Elem()
}

type ioTHubArgs struct {
	// A `cloudToDevice` block as defined below.
	CloudToDevice *IoTHubCloudToDevice `pulumi:"cloudToDevice"`
	// An `endpoint` block as defined below.
	Endpoints []IoTHubEndpoint `pulumi:"endpoints"`
	// A `enrichment` block as defined below.
	Enrichments []IoTHubEnrichment `pulumi:"enrichments"`
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount *int `pulumi:"eventHubPartitionCount"`
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays *int `pulumi:"eventHubRetentionInDays"`
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute *IoTHubFallbackRoute `pulumi:"fallbackRoute"`
	// A `fileUpload` block as defined below.
	FileUpload *IoTHubFileUpload `pulumi:"fileUpload"`
	// An `identity` block as defined below.
	Identity *IoTHubIdentity `pulumi:"identity"`
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules []IoTHubIpFilterRule `pulumi:"ipFilterRules"`
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the minimum TLS version to support for this hub. The only valid value is `1.2`. Changing this forces a new resource to be created.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `route` block as defined below.
	Routes []IoTHubRoute `pulumi:"routes"`
	// A `sku` block as defined below.
	Sku IoTHubSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IoTHub resource.
type IoTHubArgs struct {
	// A `cloudToDevice` block as defined below.
	CloudToDevice IoTHubCloudToDevicePtrInput
	// An `endpoint` block as defined below.
	Endpoints IoTHubEndpointArrayInput
	// A `enrichment` block as defined below.
	Enrichments IoTHubEnrichmentArrayInput
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount pulumi.IntPtrInput
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays pulumi.IntPtrInput
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute IoTHubFallbackRoutePtrInput
	// A `fileUpload` block as defined below.
	FileUpload IoTHubFileUploadPtrInput
	// An `identity` block as defined below.
	Identity IoTHubIdentityPtrInput
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules IoTHubIpFilterRuleArrayInput
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the minimum TLS version to support for this hub. The only valid value is `1.2`. Changing this forces a new resource to be created.
	MinTlsVersion pulumi.StringPtrInput
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `route` block as defined below.
	Routes IoTHubRouteArrayInput
	// A `sku` block as defined below.
	Sku IoTHubSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (IoTHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubArgs)(nil)).Elem()
}

type IoTHubInput interface {
	pulumi.Input

	ToIoTHubOutput() IoTHubOutput
	ToIoTHubOutputWithContext(ctx context.Context) IoTHubOutput
}

func (*IoTHub) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHub)(nil)).Elem()
}

func (i *IoTHub) ToIoTHubOutput() IoTHubOutput {
	return i.ToIoTHubOutputWithContext(context.Background())
}

func (i *IoTHub) ToIoTHubOutputWithContext(ctx context.Context) IoTHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubOutput)
}

// IoTHubArrayInput is an input type that accepts IoTHubArray and IoTHubArrayOutput values.
// You can construct a concrete instance of `IoTHubArrayInput` via:
//
//          IoTHubArray{ IoTHubArgs{...} }
type IoTHubArrayInput interface {
	pulumi.Input

	ToIoTHubArrayOutput() IoTHubArrayOutput
	ToIoTHubArrayOutputWithContext(context.Context) IoTHubArrayOutput
}

type IoTHubArray []IoTHubInput

func (IoTHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IoTHub)(nil)).Elem()
}

func (i IoTHubArray) ToIoTHubArrayOutput() IoTHubArrayOutput {
	return i.ToIoTHubArrayOutputWithContext(context.Background())
}

func (i IoTHubArray) ToIoTHubArrayOutputWithContext(ctx context.Context) IoTHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubArrayOutput)
}

// IoTHubMapInput is an input type that accepts IoTHubMap and IoTHubMapOutput values.
// You can construct a concrete instance of `IoTHubMapInput` via:
//
//          IoTHubMap{ "key": IoTHubArgs{...} }
type IoTHubMapInput interface {
	pulumi.Input

	ToIoTHubMapOutput() IoTHubMapOutput
	ToIoTHubMapOutputWithContext(context.Context) IoTHubMapOutput
}

type IoTHubMap map[string]IoTHubInput

func (IoTHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IoTHub)(nil)).Elem()
}

func (i IoTHubMap) ToIoTHubMapOutput() IoTHubMapOutput {
	return i.ToIoTHubMapOutputWithContext(context.Background())
}

func (i IoTHubMap) ToIoTHubMapOutputWithContext(ctx context.Context) IoTHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubMapOutput)
}

type IoTHubOutput struct{ *pulumi.OutputState }

func (IoTHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHub)(nil)).Elem()
}

func (o IoTHubOutput) ToIoTHubOutput() IoTHubOutput {
	return o
}

func (o IoTHubOutput) ToIoTHubOutputWithContext(ctx context.Context) IoTHubOutput {
	return o
}

type IoTHubArrayOutput struct{ *pulumi.OutputState }

func (IoTHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IoTHub)(nil)).Elem()
}

func (o IoTHubArrayOutput) ToIoTHubArrayOutput() IoTHubArrayOutput {
	return o
}

func (o IoTHubArrayOutput) ToIoTHubArrayOutputWithContext(ctx context.Context) IoTHubArrayOutput {
	return o
}

func (o IoTHubArrayOutput) Index(i pulumi.IntInput) IoTHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IoTHub {
		return vs[0].([]*IoTHub)[vs[1].(int)]
	}).(IoTHubOutput)
}

type IoTHubMapOutput struct{ *pulumi.OutputState }

func (IoTHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IoTHub)(nil)).Elem()
}

func (o IoTHubMapOutput) ToIoTHubMapOutput() IoTHubMapOutput {
	return o
}

func (o IoTHubMapOutput) ToIoTHubMapOutputWithContext(ctx context.Context) IoTHubMapOutput {
	return o
}

func (o IoTHubMapOutput) MapIndex(k pulumi.StringInput) IoTHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IoTHub {
		return vs[0].(map[string]*IoTHub)[vs[1].(string)]
	}).(IoTHubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IoTHubInput)(nil)).Elem(), &IoTHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*IoTHubArrayInput)(nil)).Elem(), IoTHubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IoTHubMapInput)(nil)).Elem(), IoTHubMap{})
	pulumi.RegisterOutputType(IoTHubOutput{})
	pulumi.RegisterOutputType(IoTHubArrayOutput{})
	pulumi.RegisterOutputType(IoTHubMapOutput{})
}
