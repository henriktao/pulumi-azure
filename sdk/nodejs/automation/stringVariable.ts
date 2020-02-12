// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a string variable in Azure Automation
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_variable_string.html.markdown.
 */
export class StringVariable extends pulumi.CustomResource {
    /**
     * Get an existing StringVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StringVariableState, opts?: pulumi.CustomResourceOptions): StringVariable {
        return new StringVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/stringVariable:StringVariable';

    /**
     * Returns true if the given object is an instance of StringVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StringVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StringVariable.__pulumiType;
    }

    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The description of the Automation Variable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The value of the Automation Variable as a `string`.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a StringVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StringVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StringVariableArgs | StringVariableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StringVariableState | undefined;
            inputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as StringVariableArgs | undefined;
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StringVariable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StringVariable resources.
 */
export interface StringVariableState {
    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName?: pulumi.Input<string>;
    /**
     * The description of the Automation Variable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The value of the Automation Variable as a `string`.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StringVariable resource.
 */
export interface StringVariableArgs {
    /**
     * The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The description of the Automation Variable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies if the Automation Variable is encrypted. Defaults to `false`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The name of the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The value of the Automation Variable as a `string`.
     */
    readonly value?: pulumi.Input<string>;
}
