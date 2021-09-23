// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Logic App Integration Account Assembly.
//
// ## Import
//
// Logic App Integration Account Assemblies can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/integrationAccountAssembly:IntegrationAccountAssembly example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/assemblies/assembly1
// ```
type IntegrationAccountAssembly struct {
	pulumi.CustomResourceState

	// The name of the Logic App Integration Account Assembly.
	AssemblyName pulumi.StringOutput `pulumi:"assemblyName"`
	// The version of the Logic App Integration Account Assembly. Defaults to `0.0.0.0`.
	AssemblyVersion pulumi.StringPtrOutput `pulumi:"assemblyVersion"`
	// The content of the Logic App Integration Account Assembly.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The content link URI of the Logic App Integration Account Assembly.
	ContentLinkUri pulumi.StringPtrOutput `pulumi:"contentLinkUri"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringOutput `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Assembly.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Assembly Artifact. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Assembly Artifact should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIntegrationAccountAssembly registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountAssembly(ctx *pulumi.Context,
	name string, args *IntegrationAccountAssemblyArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountAssembly, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssemblyName == nil {
		return nil, errors.New("invalid value for required argument 'AssemblyName'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationAccountAssembly
	err := ctx.RegisterResource("azure:logicapps/integrationAccountAssembly:IntegrationAccountAssembly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountAssembly gets an existing IntegrationAccountAssembly resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountAssembly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountAssemblyState, opts ...pulumi.ResourceOption) (*IntegrationAccountAssembly, error) {
	var resource IntegrationAccountAssembly
	err := ctx.ReadResource("azure:logicapps/integrationAccountAssembly:IntegrationAccountAssembly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountAssembly resources.
type integrationAccountAssemblyState struct {
	// The name of the Logic App Integration Account Assembly.
	AssemblyName *string `pulumi:"assemblyName"`
	// The version of the Logic App Integration Account Assembly. Defaults to `0.0.0.0`.
	AssemblyVersion *string `pulumi:"assemblyVersion"`
	// The content of the Logic App Integration Account Assembly.
	Content *string `pulumi:"content"`
	// The content link URI of the Logic App Integration Account Assembly.
	ContentLinkUri *string `pulumi:"contentLinkUri"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName *string `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Assembly.
	Metadata map[string]string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Assembly Artifact. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Assembly Artifact should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IntegrationAccountAssemblyState struct {
	// The name of the Logic App Integration Account Assembly.
	AssemblyName pulumi.StringPtrInput
	// The version of the Logic App Integration Account Assembly. Defaults to `0.0.0.0`.
	AssemblyVersion pulumi.StringPtrInput
	// The content of the Logic App Integration Account Assembly.
	Content pulumi.StringPtrInput
	// The content link URI of the Logic App Integration Account Assembly.
	ContentLinkUri pulumi.StringPtrInput
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringPtrInput
	// The metadata of the Logic App Integration Account Assembly.
	Metadata pulumi.StringMapInput
	// The name which should be used for this Logic App Integration Account Assembly Artifact. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Assembly Artifact should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IntegrationAccountAssemblyState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAssemblyState)(nil)).Elem()
}

type integrationAccountAssemblyArgs struct {
	// The name of the Logic App Integration Account Assembly.
	AssemblyName string `pulumi:"assemblyName"`
	// The version of the Logic App Integration Account Assembly. Defaults to `0.0.0.0`.
	AssemblyVersion *string `pulumi:"assemblyVersion"`
	// The content of the Logic App Integration Account Assembly.
	Content *string `pulumi:"content"`
	// The content link URI of the Logic App Integration Account Assembly.
	ContentLinkUri *string `pulumi:"contentLinkUri"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Assembly.
	Metadata map[string]string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Assembly Artifact. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Assembly Artifact should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IntegrationAccountAssembly resource.
type IntegrationAccountAssemblyArgs struct {
	// The name of the Logic App Integration Account Assembly.
	AssemblyName pulumi.StringInput
	// The version of the Logic App Integration Account Assembly. Defaults to `0.0.0.0`.
	AssemblyVersion pulumi.StringPtrInput
	// The content of the Logic App Integration Account Assembly.
	Content pulumi.StringPtrInput
	// The content link URI of the Logic App Integration Account Assembly.
	ContentLinkUri pulumi.StringPtrInput
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringInput
	// The metadata of the Logic App Integration Account Assembly.
	Metadata pulumi.StringMapInput
	// The name which should be used for this Logic App Integration Account Assembly Artifact. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Assembly Artifact should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IntegrationAccountAssemblyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAssemblyArgs)(nil)).Elem()
}

type IntegrationAccountAssemblyInput interface {
	pulumi.Input

	ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput
	ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput
}

func (*IntegrationAccountAssembly) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountAssembly)(nil))
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput {
	return i.ToIntegrationAccountAssemblyOutputWithContext(context.Background())
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyOutput)
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyPtrOutput() IntegrationAccountAssemblyPtrOutput {
	return i.ToIntegrationAccountAssemblyPtrOutputWithContext(context.Background())
}

func (i *IntegrationAccountAssembly) ToIntegrationAccountAssemblyPtrOutputWithContext(ctx context.Context) IntegrationAccountAssemblyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyPtrOutput)
}

type IntegrationAccountAssemblyPtrInput interface {
	pulumi.Input

	ToIntegrationAccountAssemblyPtrOutput() IntegrationAccountAssemblyPtrOutput
	ToIntegrationAccountAssemblyPtrOutputWithContext(ctx context.Context) IntegrationAccountAssemblyPtrOutput
}

type integrationAccountAssemblyPtrType IntegrationAccountAssemblyArgs

func (*integrationAccountAssemblyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAssembly)(nil))
}

func (i *integrationAccountAssemblyPtrType) ToIntegrationAccountAssemblyPtrOutput() IntegrationAccountAssemblyPtrOutput {
	return i.ToIntegrationAccountAssemblyPtrOutputWithContext(context.Background())
}

func (i *integrationAccountAssemblyPtrType) ToIntegrationAccountAssemblyPtrOutputWithContext(ctx context.Context) IntegrationAccountAssemblyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyPtrOutput)
}

// IntegrationAccountAssemblyArrayInput is an input type that accepts IntegrationAccountAssemblyArray and IntegrationAccountAssemblyArrayOutput values.
// You can construct a concrete instance of `IntegrationAccountAssemblyArrayInput` via:
//
//          IntegrationAccountAssemblyArray{ IntegrationAccountAssemblyArgs{...} }
type IntegrationAccountAssemblyArrayInput interface {
	pulumi.Input

	ToIntegrationAccountAssemblyArrayOutput() IntegrationAccountAssemblyArrayOutput
	ToIntegrationAccountAssemblyArrayOutputWithContext(context.Context) IntegrationAccountAssemblyArrayOutput
}

type IntegrationAccountAssemblyArray []IntegrationAccountAssemblyInput

func (IntegrationAccountAssemblyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAccountAssembly)(nil)).Elem()
}

func (i IntegrationAccountAssemblyArray) ToIntegrationAccountAssemblyArrayOutput() IntegrationAccountAssemblyArrayOutput {
	return i.ToIntegrationAccountAssemblyArrayOutputWithContext(context.Background())
}

func (i IntegrationAccountAssemblyArray) ToIntegrationAccountAssemblyArrayOutputWithContext(ctx context.Context) IntegrationAccountAssemblyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyArrayOutput)
}

// IntegrationAccountAssemblyMapInput is an input type that accepts IntegrationAccountAssemblyMap and IntegrationAccountAssemblyMapOutput values.
// You can construct a concrete instance of `IntegrationAccountAssemblyMapInput` via:
//
//          IntegrationAccountAssemblyMap{ "key": IntegrationAccountAssemblyArgs{...} }
type IntegrationAccountAssemblyMapInput interface {
	pulumi.Input

	ToIntegrationAccountAssemblyMapOutput() IntegrationAccountAssemblyMapOutput
	ToIntegrationAccountAssemblyMapOutputWithContext(context.Context) IntegrationAccountAssemblyMapOutput
}

type IntegrationAccountAssemblyMap map[string]IntegrationAccountAssemblyInput

func (IntegrationAccountAssemblyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAccountAssembly)(nil)).Elem()
}

func (i IntegrationAccountAssemblyMap) ToIntegrationAccountAssemblyMapOutput() IntegrationAccountAssemblyMapOutput {
	return i.ToIntegrationAccountAssemblyMapOutputWithContext(context.Background())
}

func (i IntegrationAccountAssemblyMap) ToIntegrationAccountAssemblyMapOutputWithContext(ctx context.Context) IntegrationAccountAssemblyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAssemblyMapOutput)
}

type IntegrationAccountAssemblyOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAssemblyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountAssembly)(nil))
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyOutput() IntegrationAccountAssemblyOutput {
	return o
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyOutputWithContext(ctx context.Context) IntegrationAccountAssemblyOutput {
	return o
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyPtrOutput() IntegrationAccountAssemblyPtrOutput {
	return o.ToIntegrationAccountAssemblyPtrOutputWithContext(context.Background())
}

func (o IntegrationAccountAssemblyOutput) ToIntegrationAccountAssemblyPtrOutputWithContext(ctx context.Context) IntegrationAccountAssemblyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationAccountAssembly) *IntegrationAccountAssembly {
		return &v
	}).(IntegrationAccountAssemblyPtrOutput)
}

type IntegrationAccountAssemblyPtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAssemblyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAssembly)(nil))
}

func (o IntegrationAccountAssemblyPtrOutput) ToIntegrationAccountAssemblyPtrOutput() IntegrationAccountAssemblyPtrOutput {
	return o
}

func (o IntegrationAccountAssemblyPtrOutput) ToIntegrationAccountAssemblyPtrOutputWithContext(ctx context.Context) IntegrationAccountAssemblyPtrOutput {
	return o
}

func (o IntegrationAccountAssemblyPtrOutput) Elem() IntegrationAccountAssemblyOutput {
	return o.ApplyT(func(v *IntegrationAccountAssembly) IntegrationAccountAssembly {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountAssembly
		return ret
	}).(IntegrationAccountAssemblyOutput)
}

type IntegrationAccountAssemblyArrayOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAssemblyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationAccountAssembly)(nil))
}

func (o IntegrationAccountAssemblyArrayOutput) ToIntegrationAccountAssemblyArrayOutput() IntegrationAccountAssemblyArrayOutput {
	return o
}

func (o IntegrationAccountAssemblyArrayOutput) ToIntegrationAccountAssemblyArrayOutputWithContext(ctx context.Context) IntegrationAccountAssemblyArrayOutput {
	return o
}

func (o IntegrationAccountAssemblyArrayOutput) Index(i pulumi.IntInput) IntegrationAccountAssemblyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationAccountAssembly {
		return vs[0].([]IntegrationAccountAssembly)[vs[1].(int)]
	}).(IntegrationAccountAssemblyOutput)
}

type IntegrationAccountAssemblyMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAssemblyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IntegrationAccountAssembly)(nil))
}

func (o IntegrationAccountAssemblyMapOutput) ToIntegrationAccountAssemblyMapOutput() IntegrationAccountAssemblyMapOutput {
	return o
}

func (o IntegrationAccountAssemblyMapOutput) ToIntegrationAccountAssemblyMapOutputWithContext(ctx context.Context) IntegrationAccountAssemblyMapOutput {
	return o
}

func (o IntegrationAccountAssemblyMapOutput) MapIndex(k pulumi.StringInput) IntegrationAccountAssemblyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IntegrationAccountAssembly {
		return vs[0].(map[string]IntegrationAccountAssembly)[vs[1].(string)]
	}).(IntegrationAccountAssemblyOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountAssemblyOutput{})
	pulumi.RegisterOutputType(IntegrationAccountAssemblyPtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountAssemblyArrayOutput{})
	pulumi.RegisterOutputType(IntegrationAccountAssemblyMapOutput{})
}
