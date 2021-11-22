// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Databox Edge Device.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databoxedge"
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
// 		_, err = databoxedge.NewDevice(ctx, "exampleDevice", &databoxedge.DeviceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("EdgeP_Base-Standard"),
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
// Databox Edge Devices can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:databoxedge/device:Device example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1
// ```
type Device struct {
	pulumi.CustomResourceState

	// A `deviceProperties` block as defined below.
	DeviceProperties DeviceDevicePropertyArrayOutput `pulumi:"deviceProperties"`
	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `skuName`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource Device
	err := ctx.RegisterResource("azure:databoxedge/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azure:databoxedge/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// A `deviceProperties` block as defined below.
	DeviceProperties []DeviceDeviceProperty `pulumi:"deviceProperties"`
	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `skuName`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags map[string]string `pulumi:"tags"`
}

type DeviceState struct {
	// A `deviceProperties` block as defined below.
	DeviceProperties DeviceDevicePropertyArrayInput
	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `skuName`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags pulumi.StringMapInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `skuName`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The Azure Region where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Databox Edge Device. Changing this forces a new Databox Edge Device to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Databox Edge Device should exist. Changing this forces a new Databox Edge Device to be created.
	ResourceGroupName pulumi.StringInput
	// The `skuName` is comprised of two segments separated by a hyphen (e.g. `TEA_1Node_UPS_Heater-Standard`). The first segment of the `skuName` defines the `name` of the sku, possible values are `Gateway`, `EdgeMR_Mini`, `EdgeP_Base`, `EdgeP_High`, `EdgePR_Base`, `EdgePR_Base_UPS`, `GPU`, `RCA_Large`, `RCA_Small`, `RDC`, `TCA_Large`, `TCA_Small`, `TDC`, `TEA_1Node`, `TEA_1Node_UPS`, `TEA_1Node_Heater`, `TEA_1Node_UPS_Heater`, `TEA_4Node_Heater`, `TEA_4Node_UPS_Heater` or `TMA`. The second segment defines the `tier` of the `skuName`, possible values are `Standard`. For more information see the product documentation. Changing this forces a new Databox Edge Device to be created.
	SkuName pulumi.StringInput
	// A mapping of tags which should be assigned to the Databox Edge Device.
	Tags pulumi.StringMapInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((*Device)(nil))
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

func (i *Device) ToDevicePtrOutput() DevicePtrOutput {
	return i.ToDevicePtrOutputWithContext(context.Background())
}

func (i *Device) ToDevicePtrOutputWithContext(ctx context.Context) DevicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePtrOutput)
}

type DevicePtrInput interface {
	pulumi.Input

	ToDevicePtrOutput() DevicePtrOutput
	ToDevicePtrOutputWithContext(ctx context.Context) DevicePtrOutput
}

type devicePtrType DeviceArgs

func (*devicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil))
}

func (i *devicePtrType) ToDevicePtrOutput() DevicePtrOutput {
	return i.ToDevicePtrOutputWithContext(context.Background())
}

func (i *devicePtrType) ToDevicePtrOutputWithContext(ctx context.Context) DevicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePtrOutput)
}

// DeviceArrayInput is an input type that accepts DeviceArray and DeviceArrayOutput values.
// You can construct a concrete instance of `DeviceArrayInput` via:
//
//          DeviceArray{ DeviceArgs{...} }
type DeviceArrayInput interface {
	pulumi.Input

	ToDeviceArrayOutput() DeviceArrayOutput
	ToDeviceArrayOutputWithContext(context.Context) DeviceArrayOutput
}

type DeviceArray []DeviceInput

func (DeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (i DeviceArray) ToDeviceArrayOutput() DeviceArrayOutput {
	return i.ToDeviceArrayOutputWithContext(context.Background())
}

func (i DeviceArray) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceArrayOutput)
}

// DeviceMapInput is an input type that accepts DeviceMap and DeviceMapOutput values.
// You can construct a concrete instance of `DeviceMapInput` via:
//
//          DeviceMap{ "key": DeviceArgs{...} }
type DeviceMapInput interface {
	pulumi.Input

	ToDeviceMapOutput() DeviceMapOutput
	ToDeviceMapOutputWithContext(context.Context) DeviceMapOutput
}

type DeviceMap map[string]DeviceInput

func (DeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (i DeviceMap) ToDeviceMapOutput() DeviceMapOutput {
	return i.ToDeviceMapOutputWithContext(context.Background())
}

func (i DeviceMap) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceMapOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Device)(nil))
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func (o DeviceOutput) ToDevicePtrOutput() DevicePtrOutput {
	return o.ToDevicePtrOutputWithContext(context.Background())
}

func (o DeviceOutput) ToDevicePtrOutputWithContext(ctx context.Context) DevicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Device) *Device {
		return &v
	}).(DevicePtrOutput)
}

type DevicePtrOutput struct{ *pulumi.OutputState }

func (DevicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil))
}

func (o DevicePtrOutput) ToDevicePtrOutput() DevicePtrOutput {
	return o
}

func (o DevicePtrOutput) ToDevicePtrOutputWithContext(ctx context.Context) DevicePtrOutput {
	return o
}

func (o DevicePtrOutput) Elem() DeviceOutput {
	return o.ApplyT(func(v *Device) Device {
		if v != nil {
			return *v
		}
		var ret Device
		return ret
	}).(DeviceOutput)
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Device)(nil))
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Device {
		return vs[0].([]Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceMapOutput struct{ *pulumi.OutputState }

func (DeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Device)(nil))
}

func (o DeviceMapOutput) ToDeviceMapOutput() DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) MapIndex(k pulumi.StringInput) DeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Device {
		return vs[0].(map[string]Device)[vs[1].(string)]
	}).(DeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePtrInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceArrayInput)(nil)).Elem(), DeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceMapInput)(nil)).Elem(), DeviceMap{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DevicePtrOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceMapOutput{})
}
