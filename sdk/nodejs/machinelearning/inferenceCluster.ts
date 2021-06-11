// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Machine Learning Inference Cluster.
 *
 * > **NOTE:** The Machine Learning Inference Cluster resource is used to attach an existing AKS cluster to the Machine Learning Workspace, it doesn't create the AKS cluster itself. Therefore it can only be created and deleted, not updated. Any change to the configuration will recreate the resource.
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
 *     addressPrefix: "10.1.0.0/24",
 * });
 * const exampleKubernetesCluster = new azure.containerservice.KubernetesCluster("exampleKubernetesCluster", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     defaultNodePool: {
 *         name: "default",
 *         nodeCount: 3,
 *         vmSize: "Standard_D3_v2",
 *         vnetSubnetId: exampleSubnet.id,
 *     },
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleInferenceCluster = new azure.machinelearning.InferenceCluster("exampleInferenceCluster", {
 *     location: exampleResourceGroup.location,
 *     clusterPurpose: "FastProd",
 *     kubernetesClusterId: exampleKubernetesCluster.id,
 *     description: "This is an example cluster used with Terraform",
 *     machineLearningWorkspaceId: exampleWorkspace.id,
 *     tags: {
 *         stage: "example",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Machine Learning Inference Clusters can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:machinelearning/inferenceCluster:InferenceCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
 * ```
 */
export class InferenceCluster extends pulumi.CustomResource {
    /**
     * Get an existing InferenceCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InferenceClusterState, opts?: pulumi.CustomResourceOptions): InferenceCluster {
        return new InferenceCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:machinelearning/inferenceCluster:InferenceCluster';

    /**
     * Returns true if the given object is an instance of InferenceCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InferenceCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InferenceCluster.__pulumiType;
    }

    /**
     * The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly clusterPurpose!: pulumi.Output<string | undefined>;
    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly kubernetesClusterId!: pulumi.Output<string>;
    /**
     * The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly machineLearningWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly ssl!: pulumi.Output<outputs.machinelearning.InferenceClusterSsl | undefined>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a InferenceCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InferenceClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InferenceClusterArgs | InferenceClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InferenceClusterState | undefined;
            inputs["clusterPurpose"] = state ? state.clusterPurpose : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["kubernetesClusterId"] = state ? state.kubernetesClusterId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["machineLearningWorkspaceId"] = state ? state.machineLearningWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ssl"] = state ? state.ssl : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InferenceClusterArgs | undefined;
            if ((!args || args.kubernetesClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesClusterId'");
            }
            if ((!args || args.machineLearningWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'machineLearningWorkspaceId'");
            }
            inputs["clusterPurpose"] = args ? args.clusterPurpose : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kubernetesClusterId"] = args ? args.kubernetesClusterId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["machineLearningWorkspaceId"] = args ? args.machineLearningWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ssl"] = args ? args.ssl : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InferenceCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InferenceCluster resources.
 */
export interface InferenceClusterState {
    /**
     * The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    clusterPurpose?: pulumi.Input<string>;
    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    kubernetesClusterId?: pulumi.Input<string>;
    /**
     * The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    machineLearningWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    ssl?: pulumi.Input<inputs.machinelearning.InferenceClusterSsl>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a InferenceCluster resource.
 */
export interface InferenceClusterArgs {
    /**
     * The purpose of the Inference Cluster. Options are `DevTest`, `DenseProd` and `FastProd`. If used for Development or Testing, use `DevTest` here. Default purpose is `FastProd`, which is recommended for production workloads. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    clusterPurpose?: pulumi.Input<string>;
    /**
     * The description of the Machine Learning compute. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Kubernetes Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    kubernetesClusterId: pulumi.Input<string>;
    /**
     * The Azure Region where the Machine Learning Inference Cluster should exist. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    machineLearningWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `ssl` block as defined below. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    ssl?: pulumi.Input<inputs.machinelearning.InferenceClusterSsl>;
    /**
     * A mapping of tags which should be assigned to the Machine Learning Inference Cluster. Changing this forces a new Machine Learning Inference Cluster to be created.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
