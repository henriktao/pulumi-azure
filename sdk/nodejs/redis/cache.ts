// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Redis Cache.
 *
 * ## Example Usage
 *
 * This example provisions a Standard Redis Cache.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * // NOTE: the Name used for Redis needs to be globally unique
 * const exampleCache = new azure.redis.Cache("exampleCache", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     capacity: 2,
 *     family: "C",
 *     skuName: "Standard",
 *     enableNonSslPort: false,
 *     minimumTlsVersion: "1.2",
 *     redisConfiguration: {},
 * });
 * ```
 * ## Default Redis Configuration Values
 *
 * | Redis Value                     | Basic        | Standard     | Premium      |
 * | ------------------------------- | ------------ | ------------ | ------------ |
 * | enableAuthentication           | true         | true         | true         |
 * | maxmemoryReserved              | 2            | 50           | 200          |
 * | maxfragmentationmemoryReserved | 2            | 50           | 200          |
 * | maxmemoryDelta                 | 2            | 50           | 200          |
 * | maxmemoryPolicy                | volatile-lru | volatile-lru | volatile-lru |
 *
 * > **NOTE:** The `maxmemoryReserved`, `maxmemoryDelta` and `maxfragmentationmemoryReserved` settings are only available for Standard and Premium caches. More details are available in the Relevant Links section below.
 *
 * ***
 *
 * A `patchSchedule` block supports the following:
 *
 * * `dayOfWeek` (Required) the Weekday name - possible values include `Monday`, `Tuesday`, `Wednesday` etc.
 *
 * * `startHourUtc` - (Optional) the Start Hour for maintenance in UTC - possible values range from `0 - 23`.
 *
 * > **Note:** The Patch Window lasts for `5` hours from the `startHourUtc`.
 *
 * * `maintenanceWindow` - (Optional) The ISO 8601 timespan which specifies the amount of time the Redis Cache can be updated. Defaults to `PT5H`.
 *
 * ## Relevant Links
 *
 *  - [Azure Redis Cache: SKU specific configuration limitations](https://azure.microsoft.com/en-us/documentation/articles/cache-configure/#advanced-settings)
 *  - [Redis: Available Configuration Settings](http://redis.io/topics/config)
 *
 * ## Import
 *
 * Redis Cache's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:redis/cache:Cache cache1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1
 * ```
 */
