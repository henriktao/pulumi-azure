// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Recovery Services Vault.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const vault = pulumi.output(azure.recoveryservices.getVault({
 *     name: "tfex-recovery_vault",
 *     resourceGroupName: "tfex-resource_group",
 * }));
 * ```
 */
export function getVault(args: GetVaultArgs, opts?: pulumi.InvokeOptions): Promise<GetVaultResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:recoveryservices/getVault:getVault", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVault.
 */
export interface GetVaultArgs {
    /**
     * Specifies the name of the Recovery Services Vault.
     */
    name: string;
    /**
     * The name of the resource group in which the Recovery Services Vault resides.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getVault.
 */
export interface GetVaultResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Azure location where the resource resides.
     */
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The vault's current SKU.
     */
    readonly sku: string;
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getVaultOutput(args: GetVaultOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVaultResult> {
    return pulumi.output(args).apply(a => getVault(a, opts))
}

/**
 * A collection of arguments for invoking getVault.
 */
export interface GetVaultOutputArgs {
    /**
     * Specifies the name of the Recovery Services Vault.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Recovery Services Vault resides.
     */
    resourceGroupName: pulumi.Input<string>;
}
