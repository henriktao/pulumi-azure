// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Application Insights Smart Detection Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appinsights.NewSmartDetectionRule(ctx, "exampleSmartDetectionRule", &appinsights.SmartDetectionRuleArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			Enabled:               pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SmartDetectionRule struct {
	pulumi.CustomResourceState

	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	AdditionalEmailRecipients pulumi.StringArrayOutput `pulumi:"additionalEmailRecipients"`
	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`,
	// `Long dependency duration`, `Degradation in server response time`, `Degradation in dependency duration`, `Degradation in trace severity ratio`, `Abnormal rise in exception volume`,
	// `Potential memory leak detected`, `Potential security issue detected`, `Abnormal rise in daily data volume`.  Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Do emails get sent to subscription owners? Defaults to `true`.
	SendEmailsToSubscriptionOwners pulumi.BoolPtrOutput `pulumi:"sendEmailsToSubscriptionOwners"`
}

// NewSmartDetectionRule registers a new resource with the given unique name, arguments, and options.
func NewSmartDetectionRule(ctx *pulumi.Context,
	name string, args *SmartDetectionRuleArgs, opts ...pulumi.ResourceOption) (*SmartDetectionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationInsightsId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationInsightsId'")
	}
	var resource SmartDetectionRule
	err := ctx.RegisterResource("azure:appinsights/smartDetectionRule:SmartDetectionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmartDetectionRule gets an existing SmartDetectionRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmartDetectionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmartDetectionRuleState, opts ...pulumi.ResourceOption) (*SmartDetectionRule, error) {
	var resource SmartDetectionRule
	err := ctx.ReadResource("azure:appinsights/smartDetectionRule:SmartDetectionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmartDetectionRule resources.
type smartDetectionRuleState struct {
	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	AdditionalEmailRecipients []string `pulumi:"additionalEmailRecipients"`
	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`,
	// `Long dependency duration`, `Degradation in server response time`, `Degradation in dependency duration`, `Degradation in trace severity ratio`, `Abnormal rise in exception volume`,
	// `Potential memory leak detected`, `Potential security issue detected`, `Abnormal rise in daily data volume`.  Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Do emails get sent to subscription owners? Defaults to `true`.
	SendEmailsToSubscriptionOwners *bool `pulumi:"sendEmailsToSubscriptionOwners"`
}

type SmartDetectionRuleState struct {
	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	AdditionalEmailRecipients pulumi.StringArrayInput
	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`,
	// `Long dependency duration`, `Degradation in server response time`, `Degradation in dependency duration`, `Degradation in trace severity ratio`, `Abnormal rise in exception volume`,
	// `Potential memory leak detected`, `Potential security issue detected`, `Abnormal rise in daily data volume`.  Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Do emails get sent to subscription owners? Defaults to `true`.
	SendEmailsToSubscriptionOwners pulumi.BoolPtrInput
}

func (SmartDetectionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectionRuleState)(nil)).Elem()
}

type smartDetectionRuleArgs struct {
	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	AdditionalEmailRecipients []string `pulumi:"additionalEmailRecipients"`
	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`,
	// `Long dependency duration`, `Degradation in server response time`, `Degradation in dependency duration`, `Degradation in trace severity ratio`, `Abnormal rise in exception volume`,
	// `Potential memory leak detected`, `Potential security issue detected`, `Abnormal rise in daily data volume`.  Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Do emails get sent to subscription owners? Defaults to `true`.
	SendEmailsToSubscriptionOwners *bool `pulumi:"sendEmailsToSubscriptionOwners"`
}

// The set of arguments for constructing a SmartDetectionRule resource.
type SmartDetectionRuleArgs struct {
	// Specifies a list of additional recipients that will be sent emails on this Application Insights Smart Detection Rule.
	AdditionalEmailRecipients pulumi.StringArrayInput
	// The ID of the Application Insights component on which the Smart Detection Rule operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// Is the Application Insights Smart Detection Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Application Insights Smart Detection Rule. Valid values include `Slow page load time`, `Slow server response time`,
	// `Long dependency duration`, `Degradation in server response time`, `Degradation in dependency duration`, `Degradation in trace severity ratio`, `Abnormal rise in exception volume`,
	// `Potential memory leak detected`, `Potential security issue detected`, `Abnormal rise in daily data volume`.  Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Do emails get sent to subscription owners? Defaults to `true`.
	SendEmailsToSubscriptionOwners pulumi.BoolPtrInput
}

