// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    /// <summary>
    /// Manages a Data Factory Managed Private Endpoint.
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
    ///         var exampleFactory = new Azure.DataFactory.Factory("exampleFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ManagedVirtualNetworkEnabled = true,
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountKind = "BlobStorage",
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleManagedPrivateEndpoint = new Azure.DataFactory.ManagedPrivateEndpoint("exampleManagedPrivateEndpoint", new Azure.DataFactory.ManagedPrivateEndpointArgs
    ///         {
    ///             DataFactoryId = exampleFactory.Id,
    ///             TargetResourceId = exampleAccount.Id,
    ///             SubresourceName = "blob",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Managed Private Endpoint can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint")]
    public partial class ManagedPrivateEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dataFactoryId")]
        public Output<string> DataFactoryId { get; private set; } = null!;

        /// <summary>
        /// Fully qualified domain names. Changing this forces a new resource to be created.
        /// </summary>
        [Output("fqdns")]
        public Output<ImmutableArray<string>> Fqdns { get; private set; } = null!;

        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subresourceName")]
        public Output<string?> SubresourceName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedPrivateEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedPrivateEndpoint(string name, ManagedPrivateEndpointArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint", name, args ?? new ManagedPrivateEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedPrivateEndpoint(string name, Input<string> id, ManagedPrivateEndpointState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedPrivateEndpoint Get(string name, Input<string> id, ManagedPrivateEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedPrivateEndpoint(name, id, state, options);
        }
    }

    public sealed class ManagedPrivateEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryId", required: true)]
        public Input<string> DataFactoryId { get; set; } = null!;

        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// Fully qualified domain names. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subresourceName")]
        public Input<string>? SubresourceName { get; set; }

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public ManagedPrivateEndpointArgs()
        {
        }
    }

    public sealed class ManagedPrivateEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryId")]
        public Input<string>? DataFactoryId { get; set; }

        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// Fully qualified domain names. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        /// <summary>
        /// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subresourceName")]
        public Input<string>? SubresourceName { get; set; }

        /// <summary>
        /// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        public ManagedPrivateEndpointState()
        {
        }
    }
}
