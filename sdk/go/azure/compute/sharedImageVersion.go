// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Version of a Shared Image within a Shared Image Gallery.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "search-api"
// 		existingImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Name:              &opt0,
// 			ResourceGroupName: "packerimages",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		existingSharedImage, err := compute.LookupSharedImage(ctx, &compute.LookupSharedImageArgs{
// 			Name:              "existing-image",
// 			GalleryName:       "existing_gallery",
// 			ResourceGroupName: "existing-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSharedImageVersion(ctx, "example", &compute.SharedImageVersionArgs{
// 			GalleryName:       pulumi.String(existingSharedImage.GalleryName),
// 			ImageName:         pulumi.String(existingSharedImage.Name),
// 			ResourceGroupName: pulumi.String(existingSharedImage.ResourceGroupName),
// 			Location:          pulumi.String(existingSharedImage.Location),
// 			ManagedImageId:    pulumi.String(existingImage.Id),
// 			TargetRegions: compute.SharedImageVersionTargetRegionArray{
// 				&compute.SharedImageVersionTargetRegionArgs{
// 					Name:                 pulumi.String(existingSharedImage.Location),
// 					RegionalReplicaCount: pulumi.Int(5),
// 					StorageAccountType:   pulumi.String("Standard_LRS"),
// 				},
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
// Shared Image Versions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/sharedImageVersion:SharedImageVersion version /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
// ```
type SharedImageVersion struct {
	pulumi.CustomResourceState

	// Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
	ExcludeFromLatest pulumi.BoolPtrOutput `pulumi:"excludeFromLatest"`
	// The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
	GalleryName pulumi.StringOutput `pulumi:"galleryName"`
	// The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	ManagedImageId pulumi.StringPtrOutput `pulumi:"managedImageId"`
	// The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	OsDiskSnapshotId pulumi.StringPtrOutput `pulumi:"osDiskSnapshotId"`
	// The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A collection of tags which should be applied to this resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// One or more `targetRegion` blocks as documented below.
	TargetRegions SharedImageVersionTargetRegionArrayOutput `pulumi:"targetRegions"`
}

// NewSharedImageVersion registers a new resource with the given unique name, arguments, and options.
func NewSharedImageVersion(ctx *pulumi.Context,
	name string, args *SharedImageVersionArgs, opts ...pulumi.ResourceOption) (*SharedImageVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetRegions == nil {
		return nil, errors.New("invalid value for required argument 'TargetRegions'")
	}
	var resource SharedImageVersion
	err := ctx.RegisterResource("azure:compute/sharedImageVersion:SharedImageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedImageVersion gets an existing SharedImageVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedImageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedImageVersionState, opts ...pulumi.ResourceOption) (*SharedImageVersion, error) {
	var resource SharedImageVersion
	err := ctx.ReadResource("azure:compute/sharedImageVersion:SharedImageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedImageVersion resources.
type sharedImageVersionState struct {
	// Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
	ExcludeFromLatest *bool `pulumi:"excludeFromLatest"`
	// The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
	GalleryName *string `pulumi:"galleryName"`
	// The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
	ImageName *string `pulumi:"imageName"`
	// The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	ManagedImageId *string `pulumi:"managedImageId"`
	// The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	OsDiskSnapshotId *string `pulumi:"osDiskSnapshotId"`
	// The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A collection of tags which should be applied to this resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `targetRegion` blocks as documented below.
	TargetRegions []SharedImageVersionTargetRegion `pulumi:"targetRegions"`
}

type SharedImageVersionState struct {
	// Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
	ExcludeFromLatest pulumi.BoolPtrInput
	// The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
	GalleryName pulumi.StringPtrInput
	// The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
	ImageName pulumi.StringPtrInput
	// The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	ManagedImageId pulumi.StringPtrInput
	// The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	OsDiskSnapshotId pulumi.StringPtrInput
	// The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A collection of tags which should be applied to this resource.
	Tags pulumi.StringMapInput
	// One or more `targetRegion` blocks as documented below.
	TargetRegions SharedImageVersionTargetRegionArrayInput
}

func (SharedImageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageVersionState)(nil)).Elem()
}

type sharedImageVersionArgs struct {
	// Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
	ExcludeFromLatest *bool `pulumi:"excludeFromLatest"`
	// The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
	GalleryName string `pulumi:"galleryName"`
	// The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
	ImageName string `pulumi:"imageName"`
	// The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	ManagedImageId *string `pulumi:"managedImageId"`
	// The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	OsDiskSnapshotId *string `pulumi:"osDiskSnapshotId"`
	// The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A collection of tags which should be applied to this resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `targetRegion` blocks as documented below.
	TargetRegions []SharedImageVersionTargetRegion `pulumi:"targetRegions"`
}

// The set of arguments for constructing a SharedImageVersion resource.
type SharedImageVersionArgs struct {
	// Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
	ExcludeFromLatest pulumi.BoolPtrInput
	// The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
	GalleryName pulumi.StringInput
	// The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
	ImageName pulumi.StringInput
	// The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The ID of the Managed Image or Virtual Machine ID which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	ManagedImageId pulumi.StringPtrInput
	// The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the OS disk snapshot which should be used for this Shared Image Version. Changing this forces a new resource to be created.
	OsDiskSnapshotId pulumi.StringPtrInput
	// The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A collection of tags which should be applied to this resource.
	Tags pulumi.StringMapInput
	// One or more `targetRegion` blocks as documented below.
	TargetRegions SharedImageVersionTargetRegionArrayInput
}

func (SharedImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageVersionArgs)(nil)).Elem()
}

