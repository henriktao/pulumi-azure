// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Relay Hybrid Connection Authorization Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.relay.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "Standard",
 *     tags: {
 *         source: "terraform",
 *     },
 * });
 * const exampleHybridConnection = new azure.relay.HybridConnection("exampleHybridConnection", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     relayNamespaceName: exampleNamespace.name,
 *     requiresClientAuthorization: false,
 *     userMetadata: "testmetadata",
 * });
 * const exampleHybridConnectionAuthorizationRule = new azure.relay.HybridConnectionAuthorizationRule("exampleHybridConnectionAuthorizationRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     hybridConnectionName: exampleHybridConnection.name,
 *     namespaceName: exampleNamespace.name,
 *     listen: true,
 *     send: true,
 *     manage: false,
 * });
 * ```
 *
 * ## Import
 *
 * Azure Relay Hybrid Connection Authorization Rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:relay/hybridConnectionAuthorizationRule:HybridConnectionAuthorizationRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/hybridConnections/connection1/authorizationRules/rule1
 * ```
 */
export class HybridConnectionAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing HybridConnectionAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HybridConnectionAuthorizationRuleState, opts?: pulumi.CustomResourceOptions): HybridConnectionAuthorizationRule {
        return new HybridConnectionAuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:relay/hybridConnectionAuthorizationRule:HybridConnectionAuthorizationRule';

    /**
     * Returns true if the given object is an instance of HybridConnectionAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridConnectionAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridConnectionAuthorizationRule.__pulumiType;
    }

    /**
     * Name of the Azure Relay Hybrid Connection for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    public readonly hybridConnectionName!: pulumi.Output<string>;
    /**
     * Grants listen access to this Authorization Rule. Defaults to `false`.
     */
    public readonly listen!: pulumi.Output<boolean | undefined>;
    /**
     * Grants manage access to this Authorization Rule. When this property is `true` - both `listen` and `send` must be set to `true` too. Defaults to `false`.
     */
    public readonly manage!: pulumi.Output<boolean | undefined>;
    /**
     * The name which should be used for this Azure Relay Hybrid Connection Authorization Rule. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the Azure Relay Namespace for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The Primary Connection String for the Azure Relay Hybrid Connection Authorization Rule.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The Primary Key for the Azure Relay Hybrid Connection Authorization Rule.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Azure Relay Hybrid Connection Authorization Rule should exist. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Connection String for the Azure Relay Hybrid Connection Authorization Rule.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The Secondary Key for the Azure Relay Hybrid Connection Authorization Rule.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * Grants send access to this Authorization Rule. Defaults to `false`.
     */
    public readonly send!: pulumi.Output<boolean | undefined>;

    /**
     * Create a HybridConnectionAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridConnectionAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HybridConnectionAuthorizationRuleArgs | HybridConnectionAuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HybridConnectionAuthorizationRuleState | undefined;
            inputs["hybridConnectionName"] = state ? state.hybridConnectionName : undefined;
            inputs["listen"] = state ? state.listen : undefined;
            inputs["manage"] = state ? state.manage : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["primaryKey"] = state ? state.primaryKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            inputs["send"] = state ? state.send : undefined;
        } else {
            const args = argsOrState as HybridConnectionAuthorizationRuleArgs | undefined;
            if ((!args || args.hybridConnectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hybridConnectionName'");
            }
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hybridConnectionName"] = args ? args.hybridConnectionName : undefined;
            inputs["listen"] = args ? args.listen : undefined;
            inputs["manage"] = args ? args.manage : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["send"] = args ? args.send : undefined;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(HybridConnectionAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HybridConnectionAuthorizationRule resources.
 */
export interface HybridConnectionAuthorizationRuleState {
    /**
     * Name of the Azure Relay Hybrid Connection for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    hybridConnectionName?: pulumi.Input<string>;
    /**
     * Grants listen access to this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this Authorization Rule. When this property is `true` - both `listen` and `send` must be set to `true` too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * The name which should be used for this Azure Relay Hybrid Connection Authorization Rule. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the Azure Relay Namespace for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * The Primary Connection String for the Azure Relay Hybrid Connection Authorization Rule.
     */
    primaryConnectionString?: pulumi.Input<string>;
    /**
     * The Primary Key for the Azure Relay Hybrid Connection Authorization Rule.
     */
    primaryKey?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Azure Relay Hybrid Connection Authorization Rule should exist. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Connection String for the Azure Relay Hybrid Connection Authorization Rule.
     */
    secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The Secondary Key for the Azure Relay Hybrid Connection Authorization Rule.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * Grants send access to this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a HybridConnectionAuthorizationRule resource.
 */
export interface HybridConnectionAuthorizationRuleArgs {
    /**
     * Name of the Azure Relay Hybrid Connection for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    hybridConnectionName: pulumi.Input<string>;
    /**
     * Grants listen access to this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this Authorization Rule. When this property is `true` - both `listen` and `send` must be set to `true` too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * The name which should be used for this Azure Relay Hybrid Connection Authorization Rule. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the Azure Relay Namespace for which this Azure Relay Hybrid Connection Authorization Rule will be created. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Azure Relay Hybrid Connection Authorization Rule should exist. Changing this forces a new Azure Relay Hybrid Connection Authorization Rule to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Grants send access to this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
}
