// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Bot Web App.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/bot"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = bot.NewWebApp(ctx, "exampleWebApp", &bot.WebAppArgs{
// 			Location:          pulumi.String("global"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("F0"),
// 			MicrosoftAppId:    pulumi.String(current.ClientId),
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
// Bot Web App's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:bot/webApp:WebApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
// ```
type WebApp struct {
	pulumi.CustomResourceState

	// The Application Insights API Key to associate with the Web App Bot.
	DeveloperAppInsightsApiKey pulumi.StringOutput `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationId pulumi.StringOutput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey pulumi.StringOutput `pulumi:"developerAppInsightsKey"`
	// The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Web App Bot endpoint.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds pulumi.StringArrayOutput `pulumi:"luisAppIds"`
	// The LUIS key to associate with the Web App Bot.
	LuisKey pulumi.StringPtrOutput `pulumi:"luisKey"`
	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringOutput `pulumi:"microsoftAppId"`
	// Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWebApp registers a new resource with the given unique name, arguments, and options.
func NewWebApp(ctx *pulumi.Context,
	name string, args *WebAppArgs, opts ...pulumi.ResourceOption) (*WebApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MicrosoftAppId == nil {
		return nil, errors.New("invalid value for required argument 'MicrosoftAppId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource WebApp
	err := ctx.RegisterResource("azure:bot/webApp:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebApp gets an existing WebApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("azure:bot/webApp:WebApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebApp resources.
type webAppState struct {
	// The Application Insights API Key to associate with the Web App Bot.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Web App Bot endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds []string `pulumi:"luisAppIds"`
	// The LUIS key to associate with the Web App Bot.
	LuisKey *string `pulumi:"luisKey"`
	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppId *string `pulumi:"microsoftAppId"`
	// Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type WebAppState struct {
	// The Application Insights API Key to associate with the Web App Bot.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Web App Bot endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds pulumi.StringArrayInput
	// The LUIS key to associate with the Web App Bot.
	LuisKey pulumi.StringPtrInput
	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringPtrInput
	// Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WebAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppState)(nil)).Elem()
}

type webAppArgs struct {
	// The Application Insights API Key to associate with the Web App Bot.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Web App Bot endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds []string `pulumi:"luisAppIds"`
	// The LUIS key to associate with the Web App Bot.
	LuisKey *string `pulumi:"luisKey"`
	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppId string `pulumi:"microsoftAppId"`
	// Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebApp resource.
type WebAppArgs struct {
	// The Application Insights API Key to associate with the Web App Bot.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Web App Bot will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Web App Bot endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds pulumi.StringArrayInput
	// The LUIS key to associate with the Web App Bot.
	LuisKey pulumi.StringPtrInput
	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput
	// Specifies the name of the Web App Bot. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Web App Bot. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WebAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppArgs)(nil)).Elem()
}

type WebAppInput interface {
	pulumi.Input

	ToWebAppOutput() WebAppOutput
	ToWebAppOutputWithContext(ctx context.Context) WebAppOutput
}

func (*WebApp) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApp)(nil))
}

func (i *WebApp) ToWebAppOutput() WebAppOutput {
	return i.ToWebAppOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppOutput)
}

func (i *WebApp) ToWebAppPtrOutput() WebAppPtrOutput {
	return i.ToWebAppPtrOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPtrOutput)
}

type WebAppPtrInput interface {
	pulumi.Input

	ToWebAppPtrOutput() WebAppPtrOutput
	ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput
}

type webAppPtrType WebAppArgs

func (*webAppPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil))
}

func (i *webAppPtrType) ToWebAppPtrOutput() WebAppPtrOutput {
	return i.ToWebAppPtrOutputWithContext(context.Background())
}

func (i *webAppPtrType) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPtrOutput)
}

// WebAppArrayInput is an input type that accepts WebAppArray and WebAppArrayOutput values.
// You can construct a concrete instance of `WebAppArrayInput` via:
//
//          WebAppArray{ WebAppArgs{...} }
type WebAppArrayInput interface {
	pulumi.Input

	ToWebAppArrayOutput() WebAppArrayOutput
	ToWebAppArrayOutputWithContext(context.Context) WebAppArrayOutput
}

type WebAppArray []WebAppInput

func (WebAppArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebApp)(nil))
}

func (i WebAppArray) ToWebAppArrayOutput() WebAppArrayOutput {
	return i.ToWebAppArrayOutputWithContext(context.Background())
}

func (i WebAppArray) ToWebAppArrayOutputWithContext(ctx context.Context) WebAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppArrayOutput)
}

// WebAppMapInput is an input type that accepts WebAppMap and WebAppMapOutput values.
// You can construct a concrete instance of `WebAppMapInput` via:
//
//          WebAppMap{ "key": WebAppArgs{...} }
type WebAppMapInput interface {
	pulumi.Input

	ToWebAppMapOutput() WebAppMapOutput
	ToWebAppMapOutputWithContext(context.Context) WebAppMapOutput
}

type WebAppMap map[string]WebAppInput

func (WebAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebApp)(nil))
}

func (i WebAppMap) ToWebAppMapOutput() WebAppMapOutput {
	return i.ToWebAppMapOutputWithContext(context.Background())
}

func (i WebAppMap) ToWebAppMapOutputWithContext(ctx context.Context) WebAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMapOutput)
}

type WebAppOutput struct {
	*pulumi.OutputState
}

func (WebAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApp)(nil))
}

func (o WebAppOutput) ToWebAppOutput() WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppPtrOutput() WebAppPtrOutput {
	return o.ToWebAppPtrOutputWithContext(context.Background())
}

func (o WebAppOutput) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return o.ApplyT(func(v WebApp) *WebApp {
		return &v
	}).(WebAppPtrOutput)
}

type WebAppPtrOutput struct {
	*pulumi.OutputState
}

func (WebAppPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil))
}

func (o WebAppPtrOutput) ToWebAppPtrOutput() WebAppPtrOutput {
	return o
}

func (o WebAppPtrOutput) ToWebAppPtrOutputWithContext(ctx context.Context) WebAppPtrOutput {
	return o
}

type WebAppArrayOutput struct{ *pulumi.OutputState }

func (WebAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApp)(nil))
}

func (o WebAppArrayOutput) ToWebAppArrayOutput() WebAppArrayOutput {
	return o
}

func (o WebAppArrayOutput) ToWebAppArrayOutputWithContext(ctx context.Context) WebAppArrayOutput {
	return o
}

func (o WebAppArrayOutput) Index(i pulumi.IntInput) WebAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApp {
		return vs[0].([]WebApp)[vs[1].(int)]
	}).(WebAppOutput)
}

type WebAppMapOutput struct{ *pulumi.OutputState }

func (WebAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebApp)(nil))
}

func (o WebAppMapOutput) ToWebAppMapOutput() WebAppMapOutput {
	return o
}

func (o WebAppMapOutput) ToWebAppMapOutputWithContext(ctx context.Context) WebAppMapOutput {
	return o
}

func (o WebAppMapOutput) MapIndex(k pulumi.StringInput) WebAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebApp {
		return vs[0].(map[string]WebApp)[vs[1].(string)]
	}).(WebAppOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
	pulumi.RegisterOutputType(WebAppPtrOutput{})
	pulumi.RegisterOutputType(WebAppArrayOutput{})
	pulumi.RegisterOutputType(WebAppMapOutput{})
}
