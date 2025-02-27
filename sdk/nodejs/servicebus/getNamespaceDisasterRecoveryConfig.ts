// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getNamespaceDisasterRecoveryConfig(args: GetNamespaceDisasterRecoveryConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceDisasterRecoveryConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:servicebus/getNamespaceDisasterRecoveryConfig:getNamespaceDisasterRecoveryConfig", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamespaceDisasterRecoveryConfig.
 */
export interface GetNamespaceDisasterRecoveryConfigArgs {
    name: string;
    namespaceName: string;
    resourceGroupName: string;
}

/**
 * A collection of values returned by getNamespaceDisasterRecoveryConfig.
 */
export interface GetNamespaceDisasterRecoveryConfigResult {
    readonly defaultPrimaryKey: string;
    readonly defaultSecondaryKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly namespaceName: string;
    readonly partnerNamespaceId: string;
    readonly primaryConnectionStringAlias: string;
    readonly resourceGroupName: string;
    readonly secondaryConnectionStringAlias: string;
}

export function getNamespaceDisasterRecoveryConfigOutput(args: GetNamespaceDisasterRecoveryConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespaceDisasterRecoveryConfigResult> {
    return pulumi.output(args).apply(a => getNamespaceDisasterRecoveryConfig(a, opts))
}

/**
 * A collection of arguments for invoking getNamespaceDisasterRecoveryConfig.
 */
export interface GetNamespaceDisasterRecoveryConfigOutputArgs {
    name: pulumi.Input<string>;
    namespaceName: pulumi.Input<string>;
    resourceGroupName: pulumi.Input<string>;
}
