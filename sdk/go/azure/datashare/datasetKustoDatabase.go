// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Share Kusto Database Dataset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datashare"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/kusto"
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
// 		exampleAccount, err := datashare.NewAccount(ctx, "exampleAccount", &datashare.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Identity: &datashare.AccountIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleShare, err := datashare.NewShare(ctx, "exampleShare", &datashare.ShareArgs{
// 			AccountId: exampleAccount.ID(),
// 			Kind:      pulumi.String("InPlace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCluster, err := kusto.NewCluster(ctx, "exampleCluster", &kusto.ClusterArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Dev(No SLA)_Standard_D11_v2"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDatabase, err := kusto.NewDatabase(ctx, "exampleDatabase", &kusto.DatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterName:       exampleCluster.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleCluster.ID(),
// 			RoleDefinitionName: pulumi.String("Contributor"),
// 			PrincipalId: exampleAccount.Identity.ApplyT(func(identity datashare.AccountIdentity) (string, error) {
// 				return identity.PrincipalId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datashare.NewDatasetKustoDatabase(ctx, "exampleDatasetKustoDatabase", &datashare.DatasetKustoDatabaseArgs{
// 			ShareId:         exampleShare.ID(),
// 			KustoDatabaseId: exampleDatabase.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAssignment,
// 		}))
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
// Data Share Kusto Database Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datashare/datasetKustoDatabase:DatasetKustoDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
// ```
type DatasetKustoDatabase struct {
	pulumi.CustomResourceState

	// The name of the Data Share Dataset.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The location of the Kusto Cluster.
	KustoClusterLocation pulumi.StringOutput `pulumi:"kustoClusterLocation"`
	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseId pulumi.StringOutput `pulumi:"kustoDatabaseId"`
	// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareId pulumi.StringOutput `pulumi:"shareId"`
}

// NewDatasetKustoDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatasetKustoDatabase(ctx *pulumi.Context,
	name string, args *DatasetKustoDatabaseArgs, opts ...pulumi.ResourceOption) (*DatasetKustoDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KustoDatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'KustoDatabaseId'")
	}
	if args.ShareId == nil {
		return nil, errors.New("invalid value for required argument 'ShareId'")
	}
	var resource DatasetKustoDatabase
	err := ctx.RegisterResource("azure:datashare/datasetKustoDatabase:DatasetKustoDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetKustoDatabase gets an existing DatasetKustoDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetKustoDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetKustoDatabaseState, opts ...pulumi.ResourceOption) (*DatasetKustoDatabase, error) {
	var resource DatasetKustoDatabase
	err := ctx.ReadResource("azure:datashare/datasetKustoDatabase:DatasetKustoDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetKustoDatabase resources.
type datasetKustoDatabaseState struct {
	// The name of the Data Share Dataset.
	DisplayName *string `pulumi:"displayName"`
	// The location of the Kusto Cluster.
	KustoClusterLocation *string `pulumi:"kustoClusterLocation"`
	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseId *string `pulumi:"kustoDatabaseId"`
	// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareId *string `pulumi:"shareId"`
}

type DatasetKustoDatabaseState struct {
	// The name of the Data Share Dataset.
	DisplayName pulumi.StringPtrInput
	// The location of the Kusto Cluster.
	KustoClusterLocation pulumi.StringPtrInput
	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseId pulumi.StringPtrInput
	// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareId pulumi.StringPtrInput
}

func (DatasetKustoDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetKustoDatabaseState)(nil)).Elem()
}

type datasetKustoDatabaseArgs struct {
	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseId string `pulumi:"kustoDatabaseId"`
	// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
	Name *string `pulumi:"name"`
	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareId string `pulumi:"shareId"`
}

// The set of arguments for constructing a DatasetKustoDatabase resource.
type DatasetKustoDatabaseArgs struct {
	// The resource ID of the Kusto Cluster Database to be shared with the receiver. Changing this forces a new Data Share Kusto Database Dataset to be created.
	KustoDatabaseId pulumi.StringInput
	// The name which should be used for this Data Share Kusto Database Dataset. Changing this forces a new Data Share Kusto Database Dataset to be created.
	Name pulumi.StringPtrInput
	// The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created. Changing this forces a new Data Share Kusto Database Dataset to be created.
	ShareId pulumi.StringInput
}

func (DatasetKustoDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetKustoDatabaseArgs)(nil)).Elem()
}

type DatasetKustoDatabaseInput interface {
	pulumi.Input

	ToDatasetKustoDatabaseOutput() DatasetKustoDatabaseOutput
	ToDatasetKustoDatabaseOutputWithContext(ctx context.Context) DatasetKustoDatabaseOutput
}

func (*DatasetKustoDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetKustoDatabase)(nil)).Elem()
}

func (i *DatasetKustoDatabase) ToDatasetKustoDatabaseOutput() DatasetKustoDatabaseOutput {
	return i.ToDatasetKustoDatabaseOutputWithContext(context.Background())
}

func (i *DatasetKustoDatabase) ToDatasetKustoDatabaseOutputWithContext(ctx context.Context) DatasetKustoDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoDatabaseOutput)
}

