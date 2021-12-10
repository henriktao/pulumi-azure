// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Pool within a NetApp Account.
 *
 * ## NetApp Pool Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.netapp.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const examplePool = new azure.netapp.Pool("examplePool", {
 *     accountName: exampleAccount.name,
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     serviceLevel: "Premium",
 *     sizeInTb: 4,
 * });
 * ```
 *
 * ## Import
 *
 * NetApp Pool can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:netapp/pool:Pool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:netapp/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the NetApp Pool. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * QoS Type of the pool. Valid values include `Auto` or `Manual`.
     */
    public readonly qosType!: pulumi.Output<string>;
    /**
     * The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
     */
    public readonly serviceLevel!: pulumi.Output<string>;
    /**
     * Provisioned size of the pool in TB. Value must be between `4` and `500`.
     */
    public readonly sizeInTb!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PoolState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["qosType"] = state ? state.qosType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceLevel"] = state ? state.serviceLevel : undefined;
            inputs["sizeInTb"] = state ? state.sizeInTb : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceLevel'");
            }
            if ((!args || args.sizeInTb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sizeInTb'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["qosType"] = args ? args.qosType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceLevel"] = args ? args.serviceLevel : undefined;
            inputs["sizeInTb"] = args ? args.sizeInTb : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Pool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the NetApp Pool. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * QoS Type of the pool. Valid values include `Auto` or `Manual`.
     */
    qosType?: pulumi.Input<string>;
    /**
     * The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
     */
    serviceLevel?: pulumi.Input<string>;
    /**
     * Provisioned size of the pool in TB. Value must be between `4` and `500`.
     */
    sizeInTb?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    accountName: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the NetApp Pool. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * QoS Type of the pool. Valid values include `Auto` or `Manual`.
     */
    qosType?: pulumi.Input<string>;
    /**
     * The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
     */
    serviceLevel: pulumi.Input<string>;
    /**
     * Provisioned size of the pool in TB. Value must be between `4` and `500`.
     */
    sizeInTb: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
