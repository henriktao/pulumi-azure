// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve a list of secret names from an existing Key Vault Secret.
 */
export function getSecrets(args: GetSecretsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getSecrets:getSecrets", {
        "keyVaultId": args.keyVaultId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecrets.
 */
export interface GetSecretsArgs {
    /**
     * Specifies the ID of the Key Vault instance to fetch secret names from, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    keyVaultId: string;
}

/**
 * A collection of values returned by getSecrets.
 */
export interface GetSecretsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Key Vault ID.
     */
    readonly keyVaultId: string;
    /**
     * List containing names of secrets that exist in this Key Vault.
     */
    readonly names: string[];
}

export function getSecretsOutput(args: GetSecretsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretsResult> {
    return pulumi.output(args).apply(a => getSecrets(a, opts))
}

/**
 * A collection of arguments for invoking getSecrets.
 */
export interface GetSecretsOutputArgs {
    /**
     * Specifies the ID of the Key Vault instance to fetch secret names from, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    keyVaultId: pulumi.Input<string>;
}
