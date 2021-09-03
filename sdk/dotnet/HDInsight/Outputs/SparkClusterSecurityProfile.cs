// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class SparkClusterSecurityProfile
    {
        /// <summary>
        /// The resource ID of the Azure Active Directory Domain Service. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string AaddsResourceId;
        /// <summary>
        /// A list of the distinguished names for the cluster user groups. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> ClusterUsersGroupDns;
        /// <summary>
        /// The name of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The user password of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string DomainUserPassword;
        /// <summary>
        /// The username of the Azure Active Directory Domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string DomainUsername;
        /// <summary>
        /// A list of the LDAPS URLs to communicate with the Azure Active Directory. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> LdapsUrls;
        /// <summary>
        /// The User Assigned Identity for the HDInsight Cluster. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string MsiResourceId;

        [OutputConstructor]
        private SparkClusterSecurityProfile(
            string aaddsResourceId,

            ImmutableArray<string> clusterUsersGroupDns,

            string domainName,

            string domainUserPassword,

            string domainUsername,

            ImmutableArray<string> ldapsUrls,

            string msiResourceId)
        {
            AaddsResourceId = aaddsResourceId;
            ClusterUsersGroupDns = clusterUsersGroupDns;
            DomainName = domainName;
            DomainUserPassword = domainUserPassword;
            DomainUsername = domainUsername;
            LdapsUrls = ldapsUrls;
            MsiResourceId = msiResourceId;
        }
    }
}
