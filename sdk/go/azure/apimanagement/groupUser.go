// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Management User Assignment to a Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleUser, err := apimanagement.LookupUser(ctx, &apimanagement.LookupUserArgs{
// 			UserId:            "my-user",
// 			ApiManagementName: "example-apim",
// 			ResourceGroupName: "search-service",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewGroupUser(ctx, "exampleGroupUser", &apimanagement.GroupUserArgs{
// 			UserId:            pulumi.String(exampleUser.Id),
// 			GroupName:         pulumi.String("example-group"),
// 			ResourceGroupName: pulumi.String(exampleUser.ResourceGroupName),
// 			ApiManagementName: pulumi.String(exampleUser.ApiManagementName),
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
// API Management Group Users can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/groupUser:GroupUser example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
// ```
type GroupUser struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewGroupUser registers a new resource with the given unique name, arguments, and options.
func NewGroupUser(ctx *pulumi.Context,
	name string, args *GroupUserArgs, opts ...pulumi.ResourceOption) (*GroupUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource GroupUser
	err := ctx.RegisterResource("azure:apimanagement/groupUser:GroupUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupUser gets an existing GroupUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupUserState, opts ...pulumi.ResourceOption) (*GroupUser, error) {
	var resource GroupUser
	err := ctx.ReadResource("azure:apimanagement/groupUser:GroupUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupUser resources.
type groupUserState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName *string `pulumi:"groupName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
	UserId *string `pulumi:"userId"`
}

type GroupUserState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
	UserId pulumi.StringPtrInput
}

func (GroupUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupUserState)(nil)).Elem()
}

type groupUserArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName string `pulumi:"groupName"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a GroupUser resource.
type GroupUserArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The Name of the API Management Group within the API Management Service. Changing this forces a new resource to be created.
	GroupName pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the API Management User which should be assigned to this API Management Group. Changing this forces a new resource to be created.
	UserId pulumi.StringInput
}

func (GroupUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupUserArgs)(nil)).Elem()
}

type GroupUserInput interface {
	pulumi.Input

	ToGroupUserOutput() GroupUserOutput
	ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput
}

func (*GroupUser) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupUser)(nil))
}

func (i *GroupUser) ToGroupUserOutput() GroupUserOutput {
	return i.ToGroupUserOutputWithContext(context.Background())
}

func (i *GroupUser) ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserOutput)
}

func (i *GroupUser) ToGroupUserPtrOutput() GroupUserPtrOutput {
	return i.ToGroupUserPtrOutputWithContext(context.Background())
}

func (i *GroupUser) ToGroupUserPtrOutputWithContext(ctx context.Context) GroupUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserPtrOutput)
}

type GroupUserPtrInput interface {
	pulumi.Input

	ToGroupUserPtrOutput() GroupUserPtrOutput
	ToGroupUserPtrOutputWithContext(ctx context.Context) GroupUserPtrOutput
}

type groupUserPtrType GroupUserArgs

func (*groupUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupUser)(nil))
}

func (i *groupUserPtrType) ToGroupUserPtrOutput() GroupUserPtrOutput {
	return i.ToGroupUserPtrOutputWithContext(context.Background())
}

func (i *groupUserPtrType) ToGroupUserPtrOutputWithContext(ctx context.Context) GroupUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserPtrOutput)
}

// GroupUserArrayInput is an input type that accepts GroupUserArray and GroupUserArrayOutput values.
// You can construct a concrete instance of `GroupUserArrayInput` via:
//
//          GroupUserArray{ GroupUserArgs{...} }
type GroupUserArrayInput interface {
	pulumi.Input

	ToGroupUserArrayOutput() GroupUserArrayOutput
	ToGroupUserArrayOutputWithContext(context.Context) GroupUserArrayOutput
}

type GroupUserArray []GroupUserInput

func (GroupUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupUser)(nil)).Elem()
}

func (i GroupUserArray) ToGroupUserArrayOutput() GroupUserArrayOutput {
	return i.ToGroupUserArrayOutputWithContext(context.Background())
}

