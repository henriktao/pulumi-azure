// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Custom Event Trigger inside an Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventgrid"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePipeline, err := datafactory.NewPipeline(ctx, "examplePipeline", &datafactory.PipelineArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := eventgrid.NewTopic(ctx, "exampleTopic", &eventgrid.TopicArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewTriggerCustomEvent(ctx, "exampleTriggerCustomEvent", &datafactory.TriggerCustomEventArgs{
// 			DataFactoryId:    exampleFactory.ID(),
// 			EventgridTopicId: exampleTopic.ID(),
// 			Events: pulumi.StringArray{
// 				pulumi.String("event1"),
// 				pulumi.String("event2"),
// 			},
// 			SubjectBeginsWith: pulumi.String("abc"),
// 			SubjectEndsWith:   pulumi.String("xyz"),
// 			Annotations: pulumi.StringArray{
// 				pulumi.String("example1"),
// 				pulumi.String("example2"),
// 				pulumi.String("example3"),
// 			},
// 			Description: pulumi.String("example description"),
// 			Pipelines: datafactory.TriggerCustomEventPipelineArray{
// 				&datafactory.TriggerCustomEventPipelineArgs{
// 					Name: examplePipeline.Name,
// 					Parameters: pulumi.StringMap{
// 						"Env": pulumi.String("Prod"),
// 					},
// 				},
// 			},
// 			AdditionalProperties: pulumi.StringMap{
// 				"foo": pulumi.String("foo1"),
// 				"bar": pulumi.String("bar2"),
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
// Data Factory Custom Event Trigger can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/triggerCustomEvent:TriggerCustomEvent example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
// ```
type TriggerCustomEvent struct {
	pulumi.CustomResourceState

	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to `true`.
	Activated pulumi.BoolPtrOutput `pulumi:"activated"`
	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The description for the Data Factory Custom Event Trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	EventgridTopicId pulumi.StringOutput `pulumi:"eventgridTopicId"`
	// List of events that will fire this trigger. At least one event must be specified.
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// Specifies the name of the Data Factory Custom Event Trigger. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `pipeline` blocks as defined below.
	Pipelines TriggerCustomEventPipelineArrayOutput `pulumi:"pipelines"`
	// The pattern that event subject starts with for trigger to fire.
	SubjectBeginsWith pulumi.StringPtrOutput `pulumi:"subjectBeginsWith"`
	// The pattern that event subject ends with for trigger to fire.
	SubjectEndsWith pulumi.StringPtrOutput `pulumi:"subjectEndsWith"`
}

