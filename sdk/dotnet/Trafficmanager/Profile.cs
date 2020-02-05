// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.TrafficManager
{
    /// <summary>
    /// Manages a Traffic Manager Profile to which multiple endpoints can be attached.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/traffic_manager_profile_legacy.html.markdown.
    /// </summary>
    public partial class Profile : Pulumi.CustomResource
    {
        /// <summary>
        /// This block specifies the DNS configuration of the
        /// Profile, it supports the fields documented below.
        /// </summary>
        [Output("dnsConfig")]
        public Output<Outputs.ProfileDnsConfig> DnsConfig { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the created Profile.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// This block specifies the Endpoint monitoring
        /// configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Output("monitorConfig")]
        public Output<Outputs.ProfileMonitorConfig> MonitorConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the virtual network. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the profile, can be set to either
        /// `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Output("profileStatus")]
        public Output<string> ProfileStatus { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to
        /// create the virtual network.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the algorithm used to route traffic, possible values are:
        /// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
        /// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
        /// - `Performance` - Traffic is routed via the User's closest Endpoint
        /// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
        /// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
        /// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
        /// </summary>
        [Output("trafficRoutingMethod")]
        public Output<string> TrafficRoutingMethod { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs args, CustomResourceOptions? options = null)
            : base("azure:trafficmanager/profile:Profile", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("azure:trafficmanager/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This block specifies the DNS configuration of the
        /// Profile, it supports the fields documented below.
        /// </summary>
        [Input("dnsConfig", required: true)]
        public Input<Inputs.ProfileDnsConfigArgs> DnsConfig { get; set; } = null!;

        /// <summary>
        /// This block specifies the Endpoint monitoring
        /// configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Input("monitorConfig", required: true)]
        public Input<Inputs.ProfileMonitorConfigArgs> MonitorConfig { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the profile, can be set to either
        /// `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("profileStatus")]
        public Input<string>? ProfileStatus { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the virtual network.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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
        /// Specifies the algorithm used to route traffic, possible values are:
        /// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
        /// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
        /// - `Performance` - Traffic is routed via the User's closest Endpoint
        /// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
        /// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
        /// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
        /// </summary>
        [Input("trafficRoutingMethod", required: true)]
        public Input<string> TrafficRoutingMethod { get; set; } = null!;

        public ProfileArgs()
        {
        }
    }

    public sealed class ProfileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This block specifies the DNS configuration of the
        /// Profile, it supports the fields documented below.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.ProfileDnsConfigGetArgs>? DnsConfig { get; set; }

        /// <summary>
        /// The FQDN of the created Profile.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// This block specifies the Endpoint monitoring
        /// configuration for the Profile, it supports the fields documented below.
        /// </summary>
        [Input("monitorConfig")]
        public Input<Inputs.ProfileMonitorConfigGetArgs>? MonitorConfig { get; set; }

        /// <summary>
        /// The name of the virtual network. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the profile, can be set to either
        /// `Enabled` or `Disabled`. Defaults to `Enabled`.
        /// </summary>
        [Input("profileStatus")]
        public Input<string>? ProfileStatus { get; set; }

        /// <summary>
        /// The name of the resource group in which to
        /// create the virtual network.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

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
        /// Specifies the algorithm used to route traffic, possible values are:
        /// - `Geographic` - Traffic is routed based on Geographic regions specified in the Endpoint.
        /// - `MultiValue`- All healthy Endpoints are returned.  MultiValue routing method works only if all the endpoints of type ‘External’ and are specified as IPv4 or IPv6 addresses.
        /// - `Performance` - Traffic is routed via the User's closest Endpoint
        /// - `Priority` - Traffic is routed to the Endpoint with the lowest `priority` value.
        /// - `Subnet` - Traffic is routed based on a mapping of sets of end-user IP address ranges to a specific Endpoint within a Traffic Manager profile.
        /// - `Weighted` - Traffic is spread across Endpoints proportional to their `weight` value.
        /// </summary>
        [Input("trafficRoutingMethod")]
        public Input<string>? TrafficRoutingMethod { get; set; }

        public ProfileState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ProfileDnsConfigArgs : Pulumi.ResourceArgs
    {
        [Input("relativeName", required: true)]
        public Input<string> RelativeName { get; set; } = null!;

        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        public ProfileDnsConfigArgs()
        {
        }
    }

    public sealed class ProfileDnsConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("relativeName", required: true)]
        public Input<string> RelativeName { get; set; } = null!;

        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        public ProfileDnsConfigGetArgs()
        {
        }
    }

    public sealed class ProfileMonitorConfigArgs : Pulumi.ResourceArgs
    {
        [Input("expectedStatusCodeRanges")]
        private InputList<string>? _expectedStatusCodeRanges;
        public InputList<string> ExpectedStatusCodeRanges
        {
            get => _expectedStatusCodeRanges ?? (_expectedStatusCodeRanges = new InputList<string>());
            set => _expectedStatusCodeRanges = value;
        }

        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        [Input("toleratedNumberOfFailures")]
        public Input<int>? ToleratedNumberOfFailures { get; set; }

        public ProfileMonitorConfigArgs()
        {
        }
    }

    public sealed class ProfileMonitorConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("expectedStatusCodeRanges")]
        private InputList<string>? _expectedStatusCodeRanges;
        public InputList<string> ExpectedStatusCodeRanges
        {
            get => _expectedStatusCodeRanges ?? (_expectedStatusCodeRanges = new InputList<string>());
            set => _expectedStatusCodeRanges = value;
        }

        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        [Input("toleratedNumberOfFailures")]
        public Input<int>? ToleratedNumberOfFailures { get; set; }

        public ProfileMonitorConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ProfileDnsConfig
    {
        public readonly string RelativeName;
        public readonly int Ttl;

        [OutputConstructor]
        private ProfileDnsConfig(
            string relativeName,
            int ttl)
        {
            RelativeName = relativeName;
            Ttl = ttl;
        }
    }

    [OutputType]
    public sealed class ProfileMonitorConfig
    {
        public readonly ImmutableArray<string> ExpectedStatusCodeRanges;
        public readonly int? IntervalInSeconds;
        public readonly string? Path;
        public readonly int Port;
        public readonly string Protocol;
        public readonly int? TimeoutInSeconds;
        public readonly int? ToleratedNumberOfFailures;

        [OutputConstructor]
        private ProfileMonitorConfig(
            ImmutableArray<string> expectedStatusCodeRanges,
            int? intervalInSeconds,
            string? path,
            int port,
            string protocol,
            int? timeoutInSeconds,
            int? toleratedNumberOfFailures)
        {
            ExpectedStatusCodeRanges = expectedStatusCodeRanges;
            IntervalInSeconds = intervalInSeconds;
            Path = path;
            Port = port;
            Protocol = protocol;
            TimeoutInSeconds = timeoutInSeconds;
            ToleratedNumberOfFailures = toleratedNumberOfFailures;
        }
    }
    }
}
