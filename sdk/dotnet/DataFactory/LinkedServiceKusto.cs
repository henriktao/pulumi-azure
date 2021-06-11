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
    /// Manages a Linked Service (connection) between a Kusto Cluster and Azure Data Factory.
    /// 
    /// ## Import
    /// 
    /// Data Factory Linked Service's can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/linkedServiceKusto:LinkedServiceKusto example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/linkedServiceKusto:LinkedServiceKusto")]
    public partial class LinkedServiceKusto : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of additional properties to associate with the Data Factory Linked Service.
        /// </summary>
        [Output("additionalProperties")]
        public Output<ImmutableDictionary<string, string>?> AdditionalProperties { get; private set; } = null!;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Linked Service.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Output("dataFactoryId")]
        public Output<string> DataFactoryId { get; private set; } = null!;

        /// <summary>
        /// The description for the Data Factory Linked Service.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The integration runtime reference to associate with the Data Factory Linked Service.
        /// </summary>
        [Output("integrationRuntimeName")]
        public Output<string?> IntegrationRuntimeName { get; private set; } = null!;

        /// <summary>
        /// The Kusto Database Name.
        /// </summary>
        [Output("kustoDatabaseName")]
        public Output<string> KustoDatabaseName { get; private set; } = null!;

        /// <summary>
        /// The URI of the Kusto Cluster endpoint.
        /// </summary>
        [Output("kustoEndpoint")]
        public Output<string> KustoEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        /// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Linked Service.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The service principal id in which to authenticate against the Kusto Database.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string?> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The service principal key in which to authenticate against the Kusto Database.
        /// </summary>
        [Output("servicePrincipalKey")]
        public Output<string?> ServicePrincipalKey { get; private set; } = null!;

        /// <summary>
        /// The service principal tenant id or name in which to authenticate against the Kusto Database.
        /// </summary>
        [Output("tenant")]
        public Output<string?> Tenant { get; private set; } = null!;

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
        /// </summary>
        [Output("useManagedIdentity")]
        public Output<bool?> UseManagedIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedServiceKusto resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedServiceKusto(string name, LinkedServiceKustoArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/linkedServiceKusto:LinkedServiceKusto", name, args ?? new LinkedServiceKustoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedServiceKusto(string name, Input<string> id, LinkedServiceKustoState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/linkedServiceKusto:LinkedServiceKusto", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LinkedServiceKusto resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkedServiceKusto Get(string name, Input<string> id, LinkedServiceKustoState? state = null, CustomResourceOptions? options = null)
        {
            return new LinkedServiceKusto(name, id, state, options);
        }
    }

    public sealed class LinkedServiceKustoArgs : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Linked Service.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Linked Service.
        /// </summary>
        public InputList<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        /// </summary>
        [Input("dataFactoryId", required: true)]
        public Input<string> DataFactoryId { get; set; } = null!;

        /// <summary>
        /// The description for the Data Factory Linked Service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The integration runtime reference to associate with the Data Factory Linked Service.
        /// </summary>
        [Input("integrationRuntimeName")]
        public Input<string>? IntegrationRuntimeName { get; set; }

        /// <summary>
        /// The Kusto Database Name.
        /// </summary>
        [Input("kustoDatabaseName", required: true)]
        public Input<string> KustoDatabaseName { get; set; } = null!;

        /// <summary>
        /// The URI of the Kusto Cluster endpoint.
        /// </summary>
        [Input("kustoEndpoint", required: true)]
        public Input<string> KustoEndpoint { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        /// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Linked Service.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The service principal id in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The service principal key in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("servicePrincipalKey")]
        public Input<string>? ServicePrincipalKey { get; set; }

        /// <summary>
        /// The service principal tenant id or name in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
        /// </summary>
        [Input("useManagedIdentity")]
        public Input<bool>? UseManagedIdentity { get; set; }

        public LinkedServiceKustoArgs()
        {
        }
    }

    public sealed class LinkedServiceKustoState : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// A map of additional properties to associate with the Data Factory Linked Service.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("annotations")]
        private InputList<string>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the Data Factory Linked Service.
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
        /// The description for the Data Factory Linked Service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The integration runtime reference to associate with the Data Factory Linked Service.
        /// </summary>
        [Input("integrationRuntimeName")]
        public Input<string>? IntegrationRuntimeName { get; set; }

        /// <summary>
        /// The Kusto Database Name.
        /// </summary>
        [Input("kustoDatabaseName")]
        public Input<string>? KustoDatabaseName { get; set; }

        /// <summary>
        /// The URI of the Kusto Cluster endpoint.
        /// </summary>
        [Input("kustoEndpoint")]
        public Input<string>? KustoEndpoint { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
        /// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Linked Service.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The service principal id in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The service principal key in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("servicePrincipalKey")]
        public Input<string>? ServicePrincipalKey { get; set; }

        /// <summary>
        /// The service principal tenant id or name in which to authenticate against the Kusto Database.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Kusto Database.
        /// </summary>
        [Input("useManagedIdentity")]
        public Input<bool>? UseManagedIdentity { get; set; }

        public LinkedServiceKustoState()
        {
        }
    }
}
