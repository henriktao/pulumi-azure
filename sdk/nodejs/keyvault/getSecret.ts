// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Secret.
 * 
 * > **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = azurerm_key_vault_existing.id.apply(id => azure.keyvault.getSecret({
 *     keyVaultId: id,
 *     name: "secret-sauce",
 * }));
 * 
 * export const secretValue = test.value;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/key_vault_secret.html.markdown.
 */
export function getSecret(args: GetSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretResult> & GetSecretResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSecretResult> = pulumi.runtime.invoke("azure:keyvault/getSecret:getSecret", {
        "keyVaultId": args.keyVaultId,
        "name": args.name,
        "vaultUri": args.vaultUri,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretArgs {
    /**
     * Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource. 
     */
    readonly keyVaultId?: string;
    /**
     * Specifies the name of the Key Vault Secret.
     */
    readonly name: string;
    readonly vaultUri?: string;
}

/**
 * A collection of values returned by getSecret.
 */
export interface GetSecretResult {
    /**
     * The content type for the Key Vault Secret.
     */
    readonly contentType: string;
    readonly keyVaultId: string;
    readonly name: string;
    /**
     * Any tags assigned to this resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The value of the Key Vault Secret.
     */
    readonly value: string;
    readonly vaultUri: string;
    /**
     * The current version of the Key Vault Secret.
     */
    readonly version: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
