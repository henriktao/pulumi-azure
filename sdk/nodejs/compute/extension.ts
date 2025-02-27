// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Machine Extension to provide post deployment configuration
 * and run automated tasks.
 *
 * > **NOTE:** Custom Script Extensions for Linux & Windows require that the `commandToExecute` returns a `0` exit code to be classified as successfully deployed. You can achieve this by appending `exit 0` to the end of your `commandToExecute`.
 *
 * > **NOTE:** Custom Script Extensions require that the Azure Virtual Machine Guest Agent is running on the Virtual Machine.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("exampleNetworkInterface", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipConfigurations: [{
 *         name: "testconfiguration1",
 *         subnetId: exampleSubnet.id,
 *         privateIpAddressAllocation: "Dynamic",
 *     }],
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const exampleContainer = new azure.storage.Container("exampleContainer", {
 *     storageAccountName: exampleAccount.name,
 *     containerAccessType: "private",
 * });
 * const exampleVirtualMachine = new azure.compute.VirtualMachine("exampleVirtualMachine", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     networkInterfaceIds: [exampleNetworkInterface.id],
 *     vmSize: "Standard_F2",
 *     storageImageReference: {
 *         publisher: "Canonical",
 *         offer: "UbuntuServer",
 *         sku: "16.04-LTS",
 *         version: "latest",
 *     },
 *     storageOsDisk: {
 *         name: "myosdisk1",
 *         vhdUri: pulumi.interpolate`${exampleAccount.primaryBlobEndpoint}${exampleContainer.name}/myosdisk1.vhd`,
 *         caching: "ReadWrite",
 *         createOption: "FromImage",
 *     },
 *     osProfile: {
 *         computerName: "hostname",
 *         adminUsername: "testadmin",
 *         adminPassword: "Password1234!",
 *     },
 *     osProfileLinuxConfig: {
 *         disablePasswordAuthentication: false,
 *     },
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const exampleExtension = new azure.compute.Extension("exampleExtension", {
 *     virtualMachineId: exampleVirtualMachine.id,
 *     publisher: "Microsoft.Azure.Extensions",
 *     type: "CustomScript",
 *     typeHandlerVersion: "2.0",
 *     settings: `	{
 * 		"commandToExecute": "hostname && uptime"
 * 	}
 * `,
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Machine Extensions can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/extension:Extension example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/myVM/extensions/extensionName
 * ```
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionState, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/extension:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    public readonly autoUpgradeMinorVersion!: pulumi.Output<boolean | undefined>;
    /**
     * Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
     */
    public readonly automaticUpgradeEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    public readonly protectedSettings!: pulumi.Output<string | undefined>;
    /**
     * The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
     */
    public readonly publisher!: pulumi.Output<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    public readonly settings!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    public readonly typeHandlerVersion!: pulumi.Output<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionArgs | ExtensionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionState | undefined;
            resourceInputs["autoUpgradeMinorVersion"] = state ? state.autoUpgradeMinorVersion : undefined;
            resourceInputs["automaticUpgradeEnabled"] = state ? state.automaticUpgradeEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protectedSettings"] = state ? state.protectedSettings : undefined;
            resourceInputs["publisher"] = state ? state.publisher : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["typeHandlerVersion"] = state ? state.typeHandlerVersion : undefined;
            resourceInputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as ExtensionArgs | undefined;
            if ((!args || args.publisher === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publisher'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.typeHandlerVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'typeHandlerVersion'");
            }
            if ((!args || args.virtualMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            resourceInputs["autoUpgradeMinorVersion"] = args ? args.autoUpgradeMinorVersion : undefined;
            resourceInputs["automaticUpgradeEnabled"] = args ? args.automaticUpgradeEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protectedSettings"] = args ? args.protectedSettings : undefined;
            resourceInputs["publisher"] = args ? args.publisher : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["typeHandlerVersion"] = args ? args.typeHandlerVersion : undefined;
            resourceInputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Extension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extension resources.
 */
export interface ExtensionState {
    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    autoUpgradeMinorVersion?: pulumi.Input<boolean>;
    /**
     * Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
     */
    automaticUpgradeEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    protectedSettings?: pulumi.Input<string>;
    /**
     * The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
     */
    publisher?: pulumi.Input<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    settings?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    typeHandlerVersion?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    /**
     * Specifies if the platform deploys
     * the latest minor version update to the `typeHandlerVersion` specified.
     */
    autoUpgradeMinorVersion?: pulumi.Input<boolean>;
    /**
     * Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
     */
    automaticUpgradeEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the virtual machine extension peering. Changing
     * this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The protectedSettings passed to the
     * extension, like settings, these are specified as a JSON object in a string.
     */
    protectedSettings?: pulumi.Input<string>;
    /**
     * The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
     */
    publisher: pulumi.Input<string>;
    /**
     * The settings passed to the extension, these are
     * specified as a JSON object in a string.
     */
    settings?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of extension, available types for a publisher can
     * be found using the Azure CLI.
     */
    type: pulumi.Input<string>;
    /**
     * Specifies the version of the extension to
     * use, available versions can be found using the Azure CLI.
     */
    typeHandlerVersion: pulumi.Input<string>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created
     */
    virtualMachineId: pulumi.Input<string>;
}
