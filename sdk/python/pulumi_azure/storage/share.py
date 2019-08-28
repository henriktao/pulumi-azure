# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Share(pulumi.CustomResource):
    acls: pulumi.Output[list]
    """
    One or more `acl` blocks as defined below.
    
      * `access_policies` (`list`)
    
        * `expiry` (`str`)
        * `permissions` (`str`)
        * `start` (`str`)
    
      * `id` (`str`) - The ID of the File Share.
    """
    metadata: pulumi.Output[dict]
    """
    A mapping of MetaData for this File Share.
    """
    name: pulumi.Output[str]
    """
    The name of the share. Must be unique within the storage account where the share is located.
    """
    quota: pulumi.Output[float]
    """
    The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB) for Standard storage accounts or 100 TB (102400 GB) for Premium storage accounts. Default is 5120.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the share. Changing this forces a new resource to be created.
    """
    storage_account_name: pulumi.Output[str]
    """
    Specifies the storage account in which to create the share.
    Changing this forces a new resource to be created.
    """
    url: pulumi.Output[str]
    """
    The URL of the File Share
    """
    def __init__(__self__, resource_name, opts=None, acls=None, metadata=None, name=None, quota=None, resource_group_name=None, storage_account_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a File Share within Azure Storage.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[dict] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[float] quota: The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB) for Standard storage accounts or 100 TB (102400 GB) for Premium storage accounts. Default is 5120.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the share. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        
        The **acls** object supports the following:
        
          * `access_policies` (`pulumi.Input[list]`)
        
            * `expiry` (`pulumi.Input[str]`)
            * `permissions` (`pulumi.Input[str]`)
            * `start` (`pulumi.Input[str]`)
        
          * `id` (`pulumi.Input[str]`) - The ID of the File Share.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_share.html.markdown.
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

            __props__['acls'] = acls
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['quota'] = quota
            __props__['resource_group_name'] = resource_group_name
            if storage_account_name is None:
                raise TypeError("Missing required property 'storage_account_name'")
            __props__['storage_account_name'] = storage_account_name
            __props__['url'] = None
        super(Share, __self__).__init__(
            'azure:storage/share:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, acls=None, metadata=None, name=None, quota=None, resource_group_name=None, storage_account_name=None, url=None):
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] acls: One or more `acl` blocks as defined below.
        :param pulumi.Input[dict] metadata: A mapping of MetaData for this File Share.
        :param pulumi.Input[str] name: The name of the share. Must be unique within the storage account where the share is located.
        :param pulumi.Input[float] quota: The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5 TB (5120 GB) for Standard storage accounts or 100 TB (102400 GB) for Premium storage accounts. Default is 5120.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the share. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_name: Specifies the storage account in which to create the share.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] url: The URL of the File Share
        
        The **acls** object supports the following:
        
          * `access_policies` (`pulumi.Input[list]`)
        
            * `expiry` (`pulumi.Input[str]`)
            * `permissions` (`pulumi.Input[str]`)
            * `start` (`pulumi.Input[str]`)
        
          * `id` (`pulumi.Input[str]`) - The ID of the File Share.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_share.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["acls"] = acls
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["quota"] = quota
        __props__["resource_group_name"] = resource_group_name
        __props__["storage_account_name"] = storage_account_name
        __props__["url"] = url
        return Share(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

