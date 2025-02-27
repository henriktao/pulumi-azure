// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages the hub settings for a Web Pubsub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "east us"});
 * const testUserAssignedIdentity = new azure.authorization.UserAssignedIdentity("testUserAssignedIdentity", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleService = new azure.webpubsub.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard_S1",
 *     capacity: 1,
 * });
 * const testHub = new azure.webpubsub.Hub("testHub", {
 *     webPubsubId: azurerm_web_pubsub.exmaple.id,
 *     eventHandlers: [
 *         {
 *             urlTemplate: "https://test.com/api/{hub}/{event}",
 *             userEventPattern: "*",
 *             systemEvents: [
 *                 "connect",
 *                 "connected",
 *             ],
 *         },
 *         {
 *             urlTemplate: "https://test.com/api/{hub}/{event}",
 *             userEventPattern: "event1, event2",
 *             systemEvents: ["connected"],
 *             auth: {
 *                 managedIdentityId: testUserAssignedIdentity.id,
 *             },
 *         },
 *     ],
 *     anonymousConnectionsEnabled: true,
 * }, {
 *     dependsOn: [azurerm_web_pubsub.test],
 * });
 * ```
 *
 * ## Import
 *
 * Web Pubsub Hub can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:webpubsub/hub:Hub example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubsub/webpubsub1/hubs/webpubsubhub1
 * ```
 */
export class Hub extends pulumi.CustomResource {
    /**
     * Get an existing Hub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HubState, opts?: pulumi.CustomResourceOptions): Hub {
        return new Hub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:webpubsub/hub:Hub';

    /**
     * Returns true if the given object is an instance of Hub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hub.__pulumiType;
    }

    /**
     * Is anonymous connections are allowed for this hub? Defaults to `false`.
     * Possible values are `true`, `false`.
     */
    public readonly anonymousConnectionsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An `eventHandler` block as defined below.
     */
    public readonly eventHandlers!: pulumi.Output<outputs.webpubsub.HubEventHandler[]>;
    /**
     * The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
     */
    public readonly webPubsubId!: pulumi.Output<string>;

    /**
     * Create a Hub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HubArgs | HubState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HubState | undefined;
            resourceInputs["anonymousConnectionsEnabled"] = state ? state.anonymousConnectionsEnabled : undefined;
            resourceInputs["eventHandlers"] = state ? state.eventHandlers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["webPubsubId"] = state ? state.webPubsubId : undefined;
        } else {
            const args = argsOrState as HubArgs | undefined;
            if ((!args || args.eventHandlers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventHandlers'");
            }
            if ((!args || args.webPubsubId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webPubsubId'");
            }
            resourceInputs["anonymousConnectionsEnabled"] = args ? args.anonymousConnectionsEnabled : undefined;
            resourceInputs["eventHandlers"] = args ? args.eventHandlers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["webPubsubId"] = args ? args.webPubsubId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Hub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hub resources.
 */
export interface HubState {
    /**
     * Is anonymous connections are allowed for this hub? Defaults to `false`.
     * Possible values are `true`, `false`.
     */
    anonymousConnectionsEnabled?: pulumi.Input<boolean>;
    /**
     * An `eventHandler` block as defined below.
     */
    eventHandlers?: pulumi.Input<pulumi.Input<inputs.webpubsub.HubEventHandler>[]>;
    /**
     * The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
     */
    webPubsubId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hub resource.
 */
export interface HubArgs {
    /**
     * Is anonymous connections are allowed for this hub? Defaults to `false`.
     * Possible values are `true`, `false`.
     */
    anonymousConnectionsEnabled?: pulumi.Input<boolean>;
    /**
     * An `eventHandler` block as defined below.
     */
    eventHandlers: pulumi.Input<pulumi.Input<inputs.webpubsub.HubEventHandler>[]>;
    /**
     * The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
     */
    webPubsubId: pulumi.Input<string>;
}
