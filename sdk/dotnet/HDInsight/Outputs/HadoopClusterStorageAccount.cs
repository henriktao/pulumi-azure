// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class HadoopClusterStorageAccount
    {
        /// <summary>
        /// Is this the Default Storage Account for the HDInsight Hadoop Cluster? Changing this forces a new resource to be created.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// The Access Key which should be used to connect to the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string StorageAccountKey;
        /// <summary>
        /// The ID of the Storage Container. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string StorageContainerId;
        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? StorageResourceId;

        [OutputConstructor]
        private HadoopClusterStorageAccount(
            bool isDefault,

            string storageAccountKey,

            string storageContainerId,

            string? storageResourceId)
        {
            IsDefault = isDefault;
            StorageAccountKey = storageAccountKey;
            StorageContainerId = storageContainerId;
            StorageResourceId = storageResourceId;
        }
    }
}
