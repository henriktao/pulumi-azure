// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Blob within a Storage Container.
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBlob(ctx, "exampleBlob", &storage.BlobArgs{
// 			StorageAccountName:   exampleAccount.Name,
// 			StorageContainerName: exampleContainer.Name,
// 			Type:                 pulumi.String("Block"),
// 			Source:               pulumi.NewFileAsset("some-local-file.zip"),
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
// Storage Blob's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/blob:Blob blob1 https://example.blob.core.windows.net/container/blob.vhd
// ```
type Blob struct {
	pulumi.CustomResourceState

	// The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
	AccessTier pulumi.StringOutput `pulumi:"accessTier"`
	// The MD5 sum of the blob contents. Cannot be defined if `sourceUri` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 pulumi.StringPtrOutput `pulumi:"contentMd5"`
	// The content type of the storage blob. Cannot be defined if `sourceUri` is defined. Defaults to `application/octet-stream`.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// A map of custom blob metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the storage blob. Must be unique within the storage container the blob is located.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
	Parallelism pulumi.IntPtrOutput `pulumi:"parallelism"`
	// Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
	Size pulumi.IntPtrOutput `pulumi:"size"`
	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `sourceContent` or `sourceUri` is specified.
	Source pulumi.AssetOrArchiveOutput `pulumi:"source"`
	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `sourceUri` is specified.
	SourceContent pulumi.StringPtrOutput `pulumi:"sourceContent"`
	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents
	// for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `sourceContent` is specified.
	SourceUri pulumi.StringPtrOutput `pulumi:"sourceUri"`
	// Specifies the storage account in which to create the storage container.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// The name of the storage container in which this blob should be created.
	StorageContainerName pulumi.StringOutput `pulumi:"storageContainerName"`
	// The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// The URL of the blob
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewBlob registers a new resource with the given unique name, arguments, and options.
func NewBlob(ctx *pulumi.Context,
	name string, args *BlobArgs, opts ...pulumi.ResourceOption) (*Blob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Blob
	err := ctx.RegisterResource("azure:storage/blob:Blob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlob gets an existing Blob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobState, opts ...pulumi.ResourceOption) (*Blob, error) {
	var resource Blob
	err := ctx.ReadResource("azure:storage/blob:Blob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Blob resources.
type blobState struct {
	// The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// The MD5 sum of the blob contents. Cannot be defined if `sourceUri` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 *string `pulumi:"contentMd5"`
	// The content type of the storage blob. Cannot be defined if `sourceUri` is defined. Defaults to `application/octet-stream`.
	ContentType *string `pulumi:"contentType"`
	// A map of custom blob metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the storage blob. Must be unique within the storage container the blob is located.
	Name *string `pulumi:"name"`
	// The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
	Parallelism *int `pulumi:"parallelism"`
	// Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
	Size *int `pulumi:"size"`
	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `sourceContent` or `sourceUri` is specified.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `sourceUri` is specified.
	SourceContent *string `pulumi:"sourceContent"`
	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents
	// for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `sourceContent` is specified.
	SourceUri *string `pulumi:"sourceUri"`
	// Specifies the storage account in which to create the storage container.
	// Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The name of the storage container in which this blob should be created.
	StorageContainerName *string `pulumi:"storageContainerName"`
	// The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The URL of the blob
	Url *string `pulumi:"url"`
}

type BlobState struct {
	// The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
	AccessTier pulumi.StringPtrInput
	// The MD5 sum of the blob contents. Cannot be defined if `sourceUri` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 pulumi.StringPtrInput
	// The content type of the storage blob. Cannot be defined if `sourceUri` is defined. Defaults to `application/octet-stream`.
	ContentType pulumi.StringPtrInput
	// A map of custom blob metadata.
	Metadata pulumi.StringMapInput
	// The name of the storage blob. Must be unique within the storage container the blob is located.
	Name pulumi.StringPtrInput
	// The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
	Parallelism pulumi.IntPtrInput
	// Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
	Size pulumi.IntPtrInput
	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `sourceContent` or `sourceUri` is specified.
	Source pulumi.AssetOrArchiveInput
	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `sourceUri` is specified.
	SourceContent pulumi.StringPtrInput
	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents
	// for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `sourceContent` is specified.
	SourceUri pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage container.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
	// The name of the storage container in which this blob should be created.
	StorageContainerName pulumi.StringPtrInput
	// The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The URL of the blob
	Url pulumi.StringPtrInput
}

func (BlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobState)(nil)).Elem()
}

type blobArgs struct {
	// The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// The MD5 sum of the blob contents. Cannot be defined if `sourceUri` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 *string `pulumi:"contentMd5"`
	// The content type of the storage blob. Cannot be defined if `sourceUri` is defined. Defaults to `application/octet-stream`.
	ContentType *string `pulumi:"contentType"`
	// A map of custom blob metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the storage blob. Must be unique within the storage container the blob is located.
	Name *string `pulumi:"name"`
	// The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
	Parallelism *int `pulumi:"parallelism"`
	// Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
	Size *int `pulumi:"size"`
	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `sourceContent` or `sourceUri` is specified.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `sourceUri` is specified.
	SourceContent *string `pulumi:"sourceContent"`
	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents
	// for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `sourceContent` is specified.
	SourceUri *string `pulumi:"sourceUri"`
	// Specifies the storage account in which to create the storage container.
	// Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the storage container in which this blob should be created.
	StorageContainerName string `pulumi:"storageContainerName"`
	// The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Blob resource.
type BlobArgs struct {
	// The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
	AccessTier pulumi.StringPtrInput
	// The MD5 sum of the blob contents. Cannot be defined if `sourceUri` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.
	ContentMd5 pulumi.StringPtrInput
	// The content type of the storage blob. Cannot be defined if `sourceUri` is defined. Defaults to `application/octet-stream`.
	ContentType pulumi.StringPtrInput
	// A map of custom blob metadata.
	Metadata pulumi.StringMapInput
	// The name of the storage blob. Must be unique within the storage container the blob is located.
	Name pulumi.StringPtrInput
	// The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
	Parallelism pulumi.IntPtrInput
	// Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
	Size pulumi.IntPtrInput
	// An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `sourceContent` or `sourceUri` is specified.
	Source pulumi.AssetOrArchiveInput
	// The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `sourceUri` is specified.
	SourceContent pulumi.StringPtrInput
	// The URI of an existing blob, or a file in the Azure File service, to use as the source contents
	// for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `sourceContent` is specified.
	SourceUri pulumi.StringPtrInput
	// Specifies the storage account in which to create the storage container.
	// Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
	// The name of the storage container in which this blob should be created.
	StorageContainerName pulumi.StringInput
	// The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
	Type pulumi.StringInput
}

func (BlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobArgs)(nil)).Elem()
}

type BlobInput interface {
	pulumi.Input

	ToBlobOutput() BlobOutput
	ToBlobOutputWithContext(ctx context.Context) BlobOutput
}

func (*Blob) ElementType() reflect.Type {
	return reflect.TypeOf((*Blob)(nil))
}

func (i *Blob) ToBlobOutput() BlobOutput {
	return i.ToBlobOutputWithContext(context.Background())
}

func (i *Blob) ToBlobOutputWithContext(ctx context.Context) BlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobOutput)
}

func (i *Blob) ToBlobPtrOutput() BlobPtrOutput {
	return i.ToBlobPtrOutputWithContext(context.Background())
}

func (i *Blob) ToBlobPtrOutputWithContext(ctx context.Context) BlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobPtrOutput)
}

