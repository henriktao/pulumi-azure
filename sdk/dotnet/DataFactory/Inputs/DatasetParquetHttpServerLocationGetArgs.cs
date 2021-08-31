// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class DatasetParquetHttpServerLocationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the `filename` using dynamic expression, function or system variables? Defaults to `false`.
        /// </summary>
        [Input("dynamicFilenameEnabled")]
        public Input<bool>? DynamicFilenameEnabled { get; set; }

        /// <summary>
        /// Is the `path` using dynamic expression, function or system variables? Defaults to `false`.
        /// </summary>
        [Input("dynamicPathEnabled")]
        public Input<bool>? DynamicPathEnabled { get; set; }

        /// <summary>
        /// The filename of the file on the web server.
        /// </summary>
        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        /// <summary>
        /// The folder path to the file on the web server.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The base URL to the web server hosting the file.
        /// </summary>
        [Input("relativeUrl", required: true)]
        public Input<string> RelativeUrl { get; set; } = null!;

        public DatasetParquetHttpServerLocationGetArgs()
        {
        }
    }
}
