// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Machine Learning Workspace can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:machinelearning/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrOutput `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The url for the discovery service to identify regional endpoints for machine learning experimentation services.
	DiscoveryUrl pulumi.StringOutput          `pulumi:"discoveryUrl"`
	Encryption   WorkspaceEncryptionPtrOutput `pulumi:"encryption"`
	// Display name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact pulumi.BoolPtrOutput `pulumi:"highBusinessImpact"`
	// An `identity` block as defined below.
	Identity WorkspaceIdentityOutput `pulumi:"identity"`
	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName pulumi.StringPtrOutput `pulumi:"imageBuildComputeName"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationInsightsId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationInsightsId'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:machinelearning/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure:machinelearning/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description *string `pulumi:"description"`
	// The url for the discovery service to identify regional endpoints for machine learning experimentation services.
	DiscoveryUrl *string              `pulumi:"discoveryUrl"`
	Encryption   *WorkspaceEncryption `pulumi:"encryption"`
	// Display name for this Machine Learning Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact *bool `pulumi:"highBusinessImpact"`
	// An `identity` block as defined below.
	Identity *WorkspaceIdentity `pulumi:"identity"`
	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName *string `pulumi:"imageBuildComputeName"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
	SkuName *string `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type WorkspaceState struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrInput
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrInput
	// The url for the discovery service to identify regional endpoints for machine learning experimentation services.
	DiscoveryUrl pulumi.StringPtrInput
	Encryption   WorkspaceEncryptionPtrInput
	// Display name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrInput
	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity WorkspaceIdentityPtrInput
	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName pulumi.StringPtrInput
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
	SkuName pulumi.StringPtrInput
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description *string              `pulumi:"description"`
	Encryption  *WorkspaceEncryption `pulumi:"encryption"`
	// Display name for this Machine Learning Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact *bool `pulumi:"highBusinessImpact"`
	// An `identity` block as defined below.
	Identity WorkspaceIdentity `pulumi:"identity"`
	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName *string `pulumi:"imageBuildComputeName"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
	SkuName *string `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrInput
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrInput
	Encryption  WorkspaceEncryptionPtrInput
	// Display name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrInput
	// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
	HighBusinessImpact pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity WorkspaceIdentityInput
	// The compute name for image build of the Machine Learning Workspace.
	ImageBuildComputeName pulumi.StringPtrInput
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Enable public access when this Machine Learning Workspace is behind VNet.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
	SkuName pulumi.StringPtrInput
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

// WorkspaceArrayInput is an input type that accepts WorkspaceArray and WorkspaceArrayOutput values.
// You can construct a concrete instance of `WorkspaceArrayInput` via:
//
//          WorkspaceArray{ WorkspaceArgs{...} }
type WorkspaceArrayInput interface {
	pulumi.Input

	ToWorkspaceArrayOutput() WorkspaceArrayOutput
	ToWorkspaceArrayOutputWithContext(context.Context) WorkspaceArrayOutput
}

type WorkspaceArray []WorkspaceInput

func (WorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (i WorkspaceArray) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return i.ToWorkspaceArrayOutputWithContext(context.Background())
}

func (i WorkspaceArray) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceArrayOutput)
}

// WorkspaceMapInput is an input type that accepts WorkspaceMap and WorkspaceMapOutput values.
// You can construct a concrete instance of `WorkspaceMapInput` via:
//
//          WorkspaceMap{ "key": WorkspaceArgs{...} }
type WorkspaceMapInput interface {
	pulumi.Input

	ToWorkspaceMapOutput() WorkspaceMapOutput
	ToWorkspaceMapOutputWithContext(context.Context) WorkspaceMapOutput
}

type WorkspaceMap map[string]WorkspaceInput

func (WorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (i WorkspaceMap) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return i.ToWorkspaceMapOutputWithContext(context.Background())
}

func (i WorkspaceMap) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceMapOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].([]*Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].(map[string]*Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceArrayInput)(nil)).Elem(), WorkspaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceMapInput)(nil)).Elem(), WorkspaceMap{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