export class Cache extends pulumi.CustomResource {
    /**
     * Get an existing Cache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CacheState, opts?: pulumi.CustomResourceOptions): Cache {
        return new Cache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:redis/cache:Cache';

    /**
     * Returns true if the given object is an instance of Cache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cache.__pulumiType;
    }

    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * Enable the non-SSL port (6379) - disabled by default.
     */
    public readonly enableNonSslPort!: pulumi.Output<boolean | undefined>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * The Hostname of the Redis Instance
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The location of the resource group.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    public readonly minimumTlsVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of `patchSchedule` blocks as defined below.
     */
    public readonly patchSchedules!: pulumi.Output<outputs.redis.CachePatchSchedule[] | undefined>;
    /**
     * The non-SSL Port of the Redis Instance
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The Primary Access Key for the Redis Instance
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * The primary connection string of the Redis Instance.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    public readonly privateStaticIpAddress!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    public readonly redisConfiguration!: pulumi.Output<outputs.redis.CacheRedisConfiguration>;
    /**
     * Redis version. Only major version needed. Valid values: `4`, `6`.
     */
    public readonly redisVersion!: pulumi.Output<string>;
    /**
     * Amount of replicas to create per master for this Redis Cache.
     */
    public readonly replicasPerMaster!: pulumi.Output<number>;
    /**
     * Amount of replicas to create per primary for this Redis Cache. If both `replicasPerPrimary` and `replicasPerMaster` are set, they need to be equal.
     */
    public readonly replicasPerPrimary!: pulumi.Output<number>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Access Key for the Redis Instance
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * The secondary connection string of the Redis Instance.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    public readonly shardCount!: pulumi.Output<number | undefined>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The SSL Port of the Redis Instance
     */
    public /*out*/ readonly sslPort!: pulumi.Output<number>;
    /**
     * *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A mapping of tenant settings to assign to the resource.
     */
    public readonly tenantSettings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of a one or more Availability Zones, where the Redis Cache should be allocated.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Cache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CacheArgs | CacheState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CacheState | undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["enableNonSslPort"] = state ? state.enableNonSslPort : undefined;
            resourceInputs["family"] = state ? state.family : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["minimumTlsVersion"] = state ? state.minimumTlsVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["patchSchedules"] = state ? state.patchSchedules : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            resourceInputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            resourceInputs["privateStaticIpAddress"] = state ? state.privateStaticIpAddress : undefined;
            resourceInputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            resourceInputs["redisConfiguration"] = state ? state.redisConfiguration : undefined;
            resourceInputs["redisVersion"] = state ? state.redisVersion : undefined;
            resourceInputs["replicasPerMaster"] = state ? state.replicasPerMaster : undefined;
            resourceInputs["replicasPerPrimary"] = state ? state.replicasPerPrimary : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            resourceInputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            resourceInputs["shardCount"] = state ? state.shardCount : undefined;
            resourceInputs["skuName"] = state ? state.skuName : undefined;
            resourceInputs["sslPort"] = state ? state.sslPort : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantSettings"] = state ? state.tenantSettings : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as CacheArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["enableNonSslPort"] = args ? args.enableNonSslPort : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["minimumTlsVersion"] = args ? args.minimumTlsVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["patchSchedules"] = args ? args.patchSchedules : undefined;
            resourceInputs["privateStaticIpAddress"] = args ? args.privateStaticIpAddress : undefined;
            resourceInputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            resourceInputs["redisConfiguration"] = args ? args.redisConfiguration : undefined;
            resourceInputs["redisVersion"] = args ? args.redisVersion : undefined;
            resourceInputs["replicasPerMaster"] = args ? args.replicasPerMaster : undefined;
            resourceInputs["replicasPerPrimary"] = args ? args.replicasPerPrimary : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["skuName"] = args ? args.skuName : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantSettings"] = args ? args.tenantSettings : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["primaryAccessKey"] = undefined /*out*/;
            resourceInputs["primaryConnectionString"] = undefined /*out*/;
            resourceInputs["secondaryAccessKey"] = undefined /*out*/;
            resourceInputs["secondaryConnectionString"] = undefined /*out*/;
            resourceInputs["sslPort"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cache.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cache resources.
 */
export interface CacheState {
    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    capacity?: pulumi.Input<number>;
    /**
     * Enable the non-SSL port (6379) - disabled by default.
     */
    enableNonSslPort?: pulumi.Input<boolean>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    family?: pulumi.Input<string>;
    /**
     * The Hostname of the Redis Instance
     */
    hostname?: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    location?: pulumi.Input<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of `patchSchedule` blocks as defined below.
     */
    patchSchedules?: pulumi.Input<pulumi.Input<inputs.redis.CachePatchSchedule>[]>;
    /**
     * The non-SSL Port of the Redis Instance
     */
    port?: pulumi.Input<number>;
    /**
     * The Primary Access Key for the Redis Instance
     */
    primaryAccessKey?: pulumi.Input<string>;
    /**
     * The primary connection string of the Redis Instance.
     */
    primaryConnectionString?: pulumi.Input<string>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    privateStaticIpAddress?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    redisConfiguration?: pulumi.Input<inputs.redis.CacheRedisConfiguration>;
    /**
     * Redis version. Only major version needed. Valid values: `4`, `6`.
     */
    redisVersion?: pulumi.Input<string>;
    /**
     * Amount of replicas to create per master for this Redis Cache.
     */
    replicasPerMaster?: pulumi.Input<number>;
    /**
     * Amount of replicas to create per primary for this Redis Cache. If both `replicasPerPrimary` and `replicasPerMaster` are set, they need to be equal.
     */
    replicasPerPrimary?: pulumi.Input<number>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Access Key for the Redis Instance
     */
    secondaryAccessKey?: pulumi.Input<string>;
    /**
     * The secondary connection string of the Redis Instance.
     */
    secondaryConnectionString?: pulumi.Input<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    skuName?: pulumi.Input<string>;
    /**
     * The SSL Port of the Redis Instance
     */
    sslPort?: pulumi.Input<number>;
    /**
     * *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A mapping of tenant settings to assign to the resource.
     */
    tenantSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of a one or more Availability Zones, where the Redis Cache should be allocated.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * The size of the Redis cache to deploy. Valid values for a SKU `family` of C (Basic/Standard) are `0, 1, 2, 3, 4, 5, 6`, and for P (Premium) `family` are `1, 2, 3, 4`.
     */
    capacity: pulumi.Input<number>;
    /**
     * Enable the non-SSL port (6379) - disabled by default.
     */
    enableNonSslPort?: pulumi.Input<boolean>;
    /**
     * The SKU family/pricing group to use. Valid values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
     */
    family: pulumi.Input<string>;
    /**
     * The location of the resource group.
     */
    location?: pulumi.Input<string>;
    /**
     * The minimum TLS version.  Defaults to `1.0`.
     */
    minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis instance. Changing this forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of `patchSchedule` blocks as defined below.
     */
    patchSchedules?: pulumi.Input<pulumi.Input<inputs.redis.CachePatchSchedule>[]>;
    /**
     * The Static IP Address to assign to the Redis Cache when hosted inside the Virtual Network. Changing this forces a new resource to be created.
     */
    privateStaticIpAddress?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this Redis Cache. `true` means this resource could be accessed by both public and private endpoint. `false` means only private endpoint access is allowed. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * A `redisConfiguration` as defined below - with some limitations by SKU - defaults/details are shown below.
     */
    redisConfiguration?: pulumi.Input<inputs.redis.CacheRedisConfiguration>;
    /**
     * Redis version. Only major version needed. Valid values: `4`, `6`.
     */
    redisVersion?: pulumi.Input<string>;
    /**
     * Amount of replicas to create per master for this Redis Cache.
     */
    replicasPerMaster?: pulumi.Input<number>;
    /**
     * Amount of replicas to create per primary for this Redis Cache. If both `replicasPerPrimary` and `replicasPerMaster` are set, they need to be equal.
     */
    replicasPerPrimary?: pulumi.Input<number>;
    /**
     * The name of the resource group in which to
     * create the Redis instance.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * *Only available when using the Premium SKU* The number of Shards to create on the Redis Cluster.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * The SKU of Redis to use. Possible values are `Basic`, `Standard` and `Premium`.
     */
    skuName: pulumi.Input<string>;
    /**
     * *Only available when using the Premium SKU* The ID of the Subnet within which the Redis Cache should be deployed. This Subnet must only contain Azure Cache for Redis instances without any other type of resources. Changing this forces a new resource to be created.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A mapping of tenant settings to assign to the resource.
     */
    tenantSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of a one or more Availability Zones, where the Redis Cache should be allocated.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
