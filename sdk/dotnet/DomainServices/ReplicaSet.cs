// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DomainServices
{
    /// <summary>
    /// Manages a Replica Set for an Active Directory Domain Service.
    /// 
    /// ## Import
    /// 
    /// Domain Services can be imported using the resource ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:domainservices/replicaSet:ReplicaSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1
    /// ```
    /// </summary>
    [AzureResourceType("azure:domainservices/replicaSet:ReplicaSet")]
    public partial class ReplicaSet : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        /// </summary>
        [Output("domainControllerIpAddresses")]
        public Output<ImmutableArray<string>> DomainControllerIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("domainServiceId")]
        public Output<string> DomainServiceId { get; private set; } = null!;

        /// <summary>
        /// The publicly routable IP address for the domain controllers in this Replica Set.
        /// </summary>
        [Output("externalAccessIpAddress")]
        public Output<string> ExternalAccessIpAddress { get; private set; } = null!;

        /// <summary>
        /// The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The current service status for the replica set.
        /// </summary>
        [Output("serviceStatus")]
        public Output<string> ServiceStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet in which to place this Replica Set.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicaSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicaSet(string name, ReplicaSetArgs args, CustomResourceOptions? options = null)
            : base("azure:domainservices/replicaSet:ReplicaSet", name, args ?? new ReplicaSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicaSet(string name, Input<string> id, ReplicaSetState? state = null, CustomResourceOptions? options = null)
            : base("azure:domainservices/replicaSet:ReplicaSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicaSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicaSet Get(string name, Input<string> id, ReplicaSetState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicaSet(name, id, state, options);
        }
    }

    public sealed class ReplicaSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("domainServiceId", required: true)]
        public Input<string> DomainServiceId { get; set; } = null!;

        /// <summary>
        /// The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the subnet in which to place this Replica Set.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ReplicaSetArgs()
        {
        }
    }

    public sealed class ReplicaSetState : Pulumi.ResourceArgs
    {
        [Input("domainControllerIpAddresses")]
        private InputList<string>? _domainControllerIpAddresses;

        /// <summary>
        /// A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        /// </summary>
        public InputList<string> DomainControllerIpAddresses
        {
            get => _domainControllerIpAddresses ?? (_domainControllerIpAddresses = new InputList<string>());
            set => _domainControllerIpAddresses = value;
        }

        /// <summary>
        /// The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("domainServiceId")]
        public Input<string>? DomainServiceId { get; set; }

        /// <summary>
        /// The publicly routable IP address for the domain controllers in this Replica Set.
        /// </summary>
        [Input("externalAccessIpAddress")]
        public Input<string>? ExternalAccessIpAddress { get; set; }

        /// <summary>
        /// The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The current service status for the replica set.
        /// </summary>
        [Input("serviceStatus")]
        public Input<string>? ServiceStatus { get; set; }

        /// <summary>
        /// The ID of the subnet in which to place this Replica Set.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public ReplicaSetState()
        {
        }
    }
}
