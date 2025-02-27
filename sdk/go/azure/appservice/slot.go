// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service Slot (within an App Service).
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// ## Example Usage
// ### Net 4.X)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.AnyMap{
// 				"azi_id": pulumi.Any(1),
// 			},
// 			ByteLength: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
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
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 			SiteConfig: &appservice.AppServiceSiteConfigArgs{
// 				DotnetFrameworkVersion: pulumi.String("v4.0"),
// 			},
// 			AppSettings: pulumi.StringMap{
// 				"SOME_KEY": pulumi.String("some-value"),
// 			},
// 			ConnectionStrings: appservice.AppServiceConnectionStringArray{
// 				&appservice.AppServiceConnectionStringArgs{
// 					Name:  pulumi.String("Database"),
// 					Type:  pulumi.String("SQLServer"),
// 					Value: pulumi.String("Server=some-server.mydomain.com;Integrated Security=SSPI"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewSlot(ctx, "exampleSlot", &appservice.SlotArgs{
// 			AppServiceName:    exampleAppService.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 			SiteConfig: &appservice.SlotSiteConfigArgs{
// 				DotnetFrameworkVersion: pulumi.String("v4.0"),
// 			},
// 			AppSettings: pulumi.StringMap{
// 				"SOME_KEY": pulumi.String("some-value"),
// 			},
// 			ConnectionStrings: appservice.SlotConnectionStringArray{
// 				&appservice.SlotConnectionStringArgs{
// 					Name:  pulumi.String("Database"),
// 					Type:  pulumi.String("SQLServer"),
// 					Value: pulumi.String("Server=some-server.mydomain.com;Integrated Security=SSPI"),
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
// ### Java 1.8)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
// 			Keepers: pulumi.AnyMap{
// 				"azi_id": pulumi.Any(1),
// 			},
// 			ByteLength: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
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
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 			SiteConfig: &appservice.AppServiceSiteConfigArgs{
// 				JavaVersion:          pulumi.String("1.8"),
// 				JavaContainer:        pulumi.String("JETTY"),
// 				JavaContainerVersion: pulumi.String("9.3"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewSlot(ctx, "exampleSlot", &appservice.SlotArgs{
// 			AppServiceName:    exampleAppService.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 			SiteConfig: &appservice.SlotSiteConfigArgs{
// 				JavaVersion:          pulumi.String("1.8"),
// 				JavaContainer:        pulumi.String("JETTY"),
// 				JavaContainerVersion: pulumi.String("9.3"),
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
// App Service Slots can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/slot:Slot instance1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/website1/slots/instance1
// ```
type Slot struct {
	pulumi.CustomResourceState

	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsOutput `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolOutput `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringOutput `pulumi:"defaultSiteHostname"`
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity SlotIdentityOutput `pulumi:"identity"`
	// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. See [Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity) for more information.
	KeyVaultReferenceIdentityId pulumi.StringOutput `pulumi:"keyVaultReferenceIdentityId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `logs` block as defined below.
	Logs SlotLogsOutput `pulumi:"logs"`
	// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service slot.
	SiteCredentials SlotSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts SlotStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSlot registers a new resource with the given unique name, arguments, and options.
func NewSlot(ctx *pulumi.Context,
	name string, args *SlotArgs, opts ...pulumi.ResourceOption) (*Slot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceName'")
	}
	if args.AppServicePlanId == nil {
		return nil, errors.New("invalid value for required argument 'AppServicePlanId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Slot
	err := ctx.RegisterResource("azure:appservice/slot:Slot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlot gets an existing Slot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlotState, opts ...pulumi.ResourceOption) (*Slot, error) {
	var resource Slot
	err := ctx.ReadResource("azure:appservice/slot:Slot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slot resources.
type slotState struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *SlotAuthSettings `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings []SlotConnectionString `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname *string `pulumi:"defaultSiteHostname"`
	// Is the App Service Slot Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *SlotIdentity `pulumi:"identity"`
	// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. See [Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity) for more information.
	KeyVaultReferenceIdentityId *string `pulumi:"keyVaultReferenceIdentityId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *SlotLogs `pulumi:"logs"`
	// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *SlotSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service slot.
	SiteCredentials []SlotSiteCredential `pulumi:"siteCredentials"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []SlotStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SlotState struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringPtrInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsPtrInput
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayInput
	// The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringPtrInput
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity SlotIdentityPtrInput
	// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. See [Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity) for more information.
	KeyVaultReferenceIdentityId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs SlotLogsPtrInput
	// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service slot.
	SiteCredentials SlotSiteCredentialArrayInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts SlotStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*slotState)(nil)).Elem()
}

type slotArgs struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *SlotAuthSettings `pulumi:"authSettings"`
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// An `connectionString` block as defined below.
	ConnectionStrings []SlotConnectionString `pulumi:"connectionStrings"`
	// Is the App Service Slot Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// An `identity` block as defined below.
	Identity *SlotIdentity `pulumi:"identity"`
	// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. See [Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity) for more information.
	KeyVaultReferenceIdentityId *string `pulumi:"keyVaultReferenceIdentityId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *SlotLogs `pulumi:"logs"`
	// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` object as defined below.
	SiteConfig *SlotSiteConfig `pulumi:"siteConfig"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []SlotStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Slot resource.
type SlotArgs struct {
	// The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings SlotAuthSettingsPtrInput
	// Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// An `connectionString` block as defined below.
	ConnectionStrings SlotConnectionStringArrayInput
	// Is the App Service Slot Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity SlotIdentityPtrInput
	// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. See [Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity) for more information.
	KeyVaultReferenceIdentityId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs SlotLogsPtrInput
	// Specifies the name of the App Service Slot component. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service Slot component.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` object as defined below.
	SiteConfig SlotSiteConfigPtrInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts SlotStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slotArgs)(nil)).Elem()
}

