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
    /// Manages a Machine Learning Compute Instance.
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
    ///         var config = new Config();
    ///         var sshKey = config.Get("sshKey") ?? "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqaZoyiz1qbdOQ8xEf6uEu1cCwYowo5FHtsBhqLoDnnp7KUTEBN+L2NxRIfQ781rxV6Iq5jSav6b2Q8z5KiseOlvKA/RF2wqU0UPYqQviQhLmW6THTpmrv/YkUCuzxDpsH7DUDhZcwySLKVVe0Qm3+5N2Ta6UYH3lsDf9R9wTP2K/+vAnflKebuypNlmocIvakFWoZda18FOmsOoIVXQ8HWFNCuw9ZCunMSN62QGamCe3dL5cXlkgHYv7ekJE15IA9aOJcM7e90oeTqo+7HTcWfdu0qQqPWY5ujyMw/llas8tsXY85LFqRnr3gJ02bAscjc477+X+j/gkpFoN1QEmt terraform@demo.tld";
    ///         var exampleComputeInstance = new Azure.MachineLearning.ComputeInstance("exampleComputeInstance", new Azure.MachineLearning.ComputeInstanceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             MachineLearningWorkspaceId = exampleWorkspace.Id,
    ///             VirtualMachineSize = "STANDARD_DS2_V2",
    ///             AuthorizationType = "personal",
    ///             Ssh = new Azure.MachineLearning.Inputs.ComputeInstanceSshArgs
    ///             {
    ///                 PublicKey = sshKey,
    ///             },
    ///             SubnetResourceId = exampleSubnet.Id,
    ///             Description = "foo",
    ///             Tags = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Machine Learning Compute Instances can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:machinelearning/computeInstance:ComputeInstance example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
    /// ```
    /// </summary>
    [AzureResourceType("azure:machinelearning/computeInstance:ComputeInstance")]
    public partial class ComputeInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// A `assign_to_user` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("assignToUser")]
        public Output<Outputs.ComputeInstanceAssignToUser?> AssignToUser { get; private set; } = null!;

        /// <summary>
        /// The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("authorizationType")]
        public Output<string?> AuthorizationType { get; private set; } = null!;

        /// <summary>
        /// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ComputeInstanceIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("machineLearningWorkspaceId")]
        public Output<string> MachineLearningWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("ssh")]
        public Output<Outputs.ComputeInstanceSsh?> Ssh { get; private set; } = null!;

        /// <summary>
        /// Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("subnetResourceId")]
        public Output<string?> SubnetResourceId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Output("virtualMachineSize")]
        public Output<string> VirtualMachineSize { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeInstance(string name, ComputeInstanceArgs args, CustomResourceOptions? options = null)
            : base("azure:machinelearning/computeInstance:ComputeInstance", name, args ?? new ComputeInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeInstance(string name, Input<string> id, ComputeInstanceState? state = null, CustomResourceOptions? options = null)
            : base("azure:machinelearning/computeInstance:ComputeInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeInstance Get(string name, Input<string> id, ComputeInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeInstance(name, id, state, options);
        }
    }

    public sealed class ComputeInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `assign_to_user` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("assignToUser")]
        public Input<Inputs.ComputeInstanceAssignToUserArgs>? AssignToUser { get; set; }

        /// <summary>
        /// The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        /// <summary>
        /// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ComputeInstanceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("machineLearningWorkspaceId", required: true)]
        public Input<string> MachineLearningWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ComputeInstanceSshArgs>? Ssh { get; set; }

        /// <summary>
        /// Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("subnetResourceId")]
        public Input<string>? SubnetResourceId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("virtualMachineSize", required: true)]
        public Input<string> VirtualMachineSize { get; set; } = null!;

        public ComputeInstanceArgs()
        {
        }
    }

    public sealed class ComputeInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `assign_to_user` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("assignToUser")]
        public Input<Inputs.ComputeInstanceAssignToUserGetArgs>? AssignToUser { get; set; }

        /// <summary>
        /// The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        /// <summary>
        /// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ComputeInstanceIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("machineLearningWorkspaceId")]
        public Input<string>? MachineLearningWorkspaceId { get; set; }

        /// <summary>
        /// The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.ComputeInstanceSshGetArgs>? Ssh { get; set; }

        /// <summary>
        /// Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("subnetResourceId")]
        public Input<string>? SubnetResourceId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        [Input("virtualMachineSize")]
        public Input<string>? VirtualMachineSize { get; set; }

        public ComputeInstanceState()
        {
        }
    }
}
