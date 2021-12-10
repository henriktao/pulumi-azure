// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Bot Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/bot"
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
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleApiKey, err := appinsights.NewApiKey(ctx, "exampleApiKey", &appinsights.ApiKeyArgs{
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			ReadPermissions: pulumi.StringArray{
// 				pulumi.String("aggregate"),
// 				pulumi.String("api"),
// 				pulumi.String("draft"),
// 				pulumi.String("extendqueries"),
// 				pulumi.String("search"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bot.NewServiceAzureBot(ctx, "exampleServiceAzureBot", &bot.ServiceAzureBotArgs{
// 			ResourceGroupName:                 exampleResourceGroup.Name,
// 			Location:                          pulumi.String("global"),
// 			MicrosoftAppId:                    pulumi.Any(data.Azurerm_client_config.Current.Client_id),
// 			Sku:                               pulumi.String("F0"),
// 			Endpoint:                          pulumi.String("https://example.com"),
// 			DeveloperAppInsightsApiKey:        exampleApiKey.ApiKey,
// 			DeveloperAppInsightsApplicationId: exampleInsights.AppId,
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("test"),
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
// Azure Bot Services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:bot/serviceAzureBot:ServiceAzureBot example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.BotService/botServices/botService1
// ```
type ServiceAzureBot struct {
	pulumi.CustomResourceState

	// The Application Insights Api Key to associate with this Azure Bot Service.
	DeveloperAppInsightsApiKey pulumi.StringPtrOutput `pulumi:"developerAppInsightsApiKey"`
	// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
	DeveloperAppInsightsApplicationId pulumi.StringPtrOutput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insight Key to associate with this Azure Bot Service.
	DeveloperAppInsightsKey pulumi.StringPtrOutput `pulumi:"developerAppInsightsKey"`
	// The name that the Azure Bot Service will be displayed as. This defaults to the value set for `name` if not specified.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Azure Bot Service endpoint.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of LUIS App IDs to associate with this Azure Bot Service.
	LuisAppIds pulumi.StringArrayOutput `pulumi:"luisAppIds"`
	// The LUIS key to associate with this Azure Bot Service.
	LuisKey pulumi.StringPtrOutput `pulumi:"luisKey"`
	// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringOutput `pulumi:"microsoftAppId"`
	// The name which should be used for this Azure Bot Service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Azure Bot Service. Accepted values are `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags which should be assigned to this Azure Bot Service.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewServiceAzureBot registers a new resource with the given unique name, arguments, and options.
func NewServiceAzureBot(ctx *pulumi.Context,
	name string, args *ServiceAzureBotArgs, opts ...pulumi.ResourceOption) (*ServiceAzureBot, error) {
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
	var resource ServiceAzureBot
	err := ctx.RegisterResource("azure:bot/serviceAzureBot:ServiceAzureBot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAzureBot gets an existing ServiceAzureBot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAzureBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAzureBotState, opts ...pulumi.ResourceOption) (*ServiceAzureBot, error) {
	var resource ServiceAzureBot
	err := ctx.ReadResource("azure:bot/serviceAzureBot:ServiceAzureBot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAzureBot resources.
type serviceAzureBotState struct {
	// The Application Insights Api Key to associate with this Azure Bot Service.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insight Key to associate with this Azure Bot Service.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name that the Azure Bot Service will be displayed as. This defaults to the value set for `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Azure Bot Service endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A list of LUIS App IDs to associate with this Azure Bot Service.
	LuisAppIds []string `pulumi:"luisAppIds"`
	// The LUIS key to associate with this Azure Bot Service.
	LuisKey *string `pulumi:"luisKey"`
	// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
	MicrosoftAppId *string `pulumi:"microsoftAppId"`
	// The name which should be used for this Azure Bot Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Azure Bot Service. Accepted values are `F0` or `S1`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags which should be assigned to this Azure Bot Service.
	Tags map[string]string `pulumi:"tags"`
}

type ServiceAzureBotState struct {
	// The Application Insights Api Key to associate with this Azure Bot Service.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insight Key to associate with this Azure Bot Service.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name that the Azure Bot Service will be displayed as. This defaults to the value set for `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Azure Bot Service endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A list of LUIS App IDs to associate with this Azure Bot Service.
	LuisAppIds pulumi.StringArrayInput
	// The LUIS key to associate with this Azure Bot Service.
	LuisKey pulumi.StringPtrInput
	// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringPtrInput
	// The name which should be used for this Azure Bot Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Azure Bot Service. Accepted values are `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags which should be assigned to this Azure Bot Service.
	Tags pulumi.StringMapInput
}

func (ServiceAzureBotState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAzureBotState)(nil)).Elem()
}

type serviceAzureBotArgs struct {
	// The Application Insights Api Key to associate with this Azure Bot Service.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insight Key to associate with this Azure Bot Service.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name that the Azure Bot Service will be displayed as. This defaults to the value set for `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Azure Bot Service endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A list of LUIS App IDs to associate with this Azure Bot Service.
	LuisAppIds []string `pulumi:"luisAppIds"`
	// The LUIS key to associate with this Azure Bot Service.
	LuisKey *string `pulumi:"luisKey"`
	// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
	MicrosoftAppId string `pulumi:"microsoftAppId"`
	// The name which should be used for this Azure Bot Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Azure Bot Service. Accepted values are `F0` or `S1`. Changing this forces a new resource to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags which should be assigned to this Azure Bot Service.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceAzureBot resource.
type ServiceAzureBotArgs struct {
	// The Application Insights Api Key to associate with this Azure Bot Service.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insight Key to associate with this Azure Bot Service.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name that the Azure Bot Service will be displayed as. This defaults to the value set for `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Azure Bot Service endpoint.
	Endpoint pulumi.StringPtrInput
	// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A list of LUIS App IDs to associate with this Azure Bot Service.
	LuisAppIds pulumi.StringArrayInput
	// The LUIS key to associate with this Azure Bot Service.
	LuisKey pulumi.StringPtrInput
	// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput
	// The name which should be used for this Azure Bot Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Azure Bot Service. Accepted values are `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput
	// A mapping of tags which should be assigned to this Azure Bot Service.
	Tags pulumi.StringMapInput
}

func (ServiceAzureBotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAzureBotArgs)(nil)).Elem()
}

type ServiceAzureBotInput interface {
	pulumi.Input

	ToServiceAzureBotOutput() ServiceAzureBotOutput
	ToServiceAzureBotOutputWithContext(ctx context.Context) ServiceAzureBotOutput
}

func (*ServiceAzureBot) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAzureBot)(nil))
}

func (i *ServiceAzureBot) ToServiceAzureBotOutput() ServiceAzureBotOutput {
	return i.ToServiceAzureBotOutputWithContext(context.Background())
}

func (i *ServiceAzureBot) ToServiceAzureBotOutputWithContext(ctx context.Context) ServiceAzureBotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAzureBotOutput)
}

