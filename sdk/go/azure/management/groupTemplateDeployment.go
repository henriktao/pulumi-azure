// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package management

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Template Deployment at a Management Group Scope.
//
// > **Note:** Deleting a Deployment at the Management Group Scope will not delete any resources created by the deployment.
//
// > **Note:** Deployments to a Management Group are always Incrementally applied. Existing resources that are not part of the template will not be removed.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00000000-0000-0000-0000-000000000000"
// 		exampleGroup, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewGroupTemplateDeployment(ctx, "exampleGroupTemplateDeployment", &management.GroupTemplateDeploymentArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ManagementGroupId: pulumi.String(exampleGroup.Id),
// 			TemplateContent:   pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"", "$", "schema\": \"https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#\",\n", "  \"contentVersion\": \"1.0.0.0\",\n", "  \"parameters\": {\n", "    \"policyAssignmentName\": {\n", "      \"type\": \"string\",\n", "      \"defaultValue\": \"[guid(parameters('policyDefinitionID'), resourceGroup().name)]\",\n", "      \"metadata\": {\n", "        \"description\": \"Specifies the name of the policy assignment, can be used defined or an idempotent name as the defaultValue provides.\"\n", "      }\n", "    },\n", "    \"policyDefinitionID\": {\n", "      \"type\": \"string\",\n", "      \"metadata\": {\n", "        \"description\": \"Specifies the ID of the policy definition or policy set definition being assigned.\"\n", "      }\n", "    }\n", "  },\n", "  \"resources\": [\n", "    {\n", "      \"type\": \"Microsoft.Authorization/policyAssignments\",\n", "      \"name\": \"[parameters('policyAssignmentName')]\",\n", "      \"apiVersion\": \"2019-09-01\",\n", "      \"properties\": {\n", "        \"scope\": \"[subscriptionResourceId('Microsoft.Resources/resourceGroups', resourceGroup().name)]\",\n", "        \"policyDefinitionId\": \"[parameters('policyDefinitionID')]\"\n", "      }\n", "    }\n", "  ]\n", "}\n")),
// 			ParametersContent: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"", "$", "schema\": \"https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#\",\n", "  \"contentVersion\": \"1.0.0.0\",\n", "  \"parameters\": {\n", "    \"policyDefinitionID\": {\n", "      \"value\": \"/providers/Microsoft.Authorization/policyDefinitions/0a914e76-4921-4c19-b460-a2d36003525a\"\n", "    }\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00000000-0000-0000-0000-000000000000"
// 		exampleGroup, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewGroupTemplateDeployment(ctx, "exampleGroupTemplateDeployment", &management.GroupTemplateDeploymentArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ManagementGroupId: pulumi.String(exampleGroup.Id),
// 			TemplateContent:   readFileOrPanic("templates/example-deploy-template.json"),
// 			ParametersContent: readFileOrPanic("templates/example-deploy-params.json"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/management"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00000000-0000-0000-0000-000000000000"
// 		exampleGroup, err := management.LookupGroup(ctx, &management.LookupGroupArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleTemplateSpecVersion, err := core.GetTemplateSpecVersion(ctx, &core.GetTemplateSpecVersionArgs{
// 			Name:              "exampleTemplateForManagementGroup",
// 			ResourceGroupName: "exampleResourceGroup",
// 			Version:           "v1.0.9",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = management.NewGroupTemplateDeployment(ctx, "exampleGroupTemplateDeployment", &management.GroupTemplateDeploymentArgs{
// 			Location:              pulumi.String("West Europe"),
// 			ManagementGroupId:     pulumi.String(exampleGroup.Id),
// 			TemplateSpecVersionId: pulumi.String(exampleTemplateSpecVersion.Id),
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
// Management Group Template Deployments can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:management/groupTemplateDeployment:GroupTemplateDeployment example /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
// ```
type GroupTemplateDeployment struct {
	pulumi.CustomResourceState

	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrOutput `pulumi:"debugLevel"`
	// The Azure Region where the Template should exist. Changing this forces a new Template to be created.
	Location          pulumi.StringOutput `pulumi:"location"`
	ManagementGroupId pulumi.StringOutput `pulumi:"managementGroupId"`
	// The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent pulumi.StringOutput `pulumi:"outputContent"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringOutput `pulumi:"parametersContent"`
	// A mapping of tags which should be assigned to the Template.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
	TemplateContent pulumi.StringOutput `pulumi:"templateContent"`
	// The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
	TemplateSpecVersionId pulumi.StringPtrOutput `pulumi:"templateSpecVersionId"`
}