// NewTriggerCustomEvent registers a new resource with the given unique name, arguments, and options.
func NewTriggerCustomEvent(ctx *pulumi.Context,
	name string, args *TriggerCustomEventArgs, opts ...pulumi.ResourceOption) (*TriggerCustomEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryId == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryId'")
	}
	if args.EventgridTopicId == nil {
		return nil, errors.New("invalid value for required argument 'EventgridTopicId'")
	}
	if args.Events == nil {
		return nil, errors.New("invalid value for required argument 'Events'")
	}
	if args.Pipelines == nil {
		return nil, errors.New("invalid value for required argument 'Pipelines'")
	}
	var resource TriggerCustomEvent
	err := ctx.RegisterResource("azure:datafactory/triggerCustomEvent:TriggerCustomEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerCustomEvent gets an existing TriggerCustomEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerCustomEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerCustomEventState, opts ...pulumi.ResourceOption) (*TriggerCustomEvent, error) {
	var resource TriggerCustomEvent
	err := ctx.ReadResource("azure:datafactory/triggerCustomEvent:TriggerCustomEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerCustomEvent resources.
type triggerCustomEventState struct {
	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to `true`.
	Activated *bool `pulumi:"activated"`
	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	Annotations []string `pulumi:"annotations"`
	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Custom Event Trigger.
	Description *string `pulumi:"description"`
	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	EventgridTopicId *string `pulumi:"eventgridTopicId"`
	// List of events that will fire this trigger. At least one event must be specified.
	Events []string `pulumi:"events"`
	// Specifies the name of the Data Factory Custom Event Trigger. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `pipeline` blocks as defined below.
	Pipelines []TriggerCustomEventPipeline `pulumi:"pipelines"`
	// The pattern that event subject starts with for trigger to fire.
	SubjectBeginsWith *string `pulumi:"subjectBeginsWith"`
	// The pattern that event subject ends with for trigger to fire.
	SubjectEndsWith *string `pulumi:"subjectEndsWith"`
}

type TriggerCustomEventState struct {
	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to `true`.
	Activated pulumi.BoolPtrInput
	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	Annotations pulumi.StringArrayInput
	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The description for the Data Factory Custom Event Trigger.
	Description pulumi.StringPtrInput
	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	EventgridTopicId pulumi.StringPtrInput
	// List of events that will fire this trigger. At least one event must be specified.
	Events pulumi.StringArrayInput
	// Specifies the name of the Data Factory Custom Event Trigger. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `pipeline` blocks as defined below.
	Pipelines TriggerCustomEventPipelineArrayInput
	// The pattern that event subject starts with for trigger to fire.
	SubjectBeginsWith pulumi.StringPtrInput
	// The pattern that event subject ends with for trigger to fire.
	SubjectEndsWith pulumi.StringPtrInput
}

func (TriggerCustomEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCustomEventState)(nil)).Elem()
}

type triggerCustomEventArgs struct {
	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to `true`.
	Activated *bool `pulumi:"activated"`
	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	Annotations []string `pulumi:"annotations"`
	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	DataFactoryId string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Custom Event Trigger.
	Description *string `pulumi:"description"`
	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	EventgridTopicId string `pulumi:"eventgridTopicId"`
	// List of events that will fire this trigger. At least one event must be specified.
	Events []string `pulumi:"events"`
	// Specifies the name of the Data Factory Custom Event Trigger. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `pipeline` blocks as defined below.
	Pipelines []TriggerCustomEventPipeline `pulumi:"pipelines"`
	// The pattern that event subject starts with for trigger to fire.
	SubjectBeginsWith *string `pulumi:"subjectBeginsWith"`
	// The pattern that event subject ends with for trigger to fire.
	SubjectEndsWith *string `pulumi:"subjectEndsWith"`
}

// The set of arguments for constructing a TriggerCustomEvent resource.
type TriggerCustomEventArgs struct {
	// Specifies if the Data Factory Custom Event Trigger is activated. Defaults to `true`.
	Activated pulumi.BoolPtrInput
	// A map of additional properties to associate with the Data Factory Custom Event Trigger.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Custom Event Trigger.
	Annotations pulumi.StringArrayInput
	// The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
	DataFactoryId pulumi.StringInput
	// The description for the Data Factory Custom Event Trigger.
	Description pulumi.StringPtrInput
	// The ID of Event Grid Topic in which event will be listened. Changing this forces a new resource.
	EventgridTopicId pulumi.StringInput
	// List of events that will fire this trigger. At least one event must be specified.
	Events pulumi.StringArrayInput
	// Specifies the name of the Data Factory Custom Event Trigger. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `pipeline` blocks as defined below.
	Pipelines TriggerCustomEventPipelineArrayInput
	// The pattern that event subject starts with for trigger to fire.
	SubjectBeginsWith pulumi.StringPtrInput
	// The pattern that event subject ends with for trigger to fire.
	SubjectEndsWith pulumi.StringPtrInput
}

func (TriggerCustomEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerCustomEventArgs)(nil)).Elem()
}

type TriggerCustomEventInput interface {
	pulumi.Input

	ToTriggerCustomEventOutput() TriggerCustomEventOutput
	ToTriggerCustomEventOutputWithContext(ctx context.Context) TriggerCustomEventOutput
}

