// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretArgs : Pulumi.ResourceArgs
    {
        [Input("certificates", required: true)]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateArgs>? _certificates;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretCertificateArgs>());
            set => _certificates = value;
        }

        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        public OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretArgs()
        {
        }
    }
}
