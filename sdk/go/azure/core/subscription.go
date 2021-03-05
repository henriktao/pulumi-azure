// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Subscriptions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/subscription:Subscription example "/providers/Microsoft.Subscription/aliases/subscription1"
// ```
//
//  In this scenario, the `subscription_id` property can be completed and Terraform will assume control of the existing subscription by creating an Alias. e.g.
type Subscription struct {
	pulumi.CustomResourceState

	// The Alias Name of the subscription. If omitted a new UUID will be generated for this property.
	Alias          pulumi.StringOutput    `pulumi:"alias"`
	BillingScopeId pulumi.StringPtrOutput `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringOutput    `pulumi:"subscriptionName"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Tenant to which the subscription belongs.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrOutput `pulumi:"workload"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	var resource Subscription
	err := ctx.RegisterResource("azure:core/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure:core/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// The Alias Name of the subscription. If omitted a new UUID will be generated for this property.
	Alias          *string `pulumi:"alias"`
	BillingScopeId *string `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName *string           `pulumi:"subscriptionName"`
	Tags             map[string]string `pulumi:"tags"`
	// The ID of the Tenant to which the subscription belongs.
	TenantId *string `pulumi:"tenantId"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload *string `pulumi:"workload"`
}

type SubscriptionState struct {
	// The Alias Name of the subscription. If omitted a new UUID will be generated for this property.
	Alias          pulumi.StringPtrInput
	BillingScopeId pulumi.StringPtrInput
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringPtrInput
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	// The ID of the Tenant to which the subscription belongs.
	TenantId pulumi.StringPtrInput
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// The Alias Name of the subscription. If omitted a new UUID will be generated for this property.
	Alias          *string `pulumi:"alias"`
	BillingScopeId *string `pulumi:"billingScopeId"`
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload *string `pulumi:"workload"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The Alias Name of the subscription. If omitted a new UUID will be generated for this property.
	Alias          pulumi.StringPtrInput
	BillingScopeId pulumi.StringPtrInput
	// The ID of the Subscription. Cannot be specified with `billingAccount`, `billingProfile`, `enrollmentAccount`, or `invoiceSection` Changing this forces a new Subscription to be created.
	SubscriptionId pulumi.StringPtrInput
	// The Name of the Subscription. This is the Display Name in the portal.
	SubscriptionName pulumi.StringInput
	// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
	Workload pulumi.StringPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

type SubscriptionPtrInput interface {
	pulumi.Input

	ToSubscriptionPtrOutput() SubscriptionPtrOutput
	ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput
}

type subscriptionPtrType SubscriptionArgs

func (*subscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return i.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (i *subscriptionPtrType) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionPtrOutput)
}

// SubscriptionArrayInput is an input type that accepts SubscriptionArray and SubscriptionArrayOutput values.
// You can construct a concrete instance of `SubscriptionArrayInput` via:
//
//          SubscriptionArray{ SubscriptionArgs{...} }
type SubscriptionArrayInput interface {
	pulumi.Input

	ToSubscriptionArrayOutput() SubscriptionArrayOutput
	ToSubscriptionArrayOutputWithContext(context.Context) SubscriptionArrayOutput
}

type SubscriptionArray []SubscriptionInput

func (SubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subscription)(nil))
}

func (i SubscriptionArray) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return i.ToSubscriptionArrayOutputWithContext(context.Background())
}

func (i SubscriptionArray) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionArrayOutput)
}

// SubscriptionMapInput is an input type that accepts SubscriptionMap and SubscriptionMapOutput values.
// You can construct a concrete instance of `SubscriptionMapInput` via:
//
//          SubscriptionMap{ "key": SubscriptionArgs{...} }
type SubscriptionMapInput interface {
	pulumi.Input

	ToSubscriptionMapOutput() SubscriptionMapOutput
	ToSubscriptionMapOutputWithContext(context.Context) SubscriptionMapOutput
}

type SubscriptionMap map[string]SubscriptionInput

func (SubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subscription)(nil))
}

func (i SubscriptionMap) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return i.ToSubscriptionMapOutputWithContext(context.Background())
}

func (i SubscriptionMap) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionMapOutput)
}

type SubscriptionOutput struct {
	*pulumi.OutputState
}

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o.ToSubscriptionPtrOutputWithContext(context.Background())
}

func (o SubscriptionOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o.ApplyT(func(v Subscription) *Subscription {
		return &v
	}).(SubscriptionPtrOutput)
}

type SubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (SubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil))
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutput() SubscriptionPtrOutput {
	return o
}

func (o SubscriptionPtrOutput) ToSubscriptionPtrOutputWithContext(ctx context.Context) SubscriptionPtrOutput {
	return o
}

type SubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subscription)(nil))
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) Index(i pulumi.IntInput) SubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].([]Subscription)[vs[1].(int)]
	}).(SubscriptionOutput)
}

type SubscriptionMapOutput struct{ *pulumi.OutputState }

func (SubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subscription)(nil))
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) MapIndex(k pulumi.StringInput) SubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subscription {
		return vs[0].(map[string]Subscription)[vs[1].(string)]
	}).(SubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
	pulumi.RegisterOutputType(SubscriptionPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionMapOutput{})
}
