// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.VideoAnalyzer.Inputs
{

    public sealed class AnalyzerStorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of the Storage Account that will be associated with the Video Analyzer instance.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Specifies the User Assigned Identity ID which should be assigned to a access this Storage Account.
        /// </summary>
        [Input("userAssignedIdentityId", required: true)]
        public Input<string> UserAssignedIdentityId { get; set; } = null!;

        public AnalyzerStorageAccountArgs()
        {
        }
    }
}
