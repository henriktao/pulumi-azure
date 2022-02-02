// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS AAAA Records within Azure Private DNS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const testResourceGroup = new azure.core.ResourceGroup("testResourceGroup", {location: "West Europe"});
 * const testZone = new azure.privatedns.Zone("testZone", {resourceGroupName: testResourceGroup.name});
 * const testAAAARecord = new azure.privatedns.AAAARecord("testAAAARecord", {
 *     zoneName: testZone.name,
 *     resourceGroupName: testResourceGroup.name,
 *     ttl: 300,
 *     records: [
 *         "fd5d:70bc:930e:d008:0000:0000:0000:7334",
 *         "fd5d:70bc:930e:d008::7335",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Private DNS AAAA Records can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:privatedns/aAAARecord:AAAARecord test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/AAAA/myrecord1
 * ```
 */
export class AAAARecord extends pulumi.CustomResource {
    /**
     * Get an existing AAAARecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AAAARecordState, opts?: pulumi.CustomResourceOptions): AAAARecord {
        return new AAAARecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/aAAARecord:AAAARecord';

    /**
     * Returns true if the given object is an instance of AAAARecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AAAARecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AAAARecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS AAAA Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS A Record.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of IPv6 Addresses.
     */
    public readonly records!: pulumi.Output<string[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a AAAARecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AAAARecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AAAARecordArgs | AAAARecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AAAARecordState | undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["records"] = state ? state.records : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as AAAARecordArgs | undefined;
            if ((!args || args.records === undefined) && !opts.urn) {
                throw new Error("Missing required property 'records'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.zoneName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["records"] = args ? args.records : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["zoneName"] = args ? args.zoneName : undefined;
            resourceInputs["fqdn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AAAARecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AAAARecord resources.
 */
export interface AAAARecordState {
    /**
     * The FQDN of the DNS AAAA Record.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS A Record.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of IPv6 Addresses.
     */
    records?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    ttl?: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AAAARecord resource.
 */
export interface AAAARecordArgs {
    /**
     * The name of the DNS A Record.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of IPv6 Addresses.
     */
    records: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    ttl: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    zoneName: pulumi.Input<string>;
}
