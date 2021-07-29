// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages the Network ACL for a SignalR service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.signalr.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         name: "Standard_S1",
 *         capacity: 1,
 *     },
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     addressSpaces: ["10.5.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.5.2.0/24"],
 *     enforcePrivateLinkEndpointNetworkPolicies: true,
 * });
 * const exampleEndpoint = new azure.privatelink.Endpoint("exampleEndpoint", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     subnetId: exampleSubnet.id,
 *     privateServiceConnection: {
 *         name: "psc-sig-test",
 *         isManualConnection: false,
 *         privateConnectionResourceId: exampleService.id,
 *         subresourceNames: ["signalr"],
 *     },
 * });
 * const exampleServiceNetworkAcl = new azure.signalr.ServiceNetworkAcl("exampleServiceNetworkAcl", {
 *     signalrServiceId: exampleService.id,
 *     defaultAction: "Deny",
 *     publicNetwork: {
 *         allowedRequestTypes: ["ClientConnection"],
 *     },
 *     privateEndpoints: [{
 *         id: exampleEndpoint.id,
 *         allowedRequestTypes: ["ServerConnection"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Network ACLs for a SignalR service can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:signalr/serviceNetworkAcl:ServiceNetworkAcl example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/SignalR/signalr1
 * ```
 */
export class ServiceNetworkAcl extends pulumi.CustomResource {
    /**
     * Get an existing ServiceNetworkAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceNetworkAclState, opts?: pulumi.CustomResourceOptions): ServiceNetworkAcl {
        return new ServiceNetworkAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:signalr/serviceNetworkAcl:ServiceNetworkAcl';

    /**
     * Returns true if the given object is an instance of ServiceNetworkAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceNetworkAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceNetworkAcl.__pulumiType;
    }

    /**
     * The default action to control the network access when no other rule matches. Possible values are `Allow` and `Deny`.
     */
    public readonly defaultAction!: pulumi.Output<string>;
    /**
     * A `privateEndpoint` block as defined below.
     */
    public readonly privateEndpoints!: pulumi.Output<outputs.signalr.ServiceNetworkAclPrivateEndpoint[] | undefined>;
    /**
     * A `publicNetwork` block as defined below.
     */
    public readonly publicNetwork!: pulumi.Output<outputs.signalr.ServiceNetworkAclPublicNetwork>;
    /**
     * The ID of the SignalR service. Changing this forces a new resource to be created.
     */
    public readonly signalrServiceId!: pulumi.Output<string>;

    /**
     * Create a ServiceNetworkAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceNetworkAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceNetworkAclArgs | ServiceNetworkAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceNetworkAclState | undefined;
            inputs["defaultAction"] = state ? state.defaultAction : undefined;
            inputs["privateEndpoints"] = state ? state.privateEndpoints : undefined;
            inputs["publicNetwork"] = state ? state.publicNetwork : undefined;
            inputs["signalrServiceId"] = state ? state.signalrServiceId : undefined;
        } else {
            const args = argsOrState as ServiceNetworkAclArgs | undefined;
            if ((!args || args.defaultAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if ((!args || args.publicNetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicNetwork'");
            }
            if ((!args || args.signalrServiceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signalrServiceId'");
            }
            inputs["defaultAction"] = args ? args.defaultAction : undefined;
            inputs["privateEndpoints"] = args ? args.privateEndpoints : undefined;
            inputs["publicNetwork"] = args ? args.publicNetwork : undefined;
            inputs["signalrServiceId"] = args ? args.signalrServiceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceNetworkAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceNetworkAcl resources.
 */
export interface ServiceNetworkAclState {
    /**
     * The default action to control the network access when no other rule matches. Possible values are `Allow` and `Deny`.
     */
    defaultAction?: pulumi.Input<string>;
    /**
     * A `privateEndpoint` block as defined below.
     */
    privateEndpoints?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceNetworkAclPrivateEndpoint>[]>;
    /**
     * A `publicNetwork` block as defined below.
     */
    publicNetwork?: pulumi.Input<inputs.signalr.ServiceNetworkAclPublicNetwork>;
    /**
     * The ID of the SignalR service. Changing this forces a new resource to be created.
     */
    signalrServiceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceNetworkAcl resource.
 */
export interface ServiceNetworkAclArgs {
    /**
     * The default action to control the network access when no other rule matches. Possible values are `Allow` and `Deny`.
     */
    defaultAction: pulumi.Input<string>;
    /**
     * A `privateEndpoint` block as defined below.
     */
    privateEndpoints?: pulumi.Input<pulumi.Input<inputs.signalr.ServiceNetworkAclPrivateEndpoint>[]>;
    /**
     * A `publicNetwork` block as defined below.
     */
    publicNetwork: pulumi.Input<inputs.signalr.ServiceNetworkAclPublicNetwork>;
    /**
     * The ID of the SignalR service. Changing this forces a new resource to be created.
     */
    signalrServiceId: pulumi.Input<string>;
}
