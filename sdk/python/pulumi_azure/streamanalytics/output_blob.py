# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OutputBlobArgs', 'OutputBlob']

@pulumi.input_type
class OutputBlobArgs:
    def __init__(__self__, *,
                 date_format: pulumi.Input[str],
                 path_pattern: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 serialization: pulumi.Input['OutputBlobSerializationArgs'],
                 storage_account_key: pulumi.Input[str],
                 storage_account_name: pulumi.Input[str],
                 storage_container_name: pulumi.Input[str],
                 stream_analytics_job_name: pulumi.Input[str],
                 time_format: pulumi.Input[str],
                 batch_max_wait_time: Optional[pulumi.Input[str]] = None,
                 batch_min_rows: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OutputBlob resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input['OutputBlobSerializationArgs'] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        :param pulumi.Input[str] batch_max_wait_time: The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        :param pulumi.Input[float] batch_min_rows: The minimum number of rows per batch (must be between `0` and `10000`).
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "date_format", date_format)
        pulumi.set(__self__, "path_pattern", path_pattern)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "serialization", serialization)
        pulumi.set(__self__, "storage_account_key", storage_account_key)
        pulumi.set(__self__, "storage_account_name", storage_account_name)
        pulumi.set(__self__, "storage_container_name", storage_container_name)
        pulumi.set(__self__, "stream_analytics_job_name", stream_analytics_job_name)
        pulumi.set(__self__, "time_format", time_format)
        if batch_max_wait_time is not None:
            pulumi.set(__self__, "batch_max_wait_time", batch_max_wait_time)
        if batch_min_rows is not None:
            pulumi.set(__self__, "batch_min_rows", batch_min_rows)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="dateFormat")
    def date_format(self) -> pulumi.Input[str]:
        """
        The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        """
        return pulumi.get(self, "date_format")

    @date_format.setter
    def date_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "date_format", value)

    @property
    @pulumi.getter(name="pathPattern")
    def path_pattern(self) -> pulumi.Input[str]:
        """
        The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        """
        return pulumi.get(self, "path_pattern")

    @path_pattern.setter
    def path_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "path_pattern", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Input['OutputBlobSerializationArgs']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @serialization.setter
    def serialization(self, value: pulumi.Input['OutputBlobSerializationArgs']):
        pulumi.set(self, "serialization", value)

    @property
    @pulumi.getter(name="storageAccountKey")
    def storage_account_key(self) -> pulumi.Input[str]:
        """
        The Access Key which should be used to connect to this Storage Account.
        """
        return pulumi.get(self, "storage_account_key")

    @storage_account_key.setter
    def storage_account_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_key", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Input[str]:
        """
        The name of the Storage Account.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter(name="storageContainerName")
    def storage_container_name(self) -> pulumi.Input[str]:
        """
        The name of the Container within the Storage Account.
        """
        return pulumi.get(self, "storage_container_name")

    @storage_container_name.setter
    def storage_container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_container_name", value)

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Input[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @stream_analytics_job_name.setter
    def stream_analytics_job_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_analytics_job_name", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> pulumi.Input[str]:
        """
        The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_format", value)

    @property
    @pulumi.getter(name="batchMaxWaitTime")
    def batch_max_wait_time(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        """
        return pulumi.get(self, "batch_max_wait_time")

    @batch_max_wait_time.setter
    def batch_max_wait_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "batch_max_wait_time", value)

    @property
    @pulumi.getter(name="batchMinRows")
    def batch_min_rows(self) -> Optional[pulumi.Input[float]]:
        """
        The minimum number of rows per batch (must be between `0` and `10000`).
        """
        return pulumi.get(self, "batch_min_rows")

    @batch_min_rows.setter
    def batch_min_rows(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "batch_min_rows", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Output. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OutputBlobState:
    def __init__(__self__, *,
                 batch_max_wait_time: Optional[pulumi.Input[str]] = None,
                 batch_min_rows: Optional[pulumi.Input[float]] = None,
                 date_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path_pattern: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input['OutputBlobSerializationArgs']] = None,
                 storage_account_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OutputBlob resources.
        :param pulumi.Input[str] batch_max_wait_time: The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        :param pulumi.Input[float] batch_min_rows: The minimum number of rows per batch (must be between `0` and `10000`).
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input['OutputBlobSerializationArgs'] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        if batch_max_wait_time is not None:
            pulumi.set(__self__, "batch_max_wait_time", batch_max_wait_time)
        if batch_min_rows is not None:
            pulumi.set(__self__, "batch_min_rows", batch_min_rows)
        if date_format is not None:
            pulumi.set(__self__, "date_format", date_format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path_pattern is not None:
            pulumi.set(__self__, "path_pattern", path_pattern)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if serialization is not None:
            pulumi.set(__self__, "serialization", serialization)
        if storage_account_key is not None:
            pulumi.set(__self__, "storage_account_key", storage_account_key)
        if storage_account_name is not None:
            pulumi.set(__self__, "storage_account_name", storage_account_name)
        if storage_container_name is not None:
            pulumi.set(__self__, "storage_container_name", storage_container_name)
        if stream_analytics_job_name is not None:
            pulumi.set(__self__, "stream_analytics_job_name", stream_analytics_job_name)
        if time_format is not None:
            pulumi.set(__self__, "time_format", time_format)

    @property
    @pulumi.getter(name="batchMaxWaitTime")
    def batch_max_wait_time(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        """
        return pulumi.get(self, "batch_max_wait_time")

    @batch_max_wait_time.setter
    def batch_max_wait_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "batch_max_wait_time", value)

    @property
    @pulumi.getter(name="batchMinRows")
    def batch_min_rows(self) -> Optional[pulumi.Input[float]]:
        """
        The minimum number of rows per batch (must be between `0` and `10000`).
        """
        return pulumi.get(self, "batch_min_rows")

    @batch_min_rows.setter
    def batch_min_rows(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "batch_min_rows", value)

    @property
    @pulumi.getter(name="dateFormat")
    def date_format(self) -> Optional[pulumi.Input[str]]:
        """
        The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        """
        return pulumi.get(self, "date_format")

    @date_format.setter
    def date_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Output. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pathPattern")
    def path_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        """
        return pulumi.get(self, "path_pattern")

    @path_pattern.setter
    def path_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path_pattern", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def serialization(self) -> Optional[pulumi.Input['OutputBlobSerializationArgs']]:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @serialization.setter
    def serialization(self, value: Optional[pulumi.Input['OutputBlobSerializationArgs']]):
        pulumi.set(self, "serialization", value)

    @property
    @pulumi.getter(name="storageAccountKey")
    def storage_account_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Access Key which should be used to connect to this Storage Account.
        """
        return pulumi.get(self, "storage_account_key")

    @storage_account_key.setter
    def storage_account_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_key", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Storage Account.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter(name="storageContainerName")
    def storage_container_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Container within the Storage Account.
        """
        return pulumi.get(self, "storage_container_name")

    @storage_container_name.setter
    def storage_container_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_container_name", value)

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @stream_analytics_job_name.setter
    def stream_analytics_job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_analytics_job_name", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> Optional[pulumi.Input[str]]:
        """
        The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_format", value)


