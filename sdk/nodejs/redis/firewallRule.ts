// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Firewall Rule associated with a Redis Cache.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 *
 * const server = new random.RandomId("server", {
 *     keepers: {
 *         azi_id: 1,
 *     },
 *     byteLength: 8,
 * });
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleCache = new azure.redis.Cache("exampleCache", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     capacity: 1,
 *     family: "P",
 *     skuName: "Premium",
 *     enableNonSslPort: false,
 *     redisConfiguration: {
 *         maxclients: 256,
 *         maxmemoryReserved: 2,
 *         maxmemoryDelta: 2,
 *         maxmemoryPolicy: "allkeys-lru",
 *     },
 * });
 * const exampleFirewallRule = new azure.redis.FirewallRule("exampleFirewallRule", {
 *     redisCacheName: exampleCache.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     startIp: "1.2.3.4",
 *     endIp: "2.3.4.5",
 * });
 * ```
 *
 * ## Import
 *
 * Redis Firewall Rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:redis/firewallRule:FirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/Redis/cache1/firewallRules/rule1
 * ```
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallRuleState, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:redis/firewallRule:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    /**
     * The highest IP address included in the range.
     */
    public readonly endIp!: pulumi.Output<string>;
    /**
     * The name of the Firewall Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Redis Cache. Changing this forces a new resource to be created.
     */
    public readonly redisCacheName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which this Redis Cache exists.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The lowest IP address included in the range
     */
    public readonly startIp!: pulumi.Output<string>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs | FirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallRuleState | undefined;
            resourceInputs["endIp"] = state ? state.endIp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redisCacheName"] = state ? state.redisCacheName : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["startIp"] = state ? state.startIp : undefined;
        } else {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if ((!args || args.endIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endIp'");
            }
            if ((!args || args.redisCacheName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'redisCacheName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.startIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startIp'");
            }
            resourceInputs["endIp"] = args ? args.endIp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redisCacheName"] = args ? args.redisCacheName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["startIp"] = args ? args.startIp : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallRule resources.
 */
export interface FirewallRuleState {
    /**
     * The highest IP address included in the range.
     */
    endIp?: pulumi.Input<string>;
    /**
     * The name of the Firewall Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Redis Cache. Changing this forces a new resource to be created.
     */
    redisCacheName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which this Redis Cache exists.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The lowest IP address included in the range
     */
    startIp?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    /**
     * The highest IP address included in the range.
     */
    endIp: pulumi.Input<string>;
    /**
     * The name of the Firewall Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Redis Cache. Changing this forces a new resource to be created.
     */
    redisCacheName: pulumi.Input<string>;
    /**
     * The name of the resource group in which this Redis Cache exists.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The lowest IP address included in the range
     */
    startIp: pulumi.Input<string>;
}
