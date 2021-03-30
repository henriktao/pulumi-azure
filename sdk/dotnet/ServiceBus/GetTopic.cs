// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceBus
{
    public static class GetTopic
    {
        /// <summary>
        /// Use this data source to access information about an existing Service Bus Topic.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.ServiceBus.GetTopic.InvokeAsync(new Azure.ServiceBus.GetTopicArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///             NamespaceName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicResult> InvokeAsync(GetTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("azure:servicebus/getTopic:getTopic", args ?? new GetTopicArgs(), options.WithVersion());
    }


    public sealed class GetTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Service Bus Topic.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Service Bus Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Service Bus Topic exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetTopicArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicResult
    {
        /// <summary>
        /// The ISO 8601 timespan duration of the idle interval after which the Topic is automatically deleted, minimum of 5 minutes.
        /// </summary>
        public readonly string AutoDeleteOnIdle;
        /// <summary>
        /// The ISO 8601 timespan duration of TTL of messages sent to this topic if no TTL value is set on the message itself.
        /// </summary>
        public readonly string DefaultMessageTtl;
        /// <summary>
        /// The ISO 8601 timespan duration during which duplicates can be detected.
        /// </summary>
        public readonly string DuplicateDetectionHistoryTimeWindow;
        /// <summary>
        /// Boolean flag which controls if server-side batched operations are enabled.
        /// </summary>
        public readonly bool EnableBatchedOperations;
        /// <summary>
        /// Boolean flag which controls whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
        /// </summary>
        public readonly bool EnableExpress;
        /// <summary>
        /// Boolean flag which controls whether to enable the topic to be partitioned across multiple message brokers.
        /// </summary>
        public readonly bool EnablePartitioning;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Integer value which controls the size of memory allocated for the topic. For supported values see the "Queue/topic size" section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
        /// </summary>
        public readonly int MaxSizeInMegabytes;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// Boolean flag which controls whether the Topic requires duplicate detection.
        /// </summary>
        public readonly bool RequiresDuplicateDetection;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Status of the Service Bus Topic. Acceptable values are Active or Disabled.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Boolean flag which controls whether the Topic supports ordering.
        /// </summary>
        public readonly bool SupportOrdering;

        [OutputConstructor]
        private GetTopicResult(
            string autoDeleteOnIdle,

            string defaultMessageTtl,

            string duplicateDetectionHistoryTimeWindow,

            bool enableBatchedOperations,

            bool enableExpress,

            bool enablePartitioning,

            string id,

            int maxSizeInMegabytes,

            string name,

            string namespaceName,

            bool requiresDuplicateDetection,

            string resourceGroupName,

            string status,

            bool supportOrdering)
        {
            AutoDeleteOnIdle = autoDeleteOnIdle;
            DefaultMessageTtl = defaultMessageTtl;
            DuplicateDetectionHistoryTimeWindow = duplicateDetectionHistoryTimeWindow;
            EnableBatchedOperations = enableBatchedOperations;
            EnableExpress = enableExpress;
            EnablePartitioning = enablePartitioning;
            Id = id;
            MaxSizeInMegabytes = maxSizeInMegabytes;
            Name = name;
            NamespaceName = namespaceName;
            RequiresDuplicateDetection = requiresDuplicateDetection;
            ResourceGroupName = resourceGroupName;
            Status = status;
            SupportOrdering = supportOrdering;
        }
    }
}
