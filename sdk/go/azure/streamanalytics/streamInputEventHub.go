// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Stream Analytics Stream Input EventHub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConsumerGroup, err := eventhub.NewConsumerGroup(ctx, "exampleConsumerGroup", &eventhub.ConsumerGroupArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewStreamInputEventHub(ctx, "exampleStreamInputEventHub", &streamanalytics.StreamInputEventHubArgs{
// 			StreamAnalyticsJobName:    pulumi.String(exampleJob.Name),
// 			ResourceGroupName:         pulumi.String(exampleJob.ResourceGroupName),
// 			EventhubConsumerGroupName: exampleConsumerGroup.Name,
// 			EventhubName:              exampleEventHub.Name,
// 			ServicebusNamespace:       exampleEventHubNamespace.Name,
// 			SharedAccessPolicyKey:     exampleEventHubNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName:    pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.StreamInputEventHubSerializationArgs{
// 				Type:     pulumi.String("Json"),
// 				Encoding: pulumi.String("UTF8"),
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
// Stream Analytics Stream Input EventHub's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/streamInputEventHub:StreamInputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
// ```
type StreamInputEventHub struct {
	pulumi.CustomResourceState

	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringOutput `pulumi:"eventhubConsumerGroupName"`
	// The name of the Event Hub.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputEventHubSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewStreamInputEventHub registers a new resource with the given unique name, arguments, and options.
func NewStreamInputEventHub(ctx *pulumi.Context,
	name string, args *StreamInputEventHubArgs, opts ...pulumi.ResourceOption) (*StreamInputEventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubConsumerGroupName'")
	}
	if args.EventhubName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Serialization == nil {
		return nil, errors.New("invalid value for required argument 'Serialization'")
	}
	if args.ServicebusNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServicebusNamespace'")
	}
	if args.SharedAccessPolicyKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyKey'")
	}
	if args.SharedAccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyName'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	var resource StreamInputEventHub
	err := ctx.RegisterResource("azure:streamanalytics/streamInputEventHub:StreamInputEventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamInputEventHub gets an existing StreamInputEventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamInputEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamInputEventHubState, opts ...pulumi.ResourceOption) (*StreamInputEventHub, error) {
	var resource StreamInputEventHub
	err := ctx.ReadResource("azure:streamanalytics/streamInputEventHub:StreamInputEventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamInputEventHub resources.
type streamInputEventHubState struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName *string `pulumi:"eventhubConsumerGroupName"`
	// The name of the Event Hub.
	EventhubName *string `pulumi:"eventhubName"`
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *StreamInputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type StreamInputEventHubState struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringPtrInput
	// The name of the Event Hub.
	EventhubName pulumi.StringPtrInput
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization StreamInputEventHubSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (StreamInputEventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputEventHubState)(nil)).Elem()
}

type streamInputEventHubArgs struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName string `pulumi:"eventhubConsumerGroupName"`
	// The name of the Event Hub.
	EventhubName string `pulumi:"eventhubName"`
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a StreamInputEventHub resource.
type StreamInputEventHubArgs struct {
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName pulumi.StringInput
	// The name of the Event Hub.
	EventhubName pulumi.StringInput
	// The name of the Stream Input EventHub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization StreamInputEventHubSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (StreamInputEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputEventHubArgs)(nil)).Elem()
}

type StreamInputEventHubInput interface {
	pulumi.Input

	ToStreamInputEventHubOutput() StreamInputEventHubOutput
	ToStreamInputEventHubOutputWithContext(ctx context.Context) StreamInputEventHubOutput
}

func (*StreamInputEventHub) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputEventHub)(nil))
}

func (i *StreamInputEventHub) ToStreamInputEventHubOutput() StreamInputEventHubOutput {
	return i.ToStreamInputEventHubOutputWithContext(context.Background())
}

func (i *StreamInputEventHub) ToStreamInputEventHubOutputWithContext(ctx context.Context) StreamInputEventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputEventHubOutput)
}

func (i *StreamInputEventHub) ToStreamInputEventHubPtrOutput() StreamInputEventHubPtrOutput {
	return i.ToStreamInputEventHubPtrOutputWithContext(context.Background())
}

func (i *StreamInputEventHub) ToStreamInputEventHubPtrOutputWithContext(ctx context.Context) StreamInputEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputEventHubPtrOutput)
}

type StreamInputEventHubPtrInput interface {
	pulumi.Input

	ToStreamInputEventHubPtrOutput() StreamInputEventHubPtrOutput
	ToStreamInputEventHubPtrOutputWithContext(ctx context.Context) StreamInputEventHubPtrOutput
}

type streamInputEventHubPtrType StreamInputEventHubArgs

func (*streamInputEventHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamInputEventHub)(nil))
}

func (i *streamInputEventHubPtrType) ToStreamInputEventHubPtrOutput() StreamInputEventHubPtrOutput {
	return i.ToStreamInputEventHubPtrOutputWithContext(context.Background())
}

func (i *streamInputEventHubPtrType) ToStreamInputEventHubPtrOutputWithContext(ctx context.Context) StreamInputEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputEventHubPtrOutput)
}

// StreamInputEventHubArrayInput is an input type that accepts StreamInputEventHubArray and StreamInputEventHubArrayOutput values.
// You can construct a concrete instance of `StreamInputEventHubArrayInput` via:
//
//          StreamInputEventHubArray{ StreamInputEventHubArgs{...} }
type StreamInputEventHubArrayInput interface {
	pulumi.Input

	ToStreamInputEventHubArrayOutput() StreamInputEventHubArrayOutput
	ToStreamInputEventHubArrayOutputWithContext(context.Context) StreamInputEventHubArrayOutput
}

