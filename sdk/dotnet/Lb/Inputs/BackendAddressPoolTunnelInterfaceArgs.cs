// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb.Inputs
{

    public sealed class BackendAddressPoolTunnelInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of this Gateway Lodbalancer Tunnel Interface.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<int> Identifier { get; set; } = null!;

        /// <summary>
        /// The port number that this Gateway Lodbalancer Tunnel Interface listens to.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol used for this Gateway Lodbalancer Tunnel Interface. Possible values are `Native` and `VXLAN`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The traffic type of this Gateway Lodbalancer Tunnel Interface. Possible values are `Internal` and `External`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BackendAddressPoolTunnelInterfaceArgs()
        {
        }
    }
}
