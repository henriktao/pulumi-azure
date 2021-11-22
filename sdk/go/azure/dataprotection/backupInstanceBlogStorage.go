// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Backup Instance Blob Storage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/dataprotection"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:               pulumi.Any(azurerm_resource_group.Example.Location),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBackupVault, err := dataprotection.NewBackupVault(ctx, "exampleBackupVault", &dataprotection.BackupVaultArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			DatastoreType:     pulumi.String("VaultStore"),
// 			Redundancy:        pulumi.String("LocallyRedundant"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Storage Account Backup Contributor Role"),
// 			PrincipalId: exampleBackupVault.Identity.ApplyT(func(identity dataprotection.BackupVaultIdentity) (string, error) {
// 				return identity.PrincipalId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBackupPolicyBlobStorage, err := dataprotection.NewBackupPolicyBlobStorage(ctx, "exampleBackupPolicyBlobStorage", &dataprotection.BackupPolicyBlobStorageArgs{
// 			VaultId:           exampleBackupVault.ID(),
// 			RetentionDuration: pulumi.String("P30D"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataprotection.NewBackupInstanceBlogStorage(ctx, "exampleBackupInstanceBlogStorage", &dataprotection.BackupInstanceBlogStorageArgs{
// 			VaultId:          exampleBackupVault.ID(),
// 			Location:         rg.Location,
// 			StorageAccountId: exampleAccount.ID(),
// 			BackupPolicyId:   exampleBackupPolicyBlobStorage.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAssignment,
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
// Backup Instance Blob Storages can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dataprotection/backupInstanceBlogStorage:BackupInstanceBlogStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
// ```
type BackupInstanceBlogStorage struct {
	pulumi.CustomResourceState

	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	Location       pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Backup Instance Blob Storage. Changing this forces a new Backup Instance Blob Storage to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the source Storage Account. Changing this forces a new Backup Instance Blob Storage to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// The ID of the Backup Vault within which the Backup Instance Blob Storage should exist. Changing this forces a new Backup Instance Blob Storage to be created.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewBackupInstanceBlogStorage registers a new resource with the given unique name, arguments, and options.
func NewBackupInstanceBlogStorage(ctx *pulumi.Context,
	name string, args *BackupInstanceBlogStorageArgs, opts ...pulumi.ResourceOption) (*BackupInstanceBlogStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyId'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	var resource BackupInstanceBlogStorage
	err := ctx.RegisterResource("azure:dataprotection/backupInstanceBlogStorage:BackupInstanceBlogStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupInstanceBlogStorage gets an existing BackupInstanceBlogStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupInstanceBlogStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupInstanceBlogStorageState, opts ...pulumi.ResourceOption) (*BackupInstanceBlogStorage, error) {
	var resource BackupInstanceBlogStorage
	err := ctx.ReadResource("azure:dataprotection/backupInstanceBlogStorage:BackupInstanceBlogStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupInstanceBlogStorage resources.
type backupInstanceBlogStorageState struct {
	// The ID of the Backup Policy.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	Location       *string `pulumi:"location"`
	// The name which should be used for this Backup Instance Blob Storage. Changing this forces a new Backup Instance Blob Storage to be created.
	Name *string `pulumi:"name"`
	// The ID of the source Storage Account. Changing this forces a new Backup Instance Blob Storage to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The ID of the Backup Vault within which the Backup Instance Blob Storage should exist. Changing this forces a new Backup Instance Blob Storage to be created.
	VaultId *string `pulumi:"vaultId"`
}

type BackupInstanceBlogStorageState struct {
	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	// The name which should be used for this Backup Instance Blob Storage. Changing this forces a new Backup Instance Blob Storage to be created.
	Name pulumi.StringPtrInput
	// The ID of the source Storage Account. Changing this forces a new Backup Instance Blob Storage to be created.
	StorageAccountId pulumi.StringPtrInput
	// The ID of the Backup Vault within which the Backup Instance Blob Storage should exist. Changing this forces a new Backup Instance Blob Storage to be created.
	VaultId pulumi.StringPtrInput
}

func (BackupInstanceBlogStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceBlogStorageState)(nil)).Elem()
}

type backupInstanceBlogStorageArgs struct {
	// The ID of the Backup Policy.
	BackupPolicyId string  `pulumi:"backupPolicyId"`
	Location       *string `pulumi:"location"`
	// The name which should be used for this Backup Instance Blob Storage. Changing this forces a new Backup Instance Blob Storage to be created.
	Name *string `pulumi:"name"`
	// The ID of the source Storage Account. Changing this forces a new Backup Instance Blob Storage to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// The ID of the Backup Vault within which the Backup Instance Blob Storage should exist. Changing this forces a new Backup Instance Blob Storage to be created.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a BackupInstanceBlogStorage resource.
type BackupInstanceBlogStorageArgs struct {
	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringInput
	Location       pulumi.StringPtrInput
	// The name which should be used for this Backup Instance Blob Storage. Changing this forces a new Backup Instance Blob Storage to be created.
	Name pulumi.StringPtrInput
	// The ID of the source Storage Account. Changing this forces a new Backup Instance Blob Storage to be created.
	StorageAccountId pulumi.StringInput
	// The ID of the Backup Vault within which the Backup Instance Blob Storage should exist. Changing this forces a new Backup Instance Blob Storage to be created.
	VaultId pulumi.StringInput
}

func (BackupInstanceBlogStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceBlogStorageArgs)(nil)).Elem()
}

type BackupInstanceBlogStorageInput interface {
	pulumi.Input

	ToBackupInstanceBlogStorageOutput() BackupInstanceBlogStorageOutput
	ToBackupInstanceBlogStorageOutputWithContext(ctx context.Context) BackupInstanceBlogStorageOutput
}

func (*BackupInstanceBlogStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceBlogStorage)(nil))
}

