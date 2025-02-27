// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing DataShareDataLakeGen1Dataset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datashare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := datashare.LookupDatasetDataLakeGen1(ctx, &datashare.LookupDatasetDataLakeGen1Args{
// 			Name:        "example-dsdsdlg1",
// 			DataShareId: "example-share-id",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupDatasetDataLakeGen1(ctx *pulumi.Context, args *LookupDatasetDataLakeGen1Args, opts ...pulumi.InvokeOption) (*LookupDatasetDataLakeGen1Result, error) {
	var rv LookupDatasetDataLakeGen1Result
	err := ctx.Invoke("azure:datashare/getDatasetDataLakeGen1:getDatasetDataLakeGen1", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatasetDataLakeGen1.
type LookupDatasetDataLakeGen1Args struct {
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created.
	DataShareId string `pulumi:"dataShareId"`
	// The name of the Data Share Data Lake Gen1 Dataset.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatasetDataLakeGen1.
type LookupDatasetDataLakeGen1Result struct {
	// The resource ID of the Data Lake Store to be shared with the receiver.
	DataLakeStoreId string `pulumi:"dataLakeStoreId"`
	DataShareId     string `pulumi:"dataShareId"`
	// The displayed name of the Data Share Dataset.
	DisplayName string `pulumi:"displayName"`
	// The file name of the data lake store to be shared with the receiver.
	FileName string `pulumi:"fileName"`
	// The folder path of the data lake store to be shared with the receiver.
	FolderPath string `pulumi:"folderPath"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupDatasetDataLakeGen1Output(ctx *pulumi.Context, args LookupDatasetDataLakeGen1OutputArgs, opts ...pulumi.InvokeOption) LookupDatasetDataLakeGen1ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatasetDataLakeGen1Result, error) {
			args := v.(LookupDatasetDataLakeGen1Args)
			r, err := LookupDatasetDataLakeGen1(ctx, &args, opts...)
			return *r, err
		}).(LookupDatasetDataLakeGen1ResultOutput)
}

// A collection of arguments for invoking getDatasetDataLakeGen1.
type LookupDatasetDataLakeGen1OutputArgs struct {
	// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created.
	DataShareId pulumi.StringInput `pulumi:"dataShareId"`
	// The name of the Data Share Data Lake Gen1 Dataset.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDatasetDataLakeGen1OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetDataLakeGen1Args)(nil)).Elem()
}

// A collection of values returned by getDatasetDataLakeGen1.
type LookupDatasetDataLakeGen1ResultOutput struct{ *pulumi.OutputState }

func (LookupDatasetDataLakeGen1ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetDataLakeGen1Result)(nil)).Elem()
}

func (o LookupDatasetDataLakeGen1ResultOutput) ToLookupDatasetDataLakeGen1ResultOutput() LookupDatasetDataLakeGen1ResultOutput {
	return o
}

func (o LookupDatasetDataLakeGen1ResultOutput) ToLookupDatasetDataLakeGen1ResultOutputWithContext(ctx context.Context) LookupDatasetDataLakeGen1ResultOutput {
	return o
}

// The resource ID of the Data Lake Store to be shared with the receiver.
func (o LookupDatasetDataLakeGen1ResultOutput) DataLakeStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.DataLakeStoreId }).(pulumi.StringOutput)
}

func (o LookupDatasetDataLakeGen1ResultOutput) DataShareId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.DataShareId }).(pulumi.StringOutput)
}

// The displayed name of the Data Share Dataset.
func (o LookupDatasetDataLakeGen1ResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The file name of the data lake store to be shared with the receiver.
func (o LookupDatasetDataLakeGen1ResultOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.FileName }).(pulumi.StringOutput)
}

// The folder path of the data lake store to be shared with the receiver.
func (o LookupDatasetDataLakeGen1ResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.FolderPath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatasetDataLakeGen1ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatasetDataLakeGen1ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetDataLakeGen1Result) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetDataLakeGen1ResultOutput{})
}
