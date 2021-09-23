// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Backup Protected File Share to enable backups for file shares within an Azure Storage Account
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/backup"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		sa, err := storage.NewAccount(ctx, "sa", &storage.AccountArgs{
// 			Location:               rg.Location,
// 			ResourceGroupName:      rg.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleShare, err := storage.NewShare(ctx, "exampleShare", &storage.ShareArgs{
// 			StorageAccountName: sa.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewContainerStorageAccount(ctx, "protection_container", &backup.ContainerStorageAccountArgs{
// 			ResourceGroupName: rg.Name,
// 			RecoveryVaultName: vault.Name,
// 			StorageAccountId:  sa.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePolicyFileShare, err := backup.NewPolicyFileShare(ctx, "examplePolicyFileShare", &backup.PolicyFileShareArgs{
// 			ResourceGroupName: rg.Name,
// 			RecoveryVaultName: vault.Name,
// 			Backup: &backup.PolicyFileShareBackupArgs{
// 				Frequency: pulumi.String("Daily"),
// 				Time:      pulumi.String("23:00"),
// 			},
// 			RetentionDaily: &backup.PolicyFileShareRetentionDailyArgs{
// 				Count: pulumi.Int(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewProtectedFileShare(ctx, "share1", &backup.ProtectedFileShareArgs{
// 			ResourceGroupName:      rg.Name,
// 			RecoveryVaultName:      vault.Name,
// 			SourceStorageAccountId: protection_container.StorageAccountId,
// 			SourceFileShareName:    exampleShare.Name,
// 			BackupPolicyId:         examplePolicyFileShare.ID(),
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
// Azure Backup Protected File Shares can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:backup/protectedFileShare:ProtectedFileShare item1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/StorageContainer;storage;group2;example-storage-account/protectedItems/AzureFileShare;3f6e3108a45793581bcbd1c61c87a3b2ceeb4ff4bc02a95ce9d1022b23722935"
// ```
//
//  -> **NOTE** The ID requires quoting as there are semicolons. This user unfriendly ID can be found in the Deployments of the used resourcegroup, look for an Deployment which starts with `ConfigureAFSProtection-`, click then `Go to resource`.
type ProtectedFileShare struct {
	pulumi.CustomResourceState

	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName pulumi.StringOutput `pulumi:"sourceFileShareName"`
	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountId pulumi.StringOutput `pulumi:"sourceStorageAccountId"`
}

// NewProtectedFileShare registers a new resource with the given unique name, arguments, and options.
func NewProtectedFileShare(ctx *pulumi.Context,
	name string, args *ProtectedFileShareArgs, opts ...pulumi.ResourceOption) (*ProtectedFileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyId'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceFileShareName == nil {
		return nil, errors.New("invalid value for required argument 'SourceFileShareName'")
	}
	if args.SourceStorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'SourceStorageAccountId'")
	}
	var resource ProtectedFileShare
	err := ctx.RegisterResource("azure:backup/protectedFileShare:ProtectedFileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectedFileShare gets an existing ProtectedFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectedFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedFileShareState, opts ...pulumi.ResourceOption) (*ProtectedFileShare, error) {
	var resource ProtectedFileShare
	err := ctx.ReadResource("azure:backup/protectedFileShare:ProtectedFileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectedFileShare resources.
type protectedFileShareState struct {
	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName *string `pulumi:"sourceFileShareName"`
	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountId *string `pulumi:"sourceStorageAccountId"`
}

type ProtectedFileShareState struct {
	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyId pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName pulumi.StringPtrInput
	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountId pulumi.StringPtrInput
}

func (ProtectedFileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedFileShareState)(nil)).Elem()
}

