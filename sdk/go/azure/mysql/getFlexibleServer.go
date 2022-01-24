// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing MySQL Flexible Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := mysql.LookupFlexibleServer(ctx, &mysql.LookupFlexibleServerArgs{
// 			Name:              "existingMySqlFlexibleServer",
// 			ResourceGroupName: "existingResGroup",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupFlexibleServer(ctx *pulumi.Context, args *LookupFlexibleServerArgs, opts ...pulumi.InvokeOption) (*LookupFlexibleServerResult, error) {
	var rv LookupFlexibleServerResult
	err := ctx.Invoke("azure:mysql/getFlexibleServer:getFlexibleServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlexibleServer.
type LookupFlexibleServerArgs struct {
	// Specifies the name of the MySQL Flexible Server.
	Name string `pulumi:"name"`
	// The name of the resource group for the MySQL Flexible Server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getFlexibleServer.
type LookupFlexibleServerResult struct {
	// The Administrator Login of the MySQL Flexible Server.
	AdministratorLogin string `pulumi:"administratorLogin"`
	// The backup retention days of the MySQL Flexible Server.
	BackupRetentionDays int `pulumi:"backupRetentionDays"`
	// The ID of the virtual network subnet the MySQL Flexible Server is created in.
	DelegatedSubnetId string `pulumi:"delegatedSubnetId"`
	// The fully qualified domain name of the MySQL Flexible Server.
	Fqdn string `pulumi:"fqdn"`
	// Is geo redundant backup enabled?
	GeoRedundantBackupEnabled bool `pulumi:"geoRedundantBackupEnabled"`
	// A `highAvailability` block for this MySQL Flexible Server as defined below.
	HighAvailabilities []GetFlexibleServerHighAvailability `pulumi:"highAvailabilities"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region of the MySQL Flexible Server.
	Location string `pulumi:"location"`
	// A `maintenanceWindow` block for this MySQL Flexible Server as defined below.
	MaintenanceWindows []GetFlexibleServerMaintenanceWindow `pulumi:"maintenanceWindows"`
	Name               string                               `pulumi:"name"`
	// The ID of the private dns zone of the MySQL Flexible Server.
	PrivateDnsZoneId string `pulumi:"privateDnsZoneId"`
	// Is the public network access enabled?
	PublicNetworkAccessEnabled bool `pulumi:"publicNetworkAccessEnabled"`
	// The maximum number of replicas that a primary MySQL Flexible Server can have.
	ReplicaCapacity int `pulumi:"replicaCapacity"`
	// The replication role of the MySQL Flexible Server.
	ReplicationRole    string `pulumi:"replicationRole"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RestorePointInTime string `pulumi:"restorePointInTime"`
	// The SKU Name of the MySQL Flexible Server.
	SkuName string `pulumi:"skuName"`
	// A `storage` block for this MySQL Flexible Server as defined below.
	Storages []GetFlexibleServerStorage `pulumi:"storages"`
	// A mapping of tags which are assigned to the MySQL Flexible Server.
	Tags map[string]string `pulumi:"tags"`
	// The version of the MySQL Flexible Server.
	Version string `pulumi:"version"`
	// The availability zones of the MySQL Flexible Server.
	Zone string `pulumi:"zone"`
}

func LookupFlexibleServerOutput(ctx *pulumi.Context, args LookupFlexibleServerOutputArgs, opts ...pulumi.InvokeOption) LookupFlexibleServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlexibleServerResult, error) {
			args := v.(LookupFlexibleServerArgs)
			r, err := LookupFlexibleServer(ctx, &args, opts...)
			return *r, err
		}).(LookupFlexibleServerResultOutput)
}

// A collection of arguments for invoking getFlexibleServer.
type LookupFlexibleServerOutputArgs struct {
	// Specifies the name of the MySQL Flexible Server.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group for the MySQL Flexible Server.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFlexibleServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlexibleServerArgs)(nil)).Elem()
}

// A collection of values returned by getFlexibleServer.
type LookupFlexibleServerResultOutput struct{ *pulumi.OutputState }

func (LookupFlexibleServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlexibleServerResult)(nil)).Elem()
}

func (o LookupFlexibleServerResultOutput) ToLookupFlexibleServerResultOutput() LookupFlexibleServerResultOutput {
	return o
}

func (o LookupFlexibleServerResultOutput) ToLookupFlexibleServerResultOutputWithContext(ctx context.Context) LookupFlexibleServerResultOutput {
	return o
}

// The Administrator Login of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) AdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.AdministratorLogin }).(pulumi.StringOutput)
}

// The backup retention days of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) BackupRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) int { return v.BackupRetentionDays }).(pulumi.IntOutput)
}

// The ID of the virtual network subnet the MySQL Flexible Server is created in.
func (o LookupFlexibleServerResultOutput) DelegatedSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.DelegatedSubnetId }).(pulumi.StringOutput)
}

// The fully qualified domain name of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Is geo redundant backup enabled?
func (o LookupFlexibleServerResultOutput) GeoRedundantBackupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) bool { return v.GeoRedundantBackupEnabled }).(pulumi.BoolOutput)
}

// A `highAvailability` block for this MySQL Flexible Server as defined below.
func (o LookupFlexibleServerResultOutput) HighAvailabilities() GetFlexibleServerHighAvailabilityArrayOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) []GetFlexibleServerHighAvailability { return v.HighAvailabilities }).(GetFlexibleServerHighAvailabilityArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFlexibleServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Region of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// A `maintenanceWindow` block for this MySQL Flexible Server as defined below.
func (o LookupFlexibleServerResultOutput) MaintenanceWindows() GetFlexibleServerMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) []GetFlexibleServerMaintenanceWindow { return v.MaintenanceWindows }).(GetFlexibleServerMaintenanceWindowArrayOutput)
}

func (o LookupFlexibleServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the private dns zone of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) PrivateDnsZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.PrivateDnsZoneId }).(pulumi.StringOutput)
}

// Is the public network access enabled?
func (o LookupFlexibleServerResultOutput) PublicNetworkAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) bool { return v.PublicNetworkAccessEnabled }).(pulumi.BoolOutput)
}

// The maximum number of replicas that a primary MySQL Flexible Server can have.
func (o LookupFlexibleServerResultOutput) ReplicaCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) int { return v.ReplicaCapacity }).(pulumi.IntOutput)
}

// The replication role of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

func (o LookupFlexibleServerResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupFlexibleServerResultOutput) RestorePointInTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.RestorePointInTime }).(pulumi.StringOutput)
}

// The SKU Name of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.SkuName }).(pulumi.StringOutput)
}

// A `storage` block for this MySQL Flexible Server as defined below.
func (o LookupFlexibleServerResultOutput) Storages() GetFlexibleServerStorageArrayOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) []GetFlexibleServerStorage { return v.Storages }).(GetFlexibleServerStorageArrayOutput)
}

// A mapping of tags which are assigned to the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The version of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Version }).(pulumi.StringOutput)
}

// The availability zones of the MySQL Flexible Server.
func (o LookupFlexibleServerResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlexibleServerResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlexibleServerResultOutput{})
}
