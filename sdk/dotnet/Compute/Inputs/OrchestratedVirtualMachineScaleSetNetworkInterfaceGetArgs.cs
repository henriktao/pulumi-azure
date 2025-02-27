// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class OrchestratedVirtualMachineScaleSetNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        [Input("enableIpForwarding")]
        public Input<bool>? EnableIpForwarding { get; set; }

        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs>? _ipConfigurations;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationGetArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkSecurityGroupId")]
        public Input<string>? NetworkSecurityGroupId { get; set; }

        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        public OrchestratedVirtualMachineScaleSetNetworkInterfaceGetArgs()
        {
        }
    }
}
