// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class FunctionAppSlotSiteConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the Function App be loaded at all times? Defaults to `false`.
        /// </summary>
        [Input("alwaysOn")]
        public Input<bool>? AlwaysOn { get; set; }

        /// <summary>
        /// The number of workers this function app can scale out to. Only applicable to apps on the Consumption and Premium plan.
        /// </summary>
        [Input("appScaleLimit")]
        public Input<int>? AppScaleLimit { get; set; }

        /// <summary>
        /// The name of the slot to automatically swap to during deployment
        /// </summary>
        [Input("autoSwapSlotName")]
        public Input<string>? AutoSwapSlotName { get; set; }

        /// <summary>
        /// A `cors` block as defined below.
        /// </summary>
        [Input("cors")]
        public Input<Inputs.FunctionAppSlotSiteConfigCorsGetArgs>? Cors { get; set; }

        /// <summary>
        /// The number of minimum instances for this function app. Only applicable to apps on the Premium plan.
        /// </summary>
        [Input("elasticInstanceMinimum")]
        public Input<int>? ElasticInstanceMinimum { get; set; }

        /// <summary>
        /// State of FTP / FTPS service for this function app. Possible values include: `AllAllowed`, `FtpsOnly` and `Disabled`.
        /// </summary>
        [Input("ftpsState")]
        public Input<string>? FtpsState { get; set; }

        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// Specifies whether or not the http2 protocol should be enabled. Defaults to `false`.
        /// </summary>
        [Input("http2Enabled")]
        public Input<bool>? Http2Enabled { get; set; }

        [Input("ipRestrictions")]
        private InputList<Inputs.FunctionAppSlotSiteConfigIpRestrictionGetArgs>? _ipRestrictions;

        /// <summary>
        /// A [List of objects](https://www.terraform.io/docs/configuration/attr-as-blocks.html) representing ip restrictions as defined below.
        /// </summary>
        public InputList<Inputs.FunctionAppSlotSiteConfigIpRestrictionGetArgs> IpRestrictions
        {
            get => _ipRestrictions ?? (_ipRestrictions = new InputList<Inputs.FunctionAppSlotSiteConfigIpRestrictionGetArgs>());
            set => _ipRestrictions = value;
        }

        [Input("javaVersion")]
        public Input<string>? JavaVersion { get; set; }

        /// <summary>
        /// Linux App Framework and version for the AppService, e.g. `DOCKER|(golang:latest)`.
        /// </summary>
        [Input("linuxFxVersion")]
        public Input<string>? LinuxFxVersion { get; set; }

        /// <summary>
        /// The minimum supported TLS version for the function app. Possible values are `1.0`, `1.1`, and `1.2`. Defaults to `1.2` for new function apps.
        /// </summary>
        [Input("minTlsVersion")]
        public Input<string>? MinTlsVersion { get; set; }

        /// <summary>
        /// The number of pre-warmed instances for this function app. Only affects apps on the Premium plan.
        /// </summary>
        [Input("preWarmedInstanceCount")]
        public Input<int>? PreWarmedInstanceCount { get; set; }

        /// <summary>
        /// Should Runtime Scale Monitoring be enabled?. Only applicable to apps on the Premium plan. Defaults to `false`.
        /// </summary>
        [Input("runtimeScaleMonitoringEnabled")]
        public Input<bool>? RuntimeScaleMonitoringEnabled { get; set; }

        [Input("scmIpRestrictions")]
        private InputList<Inputs.FunctionAppSlotSiteConfigScmIpRestrictionGetArgs>? _scmIpRestrictions;
        public InputList<Inputs.FunctionAppSlotSiteConfigScmIpRestrictionGetArgs> ScmIpRestrictions
        {
            get => _scmIpRestrictions ?? (_scmIpRestrictions = new InputList<Inputs.FunctionAppSlotSiteConfigScmIpRestrictionGetArgs>());
            set => _scmIpRestrictions = value;
        }

        [Input("scmType")]
        public Input<string>? ScmType { get; set; }

        [Input("scmUseMainIpRestriction")]
        public Input<bool>? ScmUseMainIpRestriction { get; set; }

        /// <summary>
        /// Should the Function App run in 32 bit mode, rather than 64 bit mode? Defaults to `true`.
        /// </summary>
        [Input("use32BitWorkerProcess")]
        public Input<bool>? Use32BitWorkerProcess { get; set; }

        /// <summary>
        /// Should WebSockets be enabled?
        /// </summary>
        [Input("websocketsEnabled")]
        public Input<bool>? WebsocketsEnabled { get; set; }

        public FunctionAppSlotSiteConfigGetArgs()
        {
        }
    }
}
