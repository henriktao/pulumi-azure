// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.containerservice.getKubernetesServiceVersions({
 *     location: "West Europe",
 * });
 * export const versions = current.then(current => current.versions);
 * export const latestVersion = current.then(current => current.latestVersion);
 * ```
 */
export function getKubernetesServiceVersions(args: GetKubernetesServiceVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesServiceVersionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:containerservice/getKubernetesServiceVersions:getKubernetesServiceVersions", {
        "includePreview": args.includePreview,
        "location": args.location,
        "versionPrefix": args.versionPrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesServiceVersions.
 */
export interface GetKubernetesServiceVersionsArgs {
    /**
     * Should Preview versions of Kubernetes in AKS be included? Defaults to `true`
     */
    includePreview?: boolean;
    /**
     * Specifies the location in which to query for versions.
     */
    location: string;
    /**
     * A prefix filter for the versions of Kubernetes which should be returned; for example `1.` will return `1.9` to `1.14`, whereas `1.12` will return `1.12.2`.
     */
    versionPrefix?: string;
}

/**
 * A collection of values returned by getKubernetesServiceVersions.
 */
export interface GetKubernetesServiceVersionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includePreview?: boolean;
    /**
     * The most recent version available. If `includePreview == false`, this is the most recent non-preview version available.
     */
    readonly latestVersion: string;
    readonly location: string;
    readonly versionPrefix?: string;
    /**
     * The list of all supported versions.
     */
    readonly versions: string[];
}

export function getKubernetesServiceVersionsOutput(args: GetKubernetesServiceVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubernetesServiceVersionsResult> {
    return pulumi.output(args).apply(a => getKubernetesServiceVersions(a, opts))
}

/**
 * A collection of arguments for invoking getKubernetesServiceVersions.
 */
export interface GetKubernetesServiceVersionsOutputArgs {
    /**
     * Should Preview versions of Kubernetes in AKS be included? Defaults to `true`
     */
    includePreview?: pulumi.Input<boolean>;
    /**
     * Specifies the location in which to query for versions.
     */
    location: pulumi.Input<string>;
    /**
     * A prefix filter for the versions of Kubernetes which should be returned; for example `1.` will return `1.9` to `1.14`, whereas `1.12` will return `1.12.2`.
     */
    versionPrefix?: pulumi.Input<string>;
}
