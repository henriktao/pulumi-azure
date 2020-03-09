# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Blob(pulumi.CustomResource):
    access_tier: pulumi.Output[str]
    """
    The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
    """
    content_type: pulumi.Output[str]
    """
    The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
    """
    metadata: pulumi.Output[dict]
    """
    A map of custom blob metadata.
    """
    name: pulumi.Output[str]
    """
    The name of the storage blob. Must be unique within the storage container the blob is located.
    """
    parallelism: pulumi.Output[float]
    """
    The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
    """
    size: pulumi.Output[float]
    """
    Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
    """
    source: pulumi.Output[Union[pulumi.Asset, pulumi.Archive]]
    """
    An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
    """
    source_content: pulumi.Output[str]
    """
    The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
    """
    source_uri: pulumi.Output[str]
    """
    The URI of an existing blob, or a file in the Azure File service, to use as the source contents
    for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
    """
    storage_account_name: pulumi.Output[str]
    """
    Specifies the storage account in which to create the storage container.
    Changing this forces a new resource to be created.
    """
    storage_container_name: pulumi.Output[str]
    """
    The name of the storage container in which this blob should be created.
    """
    type: pulumi.Output[str]
    """
    The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
    """
    url: pulumi.Output[str]
    """
    The URL of the blob
    """
    def __init__(__self__, resource_name, opts=None, access_tier=None, content_type=None, metadata=None, name=None, parallelism=None, size=None, source=None, source_content=None, source_uri=None, storage_account_name=None, storage_container_name=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Blob within a Storage Container.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_tier: The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
        :param pulumi.Input[str] content_type: The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
        :param pulumi.Input[dict] metadata: A map of custom blob metadata.
        :param pulumi.Input[str] name: The name of the storage blob. Must be unique within the storage container the blob is located.
        :param pulumi.Input[float] parallelism: The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
        :param pulumi.Input[float] size: Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
        :param pulumi.Input[str] source_content: The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
        :param pulumi.Input[str] source_uri: The URI of an existing blob, or a file in the Azure File service, to use as the source contents
               for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage container.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_name: The name of the storage container in which this blob should be created.
        :param pulumi.Input[str] type: The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_blob.html.markdown.
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

            __props__['access_tier'] = access_tier
            __props__['content_type'] = content_type
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['parallelism'] = parallelism
            __props__['size'] = size
            __props__['source'] = source
            __props__['source_content'] = source_content
            __props__['source_uri'] = source_uri
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            if storage_container_name is None:
                raise TypeError("Missing required property 'storage_container_name'")
            __props__['storage_container_name'] = storage_container_name
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['url'] = None
        super(Blob, __self__).__init__(
            'azure:storage/blob:Blob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_tier=None, content_type=None, metadata=None, name=None, parallelism=None, size=None, source=None, source_content=None, source_uri=None, storage_account_name=None, storage_container_name=None, type=None, url=None):
        """
        Get an existing Blob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_tier: The access tier of the storage blob. Possible values are `Archive`, `Cool` and `Hot`.
        :param pulumi.Input[str] content_type: The content type of the storage blob. Cannot be defined if `source_uri` is defined. Defaults to `application/octet-stream`.
        :param pulumi.Input[dict] metadata: A map of custom blob metadata.
        :param pulumi.Input[str] name: The name of the storage blob. Must be unique within the storage container the blob is located.
        :param pulumi.Input[float] parallelism: The number of workers per CPU core to run for concurrent uploads. Defaults to `8`.
        :param pulumi.Input[float] size: Used only for `page` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] source: An absolute path to a file on the local system. This field cannot be specified for Append blobs and cannot be specified if `source_content` or `source_uri` is specified.
        :param pulumi.Input[str] source_content: The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if `source` or `source_uri` is specified.
        :param pulumi.Input[str] source_uri: The URI of an existing blob, or a file in the Azure File service, to use as the source contents
               for the blob to be created. Changing this forces a new resource to be created. This field cannot be specified for Append blobs and cannot be specified if `source` or `source_content` is specified.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the storage container.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_name: The name of the storage container in which this blob should be created.
        :param pulumi.Input[str] type: The type of the storage blob to be created. Possible values are `Append`, `Block` or `Page`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] url: The URL of the blob

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_blob.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_tier"] = access_tier
        __props__["content_type"] = content_type
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["parallelism"] = parallelism
        __props__["size"] = size
        __props__["source"] = source
        __props__["source_content"] = source_content
        __props__["source_uri"] = source_uri
        __props__["storage_account_name"] = storage_account_name
        __props__["storage_container_name"] = storage_container_name
        __props__["type"] = type
        __props__["url"] = url
        return Blob(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

