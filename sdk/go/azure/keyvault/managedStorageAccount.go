// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Key Vault Managed Storage Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
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
// 		_ = storage.GetAccountSASOutput(ctx, storage.GetAccountSASOutputArgs{
// 			ConnectionString: exampleAccount.PrimaryConnectionString,
// 			HttpsOnly:        pulumi.Bool(true),
// 			ResourceTypes: &storage.GetAccountSASResourceTypesArgs{
// 				Service:   pulumi.Bool(true),
// 				Container: pulumi.Bool(false),
// 				Object:    pulumi.Bool(false),
// 			},
// 			Services: &storage.GetAccountSASServicesArgs{
// 				Blob:  pulumi.Bool(true),
// 				Queue: pulumi.Bool(false),
// 				Table: pulumi.Bool(false),
// 				File:  pulumi.Bool(false),
// 			},
// 			Start:  pulumi.String("2021-04-30T00:00:00Z"),
// 			Expiry: pulumi.String("2023-04-30T00:00:00Z"),
// 			Permissions: &storage.GetAccountSASPermissionsArgs{
// 				Read:    pulumi.Bool(true),
// 				Write:   pulumi.Bool(true),
// 				Delete:  pulumi.Bool(false),
// 				List:    pulumi.Bool(false),
// 				Add:     pulumi.Bool(true),
// 				Create:  pulumi.Bool(true),
// 				Update:  pulumi.Bool(false),
// 				Process: pulumi.Bool(false),
// 			},
// 		}, nil)
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.Any(data.Azurerm_client_config.Current.Tenant_id),
// 			SkuName:           pulumi.String("standard"),
// 			AccessPolicies: keyvault.KeyVaultAccessPolicyArray{
// 				&keyvault.KeyVaultAccessPolicyArgs{
// 					TenantId: pulumi.Any(data.Azurerm_client_config.Current.Tenant_id),
// 					ObjectId: pulumi.Any(data.Azurerm_client_config.Current.Object_id),
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("Get"),
// 						pulumi.String("Delete"),
// 					},
// 					StoragePermissions: pulumi.StringArray{
// 						pulumi.String("Get"),
// 						pulumi.String("List"),
// 						pulumi.String("Set"),
// 						pulumi.String("SetSAS"),
// 						pulumi.String("GetSAS"),
// 						pulumi.String("DeleteSAS"),
// 						pulumi.String("Update"),
// 						pulumi.String("RegenerateKey"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewManagedStorageAccount(ctx, "exampleManagedStorageAccount", &keyvault.ManagedStorageAccountArgs{
// 			KeyVaultId:                 exampleKeyVault.ID(),
// 			StorageAccountId:           exampleAccount.ID(),
// 			StorageAccountKey:          pulumi.String("key1"),
// 			RegenerateKeyAutomatically: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Automatically Regenerate Storage Account Access Key)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
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
// 		_ = storage.GetAccountSASOutput(ctx, storage.GetAccountSASOutputArgs{
// 			ConnectionString: exampleAccount.PrimaryConnectionString,
// 			HttpsOnly:        pulumi.Bool(true),
// 			ResourceTypes: &storage.GetAccountSASResourceTypesArgs{
// 				Service:   pulumi.Bool(true),
// 				Container: pulumi.Bool(false),
// 				Object:    pulumi.Bool(false),
// 			},
// 			Services: &storage.GetAccountSASServicesArgs{
// 				Blob:  pulumi.Bool(true),
// 				Queue: pulumi.Bool(false),
// 				Table: pulumi.Bool(false),
// 				File:  pulumi.Bool(false),
// 			},
// 			Start:  pulumi.String("2021-04-30T00:00:00Z"),
// 			Expiry: pulumi.String("2023-04-30T00:00:00Z"),
// 			Permissions: &storage.GetAccountSASPermissionsArgs{
// 				Read:    pulumi.Bool(true),
// 				Write:   pulumi.Bool(true),
// 				Delete:  pulumi.Bool(false),
// 				List:    pulumi.Bool(false),
// 				Add:     pulumi.Bool(true),
// 				Create:  pulumi.Bool(true),
// 				Update:  pulumi.Bool(false),
// 				Process: pulumi.Bool(false),
// 			},
// 		}, nil)
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.Any(data.Azurerm_client_config.Current.Tenant_id),
// 			SkuName:           pulumi.String("standard"),
// 			AccessPolicies: keyvault.KeyVaultAccessPolicyArray{
// 				&keyvault.KeyVaultAccessPolicyArgs{
// 					TenantId: pulumi.Any(data.Azurerm_client_config.Current.Tenant_id),
// 					ObjectId: pulumi.Any(data.Azurerm_client_config.Current.Object_id),
// 					SecretPermissions: pulumi.StringArray{
// 						pulumi.String("Get"),
// 						pulumi.String("Delete"),
// 					},
// 					StoragePermissions: pulumi.StringArray{
// 						pulumi.String("Get"),
// 						pulumi.String("List"),
// 						pulumi.String("Set"),
// 						pulumi.String("SetSAS"),
// 						pulumi.String("GetSAS"),
// 						pulumi.String("DeleteSAS"),
// 						pulumi.String("Update"),
// 						pulumi.String("RegenerateKey"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Storage Account Key Operator Service Role"),
// 			PrincipalId:        pulumi.String("727055f9-0386-4ccb-bcf1-9237237ee102"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewManagedStorageAccount(ctx, "exampleManagedStorageAccount", &keyvault.ManagedStorageAccountArgs{
// 			KeyVaultId:                 exampleKeyVault.ID(),
// 			StorageAccountId:           exampleAccount.ID(),
// 			StorageAccountKey:          pulumi.String("key1"),
// 			RegenerateKeyAutomatically: pulumi.Bool(true),
// 			RegenerationPeriod:         pulumi.String("P1D"),
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
// Key Vault Managed Storage Accounts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:keyvault/managedStorageAccount:ManagedStorageAccount example https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01
// ```
type ManagedStorageAccount struct {
	pulumi.CustomResourceState

	// The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Should Storage Account access key be regenerated periodically?
	RegenerateKeyAutomatically pulumi.BoolPtrOutput `pulumi:"regenerateKeyAutomatically"`
	// How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	RegenerationPeriod pulumi.StringPtrOutput `pulumi:"regenerationPeriod"`
	// The ID of the Storage Account.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.
	StorageAccountKey pulumi.StringOutput `pulumi:"storageAccountKey"`
	// A mapping of tags which should be assigned to the Key Vault Managed Storage Account.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewManagedStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewManagedStorageAccount(ctx *pulumi.Context,
	name string, args *ManagedStorageAccountArgs, opts ...pulumi.ResourceOption) (*ManagedStorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	if args.StorageAccountKey == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountKey'")
	}
	var resource ManagedStorageAccount
	err := ctx.RegisterResource("azure:keyvault/managedStorageAccount:ManagedStorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedStorageAccount gets an existing ManagedStorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedStorageAccountState, opts ...pulumi.ResourceOption) (*ManagedStorageAccount, error) {
	var resource ManagedStorageAccount
	err := ctx.ReadResource("azure:keyvault/managedStorageAccount:ManagedStorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedStorageAccount resources.
type managedStorageAccountState struct {
	// The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.
	Name *string `pulumi:"name"`
	// Should Storage Account access key be regenerated periodically?
	RegenerateKeyAutomatically *bool `pulumi:"regenerateKeyAutomatically"`
	// How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	RegenerationPeriod *string `pulumi:"regenerationPeriod"`
	// The ID of the Storage Account.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// A mapping of tags which should be assigned to the Key Vault Managed Storage Account.
	Tags map[string]string `pulumi:"tags"`
}

type ManagedStorageAccountState struct {
	// The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.
	Name pulumi.StringPtrInput
	// Should Storage Account access key be regenerated periodically?
	RegenerateKeyAutomatically pulumi.BoolPtrInput
	// How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	RegenerationPeriod pulumi.StringPtrInput
	// The ID of the Storage Account.
	StorageAccountId pulumi.StringPtrInput
	// Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.
	StorageAccountKey pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Key Vault Managed Storage Account.
	Tags pulumi.StringMapInput
}

func (ManagedStorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedStorageAccountState)(nil)).Elem()
}

