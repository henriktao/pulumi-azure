// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class ActivityLogAlertCriteriaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email address or Azure Active Directory identifier of the user who performed the operation.
        /// </summary>
        [Input("caller")]
        public Input<string>? Caller { get; set; }

        /// <summary>
        /// The category of the operation. Possible values are `Administrative`, `Autoscale`, `Policy`, `Recommendation`, `ResourceHealth`, `Security` and `ServiceHealth`.
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// The severity level of the event. Possible values are `Verbose`, `Informational`, `Warning`, `Error`, and `Critical`.
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// The Resource Manager Role-Based Access Control operation name. Supported operation should be of the form: `&lt;resourceProvider&gt;/&lt;resourceType&gt;/&lt;operation&gt;`.
        /// </summary>
        [Input("operationName")]
        public Input<string>? OperationName { get; set; }

        /// <summary>
        /// The recommendation category of the event. Possible values are `Cost`, `Reliability`, `OperationalExcellence` and `Performance`. It is only allowed when `category` is `Recommendation`.
        /// </summary>
        [Input("recommendationCategory")]
        public Input<string>? RecommendationCategory { get; set; }

        /// <summary>
        /// The recommendation impact of the event. Possible values are `High`, `Medium` and `Low`. It is only allowed when `category` is `Recommendation`.
        /// </summary>
        [Input("recommendationImpact")]
        public Input<string>? RecommendationImpact { get; set; }

        /// <summary>
        /// The recommendation type of the event. It is only allowed when `category` is `Recommendation`.
        /// </summary>
        [Input("recommendationType")]
        public Input<string>? RecommendationType { get; set; }

        /// <summary>
        /// The name of resource group monitored by the activity log alert.
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        [Input("resourceHealths")]
        private InputList<Inputs.ActivityLogAlertCriteriaResourceHealthGetArgs>? _resourceHealths;

        /// <summary>
        /// A block to define fine grain resource health settings.
        /// </summary>
        public InputList<Inputs.ActivityLogAlertCriteriaResourceHealthGetArgs> ResourceHealths
        {
            get => _resourceHealths ?? (_resourceHealths = new InputList<Inputs.ActivityLogAlertCriteriaResourceHealthGetArgs>());
            set => _resourceHealths = value;
        }

        /// <summary>
        /// The specific resource monitored by the activity log alert. It should be within one of the `scopes`.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The name of the resource provider monitored by the activity log alert.
        /// </summary>
        [Input("resourceProvider")]
        public Input<string>? ResourceProvider { get; set; }

        /// <summary>
        /// The resource type monitored by the activity log alert.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        [Input("serviceHealths")]
        private InputList<Inputs.ActivityLogAlertCriteriaServiceHealthGetArgs>? _serviceHealths;

        /// <summary>
        /// A block to define fine grain service health settings.
        /// </summary>
        public InputList<Inputs.ActivityLogAlertCriteriaServiceHealthGetArgs> ServiceHealths
        {
            get => _serviceHealths ?? (_serviceHealths = new InputList<Inputs.ActivityLogAlertCriteriaServiceHealthGetArgs>());
            set => _serviceHealths = value;
        }

        /// <summary>
        /// The status of the event. For example, `Started`, `Failed`, or `Succeeded`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The sub status of the event.
        /// </summary>
        [Input("subStatus")]
        public Input<string>? SubStatus { get; set; }

        public ActivityLogAlertCriteriaGetArgs()
        {
        }
    }
}
