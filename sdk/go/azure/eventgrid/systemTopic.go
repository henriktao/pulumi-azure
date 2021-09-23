// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Event Grid System Topic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewSystemTopic(ctx, "exampleSystemTopic", &eventgrid.SystemTopicArgs{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Location:            exampleResourceGroup.Location,
// 			SourceArmResourceId: exampleAccount.ID(),
// 			TopicType:           pulumi.String("Microsoft.Storage.StorageAccounts"),
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
// Event Grid System Topic can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventgrid/systemTopic:SystemTopic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/systemTopics/systemTopic1
// ```
type SystemTopicResource struct {
	pulumi.CustomResourceState

	// An `identity` block as defined below.
	Identity SystemTopicIdentityPtrOutput `pulumi:"identity"`
	// The Azure Region where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Metric ARM Resource ID of the Event Grid System Topic.
	MetricArmResourceId pulumi.StringOutput `pulumi:"metricArmResourceId"`
	// The name which should be used for this Event Grid System Topic. Changing this forces a new Event Grid System Topic to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Event Grid System Topic ARM Source. Changing this forces a new Event Grid System Topic to be created.
	SourceArmResourceId pulumi.StringOutput `pulumi:"sourceArmResourceId"`
	// A mapping of tags which should be assigned to the Event Grid System Topic.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType pulumi.StringOutput `pulumi:"topicType"`
}

// NewSystemTopicResource registers a new resource with the given unique name, arguments, and options.
func NewSystemTopicResource(ctx *pulumi.Context,
	name string, args *SystemTopicResourceArgs, opts ...pulumi.ResourceOption) (*SystemTopicResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceArmResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceArmResourceId'")
	}
	if args.TopicType == nil {
		return nil, errors.New("invalid value for required argument 'TopicType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventgrid/getSystemTopic:getSystemTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopicResource
	err := ctx.RegisterResource("azure:eventgrid/systemTopic:SystemTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemTopicResource gets an existing SystemTopicResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemTopicResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicResourceState, opts ...pulumi.ResourceOption) (*SystemTopicResource, error) {
	var resource SystemTopicResource
	err := ctx.ReadResource("azure:eventgrid/systemTopic:SystemTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemTopicResource resources.
type systemTopicResourceState struct {
	// An `identity` block as defined below.
	Identity *SystemTopicIdentity `pulumi:"identity"`
	// The Azure Region where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	Location *string `pulumi:"location"`
	// The Metric ARM Resource ID of the Event Grid System Topic.
	MetricArmResourceId *string `pulumi:"metricArmResourceId"`
	// The name which should be used for this Event Grid System Topic. Changing this forces a new Event Grid System Topic to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Event Grid System Topic ARM Source. Changing this forces a new Event Grid System Topic to be created.
	SourceArmResourceId *string `pulumi:"sourceArmResourceId"`
	// A mapping of tags which should be assigned to the Event Grid System Topic.
	Tags map[string]string `pulumi:"tags"`
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType *string `pulumi:"topicType"`
}

type SystemTopicResourceState struct {
	// An `identity` block as defined below.
	Identity SystemTopicIdentityPtrInput
	// The Azure Region where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	Location pulumi.StringPtrInput
	// The Metric ARM Resource ID of the Event Grid System Topic.
	MetricArmResourceId pulumi.StringPtrInput
	// The name which should be used for this Event Grid System Topic. Changing this forces a new Event Grid System Topic to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Event Grid System Topic ARM Source. Changing this forces a new Event Grid System Topic to be created.
	SourceArmResourceId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Event Grid System Topic.
	Tags pulumi.StringMapInput
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType pulumi.StringPtrInput
}

func (SystemTopicResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicResourceState)(nil)).Elem()
}

type systemTopicResourceArgs struct {
	// An `identity` block as defined below.
	Identity *SystemTopicIdentity `pulumi:"identity"`
	// The Azure Region where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Event Grid System Topic. Changing this forces a new Event Grid System Topic to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Event Grid System Topic ARM Source. Changing this forces a new Event Grid System Topic to be created.
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
	// A mapping of tags which should be assigned to the Event Grid System Topic.
	Tags map[string]string `pulumi:"tags"`
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType string `pulumi:"topicType"`
}

// The set of arguments for constructing a SystemTopicResource resource.
type SystemTopicResourceArgs struct {
	// An `identity` block as defined below.
	Identity SystemTopicIdentityPtrInput
	// The Azure Region where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Event Grid System Topic. Changing this forces a new Event Grid System Topic to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Event Grid System Topic should exist. Changing this forces a new Event Grid System Topic to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Event Grid System Topic ARM Source. Changing this forces a new Event Grid System Topic to be created.
	SourceArmResourceId pulumi.StringInput
	// A mapping of tags which should be assigned to the Event Grid System Topic.
	Tags pulumi.StringMapInput
	// The Topic Type of the Event Grid System Topic. Possible values are: `Microsoft.AppConfiguration.ConfigurationStores`, `Microsoft.Communication.CommunicationServices`
	// , `Microsoft.ContainerRegistry.Registries`, `Microsoft.Devices.IoTHubs`, `Microsoft.EventGrid.Domains`, `Microsoft.EventGrid.Topics`, `Microsoft.Eventhub.Namespaces`, `Microsoft.KeyVault.vaults`, `Microsoft.MachineLearningServices.Workspaces`, `Microsoft.Maps.Accounts`, `Microsoft.Media.MediaServices`, `Microsoft.Resources.ResourceGroups`, `Microsoft.Resources.Subscriptions`, `Microsoft.ServiceBus.Namespaces`, `Microsoft.SignalRService.SignalR`, `Microsoft.Storage.StorageAccounts`, `Microsoft.Web.ServerFarms` and `Microsoft.Web.Sites`. Changing this forces a new Event Grid System Topic to be created.
	TopicType pulumi.StringInput
}

func (SystemTopicResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicResourceArgs)(nil)).Elem()
}

