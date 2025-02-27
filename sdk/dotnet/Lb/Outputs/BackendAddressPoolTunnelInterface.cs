// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Lb.Outputs
{

    [OutputType]
    public sealed class BackendAddressPoolTunnelInterface
    {
        /// <summary>
        /// The unique identifier of this Gateway Lodbalancer Tunnel Interface.
        /// </summary>
        public readonly int Identifier;
        /// <summary>
        /// The port number that this Gateway Lodbalancer Tunnel Interface listens to.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol used for this Gateway Lodbalancer Tunnel Interface. Possible values are `Native` and `VXLAN`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The traffic type of this Gateway Lodbalancer Tunnel Interface. Possible values are `Internal` and `External`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private BackendAddressPoolTunnelInterface(
            int identifier,

            int port,

            string protocol,

            string type)
        {
            Identifier = identifier;
            Port = port;
            Protocol = protocol;
            Type = type;
        }
    }
}
