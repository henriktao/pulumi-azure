// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Dataset inside an Azure Data Factory. This is a generic resource that supports all different Dataset Types.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 			Identity: &datafactory.FactoryIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountKind:            pulumi.String("BlobStorage"),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedCustomService(ctx, "exampleLinkedCustomService", &datafactory.LinkedCustomServiceArgs{
// 			DataFactoryId:      exampleFactory.ID(),
// 			Type:               pulumi.String("AzureBlobStorage"),
// 			TypePropertiesJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"connectionString\":\"", azurerm_storage_account.Test.Primary_connection_string, "\"\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewCustomDataset(ctx, "exampleCustomDataset", &datafactory.CustomDatasetArgs{
// 			DataFactoryId: exampleFactory.ID(),
// 			Type:          pulumi.String("Json"),
// 			LinkedService: &datafactory.CustomDatasetLinkedServiceArgs{
// 				Name: pulumi.Any(azurerm_data_factory_linked_custom_service.Test.Name),
// 				Parameters: pulumi.StringMap{
// 					"key1": pulumi.String("value1"),
// 				},
// 			},
// 			TypePropertiesJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"location\": {\n", "    \"container\":\"", azurerm_storage_container.Test.Name, "\",\n", "    \"fileName\":\"foo.txt\",\n", "    \"folderPath\": \"foo/bar/\",\n", "    \"type\":\"AzureBlobStorageLocation\"\n", "  },\n", "  \"encodingName\":\"UTF-8\"\n", "}\n")),
// 			Description:        pulumi.String("test description"),
// 			Annotations: pulumi.StringArray{
// 				pulumi.String("test1"),
// 				pulumi.String("test2"),
// 				pulumi.String("test3"),
// 			},
// 			Folder: pulumi.String("testFolder"),
// 			Parameters: pulumi.StringMap{
// 				"foo": pulumi.String("test1"),
// 				"Bar": pulumi.String("Test2"),
// 			},
// 			AdditionalProperties: pulumi.StringMap{
// 				"foo": pulumi.String("test1"),
// 				"bar": pulumi.String("test2"),
// 			},
// 			SchemaJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"type\": \"object\",\n", "  \"properties\": {\n", "    \"name\": {\n", "      \"type\": \"object\",\n", "      \"properties\": {\n", "        \"firstName\": {\n", "          \"type\": \"string\"\n", "        },\n", "        \"lastName\": {\n", "          \"type\": \"string\"\n", "        }\n", "      }\n", "    },\n", "    \"age\": {\n", "      \"type\": \"integer\"\n", "    }\n", "  }\n", "}\n")),
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
//  $ pulumi import azure:datafactory/customDataset:CustomDataset example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type CustomDataset struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// A `linkedService` block as defined below.
	LinkedService CustomDatasetLinkedServiceOutput `pulumi:"linkedService"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// A JSON object that contains the schema of the Data Factory Dataset.
	SchemaJson pulumi.StringPtrOutput `pulumi:"schemaJson"`
	// The type of dataset that will be associated with Data Factory.
	Type pulumi.StringOutput `pulumi:"type"`
	// A JSON object that contains the properties of the Data Factory Dataset.
	TypePropertiesJson pulumi.StringOutput `pulumi:"typePropertiesJson"`
}

// NewCustomDataset registers a new resource with the given unique name, arguments, and options.
func NewCustomDataset(ctx *pulumi.Context,
	name string, args *CustomDatasetArgs, opts ...pulumi.ResourceOption) (*CustomDataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryId == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryId'")
	}
	if args.LinkedService == nil {
		return nil, errors.New("invalid value for required argument 'LinkedService'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.TypePropertiesJson == nil {
		return nil, errors.New("invalid value for required argument 'TypePropertiesJson'")
	}
	var resource CustomDataset
	err := ctx.RegisterResource("azure:datafactory/customDataset:CustomDataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDataset gets an existing CustomDataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDatasetState, opts ...pulumi.ResourceOption) (*CustomDataset, error) {
	var resource CustomDataset
	err := ctx.ReadResource("azure:datafactory/customDataset:CustomDataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDataset resources.
type customDatasetState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `linkedService` block as defined below.
	LinkedService *CustomDatasetLinkedService `pulumi:"linkedService"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// A JSON object that contains the schema of the Data Factory Dataset.
	SchemaJson *string `pulumi:"schemaJson"`
	// The type of dataset that will be associated with Data Factory.
	Type *string `pulumi:"type"`
	// A JSON object that contains the properties of the Data Factory Dataset.
	TypePropertiesJson *string `pulumi:"typePropertiesJson"`
}

