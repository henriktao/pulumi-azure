// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
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
// 		_, err = datafactory.NewLinkedServiceAzureFunction(ctx, "exampleLinkedServiceAzureFunction", &datafactory.LinkedServiceAzureFunctionArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			Url: exampleFunctionApp.ApplyT(func(exampleFunctionApp appservice.GetFunctionAppResult) (string, error) {
// 				return fmt.Sprintf("%v%v", "https://", exampleFunctionApp.DefaultHostname), nil
// 			}).(pulumi.StringOutput),
// 			Key: pulumi.String("foo"),
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
//  $ pulumi import azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceAzureFunction struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key pulumi.StringOutput `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewLinkedServiceAzureFunction registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceAzureFunction(ctx *pulumi.Context,
	name string, args *LinkedServiceAzureFunctionArgs, opts ...pulumi.ResourceOption) (*LinkedServiceAzureFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource LinkedServiceAzureFunction
	err := ctx.RegisterResource("azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceAzureFunction gets an existing LinkedServiceAzureFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceAzureFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceAzureFunctionState, opts ...pulumi.ResourceOption) (*LinkedServiceAzureFunction, error) {
	var resource LinkedServiceAzureFunction
	err := ctx.ReadResource("azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceAzureFunction resources.
type linkedServiceAzureFunctionState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key *string `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url *string `pulumi:"url"`
}

type LinkedServiceAzureFunctionState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The system key of the Azure Function.
	Key pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The url of the Azure Function.
	Url pulumi.StringPtrInput
}

func (LinkedServiceAzureFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureFunctionState)(nil)).Elem()
}

type linkedServiceAzureFunctionArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The system key of the Azure Function.
	Key string `pulumi:"key"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The url of the Azure Function.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a LinkedServiceAzureFunction resource.
type LinkedServiceAzureFunctionArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The system key of the Azure Function.
	Key pulumi.StringInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The url of the Azure Function.
	Url pulumi.StringInput
}

func (LinkedServiceAzureFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureFunctionArgs)(nil)).Elem()
}

type LinkedServiceAzureFunctionInput interface {
	pulumi.Input

	ToLinkedServiceAzureFunctionOutput() LinkedServiceAzureFunctionOutput
	ToLinkedServiceAzureFunctionOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionOutput
}

func (*LinkedServiceAzureFunction) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureFunction)(nil))
}

func (i *LinkedServiceAzureFunction) ToLinkedServiceAzureFunctionOutput() LinkedServiceAzureFunctionOutput {
	return i.ToLinkedServiceAzureFunctionOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureFunction) ToLinkedServiceAzureFunctionOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureFunctionOutput)
}

func (i *LinkedServiceAzureFunction) ToLinkedServiceAzureFunctionPtrOutput() LinkedServiceAzureFunctionPtrOutput {
	return i.ToLinkedServiceAzureFunctionPtrOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureFunction) ToLinkedServiceAzureFunctionPtrOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureFunctionPtrOutput)
}

type LinkedServiceAzureFunctionPtrInput interface {
	pulumi.Input

	ToLinkedServiceAzureFunctionPtrOutput() LinkedServiceAzureFunctionPtrOutput
	ToLinkedServiceAzureFunctionPtrOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionPtrOutput
}

type linkedServiceAzureFunctionPtrType LinkedServiceAzureFunctionArgs

func (*linkedServiceAzureFunctionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureFunction)(nil))
}

func (i *linkedServiceAzureFunctionPtrType) ToLinkedServiceAzureFunctionPtrOutput() LinkedServiceAzureFunctionPtrOutput {
	return i.ToLinkedServiceAzureFunctionPtrOutputWithContext(context.Background())
}

func (i *linkedServiceAzureFunctionPtrType) ToLinkedServiceAzureFunctionPtrOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureFunctionPtrOutput)
}

// LinkedServiceAzureFunctionArrayInput is an input type that accepts LinkedServiceAzureFunctionArray and LinkedServiceAzureFunctionArrayOutput values.
// You can construct a concrete instance of `LinkedServiceAzureFunctionArrayInput` via:
//
//          LinkedServiceAzureFunctionArray{ LinkedServiceAzureFunctionArgs{...} }
type LinkedServiceAzureFunctionArrayInput interface {
	pulumi.Input

	ToLinkedServiceAzureFunctionArrayOutput() LinkedServiceAzureFunctionArrayOutput
	ToLinkedServiceAzureFunctionArrayOutputWithContext(context.Context) LinkedServiceAzureFunctionArrayOutput
}

type LinkedServiceAzureFunctionArray []LinkedServiceAzureFunctionInput

