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
    /// Manages a Microsoft SQL Azure Database Server.
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
    ///         });
    ///         var exampleServer = new Azure.MSSql.Server("exampleServer", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "thisIsKat11",
    ///             MinimumTlsVersion = "1.2",
    ///             AzureadAdministrator = new Azure.MSSql.Inputs.ServerAzureadAdministratorArgs
    ///             {
    ///                 LoginUsername = "AzureAD Admin",
    ///                 ObjectId = "00000000-0000-0000-0000-000000000000",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Servers can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/server:Server example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/server:Server")]
    public partial class Server : Pulumi.CustomResource
    {
        /// <summary>
        /// The administrator login name for the new server. Changing this forces a new resource to be created.
        /// </summary>
        [Output("administratorLogin")]
        public Output<string> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
        /// </summary>
        [Output("administratorLoginPassword")]
        public Output<string> AdministratorLoginPassword { get; private set; } = null!;

        /// <summary>
        /// An `azuread_administrator` block as defined below.
        /// </summary>
        [Output("azureadAdministrator")]
        public Output<Outputs.ServerAzureadAdministrator?> AzureadAdministrator { get; private set; } = null!;

        /// <summary>
        /// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
        /// </summary>
        [Output("connectionPolicy")]
        public Output<string?> ConnectionPolicy { get; private set; } = null!;

        [Output("extendedAuditingPolicy")]
        public Output<Outputs.ServerExtendedAuditingPolicy> ExtendedAuditingPolicy { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
        /// </summary>
        [Output("fullyQualifiedDomainName")]
        public Output<string> FullyQualifiedDomainName { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServerIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        /// </summary>
        [Output("minimumTlsVersion")]
        public Output<string?> MinimumTlsVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the primary user managed identity id. Required if `type` is `UserAssigned` and should be combined with `user_assigned_identity_ids`.
        /// </summary>
        [Output("primaryUserAssignedIdentityId")]
        public Output<string> PrimaryUserAssignedIdentityId { get; private set; } = null!;

        /// <summary>
        /// Whether public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Microsoft SQL Server.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A list of dropped restorable database IDs on the server.
        /// </summary>
        [Output("restorableDroppedDatabaseIds")]
        public Output<ImmutableArray<string>> RestorableDroppedDatabaseIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/server:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/server:Server", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrator login name for the new server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("administratorLogin", required: true)]
        public Input<string> AdministratorLogin { get; set; } = null!;

        /// <summary>
        /// The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
        /// </summary>
        [Input("administratorLoginPassword", required: true)]
        public Input<string> AdministratorLoginPassword { get; set; } = null!;

        /// <summary>
        /// An `azuread_administrator` block as defined below.
        /// </summary>
        [Input("azureadAdministrator")]
        public Input<Inputs.ServerAzureadAdministratorArgs>? AzureadAdministrator { get; set; }

        /// <summary>
        /// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
        /// </summary>
        [Input("connectionPolicy")]
        public Input<string>? ConnectionPolicy { get; set; }

        [Input("extendedAuditingPolicy")]
        public Input<Inputs.ServerExtendedAuditingPolicyArgs>? ExtendedAuditingPolicy { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServerIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        /// </summary>
        [Input("minimumTlsVersion")]
        public Input<string>? MinimumTlsVersion { get; set; }

        /// <summary>
        /// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the primary user managed identity id. Required if `type` is `UserAssigned` and should be combined with `user_assigned_identity_ids`.
        /// </summary>
        [Input("primaryUserAssignedIdentityId")]
        public Input<string>? PrimaryUserAssignedIdentityId { get; set; }

        /// <summary>
        /// Whether public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Microsoft SQL Server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ServerArgs()
        {
        }
    }

    public sealed class ServerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrator login name for the new server. Changing this forces a new resource to be created.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The password associated with the `administrator_login` user. Needs to comply with Azure's [Password Policy](https://msdn.microsoft.com/library/ms161959.aspx)
        /// </summary>
        [Input("administratorLoginPassword")]
        public Input<string>? AdministratorLoginPassword { get; set; }

        /// <summary>
        /// An `azuread_administrator` block as defined below.
        /// </summary>
        [Input("azureadAdministrator")]
        public Input<Inputs.ServerAzureadAdministratorGetArgs>? AzureadAdministrator { get; set; }

        /// <summary>
        /// The connection policy the server will use. Possible values are `Default`, `Proxy`, and `Redirect`. Defaults to `Default`.
        /// </summary>
        [Input("connectionPolicy")]
        public Input<string>? ConnectionPolicy { get; set; }

        [Input("extendedAuditingPolicy")]
        public Input<Inputs.ServerExtendedAuditingPolicyGetArgs>? ExtendedAuditingPolicy { get; set; }

        /// <summary>
        /// The fully qualified domain name of the Azure SQL Server (e.g. myServerName.database.windows.net)
        /// </summary>
        [Input("fullyQualifiedDomainName")]
        public Input<string>? FullyQualifiedDomainName { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServerIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Minimum TLS Version for all SQL Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        /// </summary>
        [Input("minimumTlsVersion")]
        public Input<string>? MinimumTlsVersion { get; set; }

        /// <summary>
        /// The name of the Microsoft SQL Server. This needs to be globally unique within Azure.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the primary user managed identity id. Required if `type` is `UserAssigned` and should be combined with `user_assigned_identity_ids`.
        /// </summary>
        [Input("primaryUserAssignedIdentityId")]
        public Input<string>? PrimaryUserAssignedIdentityId { get; set; }

        /// <summary>
        /// Whether public network access is allowed for this server. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Microsoft SQL Server.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("restorableDroppedDatabaseIds")]
        private InputList<string>? _restorableDroppedDatabaseIds;

        /// <summary>
        /// A list of dropped restorable database IDs on the server.
        /// </summary>
        public InputList<string> RestorableDroppedDatabaseIds
        {
            get => _restorableDroppedDatabaseIds ?? (_restorableDroppedDatabaseIds = new InputList<string>());
            set => _restorableDroppedDatabaseIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version for the new server. Valid values are: 2.0 (for v11 server) and 12.0 (for v12 server).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerState()
        {
        }
    }
}
