// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Function App deployment Slot.
//
// ## Example Usage
// ### With App Service Plan)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		exampleFunctionApp, err := appservice.NewFunctionApp(ctx, "exampleFunctionApp", &appservice.FunctionAppArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			StorageAccountName:      exampleAccount.Name,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewFunctionAppSlot(ctx, "exampleFunctionAppSlot", &appservice.FunctionAppSlotArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			FunctionAppName:         exampleFunctionApp.Name,
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
//
// ## Import
//
// Function Apps Deployment Slots can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/functionAppSlot:FunctionAppSlot functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1/slots/staging
// ```
type FunctionAppSlot struct {
	pulumi.CustomResourceState

	// The ID of the App Service Plan within which to create this Function App Slot.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// An `authSettings` block as defined below.
	AuthSettings FunctionAppSlotAuthSettingsOutput `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolOutput `pulumi:"clientAffinityEnabled"`
	// A `connectionString` block as defined below.
	ConnectionStrings FunctionAppSlotConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrOutput `pulumi:"dailyMemoryTimeQuota"`
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// Should the built-in logging of the Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrOutput `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	FunctionAppName pulumi.StringOutput `pulumi:"functionAppName"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity FunctionAppSlotIdentityOutput `pulumi:"identity"`
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
	// The name of the resource group in which to create the Function App Slot.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSlotSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
	SiteCredentials FunctionAppSlotSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringOutput `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by the Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFunctionAppSlot registers a new resource with the given unique name, arguments, and options.
func NewFunctionAppSlot(ctx *pulumi.Context,
	name string, args *FunctionAppSlotArgs, opts ...pulumi.ResourceOption) (*FunctionAppSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServicePlanId == nil {
		return nil, errors.New("invalid value for required argument 'AppServicePlanId'")
	}
	if args.FunctionAppName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionAppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountAccessKey'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	var resource FunctionAppSlot
	err := ctx.RegisterResource("azure:appservice/functionAppSlot:FunctionAppSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionAppSlot gets an existing FunctionAppSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionAppSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionAppSlotState, opts ...pulumi.ResourceOption) (*FunctionAppSlot, error) {
	var resource FunctionAppSlot
	err := ctx.ReadResource("azure:appservice/functionAppSlot:FunctionAppSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionAppSlot resources.
type functionAppSlotState struct {
	// The ID of the App Service Plan within which to create this Function App Slot.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// An `authSettings` block as defined below.
	AuthSettings *FunctionAppSlotAuthSettings `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// A `connectionString` block as defined below.
	ConnectionStrings []FunctionAppSlotConnectionString `pulumi:"connectionStrings"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname *string `pulumi:"defaultHostname"`
	// Should the built-in logging of the Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging *bool `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled *bool `pulumi:"enabled"`
	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	FunctionAppName *string `pulumi:"functionAppName"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *FunctionAppSlotIdentity `pulumi:"identity"`
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
	// The name of the resource group in which to create the Function App Slot.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *FunctionAppSlotSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
	SiteCredentials []FunctionAppSlotSiteCredential `pulumi:"siteCredentials"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by the Function App (such as the dashboard, logs).
	StorageAccountName *string `pulumi:"storageAccountName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version *string `pulumi:"version"`
}

type FunctionAppSlotState struct {
	// The ID of the App Service Plan within which to create this Function App Slot.
	AppServicePlanId pulumi.StringPtrInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// An `authSettings` block as defined below.
	AuthSettings FunctionAppSlotAuthSettingsPtrInput
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// A `connectionString` block as defined below.
	ConnectionStrings FunctionAppSlotConnectionStringArrayInput
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
	DefaultHostname pulumi.StringPtrInput
	// Should the built-in logging of the Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrInput
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrInput
	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	FunctionAppName pulumi.StringPtrInput
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity FunctionAppSlotIdentityPtrInput
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
	// The name of the resource group in which to create the Function App Slot.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSlotSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
	SiteCredentials FunctionAppSlotSiteCredentialArrayInput
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringPtrInput
	// The backend storage account name which will be used by the Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrInput
}

func (FunctionAppSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAppSlotState)(nil)).Elem()
}

type functionAppSlotArgs struct {
	// The ID of the App Service Plan within which to create this Function App Slot.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// An `authSettings` block as defined below.
	AuthSettings *FunctionAppSlotAuthSettings `pulumi:"authSettings"`
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// A `connectionString` block as defined below.
	ConnectionStrings []FunctionAppSlotConnectionString `pulumi:"connectionStrings"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// Should the built-in logging of the Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging *bool `pulumi:"enableBuiltinLogging"`
	// Is the Function App enabled?
	Enabled *bool `pulumi:"enabled"`
	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	FunctionAppName string `pulumi:"functionAppName"`
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *FunctionAppSlotIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A string indicating the Operating System type for this function app.
	OsType *string `pulumi:"osType"`
	// The name of the resource group in which to create the Function App Slot.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *FunctionAppSlotSiteConfig `pulumi:"siteConfig"`
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey string `pulumi:"storageAccountAccessKey"`
	// The backend storage account name which will be used by the Function App (such as the dashboard, logs).
	StorageAccountName string `pulumi:"storageAccountName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a FunctionAppSlot resource.
type FunctionAppSlotArgs struct {
	// The ID of the App Service Plan within which to create this Function App Slot.
	AppServicePlanId pulumi.StringInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// An `authSettings` block as defined below.
	AuthSettings FunctionAppSlotAuthSettingsPtrInput
	// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// A `connectionString` block as defined below.
	ConnectionStrings FunctionAppSlotConnectionStringArrayInput
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// Should the built-in logging of the Function App be enabled? Defaults to `true`.
	EnableBuiltinLogging pulumi.BoolPtrInput
	// Is the Function App enabled?
	Enabled pulumi.BoolPtrInput
	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	FunctionAppName pulumi.StringInput
	// Can the Function App only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity FunctionAppSlotIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Function App. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A string indicating the Operating System type for this function app.
	OsType pulumi.StringPtrInput
	// The name of the resource group in which to create the Function App Slot.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` object as defined below.
	SiteConfig FunctionAppSlotSiteConfigPtrInput
	// The access key which will be used to access the backend storage account for the Function App.
	StorageAccountAccessKey pulumi.StringInput
	// The backend storage account name which will be used by the Function App (such as the dashboard, logs).
	StorageAccountName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The runtime version associated with the Function App. Defaults to `~1`.
	Version pulumi.StringPtrInput
}

func (FunctionAppSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAppSlotArgs)(nil)).Elem()
}

type FunctionAppSlotInput interface {
	pulumi.Input

	ToFunctionAppSlotOutput() FunctionAppSlotOutput
	ToFunctionAppSlotOutputWithContext(ctx context.Context) FunctionAppSlotOutput
}

func (*FunctionAppSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionAppSlot)(nil))
}

