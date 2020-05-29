# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TableEntity(pulumi.CustomResource):
    entity: pulumi.Output[dict]
    """
    A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
    """
    partition_key: pulumi.Output[str]
    """
    The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
    """
    row_key: pulumi.Output[str]
    """
    The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
    """
    storage_account_name: pulumi.Output[str]
    """
    Specifies the storage account in which to create the storage table entity.
    Changing this forces a new resource to be created.
    """
    table_name: pulumi.Output[str]
    """
    The name of the storage table in which to create the storage table entity.
    Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, entity=None, partition_key=None, row_key=None, storage_account_name=None, table_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Entity within a Table in an Azure Storage Account.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westus")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_table = azure.storage.Table("exampleTable",
            resource_group_name=example_resource_group.name,
            storage_account_name=example_account.name)
        example_table_entity = azure.storage.TableEntity("exampleTableEntity",
            storage_account_name=example_account.name,
            table_name=example_table.name,
            partition_key="examplepartition",
            row_key="exmamplerow",
            entity={
                "example": "example",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] entity: A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        :param pulumi.Input[str] partition_key: The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] row_key: The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage table entity.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: The name of the storage table in which to create the storage table entity.
               Changing this forces a new resource to be created.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if entity is None:
                raise TypeError("Missing required property 'entity'")
            __props__['entity'] = entity
            if partition_key is None:
                raise TypeError("Missing required property 'partition_key'")
            __props__['partition_key'] = partition_key
            if row_key is None:
                raise TypeError("Missing required property 'row_key'")
            __props__['row_key'] = row_key
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if table_name is None:
                raise TypeError("Missing required property 'table_name'")
            __props__['table_name'] = table_name
        super(TableEntity, __self__).__init__(
            'azure:storage/tableEntity:TableEntity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, entity=None, partition_key=None, row_key=None, storage_account_name=None, table_name=None):
        """
        Get an existing TableEntity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] entity: A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        :param pulumi.Input[str] partition_key: The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] row_key: The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage table entity.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] table_name: The name of the storage table in which to create the storage table entity.
               Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["entity"] = entity
        __props__["partition_key"] = partition_key
        __props__["row_key"] = row_key
        __props__["storage_account_name"] = storage_account_name
        __props__["table_name"] = table_name
        return TableEntity(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

