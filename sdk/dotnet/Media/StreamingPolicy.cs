// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media
{
    /// <summary>
    /// Manages a Streaming Policy.
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
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleServiceAccount = new Azure.Media.ServiceAccount("exampleServiceAccount", new Azure.Media.ServiceAccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.Media.Inputs.ServiceAccountStorageAccountArgs
    ///                 {
    ///                     Id = exampleAccount.Id,
    ///                     IsPrimary = true,
    ///                 },
    ///             },
    ///         });
    ///         var exampleStreamingPolicy = new Azure.Media.StreamingPolicy("exampleStreamingPolicy", new Azure.Media.StreamingPolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             MediaServicesAccountName = exampleServiceAccount.Name,
    ///             CommonEncryptionCenc = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCencArgs
    ///             {
    ///                 EnabledProtocols = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCencEnabledProtocolsArgs
    ///                 {
    ///                     Download = false,
    ///                     Dash = true,
    ///                     Hls = false,
    ///                     SmoothStreaming = false,
    ///                 },
    ///                 DrmPlayready = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCencDrmPlayreadyArgs
    ///                 {
    ///                     CustomLicenseAcquisitionUrlTemplate = "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}",
    ///                     CustomAttributes = "PlayReady CustomAttributes",
    ///                 },
    ///                 DrmWidevineCustomLicenseAcquisitionUrlTemplate = "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId}",
    ///             },
    ///             CommonEncryptionCbcs = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCbcsArgs
    ///             {
    ///                 EnabledProtocols = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCbcsEnabledProtocolsArgs
    ///                 {
    ///                     Download = false,
    ///                     Dash = true,
    ///                     Hls = false,
    ///                     SmoothStreaming = false,
    ///                 },
    ///                 DrmFairplay = new Azure.Media.Inputs.StreamingPolicyCommonEncryptionCbcsDrmFairplayArgs
    ///                 {
    ///                     CustomLicenseAcquisitionUrlTemplate = "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}",
    ///                     AllowPersistentLicense = true,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Streaming Policies can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:media/streamingPolicy:StreamingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/streamingpolicies/policy1
    /// ```
    /// </summary>
    [AzureResourceType("azure:media/streamingPolicy:StreamingPolicy")]
    public partial class StreamingPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A `common_encryption_cbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("commonEncryptionCbcs")]
        public Output<Outputs.StreamingPolicyCommonEncryptionCbcs?> CommonEncryptionCbcs { get; private set; } = null!;

        /// <summary>
        /// A `common_encryption_cenc` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("commonEncryptionCenc")]
        public Output<Outputs.StreamingPolicyCommonEncryptionCenc?> CommonEncryptionCenc { get; private set; } = null!;

        /// <summary>
        /// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("defaultContentKeyPolicyName")]
        public Output<string?> DefaultContentKeyPolicyName { get; private set; } = null!;

        /// <summary>
        /// The Media Services account name. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("mediaServicesAccountName")]
        public Output<string> MediaServicesAccountName { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `no_encryption_enabled_protocols` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("noEncryptionEnabledProtocols")]
        public Output<Outputs.StreamingPolicyNoEncryptionEnabledProtocols?> NoEncryptionEnabledProtocols { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a StreamingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamingPolicy(string name, StreamingPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:media/streamingPolicy:StreamingPolicy", name, args ?? new StreamingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamingPolicy(string name, Input<string> id, StreamingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:media/streamingPolicy:StreamingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StreamingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamingPolicy Get(string name, Input<string> id, StreamingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new StreamingPolicy(name, id, state, options);
        }
    }

    public sealed class StreamingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `common_encryption_cbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("commonEncryptionCbcs")]
        public Input<Inputs.StreamingPolicyCommonEncryptionCbcsArgs>? CommonEncryptionCbcs { get; set; }

        /// <summary>
        /// A `common_encryption_cenc` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("commonEncryptionCenc")]
        public Input<Inputs.StreamingPolicyCommonEncryptionCencArgs>? CommonEncryptionCenc { get; set; }

        /// <summary>
        /// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("defaultContentKeyPolicyName")]
        public Input<string>? DefaultContentKeyPolicyName { get; set; }

        /// <summary>
        /// The Media Services account name. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("mediaServicesAccountName", required: true)]
        public Input<string> MediaServicesAccountName { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `no_encryption_enabled_protocols` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("noEncryptionEnabledProtocols")]
        public Input<Inputs.StreamingPolicyNoEncryptionEnabledProtocolsArgs>? NoEncryptionEnabledProtocols { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public StreamingPolicyArgs()
        {
        }
    }

    public sealed class StreamingPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `common_encryption_cbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("commonEncryptionCbcs")]
        public Input<Inputs.StreamingPolicyCommonEncryptionCbcsGetArgs>? CommonEncryptionCbcs { get; set; }

        /// <summary>
        /// A `common_encryption_cenc` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("commonEncryptionCenc")]
        public Input<Inputs.StreamingPolicyCommonEncryptionCencGetArgs>? CommonEncryptionCenc { get; set; }

        /// <summary>
        /// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("defaultContentKeyPolicyName")]
        public Input<string>? DefaultContentKeyPolicyName { get; set; }

        /// <summary>
        /// The Media Services account name. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("mediaServicesAccountName")]
        public Input<string>? MediaServicesAccountName { get; set; }

        /// <summary>
        /// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `no_encryption_enabled_protocols` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("noEncryptionEnabledProtocols")]
        public Input<Inputs.StreamingPolicyNoEncryptionEnabledProtocolsGetArgs>? NoEncryptionEnabledProtocols { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        public StreamingPolicyState()
        {
        }
    }
}
