# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SharedImageVersion(pulumi.CustomResource):
    exclude_from_latest: pulumi.Output[bool]
    """
    Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
    """
    gallery_name: pulumi.Output[str]
    """
    The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
    """
    image_name: pulumi.Output[str]
    """
    The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
    """
    managed_image_id: pulumi.Output[str]
    """
    The ID of the Managed Image which should be used for this Shared Image Version. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A collection of tags which should be applied to this resource.
    """
    target_regions: pulumi.Output[list]
    """
    One or more `target_region` blocks as documented below.
    
      * `name` (`str`) - The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
      * `regionalReplicaCount` (`float`)
    """
    def __init__(__self__, resource_name, opts=None, exclude_from_latest=None, gallery_name=None, image_name=None, location=None, managed_image_id=None, name=None, resource_group_name=None, tags=None, target_regions=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Version of a Shared Image within a Shared Image Gallery.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A collection of tags which should be applied to this resource.
        :param pulumi.Input[list] target_regions: One or more `target_region` blocks as documented below.
        
        The **target_regions** object supports the following:
        
          * `name` (`pulumi.Input[str]`) - The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
          * `regionalReplicaCount` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/shared_image_version.html.markdown.
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

            __props__['exclude_from_latest'] = exclude_from_latest
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if image_name is None:
                raise TypeError("Missing required property 'image_name'")
            __props__['image_name'] = image_name
            __props__['location'] = location
            if managed_image_id is None:
                raise TypeError("Missing required property 'managed_image_id'")
            __props__['managed_image_id'] = managed_image_id
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if target_regions is None:
                raise TypeError("Missing required property 'target_regions'")
            __props__['target_regions'] = target_regions
        super(SharedImageVersion, __self__).__init__(
            'azure:compute/sharedImageVersion:SharedImageVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, exclude_from_latest=None, gallery_name=None, image_name=None, location=None, managed_image_id=None, name=None, resource_group_name=None, tags=None, target_regions=None):
        """
        Get an existing SharedImageVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclude_from_latest: Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] image_name: The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_image_id: The ID of the Managed Image which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A collection of tags which should be applied to this resource.
        :param pulumi.Input[list] target_regions: One or more `target_region` blocks as documented below.
        
        The **target_regions** object supports the following:
        
          * `name` (`pulumi.Input[str]`) - The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
          * `regionalReplicaCount` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/shared_image_version.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["exclude_from_latest"] = exclude_from_latest
        __props__["gallery_name"] = gallery_name
        __props__["image_name"] = image_name
        __props__["location"] = location
        __props__["managed_image_id"] = managed_image_id
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["target_regions"] = target_regions
        return SharedImageVersion(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

