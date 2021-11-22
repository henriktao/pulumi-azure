// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Stream Analytics Output to an EventHub.
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
// 		_, err = streamanalytics.NewOutputEventHub(ctx, "exampleOutputEventHub", &streamanalytics.OutputEventHubArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			EventhubName:           exampleEventHub.Name,
// 			ServicebusNamespace:    exampleEventHubNamespace.Name,
// 			SharedAccessPolicyKey:  exampleEventHubNamespace.DefaultPrimaryKey,
// 			SharedAccessPolicyName: pulumi.String("RootManageSharedAccessKey"),
// 			Serialization: &streamanalytics.OutputEventHubSerializationArgs{
// 				Type: pulumi.String("Avro"),
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
// Stream Analytics Outputs to an EventHub can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/outputEventHub:OutputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
// ```
type OutputEventHub struct {
	pulumi.CustomResourceState

	// The name of the Event Hub.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The column that is used for the Event Hub partition key.
	PartitionKey pulumi.StringPtrOutput `pulumi:"partitionKey"`
	// A list of property columns to add to the Event Hub output.
	PropertyColumns pulumi.StringArrayOutput `pulumi:"propertyColumns"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationOutput `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringOutput `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringOutput `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
}

// NewOutputEventHub registers a new resource with the given unique name, arguments, and options.
func NewOutputEventHub(ctx *pulumi.Context,
	name string, args *OutputEventHubArgs, opts ...pulumi.ResourceOption) (*OutputEventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource OutputEventHub
	err := ctx.RegisterResource("azure:streamanalytics/outputEventHub:OutputEventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutputEventHub gets an existing OutputEventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutputEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputEventHubState, opts ...pulumi.ResourceOption) (*OutputEventHub, error) {
	var resource OutputEventHub
	err := ctx.ReadResource("azure:streamanalytics/outputEventHub:OutputEventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutputEventHub resources.
type outputEventHubState struct {
	// The name of the Event Hub.
	EventhubName *string `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The column that is used for the Event Hub partition key.
	PartitionKey *string `pulumi:"partitionKey"`
	// A list of property columns to add to the Event Hub output.
	PropertyColumns []string `pulumi:"propertyColumns"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *OutputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace *string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey *string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
}

type OutputEventHubState struct {
	// The name of the Event Hub.
	EventhubName pulumi.StringPtrInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The column that is used for the Event Hub partition key.
	PartitionKey pulumi.StringPtrInput
	// A list of property columns to add to the Event Hub output.
	PropertyColumns pulumi.StringArrayInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationPtrInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringPtrInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringPtrInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
}

func (OutputEventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputEventHubState)(nil)).Elem()
}

type outputEventHubArgs struct {
	// The name of the Event Hub.
	EventhubName string `pulumi:"eventhubName"`
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The column that is used for the Event Hub partition key.
	PartitionKey *string `pulumi:"partitionKey"`
	// A list of property columns to add to the Event Hub output.
	PropertyColumns []string `pulumi:"propertyColumns"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerialization `pulumi:"serialization"`
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace string `pulumi:"servicebusNamespace"`
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey string `pulumi:"sharedAccessPolicyKey"`
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
}

// The set of arguments for constructing a OutputEventHub resource.
type OutputEventHubArgs struct {
	// The name of the Event Hub.
	EventhubName pulumi.StringInput
	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The column that is used for the Event Hub partition key.
	PartitionKey pulumi.StringPtrInput
	// A list of property columns to add to the Event Hub output.
	PropertyColumns pulumi.StringArrayInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization OutputEventHubSerializationInput
	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServicebusNamespace pulumi.StringInput
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey pulumi.StringInput
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
}

func (OutputEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputEventHubArgs)(nil)).Elem()
}

type OutputEventHubInput interface {
	pulumi.Input

	ToOutputEventHubOutput() OutputEventHubOutput
	ToOutputEventHubOutputWithContext(ctx context.Context) OutputEventHubOutput
}

func (*OutputEventHub) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputEventHub)(nil))
}

func (i *OutputEventHub) ToOutputEventHubOutput() OutputEventHubOutput {
	return i.ToOutputEventHubOutputWithContext(context.Background())
}

func (i *OutputEventHub) ToOutputEventHubOutputWithContext(ctx context.Context) OutputEventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputEventHubOutput)
}

func (i *OutputEventHub) ToOutputEventHubPtrOutput() OutputEventHubPtrOutput {
	return i.ToOutputEventHubPtrOutputWithContext(context.Background())
}

func (i *OutputEventHub) ToOutputEventHubPtrOutputWithContext(ctx context.Context) OutputEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputEventHubPtrOutput)
}

type OutputEventHubPtrInput interface {
	pulumi.Input

	ToOutputEventHubPtrOutput() OutputEventHubPtrOutput
	ToOutputEventHubPtrOutputWithContext(ctx context.Context) OutputEventHubPtrOutput
}

type outputEventHubPtrType OutputEventHubArgs

func (*outputEventHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputEventHub)(nil))
}

func (i *outputEventHubPtrType) ToOutputEventHubPtrOutput() OutputEventHubPtrOutput {
	return i.ToOutputEventHubPtrOutputWithContext(context.Background())
}

func (i *outputEventHubPtrType) ToOutputEventHubPtrOutputWithContext(ctx context.Context) OutputEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputEventHubPtrOutput)
}

// OutputEventHubArrayInput is an input type that accepts OutputEventHubArray and OutputEventHubArrayOutput values.
// You can construct a concrete instance of `OutputEventHubArrayInput` via:
//
//          OutputEventHubArray{ OutputEventHubArgs{...} }
type OutputEventHubArrayInput interface {
	pulumi.Input

	ToOutputEventHubArrayOutput() OutputEventHubArrayOutput
	ToOutputEventHubArrayOutputWithContext(context.Context) OutputEventHubArrayOutput
}

type OutputEventHubArray []OutputEventHubInput

func (OutputEventHubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OutputEventHub)(nil)).Elem()
}

func (i OutputEventHubArray) ToOutputEventHubArrayOutput() OutputEventHubArrayOutput {
	return i.ToOutputEventHubArrayOutputWithContext(context.Background())
}

func (i OutputEventHubArray) ToOutputEventHubArrayOutputWithContext(ctx context.Context) OutputEventHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputEventHubArrayOutput)
}

// OutputEventHubMapInput is an input type that accepts OutputEventHubMap and OutputEventHubMapOutput values.
// You can construct a concrete instance of `OutputEventHubMapInput` via:
//
//          OutputEventHubMap{ "key": OutputEventHubArgs{...} }
type OutputEventHubMapInput interface {
	pulumi.Input

	ToOutputEventHubMapOutput() OutputEventHubMapOutput
	ToOutputEventHubMapOutputWithContext(context.Context) OutputEventHubMapOutput
}

type OutputEventHubMap map[string]OutputEventHubInput

func (OutputEventHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OutputEventHub)(nil)).Elem()
}

func (i OutputEventHubMap) ToOutputEventHubMapOutput() OutputEventHubMapOutput {
	return i.ToOutputEventHubMapOutputWithContext(context.Background())
}

func (i OutputEventHubMap) ToOutputEventHubMapOutputWithContext(ctx context.Context) OutputEventHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputEventHubMapOutput)
}

type OutputEventHubOutput struct{ *pulumi.OutputState }

func (OutputEventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputEventHub)(nil))
}

func (o OutputEventHubOutput) ToOutputEventHubOutput() OutputEventHubOutput {
	return o
}

func (o OutputEventHubOutput) ToOutputEventHubOutputWithContext(ctx context.Context) OutputEventHubOutput {
	return o
}

func (o OutputEventHubOutput) ToOutputEventHubPtrOutput() OutputEventHubPtrOutput {
	return o.ToOutputEventHubPtrOutputWithContext(context.Background())
}

func (o OutputEventHubOutput) ToOutputEventHubPtrOutputWithContext(ctx context.Context) OutputEventHubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputEventHub) *OutputEventHub {
		return &v
	}).(OutputEventHubPtrOutput)
}

type OutputEventHubPtrOutput struct{ *pulumi.OutputState }

func (OutputEventHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputEventHub)(nil))
}

func (o OutputEventHubPtrOutput) ToOutputEventHubPtrOutput() OutputEventHubPtrOutput {
	return o
}

func (o OutputEventHubPtrOutput) ToOutputEventHubPtrOutputWithContext(ctx context.Context) OutputEventHubPtrOutput {
	return o
}

func (o OutputEventHubPtrOutput) Elem() OutputEventHubOutput {
	return o.ApplyT(func(v *OutputEventHub) OutputEventHub {
		if v != nil {
			return *v
		}
		var ret OutputEventHub
		return ret
	}).(OutputEventHubOutput)
}

type OutputEventHubArrayOutput struct{ *pulumi.OutputState }

func (OutputEventHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputEventHub)(nil))
}

func (o OutputEventHubArrayOutput) ToOutputEventHubArrayOutput() OutputEventHubArrayOutput {
	return o
}

func (o OutputEventHubArrayOutput) ToOutputEventHubArrayOutputWithContext(ctx context.Context) OutputEventHubArrayOutput {
	return o
}

func (o OutputEventHubArrayOutput) Index(i pulumi.IntInput) OutputEventHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputEventHub {
		return vs[0].([]OutputEventHub)[vs[1].(int)]
	}).(OutputEventHubOutput)
}

type OutputEventHubMapOutput struct{ *pulumi.OutputState }

func (OutputEventHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputEventHub)(nil))
}

func (o OutputEventHubMapOutput) ToOutputEventHubMapOutput() OutputEventHubMapOutput {
	return o
}

func (o OutputEventHubMapOutput) ToOutputEventHubMapOutputWithContext(ctx context.Context) OutputEventHubMapOutput {
	return o
}

func (o OutputEventHubMapOutput) MapIndex(k pulumi.StringInput) OutputEventHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputEventHub {
		return vs[0].(map[string]OutputEventHub)[vs[1].(string)]
	}).(OutputEventHubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OutputEventHubInput)(nil)).Elem(), &OutputEventHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputEventHubPtrInput)(nil)).Elem(), &OutputEventHub{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputEventHubArrayInput)(nil)).Elem(), OutputEventHubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputEventHubMapInput)(nil)).Elem(), OutputEventHubMap{})
	pulumi.RegisterOutputType(OutputEventHubOutput{})
	pulumi.RegisterOutputType(OutputEventHubPtrOutput{})
	pulumi.RegisterOutputType(OutputEventHubArrayOutput{})
	pulumi.RegisterOutputType(OutputEventHubMapOutput{})
}
