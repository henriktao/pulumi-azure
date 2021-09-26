// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateLink.Outputs
{

    [OutputType]
    public sealed class GetEndpointConnectionNetworkInterfaceResult
    {
        /// <summary>
        /// The ID of the network interface associated with the private endpoint.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the Name of the private endpoint.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetEndpointConnectionNetworkInterfaceResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
