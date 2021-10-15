// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Manages a Function App.
    /// 
    /// &gt; **Note:** To connect an Azure Function App and a subnet within the same region `azure.appservice.VirtualNetworkSwiftConnection` can be used.
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
    ///                 Tier = "Standard",
    ///                 Size = "S1",
    ///             },
    ///         });
    ///         var exampleFunctionApp = new Azure.AppService.FunctionApp("exampleFunctionApp", new Azure.AppService.FunctionAppArgs
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
    /// ### In A Consumption Plan)
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
    ///             Kind = "FunctionApp",
    ///             Sku = new Azure.AppService.Inputs.PlanSkuArgs
    ///             {
    ///                 Tier = "Dynamic",
    ///                 Size = "Y1",
    ///             },
    ///         });
    ///         var exampleFunctionApp = new Azure.AppService.FunctionApp("exampleFunctionApp", new Azure.AppService.FunctionAppArgs
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
    /// ### Linux)
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
    ///                 Tier = "Dynamic",
    ///                 Size = "Y1",
    ///             },
    ///         });
    ///         var exampleFunctionApp = new Azure.AppService.FunctionApp("exampleFunctionApp", new Azure.AppService.FunctionAppArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///             StorageAccountName = exampleAccount.Name,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///             OsType = "linux",
    ///             Version = "~3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// &gt; **Note:** Version `~3` is required for Linux Function Apps.
    /// 
    /// ## Import
    /// 
    /// Function Apps can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:appservice/functionApp:FunctionApp functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1
    /// ```
    /// </summary>
    [AzureResourceType("azure:appservice/functionApp:FunctionApp")]
    public partial class FunctionApp : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Function App.
        /// </summary>
        [Output("appServicePlanId")]
        public Output<string> AppServicePlanId { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
        /// </summary>
        [Output("appSettings")]
        public Output<ImmutableDictionary<string, string>> AppSettings { get; private set; } = null!;

        /// <summary>
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Output("authSettings")]
        public Output<Outputs.FunctionAppAuthSettings> AuthSettings { get; private set; } = null!;

        /// <summary>
        /// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Output("clientAffinityEnabled")]
        public Output<bool> ClientAffinityEnabled { get; private set; } = null!;

        /// <summary>
        /// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Output("clientCertMode")]
        public Output<string?> ClientCertMode { get; private set; } = null!;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        [Output("connectionStrings")]
        public Output<ImmutableArray<Outputs.FunctionAppConnectionString>> ConnectionStrings { get; private set; } = null!;

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Output("customDomainVerificationId")]
        public Output<string> CustomDomainVerificationId { get; private set; } = null!;

        /// <summary>
        /// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        /// </summary>
        [Output("dailyMemoryTimeQuota")]
        public Output<int?> DailyMemoryTimeQuota { get; private set; } = null!;

        /// <summary>
        /// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
        /// </summary>
        [Output("defaultHostname")]
        public Output<string> DefaultHostname { get; private set; } = null!;

        /// <summary>
        /// Should the built-in logging of this Function App be enabled? Defaults to `true`.
        /// </summary>
        [Output("enableBuiltinLogging")]
        public Output<bool?> EnableBuiltinLogging { get; private set; } = null!;

        /// <summary>
        /// Is the Function App enabled?
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Can the Function App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Output("httpsOnly")]
        public Output<bool?> HttpsOnly { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.FunctionAppIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// The Function App kind - such as `functionapp,linux,container`
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Function App. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A string indicating the Operating System type for this function app.
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

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
        /// The name of the resource group in which to create the Function App.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Output("siteConfig")]
        public Output<Outputs.FunctionAppSiteConfig> SiteConfig { get; private set; } = null!;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        [Output("siteCredentials")]
        public Output<ImmutableArray<Outputs.FunctionAppSiteCredential>> SiteCredentials { get; private set; } = null!;

        /// <summary>
        /// A `source_control` block, as defined below.
        /// </summary>
        [Output("sourceControl")]
        public Output<Outputs.FunctionAppSourceControl> SourceControl { get; private set; } = null!;

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Function App.
        /// </summary>
        [Output("storageAccountAccessKey")]
        public Output<string> StorageAccountAccessKey { get; private set; } = null!;

        /// <summary>
        /// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        [Output("storageConnectionString")]
        public Output<string> StorageConnectionString { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The runtime version associated with the Function App. Defaults to `~1`.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionApp(string name, FunctionAppArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/functionApp:FunctionApp", name, args ?? new FunctionAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionApp(string name, Input<string> id, FunctionAppState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/functionApp:FunctionApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionApp Get(string name, Input<string> id, FunctionAppState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionApp(name, id, state, options);
        }
    }

    public sealed class FunctionAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Function App.
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
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Input("authSettings")]
        public Input<Inputs.FunctionAppAuthSettingsArgs>? AuthSettings { get; set; }

        /// <summary>
        /// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Input("clientCertMode")]
        public Input<string>? ClientCertMode { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.FunctionAppConnectionStringArgs>? _connectionStrings;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public InputList<Inputs.FunctionAppConnectionStringArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.FunctionAppConnectionStringArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        /// </summary>
        [Input("dailyMemoryTimeQuota")]
        public Input<int>? DailyMemoryTimeQuota { get; set; }

        /// <summary>
        /// Should the built-in logging of this Function App be enabled? Defaults to `true`.
        /// </summary>
        [Input("enableBuiltinLogging")]
        public Input<bool>? EnableBuiltinLogging { get; set; }

        /// <summary>
        /// Is the Function App enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the Function App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.FunctionAppIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Function App. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string indicating the Operating System type for this function app.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Function App.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.FunctionAppSiteConfigArgs>? SiteConfig { get; set; }

        /// <summary>
        /// A `source_control` block, as defined below.
        /// </summary>
        [Input("sourceControl")]
        public Input<Inputs.FunctionAppSourceControlArgs>? SourceControl { get; set; }

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Function App.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        [Input("storageConnectionString")]
        public Input<string>? StorageConnectionString { get; set; }

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
        /// The runtime version associated with the Function App. Defaults to `~1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public FunctionAppArgs()
        {
        }
    }

    public sealed class FunctionAppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this Function App.
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
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Input("authSettings")]
        public Input<Inputs.FunctionAppAuthSettingsGetArgs>? AuthSettings { get; set; }

        /// <summary>
        /// Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// The mode of the Function App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
        /// </summary>
        [Input("clientCertMode")]
        public Input<string>? ClientCertMode { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.FunctionAppConnectionStringGetArgs>? _connectionStrings;

        /// <summary>
        /// An `connection_string` block as defined below.
        /// </summary>
        public InputList<Inputs.FunctionAppConnectionStringGetArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.FunctionAppConnectionStringGetArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Input("customDomainVerificationId")]
        public Input<string>? CustomDomainVerificationId { get; set; }

        /// <summary>
        /// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
        /// </summary>
        [Input("dailyMemoryTimeQuota")]
        public Input<int>? DailyMemoryTimeQuota { get; set; }

        /// <summary>
        /// The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
        /// </summary>
        [Input("defaultHostname")]
        public Input<string>? DefaultHostname { get; set; }

        /// <summary>
        /// Should the built-in logging of this Function App be enabled? Defaults to `true`.
        /// </summary>
        [Input("enableBuiltinLogging")]
        public Input<bool>? EnableBuiltinLogging { get; set; }

        /// <summary>
        /// Is the Function App enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the Function App only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.FunctionAppIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Function App kind - such as `functionapp,linux,container`
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Function App. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string indicating the Operating System type for this function app.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

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
        /// The name of the resource group in which to create the Function App.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `site_config` object as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.FunctionAppSiteConfigGetArgs>? SiteConfig { get; set; }

        [Input("siteCredentials")]
        private InputList<Inputs.FunctionAppSiteCredentialGetArgs>? _siteCredentials;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        public InputList<Inputs.FunctionAppSiteCredentialGetArgs> SiteCredentials
        {
            get => _siteCredentials ?? (_siteCredentials = new InputList<Inputs.FunctionAppSiteCredentialGetArgs>());
            set => _siteCredentials = value;
        }

        /// <summary>
        /// A `source_control` block, as defined below.
        /// </summary>
        [Input("sourceControl")]
        public Input<Inputs.FunctionAppSourceControlGetArgs>? SourceControl { get; set; }

        /// <summary>
        /// The access key which will be used to access the backend storage account for the Function App.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// The backend storage account name which will be used by this Function App (such as the dashboard, logs).
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        [Input("storageConnectionString")]
        public Input<string>? StorageConnectionString { get; set; }

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
        /// The runtime version associated with the Function App. Defaults to `~1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public FunctionAppState()
        {
        }
    }
}
