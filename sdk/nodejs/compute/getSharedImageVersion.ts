// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.compute.getSharedImageVersion({
 *     galleryName: "my-image-gallery",
 *     imageName: "my-image",
 *     name: "1.0.0",
 *     resourceGroupName: "example-resources",
 * }));
 * ```
 */
export function getSharedImageVersion(args: GetSharedImageVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetSharedImageVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:compute/getSharedImageVersion:getSharedImageVersion", {
        "galleryName": args.galleryName,
        "imageName": args.imageName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "sortVersionsBySemver": args.sortVersionsBySemver,
    }, opts);
}

/**
 * A collection of arguments for invoking getSharedImageVersion.
 */
export interface GetSharedImageVersionArgs {
    /**
     * The name of the Shared Image Gallery in which the Shared Image exists.
     */
    galleryName: string;
    /**
     * The name of the Shared Image in which this Version exists.
     */
    imageName: string;
    /**
     * The name of the Image Version.
     */
    name: string;
    /**
     * The name of the Resource Group in which the Shared Image Gallery exists.
     */
    resourceGroupName: string;
    /**
     * Sort available versions taking SemVer versioning scheme into account. Defaults to `false`.
     */
    sortVersionsBySemver?: boolean;
}

/**
 * A collection of values returned by getSharedImageVersion.
 */
export interface GetSharedImageVersionResult {
    /**
     * Is this Image Version excluded from the `latest` filter?
     */
    readonly excludeFromLatest: boolean;
    readonly galleryName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageName: string;
    /**
     * The supported Azure location where the Shared Image Gallery exists.
     */
    readonly location: string;
    /**
     * The ID of the Managed Image which was the source of this Shared Image Version.
     */
    readonly managedImageId: string;
    /**
     * The Azure Region in which this Image Version exists.
     */
    readonly name: string;
    /**
     * The size of the OS disk snapshot (in Gigabytes) which was the source of this Shared Image Version.
     */
    readonly osDiskImageSizeGb: number;
    /**
     * The ID of the OS disk snapshot which was the source of this Shared Image Version.
     */
    readonly osDiskSnapshotId: string;
    readonly resourceGroupName: string;
    readonly sortVersionsBySemver?: boolean;
    /**
     * A mapping of tags assigned to the Shared Image.
     */
    readonly tags: {[key: string]: string};
    /**
     * One or more `targetRegion` blocks as documented below.
     */
    readonly targetRegions: outputs.compute.GetSharedImageVersionTargetRegion[];
}

export function getSharedImageVersionOutput(args: GetSharedImageVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSharedImageVersionResult> {
    return pulumi.output(args).apply(a => getSharedImageVersion(a, opts))
}

/**
 * A collection of arguments for invoking getSharedImageVersion.
 */
export interface GetSharedImageVersionOutputArgs {
    /**
     * The name of the Shared Image Gallery in which the Shared Image exists.
     */
    galleryName: pulumi.Input<string>;
    /**
     * The name of the Shared Image in which this Version exists.
     */
    imageName: pulumi.Input<string>;
    /**
     * The name of the Image Version.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Shared Image Gallery exists.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Sort available versions taking SemVer versioning scheme into account. Defaults to `false`.
     */
    sortVersionsBySemver?: pulumi.Input<boolean>;
}
