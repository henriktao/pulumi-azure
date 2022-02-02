// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Template Deployment at a Management Group Scope.
 *
 * > **Note:** Deleting a Deployment at the Management Group Scope will not delete any resources created by the deployment.
 *
 * > **Note:** Deployments to a Management Group are always Incrementally applied. Existing resources that are not part of the template will not be removed.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleGroup = azure.management.getGroup({
 *     name: "00000000-0000-0000-0000-000000000000",
 * });
 * const exampleGroupTemplateDeployment = new azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment", {
 *     location: "West Europe",
 *     managementGroupId: exampleGroup.then(exampleGroup => exampleGroup.id),
 *     templateContent: `{
 *   "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
 *   "contentVersion": "1.0.0.0",
 *   "parameters": {
 *     "policyAssignmentName": {
 *       "type": "string",
 *       "defaultValue": "[guid(parameters('policyDefinitionID'), resourceGroup().name)]",
 *       "metadata": {
 *         "description": "Specifies the name of the policy assignment, can be used defined or an idempotent name as the defaultValue provides."
 *       }
 *     },
 *     "policyDefinitionID": {
 *       "type": "string",
 *       "metadata": {
 *         "description": "Specifies the ID of the policy definition or policy set definition being assigned."
 *       }
 *     }
 *   },
 *   "resources": [
 *     {
 *       "type": "Microsoft.Authorization/policyAssignments",
 *       "name": "[parameters('policyAssignmentName')]",
 *       "apiVersion": "2019-09-01",
 *       "properties": {
 *         "scope": "[subscriptionResourceId('Microsoft.Resources/resourceGroups', resourceGroup().name)]",
 *         "policyDefinitionId": "[parameters('policyDefinitionID')]"
 *       }
 *     }
 *   ]
 * }
 * `,
 *     parametersContent: `{
 *   "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#",
 *   "contentVersion": "1.0.0.0",
 *   "parameters": {
 *     "policyDefinitionID": {
 *       "value": "/providers/Microsoft.Authorization/policyDefinitions/0a914e76-4921-4c19-b460-a2d36003525a"
 *     }
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * from "fs";
 *
 * const exampleGroup = azure.management.getGroup({
 *     name: "00000000-0000-0000-0000-000000000000",
 * });
 * const exampleGroupTemplateDeployment = new azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment", {
 *     location: "West Europe",
 *     managementGroupId: exampleGroup.then(exampleGroup => exampleGroup.id),
 *     templateContent: fs.readFileSync("templates/example-deploy-template.json"),
 *     parametersContent: fs.readFileSync("templates/example-deploy-params.json"),
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleGroup = azure.management.getGroup({
 *     name: "00000000-0000-0000-0000-000000000000",
 * });
 * const exampleTemplateSpecVersion = azure.core.getTemplateSpecVersion({
 *     name: "exampleTemplateForManagementGroup",
 *     resourceGroupName: "exampleResourceGroup",
 *     version: "v1.0.9",
 * });
 * const exampleGroupTemplateDeployment = new azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment", {
 *     location: "West Europe",
 *     managementGroupId: exampleGroup.then(exampleGroup => exampleGroup.id),
 *     templateSpecVersionId: exampleTemplateSpecVersion.then(exampleTemplateSpecVersion => exampleTemplateSpecVersion.id),
 * });
 * ```
 *
 * ## Import
 *
 * Management Group Template Deployments can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:management/groupTemplateDeployment:GroupTemplateDeployment example /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
 * ```
 */
export class GroupTemplateDeployment extends pulumi.CustomResource {
    /**
     * Get an existing GroupTemplateDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupTemplateDeploymentState, opts?: pulumi.CustomResourceOptions): GroupTemplateDeployment {
        return new GroupTemplateDeployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:management/groupTemplateDeployment:GroupTemplateDeployment';

    /**
     * Returns true if the given object is an instance of GroupTemplateDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupTemplateDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupTemplateDeployment.__pulumiType;
    }

    /**
     * The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    public readonly debugLevel!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the Template should exist. Changing this forces a new Template to be created.
     */
    public readonly location!: pulumi.Output<string>;
    public readonly managementGroupId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The JSON Content of the Outputs of the ARM Template Deployment.
     */
    public /*out*/ readonly outputContent!: pulumi.Output<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    public readonly parametersContent!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Template.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
     */
    public readonly templateContent!: pulumi.Output<string>;
    /**
     * The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
     */
    public readonly templateSpecVersionId!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupTemplateDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupTemplateDeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupTemplateDeploymentArgs | GroupTemplateDeploymentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupTemplateDeploymentState | undefined;
            resourceInputs["debugLevel"] = state ? state.debugLevel : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managementGroupId"] = state ? state.managementGroupId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputContent"] = state ? state.outputContent : undefined;
            resourceInputs["parametersContent"] = state ? state.parametersContent : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["templateContent"] = state ? state.templateContent : undefined;
            resourceInputs["templateSpecVersionId"] = state ? state.templateSpecVersionId : undefined;
        } else {
            const args = argsOrState as GroupTemplateDeploymentArgs | undefined;
            if ((!args || args.managementGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managementGroupId'");
            }
            resourceInputs["debugLevel"] = args ? args.debugLevel : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managementGroupId"] = args ? args.managementGroupId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parametersContent"] = args ? args.parametersContent : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateContent"] = args ? args.templateContent : undefined;
            resourceInputs["templateSpecVersionId"] = args ? args.templateSpecVersionId : undefined;
            resourceInputs["outputContent"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupTemplateDeployment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupTemplateDeployment resources.
 */
export interface GroupTemplateDeploymentState {
    /**
     * The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    debugLevel?: pulumi.Input<string>;
    /**
     * The Azure Region where the Template should exist. Changing this forces a new Template to be created.
     */
    location?: pulumi.Input<string>;
    managementGroupId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The JSON Content of the Outputs of the ARM Template Deployment.
     */
    outputContent?: pulumi.Input<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    parametersContent?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Template.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
     */
    templateContent?: pulumi.Input<string>;
    /**
     * The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
     */
    templateSpecVersionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupTemplateDeployment resource.
 */
export interface GroupTemplateDeploymentArgs {
    /**
     * The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    debugLevel?: pulumi.Input<string>;
    /**
     * The Azure Region where the Template should exist. Changing this forces a new Template to be created.
     */
    location?: pulumi.Input<string>;
    managementGroupId: pulumi.Input<string>;
    /**
     * The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    parametersContent?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Template.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `templateSpecVersionId`.
     */
    templateContent?: pulumi.Input<string>;
    /**
     * The ID of the Template Spec Version to deploy. Cannot be specified with `templateContent`.
     */
    templateSpecVersionId?: pulumi.Input<string>;
}
