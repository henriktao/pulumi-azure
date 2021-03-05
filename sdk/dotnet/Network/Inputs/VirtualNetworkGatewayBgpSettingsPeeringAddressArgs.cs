// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkGatewayBgpSettingsPeeringAddressArgs : Pulumi.ResourceArgs
    {
        [Input("apipaAddresses")]
        private InputList<string>? _apipaAddresses;

        /// <summary>
        /// A list of Azure custom APIPA addresses assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public InputList<string> ApipaAddresses
        {
            get => _apipaAddresses ?? (_apipaAddresses = new InputList<string>());
            set => _apipaAddresses = value;
        }

        [Input("defaultAddresses")]
        private InputList<string>? _defaultAddresses;

        /// <summary>
        /// A list of peering address assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public InputList<string> DefaultAddresses
        {
            get => _defaultAddresses ?? (_defaultAddresses = new InputList<string>());
            set => _defaultAddresses = value;
        }

        /// <summary>
        /// The name of the IP configuration of this Virtual Network Gateway. In case there are multiple `ip_configuration` blocks defined, this property is **required** to specify.
        /// </summary>
        [Input("ipConfigurationName")]
        public Input<string>? IpConfigurationName { get; set; }

        [Input("tunnelIpAddresses")]
        private InputList<string>? _tunnelIpAddresses;

        /// <summary>
        /// A list of tunnel IP addresses assigned to the BGP peer of the Virtual Network Gateway.
        /// </summary>
        public InputList<string> TunnelIpAddresses
        {
            get => _tunnelIpAddresses ?? (_tunnelIpAddresses = new InputList<string>());
            set => _tunnelIpAddresses = value;
        }

        public VirtualNetworkGatewayBgpSettingsPeeringAddressArgs()
        {
        }
    }
}
