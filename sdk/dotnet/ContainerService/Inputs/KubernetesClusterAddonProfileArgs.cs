// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterAddonProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `aci_connector_linux` block as defined below. For more details, please visit [Create and configure an AKS cluster to use virtual nodes](https://docs.microsoft.com/en-us/azure/aks/virtual-nodes-portal).
        /// </summary>
        [Input("aciConnectorLinux")]
        public Input<Inputs.KubernetesClusterAddonProfileAciConnectorLinuxArgs>? AciConnectorLinux { get; set; }

        [Input("azureKeyvaultSecretsProvider")]
        public Input<Inputs.KubernetesClusterAddonProfileAzureKeyvaultSecretsProviderArgs>? AzureKeyvaultSecretsProvider { get; set; }

        [Input("azurePolicy")]
        public Input<Inputs.KubernetesClusterAddonProfileAzurePolicyArgs>? AzurePolicy { get; set; }

        /// <summary>
        /// A `http_application_routing` block as defined below.
        /// </summary>
        [Input("httpApplicationRouting")]
        public Input<Inputs.KubernetesClusterAddonProfileHttpApplicationRoutingArgs>? HttpApplicationRouting { get; set; }

        /// <summary>
        /// A `ingress_application_gateway` block as defined below.
        /// </summary>
        [Input("ingressApplicationGateway")]
        public Input<Inputs.KubernetesClusterAddonProfileIngressApplicationGatewayArgs>? IngressApplicationGateway { get; set; }

        [Input("kubeDashboard")]
        public Input<Inputs.KubernetesClusterAddonProfileKubeDashboardArgs>? KubeDashboard { get; set; }

        /// <summary>
        /// A `oms_agent` block as defined below.
        /// </summary>
        [Input("omsAgent")]
        public Input<Inputs.KubernetesClusterAddonProfileOmsAgentArgs>? OmsAgent { get; set; }

        [Input("openServiceMesh")]
        public Input<Inputs.KubernetesClusterAddonProfileOpenServiceMeshArgs>? OpenServiceMesh { get; set; }

        public KubernetesClusterAddonProfileArgs()
        {
        }
    }
}
