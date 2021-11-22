// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Storage Sync Cloud Endpoint.
//
// > **NOTE:** Please ensure Azure File Sync has access to the storage account in your subscription, which indicates that `Microsoft.StorageSync` is assigned role `Reader and Data Access` ( refer to details [here](https://docs.microsoft.com/en-us/azure/storage/files/storage-sync-files-troubleshoot?tabs=portal1%2Cazure-portal#common-troubleshooting-steps)).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
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
// 		exampleSync, err := storage.NewSync(ctx, "exampleSync", &storage.SyncArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSyncGroup, err := storage.NewSyncGroup(ctx, "exampleSyncGroup", &storage.SyncGroupArgs{
// 			StorageSyncId: exampleSync.ID(),
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
// 		exampleShare, err := storage.NewShare(ctx, "exampleShare", &storage.ShareArgs{
// 			StorageAccountName: exampleAccount.Name,
// 			Acls: storage.ShareAclArray{
// 				&storage.ShareAclArgs{
// 					Id: pulumi.String("GhostedRecall"),
// 					AccessPolicies: storage.ShareAclAccessPolicyArray{
// 						&storage.ShareAclAccessPolicyArgs{
// 							Permissions: pulumi.String("r"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewSyncCloudEndpoint(ctx, "exampleSyncCloudEndpoint", &storage.SyncCloudEndpointArgs{
// 			StorageSyncGroupId: exampleSyncGroup.ID(),
// 			FileShareName:      exampleShare.Name,
// 			StorageAccountId:   exampleAccount.ID(),
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
// Storage Sync Cloud Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/syncCloudEndpoint:SyncCloudEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/syncgroup1/cloudEndpoints/cloudEndpoint1
// ```
type SyncCloudEndpoint struct {
	pulumi.CustomResourceState

	// The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	FileShareName pulumi.StringOutput `pulumi:"fileShareName"`
	// The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
	StorageAccountTenantId pulumi.StringOutput `pulumi:"storageAccountTenantId"`
	// The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageSyncGroupId pulumi.StringOutput `pulumi:"storageSyncGroupId"`
}

// NewSyncCloudEndpoint registers a new resource with the given unique name, arguments, and options.
func NewSyncCloudEndpoint(ctx *pulumi.Context,
	name string, args *SyncCloudEndpointArgs, opts ...pulumi.ResourceOption) (*SyncCloudEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileShareName == nil {
		return nil, errors.New("invalid value for required argument 'FileShareName'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	if args.StorageSyncGroupId == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncGroupId'")
	}
	var resource SyncCloudEndpoint
	err := ctx.RegisterResource("azure:storage/syncCloudEndpoint:SyncCloudEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncCloudEndpoint gets an existing SyncCloudEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncCloudEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncCloudEndpointState, opts ...pulumi.ResourceOption) (*SyncCloudEndpoint, error) {
	var resource SyncCloudEndpoint
	err := ctx.ReadResource("azure:storage/syncCloudEndpoint:SyncCloudEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncCloudEndpoint resources.
type syncCloudEndpointState struct {
	// The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	FileShareName *string `pulumi:"fileShareName"`
	// The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
	StorageAccountTenantId *string `pulumi:"storageAccountTenantId"`
	// The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageSyncGroupId *string `pulumi:"storageSyncGroupId"`
}

type SyncCloudEndpointState struct {
	// The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	FileShareName pulumi.StringPtrInput
	// The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageAccountId pulumi.StringPtrInput
	// The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
	StorageAccountTenantId pulumi.StringPtrInput
	// The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageSyncGroupId pulumi.StringPtrInput
}

func (SyncCloudEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncCloudEndpointState)(nil)).Elem()
}

type syncCloudEndpointArgs struct {
	// The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	FileShareName string `pulumi:"fileShareName"`
	// The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	Name *string `pulumi:"name"`
	// The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
	StorageAccountTenantId *string `pulumi:"storageAccountTenantId"`
	// The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageSyncGroupId string `pulumi:"storageSyncGroupId"`
}

// The set of arguments for constructing a SyncCloudEndpoint resource.
type SyncCloudEndpointArgs struct {
	// The Storage Share name to be synchronized in this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	FileShareName pulumi.StringInput
	// The name which should be used for this Storage Sync Cloud Endpoint. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	Name pulumi.StringPtrInput
	// The ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageAccountId pulumi.StringInput
	// The Tenant ID of the Storage Account where the Storage Share exists. Changing this forces a new Storage Sync Cloud Endpoint to be created. Defaults to the current tenant id.
	StorageAccountTenantId pulumi.StringPtrInput
	// The ID of the Storage Sync Group where this Cloud Endpoint should be created. Changing this forces a new Storage Sync Cloud Endpoint to be created.
	StorageSyncGroupId pulumi.StringInput
}

func (SyncCloudEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncCloudEndpointArgs)(nil)).Elem()
}

