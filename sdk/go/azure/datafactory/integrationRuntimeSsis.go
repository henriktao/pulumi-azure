// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Factory Azure-SSIS Integration Runtime.
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
// 		_, err = datafactory.NewIntegrationRuntimeSsis(ctx, "exampleIntegrationRuntimeSsis", &datafactory.IntegrationRuntimeSsisArgs{
// 			DataFactoryId:     exampleFactory.ID(),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			NodeSize:          pulumi.String("Standard_D8_v3"),
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
// Data Factory Azure-SSIS Integration Runtimes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/integrationRuntimeSsis:IntegrationRuntimeSsis example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
// ```
type IntegrationRuntimeSsis struct {
	pulumi.CustomResourceState

	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeSsisCatalogInfoPtrOutput `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeSsisCustomSetupScriptPtrOutput `pulumi:"customSetupScript"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Azure-SSIS Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// An `expressCustomSetup` block as defined below.
	ExpressCustomSetup IntegrationRuntimeSsisExpressCustomSetupPtrOutput `pulumi:"expressCustomSetup"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrice`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrOutput `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The size of the nodes on which the Azure-SSIS Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringOutput `pulumi:"nodeSize"`
	// Number of nodes for the Azure-SSIS Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrOutput `pulumi:"numberOfNodes"`
	// One or more `packageStore` block as defined below.
	PackageStores IntegrationRuntimeSsisPackageStoreArrayOutput `pulumi:"packageStores"`
	// A `proxy` block as defined below.
	Proxy IntegrationRuntimeSsisProxyPtrOutput `pulumi:"proxy"`
	// The name of the resource group in which to create the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeSsisVnetIntegrationPtrOutput `pulumi:"vnetIntegration"`
}

// NewIntegrationRuntimeSsis registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntimeSsis(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeSsisArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntimeSsis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeSize == nil {
		return nil, errors.New("invalid value for required argument 'NodeSize'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationRuntimeSsis
	err := ctx.RegisterResource("azure:datafactory/integrationRuntimeSsis:IntegrationRuntimeSsis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntimeSsis gets an existing IntegrationRuntimeSsis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntimeSsis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeSsisState, opts ...pulumi.ResourceOption) (*IntegrationRuntimeSsis, error) {
	var resource IntegrationRuntimeSsis
	err := ctx.ReadResource("azure:datafactory/integrationRuntimeSsis:IntegrationRuntimeSsis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntimeSsis resources.
type integrationRuntimeSsisState struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo *IntegrationRuntimeSsisCatalogInfo `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript *IntegrationRuntimeSsisCustomSetupScript `pulumi:"customSetupScript"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The Azure-SSIS Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition *string `pulumi:"edition"`
	// An `expressCustomSetup` block as defined below.
	ExpressCustomSetup *IntegrationRuntimeSsisExpressCustomSetup `pulumi:"expressCustomSetup"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrice`. Defaults to `LicenseIncluded`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode *int `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The size of the nodes on which the Azure-SSIS Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize *string `pulumi:"nodeSize"`
	// Number of nodes for the Azure-SSIS Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// One or more `packageStore` block as defined below.
	PackageStores []IntegrationRuntimeSsisPackageStore `pulumi:"packageStores"`
	// A `proxy` block as defined below.
	Proxy *IntegrationRuntimeSsisProxy `pulumi:"proxy"`
	// The name of the resource group in which to create the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration *IntegrationRuntimeSsisVnetIntegration `pulumi:"vnetIntegration"`
}

