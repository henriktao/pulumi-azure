// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Topic authorization Rule within a ServiceBus Topic.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 *     tags: {
 *         source: "example",
 *     },
 * });
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {namespaceId: exampleNamespace.id});
 * const exampleTopicAuthorizationRule = new azure.servicebus.TopicAuthorizationRule("exampleTopicAuthorizationRule", {
 *     topicId: exampleTopic.id,
 *     listen: true,
 *     send: false,
 *     manage: false,
 * });
 * ```
 *
 * ## Import
 *
 * ServiceBus Topic authorization rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:servicebus/topicAuthorizationRule:TopicAuthorizationRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/topics/topic1/authorizationRules/rule1
 * ```
 */
export class TopicAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing TopicAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicAuthorizationRuleState, opts?: pulumi.CustomResourceOptions): TopicAuthorizationRule {
        return new TopicAuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicebus/topicAuthorizationRule:TopicAuthorizationRule';

    /**
     * Returns true if the given object is an instance of TopicAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicAuthorizationRule.__pulumiType;
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
     * Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The Primary Connection String for the ServiceBus Topic authorization Rule.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
     */
    public /*out*/ readonly primaryConnectionStringAlias!: pulumi.Output<string>;
    /**
     * The Primary Key for the ServiceBus Topic authorization Rule.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Connection String for the ServiceBus Topic authorization Rule.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The alias Secondary Connection String for the ServiceBus Namespace
     */
    public /*out*/ readonly secondaryConnectionStringAlias!: pulumi.Output<string>;
    /**
     * The Secondary Key for the ServiceBus Topic authorization Rule.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    public readonly send!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the ID of the ServiceBus Topic. Changing this forces a new resource to be created.
     */
    public readonly topicId!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a TopicAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicAuthorizationRuleArgs | TopicAuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicAuthorizationRuleState | undefined;
            resourceInputs["listen"] = state ? state.listen : undefined;
            resourceInputs["manage"] = state ? state.manage : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            resourceInputs["primaryConnectionStringAlias"] = state ? state.primaryConnectionStringAlias : undefined;
            resourceInputs["primaryKey"] = state ? state.primaryKey : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            resourceInputs["secondaryConnectionStringAlias"] = state ? state.secondaryConnectionStringAlias : undefined;
            resourceInputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            resourceInputs["send"] = state ? state.send : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as TopicAuthorizationRuleArgs | undefined;
            resourceInputs["listen"] = args ? args.listen : undefined;
            resourceInputs["manage"] = args ? args.manage : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["send"] = args ? args.send : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["primaryConnectionString"] = undefined /*out*/;
            resourceInputs["primaryConnectionStringAlias"] = undefined /*out*/;
            resourceInputs["primaryKey"] = undefined /*out*/;
            resourceInputs["secondaryConnectionString"] = undefined /*out*/;
            resourceInputs["secondaryConnectionStringAlias"] = undefined /*out*/;
            resourceInputs["secondaryKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure:eventhub/topicAuthorizationRule:TopicAuthorizationRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(TopicAuthorizationRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicAuthorizationRule resources.
 */
export interface TopicAuthorizationRuleState {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * The Primary Connection String for the ServiceBus Topic authorization Rule.
     */
    primaryConnectionString?: pulumi.Input<string>;
    /**
     * The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
     */
    primaryConnectionStringAlias?: pulumi.Input<string>;
    /**
     * The Primary Key for the ServiceBus Topic authorization Rule.
     */
    primaryKey?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Connection String for the ServiceBus Topic authorization Rule.
     */
    secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The alias Secondary Connection String for the ServiceBus Namespace
     */
    secondaryConnectionStringAlias?: pulumi.Input<string>;
    /**
     * The Secondary Key for the ServiceBus Topic authorization Rule.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the ServiceBus Topic. Changing this forces a new resource to be created.
     */
    topicId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicAuthorizationRule resource.
 */
export interface TopicAuthorizationRuleArgs {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    listen?: pulumi.Input<boolean>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the ServiceBus Topic Authorization Rule resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    send?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the ServiceBus Topic. Changing this forces a new resource to be created.
     */
    topicId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "topic_id"
     */
    topicName?: pulumi.Input<string>;
}
