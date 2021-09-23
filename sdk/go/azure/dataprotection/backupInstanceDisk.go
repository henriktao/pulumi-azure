// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Backup Instance to back up Disk.
//
// ## Import
//
// Backup Instance Disks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dataprotection/backupInstanceDisk:BackupInstanceDisk example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
// ```
type BackupInstanceDisk struct {
	pulumi.CustomResourceState

	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Backup Instance Disk. Changing this forces a new Backup Instance Disk to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName pulumi.StringOutput `pulumi:"snapshotResourceGroupName"`
	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewBackupInstanceDisk registers a new resource with the given unique name, arguments, and options.
func NewBackupInstanceDisk(ctx *pulumi.Context,
	name string, args *BackupInstanceDiskArgs, opts ...pulumi.ResourceOption) (*BackupInstanceDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyId'")
	}
	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.SnapshotResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotResourceGroupName'")
	}
	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	var resource BackupInstanceDisk
	err := ctx.RegisterResource("azure:dataprotection/backupInstanceDisk:BackupInstanceDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupInstanceDisk gets an existing BackupInstanceDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupInstanceDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupInstanceDiskState, opts ...pulumi.ResourceOption) (*BackupInstanceDisk, error) {
	var resource BackupInstanceDisk
	err := ctx.ReadResource("azure:dataprotection/backupInstanceDisk:BackupInstanceDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupInstanceDisk resources.
type backupInstanceDiskState struct {
	// The ID of the Backup Policy.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskId *string `pulumi:"diskId"`
	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Backup Instance Disk. Changing this forces a new Backup Instance Disk to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName *string `pulumi:"snapshotResourceGroupName"`
	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultId *string `pulumi:"vaultId"`
}

type BackupInstanceDiskState struct {
	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringPtrInput
	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskId pulumi.StringPtrInput
	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Backup Instance Disk. Changing this forces a new Backup Instance Disk to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName pulumi.StringPtrInput
	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultId pulumi.StringPtrInput
}

func (BackupInstanceDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceDiskState)(nil)).Elem()
}

type backupInstanceDiskArgs struct {
	// The ID of the Backup Policy.
	BackupPolicyId string `pulumi:"backupPolicyId"`
	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskId string `pulumi:"diskId"`
	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Backup Instance Disk. Changing this forces a new Backup Instance Disk to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName string `pulumi:"snapshotResourceGroupName"`
	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a BackupInstanceDisk resource.
type BackupInstanceDiskArgs struct {
	// The ID of the Backup Policy.
	BackupPolicyId pulumi.StringInput
	// The ID of the source Disk. Changing this forces a new Backup Instance Disk to be created.
	DiskId pulumi.StringInput
	// The Azure Region where the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Backup Instance Disk. Changing this forces a new Backup Instance Disk to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where snapshots are stored. Changing this forces a new Backup Instance Disk to be created.
	SnapshotResourceGroupName pulumi.StringInput
	// The ID of the Backup Vault within which the Backup Instance Disk should exist. Changing this forces a new Backup Instance Disk to be created.
	VaultId pulumi.StringInput
}

func (BackupInstanceDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceDiskArgs)(nil)).Elem()
}

type BackupInstanceDiskInput interface {
	pulumi.Input

	ToBackupInstanceDiskOutput() BackupInstanceDiskOutput
	ToBackupInstanceDiskOutputWithContext(ctx context.Context) BackupInstanceDiskOutput
}

func (*BackupInstanceDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceDisk)(nil))
}

func (i *BackupInstanceDisk) ToBackupInstanceDiskOutput() BackupInstanceDiskOutput {
	return i.ToBackupInstanceDiskOutputWithContext(context.Background())
}

func (i *BackupInstanceDisk) ToBackupInstanceDiskOutputWithContext(ctx context.Context) BackupInstanceDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceDiskOutput)
}

func (i *BackupInstanceDisk) ToBackupInstanceDiskPtrOutput() BackupInstanceDiskPtrOutput {
	return i.ToBackupInstanceDiskPtrOutputWithContext(context.Background())
}

func (i *BackupInstanceDisk) ToBackupInstanceDiskPtrOutputWithContext(ctx context.Context) BackupInstanceDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceDiskPtrOutput)
}

type BackupInstanceDiskPtrInput interface {
	pulumi.Input

	ToBackupInstanceDiskPtrOutput() BackupInstanceDiskPtrOutput
	ToBackupInstanceDiskPtrOutputWithContext(ctx context.Context) BackupInstanceDiskPtrOutput
}

type backupInstanceDiskPtrType BackupInstanceDiskArgs

func (*backupInstanceDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceDisk)(nil))
}

func (i *backupInstanceDiskPtrType) ToBackupInstanceDiskPtrOutput() BackupInstanceDiskPtrOutput {
	return i.ToBackupInstanceDiskPtrOutputWithContext(context.Background())
}

func (i *backupInstanceDiskPtrType) ToBackupInstanceDiskPtrOutputWithContext(ctx context.Context) BackupInstanceDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceDiskPtrOutput)
}

// BackupInstanceDiskArrayInput is an input type that accepts BackupInstanceDiskArray and BackupInstanceDiskArrayOutput values.
// You can construct a concrete instance of `BackupInstanceDiskArrayInput` via:
//
//          BackupInstanceDiskArray{ BackupInstanceDiskArgs{...} }
type BackupInstanceDiskArrayInput interface {
	pulumi.Input

	ToBackupInstanceDiskArrayOutput() BackupInstanceDiskArrayOutput
	ToBackupInstanceDiskArrayOutputWithContext(context.Context) BackupInstanceDiskArrayOutput
}

