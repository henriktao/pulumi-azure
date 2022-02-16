// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class GetServiceAdditionalLocationResult
    {
        /// <summary>
        /// Specifies the number of units associated with this API Management service.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// Gateway URL of the API Management service in the Region.
        /// </summary>
        public readonly string GatewayRegionalUrl;
        /// <summary>
        /// The location name of the additional region among Azure Data center regions.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Private IP addresses of the API Management service in the additional location, for instances using virtual network mode.
        /// </summary>
        public readonly ImmutableArray<string> PrivateIpAddresses;
        /// <summary>
        /// ID of the standard SKU IPv4 Public IP. Available only for Premium SKU deployed in a virtual network.
        /// </summary>
        public readonly string PublicIpAddressId;
        /// <summary>
        /// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        /// </summary>
        public readonly ImmutableArray<string> PublicIpAddresses;
        /// <summary>
        /// List of the availability zones where API Management is deployed in the additional region exists.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetServiceAdditionalLocationResult(
            int? capacity,

            string gatewayRegionalUrl,

            string location,

            ImmutableArray<string> privateIpAddresses,

            string publicIpAddressId,

            ImmutableArray<string> publicIpAddresses,

            ImmutableArray<string> zones)
        {
            Capacity = capacity;
            GatewayRegionalUrl = gatewayRegionalUrl;
            Location = location;
            PrivateIpAddresses = privateIpAddresses;
            PublicIpAddressId = publicIpAddressId;
            PublicIpAddresses = publicIpAddresses;
            Zones = zones;
        }
    }
}
