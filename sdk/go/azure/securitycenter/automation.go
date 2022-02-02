// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Security Center Automation and Continuous Export. This resource supports three types of destination in the `action`, Logic Apps, Log Analytics and Event Hubs
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/securitycenter"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 			Capacity:          pulumi.Int(2),
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
// 		exampleAuthorizationRule, err := eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
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
// 		_, err = securitycenter.NewAutomation(ctx, "exampleAutomation", &securitycenter.AutomationArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Actions: securitycenter.AutomationActionArray{
// 				&securitycenter.AutomationActionArgs{
// 					Type:             pulumi.String("EventHub"),
// 					ResourceId:       exampleEventHub.ID(),
// 					ConnectionString: exampleAuthorizationRule.PrimaryConnectionString,
// 				},
// 			},
// 			Sources: securitycenter.AutomationSourceArray{
// 				&securitycenter.AutomationSourceArgs{
// 					EventSource: pulumi.String("Alerts"),
// 					RuleSets: securitycenter.AutomationSourceRuleSetArray{
// 						&securitycenter.AutomationSourceRuleSetArgs{
// 							Rules: securitycenter.AutomationSourceRuleSetRuleArray{
// 								&securitycenter.AutomationSourceRuleSetRuleArgs{
// 									PropertyPath:  pulumi.String("properties.metadata.severity"),
// 									Operator:      pulumi.String("Equals"),
// 									ExpectedValue: pulumi.String("High"),
// 									PropertyType:  pulumi.String("String"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Scopes: pulumi.StringArray{
// 				pulumi.String(fmt.Sprintf("%v%v", "/subscriptions/", current.SubscriptionId)),
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
// Security Center Automations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:securitycenter/automation:Automation example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Security/automations/automation1
// ```
type Automation struct {
	pulumi.CustomResourceState

	// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
	Actions AutomationActionArrayOutput `pulumi:"actions"`
	// Specifies the description for the Security Center Automation.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Boolean to enable or disable this Security Center Automation.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
	Sources AutomationSourceArrayOutput `pulumi:"sources"`
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAutomation registers a new resource with the given unique name, arguments, and options.
func NewAutomation(ctx *pulumi.Context,
	name string, args *AutomationArgs, opts ...pulumi.ResourceOption) (*Automation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	var resource Automation
	err := ctx.RegisterResource("azure:securitycenter/automation:Automation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomation gets an existing Automation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationState, opts ...pulumi.ResourceOption) (*Automation, error) {
	var resource Automation
	err := ctx.ReadResource("azure:securitycenter/automation:Automation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Automation resources.
type automationState struct {
	// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
	Actions []AutomationAction `pulumi:"actions"`
	// Specifies the description for the Security Center Automation.
	Description *string `pulumi:"description"`
	// Boolean to enable or disable this Security Center Automation.
	Enabled *bool `pulumi:"enabled"`
	// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
	Scopes []string `pulumi:"scopes"`
	// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
	Sources []AutomationSource `pulumi:"sources"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AutomationState struct {
	// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
	Actions AutomationActionArrayInput
	// Specifies the description for the Security Center Automation.
	Description pulumi.StringPtrInput
	// Boolean to enable or disable this Security Center Automation.
	Enabled pulumi.BoolPtrInput
	// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
	Scopes pulumi.StringArrayInput
	// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
	Sources AutomationSourceArrayInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
}

func (AutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationState)(nil)).Elem()
}

type automationArgs struct {
	// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
	Actions []AutomationAction `pulumi:"actions"`
	// Specifies the description for the Security Center Automation.
	Description *string `pulumi:"description"`
	// Boolean to enable or disable this Security Center Automation.
	Enabled *bool `pulumi:"enabled"`
	// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
	Scopes []string `pulumi:"scopes"`
	// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
	Sources []AutomationSource `pulumi:"sources"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Automation resource.
type AutomationArgs struct {
	// One or more `action` blocks as defined below. An `action` tells this automation where the data is to be sent to upon being evaluated by the rules in the `source`.
	Actions AutomationActionArrayInput
	// Specifies the description for the Security Center Automation.
	Description pulumi.StringPtrInput
	// Boolean to enable or disable this Security Center Automation.
	Enabled pulumi.BoolPtrInput
	// The Azure Region where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Security Center Automation. Changing this forces a new Security Center Automation to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Security Center Automation should exist. Changing this forces a new Security Center Automation to be created.
	ResourceGroupName pulumi.StringInput
	// A list of scopes on which the automation logic is applied, at least one is required. Supported scopes are a subscription (in this format `/subscriptions/00000000-0000-0000-0000-000000000000`) or a resource group under that subscription (in the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example`). The automation will only apply on defined scopes.
	Scopes pulumi.StringArrayInput
	// One or more `source` blocks as defined below. A `source` defines what data types will be processed and a set of rules to filter that data.
	Sources AutomationSourceArrayInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
}

func (AutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationArgs)(nil)).Elem()
}

type AutomationInput interface {
	pulumi.Input

	ToAutomationOutput() AutomationOutput
	ToAutomationOutputWithContext(ctx context.Context) AutomationOutput
}

func (*Automation) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (i *Automation) ToAutomationOutput() AutomationOutput {
	return i.ToAutomationOutputWithContext(context.Background())
}

func (i *Automation) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationOutput)
}

// AutomationArrayInput is an input type that accepts AutomationArray and AutomationArrayOutput values.
// You can construct a concrete instance of `AutomationArrayInput` via:
//
//          AutomationArray{ AutomationArgs{...} }
type AutomationArrayInput interface {
	pulumi.Input

	ToAutomationArrayOutput() AutomationArrayOutput
	ToAutomationArrayOutputWithContext(context.Context) AutomationArrayOutput
}

type AutomationArray []AutomationInput

func (AutomationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automation)(nil)).Elem()
}

