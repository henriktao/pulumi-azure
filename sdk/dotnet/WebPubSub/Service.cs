// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.WebPubSub
{
    /// <summary>
    /// Manages an Azure Web Pubsub Service.
    /// 
    /// ## Import
    /// 
    /// Web Pubsub services can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:webpubsub/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
    /// ```
    /// </summary>
    [AzureResourceType("azure:webpubsub/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable AAD auth? Defaults to `true`.
        /// </summary>
        [Output("aadAuthEnabled")]
        public Output<bool?> AadAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
        /// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
        /// </summary>
        [Output("capacity")]
        public Output<int?> Capacity { get; private set; } = null!;

        [Output("externalIp")]
        public Output<string> ExternalIp { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the Web Pubsub service.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// A `live_trace` block as defined below.
        /// </summary>
        [Output("liveTrace")]
        public Output<Outputs.ServiceLiveTrace?> LiveTrace { get; private set; } = null!;

        /// <summary>
        /// Whether to enable local auth? Defaults to `true`.
        /// </summary>
        [Output("localAuthEnabled")]
        public Output<bool?> LocalAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Web Pubsub service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The primary access key for the Web Pubsub service.
        /// </summary>
        [Output("primaryAccessKey")]
        public Output<string> PrimaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The primary connection string for the Web Pubsub service.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the Web Pubsub service which is designed for browser/client use.
        /// </summary>
        [Output("publicPort")]
        public Output<int> PublicPort { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Web Pubsub service. Changing
        /// this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The secondary access key for the Web Pubsub service.
        /// </summary>
        [Output("secondaryAccessKey")]
        public Output<string> SecondaryAccessKey { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string for the Web Pubsub service.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the Web Pubsub service which is designed for customer server side use.
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether to request client certificate during TLS handshake? Defaults
        /// to `false`.
        /// </summary>
        [Output("tlsClientCertEnabled")]
        public Output<bool?> TlsClientCertEnabled { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:webpubsub/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:webpubsub/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// Whether to enable AAD auth? Defaults to `true`.
        /// </summary>
        [Input("aadAuthEnabled")]
        public Input<bool>? AadAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
        /// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// A `live_trace` block as defined below.
        /// </summary>
        [Input("liveTrace")]
        public Input<Inputs.ServiceLiveTraceArgs>? LiveTrace { get; set; }

        /// <summary>
        /// Whether to enable local auth? Defaults to `true`.
        /// </summary>
        [Input("localAuthEnabled")]
        public Input<bool>? LocalAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Web Pubsub service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Web Pubsub service. Changing
        /// this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

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

        /// <summary>
        /// Whether to request client certificate during TLS handshake? Defaults
        /// to `false`.
        /// </summary>
        [Input("tlsClientCertEnabled")]
        public Input<bool>? TlsClientCertEnabled { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable AAD auth? Defaults to `true`.
        /// </summary>
        [Input("aadAuthEnabled")]
        public Input<bool>? AadAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the number of units associated with this Web Pubsub resource. Valid values are:
        /// Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        [Input("externalIp")]
        public Input<string>? ExternalIp { get; set; }

        /// <summary>
        /// The FQDN of the Web Pubsub service.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// A `live_trace` block as defined below.
        /// </summary>
        [Input("liveTrace")]
        public Input<Inputs.ServiceLiveTraceGetArgs>? LiveTrace { get; set; }

        /// <summary>
        /// Whether to enable local auth? Defaults to `true`.
        /// </summary>
        [Input("localAuthEnabled")]
        public Input<bool>? LocalAuthEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Web Pubsub service exists. Changing this
        /// forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Web Pubsub service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary access key for the Web Pubsub service.
        /// </summary>
        [Input("primaryAccessKey")]
        public Input<string>? PrimaryAccessKey { get; set; }

        /// <summary>
        /// The primary connection string for the Web Pubsub service.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The publicly accessible port of the Web Pubsub service which is designed for browser/client use.
        /// </summary>
        [Input("publicPort")]
        public Input<int>? PublicPort { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Web Pubsub service. Changing
        /// this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The secondary access key for the Web Pubsub service.
        /// </summary>
        [Input("secondaryAccessKey")]
        public Input<string>? SecondaryAccessKey { get; set; }

        /// <summary>
        /// The secondary connection string for the Web Pubsub service.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The publicly accessible port of the Web Pubsub service which is designed for customer server side use.
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        /// <summary>
        /// Whether to request client certificate during TLS handshake? Defaults
        /// to `false`.
        /// </summary>
        [Input("tlsClientCertEnabled")]
        public Input<bool>? TlsClientCertEnabled { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServiceState()
        {
        }
    }
}
