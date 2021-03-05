// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an API Version Set within an API Management Service.
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
// 		_, err = apimanagement.NewApiVersionSet(ctx, "exampleApiVersionSet", &apimanagement.ApiVersionSetArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			DisplayName:       pulumi.String("ExampleAPIVersionSet"),
// 			VersioningScheme:  pulumi.String("Segment"),
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
// API Version Set can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/apiVersionSet:ApiVersionSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apiVersionSets/set1
// ```
type ApiVersionSet struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrOutput `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrOutput `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringOutput `pulumi:"versioningScheme"`
}

// NewApiVersionSet registers a new resource with the given unique name, arguments, and options.
func NewApiVersionSet(ctx *pulumi.Context,
	name string, args *ApiVersionSetArgs, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
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
	if args.VersioningScheme == nil {
		return nil, errors.New("invalid value for required argument 'VersioningScheme'")
	}
	var resource ApiVersionSet
	err := ctx.RegisterResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiVersionSet gets an existing ApiVersionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiVersionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiVersionSetState, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
	var resource ApiVersionSet
	err := ctx.ReadResource("azure:apimanagement/apiVersionSet:ApiVersionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiVersionSet resources.
type apiVersionSetState struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description *string `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName *string `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme *string `pulumi:"versioningScheme"`
}

type ApiVersionSetState struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The description of API Version Set.
	Description pulumi.StringPtrInput
	// The display name of this API Version Set.
	DisplayName pulumi.StringPtrInput
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrInput
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrInput
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringPtrInput
}

func (ApiVersionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetState)(nil)).Elem()
}

type apiVersionSetArgs struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of API Version Set.
	Description *string `pulumi:"description"`
	// The display name of this API Version Set.
	DisplayName string `pulumi:"displayName"`
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme string `pulumi:"versioningScheme"`
}

// The set of arguments for constructing a ApiVersionSet resource.
type ApiVersionSetArgs struct {
	// The name of the API Management Service in which the API Version Set should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The description of API Version Set.
	Description pulumi.StringPtrInput
	// The display name of this API Version Set.
	DisplayName pulumi.StringInput
	// The name of the API Version Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the parent API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName pulumi.StringPtrInput
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName pulumi.StringPtrInput
	// Specifies where in an Inbound HTTP Request that the API Version should be read from. Possible values are `Header`, `Query` and `Segment`.
	VersioningScheme pulumi.StringInput
}

func (ApiVersionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetArgs)(nil)).Elem()
}

type ApiVersionSetInput interface {
	pulumi.Input

	ToApiVersionSetOutput() ApiVersionSetOutput
	ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput
}

func (*ApiVersionSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSet)(nil))
}

func (i *ApiVersionSet) ToApiVersionSetOutput() ApiVersionSetOutput {
	return i.ToApiVersionSetOutputWithContext(context.Background())
}

func (i *ApiVersionSet) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetOutput)
}

func (i *ApiVersionSet) ToApiVersionSetPtrOutput() ApiVersionSetPtrOutput {
	return i.ToApiVersionSetPtrOutputWithContext(context.Background())
}

func (i *ApiVersionSet) ToApiVersionSetPtrOutputWithContext(ctx context.Context) ApiVersionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetPtrOutput)
}

type ApiVersionSetPtrInput interface {
	pulumi.Input

	ToApiVersionSetPtrOutput() ApiVersionSetPtrOutput
	ToApiVersionSetPtrOutputWithContext(ctx context.Context) ApiVersionSetPtrOutput
}

type apiVersionSetPtrType ApiVersionSetArgs

func (*apiVersionSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSet)(nil))
}

func (i *apiVersionSetPtrType) ToApiVersionSetPtrOutput() ApiVersionSetPtrOutput {
	return i.ToApiVersionSetPtrOutputWithContext(context.Background())
}

func (i *apiVersionSetPtrType) ToApiVersionSetPtrOutputWithContext(ctx context.Context) ApiVersionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetPtrOutput)
}

// ApiVersionSetArrayInput is an input type that accepts ApiVersionSetArray and ApiVersionSetArrayOutput values.
// You can construct a concrete instance of `ApiVersionSetArrayInput` via:
//
//          ApiVersionSetArray{ ApiVersionSetArgs{...} }
type ApiVersionSetArrayInput interface {
	pulumi.Input

	ToApiVersionSetArrayOutput() ApiVersionSetArrayOutput
	ToApiVersionSetArrayOutputWithContext(context.Context) ApiVersionSetArrayOutput
}

