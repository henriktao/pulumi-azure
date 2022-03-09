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
    /// Manages a Healthbot Service.
    /// 
    /// ## Import
    /// 
    /// Healthbot Service can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:bot/healthbot:Healthbot example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HealthBot/healthBots/bot1
    /// ```
    /// </summary>
    [AzureResourceType("azure:bot/healthbot:Healthbot")]
    public partial class Healthbot : Pulumi.CustomResource
    {
        /// <summary>
        /// The management portal url.
        /// </summary>
        [Output("botManagementPortalUrl")]
        public Output<string> BotManagementPortalUrl { get; private set; } = null!;

        /// <summary>
        /// Specifies The Azure Region where the resource exists. CHanging this force a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies The name of the Healthbot Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies The name of the Resource Group in which to create the Healtbot Service. CHaning this
        /// forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for the sku of the service. Possible values are "F0" and "S1".
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the service.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Healthbot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Healthbot(string name, HealthbotArgs args, CustomResourceOptions? options = null)
            : base("azure:bot/healthbot:Healthbot", name, args ?? new HealthbotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Healthbot(string name, Input<string> id, HealthbotState? state = null, CustomResourceOptions? options = null)
            : base("azure:bot/healthbot:Healthbot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Healthbot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Healthbot Get(string name, Input<string> id, HealthbotState? state = null, CustomResourceOptions? options = null)
        {
            return new Healthbot(name, id, state, options);
        }
    }

    public sealed class HealthbotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies The Azure Region where the resource exists. CHanging this force a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies The name of the Healthbot Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies The name of the Resource Group in which to create the Healtbot Service. CHaning this
        /// forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name which should be used for the sku of the service. Possible values are "F0" and "S1".
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the service.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public HealthbotArgs()
        {
        }
    }

    public sealed class HealthbotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The management portal url.
        /// </summary>
        [Input("botManagementPortalUrl")]
        public Input<string>? BotManagementPortalUrl { get; set; }

        /// <summary>
        /// Specifies The Azure Region where the resource exists. CHanging this force a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies The name of the Healthbot Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies The name of the Resource Group in which to create the Healtbot Service. CHaning this
        /// forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The name which should be used for the sku of the service. Possible values are "F0" and "S1".
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the service.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public HealthbotState()
        {
        }
    }
}
