// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Storage Sync.
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
// 		_, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewSync(ctx, "test", &storage.SyncArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 			Location:          pulumi.Any(azurerm_resource_group.Test.Location),
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// Storage Syncs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/sync:Sync example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1
// ```
type Sync struct {
	pulumi.CustomResourceState

	// Incoming traffic policy. Possible values are `AllowAllTraffic` and `AllowVirtualNetworksOnly`.
	IncomingTrafficPolicy pulumi.StringPtrOutput `pulumi:"incomingTrafficPolicy"`
	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Storage Sync. Changing this forces a new Storage Sync to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Storage Sync.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSync registers a new resource with the given unique name, arguments, and options.
func NewSync(ctx *pulumi.Context,
	name string, args *SyncArgs, opts ...pulumi.ResourceOption) (*Sync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Sync
	err := ctx.RegisterResource("azure:storage/sync:Sync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSync gets an existing Sync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncState, opts ...pulumi.ResourceOption) (*Sync, error) {
	var resource Sync
	err := ctx.ReadResource("azure:storage/sync:Sync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sync resources.
type syncState struct {
	// Incoming traffic policy. Possible values are `AllowAllTraffic` and `AllowVirtualNetworksOnly`.
	IncomingTrafficPolicy *string `pulumi:"incomingTrafficPolicy"`
	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Storage Sync. Changing this forces a new Storage Sync to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Storage Sync.
	Tags map[string]string `pulumi:"tags"`
}

type SyncState struct {
	// Incoming traffic policy. Possible values are `AllowAllTraffic` and `AllowVirtualNetworksOnly`.
	IncomingTrafficPolicy pulumi.StringPtrInput
	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Storage Sync. Changing this forces a new Storage Sync to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Storage Sync.
	Tags pulumi.StringMapInput
}

func (SyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncState)(nil)).Elem()
}

type syncArgs struct {
	// Incoming traffic policy. Possible values are `AllowAllTraffic` and `AllowVirtualNetworksOnly`.
	IncomingTrafficPolicy *string `pulumi:"incomingTrafficPolicy"`
	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Storage Sync. Changing this forces a new Storage Sync to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Storage Sync.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Sync resource.
type SyncArgs struct {
	// Incoming traffic policy. Possible values are `AllowAllTraffic` and `AllowVirtualNetworksOnly`.
	IncomingTrafficPolicy pulumi.StringPtrInput
	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Storage Sync. Changing this forces a new Storage Sync to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Storage Sync.
	Tags pulumi.StringMapInput
}

func (SyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncArgs)(nil)).Elem()
}

type SyncInput interface {
	pulumi.Input

	ToSyncOutput() SyncOutput
	ToSyncOutputWithContext(ctx context.Context) SyncOutput
}

func (*Sync) ElementType() reflect.Type {
	return reflect.TypeOf((*Sync)(nil))
}

func (i *Sync) ToSyncOutput() SyncOutput {
	return i.ToSyncOutputWithContext(context.Background())
}

func (i *Sync) ToSyncOutputWithContext(ctx context.Context) SyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncOutput)
}

func (i *Sync) ToSyncPtrOutput() SyncPtrOutput {
	return i.ToSyncPtrOutputWithContext(context.Background())
}

func (i *Sync) ToSyncPtrOutputWithContext(ctx context.Context) SyncPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPtrOutput)
}

type SyncPtrInput interface {
	pulumi.Input

	ToSyncPtrOutput() SyncPtrOutput
	ToSyncPtrOutputWithContext(ctx context.Context) SyncPtrOutput
}

type syncPtrType SyncArgs

func (*syncPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sync)(nil))
}

func (i *syncPtrType) ToSyncPtrOutput() SyncPtrOutput {
	return i.ToSyncPtrOutputWithContext(context.Background())
}

func (i *syncPtrType) ToSyncPtrOutputWithContext(ctx context.Context) SyncPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPtrOutput)
}

