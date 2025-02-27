// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

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
	case "azure:appservice/activeSlot:ActiveSlot":
		r = &ActiveSlot{}
	case "azure:appservice/appService:AppService":
		r = &AppService{}
	case "azure:appservice/certificate:Certificate":
		r = &Certificate{}
	case "azure:appservice/certificateBinding:CertificateBinding":
		r = &CertificateBinding{}
	case "azure:appservice/certificateOrder:CertificateOrder":
		r = &CertificateOrder{}
	case "azure:appservice/customHostnameBinding:CustomHostnameBinding":
		r = &CustomHostnameBinding{}
	case "azure:appservice/environment:Environment":
		r = &Environment{}
	case "azure:appservice/environmentV3:EnvironmentV3":
		r = &EnvironmentV3{}
	case "azure:appservice/functionApp:FunctionApp":
		r = &FunctionApp{}
	case "azure:appservice/functionAppSlot:FunctionAppSlot":
		r = &FunctionAppSlot{}
	case "azure:appservice/hybridConnection:HybridConnection":
		r = &HybridConnection{}
	case "azure:appservice/managedCertificate:ManagedCertificate":
		r = &ManagedCertificate{}
	case "azure:appservice/plan:Plan":
		r = &Plan{}
	case "azure:appservice/publicCertificate:PublicCertificate":
		r = &PublicCertificate{}
	case "azure:appservice/slot:Slot":
		r = &Slot{}
	case "azure:appservice/slotCustomHostnameBinding:SlotCustomHostnameBinding":
		r = &SlotCustomHostnameBinding{}
	case "azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection":
		r = &SlotVirtualNetworkSwiftConnection{}
	case "azure:appservice/sourceCodeToken:SourceCodeToken":
		r = &SourceCodeToken{}
	case "azure:appservice/staticSite:StaticSite":
		r = &StaticSite{}
	case "azure:appservice/staticSiteCustomDomain:StaticSiteCustomDomain":
		r = &StaticSiteCustomDomain{}
	case "azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection":
		r = &VirtualNetworkSwiftConnection{}
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
		"appservice/activeSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/appService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificateBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificateOrder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/customHostnameBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/environmentV3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/functionApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/functionAppSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/hybridConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/managedCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/plan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/publicCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/slot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/slotCustomHostnameBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/slotVirtualNetworkSwiftConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/sourceCodeToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/staticSite",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/staticSiteCustomDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/virtualNetworkSwiftConnection",
		&module{version},
	)
}
