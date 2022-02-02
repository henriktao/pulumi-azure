// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Action Group within Azure Monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
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
// 		_, err = monitoring.NewActionGroup(ctx, "exampleActionGroup", &monitoring.ActionGroupArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ShortName:         pulumi.String("p0action"),
// 			ArmRoleReceivers: monitoring.ActionGroupArmRoleReceiverArray{
// 				&monitoring.ActionGroupArmRoleReceiverArgs{
// 					Name:                 pulumi.String("armroleaction"),
// 					RoleId:               pulumi.String("de139f84-1756-47ae-9be6-808fbbe84772"),
// 					UseCommonAlertSchema: pulumi.Bool(true),
// 				},
// 			},
// 			AutomationRunbookReceivers: monitoring.ActionGroupAutomationRunbookReceiverArray{
// 				&monitoring.ActionGroupAutomationRunbookReceiverArgs{
// 					Name:                 pulumi.String("action_name_1"),
// 					AutomationAccountId:  pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001"),
// 					RunbookName:          pulumi.String("my runbook"),
// 					WebhookResourceId:    pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-runbooks/providers/microsoft.automation/automationaccounts/aaa001/webhooks/webhook_alert"),
// 					IsGlobalRunbook:      pulumi.Bool(true),
// 					ServiceUri:           pulumi.String("https://s13events.azure-automation.net/webhooks?token=randomtoken"),
// 					UseCommonAlertSchema: pulumi.Bool(true),
// 				},
// 			},
// 			AzureAppPushReceivers: monitoring.ActionGroupAzureAppPushReceiverArray{
// 				&monitoring.ActionGroupAzureAppPushReceiverArgs{
// 					Name:         pulumi.String("pushtoadmin"),
// 					EmailAddress: pulumi.String("admin@contoso.com"),
// 				},
// 			},
// 			AzureFunctionReceivers: monitoring.ActionGroupAzureFunctionReceiverArray{
// 				&monitoring.ActionGroupAzureFunctionReceiverArgs{
// 					Name:                  pulumi.String("funcaction"),
// 					FunctionAppResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-funcapp/providers/Microsoft.Web/sites/funcapp"),
// 					FunctionName:          pulumi.String("myfunc"),
// 					HttpTriggerUrl:        pulumi.String("https://example.com/trigger"),
// 					UseCommonAlertSchema:  pulumi.Bool(true),
// 				},
// 			},
// 			EmailReceivers: monitoring.ActionGroupEmailReceiverArray{
// 				&monitoring.ActionGroupEmailReceiverArgs{
// 					Name:         pulumi.String("sendtoadmin"),
// 					EmailAddress: pulumi.String("admin@contoso.com"),
// 				},
// 				&monitoring.ActionGroupEmailReceiverArgs{
// 					Name:                 pulumi.String("sendtodevops"),
// 					EmailAddress:         pulumi.String("devops@contoso.com"),
// 					UseCommonAlertSchema: pulumi.Bool(true),
// 				},
// 			},
// 			EventHubReceivers: monitoring.ActionGroupEventHubReceiverArray{
// 				&monitoring.ActionGroupEventHubReceiverArgs{
// 					Name:                 pulumi.String("sendtoeventhub"),
// 					EventHubId:           pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-eventhub/providers/Microsoft.EventHub/namespaces/eventhubnamespace/eventhubs/eventhub1"),
// 					UseCommonAlertSchema: pulumi.Bool(false),
// 				},
// 			},
// 			ItsmReceivers: monitoring.ActionGroupItsmReceiverArray{
// 				&monitoring.ActionGroupItsmReceiverArgs{
// 					Name:                pulumi.String("createorupdateticket"),
// 					WorkspaceId:         pulumi.String("6eee3a18-aac3-40e4-b98e-1f309f329816"),
// 					ConnectionId:        pulumi.String("53de6956-42b4-41ba-be3c-b154cdf17b13"),
// 					TicketConfiguration: pulumi.String("{}"),
// 					Region:              pulumi.String("southcentralus"),
// 				},
// 			},
// 			LogicAppReceivers: monitoring.ActionGroupLogicAppReceiverArray{
// 				&monitoring.ActionGroupLogicAppReceiverArgs{
// 					Name:                 pulumi.String("logicappaction"),
// 					ResourceId:           pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-logicapp/providers/Microsoft.Logic/workflows/logicapp"),
// 					CallbackUrl:          pulumi.String("https://logicapptriggerurl/..."),
// 					UseCommonAlertSchema: pulumi.Bool(true),
// 				},
// 			},
// 			SmsReceivers: monitoring.ActionGroupSmsReceiverArray{
// 				&monitoring.ActionGroupSmsReceiverArgs{
// 					Name:        pulumi.String("oncallmsg"),
// 					CountryCode: pulumi.String("1"),
// 					PhoneNumber: pulumi.String("1231231234"),
// 				},
// 			},
// 			VoiceReceivers: monitoring.ActionGroupVoiceReceiverArray{
// 				&monitoring.ActionGroupVoiceReceiverArgs{
// 					Name:        pulumi.String("remotesupport"),
// 					CountryCode: pulumi.String("86"),
// 					PhoneNumber: pulumi.String("13888888888"),
// 				},
// 			},
// 			WebhookReceivers: monitoring.ActionGroupWebhookReceiverArray{
// 				&monitoring.ActionGroupWebhookReceiverArgs{
// 					Name:                 pulumi.String("callmyapiaswell"),
// 					ServiceUri:           pulumi.String("http://example.com/alert"),
// 					UseCommonAlertSchema: pulumi.Bool(true),
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
// Action Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/actionGroup:ActionGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/actionGroups/myagname
// ```
type ActionGroup struct {
	pulumi.CustomResourceState

	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers ActionGroupArmRoleReceiverArrayOutput `pulumi:"armRoleReceivers"`
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers ActionGroupAutomationRunbookReceiverArrayOutput `pulumi:"automationRunbookReceivers"`
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers ActionGroupAzureAppPushReceiverArrayOutput `pulumi:"azureAppPushReceivers"`
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers ActionGroupAzureFunctionReceiverArrayOutput `pulumi:"azureFunctionReceivers"`
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers ActionGroupEmailReceiverArrayOutput `pulumi:"emailReceivers"`
	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// One or more `eventHubReceiver` blocks as defined below.
	EventHubReceivers ActionGroupEventHubReceiverArrayOutput `pulumi:"eventHubReceivers"`
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers ActionGroupItsmReceiverArrayOutput `pulumi:"itsmReceivers"`
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers ActionGroupLogicAppReceiverArrayOutput `pulumi:"logicAppReceivers"`
	// The name of the Action Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Action Group instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The short name of the action group. This will be used in SMS messages.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers ActionGroupSmsReceiverArrayOutput `pulumi:"smsReceivers"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers ActionGroupVoiceReceiverArrayOutput `pulumi:"voiceReceivers"`
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers ActionGroupWebhookReceiverArrayOutput `pulumi:"webhookReceivers"`
}

