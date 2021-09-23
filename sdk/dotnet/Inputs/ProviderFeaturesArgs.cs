// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Inputs
{

    public sealed class ProviderFeaturesArgs : Pulumi.ResourceArgs
    {
        [Input("apiManagement")]
        public Input<Inputs.ProviderFeaturesApiManagementArgs>? ApiManagement { get; set; }

        [Input("cognitiveAccount")]
        public Input<Inputs.ProviderFeaturesCognitiveAccountArgs>? CognitiveAccount { get; set; }

        [Input("keyVault")]
        public Input<Inputs.ProviderFeaturesKeyVaultArgs>? KeyVault { get; set; }

        [Input("logAnalyticsWorkspace")]
        public Input<Inputs.ProviderFeaturesLogAnalyticsWorkspaceArgs>? LogAnalyticsWorkspace { get; set; }

        [Input("network")]
        public Input<Inputs.ProviderFeaturesNetworkArgs>? Network { get; set; }

        [Input("resourceGroup")]
        public Input<Inputs.ProviderFeaturesResourceGroupArgs>? ResourceGroup { get; set; }

        [Input("templateDeployment")]
        public Input<Inputs.ProviderFeaturesTemplateDeploymentArgs>? TemplateDeployment { get; set; }

        [Input("virtualMachine")]
        public Input<Inputs.ProviderFeaturesVirtualMachineArgs>? VirtualMachine { get; set; }

        [Input("virtualMachineScaleSet")]
        public Input<Inputs.ProviderFeaturesVirtualMachineScaleSetArgs>? VirtualMachineScaleSet { get; set; }

        public ProviderFeaturesArgs()
        {
        }
    }
}
