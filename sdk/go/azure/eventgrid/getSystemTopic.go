// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing EventGrid System Topic
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eventgrid.LookupSystemTopic(ctx, &eventgrid.LookupSystemTopicArgs{
// 			Name:              "eventgrid-system-topic",
// 			ResourceGroupName: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSystemTopic(ctx *pulumi.Context, args *LookupSystemTopicArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicResult, error) {
	var rv LookupSystemTopicResult
	err := ctx.Invoke("azure:eventgrid/getSystemTopic:getSystemTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemTopic.
type LookupSystemTopicArgs struct {
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
	Identity *GetSystemTopicIdentity `pulumi:"identity"`
	// The name of the EventGrid System Topic resource.
	Name string `pulumi:"name"`
	// The name of the resource group in which the EventGrid System Topic exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSystemTopic.
type LookupSystemTopicResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
	Identity GetSystemTopicIdentity `pulumi:"identity"`
	Location string                 `pulumi:"location"`
	// The Metric ARM Resource ID of the Event Grid System Topic.
	MetricArmResourceId string `pulumi:"metricArmResourceId"`
	Name                string `pulumi:"name"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	// The ID of the Event Grid System Topic ARM Source.
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
	// A mapping of tags which are assigned to the Event Grid System Topic.
	Tags map[string]string `pulumi:"tags"`
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType string `pulumi:"topicType"`
}

func LookupSystemTopicOutput(ctx *pulumi.Context, args LookupSystemTopicOutputArgs, opts ...pulumi.InvokeOption) LookupSystemTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemTopicResult, error) {
			args := v.(LookupSystemTopicArgs)
			r, err := LookupSystemTopic(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemTopicResultOutput)
}

// A collection of arguments for invoking getSystemTopic.
type LookupSystemTopicOutputArgs struct {
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
	Identity GetSystemTopicIdentityPtrInput `pulumi:"identity"`
	// The name of the EventGrid System Topic resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group in which the EventGrid System Topic exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSystemTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicArgs)(nil)).Elem()
}

// A collection of values returned by getSystemTopic.
type LookupSystemTopicResultOutput struct{ *pulumi.OutputState }

func (LookupSystemTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicResult)(nil)).Elem()
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutput() LookupSystemTopicResultOutput {
	return o
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutputWithContext(ctx context.Context) LookupSystemTopicResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
func (o LookupSystemTopicResultOutput) Identity() GetSystemTopicIdentityOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) GetSystemTopicIdentity { return v.Identity }).(GetSystemTopicIdentityOutput)
}

func (o LookupSystemTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

// The Metric ARM Resource ID of the Event Grid System Topic.
func (o LookupSystemTopicResultOutput) MetricArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.MetricArmResourceId }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemTopicResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// The ID of the Event Grid System Topic ARM Source.
func (o LookupSystemTopicResultOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

// A mapping of tags which are assigned to the Event Grid System Topic.
func (o LookupSystemTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
func (o LookupSystemTopicResultOutput) TopicType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.TopicType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemTopicResultOutput{})
}
