// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a custom virtual machine image that can be used to create virtual machines.
//
// ## Example Usage
// ### Creating From VHD
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = compute.NewImage(ctx, "exampleImage", &compute.ImageArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			OsDisk: &compute.ImageOsDiskArgs{
// 				OsType:  pulumi.String("Linux"),
// 				OsState: pulumi.String("Generalized"),
// 				BlobUri: pulumi.String("{blob_uri}"),
// 				SizeGb:  pulumi.Int(30),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Creating From Virtual Machine (VM Must Be Generalized Beforehand)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		_, err = compute.NewImage(ctx, "exampleImage", &compute.ImageArgs{
// 			Location:               pulumi.String("West US"),
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			SourceVirtualMachineId: pulumi.String("{vm_id}"),
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
// Images can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/image:Image example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/images/image1
// ```
type Image struct {
	pulumi.CustomResourceState

	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayOutput `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrOutput `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrOutput `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrOutput `pulumi:"zoneResilient"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Image
	err := ctx.RegisterResource("azure:compute/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azure:compute/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks []ImageDataDisk `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk *ImageOsDisk `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId *string `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient *bool `pulumi:"zoneResilient"`
}

type ImageState struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayInput
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrInput
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks []ImageDataDisk `pulumi:"dataDisks"`
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `osDisk` elements as defined below.
	OsDisk *ImageOsDisk `pulumi:"osDisk"`
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId *string `pulumi:"sourceVirtualMachineId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient *bool `pulumi:"zoneResilient"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// One or more `dataDisk` elements as defined below.
	DataDisks ImageDataDiskArrayInput
	// The HyperVGenerationType of the VirtualMachine created from the image as `V1`, `V2`. The default is `V1`.
	HyperVGeneration pulumi.StringPtrInput
	// Specified the supported Azure location where the resource exists.
	// Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the image. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `osDisk` elements as defined below.
	OsDisk ImageOsDiskPtrInput
	// The name of the resource group in which to create
	// the image. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Virtual Machine ID from which to create the image.
	SourceVirtualMachineId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is zone resiliency enabled?  Defaults to `false`.  Changing this forces a new resource to be created.
	ZoneResilient pulumi.BoolPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

func (i *Image) ToImagePtrOutput() ImagePtrOutput {
	return i.ToImagePtrOutputWithContext(context.Background())
}

func (i *Image) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePtrOutput)
}

type ImagePtrInput interface {
	pulumi.Input

	ToImagePtrOutput() ImagePtrOutput
	ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput
}

type imagePtrType ImageArgs

func (*imagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil))
}

func (i *imagePtrType) ToImagePtrOutput() ImagePtrOutput {
	return i.ToImagePtrOutputWithContext(context.Background())
}

func (i *imagePtrType) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePtrOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//          ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Image)(nil))
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//          ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Image)(nil))
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) ToImagePtrOutput() ImagePtrOutput {
	return o.ToImagePtrOutputWithContext(context.Background())
}

func (o ImageOutput) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return o.ApplyT(func(v Image) *Image {
		return &v
	}).(ImagePtrOutput)
}

type ImagePtrOutput struct {
	*pulumi.OutputState
}

func (ImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil))
}

func (o ImagePtrOutput) ToImagePtrOutput() ImagePtrOutput {
	return o
}

func (o ImagePtrOutput) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return o
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Image)(nil))
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Image {
		return vs[0].([]Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Image)(nil))
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Image {
		return vs[0].(map[string]Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImagePtrOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
