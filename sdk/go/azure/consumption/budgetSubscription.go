// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package consumption

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Subscription Consumption Budget.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/consumption"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleActionGroup, err := monitoring.NewActionGroup(ctx, "exampleActionGroup", &monitoring.ActionGroupArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ShortName:         pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = consumption.NewBudgetSubscription(ctx, "exampleBudgetSubscription", &consumption.BudgetSubscriptionArgs{
// 			SubscriptionId: pulumi.String(current.Id),
// 			Amount:         pulumi.Float64(1000),
// 			TimeGrain:      pulumi.String("Monthly"),
// 			TimePeriod: &consumption.BudgetSubscriptionTimePeriodArgs{
// 				StartDate: pulumi.String("2022-06-01T00:00:00Z"),
// 				EndDate:   pulumi.String("2022-07-01T00:00:00Z"),
// 			},
// 			Filter: &consumption.BudgetSubscriptionFilterArgs{
// 				Dimensions: consumption.BudgetSubscriptionFilterDimensionArray{
// 					&consumption.BudgetSubscriptionFilterDimensionArgs{
// 						Name: pulumi.String("ResourceGroupName"),
// 						Values: pulumi.StringArray{
// 							exampleResourceGroup.Name,
// 						},
// 					},
// 				},
// 				Tags: consumption.BudgetSubscriptionFilterTagArray{
// 					&consumption.BudgetSubscriptionFilterTagArgs{
// 						Name: pulumi.String("foo"),
// 						Values: pulumi.StringArray{
// 							pulumi.String("bar"),
// 							pulumi.String("baz"),
// 						},
// 					},
// 				},
// 			},
// 			Notifications: consumption.BudgetSubscriptionNotificationArray{
// 				&consumption.BudgetSubscriptionNotificationArgs{
// 					Enabled:   pulumi.Bool(true),
// 					Threshold: pulumi.Int(90),
// 					Operator:  pulumi.String("EqualTo"),
// 					ContactEmails: pulumi.StringArray{
// 						pulumi.String("foo@example.com"),
// 						pulumi.String("bar@example.com"),
// 					},
// 					ContactGroups: pulumi.StringArray{
// 						exampleActionGroup.ID(),
// 					},
// 					ContactRoles: pulumi.StringArray{
// 						pulumi.String("Owner"),
// 					},
// 				},
// 				&consumption.BudgetSubscriptionNotificationArgs{
// 					Enabled:       pulumi.Bool(false),
// 					Threshold:     pulumi.Int(100),
// 					Operator:      pulumi.String("GreaterThan"),
// 					ThresholdType: pulumi.String("Forecasted"),
// 					ContactEmails: pulumi.StringArray{
// 						pulumi.String("foo@example.com"),
// 						pulumi.String("bar@example.com"),
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
// Subscription Consumption Budgets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:consumption/budgetSubscription:BudgetSubscription example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/subscription1
// ```
type BudgetSubscription struct {
	pulumi.CustomResourceState

	// The total amount of cost to track with the budget.
	Amount pulumi.Float64Output `pulumi:"amount"`
	// The ETag of the Subscription Consumption Budget.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter BudgetSubscriptionFilterPtrOutput `pulumi:"filter"`
	// The name which should be used for this Subscription Consumption Budget. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications BudgetSubscriptionNotificationArrayOutput `pulumi:"notifications"`
	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrOutput `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod BudgetSubscriptionTimePeriodOutput `pulumi:"timePeriod"`
}

// NewBudgetSubscription registers a new resource with the given unique name, arguments, and options.
func NewBudgetSubscription(ctx *pulumi.Context,
	name string, args *BudgetSubscriptionArgs, opts ...pulumi.ResourceOption) (*BudgetSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.Notifications == nil {
		return nil, errors.New("invalid value for required argument 'Notifications'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	var resource BudgetSubscription
	err := ctx.RegisterResource("azure:consumption/budgetSubscription:BudgetSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudgetSubscription gets an existing BudgetSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudgetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetSubscriptionState, opts ...pulumi.ResourceOption) (*BudgetSubscription, error) {
	var resource BudgetSubscription
	err := ctx.ReadResource("azure:consumption/budgetSubscription:BudgetSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BudgetSubscription resources.
type budgetSubscriptionState struct {
	// The total amount of cost to track with the budget.
	Amount *float64 `pulumi:"amount"`
	// The ETag of the Subscription Consumption Budget.
	Etag *string `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter *BudgetSubscriptionFilter `pulumi:"filter"`
	// The name which should be used for this Subscription Consumption Budget. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications []BudgetSubscriptionNotification `pulumi:"notifications"`
	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain *string `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod *BudgetSubscriptionTimePeriod `pulumi:"timePeriod"`
}

type BudgetSubscriptionState struct {
	// The total amount of cost to track with the budget.
	Amount pulumi.Float64PtrInput
	// The ETag of the Subscription Consumption Budget.
	Etag pulumi.StringPtrInput
	// A `filter` block as defined below.
	Filter BudgetSubscriptionFilterPtrInput
	// The name which should be used for this Subscription Consumption Budget. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `notification` blocks as defined below.
	Notifications BudgetSubscriptionNotificationArrayInput
	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrInput
	// A `timePeriod` block as defined below.
	TimePeriod BudgetSubscriptionTimePeriodPtrInput
}

func (BudgetSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetSubscriptionState)(nil)).Elem()
}

type budgetSubscriptionArgs struct {
	// The total amount of cost to track with the budget.
	Amount float64 `pulumi:"amount"`
	// The ETag of the Subscription Consumption Budget.
	Etag *string `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter *BudgetSubscriptionFilter `pulumi:"filter"`
	// The name which should be used for this Subscription Consumption Budget. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications []BudgetSubscriptionNotification `pulumi:"notifications"`
	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain *string `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod BudgetSubscriptionTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a BudgetSubscription resource.
type BudgetSubscriptionArgs struct {
	// The total amount of cost to track with the budget.
	Amount pulumi.Float64Input
	// The ETag of the Subscription Consumption Budget.
	Etag pulumi.StringPtrInput
	// A `filter` block as defined below.
	Filter BudgetSubscriptionFilterPtrInput
	// The name which should be used for this Subscription Consumption Budget. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `notification` blocks as defined below.
	Notifications BudgetSubscriptionNotificationArrayInput
	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrInput
	// A `timePeriod` block as defined below.
	TimePeriod BudgetSubscriptionTimePeriodInput
}

func (BudgetSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetSubscriptionArgs)(nil)).Elem()
}

type BudgetSubscriptionInput interface {
	pulumi.Input

	ToBudgetSubscriptionOutput() BudgetSubscriptionOutput
	ToBudgetSubscriptionOutputWithContext(ctx context.Context) BudgetSubscriptionOutput
}

func (*BudgetSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetSubscription)(nil))
}

func (i *BudgetSubscription) ToBudgetSubscriptionOutput() BudgetSubscriptionOutput {
	return i.ToBudgetSubscriptionOutputWithContext(context.Background())
}

func (i *BudgetSubscription) ToBudgetSubscriptionOutputWithContext(ctx context.Context) BudgetSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetSubscriptionOutput)
}

func (i *BudgetSubscription) ToBudgetSubscriptionPtrOutput() BudgetSubscriptionPtrOutput {
	return i.ToBudgetSubscriptionPtrOutputWithContext(context.Background())
}

func (i *BudgetSubscription) ToBudgetSubscriptionPtrOutputWithContext(ctx context.Context) BudgetSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetSubscriptionPtrOutput)
}

type BudgetSubscriptionPtrInput interface {
	pulumi.Input

	ToBudgetSubscriptionPtrOutput() BudgetSubscriptionPtrOutput
	ToBudgetSubscriptionPtrOutputWithContext(ctx context.Context) BudgetSubscriptionPtrOutput
}

type budgetSubscriptionPtrType BudgetSubscriptionArgs

func (*budgetSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetSubscription)(nil))
}

func (i *budgetSubscriptionPtrType) ToBudgetSubscriptionPtrOutput() BudgetSubscriptionPtrOutput {
	return i.ToBudgetSubscriptionPtrOutputWithContext(context.Background())
}

func (i *budgetSubscriptionPtrType) ToBudgetSubscriptionPtrOutputWithContext(ctx context.Context) BudgetSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetSubscriptionPtrOutput)
}

// BudgetSubscriptionArrayInput is an input type that accepts BudgetSubscriptionArray and BudgetSubscriptionArrayOutput values.
// You can construct a concrete instance of `BudgetSubscriptionArrayInput` via:
//
//          BudgetSubscriptionArray{ BudgetSubscriptionArgs{...} }
type BudgetSubscriptionArrayInput interface {
	pulumi.Input

	ToBudgetSubscriptionArrayOutput() BudgetSubscriptionArrayOutput
	ToBudgetSubscriptionArrayOutputWithContext(context.Context) BudgetSubscriptionArrayOutput
}

type BudgetSubscriptionArray []BudgetSubscriptionInput

func (BudgetSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BudgetSubscription)(nil)).Elem()
}

func (i BudgetSubscriptionArray) ToBudgetSubscriptionArrayOutput() BudgetSubscriptionArrayOutput {
	return i.ToBudgetSubscriptionArrayOutputWithContext(context.Background())
}

func (i BudgetSubscriptionArray) ToBudgetSubscriptionArrayOutputWithContext(ctx context.Context) BudgetSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetSubscriptionArrayOutput)
}

// BudgetSubscriptionMapInput is an input type that accepts BudgetSubscriptionMap and BudgetSubscriptionMapOutput values.
// You can construct a concrete instance of `BudgetSubscriptionMapInput` via:
//
//          BudgetSubscriptionMap{ "key": BudgetSubscriptionArgs{...} }
type BudgetSubscriptionMapInput interface {
	pulumi.Input

	ToBudgetSubscriptionMapOutput() BudgetSubscriptionMapOutput
	ToBudgetSubscriptionMapOutputWithContext(context.Context) BudgetSubscriptionMapOutput
}

type BudgetSubscriptionMap map[string]BudgetSubscriptionInput

func (BudgetSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BudgetSubscription)(nil)).Elem()
}

func (i BudgetSubscriptionMap) ToBudgetSubscriptionMapOutput() BudgetSubscriptionMapOutput {
	return i.ToBudgetSubscriptionMapOutputWithContext(context.Background())
}

func (i BudgetSubscriptionMap) ToBudgetSubscriptionMapOutputWithContext(ctx context.Context) BudgetSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetSubscriptionMapOutput)
}

type BudgetSubscriptionOutput struct{ *pulumi.OutputState }

func (BudgetSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetSubscription)(nil))
}

func (o BudgetSubscriptionOutput) ToBudgetSubscriptionOutput() BudgetSubscriptionOutput {
	return o
}

func (o BudgetSubscriptionOutput) ToBudgetSubscriptionOutputWithContext(ctx context.Context) BudgetSubscriptionOutput {
	return o
}

func (o BudgetSubscriptionOutput) ToBudgetSubscriptionPtrOutput() BudgetSubscriptionPtrOutput {
	return o.ToBudgetSubscriptionPtrOutputWithContext(context.Background())
}

func (o BudgetSubscriptionOutput) ToBudgetSubscriptionPtrOutputWithContext(ctx context.Context) BudgetSubscriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetSubscription) *BudgetSubscription {
		return &v
	}).(BudgetSubscriptionPtrOutput)
}

type BudgetSubscriptionPtrOutput struct{ *pulumi.OutputState }

func (BudgetSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetSubscription)(nil))
}

func (o BudgetSubscriptionPtrOutput) ToBudgetSubscriptionPtrOutput() BudgetSubscriptionPtrOutput {
	return o
}

func (o BudgetSubscriptionPtrOutput) ToBudgetSubscriptionPtrOutputWithContext(ctx context.Context) BudgetSubscriptionPtrOutput {
	return o
}

func (o BudgetSubscriptionPtrOutput) Elem() BudgetSubscriptionOutput {
	return o.ApplyT(func(v *BudgetSubscription) BudgetSubscription {
		if v != nil {
			return *v
		}
		var ret BudgetSubscription
		return ret
	}).(BudgetSubscriptionOutput)
}

type BudgetSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (BudgetSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BudgetSubscription)(nil))
}

func (o BudgetSubscriptionArrayOutput) ToBudgetSubscriptionArrayOutput() BudgetSubscriptionArrayOutput {
	return o
}

func (o BudgetSubscriptionArrayOutput) ToBudgetSubscriptionArrayOutputWithContext(ctx context.Context) BudgetSubscriptionArrayOutput {
	return o
}

func (o BudgetSubscriptionArrayOutput) Index(i pulumi.IntInput) BudgetSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BudgetSubscription {
		return vs[0].([]BudgetSubscription)[vs[1].(int)]
	}).(BudgetSubscriptionOutput)
}

type BudgetSubscriptionMapOutput struct{ *pulumi.OutputState }

func (BudgetSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BudgetSubscription)(nil))
}

func (o BudgetSubscriptionMapOutput) ToBudgetSubscriptionMapOutput() BudgetSubscriptionMapOutput {
	return o
}

func (o BudgetSubscriptionMapOutput) ToBudgetSubscriptionMapOutputWithContext(ctx context.Context) BudgetSubscriptionMapOutput {
	return o
}

func (o BudgetSubscriptionMapOutput) MapIndex(k pulumi.StringInput) BudgetSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BudgetSubscription {
		return vs[0].(map[string]BudgetSubscription)[vs[1].(string)]
	}).(BudgetSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetSubscriptionInput)(nil)).Elem(), &BudgetSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetSubscriptionPtrInput)(nil)).Elem(), &BudgetSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetSubscriptionArrayInput)(nil)).Elem(), BudgetSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetSubscriptionMapInput)(nil)).Elem(), BudgetSubscriptionMap{})
	pulumi.RegisterOutputType(BudgetSubscriptionOutput{})
	pulumi.RegisterOutputType(BudgetSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(BudgetSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(BudgetSubscriptionMapOutput{})
}
