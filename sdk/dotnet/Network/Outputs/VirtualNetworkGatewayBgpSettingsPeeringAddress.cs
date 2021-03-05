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
    public sealed class VirtualNetworkGatewayBgpSettingsPeeringAddress
    {
        /// <summary>
        /// A list of Azure custom APIPA addresses assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public readonly ImmutableArray<string> ApipaAddresses;
        /// <summary>
        /// A list of peering address assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public readonly ImmutableArray<string> DefaultAddresses;
        /// <summary>
        /// The name of the IP configuration of this Virtual Network Gateway. In case there are multiple `ip_configuration` blocks defined, this property is **required** to specify.
        /// </summary>
        public readonly string? IpConfigurationName;
        /// <summary>
        /// A list of tunnel IP addresses assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public readonly ImmutableArray<string> TunnelIpAddresses;

        [OutputConstructor]
        private VirtualNetworkGatewayBgpSettingsPeeringAddress(
            ImmutableArray<string> apipaAddresses,

            ImmutableArray<string> defaultAddresses,

            string? ipConfigurationName,

            ImmutableArray<string> tunnelIpAddresses)
        {
            ApipaAddresses = apipaAddresses;
            DefaultAddresses = defaultAddresses;
            IpConfigurationName = ipConfigurationName;
            TunnelIpAddresses = tunnelIpAddresses;
        }
    }
}