func (SmartDetectionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectionRuleArgs)(nil)).Elem()
}

type SmartDetectionRuleInput interface {
	pulumi.Input

	ToSmartDetectionRuleOutput() SmartDetectionRuleOutput
	ToSmartDetectionRuleOutputWithContext(ctx context.Context) SmartDetectionRuleOutput
}

func (*SmartDetectionRule) ElementType() reflect.Type {
	return reflect.TypeOf((*SmartDetectionRule)(nil))
}

func (i *SmartDetectionRule) ToSmartDetectionRuleOutput() SmartDetectionRuleOutput {
	return i.ToSmartDetectionRuleOutputWithContext(context.Background())
}

func (i *SmartDetectionRule) ToSmartDetectionRuleOutputWithContext(ctx context.Context) SmartDetectionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectionRuleOutput)
}

func (i *SmartDetectionRule) ToSmartDetectionRulePtrOutput() SmartDetectionRulePtrOutput {
	return i.ToSmartDetectionRulePtrOutputWithContext(context.Background())
}

func (i *SmartDetectionRule) ToSmartDetectionRulePtrOutputWithContext(ctx context.Context) SmartDetectionRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectionRulePtrOutput)
}

type SmartDetectionRulePtrInput interface {
	pulumi.Input

	ToSmartDetectionRulePtrOutput() SmartDetectionRulePtrOutput
	ToSmartDetectionRulePtrOutputWithContext(ctx context.Context) SmartDetectionRulePtrOutput
}

type smartDetectionRulePtrType SmartDetectionRuleArgs

func (*smartDetectionRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmartDetectionRule)(nil))
}

func (i *smartDetectionRulePtrType) ToSmartDetectionRulePtrOutput() SmartDetectionRulePtrOutput {
	return i.ToSmartDetectionRulePtrOutputWithContext(context.Background())
}

func (i *smartDetectionRulePtrType) ToSmartDetectionRulePtrOutputWithContext(ctx context.Context) SmartDetectionRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectionRulePtrOutput)
}

// SmartDetectionRuleArrayInput is an input type that accepts SmartDetectionRuleArray and SmartDetectionRuleArrayOutput values.
// You can construct a concrete instance of `SmartDetectionRuleArrayInput` via:
//
//          SmartDetectionRuleArray{ SmartDetectionRuleArgs{...} }
type SmartDetectionRuleArrayInput interface {
	pulumi.Input

	ToSmartDetectionRuleArrayOutput() SmartDetectionRuleArrayOutput
	ToSmartDetectionRuleArrayOutputWithContext(context.Context) SmartDetectionRuleArrayOutput
}

type SmartDetectionRuleArray []SmartDetectionRuleInput

func (SmartDetectionRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SmartDetectionRule)(nil))
}

func (i SmartDetectionRuleArray) ToSmartDetectionRuleArrayOutput() SmartDetectionRuleArrayOutput {
	return i.ToSmartDetectionRuleArrayOutputWithContext(context.Background())
}

func (i SmartDetectionRuleArray) ToSmartDetectionRuleArrayOutputWithContext(ctx context.Context) SmartDetectionRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectionRuleArrayOutput)
}

// SmartDetectionRuleMapInput is an input type that accepts SmartDetectionRuleMap and SmartDetectionRuleMapOutput values.
// You can construct a concrete instance of `SmartDetectionRuleMapInput` via:
//
//          SmartDetectionRuleMap{ "key": SmartDetectionRuleArgs{...} }
type SmartDetectionRuleMapInput interface {
	pulumi.Input

	ToSmartDetectionRuleMapOutput() SmartDetectionRuleMapOutput
	ToSmartDetectionRuleMapOutputWithContext(context.Context) SmartDetectionRuleMapOutput
}

