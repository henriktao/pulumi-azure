// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Application Insights WebTest.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
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
// 			Location:          pulumi.String("West Europe"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWebTest, err := appinsights.NewWebTest(ctx, "exampleWebTest", &appinsights.WebTestArgs{
// 			Location:              exampleInsights.Location,
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			Kind:                  pulumi.String("ping"),
// 			Frequency:             pulumi.Int(300),
// 			Timeout:               pulumi.Int(60),
// 			Enabled:               pulumi.Bool(true),
// 			GeoLocations: pulumi.StringArray{
// 				pulumi.String("us-tx-sn1-azr"),
// 				pulumi.String("us-il-ch1-azr"),
// 			},
// 			Configuration: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "<WebTest Name=\"WebTest1\" Id=\"ABD48585-0831-40CB-9069-682EA6BB3583\" Enabled=\"True\" CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"0\" WorkItemIds=\"\" xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\" CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\" Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\">\n", "  <Items>\n", "    <Request Method=\"GET\" Guid=\"a5f10126-e4cd-570d-961c-cea43999a200\" Version=\"1.1\" Url=\"http://microsoft.com\" ThinkTime=\"0\" Timeout=\"300\" ParseDependentRequests=\"True\" FollowRedirects=\"True\" RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\" ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\" IgnoreHttpStatusCode=\"False\" />\n", "  </Items>\n", "</WebTest>\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("webtestId", exampleWebTest.ID())
// 		ctx.Export("webtestsSyntheticId", exampleWebTest.SyntheticMonitorId)
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Application Insights Web Tests can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appinsights/webTest:WebTest my_test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/webTests/my_test
// ```
type WebTest struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest ([see here for more information](https://docs.microsoft.com/en-us/rest/api/application-insights/webtests/createorupdate/)).
	Configuration pulumi.StringOutput `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Valid options are `300`, `600` and `900`. Defaults to `300`.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayOutput `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`. Changing this forces a new resource to be created.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled       pulumi.BoolPtrOutput `pulumi:"retryEnabled"`
	SyntheticMonitorId pulumi.StringOutput  `pulumi:"syntheticMonitorId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewWebTest registers a new resource with the given unique name, arguments, and options.
func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationInsightsId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationInsightsId'")
	}
	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.GeoLocations == nil {
		return nil, errors.New("invalid value for required argument 'GeoLocations'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource WebTest
	err := ctx.RegisterResource("azure:appinsights/webTest:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTest gets an existing WebTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azure:appinsights/webTest:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTest resources.
type webTestState struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest ([see here for more information](https://docs.microsoft.com/en-us/rest/api/application-insights/webtests/createorupdate/)).
	Configuration *string `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Valid options are `300`, `600` and `900`. Defaults to `300`.
	Frequency *int `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []string `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled       *bool   `pulumi:"retryEnabled"`
	SyntheticMonitorId *string `pulumi:"syntheticMonitorId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout *int `pulumi:"timeout"`
}

type WebTestState struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// An XML configuration specification for a WebTest ([see here for more information](https://docs.microsoft.com/en-us/rest/api/application-insights/webtests/createorupdate/)).
	Configuration pulumi.StringPtrInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Valid options are `300`, `600` and `900`. Defaults to `300`.
	Frequency pulumi.IntPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayInput
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// Allow for retries should this WebTest fail.
	RetryEnabled       pulumi.BoolPtrInput
	SyntheticMonitorId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrInput
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// An XML configuration specification for a WebTest ([see here for more information](https://docs.microsoft.com/en-us/rest/api/application-insights/webtests/createorupdate/)).
	Configuration string `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Valid options are `300`, `600` and `900`. Defaults to `300`.
	Frequency *int `pulumi:"frequency"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations []string `pulumi:"geoLocations"`
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`. Changing this forces a new resource to be created.
	Kind string `pulumi:"kind"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location *string `pulumi:"location"`
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a WebTest resource.
type WebTestArgs struct {
	// The ID of the Application Insights component on which the WebTest operates. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// An XML configuration specification for a WebTest ([see here for more information](https://docs.microsoft.com/en-us/rest/api/application-insights/webtests/createorupdate/)).
	Configuration pulumi.StringInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Valid options are `300`, `600` and `900`. Defaults to `300`.
	Frequency pulumi.IntPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	GeoLocations pulumi.StringArrayInput
	// = (Required) The kind of web test that this web test watches. Choices are `ping` and `multistep`. Changing this forces a new resource to be created.
	Kind pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. It needs to correlate with location of parent resource (azurerm_application_insights).
	Location pulumi.StringPtrInput
	// Specifies the name of the Application Insights WebTest. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Application Insights WebTest. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default is `30`.
	Timeout pulumi.IntPtrInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}

type WebTestInput interface {
	pulumi.Input

	ToWebTestOutput() WebTestOutput
	ToWebTestOutputWithContext(ctx context.Context) WebTestOutput
}

func (*WebTest) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTest)(nil))
}

func (i *WebTest) ToWebTestOutput() WebTestOutput {
	return i.ToWebTestOutputWithContext(context.Background())
}

func (i *WebTest) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestOutput)
}

