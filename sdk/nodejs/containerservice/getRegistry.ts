// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Container Registry.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.containerservice.getRegistry({
 *     name: "testacr",
 *     resourceGroupName: "test",
 * });
 * export const loginServer = example.then(example => example.loginServer);
 * ```
 */
export function getRegistry(args: GetRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:containerservice/getRegistry:getRegistry", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegistry.
 */
export interface GetRegistryArgs {
    /**
     * The name of the Container Registry.
     */
    name: string;
    /**
     * The Name of the Resource Group where this Container Registry exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getRegistry.
 */
export interface GetRegistryResult {
    /**
     * Is the Administrator account enabled for this Container Registry.
     */
    readonly adminEnabled: boolean;
    /**
     * The Password associated with the Container Registry Admin account - if the admin account is enabled.
     */
    readonly adminPassword: string;
    /**
     * The Username associated with the Container Registry Admin account - if the admin account is enabled.
     */
    readonly adminUsername: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region in which this Container Registry exists.
     */
    readonly location: string;
    /**
     * The URL that can be used to log into the container registry.
     */
    readonly loginServer: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The SKU of this Container Registry, such as `Basic`.
     */
    readonly sku: string;
    /**
     * @deprecated this attribute is no longer recognized by the API and is not functional anymore, thus this property will be removed in v3.0
     */
    readonly storageAccountId: string;
    /**
     * A map of tags assigned to the Container Registry.
     */
    readonly tags: {[key: string]: string};
}

export function getRegistryOutput(args: GetRegistryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryResult> {
    return pulumi.output(args).apply(a => getRegistry(a, opts))
}

/**
 * A collection of arguments for invoking getRegistry.
 */
export interface GetRegistryOutputArgs {
    /**
     * The name of the Container Registry.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where this Container Registry exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
