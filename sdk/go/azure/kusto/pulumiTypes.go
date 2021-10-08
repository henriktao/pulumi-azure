// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterIdentity struct {
	// A list of IDs for User Assigned Managed Identity resources to be assigned.
	IdentityIds []string `pulumi:"identityIds"`
	// The Principal ID associated with this System Assigned Managed Service Identity.
	PrincipalId *string `pulumi:"principalId"`
	// The Tenant ID associated with this System Assigned Managed Service Identity.
	TenantId *string `pulumi:"tenantId"`
	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
	Type string `pulumi:"type"`
}

// ClusterIdentityInput is an input type that accepts ClusterIdentityArgs and ClusterIdentityOutput values.
// You can construct a concrete instance of `ClusterIdentityInput` via:
//
//          ClusterIdentityArgs{...}
type ClusterIdentityInput interface {
	pulumi.Input

	ToClusterIdentityOutput() ClusterIdentityOutput
	ToClusterIdentityOutputWithContext(context.Context) ClusterIdentityOutput
}

type ClusterIdentityArgs struct {
	// A list of IDs for User Assigned Managed Identity resources to be assigned.
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	// The Principal ID associated with this System Assigned Managed Service Identity.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The Tenant ID associated with this System Assigned Managed Service Identity.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIdentity)(nil)).Elem()
}

func (i ClusterIdentityArgs) ToClusterIdentityOutput() ClusterIdentityOutput {
	return i.ToClusterIdentityOutputWithContext(context.Background())
}

func (i ClusterIdentityArgs) ToClusterIdentityOutputWithContext(ctx context.Context) ClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityOutput)
}

func (i ClusterIdentityArgs) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return i.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ClusterIdentityArgs) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityOutput).ToClusterIdentityPtrOutputWithContext(ctx)
}

// ClusterIdentityPtrInput is an input type that accepts ClusterIdentityArgs, ClusterIdentityPtr and ClusterIdentityPtrOutput values.
// You can construct a concrete instance of `ClusterIdentityPtrInput` via:
//
//          ClusterIdentityArgs{...}
//
//  or:
//
//          nil
type ClusterIdentityPtrInput interface {
	pulumi.Input

	ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput
	ToClusterIdentityPtrOutputWithContext(context.Context) ClusterIdentityPtrOutput
}

type clusterIdentityPtrType ClusterIdentityArgs

func ClusterIdentityPtr(v *ClusterIdentityArgs) ClusterIdentityPtrInput {
	return (*clusterIdentityPtrType)(v)
}

func (*clusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIdentity)(nil)).Elem()
}

func (i *clusterIdentityPtrType) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return i.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *clusterIdentityPtrType) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIdentityPtrOutput)
}

type ClusterIdentityOutput struct{ *pulumi.OutputState }

func (ClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIdentity)(nil)).Elem()
}

func (o ClusterIdentityOutput) ToClusterIdentityOutput() ClusterIdentityOutput {
	return o
}

func (o ClusterIdentityOutput) ToClusterIdentityOutputWithContext(ctx context.Context) ClusterIdentityOutput {
	return o
}

func (o ClusterIdentityOutput) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return o.ToClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ClusterIdentityOutput) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterIdentity) *ClusterIdentity {
		return &v
	}).(ClusterIdentityPtrOutput)
}

// A list of IDs for User Assigned Managed Identity resources to be assigned.
func (o ClusterIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ClusterIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

// The Principal ID associated with this System Assigned Managed Service Identity.
func (o ClusterIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The Tenant ID associated with this System Assigned Managed Service Identity.
func (o ClusterIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
func (o ClusterIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type ClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterIdentity)(nil)).Elem()
}

func (o ClusterIdentityPtrOutput) ToClusterIdentityPtrOutput() ClusterIdentityPtrOutput {
	return o
}

func (o ClusterIdentityPtrOutput) ToClusterIdentityPtrOutputWithContext(ctx context.Context) ClusterIdentityPtrOutput {
	return o
}

func (o ClusterIdentityPtrOutput) Elem() ClusterIdentityOutput {
	return o.ApplyT(func(v *ClusterIdentity) ClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ClusterIdentity
		return ret
	}).(ClusterIdentityOutput)
}

