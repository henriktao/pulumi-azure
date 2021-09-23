// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a backend within an API Management Service.
//
// ## Import
//
// API Management backends can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/backend:Backend example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/backends/backend1
// ```
type Backend struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// A `credentials` block as documented below.
	Credentials BackendCredentialsPtrOutput `pulumi:"credentials"`
	// The description of the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// A `proxy` block as documented below.
	Proxy BackendProxyPtrOutput `pulumi:"proxy"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster BackendServiceFabricClusterPtrOutput `pulumi:"serviceFabricCluster"`
	// The title of the backend.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// A `tls` block as documented below.
	Tls BackendTlsPtrOutput `pulumi:"tls"`
	// The URL of the backend host.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource Backend
	err := ctx.RegisterResource("azure:apimanagement/backend:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("azure:apimanagement/backend:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// A `credentials` block as documented below.
	Credentials *BackendCredentials `pulumi:"credentials"`
	// The description of the backend.
	Description *string `pulumi:"description"`
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol *string `pulumi:"protocol"`
	// A `proxy` block as documented below.
	Proxy *BackendProxy `pulumi:"proxy"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId *string `pulumi:"resourceId"`
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster *BackendServiceFabricCluster `pulumi:"serviceFabricCluster"`
	// The title of the backend.
	Title *string `pulumi:"title"`
	// A `tls` block as documented below.
	Tls *BackendTls `pulumi:"tls"`
	// The URL of the backend host.
	Url *string `pulumi:"url"`
}

type BackendState struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// A `credentials` block as documented below.
	Credentials BackendCredentialsPtrInput
	// The description of the backend.
	Description pulumi.StringPtrInput
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol pulumi.StringPtrInput
	// A `proxy` block as documented below.
	Proxy BackendProxyPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId pulumi.StringPtrInput
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster BackendServiceFabricClusterPtrInput
	// The title of the backend.
	Title pulumi.StringPtrInput
	// A `tls` block as documented below.
	Tls BackendTlsPtrInput
	// The URL of the backend host.
	Url pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// A `credentials` block as documented below.
	Credentials *BackendCredentials `pulumi:"credentials"`
	// The description of the backend.
	Description *string `pulumi:"description"`
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol string `pulumi:"protocol"`
	// A `proxy` block as documented below.
	Proxy *BackendProxy `pulumi:"proxy"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId *string `pulumi:"resourceId"`
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster *BackendServiceFabricCluster `pulumi:"serviceFabricCluster"`
	// The title of the backend.
	Title *string `pulumi:"title"`
	// A `tls` block as documented below.
	Tls *BackendTls `pulumi:"tls"`
	// The URL of the backend host.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// A `credentials` block as documented below.
	Credentials BackendCredentialsPtrInput
	// The description of the backend.
	Description pulumi.StringPtrInput
	// The name of the API Management backend. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The protocol used by the backend host. Possible values are `http` or `soap`.
	Protocol pulumi.StringInput
	// A `proxy` block as documented below.
	Proxy BackendProxyPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
	ResourceId pulumi.StringPtrInput
	// A `serviceFabricCluster` block as documented below.
	ServiceFabricCluster BackendServiceFabricClusterPtrInput
	// The title of the backend.
	Title pulumi.StringPtrInput
	// A `tls` block as documented below.
	Tls BackendTlsPtrInput
	// The URL of the backend host.
	Url pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (*Backend) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil))
}

func (i *Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i *Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

func (i *Backend) ToBackendPtrOutput() BackendPtrOutput {
	return i.ToBackendPtrOutputWithContext(context.Background())
}

func (i *Backend) ToBackendPtrOutputWithContext(ctx context.Context) BackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPtrOutput)
}

type BackendPtrInput interface {
	pulumi.Input

	ToBackendPtrOutput() BackendPtrOutput
	ToBackendPtrOutputWithContext(ctx context.Context) BackendPtrOutput
}

type backendPtrType BackendArgs

func (*backendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil))
}

func (i *backendPtrType) ToBackendPtrOutput() BackendPtrOutput {
	return i.ToBackendPtrOutputWithContext(context.Background())
}

func (i *backendPtrType) ToBackendPtrOutputWithContext(ctx context.Context) BackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPtrOutput)
}

// BackendArrayInput is an input type that accepts BackendArray and BackendArrayOutput values.
// You can construct a concrete instance of `BackendArrayInput` via:
//
//          BackendArray{ BackendArgs{...} }
type BackendArrayInput interface {
	pulumi.Input

	ToBackendArrayOutput() BackendArrayOutput
	ToBackendArrayOutputWithContext(context.Context) BackendArrayOutput
}

type BackendArray []BackendInput

func (BackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backend)(nil)).Elem()
}

func (i BackendArray) ToBackendArrayOutput() BackendArrayOutput {
	return i.ToBackendArrayOutputWithContext(context.Background())
}

func (i BackendArray) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendArrayOutput)
}

// BackendMapInput is an input type that accepts BackendMap and BackendMapOutput values.
// You can construct a concrete instance of `BackendMapInput` via:
//
//          BackendMap{ "key": BackendArgs{...} }
type BackendMapInput interface {
	pulumi.Input

	ToBackendMapOutput() BackendMapOutput
	ToBackendMapOutputWithContext(context.Context) BackendMapOutput
}

type BackendMap map[string]BackendInput

func (BackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backend)(nil)).Elem()
}

func (i BackendMap) ToBackendMapOutput() BackendMapOutput {
	return i.ToBackendMapOutputWithContext(context.Background())
}

func (i BackendMap) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendMapOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil))
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

func (o BackendOutput) ToBackendPtrOutput() BackendPtrOutput {
	return o.ToBackendPtrOutputWithContext(context.Background())
}

func (o BackendOutput) ToBackendPtrOutputWithContext(ctx context.Context) BackendPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Backend) *Backend {
		return &v
	}).(BackendPtrOutput)
}

type BackendPtrOutput struct{ *pulumi.OutputState }

func (BackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil))
}

func (o BackendPtrOutput) ToBackendPtrOutput() BackendPtrOutput {
	return o
}

func (o BackendPtrOutput) ToBackendPtrOutputWithContext(ctx context.Context) BackendPtrOutput {
	return o
}

func (o BackendPtrOutput) Elem() BackendOutput {
	return o.ApplyT(func(v *Backend) Backend {
		if v != nil {
			return *v
		}
		var ret Backend
		return ret
	}).(BackendOutput)
}

type BackendArrayOutput struct{ *pulumi.OutputState }

func (BackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backend)(nil))
}

func (o BackendArrayOutput) ToBackendArrayOutput() BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) Index(i pulumi.IntInput) BackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Backend {
		return vs[0].([]Backend)[vs[1].(int)]
	}).(BackendOutput)
}

type BackendMapOutput struct{ *pulumi.OutputState }

func (BackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Backend)(nil))
}

func (o BackendMapOutput) ToBackendMapOutput() BackendMapOutput {
	return o
}

func (o BackendMapOutput) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return o
}

func (o BackendMapOutput) MapIndex(k pulumi.StringInput) BackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Backend {
		return vs[0].(map[string]Backend)[vs[1].(string)]
	}).(BackendOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendOutput{})
	pulumi.RegisterOutputType(BackendPtrOutput{})
	pulumi.RegisterOutputType(BackendArrayOutput{})
	pulumi.RegisterOutputType(BackendMapOutput{})
}