type managedStorageAccountArgs struct {
	// The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.
	Name *string `pulumi:"name"`
	// Should Storage Account access key be regenerated periodically?
	RegenerateKeyAutomatically *bool `pulumi:"regenerateKeyAutomatically"`
	// How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	RegenerationPeriod *string `pulumi:"regenerationPeriod"`
	// The ID of the Storage Account.
	StorageAccountId string `pulumi:"storageAccountId"`
	// Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.
	StorageAccountKey string `pulumi:"storageAccountKey"`
	// A mapping of tags which should be assigned to the Key Vault Managed Storage Account.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedStorageAccount resource.
type ManagedStorageAccountArgs struct {
	// The ID of the Key Vault where the Managed Storage Account should be created. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// The name which should be used for this Key Vault Managed Storage Account. Changing this forces a new Key Vault Managed Storage Account to be created.
	Name pulumi.StringPtrInput
	// Should Storage Account access key be regenerated periodically?
	RegenerateKeyAutomatically pulumi.BoolPtrInput
	// How often Storage Account access key should be regenerated. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	RegenerationPeriod pulumi.StringPtrInput
	// The ID of the Storage Account.
	StorageAccountId pulumi.StringInput
	// Which Storage Account access key that is managed by Key Vault. Possible values are `key1` and `key2`.
	StorageAccountKey pulumi.StringInput
	// A mapping of tags which should be assigned to the Key Vault Managed Storage Account.
	Tags pulumi.StringMapInput
}

func (ManagedStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedStorageAccountArgs)(nil)).Elem()
}

type ManagedStorageAccountInput interface {
	pulumi.Input

	ToManagedStorageAccountOutput() ManagedStorageAccountOutput
	ToManagedStorageAccountOutputWithContext(ctx context.Context) ManagedStorageAccountOutput
}

