// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:core/customProvider:CustomProvider":
		r = &CustomProvider{}
	case "azure:core/portalTenantConfiguration:PortalTenantConfiguration":
		r = &PortalTenantConfiguration{}
	case "azure:core/resourceGroup:ResourceGroup":
		r = &ResourceGroup{}
	case "azure:core/resourceGroupPolicyAssignment:ResourceGroupPolicyAssignment":
		r = &ResourceGroupPolicyAssignment{}
	case "azure:core/resourceGroupTemplateDeployment:ResourceGroupTemplateDeployment":
		r = &ResourceGroupTemplateDeployment{}
	case "azure:core/resourcePolicyAssignment:ResourcePolicyAssignment":
		r = &ResourcePolicyAssignment{}
	case "azure:core/resourceProviderRegistration:ResourceProviderRegistration":
		r = &ResourceProviderRegistration{}
	case "azure:core/subscription:Subscription":
		r = &Subscription{}
	case "azure:core/subscriptionPolicyAssignment:SubscriptionPolicyAssignment":
		r = &SubscriptionPolicyAssignment{}
	case "azure:core/subscriptionTemplateDeployment:SubscriptionTemplateDeployment":
		r = &SubscriptionTemplateDeployment{}
	case "azure:core/templateDeployment:TemplateDeployment":
		r = &TemplateDeployment{}
	case "azure:core/tenantTemplateDeployment:TenantTemplateDeployment":
		r = &TenantTemplateDeployment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"core/customProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/portalTenantConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceGroupPolicyAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceGroupTemplateDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourcePolicyAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/resourceProviderRegistration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/subscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/subscriptionPolicyAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/subscriptionTemplateDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/templateDeployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"core/tenantTemplateDeployment",
		&module{version},
	)
}