func (*TriggerCustomEvent) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCustomEvent)(nil))
}

func (i *TriggerCustomEvent) ToTriggerCustomEventOutput() TriggerCustomEventOutput {
	return i.ToTriggerCustomEventOutputWithContext(context.Background())
}

func (i *TriggerCustomEvent) ToTriggerCustomEventOutputWithContext(ctx context.Context) TriggerCustomEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomEventOutput)
}

func (i *TriggerCustomEvent) ToTriggerCustomEventPtrOutput() TriggerCustomEventPtrOutput {
	return i.ToTriggerCustomEventPtrOutputWithContext(context.Background())
}

func (i *TriggerCustomEvent) ToTriggerCustomEventPtrOutputWithContext(ctx context.Context) TriggerCustomEventPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomEventPtrOutput)
}

type TriggerCustomEventPtrInput interface {
	pulumi.Input

	ToTriggerCustomEventPtrOutput() TriggerCustomEventPtrOutput
	ToTriggerCustomEventPtrOutputWithContext(ctx context.Context) TriggerCustomEventPtrOutput
}

type triggerCustomEventPtrType TriggerCustomEventArgs

func (*triggerCustomEventPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCustomEvent)(nil))
}

func (i *triggerCustomEventPtrType) ToTriggerCustomEventPtrOutput() TriggerCustomEventPtrOutput {
	return i.ToTriggerCustomEventPtrOutputWithContext(context.Background())
}

func (i *triggerCustomEventPtrType) ToTriggerCustomEventPtrOutputWithContext(ctx context.Context) TriggerCustomEventPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomEventPtrOutput)
}

// TriggerCustomEventArrayInput is an input type that accepts TriggerCustomEventArray and TriggerCustomEventArrayOutput values.
// You can construct a concrete instance of `TriggerCustomEventArrayInput` via:
//
//          TriggerCustomEventArray{ TriggerCustomEventArgs{...} }
type TriggerCustomEventArrayInput interface {
	pulumi.Input

	ToTriggerCustomEventArrayOutput() TriggerCustomEventArrayOutput
	ToTriggerCustomEventArrayOutputWithContext(context.Context) TriggerCustomEventArrayOutput
}

type TriggerCustomEventArray []TriggerCustomEventInput

func (TriggerCustomEventArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerCustomEvent)(nil)).Elem()
}

func (i TriggerCustomEventArray) ToTriggerCustomEventArrayOutput() TriggerCustomEventArrayOutput {
	return i.ToTriggerCustomEventArrayOutputWithContext(context.Background())
}

func (i TriggerCustomEventArray) ToTriggerCustomEventArrayOutputWithContext(ctx context.Context) TriggerCustomEventArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomEventArrayOutput)
}

// TriggerCustomEventMapInput is an input type that accepts TriggerCustomEventMap and TriggerCustomEventMapOutput values.
// You can construct a concrete instance of `TriggerCustomEventMapInput` via:
//
//          TriggerCustomEventMap{ "key": TriggerCustomEventArgs{...} }
type TriggerCustomEventMapInput interface {
	pulumi.Input

	ToTriggerCustomEventMapOutput() TriggerCustomEventMapOutput
	ToTriggerCustomEventMapOutputWithContext(context.Context) TriggerCustomEventMapOutput
}

type TriggerCustomEventMap map[string]TriggerCustomEventInput

func (TriggerCustomEventMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerCustomEvent)(nil)).Elem()
}

func (i TriggerCustomEventMap) ToTriggerCustomEventMapOutput() TriggerCustomEventMapOutput {
	return i.ToTriggerCustomEventMapOutputWithContext(context.Background())
}

func (i TriggerCustomEventMap) ToTriggerCustomEventMapOutputWithContext(ctx context.Context) TriggerCustomEventMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerCustomEventMapOutput)
}