// DatasetKustoDatabaseArrayInput is an input type that accepts DatasetKustoDatabaseArray and DatasetKustoDatabaseArrayOutput values.
// You can construct a concrete instance of `DatasetKustoDatabaseArrayInput` via:
//
//          DatasetKustoDatabaseArray{ DatasetKustoDatabaseArgs{...} }
type DatasetKustoDatabaseArrayInput interface {
	pulumi.Input

	ToDatasetKustoDatabaseArrayOutput() DatasetKustoDatabaseArrayOutput
	ToDatasetKustoDatabaseArrayOutputWithContext(context.Context) DatasetKustoDatabaseArrayOutput
}

type DatasetKustoDatabaseArray []DatasetKustoDatabaseInput

func (DatasetKustoDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetKustoDatabase)(nil)).Elem()
}

func (i DatasetKustoDatabaseArray) ToDatasetKustoDatabaseArrayOutput() DatasetKustoDatabaseArrayOutput {
	return i.ToDatasetKustoDatabaseArrayOutputWithContext(context.Background())
}

func (i DatasetKustoDatabaseArray) ToDatasetKustoDatabaseArrayOutputWithContext(ctx context.Context) DatasetKustoDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoDatabaseArrayOutput)
}

// DatasetKustoDatabaseMapInput is an input type that accepts DatasetKustoDatabaseMap and DatasetKustoDatabaseMapOutput values.
// You can construct a concrete instance of `DatasetKustoDatabaseMapInput` via:
//
//          DatasetKustoDatabaseMap{ "key": DatasetKustoDatabaseArgs{...} }
type DatasetKustoDatabaseMapInput interface {
	pulumi.Input

	ToDatasetKustoDatabaseMapOutput() DatasetKustoDatabaseMapOutput
	ToDatasetKustoDatabaseMapOutputWithContext(context.Context) DatasetKustoDatabaseMapOutput
}

type DatasetKustoDatabaseMap map[string]DatasetKustoDatabaseInput

func (DatasetKustoDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetKustoDatabase)(nil)).Elem()
}

func (i DatasetKustoDatabaseMap) ToDatasetKustoDatabaseMapOutput() DatasetKustoDatabaseMapOutput {
	return i.ToDatasetKustoDatabaseMapOutputWithContext(context.Background())
}

func (i DatasetKustoDatabaseMap) ToDatasetKustoDatabaseMapOutputWithContext(ctx context.Context) DatasetKustoDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetKustoDatabaseMapOutput)
}

type DatasetKustoDatabaseOutput struct{ *pulumi.OutputState }

func (DatasetKustoDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetKustoDatabase)(nil)).Elem()
}

func (o DatasetKustoDatabaseOutput) ToDatasetKustoDatabaseOutput() DatasetKustoDatabaseOutput {
	return o
}

func (o DatasetKustoDatabaseOutput) ToDatasetKustoDatabaseOutputWithContext(ctx context.Context) DatasetKustoDatabaseOutput {
	return o
}

type DatasetKustoDatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatasetKustoDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetKustoDatabase)(nil)).Elem()
}

func (o DatasetKustoDatabaseArrayOutput) ToDatasetKustoDatabaseArrayOutput() DatasetKustoDatabaseArrayOutput {
	return o
}

func (o DatasetKustoDatabaseArrayOutput) ToDatasetKustoDatabaseArrayOutputWithContext(ctx context.Context) DatasetKustoDatabaseArrayOutput {
	return o
}

func (o DatasetKustoDatabaseArrayOutput) Index(i pulumi.IntInput) DatasetKustoDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetKustoDatabase {
		return vs[0].([]*DatasetKustoDatabase)[vs[1].(int)]
	}).(DatasetKustoDatabaseOutput)
}

type DatasetKustoDatabaseMapOutput struct{ *pulumi.OutputState }

func (DatasetKustoDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetKustoDatabase)(nil)).Elem()
}

func (o DatasetKustoDatabaseMapOutput) ToDatasetKustoDatabaseMapOutput() DatasetKustoDatabaseMapOutput {
	return o
}

func (o DatasetKustoDatabaseMapOutput) ToDatasetKustoDatabaseMapOutputWithContext(ctx context.Context) DatasetKustoDatabaseMapOutput {
	return o
}

func (o DatasetKustoDatabaseMapOutput) MapIndex(k pulumi.StringInput) DatasetKustoDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetKustoDatabase {
		return vs[0].(map[string]*DatasetKustoDatabase)[vs[1].(string)]
	}).(DatasetKustoDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetKustoDatabaseInput)(nil)).Elem(), &DatasetKustoDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetKustoDatabaseArrayInput)(nil)).Elem(), DatasetKustoDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetKustoDatabaseMapInput)(nil)).Elem(), DatasetKustoDatabaseMap{})
	pulumi.RegisterOutputType(DatasetKustoDatabaseOutput{})
	pulumi.RegisterOutputType(DatasetKustoDatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatasetKustoDatabaseMapOutput{})
}
