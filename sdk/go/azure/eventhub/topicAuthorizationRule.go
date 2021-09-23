// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.
//
// ## Example Usage
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
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			NamespaceName:     exampleNamespace.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewTopicAuthorizationRule(ctx, "exampleTopicAuthorizationRule", &servicebus.TopicAuthorizationRuleArgs{
// 			NamespaceName:     exampleNamespace.Name,
// 			TopicName:         exampleTopic.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Listen:            pulumi.Bool(true),
// 			Send:              pulumi.Bool(false),
// 			Manage:            pulumi.Bool(false),
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
// ServiceBus Topic authorization rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/topics/topic1/authorizationRules/rule1
// ```
//
// Deprecated: azure.eventhub.TopicAuthorizationRule has been deprecated in favor of azure.servicebus.TopicAuthorizationRule
type TopicAuthorizationRule struct {
	pulumi.CustomResourceState

	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias pulumi.StringOutput `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias pulumi.StringOutput `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopicAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewTopicAuthorizationRule(ctx *pulumi.Context,
	name string, args *TopicAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource TopicAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicAuthorizationRule gets an existing TopicAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicAuthorizationRuleState, opts ...pulumi.ResourceOption) (*TopicAuthorizationRule, error) {
	var resource TopicAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicAuthorizationRule resources.
type topicAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias *string `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias *string `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName *string `pulumi:"topicName"`
}

type TopicAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the ServiceBus Topic authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias pulumi.StringPtrInput
	// The Primary Key for the ServiceBus Topic authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the ServiceBus Topic authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias pulumi.StringPtrInput
	// The Secondary Key for the ServiceBus Topic authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringPtrInput
}

func (TopicAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleState)(nil)).Elem()
}

type topicAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicAuthorizationRule resource.
type TopicAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Topic. Changing this forces a new resource to be created.
	TopicName pulumi.StringInput
}

func (TopicAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicAuthorizationRuleArgs)(nil)).Elem()
}

type TopicAuthorizationRuleInput interface {
	pulumi.Input

	ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput
	ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput
}

func (*TopicAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRule)(nil))
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return i.ToTopicAuthorizationRuleOutputWithContext(context.Background())
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRuleOutput)
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRulePtrOutput() TopicAuthorizationRulePtrOutput {
	return i.ToTopicAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *TopicAuthorizationRule) ToTopicAuthorizationRulePtrOutputWithContext(ctx context.Context) TopicAuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRulePtrOutput)
}

type TopicAuthorizationRulePtrInput interface {
	pulumi.Input

	ToTopicAuthorizationRulePtrOutput() TopicAuthorizationRulePtrOutput
	ToTopicAuthorizationRulePtrOutputWithContext(ctx context.Context) TopicAuthorizationRulePtrOutput
}

type topicAuthorizationRulePtrType TopicAuthorizationRuleArgs

func (*topicAuthorizationRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicAuthorizationRule)(nil))
}

func (i *topicAuthorizationRulePtrType) ToTopicAuthorizationRulePtrOutput() TopicAuthorizationRulePtrOutput {
	return i.ToTopicAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *topicAuthorizationRulePtrType) ToTopicAuthorizationRulePtrOutputWithContext(ctx context.Context) TopicAuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRulePtrOutput)
}

// TopicAuthorizationRuleArrayInput is an input type that accepts TopicAuthorizationRuleArray and TopicAuthorizationRuleArrayOutput values.
// You can construct a concrete instance of `TopicAuthorizationRuleArrayInput` via:
//
//          TopicAuthorizationRuleArray{ TopicAuthorizationRuleArgs{...} }
type TopicAuthorizationRuleArrayInput interface {
	pulumi.Input

	ToTopicAuthorizationRuleArrayOutput() TopicAuthorizationRuleArrayOutput
	ToTopicAuthorizationRuleArrayOutputWithContext(context.Context) TopicAuthorizationRuleArrayOutput
}

type TopicAuthorizationRuleArray []TopicAuthorizationRuleInput

func (TopicAuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicAuthorizationRule)(nil)).Elem()
}

func (i TopicAuthorizationRuleArray) ToTopicAuthorizationRuleArrayOutput() TopicAuthorizationRuleArrayOutput {
	return i.ToTopicAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i TopicAuthorizationRuleArray) ToTopicAuthorizationRuleArrayOutputWithContext(ctx context.Context) TopicAuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRuleArrayOutput)
}

