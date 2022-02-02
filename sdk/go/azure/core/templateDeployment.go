// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a template deployment of resources
//
// > **Note on ARM Template Deployments:** Due to the way the underlying Azure API is designed, this provider can only manage the deployment of the ARM Template - and not any resources which are created by it.
// This means that when deleting the `core.TemplateDeployment` resource, this provider will only remove the reference to the deployment, whilst leaving any resources created by that ARM Template Deployment.
// One workaround for this is to use a unique Resource Group for each ARM Template Deployment, which means deleting the Resource Group would contain any resources created within it - however this isn't ideal. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).
//
// ## Example Usage
//
// > **Note:** This example uses Storage Accounts and Public IP's which are natively supported by this provider - we'd highly recommend using the Native Resources where possible instead rather than an ARM Template, for the reasons outlined above.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleTemplateDeployment, err := core.NewTemplateDeployment(ctx, "exampleTemplateDeployment", &core.TemplateDeploymentArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TemplateBody:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"", "$", "schema\": \"https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\",\n", "  \"contentVersion\": \"1.0.0.0\",\n", "  \"parameters\": {\n", "    \"storageAccountType\": {\n", "      \"type\": \"string\",\n", "      \"defaultValue\": \"Standard_LRS\",\n", "      \"allowedValues\": [\n", "        \"Standard_LRS\",\n", "        \"Standard_GRS\",\n", "        \"Standard_ZRS\"\n", "      ],\n", "      \"metadata\": {\n", "        \"description\": \"Storage Account type\"\n", "      }\n", "    }\n", "  },\n", "  \"variables\": {\n", "    \"location\": \"[resourceGroup().location]\",\n", "    \"storageAccountName\": \"[concat(uniquestring(resourceGroup().id), 'storage')]\",\n", "    \"publicIPAddressName\": \"[concat('myPublicIp', uniquestring(resourceGroup().id))]\",\n", "    \"publicIPAddressType\": \"Dynamic\",\n", "    \"apiVersion\": \"2015-06-15\",\n", "    \"dnsLabelPrefix\": \"example-acctest\"\n", "  },\n", "  \"resources\": [\n", "    {\n", "      \"type\": \"Microsoft.Storage/storageAccounts\",\n", "      \"name\": \"[variables('storageAccountName')]\",\n", "      \"apiVersion\": \"[variables('apiVersion')]\",\n", "      \"location\": \"[variables('location')]\",\n", "      \"properties\": {\n", "        \"accountType\": \"[parameters('storageAccountType')]\"\n", "      }\n", "    },\n", "    {\n", "      \"type\": \"Microsoft.Network/publicIPAddresses\",\n", "      \"apiVersion\": \"[variables('apiVersion')]\",\n", "      \"name\": \"[variables('publicIPAddressName')]\",\n", "      \"location\": \"[variables('location')]\",\n", "      \"properties\": {\n", "        \"publicIPAllocationMethod\": \"[variables('publicIPAddressType')]\",\n", "        \"dnsSettings\": {\n", "          \"domainNameLabel\": \"[variables('dnsLabelPrefix')]\"\n", "        }\n", "      }\n", "    }\n", "  ],\n", "  \"outputs\": {\n", "    \"storageAccountName\": {\n", "      \"type\": \"string\",\n", "      \"value\": \"[variables('storageAccountName')]\"\n", "    }\n", "  }\n", "}\n")),
// 			Parameters: pulumi.StringMap{
// 				"storageAccountType": pulumi.String("Standard_GRS"),
// 			},
// 			DeploymentMode: pulumi.String("Incremental"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("storageAccountName", exampleTemplateDeployment.Outputs.ApplyT(func(outputs map[string]string) (string, error) {
// 			return outputs.StorageAccountName, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
// ## Note
//
// This provider does not know about the individual resources created by Azure using a deployment template and therefore cannot delete these resources during a destroy. Destroying a template deployment removes the associated deployment operations, but will not delete the Azure resources created by the deployment. In order to delete these resources, the containing resource group must also be destroyed. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).
type TemplateDeployment struct {
	pulumi.CustomResourceState

	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode pulumi.StringOutput `pulumi:"deploymentMode"`
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
	Outputs pulumi.StringMapOutput `pulumi:"outputs"`
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody pulumi.StringPtrOutput `pulumi:"parametersBody"`
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the JSON definition for the template.
	TemplateBody pulumi.StringOutput `pulumi:"templateBody"`
}

