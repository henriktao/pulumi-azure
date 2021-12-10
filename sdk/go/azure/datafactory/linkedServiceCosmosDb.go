// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between a CosmosDB and Azure Data Factory using SQL API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
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
// 		exampleAccount, err := cosmosdb.LookupAccount(ctx, &cosmosdb.LookupAccountArgs{
// 			Name:              "tfex-cosmosdb-account",
// 			ResourceGroupName: "tfex-cosmosdb-account-rg",
// 		}, nil)
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
// 		_, err = datafactory.NewLinkedServiceCosmosDb(ctx, "exampleLinkedServiceCosmosDb", &datafactory.LinkedServiceCosmosDbArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			AccountEndpoint:   pulumi.Any(azurerm_cosmosdb_account.Example.Endpoint),
// 			AccountKey:        pulumi.String(exampleAccount.PrimaryKey),
// 			Database:          pulumi.String("foo"),
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
// Data Factory Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceCosmosDb struct {
	pulumi.CustomResourceState

	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrOutput `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrOutput `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceCosmosDb registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceCosmosDb(ctx *pulumi.Context,
	name string, args *LinkedServiceCosmosDbArgs, opts ...pulumi.ResourceOption) (*LinkedServiceCosmosDb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceCosmosDb
	err := ctx.RegisterResource("azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceCosmosDb gets an existing LinkedServiceCosmosDb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceCosmosDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceCosmosDbState, opts ...pulumi.ResourceOption) (*LinkedServiceCosmosDb, error) {
	var resource LinkedServiceCosmosDb
	err := ctx.ReadResource("azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceCosmosDb resources.
type linkedServiceCosmosDbState struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint *string `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey *string `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database *string `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceCosmosDbState struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrInput
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrInput
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceCosmosDbState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceCosmosDbState)(nil)).Elem()
}

type linkedServiceCosmosDbArgs struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint *string `pulumi:"accountEndpoint"`
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey *string `pulumi:"accountKey"`
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The name of the database. Required if `connectionString` is unspecified.
	Database *string `pulumi:"database"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceCosmosDb resource.
type LinkedServiceCosmosDbArgs struct {
	// The endpoint of the Azure CosmosDB account. Required if `connectionString` is unspecified.
	AccountEndpoint pulumi.StringPtrInput
	// The account key of the Azure Cosmos DB account. Required if `connectionString` is unspecified.
	AccountKey pulumi.StringPtrInput
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The connection string. Required if `accountEndpoint`, `accountKey`, and `database` are unspecified.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The name of the database. Required if `connectionString` is unspecified.
	Database pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceCosmosDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceCosmosDbArgs)(nil)).Elem()
}

type LinkedServiceCosmosDbInput interface {
	pulumi.Input

	ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput
	ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput
}

func (*LinkedServiceCosmosDb) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceCosmosDb)(nil))
}

func (i *LinkedServiceCosmosDb) ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput {
	return i.ToLinkedServiceCosmosDbOutputWithContext(context.Background())
}

func (i *LinkedServiceCosmosDb) ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbOutput)
}

func (i *LinkedServiceCosmosDb) ToLinkedServiceCosmosDbPtrOutput() LinkedServiceCosmosDbPtrOutput {
	return i.ToLinkedServiceCosmosDbPtrOutputWithContext(context.Background())
}

func (i *LinkedServiceCosmosDb) ToLinkedServiceCosmosDbPtrOutputWithContext(ctx context.Context) LinkedServiceCosmosDbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbPtrOutput)
}

type LinkedServiceCosmosDbPtrInput interface {
	pulumi.Input

	ToLinkedServiceCosmosDbPtrOutput() LinkedServiceCosmosDbPtrOutput
	ToLinkedServiceCosmosDbPtrOutputWithContext(ctx context.Context) LinkedServiceCosmosDbPtrOutput
}

type linkedServiceCosmosDbPtrType LinkedServiceCosmosDbArgs

func (*linkedServiceCosmosDbPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceCosmosDb)(nil))
}

func (i *linkedServiceCosmosDbPtrType) ToLinkedServiceCosmosDbPtrOutput() LinkedServiceCosmosDbPtrOutput {
	return i.ToLinkedServiceCosmosDbPtrOutputWithContext(context.Background())
}

func (i *linkedServiceCosmosDbPtrType) ToLinkedServiceCosmosDbPtrOutputWithContext(ctx context.Context) LinkedServiceCosmosDbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbPtrOutput)
}

// LinkedServiceCosmosDbArrayInput is an input type that accepts LinkedServiceCosmosDbArray and LinkedServiceCosmosDbArrayOutput values.
// You can construct a concrete instance of `LinkedServiceCosmosDbArrayInput` via:
//
//          LinkedServiceCosmosDbArray{ LinkedServiceCosmosDbArgs{...} }
type LinkedServiceCosmosDbArrayInput interface {
	pulumi.Input

	ToLinkedServiceCosmosDbArrayOutput() LinkedServiceCosmosDbArrayOutput
	ToLinkedServiceCosmosDbArrayOutputWithContext(context.Context) LinkedServiceCosmosDbArrayOutput
}

type LinkedServiceCosmosDbArray []LinkedServiceCosmosDbInput

func (LinkedServiceCosmosDbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceCosmosDb)(nil)).Elem()
}

