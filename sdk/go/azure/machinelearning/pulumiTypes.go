// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InferenceClusterSsl struct {
	// The certificate for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Cert *string `pulumi:"cert"`
	// The cname of the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Cname *string `pulumi:"cname"`
	// The key content for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Key *string `pulumi:"key"`
	// The leaf domain label for the ssl configuration. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname`. Changing this forces a new Machine Learning Inference Cluster to be created.
	LeafDomainLabel *string `pulumi:"leafDomainLabel"`
	// Whether or not to overwrite existing leaf domain. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname` Changing this forces a new Machine Learning Inference Cluster to be created.
	OverwriteExistingDomain *bool `pulumi:"overwriteExistingDomain"`
}

// InferenceClusterSslInput is an input type that accepts InferenceClusterSslArgs and InferenceClusterSslOutput values.
// You can construct a concrete instance of `InferenceClusterSslInput` via:
//
//          InferenceClusterSslArgs{...}
type InferenceClusterSslInput interface {
	pulumi.Input

	ToInferenceClusterSslOutput() InferenceClusterSslOutput
	ToInferenceClusterSslOutputWithContext(context.Context) InferenceClusterSslOutput
}

type InferenceClusterSslArgs struct {
	// The certificate for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Cert pulumi.StringPtrInput `pulumi:"cert"`
	// The cname of the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Cname pulumi.StringPtrInput `pulumi:"cname"`
	// The key content for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// The leaf domain label for the ssl configuration. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname`. Changing this forces a new Machine Learning Inference Cluster to be created.
	LeafDomainLabel pulumi.StringPtrInput `pulumi:"leafDomainLabel"`
	// Whether or not to overwrite existing leaf domain. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname` Changing this forces a new Machine Learning Inference Cluster to be created.
	OverwriteExistingDomain pulumi.BoolPtrInput `pulumi:"overwriteExistingDomain"`
}

func (InferenceClusterSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceClusterSsl)(nil)).Elem()
}

func (i InferenceClusterSslArgs) ToInferenceClusterSslOutput() InferenceClusterSslOutput {
	return i.ToInferenceClusterSslOutputWithContext(context.Background())
}

func (i InferenceClusterSslArgs) ToInferenceClusterSslOutputWithContext(ctx context.Context) InferenceClusterSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterSslOutput)
}

func (i InferenceClusterSslArgs) ToInferenceClusterSslPtrOutput() InferenceClusterSslPtrOutput {
	return i.ToInferenceClusterSslPtrOutputWithContext(context.Background())
}

func (i InferenceClusterSslArgs) ToInferenceClusterSslPtrOutputWithContext(ctx context.Context) InferenceClusterSslPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterSslOutput).ToInferenceClusterSslPtrOutputWithContext(ctx)
}

// InferenceClusterSslPtrInput is an input type that accepts InferenceClusterSslArgs, InferenceClusterSslPtr and InferenceClusterSslPtrOutput values.
// You can construct a concrete instance of `InferenceClusterSslPtrInput` via:
//
//          InferenceClusterSslArgs{...}
//
//  or:
//
//          nil
type InferenceClusterSslPtrInput interface {
	pulumi.Input

	ToInferenceClusterSslPtrOutput() InferenceClusterSslPtrOutput
	ToInferenceClusterSslPtrOutputWithContext(context.Context) InferenceClusterSslPtrOutput
}

type inferenceClusterSslPtrType InferenceClusterSslArgs

func InferenceClusterSslPtr(v *InferenceClusterSslArgs) InferenceClusterSslPtrInput {
	return (*inferenceClusterSslPtrType)(v)
}

func (*inferenceClusterSslPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceClusterSsl)(nil)).Elem()
}

func (i *inferenceClusterSslPtrType) ToInferenceClusterSslPtrOutput() InferenceClusterSslPtrOutput {
	return i.ToInferenceClusterSslPtrOutputWithContext(context.Background())
}

func (i *inferenceClusterSslPtrType) ToInferenceClusterSslPtrOutputWithContext(ctx context.Context) InferenceClusterSslPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceClusterSslPtrOutput)
}

type InferenceClusterSslOutput struct{ *pulumi.OutputState }

func (InferenceClusterSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceClusterSsl)(nil)).Elem()
}

func (o InferenceClusterSslOutput) ToInferenceClusterSslOutput() InferenceClusterSslOutput {
	return o
}

func (o InferenceClusterSslOutput) ToInferenceClusterSslOutputWithContext(ctx context.Context) InferenceClusterSslOutput {
	return o
}

func (o InferenceClusterSslOutput) ToInferenceClusterSslPtrOutput() InferenceClusterSslPtrOutput {
	return o.ToInferenceClusterSslPtrOutputWithContext(context.Background())
}

