// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a site recovery network mapping on Azure. A network mapping decides how to translate connected netwroks when a VM is migrated from one region to another.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const primaryResourceGroup = new azure.core.ResourceGroup("primaryResourceGroup", {location: "West US"});
 * const secondaryResourceGroup = new azure.core.ResourceGroup("secondaryResourceGroup", {location: "East US"});
 * const vault = new azure.recoveryservices.Vault("vault", {
 *     location: secondaryResourceGroup.location,
 *     resourceGroupName: secondaryResourceGroup.name,
 *     sku: "Standard",
 * });
 * const primaryFabric = new azure.siterecovery.Fabric("primaryFabric", {
 *     resourceGroupName: secondaryResourceGroup.name,
 *     recoveryVaultName: vault.name,
 *     location: primaryResourceGroup.location,
 * });
 * const secondaryFabric = new azure.siterecovery.Fabric("secondaryFabric", {
 *     resourceGroupName: secondaryResourceGroup.name,
 *     recoveryVaultName: vault.name,
 *     location: secondaryResourceGroup.location,
 * }, {
 *     dependsOn: [primaryFabric],
 * });
 * // Avoids issues with creating fabrics simultaneously
 * const primaryVirtualNetwork = new azure.network.VirtualNetwork("primaryVirtualNetwork", {
 *     resourceGroupName: primaryResourceGroup.name,
 *     addressSpaces: ["192.168.1.0/24"],
 *     location: primaryResourceGroup.location,
 * });
 * const secondaryVirtualNetwork = new azure.network.VirtualNetwork("secondaryVirtualNetwork", {
 *     resourceGroupName: secondaryResourceGroup.name,
 *     addressSpaces: ["192.168.2.0/24"],
 *     location: secondaryResourceGroup.location,
 * });
 * const recovery_mapping = new azure.siterecovery.NetworkMapping("recovery-mapping", {
 *     resourceGroupName: secondaryResourceGroup.name,
 *     recoveryVaultName: vault.name,
 *     sourceRecoveryFabricName: "primary-fabric",
 *     targetRecoveryFabricName: "secondary-fabric",
 *     sourceNetworkId: primaryVirtualNetwork.id,
 *     targetNetworkId: secondaryVirtualNetwork.id,
 * });
 * ```
 *
 * ## Import
 *
 * Site Recovery Network Mapping can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:siterecovery/networkMapping:NetworkMapping mymapping /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/primary-fabric-name/replicationNetworks/azureNetwork/replicationNetworkMappings/mapping-name
 * ```
 */
export class NetworkMapping extends pulumi.CustomResource {
    /**
     * Get an existing NetworkMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkMappingState, opts?: pulumi.CustomResourceOptions): NetworkMapping {
        return new NetworkMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:siterecovery/networkMapping:NetworkMapping';

    /**
     * Returns true if the given object is an instance of NetworkMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkMapping.__pulumiType;
    }

    /**
     * The name of the network mapping.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the vault that should be updated.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The id of the primary network.
     */
    public readonly sourceNetworkId!: pulumi.Output<string>;
    /**
     * Specifies the ASR fabric where mapping should be created.
     */
    public readonly sourceRecoveryFabricName!: pulumi.Output<string>;
    /**
     * The id of the recovery network.
     */
    public readonly targetNetworkId!: pulumi.Output<string>;
    /**
     * The Azure Site Recovery fabric object corresponding to the recovery Azure region.
     */
    public readonly targetRecoveryFabricName!: pulumi.Output<string>;

    /**
     * Create a NetworkMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkMappingArgs | NetworkMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkMappingState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["sourceNetworkId"] = state ? state.sourceNetworkId : undefined;
            resourceInputs["sourceRecoveryFabricName"] = state ? state.sourceRecoveryFabricName : undefined;
            resourceInputs["targetNetworkId"] = state ? state.targetNetworkId : undefined;
            resourceInputs["targetRecoveryFabricName"] = state ? state.targetRecoveryFabricName : undefined;
        } else {
            const args = argsOrState as NetworkMappingArgs | undefined;
            if ((!args || args.recoveryVaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sourceNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceNetworkId'");
            }
            if ((!args || args.sourceRecoveryFabricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceRecoveryFabricName'");
            }
            if ((!args || args.targetNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetNetworkId'");
            }
            if ((!args || args.targetRecoveryFabricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetRecoveryFabricName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sourceNetworkId"] = args ? args.sourceNetworkId : undefined;
            resourceInputs["sourceRecoveryFabricName"] = args ? args.sourceRecoveryFabricName : undefined;
            resourceInputs["targetNetworkId"] = args ? args.targetNetworkId : undefined;
            resourceInputs["targetRecoveryFabricName"] = args ? args.targetRecoveryFabricName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkMapping resources.
 */
export interface NetworkMappingState {
    /**
     * The name of the network mapping.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    recoveryVaultName?: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The id of the primary network.
     */
    sourceNetworkId?: pulumi.Input<string>;
    /**
     * Specifies the ASR fabric where mapping should be created.
     */
    sourceRecoveryFabricName?: pulumi.Input<string>;
    /**
     * The id of the recovery network.
     */
    targetNetworkId?: pulumi.Input<string>;
    /**
     * The Azure Site Recovery fabric object corresponding to the recovery Azure region.
     */
    targetRecoveryFabricName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkMapping resource.
 */
export interface NetworkMappingArgs {
    /**
     * The name of the network mapping.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    recoveryVaultName: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The id of the primary network.
     */
    sourceNetworkId: pulumi.Input<string>;
    /**
     * Specifies the ASR fabric where mapping should be created.
     */
    sourceRecoveryFabricName: pulumi.Input<string>;
    /**
     * The id of the recovery network.
     */
    targetNetworkId: pulumi.Input<string>;
    /**
     * The Azure Site Recovery fabric object corresponding to the recovery Azure region.
     */
    targetRecoveryFabricName: pulumi.Input<string>;
}
