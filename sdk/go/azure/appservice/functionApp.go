// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Function App.
//
// > **Note:** To connect an Azure Function App and a subnet within the same region `appservice.VirtualNetworkSwiftConnection` can be used.
// For an example, check the `appservice.VirtualNetworkSwiftConnection` documentation.
//
// ## Example Usage
// ### With App Service Plan)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewFunctionApp(ctx, "exampleFunctionApp", &appservice.FunctionAppArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			StorageAccountName:      exampleAccount.Name,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### In A Consumption Plan)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Kind:              pulumi.Any("FunctionApp"),
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Dynamic"),
// 				Size: pulumi.String("Y1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewFunctionApp(ctx, "exampleFunctionApp", &appservice.FunctionAppArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			StorageAccountName:      exampleAccount.Name,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Linux)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Kind:              pulumi.Any("Linux"),
// 			Reserved:          pulumi.Bool(true),
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Dynamic"),
// 				Size: pulumi.String("Y1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewFunctionApp(ctx, "exampleFunctionApp", &appservice.FunctionAppArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			StorageAccountName:      exampleAccount.Name,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 			OsType:                  pulumi.String("linux"),
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
// Function Apps can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/functionApp:FunctionApp functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1
// ```
type FunctionApp struct {
	pulumi.CustomResourceState

	// The ID of the App Service Plan within which to create this Function App.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings FunctionAppAuthSettingsOutput `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolOutput `pulumi:"clientAffinityEnabled"`
	// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
	ClientCertMode pulumi.StringPtrOutput `pulumi:"clientCertMode"`
	// An `connectionString` block as defined below.
	ConnectionStrings FunctionAppConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId pulumi.StringOutput `pulumi:"customDomainVerificationId"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrOutput `pulumi:"dailyMemoryTimeQuota"`
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// Should the built-in logging of this Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrOutput `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity FunctionAppIdentityOutput `pulumi:"identity"`
	// The Function App kind - such as `functionapp,linux,container`
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A string indicating the Operating System type for this function app.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringOutput `pulumi:"outboundIpAddresses"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringOutput `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the Function App.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials FunctionAppSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// A `sourceControl` block, as defined below.
	SourceControl FunctionAppSourceControlOutput `pulumi:"sourceControl"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringOutput `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// Deprecated: Deprecated in favour of `storage_account_name` and `storage_account_access_key`
	StorageConnectionString pulumi.StringOutput `pulumi:"storageConnectionString"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFunctionApp registers a new resource with the given unique name, arguments, and options.
func NewFunctionApp(ctx *pulumi.Context,
	name string, args *FunctionAppArgs, opts ...pulumi.ResourceOption) (*FunctionApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServicePlanId == nil {
		return nil, errors.New("invalid value for required argument 'AppServicePlanId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource FunctionApp
	err := ctx.RegisterResource("azure:appservice/functionApp:FunctionApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionApp gets an existing FunctionApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionAppState, opts ...pulumi.ResourceOption) (*FunctionApp, error) {
	var resource FunctionApp
	err := ctx.ReadResource("azure:appservice/functionApp:FunctionApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionApp resources.
type functionAppState struct {
	// The ID of the App Service Plan within which to create this Function App.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *FunctionAppAuthSettings `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
	ClientCertMode *string `pulumi:"clientCertMode"`
	// An `connectionString` block as defined below.
	ConnectionStrings []FunctionAppConnectionString `pulumi:"connectionStrings"`
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId *string `pulumi:"customDomainVerificationId"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname *string `pulumi:"defaultHostname"`
	// Should the built-in logging of this Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging *bool `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *FunctionAppIdentity `pulumi:"identity"`
	// The Function App kind - such as `functionapp,linux,container`
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A string indicating the Operating System type for this function app.
	OsType *string `pulumi:"osType"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses *string `pulumi:"outboundIpAddresses"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses *string `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the Function App.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *FunctionAppSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials []FunctionAppSiteCredential `pulumi:"siteCredentials"`
	// A `sourceControl` block, as defined below.
	SourceControl *FunctionAppSourceControl `pulumi:"sourceControl"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
	StorageAccountName *string `pulumi:"storageAccountName"`
	// Deprecated: Deprecated in favour of `storage_account_name` and `storage_account_access_key`
	StorageConnectionString *string `pulumi:"storageConnectionString"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version *string `pulumi:"version"`
}

type FunctionAppState struct {
	// The ID of the App Service Plan within which to create this Function App.
	AppServicePlanId pulumi.StringPtrInput
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings FunctionAppAuthSettingsPtrInput
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
	ClientCertMode pulumi.StringPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings FunctionAppConnectionStringArrayInput
	// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
	CustomDomainVerificationId pulumi.StringPtrInput
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname pulumi.StringPtrInput
	// Should the built-in logging of this Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrInput
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrInput
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity FunctionAppIdentityPtrInput
	// The Function App kind - such as `functionapp,linux,container`
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A string indicating the Operating System type for this function app.
	OsType pulumi.StringPtrInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringPtrInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringPtrInput
	// The name of the resource group in which to create the Function App.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials FunctionAppSiteCredentialArrayInput
	// A `sourceControl` block, as defined below.
	SourceControl FunctionAppSourceControlPtrInput
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringPtrInput
	// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringPtrInput
	// Deprecated: Deprecated in favour of `storage_account_name` and `storage_account_access_key`
	StorageConnectionString pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrInput
}

func (FunctionAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAppState)(nil)).Elem()
}

type functionAppArgs struct {
	// The ID of the App Service Plan within which to create this Function App.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *FunctionAppAuthSettings `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
	ClientCertMode *string `pulumi:"clientCertMode"`
	// An `connectionString` block as defined below.
	ConnectionStrings []FunctionAppConnectionString `pulumi:"connectionStrings"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// Should the built-in logging of this Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging *bool `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *FunctionAppIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A string indicating the Operating System type for this function app.
	OsType *string `pulumi:"osType"`
	// The name of the resource group in which to create the Function App.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *FunctionAppSiteConfig `pulumi:"siteConfig"`
	// A `sourceControl` block, as defined below.
	SourceControl *FunctionAppSourceControl `pulumi:"sourceControl"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
	StorageAccountName *string `pulumi:"storageAccountName"`
	// Deprecated: Deprecated in favour of `storage_account_name` and `storage_account_access_key`
	StorageConnectionString *string `pulumi:"storageConnectionString"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a FunctionApp resource.
type FunctionAppArgs struct {
	// The ID of the App Service Plan within which to create this Function App.
	AppServicePlanId pulumi.StringInput
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings FunctionAppAuthSettingsPtrInput
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
	ClientCertMode pulumi.StringPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings FunctionAppConnectionStringArrayInput
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// Should the built-in logging of this Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrInput
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrInput
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity FunctionAppIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A string indicating the Operating System type for this function app.
	OsType pulumi.StringPtrInput
	// The name of the resource group in which to create the Function App.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSiteConfigPtrInput
	// A `sourceControl` block, as defined below.
	SourceControl FunctionAppSourceControlPtrInput
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringPtrInput
	// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringPtrInput
	// Deprecated: Deprecated in favour of `storage_account_name` and `storage_account_access_key`
	StorageConnectionString pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrInput
}

func (FunctionAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAppArgs)(nil)).Elem()
}

