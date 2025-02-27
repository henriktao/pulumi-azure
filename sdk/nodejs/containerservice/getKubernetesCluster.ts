// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.containerservice.getKubernetesCluster({
 *     name: "myakscluster",
 *     resourceGroupName: "my-example-resource-group",
 * }));
 * ```
 */
export function getKubernetesCluster(args: GetKubernetesClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:containerservice/getKubernetesCluster:getKubernetesCluster", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesCluster.
 */
export interface GetKubernetesClusterArgs {
    /**
     * The name of the managed Kubernetes Cluster.
     */
    name: string;
    /**
     * The name of the Resource Group in which the managed Kubernetes Cluster exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getKubernetesCluster.
 */
export interface GetKubernetesClusterResult {
    /**
     * A `addonProfile` block as documented below.
     */
    readonly addonProfiles: outputs.containerservice.GetKubernetesClusterAddonProfile[];
    /**
     * An `agentPoolProfile` block as documented below.
     */
    readonly agentPoolProfiles: outputs.containerservice.GetKubernetesClusterAgentPoolProfile[];
    /**
     * The IP ranges to whitelist for incoming traffic to the primaries.
     */
    readonly apiServerAuthorizedIpRanges: string[];
    /**
     * The ID of the Disk Encryption Set used for the Nodes and Volumes.
     */
    readonly diskEncryptionSetId: string;
    /**
     * The DNS Prefix of the managed Kubernetes cluster.
     */
    readonly dnsPrefix: string;
    /**
     * The FQDN of the Azure Kubernetes Managed Cluster.
     */
    readonly fqdn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `identity` block as documented below.
     */
    readonly identities: outputs.containerservice.GetKubernetesClusterIdentity[];
    /**
     * Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled and local accounts are not disabled.
     */
    readonly kubeAdminConfigRaw: string;
    /**
     * A `kubeAdminConfig` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled and local accounts are not disabled.
     */
    readonly kubeAdminConfigs: outputs.containerservice.GetKubernetesClusterKubeAdminConfig[];
    /**
     * Base64 encoded Kubernetes configuration.
     */
    readonly kubeConfigRaw: string;
    /**
     * A `kubeConfig` block as defined below.
     */
    readonly kubeConfigs: outputs.containerservice.GetKubernetesClusterKubeConfig[];
    /**
     * A `kubeletIdentity` block as documented below.
     */
    readonly kubeletIdentities: outputs.containerservice.GetKubernetesClusterKubeletIdentity[];
    /**
     * The version of Kubernetes used on the managed Kubernetes Cluster.
     */
    readonly kubernetesVersion: string;
    /**
     * A `linuxProfile` block as documented below.
     */
    readonly linuxProfiles: outputs.containerservice.GetKubernetesClusterLinuxProfile[];
    /**
     * The Azure Region in which the managed Kubernetes Cluster exists.
     */
    readonly location: string;
    /**
     * The name assigned to this pool of agents.
     */
    readonly name: string;
    /**
     * A `networkProfile` block as documented below.
     */
    readonly networkProfiles: outputs.containerservice.GetKubernetesClusterNetworkProfile[];
    /**
     * Auto-generated Resource Group containing AKS Cluster resources.
     */
    readonly nodeResourceGroup: string;
    /**
     * If the cluster has the Kubernetes API only exposed on internal IP addresses.
     */
    readonly privateClusterEnabled: boolean;
    /**
     * The FQDN of this Kubernetes Cluster when private link has been enabled. This name is only resolvable inside the Virtual Network where the Azure Kubernetes Service is located
     */
    readonly privateFqdn: string;
    readonly privateLinkEnabled: boolean;
    readonly resourceGroupName: string;
    /**
     * A `roleBasedAccessControl` block as documented below.
     */
    readonly roleBasedAccessControls: outputs.containerservice.GetKubernetesClusterRoleBasedAccessControl[];
    /**
     * A `servicePrincipal` block as documented below.
     */
    readonly servicePrincipals: outputs.containerservice.GetKubernetesClusterServicePrincipal[];
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * A `windowsProfile` block as documented below.
     */
    readonly windowsProfiles: outputs.containerservice.GetKubernetesClusterWindowsProfile[];
}

export function getKubernetesClusterOutput(args: GetKubernetesClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubernetesClusterResult> {
    return pulumi.output(args).apply(a => getKubernetesCluster(a, opts))
}

/**
 * A collection of arguments for invoking getKubernetesCluster.
 */
export interface GetKubernetesClusterOutputArgs {
    /**
     * The name of the managed Kubernetes Cluster.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the managed Kubernetes Cluster exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
