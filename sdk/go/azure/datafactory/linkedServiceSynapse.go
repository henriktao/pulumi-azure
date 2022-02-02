// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between Synapse and Azure Data Factory.
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
// 		_, err = datafactory.NewLinkedServiceSynapse(ctx, "exampleLinkedServiceSynapse", &datafactory.LinkedServiceSynapseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
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
// 		_, err = datafactory.NewLinkedServiceSynapse(ctx, "exampleLinkedServiceSynapse", &datafactory.LinkedServiceSynapseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;"),
// 			KeyVaultPassword: &datafactory.LinkedServiceSynapseKeyVaultPasswordArgs{
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
// Data Factory Synapse Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSynapse struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrOutput `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSynapse registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSynapse(ctx *pulumi.Context,
	name string, args *LinkedServiceSynapseArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSynapse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceSynapse
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSynapse gets an existing LinkedServiceSynapse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSynapse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSynapseState, opts ...pulumi.ResourceOption) (*LinkedServiceSynapse, error) {
	var resource LinkedServiceSynapse
	err := ctx.ReadResource("azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSynapse resources.
type linkedServiceSynapseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSynapseKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSynapseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSynapseState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSynapseState)(nil)).Elem()
}

type linkedServiceSynapseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the Synapse.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Synapse.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword *LinkedServiceSynapseKeyVaultPassword `pulumi:"keyVaultPassword"`
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSynapse resource.
type LinkedServiceSynapseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Synapse.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Synapse.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the Synapse.
	ConnectionString pulumi.StringInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service Synapse.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Synapse.
	IntegrationRuntimeName pulumi.StringPtrInput
	// A `keyVaultPassword` block as defined below. Use this argument to store Synapse password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword LinkedServiceSynapseKeyVaultPasswordPtrInput
	// Specifies the name of the Data Factory Linked Service Synapse. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Synapse.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Synapse. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSynapseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSynapseArgs)(nil)).Elem()
}

type LinkedServiceSynapseInput interface {
	pulumi.Input

	ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput
	ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput
}

func (*LinkedServiceSynapse) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSynapse)(nil)).Elem()
}

func (i *LinkedServiceSynapse) ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput {
	return i.ToLinkedServiceSynapseOutputWithContext(context.Background())
}

func (i *LinkedServiceSynapse) ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSynapseOutput)
}

// LinkedServiceSynapseArrayInput is an input type that accepts LinkedServiceSynapseArray and LinkedServiceSynapseArrayOutput values.
// You can construct a concrete instance of `LinkedServiceSynapseArrayInput` via:
//
//          LinkedServiceSynapseArray{ LinkedServiceSynapseArgs{...} }
type LinkedServiceSynapseArrayInput interface {
	pulumi.Input

	ToLinkedServiceSynapseArrayOutput() LinkedServiceSynapseArrayOutput
	ToLinkedServiceSynapseArrayOutputWithContext(context.Context) LinkedServiceSynapseArrayOutput
}

type LinkedServiceSynapseArray []LinkedServiceSynapseInput

func (LinkedServiceSynapseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceSynapse)(nil)).Elem()
}

func (i LinkedServiceSynapseArray) ToLinkedServiceSynapseArrayOutput() LinkedServiceSynapseArrayOutput {
	return i.ToLinkedServiceSynapseArrayOutputWithContext(context.Background())
}

func (i LinkedServiceSynapseArray) ToLinkedServiceSynapseArrayOutputWithContext(ctx context.Context) LinkedServiceSynapseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSynapseArrayOutput)
}

// LinkedServiceSynapseMapInput is an input type that accepts LinkedServiceSynapseMap and LinkedServiceSynapseMapOutput values.
// You can construct a concrete instance of `LinkedServiceSynapseMapInput` via:
//
//          LinkedServiceSynapseMap{ "key": LinkedServiceSynapseArgs{...} }
type LinkedServiceSynapseMapInput interface {
	pulumi.Input

	ToLinkedServiceSynapseMapOutput() LinkedServiceSynapseMapOutput
	ToLinkedServiceSynapseMapOutputWithContext(context.Context) LinkedServiceSynapseMapOutput
}

type LinkedServiceSynapseMap map[string]LinkedServiceSynapseInput

func (LinkedServiceSynapseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceSynapse)(nil)).Elem()
}

func (i LinkedServiceSynapseMap) ToLinkedServiceSynapseMapOutput() LinkedServiceSynapseMapOutput {
	return i.ToLinkedServiceSynapseMapOutputWithContext(context.Background())
}

func (i LinkedServiceSynapseMap) ToLinkedServiceSynapseMapOutputWithContext(ctx context.Context) LinkedServiceSynapseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSynapseMapOutput)
}

type LinkedServiceSynapseOutput struct{ *pulumi.OutputState }

func (LinkedServiceSynapseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSynapse)(nil)).Elem()
}

func (o LinkedServiceSynapseOutput) ToLinkedServiceSynapseOutput() LinkedServiceSynapseOutput {
	return o
}

func (o LinkedServiceSynapseOutput) ToLinkedServiceSynapseOutputWithContext(ctx context.Context) LinkedServiceSynapseOutput {
	return o
}

type LinkedServiceSynapseArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceSynapseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceSynapse)(nil)).Elem()
}

func (o LinkedServiceSynapseArrayOutput) ToLinkedServiceSynapseArrayOutput() LinkedServiceSynapseArrayOutput {
	return o
}

func (o LinkedServiceSynapseArrayOutput) ToLinkedServiceSynapseArrayOutputWithContext(ctx context.Context) LinkedServiceSynapseArrayOutput {
	return o
}

func (o LinkedServiceSynapseArrayOutput) Index(i pulumi.IntInput) LinkedServiceSynapseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LinkedServiceSynapse {
		return vs[0].([]*LinkedServiceSynapse)[vs[1].(int)]
	}).(LinkedServiceSynapseOutput)
}

type LinkedServiceSynapseMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceSynapseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceSynapse)(nil)).Elem()
}

func (o LinkedServiceSynapseMapOutput) ToLinkedServiceSynapseMapOutput() LinkedServiceSynapseMapOutput {
	return o
}

func (o LinkedServiceSynapseMapOutput) ToLinkedServiceSynapseMapOutputWithContext(ctx context.Context) LinkedServiceSynapseMapOutput {
	return o
}

func (o LinkedServiceSynapseMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceSynapseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LinkedServiceSynapse {
		return vs[0].(map[string]*LinkedServiceSynapse)[vs[1].(string)]
	}).(LinkedServiceSynapseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSynapseInput)(nil)).Elem(), &LinkedServiceSynapse{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSynapseArrayInput)(nil)).Elem(), LinkedServiceSynapseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceSynapseMapInput)(nil)).Elem(), LinkedServiceSynapseMap{})
	pulumi.RegisterOutputType(LinkedServiceSynapseOutput{})
	pulumi.RegisterOutputType(LinkedServiceSynapseArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceSynapseMapOutput{})
}