func (i *ServiceAzureBot) ToServiceAzureBotPtrOutput() ServiceAzureBotPtrOutput {
	return i.ToServiceAzureBotPtrOutputWithContext(context.Background())
}

func (i *ServiceAzureBot) ToServiceAzureBotPtrOutputWithContext(ctx context.Context) ServiceAzureBotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAzureBotPtrOutput)
}

type ServiceAzureBotPtrInput interface {
	pulumi.Input

	ToServiceAzureBotPtrOutput() ServiceAzureBotPtrOutput
	ToServiceAzureBotPtrOutputWithContext(ctx context.Context) ServiceAzureBotPtrOutput
}

type serviceAzureBotPtrType ServiceAzureBotArgs

func (*serviceAzureBotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAzureBot)(nil))
}

func (i *serviceAzureBotPtrType) ToServiceAzureBotPtrOutput() ServiceAzureBotPtrOutput {
	return i.ToServiceAzureBotPtrOutputWithContext(context.Background())
}

func (i *serviceAzureBotPtrType) ToServiceAzureBotPtrOutputWithContext(ctx context.Context) ServiceAzureBotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAzureBotPtrOutput)
}

// ServiceAzureBotArrayInput is an input type that accepts ServiceAzureBotArray and ServiceAzureBotArrayOutput values.
// You can construct a concrete instance of `ServiceAzureBotArrayInput` via:
//
//          ServiceAzureBotArray{ ServiceAzureBotArgs{...} }
type ServiceAzureBotArrayInput interface {
	pulumi.Input

	ToServiceAzureBotArrayOutput() ServiceAzureBotArrayOutput
	ToServiceAzureBotArrayOutputWithContext(context.Context) ServiceAzureBotArrayOutput
}

type ServiceAzureBotArray []ServiceAzureBotInput

func (ServiceAzureBotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAzureBot)(nil)).Elem()
}

func (i ServiceAzureBotArray) ToServiceAzureBotArrayOutput() ServiceAzureBotArrayOutput {
	return i.ToServiceAzureBotArrayOutputWithContext(context.Background())
}

func (i ServiceAzureBotArray) ToServiceAzureBotArrayOutputWithContext(ctx context.Context) ServiceAzureBotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAzureBotArrayOutput)
}

