// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between a SQL Server and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceSqlServer(ctx, "exampleLinkedServiceSqlServer", &datafactory.LinkedServiceSqlServerArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;Password=test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Password In Key Vault
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			SkuName:           pulumi.String("standard"),
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
// 		exampleLinkedServiceKeyVault, err := datafactory.NewLinkedServiceKeyVault(ctx, "exampleLinkedServiceKeyVault", &datafactory.LinkedServiceKeyVaultArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			KeyVaultId:        exampleKeyVault.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceSqlServer(ctx, "exampleLinkedServiceSqlServer", &datafactory.LinkedServiceSqlServerArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;"),
// 			KeyVaultPassword: &datafactory.LinkedServiceSqlServerKeyVaultPasswordArgs{
// 				LinkedServiceName: exampleLinkedServiceKeyVault.Name,
// 				SecretName:        pulumi.String("secret"),
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
// Data Factory SQL Server Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSqlServer struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// A `keyVaultConnectionString` block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	KeyVaultConnectionString LinkedServiceSqlServerKeyVaultConnectionStringPtrOutput `pulumi:"keyVaultConnectionString"`
	// A `keyVaultPassword` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSqlServerKeyVaultPasswordPtrOutput `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSqlServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, args *LinkedServiceSqlServerArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceSqlServer
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSqlServer gets an existing LinkedServiceSqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSqlServerState, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	var resource LinkedServiceSqlServer
	err := ctx.ReadResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSqlServer resources.
type linkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultConnectionString` block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	KeyVaultConnectionString *LinkedServiceSqlServerKeyVaultConnectionString `pulumi:"keyVaultConnectionString"`
	// A `keyVaultPassword` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSqlServerKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultConnectionString` block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	KeyVaultConnectionString LinkedServiceSqlServerKeyVaultConnectionStringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSqlServerKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerState)(nil)).Elem()
}

type linkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultConnectionString` block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	KeyVaultConnectionString *LinkedServiceSqlServerKeyVaultConnectionString `pulumi:"keyVaultConnectionString"`
	// A `keyVaultPassword` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSqlServerKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSqlServer resource.
type LinkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultConnectionString` block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either `connectionString` or `keyVaultConnectionString` is required.
	KeyVaultConnectionString LinkedServiceSqlServerKeyVaultConnectionStringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSqlServerKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerArgs)(nil)).Elem()
}

type LinkedServiceSqlServerInput interface {
	pulumi.Input

	ToLinkedServiceSqlServerOutput() LinkedServiceSqlServerOutput
	ToLinkedServiceSqlServerOutputWithContext(ctx context.Context) LinkedServiceSqlServerOutput
}

func (*LinkedServiceSqlServer) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSqlServer)(nil)).Elem()
}

func (i *LinkedServiceSqlServer) ToLinkedServiceSqlServerOutput() LinkedServiceSqlServerOutput {
	return i.ToLinkedServiceSqlServerOutputWithContext(context.Background())
}

func (i *LinkedServiceSqlServer) ToLinkedServiceSqlServerOutputWithContext(ctx context.Context) LinkedServiceSqlServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSqlServerOutput)
}

// LinkedServiceSqlServerArrayInput is an input type that accepts LinkedServiceSqlServerArray and LinkedServiceSqlServerArrayOutput values.
// You can construct a concrete instance of `LinkedServiceSqlServerArrayInput` via:
//
//          LinkedServiceSqlServerArray{ LinkedServiceSqlServerArgs{...} }
type LinkedServiceSqlServerArrayInput interface {
	pulumi.Input

	ToLinkedServiceSqlServerArrayOutput() LinkedServiceSqlServerArrayOutput
	ToLinkedServiceSqlServerArrayOutputWithContext(context.Context) LinkedServiceSqlServerArrayOutput
}

type LinkedServiceSqlServerArray []LinkedServiceSqlServerInput

func (LinkedServiceSqlServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceSqlServer)(nil)).Elem()
}

