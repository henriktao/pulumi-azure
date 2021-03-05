// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Kubernetes Cluster Node Pools can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool pool1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool1
 * ```
 */
export class KubernetesClusterNodePool extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesClusterNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesClusterNodePoolState, opts?: pulumi.CustomResourceOptions): KubernetesClusterNodePool {
        return new KubernetesClusterNodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/kubernetesClusterNodePool:KubernetesClusterNodePool';

    /**
     * Returns true if the given object is an instance of KubernetesClusterNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesClusterNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesClusterNodePool.__pulumiType;
    }

    /**
     * A list of Availability Zones where the Nodes in this Node Pool should be created in. Changing this forces a new resource to be created.
     */
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
     */
    public readonly enableAutoScaling!: pulumi.Output<boolean | undefined>;
    /**
     * Should the nodes in this Node Pool have host encryption enabled? Defaults to `false`.
     */
    public readonly enableHostEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * Should each node have a Public IP Address? Defaults to `false`.
     */
    public readonly enableNodePublicIp!: pulumi.Output<boolean | undefined>;
    /**
     * The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
     */
    public readonly kubernetesClusterId!: pulumi.Output<string>;
    /**
     * The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be greater than or equal to `minCount`.
     */
    public readonly maxCount!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
     */
    public readonly maxPods!: pulumi.Output<number>;
    /**
     * The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be less than or equal to `maxCount`.
     */
    public readonly minCount!: pulumi.Output<number | undefined>;
    /**
     * Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be a value in the range `minCount` - `maxCount`.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
     */
    public readonly nodeLabels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
     */
    public readonly nodeTaints!: pulumi.Output<string[] | undefined>;
    /**
     * Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
     */
    public readonly orchestratorVersion!: pulumi.Output<string>;
    /**
     * The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
     */
    public readonly osDiskSizeGb!: pulumi.Output<number>;
    /**
     * The type of disk which should be used for the Operating System. Possible values are `Ephemeral` and `Managed`. Defaults to `Managed`. Changing this forces a new resource to be created.
     */
    public readonly osDiskType!: pulumi.Output<string | undefined>;
    /**
     * The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Proximity Placement Group where the Virtual Machine Scale Set that powers this Node Pool will be placed. Changing this forces a new resource to be created.
     */
    public readonly proximityPlacementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
     */
    public readonly spotMaxPrice!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `upgradeSettings` block as documented below.
     */
    public readonly upgradeSettings!: pulumi.Output<outputs.containerservice.KubernetesClusterNodePoolUpgradeSettings | undefined>;
    /**
     * The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
     */
    public readonly vmSize!: pulumi.Output<string>;
    /**
     * The ID of the Subnet where this Node Pool should exist.
     */
    public readonly vnetSubnetId!: pulumi.Output<string | undefined>;

    /**
     * Create a KubernetesClusterNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesClusterNodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesClusterNodePoolArgs | KubernetesClusterNodePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesClusterNodePoolState | undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["enableAutoScaling"] = state ? state.enableAutoScaling : undefined;
            inputs["enableHostEncryption"] = state ? state.enableHostEncryption : undefined;
            inputs["enableNodePublicIp"] = state ? state.enableNodePublicIp : undefined;
            inputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            inputs["kubernetesClusterId"] = state ? state.kubernetesClusterId : undefined;
            inputs["maxCount"] = state ? state.maxCount : undefined;
            inputs["maxPods"] = state ? state.maxPods : undefined;
            inputs["minCount"] = state ? state.minCount : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["nodeLabels"] = state ? state.nodeLabels : undefined;
            inputs["nodeTaints"] = state ? state.nodeTaints : undefined;
            inputs["orchestratorVersion"] = state ? state.orchestratorVersion : undefined;
            inputs["osDiskSizeGb"] = state ? state.osDiskSizeGb : undefined;
            inputs["osDiskType"] = state ? state.osDiskType : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["proximityPlacementGroupId"] = state ? state.proximityPlacementGroupId : undefined;
            inputs["spotMaxPrice"] = state ? state.spotMaxPrice : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["upgradeSettings"] = state ? state.upgradeSettings : undefined;
            inputs["vmSize"] = state ? state.vmSize : undefined;
            inputs["vnetSubnetId"] = state ? state.vnetSubnetId : undefined;
        } else {
            const args = argsOrState as KubernetesClusterNodePoolArgs | undefined;
            if ((!args || args.kubernetesClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesClusterId'");
            }
            if ((!args || args.vmSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmSize'");
            }
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["enableAutoScaling"] = args ? args.enableAutoScaling : undefined;
            inputs["enableHostEncryption"] = args ? args.enableHostEncryption : undefined;
            inputs["enableNodePublicIp"] = args ? args.enableNodePublicIp : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["kubernetesClusterId"] = args ? args.kubernetesClusterId : undefined;
            inputs["maxCount"] = args ? args.maxCount : undefined;
            inputs["maxPods"] = args ? args.maxPods : undefined;
            inputs["minCount"] = args ? args.minCount : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["nodeLabels"] = args ? args.nodeLabels : undefined;
            inputs["nodeTaints"] = args ? args.nodeTaints : undefined;
            inputs["orchestratorVersion"] = args ? args.orchestratorVersion : undefined;
            inputs["osDiskSizeGb"] = args ? args.osDiskSizeGb : undefined;
            inputs["osDiskType"] = args ? args.osDiskType : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["proximityPlacementGroupId"] = args ? args.proximityPlacementGroupId : undefined;
            inputs["spotMaxPrice"] = args ? args.spotMaxPrice : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradeSettings"] = args ? args.upgradeSettings : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["vnetSubnetId"] = args ? args.vnetSubnetId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KubernetesClusterNodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesClusterNodePool resources.
 */
export interface KubernetesClusterNodePoolState {
    /**
     * A list of Availability Zones where the Nodes in this Node Pool should be created in. Changing this forces a new resource to be created.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
     */
    readonly enableAutoScaling?: pulumi.Input<boolean>;
    /**
     * Should the nodes in this Node Pool have host encryption enabled? Defaults to `false`.
     */
    readonly enableHostEncryption?: pulumi.Input<boolean>;
    /**
     * Should each node have a Public IP Address? Defaults to `false`.
     */
    readonly enableNodePublicIp?: pulumi.Input<boolean>;
    /**
     * The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
     */
    readonly kubernetesClusterId?: pulumi.Input<string>;
    /**
     * The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be greater than or equal to `minCount`.
     */
    readonly maxCount?: pulumi.Input<number>;
    /**
     * The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
     */
    readonly maxPods?: pulumi.Input<number>;
    /**
     * The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be less than or equal to `maxCount`.
     */
    readonly minCount?: pulumi.Input<number>;
    /**
     * Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be a value in the range `minCount` - `maxCount`.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
     */
    readonly nodeLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
     */
    readonly nodeTaints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
     */
    readonly orchestratorVersion?: pulumi.Input<string>;
    /**
     * The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
     */
    readonly osDiskSizeGb?: pulumi.Input<number>;
    /**
     * The type of disk which should be used for the Operating System. Possible values are `Ephemeral` and `Managed`. Defaults to `Managed`. Changing this forces a new resource to be created.
     */
    readonly osDiskType?: pulumi.Input<string>;
    /**
     * The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group where the Virtual Machine Scale Set that powers this Node Pool will be placed. Changing this forces a new resource to be created.
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
     */
    readonly spotMaxPrice?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `upgradeSettings` block as documented below.
     */
    readonly upgradeSettings?: pulumi.Input<inputs.containerservice.KubernetesClusterNodePoolUpgradeSettings>;
    /**
     * The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
     */
    readonly vmSize?: pulumi.Input<string>;
    /**
     * The ID of the Subnet where this Node Pool should exist.
     */
    readonly vnetSubnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KubernetesClusterNodePool resource.
 */
export interface KubernetesClusterNodePoolArgs {
    /**
     * A list of Availability Zones where the Nodes in this Node Pool should be created in. Changing this forces a new resource to be created.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable [auto-scaler](https://docs.microsoft.com/en-us/azure/aks/cluster-autoscaler). Defaults to `false`.
     */
    readonly enableAutoScaling?: pulumi.Input<boolean>;
    /**
     * Should the nodes in this Node Pool have host encryption enabled? Defaults to `false`.
     */
    readonly enableHostEncryption?: pulumi.Input<boolean>;
    /**
     * Should each node have a Public IP Address? Defaults to `false`.
     */
    readonly enableNodePublicIp?: pulumi.Input<boolean>;
    /**
     * The Eviction Policy which should be used for Virtual Machines within the Virtual Machine Scale Set powering this Node Pool. Possible values are `Deallocate` and `Delete`. Changing this forces a new resource to be created.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * The ID of the Kubernetes Cluster where this Node Pool should exist. Changing this forces a new resource to be created.
     */
    readonly kubernetesClusterId: pulumi.Input<string>;
    /**
     * The maximum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be greater than or equal to `minCount`.
     */
    readonly maxCount?: pulumi.Input<number>;
    /**
     * The maximum number of pods that can run on each agent. Changing this forces a new resource to be created.
     */
    readonly maxPods?: pulumi.Input<number>;
    /**
     * The minimum number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be less than or equal to `maxCount`.
     */
    readonly minCount?: pulumi.Input<number>;
    /**
     * Should this Node Pool be used for System or User resources? Possible values are `System` and `User`. Defaults to `User`.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The name of the Node Pool which should be created within the Kubernetes Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The initial number of nodes which should exist within this Node Pool. Valid values are between `0` and `1000` and must be a value in the range `minCount` - `maxCount`.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * A map of Kubernetes labels which should be applied to nodes in this Node Pool. Changing this forces a new resource to be created.
     */
    readonly nodeLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of Kubernetes taints which should be applied to nodes in the agent pool (e.g `key=value:NoSchedule`). Changing this forces a new resource to be created.
     */
    readonly nodeTaints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Version of Kubernetes used for the Agents. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade)
     */
    readonly orchestratorVersion?: pulumi.Input<string>;
    /**
     * The Agent Operating System disk size in GB. Changing this forces a new resource to be created.
     */
    readonly osDiskSizeGb?: pulumi.Input<number>;
    /**
     * The type of disk which should be used for the Operating System. Possible values are `Ephemeral` and `Managed`. Defaults to `Managed`. Changing this forces a new resource to be created.
     */
    readonly osDiskType?: pulumi.Input<string>;
    /**
     * The Operating System which should be used for this Node Pool. Changing this forces a new resource to be created. Possible values are `Linux` and `Windows`. Defaults to `Linux`.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The Priority for Virtual Machines within the Virtual Machine Scale Set that powers this Node Pool. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group where the Virtual Machine Scale Set that powers this Node Pool will be placed. Changing this forces a new resource to be created.
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The maximum price you're willing to pay in USD per Virtual Machine. Valid values are `-1` (the current on-demand price for a Virtual Machine) or a positive value with up to five decimal places. Changing this forces a new resource to be created.
     */
    readonly spotMaxPrice?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `upgradeSettings` block as documented below.
     */
    readonly upgradeSettings?: pulumi.Input<inputs.containerservice.KubernetesClusterNodePoolUpgradeSettings>;
    /**
     * The SKU which should be used for the Virtual Machines used in this Node Pool. Changing this forces a new resource to be created.
     */
    readonly vmSize: pulumi.Input<string>;
    /**
     * The ID of the Subnet where this Node Pool should exist.
     */
    readonly vnetSubnetId?: pulumi.Input<string>;
}
