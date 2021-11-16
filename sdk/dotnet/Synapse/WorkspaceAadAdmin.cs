// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Synapse
{
    /// <summary>
    /// Manages an Azure Active Directory Administrator setting for a Synapse Workspace
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
    ///             AccountReplicationType = "LRS",
    ///             AccountKind = "StorageV2",
    ///             IsHnsEnabled = true,
    ///         });
    ///         var exampleDataLakeGen2Filesystem = new Azure.Storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", new Azure.Storage.DataLakeGen2FilesystemArgs
    ///         {
    ///             StorageAccountId = exampleAccount.Id,
    ///         });
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = azurerm_resource_group.Exampl.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "standard",
    ///             PurgeProtectionEnabled = true,
    ///         });
    ///         var deployer = new Azure.KeyVault.AccessPolicy("deployer", new Azure.KeyVault.AccessPolicyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///             KeyPermissions = 
    ///             {
    ///                 "create",
    ///                 "get",
    ///                 "delete",
    ///                 "purge",
    ///             },
    ///         });
    ///         var exampleKey = new Azure.KeyVault.Key("exampleKey", new Azure.KeyVault.KeyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             KeyType = "RSA",
    ///             KeySize = 2048,
    ///             KeyOpts = 
    ///             {
    ///                 "unwrapKey",
    ///                 "wrapKey",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 deployer,
    ///             },
    ///         });
    ///         var exampleWorkspace = new Azure.Synapse.Workspace("exampleWorkspace", new Azure.Synapse.WorkspaceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             StorageDataLakeGen2FilesystemId = exampleDataLakeGen2Filesystem.Id,
    ///             SqlAdministratorLogin = "sqladminuser",
    ///             SqlAdministratorLoginPassword = "H@Sh1CoR3!",
    ///             Tags = 
    ///             {
    ///                 { "Env", "production" },
    ///             },
    ///         });
    ///         var test = new Azure.Synapse.WorkspaceAadAdmin("test", new Azure.Synapse.WorkspaceAadAdminArgs
    ///         {
    ///             SynapseWorkspaceId = azurerm_synapse_workspace.Test.Id,
    ///             Login = "AzureAD Admin",
    ///             ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Synapse Workspace Azure AD Administrator can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:synapse/workspaceAadAdmin:WorkspaceAadAdmin example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/administrators/activeDirectory
    /// ```
    /// </summary>
    [AzureResourceType("azure:synapse/workspaceAadAdmin:WorkspaceAadAdmin")]
    public partial class WorkspaceAadAdmin : Pulumi.CustomResource
    {
        /// <summary>
        /// The login name of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// The object id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        /// </summary>
        [Output("synapseWorkspaceId")]
        public Output<string> SynapseWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The tenant id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceAadAdmin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceAadAdmin(string name, WorkspaceAadAdminArgs args, CustomResourceOptions? options = null)
            : base("azure:synapse/workspaceAadAdmin:WorkspaceAadAdmin", name, args ?? new WorkspaceAadAdminArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceAadAdmin(string name, Input<string> id, WorkspaceAadAdminState? state = null, CustomResourceOptions? options = null)
            : base("azure:synapse/workspaceAadAdmin:WorkspaceAadAdmin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkspaceAadAdmin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceAadAdmin Get(string name, Input<string> id, WorkspaceAadAdminState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkspaceAadAdmin(name, id, state, options);
        }
    }

    public sealed class WorkspaceAadAdminArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login name of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// The object id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        /// </summary>
        [Input("synapseWorkspaceId", required: true)]
        public Input<string> SynapseWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The tenant id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public WorkspaceAadAdminArgs()
        {
        }
    }

    public sealed class WorkspaceAadAdminState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The login name of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        /// <summary>
        /// The object id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        /// </summary>
        [Input("synapseWorkspaceId")]
        public Input<string>? SynapseWorkspaceId { get; set; }

        /// <summary>
        /// The tenant id of the Azure AD Administrator of this Synapse Workspace.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public WorkspaceAadAdminState()
        {
        }
    }
}
