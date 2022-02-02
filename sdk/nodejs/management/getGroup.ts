// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Management Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.management.getGroup({
 *     name: "00000000-0000-0000-0000-000000000000",
 * });
 * export const displayName = example.then(example => example.displayName);
 * ```
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:management/getGroup:getGroup", {
        "displayName": args.displayName,
        "groupId": args.groupId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * Specifies the display name of this Management Group.
     */
    displayName?: string;
    /**
     * Specifies the name or UUID of this Management Group.
     *
     * @deprecated Deprecated in favour of `name`
     */
    groupId?: string;
    /**
     * Specifies the name or UUID of this Management Group.
     */
    name?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    readonly displayName: string;
    /**
     * @deprecated Deprecated in favour of `name`
     */
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The ID of any Parent Management Group.
     */
    readonly parentManagementGroupId: string;
    /**
     * A list of Subscription IDs which are assigned to the Management Group.
     */
    readonly subscriptionIds: string[];
}

export function getGroupOutput(args?: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply(a => getGroup(a, opts))
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * Specifies the display name of this Management Group.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies the name or UUID of this Management Group.
     *
     * @deprecated Deprecated in favour of `name`
     */
    groupId?: pulumi.Input<string>;
    /**
     * Specifies the name or UUID of this Management Group.
     */
    name?: pulumi.Input<string>;
}
