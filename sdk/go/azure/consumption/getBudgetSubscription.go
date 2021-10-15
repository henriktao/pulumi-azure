// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package consumption

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Consumption Budget for a specific subscription.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/consumption"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := consumption.LookupBudgetSubscription(ctx, &consumption.LookupBudgetSubscriptionArgs{
// 			Name:           "existing",
// 			SubscriptionId: "/subscriptions/00000000-0000-0000-0000-000000000000/",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", data.Azurerm_consumption_budget.Example.Id)
// 		return nil
// 	})
// }
// ```
func LookupBudgetSubscription(ctx *pulumi.Context, args *LookupBudgetSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupBudgetSubscriptionResult, error) {
	var rv LookupBudgetSubscriptionResult
	err := ctx.Invoke("azure:consumption/getBudgetSubscription:getBudgetSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBudgetSubscription.
type LookupBudgetSubscriptionArgs struct {
	// The name of this Consumption Budget.
	Name string `pulumi:"name"`
	// The ID of the subscription.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// A collection of values returned by getBudgetSubscription.
type LookupBudgetSubscriptionResult struct {
	// The total amount of cost to track with the budget.
	Amount float64 `pulumi:"amount"`
	// A `filter` block as defined below.
	Filters []GetBudgetSubscriptionFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the tag to use for the filter.
	Name string `pulumi:"name"`
	// A `notification` block as defined below.
	Notifications  []GetBudgetSubscriptionNotification `pulumi:"notifications"`
	SubscriptionId string                              `pulumi:"subscriptionId"`
	// The time covered by a budget.
	TimeGrain string `pulumi:"timeGrain"`
	// A `timePeriod` block as defined below.
	TimePeriods []GetBudgetSubscriptionTimePeriod `pulumi:"timePeriods"`
}

func LookupBudgetSubscriptionOutput(ctx *pulumi.Context, args LookupBudgetSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetSubscriptionResult, error) {
			args := v.(LookupBudgetSubscriptionArgs)
			r, err := LookupBudgetSubscription(ctx, &args, opts...)
			return *r, err
		}).(LookupBudgetSubscriptionResultOutput)
}

// A collection of arguments for invoking getBudgetSubscription.
type LookupBudgetSubscriptionOutputArgs struct {
	// The name of this Consumption Budget.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the subscription.
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
}

func (LookupBudgetSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetSubscriptionArgs)(nil)).Elem()
}

// A collection of values returned by getBudgetSubscription.
type LookupBudgetSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetSubscriptionResult)(nil)).Elem()
}

func (o LookupBudgetSubscriptionResultOutput) ToLookupBudgetSubscriptionResultOutput() LookupBudgetSubscriptionResultOutput {
	return o
}

func (o LookupBudgetSubscriptionResultOutput) ToLookupBudgetSubscriptionResultOutputWithContext(ctx context.Context) LookupBudgetSubscriptionResultOutput {
	return o
}

// The total amount of cost to track with the budget.
func (o LookupBudgetSubscriptionResultOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) float64 { return v.Amount }).(pulumi.Float64Output)
}

// A `filter` block as defined below.
func (o LookupBudgetSubscriptionResultOutput) Filters() GetBudgetSubscriptionFilterArrayOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) []GetBudgetSubscriptionFilter { return v.Filters }).(GetBudgetSubscriptionFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBudgetSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the tag to use for the filter.
func (o LookupBudgetSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// A `notification` block as defined below.
func (o LookupBudgetSubscriptionResultOutput) Notifications() GetBudgetSubscriptionNotificationArrayOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) []GetBudgetSubscriptionNotification { return v.Notifications }).(GetBudgetSubscriptionNotificationArrayOutput)
}

func (o LookupBudgetSubscriptionResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The time covered by a budget.
func (o LookupBudgetSubscriptionResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

// A `timePeriod` block as defined below.
func (o LookupBudgetSubscriptionResultOutput) TimePeriods() GetBudgetSubscriptionTimePeriodArrayOutput {
	return o.ApplyT(func(v LookupBudgetSubscriptionResult) []GetBudgetSubscriptionTimePeriod { return v.TimePeriods }).(GetBudgetSubscriptionTimePeriodArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetSubscriptionResultOutput{})
}
