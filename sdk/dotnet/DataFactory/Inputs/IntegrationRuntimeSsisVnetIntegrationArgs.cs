// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class IntegrationRuntimeSsisVnetIntegrationArgs : Pulumi.ResourceArgs
    {
        [Input("publicIps")]
        private InputList<string>? _publicIps;

        /// <summary>
        /// Static public IP addresses for the Azure-SSIS Integration Runtime. The size must be 2.
        /// </summary>
        public InputList<string> PublicIps
        {
            get => _publicIps ?? (_publicIps = new InputList<string>());
            set => _publicIps = value;
        }

        /// <summary>
        /// Name of the subnet to which the nodes of the Azure-SSIS Integration Runtime will be added.
        /// </summary>
        [Input("subnetName", required: true)]
        public Input<string> SubnetName { get; set; } = null!;

        /// <summary>
        /// ID of the virtual network to which the nodes of the Azure-SSIS Integration Runtime will be added.
        /// </summary>
        [Input("vnetId", required: true)]
        public Input<string> VnetId { get; set; } = null!;

        public IntegrationRuntimeSsisVnetIntegrationArgs()
        {
        }
    }
}
