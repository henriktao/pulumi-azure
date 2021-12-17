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

__all__ = ['DatasetPostgresqlArgs', 'DatasetPostgresql']

@pulumi.input_type
class DatasetPostgresqlArgs:
    def __init__(__self__, *,
                 linked_service_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_id: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatasetPostgresql resource.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] data_factory_id: The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]] schema_columns: A `schema_column` block as defined below.
        :param pulumi.Input[str] table_name: The table name of the Data Factory Dataset PostgreSQL.
        """
        pulumi.set(__self__, "linked_service_name", linked_service_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if data_factory_id is not None:
            pulumi.set(__self__, "data_factory_id", data_factory_id)
        if data_factory_name is not None:
            warnings.warn("""`data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""", DeprecationWarning)
            pulumi.log.warn("""data_factory_name is deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""")
        if data_factory_name is not None:
            pulumi.set(__self__, "data_factory_name", data_factory_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if schema_columns is not None:
            pulumi.set(__self__, "schema_columns", schema_columns)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> pulumi.Input[str]:
        """
        The Data Factory Linked Service name in which to associate the Dataset with.
        """
        return pulumi.get(self, "linked_service_name")

    @linked_service_name.setter
    def linked_service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "linked_service_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="dataFactoryId")
    def data_factory_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_id")

    @data_factory_id.setter
    def data_factory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_factory_id", value)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="schemaColumns")
    def schema_columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]]:
        """
        A `schema_column` block as defined below.
        """
        return pulumi.get(self, "schema_columns")

    @schema_columns.setter
    def schema_columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]]):
        pulumi.set(self, "schema_columns", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The table name of the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


@pulumi.input_type
class _DatasetPostgresqlState:
    def __init__(__self__, *,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_id: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 linked_service_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatasetPostgresql resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] data_factory_id: The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        :param pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]] schema_columns: A `schema_column` block as defined below.
        :param pulumi.Input[str] table_name: The table name of the Data Factory Dataset PostgreSQL.
        """
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if data_factory_id is not None:
            pulumi.set(__self__, "data_factory_id", data_factory_id)
        if data_factory_name is not None:
            warnings.warn("""`data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""", DeprecationWarning)
            pulumi.log.warn("""data_factory_name is deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""")
        if data_factory_name is not None:
            pulumi.set(__self__, "data_factory_name", data_factory_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if linked_service_name is not None:
            pulumi.set(__self__, "linked_service_name", linked_service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if schema_columns is not None:
            pulumi.set(__self__, "schema_columns", schema_columns)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="dataFactoryId")
    def data_factory_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_id")

    @data_factory_id.setter
    def data_factory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_factory_id", value)

    @property
    @pulumi.getter(name="dataFactoryName")
    def data_factory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_name")

    @data_factory_name.setter
    def data_factory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_factory_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Data Factory Linked Service name in which to associate the Dataset with.
        """
        return pulumi.get(self, "linked_service_name")

    @linked_service_name.setter
    def linked_service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "linked_service_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="schemaColumns")
    def schema_columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]]:
        """
        A `schema_column` block as defined below.
        """
        return pulumi.get(self, "schema_columns")

    @schema_columns.setter
    def schema_columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetPostgresqlSchemaColumnArgs']]]]):
        pulumi.set(self, "schema_columns", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The table name of the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


class DatasetPostgresql(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_id: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 linked_service_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetPostgresqlSchemaColumnArgs']]]]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a PostgreSQL Dataset inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_postgresql = azure.datafactory.LinkedServicePostgresql("exampleLinkedServicePostgresql",
            resource_group_name=example_resource_group.name,
            data_factory_id=example_factory.id,
            connection_string="Host=example;Port=5432;Database=example;UID=example;EncryptionMethod=0;Password=example")
        example_dataset_postgresql = azure.datafactory.DatasetPostgresql("exampleDatasetPostgresql",
            resource_group_name=example_resource_group.name,
            data_factory_id=example_factory.id,
            linked_service_name=example_linked_service_postgresql.name)
        ```

        ## Import

        Data Factory PostgreSQL Datasets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/datasetPostgresql:DatasetPostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] data_factory_id: The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetPostgresqlSchemaColumnArgs']]]] schema_columns: A `schema_column` block as defined below.
        :param pulumi.Input[str] table_name: The table name of the Data Factory Dataset PostgreSQL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetPostgresqlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a PostgreSQL Dataset inside a Azure Data Factory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_factory = azure.datafactory.Factory("exampleFactory",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_linked_service_postgresql = azure.datafactory.LinkedServicePostgresql("exampleLinkedServicePostgresql",
            resource_group_name=example_resource_group.name,
            data_factory_id=example_factory.id,
            connection_string="Host=example;Port=5432;Database=example;UID=example;EncryptionMethod=0;Password=example")
        example_dataset_postgresql = azure.datafactory.DatasetPostgresql("exampleDatasetPostgresql",
            resource_group_name=example_resource_group.name,
            data_factory_id=example_factory.id,
            linked_service_name=example_linked_service_postgresql.name)
        ```

        ## Import

        Data Factory PostgreSQL Datasets can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:datafactory/datasetPostgresql:DatasetPostgresql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
        ```

        :param str resource_name: The name of the resource.
        :param DatasetPostgresqlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetPostgresqlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 data_factory_id: Optional[pulumi.Input[str]] = None,
                 data_factory_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 linked_service_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetPostgresqlSchemaColumnArgs']]]]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = DatasetPostgresqlArgs.__new__(DatasetPostgresqlArgs)

            __props__.__dict__["additional_properties"] = additional_properties
            __props__.__dict__["annotations"] = annotations
            __props__.__dict__["data_factory_id"] = data_factory_id
            if data_factory_name is not None and not opts.urn:
                warnings.warn("""`data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""", DeprecationWarning)
                pulumi.log.warn("""data_factory_name is deprecated: `data_factory_name` is deprecated in favour of `data_factory_id` and will be removed in version 3.0 of the AzureRM provider""")
            __props__.__dict__["data_factory_name"] = data_factory_name
            __props__.__dict__["description"] = description
            __props__.__dict__["folder"] = folder
            if linked_service_name is None and not opts.urn:
                raise TypeError("Missing required property 'linked_service_name'")
            __props__.__dict__["linked_service_name"] = linked_service_name
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["schema_columns"] = schema_columns
            __props__.__dict__["table_name"] = table_name
        super(DatasetPostgresql, __self__).__init__(
            'azure:datafactory/datasetPostgresql:DatasetPostgresql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            data_factory_id: Optional[pulumi.Input[str]] = None,
            data_factory_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            linked_service_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            schema_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetPostgresqlSchemaColumnArgs']]]]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'DatasetPostgresql':
        """
        Get an existing DatasetPostgresql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] additional_properties: A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] annotations: List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] data_factory_id: The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] data_factory_name: The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
        :param pulumi.Input[str] description: The description for the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] folder: The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        :param pulumi.Input[str] linked_service_name: The Data Factory Linked Service name in which to associate the Dataset with.
        :param pulumi.Input[str] name: Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetPostgresqlSchemaColumnArgs']]]] schema_columns: A `schema_column` block as defined below.
        :param pulumi.Input[str] table_name: The table name of the Data Factory Dataset PostgreSQL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatasetPostgresqlState.__new__(_DatasetPostgresqlState)

        __props__.__dict__["additional_properties"] = additional_properties
        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["data_factory_id"] = data_factory_id
        __props__.__dict__["data_factory_name"] = data_factory_name
        __props__.__dict__["description"] = description
        __props__.__dict__["folder"] = folder
        __props__.__dict__["linked_service_name"] = linked_service_name
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["schema_columns"] = schema_columns
        __props__.__dict__["table_name"] = table_name
        return DatasetPostgresql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of additional properties to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags that can be used for describing the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="dataFactoryId")
    def data_factory_id(self) -> pulumi.Output[str]:
        """
        The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
        """
        return pulumi.get(self, "data_factory_id")

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
        The description for the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[Optional[str]]:
        """
        The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> pulumi.Output[str]:
        """
        The Data Factory Linked Service name in which to associate the Dataset with.
        """
        return pulumi.get(self, "linked_service_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Data Factory Dataset PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of parameters to associate with the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Data Factory Dataset PostgreSQL. Changing this forces a new resource
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="schemaColumns")
    def schema_columns(self) -> pulumi.Output[Optional[Sequence['outputs.DatasetPostgresqlSchemaColumn']]]:
        """
        A `schema_column` block as defined below.
        """
        return pulumi.get(self, "schema_columns")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[Optional[str]]:
        """
        The table name of the Data Factory Dataset PostgreSQL.
        """
        return pulumi.get(self, "table_name")

