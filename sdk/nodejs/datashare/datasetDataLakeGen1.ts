// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Data Share Data Lake Gen1 Dataset.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.datashare.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleShare = new azure.datashare.Share("exampleShare", {
 *     accountId: exampleAccount.id,
 *     kind: "CopyBased",
 * });
 * const exampleStore = new azure.datalake.Store("exampleStore", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     firewallState: "Disabled",
 * });
 * const exampleStoreFile = new azure.datalake.StoreFile("exampleStoreFile", {
 *     accountName: exampleStore.name,
 *     localFilePath: "./example/myfile.txt",
 *     remoteFilePath: "/example/myfile.txt",
 * });
 * const exampleServicePrincipal = exampleAccount.name.apply(name => azuread.getServicePrincipal({
 *     displayName: name,
 * }));
 * const exampleAssignment = new azure.authorization.Assignment("exampleAssignment", {
 *     scope: exampleStore.id,
 *     roleDefinitionName: "Owner",
 *     principalId: exampleServicePrincipal.apply(exampleServicePrincipal => exampleServicePrincipal.objectId),
 * });
 * const exampleDatasetDataLakeGen1 = new azure.datashare.DatasetDataLakeGen1("exampleDatasetDataLakeGen1", {
 *     dataShareId: exampleShare.id,
 *     dataLakeStoreId: exampleStore.id,
 *     fileName: "myfile.txt",
 *     folderPath: "example",
 * }, {
 *     dependsOn: [exampleAssignment],
 * });
 * ```
 *
 * ## Import
 *
 * Data Share Data Lake Gen1 Datasets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
 * ```
 */
export class DatasetDataLakeGen1 extends pulumi.CustomResource {
    /**
     * Get an existing DatasetDataLakeGen1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetDataLakeGen1State, opts?: pulumi.CustomResourceOptions): DatasetDataLakeGen1 {
        return new DatasetDataLakeGen1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1';

    /**
     * Returns true if the given object is an instance of DatasetDataLakeGen1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetDataLakeGen1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetDataLakeGen1.__pulumiType;
    }

    /**
     * The resource ID of the Data Lake Store to be shared with the receiver.
     */
    public readonly dataLakeStoreId!: pulumi.Output<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    public readonly dataShareId!: pulumi.Output<string>;
    /**
     * The displayed name of the Data Share Dataset.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    public readonly fileName!: pulumi.Output<string | undefined>;
    /**
     * The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    public readonly folderPath!: pulumi.Output<string>;
    /**
     * The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a DatasetDataLakeGen1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetDataLakeGen1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetDataLakeGen1Args | DatasetDataLakeGen1State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatasetDataLakeGen1State | undefined;
            inputs["dataLakeStoreId"] = state ? state.dataLakeStoreId : undefined;
            inputs["dataShareId"] = state ? state.dataShareId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["fileName"] = state ? state.fileName : undefined;
            inputs["folderPath"] = state ? state.folderPath : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DatasetDataLakeGen1Args | undefined;
            if ((!args || args.dataLakeStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataLakeStoreId'");
            }
            if ((!args || args.dataShareId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataShareId'");
            }
            if ((!args || args.folderPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderPath'");
            }
            inputs["dataLakeStoreId"] = args ? args.dataLakeStoreId : undefined;
            inputs["dataShareId"] = args ? args.dataShareId : undefined;
            inputs["fileName"] = args ? args.fileName : undefined;
            inputs["folderPath"] = args ? args.folderPath : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["displayName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetDataLakeGen1.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetDataLakeGen1 resources.
 */
export interface DatasetDataLakeGen1State {
    /**
     * The resource ID of the Data Lake Store to be shared with the receiver.
     */
    dataLakeStoreId?: pulumi.Input<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    dataShareId?: pulumi.Input<string>;
    /**
     * The displayed name of the Data Share Dataset.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    fileName?: pulumi.Input<string>;
    /**
     * The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    folderPath?: pulumi.Input<string>;
    /**
     * The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatasetDataLakeGen1 resource.
 */
export interface DatasetDataLakeGen1Args {
    /**
     * The resource ID of the Data Lake Store to be shared with the receiver.
     */
    dataLakeStoreId: pulumi.Input<string>;
    /**
     * The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    dataShareId: pulumi.Input<string>;
    /**
     * The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    fileName?: pulumi.Input<string>;
    /**
     * The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    folderPath: pulumi.Input<string>;
    /**
     * The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
     */
    name?: pulumi.Input<string>;
}
