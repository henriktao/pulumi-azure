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
    public sealed class OrchestratedVirtualMachineScaleSetTerminationNotification
    {
        public readonly bool Enabled;
        public readonly string? Timeout;

        [OutputConstructor]
        private OrchestratedVirtualMachineScaleSetTerminationNotification(
            bool enabled,

            string? timeout)
        {
            Enabled = enabled;
            Timeout = timeout;
        }
    }
}
