// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an ExpressRoute circuit.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleExpressRouteCircuit = new azure.network.ExpressRouteCircuit("exampleExpressRouteCircuit", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     serviceProviderName: "Equinix",
 *     peeringLocation: "Silicon Valley",
 *     bandwidthInMbps: 50,
 *     sku: {
 *         tier: "Standard",
 *         family: "MeteredData",
 *     },
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ExpressRoute circuits can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/expressRouteCircuit:ExpressRouteCircuit myExpressRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute
 * ```
 */
export class ExpressRouteCircuit extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRouteCircuit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExpressRouteCircuitState, opts?: pulumi.CustomResourceOptions): ExpressRouteCircuit {
        return new ExpressRouteCircuit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/expressRouteCircuit:ExpressRouteCircuit';

    /**
     * Returns true if the given object is an instance of ExpressRouteCircuit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRouteCircuit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRouteCircuit.__pulumiType;
    }

    /**
     * Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
     */
    public readonly allowClassicOperations!: pulumi.Output<boolean | undefined>;
    /**
     * The bandwidth in Gbps of the circuit being created on the Express Route Port.
     */
    public readonly bandwidthInGbps!: pulumi.Output<number | undefined>;
    /**
     * The bandwidth in Mbps of the circuit being created on the Service Provider.
     */
    public readonly bandwidthInMbps!: pulumi.Output<number | undefined>;
    /**
     * The ID of the Express Route Port this Express Route Circuit is based on.
     */
    public readonly expressRoutePortId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
     */
    public readonly peeringLocation!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The string needed by the service provider to provision the ExpressRoute circuit.
     */
    public /*out*/ readonly serviceKey!: pulumi.Output<string>;
    /**
     * The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
     */
    public readonly serviceProviderName!: pulumi.Output<string | undefined>;
    /**
     * The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
     */
    public /*out*/ readonly serviceProviderProvisioningState!: pulumi.Output<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    public readonly sku!: pulumi.Output<outputs.network.ExpressRouteCircuitSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ExpressRouteCircuit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRouteCircuitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExpressRouteCircuitArgs | ExpressRouteCircuitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExpressRouteCircuitState | undefined;
            resourceInputs["allowClassicOperations"] = state ? state.allowClassicOperations : undefined;
            resourceInputs["bandwidthInGbps"] = state ? state.bandwidthInGbps : undefined;
            resourceInputs["bandwidthInMbps"] = state ? state.bandwidthInMbps : undefined;
            resourceInputs["expressRoutePortId"] = state ? state.expressRoutePortId : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peeringLocation"] = state ? state.peeringLocation : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["serviceKey"] = state ? state.serviceKey : undefined;
            resourceInputs["serviceProviderName"] = state ? state.serviceProviderName : undefined;
            resourceInputs["serviceProviderProvisioningState"] = state ? state.serviceProviderProvisioningState : undefined;
            resourceInputs["sku"] = state ? state.sku : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ExpressRouteCircuitArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            resourceInputs["allowClassicOperations"] = args ? args.allowClassicOperations : undefined;
            resourceInputs["bandwidthInGbps"] = args ? args.bandwidthInGbps : undefined;
            resourceInputs["bandwidthInMbps"] = args ? args.bandwidthInMbps : undefined;
            resourceInputs["expressRoutePortId"] = args ? args.expressRoutePortId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serviceProviderName"] = args ? args.serviceProviderName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["serviceKey"] = undefined /*out*/;
            resourceInputs["serviceProviderProvisioningState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExpressRouteCircuit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExpressRouteCircuit resources.
 */
export interface ExpressRouteCircuitState {
    /**
     * Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
     */
    allowClassicOperations?: pulumi.Input<boolean>;
    /**
     * The bandwidth in Gbps of the circuit being created on the Express Route Port.
     */
    bandwidthInGbps?: pulumi.Input<number>;
    /**
     * The bandwidth in Mbps of the circuit being created on the Service Provider.
     */
    bandwidthInMbps?: pulumi.Input<number>;
    /**
     * The ID of the Express Route Port this Express Route Circuit is based on.
     */
    expressRoutePortId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
     */
    peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The string needed by the service provider to provision the ExpressRoute circuit.
     */
    serviceKey?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
     */
    serviceProviderName?: pulumi.Input<string>;
    /**
     * The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
     */
    serviceProviderProvisioningState?: pulumi.Input<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    sku?: pulumi.Input<inputs.network.ExpressRouteCircuitSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ExpressRouteCircuit resource.
 */
export interface ExpressRouteCircuitArgs {
    /**
     * Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
     */
    allowClassicOperations?: pulumi.Input<boolean>;
    /**
     * The bandwidth in Gbps of the circuit being created on the Express Route Port.
     */
    bandwidthInGbps?: pulumi.Input<number>;
    /**
     * The bandwidth in Mbps of the circuit being created on the Service Provider.
     */
    bandwidthInMbps?: pulumi.Input<number>;
    /**
     * The ID of the Express Route Port this Express Route Circuit is based on.
     */
    expressRoutePortId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
     */
    peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
     */
    serviceProviderName?: pulumi.Input<string>;
    /**
     * A `sku` block for the ExpressRoute circuit as documented below.
     */
    sku: pulumi.Input<inputs.network.ExpressRouteCircuitSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
