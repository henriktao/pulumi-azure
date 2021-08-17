// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.VideoAnalyzer.Outputs
{

    [OutputType]
    public sealed class AnalyzerStorageAccount
    {
        /// <summary>
        /// Specifies the ID of the Storage Account that will be associated with the Video Analyzer instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the User Assigned Identity ID which should be assigned to a access this Storage Account.
        /// </summary>
        public readonly string UserAssignedIdentityId;

        [OutputConstructor]
        private AnalyzerStorageAccount(
            string id,

            string userAssignedIdentityId)
        {
            Id = id;
            UserAssignedIdentityId = userAssignedIdentityId;
        }
    }
}