type IntegrationRuntimeSsisState struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeSsisCatalogInfoPtrInput
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeSsisCustomSetupScriptPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The Azure-SSIS Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrInput
	// An `expressCustomSetup` block as defined below.
	ExpressCustomSetup IntegrationRuntimeSsisExpressCustomSetupPtrInput
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrice`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrInput
	// Specifies the name of the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The size of the nodes on which the Azure-SSIS Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringPtrInput
	// Number of nodes for the Azure-SSIS Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrInput
	// One or more `packageStore` block as defined below.
	PackageStores IntegrationRuntimeSsisPackageStoreArrayInput
	// A `proxy` block as defined below.
	Proxy IntegrationRuntimeSsisProxyPtrInput
	// The name of the resource group in which to create the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeSsisVnetIntegrationPtrInput
}

func (IntegrationRuntimeSsisState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeSsisState)(nil)).Elem()
}

type integrationRuntimeSsisArgs struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo *IntegrationRuntimeSsisCatalogInfo `pulumi:"catalogInfo"`
	// A `customSetupScript` block as defined below.
	CustomSetupScript *IntegrationRuntimeSsisCustomSetupScript `pulumi:"customSetupScript"`
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// The Azure-SSIS Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition *string `pulumi:"edition"`
	// An `expressCustomSetup` block as defined below.
	ExpressCustomSetup *IntegrationRuntimeSsisExpressCustomSetup `pulumi:"expressCustomSetup"`
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrice`. Defaults to `LicenseIncluded`.
	LicenseType *string `pulumi:"licenseType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode *int `pulumi:"maxParallelExecutionsPerNode"`
	// Specifies the name of the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The size of the nodes on which the Azure-SSIS Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize string `pulumi:"nodeSize"`
	// Number of nodes for the Azure-SSIS Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// One or more `packageStore` block as defined below.
	PackageStores []IntegrationRuntimeSsisPackageStore `pulumi:"packageStores"`
	// A `proxy` block as defined below.
	Proxy *IntegrationRuntimeSsisProxy `pulumi:"proxy"`
	// The name of the resource group in which to create the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `vnetIntegration` block as defined below.
	VnetIntegration *IntegrationRuntimeSsisVnetIntegration `pulumi:"vnetIntegration"`
}

// The set of arguments for constructing a IntegrationRuntimeSsis resource.
type IntegrationRuntimeSsisArgs struct {
	// A `catalogInfo` block as defined below.
	CatalogInfo IntegrationRuntimeSsisCatalogInfoPtrInput
	// A `customSetupScript` block as defined below.
	CustomSetupScript IntegrationRuntimeSsisCustomSetupScriptPtrInput
	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryId pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	//
	// Deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider
	DataFactoryName pulumi.StringPtrInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// The Azure-SSIS Integration Runtime edition. Valid values are `Standard` and `Enterprise`. Defaults to `Standard`.
	Edition pulumi.StringPtrInput
	// An `expressCustomSetup` block as defined below.
	ExpressCustomSetup IntegrationRuntimeSsisExpressCustomSetupPtrInput
	// The type of the license that is used. Valid values are `LicenseIncluded` and `BasePrice`. Defaults to `LicenseIncluded`.
	LicenseType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Defines the maximum parallel executions per node. Defaults to `1`. Max is `16`.
	MaxParallelExecutionsPerNode pulumi.IntPtrInput
	// Specifies the name of the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The size of the nodes on which the Azure-SSIS Integration Runtime runs. Valid values are: `Standard_D2_v3`, `Standard_D4_v3`, `Standard_D8_v3`, `Standard_D16_v3`, `Standard_D32_v3`, `Standard_D64_v3`, `Standard_E2_v3`, `Standard_E4_v3`, `Standard_E8_v3`, `Standard_E16_v3`, `Standard_E32_v3`, `Standard_E64_v3`, `Standard_D1_v2`, `Standard_D2_v2`, `Standard_D3_v2`, `Standard_D4_v2`, `Standard_A4_v2` and `Standard_A8_v2`
	NodeSize pulumi.StringInput
	// Number of nodes for the Azure-SSIS Integration Runtime. Max is `10`. Defaults to `1`.
	NumberOfNodes pulumi.IntPtrInput
	// One or more `packageStore` block as defined below.
	PackageStores IntegrationRuntimeSsisPackageStoreArrayInput
	// A `proxy` block as defined below.
	Proxy IntegrationRuntimeSsisProxyPtrInput
	// The name of the resource group in which to create the Azure-SSIS Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `vnetIntegration` block as defined below.
	VnetIntegration IntegrationRuntimeSsisVnetIntegrationPtrInput
}

func (IntegrationRuntimeSsisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeSsisArgs)(nil)).Elem()
}

type IntegrationRuntimeSsisInput interface {
	pulumi.Input

	ToIntegrationRuntimeSsisOutput() IntegrationRuntimeSsisOutput
	ToIntegrationRuntimeSsisOutputWithContext(ctx context.Context) IntegrationRuntimeSsisOutput
}

func (*IntegrationRuntimeSsis) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeSsis)(nil)).Elem()
}

func (i *IntegrationRuntimeSsis) ToIntegrationRuntimeSsisOutput() IntegrationRuntimeSsisOutput {
	return i.ToIntegrationRuntimeSsisOutputWithContext(context.Background())
}

func (i *IntegrationRuntimeSsis) ToIntegrationRuntimeSsisOutputWithContext(ctx context.Context) IntegrationRuntimeSsisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeSsisOutput)
}

// IntegrationRuntimeSsisArrayInput is an input type that accepts IntegrationRuntimeSsisArray and IntegrationRuntimeSsisArrayOutput values.
// You can construct a concrete instance of `IntegrationRuntimeSsisArrayInput` via:
//
//          IntegrationRuntimeSsisArray{ IntegrationRuntimeSsisArgs{...} }
type IntegrationRuntimeSsisArrayInput interface {
	pulumi.Input

	ToIntegrationRuntimeSsisArrayOutput() IntegrationRuntimeSsisArrayOutput
	ToIntegrationRuntimeSsisArrayOutputWithContext(context.Context) IntegrationRuntimeSsisArrayOutput
}

type IntegrationRuntimeSsisArray []IntegrationRuntimeSsisInput

func (IntegrationRuntimeSsisArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRuntimeSsis)(nil)).Elem()
}

func (i IntegrationRuntimeSsisArray) ToIntegrationRuntimeSsisArrayOutput() IntegrationRuntimeSsisArrayOutput {
	return i.ToIntegrationRuntimeSsisArrayOutputWithContext(context.Background())
}

func (i IntegrationRuntimeSsisArray) ToIntegrationRuntimeSsisArrayOutputWithContext(ctx context.Context) IntegrationRuntimeSsisArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeSsisArrayOutput)
}

// IntegrationRuntimeSsisMapInput is an input type that accepts IntegrationRuntimeSsisMap and IntegrationRuntimeSsisMapOutput values.
// You can construct a concrete instance of `IntegrationRuntimeSsisMapInput` via:
//
//          IntegrationRuntimeSsisMap{ "key": IntegrationRuntimeSsisArgs{...} }
type IntegrationRuntimeSsisMapInput interface {
	pulumi.Input

	ToIntegrationRuntimeSsisMapOutput() IntegrationRuntimeSsisMapOutput
	ToIntegrationRuntimeSsisMapOutputWithContext(context.Context) IntegrationRuntimeSsisMapOutput
}

type IntegrationRuntimeSsisMap map[string]IntegrationRuntimeSsisInput

func (IntegrationRuntimeSsisMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRuntimeSsis)(nil)).Elem()
}

func (i IntegrationRuntimeSsisMap) ToIntegrationRuntimeSsisMapOutput() IntegrationRuntimeSsisMapOutput {
	return i.ToIntegrationRuntimeSsisMapOutputWithContext(context.Background())
}

func (i IntegrationRuntimeSsisMap) ToIntegrationRuntimeSsisMapOutputWithContext(ctx context.Context) IntegrationRuntimeSsisMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeSsisMapOutput)
}

type IntegrationRuntimeSsisOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeSsisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeSsis)(nil)).Elem()
}

func (o IntegrationRuntimeSsisOutput) ToIntegrationRuntimeSsisOutput() IntegrationRuntimeSsisOutput {
	return o
}

func (o IntegrationRuntimeSsisOutput) ToIntegrationRuntimeSsisOutputWithContext(ctx context.Context) IntegrationRuntimeSsisOutput {
	return o
}

type IntegrationRuntimeSsisArrayOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeSsisArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationRuntimeSsis)(nil)).Elem()
}

func (o IntegrationRuntimeSsisArrayOutput) ToIntegrationRuntimeSsisArrayOutput() IntegrationRuntimeSsisArrayOutput {
	return o
}

func (o IntegrationRuntimeSsisArrayOutput) ToIntegrationRuntimeSsisArrayOutputWithContext(ctx context.Context) IntegrationRuntimeSsisArrayOutput {
	return o
}

func (o IntegrationRuntimeSsisArrayOutput) Index(i pulumi.IntInput) IntegrationRuntimeSsisOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationRuntimeSsis {
		return vs[0].([]*IntegrationRuntimeSsis)[vs[1].(int)]
	}).(IntegrationRuntimeSsisOutput)
}

type IntegrationRuntimeSsisMapOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeSsisMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationRuntimeSsis)(nil)).Elem()
}

func (o IntegrationRuntimeSsisMapOutput) ToIntegrationRuntimeSsisMapOutput() IntegrationRuntimeSsisMapOutput {
	return o
}

func (o IntegrationRuntimeSsisMapOutput) ToIntegrationRuntimeSsisMapOutputWithContext(ctx context.Context) IntegrationRuntimeSsisMapOutput {
	return o
}

func (o IntegrationRuntimeSsisMapOutput) MapIndex(k pulumi.StringInput) IntegrationRuntimeSsisOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationRuntimeSsis {
		return vs[0].(map[string]*IntegrationRuntimeSsis)[vs[1].(string)]
	}).(IntegrationRuntimeSsisOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeSsisInput)(nil)).Elem(), &IntegrationRuntimeSsis{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeSsisArrayInput)(nil)).Elem(), IntegrationRuntimeSsisArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeSsisMapInput)(nil)).Elem(), IntegrationRuntimeSsisMap{})
	pulumi.RegisterOutputType(IntegrationRuntimeSsisOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeSsisArrayOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeSsisMapOutput{})
}
