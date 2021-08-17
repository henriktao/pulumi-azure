// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Management Service Diagnostic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("My Company"),
// 			PublisherEmail:    pulumi.String("company@mycompany.io"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogger, err := apimanagement.NewLogger(ctx, "exampleLogger", &apimanagement.LoggerArgs{
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationInsights: &apimanagement.LoggerApplicationInsightsArgs{
// 				InstrumentationKey: exampleInsights.InstrumentationKey,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewDiagnostic(ctx, "exampleDiagnostic", &apimanagement.DiagnosticArgs{
// 			Identifier:              pulumi.String("applicationinsights"),
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			ApiManagementName:       exampleService.Name,
// 			ApiManagementLoggerId:   exampleLogger.ID(),
// 			SamplingPercentage:      pulumi.Float64(5),
// 			AlwaysLogErrors:         pulumi.Bool(true),
// 			LogClientIp:             pulumi.Bool(true),
// 			Verbosity:               pulumi.String("verbose"),
// 			HttpCorrelationProtocol: pulumi.String("W3C"),
// 			FrontendRequest: &apimanagement.DiagnosticFrontendRequestArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("accept"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			FrontendResponse: &apimanagement.DiagnosticFrontendResponseArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("content-length"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			BackendRequest: &apimanagement.DiagnosticBackendRequestArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("accept"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			BackendResponse: &apimanagement.DiagnosticBackendResponseArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("content-length"),
// 					pulumi.String("origin"),
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
// API Management Diagnostics can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/diagnostic:Diagnostic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/diagnostics/applicationinsights
// ```
type Diagnostic struct {
	pulumi.CustomResourceState

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolOutput `pulumi:"alwaysLogErrors"`
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringOutput `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// A `backendRequest` block as defined below.
	BackendRequest DiagnosticBackendRequestOutput `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse DiagnosticBackendResponseOutput `pulumi:"backendResponse"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A `frontendRequest` block as defined below.
	FrontendRequest DiagnosticFrontendRequestOutput `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse DiagnosticFrontendResponseOutput `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringOutput `pulumi:"httpCorrelationProtocol"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp pulumi.BoolOutput `pulumi:"logClientIp"`
	// The format of the Operation Name for Application Insights telemetries. Possible values are `Name`, and `Url`. Defaults to `Name`.
	OperationNameFormat pulumi.StringPtrOutput `pulumi:"operationNameFormat"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64Output `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringOutput `pulumi:"verbosity"`
}

// NewDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementLoggerId == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementLoggerId'")
	}
	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Diagnostic
	err := ctx.RegisterResource("azure:apimanagement/diagnostic:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnostic gets an existing Diagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azure:apimanagement/diagnostic:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Diagnostic resources.
type diagnosticState struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `pulumi:"alwaysLogErrors"`
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId *string `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// A `backendRequest` block as defined below.
	BackendRequest *DiagnosticBackendRequest `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse *DiagnosticBackendResponse `pulumi:"backendResponse"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	Enabled *bool `pulumi:"enabled"`
	// A `frontendRequest` block as defined below.
	FrontendRequest *DiagnosticFrontendRequest `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse *DiagnosticFrontendResponse `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier *string `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp *bool `pulumi:"logClientIp"`
	// The format of the Operation Name for Application Insights telemetries. Possible values are `Name`, and `Url`. Defaults to `Name`.
	OperationNameFormat *string `pulumi:"operationNameFormat"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity *string `pulumi:"verbosity"`
}

type DiagnosticState struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolPtrInput
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringPtrInput
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// A `backendRequest` block as defined below.
	BackendRequest DiagnosticBackendRequestPtrInput
	// A `backendResponse` block as defined below.
	BackendResponse DiagnosticBackendResponsePtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	Enabled pulumi.BoolPtrInput
	// A `frontendRequest` block as defined below.
	FrontendRequest DiagnosticFrontendRequestPtrInput
	// A `frontendResponse` block as defined below.
	FrontendResponse DiagnosticFrontendResponsePtrInput
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringPtrInput
	// Log client IP address.
	LogClientIp pulumi.BoolPtrInput
	// The format of the Operation Name for Application Insights telemetries. Possible values are `Name`, and `Url`. Defaults to `Name`.
	OperationNameFormat pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64PtrInput
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringPtrInput
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `pulumi:"alwaysLogErrors"`
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId string `pulumi:"apiManagementLoggerId"`
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// A `backendRequest` block as defined below.
	BackendRequest *DiagnosticBackendRequest `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse *DiagnosticBackendResponse `pulumi:"backendResponse"`
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	Enabled *bool `pulumi:"enabled"`
	// A `frontendRequest` block as defined below.
	FrontendRequest *DiagnosticFrontendRequest `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse *DiagnosticFrontendResponse `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier string `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp *bool `pulumi:"logClientIp"`
	// The format of the Operation Name for Application Insights telemetries. Possible values are `Name`, and `Url`. Defaults to `Name`.
	OperationNameFormat *string `pulumi:"operationNameFormat"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity *string `pulumi:"verbosity"`
}

