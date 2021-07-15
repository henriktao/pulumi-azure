// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

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
	case "azure:monitoring/aadDiagnosticSetting:AadDiagnosticSetting":
		r = &AadDiagnosticSetting{}
	case "azure:monitoring/actionGroup:ActionGroup":
		r = &ActionGroup{}
	case "azure:monitoring/actionRuleActionGroup:ActionRuleActionGroup":
		r = &ActionRuleActionGroup{}
	case "azure:monitoring/actionRuleSuppression:ActionRuleSuppression":
		r = &ActionRuleSuppression{}
	case "azure:monitoring/activityLogAlert:ActivityLogAlert":
		r = &ActivityLogAlert{}
	case "azure:monitoring/autoscaleSetting:AutoscaleSetting":
		r = &AutoscaleSetting{}
	case "azure:monitoring/diagnosticSetting:DiagnosticSetting":
		r = &DiagnosticSetting{}
	case "azure:monitoring/logProfile:LogProfile":
		r = &LogProfile{}
	case "azure:monitoring/metricAlert:MetricAlert":
		r = &MetricAlert{}
	case "azure:monitoring/scheduledQueryRulesAlert:ScheduledQueryRulesAlert":
		r = &ScheduledQueryRulesAlert{}
	case "azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog":
		r = &ScheduledQueryRulesLog{}
	case "azure:monitoring/smartDetectorAlertRule:SmartDetectorAlertRule":
		r = &SmartDetectorAlertRule{}
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
		"monitoring/aadDiagnosticSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/actionGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/actionRuleActionGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/actionRuleSuppression",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/activityLogAlert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/autoscaleSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/diagnosticSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/logProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/metricAlert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/scheduledQueryRulesAlert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/scheduledQueryRulesLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"monitoring/smartDetectorAlertRule",
		&module{version},
	)
}
