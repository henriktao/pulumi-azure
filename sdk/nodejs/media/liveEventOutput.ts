// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class LiveEventOutput extends pulumi.CustomResource {
    /**
     * Get an existing LiveEventOutput resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LiveEventOutputState, opts?: pulumi.CustomResourceOptions): LiveEventOutput {
        return new LiveEventOutput(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:media/liveEventOutput:LiveEventOutput';

    /**
     * Returns true if the given object is an instance of LiveEventOutput.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiveEventOutput {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiveEventOutput.__pulumiType;
    }

    public readonly archiveWindowDuration!: pulumi.Output<string>;
    public readonly assetName!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly hlsFragmentsPerTsSegment!: pulumi.Output<number | undefined>;
    public readonly liveEventId!: pulumi.Output<string>;
    public readonly manifestName!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly outputSnapTimeInSeconds!: pulumi.Output<number | undefined>;

    /**
     * Create a LiveEventOutput resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiveEventOutputArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LiveEventOutputArgs | LiveEventOutputState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LiveEventOutputState | undefined;
            resourceInputs["archiveWindowDuration"] = state ? state.archiveWindowDuration : undefined;
            resourceInputs["assetName"] = state ? state.assetName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hlsFragmentsPerTsSegment"] = state ? state.hlsFragmentsPerTsSegment : undefined;
            resourceInputs["liveEventId"] = state ? state.liveEventId : undefined;
            resourceInputs["manifestName"] = state ? state.manifestName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputSnapTimeInSeconds"] = state ? state.outputSnapTimeInSeconds : undefined;
        } else {
            const args = argsOrState as LiveEventOutputArgs | undefined;
            if ((!args || args.archiveWindowDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'archiveWindowDuration'");
            }
            if ((!args || args.assetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assetName'");
            }
            if ((!args || args.liveEventId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'liveEventId'");
            }
            resourceInputs["archiveWindowDuration"] = args ? args.archiveWindowDuration : undefined;
            resourceInputs["assetName"] = args ? args.assetName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hlsFragmentsPerTsSegment"] = args ? args.hlsFragmentsPerTsSegment : undefined;
            resourceInputs["liveEventId"] = args ? args.liveEventId : undefined;
            resourceInputs["manifestName"] = args ? args.manifestName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputSnapTimeInSeconds"] = args ? args.outputSnapTimeInSeconds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LiveEventOutput.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LiveEventOutput resources.
 */
export interface LiveEventOutputState {
    archiveWindowDuration?: pulumi.Input<string>;
    assetName?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    hlsFragmentsPerTsSegment?: pulumi.Input<number>;
    liveEventId?: pulumi.Input<string>;
    manifestName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    outputSnapTimeInSeconds?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a LiveEventOutput resource.
 */
export interface LiveEventOutputArgs {
    archiveWindowDuration: pulumi.Input<string>;
    assetName: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    hlsFragmentsPerTsSegment?: pulumi.Input<number>;
    liveEventId: pulumi.Input<string>;
    manifestName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    outputSnapTimeInSeconds?: pulumi.Input<number>;
}
