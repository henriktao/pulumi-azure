// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleKeyVault = azure.keyvault.getKeyVault({
 *     name: azurerm_key_vault.example.name,
 *     resourceGroupName: azurerm_key_vault.example.resource_group_name,
 * });
 * const exampleDiagnosticCategories = exampleKeyVault.then(exampleKeyVault => azure.monitoring.getDiagnosticCategories({
 *     resourceId: exampleKeyVault.id,
 * }));
 * ```
 */
export function getDiagnosticCategories(args: GetDiagnosticCategoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagnosticCategoriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:monitoring/getDiagnosticCategories:getDiagnosticCategories", {
        "resourceId": args.resourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDiagnosticCategories.
 */
export interface GetDiagnosticCategoriesArgs {
    /**
     * The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
     */
    resourceId: string;
}

/**
 * A collection of values returned by getDiagnosticCategories.
 */
export interface GetDiagnosticCategoriesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the Log Categories supported for this Resource.
     */
    readonly logs: string[];
    /**
     * A list of the Metric Categories supported for this Resource.
     */
    readonly metrics: string[];
    readonly resourceId: string;
}

export function getDiagnosticCategoriesOutput(args: GetDiagnosticCategoriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDiagnosticCategoriesResult> {
    return pulumi.output(args).apply(a => getDiagnosticCategories(a, opts))
}

/**
 * A collection of arguments for invoking getDiagnosticCategories.
 */
export interface GetDiagnosticCategoriesOutputArgs {
    /**
     * The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
     */
    resourceId: pulumi.Input<string>;
}
