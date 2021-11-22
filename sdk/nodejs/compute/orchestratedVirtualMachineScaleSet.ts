// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Orchestrated Virtual Machine Scale Set.
 *
 * ## Disclaimers
 *
 * > **NOTE:** All arguments including the administrator login and password will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * > **NOTE:** Orchestrated Virtual Machine Scale Sets are in Public Preview and it may receive breaking changes - [more details can be found in the Azure Documentation](https://docs.microsoft.com/azure/virtual-machine-scale-sets/orchestration-modes).
 *
 * > **NOTE:** Due to a bug in the service code `extensions` are not currently supported in the `azure.compute.OrchestratedVirtualMachineScaleSet` resource. The ETA for `extensions` support is tentatively set for January 15, 2022.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleOrchestratedVirtualMachineScaleSet = new azure.compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     platformFaultDomainCount: 1,
 *     zones: ["1"],
 * });
 * ```
 *
 * ## Import
 *
 * An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
 * ```
 */
export class OrchestratedVirtualMachineScaleSet extends pulumi.CustomResource {
    /**
     * Get an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrchestratedVirtualMachineScaleSetState, opts?: pulumi.CustomResourceOptions): OrchestratedVirtualMachineScaleSet {
        return new OrchestratedVirtualMachineScaleSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet';

    /**
     * Returns true if the given object is an instance of OrchestratedVirtualMachineScaleSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrchestratedVirtualMachineScaleSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrchestratedVirtualMachineScaleSet.__pulumiType;
    }

    public readonly automaticInstanceRepair!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair>;
    public readonly bootDiagnostics!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetBootDiagnostics | undefined>;
    public readonly dataDisks!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetDataDisk[] | undefined>;
    public readonly encryptionAtHostEnabled!: pulumi.Output<boolean | undefined>;
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    public readonly identity!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetIdentity | undefined>;
    /**
     * The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
     */
    public readonly instances!: pulumi.Output<number>;
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    public readonly maxBidPrice!: pulumi.Output<number | undefined>;
    /**
     * The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetNetworkInterface[] | undefined>;
    public readonly osDisk!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetOsDisk | undefined>;
    public readonly osProfile!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetOsProfile | undefined>;
    public readonly plan!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetPlan | undefined>;
    /**
     * Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    public readonly platformFaultDomainCount!: pulumi.Output<number>;
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    public readonly proximityPlacementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    public readonly skuName!: pulumi.Output<string | undefined>;
    public readonly sourceImageId!: pulumi.Output<string | undefined>;
    /**
     * A `sourceImageReference` block as defined below.
     */
    public readonly sourceImageReference!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetSourceImageReference | undefined>;
    /**
     * A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly terminationNotification!: pulumi.Output<outputs.compute.OrchestratedVirtualMachineScaleSetTerminationNotification>;
    /**
     * The Unique ID for the Orchestrated Virtual Machine Scale Set.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    public readonly zoneBalance!: pulumi.Output<boolean | undefined>;
    /**
     * A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
     */
    public readonly zones!: pulumi.Output<string | undefined>;

