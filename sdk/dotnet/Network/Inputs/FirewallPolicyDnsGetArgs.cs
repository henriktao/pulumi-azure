// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyDnsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the network rule fqdn be enabled?
        /// </summary>
        [Input("networkRuleFqdnEnabled")]
        public Input<bool>? NetworkRuleFqdnEnabled { get; set; }

        /// <summary>
        /// Whether to enable DNS proxy on Firewalls attached to this Firewall Policy? Defaults to `false`.
        /// </summary>
        [Input("proxyEnabled")]
        public Input<bool>? ProxyEnabled { get; set; }

        [Input("servers")]
        private InputList<string>? _servers;

        /// <summary>
        /// A list of custom DNS servers' IP addresses.
        /// </summary>
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        public FirewallPolicyDnsGetArgs()
        {
        }
    }
}
