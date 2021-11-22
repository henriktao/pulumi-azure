// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Extension for a Virtual Machine Scale Set.
//
// > **NOTE:** This resource is not intended to be used with the `compute.ScaleSet` resource - instead it's intended for this to be used with the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLinuxVirtualMachineScaleSet, err := compute.NewLinuxVirtualMachineScaleSet(ctx, "exampleLinuxVirtualMachineScaleSet", nil)
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"commandToExecute": fmt.Sprintf("%v%v%v", "echo ", "$", "HOSTNAME"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err = compute.NewVirtualMachineScaleSetExtension(ctx, "exampleVirtualMachineScaleSetExtension", &compute.VirtualMachineScaleSetExtensionArgs{
// 			VirtualMachineScaleSetId: exampleLinuxVirtualMachineScaleSet.ID(),
// 			Publisher:                pulumi.String("Microsoft.Azure.Extensions"),
// 			Type:                     pulumi.String("CustomScript"),
// 			TypeHandlerVersion:       pulumi.String("2.0"),
// 			Settings:                 pulumi.String(json0),
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
// Virtual Machine Scale Set Extensions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleSet1/extensions/extension1
// ```
type VirtualMachineScaleSetExtension struct {
	pulumi.CustomResourceState

	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
	AutomaticUpgradeEnabled pulumi.BoolPtrOutput `pulumi:"automaticUpgradeEnabled"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrOutput `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayOutput `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringOutput `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrOutput `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringOutput `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringOutput `pulumi:"virtualMachineScaleSetId"`
}

// NewVirtualMachineScaleSetExtension registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Publisher == nil {
		return nil, errors.New("invalid value for required argument 'Publisher'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.TypeHandlerVersion == nil {
		return nil, errors.New("invalid value for required argument 'TypeHandlerVersion'")
	}
	if args.VirtualMachineScaleSetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineScaleSetId'")
	}
	var resource VirtualMachineScaleSetExtension
	err := ctx.RegisterResource("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetExtension gets an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	var resource VirtualMachineScaleSetExtension
	err := ctx.ReadResource("azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetExtension resources.
type virtualMachineScaleSetExtensionState struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
	AutomaticUpgradeEnabled *bool `pulumi:"automaticUpgradeEnabled"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher *string `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings *string `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId *string `pulumi:"virtualMachineScaleSetId"`
}

type VirtualMachineScaleSetExtensionState struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
	AutomaticUpgradeEnabled pulumi.BoolPtrInput
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrInput
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringPtrInput
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrInput
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringPtrInput
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringPtrInput
}

func (VirtualMachineScaleSetExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetExtensionArgs struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
	AutomaticUpgradeEnabled *bool `pulumi:"automaticUpgradeEnabled"`
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions []string `pulumi:"provisionAfterExtensions"`
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher string `pulumi:"publisher"`
	// A JSON String which specifies Settings for the Extension.
	Settings *string `pulumi:"settings"`
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion string `pulumi:"typeHandlerVersion"`
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId string `pulumi:"virtualMachineScaleSetId"`
}

// The set of arguments for constructing a VirtualMachineScaleSetExtension resource.
type VirtualMachineScaleSetExtensionArgs struct {
	// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
	AutomaticUpgradeEnabled pulumi.BoolPtrInput
	// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
	ForceUpdateTag pulumi.StringPtrInput
	// The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
	ProtectedSettings pulumi.StringPtrInput
	// An ordered list of Extension names which this should be provisioned after.
	ProvisionAfterExtensions pulumi.StringArrayInput
	// Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
	Publisher pulumi.StringInput
	// A JSON String which specifies Settings for the Extension.
	Settings pulumi.StringPtrInput
	// Specifies the Type of the Extension. Changing this forces a new resource to be created.
	Type pulumi.StringInput
	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringInput
	// The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
	VirtualMachineScaleSetId pulumi.StringInput
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionArgs)(nil)).Elem()
}

type VirtualMachineScaleSetExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput
	ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput
}

func (*VirtualMachineScaleSetExtension) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtension)(nil))
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return i.ToVirtualMachineScaleSetExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionOutput)
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionPtrOutput() VirtualMachineScaleSetExtensionPtrOutput {
	return i.ToVirtualMachineScaleSetExtensionPtrOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionPtrOutput)
}

type VirtualMachineScaleSetExtensionPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionPtrOutput() VirtualMachineScaleSetExtensionPtrOutput
	ToVirtualMachineScaleSetExtensionPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionPtrOutput
}

type virtualMachineScaleSetExtensionPtrType VirtualMachineScaleSetExtensionArgs

func (*virtualMachineScaleSetExtensionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil))
}

func (i *virtualMachineScaleSetExtensionPtrType) ToVirtualMachineScaleSetExtensionPtrOutput() VirtualMachineScaleSetExtensionPtrOutput {
	return i.ToVirtualMachineScaleSetExtensionPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetExtensionPtrType) ToVirtualMachineScaleSetExtensionPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionPtrOutput)
}

// VirtualMachineScaleSetExtensionArrayInput is an input type that accepts VirtualMachineScaleSetExtensionArray and VirtualMachineScaleSetExtensionArrayOutput values.
// You can construct a concrete instance of `VirtualMachineScaleSetExtensionArrayInput` via:
//
//          VirtualMachineScaleSetExtensionArray{ VirtualMachineScaleSetExtensionArgs{...} }
type VirtualMachineScaleSetExtensionArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput
	ToVirtualMachineScaleSetExtensionArrayOutputWithContext(context.Context) VirtualMachineScaleSetExtensionArrayOutput
}

type VirtualMachineScaleSetExtensionArray []VirtualMachineScaleSetExtensionInput

func (VirtualMachineScaleSetExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionArray) ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput {
	return i.ToVirtualMachineScaleSetExtensionArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionArray) ToVirtualMachineScaleSetExtensionArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionArrayOutput)
}

// VirtualMachineScaleSetExtensionMapInput is an input type that accepts VirtualMachineScaleSetExtensionMap and VirtualMachineScaleSetExtensionMapOutput values.
// You can construct a concrete instance of `VirtualMachineScaleSetExtensionMapInput` via:
//
//          VirtualMachineScaleSetExtensionMap{ "key": VirtualMachineScaleSetExtensionArgs{...} }
type VirtualMachineScaleSetExtensionMapInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionMapOutput() VirtualMachineScaleSetExtensionMapOutput
	ToVirtualMachineScaleSetExtensionMapOutputWithContext(context.Context) VirtualMachineScaleSetExtensionMapOutput
}

type VirtualMachineScaleSetExtensionMap map[string]VirtualMachineScaleSetExtensionInput

func (VirtualMachineScaleSetExtensionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionMap) ToVirtualMachineScaleSetExtensionMapOutput() VirtualMachineScaleSetExtensionMapOutput {
	return i.ToVirtualMachineScaleSetExtensionMapOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionMap) ToVirtualMachineScaleSetExtensionMapOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionMapOutput)
}

type VirtualMachineScaleSetExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtension)(nil))
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionPtrOutput() VirtualMachineScaleSetExtensionPtrOutput {
	return o.ToVirtualMachineScaleSetExtensionPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetExtension) *VirtualMachineScaleSetExtension {
		return &v
	}).(VirtualMachineScaleSetExtensionPtrOutput)
}

type VirtualMachineScaleSetExtensionPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil))
}

func (o VirtualMachineScaleSetExtensionPtrOutput) ToVirtualMachineScaleSetExtensionPtrOutput() VirtualMachineScaleSetExtensionPtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionPtrOutput) ToVirtualMachineScaleSetExtensionPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionPtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionPtrOutput) Elem() VirtualMachineScaleSetExtensionOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) VirtualMachineScaleSetExtension {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtension
		return ret
	}).(VirtualMachineScaleSetExtensionOutput)
}

type VirtualMachineScaleSetExtensionArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtension)(nil))
}

func (o VirtualMachineScaleSetExtensionArrayOutput) ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionArrayOutput) ToVirtualMachineScaleSetExtensionArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtension {
		return vs[0].([]VirtualMachineScaleSetExtension)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionOutput)
}

type VirtualMachineScaleSetExtensionMapOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualMachineScaleSetExtension)(nil))
}

func (o VirtualMachineScaleSetExtensionMapOutput) ToVirtualMachineScaleSetExtensionMapOutput() VirtualMachineScaleSetExtensionMapOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionMapOutput) ToVirtualMachineScaleSetExtensionMapOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionMapOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionMapOutput) MapIndex(k pulumi.StringInput) VirtualMachineScaleSetExtensionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtension {
		return vs[0].(map[string]VirtualMachineScaleSetExtension)[vs[1].(string)]
	}).(VirtualMachineScaleSetExtensionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineScaleSetExtensionInput)(nil)).Elem(), &VirtualMachineScaleSetExtension{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineScaleSetExtensionPtrInput)(nil)).Elem(), &VirtualMachineScaleSetExtension{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineScaleSetExtensionArrayInput)(nil)).Elem(), VirtualMachineScaleSetExtensionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineScaleSetExtensionMapInput)(nil)).Elem(), VirtualMachineScaleSetExtensionMap{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionMapOutput{})
}
