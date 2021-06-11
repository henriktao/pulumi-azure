// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

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
	case "azure:postgresql/activeDirectoryAdministrator:ActiveDirectoryAdministrator":
		r = &ActiveDirectoryAdministrator{}
	case "azure:postgresql/configuration:Configuration":
		r = &Configuration{}
	case "azure:postgresql/database:Database":
		r = &Database{}
	case "azure:postgresql/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "azure:postgresql/flexibleServer:FlexibleServer":
		r = &FlexibleServer{}
	case "azure:postgresql/flexibleServerFirewallRule:FlexibleServerFirewallRule":
		r = &FlexibleServerFirewallRule{}
	case "azure:postgresql/server:Server":
		r = &Server{}
	case "azure:postgresql/serverKey:ServerKey":
		r = &ServerKey{}
	case "azure:postgresql/virtualNetworkRule:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/activeDirectoryAdministrator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/configuration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/flexibleServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/flexibleServerFirewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/serverKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"postgresql/virtualNetworkRule",
		&module{version},
	)
}
