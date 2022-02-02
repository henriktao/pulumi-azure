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
    /// Manages a Linked Service (connection) between an Azure Blob Storage Account and Azure Data Factory.
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
    ///         var exampleAccount = Azure.Storage.GetAccount.Invoke(new Azure.Storage.GetAccountInvokeArgs
    ///         {
    ///             Name = "storageaccountname",
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleFactory = new Azure.DataFactory.Factory("exampleFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleLinkedServiceAzureBlobStorage = new Azure.DataFactory.LinkedServiceAzureBlobStorage("exampleLinkedServiceAzureBlobStorage", new Azure.DataFactory.LinkedServiceAzureBlobStorageArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             DataFactoryId = exampleFactory.Id,
    ///             ConnectionString = exampleAccount.Apply(exampleAccount =&gt; exampleAccount.PrimaryConnectionString),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With SAS Uri And SAS Token.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testResourceGroup = new Azure.Core.ResourceGroup("testResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var testFactory = new Azure.DataFactory.Factory("testFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = testResourceGroup.Location,
    ///             ResourceGroupName = testResourceGroup.Name,
    ///         });
    ///         var testKeyVault = new Azure.KeyVault.KeyVault("testKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = testResourceGroup.Location,
    ///             ResourceGroupName = testResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "standard",
    ///         });
    ///         var testLinkedServiceKeyVault = new Azure.DataFactory.LinkedServiceKeyVault("testLinkedServiceKeyVault", new Azure.DataFactory.LinkedServiceKeyVaultArgs
    ///         {
    ///             ResourceGroupName = testResourceGroup.Name,
    ///             DataFactoryId = testFactory.Id,
    ///             KeyVaultId = testKeyVault.Id,
    ///         });
    ///         var testLinkedServiceAzureBlobStorage = new Azure.DataFactory.LinkedServiceAzureBlobStorage("testLinkedServiceAzureBlobStorage", new Azure.DataFactory.LinkedServiceAzureBlobStorageArgs
    ///         {
    ///             ResourceGroupName = testResourceGroup.Name,
    ///             DataFactoryId = testFactory.Id,
    ///             SasUri = "https://storageaccountname.blob.core.windows.net",
    ///             KeyVaultSasToken = new Azure.DataFactory.Inputs.LinkedServiceAzureBlobStorageKeyVaultSasTokenArgs
    ///             {
    ///                 LinkedServiceName = testLinkedServiceKeyVault.Name,
    ///                 SecretName = "secret",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Linked Service's can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage")]
    public partial class LinkedServiceAzureBlobStorage : Pulumi.CustomResource
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
        /// The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        /// </summary>
        [Output("connectionString")]
        public Output<string?> ConnectionString { get; private set; } = null!;

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
        /// A `key_vault_sas_token` block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A `sas_uri` is required.
        /// </summary>
        [Output("keyVaultSasToken")]
        public Output<Outputs.LinkedServiceAzureBlobStorageKeyVaultSasToken?> KeyVaultSasToken { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of parameters to associate with the Data Factory Linked Service.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        /// </summary>
        [Output("sasUri")]
        public Output<string?> SasUri { get; private set; } = null!;

        /// <summary>
        /// The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        /// </summary>
        [Output("serviceEndpoint")]
        public Output<string?> ServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string?> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        /// </summary>
        [Output("servicePrincipalKey")]
        public Output<string?> ServicePrincipalKey { get; private set; } = null!;

        /// <summary>
        /// The tenant id or name in which to authenticate against the Azure Blob Storage account.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        /// </summary>
        [Output("useManagedIdentity")]
        public Output<bool?> UseManagedIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedServiceAzureBlobStorage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedServiceAzureBlobStorage(string name, LinkedServiceAzureBlobStorageArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage", name, args ?? new LinkedServiceAzureBlobStorageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedServiceAzureBlobStorage(string name, Input<string> id, LinkedServiceAzureBlobStorageState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LinkedServiceAzureBlobStorage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkedServiceAzureBlobStorage Get(string name, Input<string> id, LinkedServiceAzureBlobStorageState? state = null, CustomResourceOptions? options = null)
        {
            return new LinkedServiceAzureBlobStorage(name, id, state, options);
        }
    }

    public sealed class LinkedServiceAzureBlobStorageArgs : Pulumi.ResourceArgs
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
        /// The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

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
        /// A `key_vault_sas_token` block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A `sas_uri` is required.
        /// </summary>
        [Input("keyVaultSasToken")]
        public Input<Inputs.LinkedServiceAzureBlobStorageKeyVaultSasTokenArgs>? KeyVaultSasToken { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
        /// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        /// </summary>
        [Input("sasUri")]
        public Input<string>? SasUri { get; set; }

        /// <summary>
        /// The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        /// </summary>
        [Input("serviceEndpoint")]
        public Input<string>? ServiceEndpoint { get; set; }

        /// <summary>
        /// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        /// </summary>
        [Input("servicePrincipalKey")]
        public Input<string>? ServicePrincipalKey { get; set; }

        /// <summary>
        /// The tenant id or name in which to authenticate against the Azure Blob Storage account.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        /// </summary>
        [Input("useManagedIdentity")]
        public Input<bool>? UseManagedIdentity { get; set; }

        public LinkedServiceAzureBlobStorageArgs()
        {
        }
    }

    public sealed class LinkedServiceAzureBlobStorageState : Pulumi.ResourceArgs
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
        /// The connection string. Conflicts with `sas_uri` and `service_endpoint`.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

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
        /// A `key_vault_sas_token` block as defined below. Use this argument to store SAS Token in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. A `sas_uri` is required.
        /// </summary>
        [Input("keyVaultSasToken")]
        public Input<Inputs.LinkedServiceAzureBlobStorageKeyVaultSasTokenGetArgs>? KeyVaultSasToken { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
        /// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The SAS URI. Conflicts with `connection_string` and `service_endpoint`.
        /// </summary>
        [Input("sasUri")]
        public Input<string>? SasUri { get; set; }

        /// <summary>
        /// The Service Endpoint. Conflicts with `connection_string` and `sas_uri`. Required with `use_managed_identity`.
        /// </summary>
        [Input("serviceEndpoint")]
        public Input<string>? ServiceEndpoint { get; set; }

        /// <summary>
        /// The service principal id in which to authenticate against the Azure Blob Storage account. Required if `service_principal_key` is set.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The service principal key in which to authenticate against the AAzure Blob Storage account.  Required if `service_principal_id` is set.
        /// </summary>
        [Input("servicePrincipalKey")]
        public Input<string>? ServicePrincipalKey { get; set; }

        /// <summary>
        /// The tenant id or name in which to authenticate against the Azure Blob Storage account.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Whether to use the Data Factory's managed identity to authenticate against the Azure Blob Storage account. Incompatible with `service_principal_id` and `service_principal_key`.
        /// </summary>
        [Input("useManagedIdentity")]
        public Input<bool>? UseManagedIdentity { get; set; }

        public LinkedServiceAzureBlobStorageState()
        {
        }
    }
}
