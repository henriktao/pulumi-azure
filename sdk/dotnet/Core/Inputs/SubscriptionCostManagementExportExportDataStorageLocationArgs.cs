// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core.Inputs
{

    public sealed class SubscriptionCostManagementExportExportDataStorageLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Resource Manager ID of the container where exports will be uploaded.
        /// </summary>
        [Input("containerId", required: true)]
        public Input<string> ContainerId { get; set; } = null!;

        /// <summary>
        /// The path of the directory where exports will be uploaded.
        /// </summary>
        [Input("rootFolderPath", required: true)]
        public Input<string> RootFolderPath { get; set; } = null!;

        public SubscriptionCostManagementExportExportDataStorageLocationArgs()
        {
        }
    }
}
