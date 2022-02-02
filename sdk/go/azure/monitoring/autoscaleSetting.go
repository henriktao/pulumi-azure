// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a AutoScale Setting which can be applied to Virtual Machine Scale Sets, App Services and other scalable resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
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
// 		exampleScaleSet, err := compute.NewScaleSet(ctx, "exampleScaleSet", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewAutoscaleSetting(ctx, "exampleAutoscaleSetting", &monitoring.AutoscaleSettingArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			TargetResourceId:  exampleScaleSet.ID(),
// 			Profiles: monitoring.AutoscaleSettingProfileArray{
// 				&monitoring.AutoscaleSettingProfileArgs{
// 					Name: pulumi.String("defaultProfile"),
// 					Capacity: &monitoring.AutoscaleSettingProfileCapacityArgs{
// 						Default: pulumi.Int(1),
// 						Minimum: pulumi.Int(1),
// 						Maximum: pulumi.Int(10),
// 					},
// 					Rules: monitoring.AutoscaleSettingProfileRuleArray{
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("GreaterThan"),
// 								Threshold:        pulumi.Float64(75),
// 								MetricNamespace:  pulumi.String("microsoft.compute/virtualmachinescalesets"),
// 								Dimensions: monitoring.AutoscaleSettingProfileRuleMetricTriggerDimensionArray{
// 									&monitoring.AutoscaleSettingProfileRuleMetricTriggerDimensionArgs{
// 										Name:     pulumi.String("AppName"),
// 										Operator: pulumi.String("Equals"),
// 										Values: pulumi.StringArray{
// 											pulumi.String("App1"),
// 										},
// 									},
// 								},
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Increase"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(1),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("LessThan"),
// 								Threshold:        pulumi.Float64(25),
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Decrease"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(1),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Notification: &monitoring.AutoscaleSettingNotificationArgs{
// 				Email: &monitoring.AutoscaleSettingNotificationEmailArgs{
// 					SendToSubscriptionAdministrator:   pulumi.Bool(true),
// 					SendToSubscriptionCoAdministrator: pulumi.Bool(true),
// 					CustomEmails: pulumi.StringArray{
// 						pulumi.String("admin@contoso.com"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Repeating On Weekends)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
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
// 		exampleScaleSet, err := compute.NewScaleSet(ctx, "exampleScaleSet", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewAutoscaleSetting(ctx, "exampleAutoscaleSetting", &monitoring.AutoscaleSettingArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			TargetResourceId:  exampleScaleSet.ID(),
// 			Profiles: monitoring.AutoscaleSettingProfileArray{
// 				&monitoring.AutoscaleSettingProfileArgs{
// 					Name: pulumi.String("Weekends"),
// 					Capacity: &monitoring.AutoscaleSettingProfileCapacityArgs{
// 						Default: pulumi.Int(1),
// 						Minimum: pulumi.Int(1),
// 						Maximum: pulumi.Int(10),
// 					},
// 					Rules: monitoring.AutoscaleSettingProfileRuleArray{
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("GreaterThan"),
// 								Threshold:        pulumi.Float64(90),
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Increase"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(2),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("LessThan"),
// 								Threshold:        pulumi.Float64(10),
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Decrease"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(2),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 					},
// 					Recurrence: &monitoring.AutoscaleSettingProfileRecurrenceArgs{
// 						Frequency: "Week",
// 						Timezone:  pulumi.String("Pacific Standard Time"),
// 						Days: pulumi.StringArray{
// 							pulumi.String("Saturday"),
// 							pulumi.String("Sunday"),
// 						},
// 						Hours: pulumi.Int{
// 							12,
// 						},
// 						Minutes: pulumi.Int{
// 							0,
// 						},
// 					},
// 				},
// 			},
// 			Notification: &monitoring.AutoscaleSettingNotificationArgs{
// 				Email: &monitoring.AutoscaleSettingNotificationEmailArgs{
// 					SendToSubscriptionAdministrator:   pulumi.Bool(true),
// 					SendToSubscriptionCoAdministrator: pulumi.Bool(true),
// 					CustomEmails: pulumi.StringArray{
// 						pulumi.String("admin@contoso.com"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### For Fixed Dates)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
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
// 		exampleScaleSet, err := compute.NewScaleSet(ctx, "exampleScaleSet", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewAutoscaleSetting(ctx, "exampleAutoscaleSetting", &monitoring.AutoscaleSettingArgs{
// 			Enabled:           pulumi.Bool(true),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			TargetResourceId:  exampleScaleSet.ID(),
// 			Profiles: monitoring.AutoscaleSettingProfileArray{
// 				&monitoring.AutoscaleSettingProfileArgs{
// 					Name: pulumi.String("forJuly"),
// 					Capacity: &monitoring.AutoscaleSettingProfileCapacityArgs{
// 						Default: pulumi.Int(1),
// 						Minimum: pulumi.Int(1),
// 						Maximum: pulumi.Int(10),
// 					},
// 					Rules: monitoring.AutoscaleSettingProfileRuleArray{
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("GreaterThan"),
// 								Threshold:        pulumi.Float64(90),
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Increase"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(2),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 						&monitoring.AutoscaleSettingProfileRuleArgs{
// 							MetricTrigger: &monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs{
// 								MetricName:       pulumi.String("Percentage CPU"),
// 								MetricResourceId: exampleScaleSet.ID(),
// 								TimeGrain:        pulumi.String("PT1M"),
// 								Statistic:        pulumi.String("Average"),
// 								TimeWindow:       pulumi.String("PT5M"),
// 								TimeAggregation:  pulumi.String("Average"),
// 								Operator:         pulumi.String("LessThan"),
// 								Threshold:        pulumi.Float64(10),
// 							},
// 							ScaleAction: &monitoring.AutoscaleSettingProfileRuleScaleActionArgs{
// 								Direction: pulumi.String("Decrease"),
// 								Type:      pulumi.String("ChangeCount"),
// 								Value:     pulumi.Int(2),
// 								Cooldown:  pulumi.String("PT1M"),
// 							},
// 						},
// 					},
// 					FixedDate: &monitoring.AutoscaleSettingProfileFixedDateArgs{
// 						Timezone: pulumi.String("Pacific Standard Time"),
// 						Start:    pulumi.String("2020-07-01T00:00:00Z"),
// 						End:      pulumi.String("2020-07-31T23:59:59Z"),
// 					},
// 				},
// 			},
// 			Notification: &monitoring.AutoscaleSettingNotificationArgs{
// 				Email: &monitoring.AutoscaleSettingNotificationEmailArgs{
// 					SendToSubscriptionAdministrator:   pulumi.Bool(true),
// 					SendToSubscriptionCoAdministrator: pulumi.Bool(true),
// 					CustomEmails: pulumi.StringArray{
// 						pulumi.String("admin@contoso.com"),
// 					},
// 				},
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
// AutoScale Setting can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/autoscaleSetting:AutoscaleSetting example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/autoscaleSettings/setting1
// ```
type AutoscaleSetting struct {
	pulumi.CustomResourceState

	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrOutput `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayOutput `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewAutoscaleSetting registers a new resource with the given unique name, arguments, and options.
func NewAutoscaleSetting(ctx *pulumi.Context,
	name string, args *AutoscaleSettingArgs, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Profiles == nil {
		return nil, errors.New("invalid value for required argument 'Profiles'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource AutoscaleSetting
	err := ctx.RegisterResource("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscaleSetting gets an existing AutoscaleSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscaleSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscaleSettingState, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	var resource AutoscaleSetting
	err := ctx.ReadResource("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscaleSetting resources.
type autoscaleSettingState struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification *AutoscaleSettingNotification `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles []AutoscaleSettingProfile `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type AutoscaleSettingState struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrInput
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayInput
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringPtrInput
}

func (AutoscaleSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingState)(nil)).Elem()
}

type autoscaleSettingArgs struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification *AutoscaleSettingNotification `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles []AutoscaleSettingProfile `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a AutoscaleSetting resource.
type AutoscaleSettingArgs struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrInput
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayInput
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringInput
}

func (AutoscaleSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingArgs)(nil)).Elem()
}

