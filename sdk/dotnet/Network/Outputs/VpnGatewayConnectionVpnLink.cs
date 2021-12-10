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
    public sealed class VpnGatewayConnectionVpnLink
    {
        /// <summary>
        /// The expected connection bandwidth in MBPS. Defaults to `10`.
        /// </summary>
        public readonly int? BandwidthMbps;
        /// <summary>
        /// Should the BGP be enabled? Defaults to `false`. Changing this forces a new VPN Gateway Connection to be created.
        /// </summary>
        public readonly bool? BgpEnabled;
        /// <summary>
        /// The connection mode of this VPN Link. Possible values are `Default`, `InitiatorOnly` and `ResponderOnly`. Defaults to `Default`.
        /// </summary>
        public readonly string? ConnectionMode;
        /// <summary>
        /// One or more `ipsec_policy` blocks as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnGatewayConnectionVpnLinkIpsecPolicy> IpsecPolicies;
        /// <summary>
        /// Whether to use local azure ip to initiate connection? Defaults to `false`.
        /// </summary>
        public readonly bool? LocalAzureIpAddressEnabled;
        /// <summary>
        /// The name which should be used for this VPN Link Connection.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether to enable policy-based traffic selectors? Defaults to `false`.
        /// </summary>
        public readonly bool? PolicyBasedTrafficSelectorEnabled;
        /// <summary>
        /// The protocol used for this VPN Link Connection. Possible values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Should the rate limit be enabled? Defaults to `false`.
        /// </summary>
        public readonly bool? RatelimitEnabled;
        /// <summary>
        /// Routing weight for this VPN Link Connection. Defaults to `0`.
        /// </summary>
        public readonly int? RouteWeight;
        /// <summary>
        /// SharedKey for this VPN Link Connection.
        /// </summary>
        public readonly string? SharedKey;
        /// <summary>
        /// The ID of the connected VPN Site Link. Changing this forces a new VPN Gateway Connection to be created.
        /// </summary>
        public readonly string VpnSiteLinkId;

        [OutputConstructor]
        private VpnGatewayConnectionVpnLink(
            int? bandwidthMbps,

            bool? bgpEnabled,

            string? connectionMode,

            ImmutableArray<Outputs.VpnGatewayConnectionVpnLinkIpsecPolicy> ipsecPolicies,

            bool? localAzureIpAddressEnabled,

            string name,

            bool? policyBasedTrafficSelectorEnabled,

            string? protocol,

            bool? ratelimitEnabled,

            int? routeWeight,

            string? sharedKey,

            string vpnSiteLinkId)
        {
            BandwidthMbps = bandwidthMbps;
            BgpEnabled = bgpEnabled;
            ConnectionMode = connectionMode;
            IpsecPolicies = ipsecPolicies;
            LocalAzureIpAddressEnabled = localAzureIpAddressEnabled;
            Name = name;
            PolicyBasedTrafficSelectorEnabled = policyBasedTrafficSelectorEnabled;
            Protocol = protocol;
            RatelimitEnabled = ratelimitEnabled;
            RouteWeight = routeWeight;
            SharedKey = sharedKey;
            VpnSiteLinkId = vpnSiteLinkId;
        }
    }
}
