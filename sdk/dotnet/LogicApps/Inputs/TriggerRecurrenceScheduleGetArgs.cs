// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps.Inputs
{

    public sealed class TriggerRecurrenceScheduleGetArgs : Pulumi.ResourceArgs
    {
        [Input("atTheseHours")]
        private InputList<int>? _atTheseHours;

        /// <summary>
        /// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
        /// </summary>
        public InputList<int> AtTheseHours
        {
            get => _atTheseHours ?? (_atTheseHours = new InputList<int>());
            set => _atTheseHours = value;
        }

        [Input("atTheseMinutes")]
        private InputList<int>? _atTheseMinutes;

        /// <summary>
        /// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
        /// </summary>
        public InputList<int> AtTheseMinutes
        {
            get => _atTheseMinutes ?? (_atTheseMinutes = new InputList<int>());
            set => _atTheseMinutes = value;
        }

        [Input("onTheseDays")]
        private InputList<string>? _onTheseDays;

        /// <summary>
        /// Specifies a list of days when the trigger should run. Valid values include `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`.
        /// </summary>
        public InputList<string> OnTheseDays
        {
            get => _onTheseDays ?? (_onTheseDays = new InputList<string>());
            set => _onTheseDays = value;
        }

        public TriggerRecurrenceScheduleGetArgs()
        {
        }
    }
}
