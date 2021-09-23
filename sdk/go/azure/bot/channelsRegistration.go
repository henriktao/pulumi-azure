// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Bot Channels Registration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/bot"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bot.NewChannelsRegistration(ctx, "exampleChannelsRegistration", &bot.ChannelsRegistrationArgs{
// 			Location:          pulumi.String("global"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("F0"),
// 			MicrosoftAppId:    pulumi.String(current.ClientId),
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
// Bot Channels Registration can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:bot/channelsRegistration:ChannelsRegistration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
// ```
type ChannelsRegistration struct {
	pulumi.CustomResourceState

	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultUrl pulumi.StringPtrOutput `pulumi:"cmkKeyVaultUrl"`
	// The description of the Bot Channels Registration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringOutput `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringOutput `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringOutput `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The icon URL to visually identify the Bot Channels Registration.
	IconUrl pulumi.StringOutput `pulumi:"iconUrl"`
	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled pulumi.BoolPtrOutput `pulumi:"isolatedNetworkEnabled"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringOutput `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewChannelsRegistration registers a new resource with the given unique name, arguments, and options.
func NewChannelsRegistration(ctx *pulumi.Context,
	name string, args *ChannelsRegistrationArgs, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MicrosoftAppId == nil {
		return nil, errors.New("invalid value for required argument 'MicrosoftAppId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource ChannelsRegistration
	err := ctx.RegisterResource("azure:bot/channelsRegistration:ChannelsRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelsRegistration gets an existing ChannelsRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelsRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelsRegistrationState, opts ...pulumi.ResourceOption) (*ChannelsRegistration, error) {
	var resource ChannelsRegistration
	err := ctx.ReadResource("azure:bot/channelsRegistration:ChannelsRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelsRegistration resources.
type channelsRegistrationState struct {
	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultUrl *string `pulumi:"cmkKeyVaultUrl"`
	// The description of the Bot Channels Registration.
	Description *string `pulumi:"description"`
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The icon URL to visually identify the Bot Channels Registration.
	IconUrl *string `pulumi:"iconUrl"`
	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled *bool `pulumi:"isolatedNetworkEnabled"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId *string `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ChannelsRegistrationState struct {
	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultUrl pulumi.StringPtrInput
	// The description of the Bot Channels Registration.
	Description pulumi.StringPtrInput
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrInput
	// The icon URL to visually identify the Bot Channels Registration.
	IconUrl pulumi.StringPtrInput
	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled pulumi.BoolPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringPtrInput
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelsRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelsRegistrationState)(nil)).Elem()
}

type channelsRegistrationArgs struct {
	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultUrl *string `pulumi:"cmkKeyVaultUrl"`
	// The description of the Bot Channels Registration.
	Description *string `pulumi:"description"`
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey *string `pulumi:"developerAppInsightsApiKey"`
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId *string `pulumi:"developerAppInsightsApplicationId"`
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey *string `pulumi:"developerAppInsightsKey"`
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName *string `pulumi:"displayName"`
	// The Bot Channels Registration endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// The icon URL to visually identify the Bot Channels Registration.
	IconUrl *string `pulumi:"iconUrl"`
	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled *bool `pulumi:"isolatedNetworkEnabled"`
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId string `pulumi:"microsoftAppId"`
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ChannelsRegistration resource.
type ChannelsRegistrationArgs struct {
	// The CMK Key Vault Key URL to encrypt the Bot Channels Registration with the Customer Managed Encryption Key.
	CmkKeyVaultUrl pulumi.StringPtrInput
	// The description of the Bot Channels Registration.
	Description pulumi.StringPtrInput
	// The Application Insights API Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsApiKey pulumi.StringPtrInput
	// The Application Insights Application ID to associate with the Bot Channels Registration.
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput
	// The Application Insights Key to associate with the Bot Channels Registration.
	DeveloperAppInsightsKey pulumi.StringPtrInput
	// The name of the Bot Channels Registration will be displayed as. This defaults to `name` if not specified.
	DisplayName pulumi.StringPtrInput
	// The Bot Channels Registration endpoint.
	Endpoint pulumi.StringPtrInput
	// The icon URL to visually identify the Bot Channels Registration.
	IconUrl pulumi.StringPtrInput
	// Is the Bot Channels Registration in an isolated network?
	IsolatedNetworkEnabled pulumi.BoolPtrInput
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Microsoft Application ID for the Bot Channels Registration. Changing this forces a new resource to be created.
	MicrosoftAppId pulumi.StringInput
	// Specifies the name of the Bot Channels Registration. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bot Channels Registration. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The SKU of the Bot Channels Registration. Valid values include `F0` or `S1`. Changing this forces a new resource to be created.
	Sku pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelsRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelsRegistrationArgs)(nil)).Elem()
}

type ChannelsRegistrationInput interface {
	pulumi.Input

	ToChannelsRegistrationOutput() ChannelsRegistrationOutput
	ToChannelsRegistrationOutputWithContext(ctx context.Context) ChannelsRegistrationOutput
}

func (*ChannelsRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelsRegistration)(nil))
}

func (i *ChannelsRegistration) ToChannelsRegistrationOutput() ChannelsRegistrationOutput {
	return i.ToChannelsRegistrationOutputWithContext(context.Background())
}

func (i *ChannelsRegistration) ToChannelsRegistrationOutputWithContext(ctx context.Context) ChannelsRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelsRegistrationOutput)
}

