// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.
 *
 * Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.
 *
 * Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
 * and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: "westus",
 *     accountTier: "Standard",
 *     accountReplicationType: "GRS",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const exampleAccountSAS = azure.storage.getAccountSASOutput({
 *     connectionString: exampleAccount.primaryConnectionString,
 *     httpsOnly: true,
 *     signedVersion: "2017-07-29",
 *     resourceTypes: {
 *         service: true,
 *         container: false,
 *         object: false,
 *     },
 *     services: {
 *         blob: true,
 *         queue: false,
 *         table: false,
 *         file: false,
 *     },
 *     start: "2018-03-21T00:00:00Z",
 *     expiry: "2020-03-21T00:00:00Z",
 *     permissions: {
 *         read: true,
 *         write: true,
 *         "delete": false,
 *         list: false,
 *         add: true,
 *         create: true,
 *         update: false,
 *         process: false,
 *     },
 * });
 * export const sasUrlQueryString = exampleAccountSAS.apply(exampleAccountSAS => exampleAccountSAS.sas);
 * ```
 */
export function getAccountSAS(args: GetAccountSASArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountSASResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:storage/getAccountSAS:getAccountSAS", {
        "connectionString": args.connectionString,
        "expiry": args.expiry,
        "httpsOnly": args.httpsOnly,
        "ipAddresses": args.ipAddresses,
        "permissions": args.permissions,
        "resourceTypes": args.resourceTypes,
        "services": args.services,
        "signedVersion": args.signedVersion,
        "start": args.start,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountSAS.
 */
export interface GetAccountSASArgs {
    /**
     * The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of a `azure.storage.Account` resource.
     */
    connectionString: string;
    /**
     * The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    expiry: string;
    /**
     * Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
     */
    httpsOnly?: boolean;
    /**
     * IP address, or a range of IP addresses, from which to accept requests. When specifying a range, note that the range is inclusive.
     */
    ipAddresses?: string;
    /**
     * A `permissions` block as defined below.
     */
    permissions: inputs.storage.GetAccountSASPermissions;
    /**
     * A `resourceTypes` block as defined below.
     */
    resourceTypes: inputs.storage.GetAccountSASResourceTypes;
    /**
     * A `services` block as defined below.
     */
    services: inputs.storage.GetAccountSASServices;
    /**
     * Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
     */
    signedVersion?: string;
    /**
     * The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    start: string;
}

/**
 * A collection of values returned by getAccountSAS.
 */
export interface GetAccountSASResult {
    readonly connectionString: string;
    readonly expiry: string;
    readonly httpsOnly?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddresses?: string;
    readonly permissions: outputs.storage.GetAccountSASPermissions;
    readonly resourceTypes: outputs.storage.GetAccountSASResourceTypes;
    /**
     * The computed Account Shared Access Signature (SAS).
     */
    readonly sas: string;
    readonly services: outputs.storage.GetAccountSASServices;
    readonly signedVersion?: string;
    readonly start: string;
}

export function getAccountSASOutput(args: GetAccountSASOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountSASResult> {
    return pulumi.output(args).apply(a => getAccountSAS(a, opts))
}

/**
 * A collection of arguments for invoking getAccountSAS.
 */
export interface GetAccountSASOutputArgs {
    /**
     * The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of a `azure.storage.Account` resource.
     */
    connectionString: pulumi.Input<string>;
    /**
     * The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    expiry: pulumi.Input<string>;
    /**
     * Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
     */
    httpsOnly?: pulumi.Input<boolean>;
    /**
     * IP address, or a range of IP addresses, from which to accept requests. When specifying a range, note that the range is inclusive.
     */
    ipAddresses?: pulumi.Input<string>;
    /**
     * A `permissions` block as defined below.
     */
    permissions: pulumi.Input<inputs.storage.GetAccountSASPermissionsArgs>;
    /**
     * A `resourceTypes` block as defined below.
     */
    resourceTypes: pulumi.Input<inputs.storage.GetAccountSASResourceTypesArgs>;
    /**
     * A `services` block as defined below.
     */
    services: pulumi.Input<inputs.storage.GetAccountSASServicesArgs>;
    /**
     * Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to `2017-07-29`.
     */
    signedVersion?: pulumi.Input<string>;
    /**
     * The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    start: pulumi.Input<string>;
}
