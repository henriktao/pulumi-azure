// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS zones within Azure DNS. These zones are hosted on Azure's name servers to which you can delegate the zone from the parent domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceGroup("example", {location: "West Europe"});
 * const example_public = new azure.dns.Zone("example-public", {resourceGroupName: example.name});
 * const example_private = new azure.privatedns.Zone("example-private", {resourceGroupName: example.name});
 * ```
 *
 * ## Import
 *
 * DNS Zones can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:dns/zone:Zone zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1
 * ```
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:dns/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * (Optional) Maximum number of Records in the zone. Defaults to `1000`.
     */
    public /*out*/ readonly maxNumberOfRecordSets!: pulumi.Output<number>;
    /**
     * The name of the DNS Zone. Must be a valid domain name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Optional) A list of values that make up the NS record for the zone.
     */
    public /*out*/ readonly nameServers!: pulumi.Output<string[]>;
    /**
     * (Optional) The number of records already in the zone.
     */
    public /*out*/ readonly numberOfRecordSets!: pulumi.Output<number>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * An `soaRecord` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly soaRecord!: pulumi.Output<outputs.dns.ZoneSoaRecord>;
    /**
     * A mapping of tags to assign to the Record Set.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneState | undefined;
            resourceInputs["maxNumberOfRecordSets"] = state ? state.maxNumberOfRecordSets : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameServers"] = state ? state.nameServers : undefined;
            resourceInputs["numberOfRecordSets"] = state ? state.numberOfRecordSets : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["soaRecord"] = state ? state.soaRecord : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["soaRecord"] = args ? args.soaRecord : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["maxNumberOfRecordSets"] = undefined /*out*/;
            resourceInputs["nameServers"] = undefined /*out*/;
            resourceInputs["numberOfRecordSets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Zone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * (Optional) Maximum number of Records in the zone. Defaults to `1000`.
     */
    maxNumberOfRecordSets?: pulumi.Input<number>;
    /**
     * The name of the DNS Zone. Must be a valid domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional) A list of values that make up the NS record for the zone.
     */
    nameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The number of records already in the zone.
     */
    numberOfRecordSets?: pulumi.Input<number>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * An `soaRecord` block as defined below. Changing this forces a new resource to be created.
     */
    soaRecord?: pulumi.Input<inputs.dns.ZoneSoaRecord>;
    /**
     * A mapping of tags to assign to the Record Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * The name of the DNS Zone. Must be a valid domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * An `soaRecord` block as defined below. Changing this forces a new resource to be created.
     */
    soaRecord?: pulumi.Input<inputs.dns.ZoneSoaRecord>;
    /**
     * A mapping of tags to assign to the Record Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