func (i *ChannelsRegistration) ToChannelsRegistrationPtrOutput() ChannelsRegistrationPtrOutput {
	return i.ToChannelsRegistrationPtrOutputWithContext(context.Background())
}

func (i *ChannelsRegistration) ToChannelsRegistrationPtrOutputWithContext(ctx context.Context) ChannelsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelsRegistrationPtrOutput)
}

type ChannelsRegistrationPtrInput interface {
	pulumi.Input

	ToChannelsRegistrationPtrOutput() ChannelsRegistrationPtrOutput
	ToChannelsRegistrationPtrOutputWithContext(ctx context.Context) ChannelsRegistrationPtrOutput
}

type channelsRegistrationPtrType ChannelsRegistrationArgs

func (*channelsRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelsRegistration)(nil))
}

func (i *channelsRegistrationPtrType) ToChannelsRegistrationPtrOutput() ChannelsRegistrationPtrOutput {
	return i.ToChannelsRegistrationPtrOutputWithContext(context.Background())
}

func (i *channelsRegistrationPtrType) ToChannelsRegistrationPtrOutputWithContext(ctx context.Context) ChannelsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelsRegistrationPtrOutput)
}

// ChannelsRegistrationArrayInput is an input type that accepts ChannelsRegistrationArray and ChannelsRegistrationArrayOutput values.
// You can construct a concrete instance of `ChannelsRegistrationArrayInput` via:
//
//          ChannelsRegistrationArray{ ChannelsRegistrationArgs{...} }
type ChannelsRegistrationArrayInput interface {
	pulumi.Input

	ToChannelsRegistrationArrayOutput() ChannelsRegistrationArrayOutput
	ToChannelsRegistrationArrayOutputWithContext(context.Context) ChannelsRegistrationArrayOutput
}

type ChannelsRegistrationArray []ChannelsRegistrationInput

func (ChannelsRegistrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChannelsRegistration)(nil)).Elem()
}

func (i ChannelsRegistrationArray) ToChannelsRegistrationArrayOutput() ChannelsRegistrationArrayOutput {
	return i.ToChannelsRegistrationArrayOutputWithContext(context.Background())
}

func (i ChannelsRegistrationArray) ToChannelsRegistrationArrayOutputWithContext(ctx context.Context) ChannelsRegistrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelsRegistrationArrayOutput)
}

// ChannelsRegistrationMapInput is an input type that accepts ChannelsRegistrationMap and ChannelsRegistrationMapOutput values.
// You can construct a concrete instance of `ChannelsRegistrationMapInput` via:
//
//          ChannelsRegistrationMap{ "key": ChannelsRegistrationArgs{...} }
type ChannelsRegistrationMapInput interface {
	pulumi.Input

	ToChannelsRegistrationMapOutput() ChannelsRegistrationMapOutput
	ToChannelsRegistrationMapOutputWithContext(context.Context) ChannelsRegistrationMapOutput
}

