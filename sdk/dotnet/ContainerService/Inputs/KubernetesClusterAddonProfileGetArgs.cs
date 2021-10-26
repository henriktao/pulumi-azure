// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterAddonProfileGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `aci_connector_linux` block. For more details, please visit [Create and configure an AKS cluster to use virtual nodes](https://docs.microsoft.com/en-us/azure/aks/virtual-nodes-portal).
        /// </summary>
        [Input("aciConnectorLinux")]
        public Input<Inputs.KubernetesClusterAddonProfileAciConnectorLinuxGetArgs>? AciConnectorLinux { get; set; }

        /// <summary>
        /// A `azure_policy` block as defined below. For more details please visit [Understand Azure Policy for Azure Kubernetes Service](https://docs.microsoft.com/en-ie/azure/governance/policy/concepts/rego-for-aks)
        /// </summary>
        [Input("azurePolicy")]
        public Input<Inputs.KubernetesClusterAddonProfileAzurePolicyGetArgs>? AzurePolicy { get; set; }

        /// <summary>
        /// A `http_application_routing` block as defined below.
        /// </summary>
        [Input("httpApplicationRouting")]
        public Input<Inputs.KubernetesClusterAddonProfileHttpApplicationRoutingGetArgs>? HttpApplicationRouting { get; set; }

        /// <summary>
        /// An `ingress_application_gateway` block as defined below.
        /// </summary>
        [Input("ingressApplicationGateway")]
        public Input<Inputs.KubernetesClusterAddonProfileIngressApplicationGatewayGetArgs>? IngressApplicationGateway { get; set; }

        /// <summary>
        /// A `kube_dashboard` block as defined below.
        /// </summary>
        [Input("kubeDashboard")]
        public Input<Inputs.KubernetesClusterAddonProfileKubeDashboardGetArgs>? KubeDashboard { get; set; }

        /// <summary>
        /// A `oms_agent` block as defined below. For more details, please visit [How to onboard Azure Monitor for containers](https://docs.microsoft.com/en-us/azure/monitoring/monitoring-container-insights-onboard).
        /// </summary>
        [Input("omsAgent")]
        public Input<Inputs.KubernetesClusterAddonProfileOmsAgentGetArgs>? OmsAgent { get; set; }

        /// <summary>
        /// An `open_service_mesh` block as defined below. For more details, please visit [Open Service Mesh for AKS](https://docs.microsoft.com/azure/aks/open-service-mesh-about).
        /// </summary>
        [Input("openServiceMesh")]
        public Input<Inputs.KubernetesClusterAddonProfileOpenServiceMeshGetArgs>? OpenServiceMesh { get; set; }

        public KubernetesClusterAddonProfileGetArgs()
        {
        }
    }
}