func (i LinkedServiceSqlServerArray) ToLinkedServiceSqlServerArrayOutput() LinkedServiceSqlServerArrayOutput {
	return i.ToLinkedServiceSqlServerArrayOutputWithContext(context.Background())
}

func (i LinkedServiceSqlServerArray) ToLinkedServiceSqlServerArrayOutputWithContext(ctx context.Context) LinkedServiceSqlServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSqlServerArrayOutput)
}

// LinkedServiceSqlServerMapInput is an input type that accepts LinkedServiceSqlServerMap and LinkedServiceSqlServerMapOutput values.
// You can construct a concrete instance of `LinkedServiceSqlServerMapInput` via:
//
//          LinkedServiceSqlServerMap{ "key": LinkedServiceSqlServerArgs{...} }
type LinkedServiceSqlServerMapInput interface {
	pulumi.Input

	ToLinkedServiceSqlServerMapOutput() LinkedServiceSqlServerMapOutput
	ToLinkedServiceSqlServerMapOutputWithContext(context.Context) LinkedServiceSqlServerMapOutput
}

type LinkedServiceSqlServerMap map[string]LinkedServiceSqlServerInput

func (LinkedServiceSqlServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceSqlServer)(nil)).Elem()
}

func (i LinkedServiceSqlServerMap) ToLinkedServiceSqlServerMapOutput() LinkedServiceSqlServerMapOutput {
	return i.ToLinkedServiceSqlServerMapOutputWithContext(context.Background())
}

func (i LinkedServiceSqlServerMap) ToLinkedServiceSqlServerMapOutputWithContext(ctx context.Context) LinkedServiceSqlServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSqlServerMapOutput)
}

type LinkedServiceSqlServerOutput struct{ *pulumi.OutputState }

func (LinkedServiceSqlServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSqlServer)(nil)).Elem()
}

func (o LinkedServiceSqlServerOutput) ToLinkedServiceSqlServerOutput() LinkedServiceSqlServerOutput {
	return o
}

func (o LinkedServiceSqlServerOutput) ToLinkedServiceSqlServerOutputWithContext(ctx context.Context) LinkedServiceSqlServerOutput {
	return o
}

type LinkedServiceSqlServerArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceSqlServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceSqlServer)(nil)).Elem()
}

func (o LinkedServiceSqlServerArrayOutput) ToLinkedServiceSqlServerArrayOutput() LinkedServiceSqlServerArrayOutput {
	return o
}

func (o LinkedServiceSqlServerArrayOutput) ToLinkedServiceSqlServerArrayOutputWithContext(ctx context.Context) LinkedServiceSqlServerArrayOutput {
	return o
}

func (o LinkedServiceSqlServerArrayOutput) Index(i pulumi.IntInput) LinkedServiceSqlServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LinkedServiceSqlServer {
		return vs[0].([]*LinkedServiceSqlServer)[vs[1].(int)]
	}).(LinkedServiceSqlServerOutput)
}

type LinkedServiceSqlServerMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceSqlServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceSqlServer)(nil)).Elem()
}

func (o LinkedServiceSqlServerMapOutput) ToLinkedServiceSqlServerMapOutput() LinkedServiceSqlServerMapOutput {
	return o
}

func (o LinkedServiceSqlServerMapOutput) ToLinkedServiceSqlServerMapOutputWithContext(ctx context.Context) LinkedServiceSqlServerMapOutput {
	return o
}

func (o LinkedServiceSqlServerMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceSqlServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LinkedServiceSqlServer {
		return vs[0].(map[string]*LinkedServiceSqlServer)[vs[1].(string)]
	}).(LinkedServiceSqlServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSqlServerInput)(nil)).Elem(), &LinkedServiceSqlServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSqlServerArrayInput)(nil)).Elem(), LinkedServiceSqlServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSqlServerMapInput)(nil)).Elem(), LinkedServiceSqlServerMap{})
	pulumi.RegisterOutputType(LinkedServiceSqlServerOutput{})
	pulumi.RegisterOutputType(LinkedServiceSqlServerArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceSqlServerMapOutput{})
}