type AutoscaleSettingInput interface {
	pulumi.Input

	ToAutoscaleSettingOutput() AutoscaleSettingOutput
	ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput
}

func (*AutoscaleSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSetting)(nil)).Elem()
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return i.ToAutoscaleSettingOutputWithContext(context.Background())
}

func (i *AutoscaleSetting) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingOutput)
}

// AutoscaleSettingArrayInput is an input type that accepts AutoscaleSettingArray and AutoscaleSettingArrayOutput values.
// You can construct a concrete instance of `AutoscaleSettingArrayInput` via:
//
//          AutoscaleSettingArray{ AutoscaleSettingArgs{...} }
type AutoscaleSettingArrayInput interface {
	pulumi.Input

	ToAutoscaleSettingArrayOutput() AutoscaleSettingArrayOutput
	ToAutoscaleSettingArrayOutputWithContext(context.Context) AutoscaleSettingArrayOutput
}

type AutoscaleSettingArray []AutoscaleSettingInput

func (AutoscaleSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscaleSetting)(nil)).Elem()
}

func (i AutoscaleSettingArray) ToAutoscaleSettingArrayOutput() AutoscaleSettingArrayOutput {
	return i.ToAutoscaleSettingArrayOutputWithContext(context.Background())
}

func (i AutoscaleSettingArray) ToAutoscaleSettingArrayOutputWithContext(ctx context.Context) AutoscaleSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingArrayOutput)
}

