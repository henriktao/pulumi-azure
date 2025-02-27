// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PostgreSql
{
    /// <summary>
    /// Manages a PostgreSQL Flexible Server.
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
    ///                         Name = "Microsoft.DBforPostgreSQL/flexibleServers",
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
    ///         var exampleFlexibleServer = new Azure.PostgreSql.FlexibleServer("exampleFlexibleServer", new Azure.PostgreSql.FlexibleServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12",
    ///             DelegatedSubnetId = exampleSubnet.Id,
    ///             PrivateDnsZoneId = exampleZone.Id,
    ///             AdministratorLogin = "psqladmin",
    ///             AdministratorPassword = "H@Sh1CoR3!",
    ///             Zone = "1",
    ///             StorageMb = 32768,
    ///             SkuName = "GP_Standard_D4s_v3",
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
    /// PostgreSQL Flexible Servers can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:postgresql/flexibleServer:FlexibleServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1
    /// ```
    /// </summary>
    [AzureResourceType("azure:postgresql/flexibleServer:FlexibleServer")]
    public partial class FlexibleServer : Pulumi.CustomResource
    {
        /// <summary>
        /// The Administrator Login for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("administratorLogin")]
        public Output<string> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// The Password associated with the `administrator_login` for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Output("administratorPassword")]
        public Output<string?> AdministratorPassword { get; private set; } = null!;

        /// <summary>
        /// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
        /// </summary>
        [Output("backupRetentionDays")]
        public Output<int> BackupRetentionDays { get; private set; } = null!;

        /// <summary>
        /// The status showing whether the data encryption is enabled with a customer-managed key.
        /// </summary>
        [Output("cmkEnabled")]
        public Output<string> CmkEnabled { get; private set; } = null!;

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("delegatedSubnetId")]
        public Output<string?> DelegatedSubnetId { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the PostgreSQL Flexible Server.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Is Geo-Redundant backup enabled on the PostgreSQL Flexible Server. Defaults to `false`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("geoRedundantBackupEnabled")]
        public Output<bool?> GeoRedundantBackupEnabled { get; private set; } = null!;

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Output("highAvailability")]
        public Output<Outputs.FlexibleServerHighAvailability?> HighAvailability { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.FlexibleServerMaintenanceWindow?> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("pointInTimeRestoreTimeInUtc")]
        public Output<string?> PointInTimeRestoreTimeInUtc { get; private set; } = null!;

        /// <summary>
        /// The ID of the private dns zone to create the PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("privateDnsZoneId")]
        public Output<string> PrivateDnsZoneId { get; private set; } = null!;

        /// <summary>
        /// Is public network access enabled?
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("sourceServerId")]
        public Output<string?> SourceServerId { get; private set; } = null!;

        /// <summary>
        /// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
        /// </summary>
        [Output("storageMb")]
        public Output<int> StorageMb { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
        /// *
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The version of PostgreSQL Flexible Server to use. Possible values are `11`,`12` and `13`. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone of the PostgreSQL Flexible Server. Possible values are `1`, `2` and `3`.
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
            : base("azure:postgresql/flexibleServer:FlexibleServer", name, args ?? new FlexibleServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlexibleServer(string name, Input<string> id, FlexibleServerState? state = null, CustomResourceOptions? options = null)
            : base("azure:postgresql/flexibleServer:FlexibleServer", name, state, MakeResourceOptions(options, id))
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
        /// The Administrator Login for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The Password associated with the `administrator_login` for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Input("administratorPassword")]
        public Input<string>? AdministratorPassword { get; set; }

        /// <summary>
        /// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("delegatedSubnetId")]
        public Input<string>? DelegatedSubnetId { get; set; }

        /// <summary>
        /// Is Geo-Redundant backup enabled on the PostgreSQL Flexible Server. Defaults to `false`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("geoRedundantBackupEnabled")]
        public Input<bool>? GeoRedundantBackupEnabled { get; set; }

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Input("highAvailability")]
        public Input<Inputs.FlexibleServerHighAvailabilityArgs>? HighAvailability { get; set; }

        /// <summary>
        /// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.FlexibleServerMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("pointInTimeRestoreTimeInUtc")]
        public Input<string>? PointInTimeRestoreTimeInUtc { get; set; }

        /// <summary>
        /// The ID of the private dns zone to create the PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("privateDnsZoneId")]
        public Input<string>? PrivateDnsZoneId { get; set; }

        /// <summary>
        /// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("sourceServerId")]
        public Input<string>? SourceServerId { get; set; }

        /// <summary>
        /// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
        /// </summary>
        [Input("storageMb")]
        public Input<int>? StorageMb { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
        /// *
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version of PostgreSQL Flexible Server to use. Possible values are `11`,`12` and `13`. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The Availability Zone of the PostgreSQL Flexible Server. Possible values are `1`, `2` and `3`.
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
        /// The Administrator Login for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("administratorLogin")]
        public Input<string>? AdministratorLogin { get; set; }

        /// <summary>
        /// The Password associated with the `administrator_login` for the PostgreSQL Flexible Server. Required when `create_mode` is `Default`.
        /// </summary>
        [Input("administratorPassword")]
        public Input<string>? AdministratorPassword { get; set; }

        /// <summary>
        /// The backup retention days for the PostgreSQL Flexible Server. Possible values are between `7` and `35` days.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// The status showing whether the data encryption is enabled with a customer-managed key.
        /// </summary>
        [Input("cmkEnabled")]
        public Input<string>? CmkEnabled { get; set; }

        /// <summary>
        /// The creation mode which can be used to restore or replicate existing servers. Possible values are `Default` and `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The ID of the virtual network subnet to create the PostgreSQL Flexible Server. The provided subnet should not have any other resource deployed in it and this subnet will be delegated to the PostgreSQL Flexible Server, if not already delegated. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("delegatedSubnetId")]
        public Input<string>? DelegatedSubnetId { get; set; }

        /// <summary>
        /// The FQDN of the PostgreSQL Flexible Server.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Is Geo-Redundant backup enabled on the PostgreSQL Flexible Server. Defaults to `false`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("geoRedundantBackupEnabled")]
        public Input<bool>? GeoRedundantBackupEnabled { get; set; }

        /// <summary>
        /// A `high_availability` block as defined below.
        /// </summary>
        [Input("highAvailability")]
        public Input<Inputs.FlexibleServerHighAvailabilityGetArgs>? HighAvailability { get; set; }

        /// <summary>
        /// The Azure Region where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `maintenance_window` block as defined below.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.FlexibleServerMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The name which should be used for this PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The point in time to restore from `creation_source_server_id` when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("pointInTimeRestoreTimeInUtc")]
        public Input<string>? PointInTimeRestoreTimeInUtc { get; set; }

        /// <summary>
        /// The ID of the private dns zone to create the PostgreSQL Flexible Server. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("privateDnsZoneId")]
        public Input<string>? PrivateDnsZoneId { get; set; }

        /// <summary>
        /// Is public network access enabled?
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the Resource Group where the PostgreSQL Flexible Server should exist. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The SKU Name for the PostgreSQL Flexible Server. The name of the SKU, follows the `tier` + `name` pattern (e.g. `B_Standard_B1ms`, `GP_Standard_D2s_v3`, `MO_Standard_E4s_v3`).
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The resource ID of the source PostgreSQL Flexible Server to be restored. Required when `create_mode` is `PointInTimeRestore`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("sourceServerId")]
        public Input<string>? SourceServerId { get; set; }

        /// <summary>
        /// The max storage allowed for the PostgreSQL Flexible Server. Possible values are `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`, and `33554432`.
        /// </summary>
        [Input("storageMb")]
        public Input<int>? StorageMb { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the PostgreSQL Flexible Server.
        /// *
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version of PostgreSQL Flexible Server to use. Possible values are `11`,`12` and `13`. Required when `create_mode` is `Default`. Changing this forces a new PostgreSQL Flexible Server to be created.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The Availability Zone of the PostgreSQL Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public FlexibleServerState()
        {
        }
    }
}
