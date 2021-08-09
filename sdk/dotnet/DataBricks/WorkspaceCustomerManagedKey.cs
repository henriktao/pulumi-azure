// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataBricks
{
    /// <summary>
    /// ## Import
    /// 
    /// Databricks Workspace Customer Managed Key can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:databricks/workspaceCustomerManagedKey:WorkspaceCustomerManagedKey workspace1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/customerManagedKey/workspace1
    /// ```
    /// </summary>
    [AzureResourceType("azure:databricks/workspaceCustomerManagedKey:WorkspaceCustomerManagedKey")]
    public partial class WorkspaceCustomerManagedKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Key Vault.
        /// </summary>
        [Output("keyVaultKeyId")]
        public Output<string> KeyVaultKeyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Databricks workspace.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceCustomerManagedKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceCustomerManagedKey(string name, WorkspaceCustomerManagedKeyArgs args, CustomResourceOptions? options = null)
            : base("azure:databricks/workspaceCustomerManagedKey:WorkspaceCustomerManagedKey", name, args ?? new WorkspaceCustomerManagedKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceCustomerManagedKey(string name, Input<string> id, WorkspaceCustomerManagedKeyState? state = null, CustomResourceOptions? options = null)
            : base("azure:databricks/workspaceCustomerManagedKey:WorkspaceCustomerManagedKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkspaceCustomerManagedKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceCustomerManagedKey Get(string name, Input<string> id, WorkspaceCustomerManagedKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkspaceCustomerManagedKey(name, id, state, options);
        }
    }

    public sealed class WorkspaceCustomerManagedKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Key Vault.
        /// </summary>
        [Input("keyVaultKeyId", required: true)]
        public Input<string> KeyVaultKeyId { get; set; } = null!;

        /// <summary>
        /// The ID of the Databricks workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public WorkspaceCustomerManagedKeyArgs()
        {
        }
    }

    public sealed class WorkspaceCustomerManagedKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Key Vault.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// The ID of the Databricks workspace.
        /// </summary>
        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public WorkspaceCustomerManagedKeyState()
        {
        }
    }
}
