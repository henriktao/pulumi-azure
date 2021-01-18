// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EventGrid Topic
//
// > **Note:** at this time EventGrid Topic's are only available in a limited number of regions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US 2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewTopic(ctx, "exampleTopic", &eventgrid.TopicArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// EventGrid Topic's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventGridTopic:EventGridTopic topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
// ```
//
// Deprecated: azure.eventhub.EventGridTopic has been deprecated in favor of azure.eventgrid.Topic
type EventGridTopic struct {
	pulumi.CustomResourceState

	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules EventGridTopicInboundIpRuleArrayOutput `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrOutput `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrOutput `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEventGridTopic registers a new resource with the given unique name, arguments, and options.
func NewEventGridTopic(ctx *pulumi.Context,
	name string, args *EventGridTopicArgs, opts ...pulumi.ResourceOption) (*EventGridTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EventGridTopic
	err := ctx.RegisterResource("azure:eventhub/eventGridTopic:EventGridTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventGridTopic gets an existing EventGridTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventGridTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridTopicState, opts ...pulumi.ResourceOption) (*EventGridTopic, error) {
	var resource EventGridTopic
	err := ctx.ReadResource("azure:eventhub/eventGridTopic:EventGridTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventGridTopic resources.
type eventGridTopicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint *string `pulumi:"endpoint"`
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules []EventGridTopicInboundIpRule `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *EventGridTopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *EventGridTopicInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EventGridTopicState struct {
	// The Endpoint associated with the EventGrid Topic.
	Endpoint pulumi.StringPtrInput
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules EventGridTopicInboundIpRuleArrayInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Primary Shared Access Key associated with the EventGrid Topic.
	PrimaryAccessKey pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Shared Access Key associated with the EventGrid Topic.
	SecondaryAccessKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventGridTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridTopicState)(nil)).Elem()
}

type eventGridTopicArgs struct {
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules []EventGridTopicInboundIpRule `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *EventGridTopicInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *EventGridTopicInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventGridTopic resource.
type EventGridTopicArgs struct {
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules EventGridTopicInboundIpRuleArrayInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues EventGridTopicInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields EventGridTopicInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventGridTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridTopicArgs)(nil)).Elem()
}

type EventGridTopicInput interface {
	pulumi.Input

	ToEventGridTopicOutput() EventGridTopicOutput
	ToEventGridTopicOutputWithContext(ctx context.Context) EventGridTopicOutput
}

func (EventGridTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridTopic)(nil)).Elem()
}

func (i EventGridTopic) ToEventGridTopicOutput() EventGridTopicOutput {
	return i.ToEventGridTopicOutputWithContext(context.Background())
}

func (i EventGridTopic) ToEventGridTopicOutputWithContext(ctx context.Context) EventGridTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridTopicOutput)
}

type EventGridTopicOutput struct {
	*pulumi.OutputState
}

func (EventGridTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridTopicOutput)(nil)).Elem()
}

func (o EventGridTopicOutput) ToEventGridTopicOutput() EventGridTopicOutput {
	return o
}

func (o EventGridTopicOutput) ToEventGridTopicOutputWithContext(ctx context.Context) EventGridTopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventGridTopicOutput{})
}
