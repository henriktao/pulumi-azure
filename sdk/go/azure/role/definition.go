// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package role

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewRoleDefinition(ctx, "example", &authorization.RoleDefinitionArgs{
// 			Scope:       pulumi.String(primary.Id),
// 			Description: pulumi.String("This is a custom role created"),
// 			Permissions: authorization.RoleDefinitionPermissionArray{
// 				&authorization.RoleDefinitionPermissionArgs{
// 					Actions: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					NotActions: pulumi.StringArray{},
// 				},
// 			},
// 			AssignableScopes: pulumi.StringArray{
// 				pulumi.String(primary.Id),
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
// Role Definitions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:role/definition:Definition example "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|/subscriptions/00000000-0000-0000-0000-000000000000"
// ```
//
// Deprecated: azure.role.Definition has been deprecated in favor of azure.authorization.RoleDefinition
type Definition struct {
	pulumi.CustomResourceState

	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayOutput `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions DefinitionPermissionArrayOutput `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId pulumi.StringOutput `pulumi:"roleDefinitionResourceId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource Definition
	err := ctx.RegisterResource("azure:role/definition:Definition", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:role/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []DefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId *string `pulumi:"roleDefinitionResourceId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
}

type DefinitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions DefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []DefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions DefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
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
	pulumi.RegisterOutputType(DefinitionOutput{})
	pulumi.RegisterOutputType(DefinitionPtrOutput{})
	pulumi.RegisterOutputType(DefinitionArrayOutput{})
	pulumi.RegisterOutputType(DefinitionMapOutput{})
}
