// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Backup Policy Blob Storage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/dataprotection"
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
// 		exampleBackupVault, err := dataprotection.NewBackupVault(ctx, "exampleBackupVault", &dataprotection.BackupVaultArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			DatastoreType:     pulumi.String("VaultStore"),
// 			Redundancy:        pulumi.String("LocallyRedundant"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataprotection.NewBackupPolicyBlobStorage(ctx, "exampleBackupPolicyBlobStorage", &dataprotection.BackupPolicyBlobStorageArgs{
// 			VaultId:           exampleBackupVault.ID(),
// 			RetentionDuration: pulumi.String("P30D"),
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
// Backup Policy Blob Storages can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dataprotection/backupPolicyBlobStorage:BackupPolicyBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
// ```
type BackupPolicyBlobStorage struct {
	pulumi.CustomResourceState

	// The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration pulumi.StringOutput `pulumi:"retentionDuration"`
	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewBackupPolicyBlobStorage registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicyBlobStorage(ctx *pulumi.Context,
	name string, args *BackupPolicyBlobStorageArgs, opts ...pulumi.ResourceOption) (*BackupPolicyBlobStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RetentionDuration == nil {
		return nil, errors.New("invalid value for required argument 'RetentionDuration'")
	}
	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	var resource BackupPolicyBlobStorage
	err := ctx.RegisterResource("azure:dataprotection/backupPolicyBlobStorage:BackupPolicyBlobStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicyBlobStorage gets an existing BackupPolicyBlobStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicyBlobStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyBlobStorageState, opts ...pulumi.ResourceOption) (*BackupPolicyBlobStorage, error) {
	var resource BackupPolicyBlobStorage
	err := ctx.ReadResource("azure:dataprotection/backupPolicyBlobStorage:BackupPolicyBlobStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicyBlobStorage resources.
type backupPolicyBlobStorageState struct {
	// The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
	Name *string `pulumi:"name"`
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `pulumi:"retentionDuration"`
	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultId *string `pulumi:"vaultId"`
}

type BackupPolicyBlobStorageState struct {
	// The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
	Name pulumi.StringPtrInput
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration pulumi.StringPtrInput
	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultId pulumi.StringPtrInput
}

func (BackupPolicyBlobStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyBlobStorageState)(nil)).Elem()
}

type backupPolicyBlobStorageArgs struct {
	// The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
	Name *string `pulumi:"name"`
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration string `pulumi:"retentionDuration"`
	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a BackupPolicyBlobStorage resource.
type BackupPolicyBlobStorageArgs struct {
	// The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
	Name pulumi.StringPtrInput
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration pulumi.StringInput
	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultId pulumi.StringInput
}

func (BackupPolicyBlobStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyBlobStorageArgs)(nil)).Elem()
}

type BackupPolicyBlobStorageInput interface {
	pulumi.Input

	ToBackupPolicyBlobStorageOutput() BackupPolicyBlobStorageOutput
	ToBackupPolicyBlobStorageOutputWithContext(ctx context.Context) BackupPolicyBlobStorageOutput
}

func (*BackupPolicyBlobStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyBlobStorage)(nil)).Elem()
}

func (i *BackupPolicyBlobStorage) ToBackupPolicyBlobStorageOutput() BackupPolicyBlobStorageOutput {
	return i.ToBackupPolicyBlobStorageOutputWithContext(context.Background())
}

func (i *BackupPolicyBlobStorage) ToBackupPolicyBlobStorageOutputWithContext(ctx context.Context) BackupPolicyBlobStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyBlobStorageOutput)
}

// BackupPolicyBlobStorageArrayInput is an input type that accepts BackupPolicyBlobStorageArray and BackupPolicyBlobStorageArrayOutput values.
// You can construct a concrete instance of `BackupPolicyBlobStorageArrayInput` via:
//
//          BackupPolicyBlobStorageArray{ BackupPolicyBlobStorageArgs{...} }
type BackupPolicyBlobStorageArrayInput interface {
	pulumi.Input

	ToBackupPolicyBlobStorageArrayOutput() BackupPolicyBlobStorageArrayOutput
	ToBackupPolicyBlobStorageArrayOutputWithContext(context.Context) BackupPolicyBlobStorageArrayOutput
}

type BackupPolicyBlobStorageArray []BackupPolicyBlobStorageInput

