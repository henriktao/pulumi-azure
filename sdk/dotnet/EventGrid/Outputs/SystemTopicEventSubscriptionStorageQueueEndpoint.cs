// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Outputs
{

    [OutputType]
    public sealed class SystemTopicEventSubscriptionStorageQueueEndpoint
    {
        /// <summary>
        /// Storage queue message time to live in seconds.
        /// </summary>
        public readonly int? QueueMessageTimeToLiveInSeconds;
        /// <summary>
        /// Specifies the name of the storage queue where the Event Subscription will receive events.
        /// </summary>
        public readonly string QueueName;
        /// <summary>
        /// Specifies the id of the storage account id where the storage queue is located.
        /// </summary>
        public readonly string StorageAccountId;

        [OutputConstructor]
        private SystemTopicEventSubscriptionStorageQueueEndpoint(
            int? queueMessageTimeToLiveInSeconds,

            string queueName,

            string storageAccountId)
        {
            QueueMessageTimeToLiveInSeconds = queueMessageTimeToLiveInSeconds;
            QueueName = queueName;
            StorageAccountId = storageAccountId;
        }
    }
}
