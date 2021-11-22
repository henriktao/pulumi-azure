// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Managed Application Definition.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/managedapplication"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = managedapplication.NewDefinition(ctx, "exampleDefinition", &managedapplication.DefinitionArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			LockLevel:         pulumi.String("ReadOnly"),
// 			PackageFileUri:    pulumi.String("https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip"),
// 			DisplayName:       pulumi.String("TestManagedApplicationDefinition"),
// 			Description:       pulumi.String("Test Managed Application Definition"),
// 			Authorizations: managedapplication.DefinitionAuthorizationArray{
// 				&managedapplication.DefinitionAuthorizationArgs{
// 					ServicePrincipalId: pulumi.String(current.ObjectId),
// 					RoleDefinitionId:   pulumi.String("a094b430-dad3-424d-ae58-13f72fd72591"),
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
// Managed Application Definition can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:managedapplication/definition:Definition example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applicationDefinitions/appDefinition1
// ```
type Definition struct {
	pulumi.CustomResourceState

	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayOutput `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrOutput `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringOutput `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrOutput `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrOutput `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrOutput `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.LockLevel == nil {
		return nil, errors.New("invalid value for required argument 'LockLevel'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Definition
	err := ctx.RegisterResource("azure:managedapplication/definition:Definition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefinition gets an existing Definition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefinitionState, opts ...pulumi.ResourceOption) (*Definition, error) {
	var resource Definition
	err := ctx.ReadResource("azure:managedapplication/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// One or more `authorization` block defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition *string `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description *string `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel *string `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate *string `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled *bool `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DefinitionState struct {
	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrInput
	// Specifies the managed application definition description.
	Description pulumi.StringPtrInput
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringPtrInput
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrInput
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrInput
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// One or more `authorization` block defined below.
	Authorizations []DefinitionAuthorization `pulumi:"authorizations"`
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition *string `pulumi:"createUiDefinition"`
	// Specifies the managed application definition description.
	Description *string `pulumi:"description"`
	// Specifies the managed application definition display name.
	DisplayName string `pulumi:"displayName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel string `pulumi:"lockLevel"`
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate *string `pulumi:"mainTemplate"`
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Is the package enabled? Defaults to `true`.
	PackageEnabled *bool `pulumi:"packageEnabled"`
	// Specifies the managed application definition package file Uri.
	PackageFileUri *string `pulumi:"packageFileUri"`
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// One or more `authorization` block defined below.
	Authorizations DefinitionAuthorizationArrayInput
	// Specifies the `createUiDefinition` json for the backing template with `Microsoft.Solutions/applications` resource.
	CreateUiDefinition pulumi.StringPtrInput
	// Specifies the managed application definition description.
	Description pulumi.StringPtrInput
	// Specifies the managed application definition display name.
	DisplayName pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the managed application lock level. Valid values include `CanNotDelete`, `None`, `ReadOnly`. Changing this forces a new resource to be created.
	LockLevel pulumi.StringInput
	// Specifies the inline main template json which has resources to be provisioned.
	MainTemplate pulumi.StringPtrInput
	// Specifies the name of the Managed Application Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Is the package enabled? Defaults to `true`.
	PackageEnabled pulumi.BoolPtrInput
	// Specifies the managed application definition package file Uri.
	PackageFileUri pulumi.StringPtrInput
	// The name of the Resource Group where the Managed Application Definition should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionArgs)(nil)).Elem()
}

type DefinitionInput interface {
	pulumi.Input

	ToDefinitionOutput() DefinitionOutput
	ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput
}

func (*Definition) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (i *Definition) ToDefinitionOutput() DefinitionOutput {
	return i.ToDefinitionOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionOutput)
}

func (i *Definition) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *Definition) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPtrOutput)
}

type DefinitionPtrInput interface {
	pulumi.Input

	ToDefinitionPtrOutput() DefinitionPtrOutput
	ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput
}

type definitionPtrType DefinitionArgs

func (*definitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (i *definitionPtrType) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return i.ToDefinitionPtrOutputWithContext(context.Background())
}

func (i *definitionPtrType) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPtrOutput)
}

// DefinitionArrayInput is an input type that accepts DefinitionArray and DefinitionArrayOutput values.
// You can construct a concrete instance of `DefinitionArrayInput` via:
//
//          DefinitionArray{ DefinitionArgs{...} }
type DefinitionArrayInput interface {
	pulumi.Input

	ToDefinitionArrayOutput() DefinitionArrayOutput
	ToDefinitionArrayOutputWithContext(context.Context) DefinitionArrayOutput
}

type DefinitionArray []DefinitionInput

func (DefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Definition)(nil)).Elem()
}

func (i DefinitionArray) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return i.ToDefinitionArrayOutputWithContext(context.Background())
}

func (i DefinitionArray) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionArrayOutput)
}

// DefinitionMapInput is an input type that accepts DefinitionMap and DefinitionMapOutput values.
// You can construct a concrete instance of `DefinitionMapInput` via:
//
//          DefinitionMap{ "key": DefinitionArgs{...} }
type DefinitionMapInput interface {
	pulumi.Input

	ToDefinitionMapOutput() DefinitionMapOutput
	ToDefinitionMapOutputWithContext(context.Context) DefinitionMapOutput
}

type DefinitionMap map[string]DefinitionInput

func (DefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Definition)(nil)).Elem()
}

func (i DefinitionMap) ToDefinitionMapOutput() DefinitionMapOutput {
	return i.ToDefinitionMapOutputWithContext(context.Background())
}

func (i DefinitionMap) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionMapOutput)
}

type DefinitionOutput struct{ *pulumi.OutputState }

func (DefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Definition)(nil))
}

func (o DefinitionOutput) ToDefinitionOutput() DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionOutputWithContext(ctx context.Context) DefinitionOutput {
	return o
}

func (o DefinitionOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o.ToDefinitionPtrOutputWithContext(context.Background())
}

func (o DefinitionOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Definition) *Definition {
		return &v
	}).(DefinitionPtrOutput)
}

type DefinitionPtrOutput struct{ *pulumi.OutputState }

func (DefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Definition)(nil))
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutput() DefinitionPtrOutput {
	return o
}

func (o DefinitionPtrOutput) ToDefinitionPtrOutputWithContext(ctx context.Context) DefinitionPtrOutput {
	return o
}

func (o DefinitionPtrOutput) Elem() DefinitionOutput {
	return o.ApplyT(func(v *Definition) Definition {
		if v != nil {
			return *v
		}
		var ret Definition
		return ret
	}).(DefinitionOutput)
}

type DefinitionArrayOutput struct{ *pulumi.OutputState }

func (DefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Definition)(nil))
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutput() DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) ToDefinitionArrayOutputWithContext(ctx context.Context) DefinitionArrayOutput {
	return o
}

func (o DefinitionArrayOutput) Index(i pulumi.IntInput) DefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Definition {
		return vs[0].([]Definition)[vs[1].(int)]
	}).(DefinitionOutput)
}

type DefinitionMapOutput struct{ *pulumi.OutputState }

func (DefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Definition)(nil))
}

func (o DefinitionMapOutput) ToDefinitionMapOutput() DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) ToDefinitionMapOutputWithContext(ctx context.Context) DefinitionMapOutput {
	return o
}

func (o DefinitionMapOutput) MapIndex(k pulumi.StringInput) DefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Definition {
		return vs[0].(map[string]Definition)[vs[1].(string)]
	}).(DefinitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionInput)(nil)).Elem(), &Definition{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionPtrInput)(nil)).Elem(), &Definition{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionArrayInput)(nil)).Elem(), DefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionMapInput)(nil)).Elem(), DefinitionMap{})
	pulumi.RegisterOutputType(DefinitionOutput{})
	pulumi.RegisterOutputType(DefinitionPtrOutput{})
	pulumi.RegisterOutputType(DefinitionArrayOutput{})
	pulumi.RegisterOutputType(DefinitionMapOutput{})
}
