// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages a Logic App (Standard / Single Tenant)
    /// 
    /// &gt; **Note:** To connect an Azure Logic App and a subnet within the same region `azure.appservice.VirtualNetworkSwiftConnection` can be used.
    /// For an example, check the `azure.appservice.VirtualNetworkSwiftConnection` documentation.
    /// 
    /// ## Example Usage
    /// ### With App Service Plan)
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var examplePlan = new Azure.AppService.Plan("examplePlan", new Azure.AppService.PlanArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = new Azure.AppService.Inputs.PlanSkuArgs
    ///             {
    ///                 Tier = "WorkflowStandard",
    ///                 Size = "WS1",
    ///             },
    ///         });
    ///         var exampleStandard = new Azure.LogicApps.Standard("exampleStandard", new Azure.LogicApps.StandardArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///             StorageAccountName = exampleAccount.Name,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### For Container Mode)
    /// 
    /// &gt; **Note:** You must set `azure.appservice.Plan` `kind` to `Linux` and `reserved` to `true` when used with `linux_fx_version`
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var examplePlan = new Azure.AppService.Plan("examplePlan", new Azure.AppService.PlanArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Kind = "Linux",
    ///             Reserved = true,
    ///             Sku = new Azure.AppService.Inputs.PlanSkuArgs
    ///             {
    ///                 Tier = "WorkflowStandard",
    ///                 Size = "WS1",
    ///             },
    ///         });
    ///         var exampleStandard = new Azure.LogicApps.Standard("exampleStandard", new Azure.LogicApps.StandardArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///             StorageAccountName = exampleAccount.Name,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///             SiteConfig = new Azure.LogicApps.Inputs.StandardSiteConfigArgs
    ///             {
    ///                 LinuxFxVersion = "DOCKER|mcr.microsoft.com/azure-functions/dotnet:3.0-appservice",
    ///             },
    ///             AppSettings = 
    ///             {
    ///                 { "DOCKER_REGISTRY_SERVER_URL", "https://&lt;server-name&gt;.azurecr.io" },
    ///                 { "DOCKER_REGISTRY_SERVER_USERNAME", "username" },
    ///                 { "DOCKER_REGISTRY_SERVER_PASSWORD", "password" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Logic Apps can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:logicapps/standard:Standard logicapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/logicapp1
    /// ```
    /// </summary>
    [AzureResourceType("azure:logicapps/standard:Standard")]
    public partial class Standard : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Logic App
        /// </summary>
        [Output("appServicePlanId")]
        public Output<string> AppServicePlanId { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
        /// </summary>
        [Output("appSettings")]
        public Output<ImmutableDictionary<string, string>> AppSettings { get; private set; } = null!;

        /// <summary>
        /// If `use_extension_bundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
        /// </summary>
        [Output("bundleVersion")]
        public Output<string?> BundleVersion { get; private set; } = null!;

        /// <summary>
        /// Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Output("clientAffinityEnabled")]
        public Output<bool> ClientAffinityEnabled { get; private set; } = null!;

        /// <summary>
        /// The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Output("clientCertificateMode")]
        public Output<string?> ClientCertificateMode { get; private set; } = null!;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        [Output("connectionStrings")]
        public Output<ImmutableArray<Outputs.StandardConnectionString>> ConnectionStrings { get; private set; } = null!;

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Output("customDomainVerificationId")]
        public Output<string> CustomDomainVerificationId { get; private set; } = null!;

        /// <summary>
        /// The default hostname associated with the Logic App - such as `mysite.azurewebsites.net`
        /// </summary>
        [Output("defaultHostname")]
        public Output<string> DefaultHostname { get; private set; } = null!;

        /// <summary>
        /// Is the Logic App enabled?
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Can the Logic App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Output("httpsOnly")]
        public Output<bool?> HttpsOnly { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.StandardIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Logic App kind - will be `functionapp,workflowapp`
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Logic App Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        /// </summary>
        [Output("outboundIpAddresses")]
        public Output<string> OutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        [Output("possibleOutboundIpAddresses")]
        public Output<string> PossibleOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Logic App
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Output("siteConfig")]
        public Output<Outputs.StandardSiteConfig> SiteConfig { get; private set; } = null!;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        [Output("siteCredentials")]
        public Output<ImmutableArray<Outputs.StandardSiteCredential>> SiteCredentials { get; private set; } = null!;

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Logic App
        /// </summary>
        [Output("storageAccountAccessKey")]
        public Output<string> StorageAccountAccessKey { get; private set; } = null!;

        /// <summary>
        /// The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
        /// </summary>
        [Output("storageAccountShareName")]
        public Output<string> StorageAccountShareName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
        /// </summary>
        [Output("useExtensionBundle")]
        public Output<bool?> UseExtensionBundle { get; private set; } = null!;

        /// <summary>
        /// The runtime version associated with the Logic App Defaults to `~1`.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Standard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Standard(string name, StandardArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/standard:Standard", name, args ?? new StandardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Standard(string name, Input<string> id, StandardState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/standard:Standard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Standard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Standard Get(string name, Input<string> id, StandardState? state = null, CustomResourceOptions? options = null)
        {
            return new Standard(name, id, state, options);
        }
    }

    public sealed class StandardArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Logic App
        /// </summary>
        [Input("appServicePlanId", required: true)]
        public Input<string> AppServicePlanId { get; set; } = null!;

        [Input("appSettings")]
        private InputMap<string>? _appSettings;

        /// <summary>
        /// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
        /// </summary>
        public InputMap<string> AppSettings
        {
            get => _appSettings ?? (_appSettings = new InputMap<string>());
            set => _appSettings = value;
        }

        /// <summary>
        /// If `use_extension_bundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
        /// </summary>
        [Input("bundleVersion")]
        public Input<string>? BundleVersion { get; set; }

        /// <summary>
        /// Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Input("clientCertificateMode")]
        public Input<string>? ClientCertificateMode { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.StandardConnectionStringArgs>? _connectionStrings;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public InputList<Inputs.StandardConnectionStringArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.StandardConnectionStringArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// Is the Logic App enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the Logic App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.StandardIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Logic App Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Logic App
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.StandardSiteConfigArgs>? SiteConfig { get; set; }

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Logic App
        /// </summary>
        [Input("storageAccountAccessKey", required: true)]
        public Input<string> StorageAccountAccessKey { get; set; } = null!;

        /// <summary>
        /// The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
        /// </summary>
        [Input("storageAccountShareName")]
        public Input<string>? StorageAccountShareName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
        /// </summary>
        [Input("useExtensionBundle")]
        public Input<bool>? UseExtensionBundle { get; set; }

        /// <summary>
        /// The runtime version associated with the Logic App Defaults to `~1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public StandardArgs()
        {
        }
    }

    public sealed class StandardState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Logic App
        /// </summary>
        [Input("appServicePlanId")]
        public Input<string>? AppServicePlanId { get; set; }

        [Input("appSettings")]
        private InputMap<string>? _appSettings;

        /// <summary>
        /// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
        /// </summary>
        public InputMap<string> AppSettings
        {
            get => _appSettings ?? (_appSettings = new InputMap<string>());
            set => _appSettings = value;
        }

        /// <summary>
        /// If `use_extension_bundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
        /// </summary>
        [Input("bundleVersion")]
        public Input<string>? BundleVersion { get; set; }

        /// <summary>
        /// Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Input("clientCertificateMode")]
        public Input<string>? ClientCertificateMode { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.StandardConnectionStringGetArgs>? _connectionStrings;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public InputList<Inputs.StandardConnectionStringGetArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.StandardConnectionStringGetArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Input("customDomainVerificationId")]
        public Input<string>? CustomDomainVerificationId { get; set; }

        /// <summary>
        /// The default hostname associated with the Logic App - such as `mysite.azurewebsites.net`
        /// </summary>
        [Input("defaultHostname")]
        public Input<string>? DefaultHostname { get; set; }

        /// <summary>
        /// Is the Logic App enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the Logic App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.StandardIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Logic App kind - will be `functionapp,workflowapp`
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Logic App Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        /// </summary>
        [Input("outboundIpAddresses")]
        public Input<string>? OutboundIpAddresses { get; set; }

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        [Input("possibleOutboundIpAddresses")]
        public Input<string>? PossibleOutboundIpAddresses { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Logic App
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.StandardSiteConfigGetArgs>? SiteConfig { get; set; }

        [Input("siteCredentials")]
        private InputList<Inputs.StandardSiteCredentialGetArgs>? _siteCredentials;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        public InputList<Inputs.StandardSiteCredentialGetArgs> SiteCredentials
        {
            get => _siteCredentials ?? (_siteCredentials = new InputList<Inputs.StandardSiteCredentialGetArgs>());
            set => _siteCredentials = value;
        }

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Logic App
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        /// <summary>
        /// The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
        /// </summary>
        [Input("storageAccountShareName")]
        public Input<string>? StorageAccountShareName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
        /// </summary>
        [Input("useExtensionBundle")]
        public Input<bool>? UseExtensionBundle { get; set; }

        /// <summary>
        /// The runtime version associated with the Logic App Defaults to `~1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public StandardState()
        {
        }
    }
}
