# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class IoTHub(pulumi.CustomResource):
    """
    Manages an IotHub
    """
    def __init__(__self__, __name__, __opts__=None, endpoints=None, location=None, name=None, resource_group_name=None, routes=None, sku=None, tags=None):
        """Create a IoTHub resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['endpoints'] = endpoints

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['routes'] = routes

        if not sku:
            raise TypeError('Missing required property sku')
        __props__['sku'] = sku

        __props__['tags'] = tags

        __props__['event_hub_events_endpoint'] = None
        __props__['event_hub_events_path'] = None
        __props__['event_hub_operations_endpoint'] = None
        __props__['event_hub_operations_path'] = None
        __props__['hostname'] = None
        __props__['shared_access_policies'] = None
        __props__['type'] = None

        super(IoTHub, __self__).__init__(
            'azure:iot/ioTHub:IoTHub',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

