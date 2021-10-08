// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing ServiceBus Namespace Authorization Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.servicebus.getNamespaceAuthorizationRule({
 *     name: "examplerule",
 *     namespaceName: "examplenamespace",
 *     resourceGroupName: "example-resources",
 * });
 * export const ruleId = example.then(example => example.id);
 * ```
 */
export function getNamespaceAuthorizationRule(args: GetNamespaceAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:servicebus/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule", {
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
     * Specifies the name of the ServiceBus Namespace Authorization Rule.
     */
    name: string;
    /**
     * Specifies the name of the ServiceBus Namespace.
     */
    namespaceName: string;
    /**
     * Specifies the name of the Resource Group where the ServiceBus Namespace exists.
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
    readonly name: string;
    readonly namespaceName: string;
    /**
     * The primary connection string for the authorization rule.
     */
    readonly primaryConnectionString: string;
    /**
     * The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
     */
    readonly primaryConnectionStringAlias: string;
    /**
     * The primary access key for the authorization rule.
     */
    readonly primaryKey: string;
    readonly resourceGroupName: string;
    /**
     * The secondary connection string for the authorization rule.
     */
    readonly secondaryConnectionString: string;
    /**
     * The alias Secondary Connection String for the ServiceBus Namespace
     */
    readonly secondaryConnectionStringAlias: string;
    /**
     * The secondary access key for the authorization rule.
     */
    readonly secondaryKey: string;
}

export function getNamespaceAuthorizationRuleOutput(args: GetNamespaceAuthorizationRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespaceAuthorizationRuleResult> {
    return pulumi.output(args).apply(a => getNamespaceAuthorizationRule(a, opts))
}

/**
 * A collection of arguments for invoking getNamespaceAuthorizationRule.
 */
export interface GetNamespaceAuthorizationRuleOutputArgs {
    /**
     * Specifies the name of the ServiceBus Namespace Authorization Rule.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the name of the ServiceBus Namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the ServiceBus Namespace exists.
     */
    resourceGroupName: pulumi.Input<string>;
}
