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
    /// Manages an Orchestrated Virtual Machine Scale Set.
    /// 
    /// ## Disclaimers
    /// 
    /// &gt; **NOTE:** As of the **v2.86.0** (November 19, 2021) release of the provider this resource will only create Virtual Machine Scale Sets with the **Flexible** Orchestration Mode.
    /// 
    /// &gt; **NOTE:** All arguments including the administrator login and password will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// ## Example Usage
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
    ///         var exampleOrchestratedVirtualMachineScaleSet = new Azure.Compute.OrchestratedVirtualMachineScaleSet("exampleOrchestratedVirtualMachineScaleSet", new Azure.Compute.OrchestratedVirtualMachineScaleSetArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PlatformFaultDomainCount = 1,
    ///             Zones = 
    ///             {
    ///                 "1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// An Orchestrated Virtual Machine Scale Set can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/Microsoft.Compute/virtualMachineScaleSets/scaleset1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet")]
    public partial class OrchestratedVirtualMachineScaleSet : Pulumi.CustomResource
    {
        [Output("automaticInstanceRepair")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair> AutomaticInstanceRepair { get; private set; } = null!;

        [Output("bootDiagnostics")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetBootDiagnostics?> BootDiagnostics { get; private set; } = null!;

        [Output("dataDisks")]
        public Output<ImmutableArray<Outputs.OrchestratedVirtualMachineScaleSetDataDisk>> DataDisks { get; private set; } = null!;

        [Output("encryptionAtHostEnabled")]
        public Output<bool?> EncryptionAtHostEnabled { get; private set; } = null!;

        [Output("evictionPolicy")]
        public Output<string?> EvictionPolicy { get; private set; } = null!;

        [Output("extensions")]
        public Output<ImmutableArray<Outputs.OrchestratedVirtualMachineScaleSetExtension>> Extensions { get; private set; } = null!;

        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
        /// </summary>
        [Output("extensionsTimeBudget")]
        public Output<string?> ExtensionsTimeBudget { get; private set; } = null!;

        [Output("identity")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
        /// </summary>
        [Output("instances")]
        public Output<int> Instances { get; private set; } = null!;

        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("maxBidPrice")]
        public Output<double?> MaxBidPrice { get; private set; } = null!;

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.OrchestratedVirtualMachineScaleSetNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        [Output("osDisk")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetOsDisk?> OsDisk { get; private set; } = null!;

        [Output("osProfile")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetOsProfile?> OsProfile { get; private set; } = null!;

        [Output("plan")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetPlan?> Plan { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("platformFaultDomainCount")]
        public Output<int> PlatformFaultDomainCount { get; private set; } = null!;

        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("proximityPlacementGroupId")]
        public Output<string?> ProximityPlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        [Output("skuName")]
        public Output<string?> SkuName { get; private set; } = null!;

        [Output("sourceImageId")]
        public Output<string?> SourceImageId { get; private set; } = null!;

        /// <summary>
        /// A `source_image_reference` block as defined below.
        /// </summary>
        [Output("sourceImageReference")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetSourceImageReference?> SourceImageReference { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("terminationNotification")]
        public Output<Outputs.OrchestratedVirtualMachineScaleSetTerminationNotification> TerminationNotification { get; private set; } = null!;

        /// <summary>
        /// The Unique ID for the Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;

        [Output("zoneBalance")]
        public Output<bool?> ZoneBalance { get; private set; } = null!;

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a OrchestratedVirtualMachineScaleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrchestratedVirtualMachineScaleSet(string name, OrchestratedVirtualMachineScaleSetArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, args ?? new OrchestratedVirtualMachineScaleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrchestratedVirtualMachineScaleSet(string name, Input<string> id, OrchestratedVirtualMachineScaleSetState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrchestratedVirtualMachineScaleSet Get(string name, Input<string> id, OrchestratedVirtualMachineScaleSetState? state = null, CustomResourceOptions? options = null)
        {
            return new OrchestratedVirtualMachineScaleSet(name, id, state, options);
        }
    }

    public sealed class OrchestratedVirtualMachineScaleSetArgs : Pulumi.ResourceArgs
    {
        [Input("automaticInstanceRepair")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairArgs>? AutomaticInstanceRepair { get; set; }

        [Input("bootDiagnostics")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetBootDiagnosticsArgs>? BootDiagnostics { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskArgs>? _dataDisks;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskArgs>());
            set => _dataDisks = value;
        }

        [Input("encryptionAtHostEnabled")]
        public Input<bool>? EncryptionAtHostEnabled { get; set; }

        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        [Input("extensions")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionArgs>? _extensions;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionArgs> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionArgs>());
            set => _extensions = value;
        }

        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
        /// </summary>
        [Input("extensionsTimeBudget")]
        public Input<string>? ExtensionsTimeBudget { get; set; }

        [Input("identity")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
        /// </summary>
        [Input("instances")]
        public Input<int>? Instances { get; set; }

        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("maxBidPrice")]
        public Input<double>? MaxBidPrice { get; set; }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("osDisk")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetOsDiskArgs>? OsDisk { get; set; }

        [Input("osProfile")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetOsProfileArgs>? OsProfile { get; set; }

        [Input("plan")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetPlanArgs>? Plan { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount", required: true)]
        public Input<int> PlatformFaultDomainCount { get; set; } = null!;

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// A `source_image_reference` block as defined below.
        /// </summary>
        [Input("sourceImageReference")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetSourceImageReferenceArgs>? SourceImageReference { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("terminationNotification")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetTerminationNotificationArgs>? TerminationNotification { get; set; }

        [Input("zoneBalance")]
        public Input<bool>? ZoneBalance { get; set; }

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public OrchestratedVirtualMachineScaleSetArgs()
        {
        }
    }

    public sealed class OrchestratedVirtualMachineScaleSetState : Pulumi.ResourceArgs
    {
        [Input("automaticInstanceRepair")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetAutomaticInstanceRepairGetArgs>? AutomaticInstanceRepair { get; set; }

        [Input("bootDiagnostics")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetBootDiagnosticsGetArgs>? BootDiagnostics { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskGetArgs>? _dataDisks;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.OrchestratedVirtualMachineScaleSetDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        [Input("encryptionAtHostEnabled")]
        public Input<bool>? EncryptionAtHostEnabled { get; set; }

        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        [Input("extensions")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionGetArgs>? _extensions;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionGetArgs> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<Inputs.OrchestratedVirtualMachineScaleSetExtensionGetArgs>());
            set => _extensions = value;
        }

        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M).
        /// </summary>
        [Input("extensionsTimeBudget")]
        public Input<string>? ExtensionsTimeBudget { get; set; }

        [Input("identity")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The number of Virtual Machines in the Orcestrated Virtual Machine Scale Set.
        /// </summary>
        [Input("instances")]
        public Input<int>? Instances { get; set; }

        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("maxBidPrice")]
        public Input<double>? MaxBidPrice { get; set; }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceGetArgs>? _networkInterfaces;
        public InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.OrchestratedVirtualMachineScaleSetNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        [Input("osDisk")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetOsDiskGetArgs>? OsDisk { get; set; }

        [Input("osProfile")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetOsProfileGetArgs>? OsProfile { get; set; }

        [Input("plan")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetPlanGetArgs>? Plan { get; set; }

        /// <summary>
        /// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("platformFaultDomainCount")]
        public Input<int>? PlatformFaultDomainCount { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The ID of the Proximity Placement Group which the Orchestrated Virtual Machine should be assigned to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("proximityPlacementGroupId")]
        public Input<string>? ProximityPlacementGroupId { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// A `source_image_reference` block as defined below.
        /// </summary>
        [Input("sourceImageReference")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetSourceImageReferenceGetArgs>? SourceImageReference { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("terminationNotification")]
        public Input<Inputs.OrchestratedVirtualMachineScaleSetTerminationNotificationGetArgs>? TerminationNotification { get; set; }

        /// <summary>
        /// The Unique ID for the Orchestrated Virtual Machine Scale Set.
        /// </summary>
        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        [Input("zoneBalance")]
        public Input<bool>? ZoneBalance { get; set; }

        /// <summary>
        /// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public OrchestratedVirtualMachineScaleSetState()
        {
        }
    }
}