func (i *BackupInstanceBlogStorage) ToBackupInstanceBlogStorageOutput() BackupInstanceBlogStorageOutput {
	return i.ToBackupInstanceBlogStorageOutputWithContext(context.Background())
}

func (i *BackupInstanceBlogStorage) ToBackupInstanceBlogStorageOutputWithContext(ctx context.Context) BackupInstanceBlogStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceBlogStorageOutput)
}

func (i *BackupInstanceBlogStorage) ToBackupInstanceBlogStoragePtrOutput() BackupInstanceBlogStoragePtrOutput {
	return i.ToBackupInstanceBlogStoragePtrOutputWithContext(context.Background())
}

func (i *BackupInstanceBlogStorage) ToBackupInstanceBlogStoragePtrOutputWithContext(ctx context.Context) BackupInstanceBlogStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceBlogStoragePtrOutput)
}

type BackupInstanceBlogStoragePtrInput interface {
	pulumi.Input

	ToBackupInstanceBlogStoragePtrOutput() BackupInstanceBlogStoragePtrOutput
	ToBackupInstanceBlogStoragePtrOutputWithContext(ctx context.Context) BackupInstanceBlogStoragePtrOutput
}

type backupInstanceBlogStoragePtrType BackupInstanceBlogStorageArgs

func (*backupInstanceBlogStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceBlogStorage)(nil))
}

func (i *backupInstanceBlogStoragePtrType) ToBackupInstanceBlogStoragePtrOutput() BackupInstanceBlogStoragePtrOutput {
	return i.ToBackupInstanceBlogStoragePtrOutputWithContext(context.Background())
}

func (i *backupInstanceBlogStoragePtrType) ToBackupInstanceBlogStoragePtrOutputWithContext(ctx context.Context) BackupInstanceBlogStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceBlogStoragePtrOutput)
}

// BackupInstanceBlogStorageArrayInput is an input type that accepts BackupInstanceBlogStorageArray and BackupInstanceBlogStorageArrayOutput values.
// You can construct a concrete instance of `BackupInstanceBlogStorageArrayInput` via:
//
//          BackupInstanceBlogStorageArray{ BackupInstanceBlogStorageArgs{...} }
type BackupInstanceBlogStorageArrayInput interface {
	pulumi.Input

	ToBackupInstanceBlogStorageArrayOutput() BackupInstanceBlogStorageArrayOutput
	ToBackupInstanceBlogStorageArrayOutputWithContext(context.Context) BackupInstanceBlogStorageArrayOutput
}

type BackupInstanceBlogStorageArray []BackupInstanceBlogStorageInput

func (BackupInstanceBlogStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupInstanceBlogStorage)(nil)).Elem()
}

func (i BackupInstanceBlogStorageArray) ToBackupInstanceBlogStorageArrayOutput() BackupInstanceBlogStorageArrayOutput {
	return i.ToBackupInstanceBlogStorageArrayOutputWithContext(context.Background())
}

func (i BackupInstanceBlogStorageArray) ToBackupInstanceBlogStorageArrayOutputWithContext(ctx context.Context) BackupInstanceBlogStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceBlogStorageArrayOutput)
}

// BackupInstanceBlogStorageMapInput is an input type that accepts BackupInstanceBlogStorageMap and BackupInstanceBlogStorageMapOutput values.
// You can construct a concrete instance of `BackupInstanceBlogStorageMapInput` via:
//
//          BackupInstanceBlogStorageMap{ "key": BackupInstanceBlogStorageArgs{...} }
type BackupInstanceBlogStorageMapInput interface {
	pulumi.Input

	ToBackupInstanceBlogStorageMapOutput() BackupInstanceBlogStorageMapOutput
	ToBackupInstanceBlogStorageMapOutputWithContext(context.Context) BackupInstanceBlogStorageMapOutput
}

type BackupInstanceBlogStorageMap map[string]BackupInstanceBlogStorageInput

func (BackupInstanceBlogStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupInstanceBlogStorage)(nil)).Elem()
}

