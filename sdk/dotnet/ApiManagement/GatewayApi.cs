// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages a API Management Gateway API.
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
    ///         var exampleService = Output.Create(Azure.ApiManagement.GetService.InvokeAsync(new Azure.ApiManagement.GetServiceArgs
    ///         {
    ///             Name = "example-api",
    ///             ResourceGroupName = "example-resources",
    ///         }));
    ///         var exampleApi = Output.Tuple(exampleService, exampleService).Apply(values =&gt;
    ///         {
    ///             var exampleService = values.Item1;
    ///             var exampleService1 = values.Item2;
    ///             return Output.Create(Azure.ApiManagement.GetApi.InvokeAsync(new Azure.ApiManagement.GetApiArgs
    ///             {
    ///                 Name = "search-api",
    ///                 ApiManagementName = exampleService.Name,
    ///                 ResourceGroupName = exampleService1.ResourceGroupName,
    ///                 Revision = "2",
    ///             }));
    ///         });
    ///         var exampleGateway = Output.Create(Azure.ApiManagement.GetGateway.InvokeAsync(new Azure.ApiManagement.GetGatewayArgs
    ///         {
    ///             GatewayId = "my-gateway",
    ///             ApiManagementName = exampleService.Apply(exampleService =&gt; exampleService.Name),
    ///             ResourceGroupName = exampleService.Apply(exampleService =&gt; exampleService.ResourceGroupName),
    ///         }));
    ///         var exampleGatewayApi = new Azure.ApiManagement.GatewayApi("exampleGatewayApi", new Azure.ApiManagement.GatewayApiArgs
    ///         {
    ///             GatewayId = exampleGateway.Apply(exampleGateway =&gt; exampleGateway.Id),
    ///             ApiId = exampleApi.Apply(exampleApi =&gt; exampleApi.Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management Gateway APIs can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/gatewayApi:GatewayApi example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/apis/api1
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/gatewayApi:GatewayApi")]
    public partial class GatewayApi : Pulumi.CustomResource
    {
        /// <summary>
        /// The Identifier of the API Management API within the API Management Service. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The Identifier for the API Management Gateway. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayApi(string name, GatewayApiArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/gatewayApi:GatewayApi", name, args ?? new GatewayApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayApi(string name, Input<string> id, GatewayApiState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/gatewayApi:GatewayApi", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayApi Get(string name, Input<string> id, GatewayApiState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayApi(name, id, state, options);
        }
    }

    public sealed class GatewayApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Identifier of the API Management API within the API Management Service. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The Identifier for the API Management Gateway. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        public GatewayApiArgs()
        {
        }
    }

    public sealed class GatewayApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Identifier of the API Management API within the API Management Service. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The Identifier for the API Management Gateway. Changing this forces a new API Management Gateway API to be created.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        public GatewayApiState()
        {
        }
    }
}
