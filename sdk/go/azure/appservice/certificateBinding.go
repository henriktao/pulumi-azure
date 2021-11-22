// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service Certificate Binding.
//
// ## Import
//
// App Service Certificate Bindings can be imported using the `hostname_binding_id` and the `app_service_certificate_id` , e.g.
//
// ```sh
//  $ pulumi import azure:appservice/certificateBinding:CertificateBinding example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/mywebsite.com"
// ```
type CertificateBinding struct {
	pulumi.CustomResourceState

	// The name of the App Service to which the certificate was bound.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// The hostname of the bound certificate.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
	HostnameBindingId pulumi.StringOutput `pulumi:"hostnameBindingId"`
	// The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
	SslState pulumi.StringOutput `pulumi:"sslState"`
	// The certificate thumbprint.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
}

// NewCertificateBinding registers a new resource with the given unique name, arguments, and options.
func NewCertificateBinding(ctx *pulumi.Context,
	name string, args *CertificateBindingArgs, opts ...pulumi.ResourceOption) (*CertificateBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateId'")
	}
	if args.HostnameBindingId == nil {
		return nil, errors.New("invalid value for required argument 'HostnameBindingId'")
	}
	if args.SslState == nil {
		return nil, errors.New("invalid value for required argument 'SslState'")
	}
	var resource CertificateBinding
	err := ctx.RegisterResource("azure:appservice/certificateBinding:CertificateBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateBinding gets an existing CertificateBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateBindingState, opts ...pulumi.ResourceOption) (*CertificateBinding, error) {
	var resource CertificateBinding
	err := ctx.ReadResource("azure:appservice/certificateBinding:CertificateBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateBinding resources.
type certificateBindingState struct {
	// The name of the App Service to which the certificate was bound.
	AppServiceName *string `pulumi:"appServiceName"`
	// The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
	CertificateId *string `pulumi:"certificateId"`
	// The hostname of the bound certificate.
	Hostname *string `pulumi:"hostname"`
	// The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
	HostnameBindingId *string `pulumi:"hostnameBindingId"`
	// The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
	SslState *string `pulumi:"sslState"`
	// The certificate thumbprint.
	Thumbprint *string `pulumi:"thumbprint"`
}

type CertificateBindingState struct {
	// The name of the App Service to which the certificate was bound.
	AppServiceName pulumi.StringPtrInput
	// The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
	CertificateId pulumi.StringPtrInput
	// The hostname of the bound certificate.
	Hostname pulumi.StringPtrInput
	// The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
	HostnameBindingId pulumi.StringPtrInput
	// The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
	SslState pulumi.StringPtrInput
	// The certificate thumbprint.
	Thumbprint pulumi.StringPtrInput
}

func (CertificateBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateBindingState)(nil)).Elem()
}

type certificateBindingArgs struct {
	// The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
	CertificateId string `pulumi:"certificateId"`
	// The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
	HostnameBindingId string `pulumi:"hostnameBindingId"`
	// The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
	SslState string `pulumi:"sslState"`
}

// The set of arguments for constructing a CertificateBinding resource.
type CertificateBindingArgs struct {
	// The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
	CertificateId pulumi.StringInput
	// The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
	HostnameBindingId pulumi.StringInput
	// The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
	SslState pulumi.StringInput
}

func (CertificateBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateBindingArgs)(nil)).Elem()
}

type CertificateBindingInput interface {
	pulumi.Input

	ToCertificateBindingOutput() CertificateBindingOutput
	ToCertificateBindingOutputWithContext(ctx context.Context) CertificateBindingOutput
}

func (*CertificateBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateBinding)(nil))
}

func (i *CertificateBinding) ToCertificateBindingOutput() CertificateBindingOutput {
	return i.ToCertificateBindingOutputWithContext(context.Background())
}

func (i *CertificateBinding) ToCertificateBindingOutputWithContext(ctx context.Context) CertificateBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateBindingOutput)
}

func (i *CertificateBinding) ToCertificateBindingPtrOutput() CertificateBindingPtrOutput {
	return i.ToCertificateBindingPtrOutputWithContext(context.Background())
}

func (i *CertificateBinding) ToCertificateBindingPtrOutputWithContext(ctx context.Context) CertificateBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateBindingPtrOutput)
}

type CertificateBindingPtrInput interface {
	pulumi.Input

	ToCertificateBindingPtrOutput() CertificateBindingPtrOutput
	ToCertificateBindingPtrOutputWithContext(ctx context.Context) CertificateBindingPtrOutput
}

type certificateBindingPtrType CertificateBindingArgs

func (*certificateBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateBinding)(nil))
}

func (i *certificateBindingPtrType) ToCertificateBindingPtrOutput() CertificateBindingPtrOutput {
	return i.ToCertificateBindingPtrOutputWithContext(context.Background())
}

func (i *certificateBindingPtrType) ToCertificateBindingPtrOutputWithContext(ctx context.Context) CertificateBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateBindingPtrOutput)
}