// NewGroupTemplateDeployment registers a new resource with the given unique name, arguments, and options.
func NewGroupTemplateDeployment(ctx *pulumi.Context,
	name string, args *GroupTemplateDeploymentArgs, opts ...pulumi.ResourceOption) (*GroupTemplateDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	var resource GroupTemplateDeployment
	err := ctx.RegisterResource("azure:management/groupTemplateDeployment:GroupTemplateDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupTemplateDeployment gets an existing GroupTemplateDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupTemplateDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupTemplateDeploymentState, opts ...pulumi.ResourceOption) (*GroupTemplateDeployment, error) {
	var resource GroupTemplateDeployment
	err := ctx.ReadResource("azure:management/groupTemplateDeployment:GroupTemplateDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupTemplateDeployment resources.
type groupTemplateDeploymentState struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel *string `pulumi:"debugLevel"`
	// The Azure Region where the Template should exist. Changing this forces a new Template to be created.
	Location          *string `pulumi:"location"`
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
	Name *string `pulumi:"name"`
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent *string `pulumi:"outputContent"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `pulumi:"parametersContent"`
	// A mapping of tags which should be assigned to the Template.
	Tags map[string]string `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
	TemplateContent *string `pulumi:"templateContent"`
	// The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
	TemplateSpecVersionId *string `pulumi:"templateSpecVersionId"`
}

type GroupTemplateDeploymentState struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrInput
	// The Azure Region where the Template should exist. Changing this forces a new Template to be created.
	Location          pulumi.StringPtrInput
	ManagementGroupId pulumi.StringPtrInput
	// The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
	Name pulumi.StringPtrInput
	// The JSON Content of the Outputs of the ARM Template Deployment.
	OutputContent pulumi.StringPtrInput
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Template.
	Tags pulumi.StringMapInput
	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
	TemplateContent pulumi.StringPtrInput
	// The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
	TemplateSpecVersionId pulumi.StringPtrInput
}

func (GroupTemplateDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupTemplateDeploymentState)(nil)).Elem()
}

type groupTemplateDeploymentArgs struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel *string `pulumi:"debugLevel"`
	// The Azure Region where the Template should exist. Changing this forces a new Template to be created.
	Location          *string `pulumi:"location"`
	ManagementGroupId string  `pulumi:"managementGroupId"`
	// The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
	Name *string `pulumi:"name"`
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent *string `pulumi:"parametersContent"`
	// A mapping of tags which should be assigned to the Template.
	Tags map[string]string `pulumi:"tags"`
	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
	TemplateContent *string `pulumi:"templateContent"`
	// The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
	TemplateSpecVersionId *string `pulumi:"templateSpecVersionId"`
}

// The set of arguments for constructing a GroupTemplateDeployment resource.
type GroupTemplateDeploymentArgs struct {
	// The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
	DebugLevel pulumi.StringPtrInput
	// The Azure Region where the Template should exist. Changing this forces a new Template to be created.
	Location          pulumi.StringPtrInput
	ManagementGroupId pulumi.StringInput
	// The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
	Name pulumi.StringPtrInput
	// The contents of the ARM Template parameters file - containing a JSON list of parameters.
	ParametersContent pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Template.
	Tags pulumi.StringMapInput
	// The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
	TemplateContent pulumi.StringPtrInput
	// The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
	TemplateSpecVersionId pulumi.StringPtrInput
}

func (GroupTemplateDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupTemplateDeploymentArgs)(nil)).Elem()
}

type GroupTemplateDeploymentInput interface {
	pulumi.Input

	ToGroupTemplateDeploymentOutput() GroupTemplateDeploymentOutput
	ToGroupTemplateDeploymentOutputWithContext(ctx context.Context) GroupTemplateDeploymentOutput
}

func (*GroupTemplateDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTemplateDeployment)(nil))
}

