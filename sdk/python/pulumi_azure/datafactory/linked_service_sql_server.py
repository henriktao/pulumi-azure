# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LinkedServiceSqlServer']


class LinkedServiceSqlServer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_runtime_name: Optional[pulumi.Input[str]] = None,
                 key_vault_password: Optional[pulumi.Input[pulumi.InputType['LinkedServiceSqlServerKeyVaultPasswordArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Linked Service (connection) between a SQL Server and Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_sql_server = azure.datafactory.LinkedServiceSqlServer("exampleLinkedServiceSqlServer",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            connection_string="Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;Password=test")
        ```
        ### With Password In Key Vault

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="northeurope")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_key_vault = azure.datafactory.LinkedServiceKeyVault("exampleLinkedServiceKeyVault",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            key_vault_id=example_key_vault.id)
        example_linked_service_sql_server = azure.datafactory.LinkedServiceSqlServer("exampleLinkedServiceSqlServer",
            resource_group_name=example_resource_group.name,
            data_factory_name=example_factory.name,
            connection_string="Integrated Security=False;Data Source=test;Initial Catalog=test;User ID=test;",
            key_vault_password=azure.datafactory.LinkedServiceSqlServerKeyVaultPasswordArgs(
                linked_service_name=example_linked_service_key_vault.name,
                secret_name="secret",
            ))
        ```

        ## Import

        Data Factory SQL Server Linked Service's can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] connection_string: The connection string in which to authenticate with the SQL Server.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[pulumi.InputType['LinkedServiceSqlServerKeyVaultPasswordArgs']] key_vault_password: A `key_vault_password` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['additional_properties'] = additional_properties
            __props__['annotations'] = annotations
            if connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string'")
            __props__['connection_string'] = connection_string
            if data_factory_name is None and not opts.urn:
                raise TypeError("Missing required property 'data_factory_name'")
            __props__['data_factory_name'] = data_factory_name
            __props__['description'] = description
            __props__['integration_runtime_name'] = integration_runtime_name
            __props__['key_vault_password'] = key_vault_password
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(LinkedServiceSqlServer, __self__).__init__(
            'azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            integration_runtime_name: Optional[pulumi.Input[str]] = None,
            key_vault_password: Optional[pulumi.Input[pulumi.InputType['LinkedServiceSqlServerKeyVaultPasswordArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'LinkedServiceSqlServer':
        """
        Get an existing LinkedServiceSqlServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] connection_string: The connection string in which to authenticate with the SQL Server.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] integration_runtime_name: The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[pulumi.InputType['LinkedServiceSqlServerKeyVaultPasswordArgs']] key_vault_password: A `key_vault_password` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Linked Service SQL Server.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_properties"] = additional_properties
        __props__["annotations"] = annotations
        __props__["connection_string"] = connection_string
        __props__["data_factory_name"] = data_factory_name
        __props__["description"] = description
        __props__["integration_runtime_name"] = integration_runtime_name
        __props__["key_vault_password"] = key_vault_password
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["resource_group_name"] = resource_group_name
        return LinkedServiceSqlServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional properties to associate with the Data Factory Linked Service SQL Server.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Linked Service SQL Server.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The connection string in which to authenticate with the SQL Server.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> pulumi.Output[str]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Data Factory Linked Service SQL Server.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationRuntimeName")
    def integration_runtime_name(self) -> pulumi.Output[Optional[str]]:
        """
        The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
        """
        return pulumi.get(self, "integration_runtime_name")

    @property
    @pulumi.getter(name="keyVaultPassword")
    def key_vault_password(self) -> pulumi.Output[Optional['outputs.LinkedServiceSqlServerKeyVaultPassword']]:
        """
        A `key_vault_password` block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
        """
        return pulumi.get(self, "key_vault_password")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Linked Service SQL Server.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

