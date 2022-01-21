// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql.Inputs
{

    public sealed class FlexibleServerMaintenanceWindowGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of week for maintenance window. Defaults to `0`.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<int>? DayOfWeek { get; set; }

        /// <summary>
        /// The start hour for maintenance window. Defaults to `0`.
        /// </summary>
        [Input("startHour")]
        public Input<int>? StartHour { get; set; }

        /// <summary>
        /// The start minute for maintenance window. Defaults to `0`.
        /// </summary>
        [Input("startMinute")]
        public Input<int>? StartMinute { get; set; }

        public FlexibleServerMaintenanceWindowGetArgs()
        {
        }
    }
}
