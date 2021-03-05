// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 					NotActions: []interface{}{},
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
//  $ pulumi import azure:authorization/roleDefinition:RoleDefinition example "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/00000000-0000-0000-0000-000000000000|/subscriptions/00000000-0000-0000-0000-000000000000"
// ```
type RoleDefinition struct {
	pulumi.CustomResourceState

	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayOutput `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayOutput `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId pulumi.StringOutput `pulumi:"roleDefinitionResourceId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewRoleDefinition registers a new resource with the given unique name, arguments, and options.
func NewRoleDefinition(ctx *pulumi.Context,
	name string, args *RoleDefinitionArgs, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:role/definition:Definition"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleDefinition
	err := ctx.RegisterResource("azure:authorization/roleDefinition:RoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleDefinition gets an existing RoleDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleDefinitionState, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	var resource RoleDefinition
	err := ctx.ReadResource("azure:authorization/roleDefinition:RoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleDefinition resources.
type roleDefinitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []RoleDefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId *string `pulumi:"roleDefinitionResourceId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
}

type RoleDefinitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The Azure Resource Manager ID for the resource.
	RoleDefinitionResourceId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
}

func (RoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionState)(nil)).Elem()
}

type roleDefinitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []RoleDefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleDefinition resource.
type RoleDefinitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. It is recommended to use the first entry of the `assignableScopes`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
}

func (RoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionArgs)(nil)).Elem()
}

type RoleDefinitionInput interface {
	pulumi.Input

	ToRoleDefinitionOutput() RoleDefinitionOutput
	ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput
}

func (*RoleDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinition)(nil))
}

func (i *RoleDefinition) ToRoleDefinitionOutput() RoleDefinitionOutput {
	return i.ToRoleDefinitionOutputWithContext(context.Background())
}

func (i *RoleDefinition) ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionOutput)
}

func (i *RoleDefinition) ToRoleDefinitionPtrOutput() RoleDefinitionPtrOutput {
	return i.ToRoleDefinitionPtrOutputWithContext(context.Background())
}

func (i *RoleDefinition) ToRoleDefinitionPtrOutputWithContext(ctx context.Context) RoleDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionPtrOutput)
}

type RoleDefinitionPtrInput interface {
	pulumi.Input

	ToRoleDefinitionPtrOutput() RoleDefinitionPtrOutput
	ToRoleDefinitionPtrOutputWithContext(ctx context.Context) RoleDefinitionPtrOutput
}

type roleDefinitionPtrType RoleDefinitionArgs

func (*roleDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleDefinition)(nil))
}

func (i *roleDefinitionPtrType) ToRoleDefinitionPtrOutput() RoleDefinitionPtrOutput {
	return i.ToRoleDefinitionPtrOutputWithContext(context.Background())
}

func (i *roleDefinitionPtrType) ToRoleDefinitionPtrOutputWithContext(ctx context.Context) RoleDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionPtrOutput)
}

// RoleDefinitionArrayInput is an input type that accepts RoleDefinitionArray and RoleDefinitionArrayOutput values.
// You can construct a concrete instance of `RoleDefinitionArrayInput` via:
//
//          RoleDefinitionArray{ RoleDefinitionArgs{...} }
type RoleDefinitionArrayInput interface {
	pulumi.Input

	ToRoleDefinitionArrayOutput() RoleDefinitionArrayOutput
	ToRoleDefinitionArrayOutputWithContext(context.Context) RoleDefinitionArrayOutput
}

type RoleDefinitionArray []RoleDefinitionInput

func (RoleDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RoleDefinition)(nil))
}

func (i RoleDefinitionArray) ToRoleDefinitionArrayOutput() RoleDefinitionArrayOutput {
	return i.ToRoleDefinitionArrayOutputWithContext(context.Background())
}

func (i RoleDefinitionArray) ToRoleDefinitionArrayOutputWithContext(ctx context.Context) RoleDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionArrayOutput)
}

