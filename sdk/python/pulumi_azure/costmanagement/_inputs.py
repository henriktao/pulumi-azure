# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ResourceGroupExportDeliveryInfoArgs',
    'ResourceGroupExportQueryArgs',
]

@pulumi.input_type
class ResourceGroupExportDeliveryInfoArgs:
    def __init__(__self__, *,
                 container_name: pulumi.Input[str],
                 root_folder_path: pulumi.Input[str],
                 storage_account_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] container_name: The name of the container where exports will be uploaded.
        :param pulumi.Input[str] root_folder_path: The path of the directory where exports will be uploaded.
        :param pulumi.Input[str] storage_account_id: The storage account id where exports will be delivered.
        """
        pulumi.set(__self__, "container_name", container_name)
        pulumi.set(__self__, "root_folder_path", root_folder_path)
        pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Input[str]:
        """
        The name of the container where exports will be uploaded.
        """
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> pulumi.Input[str]:
        """
        The path of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")

    @root_folder_path.setter
    def root_folder_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_folder_path", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Input[str]:
        """
        The storage account id where exports will be delivered.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_id", value)


@pulumi.input_type
class ResourceGroupExportQueryArgs:
    def __init__(__self__, *,
                 time_frame: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] time_frame: The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        :param pulumi.Input[str] type: The type of the query.
        """
        pulumi.set(__self__, "time_frame", time_frame)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="timeFrame")
    def time_frame(self) -> pulumi.Input[str]:
        """
        The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        """
        return pulumi.get(self, "time_frame")

    @time_frame.setter
    def time_frame(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_frame", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the query.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