type FunctionAppInput interface {
	pulumi.Input

	ToFunctionAppOutput() FunctionAppOutput
	ToFunctionAppOutputWithContext(ctx context.Context) FunctionAppOutput
}

func (*FunctionApp) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionApp)(nil))
}

func (i *FunctionApp) ToFunctionAppOutput() FunctionAppOutput {
	return i.ToFunctionAppOutputWithContext(context.Background())
}

func (i *FunctionApp) ToFunctionAppOutputWithContext(ctx context.Context) FunctionAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppOutput)
}

func (i *FunctionApp) ToFunctionAppPtrOutput() FunctionAppPtrOutput {
	return i.ToFunctionAppPtrOutputWithContext(context.Background())
}

func (i *FunctionApp) ToFunctionAppPtrOutputWithContext(ctx context.Context) FunctionAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppPtrOutput)
}

type FunctionAppPtrInput interface {
	pulumi.Input

	ToFunctionAppPtrOutput() FunctionAppPtrOutput
	ToFunctionAppPtrOutputWithContext(ctx context.Context) FunctionAppPtrOutput
}

type functionAppPtrType FunctionAppArgs

func (*functionAppPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionApp)(nil))
}

func (i *functionAppPtrType) ToFunctionAppPtrOutput() FunctionAppPtrOutput {
	return i.ToFunctionAppPtrOutputWithContext(context.Background())
}

func (i *functionAppPtrType) ToFunctionAppPtrOutputWithContext(ctx context.Context) FunctionAppPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppPtrOutput)
}

