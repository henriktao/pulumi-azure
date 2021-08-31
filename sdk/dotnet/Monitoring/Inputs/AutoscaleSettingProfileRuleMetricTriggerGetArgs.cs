// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class AutoscaleSettingProfileRuleMetricTriggerGetArgs : Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.AutoscaleSettingProfileRuleMetricTriggerDimensionGetArgs>? _dimensions;

        /// <summary>
        /// One or more `dimensions` block as defined below.
        /// </summary>
        public InputList<Inputs.AutoscaleSettingProfileRuleMetricTriggerDimensionGetArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.AutoscaleSettingProfileRuleMetricTriggerDimensionGetArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Whether to enable metric divide by instance count.
        /// </summary>
        [Input("divideByInstanceCount")]
        public Input<bool>? DivideByInstanceCount { get; set; }

        /// <summary>
        /// The name of the metric that defines what the rule monitors, such as `Percentage CPU` for `Virtual Machine Scale Sets` and `CpuPercentage` for `App Service Plan`.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// The namespace of the metric that defines what the rule monitors, such as `microsoft.compute/virtualmachinescalesets` for `Virtual Machine Scale Sets`.
        /// </summary>
        [Input("metricNamespace")]
        public Input<string>? MetricNamespace { get; set; }

        /// <summary>
        /// The ID of the Resource which the Rule monitors.
        /// </summary>
        [Input("metricResourceId", required: true)]
        public Input<string> MetricResourceId { get; set; } = null!;

        /// <summary>
        /// Specifies the operator used to compare the metric data and threshold. Possible values are: `Equals`, `NotEquals`, `GreaterThan`, `GreaterThanOrEqual`, `LessThan`, `LessThanOrEqual`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Specifies how the metrics from multiple instances are combined. Possible values are `Average`, `Min` and `Max`.
        /// </summary>
        [Input("statistic", required: true)]
        public Input<string> Statistic { get; set; } = null!;

        /// <summary>
        /// Specifies the threshold of the metric that triggers the scale action.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        /// <summary>
        /// Specifies how the data that's collected should be combined over time. Possible values include `Average`, `Count`, `Maximum`, `Minimum`, `Last` and `Total`. Defaults to `Average`.
        /// </summary>
        [Input("timeAggregation", required: true)]
        public Input<string> TimeAggregation { get; set; } = null!;

        /// <summary>
        /// Specifies the granularity of metrics that the rule monitors, which must be one of the pre-defined values returned from the metric definitions for the metric. This value must be between 1 minute and 12 hours an be formatted as an ISO 8601 string.
        /// </summary>
        [Input("timeGrain", required: true)]
        public Input<string> TimeGrain { get; set; } = null!;

        /// <summary>
        /// Specifies the time range for which data is collected, which must be greater than the delay in metric collection (which varies from resource to resource). This value must be between 5 minutes and 12 hours and be formatted as an ISO 8601 string.
        /// </summary>
        [Input("timeWindow", required: true)]
        public Input<string> TimeWindow { get; set; } = null!;

        public AutoscaleSettingProfileRuleMetricTriggerGetArgs()
        {
        }
    }
}
