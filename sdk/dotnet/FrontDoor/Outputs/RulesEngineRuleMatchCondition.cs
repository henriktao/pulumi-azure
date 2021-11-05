// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Outputs
{

    [OutputType]
    public sealed class RulesEngineRuleMatchCondition
    {
        /// <summary>
        /// can be set to `true` or `false` to negate the given condition. Defaults to `true`.
        /// </summary>
        public readonly bool? NegateCondition;
        /// <summary>
        /// can be set to `Any`, `IPMatch`, `GeoMatch`, `Equal`, `Contains`, `LessThan`, `GreaterThan`, `LessThanOrEqual`, `GreaterThanOrEqual`, `BeginsWith` or `EndsWith`
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// match against a specific key when `variable` is set to `PostArgs` or `RequestHeader`. It cannot be used with `QueryString` and `RequestMethod`. Defaults to `null`.
        /// </summary>
        public readonly string? Selector;
        /// <summary>
        /// can be set to one or more values out of `Lowercase`, `RemoveNulls`, `Trim`, `Uppercase`, `UrlDecode` and `UrlEncode`
        /// </summary>
        public readonly ImmutableArray<string> Transforms;
        /// <summary>
        /// can contain one or more strings.
        /// </summary>
        public readonly ImmutableArray<string> Values;
        /// <summary>
        /// can be set to `IsMobile`, `RemoteAddr`, `RequestMethod`, `QueryString`, `PostArgs`, `RequestURI`, `RequestPath`, `RequestFilename`, `RequestFilenameExtension`,`RequestHeader`,`RequestBody` or `RequestScheme`.
        /// </summary>
        public readonly string? Variable;

        [OutputConstructor]
        private RulesEngineRuleMatchCondition(
            bool? negateCondition,

            string @operator,

            string? selector,

            ImmutableArray<string> transforms,

            ImmutableArray<string> values,

            string? variable)
        {
            NegateCondition = negateCondition;
            Operator = @operator;
            Selector = selector;
            Transforms = transforms;
            Values = values;
            Variable = variable;
        }
    }
}
