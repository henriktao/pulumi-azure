// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Inputs
{

    public sealed class EventSubscriptionStorageBlobDeadLetterDestinationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the storage account id where the storage blob is located.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Storage blob container that is the destination of the deadletter events.
        /// </summary>
        [Input("storageBlobContainerName", required: true)]
        public Input<string> StorageBlobContainerName { get; set; } = null!;

        public EventSubscriptionStorageBlobDeadLetterDestinationGetArgs()
        {
        }
    }
}
