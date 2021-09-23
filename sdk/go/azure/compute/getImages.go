// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Images within a Resource Group.
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
// 		_, err := compute.GetImages(ctx, &compute.GetImagesArgs{
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("azure:compute/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	// The name of the Resource Group in which the Image exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to filter the list of images against.
	TagsFilter map[string]string `pulumi:"tagsFilter"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// One or more `images` blocks as defined below:
	Images            []GetImagesImage  `pulumi:"images"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	TagsFilter        map[string]string `pulumi:"tagsFilter"`
}

func GetImagesOutput(ctx *pulumi.Context, args GetImagesOutputArgs, opts ...pulumi.InvokeOption) GetImagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImagesResult, error) {
			args := v.(GetImagesArgs)
			r, err := GetImages(ctx, &args, opts...)
			return *r, err
		}).(GetImagesResultOutput)
}

// A collection of arguments for invoking getImages.
type GetImagesOutputArgs struct {
	// The name of the Resource Group in which the Image exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// A mapping of tags to filter the list of images against.
	TagsFilter pulumi.StringMapInput `pulumi:"tagsFilter"`
}

func (GetImagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImagesArgs)(nil)).Elem()
}

// A collection of values returned by getImages.
type GetImagesResultOutput struct{ *pulumi.OutputState }

func (GetImagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImagesResult)(nil)).Elem()
}

func (o GetImagesResultOutput) ToGetImagesResultOutput() GetImagesResultOutput {
	return o
}

func (o GetImagesResultOutput) ToGetImagesResultOutputWithContext(ctx context.Context) GetImagesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetImagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImagesResult) string { return v.Id }).(pulumi.StringOutput)
}

// One or more `images` blocks as defined below:
func (o GetImagesResultOutput) Images() GetImagesImageArrayOutput {
	return o.ApplyT(func(v GetImagesResult) []GetImagesImage { return v.Images }).(GetImagesImageArrayOutput)
}

func (o GetImagesResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetImagesResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o GetImagesResultOutput) TagsFilter() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetImagesResult) map[string]string { return v.TagsFilter }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImagesResultOutput{})
}
