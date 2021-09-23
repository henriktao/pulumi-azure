// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing File Share.
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
// 		example, err := storage.LookupShare(ctx, &storage.LookupShareArgs{
// 			Name:               "existing",
// 			StorageAccountName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure:storage/getShare:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShare.
type LookupShareArgs struct {
	// One or more acl blocks as defined below.
	Acls []GetShareAcl `pulumi:"acls"`
	// A map of custom file share metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the share.
	Name string `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// A collection of values returned by getShare.
type LookupShareResult struct {
	// One or more acl blocks as defined below.
	Acls []GetShareAcl `pulumi:"acls"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A map of custom file share metadata.
	Metadata map[string]string `pulumi:"metadata"`
	Name     string            `pulumi:"name"`
	// The quota of the File Share in GB.
	Quota              int    `pulumi:"quota"`
	ResourceManagerId  string `pulumi:"resourceManagerId"`
	StorageAccountName string `pulumi:"storageAccountName"`
}

func LookupShareOutput(ctx *pulumi.Context, args LookupShareOutputArgs, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareResult, error) {
			args := v.(LookupShareArgs)
			r, err := LookupShare(ctx, &args, opts...)
			return *r, err
		}).(LookupShareResultOutput)
}

// A collection of arguments for invoking getShare.
type LookupShareOutputArgs struct {
	// One or more acl blocks as defined below.
	Acls GetShareAclArrayInput `pulumi:"acls"`
	// A map of custom file share metadata.
	Metadata pulumi.StringMapInput `pulumi:"metadata"`
	// The name of the share.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (LookupShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}

// A collection of values returned by getShare.
type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

// One or more acl blocks as defined below.
func (o LookupShareResultOutput) Acls() GetShareAclArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []GetShareAcl { return v.Acls }).(GetShareAclArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

// A map of custom file share metadata.
func (o LookupShareResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupShareResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// The quota of the File Share in GB.
func (o LookupShareResultOutput) Quota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupShareResult) int { return v.Quota }).(pulumi.IntOutput)
}

func (o LookupShareResultOutput) ResourceManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ResourceManagerId }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