func (o InferenceClusterSslOutput) ToInferenceClusterSslPtrOutputWithContext(ctx context.Context) InferenceClusterSslPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *InferenceClusterSsl {
		return &v
	}).(InferenceClusterSslPtrOutput)
}

// The certificate for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

// The cname of the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

// The key content for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The leaf domain label for the ssl configuration. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *string { return v.LeafDomainLabel }).(pulumi.StringPtrOutput)
}

// Whether or not to overwrite existing leaf domain. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname` Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InferenceClusterSsl) *bool { return v.OverwriteExistingDomain }).(pulumi.BoolPtrOutput)
}

type InferenceClusterSslPtrOutput struct{ *pulumi.OutputState }

func (InferenceClusterSslPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceClusterSsl)(nil)).Elem()
}

func (o InferenceClusterSslPtrOutput) ToInferenceClusterSslPtrOutput() InferenceClusterSslPtrOutput {
	return o
}

func (o InferenceClusterSslPtrOutput) ToInferenceClusterSslPtrOutputWithContext(ctx context.Context) InferenceClusterSslPtrOutput {
	return o
}

func (o InferenceClusterSslPtrOutput) Elem() InferenceClusterSslOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) InferenceClusterSsl { return *v }).(InferenceClusterSslOutput)
}

// The certificate for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslPtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) *string {
		if v == nil {
			return nil
		}
		return v.Cert
	}).(pulumi.StringPtrOutput)
}

// The cname of the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

// The key content for the ssl configuration.Conflicts with `ssl.0.leaf_domain_label`,`ssl.0.overwrite_existing_domain`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

// The leaf domain label for the ssl configuration. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname`. Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslPtrOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) *string {
		if v == nil {
			return nil
		}
		return v.LeafDomainLabel
	}).(pulumi.StringPtrOutput)
}

// Whether or not to overwrite existing leaf domain. Conflicts with `ssl.0.cert`,`ssl.0.key`,`ssl.0.cname` Changing this forces a new Machine Learning Inference Cluster to be created.
func (o InferenceClusterSslPtrOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InferenceClusterSsl) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingDomain
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceIdentity struct {
	// The (Client) ID of the Service Principal.
	PrincipalId *string `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId *string `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
	Type string `pulumi:"type"`
}

// WorkspaceIdentityInput is an input type that accepts WorkspaceIdentityArgs and WorkspaceIdentityOutput values.
// You can construct a concrete instance of `WorkspaceIdentityInput` via:
//
//          WorkspaceIdentityArgs{...}
type WorkspaceIdentityInput interface {
	pulumi.Input

	ToWorkspaceIdentityOutput() WorkspaceIdentityOutput
	ToWorkspaceIdentityOutputWithContext(context.Context) WorkspaceIdentityOutput
}

type WorkspaceIdentityArgs struct {
	// The (Client) ID of the Service Principal.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The ID of the Tenant the Service Principal is assigned in.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (WorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceIdentity)(nil)).Elem()
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityOutput() WorkspaceIdentityOutput {
	return i.ToWorkspaceIdentityOutputWithContext(context.Background())
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityOutputWithContext(ctx context.Context) WorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityOutput)
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return i.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i WorkspaceIdentityArgs) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityOutput).ToWorkspaceIdentityPtrOutputWithContext(ctx)
}

// WorkspaceIdentityPtrInput is an input type that accepts WorkspaceIdentityArgs, WorkspaceIdentityPtr and WorkspaceIdentityPtrOutput values.
// You can construct a concrete instance of `WorkspaceIdentityPtrInput` via:
//
//          WorkspaceIdentityArgs{...}
//
//  or:
//
//          nil
type WorkspaceIdentityPtrInput interface {
	pulumi.Input

	ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput
	ToWorkspaceIdentityPtrOutputWithContext(context.Context) WorkspaceIdentityPtrOutput
}

type workspaceIdentityPtrType WorkspaceIdentityArgs

func WorkspaceIdentityPtr(v *WorkspaceIdentityArgs) WorkspaceIdentityPtrInput {
	return (*workspaceIdentityPtrType)(v)
}

func (*workspaceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceIdentity)(nil)).Elem()
}

func (i *workspaceIdentityPtrType) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return i.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i *workspaceIdentityPtrType) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceIdentityPtrOutput)
}

type WorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (WorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceIdentity)(nil)).Elem()
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityOutput() WorkspaceIdentityOutput {
	return o
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityOutputWithContext(ctx context.Context) WorkspaceIdentityOutput {
	return o
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return o.ToWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (o WorkspaceIdentityOutput) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *WorkspaceIdentity {
		return &v
	}).(WorkspaceIdentityPtrOutput)
}

