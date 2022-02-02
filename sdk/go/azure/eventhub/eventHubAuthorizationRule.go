// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Event Hubs authorization Rule within an Event Hub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
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
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Basic"),
// 			Capacity:          pulumi.Int(2),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
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
// EventHub Authorization Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
// ```
//
// Deprecated: azure.eventhub.EventHubAuthorizationRule has been deprecated in favor of azure.eventhub.AuthorizationRule
type EventHubAuthorizationRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias pulumi.StringOutput `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias pulumi.StringOutput `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewEventHubAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *EventHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource EventHubAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubAuthorizationRule gets an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	var resource EventHubAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubAuthorizationRule resources.
type eventHubAuthorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias *string `pulumi:"primaryConnectionStringAlias"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias *string `pulumi:"secondaryConnectionStringAlias"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type EventHubAuthorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The alias of the Primary Connection String for the Event Hubs authorization Rule, which is generated when disaster recovery is enabled.
	PrimaryConnectionStringAlias pulumi.StringPtrInput
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The alias of the Secondary Connection String for the Event Hubs Authorization Rule, which is generated when disaster recovery is enabled.
	SecondaryConnectionStringAlias pulumi.StringPtrInput
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleState)(nil)).Elem()
}

type eventHubAuthorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a EventHubAuthorizationRule resource.
type EventHubAuthorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleArgs)(nil)).Elem()
}

type EventHubAuthorizationRuleInput interface {
	pulumi.Input

	ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput
	ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput
}

func (*EventHubAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return i.ToEventHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubAuthorizationRuleOutput)
}

// EventHubAuthorizationRuleArrayInput is an input type that accepts EventHubAuthorizationRuleArray and EventHubAuthorizationRuleArrayOutput values.
// You can construct a concrete instance of `EventHubAuthorizationRuleArrayInput` via:
//
//          EventHubAuthorizationRuleArray{ EventHubAuthorizationRuleArgs{...} }
type EventHubAuthorizationRuleArrayInput interface {
	pulumi.Input

	ToEventHubAuthorizationRuleArrayOutput() EventHubAuthorizationRuleArrayOutput
	ToEventHubAuthorizationRuleArrayOutputWithContext(context.Context) EventHubAuthorizationRuleArrayOutput
}

type EventHubAuthorizationRuleArray []EventHubAuthorizationRuleInput

func (EventHubAuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventHubAuthorizationRule)(nil)).Elem()
}

func (i EventHubAuthorizationRuleArray) ToEventHubAuthorizationRuleArrayOutput() EventHubAuthorizationRuleArrayOutput {
	return i.ToEventHubAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i EventHubAuthorizationRuleArray) ToEventHubAuthorizationRuleArrayOutputWithContext(ctx context.Context) EventHubAuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubAuthorizationRuleArrayOutput)
}

// EventHubAuthorizationRuleMapInput is an input type that accepts EventHubAuthorizationRuleMap and EventHubAuthorizationRuleMapOutput values.
// You can construct a concrete instance of `EventHubAuthorizationRuleMapInput` via:
//
//          EventHubAuthorizationRuleMap{ "key": EventHubAuthorizationRuleArgs{...} }
type EventHubAuthorizationRuleMapInput interface {
	pulumi.Input

	ToEventHubAuthorizationRuleMapOutput() EventHubAuthorizationRuleMapOutput
	ToEventHubAuthorizationRuleMapOutputWithContext(context.Context) EventHubAuthorizationRuleMapOutput
}

type EventHubAuthorizationRuleMap map[string]EventHubAuthorizationRuleInput

func (EventHubAuthorizationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventHubAuthorizationRule)(nil)).Elem()
}

func (i EventHubAuthorizationRuleMap) ToEventHubAuthorizationRuleMapOutput() EventHubAuthorizationRuleMapOutput {
	return i.ToEventHubAuthorizationRuleMapOutputWithContext(context.Background())
}

func (i EventHubAuthorizationRuleMap) ToEventHubAuthorizationRuleMapOutputWithContext(ctx context.Context) EventHubAuthorizationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubAuthorizationRuleMapOutput)
}

type EventHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (EventHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return o
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return o
}

type EventHubAuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (EventHubAuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventHubAuthorizationRule)(nil)).Elem()
}

func (o EventHubAuthorizationRuleArrayOutput) ToEventHubAuthorizationRuleArrayOutput() EventHubAuthorizationRuleArrayOutput {
	return o
}

func (o EventHubAuthorizationRuleArrayOutput) ToEventHubAuthorizationRuleArrayOutputWithContext(ctx context.Context) EventHubAuthorizationRuleArrayOutput {
	return o
}

func (o EventHubAuthorizationRuleArrayOutput) Index(i pulumi.IntInput) EventHubAuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventHubAuthorizationRule {
		return vs[0].([]*EventHubAuthorizationRule)[vs[1].(int)]
	}).(EventHubAuthorizationRuleOutput)
}

type EventHubAuthorizationRuleMapOutput struct{ *pulumi.OutputState }

func (EventHubAuthorizationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventHubAuthorizationRule)(nil)).Elem()
}

func (o EventHubAuthorizationRuleMapOutput) ToEventHubAuthorizationRuleMapOutput() EventHubAuthorizationRuleMapOutput {
	return o
}

func (o EventHubAuthorizationRuleMapOutput) ToEventHubAuthorizationRuleMapOutputWithContext(ctx context.Context) EventHubAuthorizationRuleMapOutput {
	return o
}

func (o EventHubAuthorizationRuleMapOutput) MapIndex(k pulumi.StringInput) EventHubAuthorizationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventHubAuthorizationRule {
		return vs[0].(map[string]*EventHubAuthorizationRule)[vs[1].(string)]
	}).(EventHubAuthorizationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubAuthorizationRuleInput)(nil)).Elem(), &EventHubAuthorizationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubAuthorizationRuleArrayInput)(nil)).Elem(), EventHubAuthorizationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubAuthorizationRuleMapInput)(nil)).Elem(), EventHubAuthorizationRuleMap{})
	pulumi.RegisterOutputType(EventHubAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(EventHubAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(EventHubAuthorizationRuleMapOutput{})
}
