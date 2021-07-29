// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DomainServices.Outputs
{

    [OutputType]
    public sealed class ServiceInitialReplicaSet
    {
        /// <summary>
        /// A list of subnet IP addresses for the domain controllers in the initial replica set, typically two.
        /// </summary>
        public readonly ImmutableArray<string> DomainControllerIpAddresses;
        /// <summary>
        /// The publicly routable IP address for the domain controllers in the initial replica set.
        /// </summary>
        public readonly string? ExternalAccessIpAddress;
        /// <summary>
        /// The ID of the Domain Service.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The current service status for the initial replica set.
        /// </summary>
        public readonly string? ServiceStatus;
        /// <summary>
        /// The ID of the subnet in which to place the initial replica set.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private ServiceInitialReplicaSet(
            ImmutableArray<string> domainControllerIpAddresses,

            string? externalAccessIpAddress,

            string? id,

            string? location,

            string? serviceStatus,

            string subnetId)
        {
            DomainControllerIpAddresses = domainControllerIpAddresses;
            ExternalAccessIpAddress = externalAccessIpAddress;
            Id = id;
            Location = location;
            ServiceStatus = serviceStatus;
            SubnetId = subnetId;
        }
    }
}
