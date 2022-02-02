// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a VPN Server Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVpnServerConfiguration(ctx, "test", &network.VpnServerConfigurationArgs{
// 			ResourceGroupName: example.Name,
// 			Location:          example.Location,
// 			VpnAuthenticationTypes: pulumi.String{
// 				"Certificate",
// 			},
// 			ClientRootCertificates: network.VpnServerConfigurationClientRootCertificateArray{
// 				&network.VpnServerConfigurationClientRootCertificateArgs{
// 					Name:           pulumi.String("DigiCert-Federated-ID-Root-CA"),
// 					PublicCertData: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "MIIDuzCCAqOgAwIBAgIQCHTZWCM+IlfFIRXIvyKSrjANBgkqhkiG9w0BAQsFADBn\n", "MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\n", "d3cuZGlnaWNlcnQuY29tMSYwJAYDVQQDEx1EaWdpQ2VydCBGZWRlcmF0ZWQgSUQg\n", "Um9vdCBDQTAeFw0xMzAxMTUxMjAwMDBaFw0zMzAxMTUxMjAwMDBaMGcxCzAJBgNV\n", "BAYTAlVTMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdp\n", "Y2VydC5jb20xJjAkBgNVBAMTHURpZ2lDZXJ0IEZlZGVyYXRlZCBJRCBSb290IENB\n", "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvAEB4pcCqnNNOWE6Ur5j\n", "QPUH+1y1F9KdHTRSza6k5iDlXq1kGS1qAkuKtw9JsiNRrjltmFnzMZRBbX8Tlfl8\n", "zAhBmb6dDduDGED01kBsTkgywYPxXVTKec0WxYEEF0oMn4wSYNl0lt2eJAKHXjNf\n", "GTwiibdP8CUR2ghSM2sUTI8Nt1Omfc4SMHhGhYD64uJMbX98THQ/4LMGuYegou+d\n", "GTiahfHtjn7AboSEknwAMJHCh5RlYZZ6B1O4QbKJ+34Q0eKgnI3X6Vc9u0zf6DH8\n", "Dk+4zQDYRRTqTnVO3VT8jzqDlCRuNtq6YvryOWN74/dq8LQhUnXHvFyrsdMaE1X2\n", "DwIDAQABo2MwYTAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV\n", "HQ4EFgQUGRdkFnbGt1EWjKwbUne+5OaZvRYwHwYDVR0jBBgwFoAUGRdkFnbGt1EW\n", "jKwbUne+5OaZvRYwDQYJKoZIhvcNAQELBQADggEBAHcqsHkrjpESqfuVTRiptJfP\n", "9JbdtWqRTmOf6uJi2c8YVqI6XlKXsD8C1dUUaaHKLUJzvKiazibVuBwMIT84AyqR\n", "QELn3e0BtgEymEygMU569b01ZPxoFSnNXc7qDZBDef8WfqAV/sxkTi8L9BkmFYfL\n", "uGLOhRJOFprPdoDIUBB+tmCl3oDcBy3vnUeOEioz8zAkprcb3GHwHAK+vHmmfgcn\n", "WsfMLH4JCLa/tRYL+Rw/N3ybCkDp00s0WUZ+AoDywSl0Q/ZEnNY0MsFiw6LyIdbq\n", "M/s/1JRtO3bDSzD9TazRVzn2oBqzSa8VgIo5C1nOnoAKJTlsClJKvIhnRlaLQqk=\n")),
// 				},
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
// VPN Server Configurations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/vpnServerConfiguration:VpnServerConfiguration config1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/vpnServerConfigurations/config1
// ```
type VpnServerConfiguration struct {
	pulumi.CustomResourceState

	// A `azureActiveDirectoryAuthentication` block as defined below.
	AzureActiveDirectoryAuthentications VpnServerConfigurationAzureActiveDirectoryAuthenticationArrayOutput `pulumi:"azureActiveDirectoryAuthentications"`
	// One or more `clientRevokedCertificate` blocks as defined below.
	ClientRevokedCertificates VpnServerConfigurationClientRevokedCertificateArrayOutput `pulumi:"clientRevokedCertificates"`
	// One or more `clientRootCertificate` blocks as defined below.
	ClientRootCertificates VpnServerConfigurationClientRootCertificateArrayOutput `pulumi:"clientRootCertificates"`
	// A `ipsecPolicy` block as defined below.
	IpsecPolicy VpnServerConfigurationIpsecPolicyPtrOutput `pulumi:"ipsecPolicy"`
	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `radius` block as defined below.
	Radius VpnServerConfigurationRadiusPtrOutput `pulumi:"radius"`
	// A `radiusServer` block as defined below.
	//
	// Deprecated: Deprecated in favour of `radius`
	RadiusServer VpnServerConfigurationRadiusServerPtrOutput `pulumi:"radiusServer"`
	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
	VpnAuthenticationTypes pulumi.StringOutput `pulumi:"vpnAuthenticationTypes"`
	// A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
	VpnProtocols pulumi.StringArrayOutput `pulumi:"vpnProtocols"`
}

// NewVpnServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *VpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VpnAuthenticationTypes == nil {
		return nil, errors.New("invalid value for required argument 'VpnAuthenticationTypes'")
	}
	var resource VpnServerConfiguration
	err := ctx.RegisterResource("azure:network/vpnServerConfiguration:VpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnServerConfiguration gets an existing VpnServerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnServerConfigurationState, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	var resource VpnServerConfiguration
	err := ctx.ReadResource("azure:network/vpnServerConfiguration:VpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnServerConfiguration resources.
type vpnServerConfigurationState struct {
	// A `azureActiveDirectoryAuthentication` block as defined below.
	AzureActiveDirectoryAuthentications []VpnServerConfigurationAzureActiveDirectoryAuthentication `pulumi:"azureActiveDirectoryAuthentications"`
	// One or more `clientRevokedCertificate` blocks as defined below.
	ClientRevokedCertificates []VpnServerConfigurationClientRevokedCertificate `pulumi:"clientRevokedCertificates"`
	// One or more `clientRootCertificate` blocks as defined below.
	ClientRootCertificates []VpnServerConfigurationClientRootCertificate `pulumi:"clientRootCertificates"`
	// A `ipsecPolicy` block as defined below.
	IpsecPolicy *VpnServerConfigurationIpsecPolicy `pulumi:"ipsecPolicy"`
	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `radius` block as defined below.
	Radius *VpnServerConfigurationRadius `pulumi:"radius"`
	// A `radiusServer` block as defined below.
	//
	// Deprecated: Deprecated in favour of `radius`
	RadiusServer *VpnServerConfigurationRadiusServer `pulumi:"radiusServer"`
	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
	VpnAuthenticationTypes *string `pulumi:"vpnAuthenticationTypes"`
	// A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
	VpnProtocols []string `pulumi:"vpnProtocols"`
}

type VpnServerConfigurationState struct {
	// A `azureActiveDirectoryAuthentication` block as defined below.
	AzureActiveDirectoryAuthentications VpnServerConfigurationAzureActiveDirectoryAuthenticationArrayInput
	// One or more `clientRevokedCertificate` blocks as defined below.
	ClientRevokedCertificates VpnServerConfigurationClientRevokedCertificateArrayInput
	// One or more `clientRootCertificate` blocks as defined below.
	ClientRootCertificates VpnServerConfigurationClientRootCertificateArrayInput
	// A `ipsecPolicy` block as defined below.
	IpsecPolicy VpnServerConfigurationIpsecPolicyPtrInput
	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `radius` block as defined below.
	Radius VpnServerConfigurationRadiusPtrInput
	// A `radiusServer` block as defined below.
	//
	// Deprecated: Deprecated in favour of `radius`
	RadiusServer VpnServerConfigurationRadiusServerPtrInput
	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
	VpnAuthenticationTypes pulumi.StringPtrInput
	// A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
	VpnProtocols pulumi.StringArrayInput
}

func (VpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationState)(nil)).Elem()
}

type vpnServerConfigurationArgs struct {
	// A `azureActiveDirectoryAuthentication` block as defined below.
	AzureActiveDirectoryAuthentications []VpnServerConfigurationAzureActiveDirectoryAuthentication `pulumi:"azureActiveDirectoryAuthentications"`
	// One or more `clientRevokedCertificate` blocks as defined below.
	ClientRevokedCertificates []VpnServerConfigurationClientRevokedCertificate `pulumi:"clientRevokedCertificates"`
	// One or more `clientRootCertificate` blocks as defined below.
	ClientRootCertificates []VpnServerConfigurationClientRootCertificate `pulumi:"clientRootCertificates"`
	// A `ipsecPolicy` block as defined below.
	IpsecPolicy *VpnServerConfigurationIpsecPolicy `pulumi:"ipsecPolicy"`
	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `radius` block as defined below.
	Radius *VpnServerConfigurationRadius `pulumi:"radius"`
	// A `radiusServer` block as defined below.
	//
	// Deprecated: Deprecated in favour of `radius`
	RadiusServer *VpnServerConfigurationRadiusServer `pulumi:"radiusServer"`
	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
	VpnAuthenticationTypes string `pulumi:"vpnAuthenticationTypes"`
	// A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
	VpnProtocols []string `pulumi:"vpnProtocols"`
}