// RoleDefinitionMapInput is an input type that accepts RoleDefinitionMap and RoleDefinitionMapOutput values.
// You can construct a concrete instance of `RoleDefinitionMapInput` via:
//
//          RoleDefinitionMap{ "key": RoleDefinitionArgs{...} }
type RoleDefinitionMapInput interface {
	pulumi.Input

	ToRoleDefinitionMapOutput() RoleDefinitionMapOutput
	ToRoleDefinitionMapOutputWithContext(context.Context) RoleDefinitionMapOutput
}

type RoleDefinitionMap map[string]RoleDefinitionInput

func (RoleDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RoleDefinition)(nil))
}

func (i RoleDefinitionMap) ToRoleDefinitionMapOutput() RoleDefinitionMapOutput {
	return i.ToRoleDefinitionMapOutputWithContext(context.Background())
}

func (i RoleDefinitionMap) ToRoleDefinitionMapOutputWithContext(ctx context.Context) RoleDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionMapOutput)
}

type RoleDefinitionOutput struct {
	*pulumi.OutputState
}

func (RoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinition)(nil))
}

func (o RoleDefinitionOutput) ToRoleDefinitionOutput() RoleDefinitionOutput {
	return o
}

func (o RoleDefinitionOutput) ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput {
	return o
}

func (o RoleDefinitionOutput) ToRoleDefinitionPtrOutput() RoleDefinitionPtrOutput {
	return o.ToRoleDefinitionPtrOutputWithContext(context.Background())
}

func (o RoleDefinitionOutput) ToRoleDefinitionPtrOutputWithContext(ctx context.Context) RoleDefinitionPtrOutput {
	return o.ApplyT(func(v RoleDefinition) *RoleDefinition {
		return &v
	}).(RoleDefinitionPtrOutput)
}

type RoleDefinitionPtrOutput struct {
	*pulumi.OutputState
}

func (RoleDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleDefinition)(nil))
}

func (o RoleDefinitionPtrOutput) ToRoleDefinitionPtrOutput() RoleDefinitionPtrOutput {
	return o
}

func (o RoleDefinitionPtrOutput) ToRoleDefinitionPtrOutputWithContext(ctx context.Context) RoleDefinitionPtrOutput {
	return o
}

type RoleDefinitionArrayOutput struct{ *pulumi.OutputState }

func (RoleDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleDefinition)(nil))
}

func (o RoleDefinitionArrayOutput) ToRoleDefinitionArrayOutput() RoleDefinitionArrayOutput {
	return o
}

func (o RoleDefinitionArrayOutput) ToRoleDefinitionArrayOutputWithContext(ctx context.Context) RoleDefinitionArrayOutput {
	return o
}

func (o RoleDefinitionArrayOutput) Index(i pulumi.IntInput) RoleDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleDefinition {
		return vs[0].([]RoleDefinition)[vs[1].(int)]
	}).(RoleDefinitionOutput)
}

type RoleDefinitionMapOutput struct{ *pulumi.OutputState }

func (RoleDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RoleDefinition)(nil))
}

func (o RoleDefinitionMapOutput) ToRoleDefinitionMapOutput() RoleDefinitionMapOutput {
	return o
}

func (o RoleDefinitionMapOutput) ToRoleDefinitionMapOutputWithContext(ctx context.Context) RoleDefinitionMapOutput {
	return o
}

func (o RoleDefinitionMapOutput) MapIndex(k pulumi.StringInput) RoleDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RoleDefinition {
		return vs[0].(map[string]RoleDefinition)[vs[1].(string)]
	}).(RoleDefinitionOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleDefinitionOutput{})
	pulumi.RegisterOutputType(RoleDefinitionPtrOutput{})
	pulumi.RegisterOutputType(RoleDefinitionArrayOutput{})
	pulumi.RegisterOutputType(RoleDefinitionMapOutput{})
}
