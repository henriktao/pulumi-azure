// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Azure SignalR service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = pulumi.output(azure.signalr.getService({
 *     name: "test-signalr",
 *     resourceGroupName: "signalr-resource-group",
 * }));
 * ```
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:signalr/getService:getService", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * Specifies the name of the SignalR service.
     */
    name: string;
    /**
     * Specifies the name of the resource group the SignalR service is located in.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    /**
     * The FQDN of the SignalR service.
     */
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    readonly ipAddress: string;
    /**
     * Specifies the supported Azure location where the SignalR service exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The primary access key of the SignalR service.
     */
    readonly primaryAccessKey: string;
    /**
     * The primary connection string of the SignalR service.
     */
    readonly primaryConnectionString: string;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client use.
     */
    readonly publicPort: number;
    readonly resourceGroupName: string;
    /**
     * The secondary access key of the SignalR service.
     */
    readonly secondaryAccessKey: string;
    /**
     * The secondary connection string of the SignalR service.
     */
    readonly secondaryConnectionString: string;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side use.
     */
    readonly serverPort: number;
    readonly tags: {[key: string]: string};
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    /**
     * Specifies the name of the SignalR service.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group the SignalR service is located in.
     */
    resourceGroupName: pulumi.Input<string>;
}
