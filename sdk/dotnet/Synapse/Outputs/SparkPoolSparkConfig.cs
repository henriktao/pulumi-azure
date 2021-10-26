// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Synapse.Outputs
{

    [OutputType]
    public sealed class SparkPoolSparkConfig
    {
        /// <summary>
        /// The contents of a spark configuration.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// The name of the file where the spark configuration `content` will be stored.
        /// </summary>
        public readonly string Filename;

        [OutputConstructor]
        private SparkPoolSparkConfig(
            string content,

            string filename)
        {
            Content = content;
            Filename = filename;
        }
    }
}