// The set of arguments for constructing a VpnServerConfiguration resource.
type VpnServerConfigurationArgs struct {
	// A `azureActiveDirectoryAuthentication` block as defined below.
	AzureActiveDirectoryAuthentications VpnServerConfigurationAzureActiveDirectoryAuthenticationArrayInput
	// One or more `clientRevokedCertificate` blocks as defined below.
	ClientRevokedCertificates VpnServerConfigurationClientRevokedCertificateArrayInput
	// One or more `clientRootCertificate` blocks as defined below.
	ClientRootCertificates VpnServerConfigurationClientRootCertificateArrayInput
	// A `ipsecPolicy` block as defined below.
	IpsecPolicy VpnServerConfigurationIpsecPolicyPtrInput
	// The Azure location where this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Name which should be used for this VPN Server Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `radius` block as defined below.
	Radius VpnServerConfigurationRadiusPtrInput
	// A `radiusServer` block as defined below.
	//
	// Deprecated: Deprecated in favour of `radius`
	RadiusServer VpnServerConfigurationRadiusServerPtrInput
	// The Name of the Resource Group in which this VPN Server Configuration should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of Authentication Types applicable for this VPN Server Configuration. Possible values are `AAD` (Azure Active Directory), `Certificate` and `Radius`.
	VpnAuthenticationTypes pulumi.StringInput
	// A list of VPN Protocols to use for this Server Configuration. Possible values are `IkeV2` and `OpenVPN`.
	VpnProtocols pulumi.StringArrayInput
}

func (VpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationArgs)(nil)).Elem()
}

type VpnServerConfigurationInput interface {
	pulumi.Input

	ToVpnServerConfigurationOutput() VpnServerConfigurationOutput
	ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput
}

func (*VpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return i.ToVpnServerConfigurationOutputWithContext(context.Background())
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationOutput)
}

// VpnServerConfigurationArrayInput is an input type that accepts VpnServerConfigurationArray and VpnServerConfigurationArrayOutput values.
// You can construct a concrete instance of `VpnServerConfigurationArrayInput` via:
//
//          VpnServerConfigurationArray{ VpnServerConfigurationArgs{...} }
type VpnServerConfigurationArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationArrayOutput() VpnServerConfigurationArrayOutput
	ToVpnServerConfigurationArrayOutputWithContext(context.Context) VpnServerConfigurationArrayOutput
}

type VpnServerConfigurationArray []VpnServerConfigurationInput

func (VpnServerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnServerConfiguration)(nil)).Elem()
}

func (i VpnServerConfigurationArray) ToVpnServerConfigurationArrayOutput() VpnServerConfigurationArrayOutput {
	return i.ToVpnServerConfigurationArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationArray) ToVpnServerConfigurationArrayOutputWithContext(ctx context.Context) VpnServerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationArrayOutput)
}

// VpnServerConfigurationMapInput is an input type that accepts VpnServerConfigurationMap and VpnServerConfigurationMapOutput values.
// You can construct a concrete instance of `VpnServerConfigurationMapInput` via:
//
//          VpnServerConfigurationMap{ "key": VpnServerConfigurationArgs{...} }
type VpnServerConfigurationMapInput interface {
	pulumi.Input

	ToVpnServerConfigurationMapOutput() VpnServerConfigurationMapOutput
	ToVpnServerConfigurationMapOutputWithContext(context.Context) VpnServerConfigurationMapOutput
}

type VpnServerConfigurationMap map[string]VpnServerConfigurationInput

func (VpnServerConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnServerConfiguration)(nil)).Elem()
}

func (i VpnServerConfigurationMap) ToVpnServerConfigurationMapOutput() VpnServerConfigurationMapOutput {
	return i.ToVpnServerConfigurationMapOutputWithContext(context.Background())
}

func (i VpnServerConfigurationMap) ToVpnServerConfigurationMapOutputWithContext(ctx context.Context) VpnServerConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationMapOutput)
}

type VpnServerConfigurationOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return o
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return o
}

type VpnServerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnServerConfiguration)(nil)).Elem()
}

func (o VpnServerConfigurationArrayOutput) ToVpnServerConfigurationArrayOutput() VpnServerConfigurationArrayOutput {
	return o
}

func (o VpnServerConfigurationArrayOutput) ToVpnServerConfigurationArrayOutputWithContext(ctx context.Context) VpnServerConfigurationArrayOutput {
	return o
}

func (o VpnServerConfigurationArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnServerConfiguration {
		return vs[0].([]*VpnServerConfiguration)[vs[1].(int)]
	}).(VpnServerConfigurationOutput)
}

type VpnServerConfigurationMapOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnServerConfiguration)(nil)).Elem()
}

func (o VpnServerConfigurationMapOutput) ToVpnServerConfigurationMapOutput() VpnServerConfigurationMapOutput {
	return o
}

func (o VpnServerConfigurationMapOutput) ToVpnServerConfigurationMapOutputWithContext(ctx context.Context) VpnServerConfigurationMapOutput {
	return o
}

func (o VpnServerConfigurationMapOutput) MapIndex(k pulumi.StringInput) VpnServerConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnServerConfiguration {
		return vs[0].(map[string]*VpnServerConfiguration)[vs[1].(string)]
	}).(VpnServerConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnServerConfigurationInput)(nil)).Elem(), &VpnServerConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnServerConfigurationArrayInput)(nil)).Elem(), VpnServerConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnServerConfigurationMapInput)(nil)).Elem(), VpnServerConfigurationMap{})
	pulumi.RegisterOutputType(VpnServerConfigurationOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationMapOutput{})
}
