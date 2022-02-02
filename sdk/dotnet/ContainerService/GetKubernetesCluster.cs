// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService
{
    public static class GetKubernetesCluster
    {
        /// <summary>
        /// Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.ContainerService.GetKubernetesCluster.InvokeAsync(new Azure.ContainerService.GetKubernetesClusterArgs
        ///         {
        ///             Name = "myakscluster",
        ///             ResourceGroupName = "my-example-resource-group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKubernetesClusterResult> InvokeAsync(GetKubernetesClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesClusterResult>("azure:containerservice/getKubernetesCluster:getKubernetesCluster", args ?? new GetKubernetesClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.ContainerService.GetKubernetesCluster.InvokeAsync(new Azure.ContainerService.GetKubernetesClusterArgs
        ///         {
        ///             Name = "myakscluster",
        ///             ResourceGroupName = "my-example-resource-group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetKubernetesClusterResult> Invoke(GetKubernetesClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKubernetesClusterResult>("azure:containerservice/getKubernetesCluster:getKubernetesCluster", args ?? new GetKubernetesClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed Kubernetes Cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the managed Kubernetes Cluster exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetKubernetesClusterArgs()
        {
        }
    }

    public sealed class GetKubernetesClusterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed Kubernetes Cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the managed Kubernetes Cluster exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetKubernetesClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKubernetesClusterResult
    {
        /// <summary>
        /// A `addon_profile` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterAddonProfileResult> AddonProfiles;
        /// <summary>
        /// An `agent_pool_profile` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterAgentPoolProfileResult> AgentPoolProfiles;
        /// <summary>
        /// The IP ranges to whitelist for incoming traffic to the primaries.
        /// </summary>
        public readonly ImmutableArray<string> ApiServerAuthorizedIpRanges;
        /// <summary>
        /// The ID of the Disk Encryption Set used for the Nodes and Volumes.
        /// </summary>
        public readonly string DiskEncryptionSetId;
        /// <summary>
        /// The DNS Prefix of the managed Kubernetes cluster.
        /// </summary>
        public readonly string DnsPrefix;
        /// <summary>
        /// The FQDN of the Azure Kubernetes Managed Cluster.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `identity` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterIdentityResult> Identities;
        /// <summary>
        /// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled and local accounts are not disabled.
        /// </summary>
        public readonly string KubeAdminConfigRaw;
        /// <summary>
        /// A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled and local accounts are not disabled.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterKubeAdminConfigResult> KubeAdminConfigs;
        /// <summary>
        /// Base64 encoded Kubernetes configuration.
        /// </summary>
        public readonly string KubeConfigRaw;
        /// <summary>
        /// A `kube_config` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterKubeConfigResult> KubeConfigs;
        /// <summary>
        /// A `kubelet_identity` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterKubeletIdentityResult> KubeletIdentities;
        /// <summary>
        /// The version of Kubernetes used on the managed Kubernetes Cluster.
        /// </summary>
        public readonly string KubernetesVersion;
        /// <summary>
        /// A `linux_profile` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterLinuxProfileResult> LinuxProfiles;
        /// <summary>
        /// The Azure Region in which the managed Kubernetes Cluster exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name assigned to this pool of agents.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `network_profile` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterNetworkProfileResult> NetworkProfiles;
        /// <summary>
        /// Auto-generated Resource Group containing AKS Cluster resources.
        /// </summary>
        public readonly string NodeResourceGroup;
        /// <summary>
        /// If the cluster has the Kubernetes API only exposed on internal IP addresses.
        /// </summary>
        public readonly bool PrivateClusterEnabled;
        /// <summary>
        /// The FQDN of this Kubernetes Cluster when private link has been enabled. This name is only resolvable inside the Virtual Network where the Azure Kubernetes Service is located
        /// </summary>
        public readonly string PrivateFqdn;
        public readonly bool PrivateLinkEnabled;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `role_based_access_control` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterRoleBasedAccessControlResult> RoleBasedAccessControls;
        /// <summary>
        /// A `service_principal` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterServicePrincipalResult> ServicePrincipals;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// A `windows_profile` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterWindowsProfileResult> WindowsProfiles;

        [OutputConstructor]
        private GetKubernetesClusterResult(
            ImmutableArray<Outputs.GetKubernetesClusterAddonProfileResult> addonProfiles,

            ImmutableArray<Outputs.GetKubernetesClusterAgentPoolProfileResult> agentPoolProfiles,

            ImmutableArray<string> apiServerAuthorizedIpRanges,

            string diskEncryptionSetId,

            string dnsPrefix,

            string fqdn,

            string id,

            ImmutableArray<Outputs.GetKubernetesClusterIdentityResult> identities,

            string kubeAdminConfigRaw,

            ImmutableArray<Outputs.GetKubernetesClusterKubeAdminConfigResult> kubeAdminConfigs,

            string kubeConfigRaw,

            ImmutableArray<Outputs.GetKubernetesClusterKubeConfigResult> kubeConfigs,

            ImmutableArray<Outputs.GetKubernetesClusterKubeletIdentityResult> kubeletIdentities,

            string kubernetesVersion,

            ImmutableArray<Outputs.GetKubernetesClusterLinuxProfileResult> linuxProfiles,

            string location,

            string name,

            ImmutableArray<Outputs.GetKubernetesClusterNetworkProfileResult> networkProfiles,

            string nodeResourceGroup,

            bool privateClusterEnabled,

            string privateFqdn,

            bool privateLinkEnabled,

            string resourceGroupName,

            ImmutableArray<Outputs.GetKubernetesClusterRoleBasedAccessControlResult> roleBasedAccessControls,

            ImmutableArray<Outputs.GetKubernetesClusterServicePrincipalResult> servicePrincipals,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetKubernetesClusterWindowsProfileResult> windowsProfiles)
        {
            AddonProfiles = addonProfiles;
            AgentPoolProfiles = agentPoolProfiles;
            ApiServerAuthorizedIpRanges = apiServerAuthorizedIpRanges;
            DiskEncryptionSetId = diskEncryptionSetId;
            DnsPrefix = dnsPrefix;
            Fqdn = fqdn;
            Id = id;
            Identities = identities;
            KubeAdminConfigRaw = kubeAdminConfigRaw;
            KubeAdminConfigs = kubeAdminConfigs;
            KubeConfigRaw = kubeConfigRaw;
            KubeConfigs = kubeConfigs;
            KubeletIdentities = kubeletIdentities;
            KubernetesVersion = kubernetesVersion;
            LinuxProfiles = linuxProfiles;
            Location = location;
            Name = name;
            NetworkProfiles = networkProfiles;
            NodeResourceGroup = nodeResourceGroup;
            PrivateClusterEnabled = privateClusterEnabled;
            PrivateFqdn = privateFqdn;
            PrivateLinkEnabled = privateLinkEnabled;
            ResourceGroupName = resourceGroupName;
            RoleBasedAccessControls = roleBasedAccessControls;
            ServicePrincipals = servicePrincipals;
            Tags = tags;
            WindowsProfiles = windowsProfiles;
        }
    }
}