type BlobPtrInput interface {
	pulumi.Input

	ToBlobPtrOutput() BlobPtrOutput
	ToBlobPtrOutputWithContext(ctx context.Context) BlobPtrOutput
}

type blobPtrType BlobArgs

func (*blobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Blob)(nil))
}

func (i *blobPtrType) ToBlobPtrOutput() BlobPtrOutput {
	return i.ToBlobPtrOutputWithContext(context.Background())
}

func (i *blobPtrType) ToBlobPtrOutputWithContext(ctx context.Context) BlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobPtrOutput)
}

// BlobArrayInput is an input type that accepts BlobArray and BlobArrayOutput values.
// You can construct a concrete instance of `BlobArrayInput` via:
//
//          BlobArray{ BlobArgs{...} }
type BlobArrayInput interface {
	pulumi.Input

	ToBlobArrayOutput() BlobArrayOutput
	ToBlobArrayOutputWithContext(context.Context) BlobArrayOutput
}

type BlobArray []BlobInput

func (BlobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Blob)(nil)).Elem()
}

func (i BlobArray) ToBlobArrayOutput() BlobArrayOutput {
	return i.ToBlobArrayOutputWithContext(context.Background())
}

func (i BlobArray) ToBlobArrayOutputWithContext(ctx context.Context) BlobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobArrayOutput)
}