// A list of IDs for User Assigned Managed Identity resources to be assigned.
func (o ClusterIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

// The Principal ID associated with this System Assigned Managed Service Identity.
func (o ClusterIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The Tenant ID associated with this System Assigned Managed Service Identity.
func (o ClusterIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
func (o ClusterIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ClusterOptimizedAutoScale struct {
	// The maximum number of allowed instances. Must between `0` and `1000`.
	MaximumInstances int `pulumi:"maximumInstances"`
	// The minimum number of allowed instances. Must between `0` and `1000`.
	MinimumInstances int `pulumi:"minimumInstances"`
}

// ClusterOptimizedAutoScaleInput is an input type that accepts ClusterOptimizedAutoScaleArgs and ClusterOptimizedAutoScaleOutput values.
// You can construct a concrete instance of `ClusterOptimizedAutoScaleInput` via:
//
//          ClusterOptimizedAutoScaleArgs{...}
type ClusterOptimizedAutoScaleInput interface {
	pulumi.Input

	ToClusterOptimizedAutoScaleOutput() ClusterOptimizedAutoScaleOutput
	ToClusterOptimizedAutoScaleOutputWithContext(context.Context) ClusterOptimizedAutoScaleOutput
}

type ClusterOptimizedAutoScaleArgs struct {
	// The maximum number of allowed instances. Must between `0` and `1000`.
	MaximumInstances pulumi.IntInput `pulumi:"maximumInstances"`
	// The minimum number of allowed instances. Must between `0` and `1000`.
	MinimumInstances pulumi.IntInput `pulumi:"minimumInstances"`
}

func (ClusterOptimizedAutoScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterOptimizedAutoScale)(nil)).Elem()
}

func (i ClusterOptimizedAutoScaleArgs) ToClusterOptimizedAutoScaleOutput() ClusterOptimizedAutoScaleOutput {
	return i.ToClusterOptimizedAutoScaleOutputWithContext(context.Background())
}

func (i ClusterOptimizedAutoScaleArgs) ToClusterOptimizedAutoScaleOutputWithContext(ctx context.Context) ClusterOptimizedAutoScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOptimizedAutoScaleOutput)
}

func (i ClusterOptimizedAutoScaleArgs) ToClusterOptimizedAutoScalePtrOutput() ClusterOptimizedAutoScalePtrOutput {
	return i.ToClusterOptimizedAutoScalePtrOutputWithContext(context.Background())
}

func (i ClusterOptimizedAutoScaleArgs) ToClusterOptimizedAutoScalePtrOutputWithContext(ctx context.Context) ClusterOptimizedAutoScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOptimizedAutoScaleOutput).ToClusterOptimizedAutoScalePtrOutputWithContext(ctx)
}

// ClusterOptimizedAutoScalePtrInput is an input type that accepts ClusterOptimizedAutoScaleArgs, ClusterOptimizedAutoScalePtr and ClusterOptimizedAutoScalePtrOutput values.
// You can construct a concrete instance of `ClusterOptimizedAutoScalePtrInput` via:
//
//          ClusterOptimizedAutoScaleArgs{...}
//
//  or:
//
//          nil
type ClusterOptimizedAutoScalePtrInput interface {
	pulumi.Input

	ToClusterOptimizedAutoScalePtrOutput() ClusterOptimizedAutoScalePtrOutput
	ToClusterOptimizedAutoScalePtrOutputWithContext(context.Context) ClusterOptimizedAutoScalePtrOutput
}

type clusterOptimizedAutoScalePtrType ClusterOptimizedAutoScaleArgs

func ClusterOptimizedAutoScalePtr(v *ClusterOptimizedAutoScaleArgs) ClusterOptimizedAutoScalePtrInput {
	return (*clusterOptimizedAutoScalePtrType)(v)
}

func (*clusterOptimizedAutoScalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterOptimizedAutoScale)(nil)).Elem()
}

func (i *clusterOptimizedAutoScalePtrType) ToClusterOptimizedAutoScalePtrOutput() ClusterOptimizedAutoScalePtrOutput {
	return i.ToClusterOptimizedAutoScalePtrOutputWithContext(context.Background())
}

func (i *clusterOptimizedAutoScalePtrType) ToClusterOptimizedAutoScalePtrOutputWithContext(ctx context.Context) ClusterOptimizedAutoScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOptimizedAutoScalePtrOutput)
}

