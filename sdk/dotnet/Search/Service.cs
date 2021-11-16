// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Search
{
    /// <summary>
    /// Manages a Search Service.
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
    ///         var exampleService = new Azure.Search.Service("exampleService", new Azure.Search.ServiceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = "standard",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Search Services can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:search/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1
    /// ```
    /// </summary>
    [AzureResourceType("azure:search/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of IPv4 addresses or CIDRs that are allowed access to the search service endpoint.
        /// </summary>
        [Output("allowedIps")]
        public Output<ImmutableArray<string>> AllowedIps { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServiceIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of partitions which should be created.
        /// </summary>
        [Output("partitionCount")]
        public Output<int> PartitionCount { get; private set; } = null!;

        /// <summary>
        /// The Primary Key used for Search Service Administration.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this resource. Defaults to `true`.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// A `query_keys` block as defined below.
        /// </summary>
        [Output("queryKeys")]
        public Output<ImmutableArray<Outputs.ServiceQueryKey>> QueryKeys { get; private set; } = null!;

        /// <summary>
        /// The number of replica's which should be created.
        /// </summary>
        [Output("replicaCount")]
        public Output<int> ReplicaCount { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Key used for Search Service Administration.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2`, `standard3`, `storage_optimized_l1` and `storage_optimized_l2`. Changing this forces a new Search Service to be created.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Search Service.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:search/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:search/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("allowedIps")]
        private InputList<string>? _allowedIps;

        /// <summary>
        /// A list of IPv4 addresses or CIDRs that are allowed access to the search service endpoint.
        /// </summary>
        public InputList<string> AllowedIps
        {
            get => _allowedIps ?? (_allowedIps = new InputList<string>());
            set => _allowedIps = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of partitions which should be created.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this resource. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The number of replica's which should be created.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2`, `standard3`, `storage_optimized_l1` and `storage_optimized_l2`. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Search Service.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("allowedIps")]
        private InputList<string>? _allowedIps;

        /// <summary>
        /// A list of IPv4 addresses or CIDRs that are allowed access to the search service endpoint.
        /// </summary>
        public InputList<string> AllowedIps
        {
            get => _allowedIps ?? (_allowedIps = new InputList<string>());
            set => _allowedIps = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of partitions which should be created.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// The Primary Key used for Search Service Administration.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this resource. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        [Input("queryKeys")]
        private InputList<Inputs.ServiceQueryKeyGetArgs>? _queryKeys;

        /// <summary>
        /// A `query_keys` block as defined below.
        /// </summary>
        public InputList<Inputs.ServiceQueryKeyGetArgs> QueryKeys
        {
            get => _queryKeys ?? (_queryKeys = new InputList<Inputs.ServiceQueryKeyGetArgs>());
            set => _queryKeys = value;
        }

        /// <summary>
        /// The number of replica's which should be created.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Search Service should exist. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Key used for Search Service Administration.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// The SKU which should be used for this Search Service. Possible values are `basic`, `free`, `standard`, `standard2`, `standard3`, `storage_optimized_l1` and `storage_optimized_l2`. Changing this forces a new Search Service to be created.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Search Service.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceState()
        {
        }
    }
}
