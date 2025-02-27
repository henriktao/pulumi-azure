// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./availabilitySet";
export * from "./bastionHost";
export * from "./configurationPolicyAssignment";
export * from "./dataDiskAttachment";
export * from "./dedicatedHost";
export * from "./dedicatedHostGroup";
export * from "./diskAccess";
export * from "./diskEncryptionSet";
export * from "./diskPool";
export * from "./diskPoolIscsiTarget";
export * from "./diskPoolIscsiTargetLun";
export * from "./diskPoolManagedDiskAttachment";
export * from "./extension";
export * from "./getAvailabilitySet";
export * from "./getDedicatedHost";
export * from "./getDedicatedHostGroup";
export * from "./getDiskAccess";
export * from "./getDiskEncryptionSet";
export * from "./getImage";
export * from "./getImages";
export * from "./getManagedDisk";
export * from "./getPlatformImage";
export * from "./getSharedImage";
export * from "./getSharedImageGallery";
export * from "./getSharedImageVersion";
export * from "./getSharedImageVersions";
export * from "./getSnapshot";
export * from "./getSshPublicKey";
export * from "./getVirtualMachine";
export * from "./getVirtualMachineScaleSet";
export * from "./image";
export * from "./linuxVirtualMachine";
export * from "./linuxVirtualMachineScaleSet";
export * from "./managedDisk";
export * from "./orchestratedVirtualMachineScaleSet";
export * from "./scaleSet";
export * from "./sharedImage";
export * from "./sharedImageGallery";
export * from "./sharedImageVersion";
export * from "./snapshot";
export * from "./sshPublicKey";
export * from "./virtualMachine";
export * from "./virtualMachineScaleSetExtension";
export * from "./windowsVirtualMachine";
export * from "./windowsVirtualMachineScaleSet";

// Import resources to register:
import { AvailabilitySet } from "./availabilitySet";
import { BastionHost } from "./bastionHost";
import { ConfigurationPolicyAssignment } from "./configurationPolicyAssignment";
import { DataDiskAttachment } from "./dataDiskAttachment";
import { DedicatedHost } from "./dedicatedHost";
import { DedicatedHostGroup } from "./dedicatedHostGroup";
import { DiskAccess } from "./diskAccess";
import { DiskEncryptionSet } from "./diskEncryptionSet";
import { DiskPool } from "./diskPool";
import { DiskPoolIscsiTarget } from "./diskPoolIscsiTarget";
import { DiskPoolIscsiTargetLun } from "./diskPoolIscsiTargetLun";
import { DiskPoolManagedDiskAttachment } from "./diskPoolManagedDiskAttachment";
import { Extension } from "./extension";
import { Image } from "./image";
import { LinuxVirtualMachine } from "./linuxVirtualMachine";
import { LinuxVirtualMachineScaleSet } from "./linuxVirtualMachineScaleSet";
import { ManagedDisk } from "./managedDisk";
import { OrchestratedVirtualMachineScaleSet } from "./orchestratedVirtualMachineScaleSet";
import { ScaleSet } from "./scaleSet";
import { SharedImage } from "./sharedImage";
import { SharedImageGallery } from "./sharedImageGallery";
import { SharedImageVersion } from "./sharedImageVersion";
import { Snapshot } from "./snapshot";
import { SshPublicKey } from "./sshPublicKey";
import { VirtualMachine } from "./virtualMachine";
import { VirtualMachineScaleSetExtension } from "./virtualMachineScaleSetExtension";
import { WindowsVirtualMachine } from "./windowsVirtualMachine";
import { WindowsVirtualMachineScaleSet } from "./windowsVirtualMachineScaleSet";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:compute/availabilitySet:AvailabilitySet":
                return new AvailabilitySet(name, <any>undefined, { urn })
            case "azure:compute/bastionHost:BastionHost":
                return new BastionHost(name, <any>undefined, { urn })
            case "azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment":
                return new ConfigurationPolicyAssignment(name, <any>undefined, { urn })
            case "azure:compute/dataDiskAttachment:DataDiskAttachment":
                return new DataDiskAttachment(name, <any>undefined, { urn })
            case "azure:compute/dedicatedHost:DedicatedHost":
                return new DedicatedHost(name, <any>undefined, { urn })
            case "azure:compute/dedicatedHostGroup:DedicatedHostGroup":
                return new DedicatedHostGroup(name, <any>undefined, { urn })
            case "azure:compute/diskAccess:DiskAccess":
                return new DiskAccess(name, <any>undefined, { urn })
            case "azure:compute/diskEncryptionSet:DiskEncryptionSet":
                return new DiskEncryptionSet(name, <any>undefined, { urn })
            case "azure:compute/diskPool:DiskPool":
                return new DiskPool(name, <any>undefined, { urn })
            case "azure:compute/diskPoolIscsiTarget:DiskPoolIscsiTarget":
                return new DiskPoolIscsiTarget(name, <any>undefined, { urn })
            case "azure:compute/diskPoolIscsiTargetLun:DiskPoolIscsiTargetLun":
                return new DiskPoolIscsiTargetLun(name, <any>undefined, { urn })
            case "azure:compute/diskPoolManagedDiskAttachment:DiskPoolManagedDiskAttachment":
                return new DiskPoolManagedDiskAttachment(name, <any>undefined, { urn })
            case "azure:compute/extension:Extension":
                return new Extension(name, <any>undefined, { urn })
            case "azure:compute/image:Image":
                return new Image(name, <any>undefined, { urn })
            case "azure:compute/linuxVirtualMachine:LinuxVirtualMachine":
                return new LinuxVirtualMachine(name, <any>undefined, { urn })
            case "azure:compute/linuxVirtualMachineScaleSet:LinuxVirtualMachineScaleSet":
                return new LinuxVirtualMachineScaleSet(name, <any>undefined, { urn })
            case "azure:compute/managedDisk:ManagedDisk":
                return new ManagedDisk(name, <any>undefined, { urn })
            case "azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet":
                return new OrchestratedVirtualMachineScaleSet(name, <any>undefined, { urn })
            case "azure:compute/scaleSet:ScaleSet":
                return new ScaleSet(name, <any>undefined, { urn })
            case "azure:compute/sharedImage:SharedImage":
                return new SharedImage(name, <any>undefined, { urn })
            case "azure:compute/sharedImageGallery:SharedImageGallery":
                return new SharedImageGallery(name, <any>undefined, { urn })
            case "azure:compute/sharedImageVersion:SharedImageVersion":
                return new SharedImageVersion(name, <any>undefined, { urn })
            case "azure:compute/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "azure:compute/sshPublicKey:SshPublicKey":
                return new SshPublicKey(name, <any>undefined, { urn })
            case "azure:compute/virtualMachine:VirtualMachine":
                return new VirtualMachine(name, <any>undefined, { urn })
            case "azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension":
                return new VirtualMachineScaleSetExtension(name, <any>undefined, { urn })
            case "azure:compute/windowsVirtualMachine:WindowsVirtualMachine":
                return new WindowsVirtualMachine(name, <any>undefined, { urn })
            case "azure:compute/windowsVirtualMachineScaleSet:WindowsVirtualMachineScaleSet":
                return new WindowsVirtualMachineScaleSet(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "compute/availabilitySet", _module)
