// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Management Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("pub1"),
// 			PublisherEmail:    pulumi.String("pub1@email.com"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewGroup(ctx, "exampleGroup", &apimanagement.GroupArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			DisplayName:       pulumi.String("Example Group"),
// 			Description:       pulumi.String("This is an example API management group."),
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
// API Management Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/group:Group example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/groups/example-apimg
// ```
type Group struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Group
	err := ctx.RegisterResource("azure:apimanagement/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure:apimanagement/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description *string `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName *string `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId *string `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The description of this API Management Group.
	Description pulumi.StringPtrInput
	// The display name of this API Management Group.
	DisplayName pulumi.StringPtrInput
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrInput
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description *string `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName string `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId *string `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The description of this API Management Group.
	Description pulumi.StringPtrInput
	// The display name of this API Management Group.
	DisplayName pulumi.StringInput
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrInput
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *Group) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

type GroupPtrInput interface {
	pulumi.Input

	ToGroupPtrOutput() GroupPtrOutput
	ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput
}

type groupPtrType GroupArgs

func (*groupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (i *groupPtrType) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *groupPtrType) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Group)(nil))
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Group)(nil))
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o.ToGroupPtrOutputWithContext(context.Background())
}

func (o GroupOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o.ApplyT(func(v Group) *Group {
		return &v
	}).(GroupPtrOutput)
}

type GroupPtrOutput struct {
	*pulumi.OutputState
}

func (GroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (o GroupPtrOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Group)(nil))
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Group {
		return vs[0].([]Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Group)(nil))
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Group {
		return vs[0].(map[string]Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupPtrOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
