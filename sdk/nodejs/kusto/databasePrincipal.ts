// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Kusto (also known as Azure Data Explorer) Database Principal
 *
 * > **NOTE:** This resource is being **deprecated** due to API updates and should no longer be used.  Please use azure.kusto.DatabasePrincipalAssignment instead.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const cluster = new azure.kusto.Cluster("cluster", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Standard_D13_v2",
 *         capacity: 2,
 *     },
 * });
 * const database = new azure.kusto.Database("database", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: cluster.name,
 *     hotCachePeriod: "P7D",
 *     softDeletePeriod: "P31D",
 * });
 * const principal = new azure.kusto.DatabasePrincipal("principal", {
 *     resourceGroupName: rg.name,
 *     clusterName: cluster.name,
 *     databaseName: azurerm_kusto_database.test.name,
 *     role: "Viewer",
 *     type: "User",
 *     clientId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.clientId),
 * });
 * ```
 *
 * ## Import
 *
 * Kusto Database Principals can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:kusto/databasePrincipal:DatabasePrincipal example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/Role/role1/FQN/some-guid
 * ```
 */
export class DatabasePrincipal extends pulumi.CustomResource {
    /**
     * Get an existing DatabasePrincipal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabasePrincipalState, opts?: pulumi.CustomResourceOptions): DatabasePrincipal {
        return new DatabasePrincipal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:kusto/databasePrincipal:DatabasePrincipal';

    /**
     * Returns true if the given object is an instance of DatabasePrincipal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabasePrincipal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabasePrincipal.__pulumiType;
    }

    /**
     * The app id, if not empty, of the principal.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The email, if not empty, of the principal.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The fully qualified name of the principal.
     */
    public /*out*/ readonly fullyQualifiedName!: pulumi.Output<string>;
    /**
     * The name of the Kusto Database Principal.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
     */
    public readonly objectId!: pulumi.Output<string>;
    /**
     * Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabasePrincipal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabasePrincipalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabasePrincipalArgs | DatabasePrincipalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabasePrincipalState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["fullyQualifiedName"] = state ? state.fullyQualifiedName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DatabasePrincipalArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.objectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["objectId"] = args ? args.objectId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["fullyQualifiedName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabasePrincipal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabasePrincipal resources.
 */
export interface DatabasePrincipalState {
    /**
     * The app id, if not empty, of the principal.
     */
    appId?: pulumi.Input<string>;
    /**
     * The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The email, if not empty, of the principal.
     */
    email?: pulumi.Input<string>;
    /**
     * The fully qualified name of the principal.
     */
    fullyQualifiedName?: pulumi.Input<string>;
    /**
     * The name of the Kusto Database Principal.
     */
    name?: pulumi.Input<string>;
    /**
     * An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
     */
    objectId?: pulumi.Input<string>;
    /**
     * Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
     */
    role?: pulumi.Input<string>;
    /**
     * Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabasePrincipal resource.
 */
export interface DatabasePrincipalArgs {
    /**
     * The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
     */
    clientId: pulumi.Input<string>;
    /**
     * Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
     */
    databaseName: pulumi.Input<string>;
    /**
     * An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
     */
    objectId: pulumi.Input<string>;
    /**
     * Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
     */
    role: pulumi.Input<string>;
    /**
     * Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    type: pulumi.Input<string>;
}