func (i LinkedServiceCosmosDbArray) ToLinkedServiceCosmosDbArrayOutput() LinkedServiceCosmosDbArrayOutput {
	return i.ToLinkedServiceCosmosDbArrayOutputWithContext(context.Background())
}

func (i LinkedServiceCosmosDbArray) ToLinkedServiceCosmosDbArrayOutputWithContext(ctx context.Context) LinkedServiceCosmosDbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbArrayOutput)
}

// LinkedServiceCosmosDbMapInput is an input type that accepts LinkedServiceCosmosDbMap and LinkedServiceCosmosDbMapOutput values.
// You can construct a concrete instance of `LinkedServiceCosmosDbMapInput` via:
//
//          LinkedServiceCosmosDbMap{ "key": LinkedServiceCosmosDbArgs{...} }
type LinkedServiceCosmosDbMapInput interface {
	pulumi.Input

	ToLinkedServiceCosmosDbMapOutput() LinkedServiceCosmosDbMapOutput
	ToLinkedServiceCosmosDbMapOutputWithContext(context.Context) LinkedServiceCosmosDbMapOutput
}

type LinkedServiceCosmosDbMap map[string]LinkedServiceCosmosDbInput

func (LinkedServiceCosmosDbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceCosmosDb)(nil)).Elem()
}

func (i LinkedServiceCosmosDbMap) ToLinkedServiceCosmosDbMapOutput() LinkedServiceCosmosDbMapOutput {
	return i.ToLinkedServiceCosmosDbMapOutputWithContext(context.Background())
}

func (i LinkedServiceCosmosDbMap) ToLinkedServiceCosmosDbMapOutputWithContext(ctx context.Context) LinkedServiceCosmosDbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceCosmosDbMapOutput)
}

type LinkedServiceCosmosDbOutput struct{ *pulumi.OutputState }

func (LinkedServiceCosmosDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceCosmosDb)(nil))
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbOutput() LinkedServiceCosmosDbOutput {
	return o
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbOutputWithContext(ctx context.Context) LinkedServiceCosmosDbOutput {
	return o
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbPtrOutput() LinkedServiceCosmosDbPtrOutput {
	return o.ToLinkedServiceCosmosDbPtrOutputWithContext(context.Background())
}

func (o LinkedServiceCosmosDbOutput) ToLinkedServiceCosmosDbPtrOutputWithContext(ctx context.Context) LinkedServiceCosmosDbPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceCosmosDb) *LinkedServiceCosmosDb {
		return &v
	}).(LinkedServiceCosmosDbPtrOutput)
}

type LinkedServiceCosmosDbPtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceCosmosDbPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceCosmosDb)(nil))
}

func (o LinkedServiceCosmosDbPtrOutput) ToLinkedServiceCosmosDbPtrOutput() LinkedServiceCosmosDbPtrOutput {
	return o
}

func (o LinkedServiceCosmosDbPtrOutput) ToLinkedServiceCosmosDbPtrOutputWithContext(ctx context.Context) LinkedServiceCosmosDbPtrOutput {
	return o
}

func (o LinkedServiceCosmosDbPtrOutput) Elem() LinkedServiceCosmosDbOutput {
	return o.ApplyT(func(v *LinkedServiceCosmosDb) LinkedServiceCosmosDb {
		if v != nil {
			return *v
		}
		var ret LinkedServiceCosmosDb
		return ret
	}).(LinkedServiceCosmosDbOutput)
}

type LinkedServiceCosmosDbArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceCosmosDbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceCosmosDb)(nil))
}

func (o LinkedServiceCosmosDbArrayOutput) ToLinkedServiceCosmosDbArrayOutput() LinkedServiceCosmosDbArrayOutput {
	return o
}

func (o LinkedServiceCosmosDbArrayOutput) ToLinkedServiceCosmosDbArrayOutputWithContext(ctx context.Context) LinkedServiceCosmosDbArrayOutput {
	return o
}

func (o LinkedServiceCosmosDbArrayOutput) Index(i pulumi.IntInput) LinkedServiceCosmosDbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceCosmosDb {
		return vs[0].([]LinkedServiceCosmosDb)[vs[1].(int)]
	}).(LinkedServiceCosmosDbOutput)
}

type LinkedServiceCosmosDbMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceCosmosDbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceCosmosDb)(nil))
}

func (o LinkedServiceCosmosDbMapOutput) ToLinkedServiceCosmosDbMapOutput() LinkedServiceCosmosDbMapOutput {
	return o
}

func (o LinkedServiceCosmosDbMapOutput) ToLinkedServiceCosmosDbMapOutputWithContext(ctx context.Context) LinkedServiceCosmosDbMapOutput {
	return o
}

func (o LinkedServiceCosmosDbMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceCosmosDbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceCosmosDb {
		return vs[0].(map[string]LinkedServiceCosmosDb)[vs[1].(string)]
	}).(LinkedServiceCosmosDbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceCosmosDbInput)(nil)).Elem(), &LinkedServiceCosmosDb{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceCosmosDbPtrInput)(nil)).Elem(), &LinkedServiceCosmosDb{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceCosmosDbArrayInput)(nil)).Elem(), LinkedServiceCosmosDbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceCosmosDbMapInput)(nil)).Elem(), LinkedServiceCosmosDbMap{})
	pulumi.RegisterOutputType(LinkedServiceCosmosDbOutput{})
	pulumi.RegisterOutputType(LinkedServiceCosmosDbPtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceCosmosDbArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceCosmosDbMapOutput{})
}
