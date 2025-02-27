// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Snowflake Dataset inside an Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceSnowflake(ctx, "exampleLinkedServiceSnowflake", &datafactory.LinkedServiceSnowflakeArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("jdbc:snowflake://account.region.snowflakecomputing.com/?user=user&db=db&warehouse=wh"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetSnowflake(ctx, "exampleDatasetSnowflake", &datafactory.DatasetSnowflakeArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 			DataFactoryId:     pulumi.Any(azurerm_data_factory.Test.Id),
// 			LinkedServiceName: pulumi.Any(azurerm_data_factory_linked_service_snowflake.Test.Name),
// 			SchemaName:        pulumi.String("foo_schema"),
// 			TableName:         pulumi.String("foo_table"),
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
// Data Factory Snowflake Datasets can be imported using the `resource id`,
//
// e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/datasetSnowflake:DatasetSnowflake example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetSnowflake struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset Snowflake.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset Snowflake.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset Snowflake.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset Snowflake.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSnowflakeSchemaColumnArrayOutput `pulumi:"schemaColumns"`
	// The schema name of the Data Factory Dataset Snowflake.
	SchemaName pulumi.StringPtrOutput `pulumi:"schemaName"`
	// Deprecated: This block has been deprecated in favour of `schema_column` and will be removed.
	StructureColumns DatasetSnowflakeStructureColumnArrayOutput `pulumi:"structureColumns"`
	// The table name of the Data Factory Dataset Snowflake.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewDatasetSnowflake registers a new resource with the given unique name, arguments, and options.
func NewDatasetSnowflake(ctx *pulumi.Context,
	name string, args *DatasetSnowflakeArgs, opts ...pulumi.ResourceOption) (*DatasetSnowflake, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetSnowflake
	err := ctx.RegisterResource("azure:datafactory/datasetSnowflake:DatasetSnowflake", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetSnowflake gets an existing DatasetSnowflake resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetSnowflake(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetSnowflakeState, opts ...pulumi.ResourceOption) (*DatasetSnowflake, error) {
	var resource DatasetSnowflake
	err := ctx.ReadResource("azure:datafactory/datasetSnowflake:DatasetSnowflake", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetSnowflake resources.
type datasetSnowflakeState struct {
	// A map of additional properties to associate with the Data Factory Dataset Snowflake.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset Snowflake.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset Snowflake.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset Snowflake.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetSnowflakeSchemaColumn `pulumi:"schemaColumns"`
	// The schema name of the Data Factory Dataset Snowflake.
	SchemaName *string `pulumi:"schemaName"`
	// Deprecated: This block has been deprecated in favour of `schema_column` and will be removed.
	StructureColumns []DatasetSnowflakeStructureColumn `pulumi:"structureColumns"`
	// The table name of the Data Factory Dataset Snowflake.
	TableName *string `pulumi:"tableName"`
}

type DatasetSnowflakeState struct {
	// A map of additional properties to associate with the Data Factory Dataset Snowflake.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset Snowflake.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset Snowflake.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset Snowflake.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSnowflakeSchemaColumnArrayInput
	// The schema name of the Data Factory Dataset Snowflake.
	SchemaName pulumi.StringPtrInput
	// Deprecated: This block has been deprecated in favour of `schema_column` and will be removed.
	StructureColumns DatasetSnowflakeStructureColumnArrayInput
	// The table name of the Data Factory Dataset Snowflake.
	TableName pulumi.StringPtrInput
}

func (DatasetSnowflakeState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetSnowflakeState)(nil)).Elem()
}

type datasetSnowflakeArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset Snowflake.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset Snowflake.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset Snowflake.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset Snowflake.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetSnowflakeSchemaColumn `pulumi:"schemaColumns"`
	// The schema name of the Data Factory Dataset Snowflake.
	SchemaName *string `pulumi:"schemaName"`
	// Deprecated: This block has been deprecated in favour of `schema_column` and will be removed.
	StructureColumns []DatasetSnowflakeStructureColumn `pulumi:"structureColumns"`
	// The table name of the Data Factory Dataset Snowflake.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a DatasetSnowflake resource.
type DatasetSnowflakeArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset Snowflake.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset Snowflake.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset Snowflake.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset Snowflake.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetSnowflakeSchemaColumnArrayInput
	// The schema name of the Data Factory Dataset Snowflake.
	SchemaName pulumi.StringPtrInput
	// Deprecated: This block has been deprecated in favour of `schema_column` and will be removed.
	StructureColumns DatasetSnowflakeStructureColumnArrayInput
	// The table name of the Data Factory Dataset Snowflake.
	TableName pulumi.StringPtrInput
}

func (DatasetSnowflakeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetSnowflakeArgs)(nil)).Elem()
}

type DatasetSnowflakeInput interface {
	pulumi.Input

	ToDatasetSnowflakeOutput() DatasetSnowflakeOutput
	ToDatasetSnowflakeOutputWithContext(ctx context.Context) DatasetSnowflakeOutput
}

func (*DatasetSnowflake) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetSnowflake)(nil)).Elem()
}

