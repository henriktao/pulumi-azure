// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MixedReality
{
    /// <summary>
    /// Manages an Azure Spatial Anchors Account.
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
    ///         var exampleSpatialAnchorsAccount = new Azure.MixedReality.SpatialAnchorsAccount("exampleSpatialAnchorsAccount", new Azure.MixedReality.SpatialAnchorsAccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Spatial Anchors Account can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.MixedReality/spatialAnchorsAccounts/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount")]
    public partial class SpatialAnchorsAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Spatial Anchors Account.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a SpatialAnchorsAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpatialAnchorsAccount(string name, SpatialAnchorsAccountArgs args, CustomResourceOptions? options = null)
            : base("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, args ?? new SpatialAnchorsAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpatialAnchorsAccount(string name, Input<string> id, SpatialAnchorsAccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:mixedreality/spatialAnchorsAccount:SpatialAnchorsAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpatialAnchorsAccount Get(string name, Input<string> id, SpatialAnchorsAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new SpatialAnchorsAccount(name, id, state, options);
        }
    }

    public sealed class SpatialAnchorsAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Spatial Anchors Account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public SpatialAnchorsAccountArgs()
        {
        }
    }

    public sealed class SpatialAnchorsAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Spatial Anchors Account.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

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

        public SpatialAnchorsAccountState()
        {
        }
    }
}
