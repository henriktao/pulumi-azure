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
    public sealed class TrafficManagerExternalEndpointCustomHeader
    {
        /// <summary>
        /// The name of the custom header.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of custom header. Applicable for Http and Https protocol.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private TrafficManagerExternalEndpointCustomHeader(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