class OutputBlob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_max_wait_time: Optional[pulumi.Input[str]] = None,
                 batch_min_rows: Optional[pulumi.Input[float]] = None,
                 date_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path_pattern: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['OutputBlobSerializationArgs']]] = None,
                 storage_account_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Stream Analytics Output to Blob Storage.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_output_blob = azure.streamanalytics.OutputBlob("exampleOutputBlob",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            storage_account_name=example_account.name,
            storage_account_key=example_account.primary_access_key,
            storage_container_name=example_container.name,
            path_pattern="some-pattern",
            date_format="yyyy-MM-dd",
            time_format="HH",
            serialization=azure.streamanalytics.OutputBlobSerializationArgs(
                type="Csv",
                encoding="UTF8",
                field_delimiter=",",
            ))
        ```

        ## Import

        Stream Analytics Outputs to Blob Storage can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/outputBlob:OutputBlob example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] batch_max_wait_time: The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        :param pulumi.Input[float] batch_min_rows: The minimum number of rows per batch (must be between `0` and `10000`).
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['OutputBlobSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OutputBlobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Stream Analytics Output to Blob Storage.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.get_resource_group(name="example-resources")
        example_job = azure.streamanalytics.get_job(name="example-job",
            resource_group_name=azurerm_resource_group["example"]["name"])
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer",
            storage_account_name=example_account.name,
            container_access_type="private")
        example_output_blob = azure.streamanalytics.OutputBlob("exampleOutputBlob",
            stream_analytics_job_name=example_job.name,
            resource_group_name=example_job.resource_group_name,
            storage_account_name=example_account.name,
            storage_account_key=example_account.primary_access_key,
            storage_container_name=example_container.name,
            path_pattern="some-pattern",
            date_format="yyyy-MM-dd",
            time_format="HH",
            serialization=azure.streamanalytics.OutputBlobSerializationArgs(
                type="Csv",
                encoding="UTF8",
                field_delimiter=",",
            ))
        ```

        ## Import

        Stream Analytics Outputs to Blob Storage can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:streamanalytics/outputBlob:OutputBlob example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
        ```

        :param str resource_name: The name of the resource.
        :param OutputBlobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OutputBlobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_max_wait_time: Optional[pulumi.Input[str]] = None,
                 batch_min_rows: Optional[pulumi.Input[float]] = None,
                 date_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path_pattern: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 serialization: Optional[pulumi.Input[pulumi.InputType['OutputBlobSerializationArgs']]] = None,
                 storage_account_key: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_container_name: Optional[pulumi.Input[str]] = None,
                 stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OutputBlobArgs.__new__(OutputBlobArgs)

            __props__.__dict__["batch_max_wait_time"] = batch_max_wait_time
            __props__.__dict__["batch_min_rows"] = batch_min_rows
            if date_format is None and not opts.urn:
                raise TypeError("Missing required property 'date_format'")
            __props__.__dict__["date_format"] = date_format
            __props__.__dict__["name"] = name
            if path_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'path_pattern'")
            __props__.__dict__["path_pattern"] = path_pattern
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if serialization is None and not opts.urn:
                raise TypeError("Missing required property 'serialization'")
            __props__.__dict__["serialization"] = serialization
            if storage_account_key is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_key'")
            __props__.__dict__["storage_account_key"] = storage_account_key
            if storage_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__.__dict__["storage_account_name"] = storage_account_name
            if storage_container_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__.__dict__["storage_container_name"] = storage_container_name
            if stream_analytics_job_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__.__dict__["stream_analytics_job_name"] = stream_analytics_job_name
            if time_format is None and not opts.urn:
                raise TypeError("Missing required property 'time_format'")
            __props__.__dict__["time_format"] = time_format
        super(OutputBlob, __self__).__init__(
            'azure:streamanalytics/outputBlob:OutputBlob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            batch_max_wait_time: Optional[pulumi.Input[str]] = None,
            batch_min_rows: Optional[pulumi.Input[float]] = None,
            date_format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path_pattern: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            serialization: Optional[pulumi.Input[pulumi.InputType['OutputBlobSerializationArgs']]] = None,
            storage_account_key: Optional[pulumi.Input[str]] = None,
            storage_account_name: Optional[pulumi.Input[str]] = None,
            storage_container_name: Optional[pulumi.Input[str]] = None,
            stream_analytics_job_name: Optional[pulumi.Input[str]] = None,
            time_format: Optional[pulumi.Input[str]] = None) -> 'OutputBlob':
        """
        Get an existing OutputBlob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] batch_max_wait_time: The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        :param pulumi.Input[float] batch_min_rows: The minimum number of rows per batch (must be between `0` and `10000`).
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['OutputBlobSerializationArgs']] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OutputBlobState.__new__(_OutputBlobState)

        __props__.__dict__["batch_max_wait_time"] = batch_max_wait_time
        __props__.__dict__["batch_min_rows"] = batch_min_rows
        __props__.__dict__["date_format"] = date_format
        __props__.__dict__["name"] = name
        __props__.__dict__["path_pattern"] = path_pattern
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["serialization"] = serialization
        __props__.__dict__["storage_account_key"] = storage_account_key
        __props__.__dict__["storage_account_name"] = storage_account_name
        __props__.__dict__["storage_container_name"] = storage_container_name
        __props__.__dict__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__.__dict__["time_format"] = time_format
        return OutputBlob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="batchMaxWaitTime")
    def batch_max_wait_time(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum wait time per batch in `hh:mm:ss` e.g. `00:02:00` for two minutes.
        """
        return pulumi.get(self, "batch_max_wait_time")

    @property
    @pulumi.getter(name="batchMinRows")
    def batch_min_rows(self) -> pulumi.Output[Optional[float]]:
        """
        The minimum number of rows per batch (must be between `0` and `10000`).
        """
        return pulumi.get(self, "batch_min_rows")

    @property
    @pulumi.getter(name="dateFormat")
    def date_format(self) -> pulumi.Output[str]:
        """
        The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        """
        return pulumi.get(self, "date_format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Output. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pathPattern")
    def path_pattern(self) -> pulumi.Output[str]:
        """
        The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        """
        return pulumi.get(self, "path_pattern")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def serialization(self) -> pulumi.Output['outputs.OutputBlobSerialization']:
        """
        A `serialization` block as defined below.
        """
        return pulumi.get(self, "serialization")

    @property
    @pulumi.getter(name="storageAccountKey")
    def storage_account_key(self) -> pulumi.Output[str]:
        """
        The Access Key which should be used to connect to this Storage Account.
        """
        return pulumi.get(self, "storage_account_key")

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> pulumi.Output[str]:
        """
        The name of the Storage Account.
        """
        return pulumi.get(self, "storage_account_name")

    @property
    @pulumi.getter(name="storageContainerName")
    def storage_container_name(self) -> pulumi.Output[str]:
        """
        The name of the Container within the Storage Account.
        """
        return pulumi.get(self, "storage_container_name")

    @property
    @pulumi.getter(name="streamAnalyticsJobName")
    def stream_analytics_job_name(self) -> pulumi.Output[str]:
        """
        The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "stream_analytics_job_name")

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> pulumi.Output[str]:
        """
        The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        return pulumi.get(self, "time_format")

