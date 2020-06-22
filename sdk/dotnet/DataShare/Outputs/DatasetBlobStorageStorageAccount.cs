// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare.Outputs
{

    [OutputType]
    public sealed class DatasetBlobStorageStorageAccount
    {
        /// <summary>
        /// The name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource group name of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
        /// </summary>
        public readonly string ResourceGroupName;
        /// <summary>
        /// The subscription id of the storage account to be shared with the receiver. Changing this forces a new Data Share Blob Storage Dataset to be created.
        /// </summary>
        public readonly string SubscriptionId;

        [OutputConstructor]
        private DatasetBlobStorageStorageAccount(
            string name,

            string resourceGroupName,

            string subscriptionId)
        {
            Name = name;
            ResourceGroupName = resourceGroupName;
            SubscriptionId = subscriptionId;
        }
    }
}
