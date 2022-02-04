// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps.Inputs
{

    public sealed class GetStandardSiteConfigArgs : Pulumi.InvokeArgs
    {
        [Input("alwaysOn")]
        public bool? AlwaysOn { get; set; }

        [Input("appScaleLimit", required: true)]
        public int AppScaleLimit { get; set; }

        [Input("cors", required: true)]
        public Inputs.GetStandardSiteConfigCorsArgs Cors { get; set; } = null!;

        [Input("dotnetFrameworkVersion")]
        public string? DotnetFrameworkVersion { get; set; }

        [Input("elasticInstanceMinimum", required: true)]
        public int ElasticInstanceMinimum { get; set; }

        [Input("ftpsState", required: true)]
        public string FtpsState { get; set; } = null!;

        [Input("healthCheckPath")]
        public string? HealthCheckPath { get; set; }

        [Input("http2Enabled")]
        public bool? Http2Enabled { get; set; }

        [Input("ipRestrictions", required: true)]
        private List<Inputs.GetStandardSiteConfigIpRestrictionArgs>? _ipRestrictions;
        public List<Inputs.GetStandardSiteConfigIpRestrictionArgs> IpRestrictions
        {
            get => _ipRestrictions ?? (_ipRestrictions = new List<Inputs.GetStandardSiteConfigIpRestrictionArgs>());
            set => _ipRestrictions = value;
        }

        [Input("linuxFxVersion", required: true)]
        public string LinuxFxVersion { get; set; } = null!;

        [Input("minTlsVersion", required: true)]
        public string MinTlsVersion { get; set; } = null!;

        [Input("preWarmedInstanceCount", required: true)]
        public int PreWarmedInstanceCount { get; set; }

        [Input("runtimeScaleMonitoringEnabled")]
        public bool? RuntimeScaleMonitoringEnabled { get; set; }

        [Input("use32BitWorkerProcess")]
        public bool? Use32BitWorkerProcess { get; set; }

        [Input("vnetRouteAllEnabled", required: true)]
        public bool VnetRouteAllEnabled { get; set; }

        [Input("websocketsEnabled")]
        public bool? WebsocketsEnabled { get; set; }

        public GetStandardSiteConfigArgs()
        {
        }
    }
}
