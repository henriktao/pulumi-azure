// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Web Application Firewall Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.waf.getFirewallPolicy({
 *     resourceGroupName: "existing",
 *     name: "existing",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getFirewallPolicy(args: GetFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:waf/getFirewallPolicy:getFirewallPolicy", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallPolicy.
 */
export interface GetFirewallPolicyArgs {
    /**
     * The name of the Web Application Firewall Policy
     */
    name: string;
    /**
     * The name of the Resource Group where the Web Application Firewall Policy exists.
     */
    resourceGroupName: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getFirewallPolicy.
 */
export interface GetFirewallPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    readonly resourceGroupName: string;
    readonly tags?: {[key: string]: string};
}
