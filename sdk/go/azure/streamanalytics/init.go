// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

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
	case "azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF":
		r = &FunctionJavaScriptUDF{}
	case "azure:streamanalytics/job:Job":
		r = &Job{}
	case "azure:streamanalytics/outputBlob:OutputBlob":
		r = &OutputBlob{}
	case "azure:streamanalytics/outputEventHub:OutputEventHub":
		r = &OutputEventHub{}
	case "azure:streamanalytics/outputMssql:OutputMssql":
		r = &OutputMssql{}
	case "azure:streamanalytics/outputServiceBusQueue:OutputServiceBusQueue":
		r = &OutputServiceBusQueue{}
	case "azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic":
		r = &OutputServicebusTopic{}
	case "azure:streamanalytics/referenceInputBlob:ReferenceInputBlob":
		r = &ReferenceInputBlob{}
	case "azure:streamanalytics/streamInputBlob:StreamInputBlob":
		r = &StreamInputBlob{}
	case "azure:streamanalytics/streamInputEventHub:StreamInputEventHub":
		r = &StreamInputEventHub{}
	case "azure:streamanalytics/streamInputIotHub:StreamInputIotHub":
		r = &StreamInputIotHub{}
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
		"streamanalytics/functionJavaScriptUDF",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/job",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputEventHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputMssql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputServiceBusQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/outputServicebusTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/referenceInputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputEventHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"streamanalytics/streamInputIotHub",
		&module{version},
	)
}
