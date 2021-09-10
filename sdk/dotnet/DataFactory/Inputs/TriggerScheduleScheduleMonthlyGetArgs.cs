// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class TriggerScheduleScheduleMonthlyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The occurrence of the specified day during the month. For example, a `monthly` property with `weekday` and `week` values of `Sunday, -1` means the last Sunday of the month.
        /// </summary>
        [Input("week")]
        public Input<int>? Week { get; set; }

        /// <summary>
        /// The day of the week on which the trigger runs. For example, a `monthly` property with a `weekday` value of `Sunday` means every Sunday of the month.
        /// </summary>
        [Input("weekday", required: true)]
        public Input<string> Weekday { get; set; } = null!;

        public TriggerScheduleScheduleMonthlyGetArgs()
        {
        }
    }
}
