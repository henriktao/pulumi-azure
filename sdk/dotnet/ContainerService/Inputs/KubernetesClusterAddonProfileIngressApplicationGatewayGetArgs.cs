// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterAddonProfileIngressApplicationGatewayGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Application Gateway associated with the ingress controller deployed to this Kubernetes Cluster.
        /// </summary>
        [Input("effectiveGatewayId")]
        public Input<string>? EffectiveGatewayId { get; set; }

        /// <summary>
        /// Is the Kubernetes Dashboard enabled?
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The ID of the Application Gateway to integrate with the ingress controller of this Kubernetes Cluster. See [this](https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ingress-controller-add-on-existing) page for further details.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// The name of the Application Gateway to be used or created in the Nodepool Resource Group, which in turn will be integrated with the ingress controller of this Kubernetes Cluster. See [this](https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ingress-controller-add-on-new) page for further details.
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        [Input("ingressApplicationGatewayIdentities")]
        private InputList<Inputs.KubernetesClusterAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentityGetArgs>? _ingressApplicationGatewayIdentities;

        /// <summary>
        /// An `ingress_application_gateway_identity` block is exported. The exported attributes are defined below.
        /// </summary>
        [Obsolete(@"`addon_profile.0.ingress_application_gateway.0.ingress_application_gateway_identity` has been deprecated in favour of `ingress_application_gateway.0.ingress_application_gateway_identity` and will be removed in version 3.0 of the AzureRM Provider.")]
        public InputList<Inputs.KubernetesClusterAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentityGetArgs> IngressApplicationGatewayIdentities
        {
            get => _ingressApplicationGatewayIdentities ?? (_ingressApplicationGatewayIdentities = new InputList<Inputs.KubernetesClusterAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentityGetArgs>());
            set => _ingressApplicationGatewayIdentities = value;
        }

        /// <summary>
        /// The subnet CIDR to be used to create an Application Gateway, which in turn will be integrated with the ingress controller of this Kubernetes Cluster. See [this](https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ingress-controller-add-on-new) page for further details.
        /// </summary>
        [Input("subnetCidr")]
        public Input<string>? SubnetCidr { get; set; }

        /// <summary>
        /// The ID of the subnet on which to create an Application Gateway, which in turn will be integrated with the ingress controller of this Kubernetes Cluster. See [this](https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ingress-controller-add-on-new) page for further details.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public KubernetesClusterAddonProfileIngressApplicationGatewayGetArgs()
        {
        }
    }
}
