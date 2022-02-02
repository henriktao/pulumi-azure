// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Linked Service (connection) between a Database and Azure Data Factory through ODBC protocol.
 *
 * > **Note:** All arguments including the connectionString will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleFactory = new azure.datafactory.Factory("exampleFactory", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const anonymous = new azure.datafactory.LinkedServiceOdbc("anonymous", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryId: exampleFactory.id,
 *     connectionString: "Driver={SQL Server};Server=test;Database=test;Uid=test;Pwd=test;",
 * });
 * const basicAuth = new azure.datafactory.LinkedServiceOdbc("basicAuth", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     dataFactoryId: exampleFactory.id,
 *     connectionString: "Driver={SQL Server};Server=test;Database=test;Uid=test;Pwd=test;",
 *     basicAuthentication: {
 *         username: "onrylmz",
 *         password: "Ch4ngeM3!",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Data Factory ODBC Linked Service's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/linkedServiceOdbc:LinkedServiceOdbc example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
 * ```
 */
export class LinkedServiceOdbc extends pulumi.CustomResource {
    /**
     * Get an existing LinkedServiceOdbc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinkedServiceOdbcState, opts?: pulumi.CustomResourceOptions): LinkedServiceOdbc {
        return new LinkedServiceOdbc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/linkedServiceOdbc:LinkedServiceOdbc';

    /**
     * Returns true if the given object is an instance of LinkedServiceOdbc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedServiceOdbc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedServiceOdbc.__pulumiType;
    }

    /**
     * A map of additional properties to associate with the Data Factory Linked Service ODBC.
     */
    public readonly additionalProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service ODBC.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * A `basicAuthentication` block as defined below.
     */
    public readonly basicAuthentication!: pulumi.Output<outputs.datafactory.LinkedServiceOdbcBasicAuthentication | undefined>;
    /**
     * The connection string in which to authenticate with ODBC.
     */
    public readonly connectionString!: pulumi.Output<string>;
    /**
     * The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
     */
    public readonly dataFactoryId!: pulumi.Output<string>;
    /**
     * The description for the Data Factory Linked Service ODBC.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service ODBC.
     */
    public readonly integrationRuntimeName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service ODBC.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a LinkedServiceOdbc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceOdbcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkedServiceOdbcArgs | LinkedServiceOdbcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LinkedServiceOdbcState | undefined;
            resourceInputs["additionalProperties"] = state ? state.additionalProperties : undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["basicAuthentication"] = state ? state.basicAuthentication : undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["dataFactoryId"] = state ? state.dataFactoryId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["integrationRuntimeName"] = state ? state.integrationRuntimeName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as LinkedServiceOdbcArgs | undefined;
            if ((!args || args.connectionString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionString'");
            }
            if ((!args || args.dataFactoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFactoryId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["additionalProperties"] = args ? args.additionalProperties : undefined;
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["basicAuthentication"] = args ? args.basicAuthentication : undefined;
            resourceInputs["connectionString"] = args ? args.connectionString : undefined;
            resourceInputs["dataFactoryId"] = args ? args.dataFactoryId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["integrationRuntimeName"] = args ? args.integrationRuntimeName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LinkedServiceOdbc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinkedServiceOdbc resources.
 */
export interface LinkedServiceOdbcState {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service ODBC.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service ODBC.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `basicAuthentication` block as defined below.
     */
    basicAuthentication?: pulumi.Input<inputs.datafactory.LinkedServiceOdbcBasicAuthentication>;
    /**
     * The connection string in which to authenticate with ODBC.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryId?: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service ODBC.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service ODBC.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service ODBC.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinkedServiceOdbc resource.
 */
export interface LinkedServiceOdbcArgs {
    /**
     * A map of additional properties to associate with the Data Factory Linked Service ODBC.
     */
    additionalProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of tags that can be used for describing the Data Factory Linked Service ODBC.
     */
    annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `basicAuthentication` block as defined below.
     */
    basicAuthentication?: pulumi.Input<inputs.datafactory.LinkedServiceOdbcBasicAuthentication>;
    /**
     * The connection string in which to authenticate with ODBC.
     */
    connectionString: pulumi.Input<string>;
    /**
     * The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
     */
    dataFactoryId: pulumi.Input<string>;
    /**
     * The description for the Data Factory Linked Service ODBC.
     */
    description?: pulumi.Input<string>;
    /**
     * The integration runtime reference to associate with the Data Factory Linked Service ODBC.
     */
    integrationRuntimeName?: pulumi.Input<string>;
    /**
     * Specifies the name of the Data Factory Linked Service ODBC. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of parameters to associate with the Data Factory Linked Service ODBC.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Linked Service ODBC. Changing this forces a new resource
     */
    resourceGroupName: pulumi.Input<string>;
}
