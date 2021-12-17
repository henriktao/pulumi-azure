// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about Cosmos DB Restorable Database Accounts.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.cosmosdb.getRestorableDatabaseAccounts({
 *     name: "example-ca",
 *     location: "West Europe",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getRestorableDatabaseAccounts(args: GetRestorableDatabaseAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetRestorableDatabaseAccountsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:cosmosdb/getRestorableDatabaseAccounts:getRestorableDatabaseAccounts", {
        "location": args.location,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRestorableDatabaseAccounts.
 */
export interface GetRestorableDatabaseAccountsArgs {
    /**
     * The location where the Cosmos DB Database Account.
     */
    location: string;
    /**
     * The name of this Cosmos DB Database Account.
     */
    name: string;
}

/**
 * A collection of values returned by getRestorableDatabaseAccounts.
 */
export interface GetRestorableDatabaseAccountsResult {
    /**
     * One or more `accounts` blocks as defined below.
     */
    readonly accounts: outputs.cosmosdb.GetRestorableDatabaseAccountsAccount[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The location of the regional Cosmos DB Restorable Database Account.
     */
    readonly location: string;
    readonly name: string;
}

export function getRestorableDatabaseAccountsOutput(args: GetRestorableDatabaseAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRestorableDatabaseAccountsResult> {
    return pulumi.output(args).apply(a => getRestorableDatabaseAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getRestorableDatabaseAccounts.
 */
export interface GetRestorableDatabaseAccountsOutputArgs {
    /**
     * The location where the Cosmos DB Database Account.
     */
    location: pulumi.Input<string>;
    /**
     * The name of this Cosmos DB Database Account.
     */
    name: pulumi.Input<string>;
}
