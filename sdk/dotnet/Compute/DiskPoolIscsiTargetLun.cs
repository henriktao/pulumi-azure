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
    /// Manages an iSCSI Target lun.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.0.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "diskspool",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/virtualNetworks/read",
    ///                         },
    ///                         Name = "Microsoft.StoragePool/diskPools",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleDiskPool = new Azure.Compute.DiskPool("exampleDiskPool", new Azure.Compute.DiskPoolArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             SubnetId = exampleSubnet.Id,
    ///             Zones = 
    ///             {
    ///                 "1",
    ///             },
    ///             SkuName = "Basic_B1",
    ///         });
    ///         var exampleManagedDisk = new Azure.Compute.ManagedDisk("exampleManagedDisk", new Azure.Compute.ManagedDiskArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             CreateOption = "Empty",
    ///             StorageAccountType = "Premium_LRS",
    ///             DiskSizeGb = 4,
    ///             MaxShares = 2,
    ///             Zones = 
    ///             {
    ///                 "1",
    ///             },
    ///         });
    ///         var exampleServicePrincipal = Output.Create(AzureAD.GetServicePrincipal.InvokeAsync(new AzureAD.GetServicePrincipalArgs
    ///         {
    ///             DisplayName = "StoragePool Resource Provider",
    ///         }));
    ///         var roles = 
    ///         {
    ///             "Disk Pool Operator",
    ///             "Virtual Machine Contributor",
    ///         };
    ///         var exampleAssignment = new List&lt;Azure.Authorization.Assignment&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; roles.Length; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             exampleAssignment.Add(new Azure.Authorization.Assignment($"exampleAssignment-{range.Value}", new Azure.Authorization.AssignmentArgs
    ///             {
    ///                 PrincipalId = exampleServicePrincipal.Apply(exampleServicePrincipal =&gt; exampleServicePrincipal.Id),
    ///                 RoleDefinitionName = roles[range.Value],
    ///                 Scope = exampleManagedDisk.Id,
    ///             }));
    ///         }
    ///         var exampleDiskPoolManagedDiskAttachment = new Azure.Compute.DiskPoolManagedDiskAttachment("exampleDiskPoolManagedDiskAttachment", new Azure.Compute.DiskPoolManagedDiskAttachmentArgs
    ///         {
    ///             DiskPoolId = exampleDiskPool.Id,
    ///             ManagedDiskId = exampleManagedDisk.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleAssignment,
    ///             },
    ///         });
    ///         var exampleDiskPoolIscsiTarget = new Azure.Compute.DiskPoolIscsiTarget("exampleDiskPoolIscsiTarget", new Azure.Compute.DiskPoolIscsiTargetArgs
    ///         {
    ///             AclMode = "Dynamic",
    ///             DisksPoolId = exampleDiskPool.Id,
    ///             TargetIqn = "iqn.2021-11.com.microsoft:test",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleDiskPoolManagedDiskAttachment,
    ///             },
    ///         });
    ///         var exampleDiskPoolIscsiTargetLun = new Azure.Compute.DiskPoolIscsiTargetLun("exampleDiskPoolIscsiTargetLun", new Azure.Compute.DiskPoolIscsiTargetLunArgs
    ///         {
    ///             IscsiTargetId = exampleDiskPoolIscsiTarget.Id,
    ///             DiskPoolManagedDiskAttachmentId = exampleDiskPoolManagedDiskAttachment.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// iSCSI Target Luns can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/diskPoolIscsiTargetLun:DiskPoolIscsiTargetLun example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StoragePool/diskPools/diskPoolValue/iscsiTargets/iscsiTargetValue/lun|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/disks/disk1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/diskPoolIscsiTargetLun:DiskPoolIscsiTargetLun")]
    public partial class DiskPoolIscsiTargetLun : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the `azure.compute.DiskPoolManagedDiskAttachment`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Output("diskPoolManagedDiskAttachmentId")]
        public Output<string> DiskPoolManagedDiskAttachmentId { get; private set; } = null!;

        /// <summary>
        /// The ID of the iSCSI Target. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Output("iscsiTargetId")]
        public Output<string> IscsiTargetId { get; private set; } = null!;

        /// <summary>
        /// The Logical Unit Number of the iSCSI Target LUN.
        /// </summary>
        [Output("lun")]
        public Output<int> Lun { get; private set; } = null!;

        /// <summary>
        /// User defined name for iSCSI LUN. Supported characters include uppercase letters, lowercase letters, numbers, periods, underscores or hyphens. Name should end with an alphanumeric character. The length must be between `1` and `90`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a DiskPoolIscsiTargetLun resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskPoolIscsiTargetLun(string name, DiskPoolIscsiTargetLunArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/diskPoolIscsiTargetLun:DiskPoolIscsiTargetLun", name, args ?? new DiskPoolIscsiTargetLunArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DiskPoolIscsiTargetLun(string name, Input<string> id, DiskPoolIscsiTargetLunState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/diskPoolIscsiTargetLun:DiskPoolIscsiTargetLun", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskPoolIscsiTargetLun resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskPoolIscsiTargetLun Get(string name, Input<string> id, DiskPoolIscsiTargetLunState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskPoolIscsiTargetLun(name, id, state, options);
        }
    }

    public sealed class DiskPoolIscsiTargetLunArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the `azure.compute.DiskPoolManagedDiskAttachment`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("diskPoolManagedDiskAttachmentId", required: true)]
        public Input<string> DiskPoolManagedDiskAttachmentId { get; set; } = null!;

        /// <summary>
        /// The ID of the iSCSI Target. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("iscsiTargetId", required: true)]
        public Input<string> IscsiTargetId { get; set; } = null!;

        /// <summary>
        /// User defined name for iSCSI LUN. Supported characters include uppercase letters, lowercase letters, numbers, periods, underscores or hyphens. Name should end with an alphanumeric character. The length must be between `1` and `90`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DiskPoolIscsiTargetLunArgs()
        {
        }
    }

    public sealed class DiskPoolIscsiTargetLunState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the `azure.compute.DiskPoolManagedDiskAttachment`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("diskPoolManagedDiskAttachmentId")]
        public Input<string>? DiskPoolManagedDiskAttachmentId { get; set; }

        /// <summary>
        /// The ID of the iSCSI Target. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("iscsiTargetId")]
        public Input<string>? IscsiTargetId { get; set; }

        /// <summary>
        /// The Logical Unit Number of the iSCSI Target LUN.
        /// </summary>
        [Input("lun")]
        public Input<int>? Lun { get; set; }

        /// <summary>
        /// User defined name for iSCSI LUN. Supported characters include uppercase letters, lowercase letters, numbers, periods, underscores or hyphens. Name should end with an alphanumeric character. The length must be between `1` and `90`. Changing this forces a new iSCSI Target LUN to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DiskPoolIscsiTargetLunState()
        {
        }
    }
}
