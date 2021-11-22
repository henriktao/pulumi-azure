// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Inputs
{

    public sealed class GetPoolStartTaskUserIdentityAutoUserInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The elevation level of the user identity under which the start task runs.
        /// </summary>
        [Input("elevationLevel", required: true)]
        public Input<string> ElevationLevel { get; set; } = null!;

        /// <summary>
        /// The scope of the user identity under which the start task runs.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public GetPoolStartTaskUserIdentityAutoUserInputArgs()
        {
        }
    }
}
