// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class AppServiceSiteConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Are Managed Identity Credentials used for Azure Container Registry pull
        /// </summary>
        [Input("acrUseManagedIdentityCredentials")]
        public Input<bool>? AcrUseManagedIdentityCredentials { get; set; }

        /// <summary>
        /// If using User Managed Identity, the User Managed Identity Client Id
        /// </summary>
        [Input("acrUserManagedIdentityClientId")]
        public Input<string>? AcrUserManagedIdentityClientId { get; set; }

        /// <summary>
        /// Should the app be loaded at all times? Defaults to `false`.
        /// </summary>
        [Input("alwaysOn")]
        public Input<bool>? AlwaysOn { get; set; }

        /// <summary>
        /// App command line to launch, e.g. `/sbin/myserver -b 0.0.0.0`.
        /// </summary>
        [Input("appCommandLine")]
        public Input<string>? AppCommandLine { get; set; }

        [Input("autoSwapSlotName")]
        public Input<string>? AutoSwapSlotName { get; set; }

        /// <summary>
        /// A `cors` block as defined below.
        /// </summary>
        [Input("cors")]
        public Input<Inputs.AppServiceSiteConfigCorsArgs>? Cors { get; set; }

        [Input("defaultDocuments")]
        private InputList<string>? _defaultDocuments;

        /// <summary>
        /// The ordering of default documents to load, if an address isn't specified.
        /// </summary>
        public InputList<string> DefaultDocuments
        {
            get => _defaultDocuments ?? (_defaultDocuments = new InputList<string>());
            set => _defaultDocuments = value;
        }

        /// <summary>
        /// The version of the .net framework's CLR used in this App Service. Possible values are `v2.0` (which will use the latest version of the .net framework for the .net CLR v2 - currently `.net 3.5`), `v4.0` (which corresponds to the latest version of the .net CLR v4 - which at the time of writing is `.net 4.7.1`), `v5.0` and `v6.0`. [For more information on which .net CLR version to use based on the .net framework you're targeting - please see this table](https://en.wikipedia.org/wiki/.NET_Framework_version_history#Overview). Defaults to `v4.0`.
        /// </summary>
        [Input("dotnetFrameworkVersion")]
        public Input<string>? DotnetFrameworkVersion { get; set; }

        /// <summary>
        /// State of FTP / FTPS service for this App Service. Possible values include: `AllAllowed`, `FtpsOnly` and `Disabled`.
        /// </summary>
        [Input("ftpsState")]
        public Input<string>? FtpsState { get; set; }

        /// <summary>
        /// The health check path to be pinged by App Service. [For more information - please see App Service health check announcement](https://azure.github.io/AppService/2020/08/24/healthcheck-on-app-service.html).
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// Is HTTP2 Enabled on this App Service? Defaults to `false`.
        /// </summary>
        [Input("http2Enabled")]
        public Input<bool>? Http2Enabled { get; set; }

        [Input("ipRestrictions")]
        private InputList<Inputs.AppServiceSiteConfigIpRestrictionArgs>? _ipRestrictions;

        /// <summary>
        /// A list of objects representing ip restrictions as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceSiteConfigIpRestrictionArgs> IpRestrictions
        {
            get => _ipRestrictions ?? (_ipRestrictions = new InputList<Inputs.AppServiceSiteConfigIpRestrictionArgs>());
            set => _ipRestrictions = value;
        }

        /// <summary>
        /// The Java Container to use. If specified `java_version` and `java_container_version` must also be specified. Possible values are `JAVA`, `JETTY`, and `TOMCAT`.
        /// </summary>
        [Input("javaContainer")]
        public Input<string>? JavaContainer { get; set; }

        /// <summary>
        /// The version of the Java Container to use. If specified `java_version` and `java_container` must also be specified.
        /// </summary>
        [Input("javaContainerVersion")]
        public Input<string>? JavaContainerVersion { get; set; }

        /// <summary>
        /// The version of Java to use. If specified `java_container` and `java_container_version` must also be specified. Possible values are `1.7`, `1.8` and `11` and their specific versions - except for Java 11 (e.g. `1.7.0_80`, `1.8.0_181`, `11`)
        /// </summary>
        [Input("javaVersion")]
        public Input<string>? JavaVersion { get; set; }

        /// <summary>
        /// Linux App Framework and version for the App Service. Possible options are a Docker container (`DOCKER|&lt;user/image:tag&gt;`), a base-64 encoded Docker Compose file (`COMPOSE|${filebase64("compose.yml")}`) or a base-64 encoded Kubernetes Manifest (`KUBE|${filebase64("kubernetes.yml")}`).
        /// </summary>
        [Input("linuxFxVersion")]
        public Input<string>? LinuxFxVersion { get; set; }

        /// <summary>
        /// Is "MySQL In App" Enabled? This runs a local MySQL instance with your app and shares resources from the App Service plan.
        /// </summary>
        [Input("localMysqlEnabled")]
        public Input<bool>? LocalMysqlEnabled { get; set; }

        /// <summary>
        /// The Managed Pipeline Mode. Possible values are `Integrated` and `Classic`. Defaults to `Integrated`.
        /// </summary>
        [Input("managedPipelineMode")]
        public Input<string>? ManagedPipelineMode { get; set; }

        /// <summary>
        /// The minimum supported TLS version for the app service. Possible values are `1.0`, `1.1`, and `1.2`. Defaults to `1.2` for new app services.
        /// </summary>
        [Input("minTlsVersion")]
        public Input<string>? MinTlsVersion { get; set; }

        /// <summary>
        /// The scaled number of workers (for per site scaling) of this App Service. Requires that `per_site_scaling` is enabled on the `azure.appservice.Plan`. [For more information - please see Microsoft documentation on high-density hosting](https://docs.microsoft.com/en-us/azure/app-service/manage-scale-per-app).
        /// </summary>
        [Input("numberOfWorkers")]
        public Input<int>? NumberOfWorkers { get; set; }

        /// <summary>
        /// The version of PHP to use in this App Service. Possible values are `5.5`, `5.6`, `7.0`, `7.1`, `7.2`, `7.3` and `7.4`.
        /// </summary>
        [Input("phpVersion")]
        public Input<string>? PhpVersion { get; set; }

        /// <summary>
        /// The version of Python to use in this App Service. Possible values are `2.7` and `3.4`.
        /// </summary>
        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        /// <summary>
        /// Is Remote Debugging Enabled? Defaults to `false`.
        /// </summary>
        [Input("remoteDebuggingEnabled")]
        public Input<bool>? RemoteDebuggingEnabled { get; set; }

        /// <summary>
        /// Which version of Visual Studio should the Remote Debugger be compatible with? Possible values are `VS2012`, `VS2013`, `VS2015` and `VS2017`.
        /// </summary>
        [Input("remoteDebuggingVersion")]
        public Input<string>? RemoteDebuggingVersion { get; set; }

        [Input("scmIpRestrictions")]
        private InputList<Inputs.AppServiceSiteConfigScmIpRestrictionArgs>? _scmIpRestrictions;

        /// <summary>
        /// A [List of objects](https://www.terraform.io/docs/configuration/attr-as-blocks.html) representing ip restrictions as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceSiteConfigScmIpRestrictionArgs> ScmIpRestrictions
        {
            get => _scmIpRestrictions ?? (_scmIpRestrictions = new InputList<Inputs.AppServiceSiteConfigScmIpRestrictionArgs>());
            set => _scmIpRestrictions = value;
        }

        /// <summary>
        /// The type of Source Control enabled for this App Service. Defaults to `None`. Possible values are: `BitbucketGit`, `BitbucketHg`, `CodePlexGit`, `CodePlexHg`, `Dropbox`, `ExternalGit`, `ExternalHg`, `GitHub`, `LocalGit`, `None`, `OneDrive`, `Tfs`, `VSO`, and `VSTSRM`
        /// </summary>
        [Input("scmType")]
        public Input<string>? ScmType { get; set; }

        /// <summary>
        /// IP security restrictions for scm to use main. Defaults to false.
        /// </summary>
        [Input("scmUseMainIpRestriction")]
        public Input<bool>? ScmUseMainIpRestriction { get; set; }

        /// <summary>
        /// Should the App Service run in 32 bit mode, rather than 64 bit mode?
        /// </summary>
        [Input("use32BitWorkerProcess")]
        public Input<bool>? Use32BitWorkerProcess { get; set; }

        /// <summary>
        /// Should all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied.
        /// </summary>
        [Input("vnetRouteAllEnabled")]
        public Input<bool>? VnetRouteAllEnabled { get; set; }

        /// <summary>
        /// Should WebSockets be enabled?
        /// </summary>
        [Input("websocketsEnabled")]
        public Input<bool>? WebsocketsEnabled { get; set; }

        /// <summary>
        /// The Windows Docker container image (`DOCKER|&lt;user/image:tag&gt;`)
        /// </summary>
        [Input("windowsFxVersion")]
        public Input<string>? WindowsFxVersion { get; set; }

        public AppServiceSiteConfigArgs()
        {
        }
    }
}
