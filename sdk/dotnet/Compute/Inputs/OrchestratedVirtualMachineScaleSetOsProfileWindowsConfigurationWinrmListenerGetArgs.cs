// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerGetArgs : Pulumi.ResourceArgs
    {
        [Input("certificateUrl")]
        public Input<string>? CertificateUrl { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListenerGetArgs()
        {
        }
    }
}
