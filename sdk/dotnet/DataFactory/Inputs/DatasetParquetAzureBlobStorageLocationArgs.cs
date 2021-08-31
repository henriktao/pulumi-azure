// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class DatasetParquetAzureBlobStorageLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The container on the Azure Blob Storage Account hosting the file.
        /// </summary>
        [Input("container", required: true)]
        public Input<string> Container { get; set; } = null!;

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
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// The folder path to the file on the web server.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public DatasetParquetAzureBlobStorageLocationArgs()
        {
        }
    }
}
