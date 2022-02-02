// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Disk Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleVirtualNetwork.resourceGroupName,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.0.0/24"],
 *     delegations: [{
 *         name: "diskspool",
 *         serviceDelegation: {
 *             actions: ["Microsoft.Network/virtualNetworks/read"],
 *             name: "Microsoft.StoragePool/diskPools",
 *         },
 *     }],
 * });
 * const exampleDiskPool = new azure.compute.DiskPool("exampleDiskPool", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     skuName: "Basic_B1",
 *     subnetId: exampleSubnet.id,
 *     zones: ["1"],
 * });
 * ```
 *
 * ## Import
 *
 * Disk Pools can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/diskPool:DiskPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/diskPool1
 * ```
 */
export class DiskPool extends pulumi.CustomResource {
    /**
     * Get an existing DiskPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskPoolState, opts?: pulumi.CustomResourceOptions): DiskPool {
        return new DiskPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/diskPool:DiskPool';

    /**
     * Returns true if the given object is an instance of DiskPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskPool.__pulumiType;
    }

    /**
     * The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Disk Pool. Changing this forces a new Disk Pool to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Sku of the Disk Pool. Possible values are `Basic_B1`, `Standard_S1` and `Premium_P1`. Changing this forces a new Disk Pool to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * The ID of the Subnet where the Disk Pool should be created. Changing this forces a new Disk Pool to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Disk Pool.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of Zones where this Disk Pool should be deployed. Changing this forces a new Disk Pool to be created.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a DiskPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskPoolArgs | DiskPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskPoolState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["skuName"] = state ? state.skuName : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as DiskPoolArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.zones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zones'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["skuName"] = args ? args.skuName : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DiskPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskPool resources.
 */
export interface DiskPoolState {
    /**
     * The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Disk Pool. Changing this forces a new Disk Pool to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Sku of the Disk Pool. Possible values are `Basic_B1`, `Standard_S1` and `Premium_P1`. Changing this forces a new Disk Pool to be created.
     */
    skuName?: pulumi.Input<string>;
    /**
     * The ID of the Subnet where the Disk Pool should be created. Changing this forces a new Disk Pool to be created.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Disk Pool.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Zones where this Disk Pool should be deployed. Changing this forces a new Disk Pool to be created.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DiskPool resource.
 */
export interface DiskPoolArgs {
    /**
     * The Azure Region where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Disk Pool. Changing this forces a new Disk Pool to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The Sku of the Disk Pool. Possible values are `Basic_B1`, `Standard_S1` and `Premium_P1`. Changing this forces a new Disk Pool to be created.
     */
    skuName: pulumi.Input<string>;
    /**
     * The ID of the Subnet where the Disk Pool should be created. Changing this forces a new Disk Pool to be created.
     */
    subnetId: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Disk Pool.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Zones where this Disk Pool should be deployed. Changing this forces a new Disk Pool to be created.
     */
    zones: pulumi.Input<pulumi.Input<string>[]>;
}
