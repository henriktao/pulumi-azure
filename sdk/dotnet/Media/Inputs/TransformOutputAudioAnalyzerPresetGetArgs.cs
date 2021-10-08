// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class TransformOutputAudioAnalyzerPresetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Possibles value are `Basic` or `Standard`. Determines the set of audio analysis operations to be performed.
        /// </summary>
        [Input("audioAnalysisMode")]
        public Input<string>? AudioAnalysisMode { get; set; }

        /// <summary>
        /// The language for the audio payload in the input using the BCP-47 format of 'language tag-region' (e.g: 'en-US'). If you know the language of your content, it is recommended that you specify it. The language must be specified explicitly for AudioAnalysisMode:Basic, since automatic language detection is not included in basic mode. If the language isn't specified, automatic language detection will choose the first language detected and process with the selected language for the duration of the file. It does not currently support dynamically switching between languages after the first language is detected. The automatic detection works best with audio recordings with clearly discernible speech. If automatic detection fails to find the language, transcription would fallback to 'en-US'." The list of supported languages is available here: https://go.microsoft.com/fwlink/?linkid=2109463.
        /// </summary>
        [Input("audioLanguage")]
        public Input<string>? AudioLanguage { get; set; }

        public TransformOutputAudioAnalyzerPresetGetArgs()
        {
        }
    }
}
