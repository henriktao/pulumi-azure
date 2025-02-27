// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a Policy Definition, both custom and built in. Retrieves Policy Definitions from your current subscription by default.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.policy.getPolicyDefintion({
 *     displayName: "Allowed resource types",
 * });
 * export const id = example.then(example => example.id);
 * ```
 */
export function getPolicyDefintion(args?: GetPolicyDefintionArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDefintionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:policy/getPolicyDefintion:getPolicyDefintion", {
        "displayName": args.displayName,
        "managementGroupId": args.managementGroupId,
        "managementGroupName": args.managementGroupName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyDefintion.
 */
export interface GetPolicyDefintionArgs {
    /**
     * Specifies the display name of the Policy Definition. Conflicts with `name`.
     */
    displayName?: string;
    /**
     * @deprecated Deprecated in favour of `management_group_name`
     */
    managementGroupId?: string;
    /**
     * Only retrieve Policy Definitions from this Management Group.
     */
    managementGroupName?: string;
    /**
     * Specifies the name of the Policy Definition. Conflicts with `displayName`.
     */
    name?: string;
}

/**
 * A collection of values returned by getPolicyDefintion.
 */
export interface GetPolicyDefintionResult {
    /**
     * The Description of the Policy.
     */
    readonly description: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated Deprecated in favour of `management_group_name`
     */
    readonly managementGroupId?: string;
    readonly managementGroupName?: string;
    /**
     * Any Metadata defined in the Policy.
     */
    readonly metadata: string;
    readonly name: string;
    /**
     * Any Parameters defined in the Policy.
     */
    readonly parameters: string;
    /**
     * The Rule as defined (in JSON) in the Policy.
     */
    readonly policyRule: string;
    /**
     * The Type of the Policy. Possible values are "BuiltIn", "Custom" and "NotSpecified".
     */
    readonly policyType: string;
    /**
     * The Type of Policy.
     */
    readonly type: string;
}

export function getPolicyDefintionOutput(args?: GetPolicyDefintionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyDefintionResult> {
    return pulumi.output(args).apply(a => getPolicyDefintion(a, opts))
}

/**
 * A collection of arguments for invoking getPolicyDefintion.
 */
export interface GetPolicyDefintionOutputArgs {
    /**
     * Specifies the display name of the Policy Definition. Conflicts with `name`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favour of `management_group_name`
     */
    managementGroupId?: pulumi.Input<string>;
    /**
     * Only retrieve Policy Definitions from this Management Group.
     */
    managementGroupName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Policy Definition. Conflicts with `displayName`.
     */
    name?: pulumi.Input<string>;
}
