// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing EventHub Namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.eventhub.getNamespace({
 *     name: "search-eventhubns",
 *     resourceGroupName: "search-service",
 * });
 * export const eventhubNamespaceId = example.then(example => example.id);
 * ```
 */
export function getNamespace(args: GetNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:eventhub/getNamespace:getNamespace", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamespace.
 */
export interface GetNamespaceArgs {
    /**
     * The name of the EventHub Namespace.
     */
    name: string;
    /**
     * The Name of the Resource Group where the EventHub Namespace exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNamespace.
 */
export interface GetNamespaceResult {
    /**
     * Is Auto Inflate enabled for the EventHub Namespace?
     */
    readonly autoInflateEnabled: boolean;
    /**
     * The Capacity / Throughput Units for a `Standard` SKU namespace.
     */
    readonly capacity: number;
    /**
     * The ID of the EventHub Dedicated Cluster where this Namespace exists.
     */
    readonly dedicatedClusterId: string;
    /**
     * The primary connection string for the authorization
     * rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryConnectionString: string;
    /**
     * The alias of the primary connection string for the authorization
     * rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryConnectionStringAlias: string;
    /**
     * The primary access key for the authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryKey: string;
    /**
     * The secondary connection string for the
     * authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryConnectionString: string;
    /**
     * The alias of the secondary connection string for the
     * authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryConnectionStringAlias: string;
    /**
     * The secondary access key for the authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kafkaEnabled: boolean;
    /**
     * The Azure location where the EventHub Namespace exists
     */
    readonly location: string;
    /**
     * Specifies the maximum number of throughput units when Auto Inflate is Enabled.
     */
    readonly maximumThroughputUnits: number;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * Defines which tier to use.
     */
    readonly sku: string;
    /**
     * A mapping of tags to assign to the EventHub Namespace.
     */
    readonly tags: {[key: string]: string};
    /**
     * Is this EventHub Namespace deployed across Availability Zones?
     */
    readonly zoneRedundant: boolean;
}

export function getNamespaceOutput(args: GetNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespaceResult> {
    return pulumi.output(args).apply(a => getNamespace(a, opts))
}

/**
 * A collection of arguments for invoking getNamespace.
 */
export interface GetNamespaceOutputArgs {
    /**
     * The name of the EventHub Namespace.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the EventHub Namespace exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
