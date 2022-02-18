// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class KubernetesClusterAddonProfileAciConnectorLinux
    {
        /// <summary>
        /// Is the virtual node addon enabled? This field is deprecated and will be removed in version 3.0 of the AzureRM Provider.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The subnet name for the virtual nodes to run. This is required when `aci_connector_linux` `enabled` argument is set to `true`.
        /// </summary>
        public readonly string? SubnetName;

        [OutputConstructor]
        private KubernetesClusterAddonProfileAciConnectorLinux(
            bool enabled,

            string? subnetName)
        {
            Enabled = enabled;
            SubnetName = subnetName;
        }
    }
}