type SharedImageVersionInput interface {
	pulumi.Input

	ToSharedImageVersionOutput() SharedImageVersionOutput
	ToSharedImageVersionOutputWithContext(ctx context.Context) SharedImageVersionOutput
}

func (*SharedImageVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedImageVersion)(nil)).Elem()
}

func (i *SharedImageVersion) ToSharedImageVersionOutput() SharedImageVersionOutput {
	return i.ToSharedImageVersionOutputWithContext(context.Background())
}

func (i *SharedImageVersion) ToSharedImageVersionOutputWithContext(ctx context.Context) SharedImageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedImageVersionOutput)
}

// SharedImageVersionArrayInput is an input type that accepts SharedImageVersionArray and SharedImageVersionArrayOutput values.
// You can construct a concrete instance of `SharedImageVersionArrayInput` via:
//
//          SharedImageVersionArray{ SharedImageVersionArgs{...} }
type SharedImageVersionArrayInput interface {
	pulumi.Input

	ToSharedImageVersionArrayOutput() SharedImageVersionArrayOutput
	ToSharedImageVersionArrayOutputWithContext(context.Context) SharedImageVersionArrayOutput
}

type SharedImageVersionArray []SharedImageVersionInput

func (SharedImageVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SharedImageVersion)(nil)).Elem()
}

func (i SharedImageVersionArray) ToSharedImageVersionArrayOutput() SharedImageVersionArrayOutput {
	return i.ToSharedImageVersionArrayOutputWithContext(context.Background())
}

func (i SharedImageVersionArray) ToSharedImageVersionArrayOutputWithContext(ctx context.Context) SharedImageVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedImageVersionArrayOutput)
}

// SharedImageVersionMapInput is an input type that accepts SharedImageVersionMap and SharedImageVersionMapOutput values.
// You can construct a concrete instance of `SharedImageVersionMapInput` via:
//
//          SharedImageVersionMap{ "key": SharedImageVersionArgs{...} }
type SharedImageVersionMapInput interface {
	pulumi.Input

	ToSharedImageVersionMapOutput() SharedImageVersionMapOutput
	ToSharedImageVersionMapOutputWithContext(context.Context) SharedImageVersionMapOutput
}

type SharedImageVersionMap map[string]SharedImageVersionInput

func (SharedImageVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SharedImageVersion)(nil)).Elem()
}

func (i SharedImageVersionMap) ToSharedImageVersionMapOutput() SharedImageVersionMapOutput {
	return i.ToSharedImageVersionMapOutputWithContext(context.Background())
}

func (i SharedImageVersionMap) ToSharedImageVersionMapOutputWithContext(ctx context.Context) SharedImageVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedImageVersionMapOutput)
}

type SharedImageVersionOutput struct{ *pulumi.OutputState }

func (SharedImageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedImageVersion)(nil)).Elem()
}

func (o SharedImageVersionOutput) ToSharedImageVersionOutput() SharedImageVersionOutput {
	return o
}

func (o SharedImageVersionOutput) ToSharedImageVersionOutputWithContext(ctx context.Context) SharedImageVersionOutput {
	return o
}

type SharedImageVersionArrayOutput struct{ *pulumi.OutputState }

func (SharedImageVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SharedImageVersion)(nil)).Elem()
}

func (o SharedImageVersionArrayOutput) ToSharedImageVersionArrayOutput() SharedImageVersionArrayOutput {
	return o
}

func (o SharedImageVersionArrayOutput) ToSharedImageVersionArrayOutputWithContext(ctx context.Context) SharedImageVersionArrayOutput {
	return o
}

func (o SharedImageVersionArrayOutput) Index(i pulumi.IntInput) SharedImageVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SharedImageVersion {
		return vs[0].([]*SharedImageVersion)[vs[1].(int)]
	}).(SharedImageVersionOutput)
}

type SharedImageVersionMapOutput struct{ *pulumi.OutputState }

func (SharedImageVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SharedImageVersion)(nil)).Elem()
}

func (o SharedImageVersionMapOutput) ToSharedImageVersionMapOutput() SharedImageVersionMapOutput {
	return o
}

func (o SharedImageVersionMapOutput) ToSharedImageVersionMapOutputWithContext(ctx context.Context) SharedImageVersionMapOutput {
	return o
}

func (o SharedImageVersionMapOutput) MapIndex(k pulumi.StringInput) SharedImageVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SharedImageVersion {
		return vs[0].(map[string]*SharedImageVersion)[vs[1].(string)]
	}).(SharedImageVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SharedImageVersionInput)(nil)).Elem(), &SharedImageVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedImageVersionArrayInput)(nil)).Elem(), SharedImageVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedImageVersionMapInput)(nil)).Elem(), SharedImageVersionMap{})
	pulumi.RegisterOutputType(SharedImageVersionOutput{})
	pulumi.RegisterOutputType(SharedImageVersionArrayOutput{})
	pulumi.RegisterOutputType(SharedImageVersionMapOutput{})
}
