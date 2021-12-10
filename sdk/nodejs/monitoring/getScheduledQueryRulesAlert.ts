// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access the properties of an AlertingAction scheduled query rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.monitoring.getScheduledQueryRulesAlert({
 *     resourceGroupName: "example-rg",
 *     name: "tfex-queryrule",
 * });
 * export const queryRuleId = example.then(example => example.id);
 * ```
 */
export function getScheduledQueryRulesAlert(args: GetScheduledQueryRulesAlertArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledQueryRulesAlertResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:monitoring/getScheduledQueryRulesAlert:getScheduledQueryRulesAlert", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getScheduledQueryRulesAlert.
 */
export interface GetScheduledQueryRulesAlertArgs {
    /**
     * Specifies the name of the scheduled query rule.
     */
    name: string;
    /**
     * Specifies the name of the resource group where the scheduled query rule is located.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getScheduledQueryRulesAlert.
 */
export interface GetScheduledQueryRulesAlertResult {
    /**
     * An `action` block as defined below.
     */
    readonly actions: outputs.monitoring.GetScheduledQueryRulesAlertAction[];
    /**
     * The list of Resource IDs referred into query.
     */
    readonly authorizedResourceIds: string[];
    /**
     * The resource URI over which log search query is to be run.
     */
    readonly dataSourceId: string;
    /**
     * The description of the scheduled query rule.
     */
    readonly description: string;
    /**
     * Whether this scheduled query rule is enabled.
     */
    readonly enabled: boolean;
    /**
     * Frequency at which rule condition should be evaluated.
     */
    readonly frequency: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    /**
     * Log search query.
     */
    readonly query: string;
    readonly queryType: string;
    readonly resourceGroupName: string;
    /**
     * Severity of the alert.
     */
    readonly severity: number;
    readonly tags: {[key: string]: string};
    /**
     * Time for which alerts should be throttled or suppressed.
     */
    readonly throttling: number;
    /**
     * Time window for which data needs to be fetched for query.
     */
    readonly timeWindow: number;
    /**
     * A `trigger` block as defined below.
     */
    readonly triggers: outputs.monitoring.GetScheduledQueryRulesAlertTrigger[];
}

export function getScheduledQueryRulesAlertOutput(args: GetScheduledQueryRulesAlertOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScheduledQueryRulesAlertResult> {
    return pulumi.output(args).apply(a => getScheduledQueryRulesAlert(a, opts))
}

/**
 * A collection of arguments for invoking getScheduledQueryRulesAlert.
 */
export interface GetScheduledQueryRulesAlertOutputArgs {
    /**
     * Specifies the name of the scheduled query rule.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group where the scheduled query rule is located.
     */
    resourceGroupName: pulumi.Input<string>;
}