func (i *GroupTemplateDeployment) ToGroupTemplateDeploymentOutput() GroupTemplateDeploymentOutput {
	return i.ToGroupTemplateDeploymentOutputWithContext(context.Background())
}

func (i *GroupTemplateDeployment) ToGroupTemplateDeploymentOutputWithContext(ctx context.Context) GroupTemplateDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTemplateDeploymentOutput)
}

func (i *GroupTemplateDeployment) ToGroupTemplateDeploymentPtrOutput() GroupTemplateDeploymentPtrOutput {
	return i.ToGroupTemplateDeploymentPtrOutputWithContext(context.Background())
}

func (i *GroupTemplateDeployment) ToGroupTemplateDeploymentPtrOutputWithContext(ctx context.Context) GroupTemplateDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTemplateDeploymentPtrOutput)
}

type GroupTemplateDeploymentPtrInput interface {
	pulumi.Input

	ToGroupTemplateDeploymentPtrOutput() GroupTemplateDeploymentPtrOutput
	ToGroupTemplateDeploymentPtrOutputWithContext(ctx context.Context) GroupTemplateDeploymentPtrOutput
}

type groupTemplateDeploymentPtrType GroupTemplateDeploymentArgs

func (*groupTemplateDeploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupTemplateDeployment)(nil))
}

func (i *groupTemplateDeploymentPtrType) ToGroupTemplateDeploymentPtrOutput() GroupTemplateDeploymentPtrOutput {
	return i.ToGroupTemplateDeploymentPtrOutputWithContext(context.Background())
}

func (i *groupTemplateDeploymentPtrType) ToGroupTemplateDeploymentPtrOutputWithContext(ctx context.Context) GroupTemplateDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTemplateDeploymentPtrOutput)
}

// GroupTemplateDeploymentArrayInput is an input type that accepts GroupTemplateDeploymentArray and GroupTemplateDeploymentArrayOutput values.
// You can construct a concrete instance of `GroupTemplateDeploymentArrayInput` via:
//
//          GroupTemplateDeploymentArray{ GroupTemplateDeploymentArgs{...} }
type GroupTemplateDeploymentArrayInput interface {
	pulumi.Input

	ToGroupTemplateDeploymentArrayOutput() GroupTemplateDeploymentArrayOutput
	ToGroupTemplateDeploymentArrayOutputWithContext(context.Context) GroupTemplateDeploymentArrayOutput
}

type GroupTemplateDeploymentArray []GroupTemplateDeploymentInput

func (GroupTemplateDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupTemplateDeployment)(nil)).Elem()
}

func (i GroupTemplateDeploymentArray) ToGroupTemplateDeploymentArrayOutput() GroupTemplateDeploymentArrayOutput {
	return i.ToGroupTemplateDeploymentArrayOutputWithContext(context.Background())
}

func (i GroupTemplateDeploymentArray) ToGroupTemplateDeploymentArrayOutputWithContext(ctx context.Context) GroupTemplateDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTemplateDeploymentArrayOutput)
}

// GroupTemplateDeploymentMapInput is an input type that accepts GroupTemplateDeploymentMap and GroupTemplateDeploymentMapOutput values.
// You can construct a concrete instance of `GroupTemplateDeploymentMapInput` via:
//
//          GroupTemplateDeploymentMap{ "key": GroupTemplateDeploymentArgs{...} }
type GroupTemplateDeploymentMapInput interface {
	pulumi.Input

	ToGroupTemplateDeploymentMapOutput() GroupTemplateDeploymentMapOutput
	ToGroupTemplateDeploymentMapOutputWithContext(context.Context) GroupTemplateDeploymentMapOutput
}

type GroupTemplateDeploymentMap map[string]GroupTemplateDeploymentInput

func (GroupTemplateDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupTemplateDeployment)(nil)).Elem()
}

func (i GroupTemplateDeploymentMap) ToGroupTemplateDeploymentMapOutput() GroupTemplateDeploymentMapOutput {
	return i.ToGroupTemplateDeploymentMapOutputWithContext(context.Background())
}

func (i GroupTemplateDeploymentMap) ToGroupTemplateDeploymentMapOutputWithContext(ctx context.Context) GroupTemplateDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupTemplateDeploymentMapOutput)
}