type CustomDatasetState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `linkedService` block as defined below.
	LinkedService CustomDatasetLinkedServicePtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// A JSON object that contains the schema of the Data Factory Dataset.
	SchemaJson pulumi.StringPtrInput
	// The type of dataset that will be associated with Data Factory.
	Type pulumi.StringPtrInput
	// A JSON object that contains the properties of the Data Factory Dataset.
	TypePropertiesJson pulumi.StringPtrInput
}

func (CustomDatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDatasetState)(nil)).Elem()
}

type customDatasetArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryId string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// A `linkedService` block as defined below.
	LinkedService CustomDatasetLinkedService `pulumi:"linkedService"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// A JSON object that contains the schema of the Data Factory Dataset.
	SchemaJson *string `pulumi:"schemaJson"`
	// The type of dataset that will be associated with Data Factory.
	Type string `pulumi:"type"`
	// A JSON object that contains the properties of the Data Factory Dataset.
	TypePropertiesJson string `pulumi:"typePropertiesJson"`
}

// The set of arguments for constructing a CustomDataset resource.
type CustomDatasetArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryId pulumi.StringInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// A `linkedService` block as defined below.
	LinkedService CustomDatasetLinkedServiceInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// A JSON object that contains the schema of the Data Factory Dataset.
	SchemaJson pulumi.StringPtrInput
	// The type of dataset that will be associated with Data Factory.
	Type pulumi.StringInput
	// A JSON object that contains the properties of the Data Factory Dataset.
	TypePropertiesJson pulumi.StringInput
}

func (CustomDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDatasetArgs)(nil)).Elem()
}

type CustomDatasetInput interface {
	pulumi.Input

	ToCustomDatasetOutput() CustomDatasetOutput
	ToCustomDatasetOutputWithContext(ctx context.Context) CustomDatasetOutput
}

func (*CustomDataset) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDataset)(nil))
}

func (i *CustomDataset) ToCustomDatasetOutput() CustomDatasetOutput {
	return i.ToCustomDatasetOutputWithContext(context.Background())
}

func (i *CustomDataset) ToCustomDatasetOutputWithContext(ctx context.Context) CustomDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDatasetOutput)
}

func (i *CustomDataset) ToCustomDatasetPtrOutput() CustomDatasetPtrOutput {
	return i.ToCustomDatasetPtrOutputWithContext(context.Background())
}

func (i *CustomDataset) ToCustomDatasetPtrOutputWithContext(ctx context.Context) CustomDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDatasetPtrOutput)
}

type CustomDatasetPtrInput interface {
	pulumi.Input

	ToCustomDatasetPtrOutput() CustomDatasetPtrOutput
	ToCustomDatasetPtrOutputWithContext(ctx context.Context) CustomDatasetPtrOutput
}

type customDatasetPtrType CustomDatasetArgs

func (*customDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataset)(nil))
}

func (i *customDatasetPtrType) ToCustomDatasetPtrOutput() CustomDatasetPtrOutput {
	return i.ToCustomDatasetPtrOutputWithContext(context.Background())
}

func (i *customDatasetPtrType) ToCustomDatasetPtrOutputWithContext(ctx context.Context) CustomDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDatasetPtrOutput)
}

// CustomDatasetArrayInput is an input type that accepts CustomDatasetArray and CustomDatasetArrayOutput values.
// You can construct a concrete instance of `CustomDatasetArrayInput` via:
//
//          CustomDatasetArray{ CustomDatasetArgs{...} }
type CustomDatasetArrayInput interface {
	pulumi.Input

	ToCustomDatasetArrayOutput() CustomDatasetArrayOutput
	ToCustomDatasetArrayOutputWithContext(context.Context) CustomDatasetArrayOutput
}

type CustomDatasetArray []CustomDatasetInput

func (CustomDatasetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDataset)(nil)).Elem()
}

func (i CustomDatasetArray) ToCustomDatasetArrayOutput() CustomDatasetArrayOutput {
	return i.ToCustomDatasetArrayOutputWithContext(context.Background())
}

func (i CustomDatasetArray) ToCustomDatasetArrayOutputWithContext(ctx context.Context) CustomDatasetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDatasetArrayOutput)
}

// CustomDatasetMapInput is an input type that accepts CustomDatasetMap and CustomDatasetMapOutput values.
// You can construct a concrete instance of `CustomDatasetMapInput` via:
//
//          CustomDatasetMap{ "key": CustomDatasetArgs{...} }
type CustomDatasetMapInput interface {
	pulumi.Input

	ToCustomDatasetMapOutput() CustomDatasetMapOutput
	ToCustomDatasetMapOutputWithContext(context.Context) CustomDatasetMapOutput
}

type CustomDatasetMap map[string]CustomDatasetInput

func (CustomDatasetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDataset)(nil)).Elem()
}

func (i CustomDatasetMap) ToCustomDatasetMapOutput() CustomDatasetMapOutput {
	return i.ToCustomDatasetMapOutputWithContext(context.Background())
}

func (i CustomDatasetMap) ToCustomDatasetMapOutputWithContext(ctx context.Context) CustomDatasetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDatasetMapOutput)
}

type CustomDatasetOutput struct{ *pulumi.OutputState }

func (CustomDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDataset)(nil))
}

func (o CustomDatasetOutput) ToCustomDatasetOutput() CustomDatasetOutput {
	return o
}

func (o CustomDatasetOutput) ToCustomDatasetOutputWithContext(ctx context.Context) CustomDatasetOutput {
	return o
}

func (o CustomDatasetOutput) ToCustomDatasetPtrOutput() CustomDatasetPtrOutput {
	return o.ToCustomDatasetPtrOutputWithContext(context.Background())
}

func (o CustomDatasetOutput) ToCustomDatasetPtrOutputWithContext(ctx context.Context) CustomDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDataset) *CustomDataset {
		return &v
	}).(CustomDatasetPtrOutput)
}

type CustomDatasetPtrOutput struct{ *pulumi.OutputState }

func (CustomDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataset)(nil))
}

func (o CustomDatasetPtrOutput) ToCustomDatasetPtrOutput() CustomDatasetPtrOutput {
	return o
}

func (o CustomDatasetPtrOutput) ToCustomDatasetPtrOutputWithContext(ctx context.Context) CustomDatasetPtrOutput {
	return o
}

func (o CustomDatasetPtrOutput) Elem() CustomDatasetOutput {
	return o.ApplyT(func(v *CustomDataset) CustomDataset {
		if v != nil {
			return *v
		}
		var ret CustomDataset
		return ret
	}).(CustomDatasetOutput)
}

type CustomDatasetArrayOutput struct{ *pulumi.OutputState }

func (CustomDatasetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomDataset)(nil))
}

func (o CustomDatasetArrayOutput) ToCustomDatasetArrayOutput() CustomDatasetArrayOutput {
	return o
}

func (o CustomDatasetArrayOutput) ToCustomDatasetArrayOutputWithContext(ctx context.Context) CustomDatasetArrayOutput {
	return o
}

func (o CustomDatasetArrayOutput) Index(i pulumi.IntInput) CustomDatasetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomDataset {
		return vs[0].([]CustomDataset)[vs[1].(int)]
	}).(CustomDatasetOutput)
}

type CustomDatasetMapOutput struct{ *pulumi.OutputState }

func (CustomDatasetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomDataset)(nil))
}

func (o CustomDatasetMapOutput) ToCustomDatasetMapOutput() CustomDatasetMapOutput {
	return o
}

func (o CustomDatasetMapOutput) ToCustomDatasetMapOutputWithContext(ctx context.Context) CustomDatasetMapOutput {
	return o
}

func (o CustomDatasetMapOutput) MapIndex(k pulumi.StringInput) CustomDatasetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomDataset {
		return vs[0].(map[string]CustomDataset)[vs[1].(string)]
	}).(CustomDatasetOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDatasetOutput{})
	pulumi.RegisterOutputType(CustomDatasetPtrOutput{})
	pulumi.RegisterOutputType(CustomDatasetArrayOutput{})
	pulumi.RegisterOutputType(CustomDatasetMapOutput{})
}
