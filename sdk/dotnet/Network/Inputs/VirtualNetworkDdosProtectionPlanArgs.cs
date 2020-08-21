// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkDdosProtectionPlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable DDoS Protection Plan on Virtual Network.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// The ID of DDoS Protection Plan.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public VirtualNetworkDdosProtectionPlanArgs()
        {
        }
    }
}
