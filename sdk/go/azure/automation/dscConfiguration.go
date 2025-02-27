// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Automation DSC Configuration.
//
// ## Import
//
// Automation DSC Configuration's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/dscConfiguration:DscConfiguration configuration1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/configurations/configuration1
// ```
type DscConfiguration struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringOutput `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location pulumi.StringOutput `pulumi:"location"`
	// Verbose log option.
	LogVerbose pulumi.BoolPtrOutput `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	State             pulumi.StringOutput `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDscConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscConfiguration(ctx *pulumi.Context,
	name string, args *DscConfigurationArgs, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentEmbedded == nil {
		return nil, errors.New("invalid value for required argument 'ContentEmbedded'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DscConfiguration
	err := ctx.RegisterResource("azure:automation/dscConfiguration:DscConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscConfiguration gets an existing DscConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscConfigurationState, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	var resource DscConfiguration
	err := ctx.ReadResource("azure:automation/dscConfiguration:DscConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscConfiguration resources.
type dscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded *string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	State             *string `pulumi:"state"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringPtrInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	State             pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DscConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationState)(nil)).Elem()
}

type dscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DscConfiguration resource.
type DscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DscConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationArgs)(nil)).Elem()
}

type DscConfigurationInput interface {
	pulumi.Input

	ToDscConfigurationOutput() DscConfigurationOutput
	ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput
}

func (*DscConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfiguration)(nil)).Elem()
}

func (i *DscConfiguration) ToDscConfigurationOutput() DscConfigurationOutput {
	return i.ToDscConfigurationOutputWithContext(context.Background())
}

func (i *DscConfiguration) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationOutput)
}

// DscConfigurationArrayInput is an input type that accepts DscConfigurationArray and DscConfigurationArrayOutput values.
// You can construct a concrete instance of `DscConfigurationArrayInput` via:
//
//          DscConfigurationArray{ DscConfigurationArgs{...} }
type DscConfigurationArrayInput interface {
	pulumi.Input

	ToDscConfigurationArrayOutput() DscConfigurationArrayOutput
	ToDscConfigurationArrayOutputWithContext(context.Context) DscConfigurationArrayOutput
}

type DscConfigurationArray []DscConfigurationInput

func (DscConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DscConfiguration)(nil)).Elem()
}

func (i DscConfigurationArray) ToDscConfigurationArrayOutput() DscConfigurationArrayOutput {
	return i.ToDscConfigurationArrayOutputWithContext(context.Background())
}

func (i DscConfigurationArray) ToDscConfigurationArrayOutputWithContext(ctx context.Context) DscConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationArrayOutput)
}

// DscConfigurationMapInput is an input type that accepts DscConfigurationMap and DscConfigurationMapOutput values.
// You can construct a concrete instance of `DscConfigurationMapInput` via:
//
//          DscConfigurationMap{ "key": DscConfigurationArgs{...} }
type DscConfigurationMapInput interface {
	pulumi.Input

	ToDscConfigurationMapOutput() DscConfigurationMapOutput
	ToDscConfigurationMapOutputWithContext(context.Context) DscConfigurationMapOutput
}

type DscConfigurationMap map[string]DscConfigurationInput

func (DscConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DscConfiguration)(nil)).Elem()
}

func (i DscConfigurationMap) ToDscConfigurationMapOutput() DscConfigurationMapOutput {
	return i.ToDscConfigurationMapOutputWithContext(context.Background())
}

func (i DscConfigurationMap) ToDscConfigurationMapOutputWithContext(ctx context.Context) DscConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationMapOutput)
}

type DscConfigurationOutput struct{ *pulumi.OutputState }

func (DscConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfiguration)(nil)).Elem()
}

func (o DscConfigurationOutput) ToDscConfigurationOutput() DscConfigurationOutput {
	return o
}

func (o DscConfigurationOutput) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return o
}

type DscConfigurationArrayOutput struct{ *pulumi.OutputState }

func (DscConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DscConfiguration)(nil)).Elem()
}

func (o DscConfigurationArrayOutput) ToDscConfigurationArrayOutput() DscConfigurationArrayOutput {
	return o
}

func (o DscConfigurationArrayOutput) ToDscConfigurationArrayOutputWithContext(ctx context.Context) DscConfigurationArrayOutput {
	return o
}

func (o DscConfigurationArrayOutput) Index(i pulumi.IntInput) DscConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DscConfiguration {
		return vs[0].([]*DscConfiguration)[vs[1].(int)]
	}).(DscConfigurationOutput)
}

type DscConfigurationMapOutput struct{ *pulumi.OutputState }

func (DscConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DscConfiguration)(nil)).Elem()
}

func (o DscConfigurationMapOutput) ToDscConfigurationMapOutput() DscConfigurationMapOutput {
	return o
}

func (o DscConfigurationMapOutput) ToDscConfigurationMapOutputWithContext(ctx context.Context) DscConfigurationMapOutput {
	return o
}

func (o DscConfigurationMapOutput) MapIndex(k pulumi.StringInput) DscConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DscConfiguration {
		return vs[0].(map[string]*DscConfiguration)[vs[1].(string)]
	}).(DscConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationInput)(nil)).Elem(), &DscConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationArrayInput)(nil)).Elem(), DscConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationMapInput)(nil)).Elem(), DscConfigurationMap{})
	pulumi.RegisterOutputType(DscConfigurationOutput{})
	pulumi.RegisterOutputType(DscConfigurationArrayOutput{})
	pulumi.RegisterOutputType(DscConfigurationMapOutput{})
}