type BackupInstanceDiskArray []BackupInstanceDiskInput

func (BackupInstanceDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupInstanceDisk)(nil)).Elem()
}

func (i BackupInstanceDiskArray) ToBackupInstanceDiskArrayOutput() BackupInstanceDiskArrayOutput {
	return i.ToBackupInstanceDiskArrayOutputWithContext(context.Background())
}

func (i BackupInstanceDiskArray) ToBackupInstanceDiskArrayOutputWithContext(ctx context.Context) BackupInstanceDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceDiskArrayOutput)
}

// BackupInstanceDiskMapInput is an input type that accepts BackupInstanceDiskMap and BackupInstanceDiskMapOutput values.
// You can construct a concrete instance of `BackupInstanceDiskMapInput` via:
//
//          BackupInstanceDiskMap{ "key": BackupInstanceDiskArgs{...} }
type BackupInstanceDiskMapInput interface {
	pulumi.Input

	ToBackupInstanceDiskMapOutput() BackupInstanceDiskMapOutput
	ToBackupInstanceDiskMapOutputWithContext(context.Context) BackupInstanceDiskMapOutput
}

type BackupInstanceDiskMap map[string]BackupInstanceDiskInput

func (BackupInstanceDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupInstanceDisk)(nil)).Elem()
}

func (i BackupInstanceDiskMap) ToBackupInstanceDiskMapOutput() BackupInstanceDiskMapOutput {
	return i.ToBackupInstanceDiskMapOutputWithContext(context.Background())
}

func (i BackupInstanceDiskMap) ToBackupInstanceDiskMapOutputWithContext(ctx context.Context) BackupInstanceDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceDiskMapOutput)
}

type BackupInstanceDiskOutput struct{ *pulumi.OutputState }

func (BackupInstanceDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupInstanceDisk)(nil))
}

func (o BackupInstanceDiskOutput) ToBackupInstanceDiskOutput() BackupInstanceDiskOutput {
	return o
}

func (o BackupInstanceDiskOutput) ToBackupInstanceDiskOutputWithContext(ctx context.Context) BackupInstanceDiskOutput {
	return o
}

func (o BackupInstanceDiskOutput) ToBackupInstanceDiskPtrOutput() BackupInstanceDiskPtrOutput {
	return o.ToBackupInstanceDiskPtrOutputWithContext(context.Background())
}

func (o BackupInstanceDiskOutput) ToBackupInstanceDiskPtrOutputWithContext(ctx context.Context) BackupInstanceDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupInstanceDisk) *BackupInstanceDisk {
		return &v
	}).(BackupInstanceDiskPtrOutput)
}

type BackupInstanceDiskPtrOutput struct{ *pulumi.OutputState }

func (BackupInstanceDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstanceDisk)(nil))
}

func (o BackupInstanceDiskPtrOutput) ToBackupInstanceDiskPtrOutput() BackupInstanceDiskPtrOutput {
	return o
}

func (o BackupInstanceDiskPtrOutput) ToBackupInstanceDiskPtrOutputWithContext(ctx context.Context) BackupInstanceDiskPtrOutput {
	return o
}

func (o BackupInstanceDiskPtrOutput) Elem() BackupInstanceDiskOutput {
	return o.ApplyT(func(v *BackupInstanceDisk) BackupInstanceDisk {
		if v != nil {
			return *v
		}
		var ret BackupInstanceDisk
		return ret
	}).(BackupInstanceDiskOutput)
}

type BackupInstanceDiskArrayOutput struct{ *pulumi.OutputState }

func (BackupInstanceDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupInstanceDisk)(nil))
}

func (o BackupInstanceDiskArrayOutput) ToBackupInstanceDiskArrayOutput() BackupInstanceDiskArrayOutput {
	return o
}

func (o BackupInstanceDiskArrayOutput) ToBackupInstanceDiskArrayOutputWithContext(ctx context.Context) BackupInstanceDiskArrayOutput {
	return o
}

func (o BackupInstanceDiskArrayOutput) Index(i pulumi.IntInput) BackupInstanceDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupInstanceDisk {
		return vs[0].([]BackupInstanceDisk)[vs[1].(int)]
	}).(BackupInstanceDiskOutput)
}

type BackupInstanceDiskMapOutput struct{ *pulumi.OutputState }

func (BackupInstanceDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackupInstanceDisk)(nil))
}

func (o BackupInstanceDiskMapOutput) ToBackupInstanceDiskMapOutput() BackupInstanceDiskMapOutput {
	return o
}

func (o BackupInstanceDiskMapOutput) ToBackupInstanceDiskMapOutputWithContext(ctx context.Context) BackupInstanceDiskMapOutput {
	return o
}

func (o BackupInstanceDiskMapOutput) MapIndex(k pulumi.StringInput) BackupInstanceDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackupInstanceDisk {
		return vs[0].(map[string]BackupInstanceDisk)[vs[1].(string)]
	}).(BackupInstanceDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupInstanceDiskOutput{})
	pulumi.RegisterOutputType(BackupInstanceDiskPtrOutput{})
	pulumi.RegisterOutputType(BackupInstanceDiskArrayOutput{})
	pulumi.RegisterOutputType(BackupInstanceDiskMapOutput{})
}
