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
    /// Manages the Assignment of an API Management API Tag to an API.
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
    ///         var exampleService = Azure.ApiManagement.GetService.Invoke(new Azure.ApiManagement.GetServiceInvokeArgs
    ///         {
    ///             Name = "example-apim",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleApi = new Azure.ApiManagement.Api("exampleApi", new Azure.ApiManagement.ApiArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApiManagementName = azurerm_api_management.Example.Name,
    ///             Revision = "1",
    ///         });
    ///         var exampleTag = new Azure.ApiManagement.Tag("exampleTag", new Azure.ApiManagement.TagArgs
    ///         {
    ///             ApiManagementId = exampleService.Apply(exampleService =&gt; exampleService.Id),
    ///         });
    ///         var exampleApiTag = new Azure.ApiManagement.ApiTag("exampleApiTag", new Azure.ApiManagement.ApiTagArgs
    ///         {
    ///             ApiId = exampleApi.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management API Tags can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/apiTag:ApiTag example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/apiTag:ApiTag")]
    public partial class ApiTag : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ApiTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiTag(string name, ApiTagArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/apiTag:ApiTag", name, args ?? new ApiTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiTag(string name, Input<string> id, ApiTagState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/apiTag:ApiTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiTag Get(string name, Input<string> id, ApiTagState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiTag(name, id, state, options);
        }
    }

    public sealed class ApiTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApiTagArgs()
        {
        }
    }

    public sealed class ApiTagState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApiTagState()
        {
        }
    }
}
