// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SignalR
{
    /// <summary>
    /// Manages an Azure SignalR service.
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
    ///             Location = "West US",
    ///         });
    ///         var exampleService = new Azure.SignalR.Service("exampleService", new Azure.SignalR.ServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = new Azure.SignalR.Inputs.ServiceSkuArgs
    ///             {
    ///                 Name = "Free_F1",
    ///                 Capacity = 1,
    ///             },
    ///             Cors = 
    ///             {
    ///                 new Azure.SignalR.Inputs.ServiceCorArgs
    ///                 {
    ///                     AllowedOrigins = 
    ///                     {
    ///                         "http://www.example.com",
    ///                     },
    ///                 },
    ///             },
    ///             Features = 
    ///             {
    ///                 new Azure.SignalR.Inputs.ServiceFeatureArgs
    ///                 {
    ///                     Flag = "ServiceMode",
    ///                     Value = "Default",
    ///                 },
    ///             },
    ///             UpstreamEndpoints = 
    ///             {
    ///                 new Azure.SignalR.Inputs.ServiceUpstreamEndpointArgs
    ///                 {
    ///                     CategoryPatterns = 
    ///                     {
    ///                         "connections",
    ///                         "messages",
    ///                     },
    ///                     EventPatterns = 
    ///                     {
    ///                         "*",
    ///                     },
    ///                     HubPatterns = 
    ///                     {
    ///                         "hub1",
    ///                     },
    ///                     UrlTemplate = "http://foo.com",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SignalR services can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:signalr/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/SignalR/tfex-signalr
    /// ```
    /// </summary>
    [AzureResourceType("azure:signalr/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// A `cors` block as documented below.
        /// </summary>
        [Output("cors")]
        public Output<ImmutableArray<Outputs.ServiceCor>> Cors { get; private set; } = null!;

        /// <summary>
        /// A `features` block as documented below.
        /// </summary>
        [Output("features")]
        public Output<ImmutableArray<Outputs.ServiceFeature>> Features { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the SignalR service.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible IP of the SignalR service.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The primary access key for the SignalR service.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The primary connection string for the SignalR service.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for browser/client use.
        /// </summary>
        [Output("publicPort")]
        public Output<int> PublicPort { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The secondary access key for the SignalR service.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string for the SignalR service.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for customer server side use.
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// A `sku` block as documented below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ServiceSku> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        /// </summary>
        [Output("upstreamEndpoints")]
        public Output<ImmutableArray<Outputs.ServiceUpstreamEndpoint>> UpstreamEndpoints { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:signalr/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:signalr/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("cors")]
        private InputList<Inputs.ServiceCorArgs>? _cors;

        /// <summary>
        /// A `cors` block as documented below.
        /// </summary>
        public InputList<Inputs.ServiceCorArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.ServiceCorArgs>());
            set => _cors = value;
        }

        [Input("features")]
        private InputList<Inputs.ServiceFeatureArgs>? _features;

        /// <summary>
        /// A `features` block as documented below.
        /// </summary>
        public InputList<Inputs.ServiceFeatureArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.ServiceFeatureArgs>());
            set => _features = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `sku` block as documented below.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ServiceSkuArgs> Sku { get; set; } = null!;

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

        [Input("upstreamEndpoints")]
        private InputList<Inputs.ServiceUpstreamEndpointArgs>? _upstreamEndpoints;

        /// <summary>
        /// An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        /// </summary>
        public InputList<Inputs.ServiceUpstreamEndpointArgs> UpstreamEndpoints
        {
            get => _upstreamEndpoints ?? (_upstreamEndpoints = new InputList<Inputs.ServiceUpstreamEndpointArgs>());
            set => _upstreamEndpoints = value;
        }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("cors")]
        private InputList<Inputs.ServiceCorGetArgs>? _cors;

        /// <summary>
        /// A `cors` block as documented below.
        /// </summary>
        public InputList<Inputs.ServiceCorGetArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.ServiceCorGetArgs>());
            set => _cors = value;
        }

        [Input("features")]
        private InputList<Inputs.ServiceFeatureGetArgs>? _features;

        /// <summary>
        /// A `features` block as documented below.
        /// </summary>
        public InputList<Inputs.ServiceFeatureGetArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.ServiceFeatureGetArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The FQDN of the SignalR service.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The publicly accessible IP of the SignalR service.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the SignalR service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary access key for the SignalR service.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The primary connection string for the SignalR service.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for browser/client use.
        /// </summary>
        [Input("publicPort")]
        public Input<int>? PublicPort { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the SignalR service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The secondary access key for the SignalR service.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        /// <summary>
        /// The secondary connection string for the SignalR service.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for customer server side use.
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// A `sku` block as documented below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ServiceSkuGetArgs>? Sku { get; set; }

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

        [Input("upstreamEndpoints")]
        private InputList<Inputs.ServiceUpstreamEndpointGetArgs>? _upstreamEndpoints;

        /// <summary>
        /// An `upstream_endpoint` block as documented below. Using this block requires the SignalR service to be Serverless. When creating multiple blocks they will be processed in the order they are defined in.
        /// </summary>
        public InputList<Inputs.ServiceUpstreamEndpointGetArgs> UpstreamEndpoints
        {
            get => _upstreamEndpoints ?? (_upstreamEndpoints = new InputList<Inputs.ServiceUpstreamEndpointGetArgs>());
            set => _upstreamEndpoints = value;
        }

        public ServiceState()
        {
        }
    }
}
