// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package domainservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Domain Services can be imported using the resource ID, together with the Replica Set ID that you wish to designate as the initial replica set, e.g.
//
// ```sh
//  $ pulumi import azure:domainservices/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/initialReplicaSetId/00000000-0000-0000-0000-000000000000
// ```
type Service struct {
	pulumi.CustomResourceState

	// A unique ID for the managed domain deployment.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
	FilteredSyncEnabled pulumi.BoolPtrOutput `pulumi:"filteredSyncEnabled"`
	// An `initialReplicaSet` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
	InitialReplicaSet ServiceInitialReplicaSetOutput `pulumi:"initialReplicaSet"`
	// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `notifications` block as defined below.
	Notifications ServiceNotificationsOutput `pulumi:"notifications"`
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Azure resource ID for the domain service.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// A `secureLdap` block as defined below.
	SecureLdap ServiceSecureLdapOutput `pulumi:"secureLdap"`
	// A `security` block as defined below.
	Security ServiceSecurityOutput `pulumi:"security"`
	// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku       pulumi.StringOutput `pulumi:"sku"`
	SyncOwner pulumi.StringOutput `pulumi:"syncOwner"`
	// A mapping of tags assigned to the resource.
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	TenantId pulumi.StringOutput    `pulumi:"tenantId"`
	Version  pulumi.IntOutput       `pulumi:"version"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.InitialReplicaSet == nil {
		return nil, errors.New("invalid value for required argument 'InitialReplicaSet'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:domainservices/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:domainservices/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// A unique ID for the managed domain deployment.
	DeploymentId *string `pulumi:"deploymentId"`
	// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName *string `pulumi:"domainName"`
	// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
	FilteredSyncEnabled *bool `pulumi:"filteredSyncEnabled"`
	// An `initialReplicaSet` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
	InitialReplicaSet *ServiceInitialReplicaSet `pulumi:"initialReplicaSet"`
	// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `notifications` block as defined below.
	Notifications *ServiceNotifications `pulumi:"notifications"`
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Azure resource ID for the domain service.
	ResourceId *string `pulumi:"resourceId"`
	// A `secureLdap` block as defined below.
	SecureLdap *ServiceSecureLdap `pulumi:"secureLdap"`
	// A `security` block as defined below.
	Security *ServiceSecurity `pulumi:"security"`
	// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku       *string `pulumi:"sku"`
	SyncOwner *string `pulumi:"syncOwner"`
	// A mapping of tags assigned to the resource.
	Tags     map[string]string `pulumi:"tags"`
	TenantId *string           `pulumi:"tenantId"`
	Version  *int              `pulumi:"version"`
}

type ServiceState struct {
	// A unique ID for the managed domain deployment.
	DeploymentId pulumi.StringPtrInput
	// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName pulumi.StringPtrInput
	// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
	FilteredSyncEnabled pulumi.BoolPtrInput
	// An `initialReplicaSet` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
	InitialReplicaSet ServiceInitialReplicaSetPtrInput
	// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `notifications` block as defined below.
	Notifications ServiceNotificationsPtrInput
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Azure resource ID for the domain service.
	ResourceId pulumi.StringPtrInput
	// A `secureLdap` block as defined below.
	SecureLdap ServiceSecureLdapPtrInput
	// A `security` block as defined below.
	Security ServiceSecurityPtrInput
	// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku       pulumi.StringPtrInput
	SyncOwner pulumi.StringPtrInput
	// A mapping of tags assigned to the resource.
	Tags     pulumi.StringMapInput
	TenantId pulumi.StringPtrInput
	Version  pulumi.IntPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName string `pulumi:"domainName"`
	// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
	FilteredSyncEnabled *bool `pulumi:"filteredSyncEnabled"`
	// An `initialReplicaSet` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
	InitialReplicaSet ServiceInitialReplicaSet `pulumi:"initialReplicaSet"`
	// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `notifications` block as defined below.
	Notifications *ServiceNotifications `pulumi:"notifications"`
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `secureLdap` block as defined below.
	SecureLdap *ServiceSecureLdap `pulumi:"secureLdap"`
	// A `security` block as defined below.
	Security *ServiceSecurity `pulumi:"security"`
	// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
	DomainName pulumi.StringInput
	// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
	FilteredSyncEnabled pulumi.BoolPtrInput
	// An `initialReplicaSet` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
	InitialReplicaSet ServiceInitialReplicaSetInput
	// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `notifications` block as defined below.
	Notifications ServiceNotificationsPtrInput
	// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `secureLdap` block as defined below.
	SecureLdap ServiceSecureLdapPtrInput
	// A `security` block as defined below.
	Security ServiceSecurityPtrInput
	// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
	Sku pulumi.StringInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

func (i *Service) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *Service) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

type ServicePtrInput interface {
	pulumi.Input

	ToServicePtrOutput() ServicePtrOutput
	ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput
}

type servicePtrType ServiceArgs

func (*servicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (i *servicePtrType) ToServicePtrOutput() ServicePtrOutput {
	return i.ToServicePtrOutputWithContext(context.Background())
}

func (i *servicePtrType) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePtrOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) ToServicePtrOutput() ServicePtrOutput {
	return o.ToServicePtrOutputWithContext(context.Background())
}

func (o ServiceOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Service) *Service {
		return &v
	}).(ServicePtrOutput)
}

type ServicePtrOutput struct{ *pulumi.OutputState }

func (ServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil))
}

func (o ServicePtrOutput) ToServicePtrOutput() ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) ToServicePtrOutputWithContext(ctx context.Context) ServicePtrOutput {
	return o
}

func (o ServicePtrOutput) Elem() ServiceOutput {
	return o.ApplyT(func(v *Service) Service {
		if v != nil {
			return *v
		}
		var ret Service
		return ret
	}).(ServiceOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Service)(nil))
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Service {
		return vs[0].([]Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Service)(nil))
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Service {
		return vs[0].(map[string]Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServicePtrOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
