// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Subscription Rule.
 *
 * ## Example Usage
 * ### SQL Filter)
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
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {
 *     namespaceId: exampleNamespace.id,
 *     enablePartitioning: true,
 * });
 * const exampleSubscription = new azure.servicebus.Subscription("exampleSubscription", {
 *     topicId: exampleTopic.id,
 *     maxDeliveryCount: 1,
 * });
 * const exampleSubscriptionRule = new azure.servicebus.SubscriptionRule("exampleSubscriptionRule", {
 *     subscriptionId: exampleSubscription.id,
 *     filterType: "SqlFilter",
 *     sqlFilter: "colour = 'red'",
 * });
 * ```
 * ### Correlation Filter)
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
 * const exampleTopic = new azure.servicebus.Topic("exampleTopic", {
 *     namespaceId: exampleNamespace.id,
 *     enablePartitioning: true,
 * });
 * const exampleSubscription = new azure.servicebus.Subscription("exampleSubscription", {
 *     topicId: exampleTopic.id,
 *     maxDeliveryCount: 1,
 * });
 * const exampleSubscriptionRule = new azure.servicebus.SubscriptionRule("exampleSubscriptionRule", {
 *     subscriptionId: exampleSubscription.id,
 *     filterType: "CorrelationFilter",
 *     correlationFilter: {
 *         correlationId: "high",
 *         label: "red",
 *         properties: {
 *             customProperty: "value",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Service Bus Subscription Rule can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:eventhub/subscriptionRule:SubscriptionRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
 * ```
 *
 * @deprecated azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule
 */
export class SubscriptionRule extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionRuleState, opts?: pulumi.CustomResourceOptions): SubscriptionRule {
        pulumi.log.warn("SubscriptionRule is deprecated: azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule")
        return new SubscriptionRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/subscriptionRule:SubscriptionRule';

    /**
     * Returns true if the given object is an instance of SubscriptionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionRule.__pulumiType;
    }

    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    public readonly correlationFilter!: pulumi.Output<outputs.eventhub.SubscriptionRuleCorrelationFilter | undefined>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    public readonly filterType!: pulumi.Output<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    public readonly sqlFilter!: pulumi.Output<string | undefined>;
    /**
     * The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    public readonly subscriptionId!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    public readonly subscriptionName!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a SubscriptionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule */
    constructor(name: string, args: SubscriptionRuleArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule */
    constructor(name: string, argsOrState?: SubscriptionRuleArgs | SubscriptionRuleState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SubscriptionRule is deprecated: azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["correlationFilter"] = state ? state.correlationFilter : undefined;
            resourceInputs["filterType"] = state ? state.filterType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["sqlFilter"] = state ? state.sqlFilter : undefined;
            resourceInputs["subscriptionId"] = state ? state.subscriptionId : undefined;
            resourceInputs["subscriptionName"] = state ? state.subscriptionName : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as SubscriptionRuleArgs | undefined;
            if ((!args || args.filterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filterType'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["correlationFilter"] = args ? args.correlationFilter : undefined;
            resourceInputs["filterType"] = args ? args.filterType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sqlFilter"] = args ? args.sqlFilter : undefined;
            resourceInputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            resourceInputs["subscriptionName"] = args ? args.subscriptionName : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubscriptionRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionRule resources.
 */
export interface SubscriptionRuleState {
    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    action?: pulumi.Input<string>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    correlationFilter?: pulumi.Input<inputs.eventhub.SubscriptionRuleCorrelationFilter>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    sqlFilter?: pulumi.Input<string>;
    /**
     * The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    subscriptionId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    subscriptionName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionRule resource.
 */
export interface SubscriptionRuleArgs {
    /**
     * Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
     */
    action?: pulumi.Input<string>;
    /**
     * A `correlationFilter` block as documented below to be evaluated against a BrokeredMessage. Required when `filterType` is set to `CorrelationFilter`.
     */
    correlationFilter?: pulumi.Input<inputs.eventhub.SubscriptionRuleCorrelationFilter>;
    /**
     * Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
     */
    filterType: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filterType` is set to `SqlFilter`.
     */
    sqlFilter?: pulumi.Input<string>;
    /**
     * The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
     */
    subscriptionId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    subscriptionName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "subscription_id"
     */
    topicName?: pulumi.Input<string>;
}
