// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Data Factory (Version 2).
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
 * ```
 *
 * ## Import
 *
 * Data Factory can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:datafactory/factory:Factory example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example
 * ```
 */
export class Factory extends pulumi.CustomResource {
    /**
     * Get an existing Factory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FactoryState, opts?: pulumi.CustomResourceOptions): Factory {
        return new Factory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/factory:Factory';

    /**
     * Returns true if the given object is an instance of Factory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Factory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Factory.__pulumiType;
    }

    /**
     * Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
     */
    public readonly customerManagedKeyId!: pulumi.Output<string | undefined>;
    /**
     * A `githubConfiguration` block as defined below.
     */
    public readonly githubConfiguration!: pulumi.Output<outputs.datafactory.FactoryGithubConfiguration | undefined>;
    /**
     * A list of `globalParameter` blocks as defined above.
     */
    public readonly globalParameters!: pulumi.Output<outputs.datafactory.FactoryGlobalParameter[] | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.datafactory.FactoryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Is Managed Virtual Network enabled?
     */
    public readonly managedVirtualNetworkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Is the Data Factory visible to the public network? Defaults to `true`.
     */
    public readonly publicNetworkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `vstsConfiguration` block as defined below.
     */
    public readonly vstsConfiguration!: pulumi.Output<outputs.datafactory.FactoryVstsConfiguration | undefined>;

    /**
     * Create a Factory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FactoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FactoryArgs | FactoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FactoryState | undefined;
            resourceInputs["customerManagedKeyId"] = state ? state.customerManagedKeyId : undefined;
            resourceInputs["githubConfiguration"] = state ? state.githubConfiguration : undefined;
            resourceInputs["globalParameters"] = state ? state.globalParameters : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managedVirtualNetworkEnabled"] = state ? state.managedVirtualNetworkEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicNetworkEnabled"] = state ? state.publicNetworkEnabled : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vstsConfiguration"] = state ? state.vstsConfiguration : undefined;
        } else {
            const args = argsOrState as FactoryArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["customerManagedKeyId"] = args ? args.customerManagedKeyId : undefined;
            resourceInputs["githubConfiguration"] = args ? args.githubConfiguration : undefined;
            resourceInputs["globalParameters"] = args ? args.globalParameters : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedVirtualNetworkEnabled"] = args ? args.managedVirtualNetworkEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicNetworkEnabled"] = args ? args.publicNetworkEnabled : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vstsConfiguration"] = args ? args.vstsConfiguration : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Factory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Factory resources.
 */
export interface FactoryState {
    /**
     * Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
     */
    customerManagedKeyId?: pulumi.Input<string>;
    /**
     * A `githubConfiguration` block as defined below.
     */
    githubConfiguration?: pulumi.Input<inputs.datafactory.FactoryGithubConfiguration>;
    /**
     * A list of `globalParameter` blocks as defined above.
     */
    globalParameters?: pulumi.Input<pulumi.Input<inputs.datafactory.FactoryGlobalParameter>[]>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.datafactory.FactoryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Is Managed Virtual Network enabled?
     */
    managedVirtualNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * Is the Data Factory visible to the public network? Defaults to `true`.
     */
    publicNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Data Factory.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `vstsConfiguration` block as defined below.
     */
    vstsConfiguration?: pulumi.Input<inputs.datafactory.FactoryVstsConfiguration>;
}

/**
 * The set of arguments for constructing a Factory resource.
 */
export interface FactoryArgs {
    /**
     * Specifies the Azure Key Vault Key ID to be used as the Customer Managed Key (CMK) for double encryption. Required with user assigned identity.
     */
    customerManagedKeyId?: pulumi.Input<string>;
    /**
     * A `githubConfiguration` block as defined below.
     */
    githubConfiguration?: pulumi.Input<inputs.datafactory.FactoryGithubConfiguration>;
    /**
     * A list of `globalParameter` blocks as defined above.
     */
    globalParameters?: pulumi.Input<pulumi.Input<inputs.datafactory.FactoryGlobalParameter>[]>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.datafactory.FactoryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Is Managed Virtual Network enabled?
     */
    managedVirtualNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Data Factory. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    name?: pulumi.Input<string>;
    /**
     * Is the Data Factory visible to the public network? Defaults to `true`.
     */
    publicNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Data Factory.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `vstsConfiguration` block as defined below.
     */
    vstsConfiguration?: pulumi.Input<inputs.datafactory.FactoryVstsConfiguration>;
}