type GroupTemplateDeploymentOutput struct{ *pulumi.OutputState }

func (GroupTemplateDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupTemplateDeployment)(nil))
}

func (o GroupTemplateDeploymentOutput) ToGroupTemplateDeploymentOutput() GroupTemplateDeploymentOutput {
	return o
}

func (o GroupTemplateDeploymentOutput) ToGroupTemplateDeploymentOutputWithContext(ctx context.Context) GroupTemplateDeploymentOutput {
	return o
}

func (o GroupTemplateDeploymentOutput) ToGroupTemplateDeploymentPtrOutput() GroupTemplateDeploymentPtrOutput {
	return o.ToGroupTemplateDeploymentPtrOutputWithContext(context.Background())
}

func (o GroupTemplateDeploymentOutput) ToGroupTemplateDeploymentPtrOutputWithContext(ctx context.Context) GroupTemplateDeploymentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupTemplateDeployment) *GroupTemplateDeployment {
		return &v
	}).(GroupTemplateDeploymentPtrOutput)
}

type GroupTemplateDeploymentPtrOutput struct{ *pulumi.OutputState }

func (GroupTemplateDeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupTemplateDeployment)(nil))
}

func (o GroupTemplateDeploymentPtrOutput) ToGroupTemplateDeploymentPtrOutput() GroupTemplateDeploymentPtrOutput {
	return o
}

func (o GroupTemplateDeploymentPtrOutput) ToGroupTemplateDeploymentPtrOutputWithContext(ctx context.Context) GroupTemplateDeploymentPtrOutput {
	return o
}

func (o GroupTemplateDeploymentPtrOutput) Elem() GroupTemplateDeploymentOutput {
	return o.ApplyT(func(v *GroupTemplateDeployment) GroupTemplateDeployment {
		if v != nil {
			return *v
		}
		var ret GroupTemplateDeployment
		return ret
	}).(GroupTemplateDeploymentOutput)
}

type GroupTemplateDeploymentArrayOutput struct{ *pulumi.OutputState }

func (GroupTemplateDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupTemplateDeployment)(nil))
}

func (o GroupTemplateDeploymentArrayOutput) ToGroupTemplateDeploymentArrayOutput() GroupTemplateDeploymentArrayOutput {
	return o
}

func (o GroupTemplateDeploymentArrayOutput) ToGroupTemplateDeploymentArrayOutputWithContext(ctx context.Context) GroupTemplateDeploymentArrayOutput {
	return o
}

func (o GroupTemplateDeploymentArrayOutput) Index(i pulumi.IntInput) GroupTemplateDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupTemplateDeployment {
		return vs[0].([]GroupTemplateDeployment)[vs[1].(int)]
	}).(GroupTemplateDeploymentOutput)
}

type GroupTemplateDeploymentMapOutput struct{ *pulumi.OutputState }

func (GroupTemplateDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupTemplateDeployment)(nil))
}

func (o GroupTemplateDeploymentMapOutput) ToGroupTemplateDeploymentMapOutput() GroupTemplateDeploymentMapOutput {
	return o
}

func (o GroupTemplateDeploymentMapOutput) ToGroupTemplateDeploymentMapOutputWithContext(ctx context.Context) GroupTemplateDeploymentMapOutput {
	return o
}

func (o GroupTemplateDeploymentMapOutput) MapIndex(k pulumi.StringInput) GroupTemplateDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupTemplateDeployment {
		return vs[0].(map[string]GroupTemplateDeployment)[vs[1].(string)]
	}).(GroupTemplateDeploymentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTemplateDeploymentInput)(nil)).Elem(), &GroupTemplateDeployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTemplateDeploymentPtrInput)(nil)).Elem(), &GroupTemplateDeployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTemplateDeploymentArrayInput)(nil)).Elem(), GroupTemplateDeploymentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupTemplateDeploymentMapInput)(nil)).Elem(), GroupTemplateDeploymentMap{})
	pulumi.RegisterOutputType(GroupTemplateDeploymentOutput{})
	pulumi.RegisterOutputType(GroupTemplateDeploymentPtrOutput{})
	pulumi.RegisterOutputType(GroupTemplateDeploymentArrayOutput{})
	pulumi.RegisterOutputType(GroupTemplateDeploymentMapOutput{})
}
