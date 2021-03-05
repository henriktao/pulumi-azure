// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure Blob Dataset inside an Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleLinkedServiceAzureBlobStorage, err := datafactory.NewLinkedServiceAzureBlobStorage(ctx, "exampleLinkedServiceAzureBlobStorage", &datafactory.LinkedServiceAzureBlobStorageArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString: exampleAccount.ApplyT(func(exampleAccount storage.LookupAccountResult) (string, error) {
// 				return exampleAccount.PrimaryConnectionString, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetAzureBlob(ctx, "exampleDatasetAzureBlob", &datafactory.DatasetAzureBlobArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			LinkedServiceName: exampleLinkedServiceAzureBlobStorage.Name,
// 			Path:              pulumi.String("foo"),
// 			Filename:          pulumi.String("bar.png"),
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
//  $ pulumi import azure:datafactory/datasetAzureBlob:DatasetAzureBlob example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetAzureBlob struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The filename of the Azure Blob.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The path of the Azure Blob.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetAzureBlobSchemaColumnArrayOutput `pulumi:"schemaColumns"`
}

// NewDatasetAzureBlob registers a new resource with the given unique name, arguments, and options.
func NewDatasetAzureBlob(ctx *pulumi.Context,
	name string, args *DatasetAzureBlobArgs, opts ...pulumi.ResourceOption) (*DatasetAzureBlob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.LinkedServiceName == nil {
		return nil, errors.New("invalid value for required argument 'LinkedServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DatasetAzureBlob
	err := ctx.RegisterResource("azure:datafactory/datasetAzureBlob:DatasetAzureBlob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetAzureBlob gets an existing DatasetAzureBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetAzureBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetAzureBlobState, opts ...pulumi.ResourceOption) (*DatasetAzureBlob, error) {
	var resource DatasetAzureBlob
	err := ctx.ReadResource("azure:datafactory/datasetAzureBlob:DatasetAzureBlob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetAzureBlob resources.
type datasetAzureBlobState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The filename of the Azure Blob.
	Filename *string `pulumi:"filename"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The path of the Azure Blob.
	Path *string `pulumi:"path"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetAzureBlobSchemaColumn `pulumi:"schemaColumns"`
}

type DatasetAzureBlobState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The filename of the Azure Blob.
	Filename pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The path of the Azure Blob.
	Path pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetAzureBlobSchemaColumnArrayInput
}

func (DatasetAzureBlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetAzureBlobState)(nil)).Elem()
}

type datasetAzureBlobArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The filename of the Azure Blob.
	Filename *string `pulumi:"filename"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The path of the Azure Blob.
	Path *string `pulumi:"path"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetAzureBlobSchemaColumn `pulumi:"schemaColumns"`
}

// The set of arguments for constructing a DatasetAzureBlob resource.
type DatasetAzureBlobArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The filename of the Azure Blob.
	Filename pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The path of the Azure Blob.
	Path pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetAzureBlobSchemaColumnArrayInput
}

func (DatasetAzureBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetAzureBlobArgs)(nil)).Elem()
}

type DatasetAzureBlobInput interface {
	pulumi.Input

	ToDatasetAzureBlobOutput() DatasetAzureBlobOutput
	ToDatasetAzureBlobOutputWithContext(ctx context.Context) DatasetAzureBlobOutput
}

func (*DatasetAzureBlob) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetAzureBlob)(nil))
}

func (i *DatasetAzureBlob) ToDatasetAzureBlobOutput() DatasetAzureBlobOutput {
	return i.ToDatasetAzureBlobOutputWithContext(context.Background())
}

func (i *DatasetAzureBlob) ToDatasetAzureBlobOutputWithContext(ctx context.Context) DatasetAzureBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAzureBlobOutput)
}

func (i *DatasetAzureBlob) ToDatasetAzureBlobPtrOutput() DatasetAzureBlobPtrOutput {
	return i.ToDatasetAzureBlobPtrOutputWithContext(context.Background())
}

func (i *DatasetAzureBlob) ToDatasetAzureBlobPtrOutputWithContext(ctx context.Context) DatasetAzureBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAzureBlobPtrOutput)
}

type DatasetAzureBlobPtrInput interface {
	pulumi.Input

	ToDatasetAzureBlobPtrOutput() DatasetAzureBlobPtrOutput
	ToDatasetAzureBlobPtrOutputWithContext(ctx context.Context) DatasetAzureBlobPtrOutput
}

type datasetAzureBlobPtrType DatasetAzureBlobArgs

func (*datasetAzureBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetAzureBlob)(nil))
}

func (i *datasetAzureBlobPtrType) ToDatasetAzureBlobPtrOutput() DatasetAzureBlobPtrOutput {
	return i.ToDatasetAzureBlobPtrOutputWithContext(context.Background())
}

func (i *datasetAzureBlobPtrType) ToDatasetAzureBlobPtrOutputWithContext(ctx context.Context) DatasetAzureBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAzureBlobPtrOutput)
}

// DatasetAzureBlobArrayInput is an input type that accepts DatasetAzureBlobArray and DatasetAzureBlobArrayOutput values.
// You can construct a concrete instance of `DatasetAzureBlobArrayInput` via:
//
//          DatasetAzureBlobArray{ DatasetAzureBlobArgs{...} }
type DatasetAzureBlobArrayInput interface {
	pulumi.Input

	ToDatasetAzureBlobArrayOutput() DatasetAzureBlobArrayOutput
	ToDatasetAzureBlobArrayOutputWithContext(context.Context) DatasetAzureBlobArrayOutput
}