func (BackupPolicyBlobStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicyBlobStorage)(nil)).Elem()
}

func (i BackupPolicyBlobStorageArray) ToBackupPolicyBlobStorageArrayOutput() BackupPolicyBlobStorageArrayOutput {
	return i.ToBackupPolicyBlobStorageArrayOutputWithContext(context.Background())
}

func (i BackupPolicyBlobStorageArray) ToBackupPolicyBlobStorageArrayOutputWithContext(ctx context.Context) BackupPolicyBlobStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyBlobStorageArrayOutput)
}

// BackupPolicyBlobStorageMapInput is an input type that accepts BackupPolicyBlobStorageMap and BackupPolicyBlobStorageMapOutput values.
// You can construct a concrete instance of `BackupPolicyBlobStorageMapInput` via:
//
//          BackupPolicyBlobStorageMap{ "key": BackupPolicyBlobStorageArgs{...} }
type BackupPolicyBlobStorageMapInput interface {
	pulumi.Input

	ToBackupPolicyBlobStorageMapOutput() BackupPolicyBlobStorageMapOutput
	ToBackupPolicyBlobStorageMapOutputWithContext(context.Context) BackupPolicyBlobStorageMapOutput
}

type BackupPolicyBlobStorageMap map[string]BackupPolicyBlobStorageInput

func (BackupPolicyBlobStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicyBlobStorage)(nil)).Elem()
}

func (i BackupPolicyBlobStorageMap) ToBackupPolicyBlobStorageMapOutput() BackupPolicyBlobStorageMapOutput {
	return i.ToBackupPolicyBlobStorageMapOutputWithContext(context.Background())
}

func (i BackupPolicyBlobStorageMap) ToBackupPolicyBlobStorageMapOutputWithContext(ctx context.Context) BackupPolicyBlobStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyBlobStorageMapOutput)
}

type BackupPolicyBlobStorageOutput struct{ *pulumi.OutputState }

func (BackupPolicyBlobStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyBlobStorage)(nil)).Elem()
}

func (o BackupPolicyBlobStorageOutput) ToBackupPolicyBlobStorageOutput() BackupPolicyBlobStorageOutput {
	return o
}

func (o BackupPolicyBlobStorageOutput) ToBackupPolicyBlobStorageOutputWithContext(ctx context.Context) BackupPolicyBlobStorageOutput {
	return o
}

type BackupPolicyBlobStorageArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyBlobStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicyBlobStorage)(nil)).Elem()
}

func (o BackupPolicyBlobStorageArrayOutput) ToBackupPolicyBlobStorageArrayOutput() BackupPolicyBlobStorageArrayOutput {
	return o
}

func (o BackupPolicyBlobStorageArrayOutput) ToBackupPolicyBlobStorageArrayOutputWithContext(ctx context.Context) BackupPolicyBlobStorageArrayOutput {
	return o
}

func (o BackupPolicyBlobStorageArrayOutput) Index(i pulumi.IntInput) BackupPolicyBlobStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPolicyBlobStorage {
		return vs[0].([]*BackupPolicyBlobStorage)[vs[1].(int)]
	}).(BackupPolicyBlobStorageOutput)
}

type BackupPolicyBlobStorageMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyBlobStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicyBlobStorage)(nil)).Elem()
}

func (o BackupPolicyBlobStorageMapOutput) ToBackupPolicyBlobStorageMapOutput() BackupPolicyBlobStorageMapOutput {
	return o
}

func (o BackupPolicyBlobStorageMapOutput) ToBackupPolicyBlobStorageMapOutputWithContext(ctx context.Context) BackupPolicyBlobStorageMapOutput {
	return o
}

func (o BackupPolicyBlobStorageMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyBlobStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPolicyBlobStorage {
		return vs[0].(map[string]*BackupPolicyBlobStorage)[vs[1].(string)]
	}).(BackupPolicyBlobStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyBlobStorageInput)(nil)).Elem(), &BackupPolicyBlobStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyBlobStorageArrayInput)(nil)).Elem(), BackupPolicyBlobStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyBlobStorageMapInput)(nil)).Elem(), BackupPolicyBlobStorageMap{})
	pulumi.RegisterOutputType(BackupPolicyBlobStorageOutput{})
	pulumi.RegisterOutputType(BackupPolicyBlobStorageArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyBlobStorageMapOutput{})
}