type ClusterOptimizedAutoScaleOutput struct{ *pulumi.OutputState }

func (ClusterOptimizedAutoScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterOptimizedAutoScale)(nil)).Elem()
}

func (o ClusterOptimizedAutoScaleOutput) ToClusterOptimizedAutoScaleOutput() ClusterOptimizedAutoScaleOutput {
	return o
}

func (o ClusterOptimizedAutoScaleOutput) ToClusterOptimizedAutoScaleOutputWithContext(ctx context.Context) ClusterOptimizedAutoScaleOutput {
	return o
}

func (o ClusterOptimizedAutoScaleOutput) ToClusterOptimizedAutoScalePtrOutput() ClusterOptimizedAutoScalePtrOutput {
	return o.ToClusterOptimizedAutoScalePtrOutputWithContext(context.Background())
}

func (o ClusterOptimizedAutoScaleOutput) ToClusterOptimizedAutoScalePtrOutputWithContext(ctx context.Context) ClusterOptimizedAutoScalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterOptimizedAutoScale) *ClusterOptimizedAutoScale {
		return &v
	}).(ClusterOptimizedAutoScalePtrOutput)
}

// The maximum number of allowed instances. Must between `0` and `1000`.
func (o ClusterOptimizedAutoScaleOutput) MaximumInstances() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterOptimizedAutoScale) int { return v.MaximumInstances }).(pulumi.IntOutput)
}

// The minimum number of allowed instances. Must between `0` and `1000`.
func (o ClusterOptimizedAutoScaleOutput) MinimumInstances() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterOptimizedAutoScale) int { return v.MinimumInstances }).(pulumi.IntOutput)
}

type ClusterOptimizedAutoScalePtrOutput struct{ *pulumi.OutputState }

func (ClusterOptimizedAutoScalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterOptimizedAutoScale)(nil)).Elem()
}

func (o ClusterOptimizedAutoScalePtrOutput) ToClusterOptimizedAutoScalePtrOutput() ClusterOptimizedAutoScalePtrOutput {
	return o
}

func (o ClusterOptimizedAutoScalePtrOutput) ToClusterOptimizedAutoScalePtrOutputWithContext(ctx context.Context) ClusterOptimizedAutoScalePtrOutput {
	return o
}

func (o ClusterOptimizedAutoScalePtrOutput) Elem() ClusterOptimizedAutoScaleOutput {
	return o.ApplyT(func(v *ClusterOptimizedAutoScale) ClusterOptimizedAutoScale {
		if v != nil {
			return *v
		}
		var ret ClusterOptimizedAutoScale
		return ret
	}).(ClusterOptimizedAutoScaleOutput)
}

// The maximum number of allowed instances. Must between `0` and `1000`.
func (o ClusterOptimizedAutoScalePtrOutput) MaximumInstances() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterOptimizedAutoScale) *int {
		if v == nil {
			return nil
		}
		return &v.MaximumInstances
	}).(pulumi.IntPtrOutput)
}

// The minimum number of allowed instances. Must between `0` and `1000`.
func (o ClusterOptimizedAutoScalePtrOutput) MinimumInstances() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterOptimizedAutoScale) *int {
		if v == nil {
			return nil
		}
		return &v.MinimumInstances
	}).(pulumi.IntPtrOutput)
}

type ClusterSku struct {
	// Specifies the node count for the cluster. Boundaries depend on the sku name.
	Capacity *int `pulumi:"capacity"`
	// The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s`, `Standard_L8s`, `Standard_L16s_v2` and `Standard_L8s_v2`.
	Name string `pulumi:"name"`
}

// ClusterSkuInput is an input type that accepts ClusterSkuArgs and ClusterSkuOutput values.
// You can construct a concrete instance of `ClusterSkuInput` via:
//
//          ClusterSkuArgs{...}
type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	// Specifies the node count for the cluster. Boundaries depend on the sku name.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s`, `Standard_L8s`, `Standard_L16s_v2` and `Standard_L8s_v2`.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}

// ClusterSkuPtrInput is an input type that accepts ClusterSkuArgs, ClusterSkuPtr and ClusterSkuPtrOutput values.
// You can construct a concrete instance of `ClusterSkuPtrInput` via:
//
//          ClusterSkuArgs{...}
//
//  or:
//
//          nil
type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

