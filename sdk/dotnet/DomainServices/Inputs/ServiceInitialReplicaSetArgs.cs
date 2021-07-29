// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DomainServices.Inputs
{

    public sealed class ServiceInitialReplicaSetArgs : Pulumi.ResourceArgs
    {
        [Input("domainControllerIpAddresses")]
        private InputList<string>? _domainControllerIpAddresses;

        /// <summary>
        /// A list of subnet IP addresses for the domain controllers in the initial replica set, typically two.
        /// </summary>
        public InputList<string> DomainControllerIpAddresses
        {
            get => _domainControllerIpAddresses ?? (_domainControllerIpAddresses = new InputList<string>());
            set => _domainControllerIpAddresses = value;
        }

        /// <summary>
        /// The publicly routable IP address for the domain controllers in the initial replica set.
        /// </summary>
        [Input("externalAccessIpAddress")]
        public Input<string>? ExternalAccessIpAddress { get; set; }

        /// <summary>
        /// The ID of the Domain Service.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The current service status for the initial replica set.
        /// </summary>
        [Input("serviceStatus")]
        public Input<string>? ServiceStatus { get; set; }

        /// <summary>
        /// The ID of the subnet in which to place the initial replica set.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ServiceInitialReplicaSetArgs()
        {
        }
    }
}
