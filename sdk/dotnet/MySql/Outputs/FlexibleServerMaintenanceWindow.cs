// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql.Outputs
{

    [OutputType]
    public sealed class FlexibleServerMaintenanceWindow
    {
        /// <summary>
        /// The day of week for maintenance window. Defaults to `0`.
        /// </summary>
        public readonly int? DayOfWeek;
        /// <summary>
        /// The start hour for maintenance window. Defaults to `0`.
        /// </summary>
        public readonly int? StartHour;
        /// <summary>
        /// The start minute for maintenance window. Defaults to `0`.
        /// </summary>
        public readonly int? StartMinute;

        [OutputConstructor]
        private FlexibleServerMaintenanceWindow(
            int? dayOfWeek,

            int? startHour,

            int? startMinute)
        {
            DayOfWeek = dayOfWeek;
            StartHour = startHour;
            StartMinute = startMinute;
        }
    }
}
