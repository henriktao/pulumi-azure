// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service certificate.
//
// ## Import
//
// App Service Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/certificate:Certificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/certificate1
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// The ID of the associated App Service plan. Must be specified when the certificate is used inside an App Service Environment hosted App Service. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringPtrOutput `pulumi:"appServicePlanId"`
	// The expiration date for the certificate.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// The friendly name of the certificate.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// List of host names the certificate applies to.
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// The ID of the the App Service Environment where the certificate is in use.
	//
	// Deprecated: This property has been deprecated and replaced with `app_service_plan_id`
	HostingEnvironmentProfileId pulumi.StringOutput `pulumi:"hostingEnvironmentProfileId"`
	// The issue date for the certificate.
	IssueDate pulumi.StringOutput `pulumi:"issueDate"`
	// The name of the certificate issuer.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The ID of the Key Vault secret. Changing this forces a new resource to be created.
	KeyVaultSecretId pulumi.StringPtrOutput `pulumi:"keyVaultSecretId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password to access the certificate's private key. Changing this forces a new resource to be created.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
	PfxBlob pulumi.StringPtrOutput `pulumi:"pfxBlob"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The subject name of the certificate.
	SubjectName pulumi.StringOutput    `pulumi:"subjectName"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	// The thumbprint for the certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:appservice/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure:appservice/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The ID of the associated App Service plan. Must be specified when the certificate is used inside an App Service Environment hosted App Service. Changing this forces a new resource to be created.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// The expiration date for the certificate.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The friendly name of the certificate.
	FriendlyName *string `pulumi:"friendlyName"`
	// List of host names the certificate applies to.
	HostNames []string `pulumi:"hostNames"`
	// The ID of the the App Service Environment where the certificate is in use.
	//
	// Deprecated: This property has been deprecated and replaced with `app_service_plan_id`
	HostingEnvironmentProfileId *string `pulumi:"hostingEnvironmentProfileId"`
	// The issue date for the certificate.
	IssueDate *string `pulumi:"issueDate"`
	// The name of the certificate issuer.
	Issuer *string `pulumi:"issuer"`
	// The ID of the Key Vault secret. Changing this forces a new resource to be created.
	KeyVaultSecretId *string `pulumi:"keyVaultSecretId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password to access the certificate's private key. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
	PfxBlob *string `pulumi:"pfxBlob"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The subject name of the certificate.
	SubjectName *string           `pulumi:"subjectName"`
	Tags        map[string]string `pulumi:"tags"`
	// The thumbprint for the certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

type CertificateState struct {
	// The ID of the associated App Service plan. Must be specified when the certificate is used inside an App Service Environment hosted App Service. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringPtrInput
	// The expiration date for the certificate.
	ExpirationDate pulumi.StringPtrInput
	// The friendly name of the certificate.
	FriendlyName pulumi.StringPtrInput
	// List of host names the certificate applies to.
	HostNames pulumi.StringArrayInput
	// The ID of the the App Service Environment where the certificate is in use.
	//
	// Deprecated: This property has been deprecated and replaced with `app_service_plan_id`
	HostingEnvironmentProfileId pulumi.StringPtrInput
	// The issue date for the certificate.
	IssueDate pulumi.StringPtrInput
	// The name of the certificate issuer.
	Issuer pulumi.StringPtrInput
	// The ID of the Key Vault secret. Changing this forces a new resource to be created.
	KeyVaultSecretId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password to access the certificate's private key. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
	PfxBlob pulumi.StringPtrInput
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The subject name of the certificate.
	SubjectName pulumi.StringPtrInput
	Tags        pulumi.StringMapInput
	// The thumbprint for the certificate.
	Thumbprint pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The ID of the associated App Service plan. Must be specified when the certificate is used inside an App Service Environment hosted App Service. Changing this forces a new resource to be created.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// The ID of the the App Service Environment where the certificate is in use.
	//
	// Deprecated: This property has been deprecated and replaced with `app_service_plan_id`
	HostingEnvironmentProfileId *string `pulumi:"hostingEnvironmentProfileId"`
	// The ID of the Key Vault secret. Changing this forces a new resource to be created.
	KeyVaultSecretId *string `pulumi:"keyVaultSecretId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password to access the certificate's private key. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
	PfxBlob *string `pulumi:"pfxBlob"`
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The ID of the associated App Service plan. Must be specified when the certificate is used inside an App Service Environment hosted App Service. Changing this forces a new resource to be created.
	AppServicePlanId pulumi.StringPtrInput
	// The ID of the the App Service Environment where the certificate is in use.
	//
	// Deprecated: This property has been deprecated and replaced with `app_service_plan_id`
	HostingEnvironmentProfileId pulumi.StringPtrInput
	// The ID of the Key Vault secret. Changing this forces a new resource to be created.
	KeyVaultSecretId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password to access the certificate's private key. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The base64-encoded contents of the certificate. Changing this forces a new resource to be created.
	PfxBlob pulumi.StringPtrInput
	// The name of the resource group in which to create the certificate. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

func (i *Certificate) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

type CertificatePtrInput interface {
	pulumi.Input

	ToCertificatePtrOutput() CertificatePtrOutput
	ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput
}

type certificatePtrType CertificateArgs

func (*certificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (i *certificatePtrType) ToCertificatePtrOutput() CertificatePtrOutput {
	return i.ToCertificatePtrOutputWithContext(context.Background())
}

func (i *certificatePtrType) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePtrOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Certificate)(nil))
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Certificate)(nil))
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o.ToCertificatePtrOutputWithContext(context.Background())
}

func (o CertificateOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o.ApplyT(func(v Certificate) *Certificate {
		return &v
	}).(CertificatePtrOutput)
}

type CertificatePtrOutput struct {
	*pulumi.OutputState
}

func (CertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil))
}

func (o CertificatePtrOutput) ToCertificatePtrOutput() CertificatePtrOutput {
	return o
}

func (o CertificatePtrOutput) ToCertificatePtrOutputWithContext(ctx context.Context) CertificatePtrOutput {
	return o
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil))
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Certificate)(nil))
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].(map[string]Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificatePtrOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
