// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Inputs
{

    public sealed class ProviderFeaturesCognitiveAccountArgs : Pulumi.ResourceArgs
    {
        [Input("purgeSoftDeleteOnDestroy")]
        public Input<bool>? PurgeSoftDeleteOnDestroy { get; set; }

        public ProviderFeaturesCognitiveAccountArgs()
        {
        }
    }
}