// CertificateBindingArrayInput is an input type that accepts CertificateBindingArray and CertificateBindingArrayOutput values.
// You can construct a concrete instance of `CertificateBindingArrayInput` via:
//
//          CertificateBindingArray{ CertificateBindingArgs{...} }
type CertificateBindingArrayInput interface {
	pulumi.Input

	ToCertificateBindingArrayOutput() CertificateBindingArrayOutput
	ToCertificateBindingArrayOutputWithContext(context.Context) CertificateBindingArrayOutput
}

type CertificateBindingArray []CertificateBindingInput

func (CertificateBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateBinding)(nil)).Elem()
}

func (i CertificateBindingArray) ToCertificateBindingArrayOutput() CertificateBindingArrayOutput {
	return i.ToCertificateBindingArrayOutputWithContext(context.Background())
}

func (i CertificateBindingArray) ToCertificateBindingArrayOutputWithContext(ctx context.Context) CertificateBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateBindingArrayOutput)
}

// CertificateBindingMapInput is an input type that accepts CertificateBindingMap and CertificateBindingMapOutput values.
// You can construct a concrete instance of `CertificateBindingMapInput` via:
//
//          CertificateBindingMap{ "key": CertificateBindingArgs{...} }
type CertificateBindingMapInput interface {
	pulumi.Input

	ToCertificateBindingMapOutput() CertificateBindingMapOutput
	ToCertificateBindingMapOutputWithContext(context.Context) CertificateBindingMapOutput
}

type CertificateBindingMap map[string]CertificateBindingInput

func (CertificateBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateBinding)(nil)).Elem()
}

func (i CertificateBindingMap) ToCertificateBindingMapOutput() CertificateBindingMapOutput {
	return i.ToCertificateBindingMapOutputWithContext(context.Background())
}

func (i CertificateBindingMap) ToCertificateBindingMapOutputWithContext(ctx context.Context) CertificateBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateBindingMapOutput)
}

type CertificateBindingOutput struct{ *pulumi.OutputState }

func (CertificateBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateBinding)(nil))
}

func (o CertificateBindingOutput) ToCertificateBindingOutput() CertificateBindingOutput {
	return o
}

func (o CertificateBindingOutput) ToCertificateBindingOutputWithContext(ctx context.Context) CertificateBindingOutput {
	return o
}

func (o CertificateBindingOutput) ToCertificateBindingPtrOutput() CertificateBindingPtrOutput {
	return o.ToCertificateBindingPtrOutputWithContext(context.Background())
}

func (o CertificateBindingOutput) ToCertificateBindingPtrOutputWithContext(ctx context.Context) CertificateBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateBinding) *CertificateBinding {
		return &v
	}).(CertificateBindingPtrOutput)
}

type CertificateBindingPtrOutput struct{ *pulumi.OutputState }

func (CertificateBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateBinding)(nil))
}

func (o CertificateBindingPtrOutput) ToCertificateBindingPtrOutput() CertificateBindingPtrOutput {
	return o
}

func (o CertificateBindingPtrOutput) ToCertificateBindingPtrOutputWithContext(ctx context.Context) CertificateBindingPtrOutput {
	return o
}

func (o CertificateBindingPtrOutput) Elem() CertificateBindingOutput {
	return o.ApplyT(func(v *CertificateBinding) CertificateBinding {
		if v != nil {
			return *v
		}
		var ret CertificateBinding
		return ret
	}).(CertificateBindingOutput)
}

type CertificateBindingArrayOutput struct{ *pulumi.OutputState }

func (CertificateBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateBinding)(nil))
}

func (o CertificateBindingArrayOutput) ToCertificateBindingArrayOutput() CertificateBindingArrayOutput {
	return o
}

func (o CertificateBindingArrayOutput) ToCertificateBindingArrayOutputWithContext(ctx context.Context) CertificateBindingArrayOutput {
	return o
}

func (o CertificateBindingArrayOutput) Index(i pulumi.IntInput) CertificateBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateBinding {
		return vs[0].([]CertificateBinding)[vs[1].(int)]
	}).(CertificateBindingOutput)
}

type CertificateBindingMapOutput struct{ *pulumi.OutputState }

func (CertificateBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CertificateBinding)(nil))
}

func (o CertificateBindingMapOutput) ToCertificateBindingMapOutput() CertificateBindingMapOutput {
	return o
}

func (o CertificateBindingMapOutput) ToCertificateBindingMapOutputWithContext(ctx context.Context) CertificateBindingMapOutput {
	return o
}

func (o CertificateBindingMapOutput) MapIndex(k pulumi.StringInput) CertificateBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CertificateBinding {
		return vs[0].(map[string]CertificateBinding)[vs[1].(string)]
	}).(CertificateBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateBindingInput)(nil)).Elem(), &CertificateBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateBindingPtrInput)(nil)).Elem(), &CertificateBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateBindingArrayInput)(nil)).Elem(), CertificateBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateBindingMapInput)(nil)).Elem(), CertificateBindingMap{})
	pulumi.RegisterOutputType(CertificateBindingOutput{})
	pulumi.RegisterOutputType(CertificateBindingPtrOutput{})
	pulumi.RegisterOutputType(CertificateBindingArrayOutput{})
	pulumi.RegisterOutputType(CertificateBindingMapOutput{})
}
