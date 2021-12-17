// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to fetch the Host Keys of an existing Function App
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appservice.GetFunctionAppHostKeys(ctx, &appservice.GetFunctionAppHostKeysArgs{
// 			Name:              "example-function",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetFunctionAppHostKeys(ctx *pulumi.Context, args *GetFunctionAppHostKeysArgs, opts ...pulumi.InvokeOption) (*GetFunctionAppHostKeysResult, error) {
	var rv GetFunctionAppHostKeysResult
	err := ctx.Invoke("azure:appservice/getFunctionAppHostKeys:getFunctionAppHostKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunctionAppHostKeys.
type GetFunctionAppHostKeysArgs struct {
	// The name of the Function App.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Function App exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getFunctionAppHostKeys.
type GetFunctionAppHostKeysResult struct {
	// Function App resource's default function key.
	DefaultFunctionKey string `pulumi:"defaultFunctionKey"`
	// Function App resource's Durable Task Extension system key.
	DurabletaskExtensionKey string `pulumi:"durabletaskExtensionKey"`
	// Function App resource's Event Grid Extension Config system key.
	EventGridExtensionConfigKey string `pulumi:"eventGridExtensionConfigKey"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Function App resource's secret key
	//
	// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
	MasterKey         string `pulumi:"masterKey"`
	Name              string `pulumi:"name"`
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Function App resource's SignalR Extension system key.
	SignalrExtensionKey string `pulumi:"signalrExtensionKey"`
}

func GetFunctionAppHostKeysOutput(ctx *pulumi.Context, args GetFunctionAppHostKeysOutputArgs, opts ...pulumi.InvokeOption) GetFunctionAppHostKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFunctionAppHostKeysResult, error) {
			args := v.(GetFunctionAppHostKeysArgs)
			r, err := GetFunctionAppHostKeys(ctx, &args, opts...)
			return *r, err
		}).(GetFunctionAppHostKeysResultOutput)
}

// A collection of arguments for invoking getFunctionAppHostKeys.
type GetFunctionAppHostKeysOutputArgs struct {
	// The name of the Function App.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Function App exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetFunctionAppHostKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionAppHostKeysArgs)(nil)).Elem()
}

// A collection of values returned by getFunctionAppHostKeys.
type GetFunctionAppHostKeysResultOutput struct{ *pulumi.OutputState }

func (GetFunctionAppHostKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionAppHostKeysResult)(nil)).Elem()
}

func (o GetFunctionAppHostKeysResultOutput) ToGetFunctionAppHostKeysResultOutput() GetFunctionAppHostKeysResultOutput {
	return o
}

func (o GetFunctionAppHostKeysResultOutput) ToGetFunctionAppHostKeysResultOutputWithContext(ctx context.Context) GetFunctionAppHostKeysResultOutput {
	return o
}

// Function App resource's default function key.
func (o GetFunctionAppHostKeysResultOutput) DefaultFunctionKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.DefaultFunctionKey }).(pulumi.StringOutput)
}

// Function App resource's Durable Task Extension system key.
func (o GetFunctionAppHostKeysResultOutput) DurabletaskExtensionKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.DurabletaskExtensionKey }).(pulumi.StringOutput)
}

// Function App resource's Event Grid Extension Config system key.
func (o GetFunctionAppHostKeysResultOutput) EventGridExtensionConfigKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.EventGridExtensionConfigKey }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFunctionAppHostKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Function App resource's secret key
//
// Deprecated: This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
func (o GetFunctionAppHostKeysResultOutput) MasterKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.MasterKey }).(pulumi.StringOutput)
}

func (o GetFunctionAppHostKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetFunctionAppHostKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o GetFunctionAppHostKeysResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// Function App resource's SignalR Extension system key.
func (o GetFunctionAppHostKeysResultOutput) SignalrExtensionKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionAppHostKeysResult) string { return v.SignalrExtensionKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFunctionAppHostKeysResultOutput{})
}
