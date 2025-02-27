// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IotHub Device Provisioning Service Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/base64"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func filebase64OrPanic(path string) pulumi.StringPtrInput {
// 	if fileData, err := ioutil.ReadFile(path); err == nil {
// 		return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
// 	} else {
// 		panic(err.Error())
// 	}
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleIotHubDps, err := iot.NewIotHubDps(ctx, "exampleIotHubDps", &iot.IotHubDpsArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IotHubDpsSkuArgs{
// 				Name:     pulumi.String("S1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewIotHubCertificate(ctx, "exampleIotHubCertificate", &iot.IotHubCertificateArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			IotDpsName:         exampleIotHubDps.Name,
// 			CertificateContent: filebase64OrPanic("example.cer"),
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
// IoTHub Device Provisioning Service Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/iotHubCertificate:IotHubCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example/certificates/example
// ```
type IotHubCertificate struct {
	pulumi.CustomResourceState

	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringOutput `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringOutput `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIotHubCertificate registers a new resource with the given unique name, arguments, and options.
func NewIotHubCertificate(ctx *pulumi.Context,
	name string, args *IotHubCertificateArgs, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateContent == nil {
		return nil, errors.New("invalid value for required argument 'CertificateContent'")
	}
	if args.IotDpsName == nil {
		return nil, errors.New("invalid value for required argument 'IotDpsName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IotHubCertificate
	err := ctx.RegisterResource("azure:iot/iotHubCertificate:IotHubCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubCertificate gets an existing IotHubCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubCertificateState, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	var resource IotHubCertificate
	err := ctx.ReadResource("azure:iot/iotHubCertificate:IotHubCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubCertificate resources.
type iotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent *string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName *string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringPtrInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringPtrInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IotHubCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateState)(nil)).Elem()
}

type iotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IotHubCertificate resource.
type IotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IotHubCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateArgs)(nil)).Elem()
}

type IotHubCertificateInput interface {
	pulumi.Input

	ToIotHubCertificateOutput() IotHubCertificateOutput
	ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput
}

func (*IotHubCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubCertificate)(nil)).Elem()
}

func (i *IotHubCertificate) ToIotHubCertificateOutput() IotHubCertificateOutput {
	return i.ToIotHubCertificateOutputWithContext(context.Background())
}

func (i *IotHubCertificate) ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubCertificateOutput)
}

// IotHubCertificateArrayInput is an input type that accepts IotHubCertificateArray and IotHubCertificateArrayOutput values.
// You can construct a concrete instance of `IotHubCertificateArrayInput` via:
//
//          IotHubCertificateArray{ IotHubCertificateArgs{...} }
type IotHubCertificateArrayInput interface {
	pulumi.Input

	ToIotHubCertificateArrayOutput() IotHubCertificateArrayOutput
	ToIotHubCertificateArrayOutputWithContext(context.Context) IotHubCertificateArrayOutput
}

type IotHubCertificateArray []IotHubCertificateInput

func (IotHubCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHubCertificate)(nil)).Elem()
}

func (i IotHubCertificateArray) ToIotHubCertificateArrayOutput() IotHubCertificateArrayOutput {
	return i.ToIotHubCertificateArrayOutputWithContext(context.Background())
}

func (i IotHubCertificateArray) ToIotHubCertificateArrayOutputWithContext(ctx context.Context) IotHubCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubCertificateArrayOutput)
}

// IotHubCertificateMapInput is an input type that accepts IotHubCertificateMap and IotHubCertificateMapOutput values.
// You can construct a concrete instance of `IotHubCertificateMapInput` via:
//
//          IotHubCertificateMap{ "key": IotHubCertificateArgs{...} }
type IotHubCertificateMapInput interface {
	pulumi.Input

	ToIotHubCertificateMapOutput() IotHubCertificateMapOutput
	ToIotHubCertificateMapOutputWithContext(context.Context) IotHubCertificateMapOutput
}

type IotHubCertificateMap map[string]IotHubCertificateInput

func (IotHubCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHubCertificate)(nil)).Elem()
}

func (i IotHubCertificateMap) ToIotHubCertificateMapOutput() IotHubCertificateMapOutput {
	return i.ToIotHubCertificateMapOutputWithContext(context.Background())
}

func (i IotHubCertificateMap) ToIotHubCertificateMapOutputWithContext(ctx context.Context) IotHubCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubCertificateMapOutput)
}

type IotHubCertificateOutput struct{ *pulumi.OutputState }

func (IotHubCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubCertificate)(nil)).Elem()
}

func (o IotHubCertificateOutput) ToIotHubCertificateOutput() IotHubCertificateOutput {
	return o
}

func (o IotHubCertificateOutput) ToIotHubCertificateOutputWithContext(ctx context.Context) IotHubCertificateOutput {
	return o
}

type IotHubCertificateArrayOutput struct{ *pulumi.OutputState }

func (IotHubCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHubCertificate)(nil)).Elem()
}

func (o IotHubCertificateArrayOutput) ToIotHubCertificateArrayOutput() IotHubCertificateArrayOutput {
	return o
}

func (o IotHubCertificateArrayOutput) ToIotHubCertificateArrayOutputWithContext(ctx context.Context) IotHubCertificateArrayOutput {
	return o
}

func (o IotHubCertificateArrayOutput) Index(i pulumi.IntInput) IotHubCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IotHubCertificate {
		return vs[0].([]*IotHubCertificate)[vs[1].(int)]
	}).(IotHubCertificateOutput)
}

type IotHubCertificateMapOutput struct{ *pulumi.OutputState }

func (IotHubCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHubCertificate)(nil)).Elem()
}

func (o IotHubCertificateMapOutput) ToIotHubCertificateMapOutput() IotHubCertificateMapOutput {
	return o
}

func (o IotHubCertificateMapOutput) ToIotHubCertificateMapOutputWithContext(ctx context.Context) IotHubCertificateMapOutput {
	return o
}

func (o IotHubCertificateMapOutput) MapIndex(k pulumi.StringInput) IotHubCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IotHubCertificate {
		return vs[0].(map[string]*IotHubCertificate)[vs[1].(string)]
	}).(IotHubCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubCertificateInput)(nil)).Elem(), &IotHubCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubCertificateArrayInput)(nil)).Elem(), IotHubCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubCertificateMapInput)(nil)).Elem(), IotHubCertificateMap{})
	pulumi.RegisterOutputType(IotHubCertificateOutput{})
	pulumi.RegisterOutputType(IotHubCertificateArrayOutput{})
	pulumi.RegisterOutputType(IotHubCertificateMapOutput{})
}
