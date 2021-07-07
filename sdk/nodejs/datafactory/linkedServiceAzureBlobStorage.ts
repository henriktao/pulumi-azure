// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = exampleResourceGroup.name.apply(name => azure.storage.getAccount({
 *     name: "storageaccountname",
 *     resourceGroupName: name,
 * }));
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleLinkedServiceAzureBlobStorage = new azure.datafactory.LinkedServiceAzureBlobStorage("exampleLinkedServiceAzureBlobStorage", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryName: exampleFactory.name,
 *     connectionString: exampleAccount.apply(exampleAccount => exampleAccount.primaryConnectionString),
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory Linked Service's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
 * ```
 */
export class LinkedServiceAzureBlobStorage extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceAzureBlobStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceAzureBlobStorageState, opts?: pulumi.CustomResourceOptions): LinkedServiceAzureBlobStorage {
        return new LinkedServiceAzureBlobStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage';

    /**
     * Returns true if the given object is an instance of LinkedServiceAzureBlobStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceAzureBlobStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceAzureBlobStorage.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
     */
    public readonly connectionString!: pulumi.Output<string | undefined>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    public readonly integrationRuntimeName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
     */
    public readonly sasUri!: pulumi.Output<string | undefined>;
    /**
     * The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
     */
    public readonly serviceEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
     */
    public readonly servicePrincipalId!: pulumi.Output<string | undefined>;
    /**
     * The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
     */
    public readonly servicePrincipalKey!: pulumi.Output<string | undefined>;
    /**
     * The tenant id or name in which to authenticate against the Azure Blob Storage account.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
     */
    public readonly useManagedIdentity!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LinkedServiceAzureBlobStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceAzureBlobStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceAzureBlobStorageArgs | LinkedServiceAzureBlobStorageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkedServiceAzureBlobStorageState | undefined;
            inputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sasUri"] = state ? state.sasUri : undefined;
            inputs["serviceEndpoint"] = state ? state.serviceEndpoint : undefined;
            inputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            inputs["servicePrincipalKey"] = state ? state.servicePrincipalKey : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["useManagedIdentity"] = state ? state.useManagedIdentity : undefined;
        } else {
            const args = argsOrState as LinkedServiceAzureBlobStorageArgs | undefined;
            if ((!args || args.dataFactoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["connectionString"] = args ? args.connectionString : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sasUri"] = args ? args.sasUri : undefined;
            inputs["serviceEndpoint"] = args ? args.serviceEndpoint : undefined;
            inputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
            inputs["servicePrincipalKey"] = args ? args.servicePrincipalKey : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["useManagedIdentity"] = args ? args.useManagedIdentity : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LinkedServiceAzureBlobStorage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceAzureBlobStorage resources.
 */
export interface LinkedServiceAzureBlobStorageState {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
     */
    sasUri?: pulumi.Input<string>;
    /**
     * The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
     */
    serviceEndpoint?: pulumi.Input<string>;
    /**
     * The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
     */
    servicePrincipalKey?: pulumi.Input<string>;
    /**
     * The tenant id or name in which to authenticate against the Azure Blob Storage account.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
     */
    useManagedIdentity?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LinkedServiceAzureBlobStorage resource.
 */
export interface LinkedServiceAzureBlobStorageArgs {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection string. Conflicts with `sasUri` and `serviceEndpoint`.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryName: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The SAS URI. Conflicts with `connectionString` and `serviceEndpoint`.
     */
    sasUri?: pulumi.Input<string>;
    /**
     * The Service Endpoint. Conflicts with `connectionString` and `sasUri`. Required with `useManagedIdentity`.
     */
    serviceEndpoint?: pulumi.Input<string>;
    /**
     * The service principal id in which to authenticate against the Azure Blob Storage account. Required if `servicePrincipalKey` is set.
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `servicePrincipalId` is set.
     */
    servicePrincipalKey?: pulumi.Input<string>;
    /**
     * The tenant id or name in which to authenticate against the Azure Blob Storage account.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `servicePrincipalId` and `servicePrincipalKey`.
     */
    useManagedIdentity?: pulumi.Input<boolean>;
}
