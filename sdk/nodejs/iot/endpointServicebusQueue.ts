// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an IotHub ServiceBus Queue Endpoint
 *
 * > **NOTE:** Endpoints can be defined either directly on the `azure.iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `azure.iot.IoTHub` resource is not supported.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNamespace = new azure.servicebus.Namespace("exampleNamespace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 * });
 * const exampleQueue = new azure.servicebus.Queue("exampleQueue", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     namespaceName: exampleNamespace.name,
 *     enablePartitioning: true,
 * });
 * const exampleQueueAuthorizationRule = new azure.servicebus.QueueAuthorizationRule("exampleQueueAuthorizationRule", {
 *     namespaceName: exampleNamespace.name,
 *     queueName: exampleQueue.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     listen: false,
 *     send: true,
 *     manage: false,
 * });
 * const exampleIoTHub = new azure.iot.IoTHub("exampleIoTHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: {
 *         name: "B1",
 *         capacity: "1",
 *     },
 *     tags: {
 *         purpose: "example",
 *     },
 * });
 * const exampleEndpointServicebusQueue = new azure.iot.EndpointServicebusQueue("exampleEndpointServicebusQueue", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     iothubName: exampleIoTHub.name,
 *     connectionString: exampleQueueAuthorizationRule.primaryConnectionString,
 * });
 * ```
 *
 * ## Import
 *
 * IoTHub ServiceBus Queue Endpoint can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:iot/endpointServicebusQueue:EndpointServicebusQueue servicebus_queue1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/servicebusqueue_endpoint1
 * ```
 */
export class EndpointServicebusQueue extends pulumi.CustomResource {
    /**
     * Get an existing EndpointServicebusQueue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointServicebusQueueState, opts?: pulumi.CustomResourceOptions): EndpointServicebusQueue {
        return new EndpointServicebusQueue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:iot/endpointServicebusQueue:EndpointServicebusQueue';

    /**
     * Returns true if the given object is an instance of EndpointServicebusQueue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointServicebusQueue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointServicebusQueue.__pulumiType;
    }

    /**
     * Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
     */
    public readonly authenticationType!: pulumi.Output<string | undefined>;
    /**
     * The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
     */
    public readonly connectionString!: pulumi.Output<string | undefined>;
    /**
     * URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    public readonly endpointUri!: pulumi.Output<string | undefined>;
    /**
     * Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    public readonly entityPath!: pulumi.Output<string | undefined>;
    /**
     * ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
     */
    public readonly identityId!: pulumi.Output<string | undefined>;
    /**
     * The IoTHub ID for the endpoint.
     */
    public readonly iothubId!: pulumi.Output<string>;
    /**
     * The IoTHub name for the endpoint.
     *
     * @deprecated Deprecated in favour of `iothub_id`
     */
    public readonly iothubName!: pulumi.Output<string>;
    /**
     * The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a EndpointServicebusQueue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointServicebusQueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointServicebusQueueArgs | EndpointServicebusQueueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointServicebusQueueState | undefined;
            inputs["authenticationType"] = state ? state.authenticationType : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["endpointUri"] = state ? state.endpointUri : undefined;
            inputs["entityPath"] = state ? state.entityPath : undefined;
            inputs["identityId"] = state ? state.identityId : undefined;
            inputs["iothubId"] = state ? state.iothubId : undefined;
            inputs["iothubName"] = state ? state.iothubName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as EndpointServicebusQueueArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["authenticationType"] = args ? args.authenticationType : undefined;
            inputs["connectionString"] = args ? args.connectionString : undefined;
            inputs["endpointUri"] = args ? args.endpointUri : undefined;
            inputs["entityPath"] = args ? args.entityPath : undefined;
            inputs["identityId"] = args ? args.identityId : undefined;
            inputs["iothubId"] = args ? args.iothubId : undefined;
            inputs["iothubName"] = args ? args.iothubName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EndpointServicebusQueue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointServicebusQueue resources.
 */
export interface EndpointServicebusQueueState {
    /**
     * Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
     */
    authenticationType?: pulumi.Input<string>;
    /**
     * The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    endpointUri?: pulumi.Input<string>;
    /**
     * Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    entityPath?: pulumi.Input<string>;
    /**
     * ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
     */
    identityId?: pulumi.Input<string>;
    /**
     * The IoTHub ID for the endpoint.
     */
    iothubId?: pulumi.Input<string>;
    /**
     * The IoTHub name for the endpoint.
     *
     * @deprecated Deprecated in favour of `iothub_id`
     */
    iothubName?: pulumi.Input<string>;
    /**
     * The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointServicebusQueue resource.
 */
export interface EndpointServicebusQueueArgs {
    /**
     * Type used to authenticate against the Service Bus Queue endpoint. Possible values are `keyBased` and `identityBased`. Defaults to `keyBased`.
     */
    authenticationType?: pulumi.Input<string>;
    /**
     * The connection string for the endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `keyBased`.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * URI of the Service Bus endpoint. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    endpointUri?: pulumi.Input<string>;
    /**
     * Name of the Service Bus Queue. This attribute can only be specified and is mandatory when `authenticationType` is `identityBased`.
     */
    entityPath?: pulumi.Input<string>;
    /**
     * ID of the User Managed Identity used to authenticate against the Service Bus Queue endpoint.
     */
    identityId?: pulumi.Input<string>;
    /**
     * The IoTHub ID for the endpoint.
     */
    iothubId?: pulumi.Input<string>;
    /**
     * The IoTHub name for the endpoint.
     *
     * @deprecated Deprecated in favour of `iothub_id`
     */
    iothubName?: pulumi.Input<string>;
    /**
     * The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the Service Bus Queue has been created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
}
