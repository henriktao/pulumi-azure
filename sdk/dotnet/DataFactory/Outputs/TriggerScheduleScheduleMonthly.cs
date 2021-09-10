// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class TriggerScheduleScheduleMonthly
    {
        /// <summary>
        /// The occurrence of the specified day during the month. For example, a `monthly` property with `weekday` and `week` values of `Sunday, -1` means the last Sunday of the month.
        /// </summary>
        public readonly int? Week;
        /// <summary>
        /// The day of the week on which the trigger runs. For example, a `monthly` property with a `weekday` value of `Sunday` means every Sunday of the month.
        /// </summary>
        public readonly string Weekday;

        [OutputConstructor]
        private TriggerScheduleScheduleMonthly(
            int? week,

            string weekday)
        {
            Week = week;
            Weekday = weekday;
        }
    }
}
