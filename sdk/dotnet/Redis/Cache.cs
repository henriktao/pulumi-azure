// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Redis
{
    /// <summary>
    /// Manages a Redis Cache.
    /// 
    /// ## Example Usage
    /// 
    /// This example provisions a Standard Redis Cache.
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
    ///         // NOTE: the Name used for Redis needs to be globally unique
    ///         var exampleCache = new Azure.Redis.Cache("exampleCache", new Azure.Redis.CacheArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Capacity = 2,
    ///             Family = "C",
    ///             SkuName = "Standard",
    ///             EnableNonSslPort = false,
    ///             MinimumTlsVersion = "1.2",
    ///             RedisConfiguration = ,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Default Redis Configuration Values
    /// 
    /// | Redis Value                     | Basic        | Standard     | Premium      |
    /// | ------------------------------- | ------------ | ------------ | ------------ |
    /// | enable_authentication           | true         | true         | true         |
    /// | maxmemory_reserved              | 2            | 50           | 200          |
    /// | maxfragmentationmemory_reserved | 2            | 50           | 200          |
    /// | maxmemory_delta                 | 2            | 50           | 200          |
    /// | maxmemory_policy                | volatile-lru | volatile-lru | volatile-lru |
    /// 
    /// &gt; **NOTE:** The `maxmemory_reserved`, `maxmemory_delta` and `maxfragmentationmemory_reserved` settings are only available for Standard and Premium caches. More details are available in the Relevant Links section below.
    /// 
    /// ***
    /// 
    /// A `patch_schedule` block supports the following:
    /// 
    /// * `day_of_week` (Required) the Weekday name - possible values include `Monday`, `Tuesday`, `Wednesday` etc.
    /// 
    /// * `start_hour_utc` - (Optional) the Start Hour for maintenance in UTC - possible values range from `0 - 23`.
    /// 
    /// &gt; **Note:** The Patch Window lasts for `5` hours from the `start_hour_utc`.
    /// 
    /// * `maintenance_window` - (Optional) The ISO 8601 timespan which specifies the amount of time the Redis Cache can be updated. Defaults to `PT5H`.
    /// 
    /// ## Relevant Links
    /// 
    ///  - [Azure Redis Cache: SKU specific configuration limitations](https://azure.microsoft.com/en-us/documentation/articles/cache-configure/#advanced-settings)
    ///  - [Redis: Available Configuration Settings](http://redis.io/topics/config)
    /// 
    /// ## Import
    /// 
    /// Redis Cache's can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:redis/cache:Cache cache1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1
    /// ```
    /// </summary>
    [AzureResourceType("azure:redis/cache:Cache")]
    public partial class Cache : Pulumi.CustomResource
    {
        /// <summary>
        /// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// Enable the non-SSL port (6379) - disabled by default.
        /// </summary>
        [Output("enableNonSslPort")]
        public Output<bool?> EnableNonSslPort { get; private set; } = null!;

        /// <summary>
        /// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// The Hostname of the Redis Instance
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The location of the resource group.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The minimum TLS version.  Defaults to `1.0`.
        /// </summary>
        [Output("minimumTlsVersion")]
        public Output<string?> MinimumTlsVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the Redis instance. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of `patch_schedule` blocks as defined below.
        /// </summary>
        [Output("patchSchedules")]
        public Output<ImmutableArray<Outputs.CachePatchSchedule>> PatchSchedules { get; private set; } = null!;

        /// <summary>
        /// The non-SSL Port of the Redis Instance
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The Primary Access Key for the Redis Instance
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The primary connection string of the Redis Instance.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Output("privateStaticIpAddress")]
        public Output<string> PrivateStaticIpAddress { get; private set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
        /// </summary>
        [Output("redisConfiguration")]
        public Output<Outputs.CacheRedisConfiguration> RedisConfiguration { get; private set; } = null!;

        /// <summary>
        /// Redis version. Only major version needed. Valid values: `4`, `6`.
        /// </summary>
        [Output("redisVersion")]
        public Output<string> RedisVersion { get; private set; } = null!;

        /// <summary>
        /// Amount of replicas to create per master for this Redis Cache.
        /// </summary>
        [Output("replicasPerMaster")]
        public Output<int> ReplicasPerMaster { get; private set; } = null!;

        /// <summary>
        /// Amount of replicas to create per primary for this Redis Cache. If both `replicas_per_primary` and `replicas_per_master` are set, they need to be equal.
        /// </summary>
        [Output("replicasPerPrimary")]
        public Output<int> ReplicasPerPrimary { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the Redis instance.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Secondary Access Key for the Redis Instance
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string of the Redis Instance.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
        /// </summary>
        [Output("shardCount")]
        public Output<int?> ShardCount { get; private set; } = null!;

        /// <summary>
        /// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// The SSL Port of the Redis Instance
        /// </summary>
        [Output("sslPort")]
        public Output<int> SslPort { get; private set; } = null!;

        /// <summary>
        /// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A mapping of tenant settings to assign to the resource.
        /// </summary>
        [Output("tenantSettings")]
        public Output<ImmutableDictionary<string, string>?> TenantSettings { get; private set; } = null!;

        /// <summary>
        /// A list of a one or more Availability Zones, where the Redis Cache should be allocated.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Cache resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cache(string name, CacheArgs args, CustomResourceOptions? options = null)
            : base("azure:redis/cache:Cache", name, args ?? new CacheArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cache(string name, Input<string> id, CacheState? state = null, CustomResourceOptions? options = null)
            : base("azure:redis/cache:Cache", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cache resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cache Get(string name, Input<string> id, CacheState? state = null, CustomResourceOptions? options = null)
        {
            return new Cache(name, id, state, options);
        }
    }

    public sealed class CacheArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
        /// </summary>
        [Input("capacity", required: true)]
        public Input<int> Capacity { get; set; } = null!;

        /// <summary>
        /// Enable the non-SSL port (6379) - disabled by default.
        /// </summary>
        [Input("enableNonSslPort")]
        public Input<bool>? EnableNonSslPort { get; set; }

        /// <summary>
        /// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
        /// </summary>
        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        /// <summary>
        /// The location of the resource group.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The minimum TLS version.  Defaults to `1.0`.
        /// </summary>
        [Input("minimumTlsVersion")]
        public Input<string>? MinimumTlsVersion { get; set; }

        /// <summary>
        /// The name of the Redis instance. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("patchSchedules")]
        private InputList<Inputs.CachePatchScheduleArgs>? _patchSchedules;

        /// <summary>
        /// A list of `patch_schedule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CachePatchScheduleArgs> PatchSchedules
        {
            get => _patchSchedules ?? (_patchSchedules = new InputList<Inputs.CachePatchScheduleArgs>());
            set => _patchSchedules = value;
        }

        /// <summary>
        /// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("privateStaticIpAddress")]
        public Input<string>? PrivateStaticIpAddress { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
        /// </summary>
        [Input("redisConfiguration")]
        public Input<Inputs.CacheRedisConfigurationArgs>? RedisConfiguration { get; set; }

        /// <summary>
        /// Redis version. Only major version needed. Valid values: `4`, `6`.
        /// </summary>
        [Input("redisVersion")]
        public Input<string>? RedisVersion { get; set; }

        /// <summary>
        /// Amount of replicas to create per master for this Redis Cache.
        /// </summary>
        [Input("replicasPerMaster")]
        public Input<int>? ReplicasPerMaster { get; set; }

        /// <summary>
        /// Amount of replicas to create per primary for this Redis Cache. If both `replicas_per_primary` and `replicas_per_master` are set, they need to be equal.
        /// </summary>
        [Input("replicasPerPrimary")]
        public Input<int>? ReplicasPerPrimary { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Redis instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        /// <summary>
        /// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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

        [Input("tenantSettings")]
        private InputMap<string>? _tenantSettings;

        /// <summary>
        /// A mapping of tenant settings to assign to the resource.
        /// </summary>
        public InputMap<string> TenantSettings
        {
            get => _tenantSettings ?? (_tenantSettings = new InputMap<string>());
            set => _tenantSettings = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of a one or more Availability Zones, where the Redis Cache should be allocated.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public CacheArgs()
        {
        }
    }

    public sealed class CacheState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Enable the non-SSL port (6379) - disabled by default.
        /// </summary>
        [Input("enableNonSslPort")]
        public Input<bool>? EnableNonSslPort { get; set; }

        /// <summary>
        /// The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The Hostname of the Redis Instance
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The location of the resource group.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The minimum TLS version.  Defaults to `1.0`.
        /// </summary>
        [Input("minimumTlsVersion")]
        public Input<string>? MinimumTlsVersion { get; set; }

        /// <summary>
        /// The name of the Redis instance. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("patchSchedules")]
        private InputList<Inputs.CachePatchScheduleGetArgs>? _patchSchedules;

        /// <summary>
        /// A list of `patch_schedule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CachePatchScheduleGetArgs> PatchSchedules
        {
            get => _patchSchedules ?? (_patchSchedules = new InputList<Inputs.CachePatchScheduleGetArgs>());
            set => _patchSchedules = value;
        }

        /// <summary>
        /// The non-SSL Port of the Redis Instance
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The Primary Access Key for the Redis Instance
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The primary connection string of the Redis Instance.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("privateStaticIpAddress")]
        public Input<string>? PrivateStaticIpAddress { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// A `redis_configuration` as defined below - with some limitations by SKU - defaults/details are shown below.
        /// </summary>
        [Input("redisConfiguration")]
        public Input<Inputs.CacheRedisConfigurationGetArgs>? RedisConfiguration { get; set; }

        /// <summary>
        /// Redis version. Only major version needed. Valid values: `4`, `6`.
        /// </summary>
        [Input("redisVersion")]
        public Input<string>? RedisVersion { get; set; }

        /// <summary>
        /// Amount of replicas to create per master for this Redis Cache.
        /// </summary>
        [Input("replicasPerMaster")]
        public Input<int>? ReplicasPerMaster { get; set; }

        /// <summary>
        /// Amount of replicas to create per primary for this Redis Cache. If both `replicas_per_primary` and `replicas_per_master` are set, they need to be equal.
        /// </summary>
        [Input("replicasPerPrimary")]
        public Input<int>? ReplicasPerPrimary { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the Redis instance.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Secondary Access Key for the Redis Instance
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        /// <summary>
        /// The secondary connection string of the Redis Instance.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The SSL Port of the Redis Instance
        /// </summary>
        [Input("sslPort")]
        public Input<int>? SslPort { get; set; }

        /// <summary>
        /// *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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

        [Input("tenantSettings")]
        private InputMap<string>? _tenantSettings;

        /// <summary>
        /// A mapping of tenant settings to assign to the resource.
        /// </summary>
        public InputMap<string> TenantSettings
        {
            get => _tenantSettings ?? (_tenantSettings = new InputMap<string>());
            set => _tenantSettings = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of a one or more Availability Zones, where the Redis Cache should be allocated.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public CacheState()
        {
        }
    }
}
