// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a API Management API Operation Tag.
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
// 		exampleApi, err := apimanagement.LookupApi(ctx, &apimanagement.LookupApiArgs{
// 			Name:              "search-api",
// 			ApiManagementName: "search-api-management",
// 			ResourceGroupName: "search-service",
// 			Revision:          "2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleApiOperation, err := apimanagement.NewApiOperation(ctx, "exampleApiOperation", &apimanagement.ApiOperationArgs{
// 			OperationId:       pulumi.String("user-delete"),
// 			ApiName:           pulumi.String(exampleApi.Name),
// 			ApiManagementName: pulumi.String(exampleApi.ApiManagementName),
// 			ResourceGroupName: pulumi.String(exampleApi.ResourceGroupName),
// 			DisplayName:       pulumi.String("Delete User Operation"),
// 			Method:            pulumi.String("DELETE"),
// 			UrlTemplate:       pulumi.String("/users/{id}/delete"),
// 			Description:       pulumi.String("This can only be done by the logged in user."),
// 			Responses: apimanagement.ApiOperationResponseArray{
// 				&apimanagement.ApiOperationResponseArgs{
// 					StatusCode: pulumi.Int(200),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewApiOperationTag(ctx, "exampleApiOperationTag", &apimanagement.ApiOperationTagArgs{
// 			ApiOperationId: exampleApiOperation.ID(),
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
// API Management API Operation Tags can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/apiOperationTag:ApiOperationTag example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1/tags/tag1
// ```
type ApiOperationTag struct {
	pulumi.CustomResourceState

	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	ApiOperationId pulumi.StringOutput `pulumi:"apiOperationId"`
	// The display name of the API Management API Operation Tag.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name which should be used for this API Management API Operation Tag. Changing this forces a new API Management API Operation Tag to be created. The name must be unique in the API Management Service.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewApiOperationTag registers a new resource with the given unique name, arguments, and options.
func NewApiOperationTag(ctx *pulumi.Context,
	name string, args *ApiOperationTagArgs, opts ...pulumi.ResourceOption) (*ApiOperationTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiOperationId == nil {
		return nil, errors.New("invalid value for required argument 'ApiOperationId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource ApiOperationTag
	err := ctx.RegisterResource("azure:apimanagement/apiOperationTag:ApiOperationTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperationTag gets an existing ApiOperationTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperationTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationTagState, opts ...pulumi.ResourceOption) (*ApiOperationTag, error) {
	var resource ApiOperationTag
	err := ctx.ReadResource("azure:apimanagement/apiOperationTag:ApiOperationTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationTag resources.
type apiOperationTagState struct {
	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	ApiOperationId *string `pulumi:"apiOperationId"`
	// The display name of the API Management API Operation Tag.
	DisplayName *string `pulumi:"displayName"`
	// The name which should be used for this API Management API Operation Tag. Changing this forces a new API Management API Operation Tag to be created. The name must be unique in the API Management Service.
	Name *string `pulumi:"name"`
}

type ApiOperationTagState struct {
	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	ApiOperationId pulumi.StringPtrInput
	// The display name of the API Management API Operation Tag.
	DisplayName pulumi.StringPtrInput
	// The name which should be used for this API Management API Operation Tag. Changing this forces a new API Management API Operation Tag to be created. The name must be unique in the API Management Service.
	Name pulumi.StringPtrInput
}

func (ApiOperationTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationTagState)(nil)).Elem()
}

type apiOperationTagArgs struct {
	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	ApiOperationId string `pulumi:"apiOperationId"`
	// The display name of the API Management API Operation Tag.
	DisplayName string `pulumi:"displayName"`
	// The name which should be used for this API Management API Operation Tag. Changing this forces a new API Management API Operation Tag to be created. The name must be unique in the API Management Service.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ApiOperationTag resource.
type ApiOperationTagArgs struct {
	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	ApiOperationId pulumi.StringInput
	// The display name of the API Management API Operation Tag.
	DisplayName pulumi.StringInput
	// The name which should be used for this API Management API Operation Tag. Changing this forces a new API Management API Operation Tag to be created. The name must be unique in the API Management Service.
	Name pulumi.StringPtrInput
}

func (ApiOperationTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationTagArgs)(nil)).Elem()
}

type ApiOperationTagInput interface {
	pulumi.Input

	ToApiOperationTagOutput() ApiOperationTagOutput
	ToApiOperationTagOutputWithContext(ctx context.Context) ApiOperationTagOutput
}

func (*ApiOperationTag) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperationTag)(nil))
}

func (i *ApiOperationTag) ToApiOperationTagOutput() ApiOperationTagOutput {
	return i.ToApiOperationTagOutputWithContext(context.Background())
}

func (i *ApiOperationTag) ToApiOperationTagOutputWithContext(ctx context.Context) ApiOperationTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationTagOutput)
}

func (i *ApiOperationTag) ToApiOperationTagPtrOutput() ApiOperationTagPtrOutput {
	return i.ToApiOperationTagPtrOutputWithContext(context.Background())
}

func (i *ApiOperationTag) ToApiOperationTagPtrOutputWithContext(ctx context.Context) ApiOperationTagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationTagPtrOutput)
}

type ApiOperationTagPtrInput interface {
	pulumi.Input

	ToApiOperationTagPtrOutput() ApiOperationTagPtrOutput
	ToApiOperationTagPtrOutputWithContext(ctx context.Context) ApiOperationTagPtrOutput
}

type apiOperationTagPtrType ApiOperationTagArgs

func (*apiOperationTagPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationTag)(nil))
}

func (i *apiOperationTagPtrType) ToApiOperationTagPtrOutput() ApiOperationTagPtrOutput {
	return i.ToApiOperationTagPtrOutputWithContext(context.Background())
}

