// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an Authorization Rule for an Event Hub Namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.eventhub.getNamespaceAuthorizationRule({
 *     name: "navi",
 *     resourceGroupName: "example-resources",
 *     namespaceName: "example-ns",
 * });
 * export const eventhubAuthorizationRuleId = data.azurem_eventhub_namespace_authorization_rule.example.id;
 * ```
 */
export function getNamespaceAuthorizationRule(args: GetNamespaceAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:eventhub/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamespaceAuthorizationRule.
 */
export interface GetNamespaceAuthorizationRuleArgs {
    /**
     * The name of the EventHub Authorization Rule resource.
     */
    name: string;
    /**
     * Specifies the name of the EventHub Namespace.
     */
    namespaceName: string;
    /**
     * The name of the resource group in which the EventHub Namespace exists.
     */
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNamespaceAuthorizationRule.
 */
export interface GetNamespaceAuthorizationRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Does this Authorization Rule have permissions to Listen to the Event Hub?
     */
    readonly listen: boolean;
    /**
     * Does this Authorization Rule have permissions to Manage to the Event Hub?
     */
    readonly manage: boolean;
    readonly name: string;
    readonly namespaceName: string;
    /**
     * The Primary Connection String for the Event Hubs authorization Rule.
     */
    readonly primaryConnectionString: string;
    /**
     * The alias of the Primary Connection String for the Event Hubs authorization Rule.
     */
    readonly primaryConnectionStringAlias: string;
    /**
     * The Primary Key for the Event Hubs authorization Rule.
     */
    readonly primaryKey: string;
    readonly resourceGroupName: string;
    /**
     * The Secondary Connection String for the Event Hubs authorization Rule.
     */
    readonly secondaryConnectionString: string;
    /**
     * The alias of the Secondary Connection String for the Event Hubs authorization Rule.
     */
    readonly secondaryConnectionStringAlias: string;
    /**
     * The Secondary Key for the Event Hubs authorization Rule.
     */
    readonly secondaryKey: string;
    /**
     * Does this Authorization Rule have permissions to Send to the Event Hub?
     */
    readonly send: boolean;
}

export function getNamespaceAuthorizationRuleOutput(args: GetNamespaceAuthorizationRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespaceAuthorizationRuleResult> {
    return pulumi.output(args).apply(a => getNamespaceAuthorizationRule(a, opts))
}

/**
 * A collection of arguments for invoking getNamespaceAuthorizationRule.
 */
export interface GetNamespaceAuthorizationRuleOutputArgs {
    /**
     * The name of the EventHub Authorization Rule resource.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Namespace exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
