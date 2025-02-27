// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure SignalR service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleService = new azure.signalr.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         name: "Free_F1",
 *         capacity: 1,
 *     },
 *     cors: [{
 *         allowedOrigins: ["http://www.example.com"],
 *     }],
 *     connectivityLogsEnabled: true,
 *     messagingLogsEnabled: true,
 *     serviceMode: "Default",
 *     upstreamEndpoints: [{
 *         categoryPatterns: [
 *             "connections",
 *             "messages",
 *         ],
 *         eventPatterns: ["*"],
 *         hubPatterns: ["hub1"],
 *         urlTemplate: "http://foo.com",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * SignalR services can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:signalr/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/signalR/tfex-signalr
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:signalr/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Specifies if Connectivity Logs are enabled or not.
     */
    public readonly connectivityLogsEnabled!: pulumi.Output<boolean>;
    /**
     * A `cors` block as documented below.
     */
    public readonly cors!: pulumi.Output<outputs.signalr.ServiceCor[]>;
    /**
     * A `features` block as documented below.
     *
     * @deprecated Deprecated in favour of `connectivity_logs_enabled`, `messaging_logs_enabled`, `live_trace_enabled` and `service_mode`
     */
    public readonly features!: pulumi.Output<outputs.signalr.ServiceFeature[]>;
    /**
     * The FQDN of the SignalR service.
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Specifies if Live Trace is enabled or not.
     */
    public readonly liveTraceEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies if Messaging Logs are enabled or not.
     */
    public readonly messagingLogsEnabled!: pulumi.Output<boolean>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary access key for the SignalR service.
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * The primary connection string for the SignalR service.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client use.
     */
    public /*out*/ readonly publicPort!: pulumi.Output<number>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The secondary access key for the SignalR service.
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * The secondary connection string for the SignalR service.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side use.
     */
    public /*out*/ readonly serverPort!: pulumi.Output<number>;
    /**
     * Specifies the service mode. Possible values are `Classic`, `Default` and `Serverless`.
     */
    public readonly serviceMode!: pulumi.Output<string>;
    /**
     * A `sku` block as documented below.
     */
    public readonly sku!: pulumi.Output<outputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
     */
    public readonly upstreamEndpoints!: pulumi.Output<outputs.signalr.ServiceUpstreamEndpoint[] | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["connectivityLogsEnabled"] = state ? state.connectivityLogsEnabled : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["liveTraceEnabled"] = state ? state.liveTraceEnabled : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["messagingLogsEnabled"] = state ? state.messagingLogsEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            resourceInputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            resourceInputs["publicPort"] = state ? state.publicPort : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            resourceInputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            resourceInputs["serverPort"] = state ? state.serverPort : undefined;
            resourceInputs["serviceMode"] = state ? state.serviceMode : undefined;
            resourceInputs["sku"] = state ? state.sku : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["upstreamEndpoints"] = state ? state.upstreamEndpoints : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            resourceInputs["connectivityLogsEnabled"] = args ? args.connectivityLogsEnabled : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["liveTraceEnabled"] = args ? args.liveTraceEnabled : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["messagingLogsEnabled"] = args ? args.messagingLogsEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serviceMode"] = args ? args.serviceMode : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["upstreamEndpoints"] = args ? args.upstreamEndpoints : undefined;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["primaryAccessKey"] = undefined /*out*/;
            resourceInputs["primaryConnectionString"] = undefined /*out*/;
            resourceInputs["publicPort"] = undefined /*out*/;
            resourceInputs["secondaryAccessKey"] = undefined /*out*/;
            resourceInputs["secondaryConnectionString"] = undefined /*out*/;
            resourceInputs["serverPort"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * Specifies if Connectivity Logs are enabled or not.
     */
    connectivityLogsEnabled?: pulumi.Input<boolean>;
    /**
     * A `cors` block as documented below.
     */
    cors?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceCor>[]>;
    /**
     * A `features` block as documented below.
     *
     * @deprecated Deprecated in favour of `connectivity_logs_enabled`, `messaging_logs_enabled`, `live_trace_enabled` and `service_mode`
     */
    features?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceFeature>[]>;
    /**
     * The FQDN of the SignalR service.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The publicly accessible IP of the SignalR service.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Specifies if Live Trace is enabled or not.
     */
    liveTraceEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies if Messaging Logs are enabled or not.
     */
    messagingLogsEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The primary access key for the SignalR service.
     */
    primaryAccessKey?: pulumi.Input<string>;
    /**
     * The primary connection string for the SignalR service.
     */
    primaryConnectionString?: pulumi.Input<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for browser/client use.
     */
    publicPort?: pulumi.Input<number>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The secondary access key for the SignalR service.
     */
    secondaryAccessKey?: pulumi.Input<string>;
    /**
     * The secondary connection string for the SignalR service.
     */
    secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The publicly accessible port of the SignalR service which is designed for customer server side use.
     */
    serverPort?: pulumi.Input<number>;
    /**
     * Specifies the service mode. Possible values are `Classic`, `Default` and `Serverless`.
     */
    serviceMode?: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    sku?: pulumi.Input<inputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
     */
    upstreamEndpoints?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceUpstreamEndpoint>[]>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Specifies if Connectivity Logs are enabled or not.
     */
    connectivityLogsEnabled?: pulumi.Input<boolean>;
    /**
     * A `cors` block as documented below.
     */
    cors?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceCor>[]>;
    /**
     * A `features` block as documented below.
     *
     * @deprecated Deprecated in favour of `connectivity_logs_enabled`, `messaging_logs_enabled`, `live_trace_enabled` and `service_mode`
     */
    features?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceFeature>[]>;
    /**
     * Specifies if Live Trace is enabled or not.
     */
    liveTraceEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies if Messaging Logs are enabled or not.
     */
    messagingLogsEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the SignalR service. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the service mode. Possible values are `Classic`, `Default` and `Serverless`.
     */
    serviceMode?: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    sku: pulumi.Input<inputs.signalr.ServiceSku>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `upstreamEndpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
     */
    upstreamEndpoints?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceUpstreamEndpoint>[]>;
}
