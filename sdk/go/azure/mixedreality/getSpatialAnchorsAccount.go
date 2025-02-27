// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about an Azure Spatial Anchors Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mixedreality"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mixedreality.LookupSpatialAnchorsAccount(ctx, &mixedreality.LookupSpatialAnchorsAccountArgs{
// 			Name:              "example",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("accountDomain", data.Azurerm_spatial_anchors_account.Account_domain)
// 		return nil
// 	})
// }
// ```
func LookupSpatialAnchorsAccount(ctx *pulumi.Context, args *LookupSpatialAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupSpatialAnchorsAccountResult, error) {
	var rv LookupSpatialAnchorsAccountResult
	err := ctx.Invoke("azure:mixedreality/getSpatialAnchorsAccount:getSpatialAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpatialAnchorsAccount.
type LookupSpatialAnchorsAccountArgs struct {
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name string `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSpatialAnchorsAccount.
type LookupSpatialAnchorsAccountResult struct {
	// The domain of the Spatial Anchors Account.
	AccountDomain string `pulumi:"accountDomain"`
	// The account ID of the Spatial Anchors Account.
	AccountId string `pulumi:"accountId"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

func LookupSpatialAnchorsAccountOutput(ctx *pulumi.Context, args LookupSpatialAnchorsAccountOutputArgs, opts ...pulumi.InvokeOption) LookupSpatialAnchorsAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpatialAnchorsAccountResult, error) {
			args := v.(LookupSpatialAnchorsAccountArgs)
			r, err := LookupSpatialAnchorsAccount(ctx, &args, opts...)
			return *r, err
		}).(LookupSpatialAnchorsAccountResultOutput)
}

// A collection of arguments for invoking getSpatialAnchorsAccount.
type LookupSpatialAnchorsAccountOutputArgs struct {
	// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which to create the Spatial Anchors Account.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSpatialAnchorsAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpatialAnchorsAccountArgs)(nil)).Elem()
}

// A collection of values returned by getSpatialAnchorsAccount.
type LookupSpatialAnchorsAccountResultOutput struct{ *pulumi.OutputState }

func (LookupSpatialAnchorsAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpatialAnchorsAccountResult)(nil)).Elem()
}

func (o LookupSpatialAnchorsAccountResultOutput) ToLookupSpatialAnchorsAccountResultOutput() LookupSpatialAnchorsAccountResultOutput {
	return o
}

func (o LookupSpatialAnchorsAccountResultOutput) ToLookupSpatialAnchorsAccountResultOutputWithContext(ctx context.Context) LookupSpatialAnchorsAccountResultOutput {
	return o
}

// The domain of the Spatial Anchors Account.
func (o LookupSpatialAnchorsAccountResultOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.AccountDomain }).(pulumi.StringOutput)
}

// The account ID of the Spatial Anchors Account.
func (o LookupSpatialAnchorsAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSpatialAnchorsAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSpatialAnchorsAccountResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpatialAnchorsAccountResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpatialAnchorsAccountResultOutput{})
}
