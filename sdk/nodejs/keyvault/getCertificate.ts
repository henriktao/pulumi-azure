// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Key Vault Certificate.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleKeyVault = azure.keyvault.getKeyVault({
 *     name: "examplekv",
 *     resourceGroupName: "some-resource-group",
 * });
 * const exampleCertificate = exampleKeyVault.then(exampleKeyVault => azure.keyvault.getCertificate({
 *     name: "secret-sauce",
 *     keyVaultId: exampleKeyVault.id,
 * }));
 * export const certificateThumbprint = exampleCertificate.then(exampleCertificate => exampleCertificate.thumbprint);
 * ```
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getCertificate:getCertificate", {
        "keyVaultId": args.keyVaultId,
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    keyVaultId: string;
    /**
     * Specifies the name of the Key Vault Certificate.
     */
    name: string;
    /**
     * Specifies the version of the certificate to look up.  (Defaults to latest)
     */
    version?: string;
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    /**
     * The raw Key Vault Certificate data represented as a hexadecimal string.
     */
    readonly certificateData: string;
    /**
     * The raw Key Vault Certificate data represented as a base64 string.
     */
    readonly certificateDataBase64: string;
    /**
     * A `certificatePolicy` block as defined below.
     */
    readonly certificatePolicies: outputs.keyvault.GetCertificateCertificatePolicy[];
    /**
     * Expiry date of certificate in RFC3339 format.
     */
    readonly expires: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyVaultId: string;
    /**
     * The name of the Certificate Issuer.
     */
    readonly name: string;
    /**
     * Not Before date of certificate in RFC3339 format.
     */
    readonly notBefore: string;
    /**
     * The ID of the associated Key Vault Secret.
     */
    readonly secretId: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
     */
    readonly thumbprint: string;
    /**
     * The current version of the Key Vault Certificate.
     */
    readonly version: string;
    /**
     * The Base ID of the Key Vault Certificate.
     */
    readonly versionlessId: string;
    /**
     * The Base ID of the Key Vault Secret.
     */
    readonly versionlessSecretId: string;
}

export function getCertificateOutput(args: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply(a => getCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateOutputArgs {
    /**
     * Specifies the ID of the Key Vault instance where the Secret resides, available on the `azure.keyvault.KeyVault` Data Source / Resource.
     */
    keyVaultId: pulumi.Input<string>;
    /**
     * Specifies the name of the Key Vault Certificate.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the version of the certificate to look up.  (Defaults to latest)
     */
    version?: pulumi.Input<string>;
}
