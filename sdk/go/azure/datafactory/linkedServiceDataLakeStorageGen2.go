// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between Data Lake Storage Gen2 and Azure Data Factory.
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
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceDataLakeStorageGen2(ctx, "exampleLinkedServiceDataLakeStorageGen2", &datafactory.LinkedServiceDataLakeStorageGen2Args{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			DataFactoryName:     exampleFactory.Name,
// 			ServicePrincipalId:  pulumi.String(current.ClientId),
// 			ServicePrincipalKey: pulumi.String("exampleKey"),
// 			Tenant:              pulumi.String("11111111-1111-1111-1111-111111111111"),
// 			Url:                 pulumi.String("https://datalakestoragegen2"),
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
// Data Factory Data Lake Storage Gen2 Linked Services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceDataLakeStorageGen2 struct {
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
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `storageAccountKey` and `useManagedIdentity`.
	ServicePrincipalId pulumi.StringPtrOutput `pulumi:"servicePrincipalId"`
	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey pulumi.StringPtrOutput `pulumi:"servicePrincipalKey"`
	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `useManagedIdentity`.
	StorageAccountKey pulumi.StringPtrOutput `pulumi:"storageAccountKey"`
	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant pulumi.StringPtrOutput `pulumi:"tenant"`
	// The endpoint for the Azure Data Lake Storage Gen2 service.
	Url pulumi.StringOutput `pulumi:"url"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `storageAccountKey`.
	UseManagedIdentity pulumi.BoolPtrOutput `pulumi:"useManagedIdentity"`
}

// NewLinkedServiceDataLakeStorageGen2 registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceDataLakeStorageGen2(ctx *pulumi.Context,
	name string, args *LinkedServiceDataLakeStorageGen2Args, opts ...pulumi.ResourceOption) (*LinkedServiceDataLakeStorageGen2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource LinkedServiceDataLakeStorageGen2
	err := ctx.RegisterResource("azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceDataLakeStorageGen2 gets an existing LinkedServiceDataLakeStorageGen2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceDataLakeStorageGen2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceDataLakeStorageGen2State, opts ...pulumi.ResourceOption) (*LinkedServiceDataLakeStorageGen2, error) {
	var resource LinkedServiceDataLakeStorageGen2
	err := ctx.ReadResource("azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceDataLakeStorageGen2 resources.
type linkedServiceDataLakeStorageGen2State struct {
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
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `storageAccountKey` and `useManagedIdentity`.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey *string `pulumi:"servicePrincipalKey"`
	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `useManagedIdentity`.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant *string `pulumi:"tenant"`
	// The endpoint for the Azure Data Lake Storage Gen2 service.
	Url *string `pulumi:"url"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `storageAccountKey`.
	UseManagedIdentity *bool `pulumi:"useManagedIdentity"`
}

type LinkedServiceDataLakeStorageGen2State struct {
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
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `storageAccountKey` and `useManagedIdentity`.
	ServicePrincipalId pulumi.StringPtrInput
	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey pulumi.StringPtrInput
	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `useManagedIdentity`.
	StorageAccountKey pulumi.StringPtrInput
	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant pulumi.StringPtrInput
	// The endpoint for the Azure Data Lake Storage Gen2 service.
	Url pulumi.StringPtrInput
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `storageAccountKey`.
	UseManagedIdentity pulumi.BoolPtrInput
}

func (LinkedServiceDataLakeStorageGen2State) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceDataLakeStorageGen2State)(nil)).Elem()
}

type linkedServiceDataLakeStorageGen2Args struct {
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
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `storageAccountKey` and `useManagedIdentity`.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey *string `pulumi:"servicePrincipalKey"`
	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `useManagedIdentity`.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant *string `pulumi:"tenant"`
	// The endpoint for the Azure Data Lake Storage Gen2 service.
	Url string `pulumi:"url"`
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `storageAccountKey`.
	UseManagedIdentity *bool `pulumi:"useManagedIdentity"`
}

// The set of arguments for constructing a LinkedServiceDataLakeStorageGen2 resource.
type LinkedServiceDataLakeStorageGen2Args struct {
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
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `storageAccountKey` and `useManagedIdentity`.
	ServicePrincipalId pulumi.StringPtrInput
	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey pulumi.StringPtrInput
	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account.  Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `useManagedIdentity`.
	StorageAccountKey pulumi.StringPtrInput
	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant pulumi.StringPtrInput
	// The endpoint for the Azure Data Lake Storage Gen2 service.
	Url pulumi.StringInput
	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with `servicePrincipalId`, `servicePrincipalKey`, `tenant` and `storageAccountKey`.
	UseManagedIdentity pulumi.BoolPtrInput
}

