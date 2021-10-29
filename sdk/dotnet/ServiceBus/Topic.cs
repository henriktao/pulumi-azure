// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    /// <summary>
    /// Manages a ServiceBus Topic.
    /// 
    /// **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleNamespace = new Azure.ServiceBus.Namespace("exampleNamespace", new Azure.ServiceBus.NamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Tags = 
    ///             {
    ///                 { "source", "example" },
    ///             },
    ///         });
    ///         var exampleTopic = new Azure.ServiceBus.Topic("exampleTopic", new Azure.ServiceBus.TopicArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             NamespaceName = exampleNamespace.Name,
    ///             EnablePartitioning = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service Bus Topics can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:servicebus/topic:Topic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1
    /// ```
    /// </summary>
    [AzureResourceType("azure:servicebus/topic:Topic")]
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Topic is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
        /// TTL value is set on the message itself.
        /// </summary>
        [Output("defaultMessageTtl")]
        public Output<string> DefaultMessageTtl { get; private set; } = null!;

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        /// </summary>
        [Output("duplicateDetectionHistoryTimeWindow")]
        public Output<string> DuplicateDetectionHistoryTimeWindow { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls if server-side
        /// batched operations are enabled. Defaults to false.
        /// </summary>
        [Output("enableBatchedOperations")]
        public Output<bool?> EnableBatchedOperations { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express topic holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to false.
        /// </summary>
        [Output("enableExpress")]
        public Output<bool?> EnableExpress { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the topic to be partitioned across multiple message brokers. Defaults to false.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("enablePartitioning")]
        public Output<bool?> EnablePartitioning { get; private set; } = null!;

        /// <summary>
        /// Integer value which controls the maximum size of
        /// a message allowed on the topic for Premium SKU. For supported values see the "Large messages support"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
        /// </summary>
        [Output("maxMessageSizeInKilobytes")]
        public Output<int> MaxMessageSizeInKilobytes { get; private set; } = null!;

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the topic. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Output("maxSizeInMegabytes")]
        public Output<int> MaxSizeInMegabytes { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Topic resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this topic in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether
        /// the Topic requires duplicate detection. Defaults to false. Changing this forces
        /// a new resource to be created.
        /// </summary>
        [Output("requiresDuplicateDetection")]
        public Output<bool?> RequiresDuplicateDetection { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether the Topic
        /// supports ordering. Defaults to false.
        /// </summary>
        [Output("supportOrdering")]
        public Output<bool?> SupportOrdering { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("azure:servicebus/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("azure:servicebus/topic:Topic", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure:eventhub/topic:Topic"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Topic is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
        /// TTL value is set on the message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// Boolean flag which controls if server-side
        /// batched operations are enabled. Defaults to false.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express topic holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to false.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the topic to be partitioned across multiple message brokers. Defaults to false.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// Integer value which controls the maximum size of
        /// a message allowed on the topic for Premium SKU. For supported values see the "Large messages support"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
        /// </summary>
        [Input("maxMessageSizeInKilobytes")]
        public Input<int>? MaxMessageSizeInKilobytes { get; set; }

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the topic. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Topic resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this topic in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Boolean flag which controls whether
        /// the Topic requires duplicate detection. Defaults to false. Changing this forces
        /// a new resource to be created.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Topic
        /// supports ordering. Defaults to false.
        /// </summary>
        [Input("supportOrdering")]
        public Input<bool>? SupportOrdering { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the
        /// Topic is automatically deleted, minimum of 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
        /// TTL value is set on the message itself.
        /// </summary>
        [Input("defaultMessageTtl")]
        public Input<string>? DefaultMessageTtl { get; set; }

        /// <summary>
        /// The ISO 8601 timespan duration during which
        /// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// Boolean flag which controls if server-side
        /// batched operations are enabled. Defaults to false.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// Boolean flag which controls whether Express Entities
        /// are enabled. An express topic holds a message in memory temporarily before writing
        /// it to persistent storage. Defaults to false.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// Boolean flag which controls whether to enable
        /// the topic to be partitioned across multiple message brokers. Defaults to false.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// Integer value which controls the maximum size of
        /// a message allowed on the topic for Premium SKU. For supported values see the "Large messages support"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-premium-messaging#large-messages-support-preview).
        /// </summary>
        [Input("maxMessageSizeInKilobytes")]
        public Input<int>? MaxMessageSizeInKilobytes { get; set; }

        /// <summary>
        /// Integer value which controls the size of
        /// memory allocated for the topic. For supported values see the "Queue/topic size"
        /// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Topic resource. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ServiceBus Namespace to create
        /// this topic in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// Boolean flag which controls whether
        /// the Topic requires duplicate detection. Defaults to false. Changing this forces
        /// a new resource to be created.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Boolean flag which controls whether the Topic
        /// supports ordering. Defaults to false.
        /// </summary>
        [Input("supportOrdering")]
        public Input<bool>? SupportOrdering { get; set; }

        public TopicState()
        {
        }
    }
}
