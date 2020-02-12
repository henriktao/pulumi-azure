// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing EventHub Namespace.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventhub_namespace_legacy.html.markdown.
 */
export function getEventhubNamespace(args: GetEventhubNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetEventhubNamespaceResult> & GetEventhubNamespaceResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetEventhubNamespaceResult> = pulumi.runtime.invoke("azure:eventhub/getEventhubNamespace:getEventhubNamespace", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getEventhubNamespace.
 */
export interface GetEventhubNamespaceArgs {
    /**
     * The name of the EventHub Namespace.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group where the EventHub Namespace exists.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getEventhubNamespace.
 */
export interface GetEventhubNamespaceResult {
    /**
     * Is Auto Inflate enabled for the EventHub Namespace?
     */
    readonly autoInflateEnabled: boolean;
    /**
     * The Capacity / Throughput Units for a `Standard` SKU namespace.
     */
    readonly capacity: number;
    /**
     * The primary connection string for the authorization
     * rule `RootManageSharedAccessKey`.
     */
    readonly defaultPrimaryConnectionString: string;
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
     * The secondary access key for the authorization rule `RootManageSharedAccessKey`.
     */
    readonly defaultSecondaryKey: string;
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
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
