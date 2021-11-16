// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Firewall Rule for a MySQL Flexible Server.
 *
 * ## Example Usage
 * ### Single IP Address)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFlexibleServer = new azure.mysql.FlexibleServer("exampleFlexibleServer", {});
 * // ...
 * const exampleFlexibleServerFirewallRule = new azure.mysql.FlexibleServerFirewallRule("exampleFlexibleServerFirewallRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleFlexibleServer.name,
 *     startIpAddress: "40.112.8.12",
 *     endIpAddress: "40.112.8.12",
 * });
 * ```
 * ### IP Range)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFlexibleServer = new azure.mysql.FlexibleServer("exampleFlexibleServer", {});
 * // ...
 * const exampleFlexibleServerFirewallRule = new azure.mysql.FlexibleServerFirewallRule("exampleFlexibleServerFirewallRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleFlexibleServer.name,
 *     startIpAddress: "40.112.0.0",
 *     endIpAddress: "40.112.255.255",
 * });
 * ```
 * ### Allow Access To Azure Services)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFlexibleServer = new azure.mysql.FlexibleServer("exampleFlexibleServer", {});
 * // ...
 * const exampleFlexibleServerFirewallRule = new azure.mysql.FlexibleServerFirewallRule("exampleFlexibleServerFirewallRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleFlexibleServer.name,
 *     startIpAddress: "0.0.0.0",
 *     endIpAddress: "0.0.0.0",
 * });
 * ```
 *
 * ## Import
 *
 * MySQL Firewall Rule's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mysql/flexibleServerFirewallRule:FlexibleServerFirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
 * ```
 */
export class FlexibleServerFirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FlexibleServerFirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlexibleServerFirewallRuleState, opts?: pulumi.CustomResourceOptions): FlexibleServerFirewallRule {
        return new FlexibleServerFirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mysql/flexibleServerFirewallRule:FlexibleServerFirewallRule';

    /**
     * Returns true if the given object is an instance of FlexibleServerFirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlexibleServerFirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlexibleServerFirewallRule.__pulumiType;
    }

    /**
     * Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    public readonly endIpAddress!: pulumi.Output<string>;
    /**
     * Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    public readonly startIpAddress!: pulumi.Output<string>;

    /**
     * Create a FlexibleServerFirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlexibleServerFirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlexibleServerFirewallRuleArgs | FlexibleServerFirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlexibleServerFirewallRuleState | undefined;
            inputs["endIpAddress"] = state ? state.endIpAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["startIpAddress"] = state ? state.startIpAddress : undefined;
        } else {
            const args = argsOrState as FlexibleServerFirewallRuleArgs | undefined;
            if ((!args || args.endIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endIpAddress'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.startIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startIpAddress'");
            }
            inputs["endIpAddress"] = args ? args.endIpAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["startIpAddress"] = args ? args.startIpAddress : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FlexibleServerFirewallRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlexibleServerFirewallRule resources.
 */
export interface FlexibleServerFirewallRuleState {
    /**
     * Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    endIpAddress?: pulumi.Input<string>;
    /**
     * Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    startIpAddress?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlexibleServerFirewallRule resource.
 */
export interface FlexibleServerFirewallRuleArgs {
    /**
     * Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    endIpAddress: pulumi.Input<string>;
    /**
     * Specifies the name of the MySQL Firewall Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
     */
    serverName: pulumi.Input<string>;
    /**
     * Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
     */
    startIpAddress: pulumi.Input<string>;
}