// The (Client) ID of the Service Principal.
func (o WorkspaceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o WorkspaceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
func (o WorkspaceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type WorkspaceIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceIdentity)(nil)).Elem()
}

func (o WorkspaceIdentityPtrOutput) ToWorkspaceIdentityPtrOutput() WorkspaceIdentityPtrOutput {
	return o
}

func (o WorkspaceIdentityPtrOutput) ToWorkspaceIdentityPtrOutputWithContext(ctx context.Context) WorkspaceIdentityPtrOutput {
	return o
}

func (o WorkspaceIdentityPtrOutput) Elem() WorkspaceIdentityOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) WorkspaceIdentity { return *v }).(WorkspaceIdentityOutput)
}

// The (Client) ID of the Service Principal.
func (o WorkspaceIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The ID of the Tenant the Service Principal is assigned in.
func (o WorkspaceIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
func (o WorkspaceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type GetWorkspaceIdentity struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

// GetWorkspaceIdentityInput is an input type that accepts GetWorkspaceIdentityArgs and GetWorkspaceIdentityOutput values.
// You can construct a concrete instance of `GetWorkspaceIdentityInput` via:
//
//          GetWorkspaceIdentityArgs{...}
type GetWorkspaceIdentityInput interface {
	pulumi.Input

	ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput
	ToGetWorkspaceIdentityOutputWithContext(context.Context) GetWorkspaceIdentityOutput
}

type GetWorkspaceIdentityArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (GetWorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceIdentity)(nil)).Elem()
}

func (i GetWorkspaceIdentityArgs) ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput {
	return i.ToGetWorkspaceIdentityOutputWithContext(context.Background())
}

func (i GetWorkspaceIdentityArgs) ToGetWorkspaceIdentityOutputWithContext(ctx context.Context) GetWorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspaceIdentityOutput)
}

// GetWorkspaceIdentityArrayInput is an input type that accepts GetWorkspaceIdentityArray and GetWorkspaceIdentityArrayOutput values.
// You can construct a concrete instance of `GetWorkspaceIdentityArrayInput` via:
//
//          GetWorkspaceIdentityArray{ GetWorkspaceIdentityArgs{...} }
type GetWorkspaceIdentityArrayInput interface {
	pulumi.Input

	ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput
	ToGetWorkspaceIdentityArrayOutputWithContext(context.Context) GetWorkspaceIdentityArrayOutput
}

type GetWorkspaceIdentityArray []GetWorkspaceIdentityInput

func (GetWorkspaceIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspaceIdentity)(nil)).Elem()
}

func (i GetWorkspaceIdentityArray) ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput {
	return i.ToGetWorkspaceIdentityArrayOutputWithContext(context.Background())
}

func (i GetWorkspaceIdentityArray) ToGetWorkspaceIdentityArrayOutputWithContext(ctx context.Context) GetWorkspaceIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkspaceIdentityArrayOutput)
}

type GetWorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (GetWorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceIdentity)(nil)).Elem()
}

func (o GetWorkspaceIdentityOutput) ToGetWorkspaceIdentityOutput() GetWorkspaceIdentityOutput {
	return o
}

func (o GetWorkspaceIdentityOutput) ToGetWorkspaceIdentityOutputWithContext(ctx context.Context) GetWorkspaceIdentityOutput {
	return o
}

func (o GetWorkspaceIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o GetWorkspaceIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o GetWorkspaceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkspaceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type GetWorkspaceIdentityArrayOutput struct{ *pulumi.OutputState }

func (GetWorkspaceIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkspaceIdentity)(nil)).Elem()
}

func (o GetWorkspaceIdentityArrayOutput) ToGetWorkspaceIdentityArrayOutput() GetWorkspaceIdentityArrayOutput {
	return o
}

func (o GetWorkspaceIdentityArrayOutput) ToGetWorkspaceIdentityArrayOutputWithContext(ctx context.Context) GetWorkspaceIdentityArrayOutput {
	return o
}

func (o GetWorkspaceIdentityArrayOutput) Index(i pulumi.IntInput) GetWorkspaceIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkspaceIdentity {
		return vs[0].([]GetWorkspaceIdentity)[vs[1].(int)]
	}).(GetWorkspaceIdentityOutput)
}

func init() {
	pulumi.RegisterOutputType(InferenceClusterSslOutput{})
	pulumi.RegisterOutputType(InferenceClusterSslPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(WorkspaceIdentityPtrOutput{})
	pulumi.RegisterOutputType(GetWorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(GetWorkspaceIdentityArrayOutput{})
}
