// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing site recovery services protection container.
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
// 		_, err := siterecovery.LookupProtectionContainer(ctx, &siterecovery.LookupProtectionContainerArgs{
// 			Name:               "primary-container",
// 			RecoveryFabricName: "primary-fabric",
// 			RecoveryVaultName:  "tfex-recovery_vault",
// 			ResourceGroupName:  "tfex-resource_group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupProtectionContainer(ctx *pulumi.Context, args *LookupProtectionContainerArgs, opts ...pulumi.InvokeOption) (*LookupProtectionContainerResult, error) {
	var rv LookupProtectionContainerResult
	err := ctx.Invoke("azure:siterecovery/getProtectionContainer:getProtectionContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProtectionContainer.
type LookupProtectionContainerArgs struct {
	// Specifies the name of the protection container.
	Name string `pulumi:"name"`
	// The name of the fabric that contains the protection container.
	RecoveryFabricName string `pulumi:"recoveryFabricName"`
	// The name of the Recovery Services Vault that the protection container is associated witth.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the associated protection container resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getProtectionContainer.
type LookupProtectionContainerResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	Name               string `pulumi:"name"`
	RecoveryFabricName string `pulumi:"recoveryFabricName"`
	RecoveryVaultName  string `pulumi:"recoveryVaultName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}

func LookupProtectionContainerOutput(ctx *pulumi.Context, args LookupProtectionContainerOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionContainerResult, error) {
			args := v.(LookupProtectionContainerArgs)
			r, err := LookupProtectionContainer(ctx, &args, opts...)
			return *r, err
		}).(LookupProtectionContainerResultOutput)
}

// A collection of arguments for invoking getProtectionContainer.
type LookupProtectionContainerOutputArgs struct {
	// Specifies the name of the protection container.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the fabric that contains the protection container.
	RecoveryFabricName pulumi.StringInput `pulumi:"recoveryFabricName"`
	// The name of the Recovery Services Vault that the protection container is associated witth.
	RecoveryVaultName pulumi.StringInput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which the associated protection container resides.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProtectionContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionContainerArgs)(nil)).Elem()
}

// A collection of values returned by getProtectionContainer.
type LookupProtectionContainerResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionContainerResult)(nil)).Elem()
}

func (o LookupProtectionContainerResultOutput) ToLookupProtectionContainerResultOutput() LookupProtectionContainerResultOutput {
	return o
}

func (o LookupProtectionContainerResultOutput) ToLookupProtectionContainerResultOutputWithContext(ctx context.Context) LookupProtectionContainerResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProtectionContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) RecoveryFabricName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.RecoveryFabricName }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) RecoveryVaultName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.RecoveryVaultName }).(pulumi.StringOutput)
}

func (o LookupProtectionContainerResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionContainerResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionContainerResultOutput{})
}
