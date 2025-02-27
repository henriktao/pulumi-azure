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
    public sealed class ManagementPolicyRuleActionsBaseBlob
    {
        /// <summary>
        /// The age in days after last access time to delete the blob. Must be between `0` and `99999`.
        /// </summary>
        public readonly int? DeleteAfterDaysSinceLastAccessTimeGreaterThan;
        /// <summary>
        /// The age in days after last modification to delete the blob. Must be between 0 and 99999.
        /// </summary>
        public readonly int? DeleteAfterDaysSinceModificationGreaterThan;
        /// <summary>
        /// The age in days after last access time to tier blobs to archive storage. Supports blob currently at Hot or Cool tier. Must be between `0 and `99999`.
        /// </summary>
        public readonly int? TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan;
        /// <summary>
        /// The age in days after last modification to tier blobs to archive storage. Supports blob currently at Hot or Cool tier. Must be between 0 and 99999.
        /// </summary>
        public readonly int? TierToArchiveAfterDaysSinceModificationGreaterThan;
        /// <summary>
        /// The age in days after last access time to tier blobs to cool storage. Supports blob currently at Hot tier. Must be between `0` and `99999`.
        /// </summary>
        public readonly int? TierToCoolAfterDaysSinceLastAccessTimeGreaterThan;
        /// <summary>
        /// The age in days after last modification to tier blobs to cool storage. Supports blob currently at Hot tier. Must be between 0 and 99999.
        /// </summary>
        public readonly int? TierToCoolAfterDaysSinceModificationGreaterThan;

        [OutputConstructor]
        private ManagementPolicyRuleActionsBaseBlob(
            int? deleteAfterDaysSinceLastAccessTimeGreaterThan,

            int? deleteAfterDaysSinceModificationGreaterThan,

            int? tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan,

            int? tierToArchiveAfterDaysSinceModificationGreaterThan,

            int? tierToCoolAfterDaysSinceLastAccessTimeGreaterThan,

            int? tierToCoolAfterDaysSinceModificationGreaterThan)
        {
            DeleteAfterDaysSinceLastAccessTimeGreaterThan = deleteAfterDaysSinceLastAccessTimeGreaterThan;
            DeleteAfterDaysSinceModificationGreaterThan = deleteAfterDaysSinceModificationGreaterThan;
            TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan = tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan;
            TierToArchiveAfterDaysSinceModificationGreaterThan = tierToArchiveAfterDaysSinceModificationGreaterThan;
            TierToCoolAfterDaysSinceLastAccessTimeGreaterThan = tierToCoolAfterDaysSinceLastAccessTimeGreaterThan;
            TierToCoolAfterDaysSinceModificationGreaterThan = tierToCoolAfterDaysSinceModificationGreaterThan;
        }
    }
}
