// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Logic App Workflow.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleWorkflow = new azure.logicapps.Workflow("exampleWorkflow", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * ```
 *
 * ## Import
 *
 * Logic App Workflows can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:logicapps/workflow:Workflow workflow1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
 * ```
 */
export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowState, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:logicapps/workflow:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * A `accessControl` block as defined below.
     */
    public readonly accessControl!: pulumi.Output<outputs.logicapps.WorkflowAccessControl | undefined>;
    /**
     * The Access Endpoint for the Logic App Workflow.
     */
    public /*out*/ readonly accessEndpoint!: pulumi.Output<string>;
    /**
     * The list of access endpoint ip addresses of connector.
     */
    public /*out*/ readonly connectorEndpointIpAddresses!: pulumi.Output<string[]>;
    /**
     * The list of outgoing ip addresses of connector.
     */
    public /*out*/ readonly connectorOutboundIpAddresses!: pulumi.Output<string[]>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
     */
    public readonly integrationServiceEnvironmentId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the integration account linked by this Logic App Workflow.
     */
    public readonly logicAppIntegrationAccountId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of Key-Value pairs.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The list of access endpoint ip addresses of workflow.
     */
    public /*out*/ readonly workflowEndpointIpAddresses!: pulumi.Output<string[]>;
    /**
     * The list of outgoing ip addresses of workflow.
     */
    public /*out*/ readonly workflowOutboundIpAddresses!: pulumi.Output<string[]>;
    /**
     * Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a json encoded string of the parameter definition (see: https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
     */
    public readonly workflowParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
     */
    public readonly workflowSchema!: pulumi.Output<string | undefined>;
    /**
     * Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
     */
    public readonly workflowVersion!: pulumi.Output<string | undefined>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowArgs | WorkflowState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowState | undefined;
            inputs["accessControl"] = state ? state.accessControl : undefined;
            inputs["accessEndpoint"] = state ? state.accessEndpoint : undefined;
            inputs["connectorEndpointIpAddresses"] = state ? state.connectorEndpointIpAddresses : undefined;
            inputs["connectorOutboundIpAddresses"] = state ? state.connectorOutboundIpAddresses : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["integrationServiceEnvironmentId"] = state ? state.integrationServiceEnvironmentId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["logicAppIntegrationAccountId"] = state ? state.logicAppIntegrationAccountId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workflowEndpointIpAddresses"] = state ? state.workflowEndpointIpAddresses : undefined;
            inputs["workflowOutboundIpAddresses"] = state ? state.workflowOutboundIpAddresses : undefined;
            inputs["workflowParameters"] = state ? state.workflowParameters : undefined;
            inputs["workflowSchema"] = state ? state.workflowSchema : undefined;
            inputs["workflowVersion"] = state ? state.workflowVersion : undefined;
        } else {
            const args = argsOrState as WorkflowArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessControl"] = args ? args.accessControl : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["integrationServiceEnvironmentId"] = args ? args.integrationServiceEnvironmentId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logicAppIntegrationAccountId"] = args ? args.logicAppIntegrationAccountId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workflowParameters"] = args ? args.workflowParameters : undefined;
            inputs["workflowSchema"] = args ? args.workflowSchema : undefined;
            inputs["workflowVersion"] = args ? args.workflowVersion : undefined;
            inputs["accessEndpoint"] = undefined /*out*/;
            inputs["connectorEndpointIpAddresses"] = undefined /*out*/;
            inputs["connectorOutboundIpAddresses"] = undefined /*out*/;
            inputs["workflowEndpointIpAddresses"] = undefined /*out*/;
            inputs["workflowOutboundIpAddresses"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Workflow.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workflow resources.
 */
export interface WorkflowState {
    /**
     * A `accessControl` block as defined below.
     */
    accessControl?: pulumi.Input<inputs.logicapps.WorkflowAccessControl>;
    /**
     * The Access Endpoint for the Logic App Workflow.
     */
    accessEndpoint?: pulumi.Input<string>;
    /**
     * The list of access endpoint ip addresses of connector.
     */
    connectorEndpointIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of outgoing ip addresses of connector.
     */
    connectorOutboundIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
     */
    integrationServiceEnvironmentId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the integration account linked by this Logic App Workflow.
     */
    logicAppIntegrationAccountId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of Key-Value pairs.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of access endpoint ip addresses of workflow.
     */
    workflowEndpointIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of outgoing ip addresses of workflow.
     */
    workflowOutboundIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a json encoded string of the parameter definition (see: https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
     */
    workflowParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
     */
    workflowSchema?: pulumi.Input<string>;
    /**
     * Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
     */
    workflowVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * A `accessControl` block as defined below.
     */
    accessControl?: pulumi.Input<inputs.logicapps.WorkflowAccessControl>;
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
     */
    integrationServiceEnvironmentId?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the integration account linked by this Logic App Workflow.
     */
    logicAppIntegrationAccountId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of Key-Value pairs.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a map of Key-Value pairs of the Parameter Definitions to use for this Logic App Workflow. The key is the parameter name, and the value is a json encoded string of the parameter definition (see: https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-workflow-definition-language#parameters).
     */
    workflowParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
     */
    workflowSchema?: pulumi.Input<string>;
    /**
     * Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
     */
    workflowVersion?: pulumi.Input<string>;
}
