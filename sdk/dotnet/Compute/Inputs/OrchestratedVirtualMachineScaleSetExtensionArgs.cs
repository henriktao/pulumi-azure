// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class OrchestratedVirtualMachineScaleSetExtensionArgs : Pulumi.ResourceArgs
    {
        [Input("autoUpgradeMinorVersionEnabled")]
        public Input<bool>? AutoUpgradeMinorVersionEnabled { get; set; }

        [Input("extensionsToProvisionAfterVmCreations")]
        private InputList<string>? _extensionsToProvisionAfterVmCreations;

        /// <summary>
        /// An ordered list of Extension names which Orchestrated Virtual Machine Scale Set should provision after VM creation.
        /// </summary>
        public InputList<string> ExtensionsToProvisionAfterVmCreations
        {
            get => _extensionsToProvisionAfterVmCreations ?? (_extensionsToProvisionAfterVmCreations = new InputList<string>());
            set => _extensionsToProvisionAfterVmCreations = value;
        }

        [Input("forceExtensionExecutionOnChange")]
        public Input<string>? ForceExtensionExecutionOnChange { get; set; }

        /// <summary>
        /// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protectedSettings")]
        public Input<string>? ProtectedSettings { get; set; }

        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        [Input("settings")]
        public Input<string>? Settings { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("typeHandlerVersion", required: true)]
        public Input<string> TypeHandlerVersion { get; set; } = null!;

        public OrchestratedVirtualMachineScaleSetExtensionArgs()
        {
        }
    }
}
