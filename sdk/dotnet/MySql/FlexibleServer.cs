// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql
{
    /// <summary>
    /// Manages a MySQL Flexible Server.
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
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///             ServiceEndpoints = 
    ///             {
    ///                 "Microsoft.Storage",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "fs",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.DBforMySQL/flexibleServers",
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/virtualNetworks/subnets/join/action",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleZone = new Azure.PrivateDns.Zone("exampleZone", new Azure.PrivateDns.ZoneArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleZoneVirtualNetworkLink = new Azure.PrivateDns.ZoneVirtualNetworkLink("exampleZoneVirtualNetworkLink", new Azure.PrivateDns.ZoneVirtualNetworkLinkArgs
    ///         {
    ///             PrivateDnsZoneName = exampleZone.Name,
    ///             VirtualNetworkId = exampleVirtualNetwork.Id,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleFlexibleServer = new Azure.MySql.FlexibleServer("exampleFlexibleServer", new Azure.MySql.FlexibleServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AdministratorLogin = "psqladmin",
    ///             AdministratorPassword = "H@Sh1CoR3!",
    ///             BackupRetentionDays = 7,
    ///             DelegatedSubnetId = exampleSubnet.Id,
    ///             PrivateDnsZoneId = exampleZone.Id,
    ///             SkuName = "GP_Standard_D2ds_v4",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleZoneVirtualNetworkLink,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// MySQL Flexible Servers can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:mysql/flexibleServer:FlexibleServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1
    /// ```
    /// </summary>
    [AzureResourceType("azure:mysql/flexibleServer:FlexibleServer")]
    public partial class FlexibleServer : Pulumi.CustomResource
    {
        /// <summary>
        /// The Administrator Login for the MySQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("administratorLogin")]
        public Output<string> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The Password associated with the `administrator_login` for the MySQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Output("administratorPassword")]
        public Output<string?> AdministratorPassword { get; private set; } = null!;

        /// <summary>
        /// The backup retention days for the MySQL Flexible Server. Possible values are between `7` and `35` days. Defaults to `7`.
        /// </summary>
        [Output("backupRetentionDays")]
        public Output<int?> BackupRetentionDays { get; private set; } = null!;

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default`, `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual network subnet to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("delegatedSubnetId")]
        public Output<string?> DelegatedSubnetId { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name of the MySQL Flexible Server.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Should geo redundant backup enabled? Defaults to `false`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("geoRedundantBackupEnabled")]
        public Output<bool?> GeoRedundantBackupEnabled { get; private set; } = null!;

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Output("highAvailability")]
        public Output<Outputs.FlexibleServerHighAvailability?> HighAvailability { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.FlexibleServerMaintenanceWindow?> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("pointInTimeRestoreTimeInUtc")]
        public Output<string?> PointInTimeRestoreTimeInUtc { get; private set; } = null!;

        /// <summary>
        /// The ID of the private dns zone to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("privateDnsZoneId")]
        public Output<string?> PrivateDnsZoneId { get; private set; } = null!;

        /// <summary>
        /// Is the public network access enabled?
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The maximum number of replicas that a primary MySQL Flexible Server can have.
        /// </summary>
        [Output("replicaCapacity")]
        public Output<int> ReplicaCapacity { get; private set; } = null!;

        /// <summary>
        /// The replication role. Possible value is `None`.
        /// </summary>
        [Output("replicationRole")]
        public Output<string> ReplicationRole { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The SKU Name for the MySQL Flexible Server.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the source MySQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("sourceServerId")]
        public Output<string?> SourceServerId { get; private set; } = null!;

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        [Output("storage")]
        public Output<Outputs.FlexibleServerStorage> Storage { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the MySQL Flexible Server.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The version of the MySQL Flexible Server to use. Possible values are `5.7`, and `8.0.21`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The availability zone information of the MySQL Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a FlexibleServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlexibleServer(string name, FlexibleServerArgs args, CustomResourceOptions? options = null)
            : base("azure:mysql/flexibleServer:FlexibleServer", name, args ?? new FlexibleServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlexibleServer(string name, Input<string> id, FlexibleServerState? state = null, CustomResourceOptions? options = null)
            : base("azure:mysql/flexibleServer:FlexibleServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FlexibleServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlexibleServer Get(string name, Input<string> id, FlexibleServerState? state = null, CustomResourceOptions? options = null)
        {
            return new FlexibleServer(name, id, state, options);
        }
    }

    public sealed class FlexibleServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Administrator Login for the MySQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The Password associated with the `administrator_login` for the MySQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Input("administratorPassword")]
        public Input<string>? AdministratorPassword { get; set; }

        /// <summary>
        /// The backup retention days for the MySQL Flexible Server. Possible values are between `7` and `35` days. Defaults to `7`.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default`, `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The ID of the virtual network subnet to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("delegatedSubnetId")]
        public Input<string>? DelegatedSubnetId { get; set; }

        /// <summary>
        /// Should geo redundant backup enabled? Defaults to `false`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("geoRedundantBackupEnabled")]
        public Input<bool>? GeoRedundantBackupEnabled { get; set; }

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Input("highAvailability")]
        public Input<Inputs.FlexibleServerHighAvailabilityArgs>? HighAvailability { get; set; }

        /// <summary>
        /// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.FlexibleServerMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name which should be used for this MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("pointInTimeRestoreTimeInUtc")]
        public Input<string>? PointInTimeRestoreTimeInUtc { get; set; }

        /// <summary>
        /// The ID of the private dns zone to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("privateDnsZoneId")]
        public Input<string>? PrivateDnsZoneId { get; set; }

        /// <summary>
        /// The replication role. Possible value is `None`.
        /// </summary>
        [Input("replicationRole")]
        public Input<string>? ReplicationRole { get; set; }

        /// <summary>
        /// The name of the Resource Group where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU Name for the MySQL Flexible Server.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The resource ID of the source MySQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("sourceServerId")]
        public Input<string>? SourceServerId { get; set; }

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        [Input("storage")]
        public Input<Inputs.FlexibleServerStorageArgs>? Storage { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the MySQL Flexible Server.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version of the MySQL Flexible Server to use. Possible values are `5.7`, and `8.0.21`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The availability zone information of the MySQL Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public FlexibleServerArgs()
        {
        }
    }

    public sealed class FlexibleServerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Administrator Login for the MySQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The Password associated with the `administrator_login` for the MySQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Input("administratorPassword")]
        public Input<string>? AdministratorPassword { get; set; }

        /// <summary>
        /// The backup retention days for the MySQL Flexible Server. Possible values are between `7` and `35` days. Defaults to `7`.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default`, `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The ID of the virtual network subnet to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("delegatedSubnetId")]
        public Input<string>? DelegatedSubnetId { get; set; }

        /// <summary>
        /// The fully qualified domain name of the MySQL Flexible Server.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Should geo redundant backup enabled? Defaults to `false`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("geoRedundantBackupEnabled")]
        public Input<bool>? GeoRedundantBackupEnabled { get; set; }

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Input("highAvailability")]
        public Input<Inputs.FlexibleServerHighAvailabilityGetArgs>? HighAvailability { get; set; }

        /// <summary>
        /// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.FlexibleServerMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name which should be used for this MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("pointInTimeRestoreTimeInUtc")]
        public Input<string>? PointInTimeRestoreTimeInUtc { get; set; }

        /// <summary>
        /// The ID of the private dns zone to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("privateDnsZoneId")]
        public Input<string>? PrivateDnsZoneId { get; set; }

        /// <summary>
        /// Is the public network access enabled?
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The maximum number of replicas that a primary MySQL Flexible Server can have.
        /// </summary>
        [Input("replicaCapacity")]
        public Input<int>? ReplicaCapacity { get; set; }

        /// <summary>
        /// The replication role. Possible value is `None`.
        /// </summary>
        [Input("replicationRole")]
        public Input<string>? ReplicationRole { get; set; }

        /// <summary>
        /// The name of the Resource Group where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The SKU Name for the MySQL Flexible Server.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The resource ID of the source MySQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`, `GeoRestore`, and `Replica`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("sourceServerId")]
        public Input<string>? SourceServerId { get; set; }

        /// <summary>
        /// A `storage` block as defined below.
        /// </summary>
        [Input("storage")]
        public Input<Inputs.FlexibleServerStorageGetArgs>? Storage { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the MySQL Flexible Server.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version of the MySQL Flexible Server to use. Possible values are `5.7`, and `8.0.21`. Changing this forces a new MySQL Flexible Server to be created.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The availability zone information of the MySQL Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public FlexibleServerState()
        {
        }
    }
}