func (i *FunctionAppSlot) ToFunctionAppSlotOutput() FunctionAppSlotOutput {
	return i.ToFunctionAppSlotOutputWithContext(context.Background())
}

func (i *FunctionAppSlot) ToFunctionAppSlotOutputWithContext(ctx context.Context) FunctionAppSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppSlotOutput)
}

func (i *FunctionAppSlot) ToFunctionAppSlotPtrOutput() FunctionAppSlotPtrOutput {
	return i.ToFunctionAppSlotPtrOutputWithContext(context.Background())
}

func (i *FunctionAppSlot) ToFunctionAppSlotPtrOutputWithContext(ctx context.Context) FunctionAppSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppSlotPtrOutput)
}

type FunctionAppSlotPtrInput interface {
	pulumi.Input

	ToFunctionAppSlotPtrOutput() FunctionAppSlotPtrOutput
	ToFunctionAppSlotPtrOutputWithContext(ctx context.Context) FunctionAppSlotPtrOutput
}

type functionAppSlotPtrType FunctionAppSlotArgs

func (*functionAppSlotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionAppSlot)(nil))
}

func (i *functionAppSlotPtrType) ToFunctionAppSlotPtrOutput() FunctionAppSlotPtrOutput {
	return i.ToFunctionAppSlotPtrOutputWithContext(context.Background())
}

func (i *functionAppSlotPtrType) ToFunctionAppSlotPtrOutputWithContext(ctx context.Context) FunctionAppSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppSlotPtrOutput)
}

// FunctionAppSlotArrayInput is an input type that accepts FunctionAppSlotArray and FunctionAppSlotArrayOutput values.
// You can construct a concrete instance of `FunctionAppSlotArrayInput` via:
//
//          FunctionAppSlotArray{ FunctionAppSlotArgs{...} }
type FunctionAppSlotArrayInput interface {
	pulumi.Input

	ToFunctionAppSlotArrayOutput() FunctionAppSlotArrayOutput
	ToFunctionAppSlotArrayOutputWithContext(context.Context) FunctionAppSlotArrayOutput
}

type FunctionAppSlotArray []FunctionAppSlotInput

func (FunctionAppSlotArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FunctionAppSlot)(nil))
}

