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
    /// Manages a Azure Machine Learning Workspace
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
    ///             SkuName = "premium",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
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
    ///     }
    /// 
    /// }
    /// ```
    /// ### With Data Encryption
    /// 
    /// &gt; **NOTE:** The Key Vault must enable purge protection.
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
    ///             SkuName = "premium",
    ///             PurgeProtectionEnabled = true,
    ///         });
    ///         var exampleAccessPolicy = new Azure.KeyVault.AccessPolicy("exampleAccessPolicy", new Azure.KeyVault.AccessPolicyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///             KeyPermissions = 
    ///             {
    ///                 "Create",
    ///                 "Get",
    ///                 "Delete",
    ///                 "Purge",
    ///             },
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleKey = new Azure.KeyVault.Key("exampleKey", new Azure.KeyVault.KeyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             KeyType = "RSA",
    ///             KeySize = 2048,
    ///             KeyOpts = 
    ///             {
    ///                 "decrypt",
    ///                 "encrypt",
    ///                 "sign",
    ///                 "unwrapKey",
    ///                 "verify",
    ///                 "wrapKey",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleKeyVault,
    ///                 exampleAccessPolicy,
    ///             },
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
    ///             Encryption = new Azure.MachineLearning.Inputs.WorkspaceEncryptionArgs
    ///             {
    ///                 KeyVaultId = exampleKeyVault.Id,
    ///                 KeyId = exampleKey.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Machine Learning Workspace can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:machinelearning/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1
    /// ```
    /// </summary>
    [AzureResourceType("azure:machinelearning/workspace:Workspace")]
    public partial class Workspace : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationInsightsId")]
        public Output<string> ApplicationInsightsId { get; private set; } = null!;

        /// <summary>
        /// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("containerRegistryId")]
        public Output<string?> ContainerRegistryId { get; private set; } = null!;

        /// <summary>
        /// The description of this Machine Learning Workspace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The URL for the discovery service to identify regional endpoints for machine learning experimentation services.
        /// </summary>
        [Output("discoveryUrl")]
        public Output<string> DiscoveryUrl { get; private set; } = null!;

        [Output("encryption")]
        public Output<Outputs.WorkspaceEncryption?> Encryption { get; private set; } = null!;

        /// <summary>
        /// Friendly name for this Machine Learning Workspace.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
        /// </summary>
        [Output("highBusinessImpact")]
        public Output<bool?> HighBusinessImpact { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.WorkspaceIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// The compute name for image build of the Machine Learning Workspace.
        /// </summary>
        [Output("imageBuildComputeName")]
        public Output<string?> ImageBuildComputeName { get; private set; } = null!;

        /// <summary>
        /// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable public access when this Machine Learning Workspace is behind VNet.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
        /// </summary>
        [Output("skuName")]
        public Output<string?> SkuName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string> StorageAccountId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("azure:machinelearning/workspace:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:machinelearning/workspace:Workspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationInsightsId", required: true)]
        public Input<string> ApplicationInsightsId { get; set; } = null!;

        /// <summary>
        /// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("containerRegistryId")]
        public Input<string>? ContainerRegistryId { get; set; }

        /// <summary>
        /// The description of this Machine Learning Workspace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("encryption")]
        public Input<Inputs.WorkspaceEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Friendly name for this Machine Learning Workspace.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
        /// </summary>
        [Input("highBusinessImpact")]
        public Input<bool>? HighBusinessImpact { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.WorkspaceIdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// The compute name for image build of the Machine Learning Workspace.
        /// </summary>
        [Input("imageBuildComputeName")]
        public Input<string>? ImageBuildComputeName { get; set; }

        /// <summary>
        /// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable public access when this Machine Learning Workspace is behind VNet.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceArgs()
        {
        }
    }

    public sealed class WorkspaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationInsightsId")]
        public Input<string>? ApplicationInsightsId { get; set; }

        /// <summary>
        /// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("containerRegistryId")]
        public Input<string>? ContainerRegistryId { get; set; }

        /// <summary>
        /// The description of this Machine Learning Workspace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The URL for the discovery service to identify regional endpoints for machine learning experimentation services.
        /// </summary>
        [Input("discoveryUrl")]
        public Input<string>? DiscoveryUrl { get; set; }

        [Input("encryption")]
        public Input<Inputs.WorkspaceEncryptionGetArgs>? Encryption { get; set; }

        /// <summary>
        /// Friendly name for this Machine Learning Workspace.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Flag to signal High Business Impact (HBI) data in the workspace and reduce diagnostic data collected by the service
        /// </summary>
        [Input("highBusinessImpact")]
        public Input<bool>? HighBusinessImpact { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.WorkspaceIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The compute name for image build of the Machine Learning Workspace.
        /// </summary>
        [Input("imageBuildComputeName")]
        public Input<string>? ImageBuildComputeName { get; set; }

        /// <summary>
        /// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable public access when this Machine Learning Workspace is behind VNet.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// SKU/edition of the Machine Learning Workspace, possible values are `Basic`. Defaults to `Basic`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceState()
        {
        }
    }
}