type ApiVersionSetArray []ApiVersionSetInput

func (ApiVersionSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiVersionSet)(nil))
}

func (i ApiVersionSetArray) ToApiVersionSetArrayOutput() ApiVersionSetArrayOutput {
	return i.ToApiVersionSetArrayOutputWithContext(context.Background())
}

func (i ApiVersionSetArray) ToApiVersionSetArrayOutputWithContext(ctx context.Context) ApiVersionSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetArrayOutput)
}

// ApiVersionSetMapInput is an input type that accepts ApiVersionSetMap and ApiVersionSetMapOutput values.
// You can construct a concrete instance of `ApiVersionSetMapInput` via:
//
//          ApiVersionSetMap{ "key": ApiVersionSetArgs{...} }
type ApiVersionSetMapInput interface {
	pulumi.Input

	ToApiVersionSetMapOutput() ApiVersionSetMapOutput
	ToApiVersionSetMapOutputWithContext(context.Context) ApiVersionSetMapOutput
}

type ApiVersionSetMap map[string]ApiVersionSetInput

func (ApiVersionSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiVersionSet)(nil))
}

func (i ApiVersionSetMap) ToApiVersionSetMapOutput() ApiVersionSetMapOutput {
	return i.ToApiVersionSetMapOutputWithContext(context.Background())
}

func (i ApiVersionSetMap) ToApiVersionSetMapOutputWithContext(ctx context.Context) ApiVersionSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetMapOutput)
}

type ApiVersionSetOutput struct {
	*pulumi.OutputState
}

func (ApiVersionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSet)(nil))
}

func (o ApiVersionSetOutput) ToApiVersionSetOutput() ApiVersionSetOutput {
	return o
}

func (o ApiVersionSetOutput) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return o
}

func (o ApiVersionSetOutput) ToApiVersionSetPtrOutput() ApiVersionSetPtrOutput {
	return o.ToApiVersionSetPtrOutputWithContext(context.Background())
}

func (o ApiVersionSetOutput) ToApiVersionSetPtrOutputWithContext(ctx context.Context) ApiVersionSetPtrOutput {
	return o.ApplyT(func(v ApiVersionSet) *ApiVersionSet {
		return &v
	}).(ApiVersionSetPtrOutput)
}

type ApiVersionSetPtrOutput struct {
	*pulumi.OutputState
}

func (ApiVersionSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSet)(nil))
}

func (o ApiVersionSetPtrOutput) ToApiVersionSetPtrOutput() ApiVersionSetPtrOutput {
	return o
}

func (o ApiVersionSetPtrOutput) ToApiVersionSetPtrOutputWithContext(ctx context.Context) ApiVersionSetPtrOutput {
	return o
}

type ApiVersionSetArrayOutput struct{ *pulumi.OutputState }

func (ApiVersionSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiVersionSet)(nil))
}

func (o ApiVersionSetArrayOutput) ToApiVersionSetArrayOutput() ApiVersionSetArrayOutput {
	return o
}

func (o ApiVersionSetArrayOutput) ToApiVersionSetArrayOutputWithContext(ctx context.Context) ApiVersionSetArrayOutput {
	return o
}

func (o ApiVersionSetArrayOutput) Index(i pulumi.IntInput) ApiVersionSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiVersionSet {
		return vs[0].([]ApiVersionSet)[vs[1].(int)]
	}).(ApiVersionSetOutput)
}

type ApiVersionSetMapOutput struct{ *pulumi.OutputState }

func (ApiVersionSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiVersionSet)(nil))
}

func (o ApiVersionSetMapOutput) ToApiVersionSetMapOutput() ApiVersionSetMapOutput {
	return o
}

func (o ApiVersionSetMapOutput) ToApiVersionSetMapOutputWithContext(ctx context.Context) ApiVersionSetMapOutput {
	return o
}

func (o ApiVersionSetMapOutput) MapIndex(k pulumi.StringInput) ApiVersionSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiVersionSet {
		return vs[0].(map[string]ApiVersionSet)[vs[1].(string)]
	}).(ApiVersionSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiVersionSetOutput{})
	pulumi.RegisterOutputType(ApiVersionSetPtrOutput{})
	pulumi.RegisterOutputType(ApiVersionSetArrayOutput{})
	pulumi.RegisterOutputType(ApiVersionSetMapOutput{})
}