func (i AutomationArray) ToAutomationArrayOutput() AutomationArrayOutput {
	return i.ToAutomationArrayOutputWithContext(context.Background())
}

func (i AutomationArray) ToAutomationArrayOutputWithContext(ctx context.Context) AutomationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationArrayOutput)
}

// AutomationMapInput is an input type that accepts AutomationMap and AutomationMapOutput values.
// You can construct a concrete instance of `AutomationMapInput` via:
//
//          AutomationMap{ "key": AutomationArgs{...} }
type AutomationMapInput interface {
	pulumi.Input

	ToAutomationMapOutput() AutomationMapOutput
	ToAutomationMapOutputWithContext(context.Context) AutomationMapOutput
}

type AutomationMap map[string]AutomationInput

func (AutomationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automation)(nil)).Elem()
}

func (i AutomationMap) ToAutomationMapOutput() AutomationMapOutput {
	return i.ToAutomationMapOutputWithContext(context.Background())
}

func (i AutomationMap) ToAutomationMapOutputWithContext(ctx context.Context) AutomationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationMapOutput)
}

type AutomationOutput struct{ *pulumi.OutputState }

func (AutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (o AutomationOutput) ToAutomationOutput() AutomationOutput {
	return o
}

func (o AutomationOutput) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return o
}

type AutomationArrayOutput struct{ *pulumi.OutputState }

func (AutomationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Automation)(nil)).Elem()
}

func (o AutomationArrayOutput) ToAutomationArrayOutput() AutomationArrayOutput {
	return o
}

func (o AutomationArrayOutput) ToAutomationArrayOutputWithContext(ctx context.Context) AutomationArrayOutput {
	return o
}

func (o AutomationArrayOutput) Index(i pulumi.IntInput) AutomationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Automation {
		return vs[0].([]*Automation)[vs[1].(int)]
	}).(AutomationOutput)
}

type AutomationMapOutput struct{ *pulumi.OutputState }

func (AutomationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Automation)(nil)).Elem()
}

func (o AutomationMapOutput) ToAutomationMapOutput() AutomationMapOutput {
	return o
}

func (o AutomationMapOutput) ToAutomationMapOutputWithContext(ctx context.Context) AutomationMapOutput {
	return o
}

func (o AutomationMapOutput) MapIndex(k pulumi.StringInput) AutomationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Automation {
		return vs[0].(map[string]*Automation)[vs[1].(string)]
	}).(AutomationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationInput)(nil)).Elem(), &Automation{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationArrayInput)(nil)).Elem(), AutomationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationMapInput)(nil)).Elem(), AutomationMap{})
	pulumi.RegisterOutputType(AutomationOutput{})
	pulumi.RegisterOutputType(AutomationArrayOutput{})
	pulumi.RegisterOutputType(AutomationMapOutput{})
}
