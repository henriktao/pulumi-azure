// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

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
	case "azure:iot/certificate:Certificate":
		r = &Certificate{}
	case "azure:iot/consumerGroup:ConsumerGroup":
		r = &ConsumerGroup{}
	case "azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy":
		r = &DpsSharedAccessPolicy{}
	case "azure:iot/endpointEventhub:EndpointEventhub":
		r = &EndpointEventhub{}
	case "azure:iot/endpointServicebusQueue:EndpointServicebusQueue":
		r = &EndpointServicebusQueue{}
	case "azure:iot/endpointServicebusTopic:EndpointServicebusTopic":
		r = &EndpointServicebusTopic{}
	case "azure:iot/endpointStorageContainer:EndpointStorageContainer":
		r = &EndpointStorageContainer{}
	case "azure:iot/enrichment:Enrichment":
		r = &Enrichment{}
	case "azure:iot/fallbackRoute:FallbackRoute":
		r = &FallbackRoute{}
	case "azure:iot/ioTHub:IoTHub":
		r = &IoTHub{}
	case "azure:iot/iotHubCertificate:IotHubCertificate":
		r = &IotHubCertificate{}
	case "azure:iot/iotHubDps:IotHubDps":
		r = &IotHubDps{}
	case "azure:iot/route:Route":
		r = &Route{}
	case "azure:iot/securityDeviceGroup:SecurityDeviceGroup":
		r = &SecurityDeviceGroup{}
	case "azure:iot/securitySolution:SecuritySolution":
		r = &SecuritySolution{}
	case "azure:iot/sharedAccessPolicy:SharedAccessPolicy":
		r = &SharedAccessPolicy{}
	case "azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy":
		r = &TimeSeriesInsightsAccessPolicy{}
	case "azure:iot/timeSeriesInsightsEventSourceEventhub:TimeSeriesInsightsEventSourceEventhub":
		r = &TimeSeriesInsightsEventSourceEventhub{}
	case "azure:iot/timeSeriesInsightsEventSourceIothub:TimeSeriesInsightsEventSourceIothub":
		r = &TimeSeriesInsightsEventSourceIothub{}
	case "azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment":
		r = &TimeSeriesInsightsGen2Environment{}
	case "azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet":
		r = &TimeSeriesInsightsReferenceDataSet{}
	case "azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment":
		r = &TimeSeriesInsightsStandardEnvironment{}
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
		"iot/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/consumerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/dpsSharedAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointEventhub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointServicebusQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointServicebusTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointStorageContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/enrichment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/fallbackRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/ioTHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/iotHubCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/iotHubDps",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/securityDeviceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/securitySolution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/sharedAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsEventSourceEventhub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsEventSourceIothub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsGen2Environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsReferenceDataSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsStandardEnvironment",
		&module{version},
	)
}
