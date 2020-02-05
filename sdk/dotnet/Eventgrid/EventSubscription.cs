// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    /// <summary>
    /// Manages an EventGrid Event Subscription
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventgrid_event_subscription.html.markdown.
    /// </summary>
    public partial class EventSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
        /// </summary>
        [Output("eventDeliverySchema")]
        public Output<string?> EventDeliverySchema { get; private set; } = null!;

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Output("eventhubEndpoint")]
        public Output<Outputs.EventSubscriptionEventhubEndpoint?> EventhubEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Output("hybridConnectionEndpoint")]
        public Output<Outputs.EventSubscriptionHybridConnectionEndpoint?> HybridConnectionEndpoint { get; private set; } = null!;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        [Output("includedEventTypes")]
        public Output<ImmutableArray<string>> IncludedEventTypes { get; private set; } = null!;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Output("retryPolicy")]
        public Output<Outputs.EventSubscriptionRetryPolicy> RetryPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Output("storageBlobDeadLetterDestination")]
        public Output<Outputs.EventSubscriptionStorageBlobDeadLetterDestination?> StorageBlobDeadLetterDestination { get; private set; } = null!;

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Output("storageQueueEndpoint")]
        public Output<Outputs.EventSubscriptionStorageQueueEndpoint?> StorageQueueEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Output("subjectFilter")]
        public Output<Outputs.EventSubscriptionSubjectFilter?> SubjectFilter { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Output("webhookEndpoint")]
        public Output<Outputs.EventSubscriptionWebhookEndpoint?> WebhookEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a EventSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventSubscription(string name, EventSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure:eventgrid/eventSubscription:EventSubscription", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EventSubscription(string name, Input<string> id, EventSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventgrid/eventSubscription:EventSubscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "azure:eventhub/eventSubscription:EventSubscription" } },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventSubscription Get(string name, Input<string> id, EventSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new EventSubscription(name, id, state, options);
        }
    }

    public sealed class EventSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Input("eventhubEndpoint")]
        public Input<Inputs.EventSubscriptionEventhubEndpointArgs>? EventhubEndpoint { get; set; }

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Input("hybridConnectionEndpoint")]
        public Input<Inputs.EventSubscriptionHybridConnectionEndpointArgs>? HybridConnectionEndpoint { get; set; }

        [Input("includedEventTypes")]
        private InputList<string>? _includedEventTypes;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        public InputList<string> IncludedEventTypes
        {
            get => _includedEventTypes ?? (_includedEventTypes = new InputList<string>());
            set => _includedEventTypes = value;
        }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventSubscriptionRetryPolicyArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Input("storageBlobDeadLetterDestination")]
        public Input<Inputs.EventSubscriptionStorageBlobDeadLetterDestinationArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.EventSubscriptionStorageQueueEndpointArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.EventSubscriptionSubjectFilterArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.EventSubscriptionWebhookEndpointArgs>? WebhookEndpoint { get; set; }

        public EventSubscriptionArgs()
        {
        }
    }

    public sealed class EventSubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the event delivery schema for the event subscription. Possible values include: `EventGridSchema`, `CloudEventV01Schema`, `CustomInputSchema`.
        /// </summary>
        [Input("eventDeliverySchema")]
        public Input<string>? EventDeliverySchema { get; set; }

        /// <summary>
        /// A `eventhub_endpoint` block as defined below.
        /// </summary>
        [Input("eventhubEndpoint")]
        public Input<Inputs.EventSubscriptionEventhubEndpointGetArgs>? EventhubEndpoint { get; set; }

        /// <summary>
        /// A `hybrid_connection_endpoint` block as defined below.
        /// </summary>
        [Input("hybridConnectionEndpoint")]
        public Input<Inputs.EventSubscriptionHybridConnectionEndpointGetArgs>? HybridConnectionEndpoint { get; set; }

        [Input("includedEventTypes")]
        private InputList<string>? _includedEventTypes;

        /// <summary>
        /// A list of applicable event types that need to be part of the event subscription.
        /// </summary>
        public InputList<string> IncludedEventTypes
        {
            get => _includedEventTypes ?? (_includedEventTypes = new InputList<string>());
            set => _includedEventTypes = value;
        }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// A list of labels to assign to the event subscription.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the name of the EventGrid Event Subscription resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retry_policy` block as defined below.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.EventSubscriptionRetryPolicyGetArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// Specifies the scope at which the EventGrid Event Subscription should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// A `storage_blob_dead_letter_destination` block as defined below.
        /// </summary>
        [Input("storageBlobDeadLetterDestination")]
        public Input<Inputs.EventSubscriptionStorageBlobDeadLetterDestinationGetArgs>? StorageBlobDeadLetterDestination { get; set; }

        /// <summary>
        /// A `storage_queue_endpoint` block as defined below.
        /// </summary>
        [Input("storageQueueEndpoint")]
        public Input<Inputs.EventSubscriptionStorageQueueEndpointGetArgs>? StorageQueueEndpoint { get; set; }

        /// <summary>
        /// A `subject_filter` block as defined below.
        /// </summary>
        [Input("subjectFilter")]
        public Input<Inputs.EventSubscriptionSubjectFilterGetArgs>? SubjectFilter { get; set; }

        /// <summary>
        /// Specifies the name of the topic to associate with the event subscription.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// A `webhook_endpoint` block as defined below.
        /// </summary>
        [Input("webhookEndpoint")]
        public Input<Inputs.EventSubscriptionWebhookEndpointGetArgs>? WebhookEndpoint { get; set; }

        public EventSubscriptionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EventSubscriptionEventhubEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the eventhub where the Event Subscription will receive events.
        /// </summary>
        [Input("eventhubId", required: true)]
        public Input<string> EventhubId { get; set; } = null!;

        public EventSubscriptionEventhubEndpointArgs()
        {
        }
    }

    public sealed class EventSubscriptionEventhubEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the eventhub where the Event Subscription will receive events.
        /// </summary>
        [Input("eventhubId", required: true)]
        public Input<string> EventhubId { get; set; } = null!;

        public EventSubscriptionEventhubEndpointGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionHybridConnectionEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the hybrid connection where the Event Subscription will receive events.
        /// </summary>
        [Input("hybridConnectionId", required: true)]
        public Input<string> HybridConnectionId { get; set; } = null!;

        public EventSubscriptionHybridConnectionEndpointArgs()
        {
        }
    }

    public sealed class EventSubscriptionHybridConnectionEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the hybrid connection where the Event Subscription will receive events.
        /// </summary>
        [Input("hybridConnectionId", required: true)]
        public Input<string> HybridConnectionId { get; set; } = null!;

        public EventSubscriptionHybridConnectionEndpointGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionRetryPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the time to live (in minutes) for events.
        /// </summary>
        [Input("eventTimeToLive", required: true)]
        public Input<int> EventTimeToLive { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum number of delivery retry attempts for events.
        /// </summary>
        [Input("maxDeliveryAttempts", required: true)]
        public Input<int> MaxDeliveryAttempts { get; set; } = null!;

        public EventSubscriptionRetryPolicyArgs()
        {
        }
    }

    public sealed class EventSubscriptionRetryPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the time to live (in minutes) for events.
        /// </summary>
        [Input("eventTimeToLive", required: true)]
        public Input<int> EventTimeToLive { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum number of delivery retry attempts for events.
        /// </summary>
        [Input("maxDeliveryAttempts", required: true)]
        public Input<int> MaxDeliveryAttempts { get; set; } = null!;

        public EventSubscriptionRetryPolicyGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionStorageBlobDeadLetterDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Storage blob container that is the destination of the deadletter events
        /// </summary>
        [Input("storageBlobContainerName", required: true)]
        public Input<string> StorageBlobContainerName { get; set; } = null!;

        public EventSubscriptionStorageBlobDeadLetterDestinationArgs()
        {
        }
    }

    public sealed class EventSubscriptionStorageBlobDeadLetterDestinationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Storage blob container that is the destination of the deadletter events
        /// </summary>
        [Input("storageBlobContainerName", required: true)]
        public Input<string> StorageBlobContainerName { get; set; } = null!;

        public EventSubscriptionStorageBlobDeadLetterDestinationGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionStorageQueueEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the storage queue where the Event Subscriptio will receive events.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public EventSubscriptionStorageQueueEndpointArgs()
        {
        }
    }

    public sealed class EventSubscriptionStorageQueueEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the storage queue where the Event Subscriptio will receive events.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public EventSubscriptionStorageQueueEndpointGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionSubjectFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if `subject_begins_with` and `subject_ends_with` case sensitive. This value defaults to `false`.
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// A string to filter events for an event subscription based on a resource path prefix.
        /// </summary>
        [Input("subjectBeginsWith")]
        public Input<string>? SubjectBeginsWith { get; set; }

        /// <summary>
        /// A string to filter events for an event subscription based on a resource path suffix.
        /// </summary>
        [Input("subjectEndsWith")]
        public Input<string>? SubjectEndsWith { get; set; }

        public EventSubscriptionSubjectFilterArgs()
        {
        }
    }

    public sealed class EventSubscriptionSubjectFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if `subject_begins_with` and `subject_ends_with` case sensitive. This value defaults to `false`.
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// A string to filter events for an event subscription based on a resource path prefix.
        /// </summary>
        [Input("subjectBeginsWith")]
        public Input<string>? SubjectBeginsWith { get; set; }

        /// <summary>
        /// A string to filter events for an event subscription based on a resource path suffix.
        /// </summary>
        [Input("subjectEndsWith")]
        public Input<string>? SubjectEndsWith { get; set; }

        public EventSubscriptionSubjectFilterGetArgs()
        {
        }
    }

    public sealed class EventSubscriptionWebhookEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the url of the webhook where the Event Subscription will receive events.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public EventSubscriptionWebhookEndpointArgs()
        {
        }
    }

    public sealed class EventSubscriptionWebhookEndpointGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the url of the webhook where the Event Subscription will receive events.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public EventSubscriptionWebhookEndpointGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EventSubscriptionEventhubEndpoint
    {
        /// <summary>
        /// Specifies the id of the eventhub where the Event Subscription will receive events.
        /// </summary>
        public readonly string EventhubId;

        [OutputConstructor]
        private EventSubscriptionEventhubEndpoint(string eventhubId)
        {
            EventhubId = eventhubId;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionHybridConnectionEndpoint
    {
        /// <summary>
        /// Specifies the id of the hybrid connection where the Event Subscription will receive events.
        /// </summary>
        public readonly string HybridConnectionId;

        [OutputConstructor]
        private EventSubscriptionHybridConnectionEndpoint(string hybridConnectionId)
        {
            HybridConnectionId = hybridConnectionId;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionRetryPolicy
    {
        /// <summary>
        /// Specifies the time to live (in minutes) for events.
        /// </summary>
        public readonly int EventTimeToLive;
        /// <summary>
        /// Specifies the maximum number of delivery retry attempts for events.
        /// </summary>
        public readonly int MaxDeliveryAttempts;

        [OutputConstructor]
        private EventSubscriptionRetryPolicy(
            int eventTimeToLive,
            int maxDeliveryAttempts)
        {
            EventTimeToLive = eventTimeToLive;
            MaxDeliveryAttempts = maxDeliveryAttempts;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionStorageBlobDeadLetterDestination
    {
        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        public readonly string StorageAccountId;
        /// <summary>
        /// Specifies the name of the Storage blob container that is the destination of the deadletter events
        /// </summary>
        public readonly string StorageBlobContainerName;

        [OutputConstructor]
        private EventSubscriptionStorageBlobDeadLetterDestination(
            string storageAccountId,
            string storageBlobContainerName)
        {
            StorageAccountId = storageAccountId;
            StorageBlobContainerName = storageBlobContainerName;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionStorageQueueEndpoint
    {
        /// <summary>
        /// Specifies the name of the storage queue where the Event Subscriptio will receive events.
        /// </summary>
        public readonly string QueueName;
        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private EventSubscriptionStorageQueueEndpoint(
            string queueName,
            string storageAccountId)
        {
            QueueName = queueName;
            StorageAccountId = storageAccountId;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionSubjectFilter
    {
        /// <summary>
        /// Specifies if `subject_begins_with` and `subject_ends_with` case sensitive. This value defaults to `false`.
        /// </summary>
        public readonly bool? CaseSensitive;
        /// <summary>
        /// A string to filter events for an event subscription based on a resource path prefix.
        /// </summary>
        public readonly string? SubjectBeginsWith;
        /// <summary>
        /// A string to filter events for an event subscription based on a resource path suffix.
        /// </summary>
        public readonly string? SubjectEndsWith;

        [OutputConstructor]
        private EventSubscriptionSubjectFilter(
            bool? caseSensitive,
            string? subjectBeginsWith,
            string? subjectEndsWith)
        {
            CaseSensitive = caseSensitive;
            SubjectBeginsWith = subjectBeginsWith;
            SubjectEndsWith = subjectEndsWith;
        }
    }

    [OutputType]
    public sealed class EventSubscriptionWebhookEndpoint
    {
        /// <summary>
        /// Specifies the url of the webhook where the Event Subscription will receive events.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private EventSubscriptionWebhookEndpoint(string url)
        {
            Url = url;
        }
    }
    }
}
