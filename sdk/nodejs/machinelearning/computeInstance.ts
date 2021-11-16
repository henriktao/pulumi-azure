// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Machine Learning Compute Instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {
 *     location: "west europe",
 *     tags: {
 *         stage: "example",
 *     },
 * });
 * const exampleInsights = new azure.appinsights.Insights("exampleInsights", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationType: "web",
 * });
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: current.then(current => current.tenantId),
 *     skuName: "standard",
 *     purgeProtectionEnabled: true,
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleWorkspace = new azure.machinelearning.Workspace("exampleWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationInsightsId: exampleInsights.id,
 *     keyVaultId: exampleKeyVault.id,
 *     storageAccountId: exampleAccount.id,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.1.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.1.0.0/24"],
 * });
 * const config = new pulumi.Config();
 * const sshKey = config.get("sshKey") || "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqaZoyiz1qbdOQ8xEf6uEu1cCwYowo5FHtsBhqLoDnnp7KUTEBN+L2NxRIfQ781rxV6Iq5jSav6b2Q8z5KiseOlvKA/RF2wqU0UPYqQviQhLmW6THTpmrv/YkUCuzxDpsH7DUDhZcwySLKVVe0Qm3+5N2Ta6UYH3lsDf9R9wTP2K/+vAnflKebuypNlmocIvakFWoZda18FOmsOoIVXQ8HWFNCuw9ZCunMSN62QGamCe3dL5cXlkgHYv7ekJE15IA9aOJcM7e90oeTqo+7HTcWfdu0qQqPWY5ujyMw/llas8tsXY85LFqRnr3gJ02bAscjc477+X+j/gkpFoN1QEmt terraform@demo.tld";
 * const exampleComputeInstance = new azure.machinelearning.ComputeInstance("exampleComputeInstance", {
 *     location: exampleResourceGroup.location,
 *     machineLearningWorkspaceId: exampleWorkspace.id,
 *     virtualMachineSize: "STANDARD_DS2_V2",
 *     authorizationType: "personal",
 *     ssh: {
 *         publicKey: sshKey,
 *     },
 *     subnetResourceId: exampleSubnet.id,
 *     description: "foo",
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Machine Learning Compute Instances can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:machinelearning/computeInstance:ComputeInstance example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
 * ```
 */
export class ComputeInstance extends pulumi.CustomResource {
    /**
     * Get an existing ComputeInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeInstanceState, opts?: pulumi.CustomResourceOptions): ComputeInstance {
        return new ComputeInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:machinelearning/computeInstance:ComputeInstance';

    /**
     * Returns true if the given object is an instance of ComputeInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeInstance.__pulumiType;
    }

    /**
     * A `assignToUser` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly assignToUser!: pulumi.Output<outputs.machinelearning.ComputeInstanceAssignToUser | undefined>;
    /**
     * The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly authorizationType!: pulumi.Output<string | undefined>;
    /**
     * The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly identity!: pulumi.Output<outputs.machinelearning.ComputeInstanceIdentity | undefined>;
    /**
     * Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly localAuthEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly machineLearningWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly ssh!: pulumi.Output<outputs.machinelearning.ComputeInstanceSsh | undefined>;
    /**
     * Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly subnetResourceId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    public readonly virtualMachineSize!: pulumi.Output<string>;

    /**
     * Create a ComputeInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeInstanceArgs | ComputeInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeInstanceState | undefined;
            inputs["assignToUser"] = state ? state.assignToUser : undefined;
            inputs["authorizationType"] = state ? state.authorizationType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["localAuthEnabled"] = state ? state.localAuthEnabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["machineLearningWorkspaceId"] = state ? state.machineLearningWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ssh"] = state ? state.ssh : undefined;
            inputs["subnetResourceId"] = state ? state.subnetResourceId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualMachineSize"] = state ? state.virtualMachineSize : undefined;
        } else {
            const args = argsOrState as ComputeInstanceArgs | undefined;
            if ((!args || args.machineLearningWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineLearningWorkspaceId'");
            }
            if ((!args || args.virtualMachineSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineSize'");
            }
            inputs["assignToUser"] = args ? args.assignToUser : undefined;
            inputs["authorizationType"] = args ? args.authorizationType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["localAuthEnabled"] = args ? args.localAuthEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["machineLearningWorkspaceId"] = args ? args.machineLearningWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ssh"] = args ? args.ssh : undefined;
            inputs["subnetResourceId"] = args ? args.subnetResourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachineSize"] = args ? args.virtualMachineSize : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ComputeInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeInstance resources.
 */
export interface ComputeInstanceState {
    /**
     * A `assignToUser` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    assignToUser?: pulumi.Input<inputs.machinelearning.ComputeInstanceAssignToUser>;
    /**
     * The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    authorizationType?: pulumi.Input<string>;
    /**
     * The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    identity?: pulumi.Input<inputs.machinelearning.ComputeInstanceIdentity>;
    /**
     * Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    localAuthEnabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    machineLearningWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    ssh?: pulumi.Input<inputs.machinelearning.ComputeInstanceSsh>;
    /**
     * Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    subnetResourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    virtualMachineSize?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeInstance resource.
 */
export interface ComputeInstanceArgs {
    /**
     * A `assignToUser` block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    assignToUser?: pulumi.Input<inputs.machinelearning.ComputeInstanceAssignToUser>;
    /**
     * The Compute Instance Authorization type. Possible values include: `personal`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    authorizationType?: pulumi.Input<string>;
    /**
     * The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    identity?: pulumi.Input<inputs.machinelearning.ComputeInstanceIdentity>;
    /**
     * Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    localAuthEnabled?: pulumi.Input<boolean>;
    /**
     * The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    machineLearningWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `ssh` block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    ssh?: pulumi.Input<inputs.machinelearning.ComputeInstanceSsh>;
    /**
     * Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    subnetResourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
     */
    virtualMachineSize: pulumi.Input<string>;
}
