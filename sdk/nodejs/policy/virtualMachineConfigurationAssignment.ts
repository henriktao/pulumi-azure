// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Applies a Guest Configuration Policy to a Virtual Machine.
 *
 * > **NOTE:** You can create Guest Configuration Policies without defining a `azure.compute.Extension` resource, however the policies will not be executed until a `azure.compute.Extension` has been provisioned to the virtual machine.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("exampleNetworkInterface", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     ipConfigurations: [{
 *         name: "internal",
 *         subnetId: exampleSubnet.id,
 *         privateIpAddressAllocation: "Dynamic",
 *     }],
 * });
 * const exampleWindowsVirtualMachine = new azure.compute.WindowsVirtualMachine("exampleWindowsVirtualMachine", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     size: "Standard_F2",
 *     adminUsername: "adminuser",
 *     adminPassword: `P@$$w0rd1234!`,
 *     networkInterfaceIds: [exampleNetworkInterface.id],
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 *     osDisk: {
 *         caching: "ReadWrite",
 *         storageAccountType: "Standard_LRS",
 *     },
 *     sourceImageReference: {
 *         publisher: "MicrosoftWindowsServer",
 *         offer: "WindowsServer",
 *         sku: "2019-Datacenter",
 *         version: "latest",
 *     },
 * });
 * const exampleExtension = new azure.compute.Extension("exampleExtension", {
 *     virtualMachineId: exampleWindowsVirtualMachine.id,
 *     publisher: "Microsoft.GuestConfiguration",
 *     type: "ConfigurationforWindows",
 *     typeHandlerVersion: "1.0",
 *     autoUpgradeMinorVersion: "true",
 * });
 * const exampleVirtualMachineConfigurationAssignment = new azure.policy.VirtualMachineConfigurationAssignment("exampleVirtualMachineConfigurationAssignment", {
 *     location: exampleWindowsVirtualMachine.location,
 *     virtualMachineId: exampleWindowsVirtualMachine.id,
 *     configuration: {
 *         assignmentType: "ApplyAndMonitor",
 *         version: "1.*",
 *         parameters: [
 *             {
 *                 name: "Minimum Password Length;ExpectedValue",
 *                 value: "16",
 *             },
 *             {
 *                 name: "Minimum Password Age;ExpectedValue",
 *                 value: "0",
 *             },
 *             {
 *                 name: "Maximum Password Age;ExpectedValue",
 *                 value: "30,45",
 *             },
 *             {
 *                 name: "Enforce Password History;ExpectedValue",
 *                 value: "10",
 *             },
 *             {
 *                 name: "Password Must Meet Complexity Requirements;ExpectedValue",
 *                 value: "1",
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Policy Virtual Machine Configuration Assignments can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:policy/virtualMachineConfigurationAssignment:VirtualMachineConfigurationAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
 * ```
 */
export class VirtualMachineConfigurationAssignment extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineConfigurationAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualMachineConfigurationAssignmentState, opts?: pulumi.CustomResourceOptions): VirtualMachineConfigurationAssignment {
        return new VirtualMachineConfigurationAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:policy/virtualMachineConfigurationAssignment:VirtualMachineConfigurationAssignment';

    /**
     * Returns true if the given object is an instance of VirtualMachineConfigurationAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineConfigurationAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineConfigurationAssignment.__pulumiType;
    }

    /**
     * A `configuration` block as defined below.
     */
    public readonly configuration!: pulumi.Output<outputs.policy.VirtualMachineConfigurationAssignmentConfiguration>;
    /**
     * The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Guest Configuration that will be assigned in this Guest Configuration Assignment. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the Policy Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a VirtualMachineConfigurationAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineConfigurationAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMachineConfigurationAssignmentArgs | VirtualMachineConfigurationAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualMachineConfigurationAssignmentState | undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as VirtualMachineConfigurationAssignmentArgs | undefined;
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.virtualMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VirtualMachineConfigurationAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualMachineConfigurationAssignment resources.
 */
export interface VirtualMachineConfigurationAssignmentState {
    /**
     * A `configuration` block as defined below.
     */
    configuration?: pulumi.Input<inputs.policy.VirtualMachineConfigurationAssignmentConfiguration>;
    /**
     * The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Guest Configuration that will be assigned in this Guest Configuration Assignment. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource ID of the Policy Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
     */
    virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualMachineConfigurationAssignment resource.
 */
export interface VirtualMachineConfigurationAssignmentArgs {
    /**
     * A `configuration` block as defined below.
     */
    configuration: pulumi.Input<inputs.policy.VirtualMachineConfigurationAssignmentConfiguration>;
    /**
     * The Azure location where the Policy Virtual Machine Configuration Assignment should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Guest Configuration that will be assigned in this Guest Configuration Assignment. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource ID of the Policy Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
     */
    virtualMachineId: pulumi.Input<string>;
}