// NewTemplateDeployment registers a new resource with the given unique name, arguments, and options.
func NewTemplateDeployment(ctx *pulumi.Context,
	name string, args *TemplateDeploymentArgs, opts ...pulumi.ResourceOption) (*TemplateDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeploymentMode == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource TemplateDeployment
	err := ctx.RegisterResource("azure:core/templateDeployment:TemplateDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateDeployment gets an existing TemplateDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateDeploymentState, opts ...pulumi.ResourceOption) (*TemplateDeployment, error) {
	var resource TemplateDeployment
	err := ctx.ReadResource("azure:core/templateDeployment:TemplateDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateDeployment resources.
type templateDeploymentState struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
	Outputs map[string]string `pulumi:"outputs"`
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters map[string]string `pulumi:"parameters"`
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody *string `pulumi:"parametersBody"`
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the JSON definition for the template.
	TemplateBody *string `pulumi:"templateBody"`
}

type TemplateDeploymentState struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode pulumi.StringPtrInput
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using `.outputs["name"]`.
	Outputs pulumi.StringMapInput
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters pulumi.StringMapInput
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the JSON definition for the template.
	TemplateBody pulumi.StringPtrInput
}

func (TemplateDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateDeploymentState)(nil)).Elem()
}

type templateDeploymentArgs struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode string `pulumi:"deploymentMode"`
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters map[string]string `pulumi:"parameters"`
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody *string `pulumi:"parametersBody"`
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the JSON definition for the template.
	TemplateBody *string `pulumi:"templateBody"`
}

// The set of arguments for constructing a TemplateDeployment resource.
type TemplateDeploymentArgs struct {
	// Specifies the mode that is used to deploy resources. This value could be either `Incremental` or `Complete`.
	// Note that you will almost *always* want this to be set to `Incremental` otherwise the deployment will destroy all infrastructure not
	// specified within the template, and this provider will not be aware of this.
	DeploymentMode pulumi.StringInput
	// Specifies the name of the template deployment. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name and value pairs that define the deployment parameters for the template.
	Parameters pulumi.StringMapInput
	// Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references
	ParametersBody pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the template deployment.
	ResourceGroupName pulumi.StringInput
	// Specifies the JSON definition for the template.
	TemplateBody pulumi.StringPtrInput
}

func (TemplateDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateDeploymentArgs)(nil)).Elem()
}

type TemplateDeploymentInput interface {
	pulumi.Input

	ToTemplateDeploymentOutput() TemplateDeploymentOutput
	ToTemplateDeploymentOutputWithContext(ctx context.Context) TemplateDeploymentOutput
}

func (*TemplateDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateDeployment)(nil)).Elem()
}

func (i *TemplateDeployment) ToTemplateDeploymentOutput() TemplateDeploymentOutput {
	return i.ToTemplateDeploymentOutputWithContext(context.Background())
}

func (i *TemplateDeployment) ToTemplateDeploymentOutputWithContext(ctx context.Context) TemplateDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateDeploymentOutput)
}

// TemplateDeploymentArrayInput is an input type that accepts TemplateDeploymentArray and TemplateDeploymentArrayOutput values.
// You can construct a concrete instance of `TemplateDeploymentArrayInput` via:
//
//          TemplateDeploymentArray{ TemplateDeploymentArgs{...} }
type TemplateDeploymentArrayInput interface {
	pulumi.Input

	ToTemplateDeploymentArrayOutput() TemplateDeploymentArrayOutput
	ToTemplateDeploymentArrayOutputWithContext(context.Context) TemplateDeploymentArrayOutput
}

