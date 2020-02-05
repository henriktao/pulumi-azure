// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a VM replicated using Azure Site Recovery (Azure to Azure only). A replicated VM keeps a copiously updated image of the VM in another region in order to be able to start the VM in that region in case of a disaster.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const primaryResourceGroup = new azure.core.ResourceGroup("primary", {
 *     location: "West US",
 * });
 * const secondaryResourceGroup = new azure.core.ResourceGroup("secondary", {
 *     location: "East US",
 * });
 * const primaryVirtualNetwork = new azure.network.VirtualNetwork("primary", {
 *     addressSpaces: ["192.168.1.0/24"],
 *     location: primaryResourceGroup.location,
 *     resourceGroupName: primaryResourceGroup.name,
 * });
 * const primarySubnet = new azure.network.Subnet("primary", {
 *     addressPrefix: "192.168.1.0/24",
 *     resourceGroupName: primaryResourceGroup.name,
 *     virtualNetworkName: primaryVirtualNetwork.name,
 * });
 * const vmNetworkInterface = new azure.network.NetworkInterface("vm", {
 *     ipConfigurations: [{
 *         name: "vm",
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: primarySubnet.id,
 *     }],
 *     location: primaryResourceGroup.location,
 *     resourceGroupName: primaryResourceGroup.name,
 * });
 * const vmVirtualMachine = new azure.compute.VirtualMachine("vm", {
 *     location: primaryResourceGroup.location,
 *     networkInterfaceIds: [vmNetworkInterface.id],
 *     osProfile: {
 *         adminPassword: "test-pwd-123",
 *         adminUsername: "test-admin-123",
 *         computerName: "vm",
 *     },
 *     osProfileLinuxConfig: {
 *         disablePasswordAuthentication: false,
 *     },
 *     resourceGroupName: primaryResourceGroup.name,
 *     storageImageReference: {
 *         offer: "CentOS",
 *         publisher: "OpenLogic",
 *         sku: "7.5",
 *         version: "latest",
 *     },
 *     storageOsDisk: {
 *         caching: "ReadWrite",
 *         createOption: "FromImage",
 *         managedDiskType: "Premium_LRS",
 *         name: "vm-os-disk",
 *         osType: "Linux",
 *     },
 *     vmSize: "Standard_B1s",
 * });
 * const vault = new azure.recoveryservices.Vault("vault", {
 *     location: secondaryResourceGroup.location,
 *     resourceGroupName: secondaryResourceGroup.name,
 *     sku: "Standard",
 * });
 * const primaryFabric = new azure.siterecovery.Fabric("primary", {
 *     location: primaryResourceGroup.location,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const secondaryFabric = new azure.siterecovery.Fabric("secondary", {
 *     location: secondaryResourceGroup.location,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const primaryProtectionContainer = new azure.siterecovery.ProtectionContainer("primary", {
 *     recoveryFabricName: primaryFabric.name,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const secondaryProtectionContainer = new azure.siterecovery.ProtectionContainer("secondary", {
 *     recoveryFabricName: secondaryFabric.name,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const policy = new azure.siterecovery.ReplicationPolicy("policy", {
 *     applicationConsistentSnapshotFrequencyInMinutes: (4 * 60),
 *     recoveryPointRetentionInMinutes: (24 * 60),
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const containerMapping = new azure.siterecovery.ProtectionContainerMapping("container-mapping", {
 *     recoveryFabricName: primaryFabric.name,
 *     recoveryReplicationPolicyId: policy.id,
 *     recoverySourceProtectionContainerName: primaryProtectionContainer.name,
 *     recoveryTargetProtectionContainerId: secondaryProtectionContainer.id,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 * });
 * const primaryAccount = new azure.storage.Account("primary", {
 *     accountReplicationType: "LRS",
 *     accountTier: "Standard",
 *     location: primaryResourceGroup.location,
 *     resourceGroupName: primaryResourceGroup.name,
 * });
 * const vmReplication = new azure.siterecovery.ReplicatedVM("vm-replication", {
 *     managedDisks: [{
 *         diskId: vmVirtualMachine.storageOsDisk.managedDiskId!,
 *         stagingStorageAccountId: primaryAccount.id,
 *         targetDiskType: "Premium_LRS",
 *         targetReplicaDiskType: "Premium_LRS",
 *         targetResourceGroupId: secondaryResourceGroup.id,
 *     }],
 *     recoveryReplicationPolicyId: policy.id,
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondaryResourceGroup.name,
 *     sourceRecoveryFabricName: primaryFabric.name,
 *     sourceRecoveryProtectionContainerName: primaryProtectionContainer.name,
 *     sourceVmId: vmVirtualMachine.id,
 *     targetRecoveryFabricId: secondaryFabric.id,
 *     targetRecoveryProtectionContainerId: secondaryProtectionContainer.id,
 *     targetResourceGroupId: secondaryResourceGroup.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/site_recovery_replicated_vm.html.markdown.
 */
export class ReplicatedVM extends pulumi.CustomResource {
    /**
     * Get an existing ReplicatedVM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicatedVMState, opts?: pulumi.CustomResourceOptions): ReplicatedVM {
        return new ReplicatedVM(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:siterecovery/replicatedVM:ReplicatedVM';

    /**
     * Returns true if the given object is an instance of ReplicatedVM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicatedVM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicatedVM.__pulumiType;
    }

    /**
     * One or more `managedDisk` block.
     */
    public readonly managedDisks!: pulumi.Output<outputs.siterecovery.ReplicatedVMManagedDisk[] | undefined>;
    /**
     * The name of the network mapping.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly recoveryReplicationPolicyId!: pulumi.Output<string>;
    /**
     * The name of the vault that should be updated.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Name of fabric that should contains this replication.
     */
    public readonly sourceRecoveryFabricName!: pulumi.Output<string>;
    /**
     * Name of the protection container to use.
     */
    public readonly sourceRecoveryProtectionContainerName!: pulumi.Output<string>;
    /**
     * Id of the VM to replicate
     */
    public readonly sourceVmId!: pulumi.Output<string>;
    /**
     * Id of availability set that the new VM should belong to when a failover is done.
     */
    public readonly targetAvailabilitySetId!: pulumi.Output<string | undefined>;
    /**
     * Id of fabric where the VM replication should be handled when a failover is done.
     */
    public readonly targetRecoveryFabricId!: pulumi.Output<string>;
    /**
     * Id of protection container where the VM replication should be created when a failover is done.
     */
    public readonly targetRecoveryProtectionContainerId!: pulumi.Output<string>;
    /**
     * Id of resource group where the VM should be created when a failover is done.
     */
    public readonly targetResourceGroupId!: pulumi.Output<string>;

    /**
     * Create a ReplicatedVM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicatedVMArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicatedVMArgs | ReplicatedVMState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReplicatedVMState | undefined;
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
            const args = argsOrState as ReplicatedVMArgs | undefined;
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
        super(ReplicatedVM.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicatedVM resources.
 */
export interface ReplicatedVMState {
    /**
     * One or more `managedDisk` block.
     */
    readonly managedDisks?: pulumi.Input<pulumi.Input<inputs.siterecovery.ReplicatedVMManagedDisk>[]>;
    /**
     * The name of the network mapping.
     */
    readonly name?: pulumi.Input<string>;
    readonly recoveryReplicationPolicyId?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    readonly recoveryVaultName?: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Name of fabric that should contains this replication.
     */
    readonly sourceRecoveryFabricName?: pulumi.Input<string>;
    /**
     * Name of the protection container to use.
     */
    readonly sourceRecoveryProtectionContainerName?: pulumi.Input<string>;
    /**
     * Id of the VM to replicate
     */
    readonly sourceVmId?: pulumi.Input<string>;
    /**
     * Id of availability set that the new VM should belong to when a failover is done.
     */
    readonly targetAvailabilitySetId?: pulumi.Input<string>;
    /**
     * Id of fabric where the VM replication should be handled when a failover is done.
     */
    readonly targetRecoveryFabricId?: pulumi.Input<string>;
    /**
     * Id of protection container where the VM replication should be created when a failover is done.
     */
    readonly targetRecoveryProtectionContainerId?: pulumi.Input<string>;
    /**
     * Id of resource group where the VM should be created when a failover is done.
     */
    readonly targetResourceGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicatedVM resource.
 */
export interface ReplicatedVMArgs {
    /**
     * One or more `managedDisk` block.
     */
    readonly managedDisks?: pulumi.Input<pulumi.Input<inputs.siterecovery.ReplicatedVMManagedDisk>[]>;
    /**
     * The name of the network mapping.
     */
    readonly name?: pulumi.Input<string>;
    readonly recoveryReplicationPolicyId: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    readonly recoveryVaultName: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of fabric that should contains this replication.
     */
    readonly sourceRecoveryFabricName: pulumi.Input<string>;
    /**
     * Name of the protection container to use.
     */
    readonly sourceRecoveryProtectionContainerName: pulumi.Input<string>;
    /**
     * Id of the VM to replicate
     */
    readonly sourceVmId: pulumi.Input<string>;
    /**
     * Id of availability set that the new VM should belong to when a failover is done.
     */
    readonly targetAvailabilitySetId?: pulumi.Input<string>;
    /**
     * Id of fabric where the VM replication should be handled when a failover is done.
     */
    readonly targetRecoveryFabricId: pulumi.Input<string>;
    /**
     * Id of protection container where the VM replication should be created when a failover is done.
     */
    readonly targetRecoveryProtectionContainerId: pulumi.Input<string>;
    /**
     * Id of resource group where the VM should be created when a failover is done.
     */
    readonly targetResourceGroupId: pulumi.Input<string>;
}
