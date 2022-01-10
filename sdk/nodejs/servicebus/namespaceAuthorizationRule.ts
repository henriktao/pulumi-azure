// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Namespace authorization Rule within a ServiceBus.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleNamespaceAuthorizationRule = new azure.servicebus.NamespaceAuthorizationRule("exampleNamespaceAuthorizationRule", {
 *     namespaceId: exampleNamespace.id,
 *     listen: true,
 *     send: true,
 *     manage: false,
 * });
 * ```
 *
 * ## Import
 *
 * ServiceBus Namespace authorization rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/AuthorizationRules/rule1
 * ```
 */
export class NamespaceAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceAuthorizationRuleState, opts?: pulumi.CustomResourceOptions): NamespaceAuthorizationRule {
        return new NamespaceAuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicebus/namespaceAuthorizationRule:NamespaceAuthorizationRule';

    /**
     * Returns true if the given object is an instance of NamespaceAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceAuthorizationRule.__pulumiType;
    }

    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    public readonly listen!: pulumi.Output<boolean | undefined>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    public readonly manage!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the ID of the ServiceBus Namespace. Changing this forces a new resource to be created.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The Primary Connection String for the ServiceBus Namespace authorization Rule.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
     */
    public /*out*/ readonly primaryConnectionStringAlias!: pulumi.Output<string>;
    /**
     * The Primary Key for the ServiceBus Namespace authorization Rule.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Connection String for the ServiceBus Namespace authorization Rule.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The alias Secondary Connection String for the ServiceBus Namespace
     */
    public /*out*/ readonly secondaryConnectionStringAlias!: pulumi.Output<string>;
    /**
     * The Secondary Key for the ServiceBus Namespace authorization Rule.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    public readonly send!: pulumi.Output<boolean | undefined>;

    /**
     * Create a NamespaceAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NamespaceAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceAuthorizationRuleArgs | NamespaceAuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceAuthorizationRuleState | undefined;
            inputs["listen"] = state ? state.listen : undefined;
            inputs["manage"] = state ? state.manage : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["primaryConnectionStringAlias"] = state ? state.primaryConnectionStringAlias : undefined;
            inputs["primaryKey"] = state ? state.primaryKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["secondaryConnectionStringAlias"] = state ? state.secondaryConnectionStringAlias : undefined;
            inputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            inputs["send"] = state ? state.send : undefined;
        } else {
            const args = argsOrState as NamespaceAuthorizationRuleArgs | undefined;
            inputs["listen"] = args ? args.listen : undefined;
            inputs["manage"] = args ? args.manage : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["send"] = args ? args.send : undefined;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["primaryConnectionStringAlias"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["secondaryConnectionStringAlias"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure:eventhub/namespaceAuthorizationRule:NamespaceAuthorizationRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(NamespaceAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NamespaceAuthorizationRule resources.
 */
export interface NamespaceAuthorizationRuleState {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the ID of the ServiceBus Namespace. Changing this forces a new resource to be created.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * The Primary Connection String for the ServiceBus Namespace authorization Rule.
     */
    primaryConnectionString?: pulumi.Input<string>;
    /**
     * The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
     */
    primaryConnectionStringAlias?: pulumi.Input<string>;
    /**
     * The Primary Key for the ServiceBus Namespace authorization Rule.
     */
    primaryKey?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Connection String for the ServiceBus Namespace authorization Rule.
     */
    secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The alias Secondary Connection String for the ServiceBus Namespace
     */
    secondaryConnectionStringAlias?: pulumi.Input<string>;
    /**
     * The Secondary Key for the ServiceBus Namespace authorization Rule.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NamespaceAuthorizationRule resource.
 */
export interface NamespaceAuthorizationRuleArgs {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the ID of the ServiceBus Namespace. Changing this forces a new resource to be created.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
}
