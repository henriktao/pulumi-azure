// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Storage Management Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.LookupAccount(ctx, &storage.LookupAccountArgs{
// 			Name:              "storageaccountname",
// 			ResourceGroupName: pulumi.StringRef("resourcegroupname"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.GetPolicy(ctx, &storage.GetPolicyArgs{
// 			StorageAccountId: azurerm_storage_account.Example.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPolicy(ctx *pulumi.Context, args *GetPolicyArgs, opts ...pulumi.InvokeOption) (*GetPolicyResult, error) {
	var rv GetPolicyResult
	err := ctx.Invoke("azure:storage/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type GetPolicyArgs struct {
	// Specifies the id of the storage account to retrieve the management policy for.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// A collection of values returned by getPolicy.
type GetPolicyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A `rule` block as documented below.
	Rules            []GetPolicyRule `pulumi:"rules"`
	StorageAccountId string          `pulumi:"storageAccountId"`
}

func GetPolicyOutput(ctx *pulumi.Context, args GetPolicyOutputArgs, opts ...pulumi.InvokeOption) GetPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyResult, error) {
			args := v.(GetPolicyArgs)
			r, err := GetPolicy(ctx, &args, opts...)
			return *r, err
		}).(GetPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type GetPolicyOutputArgs struct {
	// Specifies the id of the storage account to retrieve the management policy for.
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (GetPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type GetPolicyResultOutput struct{ *pulumi.OutputState }

func (GetPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyResult)(nil)).Elem()
}

func (o GetPolicyResultOutput) ToGetPolicyResultOutput() GetPolicyResultOutput {
	return o
}

func (o GetPolicyResultOutput) ToGetPolicyResultOutputWithContext(ctx context.Context) GetPolicyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// A `rule` block as documented below.
func (o GetPolicyResultOutput) Rules() GetPolicyRuleArrayOutput {
	return o.ApplyT(func(v GetPolicyResult) []GetPolicyRule { return v.Rules }).(GetPolicyRuleArrayOutput)
}

func (o GetPolicyResultOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyResult) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyResultOutput{})
}