func (i FunctionAppSlotArray) ToFunctionAppSlotArrayOutput() FunctionAppSlotArrayOutput {
	return i.ToFunctionAppSlotArrayOutputWithContext(context.Background())
}

func (i FunctionAppSlotArray) ToFunctionAppSlotArrayOutputWithContext(ctx context.Context) FunctionAppSlotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppSlotArrayOutput)
}

// FunctionAppSlotMapInput is an input type that accepts FunctionAppSlotMap and FunctionAppSlotMapOutput values.
// You can construct a concrete instance of `FunctionAppSlotMapInput` via:
//
//          FunctionAppSlotMap{ "key": FunctionAppSlotArgs{...} }
type FunctionAppSlotMapInput interface {
	pulumi.Input

	ToFunctionAppSlotMapOutput() FunctionAppSlotMapOutput
	ToFunctionAppSlotMapOutputWithContext(context.Context) FunctionAppSlotMapOutput
}

type FunctionAppSlotMap map[string]FunctionAppSlotInput

func (FunctionAppSlotMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FunctionAppSlot)(nil))
}

func (i FunctionAppSlotMap) ToFunctionAppSlotMapOutput() FunctionAppSlotMapOutput {
	return i.ToFunctionAppSlotMapOutputWithContext(context.Background())
}

func (i FunctionAppSlotMap) ToFunctionAppSlotMapOutputWithContext(ctx context.Context) FunctionAppSlotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAppSlotMapOutput)
}

type FunctionAppSlotOutput struct {
	*pulumi.OutputState
}

func (FunctionAppSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionAppSlot)(nil))
}

func (o FunctionAppSlotOutput) ToFunctionAppSlotOutput() FunctionAppSlotOutput {
	return o
}

func (o FunctionAppSlotOutput) ToFunctionAppSlotOutputWithContext(ctx context.Context) FunctionAppSlotOutput {
	return o
}

func (o FunctionAppSlotOutput) ToFunctionAppSlotPtrOutput() FunctionAppSlotPtrOutput {
	return o.ToFunctionAppSlotPtrOutputWithContext(context.Background())
}

func (o FunctionAppSlotOutput) ToFunctionAppSlotPtrOutputWithContext(ctx context.Context) FunctionAppSlotPtrOutput {
	return o.ApplyT(func(v FunctionAppSlot) *FunctionAppSlot {
		return &v
	}).(FunctionAppSlotPtrOutput)
}

type FunctionAppSlotPtrOutput struct {
	*pulumi.OutputState
}

func (FunctionAppSlotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionAppSlot)(nil))
}

func (o FunctionAppSlotPtrOutput) ToFunctionAppSlotPtrOutput() FunctionAppSlotPtrOutput {
	return o
}

func (o FunctionAppSlotPtrOutput) ToFunctionAppSlotPtrOutputWithContext(ctx context.Context) FunctionAppSlotPtrOutput {
	return o
}

type FunctionAppSlotArrayOutput struct{ *pulumi.OutputState }

func (FunctionAppSlotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionAppSlot)(nil))
}

func (o FunctionAppSlotArrayOutput) ToFunctionAppSlotArrayOutput() FunctionAppSlotArrayOutput {
	return o
}

func (o FunctionAppSlotArrayOutput) ToFunctionAppSlotArrayOutputWithContext(ctx context.Context) FunctionAppSlotArrayOutput {
	return o
}

func (o FunctionAppSlotArrayOutput) Index(i pulumi.IntInput) FunctionAppSlotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionAppSlot {
		return vs[0].([]FunctionAppSlot)[vs[1].(int)]
	}).(FunctionAppSlotOutput)
}

type FunctionAppSlotMapOutput struct{ *pulumi.OutputState }

func (FunctionAppSlotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FunctionAppSlot)(nil))
}

func (o FunctionAppSlotMapOutput) ToFunctionAppSlotMapOutput() FunctionAppSlotMapOutput {
	return o
}

func (o FunctionAppSlotMapOutput) ToFunctionAppSlotMapOutputWithContext(ctx context.Context) FunctionAppSlotMapOutput {
	return o
}

func (o FunctionAppSlotMapOutput) MapIndex(k pulumi.StringInput) FunctionAppSlotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FunctionAppSlot {
		return vs[0].(map[string]FunctionAppSlot)[vs[1].(string)]
	}).(FunctionAppSlotOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionAppSlotOutput{})
	pulumi.RegisterOutputType(FunctionAppSlotPtrOutput{})
	pulumi.RegisterOutputType(FunctionAppSlotArrayOutput{})
	pulumi.RegisterOutputType(FunctionAppSlotMapOutput{})
}
