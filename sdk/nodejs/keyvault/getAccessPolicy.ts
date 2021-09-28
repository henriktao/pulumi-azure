// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about the permissions from the Management Key Vault Templates.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const contributor = azure.keyvault.getAccessPolicy({
 *     name: "Key Management",
 * });
 * export const accessPolicyKeyPermissions = contributor.then(contributor => contributor.keyPermissions);
 * ```
 */
export function getAccessPolicy(args: GetAccessPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:keyvault/getAccessPolicy:getAccessPolicy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPolicy.
 */
export interface GetAccessPolicyArgs {
    /**
     * Specifies the name of the Management Template. Possible values are: `Key Management`,
     * `Secret Management`, `Certificate Management`, `Key & Secret Management`, `Key & Certificate Management`,
     * `Secret & Certificate Management`,  `Key, Secret, & Certificate Management`
     */
    name: string;
}

/**
 * A collection of values returned by getAccessPolicy.
 */
export interface GetAccessPolicyResult {
    /**
     * the certificate permissions for the access policy
     */
    readonly certificatePermissions: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * the key permissions for the access policy
     */
    readonly keyPermissions: string[];
    readonly name: string;
    /**
     * the secret permissions for the access policy
     */
    readonly secretPermissions: string[];
}