func (LinkedServiceAzureFunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedServiceAzureFunction)(nil))
}

func (i LinkedServiceAzureFunctionArray) ToLinkedServiceAzureFunctionArrayOutput() LinkedServiceAzureFunctionArrayOutput {
	return i.ToLinkedServiceAzureFunctionArrayOutputWithContext(context.Background())
}

func (i LinkedServiceAzureFunctionArray) ToLinkedServiceAzureFunctionArrayOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureFunctionArrayOutput)
}

// LinkedServiceAzureFunctionMapInput is an input type that accepts LinkedServiceAzureFunctionMap and LinkedServiceAzureFunctionMapOutput values.
// You can construct a concrete instance of `LinkedServiceAzureFunctionMapInput` via:
//
//          LinkedServiceAzureFunctionMap{ "key": LinkedServiceAzureFunctionArgs{...} }
type LinkedServiceAzureFunctionMapInput interface {
	pulumi.Input

	ToLinkedServiceAzureFunctionMapOutput() LinkedServiceAzureFunctionMapOutput
	ToLinkedServiceAzureFunctionMapOutputWithContext(context.Context) LinkedServiceAzureFunctionMapOutput
}

type LinkedServiceAzureFunctionMap map[string]LinkedServiceAzureFunctionInput

func (LinkedServiceAzureFunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedServiceAzureFunction)(nil))
}

func (i LinkedServiceAzureFunctionMap) ToLinkedServiceAzureFunctionMapOutput() LinkedServiceAzureFunctionMapOutput {
	return i.ToLinkedServiceAzureFunctionMapOutputWithContext(context.Background())
}

func (i LinkedServiceAzureFunctionMap) ToLinkedServiceAzureFunctionMapOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureFunctionMapOutput)
}

type LinkedServiceAzureFunctionOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureFunction)(nil))
}

func (o LinkedServiceAzureFunctionOutput) ToLinkedServiceAzureFunctionOutput() LinkedServiceAzureFunctionOutput {
	return o
}

func (o LinkedServiceAzureFunctionOutput) ToLinkedServiceAzureFunctionOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionOutput {
	return o
}

func (o LinkedServiceAzureFunctionOutput) ToLinkedServiceAzureFunctionPtrOutput() LinkedServiceAzureFunctionPtrOutput {
	return o.ToLinkedServiceAzureFunctionPtrOutputWithContext(context.Background())
}

func (o LinkedServiceAzureFunctionOutput) ToLinkedServiceAzureFunctionPtrOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionPtrOutput {
	return o.ApplyT(func(v LinkedServiceAzureFunction) *LinkedServiceAzureFunction {
		return &v
	}).(LinkedServiceAzureFunctionPtrOutput)
}

type LinkedServiceAzureFunctionPtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureFunctionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureFunction)(nil))
}

func (o LinkedServiceAzureFunctionPtrOutput) ToLinkedServiceAzureFunctionPtrOutput() LinkedServiceAzureFunctionPtrOutput {
	return o
}

func (o LinkedServiceAzureFunctionPtrOutput) ToLinkedServiceAzureFunctionPtrOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionPtrOutput {
	return o
}

type LinkedServiceAzureFunctionArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureFunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceAzureFunction)(nil))
}

func (o LinkedServiceAzureFunctionArrayOutput) ToLinkedServiceAzureFunctionArrayOutput() LinkedServiceAzureFunctionArrayOutput {
	return o
}

func (o LinkedServiceAzureFunctionArrayOutput) ToLinkedServiceAzureFunctionArrayOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionArrayOutput {
	return o
}

func (o LinkedServiceAzureFunctionArrayOutput) Index(i pulumi.IntInput) LinkedServiceAzureFunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceAzureFunction {
		return vs[0].([]LinkedServiceAzureFunction)[vs[1].(int)]
	}).(LinkedServiceAzureFunctionOutput)
}

type LinkedServiceAzureFunctionMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureFunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceAzureFunction)(nil))
}

func (o LinkedServiceAzureFunctionMapOutput) ToLinkedServiceAzureFunctionMapOutput() LinkedServiceAzureFunctionMapOutput {
	return o
}

func (o LinkedServiceAzureFunctionMapOutput) ToLinkedServiceAzureFunctionMapOutputWithContext(ctx context.Context) LinkedServiceAzureFunctionMapOutput {
	return o
}

func (o LinkedServiceAzureFunctionMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceAzureFunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceAzureFunction {
		return vs[0].(map[string]LinkedServiceAzureFunction)[vs[1].(string)]
	}).(LinkedServiceAzureFunctionOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceAzureFunctionOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureFunctionPtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureFunctionArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureFunctionMapOutput{})
}