// NewActionGroup registers a new resource with the given unique name, arguments, and options.
func NewActionGroup(ctx *pulumi.Context,
	name string, args *ActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShortName == nil {
		return nil, errors.New("invalid value for required argument 'ShortName'")
	}
	var resource ActionGroup
	err := ctx.RegisterResource("azure:monitoring/actionGroup:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionGroup gets an existing ActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azure:monitoring/actionGroup:ActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionGroup resources.
type actionGroupState struct {
	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers []ActionGroupArmRoleReceiver `pulumi:"armRoleReceivers"`
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers []ActionGroupAutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers []ActionGroupAzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers []ActionGroupAzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers []ActionGroupEmailReceiver `pulumi:"emailReceivers"`
	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// One or more `eventHubReceiver` blocks as defined below.
	EventHubReceivers []ActionGroupEventHubReceiver `pulumi:"eventHubReceivers"`
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers []ActionGroupItsmReceiver `pulumi:"itsmReceivers"`
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers []ActionGroupLogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the Action Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Action Group instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The short name of the action group. This will be used in SMS messages.
	ShortName *string `pulumi:"shortName"`
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers []ActionGroupSmsReceiver `pulumi:"smsReceivers"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers []ActionGroupVoiceReceiver `pulumi:"voiceReceivers"`
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers []ActionGroupWebhookReceiver `pulumi:"webhookReceivers"`
}

type ActionGroupState struct {
	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers ActionGroupArmRoleReceiverArrayInput
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers ActionGroupAutomationRunbookReceiverArrayInput
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers ActionGroupAzureAppPushReceiverArrayInput
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers ActionGroupAzureFunctionReceiverArrayInput
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers ActionGroupEmailReceiverArrayInput
	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// One or more `eventHubReceiver` blocks as defined below.
	EventHubReceivers ActionGroupEventHubReceiverArrayInput
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers ActionGroupItsmReceiverArrayInput
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers ActionGroupLogicAppReceiverArrayInput
	// The name of the Action Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Action Group instance.
	ResourceGroupName pulumi.StringPtrInput
	// The short name of the action group. This will be used in SMS messages.
	ShortName pulumi.StringPtrInput
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers ActionGroupSmsReceiverArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers ActionGroupVoiceReceiverArrayInput
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers ActionGroupWebhookReceiverArrayInput
}

func (ActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupState)(nil)).Elem()
}

type actionGroupArgs struct {
	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers []ActionGroupArmRoleReceiver `pulumi:"armRoleReceivers"`
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers []ActionGroupAutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers []ActionGroupAzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers []ActionGroupAzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers []ActionGroupEmailReceiver `pulumi:"emailReceivers"`
	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// One or more `eventHubReceiver` blocks as defined below.
	EventHubReceivers []ActionGroupEventHubReceiver `pulumi:"eventHubReceivers"`
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers []ActionGroupItsmReceiver `pulumi:"itsmReceivers"`
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers []ActionGroupLogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the Action Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Action Group instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The short name of the action group. This will be used in SMS messages.
	ShortName string `pulumi:"shortName"`
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers []ActionGroupSmsReceiver `pulumi:"smsReceivers"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers []ActionGroupVoiceReceiver `pulumi:"voiceReceivers"`
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers []ActionGroupWebhookReceiver `pulumi:"webhookReceivers"`
}

// The set of arguments for constructing a ActionGroup resource.
type ActionGroupArgs struct {
	// One or more `armRoleReceiver` blocks as defined below.
	ArmRoleReceivers ActionGroupArmRoleReceiverArrayInput
	// One or more `automationRunbookReceiver` blocks as defined below.
	AutomationRunbookReceivers ActionGroupAutomationRunbookReceiverArrayInput
	// One or more `azureAppPushReceiver` blocks as defined below.
	AzureAppPushReceivers ActionGroupAzureAppPushReceiverArrayInput
	// One or more `azureFunctionReceiver` blocks as defined below.
	AzureFunctionReceivers ActionGroupAzureFunctionReceiverArrayInput
	// One or more `emailReceiver` blocks as defined below.
	EmailReceivers ActionGroupEmailReceiverArrayInput
	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// One or more `eventHubReceiver` blocks as defined below.
	EventHubReceivers ActionGroupEventHubReceiverArrayInput
	// One or more `itsmReceiver` blocks as defined below.
	ItsmReceivers ActionGroupItsmReceiverArrayInput
	// One or more `logicAppReceiver` blocks as defined below.
	LogicAppReceivers ActionGroupLogicAppReceiverArrayInput
	// The name of the Action Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Action Group instance.
	ResourceGroupName pulumi.StringInput
	// The short name of the action group. This will be used in SMS messages.
	ShortName pulumi.StringInput
	// One or more `smsReceiver` blocks as defined below.
	SmsReceivers ActionGroupSmsReceiverArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more `voiceReceiver` blocks as defined below.
	VoiceReceivers ActionGroupVoiceReceiverArrayInput
	// One or more `webhookReceiver` blocks as defined below.
	WebhookReceivers ActionGroupWebhookReceiverArrayInput
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupArgs)(nil)).Elem()
}

