// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Resource Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.core.getResourceGroup({
 *     name: "existing",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getResourceGroup(args: GetResourceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:core/getResourceGroup:getResourceGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceGroup.
 */
export interface GetResourceGroupArgs {
    /**
     * The Name of this Resource Group.
     */
    name: string;
}

/**
 * A collection of values returned by getResourceGroup.
 */
export interface GetResourceGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the Resource Group exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * A mapping of tags assigned to the Resource Group.
     */
    readonly tags: {[key: string]: string};
}

export function getResourceGroupOutput(args: GetResourceGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceGroupResult> {
    return pulumi.output(args).apply(a => getResourceGroup(a, opts))
}

/**
 * A collection of arguments for invoking getResourceGroup.
 */
export interface GetResourceGroupOutputArgs {
    /**
     * The Name of this Resource Group.
     */
    name: pulumi.Input<string>;
}
