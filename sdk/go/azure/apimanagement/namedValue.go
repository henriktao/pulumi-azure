// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Named Value.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
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
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("pub1"),
// 			PublisherEmail:    pulumi.String("pub1@email.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewNamedValue(ctx, "exampleNamedValue", &apimanagement.NamedValueArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			DisplayName:       pulumi.String("ExampleProperty"),
// 			Value:             pulumi.String("Example Value"),
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
// API Management Properties can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/namedValue:NamedValue example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
// ```
type NamedValue struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrOutput `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewNamedValue registers a new resource with the given unique name, arguments, and options.
func NewNamedValue(ctx *pulumi.Context,
	name string, args *NamedValueArgs, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource NamedValue
	err := ctx.RegisterResource("azure:apimanagement/namedValue:NamedValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedValue gets an existing NamedValue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedValueState, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	var resource NamedValue
	err := ctx.ReadResource("azure:apimanagement/namedValue:NamedValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedValue resources.
type namedValueState struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value *string `pulumi:"value"`
}

type NamedValueState struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringPtrInput
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayInput
	// The value of this API Management Named Value.
	Value pulumi.StringPtrInput
}

func (NamedValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueState)(nil)).Elem()
}

type namedValueArgs struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The display name of this API Management Named Value.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret *bool `pulumi:"secret"`
	// A list of tags to be applied to the API Management Named Value.
	Tags []string `pulumi:"tags"`
	// The value of this API Management Named Value.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a NamedValue resource.
type NamedValueArgs struct {
	// The name of the API Management Service in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The display name of this API Management Named Value.
	DisplayName pulumi.StringInput
	// The name of the API Management Named Value. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Named Value should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies whether the API Management Named Value is secret. Valid values are `true` or `false`. The default value is `false`.
	Secret pulumi.BoolPtrInput
	// A list of tags to be applied to the API Management Named Value.
	Tags pulumi.StringArrayInput
	// The value of this API Management Named Value.
	Value pulumi.StringInput
}

func (NamedValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueArgs)(nil)).Elem()
}

type NamedValueInput interface {
	pulumi.Input

	ToNamedValueOutput() NamedValueOutput
	ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput
}

func (*NamedValue) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedValue)(nil))
}

func (i *NamedValue) ToNamedValueOutput() NamedValueOutput {
	return i.ToNamedValueOutputWithContext(context.Background())
}

func (i *NamedValue) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValueOutput)
}

func (i *NamedValue) ToNamedValuePtrOutput() NamedValuePtrOutput {
	return i.ToNamedValuePtrOutputWithContext(context.Background())
}

func (i *NamedValue) ToNamedValuePtrOutputWithContext(ctx context.Context) NamedValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValuePtrOutput)
}

type NamedValuePtrInput interface {
	pulumi.Input

	ToNamedValuePtrOutput() NamedValuePtrOutput
	ToNamedValuePtrOutputWithContext(ctx context.Context) NamedValuePtrOutput
}

type namedValuePtrType NamedValueArgs

func (*namedValuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedValue)(nil))
}

func (i *namedValuePtrType) ToNamedValuePtrOutput() NamedValuePtrOutput {
	return i.ToNamedValuePtrOutputWithContext(context.Background())
}

func (i *namedValuePtrType) ToNamedValuePtrOutputWithContext(ctx context.Context) NamedValuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValuePtrOutput)
}

// NamedValueArrayInput is an input type that accepts NamedValueArray and NamedValueArrayOutput values.
// You can construct a concrete instance of `NamedValueArrayInput` via:
//
//          NamedValueArray{ NamedValueArgs{...} }
type NamedValueArrayInput interface {
	pulumi.Input

	ToNamedValueArrayOutput() NamedValueArrayOutput
	ToNamedValueArrayOutputWithContext(context.Context) NamedValueArrayOutput
}

type NamedValueArray []NamedValueInput

func (NamedValueArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NamedValue)(nil))
}

func (i NamedValueArray) ToNamedValueArrayOutput() NamedValueArrayOutput {
	return i.ToNamedValueArrayOutputWithContext(context.Background())
}

func (i NamedValueArray) ToNamedValueArrayOutputWithContext(ctx context.Context) NamedValueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValueArrayOutput)
}

// NamedValueMapInput is an input type that accepts NamedValueMap and NamedValueMapOutput values.
// You can construct a concrete instance of `NamedValueMapInput` via:
//
//          NamedValueMap{ "key": NamedValueArgs{...} }
type NamedValueMapInput interface {
	pulumi.Input

	ToNamedValueMapOutput() NamedValueMapOutput
	ToNamedValueMapOutputWithContext(context.Context) NamedValueMapOutput
}

type NamedValueMap map[string]NamedValueInput

func (NamedValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NamedValue)(nil))
}

func (i NamedValueMap) ToNamedValueMapOutput() NamedValueMapOutput {
	return i.ToNamedValueMapOutputWithContext(context.Background())
}

func (i NamedValueMap) ToNamedValueMapOutputWithContext(ctx context.Context) NamedValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValueMapOutput)
}

type NamedValueOutput struct {
	*pulumi.OutputState
}

func (NamedValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedValue)(nil))
}

func (o NamedValueOutput) ToNamedValueOutput() NamedValueOutput {
	return o
}

func (o NamedValueOutput) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return o
}

func (o NamedValueOutput) ToNamedValuePtrOutput() NamedValuePtrOutput {
	return o.ToNamedValuePtrOutputWithContext(context.Background())
}

func (o NamedValueOutput) ToNamedValuePtrOutputWithContext(ctx context.Context) NamedValuePtrOutput {
	return o.ApplyT(func(v NamedValue) *NamedValue {
		return &v
	}).(NamedValuePtrOutput)
}

type NamedValuePtrOutput struct {
	*pulumi.OutputState
}

func (NamedValuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedValue)(nil))
}

func (o NamedValuePtrOutput) ToNamedValuePtrOutput() NamedValuePtrOutput {
	return o
}

func (o NamedValuePtrOutput) ToNamedValuePtrOutputWithContext(ctx context.Context) NamedValuePtrOutput {
	return o
}

type NamedValueArrayOutput struct{ *pulumi.OutputState }

func (NamedValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamedValue)(nil))
}

func (o NamedValueArrayOutput) ToNamedValueArrayOutput() NamedValueArrayOutput {
	return o
}

func (o NamedValueArrayOutput) ToNamedValueArrayOutputWithContext(ctx context.Context) NamedValueArrayOutput {
	return o
}

func (o NamedValueArrayOutput) Index(i pulumi.IntInput) NamedValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamedValue {
		return vs[0].([]NamedValue)[vs[1].(int)]
	}).(NamedValueOutput)
}

type NamedValueMapOutput struct{ *pulumi.OutputState }

func (NamedValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NamedValue)(nil))
}

func (o NamedValueMapOutput) ToNamedValueMapOutput() NamedValueMapOutput {
	return o
}

func (o NamedValueMapOutput) ToNamedValueMapOutputWithContext(ctx context.Context) NamedValueMapOutput {
	return o
}

func (o NamedValueMapOutput) MapIndex(k pulumi.StringInput) NamedValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NamedValue {
		return vs[0].(map[string]NamedValue)[vs[1].(string)]
	}).(NamedValueOutput)
}

func init() {
	pulumi.RegisterOutputType(NamedValueOutput{})
	pulumi.RegisterOutputType(NamedValuePtrOutput{})
	pulumi.RegisterOutputType(NamedValueArrayOutput{})
	pulumi.RegisterOutputType(NamedValueMapOutput{})
}