func (i *DatasetSnowflake) ToDatasetSnowflakeOutput() DatasetSnowflakeOutput {
	return i.ToDatasetSnowflakeOutputWithContext(context.Background())
}

func (i *DatasetSnowflake) ToDatasetSnowflakeOutputWithContext(ctx context.Context) DatasetSnowflakeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetSnowflakeOutput)
}

// DatasetSnowflakeArrayInput is an input type that accepts DatasetSnowflakeArray and DatasetSnowflakeArrayOutput values.
// You can construct a concrete instance of `DatasetSnowflakeArrayInput` via:
//
//          DatasetSnowflakeArray{ DatasetSnowflakeArgs{...} }
type DatasetSnowflakeArrayInput interface {
	pulumi.Input

	ToDatasetSnowflakeArrayOutput() DatasetSnowflakeArrayOutput
	ToDatasetSnowflakeArrayOutputWithContext(context.Context) DatasetSnowflakeArrayOutput
}

type DatasetSnowflakeArray []DatasetSnowflakeInput

func (DatasetSnowflakeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetSnowflake)(nil)).Elem()
}

func (i DatasetSnowflakeArray) ToDatasetSnowflakeArrayOutput() DatasetSnowflakeArrayOutput {
	return i.ToDatasetSnowflakeArrayOutputWithContext(context.Background())
}

func (i DatasetSnowflakeArray) ToDatasetSnowflakeArrayOutputWithContext(ctx context.Context) DatasetSnowflakeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetSnowflakeArrayOutput)
}

// DatasetSnowflakeMapInput is an input type that accepts DatasetSnowflakeMap and DatasetSnowflakeMapOutput values.
// You can construct a concrete instance of `DatasetSnowflakeMapInput` via:
//
//          DatasetSnowflakeMap{ "key": DatasetSnowflakeArgs{...} }
type DatasetSnowflakeMapInput interface {
	pulumi.Input

	ToDatasetSnowflakeMapOutput() DatasetSnowflakeMapOutput
	ToDatasetSnowflakeMapOutputWithContext(context.Context) DatasetSnowflakeMapOutput
}

type DatasetSnowflakeMap map[string]DatasetSnowflakeInput

func (DatasetSnowflakeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetSnowflake)(nil)).Elem()
}

func (i DatasetSnowflakeMap) ToDatasetSnowflakeMapOutput() DatasetSnowflakeMapOutput {
	return i.ToDatasetSnowflakeMapOutputWithContext(context.Background())
}

func (i DatasetSnowflakeMap) ToDatasetSnowflakeMapOutputWithContext(ctx context.Context) DatasetSnowflakeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetSnowflakeMapOutput)
}

type DatasetSnowflakeOutput struct{ *pulumi.OutputState }

func (DatasetSnowflakeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetSnowflake)(nil)).Elem()
}

func (o DatasetSnowflakeOutput) ToDatasetSnowflakeOutput() DatasetSnowflakeOutput {
	return o
}

func (o DatasetSnowflakeOutput) ToDatasetSnowflakeOutputWithContext(ctx context.Context) DatasetSnowflakeOutput {
	return o
}

type DatasetSnowflakeArrayOutput struct{ *pulumi.OutputState }

func (DatasetSnowflakeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetSnowflake)(nil)).Elem()
}

func (o DatasetSnowflakeArrayOutput) ToDatasetSnowflakeArrayOutput() DatasetSnowflakeArrayOutput {
	return o
}

func (o DatasetSnowflakeArrayOutput) ToDatasetSnowflakeArrayOutputWithContext(ctx context.Context) DatasetSnowflakeArrayOutput {
	return o
}

func (o DatasetSnowflakeArrayOutput) Index(i pulumi.IntInput) DatasetSnowflakeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetSnowflake {
		return vs[0].([]*DatasetSnowflake)[vs[1].(int)]
	}).(DatasetSnowflakeOutput)
}

type DatasetSnowflakeMapOutput struct{ *pulumi.OutputState }

func (DatasetSnowflakeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetSnowflake)(nil)).Elem()
}

func (o DatasetSnowflakeMapOutput) ToDatasetSnowflakeMapOutput() DatasetSnowflakeMapOutput {
	return o
}

func (o DatasetSnowflakeMapOutput) ToDatasetSnowflakeMapOutputWithContext(ctx context.Context) DatasetSnowflakeMapOutput {
	return o
}

func (o DatasetSnowflakeMapOutput) MapIndex(k pulumi.StringInput) DatasetSnowflakeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetSnowflake {
		return vs[0].(map[string]*DatasetSnowflake)[vs[1].(string)]
	}).(DatasetSnowflakeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetSnowflakeInput)(nil)).Elem(), &DatasetSnowflake{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetSnowflakeArrayInput)(nil)).Elem(), DatasetSnowflakeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetSnowflakeMapInput)(nil)).Elem(), DatasetSnowflakeMap{})
	pulumi.RegisterOutputType(DatasetSnowflakeOutput{})
	pulumi.RegisterOutputType(DatasetSnowflakeArrayOutput{})
	pulumi.RegisterOutputType(DatasetSnowflakeMapOutput{})
}
