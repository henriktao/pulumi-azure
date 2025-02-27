// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Parquet Dataset inside an Azure Data Factory.
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
// 		exampleLinkedServiceWeb, err := datafactory.NewLinkedServiceWeb(ctx, "exampleLinkedServiceWeb", &datafactory.LinkedServiceWebArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryId:      exampleFactory.ID(),
// 			AuthenticationType: pulumi.String("Anonymous"),
// 			Url:                pulumi.String("https://www.bing.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetParquet(ctx, "exampleDatasetParquet", &datafactory.DatasetParquetArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			LinkedServiceName: exampleLinkedServiceWeb.Name,
// 			HttpServerLocation: &datafactory.DatasetParquetHttpServerLocationArgs{
// 				RelativeUrl: pulumi.String("http://www.bing.com"),
// 				Path:        pulumi.String("foo/bar/"),
// 				Filename:    pulumi.String("fizz.txt"),
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
// Data Factory Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/datasetParquet:DatasetParquet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetParquet struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetParquetAzureBlobStorageLocationPtrOutput `pulumi:"azureBlobStorageLocation"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrOutput `pulumi:"compressionCodec"`
	CompressionLevel pulumi.StringPtrOutput `pulumi:"compressionLevel"`
	DataFactoryId    pulumi.StringOutput    `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetParquetHttpServerLocationPtrOutput `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetParquetSchemaColumnArrayOutput `pulumi:"schemaColumns"`
}

// NewDatasetParquet registers a new resource with the given unique name, arguments, and options.
func NewDatasetParquet(ctx *pulumi.Context,
	name string, args *DatasetParquetArgs, opts ...pulumi.ResourceOption) (*DatasetParquet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetParquet
	err := ctx.RegisterResource("azure:datafactory/datasetParquet:DatasetParquet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetParquet gets an existing DatasetParquet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetParquet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetParquetState, opts ...pulumi.ResourceOption) (*DatasetParquet, error) {
	var resource DatasetParquet
	err := ctx.ReadResource("azure:datafactory/datasetParquet:DatasetParquet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetParquet resources.
type datasetParquetState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetParquetAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec *string `pulumi:"compressionCodec"`
	CompressionLevel *string `pulumi:"compressionLevel"`
	DataFactoryId    *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetParquetHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetParquetSchemaColumn `pulumi:"schemaColumns"`
}

type DatasetParquetState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetParquetAzureBlobStorageLocationPtrInput
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrInput
	CompressionLevel pulumi.StringPtrInput
	DataFactoryId    pulumi.StringPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetParquetHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetParquetSchemaColumnArrayInput
}

func (DatasetParquetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetParquetState)(nil)).Elem()
}

type datasetParquetArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation *DatasetParquetAzureBlobStorageLocation `pulumi:"azureBlobStorageLocation"`
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec *string `pulumi:"compressionCodec"`
	CompressionLevel *string `pulumi:"compressionLevel"`
	DataFactoryId    *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `httpServerLocation` block as defined below.
	HttpServerLocation *DatasetParquetHttpServerLocation `pulumi:"httpServerLocation"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetParquetSchemaColumn `pulumi:"schemaColumns"`
}

// The set of arguments for constructing a DatasetParquet resource.
type DatasetParquetArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// A `azureBlobStorageLocation` block as defined below.
	AzureBlobStorageLocation DatasetParquetAzureBlobStorageLocationPtrInput
	// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
	CompressionCodec pulumi.StringPtrInput
	CompressionLevel pulumi.StringPtrInput
	DataFactoryId    pulumi.StringPtrInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `httpServerLocation` block as defined below.
	HttpServerLocation DatasetParquetHttpServerLocationPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetParquetSchemaColumnArrayInput
}

func (DatasetParquetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetParquetArgs)(nil)).Elem()
}

type DatasetParquetInput interface {
	pulumi.Input

	ToDatasetParquetOutput() DatasetParquetOutput
	ToDatasetParquetOutputWithContext(ctx context.Context) DatasetParquetOutput
}

func (*DatasetParquet) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetParquet)(nil)).Elem()
}

func (i *DatasetParquet) ToDatasetParquetOutput() DatasetParquetOutput {
	return i.ToDatasetParquetOutputWithContext(context.Background())
}

func (i *DatasetParquet) ToDatasetParquetOutputWithContext(ctx context.Context) DatasetParquetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetParquetOutput)
}

// DatasetParquetArrayInput is an input type that accepts DatasetParquetArray and DatasetParquetArrayOutput values.
// You can construct a concrete instance of `DatasetParquetArrayInput` via:
//
//          DatasetParquetArray{ DatasetParquetArgs{...} }
type DatasetParquetArrayInput interface {
	pulumi.Input

	ToDatasetParquetArrayOutput() DatasetParquetArrayOutput
	ToDatasetParquetArrayOutputWithContext(context.Context) DatasetParquetArrayOutput
}

type DatasetParquetArray []DatasetParquetInput

func (DatasetParquetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetParquet)(nil)).Elem()
}

func (i DatasetParquetArray) ToDatasetParquetArrayOutput() DatasetParquetArrayOutput {
	return i.ToDatasetParquetArrayOutputWithContext(context.Background())
}

func (i DatasetParquetArray) ToDatasetParquetArrayOutputWithContext(ctx context.Context) DatasetParquetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetParquetArrayOutput)
}

// DatasetParquetMapInput is an input type that accepts DatasetParquetMap and DatasetParquetMapOutput values.
// You can construct a concrete instance of `DatasetParquetMapInput` via:
//
//          DatasetParquetMap{ "key": DatasetParquetArgs{...} }
type DatasetParquetMapInput interface {
	pulumi.Input

	ToDatasetParquetMapOutput() DatasetParquetMapOutput
	ToDatasetParquetMapOutputWithContext(context.Context) DatasetParquetMapOutput
}

type DatasetParquetMap map[string]DatasetParquetInput

func (DatasetParquetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetParquet)(nil)).Elem()
}

func (i DatasetParquetMap) ToDatasetParquetMapOutput() DatasetParquetMapOutput {
	return i.ToDatasetParquetMapOutputWithContext(context.Background())
}

func (i DatasetParquetMap) ToDatasetParquetMapOutputWithContext(ctx context.Context) DatasetParquetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetParquetMapOutput)
}

type DatasetParquetOutput struct{ *pulumi.OutputState }

func (DatasetParquetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetParquet)(nil)).Elem()
}

func (o DatasetParquetOutput) ToDatasetParquetOutput() DatasetParquetOutput {
	return o
}

func (o DatasetParquetOutput) ToDatasetParquetOutputWithContext(ctx context.Context) DatasetParquetOutput {
	return o
}

type DatasetParquetArrayOutput struct{ *pulumi.OutputState }

func (DatasetParquetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatasetParquet)(nil)).Elem()
}

func (o DatasetParquetArrayOutput) ToDatasetParquetArrayOutput() DatasetParquetArrayOutput {
	return o
}

func (o DatasetParquetArrayOutput) ToDatasetParquetArrayOutputWithContext(ctx context.Context) DatasetParquetArrayOutput {
	return o
}

func (o DatasetParquetArrayOutput) Index(i pulumi.IntInput) DatasetParquetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatasetParquet {
		return vs[0].([]*DatasetParquet)[vs[1].(int)]
	}).(DatasetParquetOutput)
}

type DatasetParquetMapOutput struct{ *pulumi.OutputState }

func (DatasetParquetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatasetParquet)(nil)).Elem()
}

func (o DatasetParquetMapOutput) ToDatasetParquetMapOutput() DatasetParquetMapOutput {
	return o
}

func (o DatasetParquetMapOutput) ToDatasetParquetMapOutputWithContext(ctx context.Context) DatasetParquetMapOutput {
	return o
}

func (o DatasetParquetMapOutput) MapIndex(k pulumi.StringInput) DatasetParquetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatasetParquet {
		return vs[0].(map[string]*DatasetParquet)[vs[1].(string)]
	}).(DatasetParquetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetParquetInput)(nil)).Elem(), &DatasetParquet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetParquetArrayInput)(nil)).Elem(), DatasetParquetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetParquetMapInput)(nil)).Elem(), DatasetParquetMap{})
	pulumi.RegisterOutputType(DatasetParquetOutput{})
	pulumi.RegisterOutputType(DatasetParquetArrayOutput{})
	pulumi.RegisterOutputType(DatasetParquetMapOutput{})
}
