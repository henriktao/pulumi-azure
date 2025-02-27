// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IotHub ServiceBus Queue Endpoint
//
// > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
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
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleQueue, err := servicebus.NewQueue(ctx, "exampleQueue", &servicebus.QueueArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			NamespaceName:      exampleNamespace.Name,
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleQueueAuthorizationRule, err := servicebus.NewQueueAuthorizationRule(ctx, "exampleQueueAuthorizationRule", &servicebus.QueueAuthorizationRuleArgs{
// 			NamespaceName:     exampleNamespace.Name,
// 			QueueName:         exampleQueue.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Listen:            pulumi.Bool(false),
// 			Send:              pulumi.Bool(true),
// 			Manage:            pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("B1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 			Tags: pulumi.StringMap{
// 				"purpose": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewEndpointServicebusQueue(ctx, "exampleEndpointServicebusQueue", &iot.EndpointServicebusQueueArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			ConnectionString:  exampleQueueAuthorizationRule.PrimaryConnectionString,
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
// IoTHub ServiceBus Queue Endpoint can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/endpointServicebusQueue:EndpointServicebusQueue servicebus_queue1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/servicebusqueue_endpoint1
// ```
type EndpointServicebusQueue struct {
	pulumi.CustomResourceState

	// Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
	AuthenticationType pulumi.StringPtrOutput `pulumi:"authenticationType"`
	// The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EndpointUri pulumi.StringPtrOutput `pulumi:"endpointUri"`
	// Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EntityPath pulumi.StringPtrOutput `pulumi:"entityPath"`
	// ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
	IdentityId pulumi.StringPtrOutput `pulumi:"identityId"`
	// The IoTHub ID for the endpoint.
	IothubId pulumi.StringOutput `pulumi:"iothubId"`
	// The IoTHub name for the endpoint.
	//
	// Deprecated: Deprecated in favour of `iothub_id`
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEndpointServicebusQueue registers a new resource with the given unique name, arguments, and options.
func NewEndpointServicebusQueue(ctx *pulumi.Context,
	name string, args *EndpointServicebusQueueArgs, opts ...pulumi.ResourceOption) (*EndpointServicebusQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EndpointServicebusQueue
	err := ctx.RegisterResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointServicebusQueue gets an existing EndpointServicebusQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointServicebusQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointServicebusQueueState, opts ...pulumi.ResourceOption) (*EndpointServicebusQueue, error) {
	var resource EndpointServicebusQueue
	err := ctx.ReadResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointServicebusQueue resources.
type endpointServicebusQueueState struct {
	// Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
	AuthenticationType *string `pulumi:"authenticationType"`
	// The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
	ConnectionString *string `pulumi:"connectionString"`
	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EndpointUri *string `pulumi:"endpointUri"`
	// Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EntityPath *string `pulumi:"entityPath"`
	// ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
	IdentityId *string `pulumi:"identityId"`
	// The IoTHub ID for the endpoint.
	IothubId *string `pulumi:"iothubId"`
	// The IoTHub name for the endpoint.
	//
	// Deprecated: Deprecated in favour of `iothub_id`
	IothubName *string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EndpointServicebusQueueState struct {
	// Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
	AuthenticationType pulumi.StringPtrInput
	// The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
	ConnectionString pulumi.StringPtrInput
	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EndpointUri pulumi.StringPtrInput
	// Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EntityPath pulumi.StringPtrInput
	// ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
	IdentityId pulumi.StringPtrInput
	// The IoTHub ID for the endpoint.
	IothubId pulumi.StringPtrInput
	// The IoTHub name for the endpoint.
	//
	// Deprecated: Deprecated in favour of `iothub_id`
	IothubName pulumi.StringPtrInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (EndpointServicebusQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusQueueState)(nil)).Elem()
}

type endpointServicebusQueueArgs struct {
	// Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
	AuthenticationType *string `pulumi:"authenticationType"`
	// The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
	ConnectionString *string `pulumi:"connectionString"`
	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EndpointUri *string `pulumi:"endpointUri"`
	// Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EntityPath *string `pulumi:"entityPath"`
	// ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
	IdentityId *string `pulumi:"identityId"`
	// The IoTHub ID for the endpoint.
	IothubId *string `pulumi:"iothubId"`
	// The IoTHub name for the endpoint.
	//
	// Deprecated: Deprecated in favour of `iothub_id`
	IothubName *string `pulumi:"iothubName"`
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EndpointServicebusQueue resource.
type EndpointServicebusQueueArgs struct {
	// Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
	AuthenticationType pulumi.StringPtrInput
	// The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
	ConnectionString pulumi.StringPtrInput
	// URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EndpointUri pulumi.StringPtrInput
	// Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
	EntityPath pulumi.StringPtrInput
	// ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
	IdentityId pulumi.StringPtrInput
	// The IoTHub ID for the endpoint.
	IothubId pulumi.StringPtrInput
	// The IoTHub name for the endpoint.
	//
	// Deprecated: Deprecated in favour of `iothub_id`
	IothubName pulumi.StringPtrInput
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (EndpointServicebusQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointServicebusQueueArgs)(nil)).Elem()
}

type EndpointServicebusQueueInput interface {
	pulumi.Input

	ToEndpointServicebusQueueOutput() EndpointServicebusQueueOutput
	ToEndpointServicebusQueueOutputWithContext(ctx context.Context) EndpointServicebusQueueOutput
}

func (*EndpointServicebusQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointServicebusQueue)(nil)).Elem()
}

func (i *EndpointServicebusQueue) ToEndpointServicebusQueueOutput() EndpointServicebusQueueOutput {
	return i.ToEndpointServicebusQueueOutputWithContext(context.Background())
}

func (i *EndpointServicebusQueue) ToEndpointServicebusQueueOutputWithContext(ctx context.Context) EndpointServicebusQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointServicebusQueueOutput)
}

// EndpointServicebusQueueArrayInput is an input type that accepts EndpointServicebusQueueArray and EndpointServicebusQueueArrayOutput values.
// You can construct a concrete instance of `EndpointServicebusQueueArrayInput` via:
//
//          EndpointServicebusQueueArray{ EndpointServicebusQueueArgs{...} }
type EndpointServicebusQueueArrayInput interface {
	pulumi.Input

	ToEndpointServicebusQueueArrayOutput() EndpointServicebusQueueArrayOutput
	ToEndpointServicebusQueueArrayOutputWithContext(context.Context) EndpointServicebusQueueArrayOutput
}

type EndpointServicebusQueueArray []EndpointServicebusQueueInput

func (EndpointServicebusQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointServicebusQueue)(nil)).Elem()
}