type DatasetAzureBlobArray []DatasetAzureBlobInput

func (DatasetAzureBlobArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatasetAzureBlob)(nil))
}

func (i DatasetAzureBlobArray) ToDatasetAzureBlobArrayOutput() DatasetAzureBlobArrayOutput {
	return i.ToDatasetAzureBlobArrayOutputWithContext(context.Background())
}

func (i DatasetAzureBlobArray) ToDatasetAzureBlobArrayOutputWithContext(ctx context.Context) DatasetAzureBlobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAzureBlobArrayOutput)
}

// DatasetAzureBlobMapInput is an input type that accepts DatasetAzureBlobMap and DatasetAzureBlobMapOutput values.
// You can construct a concrete instance of `DatasetAzureBlobMapInput` via:
//
//          DatasetAzureBlobMap{ "key": DatasetAzureBlobArgs{...} }
type DatasetAzureBlobMapInput interface {
	pulumi.Input

	ToDatasetAzureBlobMapOutput() DatasetAzureBlobMapOutput
	ToDatasetAzureBlobMapOutputWithContext(context.Context) DatasetAzureBlobMapOutput
}

type DatasetAzureBlobMap map[string]DatasetAzureBlobInput

func (DatasetAzureBlobMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatasetAzureBlob)(nil))
}

func (i DatasetAzureBlobMap) ToDatasetAzureBlobMapOutput() DatasetAzureBlobMapOutput {
	return i.ToDatasetAzureBlobMapOutputWithContext(context.Background())
}

func (i DatasetAzureBlobMap) ToDatasetAzureBlobMapOutputWithContext(ctx context.Context) DatasetAzureBlobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAzureBlobMapOutput)
}

type DatasetAzureBlobOutput struct {
	*pulumi.OutputState
}

func (DatasetAzureBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetAzureBlob)(nil))
}

func (o DatasetAzureBlobOutput) ToDatasetAzureBlobOutput() DatasetAzureBlobOutput {
	return o
}

func (o DatasetAzureBlobOutput) ToDatasetAzureBlobOutputWithContext(ctx context.Context) DatasetAzureBlobOutput {
	return o
}

func (o DatasetAzureBlobOutput) ToDatasetAzureBlobPtrOutput() DatasetAzureBlobPtrOutput {
	return o.ToDatasetAzureBlobPtrOutputWithContext(context.Background())
}

func (o DatasetAzureBlobOutput) ToDatasetAzureBlobPtrOutputWithContext(ctx context.Context) DatasetAzureBlobPtrOutput {
	return o.ApplyT(func(v DatasetAzureBlob) *DatasetAzureBlob {
		return &v
	}).(DatasetAzureBlobPtrOutput)
}

type DatasetAzureBlobPtrOutput struct {
	*pulumi.OutputState
}

func (DatasetAzureBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetAzureBlob)(nil))
}

func (o DatasetAzureBlobPtrOutput) ToDatasetAzureBlobPtrOutput() DatasetAzureBlobPtrOutput {
	return o
}

func (o DatasetAzureBlobPtrOutput) ToDatasetAzureBlobPtrOutputWithContext(ctx context.Context) DatasetAzureBlobPtrOutput {
	return o
}

type DatasetAzureBlobArrayOutput struct{ *pulumi.OutputState }

func (DatasetAzureBlobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetAzureBlob)(nil))
}

func (o DatasetAzureBlobArrayOutput) ToDatasetAzureBlobArrayOutput() DatasetAzureBlobArrayOutput {
	return o
}

func (o DatasetAzureBlobArrayOutput) ToDatasetAzureBlobArrayOutputWithContext(ctx context.Context) DatasetAzureBlobArrayOutput {
	return o
}

func (o DatasetAzureBlobArrayOutput) Index(i pulumi.IntInput) DatasetAzureBlobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetAzureBlob {
		return vs[0].([]DatasetAzureBlob)[vs[1].(int)]
	}).(DatasetAzureBlobOutput)
}

type DatasetAzureBlobMapOutput struct{ *pulumi.OutputState }

func (DatasetAzureBlobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatasetAzureBlob)(nil))
}

func (o DatasetAzureBlobMapOutput) ToDatasetAzureBlobMapOutput() DatasetAzureBlobMapOutput {
	return o
}

func (o DatasetAzureBlobMapOutput) ToDatasetAzureBlobMapOutputWithContext(ctx context.Context) DatasetAzureBlobMapOutput {
	return o
}

func (o DatasetAzureBlobMapOutput) MapIndex(k pulumi.StringInput) DatasetAzureBlobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatasetAzureBlob {
		return vs[0].(map[string]DatasetAzureBlob)[vs[1].(string)]
	}).(DatasetAzureBlobOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetAzureBlobOutput{})
	pulumi.RegisterOutputType(DatasetAzureBlobPtrOutput{})
	pulumi.RegisterOutputType(DatasetAzureBlobArrayOutput{})
	pulumi.RegisterOutputType(DatasetAzureBlobMapOutput{})
}
