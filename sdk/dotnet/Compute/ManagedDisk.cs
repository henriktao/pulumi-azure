// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages a managed disk.
    /// 
    /// ## Example Usage
    /// ### With Create Empty
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleManagedDisk = new Azure.Compute.ManagedDisk("exampleManagedDisk", new Azure.Compute.ManagedDiskArgs
    ///         {
    ///             Location = "West US 2",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StorageAccountType = "Standard_LRS",
    ///             CreateOption = "Empty",
    ///             DiskSizeGb = 1,
    ///             Tags = 
    ///             {
    ///                 { "environment", "staging" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With Create Copy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Azure.Core.ResourceGroup("example", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var source = new Azure.Compute.ManagedDisk("source", new Azure.Compute.ManagedDiskArgs
    ///         {
    ///             Location = "West US 2",
    ///             ResourceGroupName = example.Name,
    ///             StorageAccountType = "Standard_LRS",
    ///             CreateOption = "Empty",
    ///             DiskSizeGb = 1,
    ///             Tags = 
    ///             {
    ///                 { "environment", "staging" },
    ///             },
    ///         });
    ///         var copy = new Azure.Compute.ManagedDisk("copy", new Azure.Compute.ManagedDiskArgs
    ///         {
    ///             Location = "West US 2",
    ///             ResourceGroupName = example.Name,
    ///             StorageAccountType = "Standard_LRS",
    ///             CreateOption = "Copy",
    ///             SourceResourceId = source.Id,
    ///             DiskSizeGb = 1,
    ///             Tags = 
    ///             {
    ///                 { "environment", "staging" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Managed Disks can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/managedDisk:ManagedDisk example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/managedDisk:ManagedDisk")]
    public partial class ManagedDisk : Pulumi.CustomResource
    {
        /// <summary>
        /// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `source_uri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `source_resource_id`), `FromImage` (Copy a Platform Image, specified with `image_reference_id`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `source_resource_id`).
        /// </summary>
        [Output("createOption")]
        public Output<string> CreateOption { get; private set; } = null!;

        /// <summary>
        /// The ID of the disk access resource for using private endpoints on disks.
        /// </summary>
        [Output("diskAccessId")]
        public Output<string?> DiskAccessId { get; private set; } = null!;

        /// <summary>
        /// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        /// </summary>
        [Output("diskEncryptionSetId")]
        public Output<string?> DiskEncryptionSetId { get; private set; } = null!;

        /// <summary>
        /// The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Output("diskIopsReadOnly")]
        public Output<int> DiskIopsReadOnly { get; private set; } = null!;

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Output("diskIopsReadWrite")]
        public Output<int> DiskIopsReadWrite { get; private set; } = null!;

        /// <summary>
        /// The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
        /// </summary>
        [Output("diskMbpsReadOnly")]
        public Output<int> DiskMbpsReadOnly { get; private set; } = null!;

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        /// </summary>
        [Output("diskMbpsReadWrite")]
        public Output<int> DiskMbpsReadWrite { get; private set; } = null!;

        /// <summary>
        /// Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        /// </summary>
        [Output("diskSizeGb")]
        public Output<int> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// A `encryption_settings` block as defined below.
        /// </summary>
        [Output("encryptionSettings")]
        public Output<Outputs.ManagedDiskEncryptionSettings?> EncryptionSettings { get; private set; } = null!;

        /// <summary>
        /// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        /// </summary>
        [Output("imageReferenceId")]
        public Output<string?> ImageReferenceId { get; private set; } = null!;

        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("logicalSectorSize")]
        public Output<int> LogicalSectorSize { get; private set; } = null!;

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Output("maxShares")]
        public Output<int> MaxShares { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
        /// </summary>
        [Output("networkAccessPolicy")]
        public Output<string?> NetworkAccessPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Output("onDemandBurstingEnabled")]
        public Output<bool?> OnDemandBurstingEnabled { get; private set; } = null!;

        /// <summary>
        /// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Managed Disk should exist.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        /// </summary>
        [Output("sourceResourceId")]
        public Output<string?> SourceResourceId { get; private set; } = null!;

        /// <summary>
        /// URI to a valid VHD file to be used when `create_option` is `Import`.
        /// </summary>
        [Output("sourceUri")]
        public Output<string> SourceUri { get; private set; } = null!;

        /// <summary>
        /// The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string?> StorageAccountId { get; private set; } = null!;

        /// <summary>
        /// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        /// </summary>
        [Output("storageAccountType")]
        public Output<string> StorageAccountType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;

        /// <summary>
        /// Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Output("trustedLaunchEnabled")]
        public Output<bool?> TrustedLaunchEnabled { get; private set; } = null!;

        /// <summary>
        /// A collection containing the availability zone to allocate the Managed Disk in.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedDisk(string name, ManagedDiskArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/managedDisk:ManagedDisk", name, args ?? new ManagedDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedDisk(string name, Input<string> id, ManagedDiskState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/managedDisk:ManagedDisk", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedDisk Get(string name, Input<string> id, ManagedDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedDisk(name, id, state, options);
        }
    }

    public sealed class ManagedDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `source_uri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `source_resource_id`), `FromImage` (Copy a Platform Image, specified with `image_reference_id`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `source_resource_id`).
        /// </summary>
        [Input("createOption", required: true)]
        public Input<string> CreateOption { get; set; } = null!;

        /// <summary>
        /// The ID of the disk access resource for using private endpoints on disks.
        /// </summary>
        [Input("diskAccessId")]
        public Input<string>? DiskAccessId { get; set; }

        /// <summary>
        /// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIopsReadOnly")]
        public Input<int>? DiskIopsReadOnly { get; set; }

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIopsReadWrite")]
        public Input<int>? DiskIopsReadWrite { get; set; }

        /// <summary>
        /// The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
        /// </summary>
        [Input("diskMbpsReadOnly")]
        public Input<int>? DiskMbpsReadOnly { get; set; }

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        /// </summary>
        [Input("diskMbpsReadWrite")]
        public Input<int>? DiskMbpsReadWrite { get; set; }

        /// <summary>
        /// Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// A `encryption_settings` block as defined below.
        /// </summary>
        [Input("encryptionSettings")]
        public Input<Inputs.ManagedDiskEncryptionSettingsArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        /// </summary>
        [Input("imageReferenceId")]
        public Input<string>? ImageReferenceId { get; set; }

        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicalSectorSize")]
        public Input<int>? LogicalSectorSize { get; set; }

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Input("maxShares")]
        public Input<int>? MaxShares { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
        /// </summary>
        [Input("networkAccessPolicy")]
        public Input<string>? NetworkAccessPolicy { get; set; }

        /// <summary>
        /// Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Input("onDemandBurstingEnabled")]
        public Input<bool>? OnDemandBurstingEnabled { get; set; }

        /// <summary>
        /// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Disk should exist.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        /// <summary>
        /// URI to a valid VHD file to be used when `create_option` is `Import`.
        /// </summary>
        [Input("sourceUri")]
        public Input<string>? SourceUri { get; set; }

        /// <summary>
        /// The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        /// </summary>
        [Input("storageAccountType", required: true)]
        public Input<string> StorageAccountType { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Input("trustedLaunchEnabled")]
        public Input<bool>? TrustedLaunchEnabled { get; set; }

        /// <summary>
        /// A collection containing the availability zone to allocate the Managed Disk in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public ManagedDiskArgs()
        {
        }
    }

    public sealed class ManagedDiskState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method to use when creating the managed disk. Changing this forces a new resource to be created. Possible values include `Import` (Import a VHD file in to the managed disk (VHD specified with `source_uri`), `Empty` (Create an empty managed disk), `Copy` (Copy an existing managed disk or snapshot, specified with `source_resource_id`), `FromImage` (Copy a Platform Image, specified with `image_reference_id`), `Restore` (Set by Azure Backup or Site Recovery on a restored disk, specified with `source_resource_id`).
        /// </summary>
        [Input("createOption")]
        public Input<string>? CreateOption { get; set; }

        /// <summary>
        /// The ID of the disk access resource for using private endpoints on disks.
        /// </summary>
        [Input("diskAccessId")]
        public Input<string>? DiskAccessId { get; set; }

        /// <summary>
        /// The ID of a Disk Encryption Set which should be used to encrypt this Managed Disk.
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// The number of IOPS allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIopsReadOnly")]
        public Input<int>? DiskIopsReadOnly { get; set; }

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIopsReadWrite")]
        public Input<int>? DiskIopsReadWrite { get; set; }

        /// <summary>
        /// The bandwidth allowed across all VMs mounting the shared disk as read-only; only settable for UltraSSD disks with shared disk enabled. MBps means millions of bytes per second.
        /// </summary>
        [Input("diskMbpsReadOnly")]
        public Input<int>? DiskMbpsReadOnly { get; set; }

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second.
        /// </summary>
        [Input("diskMbpsReadWrite")]
        public Input<int>? DiskMbpsReadWrite { get; set; }

        /// <summary>
        /// Specifies the size of the managed disk to create in gigabytes. If `create_option` is `Copy` or `FromImage`, then the value must be equal to or greater than the source's size. The size can only be increased.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// A `encryption_settings` block as defined below.
        /// </summary>
        [Input("encryptionSettings")]
        public Input<Inputs.ManagedDiskEncryptionSettingsGetArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// ID of an existing platform/marketplace disk image to copy when `create_option` is `FromImage`.
        /// </summary>
        [Input("imageReferenceId")]
        public Input<string>? ImageReferenceId { get; set; }

        /// <summary>
        /// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Logical Sector Size. Possible values are: `512` and `4096`. Defaults to `4096`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("logicalSectorSize")]
        public Input<int>? LogicalSectorSize { get; set; }

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Input("maxShares")]
        public Input<int>? MaxShares { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Disk. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy for accessing the disk via network. Allowed values are `AllowAll`, `AllowPrivate`, and `DenyAll`.
        /// </summary>
        [Input("networkAccessPolicy")]
        public Input<string>? NetworkAccessPolicy { get; set; }

        /// <summary>
        /// Specifies if On-Demand Bursting is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Input("onDemandBurstingEnabled")]
        public Input<bool>? OnDemandBurstingEnabled { get; set; }

        /// <summary>
        /// Specify a value when the source of an `Import` or `Copy` operation targets a source that contains an operating system. Valid values are `Linux` or `Windows`.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Managed Disk should exist.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of an existing Managed Disk to copy `create_option` is `Copy` or the recovery point to restore when `create_option` is `Restore`
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        /// <summary>
        /// URI to a valid VHD file to be used when `create_option` is `Import`.
        /// </summary>
        [Input("sourceUri")]
        public Input<string>? SourceUri { get; set; }

        /// <summary>
        /// The ID of the Storage Account where the `source_uri` is located. Required when `create_option` is set to `Import`.  Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        /// <summary>
        /// The type of storage to use for the managed disk. Possible values are `Standard_LRS`, `StandardSSD_ZRS`, `Premium_LRS`, `Premium_ZRS`, `StandardSSD_LRS` or `UltraSSD_LRS`.
        /// </summary>
        [Input("storageAccountType")]
        public Input<string>? StorageAccountType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The disk performance tier to use. Possible values are documented [here](https://docs.microsoft.com/en-us/azure/virtual-machines/disks-change-performance). This feature is currently supported only for premium SSDs.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// Specifies if Trusted Launch is enabled for the Managed Disk. Defaults to `false`.
        /// </summary>
        [Input("trustedLaunchEnabled")]
        public Input<bool>? TrustedLaunchEnabled { get; set; }

        /// <summary>
        /// A collection containing the availability zone to allocate the Managed Disk in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public ManagedDiskState()
        {
        }
    }
}
