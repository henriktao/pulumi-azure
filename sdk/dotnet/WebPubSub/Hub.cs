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
    /// Manages the hub settings for a Web Pubsub.
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
    ///             Location = "east us",
    ///         });
    ///         var testUserAssignedIdentity = new Azure.Authorization.UserAssignedIdentity("testUserAssignedIdentity", new Azure.Authorization.UserAssignedIdentityArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var exampleService = new Azure.WebPubSub.Service("exampleService", new Azure.WebPubSub.ServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard_S1",
    ///             Capacity = 1,
    ///         });
    ///         var testHub = new Azure.WebPubSub.Hub("testHub", new Azure.WebPubSub.HubArgs
    ///         {
    ///             WebPubsubId = azurerm_web_pubsub.Exmaple.Id,
    ///             EventHandlers = 
    ///             {
    ///                 new Azure.WebPubSub.Inputs.HubEventHandlerArgs
    ///                 {
    ///                     UrlTemplate = "https://test.com/api/{hub}/{event}",
    ///                     UserEventPattern = "*",
    ///                     SystemEvents = 
    ///                     {
    ///                         "connect",
    ///                         "connected",
    ///                     },
    ///                 },
    ///                 new Azure.WebPubSub.Inputs.HubEventHandlerArgs
    ///                 {
    ///                     UrlTemplate = "https://test.com/api/{hub}/{event}",
    ///                     UserEventPattern = "event1, event2",
    ///                     SystemEvents = 
    ///                     {
    ///                         "connected",
    ///                     },
    ///                     Auth = new Azure.WebPubSub.Inputs.HubEventHandlerAuthArgs
    ///                     {
    ///                         ManagedIdentityId = testUserAssignedIdentity.Id,
    ///                     },
    ///                 },
    ///             },
    ///             AnonymousConnectionsEnabled = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 azurerm_web_pubsub.Test,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Web Pubsub Hub can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:webpubsub/hub:Hub example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubsub/webpubsub1/hubs/webpubsubhub1
    /// ```
    /// </summary>
    [AzureResourceType("azure:webpubsub/hub:Hub")]
    public partial class Hub : Pulumi.CustomResource
    {
        /// <summary>
        /// Is anonymous connections are allowed for this hub? Defaults to `false`.
        /// Possible values are `true`, `false`.
        /// </summary>
        [Output("anonymousConnectionsEnabled")]
        public Output<bool?> AnonymousConnectionsEnabled { get; private set; } = null!;

        /// <summary>
        /// An `event_handler` block as defined below.
        /// </summary>
        [Output("eventHandlers")]
        public Output<ImmutableArray<Outputs.HubEventHandler>> EventHandlers { get; private set; } = null!;

        /// <summary>
        /// The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
        /// </summary>
        [Output("webPubsubId")]
        public Output<string> WebPubsubId { get; private set; } = null!;


        /// <summary>
        /// Create a Hub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hub(string name, HubArgs args, CustomResourceOptions? options = null)
            : base("azure:webpubsub/hub:Hub", name, args ?? new HubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hub(string name, Input<string> id, HubState? state = null, CustomResourceOptions? options = null)
            : base("azure:webpubsub/hub:Hub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hub Get(string name, Input<string> id, HubState? state = null, CustomResourceOptions? options = null)
        {
            return new Hub(name, id, state, options);
        }
    }

    public sealed class HubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is anonymous connections are allowed for this hub? Defaults to `false`.
        /// Possible values are `true`, `false`.
        /// </summary>
        [Input("anonymousConnectionsEnabled")]
        public Input<bool>? AnonymousConnectionsEnabled { get; set; }

        [Input("eventHandlers", required: true)]
        private InputList<Inputs.HubEventHandlerArgs>? _eventHandlers;

        /// <summary>
        /// An `event_handler` block as defined below.
        /// </summary>
        public InputList<Inputs.HubEventHandlerArgs> EventHandlers
        {
            get => _eventHandlers ?? (_eventHandlers = new InputList<Inputs.HubEventHandlerArgs>());
            set => _eventHandlers = value;
        }

        /// <summary>
        /// The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("webPubsubId", required: true)]
        public Input<string> WebPubsubId { get; set; } = null!;

        public HubArgs()
        {
        }
    }

    public sealed class HubState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is anonymous connections are allowed for this hub? Defaults to `false`.
        /// Possible values are `true`, `false`.
        /// </summary>
        [Input("anonymousConnectionsEnabled")]
        public Input<bool>? AnonymousConnectionsEnabled { get; set; }

        [Input("eventHandlers")]
        private InputList<Inputs.HubEventHandlerGetArgs>? _eventHandlers;

        /// <summary>
        /// An `event_handler` block as defined below.
        /// </summary>
        public InputList<Inputs.HubEventHandlerGetArgs> EventHandlers
        {
            get => _eventHandlers ?? (_eventHandlers = new InputList<Inputs.HubEventHandlerGetArgs>());
            set => _eventHandlers = value;
        }

        /// <summary>
        /// The name of the Web Pubsub hub service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the id of the Web Pubsub. Changing this forces a new resource to be created.
        /// </summary>
        [Input("webPubsubId")]
        public Input<string>? WebPubsubId { get; set; }

        public HubState()
        {
        }
    }
}
