// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PostgreSql.Outputs
{

    [OutputType]
    public sealed class ServerStorageProfile
    {
        public readonly string? AutoGrow;
        /// <summary>
        /// Backup retention days for the server, supported values are between `7` and `35` days.
        /// </summary>
        public readonly int? BackupRetentionDays;
        public readonly string? GeoRedundantBackup;
        /// <summary>
        /// Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/azure/postgresql/concepts-pricing-tiers).
        /// </summary>
        public readonly int? StorageMb;

        [OutputConstructor]
        private ServerStorageProfile(
            string? autoGrow,

            int? backupRetentionDays,

            string? geoRedundantBackup,

            int? storageMb)
        {
            AutoGrow = autoGrow;
            BackupRetentionDays = backupRetentionDays;
            GeoRedundantBackup = geoRedundantBackup;
            StorageMb = storageMb;
        }
    }
}
