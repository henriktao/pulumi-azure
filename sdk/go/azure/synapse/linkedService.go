// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Linked Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/synapse"
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
// 			AccountKind:            pulumi.String("BlobStorage"),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkspace, err := synapse.NewWorkspace(ctx, "exampleWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 			ManagedVirtualNetworkEnabled:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFirewallRule, err := synapse.NewFirewallRule(ctx, "exampleFirewallRule", &synapse.FirewallRuleArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			StartIpAddress:     pulumi.String("0.0.0.0"),
// 			EndIpAddress:       pulumi.String("255.255.255.255"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewLinkedService(ctx, "exampleLinkedService", &synapse.LinkedServiceArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			Type:               pulumi.String("AzureBlobStorage"),
// 			TypePropertiesJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"connectionString\": \"", azurerm_storage_account.Test.Primary_connection_string, "\"\n", "}\n")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleFirewallRule,
// 		}))
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
// Synapse Linked Services can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/linkedService:LinkedService example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/linkedservices/linkedservice1
// ```
type LinkedService struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The description for the Synapse Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A `integrationRuntime` block as defined below.
	IntegrationRuntime LinkedServiceIntegrationRuntimePtrOutput `pulumi:"integrationRuntime"`
	// The name which should be used for this Synapse Linked Service. Changing this forces a new Synapse Linked Service to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Synapse Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
	// The type of data stores that will be connected to Synapse. For full list of supported data stores, please refer to [Azure Synapse connector](https://docs.microsoft.com/en-us/azure/data-factory/connector-overview). Changing this forces a new Synapse Linked Service to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJson pulumi.StringOutput `pulumi:"typePropertiesJson"`
}

