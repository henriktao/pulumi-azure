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
    /// Manages a Snowflake Dataset inside an Azure Data Factory.
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
    ///         var exampleLinkedServiceSnowflake = new Azure.DataFactory.LinkedServiceSnowflake("exampleLinkedServiceSnowflake", new Azure.DataFactory.LinkedServiceSnowflakeArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DataFactoryId = exampleFactory.Id,
    ///             ConnectionString = "jdbc:snowflake://account.region.snowflakecomputing.com/?user=user&amp;db=db&amp;warehouse=wh",
    ///         });
    ///         var exampleDatasetSnowflake = new Azure.DataFactory.DatasetSnowflake("exampleDatasetSnowflake", new Azure.DataFactory.DatasetSnowflakeArgs
    ///         {
    ///             ResourceGroupName = azurerm_resource_group.Test.Name,
    ///             DataFactoryId = azurerm_data_factory.Test.Id,
    ///             LinkedServiceName = azurerm_data_factory_linked_service_snowflake.Test.Name,
    ///             SchemaName = "foo_schema",
    ///             TableName = "foo_table",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Snowflake Datasets can be imported using the `resource id`,
    /// 
    /// e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/datasetSnowflake:DatasetSnowflake example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/datasetSnowflake:DatasetSnowflake")]
    public partial class DatasetSnowflake : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("additionalProperties")]
        public Output<ImmutableDictionary<string, string>?> AdditionalProperties { get; private set; } = null!;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Output("dataFactoryId")]
        public Output<string> DataFactoryId { get; private set; } = null!;

        /// <summary>
        /// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Output("dataFactoryName")]
        public Output<string> DataFactoryName { get; private set; } = null!;

        /// <summary>
        /// The description for the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Output("linkedServiceName")]
        public Output<string> LinkedServiceName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        [Output("schemaColumns")]
        public Output<ImmutableArray<Outputs.DatasetSnowflakeSchemaColumn>> SchemaColumns { get; private set; } = null!;

        /// <summary>
        /// The schema name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("schemaName")]
        public Output<string?> SchemaName { get; private set; } = null!;

        [Output("structureColumns")]
        public Output<ImmutableArray<Outputs.DatasetSnowflakeStructureColumn>> StructureColumns { get; private set; } = null!;

        /// <summary>
        /// The table name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetSnowflake resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetSnowflake(string name, DatasetSnowflakeArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/datasetSnowflake:DatasetSnowflake", name, args ?? new DatasetSnowflakeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetSnowflake(string name, Input<string> id, DatasetSnowflakeState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/datasetSnowflake:DatasetSnowflake", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetSnowflake resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetSnowflake Get(string name, Input<string> id, DatasetSnowflakeState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetSnowflake(name, id, state, options);
        }
    }

    public sealed class DatasetSnowflakeArgs : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset Snowflake.
        /// </summary>
        public InputList<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryId")]
        public Input<string>? DataFactoryId { get; set; }

        /// <summary>
        /// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryName")]
        public Input<string>? DataFactoryName { get; set; }

        /// <summary>
        /// The description for the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public Input<string> LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("schemaColumns")]
        private InputList<Inputs.DatasetSnowflakeSchemaColumnArgs>? _schemaColumns;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        public InputList<Inputs.DatasetSnowflakeSchemaColumnArgs> SchemaColumns
        {
            get => _schemaColumns ?? (_schemaColumns = new InputList<Inputs.DatasetSnowflakeSchemaColumnArgs>());
            set => _schemaColumns = value;
        }

        /// <summary>
        /// The schema name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        [Input("structureColumns")]
        private InputList<Inputs.DatasetSnowflakeStructureColumnArgs>? _structureColumns;
        [Obsolete(@"This block has been deprecated in favour of `schema_column` and will be removed.")]
        public InputList<Inputs.DatasetSnowflakeStructureColumnArgs> StructureColumns
        {
            get => _structureColumns ?? (_structureColumns = new InputList<Inputs.DatasetSnowflakeStructureColumnArgs>());
            set => _structureColumns = value;
        }

        /// <summary>
        /// The table name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public DatasetSnowflakeArgs()
        {
        }
    }

    public sealed class DatasetSnowflakeState : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Dataset Snowflake.
        /// </summary>
        public InputList<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryId")]
        public Input<string>? DataFactoryId { get; set; }

        /// <summary>
        /// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryName")]
        public Input<string>? DataFactoryName { get; set; }

        /// <summary>
        /// The description for the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The Data Factory Linked Service name in which to associate the Dataset with.
        /// </summary>
        [Input("linkedServiceName")]
        public Input<string>? LinkedServiceName { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory Dataset Snowflake. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Dataset Snowflake.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Dataset Snowflake. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("schemaColumns")]
        private InputList<Inputs.DatasetSnowflakeSchemaColumnGetArgs>? _schemaColumns;

        /// <summary>
        /// A `schema_column` block as defined below.
        /// </summary>
        public InputList<Inputs.DatasetSnowflakeSchemaColumnGetArgs> SchemaColumns
        {
            get => _schemaColumns ?? (_schemaColumns = new InputList<Inputs.DatasetSnowflakeSchemaColumnGetArgs>());
            set => _schemaColumns = value;
        }

        /// <summary>
        /// The schema name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        [Input("structureColumns")]
        private InputList<Inputs.DatasetSnowflakeStructureColumnGetArgs>? _structureColumns;
        [Obsolete(@"This block has been deprecated in favour of `schema_column` and will be removed.")]
        public InputList<Inputs.DatasetSnowflakeStructureColumnGetArgs> StructureColumns
        {
            get => _structureColumns ?? (_structureColumns = new InputList<Inputs.DatasetSnowflakeStructureColumnGetArgs>());
            set => _structureColumns = value;
        }

        /// <summary>
        /// The table name of the Data Factory Dataset Snowflake.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public DatasetSnowflakeState()
        {
        }
    }
}
