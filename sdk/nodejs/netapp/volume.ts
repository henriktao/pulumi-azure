// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a NetApp Volume.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/netapp_volume.html.markdown.
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:netapp/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    public readonly exportPolicyRules!: pulumi.Output<outputs.netapp.VolumeExportPolicyRule[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    public readonly poolName!: pulumi.Output<string>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    public readonly serviceLevel!: pulumi.Output<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    public readonly storageQuotaInGb!: pulumi.Output<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    public readonly volumePath!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["exportPolicyRules"] = state ? state.exportPolicyRules : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["poolName"] = state ? state.poolName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["serviceLevel"] = state ? state.serviceLevel : undefined;
            inputs["storageQuotaInGb"] = state ? state.storageQuotaInGb : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["volumePath"] = state ? state.volumePath : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.poolName === undefined) {
                throw new Error("Missing required property 'poolName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceLevel === undefined) {
                throw new Error("Missing required property 'serviceLevel'");
            }
            if (!args || args.storageQuotaInGb === undefined) {
                throw new Error("Missing required property 'storageQuotaInGb'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            if (!args || args.volumePath === undefined) {
                throw new Error("Missing required property 'volumePath'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["exportPolicyRules"] = args ? args.exportPolicyRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["poolName"] = args ? args.poolName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceLevel"] = args ? args.serviceLevel : undefined;
            inputs["storageQuotaInGb"] = args ? args.storageQuotaInGb : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["volumePath"] = args ? args.volumePath : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    readonly exportPolicyRules?: pulumi.Input<pulumi.Input<inputs.netapp.VolumeExportPolicyRule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly poolName?: pulumi.Input<string>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    readonly serviceLevel?: pulumi.Input<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    readonly storageQuotaInGb?: pulumi.Input<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    readonly volumePath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * One or more `exportPolicyRule` block defined below.
     */
    readonly exportPolicyRules?: pulumi.Input<pulumi.Input<inputs.netapp.VolumeExportPolicyRule>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the NetApp Volume. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly poolName: pulumi.Input<string>;
    /**
     * The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
     */
    readonly serviceLevel: pulumi.Input<string>;
    /**
     * The maximum Storage Quota allowed for a file system in Gigabytes.
     */
    readonly storageQuotaInGb: pulumi.Input<number>;
    /**
     * The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
     */
    readonly volumePath: pulumi.Input<string>;
}
