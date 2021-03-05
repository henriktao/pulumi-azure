// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages a Ms Sql Database Extended Auditing Policy.
    /// 
    /// &gt; **NOTE:** The Database Extended Auditing Policy Can be set inline here as well as with the mssql_database_extended_auditing_policy resource resource. You can only use one or the other and using both will cause a conflict.
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
    ///         var exampleServer = new Azure.MSSql.Server("exampleServer", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "AdminPassword123!",
    ///         });
    ///         var exampleDatabase = new Azure.MSSql.Database("exampleDatabase", new Azure.MSSql.DatabaseArgs
    ///         {
    ///             ServerId = exampleServer.Id,
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleDatabaseExtendedAuditingPolicy = new Azure.MSSql.DatabaseExtendedAuditingPolicy("exampleDatabaseExtendedAuditingPolicy", new Azure.MSSql.DatabaseExtendedAuditingPolicyArgs
    ///         {
    ///             DatabaseId = exampleDatabase.Id,
    ///             StorageEndpoint = exampleAccount.PrimaryBlobEndpoint,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///             StorageAccountAccessKeyIsSecondary = false,
    ///             RetentionInDays = 6,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ms Sql Database Extended Auditing Policys can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy")]
    public partial class DatabaseExtendedAuditingPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        [Output("logMonitoringEnabled")]
        public Output<bool?> LogMonitoringEnabled { get; private set; } = null!;

        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Output("retentionInDays")]
        public Output<int?> RetentionInDays { get; private set; } = null!;

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Output("storageAccountAccessKey")]
        public Output<string?> StorageAccountAccessKey { get; private set; } = null!;

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Output("storageAccountAccessKeyIsSecondary")]
        public Output<bool?> StorageAccountAccessKeyIsSecondary { get; private set; } = null!;

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Output("storageEndpoint")]
        public Output<string?> StorageEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseExtendedAuditingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseExtendedAuditingPolicy(string name, DatabaseExtendedAuditingPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy", name, args ?? new DatabaseExtendedAuditingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseExtendedAuditingPolicy(string name, Input<string> id, DatabaseExtendedAuditingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseExtendedAuditingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseExtendedAuditingPolicy Get(string name, Input<string> id, DatabaseExtendedAuditingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseExtendedAuditingPolicy(name, id, state, options);
        }
    }

    public sealed class DatabaseExtendedAuditingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("logMonitoringEnabled")]
        public Input<bool>? LogMonitoringEnabled { get; set; }

        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Input("storageAccountAccessKeyIsSecondary")]
        public Input<bool>? StorageAccountAccessKeyIsSecondary { get; set; }

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public DatabaseExtendedAuditingPolicyArgs()
        {
        }
    }

    public sealed class DatabaseExtendedAuditingPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        [Input("logMonitoringEnabled")]
        public Input<bool>? LogMonitoringEnabled { get; set; }

        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Input("storageAccountAccessKeyIsSecondary")]
        public Input<bool>? StorageAccountAccessKeyIsSecondary { get; set; }

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public DatabaseExtendedAuditingPolicyState()
        {
        }
    }
}
