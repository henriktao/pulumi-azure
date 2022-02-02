// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about Service Tags.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.network.getServiceTags({
 *     location: "westcentralus",
 *     service: "AzureKeyVault",
 *     locationFilter: "northeurope",
 * });
 * export const addressPrefixes = example.then(example => example.addressPrefixes);
 * export const ipv4Cidrs = example.then(example => example.ipv4Cidrs);
 * ```
 */
export function getServiceTags(args: GetServiceTagsArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceTagsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:network/getServiceTags:getServiceTags", {
        "location": args.location,
        "locationFilter": args.locationFilter,
        "service": args.service,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceTags.
 */
export interface GetServiceTagsArgs {
    /**
     * The Azure Region where the Service Tags exists. This value is not used to filter the results but for specifying the region to request. For filtering by region use `locationFilter` instead.  More information can be found here: [Service Tags URL parameters](https://docs.microsoft.com/en-us/rest/api/virtualnetwork/servicetags/list#uri-parameters).
     */
    location: string;
    /**
     * Changes the scope of the service tags. Can be any value that is also valid for `location`. If this field is empty then all address prefixes are considered instead of only location specific ones.
     */
    locationFilter?: string;
    /**
     * The type of the service for which address prefixes will be fetched. Available service tags can be found here: [Available service tags](https://docs.microsoft.com/en-us/azure/virtual-network/service-tags-overview#available-service-tags).
     */
    service: string;
}

/**
 * A collection of values returned by getServiceTags.
 */
export interface GetServiceTagsResult {
    /**
     * List of address prefixes for the service type (and optionally a specific region).
     */
    readonly addressPrefixes: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of IPv4 addresses for the service type (and optionally a specific region)
     */
    readonly ipv4Cidrs: string[];
    /**
     * List of IPv6 addresses for the service type (and optionally a specific region)
     */
    readonly ipv6Cidrs: string[];
    readonly location: string;
    readonly locationFilter?: string;
    readonly service: string;
}

export function getServiceTagsOutput(args: GetServiceTagsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceTagsResult> {
    return pulumi.output(args).apply(a => getServiceTags(a, opts))
}

/**
 * A collection of arguments for invoking getServiceTags.
 */
export interface GetServiceTagsOutputArgs {
    /**
     * The Azure Region where the Service Tags exists. This value is not used to filter the results but for specifying the region to request. For filtering by region use `locationFilter` instead.  More information can be found here: [Service Tags URL parameters](https://docs.microsoft.com/en-us/rest/api/virtualnetwork/servicetags/list#uri-parameters).
     */
    location: pulumi.Input<string>;
    /**
     * Changes the scope of the service tags. Can be any value that is also valid for `location`. If this field is empty then all address prefixes are considered instead of only location specific ones.
     */
    locationFilter?: pulumi.Input<string>;
    /**
     * The type of the service for which address prefixes will be fetched. Available service tags can be found here: [Available service tags](https://docs.microsoft.com/en-us/azure/virtual-network/service-tags-overview#available-service-tags).
     */
    service: pulumi.Input<string>;
}
