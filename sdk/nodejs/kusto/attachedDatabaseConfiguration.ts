// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Kusto (also known as Azure Data Explorer) Attached Database Configuration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const followerCluster = new azure.kusto.Cluster("followerCluster", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Dev(No SLA)_Standard_D11_v2",
 *         capacity: 1,
 *     },
 * });
 * const followedCluster = new azure.kusto.Cluster("followedCluster", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Dev(No SLA)_Standard_D11_v2",
 *         capacity: 1,
 *     },
 * });
 * const followedDatabase = new azure.kusto.Database("followedDatabase", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: azurerm_kusto_cluster.cluster2.name,
 * });
 * const exampleDatabase = new azure.kusto.Database("exampleDatabase", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: azurerm_kusto_cluster.cluster2.name,
 * });
 * const exampleAttachedDatabaseConfiguration = new azure.kusto.AttachedDatabaseConfiguration("exampleAttachedDatabaseConfiguration", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     clusterName: followerCluster.name,
 *     clusterResourceId: followedCluster.id,
 *     databaseName: exampleDatabase.name,
 *     defaultPrincipalModificationsKind: "Union",
 *     sharing: {
 *         externalTablesToExcludes: ["ExternalTable2"],
 *         externalTablesToIncludes: ["ExternalTable1"],
 *         materializedViewsToExcludes: ["MaterializedViewTable2"],
 *         materializedViewsToIncludes: ["MaterializedViewTable1"],
 *         tablesToExcludes: ["Table2"],
 *         tablesToIncludes: ["Table1"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Kusto Attached Database Configurations can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/AttachedDatabaseConfigurations/configuration1
 * ```
 */
export class AttachedDatabaseConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachedDatabaseConfigurationState, opts?: pulumi.CustomResourceOptions): AttachedDatabaseConfiguration {
        return new AttachedDatabaseConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration';

    /**
     * Returns true if the given object is an instance of AttachedDatabaseConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttachedDatabaseConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttachedDatabaseConfiguration.__pulumiType;
    }

    /**
     * The list of databases from the `clusterResourceId` which are currently attached to the cluster.
     */
    public /*out*/ readonly attachedDatabaseNames!: pulumi.Output<string[]>;
    /**
     * Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The resource id of the cluster where the databases you would like to attach reside.
     */
    public readonly clusterResourceId!: pulumi.Output<string>;
    /**
     * The name of the database which you would like to attach, use * if you want to follow all current and future databases.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
     */
    public readonly defaultPrincipalModificationKind!: pulumi.Output<string | undefined>;
    /**
     * Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `sharing` block as defined below.
     */
    public readonly sharing!: pulumi.Output<outputs.kusto.AttachedDatabaseConfigurationSharing | undefined>;

    /**
     * Create a AttachedDatabaseConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachedDatabaseConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachedDatabaseConfigurationArgs | AttachedDatabaseConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttachedDatabaseConfigurationState | undefined;
            resourceInputs["attachedDatabaseNames"] = state ? state.attachedDatabaseNames : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["clusterResourceId"] = state ? state.clusterResourceId : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["defaultPrincipalModificationKind"] = state ? state.defaultPrincipalModificationKind : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["sharing"] = state ? state.sharing : undefined;
        } else {
            const args = argsOrState as AttachedDatabaseConfigurationArgs | undefined;
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.clusterResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterResourceId'");
            }
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["clusterResourceId"] = args ? args.clusterResourceId : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["defaultPrincipalModificationKind"] = args ? args.defaultPrincipalModificationKind : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sharing"] = args ? args.sharing : undefined;
            resourceInputs["attachedDatabaseNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AttachedDatabaseConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttachedDatabaseConfiguration resources.
 */
export interface AttachedDatabaseConfigurationState {
    /**
     * The list of databases from the `clusterResourceId` which are currently attached to the cluster.
     */
    attachedDatabaseNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * The resource id of the cluster where the databases you would like to attach reside.
     */
    clusterResourceId?: pulumi.Input<string>;
    /**
     * The name of the database which you would like to attach, use * if you want to follow all current and future databases.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
     */
    defaultPrincipalModificationKind?: pulumi.Input<string>;
    /**
     * Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sharing` block as defined below.
     */
    sharing?: pulumi.Input<inputs.kusto.AttachedDatabaseConfigurationSharing>;
}

/**
 * The set of arguments for constructing a AttachedDatabaseConfiguration resource.
 */
export interface AttachedDatabaseConfigurationArgs {
    /**
     * Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The resource id of the cluster where the databases you would like to attach reside.
     */
    clusterResourceId: pulumi.Input<string>;
    /**
     * The name of the database which you would like to attach, use * if you want to follow all current and future databases.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
     */
    defaultPrincipalModificationKind?: pulumi.Input<string>;
    /**
     * Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `sharing` block as defined below.
     */
    sharing?: pulumi.Input<inputs.kusto.AttachedDatabaseConfigurationSharing>;
}
