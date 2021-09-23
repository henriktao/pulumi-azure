// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing SQL elastic pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := mssql.LookupElasticPool(ctx, &mssql.LookupElasticPoolArgs{
// 			Name:              "mssqlelasticpoolname",
// 			ResourceGroupName: "example-resources",
// 			ServerName:        "example-sql-server",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("elasticpoolId", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupElasticPool(ctx *pulumi.Context, args *LookupElasticPoolArgs, opts ...pulumi.InvokeOption) (*LookupElasticPoolResult, error) {
	var rv LookupElasticPoolResult
	err := ctx.Invoke("azure:mssql/getElasticPool:getElasticPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticPool.
type LookupElasticPoolArgs struct {
	// The name of the elastic pool.
	Name string `pulumi:"name"`
	// The name of the resource group which contains the elastic pool.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SQL Server which contains the elastic pool.
	ServerName string `pulumi:"serverName"`
}

// A collection of values returned by getElasticPool.
type LookupElasticPoolResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The license type to apply for this database.
	LicenseType string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The max data size of the elastic pool in bytes.
	MaxSizeBytes int `pulumi:"maxSizeBytes"`
	// The max data size of the elastic pool in gigabytes.
	MaxSizeGb float64 `pulumi:"maxSizeGb"`
	// Specifies the SKU Name for this Elasticpool.
	Name string `pulumi:"name"`
	// The maximum capacity any one database can consume.
	PerDbMaxCapacity int `pulumi:"perDbMaxCapacity"`
	// The minimum capacity all databases are guaranteed.
	PerDbMinCapacity  int    `pulumi:"perDbMinCapacity"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	// A `sku` block as defined below.
	Skus []GetElasticPoolSkus `pulumi:"skus"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this elastic pool is zone redundant.
	ZoneRedundant bool `pulumi:"zoneRedundant"`
}

func LookupElasticPoolOutput(ctx *pulumi.Context, args LookupElasticPoolOutputArgs, opts ...pulumi.InvokeOption) LookupElasticPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticPoolResult, error) {
			args := v.(LookupElasticPoolArgs)
			r, err := LookupElasticPool(ctx, &args, opts...)
			return *r, err
		}).(LookupElasticPoolResultOutput)
}

// A collection of arguments for invoking getElasticPool.
type LookupElasticPoolOutputArgs struct {
	// The name of the elastic pool.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group which contains the elastic pool.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the SQL Server which contains the elastic pool.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupElasticPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolArgs)(nil)).Elem()
}

// A collection of values returned by getElasticPool.
type LookupElasticPoolResultOutput struct{ *pulumi.OutputState }

func (LookupElasticPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolResult)(nil)).Elem()
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutput() LookupElasticPoolResultOutput {
	return o
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutputWithContext(ctx context.Context) LookupElasticPoolResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupElasticPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The license type to apply for this database.
func (o LookupElasticPoolResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

// Specifies the supported Azure location where the resource exists.
func (o LookupElasticPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The max data size of the elastic pool in bytes.
func (o LookupElasticPoolResultOutput) MaxSizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) int { return v.MaxSizeBytes }).(pulumi.IntOutput)
}

// The max data size of the elastic pool in gigabytes.
func (o LookupElasticPoolResultOutput) MaxSizeGb() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticPoolResult) float64 { return v.MaxSizeGb }).(pulumi.Float64Output)
}

// Specifies the SKU Name for this Elasticpool.
func (o LookupElasticPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The maximum capacity any one database can consume.
func (o LookupElasticPoolResultOutput) PerDbMaxCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) int { return v.PerDbMaxCapacity }).(pulumi.IntOutput)
}

// The minimum capacity all databases are guaranteed.
func (o LookupElasticPoolResultOutput) PerDbMinCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) int { return v.PerDbMinCapacity }).(pulumi.IntOutput)
}

func (o LookupElasticPoolResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.ServerName }).(pulumi.StringOutput)
}

// A `sku` block as defined below.
func (o LookupElasticPoolResultOutput) Skus() GetElasticPoolSkusArrayOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) []GetElasticPoolSkus { return v.Skus }).(GetElasticPoolSkusArrayOutput)
}

// A mapping of tags to assign to the resource.
func (o LookupElasticPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether or not this elastic pool is zone redundant.
func (o LookupElasticPoolResultOutput) ZoneRedundant() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) bool { return v.ZoneRedundant }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticPoolResultOutput{})
}
