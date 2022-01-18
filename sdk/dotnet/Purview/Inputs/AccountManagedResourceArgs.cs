// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Purview.Inputs
{

    public sealed class AccountManagedResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the managed event hub namespace.
        /// </summary>
        [Input("eventHubNamespaceId")]
        public Input<string>? EventHubNamespaceId { get; set; }

        /// <summary>
        /// The ID of the managed resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the managed storage account.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public AccountManagedResourceArgs()
        {
        }
    }
}
