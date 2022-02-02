// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Virtual Desktop Application Group.
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
// 		pooledbreadthfirst, err := desktopvirtualization.NewHostPool(ctx, "pooledbreadthfirst", &desktopvirtualization.HostPoolArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			Type:              pulumi.String("Pooled"),
// 			LoadBalancerType:  pulumi.String("BreadthFirst"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		personalautomatic, err := desktopvirtualization.NewHostPool(ctx, "personalautomatic", &desktopvirtualization.HostPoolArgs{
// 			Location:                      example.Location,
// 			ResourceGroupName:             example.Name,
// 			Type:                          pulumi.String("Personal"),
// 			PersonalDesktopAssignmentType: pulumi.String("Automatic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewApplicationGroup(ctx, "remoteapp", &desktopvirtualization.ApplicationGroupArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			Type:              pulumi.String("RemoteApp"),
// 			HostPoolId:        pooledbreadthfirst.ID(),
// 			FriendlyName:      pulumi.String("TestAppGroup"),
// 			Description:       pulumi.String("Acceptance Test: An application group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewApplicationGroup(ctx, "desktopapp", &desktopvirtualization.ApplicationGroupArgs{
// 			Location:          example.Location,
// 			ResourceGroupName: example.Name,
// 			Type:              pulumi.String("Desktop"),
// 			HostPoolId:        personalautomatic.ID(),
// 			FriendlyName:      pulumi.String("TestAppGroup"),
// 			Description:       pulumi.String("Acceptance Test: An application group"),
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
// Virtual Desktop Application Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/applicationGroup:ApplicationGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup
// ```
type ApplicationGroup struct {
	pulumi.CustomResourceState

	// Option to set the display name for the default sessionDesktop desktop when `type` is set to `Desktop`.
	DefaultDesktopDisplayName pulumi.StringPtrOutput `pulumi:"defaultDesktopDisplayName"`
	// Option to set a description for the Virtual Desktop Application Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Option to set a friendly name for the Virtual Desktop Application Group.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Resource ID for a Virtual Desktop Host Pool to associate with the
	// Virtual Desktop Application Group.
	HostPoolId pulumi.StringOutput `pulumi:"hostPoolId"`
	// The location/region where the Virtual Desktop Application Group is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Application Group. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of Virtual Desktop Application Group.
	// Valid options are `RemoteApp` or `Desktop` application groups.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolId == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure:desktopvirtualization/applicationGroup:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGroup gets an existing ApplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure:desktopvirtualization/applicationGroup:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGroup resources.
type applicationGroupState struct {
	// Option to set the display name for the default sessionDesktop desktop when `type` is set to `Desktop`.
	DefaultDesktopDisplayName *string `pulumi:"defaultDesktopDisplayName"`
	// Option to set a description for the Virtual Desktop Application Group.
	Description *string `pulumi:"description"`
	// Option to set a friendly name for the Virtual Desktop Application Group.
	FriendlyName *string `pulumi:"friendlyName"`
	// Resource ID for a Virtual Desktop Host Pool to associate with the
	// Virtual Desktop Application Group.
	HostPoolId *string `pulumi:"hostPoolId"`
	// The location/region where the Virtual Desktop Application Group is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Application Group. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of Virtual Desktop Application Group.
	// Valid options are `RemoteApp` or `Desktop` application groups.
	Type *string `pulumi:"type"`
}

type ApplicationGroupState struct {
	// Option to set the display name for the default sessionDesktop desktop when `type` is set to `Desktop`.
	DefaultDesktopDisplayName pulumi.StringPtrInput
	// Option to set a description for the Virtual Desktop Application Group.
	Description pulumi.StringPtrInput
	// Option to set a friendly name for the Virtual Desktop Application Group.
	FriendlyName pulumi.StringPtrInput
	// Resource ID for a Virtual Desktop Host Pool to associate with the
	// Virtual Desktop Application Group.
	HostPoolId pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Application Group is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Application Group. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Type of Virtual Desktop Application Group.
	// Valid options are `RemoteApp` or `Desktop` application groups.
	Type pulumi.StringPtrInput
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	// Option to set the display name for the default sessionDesktop desktop when `type` is set to `Desktop`.
	DefaultDesktopDisplayName *string `pulumi:"defaultDesktopDisplayName"`
	// Option to set a description for the Virtual Desktop Application Group.
	Description *string `pulumi:"description"`
	// Option to set a friendly name for the Virtual Desktop Application Group.
	FriendlyName *string `pulumi:"friendlyName"`
	// Resource ID for a Virtual Desktop Host Pool to associate with the
	// Virtual Desktop Application Group.
	HostPoolId string `pulumi:"hostPoolId"`
	// The location/region where the Virtual Desktop Application Group is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the Virtual Desktop Application Group. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of Virtual Desktop Application Group.
	// Valid options are `RemoteApp` or `Desktop` application groups.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ApplicationGroup resource.
type ApplicationGroupArgs struct {
	// Option to set the display name for the default sessionDesktop desktop when `type` is set to `Desktop`.
	DefaultDesktopDisplayName pulumi.StringPtrInput
	// Option to set a description for the Virtual Desktop Application Group.
	Description pulumi.StringPtrInput
	// Option to set a friendly name for the Virtual Desktop Application Group.
	FriendlyName pulumi.StringPtrInput
	// Resource ID for a Virtual Desktop Host Pool to associate with the
	// Virtual Desktop Application Group.
	HostPoolId pulumi.StringInput
	// The location/region where the Virtual Desktop Application Group is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Virtual Desktop Application Group. Changing the name forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Application Group. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Type of Virtual Desktop Application Group.
	// Valid options are `RemoteApp` or `Desktop` application groups.
	Type pulumi.StringInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

// ApplicationGroupArrayInput is an input type that accepts ApplicationGroupArray and ApplicationGroupArrayOutput values.
// You can construct a concrete instance of `ApplicationGroupArrayInput` via:
//
//          ApplicationGroupArray{ ApplicationGroupArgs{...} }
type ApplicationGroupArrayInput interface {
	pulumi.Input

	ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput
	ToApplicationGroupArrayOutputWithContext(context.Context) ApplicationGroupArrayOutput
}

type ApplicationGroupArray []ApplicationGroupInput

func (ApplicationGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return i.ToApplicationGroupArrayOutputWithContext(context.Background())
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupArrayOutput)
}

// ApplicationGroupMapInput is an input type that accepts ApplicationGroupMap and ApplicationGroupMapOutput values.
// You can construct a concrete instance of `ApplicationGroupMapInput` via:
//
//          ApplicationGroupMap{ "key": ApplicationGroupArgs{...} }
type ApplicationGroupMapInput interface {
	pulumi.Input

	ToApplicationGroupMapOutput() ApplicationGroupMapOutput
	ToApplicationGroupMapOutputWithContext(context.Context) ApplicationGroupMapOutput
}

type ApplicationGroupMap map[string]ApplicationGroupInput

func (ApplicationGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return i.ToApplicationGroupMapOutputWithContext(context.Background())
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupMapOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

type ApplicationGroupArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) Index(i pulumi.IntInput) ApplicationGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].([]*ApplicationGroup)[vs[1].(int)]
	}).(ApplicationGroupOutput)
}

type ApplicationGroupMapOutput struct{ *pulumi.OutputState }

func (ApplicationGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) MapIndex(k pulumi.StringInput) ApplicationGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].(map[string]*ApplicationGroup)[vs[1].(string)]
	}).(ApplicationGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupInput)(nil)).Elem(), &ApplicationGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupArrayInput)(nil)).Elem(), ApplicationGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupMapInput)(nil)).Elem(), ApplicationGroupMap{})
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
	pulumi.RegisterOutputType(ApplicationGroupArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGroupMapOutput{})
}
