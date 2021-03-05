// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class VirtualNetworkGatewayBgpSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Autonomous System Number (ASN) to use as part of the BGP.
        /// </summary>
        [Input("asn")]
        public Input<int>? Asn { get; set; }

        /// <summary>
        /// The weight added to routes which have been learned
        /// through BGP peering. Valid values can be between `0` and `100`.
        /// </summary>
        [Input("peerWeight")]
        public Input<int>? PeerWeight { get; set; }

        [Input("peeringAddress")]
        public Input<string>? PeeringAddress { get; set; }

        [Input("peeringAddresses")]
        private InputList<Inputs.VirtualNetworkGatewayBgpSettingsPeeringAddressArgs>? _peeringAddresses;

        /// <summary>
        /// A list of `peering_addresses` as defined below. Only one `peering_addresses` block can be specified except when `active_active` of this Virtual Network Gateway is `true`.
        /// </summary>
        public InputList<Inputs.VirtualNetworkGatewayBgpSettingsPeeringAddressArgs> PeeringAddresses
        {
            get => _peeringAddresses ?? (_peeringAddresses = new InputList<Inputs.VirtualNetworkGatewayBgpSettingsPeeringAddressArgs>());
            set => _peeringAddresses = value;
        }

        public VirtualNetworkGatewayBgpSettingsArgs()
        {
        }
    }
}
