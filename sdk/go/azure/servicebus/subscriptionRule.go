// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a ServiceBus Subscription Rule.
//
// ## Example Usage
// ### SQL Filter)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
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
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			NamespaceId:        exampleNamespace.ID(),
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubscription, err := servicebus.NewSubscription(ctx, "exampleSubscription", &servicebus.SubscriptionArgs{
// 			TopicId:          exampleTopic.ID(),
// 			MaxDeliveryCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewSubscriptionRule(ctx, "exampleSubscriptionRule", &servicebus.SubscriptionRuleArgs{
// 			SubscriptionId: exampleSubscription.ID(),
// 			FilterType:     pulumi.String("SqlFilter"),
// 			SqlFilter:      pulumi.String("colour = 'red'"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Correlation Filter)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/servicebus"
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
// 		exampleNamespace, err := servicebus.NewNamespace(ctx, "exampleNamespace", &servicebus.NamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Tags: pulumi.StringMap{
// 				"source": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTopic, err := servicebus.NewTopic(ctx, "exampleTopic", &servicebus.TopicArgs{
// 			NamespaceId:        exampleNamespace.ID(),
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubscription, err := servicebus.NewSubscription(ctx, "exampleSubscription", &servicebus.SubscriptionArgs{
// 			TopicId:          exampleTopic.ID(),
// 			MaxDeliveryCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewSubscriptionRule(ctx, "exampleSubscriptionRule", &servicebus.SubscriptionRuleArgs{
// 			SubscriptionId: exampleSubscription.ID(),
// 			FilterType:     pulumi.String("CorrelationFilter"),
// 			CorrelationFilter: &servicebus.SubscriptionRuleCorrelationFilterArgs{
// 				CorrelationId: pulumi.String("high"),
// 				Label:         pulumi.String("red"),
// 				Properties: pulumi.StringMap{
// 					"customProperty": pulumi.String("value"),
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
// Service Bus Subscription Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicebus/subscriptionRule:SubscriptionRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
// ```
type SubscriptionRule struct {
	pulumi.CustomResourceState

	// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
	CorrelationFilter SubscriptionRuleCorrelationFilterPtrOutput `pulumi:"correlationFilter"`
	// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated: Deprecated in favor of "subscription_id"
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
	SqlFilter pulumi.StringPtrOutput `pulumi:"sqlFilter"`
	// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Deprecated: Deprecated in favor of "subscription_id"
	SubscriptionName pulumi.StringOutput `pulumi:"subscriptionName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewSubscriptionRule registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionRule(ctx *pulumi.Context,
	name string, args *SubscriptionRuleArgs, opts ...pulumi.ResourceOption) (*SubscriptionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilterType == nil {
		return nil, errors.New("invalid value for required argument 'FilterType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/subscriptionRule:SubscriptionRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionRule
	err := ctx.RegisterResource("azure:servicebus/subscriptionRule:SubscriptionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionRule gets an existing SubscriptionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionRuleState, opts ...pulumi.ResourceOption) (*SubscriptionRule, error) {
	var resource SubscriptionRule
	err := ctx.ReadResource("azure:servicebus/subscriptionRule:SubscriptionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionRule resources.
type subscriptionRuleState struct {
	// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
	Action *string `pulumi:"action"`
	// A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
	CorrelationFilter *SubscriptionRuleCorrelationFilter `pulumi:"correlationFilter"`
	// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
	FilterType *string `pulumi:"filterType"`
	// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: Deprecated in favor of "subscription_id"
	NamespaceName *string `pulumi:"namespaceName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
	SqlFilter *string `pulumi:"sqlFilter"`
	// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Deprecated: Deprecated in favor of "subscription_id"
	SubscriptionName *string `pulumi:"subscriptionName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	TopicName *string `pulumi:"topicName"`
}

type SubscriptionRuleState struct {
	// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
	Action pulumi.StringPtrInput
	// A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
	CorrelationFilter SubscriptionRuleCorrelationFilterPtrInput
	// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
	FilterType pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	NamespaceName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	ResourceGroupName pulumi.StringPtrInput
	// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
	SqlFilter pulumi.StringPtrInput
	// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	SubscriptionName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	TopicName pulumi.StringPtrInput
}

func (SubscriptionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionRuleState)(nil)).Elem()
}

type subscriptionRuleArgs struct {
	// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
	Action *string `pulumi:"action"`
	// A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
	CorrelationFilter *SubscriptionRuleCorrelationFilter `pulumi:"correlationFilter"`
	// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
	FilterType string `pulumi:"filterType"`
	// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: Deprecated in favor of "subscription_id"
	NamespaceName *string `pulumi:"namespaceName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
	SqlFilter *string `pulumi:"sqlFilter"`
	// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Deprecated: Deprecated in favor of "subscription_id"
	SubscriptionName *string `pulumi:"subscriptionName"`
	// Deprecated: Deprecated in favor of "subscription_id"
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a SubscriptionRule resource.
type SubscriptionRuleArgs struct {
	// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
	Action pulumi.StringPtrInput
	// A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
	CorrelationFilter SubscriptionRuleCorrelationFilterPtrInput
	// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
	FilterType pulumi.StringInput
	// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	NamespaceName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	ResourceGroupName pulumi.StringPtrInput
	// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
	SqlFilter pulumi.StringPtrInput
	// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
	SubscriptionId pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	SubscriptionName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "subscription_id"
	TopicName pulumi.StringPtrInput
}

func (SubscriptionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionRuleArgs)(nil)).Elem()
}

type SubscriptionRuleInput interface {
	pulumi.Input

	ToSubscriptionRuleOutput() SubscriptionRuleOutput
	ToSubscriptionRuleOutputWithContext(ctx context.Context) SubscriptionRuleOutput
}

func (*SubscriptionRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionRule)(nil)).Elem()
}

func (i *SubscriptionRule) ToSubscriptionRuleOutput() SubscriptionRuleOutput {
	return i.ToSubscriptionRuleOutputWithContext(context.Background())
}

func (i *SubscriptionRule) ToSubscriptionRuleOutputWithContext(ctx context.Context) SubscriptionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleOutput)
}

