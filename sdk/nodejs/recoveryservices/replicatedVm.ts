// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ReplicatedVm extends pulumi.CustomResource {
    /**
     * Get an existing ReplicatedVm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicatedVmState, opts?: pulumi.CustomResourceOptions): ReplicatedVm {
        return new ReplicatedVm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:recoveryservices/replicatedVm:ReplicatedVm';

    /**
     * Returns true if the given object is an instance of ReplicatedVm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicatedVm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicatedVm.__pulumiType;
    }

    public readonly managedDisks!: pulumi.Output<outputs.recoveryservices.ReplicatedVmManagedDisk[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly recoveryReplicationPolicyId!: pulumi.Output<string>;
    public readonly recoveryVaultName!: pulumi.Output<string>;
    public readonly resourceGroupName!: pulumi.Output<string>;
    public readonly sourceRecoveryFabricName!: pulumi.Output<string>;
    public readonly sourceRecoveryProtectionContainerName!: pulumi.Output<string>;
    public readonly sourceVmId!: pulumi.Output<string>;
    public readonly targetAvailabilitySetId!: pulumi.Output<string | undefined>;
    public readonly targetRecoveryFabricId!: pulumi.Output<string>;
    public readonly targetRecoveryProtectionContainerId!: pulumi.Output<string>;
    public readonly targetResourceGroupId!: pulumi.Output<string>;

    /**
     * Create a ReplicatedVm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicatedVmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicatedVmArgs | ReplicatedVmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReplicatedVmState | undefined;
            inputs["managedDisks"] = state ? state.managedDisks : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recoveryReplicationPolicyId"] = state ? state.recoveryReplicationPolicyId : undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sourceRecoveryFabricName"] = state ? state.sourceRecoveryFabricName : undefined;
            inputs["sourceRecoveryProtectionContainerName"] = state ? state.sourceRecoveryProtectionContainerName : undefined;
            inputs["sourceVmId"] = state ? state.sourceVmId : undefined;
            inputs["targetAvailabilitySetId"] = state ? state.targetAvailabilitySetId : undefined;
            inputs["targetRecoveryFabricId"] = state ? state.targetRecoveryFabricId : undefined;
            inputs["targetRecoveryProtectionContainerId"] = state ? state.targetRecoveryProtectionContainerId : undefined;
            inputs["targetResourceGroupId"] = state ? state.targetResourceGroupId : undefined;
        } else {
            const args = argsOrState as ReplicatedVmArgs | undefined;
            if (!args || args.recoveryReplicationPolicyId === undefined) {
                throw new Error("Missing required property 'recoveryReplicationPolicyId'");
            }
            if (!args || args.recoveryVaultName === undefined) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sourceRecoveryFabricName === undefined) {
                throw new Error("Missing required property 'sourceRecoveryFabricName'");
            }
            if (!args || args.sourceRecoveryProtectionContainerName === undefined) {
                throw new Error("Missing required property 'sourceRecoveryProtectionContainerName'");
            }
            if (!args || args.sourceVmId === undefined) {
                throw new Error("Missing required property 'sourceVmId'");
            }
            if (!args || args.targetRecoveryFabricId === undefined) {
                throw new Error("Missing required property 'targetRecoveryFabricId'");
            }
            if (!args || args.targetRecoveryProtectionContainerId === undefined) {
                throw new Error("Missing required property 'targetRecoveryProtectionContainerId'");
            }
            if (!args || args.targetResourceGroupId === undefined) {
                throw new Error("Missing required property 'targetResourceGroupId'");
            }
            inputs["managedDisks"] = args ? args.managedDisks : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recoveryReplicationPolicyId"] = args ? args.recoveryReplicationPolicyId : undefined;
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceRecoveryFabricName"] = args ? args.sourceRecoveryFabricName : undefined;
            inputs["sourceRecoveryProtectionContainerName"] = args ? args.sourceRecoveryProtectionContainerName : undefined;
            inputs["sourceVmId"] = args ? args.sourceVmId : undefined;
            inputs["targetAvailabilitySetId"] = args ? args.targetAvailabilitySetId : undefined;
            inputs["targetRecoveryFabricId"] = args ? args.targetRecoveryFabricId : undefined;
            inputs["targetRecoveryProtectionContainerId"] = args ? args.targetRecoveryProtectionContainerId : undefined;
            inputs["targetResourceGroupId"] = args ? args.targetResourceGroupId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReplicatedVm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicatedVm resources.
 */
export interface ReplicatedVmState {
    readonly managedDisks?: pulumi.Input<pulumi.Input<inputs.recoveryservices.ReplicatedVmManagedDisk>[]>;
    readonly name?: pulumi.Input<string>;
    readonly recoveryReplicationPolicyId?: pulumi.Input<string>;
    readonly recoveryVaultName?: pulumi.Input<string>;
    readonly resourceGroupName?: pulumi.Input<string>;
    readonly sourceRecoveryFabricName?: pulumi.Input<string>;
    readonly sourceRecoveryProtectionContainerName?: pulumi.Input<string>;
    readonly sourceVmId?: pulumi.Input<string>;
    readonly targetAvailabilitySetId?: pulumi.Input<string>;
    readonly targetRecoveryFabricId?: pulumi.Input<string>;
    readonly targetRecoveryProtectionContainerId?: pulumi.Input<string>;
    readonly targetResourceGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicatedVm resource.
 */
export interface ReplicatedVmArgs {
    readonly managedDisks?: pulumi.Input<pulumi.Input<inputs.recoveryservices.ReplicatedVmManagedDisk>[]>;
    readonly name?: pulumi.Input<string>;
    readonly recoveryReplicationPolicyId: pulumi.Input<string>;
    readonly recoveryVaultName: pulumi.Input<string>;
    readonly resourceGroupName: pulumi.Input<string>;
    readonly sourceRecoveryFabricName: pulumi.Input<string>;
    readonly sourceRecoveryProtectionContainerName: pulumi.Input<string>;
    readonly sourceVmId: pulumi.Input<string>;
    readonly targetAvailabilitySetId?: pulumi.Input<string>;
    readonly targetRecoveryFabricId: pulumi.Input<string>;
    readonly targetRecoveryProtectionContainerId: pulumi.Input<string>;
    readonly targetResourceGroupId: pulumi.Input<string>;
}
