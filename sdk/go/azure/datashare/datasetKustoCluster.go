// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Share Kusto Cluster Dataset.
//
// ## Import
//
// Data Share Kusto Cluster Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datashare/datasetKustoCluster:DatasetKustoCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
// ```
type DatasetKustoCluster struct {
	pulumi.CustomResourceState

	// The name of the Data Share Dataset.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterId pulumi.StringOutput `pulumi:"kustoClusterId"`
	// The location of the Kusto Cluster.
	KustoClusterLocation pulumi.StringOutput `pulumi:"kustoClusterLocation"`
	// The name which should be used for this Data Share Kusto Cluster Dataset. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareId pulumi.StringOutput `pulumi:"shareId"`
}

// NewDatasetKustoCluster registers a new resource with the given unique name, arguments, and options.
func NewDatasetKustoCluster(ctx *pulumi.Context,
	name string, args *DatasetKustoClusterArgs, opts ...pulumi.ResourceOption) (*DatasetKustoCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KustoClusterId == nil {
		return nil, errors.New("invalid value for required argument 'KustoClusterId'")
	}
	if args.ShareId == nil {
		return nil, errors.New("invalid value for required argument 'ShareId'")
	}
	var resource DatasetKustoCluster
	err := ctx.RegisterResource("azure:datashare/datasetKustoCluster:DatasetKustoCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetKustoCluster gets an existing DatasetKustoCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetKustoCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetKustoClusterState, opts ...pulumi.ResourceOption) (*DatasetKustoCluster, error) {
	var resource DatasetKustoCluster
	err := ctx.ReadResource("azure:datashare/datasetKustoCluster:DatasetKustoCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetKustoCluster resources.
type datasetKustoClusterState struct {
	// The name of the Data Share Dataset.
	DisplayName *string `pulumi:"displayName"`
	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterId *string `pulumi:"kustoClusterId"`
	// The location of the Kusto Cluster.
	KustoClusterLocation *string `pulumi:"kustoClusterLocation"`
	// The name which should be used for this Data Share Kusto Cluster Dataset. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareId *string `pulumi:"shareId"`
}

type DatasetKustoClusterState struct {
	// The name of the Data Share Dataset.
	DisplayName pulumi.StringPtrInput
	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterId pulumi.StringPtrInput
	// The location of the Kusto Cluster.
	KustoClusterLocation pulumi.StringPtrInput
	// The name which should be used for this Data Share Kusto Cluster Dataset. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareId pulumi.StringPtrInput
}

func (DatasetKustoClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetKustoClusterState)(nil)).Elem()
}

type datasetKustoClusterArgs struct {
	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterId string `pulumi:"kustoClusterId"`
	// The name which should be used for this Data Share Kusto Cluster Dataset. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareId string `pulumi:"shareId"`
}

// The set of arguments for constructing a DatasetKustoCluster resource.
type DatasetKustoClusterArgs struct {
	// The resource ID of the Kusto Cluster to be shared with the receiver. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	KustoClusterId pulumi.StringInput
	// The name which should be used for this Data Share Kusto Cluster Dataset. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Kusto Cluster Dataset should be created. Changing this forces a new Data Share Kusto Cluster Dataset to be created.
	ShareId pulumi.StringInput
}

func (DatasetKustoClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetKustoClusterArgs)(nil)).Elem()
}

type DatasetKustoClusterInput interface {
	pulumi.Input

	ToDatasetKustoClusterOutput() DatasetKustoClusterOutput
	ToDatasetKustoClusterOutputWithContext(ctx context.Context) DatasetKustoClusterOutput
}

func (*DatasetKustoCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetKustoCluster)(nil))
}

func (i *DatasetKustoCluster) ToDatasetKustoClusterOutput() DatasetKustoClusterOutput {
	return i.ToDatasetKustoClusterOutputWithContext(context.Background())
}

func (i *DatasetKustoCluster) ToDatasetKustoClusterOutputWithContext(ctx context.Context) DatasetKustoClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoClusterOutput)
}

func (i *DatasetKustoCluster) ToDatasetKustoClusterPtrOutput() DatasetKustoClusterPtrOutput {
	return i.ToDatasetKustoClusterPtrOutputWithContext(context.Background())
}

func (i *DatasetKustoCluster) ToDatasetKustoClusterPtrOutputWithContext(ctx context.Context) DatasetKustoClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoClusterPtrOutput)
}

type DatasetKustoClusterPtrInput interface {
	pulumi.Input

	ToDatasetKustoClusterPtrOutput() DatasetKustoClusterPtrOutput
	ToDatasetKustoClusterPtrOutputWithContext(ctx context.Context) DatasetKustoClusterPtrOutput
}

type datasetKustoClusterPtrType DatasetKustoClusterArgs

func (*datasetKustoClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetKustoCluster)(nil))
}

func (i *datasetKustoClusterPtrType) ToDatasetKustoClusterPtrOutput() DatasetKustoClusterPtrOutput {
	return i.ToDatasetKustoClusterPtrOutputWithContext(context.Background())
}

func (i *datasetKustoClusterPtrType) ToDatasetKustoClusterPtrOutputWithContext(ctx context.Context) DatasetKustoClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoClusterPtrOutput)
}

// DatasetKustoClusterArrayInput is an input type that accepts DatasetKustoClusterArray and DatasetKustoClusterArrayOutput values.
// You can construct a concrete instance of `DatasetKustoClusterArrayInput` via:
//
//          DatasetKustoClusterArray{ DatasetKustoClusterArgs{...} }
type DatasetKustoClusterArrayInput interface {
	pulumi.Input

	ToDatasetKustoClusterArrayOutput() DatasetKustoClusterArrayOutput
	ToDatasetKustoClusterArrayOutputWithContext(context.Context) DatasetKustoClusterArrayOutput
}

