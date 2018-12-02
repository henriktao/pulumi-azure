// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Consumer Group within an IotHub
 */
export class ConsumerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ConsumerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConsumerGroupState): ConsumerGroup {
        return new ConsumerGroup(name, <any>state, { id });
    }

    /**
     * The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
     */
    public readonly eventhubEndpointName: pulumi.Output<string>;
    /**
     * The name of the IoT Hub. Changing this forces a new resource to be created.
     */
    public readonly iothubName: pulumi.Output<string>;
    /**
     * The name of this Consumer Group. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;

    /**
     * Create a ConsumerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConsumerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConsumerGroupArgs | ConsumerGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ConsumerGroupState = argsOrState as ConsumerGroupState | undefined;
            inputs["eventhubEndpointName"] = state ? state.eventhubEndpointName : undefined;
            inputs["iothubName"] = state ? state.iothubName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ConsumerGroupArgs | undefined;
            if (!args || args.eventhubEndpointName === undefined) {
                throw new Error("Missing required property 'eventhubEndpointName'");
            }
            if (!args || args.iothubName === undefined) {
                throw new Error("Missing required property 'iothubName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["eventhubEndpointName"] = args ? args.eventhubEndpointName : undefined;
            inputs["iothubName"] = args ? args.iothubName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        super("azure:iot/consumerGroup:ConsumerGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConsumerGroup resources.
 */
export interface ConsumerGroupState {
    /**
     * The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
     */
    readonly eventhubEndpointName?: pulumi.Input<string>;
    /**
     * The name of the IoT Hub. Changing this forces a new resource to be created.
     */
    readonly iothubName?: pulumi.Input<string>;
    /**
     * The name of this Consumer Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConsumerGroup resource.
 */
export interface ConsumerGroupArgs {
    /**
     * The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
     */
    readonly eventhubEndpointName: pulumi.Input<string>;
    /**
     * The name of the IoT Hub. Changing this forces a new resource to be created.
     */
    readonly iothubName: pulumi.Input<string>;
    /**
     * The name of this Consumer Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
