// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Management API Operation Policy
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApiOperation, err := apimanagement.NewApiOperation(ctx, "exampleApiOperation", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewApiOperationPolicy(ctx, "exampleApiOperationPolicy", &apimanagement.ApiOperationPolicyArgs{
// 			ApiName:           exampleApiOperation.ApiName,
// 			ApiManagementName: exampleApiOperation.ApiManagementName,
// 			ResourceGroupName: exampleApiOperation.ResourceGroupName,
// 			OperationId:       exampleApiOperation.OperationId,
// 			XmlContent:        pulumi.String(fmt.Sprintf("%v%v%v%v%v", "<policies>\n", "  <inbound>\n", "    <find-and-replace from=\"xyz\" to=\"abc\" />\n", "  </inbound>\n", "</policies>\n")),
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
// API Management API Operation Policy can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/apiOperationPolicy:ApiOperationPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1/policies/policy
// ```
type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// The operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent pulumi.StringOutput `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrOutput `pulumi:"xmlLink"`
}

// NewApiOperationPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiOperationPolicy(ctx *pulumi.Context,
	name string, args *ApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azure:apimanagement/apiOperationPolicy:ApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperationPolicy gets an existing ApiOperationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationPolicyState, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	var resource ApiOperationPolicy
	err := ctx.ReadResource("azure:apimanagement/apiOperationPolicy:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationPolicy resources.
type apiOperationPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName *string `pulumi:"apiName"`
	// The operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId *string `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

type ApiOperationPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringPtrInput
	// The operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName string `pulumi:"apiName"`
	// The operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

// The set of arguments for constructing a ApiOperationPolicy resource.
type ApiOperationPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The ID of the API Management API Operation within the API Management Service. Changing this forces a new resource to be created.
	ApiName pulumi.StringInput
	// The operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}

type ApiOperationPolicyInput interface {
	pulumi.Input

	ToApiOperationPolicyOutput() ApiOperationPolicyOutput
	ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput
}

func (*ApiOperationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return i.ToApiOperationPolicyOutputWithContext(context.Background())
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPolicyOutput)
}

// ApiOperationPolicyArrayInput is an input type that accepts ApiOperationPolicyArray and ApiOperationPolicyArrayOutput values.
// You can construct a concrete instance of `ApiOperationPolicyArrayInput` via:
//
//          ApiOperationPolicyArray{ ApiOperationPolicyArgs{...} }
type ApiOperationPolicyArrayInput interface {
	pulumi.Input

	ToApiOperationPolicyArrayOutput() ApiOperationPolicyArrayOutput
	ToApiOperationPolicyArrayOutputWithContext(context.Context) ApiOperationPolicyArrayOutput
}

type ApiOperationPolicyArray []ApiOperationPolicyInput

func (ApiOperationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiOperationPolicy)(nil)).Elem()
}

func (i ApiOperationPolicyArray) ToApiOperationPolicyArrayOutput() ApiOperationPolicyArrayOutput {
	return i.ToApiOperationPolicyArrayOutputWithContext(context.Background())
}

func (i ApiOperationPolicyArray) ToApiOperationPolicyArrayOutputWithContext(ctx context.Context) ApiOperationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPolicyArrayOutput)
}

// ApiOperationPolicyMapInput is an input type that accepts ApiOperationPolicyMap and ApiOperationPolicyMapOutput values.
// You can construct a concrete instance of `ApiOperationPolicyMapInput` via:
//
//          ApiOperationPolicyMap{ "key": ApiOperationPolicyArgs{...} }
type ApiOperationPolicyMapInput interface {
	pulumi.Input

	ToApiOperationPolicyMapOutput() ApiOperationPolicyMapOutput
	ToApiOperationPolicyMapOutputWithContext(context.Context) ApiOperationPolicyMapOutput
}

type ApiOperationPolicyMap map[string]ApiOperationPolicyInput

func (ApiOperationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiOperationPolicy)(nil)).Elem()
}

func (i ApiOperationPolicyMap) ToApiOperationPolicyMapOutput() ApiOperationPolicyMapOutput {
	return i.ToApiOperationPolicyMapOutputWithContext(context.Background())
}

func (i ApiOperationPolicyMap) ToApiOperationPolicyMapOutputWithContext(ctx context.Context) ApiOperationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPolicyMapOutput)
}

type ApiOperationPolicyOutput struct{ *pulumi.OutputState }

func (ApiOperationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return o
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return o
}

type ApiOperationPolicyArrayOutput struct{ *pulumi.OutputState }

func (ApiOperationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiOperationPolicy)(nil)).Elem()
}

func (o ApiOperationPolicyArrayOutput) ToApiOperationPolicyArrayOutput() ApiOperationPolicyArrayOutput {
	return o
}

func (o ApiOperationPolicyArrayOutput) ToApiOperationPolicyArrayOutputWithContext(ctx context.Context) ApiOperationPolicyArrayOutput {
	return o
}

func (o ApiOperationPolicyArrayOutput) Index(i pulumi.IntInput) ApiOperationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiOperationPolicy {
		return vs[0].([]*ApiOperationPolicy)[vs[1].(int)]
	}).(ApiOperationPolicyOutput)
}

type ApiOperationPolicyMapOutput struct{ *pulumi.OutputState }

func (ApiOperationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiOperationPolicy)(nil)).Elem()
}

func (o ApiOperationPolicyMapOutput) ToApiOperationPolicyMapOutput() ApiOperationPolicyMapOutput {
	return o
}

func (o ApiOperationPolicyMapOutput) ToApiOperationPolicyMapOutputWithContext(ctx context.Context) ApiOperationPolicyMapOutput {
	return o
}

func (o ApiOperationPolicyMapOutput) MapIndex(k pulumi.StringInput) ApiOperationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiOperationPolicy {
		return vs[0].(map[string]*ApiOperationPolicy)[vs[1].(string)]
	}).(ApiOperationPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationPolicyInput)(nil)).Elem(), &ApiOperationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationPolicyArrayInput)(nil)).Elem(), ApiOperationPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiOperationPolicyMapInput)(nil)).Elem(), ApiOperationPolicyMap{})
	pulumi.RegisterOutputType(ApiOperationPolicyOutput{})
	pulumi.RegisterOutputType(ApiOperationPolicyArrayOutput{})
	pulumi.RegisterOutputType(ApiOperationPolicyMapOutput{})
}
