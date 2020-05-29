// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).
func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("azure:containerservice/getKubernetesCluster:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterArgs struct {
	// The name of the managed Kubernetes Cluster.
	Name string `pulumi:"name"`
	// If the cluster has the Kubernetes API only exposed on internal IP addresses.
	PrivateClusterEnabled *bool `pulumi:"privateClusterEnabled"`
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// The name of the Resource Group in which the managed Kubernetes Cluster exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResult struct {
	// A `addonProfile` block as documented below.
	AddonProfiles []GetKubernetesClusterAddonProfile `pulumi:"addonProfiles"`
	// An `agentPoolProfile` block as documented below.
	AgentPoolProfiles []GetKubernetesClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// The DNS Prefix of the managed Kubernetes cluster.
	DnsPrefix string `pulumi:"dnsPrefix"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn string `pulumi:"fqdn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A `identity` block as documented below.
	Identities []GetKubernetesClusterIdentity `pulumi:"identities"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw string `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs []GetKubernetesClusterKubeAdminConfig `pulumi:"kubeAdminConfigs"`
	// Base64 encoded Kubernetes configuration.
	KubeConfigRaw string `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs []GetKubernetesClusterKubeConfig `pulumi:"kubeConfigs"`
	// A `kubeletIdentity` block as documented below.
	KubeletIdentities []GetKubernetesClusterKubeletIdentity `pulumi:"kubeletIdentities"`
	// The version of Kubernetes used on the managed Kubernetes Cluster.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as documented below.
	LinuxProfiles []GetKubernetesClusterLinuxProfile `pulumi:"linuxProfiles"`
	// The Azure Region in which the managed Kubernetes Cluster exists.
	Location string `pulumi:"location"`
	// The name assigned to this pool of agents.
	Name string `pulumi:"name"`
	// A `networkProfile` block as documented below.
	NetworkProfiles []GetKubernetesClusterNetworkProfile `pulumi:"networkProfiles"`
	// Auto-generated Resource Group containing AKS Cluster resources.
	NodeResourceGroup string `pulumi:"nodeResourceGroup"`
	// If the cluster has the Kubernetes API only exposed on internal IP addresses.
	PrivateClusterEnabled bool `pulumi:"privateClusterEnabled"`
	// The FQDN of this Kubernetes Cluster when private link has been enabled. This name is only resolvable inside the Virtual Network where the Azure Kubernetes Service is located
	PrivateFqdn string `pulumi:"privateFqdn"`
	// Deprecated: Deprecated in favor of `private_cluster_enabled`
	PrivateLinkEnabled bool   `pulumi:"privateLinkEnabled"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block as documented below.
	RoleBasedAccessControls []GetKubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControls"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipals []GetKubernetesClusterServicePrincipal `pulumi:"servicePrincipals"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as documented below.
	WindowsProfiles []GetKubernetesClusterWindowsProfile `pulumi:"windowsProfiles"`
}
