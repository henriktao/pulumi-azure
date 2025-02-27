// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proximity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Proximity Placement Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/proximity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := proximity.LookupPlacementGroup(ctx, &proximity.LookupPlacementGroupArgs{
// 			Name:              "tf-appsecuritygroup",
// 			ResourceGroupName: "my-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("proximityPlacementGroupId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupPlacementGroup(ctx *pulumi.Context, args *LookupPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPlacementGroupResult, error) {
	var rv LookupPlacementGroupResult
	err := ctx.Invoke("azure:proximity/getPlacementGroup:getPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlacementGroup.
type LookupPlacementGroupArgs struct {
	// The name of the Proximity Placement Group.
	Name string `pulumi:"name"`
	// The name of the resource group in which the Proximity Placement Group exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPlacementGroup.
type LookupPlacementGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

func LookupPlacementGroupOutput(ctx *pulumi.Context, args LookupPlacementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlacementGroupResult, error) {
			args := v.(LookupPlacementGroupArgs)
			r, err := LookupPlacementGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupPlacementGroupResultOutput)
}

// A collection of arguments for invoking getPlacementGroup.
type LookupPlacementGroupOutputArgs struct {
	// The name of the Proximity Placement Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the Proximity Placement Group exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPlacementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupArgs)(nil)).Elem()
}

// A collection of values returned by getPlacementGroup.
type LookupPlacementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementGroupResult)(nil)).Elem()
}

func (o LookupPlacementGroupResultOutput) ToLookupPlacementGroupResultOutput() LookupPlacementGroupResultOutput {
	return o
}

func (o LookupPlacementGroupResultOutput) ToLookupPlacementGroupResultOutputWithContext(ctx context.Context) LookupPlacementGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPlacementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPlacementGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPlacementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPlacementGroupResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupPlacementGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPlacementGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementGroupResultOutput{})
}
