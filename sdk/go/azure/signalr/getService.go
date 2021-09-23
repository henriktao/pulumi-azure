// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Azure SignalR service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/signalr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalr.LookupService(ctx, &signalr.LookupServiceArgs{
// 			Name:              "test-signalr",
// 			ResourceGroupName: "signalr-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure:signalr/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// Specifies the name of the SignalR service.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the SignalR service is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// The FQDN of the SignalR service.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The publicly accessible IP of the SignalR service.
	IpAddress string `pulumi:"ipAddress"`
	// Specifies the supported Azure location where the SignalR service exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The primary access key of the SignalR service.
	PrimaryAccessKey string `pulumi:"primaryAccessKey"`
	// The primary connection string of the SignalR service.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for browser/client use.
	PublicPort        int    `pulumi:"publicPort"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary access key of the SignalR service.
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
	// The secondary connection string of the SignalR service.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the SignalR service which is designed for customer server side use.
	ServerPort int               `pulumi:"serverPort"`
	Tags       map[string]string `pulumi:"tags"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			return *r, err
		}).(LookupServiceResultOutput)
}

// A collection of arguments for invoking getService.
type LookupServiceOutputArgs struct {
	// Specifies the name of the SignalR service.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group the SignalR service is located in.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

// A collection of values returned by getService.
type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// The FQDN of the SignalR service.
func (o LookupServiceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The publicly accessible IP of the SignalR service.
func (o LookupServiceResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// Specifies the supported Azure location where the SignalR service exists.
func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The primary access key of the SignalR service.
func (o LookupServiceResultOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

// The primary connection string of the SignalR service.
func (o LookupServiceResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// The publicly accessible port of the SignalR service which is designed for browser/client use.
func (o LookupServiceResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupServiceResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The secondary access key of the SignalR service.
func (o LookupServiceResultOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
}

// The secondary connection string of the SignalR service.
func (o LookupServiceResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// The publicly accessible port of the SignalR service which is designed for customer server side use.
func (o LookupServiceResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