// The set of arguments for constructing a Diagnostic resource.
type DiagnosticArgs struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolPtrInput
	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	ApiManagementLoggerId pulumi.StringInput
	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// A `backendRequest` block as defined below.
	BackendRequest DiagnosticBackendRequestPtrInput
	// A `backendResponse` block as defined below.
	BackendResponse DiagnosticBackendResponsePtrInput
	// Deprecated: this property has been removed from the API and will be removed in version 3.0 of the provider
	Enabled pulumi.BoolPtrInput
	// A `frontendRequest` block as defined below.
	FrontendRequest DiagnosticFrontendRequestPtrInput
	// A `frontendResponse` block as defined below.
	FrontendResponse DiagnosticFrontendResponsePtrInput
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
	Identifier pulumi.StringInput
	// Log client IP address.
	LogClientIp pulumi.BoolPtrInput
	// The format of the Operation Name for Application Insights telemetries. Possible values are `Name`, and `Url`. Defaults to `Name`.
	OperationNameFormat pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64PtrInput
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringPtrInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}

type DiagnosticInput interface {
	pulumi.Input

	ToDiagnosticOutput() DiagnosticOutput
	ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput
}

func (*Diagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((*Diagnostic)(nil))
}

func (i *Diagnostic) ToDiagnosticOutput() DiagnosticOutput {
	return i.ToDiagnosticOutputWithContext(context.Background())
}

func (i *Diagnostic) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticOutput)
}

func (i *Diagnostic) ToDiagnosticPtrOutput() DiagnosticPtrOutput {
	return i.ToDiagnosticPtrOutputWithContext(context.Background())
}

func (i *Diagnostic) ToDiagnosticPtrOutputWithContext(ctx context.Context) DiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticPtrOutput)
}

type DiagnosticPtrInput interface {
	pulumi.Input

	ToDiagnosticPtrOutput() DiagnosticPtrOutput
	ToDiagnosticPtrOutputWithContext(ctx context.Context) DiagnosticPtrOutput
}

type diagnosticPtrType DiagnosticArgs

func (*diagnosticPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil))
}

func (i *diagnosticPtrType) ToDiagnosticPtrOutput() DiagnosticPtrOutput {
	return i.ToDiagnosticPtrOutputWithContext(context.Background())
}

func (i *diagnosticPtrType) ToDiagnosticPtrOutputWithContext(ctx context.Context) DiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticPtrOutput)
}

// DiagnosticArrayInput is an input type that accepts DiagnosticArray and DiagnosticArrayOutput values.
// You can construct a concrete instance of `DiagnosticArrayInput` via:
//
//          DiagnosticArray{ DiagnosticArgs{...} }
type DiagnosticArrayInput interface {
	pulumi.Input

	ToDiagnosticArrayOutput() DiagnosticArrayOutput
	ToDiagnosticArrayOutputWithContext(context.Context) DiagnosticArrayOutput
}

