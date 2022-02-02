// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Stream Analytics Stream Input EventHub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = azure.core.getResourceGroup({
 *     name: "example-resources",
 * });
 * const exampleJob = azure.streamanalytics.getJob({
 *     name: "example-job",
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * const exampleEventHubNamespace = new azure.eventhub.EventHubNamespace("exampleEventHubNamespace", {
 *     location: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.location),
 *     resourceGroupName: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.name),
 *     sku: "Standard",
 *     capacity: 1,
 * });
 * const exampleEventHub = new azure.eventhub.EventHub("exampleEventHub", {
 *     namespaceName: exampleEventHubNamespace.name,
 *     resourceGroupName: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.name),
 *     partitionCount: 2,
 *     messageRetention: 1,
 * });
 * const exampleConsumerGroup = new azure.eventhub.ConsumerGroup("exampleConsumerGroup", {
 *     namespaceName: exampleEventHubNamespace.name,
 *     eventhubName: exampleEventHub.name,
 *     resourceGroupName: exampleResourceGroup.then(exampleResourceGroup => exampleResourceGroup.name),
 * });
 * const exampleStreamInputEventHub = new azure.streamanalytics.StreamInputEventHub("exampleStreamInputEventHub", {
 *     streamAnalyticsJobName: exampleJob.then(exampleJob => exampleJob.name),
 *     resourceGroupName: exampleJob.then(exampleJob => exampleJob.resourceGroupName),
 *     eventhubConsumerGroupName: exampleConsumerGroup.name,
 *     eventhubName: exampleEventHub.name,
 *     servicebusNamespace: exampleEventHubNamespace.name,
 *     sharedAccessPolicyKey: exampleEventHubNamespace.defaultPrimaryKey,
 *     sharedAccessPolicyName: "RootManageSharedAccessKey",
 *     serialization: {
 *         type: "Json",
 *         encoding: "UTF8",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Stream Analytics Stream Input EventHub's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:streamanalytics/streamInputEventHub:StreamInputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
 * ```
 */
export class StreamInputEventHub extends pulumi.CustomResource {
    /**
     * Get an existing StreamInputEventHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamInputEventHubState, opts?: pulumi.CustomResourceOptions): StreamInputEventHub {
        return new StreamInputEventHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:streamanalytics/streamInputEventHub:StreamInputEventHub';

    /**
     * Returns true if the given object is an instance of StreamInputEventHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamInputEventHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamInputEventHub.__pulumiType;
    }

    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub. If not set the input will use the Event Hub's default consumer group.
     */
    public readonly eventhubConsumerGroupName!: pulumi.Output<string | undefined>;
    /**
     * The name of the Event Hub.
     */
    public readonly eventhubName!: pulumi.Output<string>;
    /**
     * The name of the Stream Input EventHub. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The property the input Event Hub has been partitioned by.
     */
    public readonly partitionKey!: pulumi.Output<string | undefined>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `serialization` block as defined below.
     */
    public readonly serialization!: pulumi.Output<outputs.streamanalytics.StreamInputEventHubSerialization>;
    /**
     * The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    public readonly servicebusNamespace!: pulumi.Output<string>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    public readonly sharedAccessPolicyKey!: pulumi.Output<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    public readonly sharedAccessPolicyName!: pulumi.Output<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    public readonly streamAnalyticsJobName!: pulumi.Output<string>;

    /**
     * Create a StreamInputEventHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamInputEventHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamInputEventHubArgs | StreamInputEventHubState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamInputEventHubState | undefined;
            resourceInputs["eventhubConsumerGroupName"] = state ? state.eventhubConsumerGroupName : undefined;
            resourceInputs["eventhubName"] = state ? state.eventhubName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partitionKey"] = state ? state.partitionKey : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["serialization"] = state ? state.serialization : undefined;
            resourceInputs["servicebusNamespace"] = state ? state.servicebusNamespace : undefined;
            resourceInputs["sharedAccessPolicyKey"] = state ? state.sharedAccessPolicyKey : undefined;
            resourceInputs["sharedAccessPolicyName"] = state ? state.sharedAccessPolicyName : undefined;
            resourceInputs["streamAnalyticsJobName"] = state ? state.streamAnalyticsJobName : undefined;
        } else {
            const args = argsOrState as StreamInputEventHubArgs | undefined;
            if ((!args || args.eventhubName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventhubName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serialization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serialization'");
            }
            if ((!args || args.servicebusNamespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicebusNamespace'");
            }
            if ((!args || args.sharedAccessPolicyKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sharedAccessPolicyKey'");
            }
            if ((!args || args.sharedAccessPolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sharedAccessPolicyName'");
            }
            if ((!args || args.streamAnalyticsJobName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamAnalyticsJobName'");
            }
            resourceInputs["eventhubConsumerGroupName"] = args ? args.eventhubConsumerGroupName : undefined;
            resourceInputs["eventhubName"] = args ? args.eventhubName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partitionKey"] = args ? args.partitionKey : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serialization"] = args ? args.serialization : undefined;
            resourceInputs["servicebusNamespace"] = args ? args.servicebusNamespace : undefined;
            resourceInputs["sharedAccessPolicyKey"] = args ? args.sharedAccessPolicyKey : undefined;
            resourceInputs["sharedAccessPolicyName"] = args ? args.sharedAccessPolicyName : undefined;
            resourceInputs["streamAnalyticsJobName"] = args ? args.streamAnalyticsJobName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StreamInputEventHub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamInputEventHub resources.
 */
export interface StreamInputEventHubState {
    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub. If not set the input will use the Event Hub's default consumer group.
     */
    eventhubConsumerGroupName?: pulumi.Input<string>;
    /**
     * The name of the Event Hub.
     */
    eventhubName?: pulumi.Input<string>;
    /**
     * The name of the Stream Input EventHub. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The property the input Event Hub has been partitioned by.
     */
    partitionKey?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    serialization?: pulumi.Input<inputs.streamanalytics.StreamInputEventHubSerialization>;
    /**
     * The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    servicebusNamespace?: pulumi.Input<string>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    sharedAccessPolicyKey?: pulumi.Input<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    sharedAccessPolicyName?: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    streamAnalyticsJobName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamInputEventHub resource.
 */
export interface StreamInputEventHubArgs {
    /**
     * The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub. If not set the input will use the Event Hub's default consumer group.
     */
    eventhubConsumerGroupName?: pulumi.Input<string>;
    /**
     * The name of the Event Hub.
     */
    eventhubName: pulumi.Input<string>;
    /**
     * The name of the Stream Input EventHub. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The property the input Event Hub has been partitioned by.
     */
    partitionKey?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `serialization` block as defined below.
     */
    serialization: pulumi.Input<inputs.streamanalytics.StreamInputEventHubSerialization>;
    /**
     * The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    servicebusNamespace: pulumi.Input<string>;
    /**
     * The shared access policy key for the specified shared access policy.
     */
    sharedAccessPolicyKey: pulumi.Input<string>;
    /**
     * The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
     */
    sharedAccessPolicyName: pulumi.Input<string>;
    /**
     * The name of the Stream Analytics Job. Changing this forces a new resource to be created.
     */
    streamAnalyticsJobName: pulumi.Input<string>;
}
