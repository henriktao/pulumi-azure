// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Scheduler Job Collection.
 * 
 * > **NOTE:** Support for Scheduler Job Collections has been deprecated by Microsoft in favour of Logic Apps ([more information can be found at this link](https://docs.microsoft.com/en-us/azure/scheduler/migrate-from-scheduler-to-logic-apps)) - as such we plan to remove support for this resource as a part of version 2.0 of the AzureRM Provider.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/scheduler_job_collection.html.markdown.
 */
export class JobCollection extends pulumi.CustomResource {
    /**
     * Get an existing JobCollection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobCollectionState, opts?: pulumi.CustomResourceOptions): JobCollection {
        return new JobCollection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:scheduler/jobCollection:JobCollection';

    /**
     * Returns true if the given object is an instance of JobCollection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobCollection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobCollection.__pulumiType;
    }

    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configures the Job collection quotas as documented in the `quota` block below.
     */
    public readonly quota!: pulumi.Output<outputs.scheduler.JobCollectionQuota | undefined>;
    /**
     * The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a JobCollection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobCollectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobCollectionArgs | JobCollectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobCollectionState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["quota"] = state ? state.quota : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as JobCollectionArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["quota"] = args ? args.quota : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(JobCollection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobCollection resources.
 */
export interface JobCollectionState {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configures the Job collection quotas as documented in the `quota` block below.
     */
    readonly quota?: pulumi.Input<inputs.scheduler.JobCollectionQuota>;
    /**
     * The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a JobCollection resource.
 */
export interface JobCollectionArgs {
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configures the Job collection quotas as documented in the `quota` block below.
     */
    readonly quota?: pulumi.Input<inputs.scheduler.JobCollectionQuota>;
    /**
     * The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
     */
    readonly sku: pulumi.Input<string>;
    /**
     * Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
