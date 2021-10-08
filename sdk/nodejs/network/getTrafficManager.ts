// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.
 *
 * ## Example Usage
 * ### World)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getTrafficManager({
 *     name: "World",
 * });
 * export const locationCode = example.then(example => example.id);
 * ```
 */
export function getTrafficManager(args: GetTrafficManagerArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficManagerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getTrafficManager:getTrafficManager", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrafficManager.
 */
export interface GetTrafficManagerArgs {
    /**
     * Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
     */
    name: string;
}

/**
 * A collection of values returned by getTrafficManager.
 */
export interface GetTrafficManagerResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}

export function getTrafficManagerOutput(args: GetTrafficManagerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficManagerResult> {
    return pulumi.output(args).apply(a => getTrafficManager(a, opts))
}

/**
 * A collection of arguments for invoking getTrafficManager.
 */
export interface GetTrafficManagerOutputArgs {
    /**
     * Specifies the name of the Location, for example `World`, `Europe` or `Germany`.
     */
    name: pulumi.Input<string>;
}
