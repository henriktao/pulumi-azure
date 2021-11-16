// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cognitive
{
    /// <summary>
    /// ## Import
    /// 
    /// Cognitive Service Accounts can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:cognitive/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
    /// ```
    /// </summary>
    [AzureResourceType("azure:cognitive/account:Account")]
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        /// </summary>
        [Output("customSubdomainName")]
        public Output<string?> CustomSubdomainName { get; private set; } = null!;

        /// <summary>
        /// The endpoint used to connect to the Cognitive Service Account.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// List of FQDNs allowed for the Cognitive Account.
        /// </summary>
        [Output("fqdns")]
        public Output<ImmutableArray<string>> Fqdns { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.AccountIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Output("localAuthEnabled")]
        public Output<bool?> LocalAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("metricsAdvisorAadClientId")]
        public Output<string?> MetricsAdvisorAadClientId { get; private set; } = null!;

        /// <summary>
        /// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("metricsAdvisorAadTenantId")]
        public Output<string?> MetricsAdvisorAadTenantId { get; private set; } = null!;

        /// <summary>
        /// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("metricsAdvisorSuperUserName")]
        public Output<string?> MetricsAdvisorSuperUserName { get; private set; } = null!;

        /// <summary>
        /// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("metricsAdvisorWebsiteName")]
        public Output<string?> MetricsAdvisorWebsiteName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `network_acls` block as defined below.
        /// </summary>
        [Output("networkAcls")]
        public Output<Outputs.AccountNetworkAcls?> NetworkAcls { get; private set; } = null!;

        /// <summary>
        /// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
        /// </summary>
        [Output("outboundNetworkAccessRestrited")]
        public Output<bool?> OutboundNetworkAccessRestrited { get; private set; } = null!;

        /// <summary>
        /// A primary access key which can be used to connect to the Cognitive Service Account.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// A URL to link a QnAMaker cognitive account to a QnA runtime.
        /// </summary>
        [Output("qnaRuntimeEndpoint")]
        public Output<string?> QnaRuntimeEndpoint { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The secondary access key which can be used to connect to the Cognitive Service Account.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        [Output("storages")]
        public Output<ImmutableArray<Outputs.AccountStorage>> Storages { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azure:cognitive/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:cognitive/account:Account", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        /// </summary>
        [Input("customSubdomainName")]
        public Input<string>? CustomSubdomainName { get; set; }

        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// List of FQDNs allowed for the Cognitive Account.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AccountIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Input("localAuthEnabled")]
        public Input<bool>? LocalAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorAadClientId")]
        public Input<string>? MetricsAdvisorAadClientId { get; set; }

        /// <summary>
        /// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorAadTenantId")]
        public Input<string>? MetricsAdvisorAadTenantId { get; set; }

        /// <summary>
        /// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorSuperUserName")]
        public Input<string>? MetricsAdvisorSuperUserName { get; set; }

        /// <summary>
        /// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorWebsiteName")]
        public Input<string>? MetricsAdvisorWebsiteName { get; set; }

        /// <summary>
        /// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_acls` block as defined below.
        /// </summary>
        [Input("networkAcls")]
        public Input<Inputs.AccountNetworkAclsArgs>? NetworkAcls { get; set; }

        /// <summary>
        /// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
        /// </summary>
        [Input("outboundNetworkAccessRestrited")]
        public Input<bool>? OutboundNetworkAccessRestrited { get; set; }

        /// <summary>
        /// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// A URL to link a QnAMaker cognitive account to a QnA runtime.
        /// </summary>
        [Input("qnaRuntimeEndpoint")]
        public Input<string>? QnaRuntimeEndpoint { get; set; }

        /// <summary>
        /// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("storages")]
        private InputList<Inputs.AccountStorageArgs>? _storages;

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        public InputList<Inputs.AccountStorageArgs> Storages
        {
            get => _storages ?? (_storages = new InputList<Inputs.AccountStorageArgs>());
            set => _storages = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
        /// </summary>
        [Input("customSubdomainName")]
        public Input<string>? CustomSubdomainName { get; set; }

        /// <summary>
        /// The endpoint used to connect to the Cognitive Service Account.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// List of FQDNs allowed for the Cognitive Account.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AccountIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the type of Cognitive Service Account that should be created. Possible values are `Academic`, `AnomalyDetector`, `Bing.Autosuggest`, `Bing.Autosuggest.v7`, `Bing.CustomSearch`, `Bing.Search`, `Bing.Search.v7`, `Bing.Speech`, `Bing.SpellCheck`, `Bing.SpellCheck.v7`, `CognitiveServices`, `ComputerVision`, `ContentModerator`, `CustomSpeech`, `CustomVision.Prediction`, `CustomVision.Training`, `Emotion`, `Face`,`FormRecognizer`, `ImmersiveReader`, `LUIS`, `LUIS.Authoring`, `MetricsAdvisor`, `Personalizer`, `QnAMaker`, `Recommendations`, `SpeakerRecognition`, `Speech`, `SpeechServices`, `SpeechTranslation`, `TextAnalytics`, `TextTranslation` and `WebLM`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Whether local authentication methods is enabled for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Input("localAuthEnabled")]
        public Input<bool>? LocalAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Azure AD Client ID (Application ID). This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorAadClientId")]
        public Input<string>? MetricsAdvisorAadClientId { get; set; }

        /// <summary>
        /// The Azure AD Tenant ID. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorAadTenantId")]
        public Input<string>? MetricsAdvisorAadTenantId { get; set; }

        /// <summary>
        /// The super user of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorSuperUserName")]
        public Input<string>? MetricsAdvisorSuperUserName { get; set; }

        /// <summary>
        /// The website name of Metrics Advisor. This attribute is only set when kind is `MetricsAdvisor`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("metricsAdvisorWebsiteName")]
        public Input<string>? MetricsAdvisorWebsiteName { get; set; }

        /// <summary>
        /// Specifies the name of the Cognitive Service Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_acls` block as defined below.
        /// </summary>
        [Input("networkAcls")]
        public Input<Inputs.AccountNetworkAclsGetArgs>? NetworkAcls { get; set; }

        /// <summary>
        /// Whether outbound network access is restricted for the Cognitive Account. Defaults to `false`.
        /// </summary>
        [Input("outboundNetworkAccessRestrited")]
        public Input<bool>? OutboundNetworkAccessRestrited { get; set; }

        /// <summary>
        /// A primary access key which can be used to connect to the Cognitive Service Account.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// A URL to link a QnAMaker cognitive account to a QnA runtime.
        /// </summary>
        [Input("qnaRuntimeEndpoint")]
        public Input<string>? QnaRuntimeEndpoint { get; set; }

        /// <summary>
        /// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The secondary access key which can be used to connect to the Cognitive Service Account.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        /// <summary>
        /// Specifies the SKU Name for this Cognitive Service Account. Possible values are `F0`, `F1`, `S`, `S0`, `S1`, `S2`, `S3`, `S4`, `S5`, `S6`, `P0`, `P1`, and `P2`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("storages")]
        private InputList<Inputs.AccountStorageGetArgs>? _storages;

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        public InputList<Inputs.AccountStorageGetArgs> Storages
        {
            get => _storages ?? (_storages = new InputList<Inputs.AccountStorageGetArgs>());
            set => _storages = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AccountState()
        {
        }
    }
}
