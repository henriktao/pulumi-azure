// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a PostgreSQL Dataset inside a Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLinkedServicePostgresql, err := datafactory.NewLinkedServicePostgresql(ctx, "exampleLinkedServicePostgresql", &datafactory.LinkedServicePostgresqlArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Host=example;Port=5432;Database=example;UID=example;EncryptionMethod=0;Password=example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetPostgresql(ctx, "exampleDatasetPostgresql", &datafactory.DatasetPostgresqlArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			LinkedServiceName: exampleLinkedServicePostgresql.Name,
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
// Data Factory PostgreSQL Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/datasetPostgresql:DatasetPostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetPostgresql struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset PostgreSQL.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset PostgreSQL.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetPostgresqlSchemaColumnArrayOutput `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset PostgreSQL.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewDatasetPostgresql registers a new resource with the given unique name, arguments, and options.
func NewDatasetPostgresql(ctx *pulumi.Context,
	name string, args *DatasetPostgresqlArgs, opts ...pulumi.ResourceOption) (*DatasetPostgresql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetPostgresql
	err := ctx.RegisterResource("azure:datafactory/datasetPostgresql:DatasetPostgresql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetPostgresql gets an existing DatasetPostgresql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetPostgresql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetPostgresqlState, opts ...pulumi.ResourceOption) (*DatasetPostgresql, error) {
	var resource DatasetPostgresql
	err := ctx.ReadResource("azure:datafactory/datasetPostgresql:DatasetPostgresql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetPostgresql resources.
type datasetPostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset PostgreSQL.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetPostgresqlSchemaColumn `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset PostgreSQL.
	TableName *string `pulumi:"tableName"`
}

type DatasetPostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset PostgreSQL.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetPostgresqlSchemaColumnArrayInput
	// The table name of the Data Factory Dataset PostgreSQL.
	TableName pulumi.StringPtrInput
}

func (DatasetPostgresqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetPostgresqlState)(nil)).Elem()
}

type datasetPostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset PostgreSQL.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetPostgresqlSchemaColumn `pulumi:"schemaColumns"`
	// The table name of the Data Factory Dataset PostgreSQL.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a DatasetPostgresql resource.
type DatasetPostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset PostgreSQL.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetPostgresqlSchemaColumnArrayInput
	// The table name of the Data Factory Dataset PostgreSQL.
	TableName pulumi.StringPtrInput
}

func (DatasetPostgresqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetPostgresqlArgs)(nil)).Elem()
}

type DatasetPostgresqlInput interface {
	pulumi.Input

	ToDatasetPostgresqlOutput() DatasetPostgresqlOutput
	ToDatasetPostgresqlOutputWithContext(ctx context.Context) DatasetPostgresqlOutput
}

func (*DatasetPostgresql) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetPostgresql)(nil)).Elem()
}

func (i *DatasetPostgresql) ToDatasetPostgresqlOutput() DatasetPostgresqlOutput {
	return i.ToDatasetPostgresqlOutputWithContext(context.Background())
}

func (i *DatasetPostgresql) ToDatasetPostgresqlOutputWithContext(ctx context.Context) DatasetPostgresqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetPostgresqlOutput)
}

// DatasetPostgresqlArrayInput is an input type that accepts DatasetPostgresqlArray and DatasetPostgresqlArrayOutput values.
// You can construct a concrete instance of `DatasetPostgresqlArrayInput` via:
//
//          DatasetPostgresqlArray{ DatasetPostgresqlArgs{...} }
type DatasetPostgresqlArrayInput interface {
	pulumi.Input

	ToDatasetPostgresqlArrayOutput() DatasetPostgresqlArrayOutput
	ToDatasetPostgresqlArrayOutputWithContext(context.Context) DatasetPostgresqlArrayOutput
}

type DatasetPostgresqlArray []DatasetPostgresqlInput

func (DatasetPostgresqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetPostgresql)(nil)).Elem()
}

func (i DatasetPostgresqlArray) ToDatasetPostgresqlArrayOutput() DatasetPostgresqlArrayOutput {
	return i.ToDatasetPostgresqlArrayOutputWithContext(context.Background())
}

func (i DatasetPostgresqlArray) ToDatasetPostgresqlArrayOutputWithContext(ctx context.Context) DatasetPostgresqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetPostgresqlArrayOutput)
}

// DatasetPostgresqlMapInput is an input type that accepts DatasetPostgresqlMap and DatasetPostgresqlMapOutput values.
// You can construct a concrete instance of `DatasetPostgresqlMapInput` via:
//
//          DatasetPostgresqlMap{ "key": DatasetPostgresqlArgs{...} }
type DatasetPostgresqlMapInput interface {
	pulumi.Input

	ToDatasetPostgresqlMapOutput() DatasetPostgresqlMapOutput
	ToDatasetPostgresqlMapOutputWithContext(context.Context) DatasetPostgresqlMapOutput
}

type DatasetPostgresqlMap map[string]DatasetPostgresqlInput

func (DatasetPostgresqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetPostgresql)(nil)).Elem()
}

func (i DatasetPostgresqlMap) ToDatasetPostgresqlMapOutput() DatasetPostgresqlMapOutput {
	return i.ToDatasetPostgresqlMapOutputWithContext(context.Background())
}

func (i DatasetPostgresqlMap) ToDatasetPostgresqlMapOutputWithContext(ctx context.Context) DatasetPostgresqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetPostgresqlMapOutput)
}

type DatasetPostgresqlOutput struct{ *pulumi.OutputState }

func (DatasetPostgresqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetPostgresql)(nil)).Elem()
}

func (o DatasetPostgresqlOutput) ToDatasetPostgresqlOutput() DatasetPostgresqlOutput {
	return o
}

func (o DatasetPostgresqlOutput) ToDatasetPostgresqlOutputWithContext(ctx context.Context) DatasetPostgresqlOutput {
	return o
}

type DatasetPostgresqlArrayOutput struct{ *pulumi.OutputState }

func (DatasetPostgresqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetPostgresql)(nil)).Elem()
}

func (o DatasetPostgresqlArrayOutput) ToDatasetPostgresqlArrayOutput() DatasetPostgresqlArrayOutput {
	return o
}

func (o DatasetPostgresqlArrayOutput) ToDatasetPostgresqlArrayOutputWithContext(ctx context.Context) DatasetPostgresqlArrayOutput {
	return o
}

func (o DatasetPostgresqlArrayOutput) Index(i pulumi.IntInput) DatasetPostgresqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetPostgresql {
		return vs[0].([]*DatasetPostgresql)[vs[1].(int)]
	}).(DatasetPostgresqlOutput)
}

type DatasetPostgresqlMapOutput struct{ *pulumi.OutputState }

func (DatasetPostgresqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetPostgresql)(nil)).Elem()
}

func (o DatasetPostgresqlMapOutput) ToDatasetPostgresqlMapOutput() DatasetPostgresqlMapOutput {
	return o
}

func (o DatasetPostgresqlMapOutput) ToDatasetPostgresqlMapOutputWithContext(ctx context.Context) DatasetPostgresqlMapOutput {
	return o
}

func (o DatasetPostgresqlMapOutput) MapIndex(k pulumi.StringInput) DatasetPostgresqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetPostgresql {
		return vs[0].(map[string]*DatasetPostgresql)[vs[1].(string)]
	}).(DatasetPostgresqlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetPostgresqlInput)(nil)).Elem(), &DatasetPostgresql{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetPostgresqlArrayInput)(nil)).Elem(), DatasetPostgresqlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetPostgresqlMapInput)(nil)).Elem(), DatasetPostgresqlMap{})
	pulumi.RegisterOutputType(DatasetPostgresqlOutput{})
	pulumi.RegisterOutputType(DatasetPostgresqlArrayOutput{})
	pulumi.RegisterOutputType(DatasetPostgresqlMapOutput{})
}
