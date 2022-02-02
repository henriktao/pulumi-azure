// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Network Security Rule.
 *
 * > **NOTE on Network Security Groups and Network Security Rules:** This provider currently
 * provides both a standalone Network Security Rule resource, and allows for Network Security Rules to be defined in-line within the Network Security Group resource.
 * At this time you cannot use a Network Security Group with in-line Network Security Rules in conjunction with any Network Security Rule resources. Doing so will cause a conflict of rule settings and will overwrite rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleNetworkSecurityRule = new azure.network.NetworkSecurityRule("exampleNetworkSecurityRule", {
 *     priority: 100,
 *     direction: "Outbound",
 *     access: "Allow",
 *     protocol: "Tcp",
 *     sourcePortRange: "*",
 *     destinationPortRange: "*",
 *     sourceAddressPrefix: "*",
 *     destinationAddressPrefix: "*",
 *     resourceGroupName: exampleResourceGroup.name,
 *     networkSecurityGroupName: exampleNetworkSecurityGroup.name,
 * });
 * ```
 *
 * ## Import
 *
 * Network Security Rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/networkSecurityRule:NetworkSecurityRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1
 * ```
 */
export class NetworkSecurityRule extends pulumi.CustomResource {
    /**
     * Get an existing NetworkSecurityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions): NetworkSecurityRule {
        return new NetworkSecurityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkSecurityRule:NetworkSecurityRule';

    /**
     * Returns true if the given object is an instance of NetworkSecurityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkSecurityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkSecurityRule.__pulumiType;
    }

    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    public readonly access!: pulumi.Output<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
     */
    public readonly destinationAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
     */
    public readonly destinationAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * A List of destination Application Security Group ID's
     */
    public readonly destinationApplicationSecurityGroupIds!: pulumi.Output<string | undefined>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
     */
    public readonly destinationPortRange!: pulumi.Output<string | undefined>;
    /**
     * List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
     */
    public readonly destinationPortRanges!: pulumi.Output<string[] | undefined>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    public readonly networkSecurityGroupName!: pulumi.Output<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
     */
    public readonly sourceAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
     */
    public readonly sourceAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * A List of source Application Security Group ID's
     */
    public readonly sourceApplicationSecurityGroupIds!: pulumi.Output<string | undefined>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
     */
    public readonly sourcePortRange!: pulumi.Output<string | undefined>;
    /**
     * List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
     */
    public readonly sourcePortRanges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NetworkSecurityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkSecurityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkSecurityRuleArgs | NetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkSecurityRuleState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationAddressPrefix"] = state ? state.destinationAddressPrefix : undefined;
            resourceInputs["destinationAddressPrefixes"] = state ? state.destinationAddressPrefixes : undefined;
            resourceInputs["destinationApplicationSecurityGroupIds"] = state ? state.destinationApplicationSecurityGroupIds : undefined;
            resourceInputs["destinationPortRange"] = state ? state.destinationPortRange : undefined;
            resourceInputs["destinationPortRanges"] = state ? state.destinationPortRanges : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkSecurityGroupName"] = state ? state.networkSecurityGroupName : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["sourceAddressPrefix"] = state ? state.sourceAddressPrefix : undefined;
            resourceInputs["sourceAddressPrefixes"] = state ? state.sourceAddressPrefixes : undefined;
            resourceInputs["sourceApplicationSecurityGroupIds"] = state ? state.sourceApplicationSecurityGroupIds : undefined;
            resourceInputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
            resourceInputs["sourcePortRanges"] = state ? state.sourcePortRanges : undefined;
        } else {
            const args = argsOrState as NetworkSecurityRuleArgs | undefined;
            if ((!args || args.access === undefined) && !opts.urn) {
                throw new Error("Missing required property 'access'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.networkSecurityGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkSecurityGroupName'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationAddressPrefix"] = args ? args.destinationAddressPrefix : undefined;
            resourceInputs["destinationAddressPrefixes"] = args ? args.destinationAddressPrefixes : undefined;
            resourceInputs["destinationApplicationSecurityGroupIds"] = args ? args.destinationApplicationSecurityGroupIds : undefined;
            resourceInputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            resourceInputs["destinationPortRanges"] = args ? args.destinationPortRanges : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkSecurityGroupName"] = args ? args.networkSecurityGroupName : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sourceAddressPrefix"] = args ? args.sourceAddressPrefix : undefined;
            resourceInputs["sourceAddressPrefixes"] = args ? args.sourceAddressPrefixes : undefined;
            resourceInputs["sourceApplicationSecurityGroupIds"] = args ? args.sourceApplicationSecurityGroupIds : undefined;
            resourceInputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            resourceInputs["sourcePortRanges"] = args ? args.sourcePortRanges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkSecurityRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkSecurityRule resources.
 */
export interface NetworkSecurityRuleState {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    access?: pulumi.Input<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
     */
    destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
     */
    destinationAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of destination Application Security Group ID's
     */
    destinationApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
     */
    destinationPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    direction?: pulumi.Input<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    networkSecurityGroupName?: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    priority?: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
     */
    protocol?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
     */
    sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
     */
    sourceAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of source Application Security Group ID's
     */
    sourceApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
     */
    sourcePortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NetworkSecurityRule resource.
 */
export interface NetworkSecurityRuleArgs {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    access: pulumi.Input<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
     */
    destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
     */
    destinationAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of destination Application Security Group ID's
     */
    destinationApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
     */
    destinationPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    direction: pulumi.Input<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    networkSecurityGroupName: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    priority: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
     */
    protocol: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
     */
    sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
     */
    sourceAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of source Application Security Group ID's
     */
    sourceApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
     */
    sourcePortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}
