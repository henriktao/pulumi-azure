// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/kusto"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewCluster(ctx, "example", &kusto.ClusterArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Kusto Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringOutput `pulumi:"dataIngestionUri"`
	// Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
	DoubleEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"doubleEncryptionEnabled"`
	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop pulumi.BoolPtrOutput `pulumi:"enableAutoStop"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrOutput `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrOutput `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrOutput `pulumi:"enableStreamingIngest"`
	// . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// An `identity` block as defined below.
	Identity ClusterIdentityPtrOutput `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayOutput `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrOutput `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use `trustedExternalTenants = ["*"]` to explicitly allow all other tenants, `trustedExternalTenants = ["MyTentantOnly"]` for only your tenant or `trustedExternalTenants = ["<tenantId1>", "<tenantIdx>"]` to allow specific other tenants.
	TrustedExternalTenants pulumi.StringArrayOutput `pulumi:"trustedExternalTenants"`
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// A `virtualNetworkConfiguration` block as defined below. Changing this forces a new resource to be created.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrOutput `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Cluster
	err := ctx.RegisterResource("azure:kusto/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure:kusto/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri *string `pulumi:"dataIngestionUri"`
	// Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
	DoubleEncryptionEnabled *bool `pulumi:"doubleEncryptionEnabled"`
	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop *bool `pulumi:"enableAutoStop"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge *bool `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
	Engine *string `pulumi:"engine"`
	// An `identity` block as defined below.
	Identity *ClusterIdentity `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions []string `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale *ClusterOptimizedAutoScale `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku *ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use `trustedExternalTenants = ["*"]` to explicitly allow all other tenants, `trustedExternalTenants = ["MyTentantOnly"]` for only your tenant or `trustedExternalTenants = ["<tenantId1>", "<tenantIdx>"]` to allow specific other tenants.
	TrustedExternalTenants []string `pulumi:"trustedExternalTenants"`
	// The FQDN of the Azure Kusto Cluster.
	Uri *string `pulumi:"uri"`
	// A `virtualNetworkConfiguration` block as defined below. Changing this forces a new resource to be created.
	VirtualNetworkConfiguration *ClusterVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

type ClusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringPtrInput
	// Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
	DoubleEncryptionEnabled pulumi.BoolPtrInput
	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop pulumi.BoolPtrInput
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
	Engine pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity ClusterIdentityPtrInput
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `sku` block as defined below.
	Sku ClusterSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use `trustedExternalTenants = ["*"]` to explicitly allow all other tenants, `trustedExternalTenants = ["MyTentantOnly"]` for only your tenant or `trustedExternalTenants = ["<tenantId1>", "<tenantIdx>"]` to allow specific other tenants.
	TrustedExternalTenants pulumi.StringArrayInput
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringPtrInput
	// A `virtualNetworkConfiguration` block as defined below. Changing this forces a new resource to be created.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrInput
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
	DoubleEncryptionEnabled *bool `pulumi:"doubleEncryptionEnabled"`
	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop *bool `pulumi:"enableAutoStop"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the purge operations are enabled.
	EnablePurge *bool `pulumi:"enablePurge"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
	Engine *string `pulumi:"engine"`
	// An `identity` block as defined below.
	Identity *ClusterIdentity `pulumi:"identity"`
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions []string `pulumi:"languageExtensions"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale *ClusterOptimizedAutoScale `pulumi:"optimizedAutoScale"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use `trustedExternalTenants = ["*"]` to explicitly allow all other tenants, `trustedExternalTenants = ["MyTentantOnly"]` for only your tenant or `trustedExternalTenants = ["<tenantId1>", "<tenantIdx>"]` to allow specific other tenants.
	TrustedExternalTenants []string `pulumi:"trustedExternalTenants"`
	// A `virtualNetworkConfiguration` block as defined below. Changing this forces a new resource to be created.
	VirtualNetworkConfiguration *ClusterVirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Is the cluster's double encryption enabled? Defaults to `false`. Changing this forces a new resource to be created.
	DoubleEncryptionEnabled pulumi.BoolPtrInput
	// Specifies if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop pulumi.BoolPtrInput
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// . The engine type that should be used. Possible values are `V2` and `V3`. Defaults to `V2`.
	Engine pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity ClusterIdentityPtrInput
	// An list of `languageExtensions` to enable. Valid values are: `PYTHON` and `R`.
	LanguageExtensions pulumi.StringArrayInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// An `optimizedAutoScale` block as defined below.
	OptimizedAutoScale ClusterOptimizedAutoScalePtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as defined below.
	Sku ClusterSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a list of tenant IDs that are trusted by the cluster. Default setting trusts all other tenants. Use `trustedExternalTenants = ["*"]` to explicitly allow all other tenants, `trustedExternalTenants = ["MyTentantOnly"]` for only your tenant or `trustedExternalTenants = ["<tenantId1>", "<tenantIdx>"]` to allow specific other tenants.
	TrustedExternalTenants pulumi.StringArrayInput
	// A `virtualNetworkConfiguration` block as defined below. Changing this forces a new resource to be created.
	VirtualNetworkConfiguration ClusterVirtualNetworkConfigurationPtrInput
	// A list of Availability Zones in which the cluster instances should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