type SyncCloudEndpointInput interface {
	pulumi.Input

	ToSyncCloudEndpointOutput() SyncCloudEndpointOutput
	ToSyncCloudEndpointOutputWithContext(ctx context.Context) SyncCloudEndpointOutput
}

func (*SyncCloudEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncCloudEndpoint)(nil))
}

func (i *SyncCloudEndpoint) ToSyncCloudEndpointOutput() SyncCloudEndpointOutput {
	return i.ToSyncCloudEndpointOutputWithContext(context.Background())
}

func (i *SyncCloudEndpoint) ToSyncCloudEndpointOutputWithContext(ctx context.Context) SyncCloudEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCloudEndpointOutput)
}

func (i *SyncCloudEndpoint) ToSyncCloudEndpointPtrOutput() SyncCloudEndpointPtrOutput {
	return i.ToSyncCloudEndpointPtrOutputWithContext(context.Background())
}

func (i *SyncCloudEndpoint) ToSyncCloudEndpointPtrOutputWithContext(ctx context.Context) SyncCloudEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCloudEndpointPtrOutput)
}

type SyncCloudEndpointPtrInput interface {
	pulumi.Input

	ToSyncCloudEndpointPtrOutput() SyncCloudEndpointPtrOutput
	ToSyncCloudEndpointPtrOutputWithContext(ctx context.Context) SyncCloudEndpointPtrOutput
}

type syncCloudEndpointPtrType SyncCloudEndpointArgs

func (*syncCloudEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncCloudEndpoint)(nil))
}

func (i *syncCloudEndpointPtrType) ToSyncCloudEndpointPtrOutput() SyncCloudEndpointPtrOutput {
	return i.ToSyncCloudEndpointPtrOutputWithContext(context.Background())
}

func (i *syncCloudEndpointPtrType) ToSyncCloudEndpointPtrOutputWithContext(ctx context.Context) SyncCloudEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCloudEndpointPtrOutput)
}

// SyncCloudEndpointArrayInput is an input type that accepts SyncCloudEndpointArray and SyncCloudEndpointArrayOutput values.
// You can construct a concrete instance of `SyncCloudEndpointArrayInput` via:
//
//          SyncCloudEndpointArray{ SyncCloudEndpointArgs{...} }
type SyncCloudEndpointArrayInput interface {
	pulumi.Input

	ToSyncCloudEndpointArrayOutput() SyncCloudEndpointArrayOutput
	ToSyncCloudEndpointArrayOutputWithContext(context.Context) SyncCloudEndpointArrayOutput
}

type SyncCloudEndpointArray []SyncCloudEndpointInput

func (SyncCloudEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncCloudEndpoint)(nil)).Elem()
}

func (i SyncCloudEndpointArray) ToSyncCloudEndpointArrayOutput() SyncCloudEndpointArrayOutput {
	return i.ToSyncCloudEndpointArrayOutputWithContext(context.Background())
}

func (i SyncCloudEndpointArray) ToSyncCloudEndpointArrayOutputWithContext(ctx context.Context) SyncCloudEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCloudEndpointArrayOutput)
}

// SyncCloudEndpointMapInput is an input type that accepts SyncCloudEndpointMap and SyncCloudEndpointMapOutput values.
// You can construct a concrete instance of `SyncCloudEndpointMapInput` via:
//
//          SyncCloudEndpointMap{ "key": SyncCloudEndpointArgs{...} }
type SyncCloudEndpointMapInput interface {
	pulumi.Input

	ToSyncCloudEndpointMapOutput() SyncCloudEndpointMapOutput
	ToSyncCloudEndpointMapOutputWithContext(context.Context) SyncCloudEndpointMapOutput
}

type SyncCloudEndpointMap map[string]SyncCloudEndpointInput

func (SyncCloudEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncCloudEndpoint)(nil)).Elem()
}

