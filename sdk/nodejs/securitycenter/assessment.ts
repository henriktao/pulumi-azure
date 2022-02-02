// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages the Security Center Assessment for Azure Security Center.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * from "fs";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const internal = new azure.network.Subnet("internal", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 * });
 * const exampleLinuxVirtualMachineScaleSet = new azure.compute.LinuxVirtualMachineScaleSet("exampleLinuxVirtualMachineScaleSet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     sku: "Standard_F2",
 *     instances: 1,
 *     adminUsername: "adminuser",
 *     adminSshKeys: [{
 *         username: "adminuser",
 *         publicKey: fs.readFileSync("~/.ssh/id_rsa.pub"),
 *     }],
 *     sourceImageReference: {
 *         publisher: "Canonical",
 *         offer: "UbuntuServer",
 *         sku: "16.04-LTS",
 *         version: "latest",
 *     },
 *     osDisk: {
 *         storageAccountType: "Standard_LRS",
 *         caching: "ReadWrite",
 *     },
 *     networkInterfaces: [{
 *         name: "example",
 *         primary: true,
 *         ipConfigurations: [{
 *             name: "internal",
 *             primary: true,
 *             subnetId: internal.id,
 *         }],
 *     }],
 * });
 * const exampleAssessmentPolicy = new azure.securitycenter.AssessmentPolicy("exampleAssessmentPolicy", {
 *     displayName: "Test Display Name",
 *     severity: "Medium",
 *     description: "Test Description",
 * });
 * const exampleAssessment = new azure.securitycenter.Assessment("exampleAssessment", {
 *     assessmentPolicyId: exampleAssessmentPolicy.id,
 *     targetResourceId: exampleLinuxVirtualMachineScaleSet.id,
 *     status: {
 *         code: "Healthy",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Security Assessment can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:securitycenter/assessment:Assessment example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/00000000-0000-0000-0000-000000000000
 * ```
 */
export class Assessment extends pulumi.CustomResource {
    /**
     * Get an existing Assessment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssessmentState, opts?: pulumi.CustomResourceOptions): Assessment {
        return new Assessment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:securitycenter/assessment:Assessment';

    /**
     * Returns true if the given object is an instance of Assessment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assessment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assessment.__pulumiType;
    }

    /**
     * A map of additional data to associate with the assessment.
     */
    public readonly additionalData!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
     */
    public readonly assessmentPolicyId!: pulumi.Output<string>;
    /**
     * A `status` block as defined below.
     */
    public readonly status!: pulumi.Output<outputs.securitycenter.AssessmentStatus>;
    /**
     * The ID of the target resource. Changing this forces a new security Assessment to be created.
     */
    public readonly targetResourceId!: pulumi.Output<string>;

    /**
     * Create a Assessment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssessmentArgs | AssessmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssessmentState | undefined;
            resourceInputs["additionalData"] = state ? state.additionalData : undefined;
            resourceInputs["assessmentPolicyId"] = state ? state.assessmentPolicyId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["targetResourceId"] = state ? state.targetResourceId : undefined;
        } else {
            const args = argsOrState as AssessmentArgs | undefined;
            if ((!args || args.assessmentPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assessmentPolicyId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.targetResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            resourceInputs["additionalData"] = args ? args.additionalData : undefined;
            resourceInputs["assessmentPolicyId"] = args ? args.assessmentPolicyId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["targetResourceId"] = args ? args.targetResourceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Assessment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Assessment resources.
 */
export interface AssessmentState {
    /**
     * A map of additional data to associate with the assessment.
     */
    additionalData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
     */
    assessmentPolicyId?: pulumi.Input<string>;
    /**
     * A `status` block as defined below.
     */
    status?: pulumi.Input<inputs.securitycenter.AssessmentStatus>;
    /**
     * The ID of the target resource. Changing this forces a new security Assessment to be created.
     */
    targetResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Assessment resource.
 */
export interface AssessmentArgs {
    /**
     * A map of additional data to associate with the assessment.
     */
    additionalData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
     */
    assessmentPolicyId: pulumi.Input<string>;
    /**
     * A `status` block as defined below.
     */
    status: pulumi.Input<inputs.securitycenter.AssessmentStatus>;
    /**
     * The ID of the target resource. Changing this forces a new security Assessment to be created.
     */
    targetResourceId: pulumi.Input<string>;
}
