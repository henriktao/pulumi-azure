// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package webpubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Web Pubsub Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/webpubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("east us"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = webpubsub.NewService(ctx, "exampleService", &webpubsub.ServiceArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Sku:                        pulumi.String("Standard_S1"),
// 			Capacity:                   pulumi.Int(1),
// 			PublicNetworkAccessEnabled: pulumi.Bool(false),
// 			LiveTrace: &webpubsub.ServiceLiveTraceArgs{
// 				Enabled:                 pulumi.Bool(true),
// 				MessagingLogsEnabled:    pulumi.Bool(true),
// 				ConnectivityLogsEnabled: pulumi.Bool(false),
// 			},
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
// Web Pubsub services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:webpubsub/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
// ```
type Service struct {
	pulumi.CustomResourceState

	// Whether to enable AAD auth? Defaults to `true`.
	AadAuthEnabled pulumi.BoolPtrOutput `pulumi:"aadAuthEnabled"`
	// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
	// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
	Capacity   pulumi.IntPtrOutput `pulumi:"capacity"`
	ExternalIp pulumi.StringOutput `pulumi:"externalIp"`
	// The FQDN of the Web Pubsub service.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// A `liveTrace` block as defined below.
	LiveTrace ServiceLiveTracePtrOutput `pulumi:"liveTrace"`
	// Whether to enable local auth? Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrOutput `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
	// forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Web Pubsub service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary access key for the Web Pubsub service.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The primary connection string for the Web Pubsub service.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// Whether to enable public network access? Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The publicly accessible port of the Web Pubsub service which is designed for browser/client use.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The name of the resource group in which to create the Web Pubsub service. Changing
	// this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary access key for the Web Pubsub service.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The secondary connection string for the Web Pubsub service.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the Web Pubsub service which is designed for customer server side use.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Whether to request client certificate during TLS handshake? Defaults
	// to `false`.
	TlsClientCertEnabled pulumi.BoolPtrOutput `pulumi:"tlsClientCertEnabled"`
	Version              pulumi.StringOutput  `pulumi:"version"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:webpubsub/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:webpubsub/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Whether to enable AAD auth? Defaults to `true`.
	AadAuthEnabled *bool `pulumi:"aadAuthEnabled"`
	// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
	// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
	Capacity   *int    `pulumi:"capacity"`
	ExternalIp *string `pulumi:"externalIp"`
	// The FQDN of the Web Pubsub service.
	Hostname *string `pulumi:"hostname"`
	// A `liveTrace` block as defined below.
	LiveTrace *ServiceLiveTrace `pulumi:"liveTrace"`
	// Whether to enable local auth? Defaults to `true`.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
	// forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Web Pubsub service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary access key for the Web Pubsub service.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The primary connection string for the Web Pubsub service.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// Whether to enable public network access? Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The publicly accessible port of the Web Pubsub service which is designed for browser/client use.
	PublicPort *int `pulumi:"publicPort"`
	// The name of the resource group in which to create the Web Pubsub service. Changing
	// this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary access key for the Web Pubsub service.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The secondary connection string for the Web Pubsub service.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The publicly accessible port of the Web Pubsub service which is designed for customer server side use.
	ServerPort *int `pulumi:"serverPort"`
	// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to request client certificate during TLS handshake? Defaults
	// to `false`.
	TlsClientCertEnabled *bool   `pulumi:"tlsClientCertEnabled"`
	Version              *string `pulumi:"version"`
}

type ServiceState struct {
	// Whether to enable AAD auth? Defaults to `true`.
	AadAuthEnabled pulumi.BoolPtrInput
	// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
	// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
	Capacity   pulumi.IntPtrInput
	ExternalIp pulumi.StringPtrInput
	// The FQDN of the Web Pubsub service.
	Hostname pulumi.StringPtrInput
	// A `liveTrace` block as defined below.
	LiveTrace ServiceLiveTracePtrInput
	// Whether to enable local auth? Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
	// forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Web Pubsub service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary access key for the Web Pubsub service.
	PrimaryAccessKey pulumi.StringPtrInput
	// The primary connection string for the Web Pubsub service.
	PrimaryConnectionString pulumi.StringPtrInput
	// Whether to enable public network access? Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The publicly accessible port of the Web Pubsub service which is designed for browser/client use.
	PublicPort pulumi.IntPtrInput
	// The name of the resource group in which to create the Web Pubsub service. Changing
	// this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary access key for the Web Pubsub service.
	SecondaryAccessKey pulumi.StringPtrInput
	// The secondary connection string for the Web Pubsub service.
	SecondaryConnectionString pulumi.StringPtrInput
	// The publicly accessible port of the Web Pubsub service which is designed for customer server side use.
	ServerPort pulumi.IntPtrInput
	// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to request client certificate during TLS handshake? Defaults
	// to `false`.
	TlsClientCertEnabled pulumi.BoolPtrInput
	Version              pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Whether to enable AAD auth? Defaults to `true`.
	AadAuthEnabled *bool `pulumi:"aadAuthEnabled"`
	// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
	// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
	Capacity *int `pulumi:"capacity"`
	// A `liveTrace` block as defined below.
	LiveTrace *ServiceLiveTrace `pulumi:"liveTrace"`
	// Whether to enable local auth? Defaults to `true`.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
	// forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Web Pubsub service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Whether to enable public network access? Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which to create the Web Pubsub service. Changing
	// this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether to request client certificate during TLS handshake? Defaults
	// to `false`.
	TlsClientCertEnabled *bool `pulumi:"tlsClientCertEnabled"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Whether to enable AAD auth? Defaults to `true`.
	AadAuthEnabled pulumi.BoolPtrInput
	// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
	// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
	Capacity pulumi.IntPtrInput
	// A `liveTrace` block as defined below.
	LiveTrace ServiceLiveTracePtrInput
	// Whether to enable local auth? Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
	// forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Web Pubsub service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Whether to enable public network access? Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which to create the Web Pubsub service. Changing
	// this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Whether to request client certificate during TLS handshake? Defaults
	// to `false`.
	TlsClientCertEnabled pulumi.BoolPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