type ChannelsRegistrationMap map[string]ChannelsRegistrationInput

func (ChannelsRegistrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChannelsRegistration)(nil)).Elem()
}

func (i ChannelsRegistrationMap) ToChannelsRegistrationMapOutput() ChannelsRegistrationMapOutput {
	return i.ToChannelsRegistrationMapOutputWithContext(context.Background())
}

func (i ChannelsRegistrationMap) ToChannelsRegistrationMapOutputWithContext(ctx context.Context) ChannelsRegistrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelsRegistrationMapOutput)
}

type ChannelsRegistrationOutput struct{ *pulumi.OutputState }

func (ChannelsRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelsRegistration)(nil))
}

func (o ChannelsRegistrationOutput) ToChannelsRegistrationOutput() ChannelsRegistrationOutput {
	return o
}

func (o ChannelsRegistrationOutput) ToChannelsRegistrationOutputWithContext(ctx context.Context) ChannelsRegistrationOutput {
	return o
}

func (o ChannelsRegistrationOutput) ToChannelsRegistrationPtrOutput() ChannelsRegistrationPtrOutput {
	return o.ToChannelsRegistrationPtrOutputWithContext(context.Background())
}

func (o ChannelsRegistrationOutput) ToChannelsRegistrationPtrOutputWithContext(ctx context.Context) ChannelsRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChannelsRegistration) *ChannelsRegistration {
		return &v
	}).(ChannelsRegistrationPtrOutput)
}

type ChannelsRegistrationPtrOutput struct{ *pulumi.OutputState }

func (ChannelsRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelsRegistration)(nil))
}

func (o ChannelsRegistrationPtrOutput) ToChannelsRegistrationPtrOutput() ChannelsRegistrationPtrOutput {
	return o
}

func (o ChannelsRegistrationPtrOutput) ToChannelsRegistrationPtrOutputWithContext(ctx context.Context) ChannelsRegistrationPtrOutput {
	return o
}

func (o ChannelsRegistrationPtrOutput) Elem() ChannelsRegistrationOutput {
	return o.ApplyT(func(v *ChannelsRegistration) ChannelsRegistration {
		if v != nil {
			return *v
		}
		var ret ChannelsRegistration
		return ret
	}).(ChannelsRegistrationOutput)
}

type ChannelsRegistrationArrayOutput struct{ *pulumi.OutputState }

func (ChannelsRegistrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelsRegistration)(nil))
}

func (o ChannelsRegistrationArrayOutput) ToChannelsRegistrationArrayOutput() ChannelsRegistrationArrayOutput {
	return o
}

func (o ChannelsRegistrationArrayOutput) ToChannelsRegistrationArrayOutputWithContext(ctx context.Context) ChannelsRegistrationArrayOutput {
	return o
}

func (o ChannelsRegistrationArrayOutput) Index(i pulumi.IntInput) ChannelsRegistrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelsRegistration {
		return vs[0].([]ChannelsRegistration)[vs[1].(int)]
	}).(ChannelsRegistrationOutput)
}

type ChannelsRegistrationMapOutput struct{ *pulumi.OutputState }

func (ChannelsRegistrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ChannelsRegistration)(nil))
}

func (o ChannelsRegistrationMapOutput) ToChannelsRegistrationMapOutput() ChannelsRegistrationMapOutput {
	return o
}

func (o ChannelsRegistrationMapOutput) ToChannelsRegistrationMapOutputWithContext(ctx context.Context) ChannelsRegistrationMapOutput {
	return o
}

func (o ChannelsRegistrationMapOutput) MapIndex(k pulumi.StringInput) ChannelsRegistrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ChannelsRegistration {
		return vs[0].(map[string]ChannelsRegistration)[vs[1].(string)]
	}).(ChannelsRegistrationOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelsRegistrationOutput{})
	pulumi.RegisterOutputType(ChannelsRegistrationPtrOutput{})
	pulumi.RegisterOutputType(ChannelsRegistrationArrayOutput{})
	pulumi.RegisterOutputType(ChannelsRegistrationMapOutput{})
}
