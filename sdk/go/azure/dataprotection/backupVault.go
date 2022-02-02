// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Backup Vault.
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
// 		_, err = dataprotection.NewBackupVault(ctx, "example", &dataprotection.BackupVaultArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			DatastoreType:     pulumi.String("VaultStore"),
// 			Redundancy:        pulumi.String("LocallyRedundant"),
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
// Backup Vaults can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dataprotection/backupVault:BackupVault example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1
// ```
type BackupVault struct {
	pulumi.CustomResourceState

	// Specifies the type of the data store. Possible values are `ArchiveStore`, `SnapshotStore` and `VaultStore`.
	DatastoreType pulumi.StringOutput `pulumi:"datastoreType"`
	// An `identity` block as defined below.
	Identity BackupVaultIdentityPtrOutput `pulumi:"identity"`
	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Backup Vault. Changing this forces a new Backup Vault to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the backup storage redundancy. Possible values are `GeoRedundant` and `LocallyRedundant`. Changing this forces a new Backup Vault to be created.
	Redundancy pulumi.StringOutput `pulumi:"redundancy"`
	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Backup Vault.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewBackupVault registers a new resource with the given unique name, arguments, and options.
func NewBackupVault(ctx *pulumi.Context,
	name string, args *BackupVaultArgs, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreType == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreType'")
	}
	if args.Redundancy == nil {
		return nil, errors.New("invalid value for required argument 'Redundancy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource BackupVault
	err := ctx.RegisterResource("azure:dataprotection/backupVault:BackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupVault gets an existing BackupVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupVaultState, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	var resource BackupVault
	err := ctx.ReadResource("azure:dataprotection/backupVault:BackupVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupVault resources.
type backupVaultState struct {
	// Specifies the type of the data store. Possible values are `ArchiveStore`, `SnapshotStore` and `VaultStore`.
	DatastoreType *string `pulumi:"datastoreType"`
	// An `identity` block as defined below.
	Identity *BackupVaultIdentity `pulumi:"identity"`
	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Backup Vault. Changing this forces a new Backup Vault to be created.
	Name *string `pulumi:"name"`
	// Specifies the backup storage redundancy. Possible values are `GeoRedundant` and `LocallyRedundant`. Changing this forces a new Backup Vault to be created.
	Redundancy *string `pulumi:"redundancy"`
	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Backup Vault.
	Tags map[string]string `pulumi:"tags"`
}

type BackupVaultState struct {
	// Specifies the type of the data store. Possible values are `ArchiveStore`, `SnapshotStore` and `VaultStore`.
	DatastoreType pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity BackupVaultIdentityPtrInput
	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Backup Vault. Changing this forces a new Backup Vault to be created.
	Name pulumi.StringPtrInput
	// Specifies the backup storage redundancy. Possible values are `GeoRedundant` and `LocallyRedundant`. Changing this forces a new Backup Vault to be created.
	Redundancy pulumi.StringPtrInput
	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Backup Vault.
	Tags pulumi.StringMapInput
}

func (BackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultState)(nil)).Elem()
}

type backupVaultArgs struct {
	// Specifies the type of the data store. Possible values are `ArchiveStore`, `SnapshotStore` and `VaultStore`.
	DatastoreType string `pulumi:"datastoreType"`
	// An `identity` block as defined below.
	Identity *BackupVaultIdentity `pulumi:"identity"`
	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Backup Vault. Changing this forces a new Backup Vault to be created.
	Name *string `pulumi:"name"`
	// Specifies the backup storage redundancy. Possible values are `GeoRedundant` and `LocallyRedundant`. Changing this forces a new Backup Vault to be created.
	Redundancy string `pulumi:"redundancy"`
	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Backup Vault.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BackupVault resource.
type BackupVaultArgs struct {
	// Specifies the type of the data store. Possible values are `ArchiveStore`, `SnapshotStore` and `VaultStore`.
	DatastoreType pulumi.StringInput
	// An `identity` block as defined below.
	Identity BackupVaultIdentityPtrInput
	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Backup Vault. Changing this forces a new Backup Vault to be created.
	Name pulumi.StringPtrInput
	// Specifies the backup storage redundancy. Possible values are `GeoRedundant` and `LocallyRedundant`. Changing this forces a new Backup Vault to be created.
	Redundancy pulumi.StringInput
	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Backup Vault.
	Tags pulumi.StringMapInput
}

func (BackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultArgs)(nil)).Elem()
}

type BackupVaultInput interface {
	pulumi.Input

	ToBackupVaultOutput() BackupVaultOutput
	ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput
}

func (*BackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (i *BackupVault) ToBackupVaultOutput() BackupVaultOutput {
	return i.ToBackupVaultOutputWithContext(context.Background())
}

func (i *BackupVault) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultOutput)
}

// BackupVaultArrayInput is an input type that accepts BackupVaultArray and BackupVaultArrayOutput values.
// You can construct a concrete instance of `BackupVaultArrayInput` via:
//
//          BackupVaultArray{ BackupVaultArgs{...} }
type BackupVaultArrayInput interface {
	pulumi.Input

	ToBackupVaultArrayOutput() BackupVaultArrayOutput
	ToBackupVaultArrayOutputWithContext(context.Context) BackupVaultArrayOutput
}

type BackupVaultArray []BackupVaultInput

func (BackupVaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupVault)(nil)).Elem()
}