// ServiceAzureBotMapInput is an input type that accepts ServiceAzureBotMap and ServiceAzureBotMapOutput values.
// You can construct a concrete instance of `ServiceAzureBotMapInput` via:
//
//          ServiceAzureBotMap{ "key": ServiceAzureBotArgs{...} }
type ServiceAzureBotMapInput interface {
	pulumi.Input

	ToServiceAzureBotMapOutput() ServiceAzureBotMapOutput
	ToServiceAzureBotMapOutputWithContext(context.Context) ServiceAzureBotMapOutput
}

type ServiceAzureBotMap map[string]ServiceAzureBotInput

func (ServiceAzureBotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAzureBot)(nil)).Elem()
}

func (i ServiceAzureBotMap) ToServiceAzureBotMapOutput() ServiceAzureBotMapOutput {
	return i.ToServiceAzureBotMapOutputWithContext(context.Background())
}

func (i ServiceAzureBotMap) ToServiceAzureBotMapOutputWithContext(ctx context.Context) ServiceAzureBotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAzureBotMapOutput)
}

type ServiceAzureBotOutput struct{ *pulumi.OutputState }

func (ServiceAzureBotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAzureBot)(nil))
}

func (o ServiceAzureBotOutput) ToServiceAzureBotOutput() ServiceAzureBotOutput {
	return o
}

func (o ServiceAzureBotOutput) ToServiceAzureBotOutputWithContext(ctx context.Context) ServiceAzureBotOutput {
	return o
}

func (o ServiceAzureBotOutput) ToServiceAzureBotPtrOutput() ServiceAzureBotPtrOutput {
	return o.ToServiceAzureBotPtrOutputWithContext(context.Background())
}

func (o ServiceAzureBotOutput) ToServiceAzureBotPtrOutputWithContext(ctx context.Context) ServiceAzureBotPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAzureBot) *ServiceAzureBot {
		return &v
	}).(ServiceAzureBotPtrOutput)
}

type ServiceAzureBotPtrOutput struct{ *pulumi.OutputState }

func (ServiceAzureBotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAzureBot)(nil))
}

func (o ServiceAzureBotPtrOutput) ToServiceAzureBotPtrOutput() ServiceAzureBotPtrOutput {
	return o
}

func (o ServiceAzureBotPtrOutput) ToServiceAzureBotPtrOutputWithContext(ctx context.Context) ServiceAzureBotPtrOutput {
	return o
}

func (o ServiceAzureBotPtrOutput) Elem() ServiceAzureBotOutput {
	return o.ApplyT(func(v *ServiceAzureBot) ServiceAzureBot {
		if v != nil {
			return *v
		}
		var ret ServiceAzureBot
		return ret
	}).(ServiceAzureBotOutput)
}

type ServiceAzureBotArrayOutput struct{ *pulumi.OutputState }

func (ServiceAzureBotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAzureBot)(nil))
}

func (o ServiceAzureBotArrayOutput) ToServiceAzureBotArrayOutput() ServiceAzureBotArrayOutput {
	return o
}

func (o ServiceAzureBotArrayOutput) ToServiceAzureBotArrayOutputWithContext(ctx context.Context) ServiceAzureBotArrayOutput {
	return o
}

func (o ServiceAzureBotArrayOutput) Index(i pulumi.IntInput) ServiceAzureBotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAzureBot {
		return vs[0].([]ServiceAzureBot)[vs[1].(int)]
	}).(ServiceAzureBotOutput)
}

type ServiceAzureBotMapOutput struct{ *pulumi.OutputState }

func (ServiceAzureBotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceAzureBot)(nil))
}

func (o ServiceAzureBotMapOutput) ToServiceAzureBotMapOutput() ServiceAzureBotMapOutput {
	return o
}

func (o ServiceAzureBotMapOutput) ToServiceAzureBotMapOutputWithContext(ctx context.Context) ServiceAzureBotMapOutput {
	return o
}

func (o ServiceAzureBotMapOutput) MapIndex(k pulumi.StringInput) ServiceAzureBotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceAzureBot {
		return vs[0].(map[string]ServiceAzureBot)[vs[1].(string)]
	}).(ServiceAzureBotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAzureBotInput)(nil)).Elem(), &ServiceAzureBot{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAzureBotPtrInput)(nil)).Elem(), &ServiceAzureBot{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAzureBotArrayInput)(nil)).Elem(), ServiceAzureBotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAzureBotMapInput)(nil)).Elem(), ServiceAzureBotMap{})
	pulumi.RegisterOutputType(ServiceAzureBotOutput{})
	pulumi.RegisterOutputType(ServiceAzureBotPtrOutput{})
	pulumi.RegisterOutputType(ServiceAzureBotArrayOutput{})
	pulumi.RegisterOutputType(ServiceAzureBotMapOutput{})
}
