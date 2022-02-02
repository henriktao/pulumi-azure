// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// !> **Important** Due to capacity constraints, Microsoft requires you to pre-register your subscription IDs before you are allowed to create a Log Analytics cluster. Contact Microsoft, or open a support request to register your subscription IDs.
//
// > **Note:** Log Analytics Clusters are subject to 14-day soft delete policy. Clusters created with the same resource group & name as a previously deleted cluster will be recovered rather than creating anew.
//
// Manages a Log Analytics Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/loganalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loganalytics.NewCluster(ctx, "exampleCluster", &loganalytics.ClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Identity: &loganalytics.ClusterIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
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
// Log Analytics Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:loganalytics/cluster:Cluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The GUID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// An `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
	Identity ClusterIdentityOutput `pulumi:"identity"`
	// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
	SizeGb pulumi.IntPtrOutput `pulumi:"sizeGb"`
	// A mapping of tags which should be assigned to the Log Analytics Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Cluster
	err := ctx.RegisterResource("azure:loganalytics/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:loganalytics/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The GUID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// An `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
	Identity *ClusterIdentity `pulumi:"identity"`
	// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
	SizeGb *int `pulumi:"sizeGb"`
	// A mapping of tags which should be assigned to the Log Analytics Cluster.
	Tags map[string]string `pulumi:"tags"`
}

type ClusterState struct {
	// The GUID of the cluster.
	ClusterId pulumi.StringPtrInput
	// An `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
	Identity ClusterIdentityPtrInput
	// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
	SizeGb pulumi.IntPtrInput
	// A mapping of tags which should be assigned to the Log Analytics Cluster.
	Tags pulumi.StringMapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// An `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
	Identity ClusterIdentity `pulumi:"identity"`
	// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
	SizeGb *int `pulumi:"sizeGb"`
	// A mapping of tags which should be assigned to the Log Analytics Cluster.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// An `identity` block as defined below. Changing this forces a new Log Analytics Cluster to be created.
	Identity ClusterIdentityInput
	// The Azure Region where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Log Analytics Cluster. Changing this forces a new Log Analytics Cluster to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Cluster should exist. Changing this forces a new Log Analytics Cluster to be created.
	ResourceGroupName pulumi.StringInput
	// The capacity of the Log Analytics Cluster specified in GB/day. Defaults to 1000.
	SizeGb pulumi.IntPtrInput
	// A mapping of tags which should be assigned to the Log Analytics Cluster.
	Tags pulumi.StringMapInput
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
