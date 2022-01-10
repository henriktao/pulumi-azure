// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a ServiceBus Queue.
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
 * const exampleQueue = new azure.servicebus.Queue("exampleQueue", {
 *     namespaceId: exampleNamespace.id,
 *     enablePartitioning: true,
 * });
 * ```
 *
 * ## Import
 *
 * Service Bus Queue can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:eventhub/queue:Queue example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/queues/snqueue1
 * ```
 *
 * @deprecated azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        pulumi.log.warn("Queue is deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue")
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
     */
    public readonly autoDeleteOnIdle!: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
     */
    public readonly deadLetteringOnMessageExpiration!: pulumi.Output<boolean | undefined>;
    /**
     * The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
     */
    public readonly defaultMessageTtl!: pulumi.Output<string>;
    /**
     * The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
     */
    public readonly duplicateDetectionHistoryTimeWindow!: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
     */
    public readonly enableBatchedOperations!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
     */
    public readonly enableExpress!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
     */
    public readonly enablePartitioning!: pulumi.Output<boolean | undefined>;
    /**
     * The name of a Queue or Topic to automatically forward dead lettered messages to.
     */
    public readonly forwardDeadLetteredMessagesTo!: pulumi.Output<string | undefined>;
    /**
     * The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
     */
    public readonly forwardTo!: pulumi.Output<string | undefined>;
    /**
     * The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
     */
    public readonly lockDuration!: pulumi.Output<string>;
    /**
     * Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
     */
    public readonly maxDeliveryCount!: pulumi.Output<number | undefined>;
    /**
     * Integer value which controls the maximum size of
     * a message allowed on the queue for Premium SKU. For supported values see the "Large messages support"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
     */
    public readonly maxMessageSizeInKilobytes!: pulumi.Output<number>;
    /**
     * Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
     */
    public readonly maxSizeInMegabytes!: pulumi.Output<number>;
    /**
     * Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
     */
    public readonly requiresDuplicateDetection!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
     */
    public readonly requiresSession!: pulumi.Output<boolean | undefined>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
     */
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue */
    constructor(name: string, args?: QueueArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue */
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Queue is deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueueState | undefined;
            inputs["autoDeleteOnIdle"] = state ? state.autoDeleteOnIdle : undefined;
            inputs["deadLetteringOnMessageExpiration"] = state ? state.deadLetteringOnMessageExpiration : undefined;
            inputs["defaultMessageTtl"] = state ? state.defaultMessageTtl : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = state ? state.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = state ? state.enableBatchedOperations : undefined;
            inputs["enableExpress"] = state ? state.enableExpress : undefined;
            inputs["enablePartitioning"] = state ? state.enablePartitioning : undefined;
            inputs["forwardDeadLetteredMessagesTo"] = state ? state.forwardDeadLetteredMessagesTo : undefined;
            inputs["forwardTo"] = state ? state.forwardTo : undefined;
            inputs["lockDuration"] = state ? state.lockDuration : undefined;
            inputs["maxDeliveryCount"] = state ? state.maxDeliveryCount : undefined;
            inputs["maxMessageSizeInKilobytes"] = state ? state.maxMessageSizeInKilobytes : undefined;
            inputs["maxSizeInMegabytes"] = state ? state.maxSizeInMegabytes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["requiresDuplicateDetection"] = state ? state.requiresDuplicateDetection : undefined;
            inputs["requiresSession"] = state ? state.requiresSession : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            inputs["autoDeleteOnIdle"] = args ? args.autoDeleteOnIdle : undefined;
            inputs["deadLetteringOnMessageExpiration"] = args ? args.deadLetteringOnMessageExpiration : undefined;
            inputs["defaultMessageTtl"] = args ? args.defaultMessageTtl : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = args ? args.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = args ? args.enableBatchedOperations : undefined;
            inputs["enableExpress"] = args ? args.enableExpress : undefined;
            inputs["enablePartitioning"] = args ? args.enablePartitioning : undefined;
            inputs["forwardDeadLetteredMessagesTo"] = args ? args.forwardDeadLetteredMessagesTo : undefined;
            inputs["forwardTo"] = args ? args.forwardTo : undefined;
            inputs["lockDuration"] = args ? args.lockDuration : undefined;
            inputs["maxDeliveryCount"] = args ? args.maxDeliveryCount : undefined;
            inputs["maxMessageSizeInKilobytes"] = args ? args.maxMessageSizeInKilobytes : undefined;
            inputs["maxSizeInMegabytes"] = args ? args.maxSizeInMegabytes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["requiresDuplicateDetection"] = args ? args.requiresDuplicateDetection : undefined;
            inputs["requiresSession"] = args ? args.requiresSession : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
     */
    autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
     */
    deadLetteringOnMessageExpiration?: pulumi.Input<boolean>;
    /**
     * The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
     */
    defaultMessageTtl?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
     */
    duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
     */
    enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
     */
    enableExpress?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
     */
    enablePartitioning?: pulumi.Input<boolean>;
    /**
     * The name of a Queue or Topic to automatically forward dead lettered messages to.
     */
    forwardDeadLetteredMessagesTo?: pulumi.Input<string>;
    /**
     * The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
     */
    forwardTo?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
     */
    lockDuration?: pulumi.Input<string>;
    /**
     * Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
     */
    maxDeliveryCount?: pulumi.Input<number>;
    /**
     * Integer value which controls the maximum size of
     * a message allowed on the queue for Premium SKU. For supported values see the "Large messages support"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
     */
    maxMessageSizeInKilobytes?: pulumi.Input<number>;
    /**
     * Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
     */
    maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
     */
    requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
     */
    requiresSession?: pulumi.Input<boolean>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
     */
    autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
     */
    deadLetteringOnMessageExpiration?: pulumi.Input<boolean>;
    /**
     * The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
     */
    defaultMessageTtl?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
     */
    duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
     */
    enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
     */
    enableExpress?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
     */
    enablePartitioning?: pulumi.Input<boolean>;
    /**
     * The name of a Queue or Topic to automatically forward dead lettered messages to.
     */
    forwardDeadLetteredMessagesTo?: pulumi.Input<string>;
    /**
     * The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
     */
    forwardTo?: pulumi.Input<string>;
    /**
     * The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
     */
    lockDuration?: pulumi.Input<string>;
    /**
     * Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
     */
    maxDeliveryCount?: pulumi.Input<number>;
    /**
     * Integer value which controls the maximum size of
     * a message allowed on the queue for Premium SKU. For supported values see the "Large messages support"
     * section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
     */
    maxMessageSizeInKilobytes?: pulumi.Input<number>;
    /**
     * Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
     */
    maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
     */
    requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
     */
    requiresSession?: pulumi.Input<boolean>;
    /**
     * @deprecated Deprecated in favor of "namespace_id"
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
     */
    status?: pulumi.Input<string>;
}
