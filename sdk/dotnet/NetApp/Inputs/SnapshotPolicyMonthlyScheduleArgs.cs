// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp.Inputs
{

    public sealed class SnapshotPolicyMonthlyScheduleArgs : Pulumi.ResourceArgs
    {
        [Input("daysOfMonths", required: true)]
        private InputList<int>? _daysOfMonths;
        public InputList<int> DaysOfMonths
        {
            get => _daysOfMonths ?? (_daysOfMonths = new InputList<int>());
            set => _daysOfMonths = value;
        }

        /// <summary>
        /// Hour of the day that the snapshots will be created, valid range is from 0 to 23.
        /// </summary>
        [Input("hour", required: true)]
        public Input<int> Hour { get; set; } = null!;

        /// <summary>
        /// Minute of the hour that the snapshots will be created, valid range is from 0 to 59.
        /// </summary>
        [Input("minute", required: true)]
        public Input<int> Minute { get; set; } = null!;

        /// <summary>
        /// How many hourly snapshots to keep, valid range is from 0 to 255.
        /// </summary>
        [Input("snapshotsToKeep", required: true)]
        public Input<int> SnapshotsToKeep { get; set; } = null!;

        public SnapshotPolicyMonthlyScheduleArgs()
        {
        }
    }
}
