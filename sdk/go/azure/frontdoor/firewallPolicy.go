// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frontdoor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Front Door Web Application Firewall Policy instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/frontdoor"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = frontdoor.NewFirewallPolicy(ctx, "exampleFirewallPolicy", &frontdoor.FirewallPolicyArgs{
// 			ResourceGroupName:             exampleResourceGroup.Name,
// 			Enabled:                       pulumi.Bool(true),
// 			Mode:                          pulumi.String("Prevention"),
// 			RedirectUrl:                   pulumi.String("https://www.contoso.com"),
// 			CustomBlockResponseStatusCode: pulumi.Int(403),
// 			CustomBlockResponseBody:       pulumi.String("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
// 			CustomRules: frontdoor.FirewallPolicyCustomRuleArray{
// 				&frontdoor.FirewallPolicyCustomRuleArgs{
// 					Name:                       pulumi.String("Rule1"),
// 					Enabled:                    pulumi.Bool(true),
// 					Priority:                   pulumi.Int(1),
// 					RateLimitDurationInMinutes: pulumi.Int(1),
// 					RateLimitThreshold:         pulumi.Int(10),
// 					Type:                       pulumi.String("MatchRule"),
// 					Action:                     pulumi.String("Block"),
// 					MatchConditions: frontdoor.FirewallPolicyCustomRuleMatchConditionArray{
// 						&frontdoor.FirewallPolicyCustomRuleMatchConditionArgs{
// 							MatchVariable:     pulumi.String("RemoteAddr"),
// 							Operator:          pulumi.String("IPMatch"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("192.168.1.0/24"),
// 								pulumi.String("10.0.0.0/24"),
// 							},
// 						},
// 					},
// 				},
// 				&frontdoor.FirewallPolicyCustomRuleArgs{
// 					Name:                       pulumi.String("Rule2"),
// 					Enabled:                    pulumi.Bool(true),
// 					Priority:                   pulumi.Int(2),
// 					RateLimitDurationInMinutes: pulumi.Int(1),
// 					RateLimitThreshold:         pulumi.Int(10),
// 					Type:                       pulumi.String("MatchRule"),
// 					Action:                     pulumi.String("Block"),
// 					MatchConditions: frontdoor.FirewallPolicyCustomRuleMatchConditionArray{
// 						&frontdoor.FirewallPolicyCustomRuleMatchConditionArgs{
// 							MatchVariable:     pulumi.String("RemoteAddr"),
// 							Operator:          pulumi.String("IPMatch"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("192.168.1.0/24"),
// 							},
// 						},
// 						&frontdoor.FirewallPolicyCustomRuleMatchConditionArgs{
// 							MatchVariable:     pulumi.String("RequestHeader"),
// 							Selector:          pulumi.String("UserAgent"),
// 							Operator:          pulumi.String("Contains"),
// 							NegationCondition: pulumi.Bool(false),
// 							MatchValues: pulumi.StringArray{
// 								pulumi.String("windows"),
// 							},
// 							Transforms: pulumi.StringArray{
// 								pulumi.String("Lowercase"),
// 								pulumi.String("Trim"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			ManagedRules: frontdoor.FirewallPolicyManagedRuleArray{
// 				&frontdoor.FirewallPolicyManagedRuleArgs{
// 					Type:    pulumi.String("DefaultRuleSet"),
// 					Version: pulumi.String("1.0"),
// 					Exclusions: frontdoor.FirewallPolicyManagedRuleExclusionArray{
// 						&frontdoor.FirewallPolicyManagedRuleExclusionArgs{
// 							MatchVariable: pulumi.String("QueryStringArgNames"),
// 							Operator:      pulumi.String("Equals"),
// 							Selector:      pulumi.String("not_suspicious"),
// 						},
// 					},
// 					Overrides: frontdoor.FirewallPolicyManagedRuleOverrideArray{
// 						&frontdoor.FirewallPolicyManagedRuleOverrideArgs{
// 							RuleGroupName: pulumi.String("PHP"),
// 							Rules: frontdoor.FirewallPolicyManagedRuleOverrideRuleArray{
// 								&frontdoor.FirewallPolicyManagedRuleOverrideRuleArgs{
// 									RuleId:  pulumi.String("933100"),
// 									Enabled: pulumi.Bool(false),
// 									Action:  pulumi.String("Block"),
// 								},
// 							},
// 						},
// 						&frontdoor.FirewallPolicyManagedRuleOverrideArgs{
// 							RuleGroupName: pulumi.String("SQLI"),
// 							Exclusions: frontdoor.FirewallPolicyManagedRuleOverrideExclusionArray{
// 								&frontdoor.FirewallPolicyManagedRuleOverrideExclusionArgs{
// 									MatchVariable: pulumi.String("QueryStringArgNames"),
// 									Operator:      pulumi.String("Equals"),
// 									Selector:      pulumi.String("really_not_suspicious"),
// 								},
// 							},
// 							Rules: frontdoor.FirewallPolicyManagedRuleOverrideRuleArray{
// 								&frontdoor.FirewallPolicyManagedRuleOverrideRuleArgs{
// 									RuleId: pulumi.String("942200"),
// 									Action: pulumi.String("Block"),
// 									Exclusions: frontdoor.FirewallPolicyManagedRuleOverrideRuleExclusionArray{
// 										&frontdoor.FirewallPolicyManagedRuleOverrideRuleExclusionArgs{
// 											MatchVariable: pulumi.String("QueryStringArgNames"),
// 											Operator:      pulumi.String("Equals"),
// 											Selector:      pulumi.String("innocent"),
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 				&frontdoor.FirewallPolicyManagedRuleArgs{
// 					Type:    pulumi.String("Microsoft_BotManagerRuleSet"),
// 					Version: pulumi.String("1.0"),
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
// FrontDoor Web Application Firewall Policy can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:frontdoor/firewallPolicy:FirewallPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/examplefdwafpolicy
// ```
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrOutput `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrOutput `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayOutput `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds pulumi.StringArrayOutput `pulumi:"frontendEndpointIds"`
	// The Azure Region where this FrontDoor Firewall Policy exists.
	Location pulumi.StringOutput `pulumi:"location"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayOutput `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrOutput `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("azure:frontdoor/firewallPolicy:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azure:frontdoor/firewallPolicy:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode *int `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules []FirewallPolicyCustomRule `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds []string `pulumi:"frontendEndpointIds"`
	// The Azure Region where this FrontDoor Firewall Policy exists.
	Location *string `pulumi:"location"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules []FirewallPolicyManagedRule `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode *string `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

type FirewallPolicyState struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrInput
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrInput
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayInput
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds pulumi.StringArrayInput
	// The Azure Region where this FrontDoor Firewall Policy exists.
	Location pulumi.StringPtrInput
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayInput
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode *int `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules []FirewallPolicyCustomRule `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules []FirewallPolicyManagedRule `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode *string `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrInput
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrInput
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayInput
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayInput
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrInput interface {
	pulumi.Input

	ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput
	ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput
}

type firewallPolicyPtrType FirewallPolicyArgs

func (*firewallPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

// FirewallPolicyArrayInput is an input type that accepts FirewallPolicyArray and FirewallPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyArrayInput` via:
//
//          FirewallPolicyArray{ FirewallPolicyArgs{...} }
type FirewallPolicyArrayInput interface {
	pulumi.Input

	ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput
	ToFirewallPolicyArrayOutputWithContext(context.Context) FirewallPolicyArrayOutput
}

type FirewallPolicyArray []FirewallPolicyInput

func (FirewallPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallPolicy)(nil))
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return i.ToFirewallPolicyArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyArrayOutput)
}

