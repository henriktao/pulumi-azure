// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class InteractiveQueryClusterSecurityProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the Azure Active Directory Domain Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("aaddsResourceId", required: true)]
        public Input<string> AaddsResourceId { get; set; } = null!;

        [Input("clusterUsersGroupDns")]
        private InputList<string>? _clusterUsersGroupDns;

        /// <summary>
        /// A list of the distinguished names for the cluster user groups. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> ClusterUsersGroupDns
        {
            get => _clusterUsersGroupDns ?? (_clusterUsersGroupDns = new InputList<string>());
            set => _clusterUsersGroupDns = value;
        }

        /// <summary>
        /// The name of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The user password of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("domainUserPassword", required: true)]
        public Input<string> DomainUserPassword { get; set; } = null!;

        /// <summary>
        /// The username of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("domainUsername", required: true)]
        public Input<string> DomainUsername { get; set; } = null!;

        [Input("ldapsUrls", required: true)]
        private InputList<string>? _ldapsUrls;

        /// <summary>
        /// A list of the LDAPS URLs to communicate with the Azure Active Directory. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> LdapsUrls
        {
            get => _ldapsUrls ?? (_ldapsUrls = new InputList<string>());
            set => _ldapsUrls = value;
        }

        /// <summary>
        /// The User Assigned Identity for the HDInsight Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("msiResourceId", required: true)]
        public Input<string> MsiResourceId { get; set; } = null!;

        public InteractiveQueryClusterSecurityProfileArgs()
        {
        }
    }
}