// FunctionAppArrayInput is an input type that accepts FunctionAppArray and FunctionAppArrayOutput values.
// You can construct a concrete instance of `FunctionAppArrayInput` via:
//
//          FunctionAppArray{ FunctionAppArgs{...} }
type FunctionAppArrayInput interface {
	pulumi.Input

	ToFunctionAppArrayOutput() FunctionAppArrayOutput
	ToFunctionAppArrayOutputWithContext(context.Context) FunctionAppArrayOutput
}

type FunctionAppArray []FunctionAppInput

func (FunctionAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionApp)(nil)).Elem()
}

func (i FunctionAppArray) ToFunctionAppArrayOutput() FunctionAppArrayOutput {
	return i.ToFunctionAppArrayOutputWithContext(context.Background())
}

func (i FunctionAppArray) ToFunctionAppArrayOutputWithContext(ctx context.Context) FunctionAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppArrayOutput)
}

// FunctionAppMapInput is an input type that accepts FunctionAppMap and FunctionAppMapOutput values.
// You can construct a concrete instance of `FunctionAppMapInput` via:
//
//          FunctionAppMap{ "key": FunctionAppArgs{...} }
type FunctionAppMapInput interface {
	pulumi.Input

	ToFunctionAppMapOutput() FunctionAppMapOutput
	ToFunctionAppMapOutputWithContext(context.Context) FunctionAppMapOutput
}

type FunctionAppMap map[string]FunctionAppInput

func (FunctionAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionApp)(nil)).Elem()
}

func (i FunctionAppMap) ToFunctionAppMapOutput() FunctionAppMapOutput {
	return i.ToFunctionAppMapOutputWithContext(context.Background())
}

func (i FunctionAppMap) ToFunctionAppMapOutputWithContext(ctx context.Context) FunctionAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppMapOutput)
}

type FunctionAppOutput struct{ *pulumi.OutputState }

func (FunctionAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionApp)(nil))
}

func (o FunctionAppOutput) ToFunctionAppOutput() FunctionAppOutput {
	return o
}

func (o FunctionAppOutput) ToFunctionAppOutputWithContext(ctx context.Context) FunctionAppOutput {
	return o
}

func (o FunctionAppOutput) ToFunctionAppPtrOutput() FunctionAppPtrOutput {
	return o.ToFunctionAppPtrOutputWithContext(context.Background())
}

func (o FunctionAppOutput) ToFunctionAppPtrOutputWithContext(ctx context.Context) FunctionAppPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionApp) *FunctionApp {
		return &v
	}).(FunctionAppPtrOutput)
}

type FunctionAppPtrOutput struct{ *pulumi.OutputState }

func (FunctionAppPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionApp)(nil))
}

func (o FunctionAppPtrOutput) ToFunctionAppPtrOutput() FunctionAppPtrOutput {
	return o
}

func (o FunctionAppPtrOutput) ToFunctionAppPtrOutputWithContext(ctx context.Context) FunctionAppPtrOutput {
	return o
}

func (o FunctionAppPtrOutput) Elem() FunctionAppOutput {
	return o.ApplyT(func(v *FunctionApp) FunctionApp {
		if v != nil {
			return *v
		}
		var ret FunctionApp
		return ret
	}).(FunctionAppOutput)
}

type FunctionAppArrayOutput struct{ *pulumi.OutputState }

func (FunctionAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionApp)(nil))
}

func (o FunctionAppArrayOutput) ToFunctionAppArrayOutput() FunctionAppArrayOutput {
	return o
}

func (o FunctionAppArrayOutput) ToFunctionAppArrayOutputWithContext(ctx context.Context) FunctionAppArrayOutput {
	return o
}

func (o FunctionAppArrayOutput) Index(i pulumi.IntInput) FunctionAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionApp {
		return vs[0].([]FunctionApp)[vs[1].(int)]
	}).(FunctionAppOutput)
}

type FunctionAppMapOutput struct{ *pulumi.OutputState }

func (FunctionAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FunctionApp)(nil))
}

func (o FunctionAppMapOutput) ToFunctionAppMapOutput() FunctionAppMapOutput {
	return o
}

func (o FunctionAppMapOutput) ToFunctionAppMapOutputWithContext(ctx context.Context) FunctionAppMapOutput {
	return o
}

func (o FunctionAppMapOutput) MapIndex(k pulumi.StringInput) FunctionAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FunctionApp {
		return vs[0].(map[string]FunctionApp)[vs[1].(string)]
	}).(FunctionAppOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionAppOutput{})
	pulumi.RegisterOutputType(FunctionAppPtrOutput{})
	pulumi.RegisterOutputType(FunctionAppArrayOutput{})
	pulumi.RegisterOutputType(FunctionAppMapOutput{})
}
