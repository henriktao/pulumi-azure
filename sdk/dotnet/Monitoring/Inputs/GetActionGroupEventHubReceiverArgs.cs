// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class GetActionGroupEventHubReceiverInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the respective Event Hub.
        /// </summary>
        [Input("eventHubId", required: true)]
        public Input<string> EventHubId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Action Group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Tenant ID for the subscription containing this Event Hub.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        /// <summary>
        /// Indicates whether to use common alert schema.
        /// </summary>
        [Input("useCommonAlertSchema")]
        public Input<bool>? UseCommonAlertSchema { get; set; }

        public GetActionGroupEventHubReceiverInputArgs()
        {
        }
    }
}
