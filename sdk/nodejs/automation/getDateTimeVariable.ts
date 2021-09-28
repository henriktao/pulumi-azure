// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Automation Datetime Variable.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.automation.getDateTimeVariable({
 *     name: "tfex-example-var",
 *     resourceGroupName: "tfex-example-rg",
 *     automationAccountName: "tfex-example-account",
 * });
 * export const variableId = example.then(example => example.id);
 * ```
 */
export function getDateTimeVariable(args: GetDateTimeVariableArgs, opts?: pulumi.InvokeOptions): Promise<GetDateTimeVariableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:automation/getDateTimeVariable:getDateTimeVariable", {
        "automationAccountName": args.automationAccountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDateTimeVariable.
 */
export interface GetDateTimeVariableArgs {
    /**
     * The name of the automation account in which the Automation Variable exists.
     */
    automationAccountName: string;
    /**
     * The name of the Automation Variable.
     */
    name: string;
    /**
     * The Name of the Resource Group where the automation account exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getDateTimeVariable.
 */
export interface GetDateTimeVariableResult {
    readonly automationAccountName: string;
    /**
     * The description of the Automation Variable.
     */
    readonly description: string;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
     */
    readonly value: string;
}
