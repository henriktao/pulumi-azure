// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Automation Connection with type `Azure`.
//
// ## Import
//
// Automation Connection can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/connectionCertificate:ConnectionCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
// ```
type ConnectionCertificate struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The name of the automation certificate.
	AutomationCertificateName pulumi.StringOutput `pulumi:"automationCertificateName"`
	// A description for this Connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The id of subscription where the automation certificate exists.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
}

// NewConnectionCertificate registers a new resource with the given unique name, arguments, and options.
func NewConnectionCertificate(ctx *pulumi.Context,
	name string, args *ConnectionCertificateArgs, opts ...pulumi.ResourceOption) (*ConnectionCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.AutomationCertificateName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationCertificateName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	var resource ConnectionCertificate
	err := ctx.RegisterResource("azure:automation/connectionCertificate:ConnectionCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionCertificate gets an existing ConnectionCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionCertificateState, opts ...pulumi.ResourceOption) (*ConnectionCertificate, error) {
	var resource ConnectionCertificate
	err := ctx.ReadResource("azure:automation/connectionCertificate:ConnectionCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionCertificate resources.
type connectionCertificateState struct {
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The name of the automation certificate.
	AutomationCertificateName *string `pulumi:"automationCertificateName"`
	// A description for this Connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The id of subscription where the automation certificate exists.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

type ConnectionCertificateState struct {
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The name of the automation certificate.
	AutomationCertificateName pulumi.StringPtrInput
	// A description for this Connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The id of subscription where the automation certificate exists.
	SubscriptionId pulumi.StringPtrInput
}

func (ConnectionCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionCertificateState)(nil)).Elem()
}

type connectionCertificateArgs struct {
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of the automation certificate.
	AutomationCertificateName string `pulumi:"automationCertificateName"`
	// A description for this Connection.
	Description *string `pulumi:"description"`
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The id of subscription where the automation certificate exists.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a ConnectionCertificate resource.
type ConnectionCertificateArgs struct {
	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The name of the automation certificate.
	AutomationCertificateName pulumi.StringInput
	// A description for this Connection.
	Description pulumi.StringPtrInput
	// Specifies the name of the Connection. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The id of subscription where the automation certificate exists.
	SubscriptionId pulumi.StringInput
}

func (ConnectionCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionCertificateArgs)(nil)).Elem()
}

type ConnectionCertificateInput interface {
	pulumi.Input

	ToConnectionCertificateOutput() ConnectionCertificateOutput
	ToConnectionCertificateOutputWithContext(ctx context.Context) ConnectionCertificateOutput
}

func (*ConnectionCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionCertificate)(nil))
}

func (i *ConnectionCertificate) ToConnectionCertificateOutput() ConnectionCertificateOutput {
	return i.ToConnectionCertificateOutputWithContext(context.Background())
}

func (i *ConnectionCertificate) ToConnectionCertificateOutputWithContext(ctx context.Context) ConnectionCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionCertificateOutput)
}

func (i *ConnectionCertificate) ToConnectionCertificatePtrOutput() ConnectionCertificatePtrOutput {
	return i.ToConnectionCertificatePtrOutputWithContext(context.Background())
}

func (i *ConnectionCertificate) ToConnectionCertificatePtrOutputWithContext(ctx context.Context) ConnectionCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionCertificatePtrOutput)
}

type ConnectionCertificatePtrInput interface {
	pulumi.Input

	ToConnectionCertificatePtrOutput() ConnectionCertificatePtrOutput
	ToConnectionCertificatePtrOutputWithContext(ctx context.Context) ConnectionCertificatePtrOutput
}

type connectionCertificatePtrType ConnectionCertificateArgs

func (*connectionCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionCertificate)(nil))
}

func (i *connectionCertificatePtrType) ToConnectionCertificatePtrOutput() ConnectionCertificatePtrOutput {
	return i.ToConnectionCertificatePtrOutputWithContext(context.Background())
}

func (i *connectionCertificatePtrType) ToConnectionCertificatePtrOutputWithContext(ctx context.Context) ConnectionCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionCertificatePtrOutput)
}

// ConnectionCertificateArrayInput is an input type that accepts ConnectionCertificateArray and ConnectionCertificateArrayOutput values.
// You can construct a concrete instance of `ConnectionCertificateArrayInput` via:
//
//          ConnectionCertificateArray{ ConnectionCertificateArgs{...} }
type ConnectionCertificateArrayInput interface {
	pulumi.Input

	ToConnectionCertificateArrayOutput() ConnectionCertificateArrayOutput
	ToConnectionCertificateArrayOutputWithContext(context.Context) ConnectionCertificateArrayOutput
}

type ConnectionCertificateArray []ConnectionCertificateInput

func (ConnectionCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectionCertificate)(nil)).Elem()
}

