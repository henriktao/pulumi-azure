// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentIdentity struct {
	IdentityIds []string `pulumi:"identityIds"`
	// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
	Type string `pulumi:"type"`
}

// AssignmentIdentityInput is an input type that accepts AssignmentIdentityArgs and AssignmentIdentityOutput values.
// You can construct a concrete instance of `AssignmentIdentityInput` via:
//
//          AssignmentIdentityArgs{...}
type AssignmentIdentityInput interface {
	pulumi.Input

	ToAssignmentIdentityOutput() AssignmentIdentityOutput
	ToAssignmentIdentityOutputWithContext(context.Context) AssignmentIdentityOutput
}

type AssignmentIdentityArgs struct {
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
	Type pulumi.StringInput `pulumi:"type"`
}

func (AssignmentIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentIdentity)(nil)).Elem()
}

func (i AssignmentIdentityArgs) ToAssignmentIdentityOutput() AssignmentIdentityOutput {
	return i.ToAssignmentIdentityOutputWithContext(context.Background())
}

func (i AssignmentIdentityArgs) ToAssignmentIdentityOutputWithContext(ctx context.Context) AssignmentIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentIdentityOutput)
}

func (i AssignmentIdentityArgs) ToAssignmentIdentityPtrOutput() AssignmentIdentityPtrOutput {
	return i.ToAssignmentIdentityPtrOutputWithContext(context.Background())
}

func (i AssignmentIdentityArgs) ToAssignmentIdentityPtrOutputWithContext(ctx context.Context) AssignmentIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentIdentityOutput).ToAssignmentIdentityPtrOutputWithContext(ctx)
}

// AssignmentIdentityPtrInput is an input type that accepts AssignmentIdentityArgs, AssignmentIdentityPtr and AssignmentIdentityPtrOutput values.
// You can construct a concrete instance of `AssignmentIdentityPtrInput` via:
//
//          AssignmentIdentityArgs{...}
//
//  or:
//
//          nil
type AssignmentIdentityPtrInput interface {
	pulumi.Input

	ToAssignmentIdentityPtrOutput() AssignmentIdentityPtrOutput
	ToAssignmentIdentityPtrOutputWithContext(context.Context) AssignmentIdentityPtrOutput
}

type assignmentIdentityPtrType AssignmentIdentityArgs

func AssignmentIdentityPtr(v *AssignmentIdentityArgs) AssignmentIdentityPtrInput {
	return (*assignmentIdentityPtrType)(v)
}

func (*assignmentIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentIdentity)(nil)).Elem()
}

func (i *assignmentIdentityPtrType) ToAssignmentIdentityPtrOutput() AssignmentIdentityPtrOutput {
	return i.ToAssignmentIdentityPtrOutputWithContext(context.Background())
}

func (i *assignmentIdentityPtrType) ToAssignmentIdentityPtrOutputWithContext(ctx context.Context) AssignmentIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentIdentityPtrOutput)
}

type AssignmentIdentityOutput struct{ *pulumi.OutputState }

func (AssignmentIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentIdentity)(nil)).Elem()
}

func (o AssignmentIdentityOutput) ToAssignmentIdentityOutput() AssignmentIdentityOutput {
	return o
}

func (o AssignmentIdentityOutput) ToAssignmentIdentityOutputWithContext(ctx context.Context) AssignmentIdentityOutput {
	return o
}

func (o AssignmentIdentityOutput) ToAssignmentIdentityPtrOutput() AssignmentIdentityPtrOutput {
	return o.ToAssignmentIdentityPtrOutputWithContext(context.Background())
}

func (o AssignmentIdentityOutput) ToAssignmentIdentityPtrOutputWithContext(ctx context.Context) AssignmentIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentIdentity) *AssignmentIdentity {
		return &v
	}).(AssignmentIdentityPtrOutput)
}

func (o AssignmentIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
func (o AssignmentIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type AssignmentIdentityPtrOutput struct{ *pulumi.OutputState }

func (AssignmentIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentIdentity)(nil)).Elem()
}

func (o AssignmentIdentityPtrOutput) ToAssignmentIdentityPtrOutput() AssignmentIdentityPtrOutput {
	return o
}

func (o AssignmentIdentityPtrOutput) ToAssignmentIdentityPtrOutputWithContext(ctx context.Context) AssignmentIdentityPtrOutput {
	return o
}

func (o AssignmentIdentityPtrOutput) Elem() AssignmentIdentityOutput {
	return o.ApplyT(func(v *AssignmentIdentity) AssignmentIdentity {
		if v != nil {
			return *v
		}
		var ret AssignmentIdentity
		return ret
	}).(AssignmentIdentityOutput)
}

func (o AssignmentIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssignmentIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

// The Identity type for the Managed Service Identity. Currently only `UserAssigned` is supported.
func (o AssignmentIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentIdentityInput)(nil)).Elem(), AssignmentIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentIdentityPtrInput)(nil)).Elem(), AssignmentIdentityArgs{})
	pulumi.RegisterOutputType(AssignmentIdentityOutput{})
	pulumi.RegisterOutputType(AssignmentIdentityPtrOutput{})
}
