// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure IoT Time Series Insights IoTHub Event Source.
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
// 			Location: pulumi.String("West Eurpoe"),
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
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConsumerGroup, err := iot.NewConsumerGroup(ctx, "exampleConsumerGroup", &iot.ConsumerGroupArgs{
// 			IothubName:           exampleIoTHub.Name,
// 			EventhubEndpointName: pulumi.String("events"),
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		storage, err := storage.NewAccount(ctx, "storage", &storage.AccountArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTimeSeriesInsightsGen2Environment, err := iot.NewTimeSeriesInsightsGen2Environment(ctx, "exampleTimeSeriesInsightsGen2Environment", &iot.TimeSeriesInsightsGen2EnvironmentArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("L1"),
// 			IdProperties: pulumi.StringArray{
// 				pulumi.String("id"),
// 			},
// 			Storage: &iot.TimeSeriesInsightsGen2EnvironmentStorageArgs{
// 				Name: storage.Name,
// 				Key:  storage.PrimaryAccessKey,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewTimeSeriesInsightsEventSourceIothub(ctx, "exampleTimeSeriesInsightsEventSourceIothub", &iot.TimeSeriesInsightsEventSourceIothubArgs{
// 			Location:      exampleResourceGroup.Location,
// 			EnvironmentId: exampleTimeSeriesInsightsGen2Environment.ID(),
// 			IothubName:    exampleIoTHub.Name,
// 			SharedAccessKey: exampleIoTHub.SharedAccessPolicies.ApplyT(func(sharedAccessPolicies []iot.IoTHubSharedAccessPolicy) (string, error) {
// 				return sharedAccessPolicies[0].PrimaryKey, nil
// 			}).(pulumi.StringOutput),
// 			SharedAccessKeyName: exampleIoTHub.SharedAccessPolicies.ApplyT(func(sharedAccessPolicies []iot.IoTHubSharedAccessPolicy) (string, error) {
// 				return sharedAccessPolicies[0].KeyName, nil
// 			}).(pulumi.StringOutput),
// 			ConsumerGroupName:     exampleConsumerGroup.Name,
// 			EventSourceResourceId: exampleIoTHub.ID(),
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
// Azure IoT Time Series Insights IoTHub Event Source can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/timeSeriesInsightsEventSourceIothub:TimeSeriesInsightsEventSourceIothub example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
// ```
type TimeSeriesInsightsEventSourceIothub struct {
	pulumi.CustomResourceState

	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName pulumi.StringOutput `pulumi:"consumerGroupName"`
	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// Specifies the resource id where events will be coming from.
	EventSourceResourceId pulumi.StringOutput `pulumi:"eventSourceResourceId"`
	// Specifies the name of the IotHub which will be associated with this resource.
	IothubName pulumi.StringOutput `pulumi:"iothubName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights IoTHub Event Source. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	SharedAccessKey pulumi.StringOutput `pulumi:"sharedAccessKey"`
	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	SharedAccessKeyName pulumi.StringOutput `pulumi:"sharedAccessKeyName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName pulumi.StringOutput `pulumi:"timestampPropertyName"`
}

// NewTimeSeriesInsightsEventSourceIothub registers a new resource with the given unique name, arguments, and options.
func NewTimeSeriesInsightsEventSourceIothub(ctx *pulumi.Context,
	name string, args *TimeSeriesInsightsEventSourceIothubArgs, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsEventSourceIothub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.EventSourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceResourceId'")
	}
	if args.IothubName == nil {
		return nil, errors.New("invalid value for required argument 'IothubName'")
	}
	if args.SharedAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKey'")
	}
	if args.SharedAccessKeyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKeyName'")
	}
	var resource TimeSeriesInsightsEventSourceIothub
	err := ctx.RegisterResource("azure:iot/timeSeriesInsightsEventSourceIothub:TimeSeriesInsightsEventSourceIothub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeriesInsightsEventSourceIothub gets an existing TimeSeriesInsightsEventSourceIothub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeriesInsightsEventSourceIothub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesInsightsEventSourceIothubState, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsEventSourceIothub, error) {
	var resource TimeSeriesInsightsEventSourceIothub
	err := ctx.ReadResource("azure:iot/timeSeriesInsightsEventSourceIothub:TimeSeriesInsightsEventSourceIothub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeriesInsightsEventSourceIothub resources.
type timeSeriesInsightsEventSourceIothubState struct {
	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName *string `pulumi:"consumerGroupName"`
	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentId *string `pulumi:"environmentId"`
	// Specifies the resource id where events will be coming from.
	EventSourceResourceId *string `pulumi:"eventSourceResourceId"`
	// Specifies the name of the IotHub which will be associated with this resource.
	IothubName *string `pulumi:"iothubName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights IoTHub Event Source. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	SharedAccessKey *string `pulumi:"sharedAccessKey"`
	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	SharedAccessKeyName *string `pulumi:"sharedAccessKeyName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName *string `pulumi:"timestampPropertyName"`
}

type TimeSeriesInsightsEventSourceIothubState struct {
	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName pulumi.StringPtrInput
	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentId pulumi.StringPtrInput
	// Specifies the resource id where events will be coming from.
	EventSourceResourceId pulumi.StringPtrInput
	// Specifies the name of the IotHub which will be associated with this resource.
	IothubName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights IoTHub Event Source. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	SharedAccessKey pulumi.StringPtrInput
	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	SharedAccessKeyName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName pulumi.StringPtrInput
}

func (TimeSeriesInsightsEventSourceIothubState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsEventSourceIothubState)(nil)).Elem()
}

type timeSeriesInsightsEventSourceIothubArgs struct {
	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentId string `pulumi:"environmentId"`
	// Specifies the resource id where events will be coming from.
	EventSourceResourceId string `pulumi:"eventSourceResourceId"`
	// Specifies the name of the IotHub which will be associated with this resource.
	IothubName string `pulumi:"iothubName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights IoTHub Event Source. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	SharedAccessKey string `pulumi:"sharedAccessKey"`
	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	SharedAccessKeyName string `pulumi:"sharedAccessKeyName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName *string `pulumi:"timestampPropertyName"`
}

// The set of arguments for constructing a TimeSeriesInsightsEventSourceIothub resource.
type TimeSeriesInsightsEventSourceIothubArgs struct {
	// Specifies the name of the IotHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName pulumi.StringInput
	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentId pulumi.StringInput
	// Specifies the resource id where events will be coming from.
	EventSourceResourceId pulumi.StringInput
	// Specifies the name of the IotHub which will be associated with this resource.
	IothubName pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights IoTHub Event Source. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the IotHub.
	SharedAccessKey pulumi.StringInput
	// Specifies the name of the Shared Access key that grants the Event Source access to the IotHub.
	SharedAccessKeyName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName pulumi.StringPtrInput
}

func (TimeSeriesInsightsEventSourceIothubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsEventSourceIothubArgs)(nil)).Elem()
}