func (i GroupUserArray) ToGroupUserArrayOutputWithContext(ctx context.Context) GroupUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserArrayOutput)
}

// GroupUserMapInput is an input type that accepts GroupUserMap and GroupUserMapOutput values.
// You can construct a concrete instance of `GroupUserMapInput` via:
//
//          GroupUserMap{ "key": GroupUserArgs{...} }
type GroupUserMapInput interface {
	pulumi.Input

	ToGroupUserMapOutput() GroupUserMapOutput
	ToGroupUserMapOutputWithContext(context.Context) GroupUserMapOutput
}

type GroupUserMap map[string]GroupUserInput

func (GroupUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupUser)(nil)).Elem()
}

func (i GroupUserMap) ToGroupUserMapOutput() GroupUserMapOutput {
	return i.ToGroupUserMapOutputWithContext(context.Background())
}

func (i GroupUserMap) ToGroupUserMapOutputWithContext(ctx context.Context) GroupUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserMapOutput)
}

type GroupUserOutput struct{ *pulumi.OutputState }

func (GroupUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupUser)(nil))
}

func (o GroupUserOutput) ToGroupUserOutput() GroupUserOutput {
	return o
}

func (o GroupUserOutput) ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput {
	return o
}

func (o GroupUserOutput) ToGroupUserPtrOutput() GroupUserPtrOutput {
	return o.ToGroupUserPtrOutputWithContext(context.Background())
}

func (o GroupUserOutput) ToGroupUserPtrOutputWithContext(ctx context.Context) GroupUserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupUser) *GroupUser {
		return &v
	}).(GroupUserPtrOutput)
}

type GroupUserPtrOutput struct{ *pulumi.OutputState }

func (GroupUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupUser)(nil))
}

func (o GroupUserPtrOutput) ToGroupUserPtrOutput() GroupUserPtrOutput {
	return o
}

func (o GroupUserPtrOutput) ToGroupUserPtrOutputWithContext(ctx context.Context) GroupUserPtrOutput {
	return o
}

func (o GroupUserPtrOutput) Elem() GroupUserOutput {
	return o.ApplyT(func(v *GroupUser) GroupUser {
		if v != nil {
			return *v
		}
		var ret GroupUser
		return ret
	}).(GroupUserOutput)
}

type GroupUserArrayOutput struct{ *pulumi.OutputState }

func (GroupUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupUser)(nil))
}

func (o GroupUserArrayOutput) ToGroupUserArrayOutput() GroupUserArrayOutput {
	return o
}

func (o GroupUserArrayOutput) ToGroupUserArrayOutputWithContext(ctx context.Context) GroupUserArrayOutput {
	return o
}

func (o GroupUserArrayOutput) Index(i pulumi.IntInput) GroupUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupUser {
		return vs[0].([]GroupUser)[vs[1].(int)]
	}).(GroupUserOutput)
}

type GroupUserMapOutput struct{ *pulumi.OutputState }

func (GroupUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupUser)(nil))
}

func (o GroupUserMapOutput) ToGroupUserMapOutput() GroupUserMapOutput {
	return o
}

func (o GroupUserMapOutput) ToGroupUserMapOutputWithContext(ctx context.Context) GroupUserMapOutput {
	return o
}

func (o GroupUserMapOutput) MapIndex(k pulumi.StringInput) GroupUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupUser {
		return vs[0].(map[string]GroupUser)[vs[1].(string)]
	}).(GroupUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupUserInput)(nil)).Elem(), &GroupUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupUserPtrInput)(nil)).Elem(), &GroupUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupUserArrayInput)(nil)).Elem(), GroupUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupUserMapInput)(nil)).Elem(), GroupUserMap{})
	pulumi.RegisterOutputType(GroupUserOutput{})
	pulumi.RegisterOutputType(GroupUserPtrOutput{})
	pulumi.RegisterOutputType(GroupUserArrayOutput{})
	pulumi.RegisterOutputType(GroupUserMapOutput{})
}
