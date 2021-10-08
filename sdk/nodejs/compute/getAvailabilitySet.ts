// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Availability Set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.compute.getAvailabilitySet({
 *     name: "tf-appsecuritygroup",
 *     resourceGroupName: "my-resource-group",
 * });
 * export const availabilitySetId = example.then(example => example.id);
 * ```
 */
export function getAvailabilitySet(args: GetAvailabilitySetArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailabilitySetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:compute/getAvailabilitySet:getAvailabilitySet", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailabilitySet.
 */
export interface GetAvailabilitySetArgs {
    /**
     * The name of the Availability Set.
     */
    name: string;
    /**
     * The name of the resource group in which the Availability Set exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getAvailabilitySet.
 */
export interface GetAvailabilitySetResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The supported Azure location where the Availability Set exists.
     */
    readonly location: string;
    /**
     * Whether the availability set is managed or not.
     */
    readonly managed: boolean;
    readonly name: string;
    /**
     * The number of fault domains that are used.
     */
    readonly platformFaultDomainCount: string;
    /**
     * The number of update domains that are used.
     */
    readonly platformUpdateDomainCount: string;
    readonly resourceGroupName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getAvailabilitySetOutput(args: GetAvailabilitySetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAvailabilitySetResult> {
    return pulumi.output(args).apply(a => getAvailabilitySet(a, opts))
}

/**
 * A collection of arguments for invoking getAvailabilitySet.
 */
export interface GetAvailabilitySetOutputArgs {
    /**
     * The name of the Availability Set.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Availability Set exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