func (i ConnectionCertificateArray) ToConnectionCertificateArrayOutput() ConnectionCertificateArrayOutput {
	return i.ToConnectionCertificateArrayOutputWithContext(context.Background())
}

func (i ConnectionCertificateArray) ToConnectionCertificateArrayOutputWithContext(ctx context.Context) ConnectionCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionCertificateArrayOutput)
}

// ConnectionCertificateMapInput is an input type that accepts ConnectionCertificateMap and ConnectionCertificateMapOutput values.
// You can construct a concrete instance of `ConnectionCertificateMapInput` via:
//
//          ConnectionCertificateMap{ "key": ConnectionCertificateArgs{...} }
type ConnectionCertificateMapInput interface {
	pulumi.Input

	ToConnectionCertificateMapOutput() ConnectionCertificateMapOutput
	ToConnectionCertificateMapOutputWithContext(context.Context) ConnectionCertificateMapOutput
}

type ConnectionCertificateMap map[string]ConnectionCertificateInput

func (ConnectionCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectionCertificate)(nil)).Elem()
}

func (i ConnectionCertificateMap) ToConnectionCertificateMapOutput() ConnectionCertificateMapOutput {
	return i.ToConnectionCertificateMapOutputWithContext(context.Background())
}

func (i ConnectionCertificateMap) ToConnectionCertificateMapOutputWithContext(ctx context.Context) ConnectionCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionCertificateMapOutput)
}

type ConnectionCertificateOutput struct{ *pulumi.OutputState }

func (ConnectionCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionCertificate)(nil))
}

func (o ConnectionCertificateOutput) ToConnectionCertificateOutput() ConnectionCertificateOutput {
	return o
}

func (o ConnectionCertificateOutput) ToConnectionCertificateOutputWithContext(ctx context.Context) ConnectionCertificateOutput {
	return o
}

func (o ConnectionCertificateOutput) ToConnectionCertificatePtrOutput() ConnectionCertificatePtrOutput {
	return o.ToConnectionCertificatePtrOutputWithContext(context.Background())
}

func (o ConnectionCertificateOutput) ToConnectionCertificatePtrOutputWithContext(ctx context.Context) ConnectionCertificatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionCertificate) *ConnectionCertificate {
		return &v
	}).(ConnectionCertificatePtrOutput)
}

type ConnectionCertificatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionCertificate)(nil))
}

func (o ConnectionCertificatePtrOutput) ToConnectionCertificatePtrOutput() ConnectionCertificatePtrOutput {
	return o
}

func (o ConnectionCertificatePtrOutput) ToConnectionCertificatePtrOutputWithContext(ctx context.Context) ConnectionCertificatePtrOutput {
	return o
}

func (o ConnectionCertificatePtrOutput) Elem() ConnectionCertificateOutput {
	return o.ApplyT(func(v *ConnectionCertificate) ConnectionCertificate {
		if v != nil {
			return *v
		}
		var ret ConnectionCertificate
		return ret
	}).(ConnectionCertificateOutput)
}

type ConnectionCertificateArrayOutput struct{ *pulumi.OutputState }

func (ConnectionCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionCertificate)(nil))
}

func (o ConnectionCertificateArrayOutput) ToConnectionCertificateArrayOutput() ConnectionCertificateArrayOutput {
	return o
}

func (o ConnectionCertificateArrayOutput) ToConnectionCertificateArrayOutputWithContext(ctx context.Context) ConnectionCertificateArrayOutput {
	return o
}

func (o ConnectionCertificateArrayOutput) Index(i pulumi.IntInput) ConnectionCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionCertificate {
		return vs[0].([]ConnectionCertificate)[vs[1].(int)]
	}).(ConnectionCertificateOutput)
}

type ConnectionCertificateMapOutput struct{ *pulumi.OutputState }

func (ConnectionCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnectionCertificate)(nil))
}

func (o ConnectionCertificateMapOutput) ToConnectionCertificateMapOutput() ConnectionCertificateMapOutput {
	return o
}

func (o ConnectionCertificateMapOutput) ToConnectionCertificateMapOutputWithContext(ctx context.Context) ConnectionCertificateMapOutput {
	return o
}

func (o ConnectionCertificateMapOutput) MapIndex(k pulumi.StringInput) ConnectionCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnectionCertificate {
		return vs[0].(map[string]ConnectionCertificate)[vs[1].(string)]
	}).(ConnectionCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionCertificateInput)(nil)).Elem(), &ConnectionCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionCertificatePtrInput)(nil)).Elem(), &ConnectionCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionCertificateArrayInput)(nil)).Elem(), ConnectionCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionCertificateMapInput)(nil)).Elem(), ConnectionCertificateMap{})
	pulumi.RegisterOutputType(ConnectionCertificateOutput{})
	pulumi.RegisterOutputType(ConnectionCertificatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionCertificateArrayOutput{})
	pulumi.RegisterOutputType(ConnectionCertificateMapOutput{})
}