func (i EndpointServicebusQueueArray) ToEndpointServicebusQueueArrayOutput() EndpointServicebusQueueArrayOutput {
	return i.ToEndpointServicebusQueueArrayOutputWithContext(context.Background())
}

func (i EndpointServicebusQueueArray) ToEndpointServicebusQueueArrayOutputWithContext(ctx context.Context) EndpointServicebusQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointServicebusQueueArrayOutput)
}

// EndpointServicebusQueueMapInput is an input type that accepts EndpointServicebusQueueMap and EndpointServicebusQueueMapOutput values.
// You can construct a concrete instance of `EndpointServicebusQueueMapInput` via:
//
//          EndpointServicebusQueueMap{ "key": EndpointServicebusQueueArgs{...} }
type EndpointServicebusQueueMapInput interface {
	pulumi.Input

	ToEndpointServicebusQueueMapOutput() EndpointServicebusQueueMapOutput
	ToEndpointServicebusQueueMapOutputWithContext(context.Context) EndpointServicebusQueueMapOutput
}

type EndpointServicebusQueueMap map[string]EndpointServicebusQueueInput

func (EndpointServicebusQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointServicebusQueue)(nil)).Elem()
}

func (i EndpointServicebusQueueMap) ToEndpointServicebusQueueMapOutput() EndpointServicebusQueueMapOutput {
	return i.ToEndpointServicebusQueueMapOutputWithContext(context.Background())
}

func (i EndpointServicebusQueueMap) ToEndpointServicebusQueueMapOutputWithContext(ctx context.Context) EndpointServicebusQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointServicebusQueueMapOutput)
}

type EndpointServicebusQueueOutput struct{ *pulumi.OutputState }

func (EndpointServicebusQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointServicebusQueue)(nil)).Elem()
}

func (o EndpointServicebusQueueOutput) ToEndpointServicebusQueueOutput() EndpointServicebusQueueOutput {
	return o
}

func (o EndpointServicebusQueueOutput) ToEndpointServicebusQueueOutputWithContext(ctx context.Context) EndpointServicebusQueueOutput {
	return o
}

type EndpointServicebusQueueArrayOutput struct{ *pulumi.OutputState }

func (EndpointServicebusQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointServicebusQueue)(nil)).Elem()
}

func (o EndpointServicebusQueueArrayOutput) ToEndpointServicebusQueueArrayOutput() EndpointServicebusQueueArrayOutput {
	return o
}

func (o EndpointServicebusQueueArrayOutput) ToEndpointServicebusQueueArrayOutputWithContext(ctx context.Context) EndpointServicebusQueueArrayOutput {
	return o
}

func (o EndpointServicebusQueueArrayOutput) Index(i pulumi.IntInput) EndpointServicebusQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointServicebusQueue {
		return vs[0].([]*EndpointServicebusQueue)[vs[1].(int)]
	}).(EndpointServicebusQueueOutput)
}

type EndpointServicebusQueueMapOutput struct{ *pulumi.OutputState }

func (EndpointServicebusQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointServicebusQueue)(nil)).Elem()
}

func (o EndpointServicebusQueueMapOutput) ToEndpointServicebusQueueMapOutput() EndpointServicebusQueueMapOutput {
	return o
}

func (o EndpointServicebusQueueMapOutput) ToEndpointServicebusQueueMapOutputWithContext(ctx context.Context) EndpointServicebusQueueMapOutput {
	return o
}

func (o EndpointServicebusQueueMapOutput) MapIndex(k pulumi.StringInput) EndpointServicebusQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointServicebusQueue {
		return vs[0].(map[string]*EndpointServicebusQueue)[vs[1].(string)]
	}).(EndpointServicebusQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointServicebusQueueInput)(nil)).Elem(), &EndpointServicebusQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointServicebusQueueArrayInput)(nil)).Elem(), EndpointServicebusQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointServicebusQueueMapInput)(nil)).Elem(), EndpointServicebusQueueMap{})
	pulumi.RegisterOutputType(EndpointServicebusQueueOutput{})
	pulumi.RegisterOutputType(EndpointServicebusQueueArrayOutput{})
	pulumi.RegisterOutputType(EndpointServicebusQueueMapOutput{})
}