// SubscriptionRuleArrayInput is an input type that accepts SubscriptionRuleArray and SubscriptionRuleArrayOutput values.
// You can construct a concrete instance of `SubscriptionRuleArrayInput` via:
//
//          SubscriptionRuleArray{ SubscriptionRuleArgs{...} }
type SubscriptionRuleArrayInput interface {
	pulumi.Input

	ToSubscriptionRuleArrayOutput() SubscriptionRuleArrayOutput
	ToSubscriptionRuleArrayOutputWithContext(context.Context) SubscriptionRuleArrayOutput
}

type SubscriptionRuleArray []SubscriptionRuleInput

func (SubscriptionRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionRule)(nil)).Elem()
}

func (i SubscriptionRuleArray) ToSubscriptionRuleArrayOutput() SubscriptionRuleArrayOutput {
	return i.ToSubscriptionRuleArrayOutputWithContext(context.Background())
}

func (i SubscriptionRuleArray) ToSubscriptionRuleArrayOutputWithContext(ctx context.Context) SubscriptionRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleArrayOutput)
}

// SubscriptionRuleMapInput is an input type that accepts SubscriptionRuleMap and SubscriptionRuleMapOutput values.
// You can construct a concrete instance of `SubscriptionRuleMapInput` via:
//
//          SubscriptionRuleMap{ "key": SubscriptionRuleArgs{...} }
type SubscriptionRuleMapInput interface {
	pulumi.Input

	ToSubscriptionRuleMapOutput() SubscriptionRuleMapOutput
	ToSubscriptionRuleMapOutputWithContext(context.Context) SubscriptionRuleMapOutput
}

type SubscriptionRuleMap map[string]SubscriptionRuleInput

func (SubscriptionRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionRule)(nil)).Elem()
}

func (i SubscriptionRuleMap) ToSubscriptionRuleMapOutput() SubscriptionRuleMapOutput {
	return i.ToSubscriptionRuleMapOutputWithContext(context.Background())
}

func (i SubscriptionRuleMap) ToSubscriptionRuleMapOutputWithContext(ctx context.Context) SubscriptionRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleMapOutput)
}

type SubscriptionRuleOutput struct{ *pulumi.OutputState }

func (SubscriptionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionRule)(nil)).Elem()
}

func (o SubscriptionRuleOutput) ToSubscriptionRuleOutput() SubscriptionRuleOutput {
	return o
}

func (o SubscriptionRuleOutput) ToSubscriptionRuleOutputWithContext(ctx context.Context) SubscriptionRuleOutput {
	return o
}

type SubscriptionRuleArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionRule)(nil)).Elem()
}

func (o SubscriptionRuleArrayOutput) ToSubscriptionRuleArrayOutput() SubscriptionRuleArrayOutput {
	return o
}

func (o SubscriptionRuleArrayOutput) ToSubscriptionRuleArrayOutputWithContext(ctx context.Context) SubscriptionRuleArrayOutput {
	return o
}

func (o SubscriptionRuleArrayOutput) Index(i pulumi.IntInput) SubscriptionRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscriptionRule {
		return vs[0].([]*SubscriptionRule)[vs[1].(int)]
	}).(SubscriptionRuleOutput)
}

type SubscriptionRuleMapOutput struct{ *pulumi.OutputState }

func (SubscriptionRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionRule)(nil)).Elem()
}

func (o SubscriptionRuleMapOutput) ToSubscriptionRuleMapOutput() SubscriptionRuleMapOutput {
	return o
}

func (o SubscriptionRuleMapOutput) ToSubscriptionRuleMapOutputWithContext(ctx context.Context) SubscriptionRuleMapOutput {
	return o
}

func (o SubscriptionRuleMapOutput) MapIndex(k pulumi.StringInput) SubscriptionRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscriptionRule {
		return vs[0].(map[string]*SubscriptionRule)[vs[1].(string)]
	}).(SubscriptionRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionRuleInput)(nil)).Elem(), &SubscriptionRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionRuleArrayInput)(nil)).Elem(), SubscriptionRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionRuleMapInput)(nil)).Elem(), SubscriptionRuleMap{})
	pulumi.RegisterOutputType(SubscriptionRuleOutput{})
	pulumi.RegisterOutputType(SubscriptionRuleArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionRuleMapOutput{})
}
