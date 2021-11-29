// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Outputs
{

    [OutputType]
    public sealed class AccountBackup
    {
        /// <summary>
        /// The interval in minutes between two backups. This is configurable only when `type` is `Periodic`. Possible values are between 60 and 1440.
        /// </summary>
        public readonly int? IntervalInMinutes;
        /// <summary>
        /// The time in hours that each backup is retained. This is configurable only when `type` is `Periodic`. Possible values are between 8 and 720.
        /// </summary>
        public readonly int? RetentionInHours;
        /// <summary>
        /// The storage redundancy which is used to indicate type of backup residency. This is configurable only when `type` is `Periodic`. Possible values are `Geo`, `Local` and `Zone`.
        /// </summary>
        public readonly string? StorageRedundancy;
        /// <summary>
        /// The type of the `backup`. Possible values are `Continuous` and `Periodic`. Defaults to `Periodic`. Migration of `Periodic` to `Continuous` is one-way, changing `Continuous` to `Periodic` forces a new resource to be created.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccountBackup(
            int? intervalInMinutes,

            int? retentionInHours,

            string? storageRedundancy,

            string type)
        {
            IntervalInMinutes = intervalInMinutes;
            RetentionInHours = retentionInHours;
            StorageRedundancy = storageRedundancy;
            Type = type;
        }
    }
}
