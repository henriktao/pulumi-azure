// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Container Registry scope map.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.containerservice.getRegistryScopeMap({
 *     name: "example-scope-map",
 *     resourceGroupName: "example-resource-group",
 *     containerRegistryName: "example-registry",
 * });
 * export const actions = example.then(example => example.actions);
 * ```
 */
export function getRegistryScopeMap(args: GetRegistryScopeMapArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryScopeMapResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:containerservice/getRegistryScopeMap:getRegistryScopeMap", {
        "containerRegistryName": args.containerRegistryName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegistryScopeMap.
 */
export interface GetRegistryScopeMapArgs {
    /**
     * The Name of the Container Registry where the token exists.
     */
    containerRegistryName: string;
    /**
     * The name of the Container Registry token.
     */
    name: string;
    /**
     * The Name of the Resource Group where this Container Registry token exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getRegistryScopeMap.
 */
export interface GetRegistryScopeMapResult {
    /**
     * The actions for the Scope Map.
     */
    readonly actions: string[];
    readonly containerRegistryName: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
}

export function getRegistryScopeMapOutput(args: GetRegistryScopeMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryScopeMapResult> {
    return pulumi.output(args).apply(a => getRegistryScopeMap(a, opts))
}

/**
 * A collection of arguments for invoking getRegistryScopeMap.
 */
export interface GetRegistryScopeMapOutputArgs {
    /**
     * The Name of the Container Registry where the token exists.
     */
    containerRegistryName: pulumi.Input<string>;
    /**
     * The name of the Container Registry token.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where this Container Registry token exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