// Specifies the node count for the cluster. Boundaries depend on the sku name.
func (o ClusterSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s`, `Standard_L8s`, `Standard_L16s_v2` and `Standard_L8s_v2`.
func (o ClusterSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterSku) string { return v.Name }).(pulumi.StringOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

// Specifies the node count for the cluster. Boundaries depend on the sku name.
func (o ClusterSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s`, `Standard_L8s`, `Standard_L16s_v2` and `Standard_L8s_v2`.
func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterVirtualNetworkConfiguration struct {
	// Data management's service public IP address resource id.
	DataManagementPublicIpId string `pulumi:"dataManagementPublicIpId"`
	// Engine service's public IP address resource id.
	EnginePublicIpId string `pulumi:"enginePublicIpId"`
	// The subnet resource id.
	SubnetId string `pulumi:"subnetId"`
}

// ClusterVirtualNetworkConfigurationInput is an input type that accepts ClusterVirtualNetworkConfigurationArgs and ClusterVirtualNetworkConfigurationOutput values.
// You can construct a concrete instance of `ClusterVirtualNetworkConfigurationInput` via:
//
//          ClusterVirtualNetworkConfigurationArgs{...}
type ClusterVirtualNetworkConfigurationInput interface {
	pulumi.Input

	ToClusterVirtualNetworkConfigurationOutput() ClusterVirtualNetworkConfigurationOutput
	ToClusterVirtualNetworkConfigurationOutputWithContext(context.Context) ClusterVirtualNetworkConfigurationOutput
}

type ClusterVirtualNetworkConfigurationArgs struct {
	// Data management's service public IP address resource id.
	DataManagementPublicIpId pulumi.StringInput `pulumi:"dataManagementPublicIpId"`
	// Engine service's public IP address resource id.
	EnginePublicIpId pulumi.StringInput `pulumi:"enginePublicIpId"`
	// The subnet resource id.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (ClusterVirtualNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterVirtualNetworkConfiguration)(nil)).Elem()
}

func (i ClusterVirtualNetworkConfigurationArgs) ToClusterVirtualNetworkConfigurationOutput() ClusterVirtualNetworkConfigurationOutput {
	return i.ToClusterVirtualNetworkConfigurationOutputWithContext(context.Background())
}

func (i ClusterVirtualNetworkConfigurationArgs) ToClusterVirtualNetworkConfigurationOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterVirtualNetworkConfigurationOutput)
}

func (i ClusterVirtualNetworkConfigurationArgs) ToClusterVirtualNetworkConfigurationPtrOutput() ClusterVirtualNetworkConfigurationPtrOutput {
	return i.ToClusterVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i ClusterVirtualNetworkConfigurationArgs) ToClusterVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterVirtualNetworkConfigurationOutput).ToClusterVirtualNetworkConfigurationPtrOutputWithContext(ctx)
}

// ClusterVirtualNetworkConfigurationPtrInput is an input type that accepts ClusterVirtualNetworkConfigurationArgs, ClusterVirtualNetworkConfigurationPtr and ClusterVirtualNetworkConfigurationPtrOutput values.
// You can construct a concrete instance of `ClusterVirtualNetworkConfigurationPtrInput` via:
//
//          ClusterVirtualNetworkConfigurationArgs{...}
//
//  or:
//
//          nil
type ClusterVirtualNetworkConfigurationPtrInput interface {
	pulumi.Input

	ToClusterVirtualNetworkConfigurationPtrOutput() ClusterVirtualNetworkConfigurationPtrOutput
	ToClusterVirtualNetworkConfigurationPtrOutputWithContext(context.Context) ClusterVirtualNetworkConfigurationPtrOutput
}

type clusterVirtualNetworkConfigurationPtrType ClusterVirtualNetworkConfigurationArgs

func ClusterVirtualNetworkConfigurationPtr(v *ClusterVirtualNetworkConfigurationArgs) ClusterVirtualNetworkConfigurationPtrInput {
	return (*clusterVirtualNetworkConfigurationPtrType)(v)
}

func (*clusterVirtualNetworkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterVirtualNetworkConfiguration)(nil)).Elem()
}

