# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class HybridConnection(pulumi.CustomResource):
    app_service_name: pulumi.Output[str]
    """
    Specifies the name of the App Service. Changing this forces a new resource to be created.
    """
    hostname: pulumi.Output[str]
    """
    The hostname of the endpoint.
    """
    namespace_name: pulumi.Output[str]
    """
    The name of the Relay Namespace.
    """
    port: pulumi.Output[float]
    """
    The port of the endpoint.
    """
    relay_id: pulumi.Output[str]
    """
    The ID of the Service Bus Relay. Changing this forces a new resource to be created.
    """
    relay_name: pulumi.Output[str]
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
    """
    send_key_name: pulumi.Output[str]
    """
    The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
    """
    send_key_value: pulumi.Output[str]
    """
    The value of the Service Bus Primary Access key.
    """
    service_bus_namespace: pulumi.Output[str]
    """
    The name of the Service Bus namespace.
    """
    service_bus_suffix: pulumi.Output[str]
    """
    The suffix for the service bus endpoint.
    """
    def __init__(__self__, resource_name, opts=None, app_service_name=None, hostname=None, port=None, relay_id=None, resource_group_name=None, send_key_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an App Service Hybrid Connection for an existing App Service, Relay and Service Bus.

        ## Example Usage

        This example provisions an App Service, a Relay Hybrid Connection, and a Service Bus using their outputs to create the App Service Hybrid Connection.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_plan = azure.appservice.Plan("examplePlan",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku={
                "tier": "Standard",
                "size": "S1",
            })
        example_app_service = azure.appservice.AppService("exampleAppService",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            app_service_plan_id=example_plan.id)
        example_namespace = azure.relay.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Standard")
        example_hybrid_connection = azure.relay.HybridConnection("exampleHybridConnection",
            resource_group_name=example_resource_group.name,
            relay_namespace_name=example_namespace.name,
            user_metadata="examplemetadata")
        example_appservice_hybrid_connection_hybrid_connection = azure.appservice.HybridConnection("exampleAppservice/hybridConnectionHybridConnection",
            app_service_name=example_app_service.name,
            resource_group_name=example_resource_group.name,
            relay_id=example_hybrid_connection.id,
            hostname="testhostname.example",
            port=8080,
            send_key_name="exampleSharedAccessKey")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: Specifies the name of the App Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] hostname: The hostname of the endpoint.
        :param pulumi.Input[float] port: The port of the endpoint.
        :param pulumi.Input[str] relay_id: The ID of the Service Bus Relay. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] send_key_name: The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
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

            if app_service_name is None:
                raise TypeError("Missing required property 'app_service_name'")
            __props__['app_service_name'] = app_service_name
            if hostname is None:
                raise TypeError("Missing required property 'hostname'")
            __props__['hostname'] = hostname
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            if relay_id is None:
                raise TypeError("Missing required property 'relay_id'")
            __props__['relay_id'] = relay_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['send_key_name'] = send_key_name
            __props__['namespace_name'] = None
            __props__['relay_name'] = None
            __props__['send_key_value'] = None
            __props__['service_bus_namespace'] = None
            __props__['service_bus_suffix'] = None
        super(HybridConnection, __self__).__init__(
            'azure:appservice/hybridConnection:HybridConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_service_name=None, hostname=None, namespace_name=None, port=None, relay_id=None, relay_name=None, resource_group_name=None, send_key_name=None, send_key_value=None, service_bus_namespace=None, service_bus_suffix=None):
        """
        Get an existing HybridConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: Specifies the name of the App Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] hostname: The hostname of the endpoint.
        :param pulumi.Input[str] namespace_name: The name of the Relay Namespace.
        :param pulumi.Input[float] port: The port of the endpoint.
        :param pulumi.Input[str] relay_id: The ID of the Service Bus Relay. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the App Service.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] send_key_name: The name of the Service Bus key which has Send permissions. Defaults to `RootManageSharedAccessKey`.
        :param pulumi.Input[str] send_key_value: The value of the Service Bus Primary Access key.
        :param pulumi.Input[str] service_bus_namespace: The name of the Service Bus namespace.
        :param pulumi.Input[str] service_bus_suffix: The suffix for the service bus endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_service_name"] = app_service_name
        __props__["hostname"] = hostname
        __props__["namespace_name"] = namespace_name
        __props__["port"] = port
        __props__["relay_id"] = relay_id
        __props__["relay_name"] = relay_name
        __props__["resource_group_name"] = resource_group_name
        __props__["send_key_name"] = send_key_name
        __props__["send_key_value"] = send_key_value
        __props__["service_bus_namespace"] = service_bus_namespace
        __props__["service_bus_suffix"] = service_bus_suffix
        return HybridConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