// TopicAuthorizationRuleMapInput is an input type that accepts TopicAuthorizationRuleMap and TopicAuthorizationRuleMapOutput values.
// You can construct a concrete instance of `TopicAuthorizationRuleMapInput` via:
//
//          TopicAuthorizationRuleMap{ "key": TopicAuthorizationRuleArgs{...} }
type TopicAuthorizationRuleMapInput interface {
	pulumi.Input

	ToTopicAuthorizationRuleMapOutput() TopicAuthorizationRuleMapOutput
	ToTopicAuthorizationRuleMapOutputWithContext(context.Context) TopicAuthorizationRuleMapOutput
}

type TopicAuthorizationRuleMap map[string]TopicAuthorizationRuleInput

func (TopicAuthorizationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicAuthorizationRule)(nil)).Elem()
}

func (i TopicAuthorizationRuleMap) ToTopicAuthorizationRuleMapOutput() TopicAuthorizationRuleMapOutput {
	return i.ToTopicAuthorizationRuleMapOutputWithContext(context.Background())
}

func (i TopicAuthorizationRuleMap) ToTopicAuthorizationRuleMapOutputWithContext(ctx context.Context) TopicAuthorizationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicAuthorizationRuleMapOutput)
}

type TopicAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (TopicAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicAuthorizationRule)(nil))
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutput() TopicAuthorizationRuleOutput {
	return o
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRuleOutputWithContext(ctx context.Context) TopicAuthorizationRuleOutput {
	return o
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRulePtrOutput() TopicAuthorizationRulePtrOutput {
	return o.ToTopicAuthorizationRulePtrOutputWithContext(context.Background())
}

func (o TopicAuthorizationRuleOutput) ToTopicAuthorizationRulePtrOutputWithContext(ctx context.Context) TopicAuthorizationRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TopicAuthorizationRule) *TopicAuthorizationRule {
		return &v
	}).(TopicAuthorizationRulePtrOutput)
}

type TopicAuthorizationRulePtrOutput struct{ *pulumi.OutputState }

func (TopicAuthorizationRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicAuthorizationRule)(nil))
}

func (o TopicAuthorizationRulePtrOutput) ToTopicAuthorizationRulePtrOutput() TopicAuthorizationRulePtrOutput {
	return o
}

func (o TopicAuthorizationRulePtrOutput) ToTopicAuthorizationRulePtrOutputWithContext(ctx context.Context) TopicAuthorizationRulePtrOutput {
	return o
}

func (o TopicAuthorizationRulePtrOutput) Elem() TopicAuthorizationRuleOutput {
	return o.ApplyT(func(v *TopicAuthorizationRule) TopicAuthorizationRule {
		if v != nil {
			return *v
		}
		var ret TopicAuthorizationRule
		return ret
	}).(TopicAuthorizationRuleOutput)
}

type TopicAuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (TopicAuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicAuthorizationRule)(nil))
}

func (o TopicAuthorizationRuleArrayOutput) ToTopicAuthorizationRuleArrayOutput() TopicAuthorizationRuleArrayOutput {
	return o
}

func (o TopicAuthorizationRuleArrayOutput) ToTopicAuthorizationRuleArrayOutputWithContext(ctx context.Context) TopicAuthorizationRuleArrayOutput {
	return o
}

func (o TopicAuthorizationRuleArrayOutput) Index(i pulumi.IntInput) TopicAuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicAuthorizationRule {
		return vs[0].([]TopicAuthorizationRule)[vs[1].(int)]
	}).(TopicAuthorizationRuleOutput)
}

type TopicAuthorizationRuleMapOutput struct{ *pulumi.OutputState }

func (TopicAuthorizationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TopicAuthorizationRule)(nil))
}

func (o TopicAuthorizationRuleMapOutput) ToTopicAuthorizationRuleMapOutput() TopicAuthorizationRuleMapOutput {
	return o
}

func (o TopicAuthorizationRuleMapOutput) ToTopicAuthorizationRuleMapOutputWithContext(ctx context.Context) TopicAuthorizationRuleMapOutput {
	return o
}

func (o TopicAuthorizationRuleMapOutput) MapIndex(k pulumi.StringInput) TopicAuthorizationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TopicAuthorizationRule {
		return vs[0].(map[string]TopicAuthorizationRule)[vs[1].(string)]
	}).(TopicAuthorizationRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(TopicAuthorizationRulePtrOutput{})
	pulumi.RegisterOutputType(TopicAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(TopicAuthorizationRuleMapOutput{})
}
