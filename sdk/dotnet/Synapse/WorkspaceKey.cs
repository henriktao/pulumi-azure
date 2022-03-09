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
    /// Manages a Synapse Workspace.
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
    ///         var exampleWorkspace = new Azure.Synapse.Workspace("exampleWorkspace", new Azure.Synapse.WorkspaceArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             StorageDataLakeGen2FilesystemId = exampleDataLakeGen2Filesystem.Id,
    ///             SqlAdministratorLogin = "sqladminuser",
    ///             SqlAdministratorLoginPassword = "H@Sh1CoR3!",
    ///             AadAdmin = new Azure.Synapse.Inputs.WorkspaceAadAdminArgs
    ///             {
    ///                 Login = "AzureAD Admin",
    ///                 ObjectId = "00000000-0000-0000-0000-000000000000",
    ///                 TenantId = "00000000-0000-0000-0000-000000000000",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Env", "production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Creating A Workspace With Customer Managed Key And Azure AD Admin
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
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
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
    ///             CustomerManagedKey = new Azure.Synapse.Inputs.WorkspaceCustomerManagedKeyArgs
    ///             {
    ///                 KeyVersionlessId = exampleKey.VersionlessId,
    ///                 KeyName = "enckey",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Env", "production" },
    ///             },
    ///         });
    ///         var workspacePolicy = new Azure.KeyVault.AccessPolicy("workspacePolicy", new Azure.KeyVault.AccessPolicyArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             TenantId = exampleWorkspace.Identities.Apply(identities =&gt; identities[0].TenantId),
    ///             ObjectId = exampleWorkspace.Identities.Apply(identities =&gt; identities[0].PrincipalId),
    ///             KeyPermissions = 
    ///             {
    ///                 "Get",
    ///                 "WrapKey",
    ///                 "UnwrapKey",
    ///             },
    ///         });
    ///         var exampleWorkspaceKey = new Azure.Synapse.WorkspaceKey("exampleWorkspaceKey", new Azure.Synapse.WorkspaceKeyArgs
    ///         {
    ///             CustomerManagedKeyVersionlessId = exampleKey.VersionlessId,
    ///             SynapseWorkspaceId = exampleWorkspace.Id,
    ///             Active = true,
    ///             CustomerManagedKeyName = "enckey",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 workspacePolicy,
    ///             },
    ///         });
    ///         var exampleWorkspaceAadAdmin = new Azure.Synapse.WorkspaceAadAdmin("exampleWorkspaceAadAdmin", new Azure.Synapse.WorkspaceAadAdminArgs
    ///         {
    ///             SynapseWorkspaceId = exampleWorkspace.Id,
    ///             Login = "AzureAD Admin",
    ///             ObjectId = "00000000-0000-0000-0000-000000000000",
    ///             TenantId = "00000000-0000-0000-0000-000000000000",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleWorkspaceKey,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Synapse Workspace can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:synapse/workspaceKey:WorkspaceKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
    /// ```
    /// </summary>
    [AzureResourceType("azure:synapse/workspaceKey:WorkspaceKey")]
    public partial class WorkspaceKey : Pulumi.CustomResource
    {
        [Output("active")]
        public Output<bool> Active { get; private set; } = null!;

        [Output("cusomterManagedKeyName")]
        public Output<string> CusomterManagedKeyName { get; private set; } = null!;

        [Output("customerManagedKeyName")]
        public Output<string> CustomerManagedKeyName { get; private set; } = null!;

        [Output("customerManagedKeyVersionlessId")]
        public Output<string?> CustomerManagedKeyVersionlessId { get; private set; } = null!;

        [Output("synapseWorkspaceId")]
        public Output<string> SynapseWorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceKey(string name, WorkspaceKeyArgs args, CustomResourceOptions? options = null)
            : base("azure:synapse/workspaceKey:WorkspaceKey", name, args ?? new WorkspaceKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceKey(string name, Input<string> id, WorkspaceKeyState? state = null, CustomResourceOptions? options = null)
            : base("azure:synapse/workspaceKey:WorkspaceKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkspaceKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceKey Get(string name, Input<string> id, WorkspaceKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkspaceKey(name, id, state, options);
        }
    }

    public sealed class WorkspaceKeyArgs : Pulumi.ResourceArgs
    {
        [Input("active", required: true)]
        public Input<bool> Active { get; set; } = null!;

        [Input("cusomterManagedKeyName")]
        public Input<string>? CusomterManagedKeyName { get; set; }

        [Input("customerManagedKeyName")]
        public Input<string>? CustomerManagedKeyName { get; set; }

        [Input("customerManagedKeyVersionlessId")]
        public Input<string>? CustomerManagedKeyVersionlessId { get; set; }

        [Input("synapseWorkspaceId", required: true)]
        public Input<string> SynapseWorkspaceId { get; set; } = null!;

        public WorkspaceKeyArgs()
        {
        }
    }

    public sealed class WorkspaceKeyState : Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("cusomterManagedKeyName")]
        public Input<string>? CusomterManagedKeyName { get; set; }

        [Input("customerManagedKeyName")]
        public Input<string>? CustomerManagedKeyName { get; set; }

        [Input("customerManagedKeyVersionlessId")]
        public Input<string>? CustomerManagedKeyVersionlessId { get; set; }

        [Input("synapseWorkspaceId")]
        public Input<string>? SynapseWorkspaceId { get; set; }

        public WorkspaceKeyState()
        {
        }
    }
}
