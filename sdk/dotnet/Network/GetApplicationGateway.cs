// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Network
{
    public static class GetApplicationGateway
    {
        /// <summary>
        /// Use this data source to access information about an existing Application Gateway.
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
        ///         var example = Output.Create(Azure.Network.GetApplicationGateway.InvokeAsync(new Azure.Network.GetApplicationGatewayArgs
        ///         {
        ///             Name = "existing-app-gateway",
        ///             ResourceGroupName = "existing-resources",
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
        public static Task<GetApplicationGatewayResult> InvokeAsync(GetApplicationGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationGatewayResult>("azure:network/getApplicationGateway:getApplicationGateway", args ?? new GetApplicationGatewayArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Application Gateway.
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
        ///         var example = Output.Create(Azure.Network.GetApplicationGateway.InvokeAsync(new Azure.Network.GetApplicationGatewayArgs
        ///         {
        ///             Name = "existing-app-gateway",
        ///             ResourceGroupName = "existing-resources",
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
        public static Output<GetApplicationGatewayResult> Invoke(GetApplicationGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationGatewayResult>("azure:network/getApplicationGateway:getApplicationGateway", args ?? new GetApplicationGatewayInvokeArgs(), options.WithVersion());
    }


    public sealed class GetApplicationGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Application Gateway.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Application Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationGatewayArgs()
        {
        }
    }

    public sealed class GetApplicationGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Application Gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Application Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApplicationGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationGatewayResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationGatewayIdentityResult> Identities;
        /// <summary>
        /// The Azure Region where the Application Gateway exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the Application Gateway.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetApplicationGatewayResult(
            string id,

            ImmutableArray<Outputs.GetApplicationGatewayIdentityResult> identities,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Identities = identities;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