type DiagnosticArray []DiagnosticInput

func (DiagnosticArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Diagnostic)(nil))
}

func (i DiagnosticArray) ToDiagnosticArrayOutput() DiagnosticArrayOutput {
	return i.ToDiagnosticArrayOutputWithContext(context.Background())
}

func (i DiagnosticArray) ToDiagnosticArrayOutputWithContext(ctx context.Context) DiagnosticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticArrayOutput)
}

// DiagnosticMapInput is an input type that accepts DiagnosticMap and DiagnosticMapOutput values.
// You can construct a concrete instance of `DiagnosticMapInput` via:
//
//          DiagnosticMap{ "key": DiagnosticArgs{...} }
type DiagnosticMapInput interface {
	pulumi.Input

	ToDiagnosticMapOutput() DiagnosticMapOutput
	ToDiagnosticMapOutputWithContext(context.Context) DiagnosticMapOutput
}

type DiagnosticMap map[string]DiagnosticInput

func (DiagnosticMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Diagnostic)(nil))
}

func (i DiagnosticMap) ToDiagnosticMapOutput() DiagnosticMapOutput {
	return i.ToDiagnosticMapOutputWithContext(context.Background())
}

func (i DiagnosticMap) ToDiagnosticMapOutputWithContext(ctx context.Context) DiagnosticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticMapOutput)
}

type DiagnosticOutput struct {
	*pulumi.OutputState
}

func (DiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Diagnostic)(nil))
}

func (o DiagnosticOutput) ToDiagnosticOutput() DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) ToDiagnosticPtrOutput() DiagnosticPtrOutput {
	return o.ToDiagnosticPtrOutputWithContext(context.Background())
}

func (o DiagnosticOutput) ToDiagnosticPtrOutputWithContext(ctx context.Context) DiagnosticPtrOutput {
	return o.ApplyT(func(v Diagnostic) *Diagnostic {
		return &v
	}).(DiagnosticPtrOutput)
}

type DiagnosticPtrOutput struct {
	*pulumi.OutputState
}

func (DiagnosticPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil))
}

func (o DiagnosticPtrOutput) ToDiagnosticPtrOutput() DiagnosticPtrOutput {
	return o
}

func (o DiagnosticPtrOutput) ToDiagnosticPtrOutputWithContext(ctx context.Context) DiagnosticPtrOutput {
	return o
}

type DiagnosticArrayOutput struct{ *pulumi.OutputState }

func (DiagnosticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Diagnostic)(nil))
}

func (o DiagnosticArrayOutput) ToDiagnosticArrayOutput() DiagnosticArrayOutput {
	return o
}

func (o DiagnosticArrayOutput) ToDiagnosticArrayOutputWithContext(ctx context.Context) DiagnosticArrayOutput {
	return o
}

func (o DiagnosticArrayOutput) Index(i pulumi.IntInput) DiagnosticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Diagnostic {
		return vs[0].([]Diagnostic)[vs[1].(int)]
	}).(DiagnosticOutput)
}

type DiagnosticMapOutput struct{ *pulumi.OutputState }

func (DiagnosticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Diagnostic)(nil))
}

func (o DiagnosticMapOutput) ToDiagnosticMapOutput() DiagnosticMapOutput {
	return o
}

func (o DiagnosticMapOutput) ToDiagnosticMapOutputWithContext(ctx context.Context) DiagnosticMapOutput {
	return o
}

func (o DiagnosticMapOutput) MapIndex(k pulumi.StringInput) DiagnosticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Diagnostic {
		return vs[0].(map[string]Diagnostic)[vs[1].(string)]
	}).(DiagnosticOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticOutput{})
	pulumi.RegisterOutputType(DiagnosticPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticMapOutput{})
}
