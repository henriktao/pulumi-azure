// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Virtual Desktop Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/desktopvirtualization"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewWorkspace(ctx, "workspace", &desktopvirtualization.WorkspaceArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			FriendlyName:      pulumi.String("FriendlyName"),
// 			Description:       pulumi.String("A description of my workspace"),
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
// Virtual Desktop Workspaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:desktopvirtualization/workspace:Workspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:desktopvirtualization/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// A description for the Virtual Desktop Workspace.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type WorkspaceState struct {
	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// A description for the Virtual Desktop Workspace.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// A description for the Virtual Desktop Workspace.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Workspace.
	FriendlyName pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Workspace is located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Workspace. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Workspace. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringInput
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
	return reflect.TypeOf((*Workspace)(nil))
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

type WorkspacePtrInput interface {
	pulumi.Input

	ToWorkspacePtrOutput() WorkspacePtrOutput
	ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput
}

type workspacePtrType WorkspaceArgs

func (*workspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (i *workspacePtrType) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *workspacePtrType) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
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
	return reflect.TypeOf((*Workspace)(nil))
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o.ToWorkspacePtrOutputWithContext(context.Background())
}

func (o WorkspaceOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Workspace) *Workspace {
		return &v
	}).(WorkspacePtrOutput)
}

type WorkspacePtrOutput struct{ *pulumi.OutputState }

func (WorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o
}

func (o WorkspacePtrOutput) Elem() WorkspaceOutput {
	return o.ApplyT(func(v *Workspace) Workspace {
		if v != nil {
			return *v
		}
		var ret Workspace
		return ret
	}).(WorkspaceOutput)
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Workspace)(nil))
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].([]Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Workspace)(nil))
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].(map[string]Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspacePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