func (i *WebTest) ToWebTestPtrOutput() WebTestPtrOutput {
	return i.ToWebTestPtrOutputWithContext(context.Background())
}

func (i *WebTest) ToWebTestPtrOutputWithContext(ctx context.Context) WebTestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPtrOutput)
}

type WebTestPtrInput interface {
	pulumi.Input

	ToWebTestPtrOutput() WebTestPtrOutput
	ToWebTestPtrOutputWithContext(ctx context.Context) WebTestPtrOutput
}

type webTestPtrType WebTestArgs

func (*webTestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil))
}

func (i *webTestPtrType) ToWebTestPtrOutput() WebTestPtrOutput {
	return i.ToWebTestPtrOutputWithContext(context.Background())
}

func (i *webTestPtrType) ToWebTestPtrOutputWithContext(ctx context.Context) WebTestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPtrOutput)
}

// WebTestArrayInput is an input type that accepts WebTestArray and WebTestArrayOutput values.
// You can construct a concrete instance of `WebTestArrayInput` via:
//
//          WebTestArray{ WebTestArgs{...} }
type WebTestArrayInput interface {
	pulumi.Input

	ToWebTestArrayOutput() WebTestArrayOutput
	ToWebTestArrayOutputWithContext(context.Context) WebTestArrayOutput
}

type WebTestArray []WebTestInput

func (WebTestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebTest)(nil)).Elem()
}

func (i WebTestArray) ToWebTestArrayOutput() WebTestArrayOutput {
	return i.ToWebTestArrayOutputWithContext(context.Background())
}

func (i WebTestArray) ToWebTestArrayOutputWithContext(ctx context.Context) WebTestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestArrayOutput)
}

// WebTestMapInput is an input type that accepts WebTestMap and WebTestMapOutput values.
// You can construct a concrete instance of `WebTestMapInput` via:
//
//          WebTestMap{ "key": WebTestArgs{...} }
type WebTestMapInput interface {
	pulumi.Input

	ToWebTestMapOutput() WebTestMapOutput
	ToWebTestMapOutputWithContext(context.Context) WebTestMapOutput
}

type WebTestMap map[string]WebTestInput

func (WebTestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebTest)(nil)).Elem()
}

func (i WebTestMap) ToWebTestMapOutput() WebTestMapOutput {
	return i.ToWebTestMapOutputWithContext(context.Background())
}

func (i WebTestMap) ToWebTestMapOutputWithContext(ctx context.Context) WebTestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestMapOutput)
}

type WebTestOutput struct{ *pulumi.OutputState }

func (WebTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTest)(nil))
}

func (o WebTestOutput) ToWebTestOutput() WebTestOutput {
	return o
}

func (o WebTestOutput) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return o
}

func (o WebTestOutput) ToWebTestPtrOutput() WebTestPtrOutput {
	return o.ToWebTestPtrOutputWithContext(context.Background())
}

func (o WebTestOutput) ToWebTestPtrOutputWithContext(ctx context.Context) WebTestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTest) *WebTest {
		return &v
	}).(WebTestPtrOutput)
}

type WebTestPtrOutput struct{ *pulumi.OutputState }

func (WebTestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil))
}

func (o WebTestPtrOutput) ToWebTestPtrOutput() WebTestPtrOutput {
	return o
}

func (o WebTestPtrOutput) ToWebTestPtrOutputWithContext(ctx context.Context) WebTestPtrOutput {
	return o
}

func (o WebTestPtrOutput) Elem() WebTestOutput {
	return o.ApplyT(func(v *WebTest) WebTest {
		if v != nil {
			return *v
		}
		var ret WebTest
		return ret
	}).(WebTestOutput)
}

type WebTestArrayOutput struct{ *pulumi.OutputState }

func (WebTestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTest)(nil))
}

func (o WebTestArrayOutput) ToWebTestArrayOutput() WebTestArrayOutput {
	return o
}

func (o WebTestArrayOutput) ToWebTestArrayOutputWithContext(ctx context.Context) WebTestArrayOutput {
	return o
}

func (o WebTestArrayOutput) Index(i pulumi.IntInput) WebTestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTest {
		return vs[0].([]WebTest)[vs[1].(int)]
	}).(WebTestOutput)
}

type WebTestMapOutput struct{ *pulumi.OutputState }

func (WebTestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebTest)(nil))
}

func (o WebTestMapOutput) ToWebTestMapOutput() WebTestMapOutput {
	return o
}

func (o WebTestMapOutput) ToWebTestMapOutputWithContext(ctx context.Context) WebTestMapOutput {
	return o
}

func (o WebTestMapOutput) MapIndex(k pulumi.StringInput) WebTestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebTest {
		return vs[0].(map[string]WebTest)[vs[1].(string)]
	}).(WebTestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestInput)(nil)).Elem(), &WebTest{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestPtrInput)(nil)).Elem(), &WebTest{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestArrayInput)(nil)).Elem(), WebTestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebTestMapInput)(nil)).Elem(), WebTestMap{})
	pulumi.RegisterOutputType(WebTestOutput{})
	pulumi.RegisterOutputType(WebTestPtrOutput{})
	pulumi.RegisterOutputType(WebTestArrayOutput{})
	pulumi.RegisterOutputType(WebTestMapOutput{})
}
