// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Microsoft SQL Server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.mssql.getServer({
 *     name: "existingMsSqlServer",
 *     resourceGroupName: "existingResGroup",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:mssql/getServer:getServer", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerArgs {
    /**
     * The name of this Microsoft SQL Server.
     */
    name: string;
    /**
     * The name of the Resource Group where the Microsoft SQL Server exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getServer.
 */
export interface GetServerResult {
    /**
     * The server's administrator login name.
     */
    readonly administratorLogin: string;
    /**
     * The fully qualified domain name of the Azure SQL Server.
     */
    readonly fullyQualifiedDomainName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `identity` block as defined below.
     */
    readonly identities: outputs.mssql.GetServerIdentity[];
    /**
     * The Azure Region where the Microsoft SQL Server exists.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * A list of dropped restorable database IDs on the server.
     */
    readonly restorableDroppedDatabaseIds: string[];
    /**
     * A mapping of tags assigned to this Microsoft SQL Server.
     */
    readonly tags: {[key: string]: string};
    /**
     * This servers MS SQL version.
     */
    readonly version: string;
}

export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerResult> {
    return pulumi.output(args).apply(a => getServer(a, opts))
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerOutputArgs {
    /**
     * The name of this Microsoft SQL Server.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Microsoft SQL Server exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
