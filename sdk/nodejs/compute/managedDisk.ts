// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a managed disk.
 *
 * ## Example Usage
 * ### With Create Empty
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleManagedDisk = new azure.compute.ManagedDisk("exampleManagedDisk", {
 *     location: "West US 2",
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccountType: "Standard_LRS",
 *     createOption: "Empty",
 *     diskSizeGb: "1",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * ```
 * ### With Create Copy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceGroup("example", {location: "West Europe"});
 * const source = new azure.compute.ManagedDisk("source", {
 *     location: "West US 2",
 *     resourceGroupName: example.name,
 *     storageAccountType: "Standard_LRS",
 *     createOption: "Empty",
 *     diskSizeGb: "1",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const copy = new azure.compute.ManagedDisk("copy", {
 *     location: "West US 2",
 *     resourceGroupName: example.name,
 *     storageAccountType: "Standard_LRS",
 *     createOption: "Copy",
 *     sourceResourceId: source.id,
 *     diskSizeGb: "1",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Managed Disks can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/managedDisk:ManagedDisk example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1
 * ```
 */
export class ManagedDisk extends pulumi.CustomResource {
    /**
     * Get an existing ManagedDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedDiskState, opts?: pulumi.CustomResourceOptions): ManagedDisk {
        return new ManagedDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/managedDisk:ManagedDisk';

    /**
     * Returns true if the given object is an instance of ManagedDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedDisk.__pulumiType;
    }

    /**
     * The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
     */
    public readonly createOption!: pulumi.Output<string>;
    /**
     * The ID of the disk access resource for using private endpoints on disks.
     */
    public readonly diskAccessId!: pulumi.Output<string | undefined>;
    /**
     * The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
     */
    public readonly diskEncryptionSetId!: pulumi.Output<string | undefined>;
    /**
     * The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
     */
    public readonly diskIopsReadOnly!: pulumi.Output<number>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    public readonly diskIopsReadWrite!: pulumi.Output<number>;
    /**
     * The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
     */
    public readonly diskMbpsReadOnly!: pulumi.Output<number>;
    /**
     * The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
     */
    public readonly diskMbpsReadWrite!: pulumi.Output<number>;
    /**
     * Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
     */
    public readonly diskSizeGb!: pulumi.Output<number>;
    /**
     * A `encryptionSettings` block as defined below.
     */
    public readonly encryptionSettings!: pulumi.Output<outputs.compute.ManagedDiskEncryptionSettings | undefined>;
    /**
     * The HyperV Generation of the Disk when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Possible values are `V1` and `V2`. Changing this forces a new resource to be created.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
     */
    public readonly imageReferenceId!: pulumi.Output<string | undefined>;
    /**
     * Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
     */
    public readonly logicalSectorSize!: pulumi.Output<number>;
    /**
     * The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
     */
    public readonly maxShares!: pulumi.Output<number>;
    /**
     * Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
     */
    public readonly networkAccessPolicy!: pulumi.Output<string | undefined>;
    /**
     * Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
     */
    public readonly onDemandBurstingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * Whether it is allowed to access the disk via public network. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Resource Group where the Managed Disk should exist.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
     */
    public readonly sourceResourceId!: pulumi.Output<string | undefined>;
    /**
     * URI to a valid VHD file to be used when `createOption` is `Import`.
     */
    public readonly sourceUri!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;
    /**
     * The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
     */
    public readonly storageAccountType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
     */
    public readonly trustedLaunchEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A collection containing the availability zone to allocate the Managed Disk in.
     */
    public readonly zones!: pulumi.Output<string | undefined>;

    /**
     * Create a ManagedDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedDiskArgs | ManagedDiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedDiskState | undefined;
            inputs["createOption"] = state ? state.createOption : undefined;
            inputs["diskAccessId"] = state ? state.diskAccessId : undefined;
            inputs["diskEncryptionSetId"] = state ? state.diskEncryptionSetId : undefined;
            inputs["diskIopsReadOnly"] = state ? state.diskIopsReadOnly : undefined;
            inputs["diskIopsReadWrite"] = state ? state.diskIopsReadWrite : undefined;
            inputs["diskMbpsReadOnly"] = state ? state.diskMbpsReadOnly : undefined;
            inputs["diskMbpsReadWrite"] = state ? state.diskMbpsReadWrite : undefined;
            inputs["diskSizeGb"] = state ? state.diskSizeGb : undefined;
            inputs["encryptionSettings"] = state ? state.encryptionSettings : undefined;
            inputs["hyperVGeneration"] = state ? state.hyperVGeneration : undefined;
            inputs["imageReferenceId"] = state ? state.imageReferenceId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["logicalSectorSize"] = state ? state.logicalSectorSize : undefined;
            inputs["maxShares"] = state ? state.maxShares : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkAccessPolicy"] = state ? state.networkAccessPolicy : undefined;
            inputs["onDemandBurstingEnabled"] = state ? state.onDemandBurstingEnabled : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sourceResourceId"] = state ? state.sourceResourceId : undefined;
            inputs["sourceUri"] = state ? state.sourceUri : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            inputs["storageAccountType"] = state ? state.storageAccountType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["trustedLaunchEnabled"] = state ? state.trustedLaunchEnabled : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as ManagedDiskArgs | undefined;
            if ((!args || args.createOption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'createOption'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.storageAccountType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountType'");
            }
            inputs["createOption"] = args ? args.createOption : undefined;
            inputs["diskAccessId"] = args ? args.diskAccessId : undefined;
            inputs["diskEncryptionSetId"] = args ? args.diskEncryptionSetId : undefined;
            inputs["diskIopsReadOnly"] = args ? args.diskIopsReadOnly : undefined;
            inputs["diskIopsReadWrite"] = args ? args.diskIopsReadWrite : undefined;
            inputs["diskMbpsReadOnly"] = args ? args.diskMbpsReadOnly : undefined;
            inputs["diskMbpsReadWrite"] = args ? args.diskMbpsReadWrite : undefined;
            inputs["diskSizeGb"] = args ? args.diskSizeGb : undefined;
            inputs["encryptionSettings"] = args ? args.encryptionSettings : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["imageReferenceId"] = args ? args.imageReferenceId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logicalSectorSize"] = args ? args.logicalSectorSize : undefined;
            inputs["maxShares"] = args ? args.maxShares : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkAccessPolicy"] = args ? args.networkAccessPolicy : undefined;
            inputs["onDemandBurstingEnabled"] = args ? args.onDemandBurstingEnabled : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceResourceId"] = args ? args.sourceResourceId : undefined;
            inputs["sourceUri"] = args ? args.sourceUri : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["storageAccountType"] = args ? args.storageAccountType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["trustedLaunchEnabled"] = args ? args.trustedLaunchEnabled : undefined;
            inputs["zones"] = args ? args.zones : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ManagedDisk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedDisk resources.
 */
export interface ManagedDiskState {
    /**
     * The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
     */
    createOption?: pulumi.Input<string>;
    /**
     * The ID of the disk access resource for using private endpoints on disks.
     */
    diskAccessId?: pulumi.Input<string>;
    /**
     * The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
     */
    diskEncryptionSetId?: pulumi.Input<string>;
    /**
     * The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
     */
    diskIopsReadOnly?: pulumi.Input<number>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    diskIopsReadWrite?: pulumi.Input<number>;
    /**
     * The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
     */
    diskMbpsReadOnly?: pulumi.Input<number>;
    /**
     * The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
     */
    diskMbpsReadWrite?: pulumi.Input<number>;
    /**
     * Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
     */
    diskSizeGb?: pulumi.Input<number>;
    /**
     * A `encryptionSettings` block as defined below.
     */
    encryptionSettings?: pulumi.Input<inputs.compute.ManagedDiskEncryptionSettings>;
    /**
     * The HyperV Generation of the Disk when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Possible values are `V1` and `V2`. Changing this forces a new resource to be created.
     */
    hyperVGeneration?: pulumi.Input<string>;
    /**
     * ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
     */
    imageReferenceId?: pulumi.Input<string>;
    /**
     * Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
     */
    logicalSectorSize?: pulumi.Input<number>;
    /**
     * The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
     */
    maxShares?: pulumi.Input<number>;
    /**
     * Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
     */
    networkAccessPolicy?: pulumi.Input<string>;
    /**
     * Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
     */
    onDemandBurstingEnabled?: pulumi.Input<boolean>;
    /**
     * Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Whether it is allowed to access the disk via public network. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the Managed Disk should exist.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
     */
    sourceResourceId?: pulumi.Input<string>;
    /**
     * URI to a valid VHD file to be used when `createOption` is `Import`.
     */
    sourceUri?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
     */
    storageAccountType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
     */
    tier?: pulumi.Input<string>;
    /**
     * Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
     */
    trustedLaunchEnabled?: pulumi.Input<boolean>;
    /**
     * A collection containing the availability zone to allocate the Managed Disk in.
     */
    zones?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedDisk resource.
 */
export interface ManagedDiskArgs {
    /**
     * The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `sourceUri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `sourceResourceId`), `FromImage` (Copy a Platform Image, specified with `imageReferenceId`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `sourceResourceId`).
     */
    createOption: pulumi.Input<string>;
    /**
     * The ID of the disk access resource for using private endpoints on disks.
     */
    diskAccessId?: pulumi.Input<string>;
    /**
     * The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
     */
    diskEncryptionSetId?: pulumi.Input<string>;
    /**
     * The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
     */
    diskIopsReadOnly?: pulumi.Input<number>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    diskIopsReadWrite?: pulumi.Input<number>;
    /**
     * The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
     */
    diskMbpsReadOnly?: pulumi.Input<number>;
    /**
     * The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
     */
    diskMbpsReadWrite?: pulumi.Input<number>;
    /**
     * Specifies the size of the managed disk to create in gigabytes. If `createOption` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
     */
    diskSizeGb?: pulumi.Input<number>;
    /**
     * A `encryptionSettings` block as defined below.
     */
    encryptionSettings?: pulumi.Input<inputs.compute.ManagedDiskEncryptionSettings>;
    /**
     * The HyperV Generation of the Disk when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Possible values are `V1` and `V2`. Changing this forces a new resource to be created.
     */
    hyperVGeneration?: pulumi.Input<string>;
    /**
     * ID of an existing platform/marketplace disk image to copy when `createOption` is `FromImage`.
     */
    imageReferenceId?: pulumi.Input<string>;
    /**
     * Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
     */
    logicalSectorSize?: pulumi.Input<number>;
    /**
     * The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
     */
    maxShares?: pulumi.Input<number>;
    /**
     * Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
     */
    networkAccessPolicy?: pulumi.Input<string>;
    /**
     * Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
     */
    onDemandBurstingEnabled?: pulumi.Input<boolean>;
    /**
     * Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Whether it is allowed to access the disk via public network. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the Managed Disk should exist.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The ID of an existing Managed Disk to copy `createOption` is `Copy` or the recovery point to restore when `createOption` is `Restore`
     */
    sourceResourceId?: pulumi.Input<string>;
    /**
     * URI to a valid VHD file to be used when `createOption` is `Import`.
     */
    sourceUri?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where the `sourceUri` is located. Required when `createOption` is set to `Import`.  Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
     */
    storageAccountType: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
     */
    tier?: pulumi.Input<string>;
    /**
     * Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
     */
    trustedLaunchEnabled?: pulumi.Input<boolean>;
    /**
     * A collection containing the availability zone to allocate the Managed Disk in.
     */
    zones?: pulumi.Input<string>;
}
