// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package consumption

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Resource Group Consumption Budget.
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
// 		_, err = consumption.NewBudgetResourceGroup(ctx, "exampleBudgetResourceGroup", &consumption.BudgetResourceGroupArgs{
// 			ResourceGroupId: exampleResourceGroup.ID(),
// 			Amount:          pulumi.Float64(1000),
// 			TimeGrain:       pulumi.String("Monthly"),
// 			TimePeriod: &consumption.BudgetResourceGroupTimePeriodArgs{
// 				StartDate: pulumi.String("2022-06-01T00:00:00Z"),
// 				EndDate:   pulumi.String("2022-07-01T00:00:00Z"),
// 			},
// 			Filter: &consumption.BudgetResourceGroupFilterArgs{
// 				Dimensions: consumption.BudgetResourceGroupFilterDimensionArray{
// 					&consumption.BudgetResourceGroupFilterDimensionArgs{
// 						Name: pulumi.String("ResourceId"),
// 						Values: pulumi.StringArray{
// 							exampleActionGroup.ID(),
// 						},
// 					},
// 				},
// 				Tags: consumption.BudgetResourceGroupFilterTagArray{
// 					&consumption.BudgetResourceGroupFilterTagArgs{
// 						Name: pulumi.String("foo"),
// 						Values: pulumi.StringArray{
// 							pulumi.String("bar"),
// 							pulumi.String("baz"),
// 						},
// 					},
// 				},
// 			},
// 			Notifications: consumption.BudgetResourceGroupNotificationArray{
// 				&consumption.BudgetResourceGroupNotificationArgs{
// 					Enabled:       pulumi.Bool(true),
// 					Threshold:     pulumi.Int(90),
// 					Operator:      pulumi.String("EqualTo"),
// 					ThresholdType: pulumi.String("Forecasted"),
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
// 				&consumption.BudgetResourceGroupNotificationArgs{
// 					Enabled:   pulumi.Bool(false),
// 					Threshold: pulumi.Int(100),
// 					Operator:  pulumi.String("GreaterThan"),
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
// Resource Group Consumption Budgets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:consumption/budgetResourceGroup:BudgetResourceGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Consumption/budgets/resourceGroup1
// ```
type BudgetResourceGroup struct {
	pulumi.CustomResourceState

	// The total amount of cost to track with the budget.
	Amount pulumi.Float64Output `pulumi:"amount"`
	// The ETag of the Resource Group Consumption Budget
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter BudgetResourceGroupFilterPtrOutput `pulumi:"filter"`
	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications BudgetResourceGroupNotificationArrayOutput `pulumi:"notifications"`
	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrOutput `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod BudgetResourceGroupTimePeriodOutput `pulumi:"timePeriod"`
}

// NewBudgetResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewBudgetResourceGroup(ctx *pulumi.Context,
	name string, args *BudgetResourceGroupArgs, opts ...pulumi.ResourceOption) (*BudgetResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.Notifications == nil {
		return nil, errors.New("invalid value for required argument 'Notifications'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	var resource BudgetResourceGroup
	err := ctx.RegisterResource("azure:consumption/budgetResourceGroup:BudgetResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudgetResourceGroup gets an existing BudgetResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudgetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetResourceGroupState, opts ...pulumi.ResourceOption) (*BudgetResourceGroup, error) {
	var resource BudgetResourceGroup
	err := ctx.ReadResource("azure:consumption/budgetResourceGroup:BudgetResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BudgetResourceGroup resources.
type budgetResourceGroupState struct {
	// The total amount of cost to track with the budget.
	Amount *float64 `pulumi:"amount"`
	// The ETag of the Resource Group Consumption Budget
	Etag *string `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter *BudgetResourceGroupFilter `pulumi:"filter"`
	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name *string `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications []BudgetResourceGroupNotification `pulumi:"notifications"`
	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain *string `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod *BudgetResourceGroupTimePeriod `pulumi:"timePeriod"`
}

type BudgetResourceGroupState struct {
	// The total amount of cost to track with the budget.
	Amount pulumi.Float64PtrInput
	// The ETag of the Resource Group Consumption Budget
	Etag pulumi.StringPtrInput
	// A `filter` block as defined below.
	Filter BudgetResourceGroupFilterPtrInput
	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name pulumi.StringPtrInput
	// One or more `notification` blocks as defined below.
	Notifications BudgetResourceGroupNotificationArrayInput
	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupId pulumi.StringPtrInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrInput
	// A `timePeriod` block as defined below.
	TimePeriod BudgetResourceGroupTimePeriodPtrInput
}

func (BudgetResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetResourceGroupState)(nil)).Elem()
}

