// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing App Service Plan (formerly known as a `Server Farm`).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appservice.getAppServicePlan({
 *     name: "search-app-service-plan",
 *     resourceGroupName: "search-service",
 * });
 * export const appServicePlanId = example.then(example => example.id);
 * ```
 */
export function getAppServicePlan(args: GetAppServicePlanArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServicePlanResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:appservice/getAppServicePlan:getAppServicePlan", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppServicePlan.
 */
export interface GetAppServicePlanArgs {
    /**
     * The name of the App Service Plan.
     */
    name: string;
    /**
     * The Name of the Resource Group where the App Service Plan exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getAppServicePlan.
 */
export interface GetAppServicePlanResult {
    /**
     * The ID of the App Service Environment where the App Service Plan is located.
     */
    readonly appServiceEnvironmentId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A flag that indicates if it's a xenon plan (support for Windows Container)
     */
    readonly isXenon: boolean;
    /**
     * The Operating System type of the App Service Plan
     */
    readonly kind: string;
    /**
     * The Azure location where the App Service Plan exists
     */
    readonly location: string;
    /**
     * The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
     */
    readonly maximumElasticWorkerCount: number;
    /**
     * The maximum number of workers supported with the App Service Plan's sku.
     */
    readonly maximumNumberOfWorkers: number;
    readonly name: string;
    /**
     * Can Apps assigned to this App Service Plan be scaled independently?
     */
    readonly perSiteScaling: boolean;
    /**
     * Is this App Service Plan `Reserved`?
     */
    readonly reserved: boolean;
    readonly resourceGroupName: string;
    /**
     * A `sku` block as documented below.
     */
    readonly sku: outputs.appservice.GetAppServicePlanSku;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * App Service Plan perform availability zone balancing.
     */
    readonly zoneRedundant: boolean;
}

export function getAppServicePlanOutput(args: GetAppServicePlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppServicePlanResult> {
    return pulumi.output(args).apply(a => getAppServicePlan(a, opts))
}

/**
 * A collection of arguments for invoking getAppServicePlan.
 */
export interface GetAppServicePlanOutputArgs {
    /**
     * The name of the App Service Plan.
     */
    name: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the App Service Plan exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
