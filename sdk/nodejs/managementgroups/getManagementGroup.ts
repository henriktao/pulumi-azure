// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Management Group.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/management_group.html.markdown.
 */
/** @deprecated azure.managementgroups.getManagementGroup has been deprecated in favour of azure.management.getGroup */
export function getManagementGroup(args?: GetManagementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementGroupResult> {
    pulumi.log.warn("getManagementGroup is deprecated: azure.managementgroups.getManagementGroup has been deprecated in favour of azure.management.getGroup")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:managementgroups/getManagementGroup:getManagementGroup", {
        "groupId": args.groupId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagementGroup.
 */
export interface GetManagementGroupArgs {
    /**
     * Specifies the name or UUID of this Management Group.
     * 
     * @deprecated Deprecated in favour of `name`
     */
    readonly groupId?: string;
    /**
     * Specifies the name or UUID of this Management Group.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getManagementGroup.
 */
export interface GetManagementGroupResult {
    /**
     * A friendly name for the Management Group.
     */
    readonly displayName: string;
    readonly groupId: string;
    readonly name: string;
    /**
     * The ID of any Parent Management Group.
     */
    readonly parentManagementGroupId: string;
    /**
     * A list of Subscription ID's which are assigned to the Management Group.
     */
    readonly subscriptionIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