type SmartDetectionRuleMap map[string]SmartDetectionRuleInput

func (SmartDetectionRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SmartDetectionRule)(nil))
}

func (i SmartDetectionRuleMap) ToSmartDetectionRuleMapOutput() SmartDetectionRuleMapOutput {
	return i.ToSmartDetectionRuleMapOutputWithContext(context.Background())
}

func (i SmartDetectionRuleMap) ToSmartDetectionRuleMapOutputWithContext(ctx context.Context) SmartDetectionRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectionRuleMapOutput)
}

type SmartDetectionRuleOutput struct {
	*pulumi.OutputState
}

func (SmartDetectionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmartDetectionRule)(nil))
}

func (o SmartDetectionRuleOutput) ToSmartDetectionRuleOutput() SmartDetectionRuleOutput {
	return o
}

func (o SmartDetectionRuleOutput) ToSmartDetectionRuleOutputWithContext(ctx context.Context) SmartDetectionRuleOutput {
	return o
}

func (o SmartDetectionRuleOutput) ToSmartDetectionRulePtrOutput() SmartDetectionRulePtrOutput {
	return o.ToSmartDetectionRulePtrOutputWithContext(context.Background())
}

func (o SmartDetectionRuleOutput) ToSmartDetectionRulePtrOutputWithContext(ctx context.Context) SmartDetectionRulePtrOutput {
	return o.ApplyT(func(v SmartDetectionRule) *SmartDetectionRule {
		return &v
	}).(SmartDetectionRulePtrOutput)
}

type SmartDetectionRulePtrOutput struct {
	*pulumi.OutputState
}

func (SmartDetectionRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmartDetectionRule)(nil))
}

func (o SmartDetectionRulePtrOutput) ToSmartDetectionRulePtrOutput() SmartDetectionRulePtrOutput {
	return o
}

func (o SmartDetectionRulePtrOutput) ToSmartDetectionRulePtrOutputWithContext(ctx context.Context) SmartDetectionRulePtrOutput {
	return o
}

type SmartDetectionRuleArrayOutput struct{ *pulumi.OutputState }

func (SmartDetectionRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmartDetectionRule)(nil))
}

func (o SmartDetectionRuleArrayOutput) ToSmartDetectionRuleArrayOutput() SmartDetectionRuleArrayOutput {
	return o
}

func (o SmartDetectionRuleArrayOutput) ToSmartDetectionRuleArrayOutputWithContext(ctx context.Context) SmartDetectionRuleArrayOutput {
	return o
}

func (o SmartDetectionRuleArrayOutput) Index(i pulumi.IntInput) SmartDetectionRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SmartDetectionRule {
		return vs[0].([]SmartDetectionRule)[vs[1].(int)]
	}).(SmartDetectionRuleOutput)
}

type SmartDetectionRuleMapOutput struct{ *pulumi.OutputState }

func (SmartDetectionRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SmartDetectionRule)(nil))
}

func (o SmartDetectionRuleMapOutput) ToSmartDetectionRuleMapOutput() SmartDetectionRuleMapOutput {
	return o
}

func (o SmartDetectionRuleMapOutput) ToSmartDetectionRuleMapOutputWithContext(ctx context.Context) SmartDetectionRuleMapOutput {
	return o
}

func (o SmartDetectionRuleMapOutput) MapIndex(k pulumi.StringInput) SmartDetectionRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SmartDetectionRule {
		return vs[0].(map[string]SmartDetectionRule)[vs[1].(string)]
	}).(SmartDetectionRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(SmartDetectionRuleOutput{})
	pulumi.RegisterOutputType(SmartDetectionRulePtrOutput{})
	pulumi.RegisterOutputType(SmartDetectionRuleArrayOutput{})
	pulumi.RegisterOutputType(SmartDetectionRuleMapOutput{})
}
