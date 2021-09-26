// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class FunctionAppSlotSiteConfig
    {
        /// <summary>
        /// Should the Function App be loaded at all times? Defaults to `false`.
        /// </summary>
        public readonly bool? AlwaysOn;
        /// <summary>
        /// The number of workers this function app can scale out to. Only applicable to apps on the Consumption and Premium plan.
        /// </summary>
        public readonly int? AppScaleLimit;
        /// <summary>
        /// The name of the slot to automatically swap to during deployment
        /// </summary>
        public readonly string? AutoSwapSlotName;
        /// <summary>
        /// A `cors` block as defined below.
        /// </summary>
        public readonly Outputs.FunctionAppSlotSiteConfigCors? Cors;
        /// <summary>
        /// The version of the .net framework's CLR used in this function app. Possible values are `v4.0` (including .NET Core 2.1 and 3.1), `v5.0` and `v6.0`. [For more information on which .net Framework version to use based on the runtime version you're targeting - please see this table](https://docs.microsoft.com/en-us/azure/azure-functions/functions-dotnet-class-library#supported-versions). Defaults to `v4.0`.
        /// </summary>
        public readonly string? DotnetFrameworkVersion;
        /// <summary>
        /// The number of minimum instances for this function app. Only applicable to apps on the Premium plan.
        /// </summary>
        public readonly int? ElasticInstanceMinimum;
        /// <summary>
        /// State of FTP / FTPS service for this function app. Possible values include: `AllAllowed`, `FtpsOnly` and `Disabled`.
        /// </summary>
        public readonly string? FtpsState;
        public readonly string? HealthCheckPath;
        /// <summary>
        /// Specifies whether or not the http2 protocol should be enabled. Defaults to `false`.
        /// </summary>
        public readonly bool? Http2Enabled;
        /// <summary>
        /// A [List of objects](https://www.terraform.io/docs/configuration/attr-as-blocks.html) representing ip restrictions as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionAppSlotSiteConfigIpRestriction> IpRestrictions;
        public readonly string? JavaVersion;
        /// <summary>
        /// Linux App Framework and version for the AppService, e.g. `DOCKER|(golang:latest)`.
        /// </summary>
        public readonly string? LinuxFxVersion;
        /// <summary>
        /// The minimum supported TLS version for the function app. Possible values are `1.0`, `1.1`, and `1.2`. Defaults to `1.2` for new function apps.
        /// </summary>
        public readonly string? MinTlsVersion;
        /// <summary>
        /// The number of pre-warmed instances for this function app. Only affects apps on the Premium plan.
        /// </summary>
        public readonly int? PreWarmedInstanceCount;
        /// <summary>
        /// Should Runtime Scale Monitoring be enabled?. Only applicable to apps on the Premium plan. Defaults to `false`.
        /// </summary>
        public readonly bool? RuntimeScaleMonitoringEnabled;
        public readonly ImmutableArray<Outputs.FunctionAppSlotSiteConfigScmIpRestriction> ScmIpRestrictions;
        public readonly string? ScmType;
        public readonly bool? ScmUseMainIpRestriction;
        /// <summary>
        /// Should the Function App run in 32 bit mode, rather than 64 bit mode? Defaults to `true`.
        /// </summary>
        public readonly bool? Use32BitWorkerProcess;
        public readonly bool? VnetRouteAllEnabled;
        /// <summary>
        /// Should WebSockets be enabled?
        /// </summary>
        public readonly bool? WebsocketsEnabled;

        [OutputConstructor]
        private FunctionAppSlotSiteConfig(
            bool? alwaysOn,

            int? appScaleLimit,

            string? autoSwapSlotName,

            Outputs.FunctionAppSlotSiteConfigCors? cors,

            string? dotnetFrameworkVersion,

            int? elasticInstanceMinimum,

            string? ftpsState,

            string? healthCheckPath,

            bool? http2Enabled,

            ImmutableArray<Outputs.FunctionAppSlotSiteConfigIpRestriction> ipRestrictions,

            string? javaVersion,

            string? linuxFxVersion,

            string? minTlsVersion,

            int? preWarmedInstanceCount,

            bool? runtimeScaleMonitoringEnabled,

            ImmutableArray<Outputs.FunctionAppSlotSiteConfigScmIpRestriction> scmIpRestrictions,

            string? scmType,

            bool? scmUseMainIpRestriction,

            bool? use32BitWorkerProcess,

            bool? vnetRouteAllEnabled,

            bool? websocketsEnabled)
        {
            AlwaysOn = alwaysOn;
            AppScaleLimit = appScaleLimit;
            AutoSwapSlotName = autoSwapSlotName;
            Cors = cors;
            DotnetFrameworkVersion = dotnetFrameworkVersion;
            ElasticInstanceMinimum = elasticInstanceMinimum;
            FtpsState = ftpsState;
            HealthCheckPath = healthCheckPath;
            Http2Enabled = http2Enabled;
            IpRestrictions = ipRestrictions;
            JavaVersion = javaVersion;
            LinuxFxVersion = linuxFxVersion;
            MinTlsVersion = minTlsVersion;
            PreWarmedInstanceCount = preWarmedInstanceCount;
            RuntimeScaleMonitoringEnabled = runtimeScaleMonitoringEnabled;
            ScmIpRestrictions = scmIpRestrictions;
            ScmType = scmType;
            ScmUseMainIpRestriction = scmUseMainIpRestriction;
            Use32BitWorkerProcess = use32BitWorkerProcess;
            VnetRouteAllEnabled = vnetRouteAllEnabled;
            WebsocketsEnabled = websocketsEnabled;
        }
    }
}
