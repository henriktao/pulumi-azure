// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getNamespaceDisasterRecoveryConfig(args: GetNamespaceDisasterRecoveryConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceDisasterRecoveryConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
    readonly name: string;
    readonly namespaceName: string;
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getNamespaceDisasterRecoveryConfig.
 */
export interface GetNamespaceDisasterRecoveryConfigResult {
    readonly aliasPrimaryConnectionString: string;
    readonly aliasSecondaryConnectionString: string;
    readonly defaultPrimaryKey: string;
    readonly defaultSecondaryKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly namespaceName: string;
    readonly partnerNamespaceId: string;
    readonly resourceGroupName: string;
}
