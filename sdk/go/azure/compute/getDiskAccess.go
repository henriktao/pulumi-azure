// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Disk Access.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := compute.LookupDiskAccess(ctx, &compute.LookupDiskAccessArgs{
// 			Name:              "existing",
// 			ResourceGroupName: "existing",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupDiskAccess(ctx *pulumi.Context, args *LookupDiskAccessArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessResult, error) {
	var rv LookupDiskAccessResult
	err := ctx.Invoke("azure:compute/getDiskAccess:getDiskAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDiskAccess.
type LookupDiskAccessArgs struct {
	// The name of this Disk Access.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Disk Access exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Disk Access.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDiskAccess.
type LookupDiskAccessResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string            `pulumi:"id"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

func LookupDiskAccessOutput(ctx *pulumi.Context, args LookupDiskAccessOutputArgs, opts ...pulumi.InvokeOption) LookupDiskAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskAccessResult, error) {
			args := v.(LookupDiskAccessArgs)
			r, err := LookupDiskAccess(ctx, &args, opts...)
			return *r, err
		}).(LookupDiskAccessResultOutput)
}

// A collection of arguments for invoking getDiskAccess.
type LookupDiskAccessOutputArgs struct {
	// The name of this Disk Access.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Disk Access exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Disk Access.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupDiskAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessArgs)(nil)).Elem()
}

// A collection of values returned by getDiskAccess.
type LookupDiskAccessResultOutput struct{ *pulumi.OutputState }

func (LookupDiskAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessResult)(nil)).Elem()
}

func (o LookupDiskAccessResultOutput) ToLookupDiskAccessResultOutput() LookupDiskAccessResultOutput {
	return o
}

func (o LookupDiskAccessResultOutput) ToLookupDiskAccessResultOutputWithContext(ctx context.Context) LookupDiskAccessResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDiskAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskAccessResultOutput{})
}
