// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class KafkaClusterStorageAccountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is this the Default Storage Account for the HDInsight Hadoop Cluster? Changing this forces a new resource to be created.
        /// </summary>
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        /// <summary>
        /// The Access Key which should be used to connect to the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        /// <summary>
        /// The ID of the Storage Container. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        /// <summary>
        /// The ID of the Storage Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageResourceId")]
        public Input<string>? StorageResourceId { get; set; }

        public KafkaClusterStorageAccountGetArgs()
        {
        }
    }
}