func (LinkedServiceDataLakeStorageGen2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceDataLakeStorageGen2Args)(nil)).Elem()
}

type LinkedServiceDataLakeStorageGen2Input interface {
	pulumi.Input

	ToLinkedServiceDataLakeStorageGen2Output() LinkedServiceDataLakeStorageGen2Output
	ToLinkedServiceDataLakeStorageGen2OutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2Output
}

func (*LinkedServiceDataLakeStorageGen2) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceDataLakeStorageGen2)(nil))
}

func (i *LinkedServiceDataLakeStorageGen2) ToLinkedServiceDataLakeStorageGen2Output() LinkedServiceDataLakeStorageGen2Output {
	return i.ToLinkedServiceDataLakeStorageGen2OutputWithContext(context.Background())
}

func (i *LinkedServiceDataLakeStorageGen2) ToLinkedServiceDataLakeStorageGen2OutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2Output {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceDataLakeStorageGen2Output)
}

func (i *LinkedServiceDataLakeStorageGen2) ToLinkedServiceDataLakeStorageGen2PtrOutput() LinkedServiceDataLakeStorageGen2PtrOutput {
	return i.ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(context.Background())
}

func (i *LinkedServiceDataLakeStorageGen2) ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceDataLakeStorageGen2PtrOutput)
}

type LinkedServiceDataLakeStorageGen2PtrInput interface {
	pulumi.Input

	ToLinkedServiceDataLakeStorageGen2PtrOutput() LinkedServiceDataLakeStorageGen2PtrOutput
	ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2PtrOutput
}

type linkedServiceDataLakeStorageGen2PtrType LinkedServiceDataLakeStorageGen2Args

func (*linkedServiceDataLakeStorageGen2PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceDataLakeStorageGen2)(nil))
}

func (i *linkedServiceDataLakeStorageGen2PtrType) ToLinkedServiceDataLakeStorageGen2PtrOutput() LinkedServiceDataLakeStorageGen2PtrOutput {
	return i.ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(context.Background())
}

func (i *linkedServiceDataLakeStorageGen2PtrType) ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceDataLakeStorageGen2PtrOutput)
}

// LinkedServiceDataLakeStorageGen2ArrayInput is an input type that accepts LinkedServiceDataLakeStorageGen2Array and LinkedServiceDataLakeStorageGen2ArrayOutput values.
// You can construct a concrete instance of `LinkedServiceDataLakeStorageGen2ArrayInput` via:
//
//          LinkedServiceDataLakeStorageGen2Array{ LinkedServiceDataLakeStorageGen2Args{...} }
type LinkedServiceDataLakeStorageGen2ArrayInput interface {
	pulumi.Input

	ToLinkedServiceDataLakeStorageGen2ArrayOutput() LinkedServiceDataLakeStorageGen2ArrayOutput
	ToLinkedServiceDataLakeStorageGen2ArrayOutputWithContext(context.Context) LinkedServiceDataLakeStorageGen2ArrayOutput
}

type LinkedServiceDataLakeStorageGen2Array []LinkedServiceDataLakeStorageGen2Input

func (LinkedServiceDataLakeStorageGen2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceDataLakeStorageGen2)(nil)).Elem()
}

func (i LinkedServiceDataLakeStorageGen2Array) ToLinkedServiceDataLakeStorageGen2ArrayOutput() LinkedServiceDataLakeStorageGen2ArrayOutput {
	return i.ToLinkedServiceDataLakeStorageGen2ArrayOutputWithContext(context.Background())
}

func (i LinkedServiceDataLakeStorageGen2Array) ToLinkedServiceDataLakeStorageGen2ArrayOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceDataLakeStorageGen2ArrayOutput)
}

// LinkedServiceDataLakeStorageGen2MapInput is an input type that accepts LinkedServiceDataLakeStorageGen2Map and LinkedServiceDataLakeStorageGen2MapOutput values.
// You can construct a concrete instance of `LinkedServiceDataLakeStorageGen2MapInput` via:
//
//          LinkedServiceDataLakeStorageGen2Map{ "key": LinkedServiceDataLakeStorageGen2Args{...} }
type LinkedServiceDataLakeStorageGen2MapInput interface {
	pulumi.Input

	ToLinkedServiceDataLakeStorageGen2MapOutput() LinkedServiceDataLakeStorageGen2MapOutput
	ToLinkedServiceDataLakeStorageGen2MapOutputWithContext(context.Context) LinkedServiceDataLakeStorageGen2MapOutput
}

type LinkedServiceDataLakeStorageGen2Map map[string]LinkedServiceDataLakeStorageGen2Input

func (LinkedServiceDataLakeStorageGen2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceDataLakeStorageGen2)(nil)).Elem()
}

