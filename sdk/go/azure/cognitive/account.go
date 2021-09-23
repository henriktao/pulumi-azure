// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognitive

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Cognitive Service Accounts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cognitive/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
// ```
type Account struct {
	pulumi.CustomResourceState

	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName pulumi.StringPtrOutput `pulumi:"customSubdomainName"`
	// The endpoint used to connect to the Cognitive Service Account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// List of FQDNs allowed for the Cognitive Account.
	Fqdns pulumi.StringArrayOutput `pulumi:"fqdns"`
	// An `identity` block is documented below.
	Identity AccountIdentityPtrOutput `pulumi:"identity"`
	// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrOutput `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadClientId pulumi.StringPtrOutput `pulumi:"metricsAdvisorAadClientId"`
	// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadTenantId pulumi.StringPtrOutput `pulumi:"metricsAdvisorAadTenantId"`
	// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName pulumi.StringPtrOutput `pulumi:"metricsAdvisorSuperUserName"`
	// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName pulumi.StringPtrOutput `pulumi:"metricsAdvisorWebsiteName"`
	// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls AccountNetworkAclsPtrOutput `pulumi:"networkAcls"`
	// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
	OutboundNetworkAccessRestrited pulumi.BoolPtrOutput `pulumi:"outboundNetworkAccessRestrited"`
	// A primary access key which can be used to connect to the Cognitive Service Account.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint pulumi.StringPtrOutput `pulumi:"qnaRuntimeEndpoint"`
	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary access key which can be used to connect to the Cognitive Service Account.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// An `identity` block is documented below.
	Storages AccountStorageArrayOutput `pulumi:"storages"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	var resource Account
	err := ctx.RegisterResource("azure:cognitive/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:cognitive/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName *string `pulumi:"customSubdomainName"`
	// The endpoint used to connect to the Cognitive Service Account.
	Endpoint *string `pulumi:"endpoint"`
	// List of FQDNs allowed for the Cognitive Account.
	Fqdns []string `pulumi:"fqdns"`
	// An `identity` block is documented below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadClientId *string `pulumi:"metricsAdvisorAadClientId"`
	// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadTenantId *string `pulumi:"metricsAdvisorAadTenantId"`
	// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName *string `pulumi:"metricsAdvisorSuperUserName"`
	// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName *string `pulumi:"metricsAdvisorWebsiteName"`
	// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *AccountNetworkAcls `pulumi:"networkAcls"`
	// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
	OutboundNetworkAccessRestrited *bool `pulumi:"outboundNetworkAccessRestrited"`
	// A primary access key which can be used to connect to the Cognitive Service Account.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint *string `pulumi:"qnaRuntimeEndpoint"`
	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary access key which can be used to connect to the Cognitive Service Account.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
	SkuName *string `pulumi:"skuName"`
	// An `identity` block is documented below.
	Storages []AccountStorage `pulumi:"storages"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AccountState struct {
	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName pulumi.StringPtrInput
	// The endpoint used to connect to the Cognitive Service Account.
	Endpoint pulumi.StringPtrInput
	// List of FQDNs allowed for the Cognitive Account.
	Fqdns pulumi.StringArrayInput
	// An `identity` block is documented below.
	Identity AccountIdentityPtrInput
	// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadClientId pulumi.StringPtrInput
	// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadTenantId pulumi.StringPtrInput
	// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName pulumi.StringPtrInput
	// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName pulumi.StringPtrInput
	// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls AccountNetworkAclsPtrInput
	// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
	OutboundNetworkAccessRestrited pulumi.BoolPtrInput
	// A primary access key which can be used to connect to the Cognitive Service Account.
	PrimaryAccessKey pulumi.StringPtrInput
	// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint pulumi.StringPtrInput
	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary access key which can be used to connect to the Cognitive Service Account.
	SecondaryAccessKey pulumi.StringPtrInput
	// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
	SkuName pulumi.StringPtrInput
	// An `identity` block is documented below.
	Storages AccountStorageArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName *string `pulumi:"customSubdomainName"`
	// List of FQDNs allowed for the Cognitive Account.
	Fqdns []string `pulumi:"fqdns"`
	// An `identity` block is documented below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
	Kind string `pulumi:"kind"`
	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadClientId *string `pulumi:"metricsAdvisorAadClientId"`
	// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadTenantId *string `pulumi:"metricsAdvisorAadTenantId"`
	// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName *string `pulumi:"metricsAdvisorSuperUserName"`
	// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName *string `pulumi:"metricsAdvisorWebsiteName"`
	// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkAcls` block as defined below.
	NetworkAcls *AccountNetworkAcls `pulumi:"networkAcls"`
	// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
	OutboundNetworkAccessRestrited *bool `pulumi:"outboundNetworkAccessRestrited"`
	// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint *string `pulumi:"qnaRuntimeEndpoint"`
	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
	SkuName string `pulumi:"skuName"`
	// An `identity` block is documented below.
	Storages []AccountStorage `pulumi:"storages"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName pulumi.StringPtrInput
	// List of FQDNs allowed for the Cognitive Account.
	Fqdns pulumi.StringArrayInput
	// An `identity` block is documented below.
	Identity AccountIdentityPtrInput
	// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
	Kind pulumi.StringInput
	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
	LocalAuthEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadClientId pulumi.StringPtrInput
	// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorAadTenantId pulumi.StringPtrInput
	// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName pulumi.StringPtrInput
	// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName pulumi.StringPtrInput
	// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkAcls` block as defined below.
	NetworkAcls AccountNetworkAclsPtrInput
	// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
	OutboundNetworkAccessRestrited pulumi.BoolPtrInput
	// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint pulumi.StringPtrInput
	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
	SkuName pulumi.StringInput
	// An `identity` block is documented below.
	Storages AccountStorageArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

func (i *Account) ToAccountPtrOutput() AccountPtrOutput {
	return i.ToAccountPtrOutputWithContext(context.Background())
}

func (i *Account) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPtrOutput)
}

type AccountPtrInput interface {
	pulumi.Input

	ToAccountPtrOutput() AccountPtrOutput
	ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput
}

type accountPtrType AccountArgs

func (*accountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil))
}

func (i *accountPtrType) ToAccountPtrOutput() AccountPtrOutput {
	return i.ToAccountPtrOutputWithContext(context.Background())
}

func (i *accountPtrType) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPtrOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//          AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//          AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) ToAccountPtrOutput() AccountPtrOutput {
	return o.ToAccountPtrOutputWithContext(context.Background())
}

func (o AccountOutput) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Account) *Account {
		return &v
	}).(AccountPtrOutput)
}

type AccountPtrOutput struct{ *pulumi.OutputState }

func (AccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil))
}

func (o AccountPtrOutput) ToAccountPtrOutput() AccountPtrOutput {
	return o
}

func (o AccountPtrOutput) ToAccountPtrOutputWithContext(ctx context.Context) AccountPtrOutput {
	return o
}

func (o AccountPtrOutput) Elem() AccountOutput {
	return o.ApplyT(func(v *Account) Account {
		if v != nil {
			return *v
		}
		var ret Account
		return ret
	}).(AccountOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Account)(nil))
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Account {
		return vs[0].([]Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Account)(nil))
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Account {
		return vs[0].(map[string]Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountPtrOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
