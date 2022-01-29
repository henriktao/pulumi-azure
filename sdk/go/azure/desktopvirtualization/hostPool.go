// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Virtual Desktop Host Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/hostPool:HostPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostpools/myhostpool
// ```
type HostPool struct {
	pulumi.CustomResourceState

	// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
	CustomRdpProperties pulumi.StringPtrOutput `pulumi:"customRdpProperties"`
	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrOutput `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrOutput `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrOutput `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrOutput `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Enables or disables the Start VM on Connection Feature. Defaults to `false`.
	StartVmOnConnect pulumi.BoolPtrOutput `pulumi:"startVmOnConnect"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// Allows you to test service changes before they are deployed to production. Defaults to `false`.
	ValidateEnvironment pulumi.BoolPtrOutput `pulumi:"validateEnvironment"`
}

// NewHostPool registers a new resource with the given unique name, arguments, and options.
func NewHostPool(ctx *pulumi.Context,
	name string, args *HostPoolArgs, opts ...pulumi.ResourceOption) (*HostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerType == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource HostPool
	err := ctx.RegisterResource("azure:desktopvirtualization/hostPool:HostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostPool gets an existing HostPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPoolState, opts ...pulumi.ResourceOption) (*HostPool, error) {
	var resource HostPool
	err := ctx.ReadResource("azure:desktopvirtualization/hostPool:HostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostPool resources.
type hostPoolState struct {
	// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
	CustomRdpProperties *string `pulumi:"customRdpProperties"`
	// A description for the Virtual Desktop Host Pool.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName *string `pulumi:"friendlyName"`
	// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed *int `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType *string `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType *string `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo *HostPoolRegistrationInfo `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Enables or disables the Start VM on Connection Feature. Defaults to `false`.
	StartVmOnConnect *bool `pulumi:"startVmOnConnect"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type *string `pulumi:"type"`
	// Allows you to test service changes before they are deployed to production. Defaults to `false`.
	ValidateEnvironment *bool `pulumi:"validateEnvironment"`
}

type HostPoolState struct {
	// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
	CustomRdpProperties pulumi.StringPtrInput
	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrInput
	// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringPtrInput
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrInput
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrInput
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Enables or disables the Start VM on Connection Feature. Defaults to `false`.
	StartVmOnConnect pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type pulumi.StringPtrInput
	// Allows you to test service changes before they are deployed to production. Defaults to `false`.
	ValidateEnvironment pulumi.BoolPtrInput
}

func (HostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolState)(nil)).Elem()
}

type hostPoolArgs struct {
	// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
	CustomRdpProperties *string `pulumi:"customRdpProperties"`
	// A description for the Virtual Desktop Host Pool.
	Description *string `pulumi:"description"`
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName *string `pulumi:"friendlyName"`
	// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType string `pulumi:"loadBalancerType"`
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed *int `pulumi:"maximumSessionsAllowed"`
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType *string `pulumi:"personalDesktopAssignmentType"`
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType *string `pulumi:"preferredAppGroupType"`
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo *HostPoolRegistrationInfo `pulumi:"registrationInfo"`
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enables or disables the Start VM on Connection Feature. Defaults to `false`.
	StartVmOnConnect *bool `pulumi:"startVmOnConnect"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type string `pulumi:"type"`
	// Allows you to test service changes before they are deployed to production. Defaults to `false`.
	ValidateEnvironment *bool `pulumi:"validateEnvironment"`
}

// The set of arguments for constructing a HostPool resource.
type HostPoolArgs struct {
	// A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
	CustomRdpProperties pulumi.StringPtrInput
	// A description for the Virtual Desktop Host Pool.
	Description pulumi.StringPtrInput
	// A friendly name for the Virtual Desktop Host Pool.
	FriendlyName pulumi.StringPtrInput
	// `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
	// `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
	// `Persistent` should be used if the host pool type is `Personal`
	LoadBalancerType pulumi.StringInput
	// The location/region where the Virtual Desktop Host Pool is
	// located. Changing the location/region forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
	// Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
	MaximumSessionsAllowed pulumi.IntPtrInput
	// The name of the Virtual Desktop Host Pool. Changing the name
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// `Automatic` assignment – The service will select an available host and assign it to an user.
	// `Direct` Assignment – Admin selects a specific host to assign to an user.
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	// Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
	// Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
	PreferredAppGroupType pulumi.StringPtrInput
	// A `registrationInfo` block which is documented below. Specifies configuration on the registration information of the Virtual Desktop Host Pool.
	RegistrationInfo HostPoolRegistrationInfoPtrInput
	// The name of the resource group in which to
	// create the Virtual Desktop Host Pool. Changing the resource group name forces
	// a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Enables or disables the Start VM on Connection Feature. Defaults to `false`.
	StartVmOnConnect pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the Virtual Desktop Host Pool. Valid options are
	// `Personal` or `Pooled`. Changing the type forces a new resource to be created.
	Type pulumi.StringInput
	// Allows you to test service changes before they are deployed to production. Defaults to `false`.
	ValidateEnvironment pulumi.BoolPtrInput
}

func (HostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolArgs)(nil)).Elem()
}

type HostPoolInput interface {
	pulumi.Input

	ToHostPoolOutput() HostPoolOutput
	ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput
}

func (*HostPool) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPool)(nil))
}

func (i *HostPool) ToHostPoolOutput() HostPoolOutput {
	return i.ToHostPoolOutputWithContext(context.Background())
}