// FirewallPolicyMapInput is an input type that accepts FirewallPolicyMap and FirewallPolicyMapOutput values.
// You can construct a concrete instance of `FirewallPolicyMapInput` via:
//
//          FirewallPolicyMap{ "key": FirewallPolicyArgs{...} }
type FirewallPolicyMapInput interface {
	pulumi.Input

	ToFirewallPolicyMapOutput() FirewallPolicyMapOutput
	ToFirewallPolicyMapOutputWithContext(context.Context) FirewallPolicyMapOutput
}

type FirewallPolicyMap map[string]FirewallPolicyInput

func (FirewallPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallPolicy)(nil))
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return i.ToFirewallPolicyMapOutputWithContext(context.Background())
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyMapOutput)
}

type FirewallPolicyOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o.ApplyT(func(v FirewallPolicy) *FirewallPolicy {
		return &v
	}).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o
}

type FirewallPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPolicy)(nil))
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) Index(i pulumi.IntInput) FirewallPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].([]FirewallPolicy)[vs[1].(int)]
	}).(FirewallPolicyOutput)
}

type FirewallPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallPolicy)(nil))
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].(map[string]FirewallPolicy)[vs[1].(string)]
	}).(FirewallPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
	pulumi.RegisterOutputType(FirewallPolicyPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyMapOutput{})
}
