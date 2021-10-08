// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefinitionAuthorization struct {
	// The set of role definition ids which define all the permissions that the principal id can assign.
	DelegatedRoleDefinitionIds []string `pulumi:"delegatedRoleDefinitionIds"`
	// The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
	PrincipalDisplayName *string `pulumi:"principalDisplayName"`
	// Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
	PrincipalId string `pulumi:"principalId"`
	// The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

// DefinitionAuthorizationInput is an input type that accepts DefinitionAuthorizationArgs and DefinitionAuthorizationOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationInput` via:
//
//          DefinitionAuthorizationArgs{...}
type DefinitionAuthorizationInput interface {
	pulumi.Input

	ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput
	ToDefinitionAuthorizationOutputWithContext(context.Context) DefinitionAuthorizationOutput
}

type DefinitionAuthorizationArgs struct {
	// The set of role definition ids which define all the permissions that the principal id can assign.
	DelegatedRoleDefinitionIds pulumi.StringArrayInput `pulumi:"delegatedRoleDefinitionIds"`
	// The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
	PrincipalDisplayName pulumi.StringPtrInput `pulumi:"principalDisplayName"`
	// Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (DefinitionAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return i.ToDefinitionAuthorizationOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationOutput)
}

// DefinitionAuthorizationArrayInput is an input type that accepts DefinitionAuthorizationArray and DefinitionAuthorizationArrayOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationArrayInput` via:
//
//          DefinitionAuthorizationArray{ DefinitionAuthorizationArgs{...} }
type DefinitionAuthorizationArrayInput interface {
	pulumi.Input

	ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput
	ToDefinitionAuthorizationArrayOutputWithContext(context.Context) DefinitionAuthorizationArrayOutput
}

type DefinitionAuthorizationArray []DefinitionAuthorizationInput

func (DefinitionAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return i.ToDefinitionAuthorizationArrayOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationArrayOutput)
}

type DefinitionAuthorizationOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return o
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return o
}

// The set of role definition ids which define all the permissions that the principal id can assign.
func (o DefinitionAuthorizationOutput) DelegatedRoleDefinitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DefinitionAuthorization) []string { return v.DelegatedRoleDefinitionIds }).(pulumi.StringArrayOutput)
}

// The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
func (o DefinitionAuthorizationOutput) PrincipalDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefinitionAuthorization) *string { return v.PrincipalDisplayName }).(pulumi.StringPtrOutput)
}

// Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
func (o DefinitionAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
func (o DefinitionAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type DefinitionAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) Index(i pulumi.IntInput) DefinitionAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefinitionAuthorization {
		return vs[0].([]DefinitionAuthorization)[vs[1].(int)]
	}).(DefinitionAuthorizationOutput)
}

type DefinitionPlan struct {
	// The plan name of the marketplace offer.
	Name string `pulumi:"name"`
	// The product code of the plan.
	Product string `pulumi:"product"`
	// The publisher ID of the plan.
	Publisher string `pulumi:"publisher"`
	// The version of the plan.
	Version string `pulumi:"version"`
}

// DefinitionPlanInput is an input type that accepts DefinitionPlanArgs and DefinitionPlanOutput values.
// You can construct a concrete instance of `DefinitionPlanInput` via:
//
//          DefinitionPlanArgs{...}
type DefinitionPlanInput interface {
	pulumi.Input

	ToDefinitionPlanOutput() DefinitionPlanOutput
	ToDefinitionPlanOutputWithContext(context.Context) DefinitionPlanOutput
}

type DefinitionPlanArgs struct {
	// The plan name of the marketplace offer.
	Name pulumi.StringInput `pulumi:"name"`
	// The product code of the plan.
	Product pulumi.StringInput `pulumi:"product"`
	// The publisher ID of the plan.
	Publisher pulumi.StringInput `pulumi:"publisher"`
	// The version of the plan.
	Version pulumi.StringInput `pulumi:"version"`
}

func (DefinitionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionPlan)(nil)).Elem()
}

func (i DefinitionPlanArgs) ToDefinitionPlanOutput() DefinitionPlanOutput {
	return i.ToDefinitionPlanOutputWithContext(context.Background())
}

func (i DefinitionPlanArgs) ToDefinitionPlanOutputWithContext(ctx context.Context) DefinitionPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPlanOutput)
}

func (i DefinitionPlanArgs) ToDefinitionPlanPtrOutput() DefinitionPlanPtrOutput {
	return i.ToDefinitionPlanPtrOutputWithContext(context.Background())
}

