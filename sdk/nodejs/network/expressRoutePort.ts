// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Express Route Port.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleExpressRoutePort = new azure.network.ExpressRoutePort("exampleExpressRoutePort", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     peeringLocation: "Airtel-Chennai-CLS",
 *     bandwidthInGbps: 10,
 *     encapsulation: "Dot1Q",
 * });
 * ```
 *
 * ## Import
 *
 * Express Route Ports can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/expressRoutePort:ExpressRoutePort example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/expressRoutePorts/port1
 * ```
 */
export class ExpressRoutePort extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRoutePort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExpressRoutePortState, opts?: pulumi.CustomResourceOptions): ExpressRoutePort {
        return new ExpressRoutePort(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/expressRoutePort:ExpressRoutePort';

    /**
     * Returns true if the given object is an instance of ExpressRoutePort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRoutePort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRoutePort.__pulumiType;
    }

    /**
     * Bandwidth of the Express Route Port in Gbps. Changing this forces a new Express Route Port to be created.
     */
    public readonly bandwidthInGbps!: pulumi.Output<number>;
    /**
     * The encapsulation method used for the Express Route Port. Changing this forces a new Express Route Port to be created. Possible values are: `Dot1Q`, `QinQ`.
     */
    public readonly encapsulation!: pulumi.Output<string>;
    /**
     * The EtherType of the Express Route Port.
     */
    public /*out*/ readonly ethertype!: pulumi.Output<string>;
    /**
     * The resource GUID of the Express Route Port.
     */
    public /*out*/ readonly guid!: pulumi.Output<string>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.network.ExpressRoutePortIdentity | undefined>;
    /**
     * A list of `link` blocks as defined below.
     */
    public readonly link1!: pulumi.Output<outputs.network.ExpressRoutePortLink1>;
    /**
     * A list of `link` blocks as defined below.
     */
    public readonly link2!: pulumi.Output<outputs.network.ExpressRoutePortLink2>;
    /**
     * The Azure Region where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The maximum transmission unit of the Express Route Port.
     */
    public /*out*/ readonly mtu!: pulumi.Output<string>;
    /**
     * The name which should be used for this Express Route Port. Changing this forces a new Express Route Port to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the peering location that this Express Route Port is physically mapped to. Changing this forces a new Express Route Port to be created.
     */
    public readonly peeringLocation!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Express Route Port.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ExpressRoutePort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRoutePortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExpressRoutePortArgs | ExpressRoutePortState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExpressRoutePortState | undefined;
            inputs["bandwidthInGbps"] = state ? state.bandwidthInGbps : undefined;
            inputs["encapsulation"] = state ? state.encapsulation : undefined;
            inputs["ethertype"] = state ? state.ethertype : undefined;
            inputs["guid"] = state ? state.guid : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["link1"] = state ? state.link1 : undefined;
            inputs["link2"] = state ? state.link2 : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["peeringLocation"] = state ? state.peeringLocation : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ExpressRoutePortArgs | undefined;
            if ((!args || args.bandwidthInGbps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthInGbps'");
            }
            if ((!args || args.encapsulation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encapsulation'");
            }
            if ((!args || args.peeringLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peeringLocation'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bandwidthInGbps"] = args ? args.bandwidthInGbps : undefined;
            inputs["encapsulation"] = args ? args.encapsulation : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["link1"] = args ? args.link1 : undefined;
            inputs["link2"] = args ? args.link2 : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ethertype"] = undefined /*out*/;
            inputs["guid"] = undefined /*out*/;
            inputs["mtu"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ExpressRoutePort.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExpressRoutePort resources.
 */
export interface ExpressRoutePortState {
    /**
     * Bandwidth of the Express Route Port in Gbps. Changing this forces a new Express Route Port to be created.
     */
    bandwidthInGbps?: pulumi.Input<number>;
    /**
     * The encapsulation method used for the Express Route Port. Changing this forces a new Express Route Port to be created. Possible values are: `Dot1Q`, `QinQ`.
     */
    encapsulation?: pulumi.Input<string>;
    /**
     * The EtherType of the Express Route Port.
     */
    ethertype?: pulumi.Input<string>;
    /**
     * The resource GUID of the Express Route Port.
     */
    guid?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.network.ExpressRoutePortIdentity>;
    /**
     * A list of `link` blocks as defined below.
     */
    link1?: pulumi.Input<inputs.network.ExpressRoutePortLink1>;
    /**
     * A list of `link` blocks as defined below.
     */
    link2?: pulumi.Input<inputs.network.ExpressRoutePortLink2>;
    /**
     * The Azure Region where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The maximum transmission unit of the Express Route Port.
     */
    mtu?: pulumi.Input<string>;
    /**
     * The name which should be used for this Express Route Port. Changing this forces a new Express Route Port to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the peering location that this Express Route Port is physically mapped to. Changing this forces a new Express Route Port to be created.
     */
    peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Express Route Port.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ExpressRoutePort resource.
 */
export interface ExpressRoutePortArgs {
    /**
     * Bandwidth of the Express Route Port in Gbps. Changing this forces a new Express Route Port to be created.
     */
    bandwidthInGbps: pulumi.Input<number>;
    /**
     * The encapsulation method used for the Express Route Port. Changing this forces a new Express Route Port to be created. Possible values are: `Dot1Q`, `QinQ`.
     */
    encapsulation: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.network.ExpressRoutePortIdentity>;
    /**
     * A list of `link` blocks as defined below.
     */
    link1?: pulumi.Input<inputs.network.ExpressRoutePortLink1>;
    /**
     * A list of `link` blocks as defined below.
     */
    link2?: pulumi.Input<inputs.network.ExpressRoutePortLink2>;
    /**
     * The Azure Region where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Express Route Port. Changing this forces a new Express Route Port to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the peering location that this Express Route Port is physically mapped to. Changing this forces a new Express Route Port to be created.
     */
    peeringLocation: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Express Route Port.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