// AutoscaleSettingMapInput is an input type that accepts AutoscaleSettingMap and AutoscaleSettingMapOutput values.
// You can construct a concrete instance of `AutoscaleSettingMapInput` via:
//
//          AutoscaleSettingMap{ "key": AutoscaleSettingArgs{...} }
type AutoscaleSettingMapInput interface {
	pulumi.Input

	ToAutoscaleSettingMapOutput() AutoscaleSettingMapOutput
	ToAutoscaleSettingMapOutputWithContext(context.Context) AutoscaleSettingMapOutput
}

type AutoscaleSettingMap map[string]AutoscaleSettingInput

func (AutoscaleSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscaleSetting)(nil)).Elem()
}

func (i AutoscaleSettingMap) ToAutoscaleSettingMapOutput() AutoscaleSettingMapOutput {
	return i.ToAutoscaleSettingMapOutputWithContext(context.Background())
}

func (i AutoscaleSettingMap) ToAutoscaleSettingMapOutputWithContext(ctx context.Context) AutoscaleSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingMapOutput)
}

type AutoscaleSettingOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSetting)(nil)).Elem()
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutput() AutoscaleSettingOutput {
	return o
}

func (o AutoscaleSettingOutput) ToAutoscaleSettingOutputWithContext(ctx context.Context) AutoscaleSettingOutput {
	return o
}

type AutoscaleSettingArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoscaleSetting)(nil)).Elem()
}

func (o AutoscaleSettingArrayOutput) ToAutoscaleSettingArrayOutput() AutoscaleSettingArrayOutput {
	return o
}

func (o AutoscaleSettingArrayOutput) ToAutoscaleSettingArrayOutputWithContext(ctx context.Context) AutoscaleSettingArrayOutput {
	return o
}

func (o AutoscaleSettingArrayOutput) Index(i pulumi.IntInput) AutoscaleSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoscaleSetting {
		return vs[0].([]*AutoscaleSetting)[vs[1].(int)]
	}).(AutoscaleSettingOutput)
}

type AutoscaleSettingMapOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoscaleSetting)(nil)).Elem()
}

func (o AutoscaleSettingMapOutput) ToAutoscaleSettingMapOutput() AutoscaleSettingMapOutput {
	return o
}

func (o AutoscaleSettingMapOutput) ToAutoscaleSettingMapOutputWithContext(ctx context.Context) AutoscaleSettingMapOutput {
	return o
}

func (o AutoscaleSettingMapOutput) MapIndex(k pulumi.StringInput) AutoscaleSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoscaleSetting {
		return vs[0].(map[string]*AutoscaleSetting)[vs[1].(string)]
	}).(AutoscaleSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscaleSettingInput)(nil)).Elem(), &AutoscaleSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscaleSettingArrayInput)(nil)).Elem(), AutoscaleSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscaleSettingMapInput)(nil)).Elem(), AutoscaleSettingMap{})
	pulumi.RegisterOutputType(AutoscaleSettingOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingMapOutput{})
}
