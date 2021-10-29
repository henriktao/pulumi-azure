// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

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
	case "azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration":
		r = &AttachedDatabaseConfiguration{}
	case "azure:kusto/cluster:Cluster":
		r = &Cluster{}
	case "azure:kusto/clusterCustomerManagedKey:ClusterCustomerManagedKey":
		r = &ClusterCustomerManagedKey{}
	case "azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment":
		r = &ClusterPrincipalAssignment{}
	case "azure:kusto/database:Database":
		r = &Database{}
	case "azure:kusto/databasePrincipal:DatabasePrincipal":
		r = &DatabasePrincipal{}
	case "azure:kusto/databasePrincipalAssignment:DatabasePrincipalAssignment":
		r = &DatabasePrincipalAssignment{}
	case "azure:kusto/eventGridDataConnection:EventGridDataConnection":
		r = &EventGridDataConnection{}
	case "azure:kusto/eventhubDataConnection:EventhubDataConnection":
		r = &EventhubDataConnection{}
	case "azure:kusto/iotHubDataConnection:IotHubDataConnection":
		r = &IotHubDataConnection{}
	case "azure:kusto/script:Script":
		r = &Script{}
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
		"kusto/attachedDatabaseConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/clusterCustomerManagedKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/clusterPrincipalAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/databasePrincipal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/databasePrincipalAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/eventGridDataConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/eventhubDataConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/iotHubDataConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"kusto/script",
		&module{version},
	)
}
