// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class OrchestratedVirtualMachineScaleSetNetworkInterface
    {
        public readonly ImmutableArray<string> DnsServers;
        public readonly bool? EnableAcceleratedNetworking;
        public readonly bool? EnableIpForwarding;
        public readonly ImmutableArray<Outputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfiguration> IpConfigurations;
        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        public readonly string? NetworkSecurityGroupId;
        public readonly bool? Primary;

        [OutputConstructor]
        private OrchestratedVirtualMachineScaleSetNetworkInterface(
            ImmutableArray<string> dnsServers,

            bool? enableAcceleratedNetworking,

            bool? enableIpForwarding,

            ImmutableArray<Outputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfiguration> ipConfigurations,

            string name,

            string? networkSecurityGroupId,

            bool? primary)
        {
            DnsServers = dnsServers;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableIpForwarding = enableIpForwarding;
            IpConfigurations = ipConfigurations;
            Name = name;
            NetworkSecurityGroupId = networkSecurityGroupId;
            Primary = primary;
        }
    }
}
