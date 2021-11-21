// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core.Inputs
{

    public sealed class ResourceGroupCostManagementExportExportDataOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        /// </summary>
        [Input("timeFrame", required: true)]
        public Input<string> TimeFrame { get; set; } = null!;

        /// <summary>
        /// The type of the query.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourceGroupCostManagementExportExportDataOptionsGetArgs()
        {
        }
    }
}