type budgetResourceGroupArgs struct {
	// The total amount of cost to track with the budget.
	Amount float64 `pulumi:"amount"`
	// The ETag of the Resource Group Consumption Budget
	Etag *string `pulumi:"etag"`
	// A `filter` block as defined below.
	Filter *BudgetResourceGroupFilter `pulumi:"filter"`
	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name *string `pulumi:"name"`
	// One or more `notification` blocks as defined below.
	Notifications []BudgetResourceGroupNotification `pulumi:"notifications"`
	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain *string `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriod BudgetResourceGroupTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a BudgetResourceGroup resource.
type BudgetResourceGroupArgs struct {
	// The total amount of cost to track with the budget.
	Amount pulumi.Float64Input
	// The ETag of the Resource Group Consumption Budget
	Etag pulumi.StringPtrInput
	// A `filter` block as defined below.
	Filter BudgetResourceGroupFilterPtrInput
	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name pulumi.StringPtrInput
	// One or more `notification` blocks as defined below.
	Notifications BudgetResourceGroupNotificationArrayInput
	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupId pulumi.StringInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of `Monthly`, `Quarterly`, `Annually`, `BillingMonth`, `BillingQuarter`, or `BillingYear`. Defaults to `Monthly`.
	TimeGrain pulumi.StringPtrInput
	// A `timePeriod` block as defined below.
	TimePeriod BudgetResourceGroupTimePeriodInput
}

func (BudgetResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetResourceGroupArgs)(nil)).Elem()
}

type BudgetResourceGroupInput interface {
	pulumi.Input

	ToBudgetResourceGroupOutput() BudgetResourceGroupOutput
	ToBudgetResourceGroupOutputWithContext(ctx context.Context) BudgetResourceGroupOutput
}

func (*BudgetResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetResourceGroup)(nil)).Elem()
}

func (i *BudgetResourceGroup) ToBudgetResourceGroupOutput() BudgetResourceGroupOutput {
	return i.ToBudgetResourceGroupOutputWithContext(context.Background())
}

func (i *BudgetResourceGroup) ToBudgetResourceGroupOutputWithContext(ctx context.Context) BudgetResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceGroupOutput)
}

// BudgetResourceGroupArrayInput is an input type that accepts BudgetResourceGroupArray and BudgetResourceGroupArrayOutput values.
// You can construct a concrete instance of `BudgetResourceGroupArrayInput` via:
//
//          BudgetResourceGroupArray{ BudgetResourceGroupArgs{...} }
type BudgetResourceGroupArrayInput interface {
	pulumi.Input

	ToBudgetResourceGroupArrayOutput() BudgetResourceGroupArrayOutput
	ToBudgetResourceGroupArrayOutputWithContext(context.Context) BudgetResourceGroupArrayOutput
}

type BudgetResourceGroupArray []BudgetResourceGroupInput

func (BudgetResourceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BudgetResourceGroup)(nil)).Elem()
}

func (i BudgetResourceGroupArray) ToBudgetResourceGroupArrayOutput() BudgetResourceGroupArrayOutput {
	return i.ToBudgetResourceGroupArrayOutputWithContext(context.Background())
}

func (i BudgetResourceGroupArray) ToBudgetResourceGroupArrayOutputWithContext(ctx context.Context) BudgetResourceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceGroupArrayOutput)
}

// BudgetResourceGroupMapInput is an input type that accepts BudgetResourceGroupMap and BudgetResourceGroupMapOutput values.
// You can construct a concrete instance of `BudgetResourceGroupMapInput` via:
//
//          BudgetResourceGroupMap{ "key": BudgetResourceGroupArgs{...} }
type BudgetResourceGroupMapInput interface {
	pulumi.Input

	ToBudgetResourceGroupMapOutput() BudgetResourceGroupMapOutput
	ToBudgetResourceGroupMapOutputWithContext(context.Context) BudgetResourceGroupMapOutput
}

type BudgetResourceGroupMap map[string]BudgetResourceGroupInput

func (BudgetResourceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BudgetResourceGroup)(nil)).Elem()
}

func (i BudgetResourceGroupMap) ToBudgetResourceGroupMapOutput() BudgetResourceGroupMapOutput {
	return i.ToBudgetResourceGroupMapOutputWithContext(context.Background())
}

func (i BudgetResourceGroupMap) ToBudgetResourceGroupMapOutputWithContext(ctx context.Context) BudgetResourceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceGroupMapOutput)
}

type BudgetResourceGroupOutput struct{ *pulumi.OutputState }

func (BudgetResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetResourceGroup)(nil)).Elem()
}

func (o BudgetResourceGroupOutput) ToBudgetResourceGroupOutput() BudgetResourceGroupOutput {
	return o
}

func (o BudgetResourceGroupOutput) ToBudgetResourceGroupOutputWithContext(ctx context.Context) BudgetResourceGroupOutput {
	return o
}

type BudgetResourceGroupArrayOutput struct{ *pulumi.OutputState }

func (BudgetResourceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BudgetResourceGroup)(nil)).Elem()
}

func (o BudgetResourceGroupArrayOutput) ToBudgetResourceGroupArrayOutput() BudgetResourceGroupArrayOutput {
	return o
}

func (o BudgetResourceGroupArrayOutput) ToBudgetResourceGroupArrayOutputWithContext(ctx context.Context) BudgetResourceGroupArrayOutput {
	return o
}

func (o BudgetResourceGroupArrayOutput) Index(i pulumi.IntInput) BudgetResourceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BudgetResourceGroup {
		return vs[0].([]*BudgetResourceGroup)[vs[1].(int)]
	}).(BudgetResourceGroupOutput)
}

type BudgetResourceGroupMapOutput struct{ *pulumi.OutputState }

func (BudgetResourceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BudgetResourceGroup)(nil)).Elem()
}

func (o BudgetResourceGroupMapOutput) ToBudgetResourceGroupMapOutput() BudgetResourceGroupMapOutput {
	return o
}

func (o BudgetResourceGroupMapOutput) ToBudgetResourceGroupMapOutputWithContext(ctx context.Context) BudgetResourceGroupMapOutput {
	return o
}

func (o BudgetResourceGroupMapOutput) MapIndex(k pulumi.StringInput) BudgetResourceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BudgetResourceGroup {
		return vs[0].(map[string]*BudgetResourceGroup)[vs[1].(string)]
	}).(BudgetResourceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceGroupInput)(nil)).Elem(), &BudgetResourceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceGroupArrayInput)(nil)).Elem(), BudgetResourceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceGroupMapInput)(nil)).Elem(), BudgetResourceGroupMap{})
	pulumi.RegisterOutputType(BudgetResourceGroupOutput{})
	pulumi.RegisterOutputType(BudgetResourceGroupArrayOutput{})
	pulumi.RegisterOutputType(BudgetResourceGroupMapOutput{})
}