type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput
}

func (*ActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (i *ActionGroup) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i *ActionGroup) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

// ActionGroupArrayInput is an input type that accepts ActionGroupArray and ActionGroupArrayOutput values.
// You can construct a concrete instance of `ActionGroupArrayInput` via:
//
//          ActionGroupArray{ ActionGroupArgs{...} }
type ActionGroupArrayInput interface {
	pulumi.Input

	ToActionGroupArrayOutput() ActionGroupArrayOutput
	ToActionGroupArrayOutputWithContext(context.Context) ActionGroupArrayOutput
}

type ActionGroupArray []ActionGroupInput

func (ActionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionGroup)(nil)).Elem()
}

func (i ActionGroupArray) ToActionGroupArrayOutput() ActionGroupArrayOutput {
	return i.ToActionGroupArrayOutputWithContext(context.Background())
}

func (i ActionGroupArray) ToActionGroupArrayOutputWithContext(ctx context.Context) ActionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupArrayOutput)
}

// ActionGroupMapInput is an input type that accepts ActionGroupMap and ActionGroupMapOutput values.
// You can construct a concrete instance of `ActionGroupMapInput` via:
//
//          ActionGroupMap{ "key": ActionGroupArgs{...} }
type ActionGroupMapInput interface {
	pulumi.Input

	ToActionGroupMapOutput() ActionGroupMapOutput
	ToActionGroupMapOutputWithContext(context.Context) ActionGroupMapOutput
}

