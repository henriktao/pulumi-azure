// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages an IotHub Device Provisioning Service.
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
    ///         var exampleIotHubDps = new Azure.Iot.IotHubDps("exampleIotHubDps", new Azure.Iot.IotHubDpsArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IotHubDpsSkuArgs
    ///             {
    ///                 Name = "S1",
    ///                 Capacity = 1,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoT Device Provisioning Service can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:iot/iotHubDps:IotHubDps example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:iot/iotHubDps:IotHubDps")]
    public partial class IotHubDps : Pulumi.CustomResource
    {
        /// <summary>
        /// The allocation policy of the IoT Device Provisioning Service.
        /// </summary>
        [Output("allocationPolicy")]
        public Output<string> AllocationPolicy { get; private set; } = null!;

        /// <summary>
        /// The device endpoint of the IoT Device Provisioning Service.
        /// </summary>
        [Output("deviceProvisioningHostName")]
        public Output<string> DeviceProvisioningHostName { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the IoT Device Provisioning Service.
        /// </summary>
        [Output("idScope")]
        public Output<string> IdScope { get; private set; } = null!;

        /// <summary>
        /// A `linked_hub` block as defined below.
        /// </summary>
        [Output("linkedHubs")]
        public Output<ImmutableArray<Outputs.IotHubDpsLinkedHub>> LinkedHubs { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The service endpoint of the IoT Device Provisioning Service.
        /// </summary>
        [Output("serviceOperationsHostName")]
        public Output<string> ServiceOperationsHostName { get; private set; } = null!;

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.IotHubDpsSku> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a IotHubDps resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHubDps(string name, IotHubDpsArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/iotHubDps:IotHubDps", name, args ?? new IotHubDpsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotHubDps(string name, Input<string> id, IotHubDpsState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/iotHubDps:IotHubDps", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IotHubDps resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHubDps Get(string name, Input<string> id, IotHubDpsState? state = null, CustomResourceOptions? options = null)
        {
            return new IotHubDps(name, id, state, options);
        }
    }

    public sealed class IotHubDpsArgs : Pulumi.ResourceArgs
    {
        [Input("linkedHubs")]
        private InputList<Inputs.IotHubDpsLinkedHubArgs>? _linkedHubs;

        /// <summary>
        /// A `linked_hub` block as defined below.
        /// </summary>
        public InputList<Inputs.IotHubDpsLinkedHubArgs> LinkedHubs
        {
            get => _linkedHubs ?? (_linkedHubs = new InputList<Inputs.IotHubDpsLinkedHubArgs>());
            set => _linkedHubs = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.IotHubDpsSkuArgs> Sku { get; set; } = null!;

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

        public IotHubDpsArgs()
        {
        }
    }

    public sealed class IotHubDpsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation policy of the IoT Device Provisioning Service.
        /// </summary>
        [Input("allocationPolicy")]
        public Input<string>? AllocationPolicy { get; set; }

        /// <summary>
        /// The device endpoint of the IoT Device Provisioning Service.
        /// </summary>
        [Input("deviceProvisioningHostName")]
        public Input<string>? DeviceProvisioningHostName { get; set; }

        /// <summary>
        /// The unique identifier of the IoT Device Provisioning Service.
        /// </summary>
        [Input("idScope")]
        public Input<string>? IdScope { get; set; }

        [Input("linkedHubs")]
        private InputList<Inputs.IotHubDpsLinkedHubGetArgs>? _linkedHubs;

        /// <summary>
        /// A `linked_hub` block as defined below.
        /// </summary>
        public InputList<Inputs.IotHubDpsLinkedHubGetArgs> LinkedHubs
        {
            get => _linkedHubs ?? (_linkedHubs = new InputList<Inputs.IotHubDpsLinkedHubGetArgs>());
            set => _linkedHubs = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The service endpoint of the IoT Device Provisioning Service.
        /// </summary>
        [Input("serviceOperationsHostName")]
        public Input<string>? ServiceOperationsHostName { get; set; }

        /// <summary>
        /// A `sku` block as defined below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.IotHubDpsSkuGetArgs>? Sku { get; set; }

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

        public IotHubDpsState()
        {
        }
    }
}
