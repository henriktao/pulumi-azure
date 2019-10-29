# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SharedImage(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of this Shared Image.
    """
    eula: pulumi.Output[str]
    """
    The End User Licence Agreement for the Shared Image.
    """
    gallery_name: pulumi.Output[str]
    """
    Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
    """
    identifier: pulumi.Output[dict]
    """
    An `identifier` block as defined below.
    
      * `offer` (`str`)
      * `publisher` (`str`)
      * `sku` (`str`)
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Shared Image. Changing this forces a new resource to be created.
    """
    os_type: pulumi.Output[str]
    """
    The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
    """
    privacy_statement_uri: pulumi.Output[str]
    """
    The URI containing the Privacy Statement associated with this Shared Image.
    """
    release_note_uri: pulumi.Output[str]
    """
    The URI containing the Release Notes associated with this Shared Image.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Shared Image.
    """
    def __init__(__self__, resource_name, opts=None, description=None, eula=None, gallery_name=None, identifier=None, location=None, name=None, os_type=None, privacy_statement_uri=None, release_note_uri=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Shared Image within a Shared Image Gallery.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of this Shared Image.
        :param pulumi.Input[str] eula: The End User Licence Agreement for the Shared Image.
        :param pulumi.Input[str] gallery_name: Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] identifier: An `identifier` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
        :param pulumi.Input[str] privacy_statement_uri: The URI containing the Privacy Statement associated with this Shared Image.
        :param pulumi.Input[str] release_note_uri: The URI containing the Release Notes associated with this Shared Image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Shared Image.
        
        The **identifier** object supports the following:
        
          * `offer` (`pulumi.Input[str]`)
          * `publisher` (`pulumi.Input[str]`)
          * `sku` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/shared_image.html.markdown.
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

            __props__['description'] = description
            __props__['eula'] = eula
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if identifier is None:
                raise TypeError("Missing required property 'identifier'")
            __props__['identifier'] = identifier
            __props__['location'] = location
            __props__['name'] = name
            if os_type is None:
                raise TypeError("Missing required property 'os_type'")
            __props__['os_type'] = os_type
            __props__['privacy_statement_uri'] = privacy_statement_uri
            __props__['release_note_uri'] = release_note_uri
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(SharedImage, __self__).__init__(
            'azure:compute/sharedImage:SharedImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, eula=None, gallery_name=None, identifier=None, location=None, name=None, os_type=None, privacy_statement_uri=None, release_note_uri=None, resource_group_name=None, tags=None):
        """
        Get an existing SharedImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of this Shared Image.
        :param pulumi.Input[str] eula: The End User Licence Agreement for the Shared Image.
        :param pulumi.Input[str] gallery_name: Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] identifier: An `identifier` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Shared Image. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
        :param pulumi.Input[str] privacy_statement_uri: The URI containing the Privacy Statement associated with this Shared Image.
        :param pulumi.Input[str] release_note_uri: The URI containing the Release Notes associated with this Shared Image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Shared Image.
        
        The **identifier** object supports the following:
        
          * `offer` (`pulumi.Input[str]`)
          * `publisher` (`pulumi.Input[str]`)
          * `sku` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/shared_image.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["eula"] = eula
        __props__["gallery_name"] = gallery_name
        __props__["identifier"] = identifier
        __props__["location"] = location
        __props__["name"] = name
        __props__["os_type"] = os_type
        __props__["privacy_statement_uri"] = privacy_statement_uri
        __props__["release_note_uri"] = release_note_uri
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return SharedImage(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

