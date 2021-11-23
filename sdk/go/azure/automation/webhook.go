// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Automation Runbook's Webhook.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/automation"
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
// 		exampleAccount, err := automation.NewAccount(ctx, "exampleAccount", &automation.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Basic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleRunBook, err := automation.NewRunBook(ctx, "exampleRunBook", &automation.RunBookArgs{
// 			Location:              exampleResourceGroup.Location,
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			AutomationAccountName: exampleAccount.Name,
// 			LogVerbose:            pulumi.Bool(true),
// 			LogProgress:           pulumi.Bool(true),
// 			Description:           pulumi.String("This is an example runbook"),
// 			RunbookType:           pulumi.String("PowerShellWorkflow"),
// 			PublishContentLink: &automation.RunBookPublishContentLinkArgs{
// 				Uri: pulumi.String("https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = automation.NewWebhook(ctx, "exampleWebhook", &automation.WebhookArgs{
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			AutomationAccountName: exampleAccount.Name,
// 			ExpiryTime:            pulumi.String("2021-12-31T00:00:00Z"),
// 			Enabled:               pulumi.Bool(true),
// 			RunbookName:           exampleRunBook.Name,
// 			Parameters: pulumi.StringMap{
// 				"input": pulumi.String("parameter"),
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
// Automation Webhooks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/webhook:Webhook TestRunbook_webhook /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webhooks/TestRunbook_webhook
// ```
type Webhook struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// Controls if Webhook is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Map of input parameters passed to runbook.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup pulumi.StringPtrOutput `pulumi:"runOnWorkerGroup"`
	// Name of the Automation Runbook to execute by Webhook.
	RunbookName pulumi.StringOutput `pulumi:"runbookName"`
	// URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ExpiryTime == nil {
		return nil, errors.New("invalid value for required argument 'ExpiryTime'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookName == nil {
		return nil, errors.New("invalid value for required argument 'RunbookName'")
	}
	var resource Webhook
	err := ctx.RegisterResource("azure:automation/webhook:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azure:automation/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// Controls if Webhook is enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime *string `pulumi:"expiryTime"`
	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Map of input parameters passed to runbook.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup *string `pulumi:"runOnWorkerGroup"`
	// Name of the Automation Runbook to execute by Webhook.
	RunbookName *string `pulumi:"runbookName"`
	// URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
	Uri *string `pulumi:"uri"`
}

type WebhookState struct {
	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// Controls if Webhook is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime pulumi.StringPtrInput
	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Map of input parameters passed to runbook.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup pulumi.StringPtrInput
	// Name of the Automation Runbook to execute by Webhook.
	RunbookName pulumi.StringPtrInput
	// URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
	Uri pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Controls if Webhook is enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime string `pulumi:"expiryTime"`
	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Map of input parameters passed to runbook.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup *string `pulumi:"runOnWorkerGroup"`
	// Name of the Automation Runbook to execute by Webhook.
	RunbookName string `pulumi:"runbookName"`
	// URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// Controls if Webhook is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime pulumi.StringInput
	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Map of input parameters passed to runbook.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup pulumi.StringPtrInput
	// Name of the Automation Runbook to execute by Webhook.
	RunbookName pulumi.StringInput
	// URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
	Uri pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

func (i *Webhook) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

type WebhookPtrInput interface {
	pulumi.Input

	ToWebhookPtrOutput() WebhookPtrOutput
	ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput
}

type webhookPtrType WebhookArgs

func (*webhookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil))
}

func (i *webhookPtrType) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *webhookPtrType) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

// WebhookArrayInput is an input type that accepts WebhookArray and WebhookArrayOutput values.
// You can construct a concrete instance of `WebhookArrayInput` via:
//
//          WebhookArray{ WebhookArgs{...} }
type WebhookArrayInput interface {
	pulumi.Input

	ToWebhookArrayOutput() WebhookArrayOutput
	ToWebhookArrayOutputWithContext(context.Context) WebhookArrayOutput
}

type WebhookArray []WebhookInput

func (WebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhook)(nil)).Elem()
}

func (i WebhookArray) ToWebhookArrayOutput() WebhookArrayOutput {
	return i.ToWebhookArrayOutputWithContext(context.Background())
}

func (i WebhookArray) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookArrayOutput)
}

// WebhookMapInput is an input type that accepts WebhookMap and WebhookMapOutput values.
// You can construct a concrete instance of `WebhookMapInput` via:
//
//          WebhookMap{ "key": WebhookArgs{...} }
type WebhookMapInput interface {
	pulumi.Input

	ToWebhookMapOutput() WebhookMapOutput
	ToWebhookMapOutputWithContext(context.Context) WebhookMapOutput
}

type WebhookMap map[string]WebhookInput

func (WebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhook)(nil)).Elem()
}

func (i WebhookMap) ToWebhookMapOutput() WebhookMapOutput {
	return i.ToWebhookMapOutputWithContext(context.Background())
}

func (i WebhookMap) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookMapOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o.ToWebhookPtrOutputWithContext(context.Background())
}

func (o WebhookOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Webhook) *Webhook {
		return &v
	}).(WebhookPtrOutput)
}

type WebhookPtrOutput struct{ *pulumi.OutputState }

func (WebhookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil))
}

func (o WebhookPtrOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) Elem() WebhookOutput {
	return o.ApplyT(func(v *Webhook) Webhook {
		if v != nil {
			return *v
		}
		var ret Webhook
		return ret
	}).(WebhookOutput)
}

type WebhookArrayOutput struct{ *pulumi.OutputState }

func (WebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Webhook)(nil))
}

func (o WebhookArrayOutput) ToWebhookArrayOutput() WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) ToWebhookArrayOutputWithContext(ctx context.Context) WebhookArrayOutput {
	return o
}

func (o WebhookArrayOutput) Index(i pulumi.IntInput) WebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Webhook {
		return vs[0].([]Webhook)[vs[1].(int)]
	}).(WebhookOutput)
}

type WebhookMapOutput struct{ *pulumi.OutputState }

func (WebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Webhook)(nil))
}

func (o WebhookMapOutput) ToWebhookMapOutput() WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) ToWebhookMapOutputWithContext(ctx context.Context) WebhookMapOutput {
	return o
}

func (o WebhookMapOutput) MapIndex(k pulumi.StringInput) WebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Webhook {
		return vs[0].(map[string]Webhook)[vs[1].(string)]
	}).(WebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookInput)(nil)).Elem(), &Webhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookPtrInput)(nil)).Elem(), &Webhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookArrayInput)(nil)).Elem(), WebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookMapInput)(nil)).Elem(), WebhookMap{})
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookPtrOutput{})
	pulumi.RegisterOutputType(WebhookArrayOutput{})
	pulumi.RegisterOutputType(WebhookMapOutput{})
}
