// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class DatasetSnowflakeSchemaColumn
    {
        /// <summary>
        /// The name of the column.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The total number of digits allowed.
        /// </summary>
        public readonly int? Precision;
        /// <summary>
        /// The number of digits allowed to the right of the decimal point.
        /// </summary>
        public readonly int? Scale;
        /// <summary>
        /// Type of the column. Valid values are `NUMBER`, `DECIMAL`, `NUMERIC`, `INT`, `INTEGER`, `BIGINT`, `SMALLINT`, `FLOAT``FLOAT4`, `FLOAT8`, `DOUBLE`, `DOUBLE PRECISION`, `REAL`, `VARCHAR`, `CHAR`, `CHARACTER`, `STRING`, `TEXT`, `BINARY`, `VARBINARY`, `BOOLEAN`, `DATE`, `DATETIME`, `TIME`, `TIMESTAMP`, `TIMESTAMP_LTZ`, `TIMESTAMP_NTZ`, `TIMESTAMP_TZ`, `VARIANT`, `OBJECT`, `ARRAY`, `GEOGRAPHY`. Please note these values are case sensitive.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DatasetSnowflakeSchemaColumn(
            string name,

            int? precision,

            int? scale,

            string? type)
        {
            Name = name;
            Precision = precision;
            Scale = scale;
            Type = type;
        }
    }
}