func (*ManagedStorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedStorageAccount)(nil)).Elem()
}

func (i *ManagedStorageAccount) ToManagedStorageAccountOutput() ManagedStorageAccountOutput {
	return i.ToManagedStorageAccountOutputWithContext(context.Background())
}

func (i *ManagedStorageAccount) ToManagedStorageAccountOutputWithContext(ctx context.Context) ManagedStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedStorageAccountOutput)
}

// ManagedStorageAccountArrayInput is an input type that accepts ManagedStorageAccountArray and ManagedStorageAccountArrayOutput values.
// You can construct a concrete instance of `ManagedStorageAccountArrayInput` via:
//
//          ManagedStorageAccountArray{ ManagedStorageAccountArgs{...} }
type ManagedStorageAccountArrayInput interface {
	pulumi.Input

	ToManagedStorageAccountArrayOutput() ManagedStorageAccountArrayOutput
	ToManagedStorageAccountArrayOutputWithContext(context.Context) ManagedStorageAccountArrayOutput
}

type ManagedStorageAccountArray []ManagedStorageAccountInput

func (ManagedStorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedStorageAccount)(nil)).Elem()
}

func (i ManagedStorageAccountArray) ToManagedStorageAccountArrayOutput() ManagedStorageAccountArrayOutput {
	return i.ToManagedStorageAccountArrayOutputWithContext(context.Background())
}

func (i ManagedStorageAccountArray) ToManagedStorageAccountArrayOutputWithContext(ctx context.Context) ManagedStorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedStorageAccountArrayOutput)
}

// ManagedStorageAccountMapInput is an input type that accepts ManagedStorageAccountMap and ManagedStorageAccountMapOutput values.
// You can construct a concrete instance of `ManagedStorageAccountMapInput` via:
//
//          ManagedStorageAccountMap{ "key": ManagedStorageAccountArgs{...} }
type ManagedStorageAccountMapInput interface {
	pulumi.Input

	ToManagedStorageAccountMapOutput() ManagedStorageAccountMapOutput
	ToManagedStorageAccountMapOutputWithContext(context.Context) ManagedStorageAccountMapOutput
}

type ManagedStorageAccountMap map[string]ManagedStorageAccountInput

func (ManagedStorageAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedStorageAccount)(nil)).Elem()
}

func (i ManagedStorageAccountMap) ToManagedStorageAccountMapOutput() ManagedStorageAccountMapOutput {
	return i.ToManagedStorageAccountMapOutputWithContext(context.Background())
}

func (i ManagedStorageAccountMap) ToManagedStorageAccountMapOutputWithContext(ctx context.Context) ManagedStorageAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedStorageAccountMapOutput)
}

type ManagedStorageAccountOutput struct{ *pulumi.OutputState }

func (ManagedStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedStorageAccount)(nil)).Elem()
}

func (o ManagedStorageAccountOutput) ToManagedStorageAccountOutput() ManagedStorageAccountOutput {
	return o
}

func (o ManagedStorageAccountOutput) ToManagedStorageAccountOutputWithContext(ctx context.Context) ManagedStorageAccountOutput {
	return o
}

type ManagedStorageAccountArrayOutput struct{ *pulumi.OutputState }

func (ManagedStorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedStorageAccount)(nil)).Elem()
}

func (o ManagedStorageAccountArrayOutput) ToManagedStorageAccountArrayOutput() ManagedStorageAccountArrayOutput {
	return o
}

func (o ManagedStorageAccountArrayOutput) ToManagedStorageAccountArrayOutputWithContext(ctx context.Context) ManagedStorageAccountArrayOutput {
	return o
}

func (o ManagedStorageAccountArrayOutput) Index(i pulumi.IntInput) ManagedStorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedStorageAccount {
		return vs[0].([]*ManagedStorageAccount)[vs[1].(int)]
	}).(ManagedStorageAccountOutput)
}

type ManagedStorageAccountMapOutput struct{ *pulumi.OutputState }

func (ManagedStorageAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedStorageAccount)(nil)).Elem()
}

func (o ManagedStorageAccountMapOutput) ToManagedStorageAccountMapOutput() ManagedStorageAccountMapOutput {
	return o
}

func (o ManagedStorageAccountMapOutput) ToManagedStorageAccountMapOutputWithContext(ctx context.Context) ManagedStorageAccountMapOutput {
	return o
}

func (o ManagedStorageAccountMapOutput) MapIndex(k pulumi.StringInput) ManagedStorageAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedStorageAccount {
		return vs[0].(map[string]*ManagedStorageAccount)[vs[1].(string)]
	}).(ManagedStorageAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedStorageAccountInput)(nil)).Elem(), &ManagedStorageAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedStorageAccountArrayInput)(nil)).Elem(), ManagedStorageAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedStorageAccountMapInput)(nil)).Elem(), ManagedStorageAccountMap{})
	pulumi.RegisterOutputType(ManagedStorageAccountOutput{})
	pulumi.RegisterOutputType(ManagedStorageAccountArrayOutput{})
	pulumi.RegisterOutputType(ManagedStorageAccountMapOutput{})
}