func (i *HostPool) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolOutput)
}

func (i *HostPool) ToHostPoolPtrOutput() HostPoolPtrOutput {
	return i.ToHostPoolPtrOutputWithContext(context.Background())
}

func (i *HostPool) ToHostPoolPtrOutputWithContext(ctx context.Context) HostPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolPtrOutput)
}

type HostPoolPtrInput interface {
	pulumi.Input

	ToHostPoolPtrOutput() HostPoolPtrOutput
	ToHostPoolPtrOutputWithContext(ctx context.Context) HostPoolPtrOutput
}

type hostPoolPtrType HostPoolArgs

func (*hostPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil))
}

func (i *hostPoolPtrType) ToHostPoolPtrOutput() HostPoolPtrOutput {
	return i.ToHostPoolPtrOutputWithContext(context.Background())
}

func (i *hostPoolPtrType) ToHostPoolPtrOutputWithContext(ctx context.Context) HostPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolPtrOutput)
}

// HostPoolArrayInput is an input type that accepts HostPoolArray and HostPoolArrayOutput values.
// You can construct a concrete instance of `HostPoolArrayInput` via:
//
//          HostPoolArray{ HostPoolArgs{...} }
type HostPoolArrayInput interface {
	pulumi.Input

	ToHostPoolArrayOutput() HostPoolArrayOutput
	ToHostPoolArrayOutputWithContext(context.Context) HostPoolArrayOutput
}

type HostPoolArray []HostPoolInput

func (HostPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostPool)(nil)).Elem()
}

func (i HostPoolArray) ToHostPoolArrayOutput() HostPoolArrayOutput {
	return i.ToHostPoolArrayOutputWithContext(context.Background())
}

func (i HostPoolArray) ToHostPoolArrayOutputWithContext(ctx context.Context) HostPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolArrayOutput)
}

// HostPoolMapInput is an input type that accepts HostPoolMap and HostPoolMapOutput values.
// You can construct a concrete instance of `HostPoolMapInput` via:
//
//          HostPoolMap{ "key": HostPoolArgs{...} }
type HostPoolMapInput interface {
	pulumi.Input

	ToHostPoolMapOutput() HostPoolMapOutput
	ToHostPoolMapOutputWithContext(context.Context) HostPoolMapOutput
}

type HostPoolMap map[string]HostPoolInput

func (HostPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostPool)(nil)).Elem()
}

func (i HostPoolMap) ToHostPoolMapOutput() HostPoolMapOutput {
	return i.ToHostPoolMapOutputWithContext(context.Background())
}

func (i HostPoolMap) ToHostPoolMapOutputWithContext(ctx context.Context) HostPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolMapOutput)
}

type HostPoolOutput struct{ *pulumi.OutputState }

func (HostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPool)(nil))
}

func (o HostPoolOutput) ToHostPoolOutput() HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToHostPoolPtrOutput() HostPoolPtrOutput {
	return o.ToHostPoolPtrOutputWithContext(context.Background())
}

func (o HostPoolOutput) ToHostPoolPtrOutputWithContext(ctx context.Context) HostPoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostPool) *HostPool {
		return &v
	}).(HostPoolPtrOutput)
}

type HostPoolPtrOutput struct{ *pulumi.OutputState }

func (HostPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil))
}

func (o HostPoolPtrOutput) ToHostPoolPtrOutput() HostPoolPtrOutput {
	return o
}

func (o HostPoolPtrOutput) ToHostPoolPtrOutputWithContext(ctx context.Context) HostPoolPtrOutput {
	return o
}

func (o HostPoolPtrOutput) Elem() HostPoolOutput {
	return o.ApplyT(func(v *HostPool) HostPool {
		if v != nil {
			return *v
		}
		var ret HostPool
		return ret
	}).(HostPoolOutput)
}

type HostPoolArrayOutput struct{ *pulumi.OutputState }

func (HostPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostPool)(nil))
}

func (o HostPoolArrayOutput) ToHostPoolArrayOutput() HostPoolArrayOutput {
	return o
}

func (o HostPoolArrayOutput) ToHostPoolArrayOutputWithContext(ctx context.Context) HostPoolArrayOutput {
	return o
}

func (o HostPoolArrayOutput) Index(i pulumi.IntInput) HostPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostPool {
		return vs[0].([]HostPool)[vs[1].(int)]
	}).(HostPoolOutput)
}

type HostPoolMapOutput struct{ *pulumi.OutputState }

func (HostPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostPool)(nil))
}

func (o HostPoolMapOutput) ToHostPoolMapOutput() HostPoolMapOutput {
	return o
}

func (o HostPoolMapOutput) ToHostPoolMapOutputWithContext(ctx context.Context) HostPoolMapOutput {
	return o
}

func (o HostPoolMapOutput) MapIndex(k pulumi.StringInput) HostPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostPool {
		return vs[0].(map[string]HostPool)[vs[1].(string)]
	}).(HostPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostPoolInput)(nil)).Elem(), &HostPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostPoolPtrInput)(nil)).Elem(), &HostPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostPoolArrayInput)(nil)).Elem(), HostPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostPoolMapInput)(nil)).Elem(), HostPoolMap{})
	pulumi.RegisterOutputType(HostPoolOutput{})
	pulumi.RegisterOutputType(HostPoolPtrOutput{})
	pulumi.RegisterOutputType(HostPoolArrayOutput{})
	pulumi.RegisterOutputType(HostPoolMapOutput{})
}
