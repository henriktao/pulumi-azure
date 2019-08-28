// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Azure recovery vault fabric.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const primary = new azure.core.ResourceGroup("primary", {
 *     location: "West US",
 *     name: "tfex-network-mapping-primary",
 * });
 * const secondary = new azure.core.ResourceGroup("secondary", {
 *     location: "East US",
 *     name: "tfex-network-mapping-secondary",
 * });
 * const vault = new azure.recoveryservices.Vault("vault", {
 *     location: secondary.location,
 *     name: "example-recovery-vault",
 *     resourceGroupName: secondary.name,
 *     sku: "Standard",
 * });
 * const fabric = new azure.recoveryservices.Fabric("fabric", {
 *     location: primary.location,
 *     name: "primary-fabric",
 *     recoveryVaultName: vault.name,
 *     resourceGroupName: secondary.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/recovery_services_fabric.html.markdown.
 */
export class Fabric extends pulumi.CustomResource {
    /**
     * Get an existing Fabric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FabricState, opts?: pulumi.CustomResourceOptions): Fabric {
        return new Fabric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:recoveryservices/fabric:Fabric';

    /**
     * Returns true if the given object is an instance of Fabric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fabric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fabric.__pulumiType;
    }

    /**
     * In what region should the fabric be located.
     */
    public readonly location!: pulumi.Output<string>;
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
     * Create a Fabric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FabricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FabricArgs | FabricState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FabricState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as FabricArgs | undefined;
            if (!args || args.recoveryVaultName === undefined) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Fabric.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fabric resources.
 */
export interface FabricState {
    /**
     * In what region should the fabric be located.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the network mapping.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    readonly recoveryVaultName?: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fabric resource.
 */
export interface FabricArgs {
    /**
     * In what region should the fabric be located.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the network mapping.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the vault that should be updated.
     */
    readonly recoveryVaultName: pulumi.Input<string>;
    /**
     * Name of the resource group where the vault that should be updated is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
