// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing ServiceBus Subscription.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.servicebus.getSubscription({
 *     name: "examplesubscription",
 *     resourceGroupName: "exampleresources",
 *     namespaceName: "examplenamespace",
 *     topicName: "exampletopic",
 * });
 * export const servicebusSubscription = data.azurerm_servicebus_namespace.example;
 * ```
 */
export function getSubscription(args: GetSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:servicebus/getSubscription:getSubscription", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubscription.
 */
export interface GetSubscriptionArgs {
    /**
     * Specifies the name of the ServiceBus Subscription.
     */
    name: string;
    /**
     * The name of the ServiceBus Namespace.
     */
    namespaceName: string;
    /**
     * Specifies the name of the Resource Group where the ServiceBus Namespace exists.
     */
    resourceGroupName: string;
    /**
     * The name of the ServiceBus Topic.
     */
    topicName: string;
}

/**
 * A collection of values returned by getSubscription.
 */
export interface GetSubscriptionResult {
    /**
     * The idle interval after which the topic is automatically deleted.
     */
    readonly autoDeleteOnIdle: string;
    /**
     * Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?
     */
    readonly deadLetteringOnFilterEvaluationError: boolean;
    /**
     * Does the Service Bus Subscription have dead letter support when a message expires?
     */
    readonly deadLetteringOnMessageExpiration: boolean;
    /**
     * The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    readonly defaultMessageTtl: string;
    /**
     * Are batched operations enabled on this ServiceBus Subscription?
     */
    readonly enableBatchedOperations: boolean;
    /**
     * The name of a Queue or Topic to automatically forward Dead Letter messages to.
     */
    readonly forwardDeadLetteredMessagesTo: string;
    /**
     * The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.
     */
    readonly forwardTo: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The lock duration for the subscription.
     */
    readonly lockDuration: string;
    /**
     * The maximum number of deliveries.
     */
    readonly maxDeliveryCount: number;
    readonly name: string;
    readonly namespaceName: string;
    /**
     * Whether or not this ServiceBus Subscription supports session.
     */
    readonly requiresSession: boolean;
    readonly resourceGroupName: string;
    readonly topicName: string;
}

export function getSubscriptionOutput(args: GetSubscriptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubscriptionResult> {
    return pulumi.output(args).apply(a => getSubscription(a, opts))
}

/**
 * A collection of arguments for invoking getSubscription.
 */
export interface GetSubscriptionOutputArgs {
    /**
     * Specifies the name of the ServiceBus Subscription.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the ServiceBus Namespace exists.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the ServiceBus Topic.
     */
    topicName: pulumi.Input<string>;
}
