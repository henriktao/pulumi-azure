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
    public sealed class KubernetesClusterNetworkProfile
    {
        /// <summary>
        /// IP address within the Kubernetes service address range that will be used by cluster service discovery (kube-dns). Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DnsServiceIp;
        /// <summary>
        /// IP address (in CIDR notation) used as the Docker bridge IP address on nodes. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DockerBridgeCidr;
        /// <summary>
        /// A `load_balancer_profile` block. This can only be specified when `load_balancer_sku` is set to `Standard`.
        /// </summary>
        public readonly Outputs.KubernetesClusterNetworkProfileLoadBalancerProfile? LoadBalancerProfile;
        /// <summary>
        /// Specifies the SKU of the Load Balancer used for this Kubernetes Cluster. Possible values are `Basic` and `Standard`. Defaults to `Standard`.
        /// </summary>
        public readonly string? LoadBalancerSku;
        /// <summary>
        /// A `nat_gateway_profile` block. This can only be specified when `load_balancer_sku` is set to `Standard` and `outbound_type` is set to `managedNATGateway` or `userAssignedNATGateway`.
        /// </summary>
        public readonly Outputs.KubernetesClusterNetworkProfileNatGatewayProfile? NatGatewayProfile;
        /// <summary>
        /// Network mode to be used with Azure CNI. Possible values are `bridge` and `transparent`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? NetworkMode;
        /// <summary>
        /// Network plugin to use for networking. Currently supported values are `azure` and `kubenet`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string NetworkPlugin;
        /// <summary>
        /// Sets up network policy to be used with Azure CNI. [Network policy allows us to control the traffic flow between pods](https://docs.microsoft.com/en-us/azure/aks/use-network-policies). Currently supported values are `calico` and `azure`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? NetworkPolicy;
        /// <summary>
        /// The outbound (egress) routing method which should be used for this Kubernetes Cluster. Possible values are `loadBalancer`, `userDefinedRouting`, `managedNATGateway` and `userAssignedNATGateway`. Defaults to `loadBalancer`.
        /// </summary>
        public readonly string? OutboundType;
        /// <summary>
        /// The CIDR to use for pod IP addresses. This field can only be set when `network_plugin` is set to `kubenet`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? PodCidr;
        /// <summary>
        /// The Network Range used by the Kubernetes service. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? ServiceCidr;

        [OutputConstructor]
        private KubernetesClusterNetworkProfile(
            string? dnsServiceIp,

            string? dockerBridgeCidr,

            Outputs.KubernetesClusterNetworkProfileLoadBalancerProfile? loadBalancerProfile,

            string? loadBalancerSku,

            Outputs.KubernetesClusterNetworkProfileNatGatewayProfile? natGatewayProfile,

            string? networkMode,

            string networkPlugin,

            string? networkPolicy,

            string? outboundType,

            string? podCidr,

            string? serviceCidr)
        {
            DnsServiceIp = dnsServiceIp;
            DockerBridgeCidr = dockerBridgeCidr;
            LoadBalancerProfile = loadBalancerProfile;
            LoadBalancerSku = loadBalancerSku;
            NatGatewayProfile = natGatewayProfile;
            NetworkMode = networkMode;
            NetworkPlugin = networkPlugin;
            NetworkPolicy = networkPolicy;
            OutboundType = outboundType;
            PodCidr = podCidr;
            ServiceCidr = serviceCidr;
        }
    }
}
