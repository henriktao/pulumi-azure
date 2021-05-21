// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Redis.Outputs
{

    [OutputType]
    public sealed class CacheRedisConfiguration
    {
        /// <summary>
        /// Enable or disable AOF persistence for this Redis Cache.
        /// </summary>
        public readonly bool? AofBackupEnabled;
        /// <summary>
        /// First Storage Account connection string for AOF persistence.
        /// </summary>
        public readonly string? AofStorageConnectionString0;
        /// <summary>
        /// Second Storage Account connection string for AOF persistence.
        /// </summary>
        public readonly string? AofStorageConnectionString1;
        /// <summary>
        /// If set to `false`, the Redis instance will be accessible without authentication. Defaults to `true`.
        /// </summary>
        public readonly bool? EnableAuthentication;
        /// <summary>
        /// Returns the max number of connected clients at the same time.
        /// </summary>
        public readonly int? Maxclients;
        /// <summary>
        /// Value in megabytes reserved to accommodate for memory fragmentation. Defaults are shown below.
        /// </summary>
        public readonly int? MaxfragmentationmemoryReserved;
        /// <summary>
        /// The max-memory delta for this Redis instance. Defaults are shown below.
        /// </summary>
        public readonly int? MaxmemoryDelta;
        /// <summary>
        /// How Redis will select what to remove when `maxmemory` is reached. Defaults are shown below.
        /// </summary>
        public readonly string? MaxmemoryPolicy;
        /// <summary>
        /// Value in megabytes reserved for non-cache usage e.g. failover. Defaults are shown below.
        /// </summary>
        public readonly int? MaxmemoryReserved;
        /// <summary>
        /// Keyspace notifications allows clients to subscribe to Pub/Sub channels in order to receive events affecting the Redis data set in some way. [Reference](https://redis.io/topics/notifications#configuration)
        /// </summary>
        public readonly string? NotifyKeyspaceEvents;
        /// <summary>
        /// Is Backup Enabled? Only supported on Premium SKU's.
        /// </summary>
        public readonly bool? RdbBackupEnabled;
        /// <summary>
        /// The Backup Frequency in Minutes. Only supported on Premium SKU's. Possible values are: `15`, `30`, `60`, `360`, `720` and `1440`.
        /// </summary>
        public readonly int? RdbBackupFrequency;
        /// <summary>
        /// The maximum number of snapshots to create as a backup. Only supported for Premium SKU's.
        /// </summary>
        public readonly int? RdbBackupMaxSnapshotCount;
        /// <summary>
        /// The Connection String to the Storage Account. Only supported for Premium SKU's. In the format: `DefaultEndpointsProtocol=https;BlobEndpoint=${azurerm_storage_account.example.primary_blob_endpoint};AccountName=${azurerm_storage_account.example.name};AccountKey=${azurerm_storage_account.example.primary_access_key}`.
        /// </summary>
        public readonly string? RdbStorageConnectionString;

        [OutputConstructor]
        private CacheRedisConfiguration(
            bool? aofBackupEnabled,

            string? aofStorageConnectionString0,

            string? aofStorageConnectionString1,

            bool? enableAuthentication,

            int? maxclients,

            int? maxfragmentationmemoryReserved,

            int? maxmemoryDelta,

            string? maxmemoryPolicy,

            int? maxmemoryReserved,

            string? notifyKeyspaceEvents,

            bool? rdbBackupEnabled,

            int? rdbBackupFrequency,

            int? rdbBackupMaxSnapshotCount,

            string? rdbStorageConnectionString)
        {
            AofBackupEnabled = aofBackupEnabled;
            AofStorageConnectionString0 = aofStorageConnectionString0;
            AofStorageConnectionString1 = aofStorageConnectionString1;
            EnableAuthentication = enableAuthentication;
            Maxclients = maxclients;
            MaxfragmentationmemoryReserved = maxfragmentationmemoryReserved;
            MaxmemoryDelta = maxmemoryDelta;
            MaxmemoryPolicy = maxmemoryPolicy;
            MaxmemoryReserved = maxmemoryReserved;
            NotifyKeyspaceEvents = notifyKeyspaceEvents;
            RdbBackupEnabled = rdbBackupEnabled;
            RdbBackupFrequency = rdbBackupFrequency;
            RdbBackupMaxSnapshotCount = rdbBackupMaxSnapshotCount;
            RdbStorageConnectionString = rdbStorageConnectionString;
        }
    }
}