func (i DefinitionPlanArgs) ToDefinitionPlanPtrOutputWithContext(ctx context.Context) DefinitionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPlanOutput).ToDefinitionPlanPtrOutputWithContext(ctx)
}

// DefinitionPlanPtrInput is an input type that accepts DefinitionPlanArgs, DefinitionPlanPtr and DefinitionPlanPtrOutput values.
// You can construct a concrete instance of `DefinitionPlanPtrInput` via:
//
//          DefinitionPlanArgs{...}
//
//  or:
//
//          nil
type DefinitionPlanPtrInput interface {
	pulumi.Input

	ToDefinitionPlanPtrOutput() DefinitionPlanPtrOutput
	ToDefinitionPlanPtrOutputWithContext(context.Context) DefinitionPlanPtrOutput
}

type definitionPlanPtrType DefinitionPlanArgs

func DefinitionPlanPtr(v *DefinitionPlanArgs) DefinitionPlanPtrInput {
	return (*definitionPlanPtrType)(v)
}

func (*definitionPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefinitionPlan)(nil)).Elem()
}

func (i *definitionPlanPtrType) ToDefinitionPlanPtrOutput() DefinitionPlanPtrOutput {
	return i.ToDefinitionPlanPtrOutputWithContext(context.Background())
}

func (i *definitionPlanPtrType) ToDefinitionPlanPtrOutputWithContext(ctx context.Context) DefinitionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionPlanPtrOutput)
}

type DefinitionPlanOutput struct{ *pulumi.OutputState }

func (DefinitionPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionPlan)(nil)).Elem()
}

func (o DefinitionPlanOutput) ToDefinitionPlanOutput() DefinitionPlanOutput {
	return o
}

func (o DefinitionPlanOutput) ToDefinitionPlanOutputWithContext(ctx context.Context) DefinitionPlanOutput {
	return o
}

func (o DefinitionPlanOutput) ToDefinitionPlanPtrOutput() DefinitionPlanPtrOutput {
	return o.ToDefinitionPlanPtrOutputWithContext(context.Background())
}

func (o DefinitionPlanOutput) ToDefinitionPlanPtrOutputWithContext(ctx context.Context) DefinitionPlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefinitionPlan) *DefinitionPlan {
		return &v
	}).(DefinitionPlanPtrOutput)
}

// The plan name of the marketplace offer.
func (o DefinitionPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionPlan) string { return v.Name }).(pulumi.StringOutput)
}

// The product code of the plan.
func (o DefinitionPlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionPlan) string { return v.Product }).(pulumi.StringOutput)
}

// The publisher ID of the plan.
func (o DefinitionPlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionPlan) string { return v.Publisher }).(pulumi.StringOutput)
}

// The version of the plan.
func (o DefinitionPlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionPlan) string { return v.Version }).(pulumi.StringOutput)
}

type DefinitionPlanPtrOutput struct{ *pulumi.OutputState }

func (DefinitionPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefinitionPlan)(nil)).Elem()
}

func (o DefinitionPlanPtrOutput) ToDefinitionPlanPtrOutput() DefinitionPlanPtrOutput {
	return o
}

func (o DefinitionPlanPtrOutput) ToDefinitionPlanPtrOutputWithContext(ctx context.Context) DefinitionPlanPtrOutput {
	return o
}

func (o DefinitionPlanPtrOutput) Elem() DefinitionPlanOutput {
	return o.ApplyT(func(v *DefinitionPlan) DefinitionPlan {
		if v != nil {
			return *v
		}
		var ret DefinitionPlan
		return ret
	}).(DefinitionPlanOutput)
}

// The plan name of the marketplace offer.
func (o DefinitionPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefinitionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The product code of the plan.
func (o DefinitionPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefinitionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

// The publisher ID of the plan.
func (o DefinitionPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefinitionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

// The version of the plan.
func (o DefinitionPlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefinitionPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionAuthorizationInput)(nil)).Elem(), DefinitionAuthorizationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionAuthorizationArrayInput)(nil)).Elem(), DefinitionAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionPlanInput)(nil)).Elem(), DefinitionPlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefinitionPlanPtrInput)(nil)).Elem(), DefinitionPlanArgs{})
	pulumi.RegisterOutputType(DefinitionAuthorizationOutput{})
	pulumi.RegisterOutputType(DefinitionAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(DefinitionPlanOutput{})
	pulumi.RegisterOutputType(DefinitionPlanPtrOutput{})
}
