// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Authorization Rule for a ServiceBus Queue.
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
// 			Location: pulumi.String("West US"),
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
// 		exampleQueue, err := servicebus.NewQueue(ctx, "exampleQueue", &servicebus.QueueArgs{
// 			NamespaceId:        exampleNamespace.ID(),
// 			EnablePartitioning: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicebus.NewQueueAuthorizationRule(ctx, "exampleQueueAuthorizationRule", &servicebus.QueueAuthorizationRuleArgs{
// 			QueueId: exampleQueue.ID(),
// 			Listen:  pulumi.Bool(true),
// 			Send:    pulumi.Bool(true),
// 			Manage:  pulumi.Bool(false),
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
// ServiceBus Queue Authorization Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/queues/queue1/authorizationRules/rule1
// ```
type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated: Deprecated in favor of "queue_id"
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias pulumi.StringOutput `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	QueueId pulumi.StringOutput `pulumi:"queueId"`
	// Deprecated: Deprecated in favor of "queue_id"
	QueueName pulumi.StringOutput `pulumi:"queueName"`
	// Deprecated: Deprecated in favor of "queue_id"
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias pulumi.StringOutput `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewQueueAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil {
		args = &QueueAuthorizationRuleArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/queueAuthorizationRule:QueueAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueueAuthorizationRule gets an existing QueueAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueueAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueAuthorizationRuleState, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	var resource QueueAuthorizationRule
	err := ctx.ReadResource("azure:servicebus/queueAuthorizationRule:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueueAuthorizationRule resources.
type queueAuthorizationRuleState struct {
	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: Deprecated in favor of "queue_id"
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias *string `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	QueueId *string `pulumi:"queueId"`
	// Deprecated: Deprecated in favor of "queue_id"
	QueueName *string `pulumi:"queueName"`
	// Deprecated: Deprecated in favor of "queue_id"
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias *string `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type QueueAuthorizationRuleState struct {
	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
	PrimaryConnectionStringAlias pulumi.StringPtrInput
	// The Primary Key for the Authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	QueueId pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	QueueName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The alias Secondary Connection String for the ServiceBus Namespace
	SecondaryConnectionStringAlias pulumi.StringPtrInput
	// The Secondary Key for the Authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Deprecated: Deprecated in favor of "queue_id"
	NamespaceName *string `pulumi:"namespaceName"`
	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	QueueId *string `pulumi:"queueId"`
	// Deprecated: Deprecated in favor of "queue_id"
	QueueName *string `pulumi:"queueName"`
	// Deprecated: Deprecated in favor of "queue_id"
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a QueueAuthorizationRule resource.
type QueueAuthorizationRuleArgs struct {
	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	NamespaceName pulumi.StringPtrInput
	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	QueueId pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	QueueName pulumi.StringPtrInput
	// Deprecated: Deprecated in favor of "queue_id"
	ResourceGroupName pulumi.StringPtrInput
	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}

type QueueAuthorizationRuleInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput
	ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput
}

func (*QueueAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAuthorizationRule)(nil))
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return i.ToQueueAuthorizationRuleOutputWithContext(context.Background())
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleOutput)
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRulePtrOutput() QueueAuthorizationRulePtrOutput {
	return i.ToQueueAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRulePtrOutputWithContext(ctx context.Context) QueueAuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRulePtrOutput)
}

type QueueAuthorizationRulePtrInput interface {
	pulumi.Input

	ToQueueAuthorizationRulePtrOutput() QueueAuthorizationRulePtrOutput
	ToQueueAuthorizationRulePtrOutputWithContext(ctx context.Context) QueueAuthorizationRulePtrOutput
}

type queueAuthorizationRulePtrType QueueAuthorizationRuleArgs

func (*queueAuthorizationRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil))
}

func (i *queueAuthorizationRulePtrType) ToQueueAuthorizationRulePtrOutput() QueueAuthorizationRulePtrOutput {
	return i.ToQueueAuthorizationRulePtrOutputWithContext(context.Background())
}

func (i *queueAuthorizationRulePtrType) ToQueueAuthorizationRulePtrOutputWithContext(ctx context.Context) QueueAuthorizationRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRulePtrOutput)
}

// QueueAuthorizationRuleArrayInput is an input type that accepts QueueAuthorizationRuleArray and QueueAuthorizationRuleArrayOutput values.
// You can construct a concrete instance of `QueueAuthorizationRuleArrayInput` via:
//
//          QueueAuthorizationRuleArray{ QueueAuthorizationRuleArgs{...} }
type QueueAuthorizationRuleArrayInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleArrayOutput() QueueAuthorizationRuleArrayOutput
	ToQueueAuthorizationRuleArrayOutputWithContext(context.Context) QueueAuthorizationRuleArrayOutput
}