type TriggerCustomEventOutput struct{ *pulumi.OutputState }

func (TriggerCustomEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCustomEvent)(nil))
}

func (o TriggerCustomEventOutput) ToTriggerCustomEventOutput() TriggerCustomEventOutput {
	return o
}

func (o TriggerCustomEventOutput) ToTriggerCustomEventOutputWithContext(ctx context.Context) TriggerCustomEventOutput {
	return o
}

func (o TriggerCustomEventOutput) ToTriggerCustomEventPtrOutput() TriggerCustomEventPtrOutput {
	return o.ToTriggerCustomEventPtrOutputWithContext(context.Background())
}

func (o TriggerCustomEventOutput) ToTriggerCustomEventPtrOutputWithContext(ctx context.Context) TriggerCustomEventPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerCustomEvent) *TriggerCustomEvent {
		return &v
	}).(TriggerCustomEventPtrOutput)
}

type TriggerCustomEventPtrOutput struct{ *pulumi.OutputState }

func (TriggerCustomEventPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerCustomEvent)(nil))
}

func (o TriggerCustomEventPtrOutput) ToTriggerCustomEventPtrOutput() TriggerCustomEventPtrOutput {
	return o
}

func (o TriggerCustomEventPtrOutput) ToTriggerCustomEventPtrOutputWithContext(ctx context.Context) TriggerCustomEventPtrOutput {
	return o
}

func (o TriggerCustomEventPtrOutput) Elem() TriggerCustomEventOutput {
	return o.ApplyT(func(v *TriggerCustomEvent) TriggerCustomEvent {
		if v != nil {
			return *v
		}
		var ret TriggerCustomEvent
		return ret
	}).(TriggerCustomEventOutput)
}

type TriggerCustomEventArrayOutput struct{ *pulumi.OutputState }

func (TriggerCustomEventArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerCustomEvent)(nil))
}

func (o TriggerCustomEventArrayOutput) ToTriggerCustomEventArrayOutput() TriggerCustomEventArrayOutput {
	return o
}

func (o TriggerCustomEventArrayOutput) ToTriggerCustomEventArrayOutputWithContext(ctx context.Context) TriggerCustomEventArrayOutput {
	return o
}

func (o TriggerCustomEventArrayOutput) Index(i pulumi.IntInput) TriggerCustomEventOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerCustomEvent {
		return vs[0].([]TriggerCustomEvent)[vs[1].(int)]
	}).(TriggerCustomEventOutput)
}

type TriggerCustomEventMapOutput struct{ *pulumi.OutputState }

func (TriggerCustomEventMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TriggerCustomEvent)(nil))
}

func (o TriggerCustomEventMapOutput) ToTriggerCustomEventMapOutput() TriggerCustomEventMapOutput {
	return o
}

func (o TriggerCustomEventMapOutput) ToTriggerCustomEventMapOutputWithContext(ctx context.Context) TriggerCustomEventMapOutput {
	return o
}

func (o TriggerCustomEventMapOutput) MapIndex(k pulumi.StringInput) TriggerCustomEventOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TriggerCustomEvent {
		return vs[0].(map[string]TriggerCustomEvent)[vs[1].(string)]
	}).(TriggerCustomEventOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCustomEventInput)(nil)).Elem(), &TriggerCustomEvent{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCustomEventPtrInput)(nil)).Elem(), &TriggerCustomEvent{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCustomEventArrayInput)(nil)).Elem(), TriggerCustomEventArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerCustomEventMapInput)(nil)).Elem(), TriggerCustomEventMap{})
	pulumi.RegisterOutputType(TriggerCustomEventOutput{})
	pulumi.RegisterOutputType(TriggerCustomEventPtrOutput{})
	pulumi.RegisterOutputType(TriggerCustomEventArrayOutput{})
	pulumi.RegisterOutputType(TriggerCustomEventMapOutput{})
}