// SyncArrayInput is an input type that accepts SyncArray and SyncArrayOutput values.
// You can construct a concrete instance of `SyncArrayInput` via:
//
//          SyncArray{ SyncArgs{...} }
type SyncArrayInput interface {
	pulumi.Input

	ToSyncArrayOutput() SyncArrayOutput
	ToSyncArrayOutputWithContext(context.Context) SyncArrayOutput
}

type SyncArray []SyncInput

func (SyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sync)(nil)).Elem()
}

func (i SyncArray) ToSyncArrayOutput() SyncArrayOutput {
	return i.ToSyncArrayOutputWithContext(context.Background())
}

func (i SyncArray) ToSyncArrayOutputWithContext(ctx context.Context) SyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncArrayOutput)
}

// SyncMapInput is an input type that accepts SyncMap and SyncMapOutput values.
// You can construct a concrete instance of `SyncMapInput` via:
//
//          SyncMap{ "key": SyncArgs{...} }
type SyncMapInput interface {
	pulumi.Input

	ToSyncMapOutput() SyncMapOutput
	ToSyncMapOutputWithContext(context.Context) SyncMapOutput
}

type SyncMap map[string]SyncInput

func (SyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sync)(nil)).Elem()
}

func (i SyncMap) ToSyncMapOutput() SyncMapOutput {
	return i.ToSyncMapOutputWithContext(context.Background())
}

func (i SyncMap) ToSyncMapOutputWithContext(ctx context.Context) SyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncMapOutput)
}

type SyncOutput struct{ *pulumi.OutputState }

func (SyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sync)(nil))
}

func (o SyncOutput) ToSyncOutput() SyncOutput {
	return o
}

func (o SyncOutput) ToSyncOutputWithContext(ctx context.Context) SyncOutput {
	return o
}

func (o SyncOutput) ToSyncPtrOutput() SyncPtrOutput {
	return o.ToSyncPtrOutputWithContext(context.Background())
}

func (o SyncOutput) ToSyncPtrOutputWithContext(ctx context.Context) SyncPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sync) *Sync {
		return &v
	}).(SyncPtrOutput)
}

type SyncPtrOutput struct{ *pulumi.OutputState }

func (SyncPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sync)(nil))
}

func (o SyncPtrOutput) ToSyncPtrOutput() SyncPtrOutput {
	return o
}

func (o SyncPtrOutput) ToSyncPtrOutputWithContext(ctx context.Context) SyncPtrOutput {
	return o
}

func (o SyncPtrOutput) Elem() SyncOutput {
	return o.ApplyT(func(v *Sync) Sync {
		if v != nil {
			return *v
		}
		var ret Sync
		return ret
	}).(SyncOutput)
}

type SyncArrayOutput struct{ *pulumi.OutputState }

func (SyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Sync)(nil))
}

func (o SyncArrayOutput) ToSyncArrayOutput() SyncArrayOutput {
	return o
}

func (o SyncArrayOutput) ToSyncArrayOutputWithContext(ctx context.Context) SyncArrayOutput {
	return o
}

func (o SyncArrayOutput) Index(i pulumi.IntInput) SyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Sync {
		return vs[0].([]Sync)[vs[1].(int)]
	}).(SyncOutput)
}

type SyncMapOutput struct{ *pulumi.OutputState }

func (SyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Sync)(nil))
}

func (o SyncMapOutput) ToSyncMapOutput() SyncMapOutput {
	return o
}

func (o SyncMapOutput) ToSyncMapOutputWithContext(ctx context.Context) SyncMapOutput {
	return o
}

func (o SyncMapOutput) MapIndex(k pulumi.StringInput) SyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Sync {
		return vs[0].(map[string]Sync)[vs[1].(string)]
	}).(SyncOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncInput)(nil)).Elem(), &Sync{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncPtrInput)(nil)).Elem(), &Sync{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncArrayInput)(nil)).Elem(), SyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncMapInput)(nil)).Elem(), SyncMap{})
	pulumi.RegisterOutputType(SyncOutput{})
	pulumi.RegisterOutputType(SyncPtrOutput{})
	pulumi.RegisterOutputType(SyncArrayOutput{})
	pulumi.RegisterOutputType(SyncMapOutput{})
}
