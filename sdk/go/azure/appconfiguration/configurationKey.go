// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfiguration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure App Configuration Key.
//
// > **Note:** App Configuration Keys are provisioned using a Data Plane API which requires the role `App Configuration Data Owner` on either the App Configuration or a parent scope (such as the Resource Group/Subscription). [More information can be found in the Azure Documentation for App Configuration](https://docs.microsoft.com/azure/azure-app-configuration/concept-enable-rbac#azure-built-in-roles-for-azure-app-configuration).
//
// ## Example Usage
// ### `Kv` Type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appconfiguration"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appconf, err := appconfiguration.NewConfigurationStore(ctx, "appconf", &appconfiguration.ConfigurationStoreArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		appconfDataowner, err := authorization.NewAssignment(ctx, "appconfDataowner", &authorization.AssignmentArgs{
// 			Scope:              appconf.ID(),
// 			RoleDefinitionName: pulumi.String("App Configuration Data Owner"),
// 			PrincipalId:        pulumi.String(current.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appconfiguration.NewConfigurationKey(ctx, "test", &appconfiguration.ConfigurationKeyArgs{
// 			ConfigurationStoreId: appconf.ID(),
// 			Key:                  pulumi.String("appConfKey1"),
// 			Label:                pulumi.String("somelabel"),
// 			Value:                pulumi.String("a test"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			appconfDataowner,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### `Vault` Type
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appconfiguration"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appconf, err := appconfiguration.NewConfigurationStore(ctx, "appconf", &appconfiguration.ConfigurationStoreArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		kv, err := keyvault.NewKeyVault(ctx, "kv", &keyvault.KeyVaultArgs{
// 			Location:                pulumi.Any(azurerm_resource_group.Test.Location),
// 			ResourceGroupName:       pulumi.Any(azurerm_resource_group.Test.Name),
// 			TenantId:                pulumi.String(current.TenantId),
// 			SkuName:                 pulumi.String("premium"),
// 			SoftDeleteRetentionDays: pulumi.Int(7),
// 			AccessPolicies: keyvault.KeyVaultAccessPolicyArray{
// 				&keyvault.KeyVaultAccessPolicyArgs{
// 					TenantId: pulumi.String(current.TenantId),
// 					ObjectId: pulumi.String(current.ObjectId),
// 					KeyPermissions: pulumi.StringArray{
// 						pulumi.String("create"),
// 						pulumi.String("get"),
// 					},
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("set"),
// 						pulumi.String("get"),
// 						pulumi.String("delete"),
// 						pulumi.String("purge"),
// 						pulumi.String("recover"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		kvs, err := keyvault.NewSecret(ctx, "kvs", &keyvault.SecretArgs{
// 			Value:      pulumi.String("szechuan"),
// 			KeyVaultId: kv.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appconfDataowner, err := authorization.NewAssignment(ctx, "appconfDataowner", &authorization.AssignmentArgs{
// 			Scope:              appconf.ID(),
// 			RoleDefinitionName: pulumi.String("App Configuration Data Owner"),
// 			PrincipalId:        pulumi.String(current.ObjectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appconfiguration.NewConfigurationKey(ctx, "test", &appconfiguration.ConfigurationKeyArgs{
// 			ConfigurationStoreId: pulumi.Any(azurerm_app_configuration.Test.Id),
// 			Key:                  pulumi.String("key1"),
// 			Type:                 pulumi.String("vault"),
// 			Label:                pulumi.String("label1"),
// 			VaultKeyReference:    kvs.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			appconfDataowner,
// 		}))
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
// App Configuration Keys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/label1
// ```
//
//  If you wish to import a key with an empty label then sustitute the label's name with `%00`, like this
//
// ```sh
//  $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/%00
// ```
type ConfigurationKey struct {
	pulumi.CustomResourceState

	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringOutput `pulumi:"configurationStoreId"`
	// The content type of the App Configuration Key. This should only be set when type is set to `kv`.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The ETag of the key.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the App Configuration Key to create. Changing this forces a new resource to be created.
	Key pulumi.StringOutput `pulumi:"key"`
	// The label of the App Configuration Key.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// Should this App Configuration Key be Locked to prevent changes?
	Locked pulumi.BoolPtrOutput `pulumi:"locked"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The value of the App Configuration Key. This should only be set when type is set to `kv`.
	Value pulumi.StringOutput `pulumi:"value"`
	// The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
	VaultKeyReference pulumi.StringPtrOutput `pulumi:"vaultKeyReference"`
}

// NewConfigurationKey registers a new resource with the given unique name, arguments, and options.
func NewConfigurationKey(ctx *pulumi.Context,
	name string, args *ConfigurationKeyArgs, opts ...pulumi.ResourceOption) (*ConfigurationKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationStoreId'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource ConfigurationKey
	err := ctx.RegisterResource("azure:appconfiguration/configurationKey:ConfigurationKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationKey gets an existing ConfigurationKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationKeyState, opts ...pulumi.ResourceOption) (*ConfigurationKey, error) {
	var resource ConfigurationKey
	err := ctx.ReadResource("azure:appconfiguration/configurationKey:ConfigurationKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationKey resources.
type configurationKeyState struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId *string `pulumi:"configurationStoreId"`
	// The content type of the App Configuration Key. This should only be set when type is set to `kv`.
	ContentType *string `pulumi:"contentType"`
	// The ETag of the key.
	Etag *string `pulumi:"etag"`
	// The name of the App Configuration Key to create. Changing this forces a new resource to be created.
	Key *string `pulumi:"key"`
	// The label of the App Configuration Key.  Changing this forces a new resource to be created.
	Label *string `pulumi:"label"`
	// Should this App Configuration Key be Locked to prevent changes?
	Locked *bool `pulumi:"locked"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
	Type *string `pulumi:"type"`
	// The value of the App Configuration Key. This should only be set when type is set to `kv`.
	Value *string `pulumi:"value"`
	// The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
	VaultKeyReference *string `pulumi:"vaultKeyReference"`
}

