// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package containerservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)
// 
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kubernetes_cluster.html.markdown.
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfileOutput `pulumi:"addonProfile"`
	// One or more `agentPoolProfile` blocks as defined below.
	AgentPoolProfiles KubernetesClusterAgentPoolProfileArrayOutput `pulumi:"agentPoolProfiles"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayOutput `pulumi:"apiServerAuthorizedIpRanges"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolPtrOutput `pulumi:"defaultNodePool"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringOutput `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolOutput `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrOutput `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringOutput `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayOutput `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringOutput `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayOutput `pulumi:"kubeConfigs"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrOutput `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfileOutput `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringOutput `pulumi:"nodeResourceGroup"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringOutput `pulumi:"privateFqdn"`
	PrivateLinkEnabled pulumi.BoolPtrOutput `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlOutput `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalOutput `pulumi:"servicePrincipal"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrOutput `pulumi:"windowsProfile"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil || args.DnsPrefix == nil {
		return nil, errors.New("missing required argument 'DnsPrefix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServicePrincipal == nil {
		return nil, errors.New("missing required argument 'ServicePrincipal'")
	}
	if args == nil {
		args = &KubernetesClusterArgs{}
	}
	var resource KubernetesCluster
	err := ctx.RegisterResource("azure:containerservice/kubernetesCluster:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("azure:containerservice/kubernetesCluster:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
	// A `addonProfile` block as defined below.
	AddonProfile *KubernetesClusterAddonProfile `pulumi:"addonProfile"`
	// One or more `agentPoolProfile` blocks as defined below.
	AgentPoolProfiles []KubernetesClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool *KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn *string `pulumi:"fqdn"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw *string `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs []KubernetesClusterKubeAdminConfig `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw *string `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs []KubernetesClusterKubeConfig `pulumi:"kubeConfigs"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn *string `pulumi:"privateFqdn"`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal *KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

type KubernetesClusterState struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// One or more `agentPoolProfile` blocks as defined below.
	AgentPoolProfiles KubernetesClusterAgentPoolProfileArrayInput
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringPtrInput
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrInput
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringPtrInput
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayInput
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringPtrInput
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringPtrInput
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roleBasedAccessControl` block.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrInput
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// A `addonProfile` block as defined below.
	AddonProfile *KubernetesClusterAddonProfile `pulumi:"addonProfile"`
	// One or more `agentPoolProfile` blocks as defined below.
	AgentPoolProfiles []KubernetesClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool *KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix string `pulumi:"dnsPrefix"`
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// One or more `agentPoolProfile` blocks as defined below.
	AgentPoolProfiles KubernetesClusterAgentPoolProfileArrayInput
	// The IP ranges to whitelist for incoming traffic to the masters.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringInput
	// Whether Pod Security Policies are enabled. Note that this also requires role based access control to be enabled.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// A `identity` block as defined below. Changing this forces a new resource to be created.
	Identity KubernetesClusterIdentityPtrInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roleBasedAccessControl` block.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below.
	ServicePrincipal KubernetesClusterServicePrincipalInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

