// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EventGrid Event Subscription
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const defaultResourceGroup = new azure.core.ResourceGroup("defaultResourceGroup", {location: "West Europe"});
 * const defaultAccount = new azure.storage.Account("defaultAccount", {
 *     resourceGroupName: defaultResourceGroup.name,
 *     location: defaultResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const defaultQueue = new azure.storage.Queue("defaultQueue", {storageAccountName: defaultAccount.name});
 * const defaultEventSubscription = new azure.eventgrid.EventSubscription("defaultEventSubscription", {
 *     scope: defaultResourceGroup.id,
 *     storageQueueEndpoint: {
 *         storageAccountId: defaultAccount.id,
 *         queueName: defaultQueue.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * EventGrid Event Subscription's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:eventhub/eventSubscription:EventSubscription eventSubscription1
 * ```
 *
 *  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscription1
 *
 * @deprecated azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription
 */
export class EventSubscription extends pulumi.CustomResource {
    /**
     * Get an existing EventSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventSubscriptionState, opts?: pulumi.CustomResourceOptions): EventSubscription {
        pulumi.log.warn("EventSubscription is deprecated: azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription")
        return new EventSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/eventSubscription:EventSubscription';

    /**
     * Returns true if the given object is an instance of EventSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSubscription.__pulumiType;
    }

    /**
     * A `advancedFilter` block as defined below.
     */
    public readonly advancedFilter!: pulumi.Output<outputs.eventhub.EventSubscriptionAdvancedFilter | undefined>;
    /**
     * Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
     */
    public readonly advancedFilteringOnArraysEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An `azureFunctionEndpoint` block as defined below.
     */
    public readonly azureFunctionEndpoint!: pulumi.Output<outputs.eventhub.EventSubscriptionAzureFunctionEndpoint | undefined>;
    /**
     * A `deadLetterIdentity` block as defined below.
     */
    public readonly deadLetterIdentity!: pulumi.Output<outputs.eventhub.EventSubscriptionDeadLetterIdentity | undefined>;
    /**
     * A `deliveryIdentity` block as defined below.
     */
    public readonly deliveryIdentity!: pulumi.Output<outputs.eventhub.EventSubscriptionDeliveryIdentity | undefined>;
    /**
     * A `deliveryProperty` block as defined below.
     */
    public readonly deliveryProperties!: pulumi.Output<outputs.eventhub.EventSubscriptionDeliveryProperty[] | undefined>;
    /**
     * Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
     */
    public readonly eventDeliverySchema!: pulumi.Output<string | undefined>;
    /**
     * A `eventhubEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `eventhub_endpoint_id`
     */
    public readonly eventhubEndpoint!: pulumi.Output<outputs.eventhub.EventSubscriptionEventhubEndpoint>;
    /**
     * Specifies the id where the Event Hub is located.
     */
    public readonly eventhubEndpointId!: pulumi.Output<string>;
    /**
     * Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
     */
    public readonly expirationTimeUtc!: pulumi.Output<string | undefined>;
    /**
     * A `hybridConnectionEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `hybrid_connection_endpoint_id`
     */
    public readonly hybridConnectionEndpoint!: pulumi.Output<outputs.eventhub.EventSubscriptionHybridConnectionEndpoint>;
    /**
     * Specifies the id where the Hybrid Connection is located.
     */
    public readonly hybridConnectionEndpointId!: pulumi.Output<string>;
    /**
     * A list of applicable event types that need to be part of the event subscription.
     */
    public readonly includedEventTypes!: pulumi.Output<string[]>;
    /**
     * A list of labels to assign to the event subscription.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `retryPolicy` block as defined below.
     */
    public readonly retryPolicy!: pulumi.Output<outputs.eventhub.EventSubscriptionRetryPolicy>;
    /**
     * Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Specifies the id where the Service Bus Queue is located.
     */
    public readonly serviceBusQueueEndpointId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the id where the Service Bus Topic is located.
     */
    public readonly serviceBusTopicEndpointId!: pulumi.Output<string | undefined>;
    /**
     * A `storageBlobDeadLetterDestination` block as defined below.
     */
    public readonly storageBlobDeadLetterDestination!: pulumi.Output<outputs.eventhub.EventSubscriptionStorageBlobDeadLetterDestination | undefined>;
    /**
     * A `storageQueueEndpoint` block as defined below.
     */
    public readonly storageQueueEndpoint!: pulumi.Output<outputs.eventhub.EventSubscriptionStorageQueueEndpoint | undefined>;
    /**
     * A `subjectFilter` block as defined below.
     */
    public readonly subjectFilter!: pulumi.Output<outputs.eventhub.EventSubscriptionSubjectFilter | undefined>;
    /**
     * (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
     *
     * @deprecated This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
     */
    public readonly topicName!: pulumi.Output<string>;
    /**
     * A `webhookEndpoint` block as defined below.
     */
    public readonly webhookEndpoint!: pulumi.Output<outputs.eventhub.EventSubscriptionWebhookEndpoint | undefined>;

    /**
     * Create a EventSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription */
    constructor(name: string, args: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription */
    constructor(name: string, argsOrState?: EventSubscriptionArgs | EventSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EventSubscription is deprecated: azure.eventhub.EventSubscription has been deprecated in favor of azure.eventgrid.EventSubscription")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventSubscriptionState | undefined;
            resourceInputs["advancedFilter"] = state ? state.advancedFilter : undefined;
            resourceInputs["advancedFilteringOnArraysEnabled"] = state ? state.advancedFilteringOnArraysEnabled : undefined;
            resourceInputs["azureFunctionEndpoint"] = state ? state.azureFunctionEndpoint : undefined;
            resourceInputs["deadLetterIdentity"] = state ? state.deadLetterIdentity : undefined;
            resourceInputs["deliveryIdentity"] = state ? state.deliveryIdentity : undefined;
            resourceInputs["deliveryProperties"] = state ? state.deliveryProperties : undefined;
            resourceInputs["eventDeliverySchema"] = state ? state.eventDeliverySchema : undefined;
            resourceInputs["eventhubEndpoint"] = state ? state.eventhubEndpoint : undefined;
            resourceInputs["eventhubEndpointId"] = state ? state.eventhubEndpointId : undefined;
            resourceInputs["expirationTimeUtc"] = state ? state.expirationTimeUtc : undefined;
            resourceInputs["hybridConnectionEndpoint"] = state ? state.hybridConnectionEndpoint : undefined;
            resourceInputs["hybridConnectionEndpointId"] = state ? state.hybridConnectionEndpointId : undefined;
            resourceInputs["includedEventTypes"] = state ? state.includedEventTypes : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retryPolicy"] = state ? state.retryPolicy : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["serviceBusQueueEndpointId"] = state ? state.serviceBusQueueEndpointId : undefined;
            resourceInputs["serviceBusTopicEndpointId"] = state ? state.serviceBusTopicEndpointId : undefined;
            resourceInputs["storageBlobDeadLetterDestination"] = state ? state.storageBlobDeadLetterDestination : undefined;
            resourceInputs["storageQueueEndpoint"] = state ? state.storageQueueEndpoint : undefined;
            resourceInputs["subjectFilter"] = state ? state.subjectFilter : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
            resourceInputs["webhookEndpoint"] = state ? state.webhookEndpoint : undefined;
        } else {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["advancedFilter"] = args ? args.advancedFilter : undefined;
            resourceInputs["advancedFilteringOnArraysEnabled"] = args ? args.advancedFilteringOnArraysEnabled : undefined;
            resourceInputs["azureFunctionEndpoint"] = args ? args.azureFunctionEndpoint : undefined;
            resourceInputs["deadLetterIdentity"] = args ? args.deadLetterIdentity : undefined;
            resourceInputs["deliveryIdentity"] = args ? args.deliveryIdentity : undefined;
            resourceInputs["deliveryProperties"] = args ? args.deliveryProperties : undefined;
            resourceInputs["eventDeliverySchema"] = args ? args.eventDeliverySchema : undefined;
            resourceInputs["eventhubEndpoint"] = args ? args.eventhubEndpoint : undefined;
            resourceInputs["eventhubEndpointId"] = args ? args.eventhubEndpointId : undefined;
            resourceInputs["expirationTimeUtc"] = args ? args.expirationTimeUtc : undefined;
            resourceInputs["hybridConnectionEndpoint"] = args ? args.hybridConnectionEndpoint : undefined;
            resourceInputs["hybridConnectionEndpointId"] = args ? args.hybridConnectionEndpointId : undefined;
            resourceInputs["includedEventTypes"] = args ? args.includedEventTypes : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["serviceBusQueueEndpointId"] = args ? args.serviceBusQueueEndpointId : undefined;
            resourceInputs["serviceBusTopicEndpointId"] = args ? args.serviceBusTopicEndpointId : undefined;
            resourceInputs["storageBlobDeadLetterDestination"] = args ? args.storageBlobDeadLetterDestination : undefined;
            resourceInputs["storageQueueEndpoint"] = args ? args.storageQueueEndpoint : undefined;
            resourceInputs["subjectFilter"] = args ? args.subjectFilter : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["webhookEndpoint"] = args ? args.webhookEndpoint : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSubscription resources.
 */
export interface EventSubscriptionState {
    /**
     * A `advancedFilter` block as defined below.
     */
    advancedFilter?: pulumi.Input<inputs.eventhub.EventSubscriptionAdvancedFilter>;
    /**
     * Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
     */
    advancedFilteringOnArraysEnabled?: pulumi.Input<boolean>;
    /**
     * An `azureFunctionEndpoint` block as defined below.
     */
    azureFunctionEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionAzureFunctionEndpoint>;
    /**
     * A `deadLetterIdentity` block as defined below.
     */
    deadLetterIdentity?: pulumi.Input<inputs.eventhub.EventSubscriptionDeadLetterIdentity>;
    /**
     * A `deliveryIdentity` block as defined below.
     */
    deliveryIdentity?: pulumi.Input<inputs.eventhub.EventSubscriptionDeliveryIdentity>;
    /**
     * A `deliveryProperty` block as defined below.
     */
    deliveryProperties?: pulumi.Input<pulumi.Input<inputs.eventhub.EventSubscriptionDeliveryProperty>[]>;
    /**
     * Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
     */
    eventDeliverySchema?: pulumi.Input<string>;
    /**
     * A `eventhubEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `eventhub_endpoint_id`
     */
    eventhubEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionEventhubEndpoint>;
    /**
     * Specifies the id where the Event Hub is located.
     */
    eventhubEndpointId?: pulumi.Input<string>;
    /**
     * Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
     */
    expirationTimeUtc?: pulumi.Input<string>;
    /**
     * A `hybridConnectionEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `hybrid_connection_endpoint_id`
     */
    hybridConnectionEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionHybridConnectionEndpoint>;
    /**
     * Specifies the id where the Hybrid Connection is located.
     */
    hybridConnectionEndpointId?: pulumi.Input<string>;
    /**
     * A list of applicable event types that need to be part of the event subscription.
     */
    includedEventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of labels to assign to the event subscription.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `retryPolicy` block as defined below.
     */
    retryPolicy?: pulumi.Input<inputs.eventhub.EventSubscriptionRetryPolicy>;
    /**
     * Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
     */
    scope?: pulumi.Input<string>;
    /**
     * Specifies the id where the Service Bus Queue is located.
     */
    serviceBusQueueEndpointId?: pulumi.Input<string>;
    /**
     * Specifies the id where the Service Bus Topic is located.
     */
    serviceBusTopicEndpointId?: pulumi.Input<string>;
    /**
     * A `storageBlobDeadLetterDestination` block as defined below.
     */
    storageBlobDeadLetterDestination?: pulumi.Input<inputs.eventhub.EventSubscriptionStorageBlobDeadLetterDestination>;
    /**
     * A `storageQueueEndpoint` block as defined below.
     */
    storageQueueEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionStorageQueueEndpoint>;
    /**
     * A `subjectFilter` block as defined below.
     */
    subjectFilter?: pulumi.Input<inputs.eventhub.EventSubscriptionSubjectFilter>;
    /**
     * (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
     *
     * @deprecated This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
     */
    topicName?: pulumi.Input<string>;
    /**
     * A `webhookEndpoint` block as defined below.
     */
    webhookEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionWebhookEndpoint>;
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    /**
     * A `advancedFilter` block as defined below.
     */
    advancedFilter?: pulumi.Input<inputs.eventhub.EventSubscriptionAdvancedFilter>;
    /**
     * Specifies whether advanced filters should be evaluated against an array of values instead of expecting a singular value. Defaults to `false`.
     */
    advancedFilteringOnArraysEnabled?: pulumi.Input<boolean>;
    /**
     * An `azureFunctionEndpoint` block as defined below.
     */
    azureFunctionEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionAzureFunctionEndpoint>;
    /**
     * A `deadLetterIdentity` block as defined below.
     */
    deadLetterIdentity?: pulumi.Input<inputs.eventhub.EventSubscriptionDeadLetterIdentity>;
    /**
     * A `deliveryIdentity` block as defined below.
     */
    deliveryIdentity?: pulumi.Input<inputs.eventhub.EventSubscriptionDeliveryIdentity>;
    /**
     * A `deliveryProperty` block as defined below.
     */
    deliveryProperties?: pulumi.Input<pulumi.Input<inputs.eventhub.EventSubscriptionDeliveryProperty>[]>;
    /**
     * Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventSchemaV1_0`, `CustomInputSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.
     */
    eventDeliverySchema?: pulumi.Input<string>;
    /**
     * A `eventhubEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `eventhub_endpoint_id`
     */
    eventhubEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionEventhubEndpoint>;
    /**
     * Specifies the id where the Event Hub is located.
     */
    eventhubEndpointId?: pulumi.Input<string>;
    /**
     * Specifies the expiration time of the event subscription (Datetime Format `RFC 3339`).
     */
    expirationTimeUtc?: pulumi.Input<string>;
    /**
     * A `hybridConnectionEndpoint` block as defined below.
     *
     * @deprecated Deprecated in favour of `hybrid_connection_endpoint_id`
     */
    hybridConnectionEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionHybridConnectionEndpoint>;
    /**
     * Specifies the id where the Hybrid Connection is located.
     */
    hybridConnectionEndpointId?: pulumi.Input<string>;
    /**
     * A list of applicable event types that need to be part of the event subscription.
     */
    includedEventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of labels to assign to the event subscription.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `retryPolicy` block as defined below.
     */
    retryPolicy?: pulumi.Input<inputs.eventhub.EventSubscriptionRetryPolicy>;
    /**
     * Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
     */
    scope: pulumi.Input<string>;
    /**
     * Specifies the id where the Service Bus Queue is located.
     */
    serviceBusQueueEndpointId?: pulumi.Input<string>;
    /**
     * Specifies the id where the Service Bus Topic is located.
     */
    serviceBusTopicEndpointId?: pulumi.Input<string>;
    /**
     * A `storageBlobDeadLetterDestination` block as defined below.
     */
    storageBlobDeadLetterDestination?: pulumi.Input<inputs.eventhub.EventSubscriptionStorageBlobDeadLetterDestination>;
    /**
     * A `storageQueueEndpoint` block as defined below.
     */
    storageQueueEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionStorageQueueEndpoint>;
    /**
     * A `subjectFilter` block as defined below.
     */
    subjectFilter?: pulumi.Input<inputs.eventhub.EventSubscriptionSubjectFilter>;
    /**
     * (Optional/ **Deprecated) Specifies the name of the topic to associate with the event subscription.
     *
     * @deprecated This field has been updated to readonly field since Apr 25, 2019 so no longer has any affect and will be removed in version 3.0 of the provider.
     */
    topicName?: pulumi.Input<string>;
    /**
     * A `webhookEndpoint` block as defined below.
     */
    webhookEndpoint?: pulumi.Input<inputs.eventhub.EventSubscriptionWebhookEndpoint>;
}
