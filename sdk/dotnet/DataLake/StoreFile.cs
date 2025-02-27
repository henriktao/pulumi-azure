// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataLake
{
    /// <summary>
    /// Manages a Azure Data Lake Store File.
    /// 
    /// &gt; **Note:** This resoruce manages an `Azure Data Lake Storage Gen1`, previously known as `Azure Data Lake Store`.
    /// 
    /// &gt; **Note:** If you want to change the data in the remote file without changing the `local_file_path`, then
    /// taint the resource so the `azure.datalake.StoreFile` gets recreated with the new data.
    /// 
    /// ## Import
    /// 
    /// Data Lake Store File's can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datalake/storeFile:StoreFile example example.azuredatalakestore.net/test/example.txt
    /// ```
    /// </summary>
    [AzureResourceType("azure:datalake/storeFile:StoreFile")]
    public partial class StoreFile : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the File should created.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The path to the local file to be added to the Data Lake Store.
        /// </summary>
        [Output("localFilePath")]
        public Output<string> LocalFilePath { get; private set; } = null!;

        /// <summary>
        /// The path created for the file on the Data Lake Store.
        /// </summary>
        [Output("remoteFilePath")]
        public Output<string> RemoteFilePath { get; private set; } = null!;


        /// <summary>
        /// Create a StoreFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StoreFile(string name, StoreFileArgs args, CustomResourceOptions? options = null)
            : base("azure:datalake/storeFile:StoreFile", name, args ?? new StoreFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StoreFile(string name, Input<string> id, StoreFileState? state = null, CustomResourceOptions? options = null)
            : base("azure:datalake/storeFile:StoreFile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StoreFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StoreFile Get(string name, Input<string> id, StoreFileState? state = null, CustomResourceOptions? options = null)
        {
            return new StoreFile(name, id, state, options);
        }
    }

    public sealed class StoreFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the File should created.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The path to the local file to be added to the Data Lake Store.
        /// </summary>
        [Input("localFilePath", required: true)]
        public Input<string> LocalFilePath { get; set; } = null!;

        /// <summary>
        /// The path created for the file on the Data Lake Store.
        /// </summary>
        [Input("remoteFilePath", required: true)]
        public Input<string> RemoteFilePath { get; set; } = null!;

        public StoreFileArgs()
        {
        }
    }

    public sealed class StoreFileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the File should created.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The path to the local file to be added to the Data Lake Store.
        /// </summary>
        [Input("localFilePath")]
        public Input<string>? LocalFilePath { get; set; }

        /// <summary>
        /// The path created for the file on the Data Lake Store.
        /// </summary>
        [Input("remoteFilePath")]
        public Input<string>? RemoteFilePath { get; set; }

        public StoreFileState()
        {
        }
    }
}
