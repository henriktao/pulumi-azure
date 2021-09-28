// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing NAT Gateway.
 */
export function getNatGateway(args: GetNatGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetNatGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:network/getNatGateway:getNatGateway", {
        "name": args.name,
        "publicIpAddressIds": args.publicIpAddressIds,
        "publicIpPrefixIds": args.publicIpPrefixIds,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatGateway.
 */
export interface GetNatGatewayArgs {
    /**
     * Specifies the Name of the NAT Gateway.
     */
    name: string;
    /**
     * A list of existing Public IP Address resource IDs which the NAT Gateway is using.
     */
    publicIpAddressIds?: string[];
    /**
     * A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
     */
    publicIpPrefixIds?: string[];
    /**
     * Specifies the name of the Resource Group where the NAT Gateway exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNatGateway.
 */
export interface GetNatGatewayResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The idle timeout in minutes which is used for the NAT Gateway.
     */
    readonly idleTimeoutInMinutes: number;
    /**
     * The location where the NAT Gateway exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * A list of existing Public IP Address resource IDs which the NAT Gateway is using.
     */
    readonly publicIpAddressIds: string[];
    /**
     * A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.
     */
    readonly publicIpPrefixIds: string[];
    readonly resourceGroupName: string;
    /**
     * The Resource GUID of the NAT Gateway.
     */
    readonly resourceGuid: string;
    /**
     * The SKU used by the NAT Gateway.
     */
    readonly skuName: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * A list of Availability Zones which the NAT Gateway exists in.
     */
    readonly zones: string[];
}
