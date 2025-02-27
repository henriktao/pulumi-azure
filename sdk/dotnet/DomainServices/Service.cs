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
    /// ## Import
    /// 
    /// Domain Services can be imported using the resource ID, together with the Replica Set ID that you wish to designate as the initial replica set, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:domainservices/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/initialReplicaSetId/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureResourceType("azure:domainservices/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique ID for the managed domain deployment.
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
        /// </summary>
        [Output("filteredSyncEnabled")]
        public Output<bool?> FilteredSyncEnabled { get; private set; } = null!;

        /// <summary>
        /// An `initial_replica_set` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
        /// </summary>
        [Output("initialReplicaSet")]
        public Output<Outputs.ServiceInitialReplicaSet> InitialReplicaSet { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `notifications` block as defined below.
        /// </summary>
        [Output("notifications")]
        public Output<Outputs.ServiceNotifications> Notifications { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Azure resource ID for the domain service.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// A `secure_ldap` block as defined below.
        /// </summary>
        [Output("secureLdap")]
        public Output<Outputs.ServiceSecureLdap> SecureLdap { get; private set; } = null!;

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Output("security")]
        public Output<Outputs.ServiceSecurity> Security { get; private set; } = null!;

        /// <summary>
        /// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        [Output("syncOwner")]
        public Output<string> SyncOwner { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:domainservices/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:domainservices/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
        /// </summary>
        [Input("filteredSyncEnabled")]
        public Input<bool>? FilteredSyncEnabled { get; set; }

        /// <summary>
        /// An `initial_replica_set` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
        /// </summary>
        [Input("initialReplicaSet", required: true)]
        public Input<Inputs.ServiceInitialReplicaSetArgs> InitialReplicaSet { get; set; } = null!;

        /// <summary>
        /// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `notifications` block as defined below.
        /// </summary>
        [Input("notifications")]
        public Input<Inputs.ServiceNotificationsArgs>? Notifications { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `secure_ldap` block as defined below.
        /// </summary>
        [Input("secureLdap")]
        public Input<Inputs.ServiceSecureLdapArgs>? SecureLdap { get; set; }

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Input("security")]
        public Input<Inputs.ServiceSecurityArgs>? Security { get; set; }

        /// <summary>
        /// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique ID for the managed domain deployment.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The Active Directory domain to use. See [official documentation](https://docs.microsoft.com/en-us/azure/active-directory-domain-services/tutorial-create-instance#create-a-managed-domain) for constraints and recommendations.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Whether to enable group-based filtered sync (also called scoped synchronisation). Defaults to `false`.
        /// </summary>
        [Input("filteredSyncEnabled")]
        public Input<bool>? FilteredSyncEnabled { get; set; }

        /// <summary>
        /// An `initial_replica_set` block as defined below. The initial replica set inherits the same location as the Domain Service resource.
        /// </summary>
        [Input("initialReplicaSet")]
        public Input<Inputs.ServiceInitialReplicaSetGetArgs>? InitialReplicaSet { get; set; }

        /// <summary>
        /// The Azure location where the Domain Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The display name for your managed Active Directory Domain Service resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `notifications` block as defined below.
        /// </summary>
        [Input("notifications")]
        public Input<Inputs.ServiceNotificationsGetArgs>? Notifications { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Domain Service should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Azure resource ID for the domain service.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// A `secure_ldap` block as defined below.
        /// </summary>
        [Input("secureLdap")]
        public Input<Inputs.ServiceSecureLdapGetArgs>? SecureLdap { get; set; }

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Input("security")]
        public Input<Inputs.ServiceSecurityGetArgs>? Security { get; set; }

        /// <summary>
        /// The SKU to use when provisioning the Domain Service resource. One of `Standard`, `Enterprise` or `Premium`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("syncOwner")]
        public Input<string>? SyncOwner { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public ServiceState()
        {
        }
    }
}
