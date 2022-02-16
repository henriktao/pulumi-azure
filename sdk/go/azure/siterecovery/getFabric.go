// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Site Recovery Replication Fabric.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := siterecovery.LookupFabric(ctx, &siterecovery.LookupFabricArgs{
// 			Name:              "primary-fabric",
// 			RecoveryVaultName: "tfex-recovery_vault",
// 			ResourceGroupName: "tfex-resource_group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFabric(ctx *pulumi.Context, args *LookupFabricArgs, opts ...pulumi.InvokeOption) (*LookupFabricResult, error) {
	var rv LookupFabricResult
	err := ctx.Invoke("azure:siterecovery/getFabric:getFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFabric.
type LookupFabricArgs struct {
	// Specifies the name of the Site Recovery Replication Fabric.
	Name string `pulumi:"name"`
	// The name of the Recovery Services Vault that the Site Recovery Replication Fabric is associated witth.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the associated Recovery Services Vault resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getFabric.
type LookupFabricResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure location where the Site Recovery Replication Fabric resides.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

func LookupFabricOutput(ctx *pulumi.Context, args LookupFabricOutputArgs, opts ...pulumi.InvokeOption) LookupFabricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFabricResult, error) {
			args := v.(LookupFabricArgs)
			r, err := LookupFabric(ctx, &args, opts...)
			return *r, err
		}).(LookupFabricResultOutput)
}

// A collection of arguments for invoking getFabric.
type LookupFabricOutputArgs struct {
	// Specifies the name of the Site Recovery Replication Fabric.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Recovery Services Vault that the Site Recovery Replication Fabric is associated witth.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the associated Recovery Services Vault resides.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFabricArgs)(nil)).Elem()
}

// A collection of values returned by getFabric.
type LookupFabricResultOutput struct{ *pulumi.OutputState }

func (LookupFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFabricResult)(nil)).Elem()
}

func (o LookupFabricResultOutput) ToLookupFabricResultOutput() LookupFabricResultOutput {
	return o
}

func (o LookupFabricResultOutput) ToLookupFabricResultOutputWithContext(ctx context.Context) LookupFabricResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure location where the Site Recovery Replication Fabric resides.
func (o LookupFabricResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFabricResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFabricResultOutput) RecoveryVaultName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFabricResult) string { return v.RecoveryVaultName }).(pulumi.StringOutput)
}

func (o LookupFabricResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFabricResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFabricResultOutput{})
}
