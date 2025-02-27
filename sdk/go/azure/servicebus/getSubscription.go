// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing ServiceBus Subscription.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicebus.LookupSubscription(ctx, &servicebus.LookupSubscriptionArgs{
// 			Name:              "examplesubscription",
// 			ResourceGroupName: "exampleresources",
// 			NamespaceName:     "examplenamespace",
// 			TopicName:         "exampletopic",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("servicebusSubscription", data.Azurerm_servicebus_namespace.Example)
// 		return nil
// 	})
// }
// ```
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("azure:servicebus/getSubscription:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubscription.
type LookupSubscriptionArgs struct {
	// Specifies the name of the ServiceBus Subscription.
	Name string `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the ServiceBus Topic.
	TopicName string `pulumi:"topicName"`
}

// A collection of values returned by getSubscription.
type LookupSubscriptionResult struct {
	// The idle interval after which the topic is automatically deleted.
	AutoDeleteOnIdle string `pulumi:"autoDeleteOnIdle"`
	// Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?
	DeadLetteringOnFilterEvaluationError bool `pulumi:"deadLetteringOnFilterEvaluationError"`
	// Does the Service Bus Subscription have dead letter support when a message expires?
	DeadLetteringOnMessageExpiration bool `pulumi:"deadLetteringOnMessageExpiration"`
	// The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTtl string `pulumi:"defaultMessageTtl"`
	// Are batched operations enabled on this ServiceBus Subscription?
	EnableBatchedOperations bool `pulumi:"enableBatchedOperations"`
	// The name of a Queue or Topic to automatically forward Dead Letter messages to.
	ForwardDeadLetteredMessagesTo string `pulumi:"forwardDeadLetteredMessagesTo"`
	// The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.
	ForwardTo string `pulumi:"forwardTo"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The lock duration for the subscription.
	LockDuration string `pulumi:"lockDuration"`
	// The maximum number of deliveries.
	MaxDeliveryCount int    `pulumi:"maxDeliveryCount"`
	Name             string `pulumi:"name"`
	NamespaceName    string `pulumi:"namespaceName"`
	// Whether or not this ServiceBus Subscription supports session.
	RequiresSession   bool   `pulumi:"requiresSession"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			return *r, err
		}).(LookupSubscriptionResultOutput)
}

// A collection of arguments for invoking getSubscription.
type LookupSubscriptionOutputArgs struct {
	// Specifies the name of the ServiceBus Subscription.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the ServiceBus Namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the ServiceBus Topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}

// A collection of values returned by getSubscription.
type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

// The idle interval after which the topic is automatically deleted.
func (o LookupSubscriptionResultOutput) AutoDeleteOnIdle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.AutoDeleteOnIdle }).(pulumi.StringOutput)
}

// Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?
func (o LookupSubscriptionResultOutput) DeadLetteringOnFilterEvaluationError() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.DeadLetteringOnFilterEvaluationError }).(pulumi.BoolOutput)
}

// Does the Service Bus Subscription have dead letter support when a message expires?
func (o LookupSubscriptionResultOutput) DeadLetteringOnMessageExpiration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolOutput)
}

// The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
func (o LookupSubscriptionResultOutput) DefaultMessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.DefaultMessageTtl }).(pulumi.StringOutput)
}

// Are batched operations enabled on this ServiceBus Subscription?
func (o LookupSubscriptionResultOutput) EnableBatchedOperations() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.EnableBatchedOperations }).(pulumi.BoolOutput)
}

// The name of a Queue or Topic to automatically forward Dead Letter messages to.
func (o LookupSubscriptionResultOutput) ForwardDeadLetteredMessagesTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringOutput)
}

// The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.
func (o LookupSubscriptionResultOutput) ForwardTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.ForwardTo }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The lock duration for the subscription.
func (o LookupSubscriptionResultOutput) LockDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.LockDuration }).(pulumi.StringOutput)
}

// The maximum number of deliveries.
func (o LookupSubscriptionResultOutput) MaxDeliveryCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) int { return v.MaxDeliveryCount }).(pulumi.IntOutput)
}

func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

// Whether or not this ServiceBus Subscription supports session.
func (o LookupSubscriptionResultOutput) RequiresSession() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.RequiresSession }).(pulumi.BoolOutput)
}

func (o LookupSubscriptionResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.TopicName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