type TimeSeriesInsightsEventSourceIothubInput interface {
	pulumi.Input

	ToTimeSeriesInsightsEventSourceIothubOutput() TimeSeriesInsightsEventSourceIothubOutput
	ToTimeSeriesInsightsEventSourceIothubOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubOutput
}

func (*TimeSeriesInsightsEventSourceIothub) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsEventSourceIothub)(nil))
}

func (i *TimeSeriesInsightsEventSourceIothub) ToTimeSeriesInsightsEventSourceIothubOutput() TimeSeriesInsightsEventSourceIothubOutput {
	return i.ToTimeSeriesInsightsEventSourceIothubOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsEventSourceIothub) ToTimeSeriesInsightsEventSourceIothubOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsEventSourceIothubOutput)
}

func (i *TimeSeriesInsightsEventSourceIothub) ToTimeSeriesInsightsEventSourceIothubPtrOutput() TimeSeriesInsightsEventSourceIothubPtrOutput {
	return i.ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsEventSourceIothub) ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsEventSourceIothubPtrOutput)
}

type TimeSeriesInsightsEventSourceIothubPtrInput interface {
	pulumi.Input

	ToTimeSeriesInsightsEventSourceIothubPtrOutput() TimeSeriesInsightsEventSourceIothubPtrOutput
	ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubPtrOutput
}

type timeSeriesInsightsEventSourceIothubPtrType TimeSeriesInsightsEventSourceIothubArgs

func (*timeSeriesInsightsEventSourceIothubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsEventSourceIothub)(nil))
}

func (i *timeSeriesInsightsEventSourceIothubPtrType) ToTimeSeriesInsightsEventSourceIothubPtrOutput() TimeSeriesInsightsEventSourceIothubPtrOutput {
	return i.ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(context.Background())
}

func (i *timeSeriesInsightsEventSourceIothubPtrType) ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsEventSourceIothubPtrOutput)
}

// TimeSeriesInsightsEventSourceIothubArrayInput is an input type that accepts TimeSeriesInsightsEventSourceIothubArray and TimeSeriesInsightsEventSourceIothubArrayOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsEventSourceIothubArrayInput` via:
//
//          TimeSeriesInsightsEventSourceIothubArray{ TimeSeriesInsightsEventSourceIothubArgs{...} }
type TimeSeriesInsightsEventSourceIothubArrayInput interface {
	pulumi.Input

	ToTimeSeriesInsightsEventSourceIothubArrayOutput() TimeSeriesInsightsEventSourceIothubArrayOutput
	ToTimeSeriesInsightsEventSourceIothubArrayOutputWithContext(context.Context) TimeSeriesInsightsEventSourceIothubArrayOutput
}

type TimeSeriesInsightsEventSourceIothubArray []TimeSeriesInsightsEventSourceIothubInput

func (TimeSeriesInsightsEventSourceIothubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TimeSeriesInsightsEventSourceIothub)(nil)).Elem()
}

func (i TimeSeriesInsightsEventSourceIothubArray) ToTimeSeriesInsightsEventSourceIothubArrayOutput() TimeSeriesInsightsEventSourceIothubArrayOutput {
	return i.ToTimeSeriesInsightsEventSourceIothubArrayOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsEventSourceIothubArray) ToTimeSeriesInsightsEventSourceIothubArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsEventSourceIothubArrayOutput)
}

