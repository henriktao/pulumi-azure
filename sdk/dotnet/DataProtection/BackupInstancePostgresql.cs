// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataProtection
{
    /// <summary>
    /// Manages a Backup Instance to back up PostgreSQL.
    /// 
    /// &gt; **Note**: Before using this resource, there are some prerequisite permissions for configure backup and restore. See more details from https://docs.microsoft.com/en-us/azure/backup/backup-azure-database-postgresql#prerequisite-permissions-for-configure-backup-and-restore.
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
    ///         var exampleServer = new Azure.PostgreSql.Server("exampleServer", new Azure.PostgreSql.ServerArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "B_Gen5_2",
    ///             StorageMb = 5120,
    ///             BackupRetentionDays = 7,
    ///             GeoRedundantBackupEnabled = false,
    ///             AutoGrowEnabled = true,
    ///             AdministratorLogin = "psqladmin",
    ///             AdministratorLoginPassword = "H@Sh1CoR3!",
    ///             Version = "9.5",
    ///             SslEnforcementEnabled = true,
    ///         });
    ///         var exampleFirewallRule = new Azure.PostgreSql.FirewallRule("exampleFirewallRule", new Azure.PostgreSql.FirewallRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServerName = exampleServer.Name,
    ///             StartIpAddress = "0.0.0.0",
    ///             EndIpAddress = "0.0.0.0",
    ///         });
    ///         var exampleDatabase = new Azure.PostgreSql.Database("exampleDatabase", new Azure.PostgreSql.DatabaseArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ServerName = exampleServer.Name,
    ///             Charset = "UTF8",
    ///             Collation = "English_United States.1252",
    ///         });
    ///         var exampleBackupVault = new Azure.DataProtection.BackupVault("exampleBackupVault", new Azure.DataProtection.BackupVaultArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             DatastoreType = "VaultStore",
    ///             Redundancy = "LocallyRedundant",
    ///             Identity = new Azure.DataProtection.Inputs.BackupVaultIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "premium",
    ///             SoftDeleteRetentionDays = 7,
    ///             AccessPolicies = 
    ///             {
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///                     KeyPermissions = 
    ///                     {
    ///                         "create",
    ///                         "get",
    ///                     },
    ///                     SecretPermissions = 
    ///                     {
    ///                         "set",
    ///                         "get",
    ///                         "delete",
    ///                         "purge",
    ///                         "recover",
    ///                     },
    ///                 },
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = exampleBackupVault.Identity.Apply(identity =&gt; identity?.TenantId),
    ///                     ObjectId = exampleBackupVault.Identity.Apply(identity =&gt; identity?.PrincipalId),
    ///                     KeyPermissions = 
    ///                     {
    ///                         "create",
    ///                         "get",
    ///                     },
    ///                     SecretPermissions = 
    ///                     {
    ///                         "set",
    ///                         "get",
    ///                         "delete",
    ///                         "purge",
    ///                         "recover",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleSecret = new Azure.KeyVault.Secret("exampleSecret", new Azure.KeyVault.SecretArgs
    ///         {
    ///             Value = Output.Tuple(exampleServer.Name, exampleDatabase.Name, exampleServer.Name).Apply(values =&gt;
    ///             {
    ///                 var exampleServerName = values.Item1;
    ///                 var exampleDatabaseName = values.Item2;
    ///                 var exampleServerName1 = values.Item3;
    ///                 return $"Server={exampleServerName}.postgres.database.azure.com;Database={exampleDatabaseName};Port=5432;User Id=psqladmin@{exampleServerName1};Password=H@Sh1CoR3!;Ssl Mode=Require;";
    ///             }),
    ///             KeyVaultId = exampleKeyVault.Id,
    ///         });
    ///         var exampleBackupPolicyPostgresql = new Azure.DataProtection.BackupPolicyPostgresql("exampleBackupPolicyPostgresql", new Azure.DataProtection.BackupPolicyPostgresqlArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VaultName = exampleBackupVault.Name,
    ///             BackupRepeatingTimeIntervals = 
    ///             {
    ///                 "R/2021-05-23T02:30:00+00:00/P1W",
    ///             },
    ///             DefaultRetentionDuration = "P4M",
    ///         });
    ///         var exampleAssignment = new Azure.Authorization.Assignment("exampleAssignment", new Azure.Authorization.AssignmentArgs
    ///         {
    ///             Scope = exampleServer.Id,
    ///             RoleDefinitionName = "Reader",
    ///             PrincipalId = exampleBackupVault.Identity.Apply(identity =&gt; identity?.PrincipalId),
    ///         });
    ///         var exampleBackupInstancePostgresql = new Azure.DataProtection.BackupInstancePostgresql("exampleBackupInstancePostgresql", new Azure.DataProtection.BackupInstancePostgresqlArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             VaultId = exampleBackupVault.Id,
    ///             DatabaseId = exampleDatabase.Id,
    ///             BackupPolicyId = exampleBackupPolicyPostgresql.Id,
    ///             DatabaseCredentialKeyVaultSecretId = exampleSecret.VersionlessId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Backup Instance PostgreSQL can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:dataprotection/backupInstancePostgresql:BackupInstancePostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
    /// ```
    /// </summary>
    [AzureResourceType("azure:dataprotection/backupInstancePostgresql:BackupInstancePostgresql")]
    public partial class BackupInstancePostgresql : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Backup Policy.
        /// </summary>
        [Output("backupPolicyId")]
        public Output<string> BackupPolicyId { get; private set; } = null!;

        /// <summary>
        /// The ID or versionless ID of the key vault secret which stores the connection string of the database.
        /// </summary>
        [Output("databaseCredentialKeyVaultSecretId")]
        public Output<string?> DatabaseCredentialKeyVaultSecretId { get; private set; } = null!;

        /// <summary>
        /// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Backup Instance PostgreSQL. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Backup Vault within which the PostgreSQL Backup Instance should exist. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Output("vaultId")]
        public Output<string> VaultId { get; private set; } = null!;


        /// <summary>
        /// Create a BackupInstancePostgresql resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupInstancePostgresql(string name, BackupInstancePostgresqlArgs args, CustomResourceOptions? options = null)
            : base("azure:dataprotection/backupInstancePostgresql:BackupInstancePostgresql", name, args ?? new BackupInstancePostgresqlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupInstancePostgresql(string name, Input<string> id, BackupInstancePostgresqlState? state = null, CustomResourceOptions? options = null)
            : base("azure:dataprotection/backupInstancePostgresql:BackupInstancePostgresql", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupInstancePostgresql resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupInstancePostgresql Get(string name, Input<string> id, BackupInstancePostgresqlState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupInstancePostgresql(name, id, state, options);
        }
    }

    public sealed class BackupInstancePostgresqlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Backup Policy.
        /// </summary>
        [Input("backupPolicyId", required: true)]
        public Input<string> BackupPolicyId { get; set; } = null!;

        /// <summary>
        /// The ID or versionless ID of the key vault secret which stores the connection string of the database.
        /// </summary>
        [Input("databaseCredentialKeyVaultSecretId")]
        public Input<string>? DatabaseCredentialKeyVaultSecretId { get; set; }

        /// <summary>
        /// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Backup Instance PostgreSQL. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Backup Vault within which the PostgreSQL Backup Instance should exist. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("vaultId", required: true)]
        public Input<string> VaultId { get; set; } = null!;

        public BackupInstancePostgresqlArgs()
        {
        }
    }

    public sealed class BackupInstancePostgresqlState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Backup Policy.
        /// </summary>
        [Input("backupPolicyId")]
        public Input<string>? BackupPolicyId { get; set; }

        /// <summary>
        /// The ID or versionless ID of the key vault secret which stores the connection string of the database.
        /// </summary>
        [Input("databaseCredentialKeyVaultSecretId")]
        public Input<string>? DatabaseCredentialKeyVaultSecretId { get; set; }

        /// <summary>
        /// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Backup Instance PostgreSQL. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Backup Vault within which the PostgreSQL Backup Instance should exist. Changing this forces a new Backup Instance PostgreSQL to be created.
        /// </summary>
        [Input("vaultId")]
        public Input<string>? VaultId { get; set; }

        public BackupInstancePostgresqlState()
        {
        }
    }
}
