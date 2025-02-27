// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Stream Analytics Job.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.streamanalytics.getJob({
 *     name: "example-job",
 *     resourceGroupName: "example-resources",
 * });
 * export const jobId = example.then(example => example.jobId);
 * ```
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:streamanalytics/getJob:getJob", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getJob.
 */
export interface GetJobArgs {
    /**
     * Specifies the name of the Stream Analytics Job.
     */
    name: string;
    /**
     * Specifies the name of the resource group the Stream Analytics Job is located in.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getJob.
 */
export interface GetJobResult {
    /**
     * The compatibility level for this job.
     */
    readonly compatibilityLevel: string;
    /**
     * The Data Locale of the Job.
     */
    readonly dataLocale: string;
    /**
     * The maximum tolerable delay in seconds where events arriving late could be included.
     */
    readonly eventsLateArrivalMaxDelayInSeconds: number;
    /**
     * The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
     */
    readonly eventsOutOfOrderMaxDelayInSeconds: number;
    /**
     * The policy which should be applied to events which arrive out of order in the input event stream.
     */
    readonly eventsOutOfOrderPolicy: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Optional) An `identity` block as defined below.
     */
    readonly identities: outputs.streamanalytics.GetJobIdentity[];
    /**
     * The Job ID assigned by the Stream Analytics Job.
     */
    readonly jobId: string;
    /**
     * The Azure location where the Stream Analytics Job exists.
     */
    readonly location: string;
    readonly name: string;
    /**
     * The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).
     */
    readonly outputErrorPolicy: string;
    readonly resourceGroupName: string;
    /**
     * The number of streaming units that the streaming job uses.
     */
    readonly streamingUnits: number;
    /**
     * The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
     */
    readonly transformationQuery: string;
}

export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply(a => getJob(a, opts))
}

/**
 * A collection of arguments for invoking getJob.
 */
export interface GetJobOutputArgs {
    /**
     * Specifies the name of the Stream Analytics Job.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the resource group the Stream Analytics Job is located in.
     */
    resourceGroupName: pulumi.Input<string>;
}
