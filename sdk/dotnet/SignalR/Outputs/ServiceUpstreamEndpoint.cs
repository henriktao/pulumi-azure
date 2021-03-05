// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SignalR.Outputs
{

    [OutputType]
    public sealed class ServiceUpstreamEndpoint
    {
        /// <summary>
        /// The categories to match on, or `*` for all.
        /// </summary>
        public readonly ImmutableArray<string> CategoryPatterns;
        /// <summary>
        /// The events to match on, or `*` for all.
        /// </summary>
        public readonly ImmutableArray<string> EventPatterns;
        /// <summary>
        /// The hubs to match on, or `*` for all.
        /// </summary>
        public readonly ImmutableArray<string> HubPatterns;
        /// <summary>
        /// The upstream URL Template. This can be a url or a template such as `http://host.com/{hub}/api/{category}/{event}`.
        /// </summary>
        public readonly string UrlTemplate;

        [OutputConstructor]
        private ServiceUpstreamEndpoint(
            ImmutableArray<string> categoryPatterns,

            ImmutableArray<string> eventPatterns,

            ImmutableArray<string> hubPatterns,

            string urlTemplate)
        {
            CategoryPatterns = categoryPatterns;
            EventPatterns = eventPatterns;
            HubPatterns = hubPatterns;
            UrlTemplate = urlTemplate;
        }
    }
}
