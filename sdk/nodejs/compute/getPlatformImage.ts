// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a Platform Image.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.compute.getPlatformImage({
 *     location: "West Europe",
 *     publisher: "Canonical",
 *     offer: "UbuntuServer",
 *     sku: "16.04-LTS",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getPlatformImage(args: GetPlatformImageArgs, opts?: pulumi.InvokeOptions): Promise<GetPlatformImageResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:compute/getPlatformImage:getPlatformImage", {
        "location": args.location,
        "offer": args.offer,
        "publisher": args.publisher,
        "sku": args.sku,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlatformImage.
 */
export interface GetPlatformImageArgs {
    /**
     * Specifies the Location to pull information about this Platform Image from.
     */
    location: string;
    /**
     * Specifies the Offer associated with the Platform Image.
     */
    offer: string;
    /**
     * Specifies the Publisher associated with the Platform Image.
     */
    publisher: string;
    /**
     * Specifies the SKU of the Platform Image.
     */
    sku: string;
    /**
     * The version of the Platform Image.
     */
    version?: string;
}

/**
 * A collection of values returned by getPlatformImage.
 */
export interface GetPlatformImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly offer: string;
    readonly publisher: string;
    readonly sku: string;
    readonly version: string;
}

export function getPlatformImageOutput(args: GetPlatformImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlatformImageResult> {
    return pulumi.output(args).apply(a => getPlatformImage(a, opts))
}

/**
 * A collection of arguments for invoking getPlatformImage.
 */
export interface GetPlatformImageOutputArgs {
    /**
     * Specifies the Location to pull information about this Platform Image from.
     */
    location: pulumi.Input<string>;
    /**
     * Specifies the Offer associated with the Platform Image.
     */
    offer: pulumi.Input<string>;
    /**
     * Specifies the Publisher associated with the Platform Image.
     */
    publisher: pulumi.Input<string>;
    /**
     * Specifies the SKU of the Platform Image.
     */
    sku: pulumi.Input<string>;
    /**
     * The version of the Platform Image.
     */
    version?: pulumi.Input<string>;
}
