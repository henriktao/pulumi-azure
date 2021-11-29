# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EndpointServicebusQueueArgs', 'EndpointServicebusQueue']

@pulumi.input_type
class EndpointServicebusQueueArgs:
    def __init__(__self__, *,
                 connection_string: pulumi.Input[str],
                 iothub_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointServicebusQueue resource.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "iothub_name", iothub_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Input[str]:
        """
        The connection string for the endpoint.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter(name="iothubName")
    def iothub_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "iothub_name")

    @iothub_name.setter
    def iothub_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "iothub_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EndpointServicebusQueueState:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 iothub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointServicebusQueue resources.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if iothub_name is not None:
            pulumi.set(__self__, "iothub_name", iothub_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The connection string for the endpoint.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter(name="iothubName")
    def iothub_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iothub_name")

    @iothub_name.setter
    def iothub_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iothub_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)


class EndpointServicebusQueue(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 iothub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an IotHub ServiceBus Queue Endpoint

        > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard")
        example_queue = azure.servicebus.Queue("exampleQueue",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_queue_authorization_rule = azure.servicebus.QueueAuthorizationRule("exampleQueueAuthorizationRule",
            namespace_name=example_namespace.name,
            queue_name=example_queue.name,
            resource_group_name=example_resource_group.name,
            listen=False,
            send=True,
            manage=False)
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="B1",
                capacity=1,
            ),
            tags={
                "purpose": "example",
            })
        example_endpoint_servicebus_queue = azure.iot.EndpointServicebusQueue("exampleEndpointServicebusQueue",
            resource_group_name=example_resource_group.name,
            iothub_name=example_io_t_hub.name,
            connection_string=example_queue_authorization_rule.primary_connection_string)
        ```

        ## Import

        IoTHub ServiceBus Queue Endpoint can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/endpointServicebusQueue:EndpointServicebusQueue servicebus_queue1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/servicebusqueue_endpoint1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointServicebusQueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an IotHub ServiceBus Queue Endpoint

        > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard")
        example_queue = azure.servicebus.Queue("exampleQueue",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        example_queue_authorization_rule = azure.servicebus.QueueAuthorizationRule("exampleQueueAuthorizationRule",
            namespace_name=example_namespace.name,
            queue_name=example_queue.name,
            resource_group_name=example_resource_group.name,
            listen=False,
            send=True,
            manage=False)
        example_io_t_hub = azure.iot.IoTHub("exampleIoTHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku=azure.iot.IoTHubSkuArgs(
                name="B1",
                capacity=1,
            ),
            tags={
                "purpose": "example",
            })
        example_endpoint_servicebus_queue = azure.iot.EndpointServicebusQueue("exampleEndpointServicebusQueue",
            resource_group_name=example_resource_group.name,
            iothub_name=example_io_t_hub.name,
            connection_string=example_queue_authorization_rule.primary_connection_string)
        ```

        ## Import

        IoTHub ServiceBus Queue Endpoint can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/endpointServicebusQueue:EndpointServicebusQueue servicebus_queue1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/servicebusqueue_endpoint1
        ```

        :param str resource_name: The name of the resource.
        :param EndpointServicebusQueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointServicebusQueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 iothub_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = EndpointServicebusQueueArgs.__new__(EndpointServicebusQueueArgs)

            if connection_string is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string'")
            __props__.__dict__["connection_string"] = connection_string
            if iothub_name is None and not opts.urn:
                raise TypeError("Missing required property 'iothub_name'")
            __props__.__dict__["iothub_name"] = iothub_name
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
        super(EndpointServicebusQueue, __self__).__init__(
            'azure:iot/endpointServicebusQueue:EndpointServicebusQueue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            iothub_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'EndpointServicebusQueue':
        """
        Get an existing EndpointServicebusQueue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The connection string for the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointServicebusQueueState.__new__(_EndpointServicebusQueueState)

        __props__.__dict__["connection_string"] = connection_string
        __props__.__dict__["iothub_name"] = iothub_name
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        return EndpointServicebusQueue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The connection string for the endpoint.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="iothubName")
    def iothub_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "iothub_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_group_name")

