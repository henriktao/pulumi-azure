// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)
//
// ## Example Usage
//
// This example provisions a basic Managed Kubernetes Cluster.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKubernetesCluster, err := containerservice.NewKubernetesCluster(ctx, "exampleKubernetesCluster", &containerservice.KubernetesClusterArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DnsPrefix:         pulumi.String("exampleaks1"),
// 			DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
// 				Name:      pulumi.String("default"),
// 				NodeCount: pulumi.Int(1),
// 				VmSize:    pulumi.String("Standard_D2_v2"),
// 			},
// 			Identity: &containerservice.KubernetesClusterIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("clientCertificate", exampleKubernetesCluster.KubeConfigs.ApplyT(func(kubeConfigs []containerservice.KubernetesClusterKubeConfig) (string, error) {
// 			return kubeConfigs[0].ClientCertificate, nil
// 		}).(pulumi.StringOutput))
// 		ctx.Export("kubeConfig", exampleKubernetesCluster.KubeConfigRaw)
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Managed Kubernetes Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:containerservice/kubernetesCluster:KubernetesCluster cluster1 /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
// ```
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfileOutput `pulumi:"addonProfile"`
	// The IP ranges to allow for incoming traffic to the server nodes.
	ApiServerAuthorizedIpRanges pulumi.StringArrayOutput `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfileOutput `pulumi:"autoScalerProfile"`
	// The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, `node-image` and `stable`. Omitting this field sets this value to `none`.
	AutomaticChannelUpgrade pulumi.StringPtrOutput `pulumi:"automaticChannelUpgrade"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolOutput `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrOutput `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringPtrOutput `pulumi:"dnsPrefix"`
	// Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
	DnsPrefixPrivateCluster pulumi.StringPtrOutput `pulumi:"dnsPrefixPrivateCluster"`
	EnablePodSecurityPolicy pulumi.BoolPtrOutput   `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
	Identity KubernetesClusterIdentityPtrOutput `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringOutput `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayOutput `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringOutput `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayOutput `pulumi:"kubeConfigs"`
	// A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
	KubeletIdentities KubernetesClusterKubeletIdentityArrayOutput `pulumi:"kubeletIdentities"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrOutput `pulumi:"linuxProfile"`
	// Is local account disabled for AAD integrated kubernetes cluster?
	LocalAccountDisabled pulumi.BoolPtrOutput `pulumi:"localAccountDisabled"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow KubernetesClusterMaintenanceWindowPtrOutput `pulumi:"maintenanceWindow"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfileOutput `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringOutput `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolOutput `pulumi:"privateClusterEnabled"`
	// Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`.
	PrivateClusterPublicFqdnEnabled pulumi.BoolPtrOutput `pulumi:"privateClusterPublicFqdnEnabled"`
	// Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
	PrivateDnsZoneId pulumi.StringOutput `pulumi:"privateDnsZoneId"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringOutput `pulumi:"privateFqdn"`
	// Deprecated: Deprecated in favour of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolOutput `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlOutput `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
	ServicePrincipal KubernetesClusterServicePrincipalPtrOutput `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrOutput `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfileOutput `pulumi:"windowsProfile"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultNodePool == nil {
		return nil, errors.New("invalid value for required argument 'DefaultNodePool'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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
	// The IP ranges to allow for incoming traffic to the server nodes.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `pulumi:"autoScalerProfile"`
	// The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, `node-image` and `stable`. Omitting this field sets this value to `none`.
	AutomaticChannelUpgrade *string `pulumi:"automaticChannelUpgrade"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool *KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
	DnsPrefixPrivateCluster *string `pulumi:"dnsPrefixPrivateCluster"`
	EnablePodSecurityPolicy *bool   `pulumi:"enablePodSecurityPolicy"`
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn *string `pulumi:"fqdn"`
	// An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw *string `pulumi:"kubeAdminConfigRaw"`
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs []KubernetesClusterKubeAdminConfig `pulumi:"kubeAdminConfigs"`
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw *string `pulumi:"kubeConfigRaw"`
	// A `kubeConfig` block as defined below.
	KubeConfigs []KubernetesClusterKubeConfig `pulumi:"kubeConfigs"`
	// A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
	KubeletIdentities []KubernetesClusterKubeletIdentity `pulumi:"kubeletIdentities"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// Is local account disabled for AAD integrated kubernetes cluster?
	LocalAccountDisabled *bool `pulumi:"localAccountDisabled"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow *KubernetesClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled *bool `pulumi:"privateClusterEnabled"`
	// Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`.
	PrivateClusterPublicFqdnEnabled *bool `pulumi:"privateClusterPublicFqdnEnabled"`
	// Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
	PrivateDnsZoneId *string `pulumi:"privateDnsZoneId"`
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn *string `pulumi:"privateFqdn"`
	// Deprecated: Deprecated in favour of `private_cluster_enabled`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
	ServicePrincipal *KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

type KubernetesClusterState struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// The IP ranges to allow for incoming traffic to the server nodes.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfilePtrInput
	// The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, `node-image` and `stable`. Omitting this field sets this value to `none`.
	AutomaticChannelUpgrade pulumi.StringPtrInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolPtrInput
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringPtrInput
	// Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
	DnsPrefixPrivateCluster pulumi.StringPtrInput
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// The FQDN of the Azure Kubernetes Managed Cluster.
	Fqdn pulumi.StringPtrInput
	// An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
	Identity KubernetesClusterIdentityPtrInput
	// Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigRaw pulumi.StringPtrInput
	// A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
	KubeAdminConfigs KubernetesClusterKubeAdminConfigArrayInput
	// Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
	KubeConfigRaw pulumi.StringPtrInput
	// A `kubeConfig` block as defined below.
	KubeConfigs KubernetesClusterKubeConfigArrayInput
	// A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
	KubeletIdentities KubernetesClusterKubeletIdentityArrayInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// Is local account disabled for AAD integrated kubernetes cluster?
	LocalAccountDisabled pulumi.BoolPtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow KubernetesClusterMaintenanceWindowPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolPtrInput
	// Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`.
	PrivateClusterPublicFqdnEnabled pulumi.BoolPtrInput
	// Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
	PrivateDnsZoneId pulumi.StringPtrInput
	// The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
	PrivateFqdn pulumi.StringPtrInput
	// Deprecated: Deprecated in favour of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
	ServicePrincipal KubernetesClusterServicePrincipalPtrInput
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrInput
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
	// The IP ranges to allow for incoming traffic to the server nodes.
	ApiServerAuthorizedIpRanges []string `pulumi:"apiServerAuthorizedIpRanges"`
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `pulumi:"autoScalerProfile"`
	// The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, `node-image` and `stable`. Omitting this field sets this value to `none`.
	AutomaticChannelUpgrade *string `pulumi:"automaticChannelUpgrade"`
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePool `pulumi:"defaultNodePool"`
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
	DnsPrefixPrivateCluster *string `pulumi:"dnsPrefixPrivateCluster"`
	EnablePodSecurityPolicy *bool   `pulumi:"enablePodSecurityPolicy"`
	// An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
	Identity *KubernetesClusterIdentity `pulumi:"identity"`
	// A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
	KubeletIdentities []KubernetesClusterKubeletIdentity `pulumi:"kubeletIdentities"`
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A `linuxProfile` block as defined below.
	LinuxProfile *KubernetesClusterLinuxProfile `pulumi:"linuxProfile"`
	// Is local account disabled for AAD integrated kubernetes cluster?
	LocalAccountDisabled *bool `pulumi:"localAccountDisabled"`
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow *KubernetesClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `networkProfile` block as defined below.
	NetworkProfile *KubernetesClusterNetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled *bool `pulumi:"privateClusterEnabled"`
	// Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`.
	PrivateClusterPublicFqdnEnabled *bool `pulumi:"privateClusterPublicFqdnEnabled"`
	// Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
	PrivateDnsZoneId *string `pulumi:"privateDnsZoneId"`
	// Deprecated: Deprecated in favour of `private_cluster_enabled`
	PrivateLinkEnabled *bool `pulumi:"privateLinkEnabled"`
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl *KubernetesClusterRoleBasedAccessControl `pulumi:"roleBasedAccessControl"`
	// A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
	ServicePrincipal *KubernetesClusterServicePrincipal `pulumi:"servicePrincipal"`
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier *string `pulumi:"skuTier"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `windowsProfile` block as defined below.
	WindowsProfile *KubernetesClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// A `addonProfile` block as defined below.
	AddonProfile KubernetesClusterAddonProfilePtrInput
	// The IP ranges to allow for incoming traffic to the server nodes.
	ApiServerAuthorizedIpRanges pulumi.StringArrayInput
	// A `autoScalerProfile` block as defined below.
	AutoScalerProfile KubernetesClusterAutoScalerProfilePtrInput
	// The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, `node-image` and `stable`. Omitting this field sets this value to `none`.
	AutomaticChannelUpgrade pulumi.StringPtrInput
	// A `defaultNodePool` block as defined below.
	DefaultNodePool KubernetesClusterDefaultNodePoolInput
	// The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
	DiskEncryptionSetId pulumi.StringPtrInput
	// DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
	DnsPrefix pulumi.StringPtrInput
	// Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
	DnsPrefixPrivateCluster pulumi.StringPtrInput
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
	Identity KubernetesClusterIdentityPtrInput
	// A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
	KubeletIdentities KubernetesClusterKubeletIdentityArrayInput
	// Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
	KubernetesVersion pulumi.StringPtrInput
	// A `linuxProfile` block as defined below.
	LinuxProfile KubernetesClusterLinuxProfilePtrInput
	// Is local account disabled for AAD integrated kubernetes cluster?
	LocalAccountDisabled pulumi.BoolPtrInput
	// The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `maintenanceWindow` block as defined below.
	MaintenanceWindow KubernetesClusterMaintenanceWindowPtrInput
	// The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `networkProfile` block as defined below.
	NetworkProfile KubernetesClusterNetworkProfilePtrInput
	// The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
	NodeResourceGroup pulumi.StringPtrInput
	// Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
	PrivateClusterEnabled pulumi.BoolPtrInput
	// Specifies whether a Public FQDN for this Private Cluster should be added. Defaults to `false`.
	PrivateClusterPublicFqdnEnabled pulumi.BoolPtrInput
	// Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
	PrivateDnsZoneId pulumi.StringPtrInput
	// Deprecated: Deprecated in favour of `private_cluster_enabled`
	PrivateLinkEnabled pulumi.BoolPtrInput
	// Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
	RoleBasedAccessControl KubernetesClusterRoleBasedAccessControlPtrInput
	// A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
	ServicePrincipal KubernetesClusterServicePrincipalPtrInput
	// The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
	SkuTier pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `windowsProfile` block as defined below.
	WindowsProfile KubernetesClusterWindowsProfilePtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

type KubernetesClusterInput interface {
	pulumi.Input

	ToKubernetesClusterOutput() KubernetesClusterOutput
	ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput
}

func (*KubernetesCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesCluster)(nil))
}

func (i *KubernetesCluster) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return i.ToKubernetesClusterOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterOutput)
}

func (i *KubernetesCluster) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return i.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPtrOutput)
}

type KubernetesClusterPtrInput interface {
	pulumi.Input

	ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput
	ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput
}

type kubernetesClusterPtrType KubernetesClusterArgs

func (*kubernetesClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil))
}

func (i *kubernetesClusterPtrType) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return i.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterPtrType) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPtrOutput)
}

// KubernetesClusterArrayInput is an input type that accepts KubernetesClusterArray and KubernetesClusterArrayOutput values.
// You can construct a concrete instance of `KubernetesClusterArrayInput` via:
//
//          KubernetesClusterArray{ KubernetesClusterArgs{...} }
type KubernetesClusterArrayInput interface {
	pulumi.Input

	ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput
	ToKubernetesClusterArrayOutputWithContext(context.Context) KubernetesClusterArrayOutput
}

type KubernetesClusterArray []KubernetesClusterInput

func (KubernetesClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesCluster)(nil)).Elem()
}

func (i KubernetesClusterArray) ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput {
	return i.ToKubernetesClusterArrayOutputWithContext(context.Background())
}

func (i KubernetesClusterArray) ToKubernetesClusterArrayOutputWithContext(ctx context.Context) KubernetesClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterArrayOutput)
}

// KubernetesClusterMapInput is an input type that accepts KubernetesClusterMap and KubernetesClusterMapOutput values.
// You can construct a concrete instance of `KubernetesClusterMapInput` via:
//
//          KubernetesClusterMap{ "key": KubernetesClusterArgs{...} }
type KubernetesClusterMapInput interface {
	pulumi.Input

	ToKubernetesClusterMapOutput() KubernetesClusterMapOutput
	ToKubernetesClusterMapOutputWithContext(context.Context) KubernetesClusterMapOutput
}

type KubernetesClusterMap map[string]KubernetesClusterInput

func (KubernetesClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesCluster)(nil)).Elem()
}

func (i KubernetesClusterMap) ToKubernetesClusterMapOutput() KubernetesClusterMapOutput {
	return i.ToKubernetesClusterMapOutputWithContext(context.Background())
}

func (i KubernetesClusterMap) ToKubernetesClusterMapOutputWithContext(ctx context.Context) KubernetesClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterMapOutput)
}

type KubernetesClusterOutput struct{ *pulumi.OutputState }

func (KubernetesClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesCluster)(nil))
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return o.ToKubernetesClusterPtrOutputWithContext(context.Background())
}

func (o KubernetesClusterOutput) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesCluster) *KubernetesCluster {
		return &v
	}).(KubernetesClusterPtrOutput)
}

type KubernetesClusterPtrOutput struct{ *pulumi.OutputState }

func (KubernetesClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil))
}

func (o KubernetesClusterPtrOutput) ToKubernetesClusterPtrOutput() KubernetesClusterPtrOutput {
	return o
}

func (o KubernetesClusterPtrOutput) ToKubernetesClusterPtrOutputWithContext(ctx context.Context) KubernetesClusterPtrOutput {
	return o
}

func (o KubernetesClusterPtrOutput) Elem() KubernetesClusterOutput {
	return o.ApplyT(func(v *KubernetesCluster) KubernetesCluster {
		if v != nil {
			return *v
		}
		var ret KubernetesCluster
		return ret
	}).(KubernetesClusterOutput)
}

type KubernetesClusterArrayOutput struct{ *pulumi.OutputState }

func (KubernetesClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KubernetesCluster)(nil))
}

func (o KubernetesClusterArrayOutput) ToKubernetesClusterArrayOutput() KubernetesClusterArrayOutput {
	return o
}

func (o KubernetesClusterArrayOutput) ToKubernetesClusterArrayOutputWithContext(ctx context.Context) KubernetesClusterArrayOutput {
	return o
}

func (o KubernetesClusterArrayOutput) Index(i pulumi.IntInput) KubernetesClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KubernetesCluster {
		return vs[0].([]KubernetesCluster)[vs[1].(int)]
	}).(KubernetesClusterOutput)
}

type KubernetesClusterMapOutput struct{ *pulumi.OutputState }

func (KubernetesClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KubernetesCluster)(nil))
}

func (o KubernetesClusterMapOutput) ToKubernetesClusterMapOutput() KubernetesClusterMapOutput {
	return o
}

func (o KubernetesClusterMapOutput) ToKubernetesClusterMapOutputWithContext(ctx context.Context) KubernetesClusterMapOutput {
	return o
}

func (o KubernetesClusterMapOutput) MapIndex(k pulumi.StringInput) KubernetesClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KubernetesCluster {
		return vs[0].(map[string]KubernetesCluster)[vs[1].(string)]
	}).(KubernetesClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesClusterOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterArrayOutput{})
	pulumi.RegisterOutputType(KubernetesClusterMapOutput{})
}