func (i *apiOperationTagPtrType) ToApiOperationTagPtrOutputWithContext(ctx context.Context) ApiOperationTagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationTagPtrOutput)
}

// ApiOperationTagArrayInput is an input type that accepts ApiOperationTagArray and ApiOperationTagArrayOutput values.
// You can construct a concrete instance of `ApiOperationTagArrayInput` via:
//
//          ApiOperationTagArray{ ApiOperationTagArgs{...} }
type ApiOperationTagArrayInput interface {
	pulumi.Input

	ToApiOperationTagArrayOutput() ApiOperationTagArrayOutput
	ToApiOperationTagArrayOutputWithContext(context.Context) ApiOperationTagArrayOutput
}

type ApiOperationTagArray []ApiOperationTagInput

func (ApiOperationTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiOperationTag)(nil)).Elem()
}

func (i ApiOperationTagArray) ToApiOperationTagArrayOutput() ApiOperationTagArrayOutput {
	return i.ToApiOperationTagArrayOutputWithContext(context.Background())
}

func (i ApiOperationTagArray) ToApiOperationTagArrayOutputWithContext(ctx context.Context) ApiOperationTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationTagArrayOutput)
}

// ApiOperationTagMapInput is an input type that accepts ApiOperationTagMap and ApiOperationTagMapOutput values.
// You can construct a concrete instance of `ApiOperationTagMapInput` via:
//
//          ApiOperationTagMap{ "key": ApiOperationTagArgs{...} }
type ApiOperationTagMapInput interface {
	pulumi.Input

	ToApiOperationTagMapOutput() ApiOperationTagMapOutput
	ToApiOperationTagMapOutputWithContext(context.Context) ApiOperationTagMapOutput
}

type ApiOperationTagMap map[string]ApiOperationTagInput

func (ApiOperationTagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiOperationTag)(nil)).Elem()
}

func (i ApiOperationTagMap) ToApiOperationTagMapOutput() ApiOperationTagMapOutput {
	return i.ToApiOperationTagMapOutputWithContext(context.Background())
}

func (i ApiOperationTagMap) ToApiOperationTagMapOutputWithContext(ctx context.Context) ApiOperationTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationTagMapOutput)
}

type ApiOperationTagOutput struct{ *pulumi.OutputState }

func (ApiOperationTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperationTag)(nil))
}

func (o ApiOperationTagOutput) ToApiOperationTagOutput() ApiOperationTagOutput {
	return o
}

func (o ApiOperationTagOutput) ToApiOperationTagOutputWithContext(ctx context.Context) ApiOperationTagOutput {
	return o
}

func (o ApiOperationTagOutput) ToApiOperationTagPtrOutput() ApiOperationTagPtrOutput {
	return o.ToApiOperationTagPtrOutputWithContext(context.Background())
}

func (o ApiOperationTagOutput) ToApiOperationTagPtrOutputWithContext(ctx context.Context) ApiOperationTagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiOperationTag) *ApiOperationTag {
		return &v
	}).(ApiOperationTagPtrOutput)
}

type ApiOperationTagPtrOutput struct{ *pulumi.OutputState }

func (ApiOperationTagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationTag)(nil))
}

func (o ApiOperationTagPtrOutput) ToApiOperationTagPtrOutput() ApiOperationTagPtrOutput {
	return o
}

func (o ApiOperationTagPtrOutput) ToApiOperationTagPtrOutputWithContext(ctx context.Context) ApiOperationTagPtrOutput {
	return o
}

func (o ApiOperationTagPtrOutput) Elem() ApiOperationTagOutput {
	return o.ApplyT(func(v *ApiOperationTag) ApiOperationTag {
		if v != nil {
			return *v
		}
		var ret ApiOperationTag
		return ret
	}).(ApiOperationTagOutput)
}

type ApiOperationTagArrayOutput struct{ *pulumi.OutputState }

func (ApiOperationTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiOperationTag)(nil))
}

func (o ApiOperationTagArrayOutput) ToApiOperationTagArrayOutput() ApiOperationTagArrayOutput {
	return o
}

func (o ApiOperationTagArrayOutput) ToApiOperationTagArrayOutputWithContext(ctx context.Context) ApiOperationTagArrayOutput {
	return o
}

func (o ApiOperationTagArrayOutput) Index(i pulumi.IntInput) ApiOperationTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiOperationTag {
		return vs[0].([]ApiOperationTag)[vs[1].(int)]
	}).(ApiOperationTagOutput)
}

type ApiOperationTagMapOutput struct{ *pulumi.OutputState }

func (ApiOperationTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiOperationTag)(nil))
}

func (o ApiOperationTagMapOutput) ToApiOperationTagMapOutput() ApiOperationTagMapOutput {
	return o
}

func (o ApiOperationTagMapOutput) ToApiOperationTagMapOutputWithContext(ctx context.Context) ApiOperationTagMapOutput {
	return o
}

func (o ApiOperationTagMapOutput) MapIndex(k pulumi.StringInput) ApiOperationTagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiOperationTag {
		return vs[0].(map[string]ApiOperationTag)[vs[1].(string)]
	}).(ApiOperationTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationTagInput)(nil)).Elem(), &ApiOperationTag{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationTagPtrInput)(nil)).Elem(), &ApiOperationTag{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationTagArrayInput)(nil)).Elem(), ApiOperationTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationTagMapInput)(nil)).Elem(), ApiOperationTagMap{})
	pulumi.RegisterOutputType(ApiOperationTagOutput{})
	pulumi.RegisterOutputType(ApiOperationTagPtrOutput{})
	pulumi.RegisterOutputType(ApiOperationTagArrayOutput{})
	pulumi.RegisterOutputType(ApiOperationTagMapOutput{})
}