type StreamInputEventHubArray []StreamInputEventHubInput

func (StreamInputEventHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamInputEventHub)(nil)).Elem()
}

func (i StreamInputEventHubArray) ToStreamInputEventHubArrayOutput() StreamInputEventHubArrayOutput {
	return i.ToStreamInputEventHubArrayOutputWithContext(context.Background())
}

func (i StreamInputEventHubArray) ToStreamInputEventHubArrayOutputWithContext(ctx context.Context) StreamInputEventHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputEventHubArrayOutput)
}

// StreamInputEventHubMapInput is an input type that accepts StreamInputEventHubMap and StreamInputEventHubMapOutput values.
// You can construct a concrete instance of `StreamInputEventHubMapInput` via:
//
//          StreamInputEventHubMap{ "key": StreamInputEventHubArgs{...} }
type StreamInputEventHubMapInput interface {
	pulumi.Input

	ToStreamInputEventHubMapOutput() StreamInputEventHubMapOutput
	ToStreamInputEventHubMapOutputWithContext(context.Context) StreamInputEventHubMapOutput
}

type StreamInputEventHubMap map[string]StreamInputEventHubInput

func (StreamInputEventHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamInputEventHub)(nil)).Elem()
}

func (i StreamInputEventHubMap) ToStreamInputEventHubMapOutput() StreamInputEventHubMapOutput {
	return i.ToStreamInputEventHubMapOutputWithContext(context.Background())
}

func (i StreamInputEventHubMap) ToStreamInputEventHubMapOutputWithContext(ctx context.Context) StreamInputEventHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputEventHubMapOutput)
}

type StreamInputEventHubOutput struct{ *pulumi.OutputState }

func (StreamInputEventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputEventHub)(nil))
}

func (o StreamInputEventHubOutput) ToStreamInputEventHubOutput() StreamInputEventHubOutput {
	return o
}

func (o StreamInputEventHubOutput) ToStreamInputEventHubOutputWithContext(ctx context.Context) StreamInputEventHubOutput {
	return o
}

func (o StreamInputEventHubOutput) ToStreamInputEventHubPtrOutput() StreamInputEventHubPtrOutput {
	return o.ToStreamInputEventHubPtrOutputWithContext(context.Background())
}

func (o StreamInputEventHubOutput) ToStreamInputEventHubPtrOutputWithContext(ctx context.Context) StreamInputEventHubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamInputEventHub) *StreamInputEventHub {
		return &v
	}).(StreamInputEventHubPtrOutput)
}

type StreamInputEventHubPtrOutput struct{ *pulumi.OutputState }

func (StreamInputEventHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamInputEventHub)(nil))
}

func (o StreamInputEventHubPtrOutput) ToStreamInputEventHubPtrOutput() StreamInputEventHubPtrOutput {
	return o
}

func (o StreamInputEventHubPtrOutput) ToStreamInputEventHubPtrOutputWithContext(ctx context.Context) StreamInputEventHubPtrOutput {
	return o
}

func (o StreamInputEventHubPtrOutput) Elem() StreamInputEventHubOutput {
	return o.ApplyT(func(v *StreamInputEventHub) StreamInputEventHub {
		if v != nil {
			return *v
		}
		var ret StreamInputEventHub
		return ret
	}).(StreamInputEventHubOutput)
}

type StreamInputEventHubArrayOutput struct{ *pulumi.OutputState }

func (StreamInputEventHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamInputEventHub)(nil))
}

func (o StreamInputEventHubArrayOutput) ToStreamInputEventHubArrayOutput() StreamInputEventHubArrayOutput {
	return o
}

func (o StreamInputEventHubArrayOutput) ToStreamInputEventHubArrayOutputWithContext(ctx context.Context) StreamInputEventHubArrayOutput {
	return o
}

func (o StreamInputEventHubArrayOutput) Index(i pulumi.IntInput) StreamInputEventHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamInputEventHub {
		return vs[0].([]StreamInputEventHub)[vs[1].(int)]
	}).(StreamInputEventHubOutput)
}

type StreamInputEventHubMapOutput struct{ *pulumi.OutputState }

func (StreamInputEventHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamInputEventHub)(nil))
}

func (o StreamInputEventHubMapOutput) ToStreamInputEventHubMapOutput() StreamInputEventHubMapOutput {
	return o
}

func (o StreamInputEventHubMapOutput) ToStreamInputEventHubMapOutputWithContext(ctx context.Context) StreamInputEventHubMapOutput {
	return o
}

func (o StreamInputEventHubMapOutput) MapIndex(k pulumi.StringInput) StreamInputEventHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StreamInputEventHub {
		return vs[0].(map[string]StreamInputEventHub)[vs[1].(string)]
	}).(StreamInputEventHubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInputEventHubInput)(nil)).Elem(), &StreamInputEventHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInputEventHubPtrInput)(nil)).Elem(), &StreamInputEventHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInputEventHubArrayInput)(nil)).Elem(), StreamInputEventHubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInputEventHubMapInput)(nil)).Elem(), StreamInputEventHubMap{})
	pulumi.RegisterOutputType(StreamInputEventHubOutput{})
	pulumi.RegisterOutputType(StreamInputEventHubPtrOutput{})
	pulumi.RegisterOutputType(StreamInputEventHubArrayOutput{})
	pulumi.RegisterOutputType(StreamInputEventHubMapOutput{})
}
