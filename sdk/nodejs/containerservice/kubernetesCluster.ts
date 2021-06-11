// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)
 *
 * ## Example Usage
 *
 * This example provisions a basic Managed Kubernetes Cluster.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleKubernetesCluster = new azure.containerservice.KubernetesCluster("exampleKubernetesCluster", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     dnsPrefix: "exampleaks1",
 *     defaultNodePool: {
 *         name: "default",
 *         nodeCount: 1,
 *         vmSize: "Standard_D2_v2",
 *     },
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 *     tags: {
 *         Environment: "Production",
 *     },
 * });
 * export const clientCertificate = exampleKubernetesCluster.kubeConfigs.apply(kubeConfigs => kubeConfigs[0].clientCertificate);
 * export const kubeConfig = exampleKubernetesCluster.kubeConfigRaw;
 * ```
 *
 * ## Import
 *
 * Managed Kubernetes Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:containerservice/kubernetesCluster:KubernetesCluster cluster1 /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
 * ```
 */
export class KubernetesCluster extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesClusterState, opts?: pulumi.CustomResourceOptions): KubernetesCluster {
        return new KubernetesCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/kubernetesCluster:KubernetesCluster';

    /**
     * Returns true if the given object is an instance of KubernetesCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesCluster.__pulumiType;
    }

    /**
     * A `addonProfile` block as defined below.
     */
    public readonly addonProfile!: pulumi.Output<outputs.containerservice.KubernetesClusterAddonProfile>;
    /**
     * The IP ranges to whitelist for incoming traffic to the masters.
     */
    public readonly apiServerAuthorizedIpRanges!: pulumi.Output<string[] | undefined>;
    /**
     * A `autoScalerProfile` block as defined below.
     */
    public readonly autoScalerProfile!: pulumi.Output<outputs.containerservice.KubernetesClusterAutoScalerProfile>;
    /**
     * The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, and `stable`.
     */
    public readonly automaticChannelUpgrade!: pulumi.Output<string | undefined>;
    /**
     * A `defaultNodePool` block as defined below.
     */
    public readonly defaultNodePool!: pulumi.Output<outputs.containerservice.KubernetesClusterDefaultNodePool>;
    /**
     * The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
     */
    public readonly diskEncryptionSetId!: pulumi.Output<string | undefined>;
    /**
     * DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
     */
    public readonly dnsPrefix!: pulumi.Output<string | undefined>;
    /**
     * Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
     */
    public readonly dnsPrefixPrivateCluster!: pulumi.Output<string | undefined>;
    public readonly enablePodSecurityPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The FQDN of the Azure Kubernetes Managed Cluster.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
     */
    public readonly identity!: pulumi.Output<outputs.containerservice.KubernetesClusterIdentity | undefined>;
    /**
     * Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
     */
    public /*out*/ readonly kubeAdminConfigRaw!: pulumi.Output<string>;
    /**
     * A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
     */
    public /*out*/ readonly kubeAdminConfigs!: pulumi.Output<outputs.containerservice.KubernetesClusterKubeAdminConfig[]>;
    /**
     * Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
     */
    public /*out*/ readonly kubeConfigRaw!: pulumi.Output<string>;
    /**
     * A `kubeConfig` block as defined below.
     */
    public /*out*/ readonly kubeConfigs!: pulumi.Output<outputs.containerservice.KubernetesClusterKubeConfig[]>;
    /**
     * A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly kubeletIdentities!: pulumi.Output<outputs.containerservice.KubernetesClusterKubeletIdentity[]>;
    /**
     * Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
     */
    public readonly kubernetesVersion!: pulumi.Output<string>;
    /**
     * A `linuxProfile` block as defined below.
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.KubernetesClusterLinuxProfile | undefined>;
    /**
     * The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    public readonly networkProfile!: pulumi.Output<outputs.containerservice.KubernetesClusterNetworkProfile>;
    /**
     * The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
     */
    public readonly nodeResourceGroup!: pulumi.Output<string>;
    /**
     * Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly privateClusterEnabled!: pulumi.Output<boolean>;
    /**
     * Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
     */
    public readonly privateDnsZoneId!: pulumi.Output<string>;
    /**
     * The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
     */
    public /*out*/ readonly privateFqdn!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favour of `private_cluster_enabled`
     */
    public readonly privateLinkEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
     */
    public readonly roleBasedAccessControl!: pulumi.Output<outputs.containerservice.KubernetesClusterRoleBasedAccessControl>;
    /**
     * A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
     */
    public readonly servicePrincipal!: pulumi.Output<outputs.containerservice.KubernetesClusterServicePrincipal | undefined>;
    /**
     * The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
     */
    public readonly skuTier!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `windowsProfile` block as defined below.
     */
    public readonly windowsProfile!: pulumi.Output<outputs.containerservice.KubernetesClusterWindowsProfile>;

