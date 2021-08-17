// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    /// <summary>
    /// Manages an Azure Delimited Text Dataset inside an Azure Data Factory.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleFactory = new Azure.DataFactory.Factory("exampleFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleLinkedServiceWeb = new Azure.DataFactory.LinkedServiceWeb("exampleLinkedServiceWeb", new Azure.DataFactory.LinkedServiceWebArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DataFactoryName = exampleFactory.Name,
    ///             AuthenticationType = "Anonymous",
    ///             Url = "https://www.bing.com",
    ///         });
    ///         var exampleDatasetDelimitedText = new Azure.DataFactory.DatasetDelimitedText("exampleDatasetDelimitedText", new Azure.DataFactory.DatasetDelimitedTextArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DataFactoryName = exampleFactory.Name,
    ///             LinkedServiceName = exampleLinkedServiceWeb.Name,
    ///             HttpServerLocation = new Azure.DataFactory.Inputs.DatasetDelimitedTextHttpServerLocationArgs
    ///             {
    ///                 RelativeUrl = "http://www.bing.com",
    ///                 Path = "foo/bar/",
    ///                 Filename = "fizz.txt",
    ///             },
    ///             ColumnDelimiter = ",",
    ///             RowDelimiter = "NEW",
    ///             Encoding = "UTF-8",
    ///             QuoteCharacter = "x",
    ///             EscapeCharacter = "f",
    ///             FirstRowAsHeader = true,
    ///             NullValue = "NULL",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Datasets can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/datasetDelimitedText:DatasetDelimitedText example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/datasetDelimitedText:DatasetDelimitedText")]
    public partial class DatasetDelimitedText : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset.
        /// </summary>
        [Output("additionalProperties")]
        public Output<ImmutableDictionary<string, string>?> AdditionalProperties { get; private set; } = null!;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// An `azure_blob_fs_location` block as defined below.
        /// </summary>
        [Output("azureBlobFsLocation")]
        public Output<Outputs.DatasetDelimitedTextAzureBlobFsLocation?> AzureBlobFsLocation { get; private set; } = null!;

        /// <summary>
        /// An `azure_blob_storage_location` block as defined below.
        /// </summary>
        [Output("azureBlobStorageLocation")]
        public Output<Outputs.DatasetDelimitedTextAzureBlobStorageLocation?> AzureBlobStorageLocation { get; private set; } = null!;

        /// <summary>
        /// The column delimiter. Defaults to `,`.
        /// </summary>
        [Output("columnDelimiter")]
        public Output<string?> ColumnDelimiter { get; private set; } = null!;

        /// <summary>
        /// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
        /// </summary>
        [Output("compressionCodec")]
        public Output<string?> CompressionCodec { get; private set; } = null!;

        /// <summary>
        /// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
        /// </summary>
        [Output("compressionLevel")]
        public Output<string?> CompressionLevel { get; private set; } = null!;

        /// <summary>
        /// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        /// </summary>
        [Output("dataFactoryName")]
        public Output<string> DataFactoryName { get; private set; } = null!;

        /// <summary>
        /// The description for the Data Factory Dataset.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The encoding format for the file.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// The escape character. Defaults to `\`.
        /// </summary>
        [Output("escapeCharacter")]
        public Output<string?> EscapeCharacter { get; private set; } = null!;

        /// <summary>
        /// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to `false`.
        /// </summary>
        [Output("firstRowAsHeader")]
        public Output<bool?> FirstRowAsHeader { get; private set; } = null!;

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// A `http_server_location` block as defined below.
        /// </summary>
        [Output("httpServerLocation")]
        public Output<Outputs.DatasetDelimitedTextHttpServerLocation?> HttpServerLocation { get; private set; } = null!;

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Output("linkedServiceName")]
        public Output<string> LinkedServiceName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The null value string. Defaults to an empty string.
        /// </summary>
        [Output("nullValue")]
        public Output<string?> NullValue { get; private set; } = null!;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The quote character. Defaults to `"`.
        /// </summary>
        [Output("quoteCharacter")]
        public Output<string?> QuoteCharacter { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The row delimiter. Defaults to any of the following values on read: `\r\n`, `\r`, `\n`, and `\n` or `\r\n` on write by mapping data flow and Copy activity respectively.
        /// </summary>
        [Output("rowDelimiter")]
        public Output<string?> RowDelimiter { get; private set; } = null!;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        [Output("schemaColumns")]
        public Output<ImmutableArray<Outputs.DatasetDelimitedTextSchemaColumn>> SchemaColumns { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetDelimitedText resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetDelimitedText(string name, DatasetDelimitedTextArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/datasetDelimitedText:DatasetDelimitedText", name, args ?? new DatasetDelimitedTextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetDelimitedText(string name, Input<string> id, DatasetDelimitedTextState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/datasetDelimitedText:DatasetDelimitedText", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetDelimitedText resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetDelimitedText Get(string name, Input<string> id, DatasetDelimitedTextState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetDelimitedText(name, id, state, options);
        }
    }

    public sealed class DatasetDelimitedTextArgs : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset.
        /// </summary>
        public InputList<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// An `azure_blob_fs_location` block as defined below.
        /// </summary>
        [Input("azureBlobFsLocation")]
        public Input<Inputs.DatasetDelimitedTextAzureBlobFsLocationArgs>? AzureBlobFsLocation { get; set; }

        /// <summary>
        /// An `azure_blob_storage_location` block as defined below.
        /// </summary>
        [Input("azureBlobStorageLocation")]
        public Input<Inputs.DatasetDelimitedTextAzureBlobStorageLocationArgs>? AzureBlobStorageLocation { get; set; }

        /// <summary>
        /// The column delimiter. Defaults to `,`.
        /// </summary>
        [Input("columnDelimiter")]
        public Input<string>? ColumnDelimiter { get; set; }

        /// <summary>
        /// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
        /// </summary>
        [Input("compressionCodec")]
        public Input<string>? CompressionCodec { get; set; }

        /// <summary>
        /// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
        /// </summary>
        [Input("compressionLevel")]
        public Input<string>? CompressionLevel { get; set; }

        /// <summary>
        /// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryName", required: true)]
        public Input<string> DataFactoryName { get; set; } = null!;

        /// <summary>
        /// The description for the Data Factory Dataset.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The encoding format for the file.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The escape character. Defaults to `\`.
        /// </summary>
        [Input("escapeCharacter")]
        public Input<string>? EscapeCharacter { get; set; }

        /// <summary>
        /// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to `false`.
        /// </summary>
        [Input("firstRowAsHeader")]
        public Input<bool>? FirstRowAsHeader { get; set; }

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// A `http_server_location` block as defined below.
        /// </summary>
        [Input("httpServerLocation")]
        public Input<Inputs.DatasetDelimitedTextHttpServerLocationArgs>? HttpServerLocation { get; set; }

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public Input<string> LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The null value string. Defaults to an empty string.
        /// </summary>
        [Input("nullValue")]
        public Input<string>? NullValue { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The quote character. Defaults to `"`.
        /// </summary>
        [Input("quoteCharacter")]
        public Input<string>? QuoteCharacter { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The row delimiter. Defaults to any of the following values on read: `\r\n`, `\r`, `\n`, and `\n` or `\r\n` on write by mapping data flow and Copy activity respectively.
        /// </summary>
        [Input("rowDelimiter")]
        public Input<string>? RowDelimiter { get; set; }

        [Input("schemaColumns")]
        private InputList<Inputs.DatasetDelimitedTextSchemaColumnArgs>? _schemaColumns;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        public InputList<Inputs.DatasetDelimitedTextSchemaColumnArgs> SchemaColumns
        {
            get => _schemaColumns ?? (_schemaColumns = new InputList<Inputs.DatasetDelimitedTextSchemaColumnArgs>());
            set => _schemaColumns = value;
        }

        public DatasetDelimitedTextArgs()
        {
        }
    }

    public sealed class DatasetDelimitedTextState : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset.
        /// </summary>
        public InputList<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// An `azure_blob_fs_location` block as defined below.
        /// </summary>
        [Input("azureBlobFsLocation")]
        public Input<Inputs.DatasetDelimitedTextAzureBlobFsLocationGetArgs>? AzureBlobFsLocation { get; set; }

        /// <summary>
        /// An `azure_blob_storage_location` block as defined below.
        /// </summary>
        [Input("azureBlobStorageLocation")]
        public Input<Inputs.DatasetDelimitedTextAzureBlobStorageLocationGetArgs>? AzureBlobStorageLocation { get; set; }

        /// <summary>
        /// The column delimiter. Defaults to `,`.
        /// </summary>
        [Input("columnDelimiter")]
        public Input<string>? ColumnDelimiter { get; set; }

        /// <summary>
        /// The compression codec used to read/write text files. Valid values are `bzip2`, `gzip`, `deflate`, `ZipDeflate`, `TarGzip`, `Tar`, `snappy`, or `lz4`. Please note these values are case sensitive.
        /// </summary>
        [Input("compressionCodec")]
        public Input<string>? CompressionCodec { get; set; }

        /// <summary>
        /// The compression ratio for the Data Factory Dataset. Valid values are `Fastest` or `Optimal`. Please note these values are case sensitive.
        /// </summary>
        [Input("compressionLevel")]
        public Input<string>? CompressionLevel { get; set; }

        /// <summary>
        /// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryName")]
        public Input<string>? DataFactoryName { get; set; }

        /// <summary>
        /// The description for the Data Factory Dataset.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The encoding format for the file.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The escape character. Defaults to `\`.
        /// </summary>
        [Input("escapeCharacter")]
        public Input<string>? EscapeCharacter { get; set; }

        /// <summary>
        /// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to `false`.
        /// </summary>
        [Input("firstRowAsHeader")]
        public Input<bool>? FirstRowAsHeader { get; set; }

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// A `http_server_location` block as defined below.
        /// </summary>
        [Input("httpServerLocation")]
        public Input<Inputs.DatasetDelimitedTextHttpServerLocationGetArgs>? HttpServerLocation { get; set; }

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Input("linkedServiceName")]
        public Input<string>? LinkedServiceName { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The null value string. Defaults to an empty string.
        /// </summary>
        [Input("nullValue")]
        public Input<string>? NullValue { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The quote character. Defaults to `"`.
        /// </summary>
        [Input("quoteCharacter")]
        public Input<string>? QuoteCharacter { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The row delimiter. Defaults to any of the following values on read: `\r\n`, `\r`, `\n`, and `\n` or `\r\n` on write by mapping data flow and Copy activity respectively.
        /// </summary>
        [Input("rowDelimiter")]
        public Input<string>? RowDelimiter { get; set; }

        [Input("schemaColumns")]
        private InputList<Inputs.DatasetDelimitedTextSchemaColumnGetArgs>? _schemaColumns;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        public InputList<Inputs.DatasetDelimitedTextSchemaColumnGetArgs> SchemaColumns
        {
            get => _schemaColumns ?? (_schemaColumns = new InputList<Inputs.DatasetDelimitedTextSchemaColumnGetArgs>());
            set => _schemaColumns = value;
        }

        public DatasetDelimitedTextState()
        {
        }
    }
}