func (i SyncCloudEndpointMap) ToSyncCloudEndpointMapOutput() SyncCloudEndpointMapOutput {
	return i.ToSyncCloudEndpointMapOutputWithContext(context.Background())
}

func (i SyncCloudEndpointMap) ToSyncCloudEndpointMapOutputWithContext(ctx context.Context) SyncCloudEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncCloudEndpointMapOutput)
}

type SyncCloudEndpointOutput struct{ *pulumi.OutputState }

func (SyncCloudEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncCloudEndpoint)(nil))
}

func (o SyncCloudEndpointOutput) ToSyncCloudEndpointOutput() SyncCloudEndpointOutput {
	return o
}

func (o SyncCloudEndpointOutput) ToSyncCloudEndpointOutputWithContext(ctx context.Context) SyncCloudEndpointOutput {
	return o
}

func (o SyncCloudEndpointOutput) ToSyncCloudEndpointPtrOutput() SyncCloudEndpointPtrOutput {
	return o.ToSyncCloudEndpointPtrOutputWithContext(context.Background())
}

func (o SyncCloudEndpointOutput) ToSyncCloudEndpointPtrOutputWithContext(ctx context.Context) SyncCloudEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncCloudEndpoint) *SyncCloudEndpoint {
		return &v
	}).(SyncCloudEndpointPtrOutput)
}

type SyncCloudEndpointPtrOutput struct{ *pulumi.OutputState }

func (SyncCloudEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncCloudEndpoint)(nil))
}

func (o SyncCloudEndpointPtrOutput) ToSyncCloudEndpointPtrOutput() SyncCloudEndpointPtrOutput {
	return o
}

func (o SyncCloudEndpointPtrOutput) ToSyncCloudEndpointPtrOutputWithContext(ctx context.Context) SyncCloudEndpointPtrOutput {
	return o
}

func (o SyncCloudEndpointPtrOutput) Elem() SyncCloudEndpointOutput {
	return o.ApplyT(func(v *SyncCloudEndpoint) SyncCloudEndpoint {
		if v != nil {
			return *v
		}
		var ret SyncCloudEndpoint
		return ret
	}).(SyncCloudEndpointOutput)
}

type SyncCloudEndpointArrayOutput struct{ *pulumi.OutputState }

func (SyncCloudEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncCloudEndpoint)(nil))
}

func (o SyncCloudEndpointArrayOutput) ToSyncCloudEndpointArrayOutput() SyncCloudEndpointArrayOutput {
	return o
}

func (o SyncCloudEndpointArrayOutput) ToSyncCloudEndpointArrayOutputWithContext(ctx context.Context) SyncCloudEndpointArrayOutput {
	return o
}

func (o SyncCloudEndpointArrayOutput) Index(i pulumi.IntInput) SyncCloudEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncCloudEndpoint {
		return vs[0].([]SyncCloudEndpoint)[vs[1].(int)]
	}).(SyncCloudEndpointOutput)
}

type SyncCloudEndpointMapOutput struct{ *pulumi.OutputState }

func (SyncCloudEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SyncCloudEndpoint)(nil))
}

func (o SyncCloudEndpointMapOutput) ToSyncCloudEndpointMapOutput() SyncCloudEndpointMapOutput {
	return o
}

func (o SyncCloudEndpointMapOutput) ToSyncCloudEndpointMapOutputWithContext(ctx context.Context) SyncCloudEndpointMapOutput {
	return o
}

func (o SyncCloudEndpointMapOutput) MapIndex(k pulumi.StringInput) SyncCloudEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SyncCloudEndpoint {
		return vs[0].(map[string]SyncCloudEndpoint)[vs[1].(string)]
	}).(SyncCloudEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCloudEndpointInput)(nil)).Elem(), &SyncCloudEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCloudEndpointPtrInput)(nil)).Elem(), &SyncCloudEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCloudEndpointArrayInput)(nil)).Elem(), SyncCloudEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncCloudEndpointMapInput)(nil)).Elem(), SyncCloudEndpointMap{})
	pulumi.RegisterOutputType(SyncCloudEndpointOutput{})
	pulumi.RegisterOutputType(SyncCloudEndpointPtrOutput{})
	pulumi.RegisterOutputType(SyncCloudEndpointArrayOutput{})
	pulumi.RegisterOutputType(SyncCloudEndpointMapOutput{})
}