type ConfigurationKeyState struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringPtrInput
	// The content type of the App Configuration Key. This should only be set when type is set to `kv`.
	ContentType pulumi.StringPtrInput
	// The ETag of the key.
	Etag pulumi.StringPtrInput
	// The name of the App Configuration Key to create. Changing this forces a new resource to be created.
	Key pulumi.StringPtrInput
	// The label of the App Configuration Key.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrInput
	// Should this App Configuration Key be Locked to prevent changes?
	Locked pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
	Type pulumi.StringPtrInput
	// The value of the App Configuration Key. This should only be set when type is set to `kv`.
	Value pulumi.StringPtrInput
	// The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
	VaultKeyReference pulumi.StringPtrInput
}

func (ConfigurationKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationKeyState)(nil)).Elem()
}

type configurationKeyArgs struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId string `pulumi:"configurationStoreId"`
	// The content type of the App Configuration Key. This should only be set when type is set to `kv`.
	ContentType *string `pulumi:"contentType"`
	// The ETag of the key.
	Etag *string `pulumi:"etag"`
	// The name of the App Configuration Key to create. Changing this forces a new resource to be created.
	Key string `pulumi:"key"`
	// The label of the App Configuration Key.  Changing this forces a new resource to be created.
	Label *string `pulumi:"label"`
	// Should this App Configuration Key be Locked to prevent changes?
	Locked *bool `pulumi:"locked"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
	Type *string `pulumi:"type"`
	// The value of the App Configuration Key. This should only be set when type is set to `kv`.
	Value *string `pulumi:"value"`
	// The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
	VaultKeyReference *string `pulumi:"vaultKeyReference"`
}

// The set of arguments for constructing a ConfigurationKey resource.
type ConfigurationKeyArgs struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringInput
	// The content type of the App Configuration Key. This should only be set when type is set to `kv`.
	ContentType pulumi.StringPtrInput
	// The ETag of the key.
	Etag pulumi.StringPtrInput
	// The name of the App Configuration Key to create. Changing this forces a new resource to be created.
	Key pulumi.StringInput
	// The label of the App Configuration Key.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrInput
	// Should this App Configuration Key be Locked to prevent changes?
	Locked pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
	Type pulumi.StringPtrInput
	// The value of the App Configuration Key. This should only be set when type is set to `kv`.
	Value pulumi.StringPtrInput
	// The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
	VaultKeyReference pulumi.StringPtrInput
}

func (ConfigurationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationKeyArgs)(nil)).Elem()
}

type ConfigurationKeyInput interface {
	pulumi.Input

	ToConfigurationKeyOutput() ConfigurationKeyOutput
	ToConfigurationKeyOutputWithContext(ctx context.Context) ConfigurationKeyOutput
}

func (*ConfigurationKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationKey)(nil)).Elem()
}

func (i *ConfigurationKey) ToConfigurationKeyOutput() ConfigurationKeyOutput {
	return i.ToConfigurationKeyOutputWithContext(context.Background())
}

func (i *ConfigurationKey) ToConfigurationKeyOutputWithContext(ctx context.Context) ConfigurationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationKeyOutput)
}

// ConfigurationKeyArrayInput is an input type that accepts ConfigurationKeyArray and ConfigurationKeyArrayOutput values.
// You can construct a concrete instance of `ConfigurationKeyArrayInput` via:
//
//          ConfigurationKeyArray{ ConfigurationKeyArgs{...} }
type ConfigurationKeyArrayInput interface {
	pulumi.Input

	ToConfigurationKeyArrayOutput() ConfigurationKeyArrayOutput
	ToConfigurationKeyArrayOutputWithContext(context.Context) ConfigurationKeyArrayOutput
}

type ConfigurationKeyArray []ConfigurationKeyInput

func (ConfigurationKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigurationKey)(nil)).Elem()
}

func (i ConfigurationKeyArray) ToConfigurationKeyArrayOutput() ConfigurationKeyArrayOutput {
	return i.ToConfigurationKeyArrayOutputWithContext(context.Background())
}

func (i ConfigurationKeyArray) ToConfigurationKeyArrayOutputWithContext(ctx context.Context) ConfigurationKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationKeyArrayOutput)
}

// ConfigurationKeyMapInput is an input type that accepts ConfigurationKeyMap and ConfigurationKeyMapOutput values.
// You can construct a concrete instance of `ConfigurationKeyMapInput` via:
//
//          ConfigurationKeyMap{ "key": ConfigurationKeyArgs{...} }
type ConfigurationKeyMapInput interface {
	pulumi.Input

	ToConfigurationKeyMapOutput() ConfigurationKeyMapOutput
	ToConfigurationKeyMapOutputWithContext(context.Context) ConfigurationKeyMapOutput
}

type ConfigurationKeyMap map[string]ConfigurationKeyInput

func (ConfigurationKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigurationKey)(nil)).Elem()
}

func (i ConfigurationKeyMap) ToConfigurationKeyMapOutput() ConfigurationKeyMapOutput {
	return i.ToConfigurationKeyMapOutputWithContext(context.Background())
}

func (i ConfigurationKeyMap) ToConfigurationKeyMapOutputWithContext(ctx context.Context) ConfigurationKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationKeyMapOutput)
}

type ConfigurationKeyOutput struct{ *pulumi.OutputState }

func (ConfigurationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationKey)(nil)).Elem()
}

func (o ConfigurationKeyOutput) ToConfigurationKeyOutput() ConfigurationKeyOutput {
	return o
}

func (o ConfigurationKeyOutput) ToConfigurationKeyOutputWithContext(ctx context.Context) ConfigurationKeyOutput {
	return o
}

type ConfigurationKeyArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigurationKey)(nil)).Elem()
}

func (o ConfigurationKeyArrayOutput) ToConfigurationKeyArrayOutput() ConfigurationKeyArrayOutput {
	return o
}

func (o ConfigurationKeyArrayOutput) ToConfigurationKeyArrayOutputWithContext(ctx context.Context) ConfigurationKeyArrayOutput {
	return o
}

func (o ConfigurationKeyArrayOutput) Index(i pulumi.IntInput) ConfigurationKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigurationKey {
		return vs[0].([]*ConfigurationKey)[vs[1].(int)]
	}).(ConfigurationKeyOutput)
}

type ConfigurationKeyMapOutput struct{ *pulumi.OutputState }

func (ConfigurationKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigurationKey)(nil)).Elem()
}

func (o ConfigurationKeyMapOutput) ToConfigurationKeyMapOutput() ConfigurationKeyMapOutput {
	return o
}

func (o ConfigurationKeyMapOutput) ToConfigurationKeyMapOutputWithContext(ctx context.Context) ConfigurationKeyMapOutput {
	return o
}

func (o ConfigurationKeyMapOutput) MapIndex(k pulumi.StringInput) ConfigurationKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigurationKey {
		return vs[0].(map[string]*ConfigurationKey)[vs[1].(string)]
	}).(ConfigurationKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationKeyInput)(nil)).Elem(), &ConfigurationKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationKeyArrayInput)(nil)).Elem(), ConfigurationKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationKeyMapInput)(nil)).Elem(), ConfigurationKeyMap{})
	pulumi.RegisterOutputType(ConfigurationKeyOutput{})
	pulumi.RegisterOutputType(ConfigurationKeyArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationKeyMapOutput{})
}