type SlotInput interface {
	pulumi.Input

	ToSlotOutput() SlotOutput
	ToSlotOutputWithContext(ctx context.Context) SlotOutput
}

func (*Slot) ElementType() reflect.Type {
	return reflect.TypeOf((**Slot)(nil)).Elem()
}

func (i *Slot) ToSlotOutput() SlotOutput {
	return i.ToSlotOutputWithContext(context.Background())
}

func (i *Slot) ToSlotOutputWithContext(ctx context.Context) SlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotOutput)
}

// SlotArrayInput is an input type that accepts SlotArray and SlotArrayOutput values.
// You can construct a concrete instance of `SlotArrayInput` via:
//
//          SlotArray{ SlotArgs{...} }
type SlotArrayInput interface {
	pulumi.Input

	ToSlotArrayOutput() SlotArrayOutput
	ToSlotArrayOutputWithContext(context.Context) SlotArrayOutput
}

type SlotArray []SlotInput

func (SlotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slot)(nil)).Elem()
}

func (i SlotArray) ToSlotArrayOutput() SlotArrayOutput {
	return i.ToSlotArrayOutputWithContext(context.Background())
}

func (i SlotArray) ToSlotArrayOutputWithContext(ctx context.Context) SlotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotArrayOutput)
}

// SlotMapInput is an input type that accepts SlotMap and SlotMapOutput values.
// You can construct a concrete instance of `SlotMapInput` via:
//
//          SlotMap{ "key": SlotArgs{...} }
type SlotMapInput interface {
	pulumi.Input

	ToSlotMapOutput() SlotMapOutput
	ToSlotMapOutputWithContext(context.Context) SlotMapOutput
}

type SlotMap map[string]SlotInput

func (SlotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slot)(nil)).Elem()
}

func (i SlotMap) ToSlotMapOutput() SlotMapOutput {
	return i.ToSlotMapOutputWithContext(context.Background())
}

func (i SlotMap) ToSlotMapOutputWithContext(ctx context.Context) SlotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotMapOutput)
}

type SlotOutput struct{ *pulumi.OutputState }

func (SlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slot)(nil)).Elem()
}

func (o SlotOutput) ToSlotOutput() SlotOutput {
	return o
}

func (o SlotOutput) ToSlotOutputWithContext(ctx context.Context) SlotOutput {
	return o
}

type SlotArrayOutput struct{ *pulumi.OutputState }

func (SlotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slot)(nil)).Elem()
}

func (o SlotArrayOutput) ToSlotArrayOutput() SlotArrayOutput {
	return o
}

func (o SlotArrayOutput) ToSlotArrayOutputWithContext(ctx context.Context) SlotArrayOutput {
	return o
}

func (o SlotArrayOutput) Index(i pulumi.IntInput) SlotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Slot {
		return vs[0].([]*Slot)[vs[1].(int)]
	}).(SlotOutput)
}

type SlotMapOutput struct{ *pulumi.OutputState }

func (SlotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slot)(nil)).Elem()
}

func (o SlotMapOutput) ToSlotMapOutput() SlotMapOutput {
	return o
}

func (o SlotMapOutput) ToSlotMapOutputWithContext(ctx context.Context) SlotMapOutput {
	return o
}

func (o SlotMapOutput) MapIndex(k pulumi.StringInput) SlotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Slot {
		return vs[0].(map[string]*Slot)[vs[1].(string)]
	}).(SlotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SlotInput)(nil)).Elem(), &Slot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotArrayInput)(nil)).Elem(), SlotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotMapInput)(nil)).Elem(), SlotMap{})
	pulumi.RegisterOutputType(SlotOutput{})
	pulumi.RegisterOutputType(SlotArrayOutput{})
	pulumi.RegisterOutputType(SlotMapOutput{})
}
