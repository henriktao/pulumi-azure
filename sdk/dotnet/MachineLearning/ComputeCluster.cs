// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MachineLearning
{
    /// <summary>
    /// Manages a Machine Learning Compute Cluster.
    /// **NOTE:** At this point in time the resource cannot be updated (not supported by the backend Azure Go SDK). Therefore it can only be created and deleted, not updated. At the moment, there is also no possibility to specify ssh User Account Credentials to ssh into the compute cluster.
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
    ///             Location = "west europe",
    ///             Tags = 
    ///             {
    ///                 { "stage", "example" },
    ///             },
    ///         });
    ///         var exampleInsights = new Azure.AppInsights.Insights("exampleInsights", new Azure.AppInsights.InsightsArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApplicationType = "web",
    ///         });
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "standard",
    ///             PurgeProtectionEnabled = true,
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleWorkspace = new Azure.MachineLearning.Workspace("exampleWorkspace", new Azure.MachineLearning.WorkspaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApplicationInsightsId = exampleInsights.Id,
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             StorageAccountId = exampleAccount.Id,
    ///             Identity = new Azure.MachineLearning.Inputs.WorkspaceIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.1.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.1.0.0/24",
    ///             },
    ///         });
    ///         var test = new Azure.MachineLearning.ComputeCluster("test", new Azure.MachineLearning.ComputeClusterArgs
    ///         {
    ///             Location = "West Europe",
    ///             VmPriority = "LowPriority",
    ///             VmSize = "Standard_DS2_v2",
    ///             MachineLearningWorkspaceId = exampleWorkspace.Id,
    ///             SubnetResourceId = exampleSubnet.Id,
    ///             ScaleSettings = new Azure.MachineLearning.Inputs.ComputeClusterScaleSettingsArgs
    ///             {
    ///                 MinNodeCount = 0,
    ///                 MaxNodeCount = 1,
    ///                 ScaleDownNodesAfterIdleDuration = "PT30S",
    ///             },
    ///             Identity = new Azure.MachineLearning.Inputs.ComputeClusterIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Machine Learning Compute Clusters can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:machinelearning/computeCluster:ComputeCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
    /// ```
    /// </summary>
    [AzureResourceType("azure:machinelearning/computeCluster:ComputeCluster")]
    public partial class ComputeCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ComputeClusterIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("machineLearningWorkspaceId")]
        public Output<string> MachineLearningWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `scale_settings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("scaleSettings")]
        public Output<Outputs.ComputeClusterScaleSettings> ScaleSettings { get; private set; } = null!;

        /// <summary>
        /// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("ssh")]
        public Output<Outputs.ComputeClusterSsh?> Ssh { get; private set; } = null!;

        /// <summary>
        /// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("sshPublicAccessEnabled")]
        public Output<bool> SshPublicAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("subnetResourceId")]
        public Output<string?> SubnetResourceId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("vmPriority")]
        public Output<string> VmPriority { get; private set; } = null!;

        /// <summary>
        /// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Output("vmSize")]
        public Output<string> VmSize { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeCluster(string name, ComputeClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:machinelearning/computeCluster:ComputeCluster", name, args ?? new ComputeClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeCluster(string name, Input<string> id, ComputeClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:machinelearning/computeCluster:ComputeCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeCluster Get(string name, Input<string> id, ComputeClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeCluster(name, id, state, options);
        }
    }

    public sealed class ComputeClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ComputeClusterIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("machineLearningWorkspaceId", required: true)]
        public Input<string> MachineLearningWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `scale_settings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("scaleSettings", required: true)]
        public Input<Inputs.ComputeClusterScaleSettingsArgs> ScaleSettings { get; set; } = null!;

        /// <summary>
        /// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ComputeClusterSshArgs>? Ssh { get; set; }

        /// <summary>
        /// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("sshPublicAccessEnabled")]
        public Input<bool>? SshPublicAccessEnabled { get; set; }

        /// <summary>
        /// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("subnetResourceId")]
        public Input<string>? SubnetResourceId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("vmPriority", required: true)]
        public Input<string> VmPriority { get; set; } = null!;

        /// <summary>
        /// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public ComputeClusterArgs()
        {
        }
    }

    public sealed class ComputeClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ComputeClusterIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("machineLearningWorkspaceId")]
        public Input<string>? MachineLearningWorkspaceId { get; set; }

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `scale_settings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("scaleSettings")]
        public Input<Inputs.ComputeClusterScaleSettingsGetArgs>? ScaleSettings { get; set; }

        /// <summary>
        /// Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ComputeClusterSshGetArgs>? Ssh { get; set; }

        /// <summary>
        /// A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("sshPublicAccessEnabled")]
        public Input<bool>? SshPublicAccessEnabled { get; set; }

        /// <summary>
        /// The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("subnetResourceId")]
        public Input<string>? SubnetResourceId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("vmPriority")]
        public Input<string>? VmPriority { get; set; }

        /// <summary>
        /// The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        public ComputeClusterState()
        {
        }
    }
}
