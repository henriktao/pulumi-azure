// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between a Database and Azure Data Factory through ODBC protocol.
//
// > **Note:** All arguments including the connectionString will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
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
// 		_, err = datafactory.NewLinkedServiceOdbc(ctx, "anonymous", &datafactory.LinkedServiceOdbcArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Driver={SQL Server};Server=test;Database=test;Uid=test;Pwd=test;"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceOdbc(ctx, "basicAuth", &datafactory.LinkedServiceOdbcArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryId:     exampleFactory.ID(),
// 			ConnectionString:  pulumi.String("Driver={SQL Server};Server=test;Database=test;Uid=test;Pwd=test;"),
// 			BasicAuthentication: &datafactory.LinkedServiceOdbcBasicAuthenticationArgs{
// 				Username: pulumi.String("onrylmz"),
// 				Password: pulumi.String("Ch4ngeM3!"),
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
// Data Factory ODBC Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceOdbc:LinkedServiceOdbc example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceOdbc struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// A `basicAuthentication` block as defined below.
	BasicAuthentication LinkedServiceOdbcBasicAuthenticationPtrOutput `pulumi:"basicAuthentication"`
	// The connection string in which to authenticate with ODBC.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The description for the Data Factory Linked Service ODBC.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceOdbc registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceOdbc(ctx *pulumi.Context,
	name string, args *LinkedServiceOdbcArgs, opts ...pulumi.ResourceOption) (*LinkedServiceOdbc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.DataFactoryId == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceOdbc
	err := ctx.RegisterResource("azure:datafactory/linkedServiceOdbc:LinkedServiceOdbc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceOdbc gets an existing LinkedServiceOdbc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceOdbc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceOdbcState, opts ...pulumi.ResourceOption) (*LinkedServiceOdbc, error) {
	var resource LinkedServiceOdbc
	err := ctx.ReadResource("azure:datafactory/linkedServiceOdbc:LinkedServiceOdbc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceOdbc resources.
type linkedServiceOdbcState struct {
	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations []string `pulumi:"annotations"`
	// A `basicAuthentication` block as defined below.
	BasicAuthentication *LinkedServiceOdbcBasicAuthentication `pulumi:"basicAuthentication"`
	// The connection string in which to authenticate with ODBC.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Linked Service ODBC.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceOdbcState struct {
	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations pulumi.StringArrayInput
	// A `basicAuthentication` block as defined below.
	BasicAuthentication LinkedServiceOdbcBasicAuthenticationPtrInput
	// The connection string in which to authenticate with ODBC.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The description for the Data Factory Linked Service ODBC.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceOdbcState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceOdbcState)(nil)).Elem()
}

type linkedServiceOdbcArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations []string `pulumi:"annotations"`
	// A `basicAuthentication` block as defined below.
	BasicAuthentication *LinkedServiceOdbcBasicAuthentication `pulumi:"basicAuthentication"`
	// The connection string in which to authenticate with ODBC.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId string `pulumi:"dataFactoryId"`
	// The description for the Data Factory Linked Service ODBC.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceOdbc resource.
type LinkedServiceOdbcArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations pulumi.StringArrayInput
	// A `basicAuthentication` block as defined below.
	BasicAuthentication LinkedServiceOdbcBasicAuthenticationPtrInput
	// The connection string in which to authenticate with ODBC.
	ConnectionString pulumi.StringInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringInput
	// The description for the Data Factory Linked Service ODBC.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceOdbcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceOdbcArgs)(nil)).Elem()
}

type LinkedServiceOdbcInput interface {
	pulumi.Input

	ToLinkedServiceOdbcOutput() LinkedServiceOdbcOutput
	ToLinkedServiceOdbcOutputWithContext(ctx context.Context) LinkedServiceOdbcOutput
}

func (*LinkedServiceOdbc) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceOdbc)(nil)).Elem()
}

func (i *LinkedServiceOdbc) ToLinkedServiceOdbcOutput() LinkedServiceOdbcOutput {
	return i.ToLinkedServiceOdbcOutputWithContext(context.Background())
}

func (i *LinkedServiceOdbc) ToLinkedServiceOdbcOutputWithContext(ctx context.Context) LinkedServiceOdbcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOdbcOutput)
}