type protectedFileShareArgs struct {
	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyId string `pulumi:"backupPolicyId"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName string `pulumi:"sourceFileShareName"`
	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountId string `pulumi:"sourceStorageAccountId"`
}

// The set of arguments for constructing a ProtectedFileShare resource.
type ProtectedFileShareArgs struct {
	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyId pulumi.StringInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringInput
	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName pulumi.StringInput
	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountId pulumi.StringInput
}

func (ProtectedFileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedFileShareArgs)(nil)).Elem()
}

type ProtectedFileShareInput interface {
	pulumi.Input

	ToProtectedFileShareOutput() ProtectedFileShareOutput
	ToProtectedFileShareOutputWithContext(ctx context.Context) ProtectedFileShareOutput
}

func (*ProtectedFileShare) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedFileShare)(nil))
}

func (i *ProtectedFileShare) ToProtectedFileShareOutput() ProtectedFileShareOutput {
	return i.ToProtectedFileShareOutputWithContext(context.Background())
}

func (i *ProtectedFileShare) ToProtectedFileShareOutputWithContext(ctx context.Context) ProtectedFileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedFileShareOutput)
}

func (i *ProtectedFileShare) ToProtectedFileSharePtrOutput() ProtectedFileSharePtrOutput {
	return i.ToProtectedFileSharePtrOutputWithContext(context.Background())
}

func (i *ProtectedFileShare) ToProtectedFileSharePtrOutputWithContext(ctx context.Context) ProtectedFileSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedFileSharePtrOutput)
}

type ProtectedFileSharePtrInput interface {
	pulumi.Input

	ToProtectedFileSharePtrOutput() ProtectedFileSharePtrOutput
	ToProtectedFileSharePtrOutputWithContext(ctx context.Context) ProtectedFileSharePtrOutput
}

type protectedFileSharePtrType ProtectedFileShareArgs

func (*protectedFileSharePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedFileShare)(nil))
}

func (i *protectedFileSharePtrType) ToProtectedFileSharePtrOutput() ProtectedFileSharePtrOutput {
	return i.ToProtectedFileSharePtrOutputWithContext(context.Background())
}

func (i *protectedFileSharePtrType) ToProtectedFileSharePtrOutputWithContext(ctx context.Context) ProtectedFileSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedFileSharePtrOutput)
}

// ProtectedFileShareArrayInput is an input type that accepts ProtectedFileShareArray and ProtectedFileShareArrayOutput values.
// You can construct a concrete instance of `ProtectedFileShareArrayInput` via:
//
//          ProtectedFileShareArray{ ProtectedFileShareArgs{...} }
type ProtectedFileShareArrayInput interface {
	pulumi.Input

	ToProtectedFileShareArrayOutput() ProtectedFileShareArrayOutput
	ToProtectedFileShareArrayOutputWithContext(context.Context) ProtectedFileShareArrayOutput
}

type ProtectedFileShareArray []ProtectedFileShareInput

func (ProtectedFileShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProtectedFileShare)(nil)).Elem()
}

func (i ProtectedFileShareArray) ToProtectedFileShareArrayOutput() ProtectedFileShareArrayOutput {
	return i.ToProtectedFileShareArrayOutputWithContext(context.Background())
}

func (i ProtectedFileShareArray) ToProtectedFileShareArrayOutputWithContext(ctx context.Context) ProtectedFileShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedFileShareArrayOutput)
}

// ProtectedFileShareMapInput is an input type that accepts ProtectedFileShareMap and ProtectedFileShareMapOutput values.
// You can construct a concrete instance of `ProtectedFileShareMapInput` via:
//
//          ProtectedFileShareMap{ "key": ProtectedFileShareArgs{...} }
type ProtectedFileShareMapInput interface {
	pulumi.Input

	ToProtectedFileShareMapOutput() ProtectedFileShareMapOutput
	ToProtectedFileShareMapOutputWithContext(context.Context) ProtectedFileShareMapOutput
}

type ProtectedFileShareMap map[string]ProtectedFileShareInput

func (ProtectedFileShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProtectedFileShare)(nil)).Elem()
}

func (i ProtectedFileShareMap) ToProtectedFileShareMapOutput() ProtectedFileShareMapOutput {
	return i.ToProtectedFileShareMapOutputWithContext(context.Background())
}

func (i ProtectedFileShareMap) ToProtectedFileShareMapOutputWithContext(ctx context.Context) ProtectedFileShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedFileShareMapOutput)
}

type ProtectedFileShareOutput struct{ *pulumi.OutputState }

func (ProtectedFileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedFileShare)(nil))
}

func (o ProtectedFileShareOutput) ToProtectedFileShareOutput() ProtectedFileShareOutput {
	return o
}

func (o ProtectedFileShareOutput) ToProtectedFileShareOutputWithContext(ctx context.Context) ProtectedFileShareOutput {
	return o
}

func (o ProtectedFileShareOutput) ToProtectedFileSharePtrOutput() ProtectedFileSharePtrOutput {
	return o.ToProtectedFileSharePtrOutputWithContext(context.Background())
}

func (o ProtectedFileShareOutput) ToProtectedFileSharePtrOutputWithContext(ctx context.Context) ProtectedFileSharePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectedFileShare) *ProtectedFileShare {
		return &v
	}).(ProtectedFileSharePtrOutput)
}

type ProtectedFileSharePtrOutput struct{ *pulumi.OutputState }

func (ProtectedFileSharePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedFileShare)(nil))
}

func (o ProtectedFileSharePtrOutput) ToProtectedFileSharePtrOutput() ProtectedFileSharePtrOutput {
	return o
}

func (o ProtectedFileSharePtrOutput) ToProtectedFileSharePtrOutputWithContext(ctx context.Context) ProtectedFileSharePtrOutput {
	return o
}

func (o ProtectedFileSharePtrOutput) Elem() ProtectedFileShareOutput {
	return o.ApplyT(func(v *ProtectedFileShare) ProtectedFileShare {
		if v != nil {
			return *v
		}
		var ret ProtectedFileShare
		return ret
	}).(ProtectedFileShareOutput)
}

type ProtectedFileShareArrayOutput struct{ *pulumi.OutputState }

func (ProtectedFileShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProtectedFileShare)(nil))
}

func (o ProtectedFileShareArrayOutput) ToProtectedFileShareArrayOutput() ProtectedFileShareArrayOutput {
	return o
}

func (o ProtectedFileShareArrayOutput) ToProtectedFileShareArrayOutputWithContext(ctx context.Context) ProtectedFileShareArrayOutput {
	return o
}

func (o ProtectedFileShareArrayOutput) Index(i pulumi.IntInput) ProtectedFileShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProtectedFileShare {
		return vs[0].([]ProtectedFileShare)[vs[1].(int)]
	}).(ProtectedFileShareOutput)
}

type ProtectedFileShareMapOutput struct{ *pulumi.OutputState }

func (ProtectedFileShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProtectedFileShare)(nil))
}

func (o ProtectedFileShareMapOutput) ToProtectedFileShareMapOutput() ProtectedFileShareMapOutput {
	return o
}

func (o ProtectedFileShareMapOutput) ToProtectedFileShareMapOutputWithContext(ctx context.Context) ProtectedFileShareMapOutput {
	return o
}

func (o ProtectedFileShareMapOutput) MapIndex(k pulumi.StringInput) ProtectedFileShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProtectedFileShare {
		return vs[0].(map[string]ProtectedFileShare)[vs[1].(string)]
	}).(ProtectedFileShareOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectedFileShareOutput{})
	pulumi.RegisterOutputType(ProtectedFileSharePtrOutput{})
	pulumi.RegisterOutputType(ProtectedFileShareArrayOutput{})
	pulumi.RegisterOutputType(ProtectedFileShareMapOutput{})
}
