// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Machine Learning Compute Cluster.
 * **NOTE:** At this point in time the resource cannot be updated (not supported by the backend Azure Go SDK). Therefore it can only be created and deleted, not updated. At the moment, there is also no possibility to specify ssh User Account Credentials to ssh into the compute cluster.
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
 * const test = new azure.machinelearning.ComputeCluster("test", {
 *     location: "West Europe",
 *     vmPriority: "LowPriority",
 *     vmSize: "Standard_DS2_v2",
 *     machineLearningWorkspaceId: exampleWorkspace.id,
 *     subnetResourceId: exampleSubnet.id,
 *     scaleSettings: {
 *         minNodeCount: 0,
 *         maxNodeCount: 1,
 *         scaleDownNodesAfterIdleDuration: "PT30S",
 *     },
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Machine Learning Compute Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:machinelearning/computeCluster:ComputeCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
 * ```
 */
export class ComputeCluster extends pulumi.CustomResource {
    /**
     * Get an existing ComputeCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterState, opts?: pulumi.CustomResourceOptions): ComputeCluster {
        return new ComputeCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:machinelearning/computeCluster:ComputeCluster';

    /**
     * Returns true if the given object is an instance of ComputeCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeCluster.__pulumiType;
    }

    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly identity!: pulumi.Output<outputs.machinelearning.ComputeClusterIdentity | undefined>;
    /**
     * The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly machineLearningWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly scaleSettings!: pulumi.Output<outputs.machinelearning.ComputeClusterScaleSettings>;
    /**
     * Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly ssh!: pulumi.Output<outputs.machinelearning.ComputeClusterSsh | undefined>;
    /**
     * A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly sshPublicAccessEnabled!: pulumi.Output<boolean>;
    /**
     * The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly subnetResourceId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly vmPriority!: pulumi.Output<string>;
    /**
     * The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    public readonly vmSize!: pulumi.Output<string>;

    /**
     * Create a ComputeCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterArgs | ComputeClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeClusterState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["machineLearningWorkspaceId"] = state ? state.machineLearningWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scaleSettings"] = state ? state.scaleSettings : undefined;
            inputs["ssh"] = state ? state.ssh : undefined;
            inputs["sshPublicAccessEnabled"] = state ? state.sshPublicAccessEnabled : undefined;
            inputs["subnetResourceId"] = state ? state.subnetResourceId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vmPriority"] = state ? state.vmPriority : undefined;
            inputs["vmSize"] = state ? state.vmSize : undefined;
        } else {
            const args = argsOrState as ComputeClusterArgs | undefined;
            if ((!args || args.machineLearningWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineLearningWorkspaceId'");
            }
            if ((!args || args.scaleSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scaleSettings'");
            }
            if ((!args || args.vmPriority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmPriority'");
            }
            if ((!args || args.vmSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmSize'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["machineLearningWorkspaceId"] = args ? args.machineLearningWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scaleSettings"] = args ? args.scaleSettings : undefined;
            inputs["ssh"] = args ? args.ssh : undefined;
            inputs["sshPublicAccessEnabled"] = args ? args.sshPublicAccessEnabled : undefined;
            inputs["subnetResourceId"] = args ? args.subnetResourceId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vmPriority"] = args ? args.vmPriority : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ComputeCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeCluster resources.
 */
export interface ComputeClusterState {
    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    identity?: pulumi.Input<inputs.machinelearning.ComputeClusterIdentity>;
    /**
     * The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    machineLearningWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    scaleSettings?: pulumi.Input<inputs.machinelearning.ComputeClusterScaleSettings>;
    /**
     * Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    ssh?: pulumi.Input<inputs.machinelearning.ComputeClusterSsh>;
    /**
     * A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    sshPublicAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    subnetResourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    vmPriority?: pulumi.Input<string>;
    /**
     * The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    vmSize?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeCluster resource.
 */
export interface ComputeClusterArgs {
    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * A `identity` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    identity?: pulumi.Input<inputs.machinelearning.ComputeClusterIdentity>;
    /**
     * The Azure Region where the Machine Learning Compute Cluster should exist. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    machineLearningWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `scaleSettings` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    scaleSettings: pulumi.Input<inputs.machinelearning.ComputeClusterScaleSettings>;
    /**
     * Credentials for an administrator user account that will be created on each compute node. A `ssh` block as defined below. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    ssh?: pulumi.Input<inputs.machinelearning.ComputeClusterSsh>;
    /**
     * A boolean value indicating whether enable the public SSH port. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    sshPublicAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Subnet that the Compute Cluster should reside in. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    subnetResourceId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The priority of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    vmPriority: pulumi.Input<string>;
    /**
     * The size of the VM. Changing this forces a new Machine Learning Compute Cluster to be created.
     */
    vmSize: pulumi.Input<string>;
}