type SystemTopicResourceInput interface {
	pulumi.Input

	ToSystemTopicResourceOutput() SystemTopicResourceOutput
	ToSystemTopicResourceOutputWithContext(ctx context.Context) SystemTopicResourceOutput
}

func (*SystemTopicResource) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicResource)(nil))
}

func (i *SystemTopicResource) ToSystemTopicResourceOutput() SystemTopicResourceOutput {
	return i.ToSystemTopicResourceOutputWithContext(context.Background())
}

func (i *SystemTopicResource) ToSystemTopicResourceOutputWithContext(ctx context.Context) SystemTopicResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicResourceOutput)
}

func (i *SystemTopicResource) ToSystemTopicResourcePtrOutput() SystemTopicResourcePtrOutput {
	return i.ToSystemTopicResourcePtrOutputWithContext(context.Background())
}

func (i *SystemTopicResource) ToSystemTopicResourcePtrOutputWithContext(ctx context.Context) SystemTopicResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicResourcePtrOutput)
}

type SystemTopicResourcePtrInput interface {
	pulumi.Input

	ToSystemTopicResourcePtrOutput() SystemTopicResourcePtrOutput
	ToSystemTopicResourcePtrOutputWithContext(ctx context.Context) SystemTopicResourcePtrOutput
}

type systemTopicResourcePtrType SystemTopicResourceArgs

func (*systemTopicResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicResource)(nil))
}

func (i *systemTopicResourcePtrType) ToSystemTopicResourcePtrOutput() SystemTopicResourcePtrOutput {
	return i.ToSystemTopicResourcePtrOutputWithContext(context.Background())
}

func (i *systemTopicResourcePtrType) ToSystemTopicResourcePtrOutputWithContext(ctx context.Context) SystemTopicResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicResourcePtrOutput)
}

// SystemTopicResourceArrayInput is an input type that accepts SystemTopicResourceArray and SystemTopicResourceArrayOutput values.
// You can construct a concrete instance of `SystemTopicResourceArrayInput` via:
//
//          SystemTopicResourceArray{ SystemTopicResourceArgs{...} }
type SystemTopicResourceArrayInput interface {
	pulumi.Input

	ToSystemTopicResourceArrayOutput() SystemTopicResourceArrayOutput
	ToSystemTopicResourceArrayOutputWithContext(context.Context) SystemTopicResourceArrayOutput
}

type SystemTopicResourceArray []SystemTopicResourceInput

func (SystemTopicResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemTopicResource)(nil)).Elem()
}

func (i SystemTopicResourceArray) ToSystemTopicResourceArrayOutput() SystemTopicResourceArrayOutput {
	return i.ToSystemTopicResourceArrayOutputWithContext(context.Background())
}

func (i SystemTopicResourceArray) ToSystemTopicResourceArrayOutputWithContext(ctx context.Context) SystemTopicResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicResourceArrayOutput)
}

