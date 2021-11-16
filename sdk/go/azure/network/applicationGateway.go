// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Application Gateway.
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
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.254.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		frontend, err := network.NewSubnet(ctx, "frontend", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.254.0.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnet(ctx, "backend", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.254.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewApplicationGateway(ctx, "network", &network.ApplicationGatewayArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &network.ApplicationGatewaySkuArgs{
// 				Name:     pulumi.String("Standard_Small"),
// 				Tier:     pulumi.String("Standard"),
// 				Capacity: pulumi.Int(2),
// 			},
// 			GatewayIpConfigurations: network.ApplicationGatewayGatewayIpConfigurationArray{
// 				&network.ApplicationGatewayGatewayIpConfigurationArgs{
// 					Name:     pulumi.String("my-gateway-ip-configuration"),
// 					SubnetId: frontend.ID(),
// 				},
// 			},
// 			FrontendPorts: network.ApplicationGatewayFrontendPortArray{
// 				&network.ApplicationGatewayFrontendPortArgs{
// 					Name: pulumi.String(frontendPortName),
// 					Port: pulumi.Int(80),
// 				},
// 			},
// 			FrontendIpConfigurations: network.ApplicationGatewayFrontendIpConfigurationArray{
// 				&network.ApplicationGatewayFrontendIpConfigurationArgs{
// 					Name:              pulumi.String(frontendIpConfigurationName),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 			BackendAddressPools: network.ApplicationGatewayBackendAddressPoolArray{
// 				&network.ApplicationGatewayBackendAddressPoolArgs{
// 					Name: pulumi.String(backendAddressPoolName),
// 				},
// 			},
// 			BackendHttpSettings: network.ApplicationGatewayBackendHttpSettingArray{
// 				&network.ApplicationGatewayBackendHttpSettingArgs{
// 					Name:                pulumi.String(httpSettingName),
// 					CookieBasedAffinity: pulumi.String("Disabled"),
// 					Path:                pulumi.String("/path1/"),
// 					Port:                pulumi.Int(80),
// 					Protocol:            pulumi.String("Http"),
// 					RequestTimeout:      pulumi.Int(60),
// 				},
// 			},
// 			HttpListeners: network.ApplicationGatewayHttpListenerArray{
// 				&network.ApplicationGatewayHttpListenerArgs{
// 					Name:                        pulumi.String(listenerName),
// 					FrontendIpConfigurationName: pulumi.String(frontendIpConfigurationName),
// 					FrontendPortName:            pulumi.String(frontendPortName),
// 					Protocol:                    pulumi.String("Http"),
// 				},
// 			},
// 			RequestRoutingRules: network.ApplicationGatewayRequestRoutingRuleArray{
// 				&network.ApplicationGatewayRequestRoutingRuleArgs{
// 					Name:                    pulumi.String(requestRoutingRuleName),
// 					RuleType:                pulumi.String("Basic"),
// 					HttpListenerName:        pulumi.String(listenerName),
// 					BackendAddressPoolName:  pulumi.String(backendAddressPoolName),
// 					BackendHttpSettingsName: pulumi.String(httpSettingName),
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
// Application Gateway's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/applicationGateway:ApplicationGateway example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/applicationGateways/myGateway1
// ```
type ApplicationGateway struct {
	pulumi.CustomResourceState

	// One or more `authenticationCertificate` blocks as defined below.
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateArrayOutput `pulumi:"authenticationCertificates"`
	// A `autoscaleConfiguration` block as defined below.
	AutoscaleConfiguration ApplicationGatewayAutoscaleConfigurationPtrOutput `pulumi:"autoscaleConfiguration"`
	// One or more `backendAddressPool` blocks as defined below.
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayOutput `pulumi:"backendAddressPools"`
	// One or more `backendHttpSettings` blocks as defined below.
	BackendHttpSettings ApplicationGatewayBackendHttpSettingArrayOutput `pulumi:"backendHttpSettings"`
	// One or more `customErrorConfiguration` blocks as defined below.
	CustomErrorConfigurations ApplicationGatewayCustomErrorConfigurationArrayOutput `pulumi:"customErrorConfigurations"`
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 pulumi.BoolPtrOutput `pulumi:"enableHttp2"`
	// The ID of the Web Application Firewall Policy.
	FirewallPolicyId pulumi.StringPtrOutput `pulumi:"firewallPolicyId"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations ApplicationGatewayFrontendIpConfigurationArrayOutput `pulumi:"frontendIpConfigurations"`
	// One or more `frontendPort` blocks as defined below.
	FrontendPorts ApplicationGatewayFrontendPortArrayOutput `pulumi:"frontendPorts"`
	// One or more `gatewayIpConfiguration` blocks as defined below.
	GatewayIpConfigurations ApplicationGatewayGatewayIpConfigurationArrayOutput `pulumi:"gatewayIpConfigurations"`
	// One or more `httpListener` blocks as defined below.
	HttpListeners ApplicationGatewayHttpListenerArrayOutput `pulumi:"httpListeners"`
	// An `identity` block as defined below.
	Identity ApplicationGatewayIdentityPtrOutput `pulumi:"identity"`
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `probe` blocks as defined below.
	Probes ApplicationGatewayProbeArrayOutput `pulumi:"probes"`
	// One or more `redirectConfiguration` blocks as defined below.
	RedirectConfigurations ApplicationGatewayRedirectConfigurationArrayOutput `pulumi:"redirectConfigurations"`
	// One or more `requestRoutingRule` blocks as defined below.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayOutput `pulumi:"requestRoutingRules"`
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `rewriteRuleSet` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets ApplicationGatewayRewriteRuleSetArrayOutput `pulumi:"rewriteRuleSets"`
	// A `sku` block as defined below.
	Sku ApplicationGatewaySkuOutput `pulumi:"sku"`
	// One or more `sslCertificate` blocks as defined below.
	SslCertificates ApplicationGatewaySslCertificateArrayOutput `pulumi:"sslCertificates"`
	// a `ssl policy` block as defined below.
	SslPolicies ApplicationGatewaySslPolicyArrayOutput `pulumi:"sslPolicies"`
	// One or more `sslProfile` blocks as defined below.
	SslProfiles ApplicationGatewaySslProfileArrayOutput `pulumi:"sslProfiles"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// One or more `trustedClientCertificate` blocks as defined below.
	TrustedClientCertificates ApplicationGatewayTrustedClientCertificateArrayOutput `pulumi:"trustedClientCertificates"`
	// One or more `trustedRootCertificate` blocks as defined below.
	TrustedRootCertificates ApplicationGatewayTrustedRootCertificateArrayOutput `pulumi:"trustedRootCertificates"`
	// One or more `urlPathMap` blocks as defined below.
	UrlPathMaps ApplicationGatewayUrlPathMapArrayOutput `pulumi:"urlPathMaps"`
	// A `wafConfiguration` block as defined below.
	WafConfiguration ApplicationGatewayWafConfigurationPtrOutput `pulumi:"wafConfiguration"`
	// A collection of availability zones to spread the Application Gateway over.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewApplicationGateway registers a new resource with the given unique name, arguments, and options.
func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendAddressPools == nil {
		return nil, errors.New("invalid value for required argument 'BackendAddressPools'")
	}
	if args.BackendHttpSettings == nil {
		return nil, errors.New("invalid value for required argument 'BackendHttpSettings'")
	}
	if args.FrontendIpConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'FrontendIpConfigurations'")
	}
	if args.FrontendPorts == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPorts'")
	}
	if args.GatewayIpConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'GatewayIpConfigurations'")
	}
	if args.HttpListeners == nil {
		return nil, errors.New("invalid value for required argument 'HttpListeners'")
	}
	if args.RequestRoutingRules == nil {
		return nil, errors.New("invalid value for required argument 'RequestRoutingRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource ApplicationGateway
	err := ctx.RegisterResource("azure:network/applicationGateway:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGateway gets an existing ApplicationGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azure:network/applicationGateway:ApplicationGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGateway resources.
type applicationGatewayState struct {
	// One or more `authenticationCertificate` blocks as defined below.
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificate `pulumi:"authenticationCertificates"`
	// A `autoscaleConfiguration` block as defined below.
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfiguration `pulumi:"autoscaleConfiguration"`
	// One or more `backendAddressPool` blocks as defined below.
	BackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"backendAddressPools"`
	// One or more `backendHttpSettings` blocks as defined below.
	BackendHttpSettings []ApplicationGatewayBackendHttpSetting `pulumi:"backendHttpSettings"`
	// One or more `customErrorConfiguration` blocks as defined below.
	CustomErrorConfigurations []ApplicationGatewayCustomErrorConfiguration `pulumi:"customErrorConfigurations"`
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// The ID of the Web Application Firewall Policy.
	FirewallPolicyId *string `pulumi:"firewallPolicyId"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations []ApplicationGatewayFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// One or more `frontendPort` blocks as defined below.
	FrontendPorts []ApplicationGatewayFrontendPort `pulumi:"frontendPorts"`
	// One or more `gatewayIpConfiguration` blocks as defined below.
	GatewayIpConfigurations []ApplicationGatewayGatewayIpConfiguration `pulumi:"gatewayIpConfigurations"`
	// One or more `httpListener` blocks as defined below.
	HttpListeners []ApplicationGatewayHttpListener `pulumi:"httpListeners"`
	// An `identity` block as defined below.
	Identity *ApplicationGatewayIdentity `pulumi:"identity"`
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `probe` blocks as defined below.
	Probes []ApplicationGatewayProbe `pulumi:"probes"`
	// One or more `redirectConfiguration` blocks as defined below.
	RedirectConfigurations []ApplicationGatewayRedirectConfiguration `pulumi:"redirectConfigurations"`
	// One or more `requestRoutingRule` blocks as defined below.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule `pulumi:"requestRoutingRules"`
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `rewriteRuleSet` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets []ApplicationGatewayRewriteRuleSet `pulumi:"rewriteRuleSets"`
	// A `sku` block as defined below.
	Sku *ApplicationGatewaySku `pulumi:"sku"`
	// One or more `sslCertificate` blocks as defined below.
	SslCertificates []ApplicationGatewaySslCertificate `pulumi:"sslCertificates"`
	// a `ssl policy` block as defined below.
	SslPolicies []ApplicationGatewaySslPolicy `pulumi:"sslPolicies"`
	// One or more `sslProfile` blocks as defined below.
	SslProfiles []ApplicationGatewaySslProfile `pulumi:"sslProfiles"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `trustedClientCertificate` blocks as defined below.
	TrustedClientCertificates []ApplicationGatewayTrustedClientCertificate `pulumi:"trustedClientCertificates"`
	// One or more `trustedRootCertificate` blocks as defined below.
	TrustedRootCertificates []ApplicationGatewayTrustedRootCertificate `pulumi:"trustedRootCertificates"`
	// One or more `urlPathMap` blocks as defined below.
	UrlPathMaps []ApplicationGatewayUrlPathMap `pulumi:"urlPathMaps"`
	// A `wafConfiguration` block as defined below.
	WafConfiguration *ApplicationGatewayWafConfiguration `pulumi:"wafConfiguration"`
	// A collection of availability zones to spread the Application Gateway over.
	Zones []string `pulumi:"zones"`
}

type ApplicationGatewayState struct {
	// One or more `authenticationCertificate` blocks as defined below.
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateArrayInput
	// A `autoscaleConfiguration` block as defined below.
	AutoscaleConfiguration ApplicationGatewayAutoscaleConfigurationPtrInput
	// One or more `backendAddressPool` blocks as defined below.
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput
	// One or more `backendHttpSettings` blocks as defined below.
	BackendHttpSettings ApplicationGatewayBackendHttpSettingArrayInput
	// One or more `customErrorConfiguration` blocks as defined below.
	CustomErrorConfigurations ApplicationGatewayCustomErrorConfigurationArrayInput
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 pulumi.BoolPtrInput
	// The ID of the Web Application Firewall Policy.
	FirewallPolicyId pulumi.StringPtrInput
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations ApplicationGatewayFrontendIpConfigurationArrayInput
	// One or more `frontendPort` blocks as defined below.
	FrontendPorts ApplicationGatewayFrontendPortArrayInput
	// One or more `gatewayIpConfiguration` blocks as defined below.
	GatewayIpConfigurations ApplicationGatewayGatewayIpConfigurationArrayInput
	// One or more `httpListener` blocks as defined below.
	HttpListeners ApplicationGatewayHttpListenerArrayInput
	// An `identity` block as defined below.
	Identity ApplicationGatewayIdentityPtrInput
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `probe` blocks as defined below.
	Probes ApplicationGatewayProbeArrayInput
	// One or more `redirectConfiguration` blocks as defined below.
	RedirectConfigurations ApplicationGatewayRedirectConfigurationArrayInput
	// One or more `requestRoutingRule` blocks as defined below.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayInput
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `rewriteRuleSet` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets ApplicationGatewayRewriteRuleSetArrayInput
	// A `sku` block as defined below.
	Sku ApplicationGatewaySkuPtrInput
	// One or more `sslCertificate` blocks as defined below.
	SslCertificates ApplicationGatewaySslCertificateArrayInput
	// a `ssl policy` block as defined below.
	SslPolicies ApplicationGatewaySslPolicyArrayInput
	// One or more `sslProfile` blocks as defined below.
	SslProfiles ApplicationGatewaySslProfileArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more `trustedClientCertificate` blocks as defined below.
	TrustedClientCertificates ApplicationGatewayTrustedClientCertificateArrayInput
	// One or more `trustedRootCertificate` blocks as defined below.
	TrustedRootCertificates ApplicationGatewayTrustedRootCertificateArrayInput
	// One or more `urlPathMap` blocks as defined below.
	UrlPathMaps ApplicationGatewayUrlPathMapArrayInput
	// A `wafConfiguration` block as defined below.
	WafConfiguration ApplicationGatewayWafConfigurationPtrInput
	// A collection of availability zones to spread the Application Gateway over.
	Zones pulumi.StringArrayInput
}

func (ApplicationGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayState)(nil)).Elem()
}

type applicationGatewayArgs struct {
	// One or more `authenticationCertificate` blocks as defined below.
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificate `pulumi:"authenticationCertificates"`
	// A `autoscaleConfiguration` block as defined below.
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfiguration `pulumi:"autoscaleConfiguration"`
	// One or more `backendAddressPool` blocks as defined below.
	BackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"backendAddressPools"`
	// One or more `backendHttpSettings` blocks as defined below.
	BackendHttpSettings []ApplicationGatewayBackendHttpSetting `pulumi:"backendHttpSettings"`
	// One or more `customErrorConfiguration` blocks as defined below.
	CustomErrorConfigurations []ApplicationGatewayCustomErrorConfiguration `pulumi:"customErrorConfigurations"`
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// The ID of the Web Application Firewall Policy.
	FirewallPolicyId *string `pulumi:"firewallPolicyId"`
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations []ApplicationGatewayFrontendIpConfiguration `pulumi:"frontendIpConfigurations"`
	// One or more `frontendPort` blocks as defined below.
	FrontendPorts []ApplicationGatewayFrontendPort `pulumi:"frontendPorts"`
	// One or more `gatewayIpConfiguration` blocks as defined below.
	GatewayIpConfigurations []ApplicationGatewayGatewayIpConfiguration `pulumi:"gatewayIpConfigurations"`
	// One or more `httpListener` blocks as defined below.
	HttpListeners []ApplicationGatewayHttpListener `pulumi:"httpListeners"`
	// An `identity` block as defined below.
	Identity *ApplicationGatewayIdentity `pulumi:"identity"`
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `probe` blocks as defined below.
	Probes []ApplicationGatewayProbe `pulumi:"probes"`
	// One or more `redirectConfiguration` blocks as defined below.
	RedirectConfigurations []ApplicationGatewayRedirectConfiguration `pulumi:"redirectConfigurations"`
	// One or more `requestRoutingRule` blocks as defined below.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule `pulumi:"requestRoutingRules"`
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `rewriteRuleSet` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets []ApplicationGatewayRewriteRuleSet `pulumi:"rewriteRuleSets"`
	// A `sku` block as defined below.
	Sku ApplicationGatewaySku `pulumi:"sku"`
	// One or more `sslCertificate` blocks as defined below.
	SslCertificates []ApplicationGatewaySslCertificate `pulumi:"sslCertificates"`
	// a `ssl policy` block as defined below.
	SslPolicies []ApplicationGatewaySslPolicy `pulumi:"sslPolicies"`
	// One or more `sslProfile` blocks as defined below.
	SslProfiles []ApplicationGatewaySslProfile `pulumi:"sslProfiles"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more `trustedClientCertificate` blocks as defined below.
	TrustedClientCertificates []ApplicationGatewayTrustedClientCertificate `pulumi:"trustedClientCertificates"`
	// One or more `trustedRootCertificate` blocks as defined below.
	TrustedRootCertificates []ApplicationGatewayTrustedRootCertificate `pulumi:"trustedRootCertificates"`
	// One or more `urlPathMap` blocks as defined below.
	UrlPathMaps []ApplicationGatewayUrlPathMap `pulumi:"urlPathMaps"`
	// A `wafConfiguration` block as defined below.
	WafConfiguration *ApplicationGatewayWafConfiguration `pulumi:"wafConfiguration"`
	// A collection of availability zones to spread the Application Gateway over.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ApplicationGateway resource.
type ApplicationGatewayArgs struct {
	// One or more `authenticationCertificate` blocks as defined below.
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateArrayInput
	// A `autoscaleConfiguration` block as defined below.
	AutoscaleConfiguration ApplicationGatewayAutoscaleConfigurationPtrInput
	// One or more `backendAddressPool` blocks as defined below.
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput
	// One or more `backendHttpSettings` blocks as defined below.
	BackendHttpSettings ApplicationGatewayBackendHttpSettingArrayInput
	// One or more `customErrorConfiguration` blocks as defined below.
	CustomErrorConfigurations ApplicationGatewayCustomErrorConfigurationArrayInput
	// Is HTTP2 enabled on the application gateway resource? Defaults to `false`.
	EnableHttp2 pulumi.BoolPtrInput
	// The ID of the Web Application Firewall Policy.
	FirewallPolicyId pulumi.StringPtrInput
	// One or more `frontendIpConfiguration` blocks as defined below.
	FrontendIpConfigurations ApplicationGatewayFrontendIpConfigurationArrayInput
	// One or more `frontendPort` blocks as defined below.
	FrontendPorts ApplicationGatewayFrontendPortArrayInput
	// One or more `gatewayIpConfiguration` blocks as defined below.
	GatewayIpConfigurations ApplicationGatewayGatewayIpConfigurationArrayInput
	// One or more `httpListener` blocks as defined below.
	HttpListeners ApplicationGatewayHttpListenerArrayInput
	// An `identity` block as defined below.
	Identity ApplicationGatewayIdentityPtrInput
	// The Azure region where the Application Gateway should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Application Gateway. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `probe` blocks as defined below.
	Probes ApplicationGatewayProbeArrayInput
	// One or more `redirectConfiguration` blocks as defined below.
	RedirectConfigurations ApplicationGatewayRedirectConfigurationArrayInput
	// One or more `requestRoutingRule` blocks as defined below.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayInput
	// The name of the resource group in which to the Application Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `rewriteRuleSet` blocks as defined below. Only valid for v2 SKUs.
	RewriteRuleSets ApplicationGatewayRewriteRuleSetArrayInput
	// A `sku` block as defined below.
	Sku ApplicationGatewaySkuInput
	// One or more `sslCertificate` blocks as defined below.
	SslCertificates ApplicationGatewaySslCertificateArrayInput
	// a `ssl policy` block as defined below.
	SslPolicies ApplicationGatewaySslPolicyArrayInput
	// One or more `sslProfile` blocks as defined below.
	SslProfiles ApplicationGatewaySslProfileArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more `trustedClientCertificate` blocks as defined below.
	TrustedClientCertificates ApplicationGatewayTrustedClientCertificateArrayInput
	// One or more `trustedRootCertificate` blocks as defined below.
	TrustedRootCertificates ApplicationGatewayTrustedRootCertificateArrayInput
	// One or more `urlPathMap` blocks as defined below.
	UrlPathMaps ApplicationGatewayUrlPathMapArrayInput
	// A `wafConfiguration` block as defined below.
	WafConfiguration ApplicationGatewayWafConfigurationPtrInput
	// A collection of availability zones to spread the Application Gateway over.
	Zones pulumi.StringArrayInput
}

func (ApplicationGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayArgs)(nil)).Elem()
}

type ApplicationGatewayInput interface {
	pulumi.Input

	ToApplicationGatewayOutput() ApplicationGatewayOutput
	ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput
}

func (*ApplicationGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGateway)(nil))
}

func (i *ApplicationGateway) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return i.ToApplicationGatewayOutputWithContext(context.Background())
}

func (i *ApplicationGateway) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayOutput)
}

func (i *ApplicationGateway) ToApplicationGatewayPtrOutput() ApplicationGatewayPtrOutput {
	return i.ToApplicationGatewayPtrOutputWithContext(context.Background())
}

func (i *ApplicationGateway) ToApplicationGatewayPtrOutputWithContext(ctx context.Context) ApplicationGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPtrOutput)
}

type ApplicationGatewayPtrInput interface {
	pulumi.Input

	ToApplicationGatewayPtrOutput() ApplicationGatewayPtrOutput
	ToApplicationGatewayPtrOutputWithContext(ctx context.Context) ApplicationGatewayPtrOutput
}

type applicationGatewayPtrType ApplicationGatewayArgs

func (*applicationGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil))
}

func (i *applicationGatewayPtrType) ToApplicationGatewayPtrOutput() ApplicationGatewayPtrOutput {
	return i.ToApplicationGatewayPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayPtrType) ToApplicationGatewayPtrOutputWithContext(ctx context.Context) ApplicationGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPtrOutput)
}

// ApplicationGatewayArrayInput is an input type that accepts ApplicationGatewayArray and ApplicationGatewayArrayOutput values.
// You can construct a concrete instance of `ApplicationGatewayArrayInput` via:
//
//          ApplicationGatewayArray{ ApplicationGatewayArgs{...} }
type ApplicationGatewayArrayInput interface {
	pulumi.Input

	ToApplicationGatewayArrayOutput() ApplicationGatewayArrayOutput
	ToApplicationGatewayArrayOutputWithContext(context.Context) ApplicationGatewayArrayOutput
}

type ApplicationGatewayArray []ApplicationGatewayInput

func (ApplicationGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGateway)(nil)).Elem()
}

func (i ApplicationGatewayArray) ToApplicationGatewayArrayOutput() ApplicationGatewayArrayOutput {
	return i.ToApplicationGatewayArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayArray) ToApplicationGatewayArrayOutputWithContext(ctx context.Context) ApplicationGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayArrayOutput)
}

// ApplicationGatewayMapInput is an input type that accepts ApplicationGatewayMap and ApplicationGatewayMapOutput values.
// You can construct a concrete instance of `ApplicationGatewayMapInput` via:
//
//          ApplicationGatewayMap{ "key": ApplicationGatewayArgs{...} }
type ApplicationGatewayMapInput interface {
	pulumi.Input

	ToApplicationGatewayMapOutput() ApplicationGatewayMapOutput
	ToApplicationGatewayMapOutputWithContext(context.Context) ApplicationGatewayMapOutput
}

type ApplicationGatewayMap map[string]ApplicationGatewayInput

func (ApplicationGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGateway)(nil)).Elem()
}

func (i ApplicationGatewayMap) ToApplicationGatewayMapOutput() ApplicationGatewayMapOutput {
	return i.ToApplicationGatewayMapOutputWithContext(context.Background())
}

func (i ApplicationGatewayMap) ToApplicationGatewayMapOutputWithContext(ctx context.Context) ApplicationGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayMapOutput)
}

type ApplicationGatewayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGateway)(nil))
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return o
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return o
}

func (o ApplicationGatewayOutput) ToApplicationGatewayPtrOutput() ApplicationGatewayPtrOutput {
	return o.ToApplicationGatewayPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayOutput) ToApplicationGatewayPtrOutputWithContext(ctx context.Context) ApplicationGatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGateway) *ApplicationGateway {
		return &v
	}).(ApplicationGatewayPtrOutput)
}

type ApplicationGatewayPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil))
}

func (o ApplicationGatewayPtrOutput) ToApplicationGatewayPtrOutput() ApplicationGatewayPtrOutput {
	return o
}

func (o ApplicationGatewayPtrOutput) ToApplicationGatewayPtrOutputWithContext(ctx context.Context) ApplicationGatewayPtrOutput {
	return o
}

func (o ApplicationGatewayPtrOutput) Elem() ApplicationGatewayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGateway {
		if v != nil {
			return *v
		}
		var ret ApplicationGateway
		return ret
	}).(ApplicationGatewayOutput)
}

type ApplicationGatewayArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGateway)(nil))
}

func (o ApplicationGatewayArrayOutput) ToApplicationGatewayArrayOutput() ApplicationGatewayArrayOutput {
	return o
}

func (o ApplicationGatewayArrayOutput) ToApplicationGatewayArrayOutputWithContext(ctx context.Context) ApplicationGatewayArrayOutput {
	return o
}

func (o ApplicationGatewayArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGateway {
		return vs[0].([]ApplicationGateway)[vs[1].(int)]
	}).(ApplicationGatewayOutput)
}

type ApplicationGatewayMapOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationGateway)(nil))
}

func (o ApplicationGatewayMapOutput) ToApplicationGatewayMapOutput() ApplicationGatewayMapOutput {
	return o
}

func (o ApplicationGatewayMapOutput) ToApplicationGatewayMapOutputWithContext(ctx context.Context) ApplicationGatewayMapOutput {
	return o
}

func (o ApplicationGatewayMapOutput) MapIndex(k pulumi.StringInput) ApplicationGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationGateway {
		return vs[0].(map[string]ApplicationGateway)[vs[1].(string)]
	}).(ApplicationGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayMapOutput{})
}
