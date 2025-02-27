// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetNetworkSecurityGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Network Security Group.
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
        ///         var example = Output.Create(Azure.Network.GetNetworkSecurityGroup.InvokeAsync(new Azure.Network.GetNetworkSecurityGroupArgs
        ///         {
        ///             Name = "example",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///         this.Location = example.Apply(example =&gt; example.Location);
        ///     }
        /// 
        ///     [Output("location")]
        ///     public Output&lt;string&gt; Location { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkSecurityGroupResult> InvokeAsync(GetNetworkSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkSecurityGroupResult>("azure:network/getNetworkSecurityGroup:getNetworkSecurityGroup", args ?? new GetNetworkSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Network Security Group.
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
        ///         var example = Output.Create(Azure.Network.GetNetworkSecurityGroup.InvokeAsync(new Azure.Network.GetNetworkSecurityGroupArgs
        ///         {
        ///             Name = "example",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///         this.Location = example.Apply(example =&gt; example.Location);
        ///     }
        /// 
        ///     [Output("location")]
        ///     public Output&lt;string&gt; Location { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkSecurityGroupResult> Invoke(GetNetworkSecurityGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkSecurityGroupResult>("azure:network/getNetworkSecurityGroup:getNetworkSecurityGroup", args ?? new GetNetworkSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Network Security Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Network Security Group exists
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkSecurityGroupArgs()
        {
        }
    }

    public sealed class GetNetworkSecurityGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Network Security Group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Network Security Group exists
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetNetworkSecurityGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkSecurityGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the security rule.
        /// </summary>
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// One or more `security_rule` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkSecurityGroupSecurityRuleResult> SecurityRules;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetNetworkSecurityGroupResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableArray<Outputs.GetNetworkSecurityGroupSecurityRuleResult> securityRules,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SecurityRules = securityRules;
            Tags = tags;
        }
    }
}