type DatasetKustoClusterArray []DatasetKustoClusterInput

func (DatasetKustoClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetKustoCluster)(nil)).Elem()
}

func (i DatasetKustoClusterArray) ToDatasetKustoClusterArrayOutput() DatasetKustoClusterArrayOutput {
	return i.ToDatasetKustoClusterArrayOutputWithContext(context.Background())
}

func (i DatasetKustoClusterArray) ToDatasetKustoClusterArrayOutputWithContext(ctx context.Context) DatasetKustoClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoClusterArrayOutput)
}

// DatasetKustoClusterMapInput is an input type that accepts DatasetKustoClusterMap and DatasetKustoClusterMapOutput values.
// You can construct a concrete instance of `DatasetKustoClusterMapInput` via:
//
//          DatasetKustoClusterMap{ "key": DatasetKustoClusterArgs{...} }
type DatasetKustoClusterMapInput interface {
	pulumi.Input

	ToDatasetKustoClusterMapOutput() DatasetKustoClusterMapOutput
	ToDatasetKustoClusterMapOutputWithContext(context.Context) DatasetKustoClusterMapOutput
}

type DatasetKustoClusterMap map[string]DatasetKustoClusterInput

func (DatasetKustoClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetKustoCluster)(nil)).Elem()
}

func (i DatasetKustoClusterMap) ToDatasetKustoClusterMapOutput() DatasetKustoClusterMapOutput {
	return i.ToDatasetKustoClusterMapOutputWithContext(context.Background())
}

func (i DatasetKustoClusterMap) ToDatasetKustoClusterMapOutputWithContext(ctx context.Context) DatasetKustoClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoClusterMapOutput)
}

type DatasetKustoClusterOutput struct{ *pulumi.OutputState }

func (DatasetKustoClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetKustoCluster)(nil))
}

func (o DatasetKustoClusterOutput) ToDatasetKustoClusterOutput() DatasetKustoClusterOutput {
	return o
}

func (o DatasetKustoClusterOutput) ToDatasetKustoClusterOutputWithContext(ctx context.Context) DatasetKustoClusterOutput {
	return o
}

func (o DatasetKustoClusterOutput) ToDatasetKustoClusterPtrOutput() DatasetKustoClusterPtrOutput {
	return o.ToDatasetKustoClusterPtrOutputWithContext(context.Background())
}

func (o DatasetKustoClusterOutput) ToDatasetKustoClusterPtrOutputWithContext(ctx context.Context) DatasetKustoClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetKustoCluster) *DatasetKustoCluster {
		return &v
	}).(DatasetKustoClusterPtrOutput)
}

type DatasetKustoClusterPtrOutput struct{ *pulumi.OutputState }

func (DatasetKustoClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetKustoCluster)(nil))
}

func (o DatasetKustoClusterPtrOutput) ToDatasetKustoClusterPtrOutput() DatasetKustoClusterPtrOutput {
	return o
}

func (o DatasetKustoClusterPtrOutput) ToDatasetKustoClusterPtrOutputWithContext(ctx context.Context) DatasetKustoClusterPtrOutput {
	return o
}

func (o DatasetKustoClusterPtrOutput) Elem() DatasetKustoClusterOutput {
	return o.ApplyT(func(v *DatasetKustoCluster) DatasetKustoCluster {
		if v != nil {
			return *v
		}
		var ret DatasetKustoCluster
		return ret
	}).(DatasetKustoClusterOutput)
}

type DatasetKustoClusterArrayOutput struct{ *pulumi.OutputState }

func (DatasetKustoClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetKustoCluster)(nil))
}

func (o DatasetKustoClusterArrayOutput) ToDatasetKustoClusterArrayOutput() DatasetKustoClusterArrayOutput {
	return o
}

func (o DatasetKustoClusterArrayOutput) ToDatasetKustoClusterArrayOutputWithContext(ctx context.Context) DatasetKustoClusterArrayOutput {
	return o
}

func (o DatasetKustoClusterArrayOutput) Index(i pulumi.IntInput) DatasetKustoClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetKustoCluster {
		return vs[0].([]DatasetKustoCluster)[vs[1].(int)]
	}).(DatasetKustoClusterOutput)
}

type DatasetKustoClusterMapOutput struct{ *pulumi.OutputState }

func (DatasetKustoClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatasetKustoCluster)(nil))
}

func (o DatasetKustoClusterMapOutput) ToDatasetKustoClusterMapOutput() DatasetKustoClusterMapOutput {
	return o
}

func (o DatasetKustoClusterMapOutput) ToDatasetKustoClusterMapOutputWithContext(ctx context.Context) DatasetKustoClusterMapOutput {
	return o
}

func (o DatasetKustoClusterMapOutput) MapIndex(k pulumi.StringInput) DatasetKustoClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatasetKustoCluster {
		return vs[0].(map[string]DatasetKustoCluster)[vs[1].(string)]
	}).(DatasetKustoClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetKustoClusterOutput{})
	pulumi.RegisterOutputType(DatasetKustoClusterPtrOutput{})
	pulumi.RegisterOutputType(DatasetKustoClusterArrayOutput{})
	pulumi.RegisterOutputType(DatasetKustoClusterMapOutput{})
}
