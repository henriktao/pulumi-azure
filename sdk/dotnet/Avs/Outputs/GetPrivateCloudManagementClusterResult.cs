// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Avs.Outputs
{

    [OutputType]
    public sealed class GetPrivateCloudManagementClusterResult
    {
        /// <summary>
        /// The list of the hosts in the management cluster.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// The ID of the management cluster.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The size of the management cluster.
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private GetPrivateCloudManagementClusterResult(
            ImmutableArray<string> hosts,

            int id,

            int size)
        {
            Hosts = hosts;
            Id = id;
            Size = size;
        }
    }
}