type ActionGroupMap map[string]ActionGroupInput

func (ActionGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionGroup)(nil)).Elem()
}

func (i ActionGroupMap) ToActionGroupMapOutput() ActionGroupMapOutput {
	return i.ToActionGroupMapOutputWithContext(context.Background())
}

func (i ActionGroupMap) ToActionGroupMapOutputWithContext(ctx context.Context) ActionGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupMapOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroup)(nil)).Elem()
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

type ActionGroupArrayOutput struct{ *pulumi.OutputState }

func (ActionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionGroup)(nil)).Elem()
}

func (o ActionGroupArrayOutput) ToActionGroupArrayOutput() ActionGroupArrayOutput {
	return o
}

func (o ActionGroupArrayOutput) ToActionGroupArrayOutputWithContext(ctx context.Context) ActionGroupArrayOutput {
	return o
}

func (o ActionGroupArrayOutput) Index(i pulumi.IntInput) ActionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionGroup {
		return vs[0].([]*ActionGroup)[vs[1].(int)]
	}).(ActionGroupOutput)
}

type ActionGroupMapOutput struct{ *pulumi.OutputState }

func (ActionGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionGroup)(nil)).Elem()
}

func (o ActionGroupMapOutput) ToActionGroupMapOutput() ActionGroupMapOutput {
	return o
}

func (o ActionGroupMapOutput) ToActionGroupMapOutputWithContext(ctx context.Context) ActionGroupMapOutput {
	return o
}

func (o ActionGroupMapOutput) MapIndex(k pulumi.StringInput) ActionGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionGroup {
		return vs[0].(map[string]*ActionGroup)[vs[1].(string)]
	}).(ActionGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupInput)(nil)).Elem(), &ActionGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupArrayInput)(nil)).Elem(), ActionGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupMapInput)(nil)).Elem(), ActionGroupMap{})
	pulumi.RegisterOutputType(ActionGroupOutput{})
	pulumi.RegisterOutputType(ActionGroupArrayOutput{})
	pulumi.RegisterOutputType(ActionGroupMapOutput{})
}
