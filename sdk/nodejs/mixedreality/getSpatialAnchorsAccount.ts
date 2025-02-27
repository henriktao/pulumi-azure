// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information about an Azure Spatial Anchors Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.mixedreality.getSpatialAnchorsAccount({
 *     name: "example",
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * export const accountDomain = data.azurerm_spatial_anchors_account.account_domain;
 * ```
 */
export function getSpatialAnchorsAccount(args: GetSpatialAnchorsAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetSpatialAnchorsAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:mixedreality/getSpatialAnchorsAccount:getSpatialAnchorsAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpatialAnchorsAccount.
 */
export interface GetSpatialAnchorsAccountArgs {
    /**
     * Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
     */
    name: string;
    /**
     * The name of the resource group in which to create the Spatial Anchors Account.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getSpatialAnchorsAccount.
 */
export interface GetSpatialAnchorsAccountResult {
    /**
     * The domain of the Spatial Anchors Account.
     */
    readonly accountDomain: string;
    /**
     * The account ID of the Spatial Anchors Account.
     */
    readonly accountId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
}

export function getSpatialAnchorsAccountOutput(args: GetSpatialAnchorsAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpatialAnchorsAccountResult> {
    return pulumi.output(args).apply(a => getSpatialAnchorsAccount(a, opts))
}

/**
 * A collection of arguments for invoking getSpatialAnchorsAccount.
 */
export interface GetSpatialAnchorsAccountOutputArgs {
    /**
     * Specifies the name of the Spatial Anchors Account. Changing this forces a new resource to be created. Must be globally unique.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Spatial Anchors Account.
     */
    resourceGroupName: pulumi.Input<string>;
}
