// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Subnet Service Endpoint Storage Policy.
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
    ///         var exampleSubnetServiceEndpointStoragePolicy = new Azure.Network.SubnetServiceEndpointStoragePolicy("exampleSubnetServiceEndpointStoragePolicy", new Azure.Network.SubnetServiceEndpointStoragePolicyArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Definition = new Azure.Network.Inputs.SubnetServiceEndpointStoragePolicyDefinitionArgs
    ///             {
    ///                 Name = "name1",
    ///                 Description = "definition1",
    ///                 ServiceResources = 
    ///                 {
    ///                     exampleResourceGroup.Id,
    ///                     exampleAccount.Id,
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
    /// Subnet Service Endpoint Policies can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/subnetServiceEndpointStoragePolicy:SubnetServiceEndpointStoragePolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/serviceEndpointPolicies/policy1
    /// ```
    /// </summary>
    [AzureResourceType("azure:network/subnetServiceEndpointStoragePolicy:SubnetServiceEndpointStoragePolicy")]
    public partial class SubnetServiceEndpointStoragePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A `definition` block as defined below
        /// </summary>
        [Output("definition")]
        public Output<Outputs.SubnetServiceEndpointStoragePolicyDefinition?> Definition { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Subnet Service Endpoint Storage Policy. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Subnet Service Endpoint Storage Policy.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetServiceEndpointStoragePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetServiceEndpointStoragePolicy(string name, SubnetServiceEndpointStoragePolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:network/subnetServiceEndpointStoragePolicy:SubnetServiceEndpointStoragePolicy", name, args ?? new SubnetServiceEndpointStoragePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubnetServiceEndpointStoragePolicy(string name, Input<string> id, SubnetServiceEndpointStoragePolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/subnetServiceEndpointStoragePolicy:SubnetServiceEndpointStoragePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubnetServiceEndpointStoragePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetServiceEndpointStoragePolicy Get(string name, Input<string> id, SubnetServiceEndpointStoragePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SubnetServiceEndpointStoragePolicy(name, id, state, options);
        }
    }

    public sealed class SubnetServiceEndpointStoragePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `definition` block as defined below
        /// </summary>
        [Input("definition")]
        public Input<Inputs.SubnetServiceEndpointStoragePolicyDefinitionArgs>? Definition { get; set; }

        /// <summary>
        /// The Azure Region where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Subnet Service Endpoint Storage Policy. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Subnet Service Endpoint Storage Policy.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SubnetServiceEndpointStoragePolicyArgs()
        {
        }
    }

    public sealed class SubnetServiceEndpointStoragePolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `definition` block as defined below
        /// </summary>
        [Input("definition")]
        public Input<Inputs.SubnetServiceEndpointStoragePolicyDefinitionGetArgs>? Definition { get; set; }

        /// <summary>
        /// The Azure Region where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Subnet Service Endpoint Storage Policy. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Subnet Service Endpoint Storage Policy should exist. Changing this forces a new Subnet Service Endpoint Storage Policy to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Subnet Service Endpoint Storage Policy.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SubnetServiceEndpointStoragePolicyState()
        {
        }
    }
}
