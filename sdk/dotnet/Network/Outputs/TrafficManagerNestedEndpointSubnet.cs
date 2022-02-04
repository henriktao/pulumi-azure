// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class TrafficManagerNestedEndpointSubnet
    {
        /// <summary>
        /// The first IP Address in this subnet.
        /// </summary>
        public readonly string First;
        /// <summary>
        /// The last IP Address in this subnet.
        /// </summary>
        public readonly string? Last;
        /// <summary>
        /// The block size (number of leading bits in the subnet mask).
        /// </summary>
        public readonly int? Scope;

        [OutputConstructor]
        private TrafficManagerNestedEndpointSubnet(
            string first,

            string? last,

            int? scope)
        {
            First = first;
            Last = last;
            Scope = scope;
        }
    }
}
