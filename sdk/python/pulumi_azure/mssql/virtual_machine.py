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

__all__ = ['VirtualMachineArgs', 'VirtualMachine']

@pulumi.input_type
class VirtualMachineArgs:
    def __init__(__self__, *,
                 sql_license_type: pulumi.Input[str],
                 virtual_machine_id: pulumi.Input[str],
                 auto_backup: Optional[pulumi.Input['VirtualMachineAutoBackupArgs']] = None,
                 auto_patching: Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']] = None,
                 key_vault_credential: Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']] = None,
                 r_services_enabled: Optional[pulumi.Input[bool]] = None,
                 sql_connectivity_port: Optional[pulumi.Input[int]] = None,
                 sql_connectivity_type: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
                 storage_configuration: Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VirtualMachine resource.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
        :param pulumi.Input['VirtualMachineAutoBackupArgs'] auto_backup: An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        :param pulumi.Input['VirtualMachineAutoPatchingArgs'] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input['VirtualMachineKeyVaultCredentialArgs'] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[int] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input['VirtualMachineStorageConfigurationArgs'] storage_configuration: An `storage_configuration` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "sql_license_type", sql_license_type)
        pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)
        if auto_backup is not None:
            pulumi.set(__self__, "auto_backup", auto_backup)
        if auto_patching is not None:
            pulumi.set(__self__, "auto_patching", auto_patching)
        if key_vault_credential is not None:
            pulumi.set(__self__, "key_vault_credential", key_vault_credential)
        if r_services_enabled is not None:
            pulumi.set(__self__, "r_services_enabled", r_services_enabled)
        if sql_connectivity_port is not None:
            pulumi.set(__self__, "sql_connectivity_port", sql_connectivity_port)
        if sql_connectivity_type is not None:
            pulumi.set(__self__, "sql_connectivity_type", sql_connectivity_type)
        if sql_connectivity_update_password is not None:
            pulumi.set(__self__, "sql_connectivity_update_password", sql_connectivity_update_password)
        if sql_connectivity_update_username is not None:
            pulumi.set(__self__, "sql_connectivity_update_username", sql_connectivity_update_username)
        if storage_configuration is not None:
            pulumi.set(__self__, "storage_configuration", storage_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="sqlLicenseType")
    def sql_license_type(self) -> pulumi.Input[str]:
        """
        The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sql_license_type")

    @sql_license_type.setter
    def sql_license_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "sql_license_type", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Input[str]:
        """
        The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_machine_id", value)

    @property
    @pulumi.getter(name="autoBackup")
    def auto_backup(self) -> Optional[pulumi.Input['VirtualMachineAutoBackupArgs']]:
        """
        An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        """
        return pulumi.get(self, "auto_backup")

    @auto_backup.setter
    def auto_backup(self, value: Optional[pulumi.Input['VirtualMachineAutoBackupArgs']]):
        pulumi.set(self, "auto_backup", value)

    @property
    @pulumi.getter(name="autoPatching")
    def auto_patching(self) -> Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']]:
        """
        An `auto_patching` block as defined below.
        """
        return pulumi.get(self, "auto_patching")

    @auto_patching.setter
    def auto_patching(self, value: Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']]):
        pulumi.set(self, "auto_patching", value)

    @property
    @pulumi.getter(name="keyVaultCredential")
    def key_vault_credential(self) -> Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']]:
        """
        (Optional) An `key_vault_credential` block as defined below.
        """
        return pulumi.get(self, "key_vault_credential")

    @key_vault_credential.setter
    def key_vault_credential(self, value: Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']]):
        pulumi.set(self, "key_vault_credential", value)

    @property
    @pulumi.getter(name="rServicesEnabled")
    def r_services_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should R Services be enabled?
        """
        return pulumi.get(self, "r_services_enabled")

    @r_services_enabled.setter
    def r_services_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "r_services_enabled", value)

    @property
    @pulumi.getter(name="sqlConnectivityPort")
    def sql_connectivity_port(self) -> Optional[pulumi.Input[int]]:
        """
        The SQL Server port. Defaults to `1433`.
        """
        return pulumi.get(self, "sql_connectivity_port")

    @sql_connectivity_port.setter
    def sql_connectivity_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sql_connectivity_port", value)

    @property
    @pulumi.getter(name="sqlConnectivityType")
    def sql_connectivity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        """
        return pulumi.get(self, "sql_connectivity_type")

    @sql_connectivity_type.setter
    def sql_connectivity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_type", value)

    @property
    @pulumi.getter(name="sqlConnectivityUpdatePassword")
    def sql_connectivity_update_password(self) -> Optional[pulumi.Input[str]]:
        """
        The SQL Server sysadmin login password.
        """
        return pulumi.get(self, "sql_connectivity_update_password")

    @sql_connectivity_update_password.setter
    def sql_connectivity_update_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_update_password", value)

    @property
    @pulumi.getter(name="sqlConnectivityUpdateUsername")
    def sql_connectivity_update_username(self) -> Optional[pulumi.Input[str]]:
        """
        The SQL Server sysadmin login to create.
        """
        return pulumi.get(self, "sql_connectivity_update_username")

    @sql_connectivity_update_username.setter
    def sql_connectivity_update_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_update_username", value)

    @property
    @pulumi.getter(name="storageConfiguration")
    def storage_configuration(self) -> Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']]:
        """
        An `storage_configuration` block as defined below.
        """
        return pulumi.get(self, "storage_configuration")

    @storage_configuration.setter
    def storage_configuration(self, value: Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']]):
        pulumi.set(self, "storage_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VirtualMachineState:
    def __init__(__self__, *,
                 auto_backup: Optional[pulumi.Input['VirtualMachineAutoBackupArgs']] = None,
                 auto_patching: Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']] = None,
                 key_vault_credential: Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']] = None,
                 r_services_enabled: Optional[pulumi.Input[bool]] = None,
                 sql_connectivity_port: Optional[pulumi.Input[int]] = None,
                 sql_connectivity_type: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
                 sql_license_type: Optional[pulumi.Input[str]] = None,
                 storage_configuration: Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VirtualMachine resources.
        :param pulumi.Input['VirtualMachineAutoBackupArgs'] auto_backup: An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        :param pulumi.Input['VirtualMachineAutoPatchingArgs'] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input['VirtualMachineKeyVaultCredentialArgs'] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[int] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input['VirtualMachineStorageConfigurationArgs'] storage_configuration: An `storage_configuration` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        if auto_backup is not None:
            pulumi.set(__self__, "auto_backup", auto_backup)
        if auto_patching is not None:
            pulumi.set(__self__, "auto_patching", auto_patching)
        if key_vault_credential is not None:
            pulumi.set(__self__, "key_vault_credential", key_vault_credential)
        if r_services_enabled is not None:
            pulumi.set(__self__, "r_services_enabled", r_services_enabled)
        if sql_connectivity_port is not None:
            pulumi.set(__self__, "sql_connectivity_port", sql_connectivity_port)
        if sql_connectivity_type is not None:
            pulumi.set(__self__, "sql_connectivity_type", sql_connectivity_type)
        if sql_connectivity_update_password is not None:
            pulumi.set(__self__, "sql_connectivity_update_password", sql_connectivity_update_password)
        if sql_connectivity_update_username is not None:
            pulumi.set(__self__, "sql_connectivity_update_username", sql_connectivity_update_username)
        if sql_license_type is not None:
            pulumi.set(__self__, "sql_license_type", sql_license_type)
        if storage_configuration is not None:
            pulumi.set(__self__, "storage_configuration", storage_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter(name="autoBackup")
    def auto_backup(self) -> Optional[pulumi.Input['VirtualMachineAutoBackupArgs']]:
        """
        An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        """
        return pulumi.get(self, "auto_backup")

    @auto_backup.setter
    def auto_backup(self, value: Optional[pulumi.Input['VirtualMachineAutoBackupArgs']]):
        pulumi.set(self, "auto_backup", value)

    @property
    @pulumi.getter(name="autoPatching")
    def auto_patching(self) -> Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']]:
        """
        An `auto_patching` block as defined below.
        """
        return pulumi.get(self, "auto_patching")

    @auto_patching.setter
    def auto_patching(self, value: Optional[pulumi.Input['VirtualMachineAutoPatchingArgs']]):
        pulumi.set(self, "auto_patching", value)

    @property
    @pulumi.getter(name="keyVaultCredential")
    def key_vault_credential(self) -> Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']]:
        """
        (Optional) An `key_vault_credential` block as defined below.
        """
        return pulumi.get(self, "key_vault_credential")

    @key_vault_credential.setter
    def key_vault_credential(self, value: Optional[pulumi.Input['VirtualMachineKeyVaultCredentialArgs']]):
        pulumi.set(self, "key_vault_credential", value)

    @property
    @pulumi.getter(name="rServicesEnabled")
    def r_services_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should R Services be enabled?
        """
        return pulumi.get(self, "r_services_enabled")

    @r_services_enabled.setter
    def r_services_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "r_services_enabled", value)

    @property
    @pulumi.getter(name="sqlConnectivityPort")
    def sql_connectivity_port(self) -> Optional[pulumi.Input[int]]:
        """
        The SQL Server port. Defaults to `1433`.
        """
        return pulumi.get(self, "sql_connectivity_port")

    @sql_connectivity_port.setter
    def sql_connectivity_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sql_connectivity_port", value)

    @property
    @pulumi.getter(name="sqlConnectivityType")
    def sql_connectivity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        """
        return pulumi.get(self, "sql_connectivity_type")

    @sql_connectivity_type.setter
    def sql_connectivity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_type", value)

    @property
    @pulumi.getter(name="sqlConnectivityUpdatePassword")
    def sql_connectivity_update_password(self) -> Optional[pulumi.Input[str]]:
        """
        The SQL Server sysadmin login password.
        """
        return pulumi.get(self, "sql_connectivity_update_password")

    @sql_connectivity_update_password.setter
    def sql_connectivity_update_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_update_password", value)

    @property
    @pulumi.getter(name="sqlConnectivityUpdateUsername")
    def sql_connectivity_update_username(self) -> Optional[pulumi.Input[str]]:
        """
        The SQL Server sysadmin login to create.
        """
        return pulumi.get(self, "sql_connectivity_update_username")

    @sql_connectivity_update_username.setter
    def sql_connectivity_update_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_connectivity_update_username", value)

    @property
    @pulumi.getter(name="sqlLicenseType")
    def sql_license_type(self) -> Optional[pulumi.Input[str]]:
        """
        The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sql_license_type")

    @sql_license_type.setter
    def sql_license_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_license_type", value)

    @property
    @pulumi.getter(name="storageConfiguration")
    def storage_configuration(self) -> Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']]:
        """
        An `storage_configuration` block as defined below.
        """
        return pulumi.get(self, "storage_configuration")

    @storage_configuration.setter
    def storage_configuration(self, value: Optional[pulumi.Input['VirtualMachineStorageConfigurationArgs']]):
        pulumi.set(self, "storage_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_machine_id", value)


class VirtualMachine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_backup: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoBackupArgs']]] = None,
                 auto_patching: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']]] = None,
                 key_vault_credential: Optional[pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']]] = None,
                 r_services_enabled: Optional[pulumi.Input[bool]] = None,
                 sql_connectivity_port: Optional[pulumi.Input[int]] = None,
                 sql_connectivity_type: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
                 sql_license_type: Optional[pulumi.Input[str]] = None,
                 storage_configuration: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Microsoft SQL Virtual Machine

        ## Example Usage

        This example provisions a brief Managed MsSql Virtual Machine.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_virtual_machine = azure.compute.get_virtual_machine(name="example-vm",
            resource_group_name="example-resources")
        example_mssql_virtual_machine_virtual_machine = azure.mssql.VirtualMachine("exampleMssql/virtualMachineVirtualMachine",
            virtual_machine_id=example_virtual_machine.id,
            sql_license_type="PAYG",
            r_services_enabled=True,
            sql_connectivity_port=1433,
            sql_connectivity_type="PRIVATE",
            sql_connectivity_update_password="Password1234!",
            sql_connectivity_update_username="sqllogin",
            auto_patching=azure.mssql.VirtualMachineAutoPatchingArgs(
                day_of_week="Sunday",
                maintenance_window_duration_in_minutes=60,
                maintenance_window_starting_hour=2,
            ))
        ```

        ## Import

        Sql Virtual Machines can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mssql/virtualMachine:VirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoBackupArgs']] auto_backup: An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[int] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageConfigurationArgs']] storage_configuration: An `storage_configuration` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualMachineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Microsoft SQL Virtual Machine

        ## Example Usage

        This example provisions a brief Managed MsSql Virtual Machine.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_virtual_machine = azure.compute.get_virtual_machine(name="example-vm",
            resource_group_name="example-resources")
        example_mssql_virtual_machine_virtual_machine = azure.mssql.VirtualMachine("exampleMssql/virtualMachineVirtualMachine",
            virtual_machine_id=example_virtual_machine.id,
            sql_license_type="PAYG",
            r_services_enabled=True,
            sql_connectivity_port=1433,
            sql_connectivity_type="PRIVATE",
            sql_connectivity_update_password="Password1234!",
            sql_connectivity_update_username="sqllogin",
            auto_patching=azure.mssql.VirtualMachineAutoPatchingArgs(
                day_of_week="Sunday",
                maintenance_window_duration_in_minutes=60,
                maintenance_window_starting_hour=2,
            ))
        ```

        ## Import

        Sql Virtual Machines can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:mssql/virtualMachine:VirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
        ```

        :param str resource_name: The name of the resource.
        :param VirtualMachineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualMachineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_backup: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoBackupArgs']]] = None,
                 auto_patching: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']]] = None,
                 key_vault_credential: Optional[pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']]] = None,
                 r_services_enabled: Optional[pulumi.Input[bool]] = None,
                 sql_connectivity_port: Optional[pulumi.Input[int]] = None,
                 sql_connectivity_type: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
                 sql_license_type: Optional[pulumi.Input[str]] = None,
                 storage_configuration: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VirtualMachineArgs.__new__(VirtualMachineArgs)

            __props__.__dict__["auto_backup"] = auto_backup
            __props__.__dict__["auto_patching"] = auto_patching
            __props__.__dict__["key_vault_credential"] = key_vault_credential
            __props__.__dict__["r_services_enabled"] = r_services_enabled
            __props__.__dict__["sql_connectivity_port"] = sql_connectivity_port
            __props__.__dict__["sql_connectivity_type"] = sql_connectivity_type
            __props__.__dict__["sql_connectivity_update_password"] = sql_connectivity_update_password
            __props__.__dict__["sql_connectivity_update_username"] = sql_connectivity_update_username
            if sql_license_type is None and not opts.urn:
                raise TypeError("Missing required property 'sql_license_type'")
            __props__.__dict__["sql_license_type"] = sql_license_type
            __props__.__dict__["storage_configuration"] = storage_configuration
            __props__.__dict__["tags"] = tags
            if virtual_machine_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        super(VirtualMachine, __self__).__init__(
            'azure:mssql/virtualMachine:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_backup: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoBackupArgs']]] = None,
            auto_patching: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']]] = None,
            key_vault_credential: Optional[pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']]] = None,
            r_services_enabled: Optional[pulumi.Input[bool]] = None,
            sql_connectivity_port: Optional[pulumi.Input[int]] = None,
            sql_connectivity_type: Optional[pulumi.Input[str]] = None,
            sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
            sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
            sql_license_type: Optional[pulumi.Input[str]] = None,
            storage_configuration: Optional[pulumi.Input[pulumi.InputType['VirtualMachineStorageConfigurationArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoBackupArgs']] auto_backup: An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[int] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['VirtualMachineStorageConfigurationArgs']] storage_configuration: An `storage_configuration` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualMachineState.__new__(_VirtualMachineState)

        __props__.__dict__["auto_backup"] = auto_backup
        __props__.__dict__["auto_patching"] = auto_patching
        __props__.__dict__["key_vault_credential"] = key_vault_credential
        __props__.__dict__["r_services_enabled"] = r_services_enabled
        __props__.__dict__["sql_connectivity_port"] = sql_connectivity_port
        __props__.__dict__["sql_connectivity_type"] = sql_connectivity_type
        __props__.__dict__["sql_connectivity_update_password"] = sql_connectivity_update_password
        __props__.__dict__["sql_connectivity_update_username"] = sql_connectivity_update_username
        __props__.__dict__["sql_license_type"] = sql_license_type
        __props__.__dict__["storage_configuration"] = storage_configuration
        __props__.__dict__["tags"] = tags
        __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoBackup")
    def auto_backup(self) -> pulumi.Output[Optional['outputs.VirtualMachineAutoBackup']]:
        """
        An `auto_backup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
        """
        return pulumi.get(self, "auto_backup")

    @property
    @pulumi.getter(name="autoPatching")
    def auto_patching(self) -> pulumi.Output[Optional['outputs.VirtualMachineAutoPatching']]:
        """
        An `auto_patching` block as defined below.
        """
        return pulumi.get(self, "auto_patching")

    @property
    @pulumi.getter(name="keyVaultCredential")
    def key_vault_credential(self) -> pulumi.Output[Optional['outputs.VirtualMachineKeyVaultCredential']]:
        """
        (Optional) An `key_vault_credential` block as defined below.
        """
        return pulumi.get(self, "key_vault_credential")

    @property
    @pulumi.getter(name="rServicesEnabled")
    def r_services_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Should R Services be enabled?
        """
        return pulumi.get(self, "r_services_enabled")

    @property
    @pulumi.getter(name="sqlConnectivityPort")
    def sql_connectivity_port(self) -> pulumi.Output[Optional[int]]:
        """
        The SQL Server port. Defaults to `1433`.
        """
        return pulumi.get(self, "sql_connectivity_port")

    @property
    @pulumi.getter(name="sqlConnectivityType")
    def sql_connectivity_type(self) -> pulumi.Output[Optional[str]]:
        """
        The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        """
        return pulumi.get(self, "sql_connectivity_type")

    @property
    @pulumi.getter(name="sqlConnectivityUpdatePassword")
    def sql_connectivity_update_password(self) -> pulumi.Output[Optional[str]]:
        """
        The SQL Server sysadmin login password.
        """
        return pulumi.get(self, "sql_connectivity_update_password")

    @property
    @pulumi.getter(name="sqlConnectivityUpdateUsername")
    def sql_connectivity_update_username(self) -> pulumi.Output[Optional[str]]:
        """
        The SQL Server sysadmin login to create.
        """
        return pulumi.get(self, "sql_connectivity_update_username")

    @property
    @pulumi.getter(name="sqlLicenseType")
    def sql_license_type(self) -> pulumi.Output[str]:
        """
        The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit), `DR` (Disaster Recovery), and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sql_license_type")

    @property
    @pulumi.getter(name="storageConfiguration")
    def storage_configuration(self) -> pulumi.Output[Optional['outputs.VirtualMachineStorageConfiguration']]:
        """
        An `storage_configuration` block as defined below.
        """
        return pulumi.get(self, "storage_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