// NewLinkedService registers a new resource with the given unique name, arguments, and options.
func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.TypePropertiesJson == nil {
		return nil, errors.New("invalid value for required argument 'TypePropertiesJson'")
	}
	var resource LinkedService
	err := ctx.RegisterResource("azure:synapse/linkedService:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedService gets an existing LinkedService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure:synapse/linkedService:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedService resources.
type linkedServiceState struct {
	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The description for the Synapse Linked Service.
	Description *string `pulumi:"description"`
	// A `integrationRuntime` block as defined below.
	IntegrationRuntime *LinkedServiceIntegrationRuntime `pulumi:"integrationRuntime"`
	// The name which should be used for this Synapse Linked Service. Changing this forces a new Synapse Linked Service to be created.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Synapse Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
	// The type of data stores that will be connected to Synapse. For full list of supported data stores, please refer to [Azure Synapse connector](https://docs.microsoft.com/en-us/azure/data-factory/connector-overview). Changing this forces a new Synapse Linked Service to be created.
	Type *string `pulumi:"type"`
	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJson *string `pulumi:"typePropertiesJson"`
}

type LinkedServiceState struct {
	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations pulumi.StringArrayInput
	// The description for the Synapse Linked Service.
	Description pulumi.StringPtrInput
	// A `integrationRuntime` block as defined below.
	IntegrationRuntime LinkedServiceIntegrationRuntimePtrInput
	// The name which should be used for this Synapse Linked Service. Changing this forces a new Synapse Linked Service to be created.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Synapse Linked Service.
	Parameters pulumi.StringMapInput
	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
	// The type of data stores that will be connected to Synapse. For full list of supported data stores, please refer to [Azure Synapse connector](https://docs.microsoft.com/en-us/azure/data-factory/connector-overview). Changing this forces a new Synapse Linked Service to be created.
	Type pulumi.StringPtrInput
	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJson pulumi.StringPtrInput
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The description for the Synapse Linked Service.
	Description *string `pulumi:"description"`
	// A `integrationRuntime` block as defined below.
	IntegrationRuntime *LinkedServiceIntegrationRuntime `pulumi:"integrationRuntime"`
	// The name which should be used for this Synapse Linked Service. Changing this forces a new Synapse Linked Service to be created.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Synapse Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
	// The type of data stores that will be connected to Synapse. For full list of supported data stores, please refer to [Azure Synapse connector](https://docs.microsoft.com/en-us/azure/data-factory/connector-overview). Changing this forces a new Synapse Linked Service to be created.
	Type string `pulumi:"type"`
	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJson string `pulumi:"typePropertiesJson"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// A map of additional properties to associate with the Synapse Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Synapse Linked Service.
	Annotations pulumi.StringArrayInput
	// The description for the Synapse Linked Service.
	Description pulumi.StringPtrInput
	// A `integrationRuntime` block as defined below.
	IntegrationRuntime LinkedServiceIntegrationRuntimePtrInput
	// The name which should be used for this Synapse Linked Service. Changing this forces a new Synapse Linked Service to be created.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Synapse Linked Service.
	Parameters pulumi.StringMapInput
	// The Synapse Workspace ID in which to associate the Linked Service with. Changing this forces a new Synapse Linked Service to be created.
	SynapseWorkspaceId pulumi.StringInput
	// The type of data stores that will be connected to Synapse. For full list of supported data stores, please refer to [Azure Synapse connector](https://docs.microsoft.com/en-us/azure/data-factory/connector-overview). Changing this forces a new Synapse Linked Service to be created.
	Type pulumi.StringInput
	// A JSON object that contains the properties of the Synapse Linked Service.
	TypePropertiesJson pulumi.StringInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}

type LinkedServiceInput interface {
	pulumi.Input

	ToLinkedServiceOutput() LinkedServiceOutput
	ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput
}

func (*LinkedService) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedService)(nil))
}

func (i *LinkedService) ToLinkedServiceOutput() LinkedServiceOutput {
	return i.ToLinkedServiceOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOutput)
}

func (i *LinkedService) ToLinkedServicePtrOutput() LinkedServicePtrOutput {
	return i.ToLinkedServicePtrOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServicePtrOutputWithContext(ctx context.Context) LinkedServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePtrOutput)
}

type LinkedServicePtrInput interface {
	pulumi.Input

	ToLinkedServicePtrOutput() LinkedServicePtrOutput
	ToLinkedServicePtrOutputWithContext(ctx context.Context) LinkedServicePtrOutput
}

type linkedServicePtrType LinkedServiceArgs

func (*linkedServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil))
}

func (i *linkedServicePtrType) ToLinkedServicePtrOutput() LinkedServicePtrOutput {
	return i.ToLinkedServicePtrOutputWithContext(context.Background())
}

func (i *linkedServicePtrType) ToLinkedServicePtrOutputWithContext(ctx context.Context) LinkedServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePtrOutput)
}

// LinkedServiceArrayInput is an input type that accepts LinkedServiceArray and LinkedServiceArrayOutput values.
// You can construct a concrete instance of `LinkedServiceArrayInput` via:
//
//          LinkedServiceArray{ LinkedServiceArgs{...} }
type LinkedServiceArrayInput interface {
	pulumi.Input

	ToLinkedServiceArrayOutput() LinkedServiceArrayOutput
	ToLinkedServiceArrayOutputWithContext(context.Context) LinkedServiceArrayOutput
}

type LinkedServiceArray []LinkedServiceInput

func (LinkedServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedService)(nil))
}

func (i LinkedServiceArray) ToLinkedServiceArrayOutput() LinkedServiceArrayOutput {
	return i.ToLinkedServiceArrayOutputWithContext(context.Background())
}

func (i LinkedServiceArray) ToLinkedServiceArrayOutputWithContext(ctx context.Context) LinkedServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceArrayOutput)
}

// LinkedServiceMapInput is an input type that accepts LinkedServiceMap and LinkedServiceMapOutput values.
// You can construct a concrete instance of `LinkedServiceMapInput` via:
//
//          LinkedServiceMap{ "key": LinkedServiceArgs{...} }
type LinkedServiceMapInput interface {
	pulumi.Input

	ToLinkedServiceMapOutput() LinkedServiceMapOutput
	ToLinkedServiceMapOutputWithContext(context.Context) LinkedServiceMapOutput
}

type LinkedServiceMap map[string]LinkedServiceInput

func (LinkedServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedService)(nil))
}

func (i LinkedServiceMap) ToLinkedServiceMapOutput() LinkedServiceMapOutput {
	return i.ToLinkedServiceMapOutputWithContext(context.Background())
}

func (i LinkedServiceMap) ToLinkedServiceMapOutputWithContext(ctx context.Context) LinkedServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceMapOutput)
}

type LinkedServiceOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedService)(nil))
}

func (o LinkedServiceOutput) ToLinkedServiceOutput() LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToLinkedServicePtrOutput() LinkedServicePtrOutput {
	return o.ToLinkedServicePtrOutputWithContext(context.Background())
}

func (o LinkedServiceOutput) ToLinkedServicePtrOutputWithContext(ctx context.Context) LinkedServicePtrOutput {
	return o.ApplyT(func(v LinkedService) *LinkedService {
		return &v
	}).(LinkedServicePtrOutput)
}

type LinkedServicePtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil))
}

func (o LinkedServicePtrOutput) ToLinkedServicePtrOutput() LinkedServicePtrOutput {
	return o
}

func (o LinkedServicePtrOutput) ToLinkedServicePtrOutputWithContext(ctx context.Context) LinkedServicePtrOutput {
	return o
}

type LinkedServiceArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedService)(nil))
}

func (o LinkedServiceArrayOutput) ToLinkedServiceArrayOutput() LinkedServiceArrayOutput {
	return o
}

func (o LinkedServiceArrayOutput) ToLinkedServiceArrayOutputWithContext(ctx context.Context) LinkedServiceArrayOutput {
	return o
}

func (o LinkedServiceArrayOutput) Index(i pulumi.IntInput) LinkedServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedService {
		return vs[0].([]LinkedService)[vs[1].(int)]
	}).(LinkedServiceOutput)
}

type LinkedServiceMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedService)(nil))
}

func (o LinkedServiceMapOutput) ToLinkedServiceMapOutput() LinkedServiceMapOutput {
	return o
}

func (o LinkedServiceMapOutput) ToLinkedServiceMapOutputWithContext(ctx context.Context) LinkedServiceMapOutput {
	return o
}

func (o LinkedServiceMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedService {
		return vs[0].(map[string]LinkedService)[vs[1].(string)]
	}).(LinkedServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
	pulumi.RegisterOutputType(LinkedServicePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceMapOutput{})
}