// TimeSeriesInsightsEventSourceIothubMapInput is an input type that accepts TimeSeriesInsightsEventSourceIothubMap and TimeSeriesInsightsEventSourceIothubMapOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsEventSourceIothubMapInput` via:
//
//          TimeSeriesInsightsEventSourceIothubMap{ "key": TimeSeriesInsightsEventSourceIothubArgs{...} }
type TimeSeriesInsightsEventSourceIothubMapInput interface {
	pulumi.Input

	ToTimeSeriesInsightsEventSourceIothubMapOutput() TimeSeriesInsightsEventSourceIothubMapOutput
	ToTimeSeriesInsightsEventSourceIothubMapOutputWithContext(context.Context) TimeSeriesInsightsEventSourceIothubMapOutput
}

type TimeSeriesInsightsEventSourceIothubMap map[string]TimeSeriesInsightsEventSourceIothubInput

func (TimeSeriesInsightsEventSourceIothubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TimeSeriesInsightsEventSourceIothub)(nil)).Elem()
}

func (i TimeSeriesInsightsEventSourceIothubMap) ToTimeSeriesInsightsEventSourceIothubMapOutput() TimeSeriesInsightsEventSourceIothubMapOutput {
	return i.ToTimeSeriesInsightsEventSourceIothubMapOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsEventSourceIothubMap) ToTimeSeriesInsightsEventSourceIothubMapOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsEventSourceIothubMapOutput)
}

type TimeSeriesInsightsEventSourceIothubOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsEventSourceIothubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsEventSourceIothub)(nil))
}

func (o TimeSeriesInsightsEventSourceIothubOutput) ToTimeSeriesInsightsEventSourceIothubOutput() TimeSeriesInsightsEventSourceIothubOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubOutput) ToTimeSeriesInsightsEventSourceIothubOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubOutput) ToTimeSeriesInsightsEventSourceIothubPtrOutput() TimeSeriesInsightsEventSourceIothubPtrOutput {
	return o.ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(context.Background())
}

func (o TimeSeriesInsightsEventSourceIothubOutput) ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeSeriesInsightsEventSourceIothub) *TimeSeriesInsightsEventSourceIothub {
		return &v
	}).(TimeSeriesInsightsEventSourceIothubPtrOutput)
}

type TimeSeriesInsightsEventSourceIothubPtrOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsEventSourceIothubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsEventSourceIothub)(nil))
}

func (o TimeSeriesInsightsEventSourceIothubPtrOutput) ToTimeSeriesInsightsEventSourceIothubPtrOutput() TimeSeriesInsightsEventSourceIothubPtrOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubPtrOutput) ToTimeSeriesInsightsEventSourceIothubPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubPtrOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubPtrOutput) Elem() TimeSeriesInsightsEventSourceIothubOutput {
	return o.ApplyT(func(v *TimeSeriesInsightsEventSourceIothub) TimeSeriesInsightsEventSourceIothub {
		if v != nil {
			return *v
		}
		var ret TimeSeriesInsightsEventSourceIothub
		return ret
	}).(TimeSeriesInsightsEventSourceIothubOutput)
}

type TimeSeriesInsightsEventSourceIothubArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsEventSourceIothubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesInsightsEventSourceIothub)(nil))
}

func (o TimeSeriesInsightsEventSourceIothubArrayOutput) ToTimeSeriesInsightsEventSourceIothubArrayOutput() TimeSeriesInsightsEventSourceIothubArrayOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubArrayOutput) ToTimeSeriesInsightsEventSourceIothubArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubArrayOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubArrayOutput) Index(i pulumi.IntInput) TimeSeriesInsightsEventSourceIothubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesInsightsEventSourceIothub {
		return vs[0].([]TimeSeriesInsightsEventSourceIothub)[vs[1].(int)]
	}).(TimeSeriesInsightsEventSourceIothubOutput)
}

type TimeSeriesInsightsEventSourceIothubMapOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsEventSourceIothubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TimeSeriesInsightsEventSourceIothub)(nil))
}

func (o TimeSeriesInsightsEventSourceIothubMapOutput) ToTimeSeriesInsightsEventSourceIothubMapOutput() TimeSeriesInsightsEventSourceIothubMapOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubMapOutput) ToTimeSeriesInsightsEventSourceIothubMapOutputWithContext(ctx context.Context) TimeSeriesInsightsEventSourceIothubMapOutput {
	return o
}

func (o TimeSeriesInsightsEventSourceIothubMapOutput) MapIndex(k pulumi.StringInput) TimeSeriesInsightsEventSourceIothubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TimeSeriesInsightsEventSourceIothub {
		return vs[0].(map[string]TimeSeriesInsightsEventSourceIothub)[vs[1].(string)]
	}).(TimeSeriesInsightsEventSourceIothubOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsEventSourceIothubOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsEventSourceIothubPtrOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsEventSourceIothubArrayOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsEventSourceIothubMapOutput{})
}
