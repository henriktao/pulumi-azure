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
    /// Manages a Microsoft Azure SQL Failover Group.
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
    ///         var primary = new Azure.MSSql.Server("primary", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "thisIsKat11",
    ///         });
    ///         var secondary = new Azure.MSSql.Server("secondary", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "thisIsKat12",
    ///         });
    ///         var exampleDatabase = new Azure.MSSql.Database("exampleDatabase", new Azure.MSSql.DatabaseArgs
    ///         {
    ///             ServerId = primary.Id,
    ///             SkuName = "S1",
    ///             Collation = "SQL_Latin1_General_CP1_CI_AS",
    ///             MaxSizeGb = 200,
    ///         });
    ///         var exampleFailoverGroup = new Azure.MSSql.FailoverGroup("exampleFailoverGroup", new Azure.MSSql.FailoverGroupArgs
    ///         {
    ///             ServerId = primary.Id,
    ///             Databases = 
    ///             {
    ///                 exampleDatabase.Id,
    ///             },
    ///             PartnerServers = 
    ///             {
    ///                 new Azure.MSSql.Inputs.FailoverGroupPartnerServerArgs
    ///                 {
    ///                     Id = secondary.Id,
    ///                 },
    ///             },
    ///             ReadWriteEndpointFailoverPolicy = new Azure.MSSql.Inputs.FailoverGroupReadWriteEndpointFailoverPolicyArgs
    ///             {
    ///                 Mode = "Automatic",
    ///                 GraceMinutes = 80,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "prod" },
    ///                 { "database", "example" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Failover Groups can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mssql/failoverGroup:FailoverGroup example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/servers/server1/failoverGroups/failoverGroup1
    /// ```
    /// </summary>
    [AzureResourceType("azure:mssql/failoverGroup:FailoverGroup")]
    public partial class FailoverGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A set of database names to include in the failover group.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<string>> Databases { get; private set; } = null!;

        /// <summary>
        /// The name of the Failover Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `partner_server` block as defined below.
        /// </summary>
        [Output("partnerServers")]
        public Output<ImmutableArray<Outputs.FailoverGroupPartnerServer>> PartnerServers { get; private set; } = null!;

        /// <summary>
        /// A `read_write_endpoint_failover_policy` block as defined below.
        /// </summary>
        [Output("readWriteEndpointFailoverPolicy")]
        public Output<Outputs.FailoverGroupReadWriteEndpointFailoverPolicy> ReadWriteEndpointFailoverPolicy { get; private set; } = null!;

        /// <summary>
        /// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
        /// </summary>
        [Output("readonlyEndpointFailoverPolicyEnabled")]
        public Output<bool> ReadonlyEndpointFailoverPolicyEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a FailoverGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FailoverGroup(string name, FailoverGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/failoverGroup:FailoverGroup", name, args ?? new FailoverGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FailoverGroup(string name, Input<string> id, FailoverGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/failoverGroup:FailoverGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FailoverGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FailoverGroup Get(string name, Input<string> id, FailoverGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new FailoverGroup(name, id, state, options);
        }
    }

    public sealed class FailoverGroupArgs : Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;

        /// <summary>
        /// A set of database names to include in the failover group.
        /// </summary>
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        /// <summary>
        /// The name of the Failover Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("partnerServers", required: true)]
        private InputList<Inputs.FailoverGroupPartnerServerArgs>? _partnerServers;

        /// <summary>
        /// A `partner_server` block as defined below.
        /// </summary>
        public InputList<Inputs.FailoverGroupPartnerServerArgs> PartnerServers
        {
            get => _partnerServers ?? (_partnerServers = new InputList<Inputs.FailoverGroupPartnerServerArgs>());
            set => _partnerServers = value;
        }

        /// <summary>
        /// A `read_write_endpoint_failover_policy` block as defined below.
        /// </summary>
        [Input("readWriteEndpointFailoverPolicy", required: true)]
        public Input<Inputs.FailoverGroupReadWriteEndpointFailoverPolicyArgs> ReadWriteEndpointFailoverPolicy { get; set; } = null!;

        /// <summary>
        /// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
        /// </summary>
        [Input("readonlyEndpointFailoverPolicyEnabled")]
        public Input<bool>? ReadonlyEndpointFailoverPolicyEnabled { get; set; }

        /// <summary>
        /// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

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

        public FailoverGroupArgs()
        {
        }
    }

    public sealed class FailoverGroupState : Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;

        /// <summary>
        /// A set of database names to include in the failover group.
        /// </summary>
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        /// <summary>
        /// The name of the Failover Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("partnerServers")]
        private InputList<Inputs.FailoverGroupPartnerServerGetArgs>? _partnerServers;

        /// <summary>
        /// A `partner_server` block as defined below.
        /// </summary>
        public InputList<Inputs.FailoverGroupPartnerServerGetArgs> PartnerServers
        {
            get => _partnerServers ?? (_partnerServers = new InputList<Inputs.FailoverGroupPartnerServerGetArgs>());
            set => _partnerServers = value;
        }

        /// <summary>
        /// A `read_write_endpoint_failover_policy` block as defined below.
        /// </summary>
        [Input("readWriteEndpointFailoverPolicy")]
        public Input<Inputs.FailoverGroupReadWriteEndpointFailoverPolicyGetArgs>? ReadWriteEndpointFailoverPolicy { get; set; }

        /// <summary>
        /// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
        /// </summary>
        [Input("readonlyEndpointFailoverPolicyEnabled")]
        public Input<bool>? ReadonlyEndpointFailoverPolicyEnabled { get; set; }

        /// <summary>
        /// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

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

        public FailoverGroupState()
        {
        }
    }
}