    /**
     * Create a OrchestratedVirtualMachineScaleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrchestratedVirtualMachineScaleSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrchestratedVirtualMachineScaleSetArgs | OrchestratedVirtualMachineScaleSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrchestratedVirtualMachineScaleSetState | undefined;
            inputs["automaticInstanceRepair"] = state ? state.automaticInstanceRepair : undefined;
            inputs["bootDiagnostics"] = state ? state.bootDiagnostics : undefined;
            inputs["dataDisks"] = state ? state.dataDisks : undefined;
            inputs["encryptionAtHostEnabled"] = state ? state.encryptionAtHostEnabled : undefined;
            inputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["licenseType"] = state ? state.licenseType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maxBidPrice"] = state ? state.maxBidPrice : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["osDisk"] = state ? state.osDisk : undefined;
            inputs["osProfile"] = state ? state.osProfile : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["platformFaultDomainCount"] = state ? state.platformFaultDomainCount : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["proximityPlacementGroupId"] = state ? state.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["sourceImageId"] = state ? state.sourceImageId : undefined;
            inputs["sourceImageReference"] = state ? state.sourceImageReference : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["terminationNotification"] = state ? state.terminationNotification : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
            inputs["zoneBalance"] = state ? state.zoneBalance : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as OrchestratedVirtualMachineScaleSetArgs | undefined;
            if ((!args || args.platformFaultDomainCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platformFaultDomainCount'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automaticInstanceRepair"] = args ? args.automaticInstanceRepair : undefined;
            inputs["bootDiagnostics"] = args ? args.bootDiagnostics : undefined;
            inputs["dataDisks"] = args ? args.dataDisks : undefined;
            inputs["encryptionAtHostEnabled"] = args ? args.encryptionAtHostEnabled : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxBidPrice"] = args ? args.maxBidPrice : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["osDisk"] = args ? args.osDisk : undefined;
            inputs["osProfile"] = args ? args.osProfile : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["platformFaultDomainCount"] = args ? args.platformFaultDomainCount : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["proximityPlacementGroupId"] = args ? args.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["sourceImageId"] = args ? args.sourceImageId : undefined;
            inputs["sourceImageReference"] = args ? args.sourceImageReference : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["terminationNotification"] = args ? args.terminationNotification : undefined;
            inputs["zoneBalance"] = args ? args.zoneBalance : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrchestratedVirtualMachineScaleSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrchestratedVirtualMachineScaleSet resources.
 */
export interface OrchestratedVirtualMachineScaleSetState {
    automaticInstanceRepair?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair>;
    bootDiagnostics?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetBootDiagnostics>;
    dataDisks?: pulumi.Input<pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetDataDisk>[]>;
    encryptionAtHostEnabled?: pulumi.Input<boolean>;
    evictionPolicy?: pulumi.Input<string>;
    identity?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetIdentity>;
    /**
     * The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
     */
    instances?: pulumi.Input<number>;
    licenseType?: pulumi.Input<string>;
    /**
     * The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    maxBidPrice?: pulumi.Input<number>;
    /**
     * The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetNetworkInterface>[]>;
    osDisk?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetOsDisk>;
    osProfile?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetOsProfile>;
    plan?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetPlan>;
    /**
     * Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    platformFaultDomainCount?: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    skuName?: pulumi.Input<string>;
    sourceImageId?: pulumi.Input<string>;
    /**
     * A `sourceImageReference` block as defined below.
     */
    sourceImageReference?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetSourceImageReference>;
    /**
     * A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    terminationNotification?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetTerminationNotification>;
    /**
     * The Unique ID for the Orchestrated Virtual Machine Scale Set.
     */
    uniqueId?: pulumi.Input<string>;
    zoneBalance?: pulumi.Input<boolean>;
    /**
     * A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
     */
    zones?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrchestratedVirtualMachineScaleSet resource.
 */
export interface OrchestratedVirtualMachineScaleSetArgs {
    automaticInstanceRepair?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair>;
    bootDiagnostics?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetBootDiagnostics>;
    dataDisks?: pulumi.Input<pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetDataDisk>[]>;
    encryptionAtHostEnabled?: pulumi.Input<boolean>;
    evictionPolicy?: pulumi.Input<string>;
    identity?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetIdentity>;
    /**
     * The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
     */
    instances?: pulumi.Input<number>;
    licenseType?: pulumi.Input<string>;
    /**
     * The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    maxBidPrice?: pulumi.Input<number>;
    /**
     * The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetNetworkInterface>[]>;
    osDisk?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetOsDisk>;
    osProfile?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetOsProfile>;
    plan?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetPlan>;
    /**
     * Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
     */
    platformFaultDomainCount: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    skuName?: pulumi.Input<string>;
    sourceImageId?: pulumi.Input<string>;
    /**
     * A `sourceImageReference` block as defined below.
     */
    sourceImageReference?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetSourceImageReference>;
    /**
     * A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    terminationNotification?: pulumi.Input<inputs.compute.OrchestratedVirtualMachineScaleSetTerminationNotification>;
    zoneBalance?: pulumi.Input<boolean>;
    /**
     * A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
     */
    zones?: pulumi.Input<string>;
}