func (i BackupInstanceBlogStorageMap) ToBackupInstanceBlogStorageMapOutput() BackupInstanceBlogStorageMapOutput {
	return i.ToBackupInstanceBlogStorageMapOutputWithContext(context.Background())
}

func (i BackupInstanceBlogStorageMap) ToBackupInstanceBlogStorageMapOutputWithContext(ctx context.Context) BackupInstanceBlogStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceBlogStorageMapOutput)
}

type BackupInstanceBlogStorageOutput struct{ *pulumi.OutputState }

func (BackupInstanceBlogStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceBlogStorage)(nil))
}

func (o BackupInstanceBlogStorageOutput) ToBackupInstanceBlogStorageOutput() BackupInstanceBlogStorageOutput {
	return o
}

func (o BackupInstanceBlogStorageOutput) ToBackupInstanceBlogStorageOutputWithContext(ctx context.Context) BackupInstanceBlogStorageOutput {
	return o
}

func (o BackupInstanceBlogStorageOutput) ToBackupInstanceBlogStoragePtrOutput() BackupInstanceBlogStoragePtrOutput {
	return o.ToBackupInstanceBlogStoragePtrOutputWithContext(context.Background())
}

func (o BackupInstanceBlogStorageOutput) ToBackupInstanceBlogStoragePtrOutputWithContext(ctx context.Context) BackupInstanceBlogStoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupInstanceBlogStorage) *BackupInstanceBlogStorage {
		return &v
	}).(BackupInstanceBlogStoragePtrOutput)
}

type BackupInstanceBlogStoragePtrOutput struct{ *pulumi.OutputState }

func (BackupInstanceBlogStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceBlogStorage)(nil))
}

func (o BackupInstanceBlogStoragePtrOutput) ToBackupInstanceBlogStoragePtrOutput() BackupInstanceBlogStoragePtrOutput {
	return o
}

func (o BackupInstanceBlogStoragePtrOutput) ToBackupInstanceBlogStoragePtrOutputWithContext(ctx context.Context) BackupInstanceBlogStoragePtrOutput {
	return o
}

func (o BackupInstanceBlogStoragePtrOutput) Elem() BackupInstanceBlogStorageOutput {
	return o.ApplyT(func(v *BackupInstanceBlogStorage) BackupInstanceBlogStorage {
		if v != nil {
			return *v
		}
		var ret BackupInstanceBlogStorage
		return ret
	}).(BackupInstanceBlogStorageOutput)
}

type BackupInstanceBlogStorageArrayOutput struct{ *pulumi.OutputState }

func (BackupInstanceBlogStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupInstanceBlogStorage)(nil))
}

func (o BackupInstanceBlogStorageArrayOutput) ToBackupInstanceBlogStorageArrayOutput() BackupInstanceBlogStorageArrayOutput {
	return o
}

func (o BackupInstanceBlogStorageArrayOutput) ToBackupInstanceBlogStorageArrayOutputWithContext(ctx context.Context) BackupInstanceBlogStorageArrayOutput {
	return o
}

func (o BackupInstanceBlogStorageArrayOutput) Index(i pulumi.IntInput) BackupInstanceBlogStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupInstanceBlogStorage {
		return vs[0].([]BackupInstanceBlogStorage)[vs[1].(int)]
	}).(BackupInstanceBlogStorageOutput)
}

type BackupInstanceBlogStorageMapOutput struct{ *pulumi.OutputState }

func (BackupInstanceBlogStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackupInstanceBlogStorage)(nil))
}

func (o BackupInstanceBlogStorageMapOutput) ToBackupInstanceBlogStorageMapOutput() BackupInstanceBlogStorageMapOutput {
	return o
}

func (o BackupInstanceBlogStorageMapOutput) ToBackupInstanceBlogStorageMapOutputWithContext(ctx context.Context) BackupInstanceBlogStorageMapOutput {
	return o
}

func (o BackupInstanceBlogStorageMapOutput) MapIndex(k pulumi.StringInput) BackupInstanceBlogStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackupInstanceBlogStorage {
		return vs[0].(map[string]BackupInstanceBlogStorage)[vs[1].(string)]
	}).(BackupInstanceBlogStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInstanceBlogStorageInput)(nil)).Elem(), &BackupInstanceBlogStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInstanceBlogStoragePtrInput)(nil)).Elem(), &BackupInstanceBlogStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInstanceBlogStorageArrayInput)(nil)).Elem(), BackupInstanceBlogStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInstanceBlogStorageMapInput)(nil)).Elem(), BackupInstanceBlogStorageMap{})
	pulumi.RegisterOutputType(BackupInstanceBlogStorageOutput{})
	pulumi.RegisterOutputType(BackupInstanceBlogStoragePtrOutput{})
	pulumi.RegisterOutputType(BackupInstanceBlogStorageArrayOutput{})
	pulumi.RegisterOutputType(BackupInstanceBlogStorageMapOutput{})
}