// LinkedServiceOdbcArrayInput is an input type that accepts LinkedServiceOdbcArray and LinkedServiceOdbcArrayOutput values.
// You can construct a concrete instance of `LinkedServiceOdbcArrayInput` via:
//
//          LinkedServiceOdbcArray{ LinkedServiceOdbcArgs{...} }
type LinkedServiceOdbcArrayInput interface {
	pulumi.Input

	ToLinkedServiceOdbcArrayOutput() LinkedServiceOdbcArrayOutput
	ToLinkedServiceOdbcArrayOutputWithContext(context.Context) LinkedServiceOdbcArrayOutput
}

type LinkedServiceOdbcArray []LinkedServiceOdbcInput

func (LinkedServiceOdbcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceOdbc)(nil)).Elem()
}

func (i LinkedServiceOdbcArray) ToLinkedServiceOdbcArrayOutput() LinkedServiceOdbcArrayOutput {
	return i.ToLinkedServiceOdbcArrayOutputWithContext(context.Background())
}

func (i LinkedServiceOdbcArray) ToLinkedServiceOdbcArrayOutputWithContext(ctx context.Context) LinkedServiceOdbcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOdbcArrayOutput)
}

// LinkedServiceOdbcMapInput is an input type that accepts LinkedServiceOdbcMap and LinkedServiceOdbcMapOutput values.
// You can construct a concrete instance of `LinkedServiceOdbcMapInput` via:
//
//          LinkedServiceOdbcMap{ "key": LinkedServiceOdbcArgs{...} }
type LinkedServiceOdbcMapInput interface {
	pulumi.Input

	ToLinkedServiceOdbcMapOutput() LinkedServiceOdbcMapOutput
	ToLinkedServiceOdbcMapOutputWithContext(context.Context) LinkedServiceOdbcMapOutput
}

type LinkedServiceOdbcMap map[string]LinkedServiceOdbcInput

func (LinkedServiceOdbcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceOdbc)(nil)).Elem()
}

func (i LinkedServiceOdbcMap) ToLinkedServiceOdbcMapOutput() LinkedServiceOdbcMapOutput {
	return i.ToLinkedServiceOdbcMapOutputWithContext(context.Background())
}

func (i LinkedServiceOdbcMap) ToLinkedServiceOdbcMapOutputWithContext(ctx context.Context) LinkedServiceOdbcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOdbcMapOutput)
}

type LinkedServiceOdbcOutput struct{ *pulumi.OutputState }

func (LinkedServiceOdbcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceOdbc)(nil)).Elem()
}

func (o LinkedServiceOdbcOutput) ToLinkedServiceOdbcOutput() LinkedServiceOdbcOutput {
	return o
}

func (o LinkedServiceOdbcOutput) ToLinkedServiceOdbcOutputWithContext(ctx context.Context) LinkedServiceOdbcOutput {
	return o
}

type LinkedServiceOdbcArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceOdbcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceOdbc)(nil)).Elem()
}

func (o LinkedServiceOdbcArrayOutput) ToLinkedServiceOdbcArrayOutput() LinkedServiceOdbcArrayOutput {
	return o
}

func (o LinkedServiceOdbcArrayOutput) ToLinkedServiceOdbcArrayOutputWithContext(ctx context.Context) LinkedServiceOdbcArrayOutput {
	return o
}

func (o LinkedServiceOdbcArrayOutput) Index(i pulumi.IntInput) LinkedServiceOdbcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LinkedServiceOdbc {
		return vs[0].([]*LinkedServiceOdbc)[vs[1].(int)]
	}).(LinkedServiceOdbcOutput)
}

type LinkedServiceOdbcMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceOdbcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceOdbc)(nil)).Elem()
}

func (o LinkedServiceOdbcMapOutput) ToLinkedServiceOdbcMapOutput() LinkedServiceOdbcMapOutput {
	return o
}

func (o LinkedServiceOdbcMapOutput) ToLinkedServiceOdbcMapOutputWithContext(ctx context.Context) LinkedServiceOdbcMapOutput {
	return o
}

func (o LinkedServiceOdbcMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceOdbcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LinkedServiceOdbc {
		return vs[0].(map[string]*LinkedServiceOdbc)[vs[1].(string)]
	}).(LinkedServiceOdbcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceOdbcInput)(nil)).Elem(), &LinkedServiceOdbc{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceOdbcArrayInput)(nil)).Elem(), LinkedServiceOdbcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceOdbcMapInput)(nil)).Elem(), LinkedServiceOdbcMap{})
	pulumi.RegisterOutputType(LinkedServiceOdbcOutput{})
	pulumi.RegisterOutputType(LinkedServiceOdbcArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceOdbcMapOutput{})
}
