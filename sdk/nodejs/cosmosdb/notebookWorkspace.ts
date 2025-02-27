// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an SQL Notebook Workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.cosmosdb.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     offerType: "Standard",
 *     kind: "GlobalDocumentDB",
 *     consistencyPolicy: {
 *         consistencyLevel: "BoundedStaleness",
 *     },
 *     geoLocations: [{
 *         location: exampleResourceGroup.location,
 *         failoverPriority: 0,
 *     }],
 * });
 * const exampleNotebookWorkspace = new azure.cosmosdb.NotebookWorkspace("exampleNotebookWorkspace", {
 *     resourceGroupName: exampleAccount.resourceGroupName,
 *     accountName: exampleAccount.name,
 * });
 * ```
 *
 * ## Import
 *
 * =SQL Notebook Workspaces can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:cosmosdb/notebookWorkspace:NotebookWorkspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/notebookWorkspaces/notebookWorkspace1
 * ```
 */
export class NotebookWorkspace extends pulumi.CustomResource {
    /**
     * Get an existing NotebookWorkspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotebookWorkspaceState, opts?: pulumi.CustomResourceOptions): NotebookWorkspace {
        return new NotebookWorkspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cosmosdb/notebookWorkspace:NotebookWorkspace';

    /**
     * Returns true if the given object is an instance of NotebookWorkspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotebookWorkspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotebookWorkspace.__pulumiType;
    }

    /**
     * The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the endpoint of Notebook server.
     */
    public /*out*/ readonly serverEndpoint!: pulumi.Output<string>;

    /**
     * Create a NotebookWorkspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotebookWorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotebookWorkspaceArgs | NotebookWorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotebookWorkspaceState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["serverEndpoint"] = state ? state.serverEndpoint : undefined;
        } else {
            const args = argsOrState as NotebookWorkspaceArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serverEndpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotebookWorkspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotebookWorkspace resources.
 */
export interface NotebookWorkspaceState {
    /**
     * The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the endpoint of Notebook server.
     */
    serverEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotebookWorkspace resource.
 */
export interface NotebookWorkspaceArgs {
    /**
     * The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
     */
    resourceGroupName: pulumi.Input<string>;
}
