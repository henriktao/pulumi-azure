// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Blob Event Trigger inside an Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const examplePipeline = new azure.datafactory.Pipeline("examplePipeline", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryId: exampleFactory.id,
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleTriggerBlobEvent = new azure.datafactory.TriggerBlobEvent("exampleTriggerBlobEvent", {
 *     dataFactoryId: exampleFactory.id,
 *     storageAccountId: exampleAccount.id,
 *     events: [
 *         "Microsoft.Storage.BlobCreated",
 *         "Microsoft.Storage.BlobDeleted",
 *     ],
 *     blobPathEndsWith: ".txt",
 *     ignoreEmptyBlobs: true,
 *     activated: true,
 *     annotations: [
 *         "test1",
 *         "test2",
 *         "test3",
 *     ],
 *     description: "example description",
 *     pipelines: [{
 *         name: examplePipeline.name,
 *         parameters: {
 *             Env: "Prod",
 *         },
 *     }],
 *     additionalProperties: {
 *         foo: "foo1",
 *         bar: "bar2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Blob Event Trigger can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/triggerBlobEvent:TriggerBlobEvent example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
 * ```
 */
export class TriggerBlobEvent extends pulumi.CustomResource {
    /**
     * Get an existing TriggerBlobEvent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerBlobEventState, opts?: pulumi.CustomResourceOptions): TriggerBlobEvent {
        return new TriggerBlobEvent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/triggerBlobEvent:TriggerBlobEvent';

    /**
     * Returns true if the given object is an instance of TriggerBlobEvent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TriggerBlobEvent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TriggerBlobEvent.__pulumiType;
    }

    /**
     * Specifies if the Data Factory Blob Event Trigger is activated. Defaults to `true`.
     */
    public readonly activated!: pulumi.Output<boolean | undefined>;
    /**
     * A map of additional properties to associate with the Data Factory Blob Event Trigger.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Blob Event Trigger.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The pattern that blob path starts with for trigger to fire.
     */
    public readonly blobPathBeginsWith!: pulumi.Output<string | undefined>;
    /**
     * The pattern that blob path ends with for trigger to fire.
     */
    public readonly blobPathEndsWith!: pulumi.Output<string | undefined>;
    /**
     * The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
     */
    public readonly dataFactoryId!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Blob Event Trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of events that will fire this trigger. Possible values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobDeleted`.
     */
    public readonly events!: pulumi.Output<string[]>;
    /**
     * are blobs with zero bytes ignored?
     */
    public readonly ignoreEmptyBlobs!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Data Factory Blob Event Trigger. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `pipeline` blocks as defined below.
     */
    public readonly pipelines!: pulumi.Output<outputs.datafactory.TriggerBlobEventPipeline[]>;
    /**
     * The ID of Storage Account in which blob event will be listened. Changing this forces a new resource.
     */
    public readonly storageAccountId!: pulumi.Output<string>;

    /**
     * Create a TriggerBlobEvent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerBlobEventArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerBlobEventArgs | TriggerBlobEventState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerBlobEventState | undefined;
            resourceInputs["activated"] = state ? state.activated : undefined;
            resourceInputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["blobPathBeginsWith"] = state ? state.blobPathBeginsWith : undefined;
            resourceInputs["blobPathEndsWith"] = state ? state.blobPathEndsWith : undefined;
            resourceInputs["dataFactoryId"] = state ? state.dataFactoryId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["events"] = state ? state.events : undefined;
            resourceInputs["ignoreEmptyBlobs"] = state ? state.ignoreEmptyBlobs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pipelines"] = state ? state.pipelines : undefined;
            resourceInputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as TriggerBlobEventArgs | undefined;
            if ((!args || args.dataFactoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryId'");
            }
            if ((!args || args.events === undefined) && !opts.urn) {
                throw new Error("Missing required property 'events'");
            }
            if ((!args || args.pipelines === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelines'");
            }
            if ((!args || args.storageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            resourceInputs["activated"] = args ? args.activated : undefined;
            resourceInputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["blobPathBeginsWith"] = args ? args.blobPathBeginsWith : undefined;
            resourceInputs["blobPathEndsWith"] = args ? args.blobPathEndsWith : undefined;
            resourceInputs["dataFactoryId"] = args ? args.dataFactoryId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["events"] = args ? args.events : undefined;
            resourceInputs["ignoreEmptyBlobs"] = args ? args.ignoreEmptyBlobs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pipelines"] = args ? args.pipelines : undefined;
            resourceInputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TriggerBlobEvent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TriggerBlobEvent resources.
 */
export interface TriggerBlobEventState {
    /**
     * Specifies if the Data Factory Blob Event Trigger is activated. Defaults to `true`.
     */
    activated?: pulumi.Input<boolean>;
    /**
     * A map of additional properties to associate with the Data Factory Blob Event Trigger.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Blob Event Trigger.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The pattern that blob path starts with for trigger to fire.
     */
    blobPathBeginsWith?: pulumi.Input<string>;
    /**
     * The pattern that blob path ends with for trigger to fire.
     */
    blobPathEndsWith?: pulumi.Input<string>;
    /**
     * The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
     */
    dataFactoryId?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Blob Event Trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * List of events that will fire this trigger. Possible values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobDeleted`.
     */
    events?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * are blobs with zero bytes ignored?
     */
    ignoreEmptyBlobs?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Data Factory Blob Event Trigger. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more `pipeline` blocks as defined below.
     */
    pipelines?: pulumi.Input<pulumi.Input<inputs.datafactory.TriggerBlobEventPipeline>[]>;
    /**
     * The ID of Storage Account in which blob event will be listened. Changing this forces a new resource.
     */
    storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TriggerBlobEvent resource.
 */
export interface TriggerBlobEventArgs {
    /**
     * Specifies if the Data Factory Blob Event Trigger is activated. Defaults to `true`.
     */
    activated?: pulumi.Input<boolean>;
    /**
     * A map of additional properties to associate with the Data Factory Blob Event Trigger.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Blob Event Trigger.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The pattern that blob path starts with for trigger to fire.
     */
    blobPathBeginsWith?: pulumi.Input<string>;
    /**
     * The pattern that blob path ends with for trigger to fire.
     */
    blobPathEndsWith?: pulumi.Input<string>;
    /**
     * The ID of Data Factory in which to associate the Trigger with. Changing this forces a new resource.
     */
    dataFactoryId: pulumi.Input<string>;
    /**
     * The description for the Data Factory Blob Event Trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * List of events that will fire this trigger. Possible values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobDeleted`.
     */
    events: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * are blobs with zero bytes ignored?
     */
    ignoreEmptyBlobs?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Data Factory Blob Event Trigger. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more `pipeline` blocks as defined below.
     */
    pipelines: pulumi.Input<pulumi.Input<inputs.datafactory.TriggerBlobEventPipeline>[]>;
    /**
     * The ID of Storage Account in which blob event will be listened. Changing this forces a new resource.
     */
    storageAccountId: pulumi.Input<string>;
}
