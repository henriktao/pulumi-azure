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
    /// Manages a Virtual Machine.
    /// 
    /// ## Disclaimers
    /// 
    /// &gt; **Note:** The `azure.compute.VirtualMachine` resource has been superseded by the `azure.compute.LinuxVirtualMachine` and `azure.compute.WindowsVirtualMachine` resources. The existing `azure.compute.VirtualMachine` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `azure.compute.LinuxVirtualMachine` and `azure.compute.WindowsVirtualMachine` resources.
    /// 
    /// &gt; **Note:** Data Disks can be attached either directly on the `azure.compute.VirtualMachine` resource, or using the `azure.compute.DataDiskAttachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
    /// 
    /// ## Example Usage
    /// ### From An Azure Platform Image)
    /// 
    /// This example provisions a Virtual Machine with Managed Disks.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var prefix = config.Get("prefix") ?? "tfvmex";
    ///         var mainResourceGroup = new Azure.Core.ResourceGroup("mainResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var mainVirtualNetwork = new Azure.Network.VirtualNetwork("mainVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///         });
    ///         var @internal = new Azure.Network.Subnet("internal", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             VirtualNetworkName = mainVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///         });
    ///         var mainNetworkInterface = new Azure.Network.NetworkInterface("mainNetworkInterface", new Azure.Network.NetworkInterfaceArgs
    ///         {
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
    ///                 {
    ///                     Name = "testconfiguration1",
    ///                     SubnetId = @internal.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                 },
    ///             },
    ///         });
    ///         var mainVirtualMachine = new Azure.Compute.VirtualMachine("mainVirtualMachine", new Azure.Compute.VirtualMachineArgs
    ///         {
    ///             Location = mainResourceGroup.Location,
    ///             ResourceGroupName = mainResourceGroup.Name,
    ///             NetworkInterfaceIds = 
    ///             {
    ///                 mainNetworkInterface.Id,
    ///             },
    ///             VmSize = "Standard_DS1_v2",
    ///             StorageImageReference = new Azure.Compute.Inputs.VirtualMachineStorageImageReferenceArgs
    ///             {
    ///                 Publisher = "Canonical",
    ///                 Offer = "UbuntuServer",
    ///                 Sku = "16.04-LTS",
    ///                 Version = "latest",
    ///             },
    ///             StorageOsDisk = new Azure.Compute.Inputs.VirtualMachineStorageOsDiskArgs
    ///             {
    ///                 Name = "myosdisk1",
    ///                 Caching = "ReadWrite",
    ///                 CreateOption = "FromImage",
    ///                 ManagedDiskType = "Standard_LRS",
    ///             },
    ///             OsProfile = new Azure.Compute.Inputs.VirtualMachineOsProfileArgs
    ///             {
    ///                 ComputerName = "hostname",
    ///                 AdminUsername = "testadmin",
    ///                 AdminPassword = "Password1234!",
    ///             },
    ///             OsProfileLinuxConfig = new Azure.Compute.Inputs.VirtualMachineOsProfileLinuxConfigArgs
    ///             {
    ///                 DisablePasswordAuthentication = false,
    ///             },
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
    /// Virtual Machines can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/virtualMachine:VirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/virtualMachine:VirtualMachine")]
    public partial class VirtualMachine : Pulumi.CustomResource
    {
        /// <summary>
        /// An `additional_capabilities` block as defined below.
        /// </summary>
        [Output("additionalCapabilities")]
        public Output<Outputs.VirtualMachineAdditionalCapabilities?> AdditionalCapabilities { get; private set; } = null!;

        /// <summary>
        /// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("availabilitySetId")]
        public Output<string> AvailabilitySetId { get; private set; } = null!;

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Output("bootDiagnostics")]
        public Output<Outputs.VirtualMachineBootDiagnostics?> BootDiagnostics { get; private set; } = null!;

        /// <summary>
        /// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Output("deleteDataDisksOnTermination")]
        public Output<bool?> DeleteDataDisksOnTermination { get; private set; } = null!;

        /// <summary>
        /// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Output("deleteOsDiskOnTermination")]
        public Output<bool?> DeleteOsDiskOnTermination { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.VirtualMachineIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        /// </summary>
        [Output("licenseType")]
        public Output<string> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of Network Interface ID's which should be associated with the Virtual Machine.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<ImmutableArray<string>> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// An `os_profile` block as defined below. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        /// </summary>
        [Output("osProfile")]
        public Output<Outputs.VirtualMachineOsProfile?> OsProfile { get; private set; } = null!;

        /// <summary>
        /// An `os_profile_linux_config` block as defined below.
        /// </summary>
        [Output("osProfileLinuxConfig")]
        public Output<Outputs.VirtualMachineOsProfileLinuxConfig?> OsProfileLinuxConfig { get; private set; } = null!;

        /// <summary>
        /// One or more `os_profile_secrets` blocks.
        /// </summary>
        [Output("osProfileSecrets")]
        public Output<ImmutableArray<Outputs.VirtualMachineOsProfileSecret>> OsProfileSecrets { get; private set; } = null!;

        /// <summary>
        /// An `os_profile_windows_config` block as defined below.
        /// </summary>
        [Output("osProfileWindowsConfig")]
        public Output<Outputs.VirtualMachineOsProfileWindowsConfig?> OsProfileWindowsConfig { get; private set; } = null!;

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.VirtualMachinePlan?> Plan { get; private set; } = null!;

        /// <summary>
        /// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        /// </summary>
        [Output("primaryNetworkInterfaceId")]
        public Output<string?> PrimaryNetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Output("proximityPlacementGroupId")]
        public Output<string?> ProximityPlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_data_disk` blocks.
        /// </summary>
        [Output("storageDataDisks")]
        public Output<ImmutableArray<Outputs.VirtualMachineStorageDataDisk>> StorageDataDisks { get; private set; } = null!;

        /// <summary>
        /// A `storage_image_reference` block as defined below.
        /// </summary>
        [Output("storageImageReference")]
        public Output<Outputs.VirtualMachineStorageImageReference> StorageImageReference { get; private set; } = null!;

        /// <summary>
        /// A `storage_os_disk` block as defined below.
        /// </summary>
        [Output("storageOsDisk")]
        public Output<Outputs.VirtualMachineStorageOsDisk> StorageOsDisk { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Machine.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the [size of the Virtual Machine](https://docs.microsoft.com/azure/virtual-machines/sizes-general). See also [Azure VM Naming Conventions](https://docs.microsoft.com/azure/virtual-machines/vm-naming-conventions).
        /// </summary>
        [Output("vmSize")]
        public Output<string> VmSize { get; private set; } = null!;

        /// <summary>
        /// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/virtualMachine:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/virtualMachine:VirtualMachine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, state, options);
        }
    }

    public sealed class VirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `additional_capabilities` block as defined below.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.VirtualMachineAdditionalCapabilitiesArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("availabilitySetId")]
        public Input<string>? AvailabilitySetId { get; set; }

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Input("bootDiagnostics")]
        public Input<Inputs.VirtualMachineBootDiagnosticsArgs>? BootDiagnostics { get; set; }

        /// <summary>
        /// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Input("deleteDataDisksOnTermination")]
        public Input<bool>? DeleteDataDisksOnTermination { get; set; }

        /// <summary>
        /// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Input("deleteOsDiskOnTermination")]
        public Input<bool>? DeleteOsDiskOnTermination { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.VirtualMachineIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceIds", required: true)]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// A list of Network Interface ID's which should be associated with the Virtual Machine.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// An `os_profile` block as defined below. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        /// </summary>
        [Input("osProfile")]
        public Input<Inputs.VirtualMachineOsProfileArgs>? OsProfile { get; set; }

        /// <summary>
        /// An `os_profile_linux_config` block as defined below.
        /// </summary>
        [Input("osProfileLinuxConfig")]
        public Input<Inputs.VirtualMachineOsProfileLinuxConfigArgs>? OsProfileLinuxConfig { get; set; }

        [Input("osProfileSecrets")]
        private InputList<Inputs.VirtualMachineOsProfileSecretArgs>? _osProfileSecrets;

        /// <summary>
        /// One or more `os_profile_secrets` blocks.
        /// </summary>
        public InputList<Inputs.VirtualMachineOsProfileSecretArgs> OsProfileSecrets
        {
            get => _osProfileSecrets ?? (_osProfileSecrets = new InputList<Inputs.VirtualMachineOsProfileSecretArgs>());
            set => _osProfileSecrets = value;
        }

        /// <summary>
        /// An `os_profile_windows_config` block as defined below.
        /// </summary>
        [Input("osProfileWindowsConfig")]
        public Input<Inputs.VirtualMachineOsProfileWindowsConfigArgs>? OsProfileWindowsConfig { get; set; }

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.VirtualMachinePlanArgs>? Plan { get; set; }

        /// <summary>
        /// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        /// </summary>
        [Input("primaryNetworkInterfaceId")]
        public Input<string>? PrimaryNetworkInterfaceId { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("storageDataDisks")]
        private InputList<Inputs.VirtualMachineStorageDataDiskArgs>? _storageDataDisks;

        /// <summary>
        /// One or more `storage_data_disk` blocks.
        /// </summary>
        public InputList<Inputs.VirtualMachineStorageDataDiskArgs> StorageDataDisks
        {
            get => _storageDataDisks ?? (_storageDataDisks = new InputList<Inputs.VirtualMachineStorageDataDiskArgs>());
            set => _storageDataDisks = value;
        }

        /// <summary>
        /// A `storage_image_reference` block as defined below.
        /// </summary>
        [Input("storageImageReference")]
        public Input<Inputs.VirtualMachineStorageImageReferenceArgs>? StorageImageReference { get; set; }

        /// <summary>
        /// A `storage_os_disk` block as defined below.
        /// </summary>
        [Input("storageOsDisk", required: true)]
        public Input<Inputs.VirtualMachineStorageOsDiskArgs> StorageOsDisk { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Machine.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the [size of the Virtual Machine](https://docs.microsoft.com/azure/virtual-machines/sizes-general). See also [Azure VM Naming Conventions](https://docs.microsoft.com/azure/virtual-machines/vm-naming-conventions).
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        /// <summary>
        /// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public VirtualMachineArgs()
        {
        }
    }

    public sealed class VirtualMachineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `additional_capabilities` block as defined below.
        /// </summary>
        [Input("additionalCapabilities")]
        public Input<Inputs.VirtualMachineAdditionalCapabilitiesGetArgs>? AdditionalCapabilities { get; set; }

        /// <summary>
        /// The ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("availabilitySetId")]
        public Input<string>? AvailabilitySetId { get; set; }

        /// <summary>
        /// A `boot_diagnostics` block as defined below.
        /// </summary>
        [Input("bootDiagnostics")]
        public Input<Inputs.VirtualMachineBootDiagnosticsGetArgs>? BootDiagnostics { get; set; }

        /// <summary>
        /// Should the Data Disks (either the Managed Disks / VHD Blobs) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Input("deleteDataDisksOnTermination")]
        public Input<bool>? DeleteDataDisksOnTermination { get; set; }

        /// <summary>
        /// Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed? Defaults to `false`.
        /// </summary>
        [Input("deleteOsDiskOnTermination")]
        public Input<bool>? DeleteOsDiskOnTermination { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.VirtualMachineIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the BYOL Type for this Virtual Machine. This is only applicable to Windows Virtual Machines. Possible values are `Windows_Client` and `Windows_Server`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the Virtual Machine exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Virtual Machine. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// A list of Network Interface ID's which should be associated with the Virtual Machine.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// An `os_profile` block as defined below. Required when `create_option` in the `storage_os_disk` block is set to `FromImage`.
        /// </summary>
        [Input("osProfile")]
        public Input<Inputs.VirtualMachineOsProfileGetArgs>? OsProfile { get; set; }

        /// <summary>
        /// An `os_profile_linux_config` block as defined below.
        /// </summary>
        [Input("osProfileLinuxConfig")]
        public Input<Inputs.VirtualMachineOsProfileLinuxConfigGetArgs>? OsProfileLinuxConfig { get; set; }

        [Input("osProfileSecrets")]
        private InputList<Inputs.VirtualMachineOsProfileSecretGetArgs>? _osProfileSecrets;

        /// <summary>
        /// One or more `os_profile_secrets` blocks.
        /// </summary>
        public InputList<Inputs.VirtualMachineOsProfileSecretGetArgs> OsProfileSecrets
        {
            get => _osProfileSecrets ?? (_osProfileSecrets = new InputList<Inputs.VirtualMachineOsProfileSecretGetArgs>());
            set => _osProfileSecrets = value;
        }

        /// <summary>
        /// An `os_profile_windows_config` block as defined below.
        /// </summary>
        [Input("osProfileWindowsConfig")]
        public Input<Inputs.VirtualMachineOsProfileWindowsConfigGetArgs>? OsProfileWindowsConfig { get; set; }

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.VirtualMachinePlanGetArgs>? Plan { get; set; }

        /// <summary>
        /// The ID of the Network Interface (which must be attached to the Virtual Machine) which should be the Primary Network Interface for this Virtual Machine.
        /// </summary>
        [Input("primaryNetworkInterfaceId")]
        public Input<string>? PrimaryNetworkInterfaceId { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which the Virtual Machine should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("storageDataDisks")]
        private InputList<Inputs.VirtualMachineStorageDataDiskGetArgs>? _storageDataDisks;

        /// <summary>
        /// One or more `storage_data_disk` blocks.
        /// </summary>
        public InputList<Inputs.VirtualMachineStorageDataDiskGetArgs> StorageDataDisks
        {
            get => _storageDataDisks ?? (_storageDataDisks = new InputList<Inputs.VirtualMachineStorageDataDiskGetArgs>());
            set => _storageDataDisks = value;
        }

        /// <summary>
        /// A `storage_image_reference` block as defined below.
        /// </summary>
        [Input("storageImageReference")]
        public Input<Inputs.VirtualMachineStorageImageReferenceGetArgs>? StorageImageReference { get; set; }

        /// <summary>
        /// A `storage_os_disk` block as defined below.
        /// </summary>
        [Input("storageOsDisk")]
        public Input<Inputs.VirtualMachineStorageOsDiskGetArgs>? StorageOsDisk { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Virtual Machine.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the [size of the Virtual Machine](https://docs.microsoft.com/azure/virtual-machines/sizes-general). See also [Azure VM Naming Conventions](https://docs.microsoft.com/azure/virtual-machines/vm-naming-conventions).
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        /// <summary>
        /// A list of a single item of the Availability Zone which the Virtual Machine should be allocated in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public VirtualMachineState()
        {
        }
    }
}
