// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Bot
{
    /// <summary>
    /// Manages a Direct Line Speech integration for a Bot Channel
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleAccount = new Azure.Cognitive.Account("exampleAccount", new Azure.Cognitive.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Kind = "SpeechServices",
    ///             SkuName = "S0",
    ///         });
    ///         var exampleChannelsRegistration = new Azure.Bot.ChannelsRegistration("exampleChannelsRegistration", new Azure.Bot.ChannelsRegistrationArgs
    ///         {
    ///             Location = "global",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "F0",
    ///             MicrosoftAppId = current.Apply(current =&gt; current.ClientId),
    ///         });
    ///         var exampleChannelDirectLineSpeech = new Azure.Bot.ChannelDirectLineSpeech("exampleChannelDirectLineSpeech", new Azure.Bot.ChannelDirectLineSpeechArgs
    ///         {
    ///             BotName = exampleChannelsRegistration.Name,
    ///             Location = exampleChannelsRegistration.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             CognitiveServiceLocation = exampleAccount.Location,
    ///             CognitiveServiceAccessKey = exampleAccount.PrimaryAccessKey,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Direct Line Speech Channels can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
    /// ```
    /// </summary>
    [AzureResourceType("azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech")]
    public partial class ChannelDirectLineSpeech : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Output("botName")]
        public Output<string> BotName { get; private set; } = null!;

        /// <summary>
        /// The access key to access the Cognitive Service.
        /// </summary>
        [Output("cognitiveServiceAccessKey")]
        public Output<string> CognitiveServiceAccessKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Cognitive Service resource exists.
        /// </summary>
        [Output("cognitiveServiceLocation")]
        public Output<string> CognitiveServiceLocation { get; private set; } = null!;

        /// <summary>
        /// The custom speech model id for the Direct Line Speech Channel.
        /// </summary>
        [Output("customSpeechModelId")]
        public Output<string?> CustomSpeechModelId { get; private set; } = null!;

        /// <summary>
        /// The custom voice deployment id for the Direct Line Speech Channel.
        /// </summary>
        [Output("customVoiceDeploymentId")]
        public Output<string?> CustomVoiceDeploymentId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a ChannelDirectLineSpeech resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChannelDirectLineSpeech(string name, ChannelDirectLineSpeechArgs args, CustomResourceOptions? options = null)
            : base("azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech", name, args ?? new ChannelDirectLineSpeechArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChannelDirectLineSpeech(string name, Input<string> id, ChannelDirectLineSpeechState? state = null, CustomResourceOptions? options = null)
            : base("azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ChannelDirectLineSpeech resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChannelDirectLineSpeech Get(string name, Input<string> id, ChannelDirectLineSpeechState? state = null, CustomResourceOptions? options = null)
        {
            return new ChannelDirectLineSpeech(name, id, state, options);
        }
    }

    public sealed class ChannelDirectLineSpeechArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName", required: true)]
        public Input<string> BotName { get; set; } = null!;

        /// <summary>
        /// The access key to access the Cognitive Service.
        /// </summary>
        [Input("cognitiveServiceAccessKey", required: true)]
        public Input<string> CognitiveServiceAccessKey { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Cognitive Service resource exists.
        /// </summary>
        [Input("cognitiveServiceLocation", required: true)]
        public Input<string> CognitiveServiceLocation { get; set; } = null!;

        /// <summary>
        /// The custom speech model id for the Direct Line Speech Channel.
        /// </summary>
        [Input("customSpeechModelId")]
        public Input<string>? CustomSpeechModelId { get; set; }

        /// <summary>
        /// The custom voice deployment id for the Direct Line Speech Channel.
        /// </summary>
        [Input("customVoiceDeploymentId")]
        public Input<string>? CustomVoiceDeploymentId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ChannelDirectLineSpeechArgs()
        {
        }
    }

    public sealed class ChannelDirectLineSpeechState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
        /// </summary>
        [Input("botName")]
        public Input<string>? BotName { get; set; }

        /// <summary>
        /// The access key to access the Cognitive Service.
        /// </summary>
        [Input("cognitiveServiceAccessKey")]
        public Input<string>? CognitiveServiceAccessKey { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Cognitive Service resource exists.
        /// </summary>
        [Input("cognitiveServiceLocation")]
        public Input<string>? CognitiveServiceLocation { get; set; }

        /// <summary>
        /// The custom speech model id for the Direct Line Speech Channel.
        /// </summary>
        [Input("customSpeechModelId")]
        public Input<string>? CustomSpeechModelId { get; set; }

        /// <summary>
        /// The custom voice deployment id for the Direct Line Speech Channel.
        /// </summary>
        [Input("customVoiceDeploymentId")]
        public Input<string>? CustomVoiceDeploymentId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public ChannelDirectLineSpeechState()
        {
        }
    }
}