// BlobMapInput is an input type that accepts BlobMap and BlobMapOutput values.
// You can construct a concrete instance of `BlobMapInput` via:
//
//          BlobMap{ "key": BlobArgs{...} }
type BlobMapInput interface {
	pulumi.Input

	ToBlobMapOutput() BlobMapOutput
	ToBlobMapOutputWithContext(context.Context) BlobMapOutput
}

type BlobMap map[string]BlobInput

func (BlobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Blob)(nil)).Elem()
}

func (i BlobMap) ToBlobMapOutput() BlobMapOutput {
	return i.ToBlobMapOutputWithContext(context.Background())
}

func (i BlobMap) ToBlobMapOutputWithContext(ctx context.Context) BlobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobMapOutput)
}

type BlobOutput struct{ *pulumi.OutputState }

func (BlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Blob)(nil))
}

func (o BlobOutput) ToBlobOutput() BlobOutput {
	return o
}

func (o BlobOutput) ToBlobOutputWithContext(ctx context.Context) BlobOutput {
	return o
}

func (o BlobOutput) ToBlobPtrOutput() BlobPtrOutput {
	return o.ToBlobPtrOutputWithContext(context.Background())
}

func (o BlobOutput) ToBlobPtrOutputWithContext(ctx context.Context) BlobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Blob) *Blob {
		return &v
	}).(BlobPtrOutput)
}

type BlobPtrOutput struct{ *pulumi.OutputState }

func (BlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Blob)(nil))
}

func (o BlobPtrOutput) ToBlobPtrOutput() BlobPtrOutput {
	return o
}

func (o BlobPtrOutput) ToBlobPtrOutputWithContext(ctx context.Context) BlobPtrOutput {
	return o
}

func (o BlobPtrOutput) Elem() BlobOutput {
	return o.ApplyT(func(v *Blob) Blob {
		if v != nil {
			return *v
		}
		var ret Blob
		return ret
	}).(BlobOutput)
}

type BlobArrayOutput struct{ *pulumi.OutputState }

func (BlobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Blob)(nil))
}

func (o BlobArrayOutput) ToBlobArrayOutput() BlobArrayOutput {
	return o
}

func (o BlobArrayOutput) ToBlobArrayOutputWithContext(ctx context.Context) BlobArrayOutput {
	return o
}

func (o BlobArrayOutput) Index(i pulumi.IntInput) BlobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Blob {
		return vs[0].([]Blob)[vs[1].(int)]
	}).(BlobOutput)
}

type BlobMapOutput struct{ *pulumi.OutputState }

func (BlobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Blob)(nil))
}

func (o BlobMapOutput) ToBlobMapOutput() BlobMapOutput {
	return o
}

func (o BlobMapOutput) ToBlobMapOutputWithContext(ctx context.Context) BlobMapOutput {
	return o
}

func (o BlobMapOutput) MapIndex(k pulumi.StringInput) BlobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Blob {
		return vs[0].(map[string]Blob)[vs[1].(string)]
	}).(BlobOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobOutput{})
	pulumi.RegisterOutputType(BlobPtrOutput{})
	pulumi.RegisterOutputType(BlobArrayOutput{})
	pulumi.RegisterOutputType(BlobMapOutput{})
}
