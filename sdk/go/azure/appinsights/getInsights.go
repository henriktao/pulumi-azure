// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Application Insights component.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := appinsights.LookupInsights(ctx, &appinsights.LookupInsightsArgs{
// 			Name:              "production",
// 			ResourceGroupName: "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("applicationInsightsInstrumentationKey", example.InstrumentationKey)
// 		return nil
// 	})
// }
// ```
func LookupInsights(ctx *pulumi.Context, args *LookupInsightsArgs, opts ...pulumi.InvokeOption) (*LookupInsightsResult, error) {
	var rv LookupInsightsResult
	err := ctx.Invoke("azure:appinsights/getInsights:getInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInsights.
type LookupInsightsArgs struct {
	// Specifies the name of the Application Insights component.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Application Insights component is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getInsights.
type LookupInsightsResult struct {
	// The App ID associated with this Application Insights component.
	AppId string `pulumi:"appId"`
	// The type of the component.
	ApplicationType string `pulumi:"applicationType"`
	// The connection string of the Application Insights component. (Sensitive)
	ConnectionString string `pulumi:"connectionString"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instrumentation key of the Application Insights component.
	InstrumentationKey string `pulumi:"instrumentationKey"`
	// The Azure location where the component exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The retention period in days.
	RetentionInDays int `pulumi:"retentionInDays"`
	// Tags applied to the component.
	Tags map[string]string `pulumi:"tags"`
	// The id of the associated Log Analytics workspace
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupInsightsOutput(ctx *pulumi.Context, args LookupInsightsOutputArgs, opts ...pulumi.InvokeOption) LookupInsightsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInsightsResult, error) {
			args := v.(LookupInsightsArgs)
			r, err := LookupInsights(ctx, &args, opts...)
			return *r, err
		}).(LookupInsightsResultOutput)
}

// A collection of arguments for invoking getInsights.
type LookupInsightsOutputArgs struct {
	// Specifies the name of the Application Insights component.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the name of the resource group the Application Insights component is located in.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInsightsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInsightsArgs)(nil)).Elem()
}

// A collection of values returned by getInsights.
type LookupInsightsResultOutput struct{ *pulumi.OutputState }

func (LookupInsightsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInsightsResult)(nil)).Elem()
}

func (o LookupInsightsResultOutput) ToLookupInsightsResultOutput() LookupInsightsResultOutput {
	return o
}

func (o LookupInsightsResultOutput) ToLookupInsightsResultOutputWithContext(ctx context.Context) LookupInsightsResultOutput {
	return o
}

// The App ID associated with this Application Insights component.
func (o LookupInsightsResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.AppId }).(pulumi.StringOutput)
}

// The type of the component.
func (o LookupInsightsResultOutput) ApplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.ApplicationType }).(pulumi.StringOutput)
}

// The connection string of the Application Insights component. (Sensitive)
func (o LookupInsightsResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInsightsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The instrumentation key of the Application Insights component.
func (o LookupInsightsResultOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.InstrumentationKey }).(pulumi.StringOutput)
}

// The Azure location where the component exists.
func (o LookupInsightsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupInsightsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInsightsResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The retention period in days.
func (o LookupInsightsResultOutput) RetentionInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInsightsResult) int { return v.RetentionInDays }).(pulumi.IntOutput)
}

// Tags applied to the component.
func (o LookupInsightsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInsightsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The id of the associated Log Analytics workspace
func (o LookupInsightsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInsightsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInsightsResultOutput{})
}