pulumi.runtime.registerResourceModule("azure", "compute/bastionHost", _module)
pulumi.runtime.registerResourceModule("azure", "compute/configurationPolicyAssignment", _module)
pulumi.runtime.registerResourceModule("azure", "compute/dataDiskAttachment", _module)
pulumi.runtime.registerResourceModule("azure", "compute/dedicatedHost", _module)
pulumi.runtime.registerResourceModule("azure", "compute/dedicatedHostGroup", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskAccess", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskEncryptionSet", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskPool", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskPoolIscsiTarget", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskPoolIscsiTargetLun", _module)
pulumi.runtime.registerResourceModule("azure", "compute/diskPoolManagedDiskAttachment", _module)
pulumi.runtime.registerResourceModule("azure", "compute/extension", _module)
pulumi.runtime.registerResourceModule("azure", "compute/image", _module)
pulumi.runtime.registerResourceModule("azure", "compute/linuxVirtualMachine", _module)
pulumi.runtime.registerResourceModule("azure", "compute/linuxVirtualMachineScaleSet", _module)
pulumi.runtime.registerResourceModule("azure", "compute/managedDisk", _module)
pulumi.runtime.registerResourceModule("azure", "compute/orchestratedVirtualMachineScaleSet", _module)
pulumi.runtime.registerResourceModule("azure", "compute/scaleSet", _module)
pulumi.runtime.registerResourceModule("azure", "compute/sharedImage", _module)
pulumi.runtime.registerResourceModule("azure", "compute/sharedImageGallery", _module)
pulumi.runtime.registerResourceModule("azure", "compute/sharedImageVersion", _module)
pulumi.runtime.registerResourceModule("azure", "compute/snapshot", _module)
pulumi.runtime.registerResourceModule("azure", "compute/sshPublicKey", _module)
pulumi.runtime.registerResourceModule("azure", "compute/virtualMachine", _module)
pulumi.runtime.registerResourceModule("azure", "compute/virtualMachineScaleSetExtension", _module)
pulumi.runtime.registerResourceModule("azure", "compute/windowsVirtualMachine", _module)
pulumi.runtime.registerResourceModule("azure", "compute/windowsVirtualMachineScaleSet", _module)