type TemplateDeploymentArray []TemplateDeploymentInput

func (TemplateDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateDeployment)(nil)).Elem()
}

func (i TemplateDeploymentArray) ToTemplateDeploymentArrayOutput() TemplateDeploymentArrayOutput {
	return i.ToTemplateDeploymentArrayOutputWithContext(context.Background())
}

func (i TemplateDeploymentArray) ToTemplateDeploymentArrayOutputWithContext(ctx context.Context) TemplateDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateDeploymentArrayOutput)
}

// TemplateDeploymentMapInput is an input type that accepts TemplateDeploymentMap and TemplateDeploymentMapOutput values.
// You can construct a concrete instance of `TemplateDeploymentMapInput` via:
//
//          TemplateDeploymentMap{ "key": TemplateDeploymentArgs{...} }
type TemplateDeploymentMapInput interface {
	pulumi.Input

	ToTemplateDeploymentMapOutput() TemplateDeploymentMapOutput
	ToTemplateDeploymentMapOutputWithContext(context.Context) TemplateDeploymentMapOutput
}

type TemplateDeploymentMap map[string]TemplateDeploymentInput

func (TemplateDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateDeployment)(nil)).Elem()
}

func (i TemplateDeploymentMap) ToTemplateDeploymentMapOutput() TemplateDeploymentMapOutput {
	return i.ToTemplateDeploymentMapOutputWithContext(context.Background())
}

func (i TemplateDeploymentMap) ToTemplateDeploymentMapOutputWithContext(ctx context.Context) TemplateDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateDeploymentMapOutput)
}

type TemplateDeploymentOutput struct{ *pulumi.OutputState }

func (TemplateDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateDeployment)(nil)).Elem()
}

func (o TemplateDeploymentOutput) ToTemplateDeploymentOutput() TemplateDeploymentOutput {
	return o
}

func (o TemplateDeploymentOutput) ToTemplateDeploymentOutputWithContext(ctx context.Context) TemplateDeploymentOutput {
	return o
}

type TemplateDeploymentArrayOutput struct{ *pulumi.OutputState }

func (TemplateDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateDeployment)(nil)).Elem()
}

func (o TemplateDeploymentArrayOutput) ToTemplateDeploymentArrayOutput() TemplateDeploymentArrayOutput {
	return o
}

func (o TemplateDeploymentArrayOutput) ToTemplateDeploymentArrayOutputWithContext(ctx context.Context) TemplateDeploymentArrayOutput {
	return o
}

func (o TemplateDeploymentArrayOutput) Index(i pulumi.IntInput) TemplateDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemplateDeployment {
		return vs[0].([]*TemplateDeployment)[vs[1].(int)]
	}).(TemplateDeploymentOutput)
}

type TemplateDeploymentMapOutput struct{ *pulumi.OutputState }

func (TemplateDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateDeployment)(nil)).Elem()
}

func (o TemplateDeploymentMapOutput) ToTemplateDeploymentMapOutput() TemplateDeploymentMapOutput {
	return o
}

func (o TemplateDeploymentMapOutput) ToTemplateDeploymentMapOutputWithContext(ctx context.Context) TemplateDeploymentMapOutput {
	return o
}

func (o TemplateDeploymentMapOutput) MapIndex(k pulumi.StringInput) TemplateDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemplateDeployment {
		return vs[0].(map[string]*TemplateDeployment)[vs[1].(string)]
	}).(TemplateDeploymentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateDeploymentInput)(nil)).Elem(), &TemplateDeployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateDeploymentArrayInput)(nil)).Elem(), TemplateDeploymentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateDeploymentMapInput)(nil)).Elem(), TemplateDeploymentMap{})
	pulumi.RegisterOutputType(TemplateDeploymentOutput{})
	pulumi.RegisterOutputType(TemplateDeploymentArrayOutput{})
	pulumi.RegisterOutputType(TemplateDeploymentMapOutput{})
}
