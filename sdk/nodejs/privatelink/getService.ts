// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Private Link Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.privatelink.getService({
 *     name: "myPrivateLinkService",
 *     resourceGroupName: "PrivateLinkServiceRG",
 * });
 * export const privateLinkServiceId = example.then(example => example.id);
 * ```
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:privatelink/getService:getService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * The name of the private link service.
     */
    name: string;
    /**
     * The name of the resource group in which the private link service resides.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    /**
     * The alias is a globally unique name for your private link service which Azure generates for you. Your can use this alias to request a connection to your private link service.
     */
    readonly alias: string;
    /**
     * The list of subscription(s) globally unique identifiers that will be auto approved to use the private link service.
     */
    readonly autoApprovalSubscriptionIds: string[];
    /**
     * Does the Private Link Service support the Proxy Protocol?
     */
    readonly enableProxyProtocol: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of Standard Load Balancer(SLB) resource IDs. The Private Link service is tied to the frontend IP address of a SLB. All traffic destined for the private link service will reach the frontend of the SLB. You can configure SLB rules to direct this traffic to appropriate backend pools where your applications are running.
     */
    readonly loadBalancerFrontendIpConfigurationIds: string[];
    /**
     * The supported Azure location where the resource exists.
     */
    readonly location: string;
    /**
     * The name of private link service NAT IP configuration.
     */
    readonly name: string;
    /**
     * The `natIpConfiguration` block as defined below.
     */
    readonly natIpConfigurations: outputs.privatelink.GetServiceNatIpConfiguration[];
    readonly resourceGroupName: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The list of subscription(s) globally unique identifiers(GUID) that will be able to see the private link service.
     */
    readonly visibilitySubscriptionIds: string[];
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    /**
     * The name of the private link service.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group in which the private link service resides.
     */
    resourceGroupName: pulumi.Input<string>;
}
