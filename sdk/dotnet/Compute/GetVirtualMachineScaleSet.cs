// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static class GetVirtualMachineScaleSet
    {
        /// <summary>
        /// Use this data source to access information about an existing Virtual Machine Scale Set.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Compute.GetVirtualMachineScaleSet.InvokeAsync(new Azure.Compute.GetVirtualMachineScaleSetArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualMachineScaleSetResult> InvokeAsync(GetVirtualMachineScaleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineScaleSetResult>("azure:compute/getVirtualMachineScaleSet:getVirtualMachineScaleSet", args ?? new GetVirtualMachineScaleSetArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Virtual Machine Scale Set.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Compute.GetVirtualMachineScaleSet.InvokeAsync(new Azure.Compute.GetVirtualMachineScaleSetArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVirtualMachineScaleSetResult> Invoke(GetVirtualMachineScaleSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVirtualMachineScaleSetResult>("azure:compute/getVirtualMachineScaleSet:getVirtualMachineScaleSet", args ?? new GetVirtualMachineScaleSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineScaleSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Virtual Machine Scale Set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Virtual Machine Scale Set exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineScaleSetArgs()
        {
        }
    }

    public sealed class GetVirtualMachineScaleSetInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Virtual Machine Scale Set.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Virtual Machine Scale Set exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineScaleSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineScaleSetResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineScaleSetIdentityResult> Identities;
        public readonly string Location;
        /// <summary>
        /// The name of the public ip address configuration
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of `network_interface` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceResult> NetworkInterfaces;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetVirtualMachineScaleSetResult(
            string id,

            ImmutableArray<Outputs.GetVirtualMachineScaleSetIdentityResult> identities,

            string location,

            string name,

            ImmutableArray<Outputs.GetVirtualMachineScaleSetNetworkInterfaceResult> networkInterfaces,

            string resourceGroupName)
        {
            Id = id;
            Identities = identities;
            Location = location;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            ResourceGroupName = resourceGroupName;
        }
    }
}
