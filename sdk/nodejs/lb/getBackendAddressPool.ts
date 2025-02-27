// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Load Balancer's Backend Address Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleLB = azure.lb.getLB({
 *     name: "example-lb",
 *     resourceGroupName: "example-resources",
 * });
 * const exampleBackendAddressPool = exampleLB.then(exampleLB => azure.lb.getBackendAddressPool({
 *     name: "first",
 *     loadbalancerId: exampleLB.id,
 * }));
 * export const backendAddressPoolId = exampleBackendAddressPool.then(exampleBackendAddressPool => exampleBackendAddressPool.id);
 * export const backendIpConfigurationIds = data.azurerm_lb_backend_address_pool.beap.backend_ip_configurations.map(__item => __item.id);
 * ```
 */
export function getBackendAddressPool(args: GetBackendAddressPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendAddressPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure:lb/getBackendAddressPool:getBackendAddressPool", {
        "loadbalancerId": args.loadbalancerId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackendAddressPool.
 */
export interface GetBackendAddressPoolArgs {
    /**
     * The ID of the Load Balancer in which the Backend Address Pool exists.
     */
    loadbalancerId: string;
    /**
     * Specifies the name of the Backend Address Pool.
     */
    name: string;
}

/**
 * A collection of values returned by getBackendAddressPool.
 */
export interface GetBackendAddressPoolResult {
    /**
     * A list of `backendAddress` block as defined below.
     */
    readonly backendAddresses: outputs.lb.GetBackendAddressPoolBackendAddress[];
    /**
     * A list of references to IP addresses defined in network interfaces.
     */
    readonly backendIpConfigurations: outputs.lb.GetBackendAddressPoolBackendIpConfiguration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the Load Balancing Rules associated with this Backend Address Pool.
     */
    readonly loadBalancingRules: string[];
    readonly loadbalancerId: string;
    /**
     * The name of the Backend Address.
     */
    readonly name: string;
    /**
     * A list of the Load Balancing Outbound Rules associated with this Backend Address Pool.
     */
    readonly outboundRules: string[];
}

export function getBackendAddressPoolOutput(args: GetBackendAddressPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackendAddressPoolResult> {
    return pulumi.output(args).apply(a => getBackendAddressPool(a, opts))
}

/**
 * A collection of arguments for invoking getBackendAddressPool.
 */
export interface GetBackendAddressPoolOutputArgs {
    /**
     * The ID of the Load Balancer in which the Backend Address Pool exists.
     */
    loadbalancerId: pulumi.Input<string>;
    /**
     * Specifies the name of the Backend Address Pool.
     */
    name: pulumi.Input<string>;
}
