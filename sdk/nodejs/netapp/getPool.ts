// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Uses this data source to access information about an existing NetApp Pool.
 *
 * ## NetApp Pool Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.netapp.getPool({
 *     resourceGroupName: "acctestRG",
 *     accountName: "acctestnetappaccount",
 *     name: "acctestnetapppool",
 * });
 * export const netappPoolId = example.then(example => example.id);
 * ```
 */
export function getPool(args: GetPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:netapp/getPool:getPool", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPool.
 */
export interface GetPoolArgs {
    /**
     * The name of the NetApp account where the NetApp pool exists.
     */
    accountName: string;
    /**
     * The name of the NetApp Pool.
     */
    name: string;
    /**
     * The Name of the Resource Group where the NetApp Pool exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getPool.
 */
export interface GetPoolResult {
    readonly accountName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure Region where the NetApp Pool exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The service level of the file system.
     */
    readonly serviceLevel: string;
    /**
     * Provisioned size of the pool in TB.
     */
    readonly sizeInTb: number;
}
