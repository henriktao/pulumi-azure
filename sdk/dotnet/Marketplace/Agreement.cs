// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Marketplace
{
    /// <summary>
    /// Allows accepting the Legal Terms for a Marketplace Image.
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
    ///         var barracuda = new Azure.Marketplace.Agreement("barracuda", new Azure.Marketplace.AgreementArgs
    ///         {
    ///             Offer = "waf",
    ///             Plan = "hourly",
    ///             Publisher = "barracudanetworks",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Agreement : Pulumi.CustomResource
    {
        [Output("licenseTextLink")]
        public Output<string> LicenseTextLink { get; private set; } = null!;

        /// <summary>
        /// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Output("offer")]
        public Output<string> Offer { get; private set; } = null!;

        /// <summary>
        /// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("privacyPolicyLink")]
        public Output<string> PrivacyPolicyLink { get; private set; } = null!;

        /// <summary>
        /// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Output("publisher")]
        public Output<string> Publisher { get; private set; } = null!;


        /// <summary>
        /// Create a Agreement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Agreement(string name, AgreementArgs args, CustomResourceOptions? options = null)
            : base("azure:marketplace/agreement:Agreement", name, args ?? new AgreementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Agreement(string name, Input<string> id, AgreementState? state = null, CustomResourceOptions? options = null)
            : base("azure:marketplace/agreement:Agreement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Agreement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Agreement Get(string name, Input<string> id, AgreementState? state = null, CustomResourceOptions? options = null)
        {
            return new Agreement(name, id, state, options);
        }
    }

    public sealed class AgreementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("offer", required: true)]
        public Input<string> Offer { get; set; } = null!;

        /// <summary>
        /// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        public AgreementArgs()
        {
        }
    }

    public sealed class AgreementState : Pulumi.ResourceArgs
    {
        [Input("licenseTextLink")]
        public Input<string>? LicenseTextLink { get; set; }

        /// <summary>
        /// The Offer of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        /// <summary>
        /// The Plan of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("privacyPolicyLink")]
        public Input<string>? PrivacyPolicyLink { get; set; }

        /// <summary>
        /// The Publisher of the Marketplace Image. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        public AgreementState()
        {
        }
    }
}
