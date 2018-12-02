// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.
func LookupDiagnosticCategories(ctx *pulumi.Context, args *GetDiagnosticCategoriesArgs) (*GetDiagnosticCategoriesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["resourceId"] = args.ResourceId
	}
	outputs, err := ctx.Invoke("azure:monitoring/getDiagnosticCategories:getDiagnosticCategories", inputs)
	if err != nil {
		return nil, err
	}
	return &GetDiagnosticCategoriesResult{
		Logs: outputs["logs"],
		Metrics: outputs["metrics"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getDiagnosticCategories.
type GetDiagnosticCategoriesArgs struct {
	// The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
	ResourceId interface{}
}

// A collection of values returned by getDiagnosticCategories.
type GetDiagnosticCategoriesResult struct {
	// A list of the Log Categories supported for this Resource.
	Logs interface{}
	// A list of the Metric Categories supported for this Resource.
	Metrics interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
