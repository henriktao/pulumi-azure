// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing App Service Environment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appservice.getAppServiceEnvironment({
 *     name: "existing-ase",
 *     resourceGroupName: "existing-rg",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getAppServiceEnvironment(args: GetAppServiceEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppServiceEnvironment.
 */
export interface GetAppServiceEnvironmentArgs {
    /**
     * The name of this App Service Environment.
     */
    name: string;
    /**
     * The name of the Resource Group where the App Service Environment exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getAppServiceEnvironment.
 */
export interface GetAppServiceEnvironmentResult {
    /**
     * Zero or more `clusterSetting` blocks as defined below.
     */
    readonly clusterSettings: outputs.appservice.GetAppServiceEnvironmentClusterSetting[];
    /**
     * The number of app instances per App Service Environment Front End.
     */
    readonly frontEndScaleFactor: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IP address of internal load balancer of the App Service Environment.
     */
    readonly internalIpAddress: string;
    /**
     * The Azure Region where the App Service Environment exists.
     */
    readonly location: string;
    /**
     * The name of the Cluster Setting.
     */
    readonly name: string;
    /**
     * List of outbound IP addresses of the App Service Environment.
     */
    readonly outboundIpAddresses: string[];
    /**
     * The Pricing Tier (Isolated SKU) of the App Service Environment.
     */
    readonly pricingTier: string;
    readonly resourceGroupName: string;
    /**
     * IP address of service endpoint of the App Service Environment.
     */
    readonly serviceIpAddress: string;
    /**
     * A mapping of tags assigned to the App Service Environment.
     */
    readonly tags: {[key: string]: string};
}

export function getAppServiceEnvironmentOutput(args: GetAppServiceEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppServiceEnvironmentResult> {
    return pulumi.output(args).apply(a => getAppServiceEnvironment(a, opts))
}

/**
 * A collection of arguments for invoking getAppServiceEnvironment.
 */
export interface GetAppServiceEnvironmentOutputArgs {
    /**
     * The name of this App Service Environment.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the App Service Environment exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
