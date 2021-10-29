// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachineScaleSetExtensionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<bool>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension? Defaults to `false`.
        /// </summary>
        [Input("automaticUpgradeEnabled")]
        public Input<bool>? AutomaticUpgradeEnabled { get; set; }

        /// <summary>
        /// A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        /// </summary>
        [Input("forceUpdateTag")]
        public Input<string>? ForceUpdateTag { get; set; }

        /// <summary>
        /// The name for the Virtual Machine Scale Set Extension.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        /// </summary>
        [Input("protectedSettings")]
        public Input<string>? ProtectedSettings { get; set; }

        [Input("provisionAfterExtensions")]
        private InputList<string>? _provisionAfterExtensions;

        /// <summary>
        /// An ordered list of Extension names which this should be provisioned after.
        /// </summary>
        public InputList<string> ProvisionAfterExtensions
        {
            get => _provisionAfterExtensions ?? (_provisionAfterExtensions = new InputList<string>());
            set => _provisionAfterExtensions = value;
        }

        /// <summary>
        /// Specifies the Publisher of the Extension.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        /// <summary>
        /// A JSON String which specifies Settings for the Extension.
        /// </summary>
        [Input("settings")]
        public Input<string>? Settings { get; set; }

        /// <summary>
        /// Specifies the Type of the Extension.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
        /// </summary>
        [Input("typeHandlerVersion", required: true)]
        public Input<string> TypeHandlerVersion { get; set; } = null!;

        public LinuxVirtualMachineScaleSetExtensionGetArgs()
        {
        }
    }
}