func (i *clusterVirtualNetworkConfigurationPtrType) ToClusterVirtualNetworkConfigurationPtrOutput() ClusterVirtualNetworkConfigurationPtrOutput {
	return i.ToClusterVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *clusterVirtualNetworkConfigurationPtrType) ToClusterVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterVirtualNetworkConfigurationPtrOutput)
}

type ClusterVirtualNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (ClusterVirtualNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterVirtualNetworkConfiguration)(nil)).Elem()
}

func (o ClusterVirtualNetworkConfigurationOutput) ToClusterVirtualNetworkConfigurationOutput() ClusterVirtualNetworkConfigurationOutput {
	return o
}

func (o ClusterVirtualNetworkConfigurationOutput) ToClusterVirtualNetworkConfigurationOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationOutput {
	return o
}

func (o ClusterVirtualNetworkConfigurationOutput) ToClusterVirtualNetworkConfigurationPtrOutput() ClusterVirtualNetworkConfigurationPtrOutput {
	return o.ToClusterVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o ClusterVirtualNetworkConfigurationOutput) ToClusterVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterVirtualNetworkConfiguration) *ClusterVirtualNetworkConfiguration {
		return &v
	}).(ClusterVirtualNetworkConfigurationPtrOutput)
}

// Data management's service public IP address resource id.
func (o ClusterVirtualNetworkConfigurationOutput) DataManagementPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterVirtualNetworkConfiguration) string { return v.DataManagementPublicIpId }).(pulumi.StringOutput)
}

// Engine service's public IP address resource id.
func (o ClusterVirtualNetworkConfigurationOutput) EnginePublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterVirtualNetworkConfiguration) string { return v.EnginePublicIpId }).(pulumi.StringOutput)
}

// The subnet resource id.
func (o ClusterVirtualNetworkConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterVirtualNetworkConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type ClusterVirtualNetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ClusterVirtualNetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterVirtualNetworkConfiguration)(nil)).Elem()
}

func (o ClusterVirtualNetworkConfigurationPtrOutput) ToClusterVirtualNetworkConfigurationPtrOutput() ClusterVirtualNetworkConfigurationPtrOutput {
	return o
}

func (o ClusterVirtualNetworkConfigurationPtrOutput) ToClusterVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) ClusterVirtualNetworkConfigurationPtrOutput {
	return o
}

func (o ClusterVirtualNetworkConfigurationPtrOutput) Elem() ClusterVirtualNetworkConfigurationOutput {
	return o.ApplyT(func(v *ClusterVirtualNetworkConfiguration) ClusterVirtualNetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret ClusterVirtualNetworkConfiguration
		return ret
	}).(ClusterVirtualNetworkConfigurationOutput)
}

// Data management's service public IP address resource id.
func (o ClusterVirtualNetworkConfigurationPtrOutput) DataManagementPublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterVirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.DataManagementPublicIpId
	}).(pulumi.StringPtrOutput)
}

// Engine service's public IP address resource id.
func (o ClusterVirtualNetworkConfigurationPtrOutput) EnginePublicIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterVirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.EnginePublicIpId
	}).(pulumi.StringPtrOutput)
}

// The subnet resource id.
func (o ClusterVirtualNetworkConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterVirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIdentityInput)(nil)).Elem(), ClusterIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterIdentityPtrInput)(nil)).Elem(), ClusterIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterOptimizedAutoScaleInput)(nil)).Elem(), ClusterOptimizedAutoScaleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterOptimizedAutoScalePtrInput)(nil)).Elem(), ClusterOptimizedAutoScaleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuInput)(nil)).Elem(), ClusterSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuPtrInput)(nil)).Elem(), ClusterSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterVirtualNetworkConfigurationInput)(nil)).Elem(), ClusterVirtualNetworkConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterVirtualNetworkConfigurationPtrInput)(nil)).Elem(), ClusterVirtualNetworkConfigurationArgs{})
	pulumi.RegisterOutputType(ClusterIdentityOutput{})
	pulumi.RegisterOutputType(ClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ClusterOptimizedAutoScaleOutput{})
	pulumi.RegisterOutputType(ClusterOptimizedAutoScalePtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterVirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(ClusterVirtualNetworkConfigurationPtrOutput{})
}