type QueueAuthorizationRuleArray []QueueAuthorizationRuleInput

func (QueueAuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueueAuthorizationRule)(nil)).Elem()
}

func (i QueueAuthorizationRuleArray) ToQueueAuthorizationRuleArrayOutput() QueueAuthorizationRuleArrayOutput {
	return i.ToQueueAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i QueueAuthorizationRuleArray) ToQueueAuthorizationRuleArrayOutputWithContext(ctx context.Context) QueueAuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleArrayOutput)
}

// QueueAuthorizationRuleMapInput is an input type that accepts QueueAuthorizationRuleMap and QueueAuthorizationRuleMapOutput values.
// You can construct a concrete instance of `QueueAuthorizationRuleMapInput` via:
//
//          QueueAuthorizationRuleMap{ "key": QueueAuthorizationRuleArgs{...} }
type QueueAuthorizationRuleMapInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleMapOutput() QueueAuthorizationRuleMapOutput
	ToQueueAuthorizationRuleMapOutputWithContext(context.Context) QueueAuthorizationRuleMapOutput
}

type QueueAuthorizationRuleMap map[string]QueueAuthorizationRuleInput

func (QueueAuthorizationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueueAuthorizationRule)(nil)).Elem()
}

func (i QueueAuthorizationRuleMap) ToQueueAuthorizationRuleMapOutput() QueueAuthorizationRuleMapOutput {
	return i.ToQueueAuthorizationRuleMapOutputWithContext(context.Background())
}

func (i QueueAuthorizationRuleMap) ToQueueAuthorizationRuleMapOutputWithContext(ctx context.Context) QueueAuthorizationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleMapOutput)
}

type QueueAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueAuthorizationRule)(nil))
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRulePtrOutput() QueueAuthorizationRulePtrOutput {
	return o.ToQueueAuthorizationRulePtrOutputWithContext(context.Background())
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRulePtrOutputWithContext(ctx context.Context) QueueAuthorizationRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueueAuthorizationRule) *QueueAuthorizationRule {
		return &v
	}).(QueueAuthorizationRulePtrOutput)
}

type QueueAuthorizationRulePtrOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil))
}

func (o QueueAuthorizationRulePtrOutput) ToQueueAuthorizationRulePtrOutput() QueueAuthorizationRulePtrOutput {
	return o
}

func (o QueueAuthorizationRulePtrOutput) ToQueueAuthorizationRulePtrOutputWithContext(ctx context.Context) QueueAuthorizationRulePtrOutput {
	return o
}

func (o QueueAuthorizationRulePtrOutput) Elem() QueueAuthorizationRuleOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) QueueAuthorizationRule {
		if v != nil {
			return *v
		}
		var ret QueueAuthorizationRule
		return ret
	}).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueueAuthorizationRule)(nil))
}

func (o QueueAuthorizationRuleArrayOutput) ToQueueAuthorizationRuleArrayOutput() QueueAuthorizationRuleArrayOutput {
	return o
}

func (o QueueAuthorizationRuleArrayOutput) ToQueueAuthorizationRuleArrayOutputWithContext(ctx context.Context) QueueAuthorizationRuleArrayOutput {
	return o
}

func (o QueueAuthorizationRuleArrayOutput) Index(i pulumi.IntInput) QueueAuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueueAuthorizationRule {
		return vs[0].([]QueueAuthorizationRule)[vs[1].(int)]
	}).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleMapOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueueAuthorizationRule)(nil))
}

func (o QueueAuthorizationRuleMapOutput) ToQueueAuthorizationRuleMapOutput() QueueAuthorizationRuleMapOutput {
	return o
}

func (o QueueAuthorizationRuleMapOutput) ToQueueAuthorizationRuleMapOutputWithContext(ctx context.Context) QueueAuthorizationRuleMapOutput {
	return o
}

func (o QueueAuthorizationRuleMapOutput) MapIndex(k pulumi.StringInput) QueueAuthorizationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueueAuthorizationRule {
		return vs[0].(map[string]QueueAuthorizationRule)[vs[1].(string)]
	}).(QueueAuthorizationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueAuthorizationRuleInput)(nil)).Elem(), &QueueAuthorizationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueAuthorizationRulePtrInput)(nil)).Elem(), &QueueAuthorizationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueAuthorizationRuleArrayInput)(nil)).Elem(), QueueAuthorizationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueAuthorizationRuleMapInput)(nil)).Elem(), QueueAuthorizationRuleMap{})
	pulumi.RegisterOutputType(QueueAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(QueueAuthorizationRulePtrOutput{})
	pulumi.RegisterOutputType(QueueAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(QueueAuthorizationRuleMapOutput{})
}