func (i BackupVaultArray) ToBackupVaultArrayOutput() BackupVaultArrayOutput {
	return i.ToBackupVaultArrayOutputWithContext(context.Background())
}

func (i BackupVaultArray) ToBackupVaultArrayOutputWithContext(ctx context.Context) BackupVaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultArrayOutput)
}

// BackupVaultMapInput is an input type that accepts BackupVaultMap and BackupVaultMapOutput values.
// You can construct a concrete instance of `BackupVaultMapInput` via:
//
//          BackupVaultMap{ "key": BackupVaultArgs{...} }
type BackupVaultMapInput interface {
	pulumi.Input

	ToBackupVaultMapOutput() BackupVaultMapOutput
	ToBackupVaultMapOutputWithContext(context.Context) BackupVaultMapOutput
}

type BackupVaultMap map[string]BackupVaultInput

func (BackupVaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupVault)(nil)).Elem()
}

func (i BackupVaultMap) ToBackupVaultMapOutput() BackupVaultMapOutput {
	return i.ToBackupVaultMapOutputWithContext(context.Background())
}

func (i BackupVaultMap) ToBackupVaultMapOutputWithContext(ctx context.Context) BackupVaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultMapOutput)
}

type BackupVaultOutput struct{ *pulumi.OutputState }

func (BackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (o BackupVaultOutput) ToBackupVaultOutput() BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return o
}

type BackupVaultArrayOutput struct{ *pulumi.OutputState }

func (BackupVaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupVault)(nil)).Elem()
}

func (o BackupVaultArrayOutput) ToBackupVaultArrayOutput() BackupVaultArrayOutput {
	return o
}

func (o BackupVaultArrayOutput) ToBackupVaultArrayOutputWithContext(ctx context.Context) BackupVaultArrayOutput {
	return o
}

func (o BackupVaultArrayOutput) Index(i pulumi.IntInput) BackupVaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupVault {
		return vs[0].([]*BackupVault)[vs[1].(int)]
	}).(BackupVaultOutput)
}

type BackupVaultMapOutput struct{ *pulumi.OutputState }

func (BackupVaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupVault)(nil)).Elem()
}

func (o BackupVaultMapOutput) ToBackupVaultMapOutput() BackupVaultMapOutput {
	return o
}

func (o BackupVaultMapOutput) ToBackupVaultMapOutputWithContext(ctx context.Context) BackupVaultMapOutput {
	return o
}

func (o BackupVaultMapOutput) MapIndex(k pulumi.StringInput) BackupVaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupVault {
		return vs[0].(map[string]*BackupVault)[vs[1].(string)]
	}).(BackupVaultOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupVaultInput)(nil)).Elem(), &BackupVault{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupVaultArrayInput)(nil)).Elem(), BackupVaultArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupVaultMapInput)(nil)).Elem(), BackupVaultMap{})
	pulumi.RegisterOutputType(BackupVaultOutput{})
	pulumi.RegisterOutputType(BackupVaultArrayOutput{})
	pulumi.RegisterOutputType(BackupVaultMapOutput{})
}