// SystemTopicResourceMapInput is an input type that accepts SystemTopicResourceMap and SystemTopicResourceMapOutput values.
// You can construct a concrete instance of `SystemTopicResourceMapInput` via:
//
//          SystemTopicResourceMap{ "key": SystemTopicResourceArgs{...} }
type SystemTopicResourceMapInput interface {
	pulumi.Input

	ToSystemTopicResourceMapOutput() SystemTopicResourceMapOutput
	ToSystemTopicResourceMapOutputWithContext(context.Context) SystemTopicResourceMapOutput
}

type SystemTopicResourceMap map[string]SystemTopicResourceInput

func (SystemTopicResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemTopicResource)(nil)).Elem()
}

func (i SystemTopicResourceMap) ToSystemTopicResourceMapOutput() SystemTopicResourceMapOutput {
	return i.ToSystemTopicResourceMapOutputWithContext(context.Background())
}

func (i SystemTopicResourceMap) ToSystemTopicResourceMapOutputWithContext(ctx context.Context) SystemTopicResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicResourceMapOutput)
}

type SystemTopicResourceOutput struct{ *pulumi.OutputState }

func (SystemTopicResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicResource)(nil))
}

func (o SystemTopicResourceOutput) ToSystemTopicResourceOutput() SystemTopicResourceOutput {
	return o
}

func (o SystemTopicResourceOutput) ToSystemTopicResourceOutputWithContext(ctx context.Context) SystemTopicResourceOutput {
	return o
}

func (o SystemTopicResourceOutput) ToSystemTopicResourcePtrOutput() SystemTopicResourcePtrOutput {
	return o.ToSystemTopicResourcePtrOutputWithContext(context.Background())
}

func (o SystemTopicResourceOutput) ToSystemTopicResourcePtrOutputWithContext(ctx context.Context) SystemTopicResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemTopicResource) *SystemTopicResource {
		return &v
	}).(SystemTopicResourcePtrOutput)
}

type SystemTopicResourcePtrOutput struct{ *pulumi.OutputState }

func (SystemTopicResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicResource)(nil))
}

func (o SystemTopicResourcePtrOutput) ToSystemTopicResourcePtrOutput() SystemTopicResourcePtrOutput {
	return o
}

func (o SystemTopicResourcePtrOutput) ToSystemTopicResourcePtrOutputWithContext(ctx context.Context) SystemTopicResourcePtrOutput {
	return o
}

func (o SystemTopicResourcePtrOutput) Elem() SystemTopicResourceOutput {
	return o.ApplyT(func(v *SystemTopicResource) SystemTopicResource {
		if v != nil {
			return *v
		}
		var ret SystemTopicResource
		return ret
	}).(SystemTopicResourceOutput)
}

type SystemTopicResourceArrayOutput struct{ *pulumi.OutputState }

func (SystemTopicResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemTopicResource)(nil))
}

func (o SystemTopicResourceArrayOutput) ToSystemTopicResourceArrayOutput() SystemTopicResourceArrayOutput {
	return o
}

func (o SystemTopicResourceArrayOutput) ToSystemTopicResourceArrayOutputWithContext(ctx context.Context) SystemTopicResourceArrayOutput {
	return o
}

func (o SystemTopicResourceArrayOutput) Index(i pulumi.IntInput) SystemTopicResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemTopicResource {
		return vs[0].([]SystemTopicResource)[vs[1].(int)]
	}).(SystemTopicResourceOutput)
}

type SystemTopicResourceMapOutput struct{ *pulumi.OutputState }

func (SystemTopicResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemTopicResource)(nil))
}

func (o SystemTopicResourceMapOutput) ToSystemTopicResourceMapOutput() SystemTopicResourceMapOutput {
	return o
}

func (o SystemTopicResourceMapOutput) ToSystemTopicResourceMapOutputWithContext(ctx context.Context) SystemTopicResourceMapOutput {
	return o
}

func (o SystemTopicResourceMapOutput) MapIndex(k pulumi.StringInput) SystemTopicResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemTopicResource {
		return vs[0].(map[string]SystemTopicResource)[vs[1].(string)]
	}).(SystemTopicResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemTopicResourceOutput{})
	pulumi.RegisterOutputType(SystemTopicResourcePtrOutput{})
	pulumi.RegisterOutputType(SystemTopicResourceArrayOutput{})
	pulumi.RegisterOutputType(SystemTopicResourceMapOutput{})
}
