// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EventGrid System Topic Event Subscription.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventgrid"
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
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleQueue, err := storage.NewQueue(ctx, "exampleQueue", &storage.QueueArgs{
// 			StorageAccountName: exampleAccount.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewSystemTopic(ctx, "exampleSystemTopic", &eventgrid.SystemTopicArgs{
// 			Location:            pulumi.String("Global"),
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			SourceArmResourceId: exampleResourceGroup.ID(),
// 			TopicType:           pulumi.String("Microsoft.Resources.ResourceGroups"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewSystemTopicEventSubscription(ctx, "exampleSystemTopicEventSubscription", &eventgrid.SystemTopicEventSubscriptionArgs{
// 			SystemTopic:       pulumi.Any(azurerm_system_topic.Example.Name),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StorageQueueEndpoint: &eventgrid.SystemTopicEventSubscriptionStorageQueueEndpointArgs{
// 				StorageAccountId: exampleAccount.ID(),
// 				QueueName:        exampleQueue.Name,
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
// EventGrid System Topic Event Subscriptions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/systemTopics/topic1/eventSubscriptions/subscription1
// ```
type SystemTopicEventSubscription struct {
	pulumi.CustomResourceState

	// A `advancedFilter` block as defined below.
	AdvancedFilter SystemTopicEventSubscriptionAdvancedFilterPtrOutput `pulumi:"advancedFilter"`
	// Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
	AdvancedFilteringOnArraysEnabled pulumi.BoolPtrOutput `pulumi:"advancedFilteringOnArraysEnabled"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint SystemTopicEventSubscriptionAzureFunctionEndpointPtrOutput `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringOutput `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrOutput `pulumi:"expirationTimeUtc"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringOutput `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayOutput `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `retryPolicy` block as defined below.
	RetryPolicy SystemTopicEventSubscriptionRetryPolicyOutput `pulumi:"retryPolicy"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrOutput `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrOutput `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationPtrOutput `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint SystemTopicEventSubscriptionStorageQueueEndpointPtrOutput `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter SystemTopicEventSubscriptionSubjectFilterPtrOutput `pulumi:"subjectFilter"`
	// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
	SystemTopic pulumi.StringOutput `pulumi:"systemTopic"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint SystemTopicEventSubscriptionWebhookEndpointPtrOutput `pulumi:"webhookEndpoint"`
}

// NewSystemTopicEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, args *SystemTopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SystemTopic == nil {
		return nil, errors.New("invalid value for required argument 'SystemTopic'")
	}
	var resource SystemTopicEventSubscription
	err := ctx.RegisterResource("azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemTopicEventSubscription gets an existing SystemTopicEventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	var resource SystemTopicEventSubscription
	err := ctx.ReadResource("azure:eventgrid/systemTopicEventSubscription:SystemTopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemTopicEventSubscription resources.
type systemTopicEventSubscriptionState struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter *SystemTopicEventSubscriptionAdvancedFilter `pulumi:"advancedFilter"`
	// Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
	AdvancedFilteringOnArraysEnabled *bool `pulumi:"advancedFilteringOnArraysEnabled"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint *SystemTopicEventSubscriptionAzureFunctionEndpoint `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId *string `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId *string `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *SystemTopicEventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId *string `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId *string `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *SystemTopicEventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *SystemTopicEventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *SystemTopicEventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
	SystemTopic *string `pulumi:"systemTopic"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *SystemTopicEventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

type SystemTopicEventSubscriptionState struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter SystemTopicEventSubscriptionAdvancedFilterPtrInput
	// Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
	AdvancedFilteringOnArraysEnabled pulumi.BoolPtrInput
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint SystemTopicEventSubscriptionAzureFunctionEndpointPtrInput
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrInput
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringPtrInput
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrInput
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `retryPolicy` block as defined below.
	RetryPolicy SystemTopicEventSubscriptionRetryPolicyPtrInput
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrInput
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint SystemTopicEventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter SystemTopicEventSubscriptionSubjectFilterPtrInput
	// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
	SystemTopic pulumi.StringPtrInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint SystemTopicEventSubscriptionWebhookEndpointPtrInput
}

func (SystemTopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionState)(nil)).Elem()
}

type systemTopicEventSubscriptionArgs struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter *SystemTopicEventSubscriptionAdvancedFilter `pulumi:"advancedFilter"`
	// Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
	AdvancedFilteringOnArraysEnabled *bool `pulumi:"advancedFilteringOnArraysEnabled"`
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint *SystemTopicEventSubscriptionAzureFunctionEndpoint `pulumi:"azureFunctionEndpoint"`
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId *string `pulumi:"eventhubEndpointId"`
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId *string `pulumi:"hybridConnectionEndpointId"`
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// A list of labels to assign to the event subscription.
	Labels []string `pulumi:"labels"`
	// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `retryPolicy` block as defined below.
	RetryPolicy *SystemTopicEventSubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId *string `pulumi:"serviceBusQueueEndpointId"`
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId *string `pulumi:"serviceBusTopicEndpointId"`
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination *SystemTopicEventSubscriptionStorageBlobDeadLetterDestination `pulumi:"storageBlobDeadLetterDestination"`
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint *SystemTopicEventSubscriptionStorageQueueEndpoint `pulumi:"storageQueueEndpoint"`
	// A `subjectFilter` block as defined below.
	SubjectFilter *SystemTopicEventSubscriptionSubjectFilter `pulumi:"subjectFilter"`
	// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
	SystemTopic string `pulumi:"systemTopic"`
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint *SystemTopicEventSubscriptionWebhookEndpoint `pulumi:"webhookEndpoint"`
}

// The set of arguments for constructing a SystemTopicEventSubscription resource.
type SystemTopicEventSubscriptionArgs struct {
	// A `advancedFilter` block as defined below.
	AdvancedFilter SystemTopicEventSubscriptionAdvancedFilterPtrInput
	// Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
	AdvancedFilteringOnArraysEnabled pulumi.BoolPtrInput
	// An `azureFunctionEndpoint` block as defined below.
	AzureFunctionEndpoint SystemTopicEventSubscriptionAzureFunctionEndpointPtrInput
	// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	EventDeliverySchema pulumi.StringPtrInput
	// Specifies the id where the Event Hub is located.
	EventhubEndpointId pulumi.StringPtrInput
	// Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
	ExpirationTimeUtc pulumi.StringPtrInput
	// Specifies the id where the Hybrid Connection is located.
	HybridConnectionEndpointId pulumi.StringPtrInput
	// A list of applicable event types that need to be part of the event subscription.
	IncludedEventTypes pulumi.StringArrayInput
	// A list of labels to assign to the event subscription.
	Labels pulumi.StringArrayInput
	// The name which should be used for this Event Subscription. Changing this forces a new Event Subscription to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the System Topic exists. Changing this forces a new Event Subscription to be created.
	ResourceGroupName pulumi.StringInput
	// A `retryPolicy` block as defined below.
	RetryPolicy SystemTopicEventSubscriptionRetryPolicyPtrInput
	// Specifies the id where the Service Bus Queue is located.
	ServiceBusQueueEndpointId pulumi.StringPtrInput
	// Specifies the id where the Service Bus Topic is located.
	ServiceBusTopicEndpointId pulumi.StringPtrInput
	// A `storageBlobDeadLetterDestination` block as defined below.
	StorageBlobDeadLetterDestination SystemTopicEventSubscriptionStorageBlobDeadLetterDestinationPtrInput
	// A `storageQueueEndpoint` block as defined below.
	StorageQueueEndpoint SystemTopicEventSubscriptionStorageQueueEndpointPtrInput
	// A `subjectFilter` block as defined below.
	SubjectFilter SystemTopicEventSubscriptionSubjectFilterPtrInput
	// The System Topic where the Event Subscription should be created in. Changing this forces a new Event Subscription to be created.
	SystemTopic pulumi.StringInput
	// A `webhookEndpoint` block as defined below.
	WebhookEndpoint SystemTopicEventSubscriptionWebhookEndpointPtrInput
}

func (SystemTopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionArgs)(nil)).Elem()
}

type SystemTopicEventSubscriptionInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput
	ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput
}

func (*SystemTopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicEventSubscription)(nil))
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return i.ToSystemTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionOutput)
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionPtrOutput() SystemTopicEventSubscriptionPtrOutput {
	return i.ToSystemTopicEventSubscriptionPtrOutputWithContext(context.Background())
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionPtrOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionPtrOutput)
}

type SystemTopicEventSubscriptionPtrInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionPtrOutput() SystemTopicEventSubscriptionPtrOutput
	ToSystemTopicEventSubscriptionPtrOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionPtrOutput
}

type systemTopicEventSubscriptionPtrType SystemTopicEventSubscriptionArgs

func (*systemTopicEventSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicEventSubscription)(nil))
}

func (i *systemTopicEventSubscriptionPtrType) ToSystemTopicEventSubscriptionPtrOutput() SystemTopicEventSubscriptionPtrOutput {
	return i.ToSystemTopicEventSubscriptionPtrOutputWithContext(context.Background())
}

func (i *systemTopicEventSubscriptionPtrType) ToSystemTopicEventSubscriptionPtrOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionPtrOutput)
}

// SystemTopicEventSubscriptionArrayInput is an input type that accepts SystemTopicEventSubscriptionArray and SystemTopicEventSubscriptionArrayOutput values.
// You can construct a concrete instance of `SystemTopicEventSubscriptionArrayInput` via:
//
//          SystemTopicEventSubscriptionArray{ SystemTopicEventSubscriptionArgs{...} }
type SystemTopicEventSubscriptionArrayInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionArrayOutput() SystemTopicEventSubscriptionArrayOutput
	ToSystemTopicEventSubscriptionArrayOutputWithContext(context.Context) SystemTopicEventSubscriptionArrayOutput
}

type SystemTopicEventSubscriptionArray []SystemTopicEventSubscriptionInput

func (SystemTopicEventSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemTopicEventSubscription)(nil))
}

func (i SystemTopicEventSubscriptionArray) ToSystemTopicEventSubscriptionArrayOutput() SystemTopicEventSubscriptionArrayOutput {
	return i.ToSystemTopicEventSubscriptionArrayOutputWithContext(context.Background())
}

func (i SystemTopicEventSubscriptionArray) ToSystemTopicEventSubscriptionArrayOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionArrayOutput)
}

// SystemTopicEventSubscriptionMapInput is an input type that accepts SystemTopicEventSubscriptionMap and SystemTopicEventSubscriptionMapOutput values.
// You can construct a concrete instance of `SystemTopicEventSubscriptionMapInput` via:
//
//          SystemTopicEventSubscriptionMap{ "key": SystemTopicEventSubscriptionArgs{...} }
type SystemTopicEventSubscriptionMapInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionMapOutput() SystemTopicEventSubscriptionMapOutput
	ToSystemTopicEventSubscriptionMapOutputWithContext(context.Context) SystemTopicEventSubscriptionMapOutput
}

type SystemTopicEventSubscriptionMap map[string]SystemTopicEventSubscriptionInput

func (SystemTopicEventSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemTopicEventSubscription)(nil))
}

func (i SystemTopicEventSubscriptionMap) ToSystemTopicEventSubscriptionMapOutput() SystemTopicEventSubscriptionMapOutput {
	return i.ToSystemTopicEventSubscriptionMapOutputWithContext(context.Background())
}

func (i SystemTopicEventSubscriptionMap) ToSystemTopicEventSubscriptionMapOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionMapOutput)
}

type SystemTopicEventSubscriptionOutput struct {
	*pulumi.OutputState
}

func (SystemTopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicEventSubscription)(nil))
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return o
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return o
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionPtrOutput() SystemTopicEventSubscriptionPtrOutput {
	return o.ToSystemTopicEventSubscriptionPtrOutputWithContext(context.Background())
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionPtrOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionPtrOutput {
	return o.ApplyT(func(v SystemTopicEventSubscription) *SystemTopicEventSubscription {
		return &v
	}).(SystemTopicEventSubscriptionPtrOutput)
}

type SystemTopicEventSubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (SystemTopicEventSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicEventSubscription)(nil))
}

func (o SystemTopicEventSubscriptionPtrOutput) ToSystemTopicEventSubscriptionPtrOutput() SystemTopicEventSubscriptionPtrOutput {
	return o
}

func (o SystemTopicEventSubscriptionPtrOutput) ToSystemTopicEventSubscriptionPtrOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionPtrOutput {
	return o
}

type SystemTopicEventSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SystemTopicEventSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemTopicEventSubscription)(nil))
}

func (o SystemTopicEventSubscriptionArrayOutput) ToSystemTopicEventSubscriptionArrayOutput() SystemTopicEventSubscriptionArrayOutput {
	return o
}

func (o SystemTopicEventSubscriptionArrayOutput) ToSystemTopicEventSubscriptionArrayOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionArrayOutput {
	return o
}

func (o SystemTopicEventSubscriptionArrayOutput) Index(i pulumi.IntInput) SystemTopicEventSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemTopicEventSubscription {
		return vs[0].([]SystemTopicEventSubscription)[vs[1].(int)]
	}).(SystemTopicEventSubscriptionOutput)
}

type SystemTopicEventSubscriptionMapOutput struct{ *pulumi.OutputState }

func (SystemTopicEventSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemTopicEventSubscription)(nil))
}

func (o SystemTopicEventSubscriptionMapOutput) ToSystemTopicEventSubscriptionMapOutput() SystemTopicEventSubscriptionMapOutput {
	return o
}

func (o SystemTopicEventSubscriptionMapOutput) ToSystemTopicEventSubscriptionMapOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionMapOutput {
	return o
}

func (o SystemTopicEventSubscriptionMapOutput) MapIndex(k pulumi.StringInput) SystemTopicEventSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemTopicEventSubscription {
		return vs[0].(map[string]SystemTopicEventSubscription)[vs[1].(string)]
	}).(SystemTopicEventSubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionOutput{})
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionMapOutput{})
}
