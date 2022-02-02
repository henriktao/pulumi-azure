// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Synapse Managed Private Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     isHnsEnabled: "true",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleAccount.id});
 * const exampleWorkspace = new azure.synapse.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     storageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.id,
 *     sqlAdministratorLogin: "sqladminuser",
 *     sqlAdministratorLoginPassword: "H@Sh1CoR3!",
 *     managedVirtualNetworkEnabled: true,
 * });
 * const exampleFirewallRule = new azure.synapse.FirewallRule("exampleFirewallRule", {
 *     synapseWorkspaceId: azurerm_synapse_workspace.test.id,
 *     startIpAddress: "0.0.0.0",
 *     endIpAddress: "255.255.255.255",
 * });
 * const exampleConnect = new azure.storage.Account("exampleConnect", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "BlobStorage",
 * });
 * const exampleManagedPrivateEndpoint = new azure.synapse.ManagedPrivateEndpoint("exampleManagedPrivateEndpoint", {
 *     synapseWorkspaceId: exampleWorkspace.id,
 *     targetResourceId: exampleConnect.id,
 *     subresourceName: "blob",
 * }, {
 *     dependsOn: [exampleFirewallRule],
 * });
 * ```
 *
 * ## Import
 *
 * Synapse Managed Private Endpoint can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
 * ```
 */
export class ManagedPrivateEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedPrivateEndpointState, opts?: pulumi.CustomResourceOptions): ManagedPrivateEndpoint {
        return new ManagedPrivateEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/managedPrivateEndpoint:ManagedPrivateEndpoint';

    /**
     * Returns true if the given object is an instance of ManagedPrivateEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedPrivateEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedPrivateEndpoint.__pulumiType;
    }

    /**
     * Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
     */
    public readonly subresourceName!: pulumi.Output<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    public readonly synapseWorkspaceId!: pulumi.Output<string>;
    /**
     * The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
     */
    public readonly targetResourceId!: pulumi.Output<string>;

    /**
     * Create a ManagedPrivateEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedPrivateEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedPrivateEndpointArgs | ManagedPrivateEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedPrivateEndpointState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["subresourceName"] = state ? state.subresourceName : undefined;
            resourceInputs["synapseWorkspaceId"] = state ? state.synapseWorkspaceId : undefined;
            resourceInputs["targetResourceId"] = state ? state.targetResourceId : undefined;
        } else {
            const args = argsOrState as ManagedPrivateEndpointArgs | undefined;
            if ((!args || args.subresourceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subresourceName'");
            }
            if ((!args || args.synapseWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'synapseWorkspaceId'");
            }
            if ((!args || args.targetResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subresourceName"] = args ? args.subresourceName : undefined;
            resourceInputs["synapseWorkspaceId"] = args ? args.synapseWorkspaceId : undefined;
            resourceInputs["targetResourceId"] = args ? args.targetResourceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ManagedPrivateEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
 */
export interface ManagedPrivateEndpointState {
    /**
     * Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
     */
    subresourceName?: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    synapseWorkspaceId?: pulumi.Input<string>;
    /**
     * The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
     */
    targetResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedPrivateEndpoint resource.
 */
export interface ManagedPrivateEndpointArgs {
    /**
     * Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the sub resource name which the Synapse Private Endpoint is able to connect to. Changing this forces a new resource to be created.
     */
    subresourceName: pulumi.Input<string>;
    /**
     * The ID of the Synapse Workspace on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
     */
    synapseWorkspaceId: pulumi.Input<string>;
    /**
     * The ID of the Private Link Enabled Remote Resource which this Synapse Private Endpoint should be connected to. Changing this forces a new resource to be created.
     */
    targetResourceId: pulumi.Input<string>;
}