func (i LinkedServiceDataLakeStorageGen2Map) ToLinkedServiceDataLakeStorageGen2MapOutput() LinkedServiceDataLakeStorageGen2MapOutput {
	return i.ToLinkedServiceDataLakeStorageGen2MapOutputWithContext(context.Background())
}

func (i LinkedServiceDataLakeStorageGen2Map) ToLinkedServiceDataLakeStorageGen2MapOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceDataLakeStorageGen2MapOutput)
}

type LinkedServiceDataLakeStorageGen2Output struct{ *pulumi.OutputState }

func (LinkedServiceDataLakeStorageGen2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceDataLakeStorageGen2)(nil))
}

func (o LinkedServiceDataLakeStorageGen2Output) ToLinkedServiceDataLakeStorageGen2Output() LinkedServiceDataLakeStorageGen2Output {
	return o
}

func (o LinkedServiceDataLakeStorageGen2Output) ToLinkedServiceDataLakeStorageGen2OutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2Output {
	return o
}

func (o LinkedServiceDataLakeStorageGen2Output) ToLinkedServiceDataLakeStorageGen2PtrOutput() LinkedServiceDataLakeStorageGen2PtrOutput {
	return o.ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(context.Background())
}

func (o LinkedServiceDataLakeStorageGen2Output) ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceDataLakeStorageGen2) *LinkedServiceDataLakeStorageGen2 {
		return &v
	}).(LinkedServiceDataLakeStorageGen2PtrOutput)
}

type LinkedServiceDataLakeStorageGen2PtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceDataLakeStorageGen2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceDataLakeStorageGen2)(nil))
}

func (o LinkedServiceDataLakeStorageGen2PtrOutput) ToLinkedServiceDataLakeStorageGen2PtrOutput() LinkedServiceDataLakeStorageGen2PtrOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2PtrOutput) ToLinkedServiceDataLakeStorageGen2PtrOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2PtrOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2PtrOutput) Elem() LinkedServiceDataLakeStorageGen2Output {
	return o.ApplyT(func(v *LinkedServiceDataLakeStorageGen2) LinkedServiceDataLakeStorageGen2 {
		if v != nil {
			return *v
		}
		var ret LinkedServiceDataLakeStorageGen2
		return ret
	}).(LinkedServiceDataLakeStorageGen2Output)
}

type LinkedServiceDataLakeStorageGen2ArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceDataLakeStorageGen2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceDataLakeStorageGen2)(nil))
}

func (o LinkedServiceDataLakeStorageGen2ArrayOutput) ToLinkedServiceDataLakeStorageGen2ArrayOutput() LinkedServiceDataLakeStorageGen2ArrayOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2ArrayOutput) ToLinkedServiceDataLakeStorageGen2ArrayOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2ArrayOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2ArrayOutput) Index(i pulumi.IntInput) LinkedServiceDataLakeStorageGen2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceDataLakeStorageGen2 {
		return vs[0].([]LinkedServiceDataLakeStorageGen2)[vs[1].(int)]
	}).(LinkedServiceDataLakeStorageGen2Output)
}

type LinkedServiceDataLakeStorageGen2MapOutput struct{ *pulumi.OutputState }

func (LinkedServiceDataLakeStorageGen2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceDataLakeStorageGen2)(nil))
}

func (o LinkedServiceDataLakeStorageGen2MapOutput) ToLinkedServiceDataLakeStorageGen2MapOutput() LinkedServiceDataLakeStorageGen2MapOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2MapOutput) ToLinkedServiceDataLakeStorageGen2MapOutputWithContext(ctx context.Context) LinkedServiceDataLakeStorageGen2MapOutput {
	return o
}

func (o LinkedServiceDataLakeStorageGen2MapOutput) MapIndex(k pulumi.StringInput) LinkedServiceDataLakeStorageGen2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceDataLakeStorageGen2 {
		return vs[0].(map[string]LinkedServiceDataLakeStorageGen2)[vs[1].(string)]
	}).(LinkedServiceDataLakeStorageGen2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceDataLakeStorageGen2Input)(nil)).Elem(), &LinkedServiceDataLakeStorageGen2{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceDataLakeStorageGen2PtrInput)(nil)).Elem(), &LinkedServiceDataLakeStorageGen2{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceDataLakeStorageGen2ArrayInput)(nil)).Elem(), LinkedServiceDataLakeStorageGen2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceDataLakeStorageGen2MapInput)(nil)).Elem(), LinkedServiceDataLakeStorageGen2Map{})
	pulumi.RegisterOutputType(LinkedServiceDataLakeStorageGen2Output{})
	pulumi.RegisterOutputType(LinkedServiceDataLakeStorageGen2PtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceDataLakeStorageGen2ArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceDataLakeStorageGen2MapOutput{})
}