    /**
     * Create a KubernetesCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesClusterArgs | KubernetesClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesClusterState | undefined;
            inputs["addonProfile"] = state ? state.addonProfile : undefined;
            inputs["apiServerAuthorizedIpRanges"] = state ? state.apiServerAuthorizedIpRanges : undefined;
            inputs["autoScalerProfile"] = state ? state.autoScalerProfile : undefined;
            inputs["automaticChannelUpgrade"] = state ? state.automaticChannelUpgrade : undefined;
            inputs["defaultNodePool"] = state ? state.defaultNodePool : undefined;
            inputs["diskEncryptionSetId"] = state ? state.diskEncryptionSetId : undefined;
            inputs["dnsPrefix"] = state ? state.dnsPrefix : undefined;
            inputs["dnsPrefixPrivateCluster"] = state ? state.dnsPrefixPrivateCluster : undefined;
            inputs["enablePodSecurityPolicy"] = state ? state.enablePodSecurityPolicy : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["kubeAdminConfigRaw"] = state ? state.kubeAdminConfigRaw : undefined;
            inputs["kubeAdminConfigs"] = state ? state.kubeAdminConfigs : undefined;
            inputs["kubeConfigRaw"] = state ? state.kubeConfigRaw : undefined;
            inputs["kubeConfigs"] = state ? state.kubeConfigs : undefined;
            inputs["kubeletIdentities"] = state ? state.kubeletIdentities : undefined;
            inputs["kubernetesVersion"] = state ? state.kubernetesVersion : undefined;
            inputs["linuxProfile"] = state ? state.linuxProfile : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkProfile"] = state ? state.networkProfile : undefined;
            inputs["nodeResourceGroup"] = state ? state.nodeResourceGroup : undefined;
            inputs["privateClusterEnabled"] = state ? state.privateClusterEnabled : undefined;
            inputs["privateDnsZoneId"] = state ? state.privateDnsZoneId : undefined;
            inputs["privateFqdn"] = state ? state.privateFqdn : undefined;
            inputs["privateLinkEnabled"] = state ? state.privateLinkEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["roleBasedAccessControl"] = state ? state.roleBasedAccessControl : undefined;
            inputs["servicePrincipal"] = state ? state.servicePrincipal : undefined;
            inputs["skuTier"] = state ? state.skuTier : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["windowsProfile"] = state ? state.windowsProfile : undefined;
        } else {
            const args = argsOrState as KubernetesClusterArgs | undefined;
            if ((!args || args.defaultNodePool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultNodePool'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["addonProfile"] = args ? args.addonProfile : undefined;
            inputs["apiServerAuthorizedIpRanges"] = args ? args.apiServerAuthorizedIpRanges : undefined;
            inputs["autoScalerProfile"] = args ? args.autoScalerProfile : undefined;
            inputs["automaticChannelUpgrade"] = args ? args.automaticChannelUpgrade : undefined;
            inputs["defaultNodePool"] = args ? args.defaultNodePool : undefined;
            inputs["diskEncryptionSetId"] = args ? args.diskEncryptionSetId : undefined;
            inputs["dnsPrefix"] = args ? args.dnsPrefix : undefined;
            inputs["dnsPrefixPrivateCluster"] = args ? args.dnsPrefixPrivateCluster : undefined;
            inputs["enablePodSecurityPolicy"] = args ? args.enablePodSecurityPolicy : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kubeletIdentities"] = args ? args.kubeletIdentities : undefined;
            inputs["kubernetesVersion"] = args ? args.kubernetesVersion : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["nodeResourceGroup"] = args ? args.nodeResourceGroup : undefined;
            inputs["privateClusterEnabled"] = args ? args.privateClusterEnabled : undefined;
            inputs["privateDnsZoneId"] = args ? args.privateDnsZoneId : undefined;
            inputs["privateLinkEnabled"] = args ? args.privateLinkEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["roleBasedAccessControl"] = args ? args.roleBasedAccessControl : undefined;
            inputs["servicePrincipal"] = args ? args.servicePrincipal : undefined;
            inputs["skuTier"] = args ? args.skuTier : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["windowsProfile"] = args ? args.windowsProfile : undefined;
            inputs["fqdn"] = undefined /*out*/;
            inputs["kubeAdminConfigRaw"] = undefined /*out*/;
            inputs["kubeAdminConfigs"] = undefined /*out*/;
            inputs["kubeConfigRaw"] = undefined /*out*/;
            inputs["kubeConfigs"] = undefined /*out*/;
            inputs["privateFqdn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KubernetesCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesCluster resources.
 */
export interface KubernetesClusterState {
    /**
     * A `addonProfile` block as defined below.
     */
    addonProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterAddonProfile>;
    /**
     * The IP ranges to whitelist for incoming traffic to the masters.
     */
    apiServerAuthorizedIpRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `autoScalerProfile` block as defined below.
     */
    autoScalerProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterAutoScalerProfile>;
    /**
     * The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, and `stable`.
     */
    automaticChannelUpgrade?: pulumi.Input<string>;
    /**
     * A `defaultNodePool` block as defined below.
     */
    defaultNodePool?: pulumi.Input<inputs.containerservice.KubernetesClusterDefaultNodePool>;
    /**
     * The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
     */
    diskEncryptionSetId?: pulumi.Input<string>;
    /**
     * DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
     */
    dnsPrefix?: pulumi.Input<string>;
    /**
     * Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
     */
    dnsPrefixPrivateCluster?: pulumi.Input<string>;
    enablePodSecurityPolicy?: pulumi.Input<boolean>;
    /**
     * The FQDN of the Azure Kubernetes Managed Cluster.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
     */
    identity?: pulumi.Input<inputs.containerservice.KubernetesClusterIdentity>;
    /**
     * Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
     */
    kubeAdminConfigRaw?: pulumi.Input<string>;
    /**
     * A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
     */
    kubeAdminConfigs?: pulumi.Input<pulumi.Input<inputs.containerservice.KubernetesClusterKubeAdminConfig>[]>;
    /**
     * Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
     */
    kubeConfigRaw?: pulumi.Input<string>;
    /**
     * A `kubeConfig` block as defined below.
     */
    kubeConfigs?: pulumi.Input<pulumi.Input<inputs.containerservice.KubernetesClusterKubeConfig>[]>;
    /**
     * A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
     */
    kubeletIdentities?: pulumi.Input<pulumi.Input<inputs.containerservice.KubernetesClusterKubeletIdentity>[]>;
    /**
     * Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
     */
    kubernetesVersion?: pulumi.Input<string>;
    /**
     * A `linuxProfile` block as defined below.
     */
    linuxProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterLinuxProfile>;
    /**
     * The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    networkProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterNetworkProfile>;
    /**
     * The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
     */
    nodeResourceGroup?: pulumi.Input<string>;
    /**
     * Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
     */
    privateClusterEnabled?: pulumi.Input<boolean>;
    /**
     * Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
     */
    privateDnsZoneId?: pulumi.Input<string>;
    /**
     * The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
     */
    privateFqdn?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favour of `private_cluster_enabled`
     */
    privateLinkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
     */
    roleBasedAccessControl?: pulumi.Input<inputs.containerservice.KubernetesClusterRoleBasedAccessControl>;
    /**
     * A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
     */
    servicePrincipal?: pulumi.Input<inputs.containerservice.KubernetesClusterServicePrincipal>;
    /**
     * The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
     */
    skuTier?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `windowsProfile` block as defined below.
     */
    windowsProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterWindowsProfile>;
}

/**
 * The set of arguments for constructing a KubernetesCluster resource.
 */
export interface KubernetesClusterArgs {
    /**
     * A `addonProfile` block as defined below.
     */
    addonProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterAddonProfile>;
    /**
     * The IP ranges to whitelist for incoming traffic to the masters.
     */
    apiServerAuthorizedIpRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `autoScalerProfile` block as defined below.
     */
    autoScalerProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterAutoScalerProfile>;
    /**
     * The upgrade channel for this Kubernetes Cluster. Possible values are `patch`, `rapid`, and `stable`.
     */
    automaticChannelUpgrade?: pulumi.Input<string>;
    /**
     * A `defaultNodePool` block as defined below.
     */
    defaultNodePool: pulumi.Input<inputs.containerservice.KubernetesClusterDefaultNodePool>;
    /**
     * The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
     */
    diskEncryptionSetId?: pulumi.Input<string>;
    /**
     * DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
     */
    dnsPrefix?: pulumi.Input<string>;
    /**
     * Specifies the DNS prefix to use with private clusters. Changing this forces a new resource to be created.
     */
    dnsPrefixPrivateCluster?: pulumi.Input<string>;
    enablePodSecurityPolicy?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below. One of either `identity` or `servicePrincipal` must be specified.
     */
    identity?: pulumi.Input<inputs.containerservice.KubernetesClusterIdentity>;
    /**
     * A `kubeletIdentity` block as defined below. Changing this forces a new resource to be created.
     */
    kubeletIdentities?: pulumi.Input<pulumi.Input<inputs.containerservice.KubernetesClusterKubeletIdentity>[]>;
    /**
     * Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
     */
    kubernetesVersion?: pulumi.Input<string>;
    /**
     * A `linuxProfile` block as defined below.
     */
    linuxProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterLinuxProfile>;
    /**
     * The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkProfile` block as defined below.
     */
    networkProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterNetworkProfile>;
    /**
     * The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
     */
    nodeResourceGroup?: pulumi.Input<string>;
    /**
     * Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
     */
    privateClusterEnabled?: pulumi.Input<boolean>;
    /**
     * Either the ID of Private DNS Zone which should be delegated to this Cluster, `System` to have AKS manage this or `None`. In case of `None` you will need to bring your own DNS server and set up resolving, otherwise cluster will have issues after provisioning.
     */
    privateDnsZoneId?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favour of `private_cluster_enabled`
     */
    privateLinkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `roleBasedAccessControl` block. Changing this forces a new resource to be created.
     */
    roleBasedAccessControl?: pulumi.Input<inputs.containerservice.KubernetesClusterRoleBasedAccessControl>;
    /**
     * A `servicePrincipal` block as documented below. One of either `identity` or `servicePrincipal` must be specified.
     */
    servicePrincipal?: pulumi.Input<inputs.containerservice.KubernetesClusterServicePrincipal>;
    /**
     * The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
     */
    skuTier?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `windowsProfile` block as defined below.
     */
    windowsProfile?: pulumi.Input<inputs.containerservice.KubernetesClusterWindowsProfile>;
}
