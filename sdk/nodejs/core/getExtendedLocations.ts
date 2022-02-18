// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source return the available Extended Locations for a specific Azure Region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.core.getExtendedLocations({
 *     location: "West Europe",
 * }));
 * ```
 */
export function getExtendedLocations(args: GetExtendedLocationsArgs, opts?: pulumi.InvokeOptions): Promise<GetExtendedLocationsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:core/getExtendedLocations:getExtendedLocations", {
        "location": args.location,
    }, opts);
}

/**
 * A collection of arguments for invoking getExtendedLocations.
 */
export interface GetExtendedLocationsArgs {
    /**
     * The Azure location to retrieve the Extended Locations for.
     */
    location: string;
}

/**
 * A collection of values returned by getExtendedLocations.
 */
export interface GetExtendedLocationsResult {
    /**
     * The available extended locations for the Azure Location.
     */
    readonly extendedLocations: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
}

export function getExtendedLocationsOutput(args: GetExtendedLocationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExtendedLocationsResult> {
    return pulumi.output(args).apply(a => getExtendedLocations(a, opts))
}

/**
 * A collection of arguments for invoking getExtendedLocations.
 */
export interface GetExtendedLocationsOutputArgs {
    /**
     * The Azure location to retrieve the Extended Locations for.
     */
    location: pulumi.Input<string>;
}
