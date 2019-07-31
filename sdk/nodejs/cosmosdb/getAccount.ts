// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.cosmosdb.getAccount({
 *     name: "tfex-cosmosdb-account",
 *     resourceGroupName: "tfex-cosmosdb-account-rg",
 * }));
 * 
 * export const cosmosdbAccountEndpoint = azurerm_cosmosdb_account_jobs.endpoint;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/cosmosdb_account.html.markdown.
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> & GetAccountResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAccountResult> = pulumi.runtime.invoke("azure:cosmosdb/getAccount:getAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * Specifies the name of the CosmosDB Account.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group in which the CosmosDB Account resides.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * Capabilities enabled on this Cosmos DB account.
     */
    readonly capabilities: { name: string }[];
    readonly consistencyPolicies: { consistencyLevel: string, maxIntervalInSeconds: number, maxStalenessPrefix: number }[];
    /**
     * If automatic failover is enabled for this CosmosDB Account.
     */
    readonly enableAutomaticFailover: boolean;
    /**
     * If multi-master is enabled for this Cosmos DB account.
     */
    readonly enableMultipleWriteLocations: boolean;
    /**
     * The endpoint used to connect to the CosmosDB account.
     */
    readonly endpoint: string;
    readonly geoLocations: { failoverPriority: number, id: string, location: string }[];
    /**
     * The current IP Filter for this CosmosDB account
     */
    readonly ipRangeFilter: string;
    /**
     * If virtual network filtering is enabled for this Cosmos DB account.
     */
    readonly isVirtualNetworkFilterEnabled: boolean;
    /**
     * The Kind of the CosmosDB account.
     */
    readonly kind: string;
    /**
     * The name of the Azure region hosting replicated data.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The Offer Type to used by this CosmosDB Account.
     */
    readonly offerType: string;
    /**
     * The Primary master key for the CosmosDB Account.
     */
    readonly primaryMasterKey: string;
    /**
     * The Primary read-only master Key for the CosmosDB Account.
     */
    readonly primaryReadonlyMasterKey: string;
    /**
     * A list of read endpoints available for this CosmosDB account.
     */
    readonly readEndpoints: string[];
    readonly resourceGroupName: string;
    /**
     * The Secondary master key for the CosmosDB Account.
     */
    readonly secondaryMasterKey: string;
    /**
     * The Secondary read-only master key for the CosmosDB Account.
     */
    readonly secondaryReadonlyMasterKey: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * Subnets that are allowed to access this CosmosDB account.
     */
    readonly virtualNetworkRules: { id: string }[];
    /**
     * A list of write endpoints available for this CosmosDB account.
     */
    readonly writeEndpoints: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
