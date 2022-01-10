// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Azure Site Recovery protection container. Protection containers serve as containers for replicated VMs and belong to a single region / recovery fabric. Protection containers can contain more than one replicated VM. To replicate a VM, a container must exist in both the source and target Azure regions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const primary = new azure.core.ResourceGroup("primary", {location: "West US"});
 * const secondary = new azure.core.ResourceGroup("secondary", {location: "East US"});
 * const vault = new azure.recoveryservices.Vault("vault", {
 *     location: secondary.location,
 *     resourceGroupName: secondary.name,
 *     sku: "Standard",
 * });
 * const fabric = new azure.siterecovery.Fabric("fabric", {
 *     resourceGroupName: secondary.name,
 *     recoveryVaultName: vault.name,
 *     location: primary.location,
 * });
 * const protection_container = new azure.siterecovery.ProtectionContainer("protection-container", {
 *     resourceGroupName: secondary.name,
 *     recoveryVaultName: vault.name,
 *     recoveryFabricName: fabric.name,
 * });
 * ```
 *
 * ## Import
 *
 * Site Recovery Protection Containers can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:siterecovery/protectionContainer:ProtectionContainer mycontainer /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name
 * ```
 */
export class ProtectionContainer extends pulumi.CustomResource {
    /**
     * Get an existing ProtectionContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProtectionContainerState, opts?: pulumi.CustomResourceOptions): ProtectionContainer {
        return new ProtectionContainer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:siterecovery/protectionContainer:ProtectionContainer';

    /**
     * Returns true if the given object is an instance of ProtectionContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProtectionContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProtectionContainer.__pulumiType;
    }

    /**
     * The name of the protection container.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of fabric that should contain this protection container.
     */
    public readonly recoveryFabricName!: pulumi.Output<string>;
    /**
     * The name of the vault that should be updated.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a ProtectionContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProtectionContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProtectionContainerArgs | ProtectionContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProtectionContainerState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recoveryFabricName"] = state ? state.recoveryFabricName : undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ProtectionContainerArgs | undefined;
            if ((!args || args.recoveryFabricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recoveryFabricName'");
            }
            if ((!args || args.recoveryVaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["recoveryFabricName"] = args ? args.recoveryFabricName : undefined;
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProtectionContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProtectionContainer resources.
 */
export interface ProtectionContainerState {
    /**
     * The name of the protection container.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of fabric that should contain this protection container.
     */
    recoveryFabricName?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    recoveryVaultName?: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProtectionContainer resource.
 */
export interface ProtectionContainerArgs {
    /**
     * The name of the protection container.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of fabric that should contain this protection container.
     */
    recoveryFabricName: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    recoveryVaultName: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    resourceGroupName: pulumi.Input<string>;
}
