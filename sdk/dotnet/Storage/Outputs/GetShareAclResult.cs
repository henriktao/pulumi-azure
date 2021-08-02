// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Outputs
{

    [OutputType]
    public sealed class GetShareAclResult
    {
        /// <summary>
        /// An `access_policy` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetShareAclAccessPolicyResult> AccessPolicies;
        /// <summary>
        /// The ID which should be used for this Shared Identifier.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetShareAclResult(
            ImmutableArray<Outputs.GetShareAclAccessPolicyResult> accessPolicies,

            string id)
        {
            AccessPolicies = accessPolicies;
            Id = id;
        }
    }
}
