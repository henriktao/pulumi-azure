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

__all__ = ['IotHubDpsArgs', 'IotHubDps']

@pulumi.input_type
class IotHubDpsArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 sku: pulumi.Input['IotHubDpsSkuArgs'],
                 allocation_policy: Optional[pulumi.Input[str]] = None,
                 linked_hubs: Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a IotHubDps resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input['IotHubDpsSkuArgs'] sku: A `sku` block as defined below.
        :param pulumi.Input[str] allocation_policy: The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        :param pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if allocation_policy is not None:
            pulumi.set(__self__, "allocation_policy", allocation_policy)
        if linked_hubs is not None:
            pulumi.set(__self__, "linked_hubs", linked_hubs)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input['IotHubDpsSkuArgs']:
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['IotHubDpsSkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        """
        return pulumi.get(self, "allocation_policy")

    @allocation_policy.setter
    def allocation_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_policy", value)

    @property
    @pulumi.getter(name="linkedHubs")
    def linked_hubs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]]:
        """
        A `linked_hub` block as defined below.
        """
        return pulumi.get(self, "linked_hubs")

    @linked_hubs.setter
    def linked_hubs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]]):
        pulumi.set(self, "linked_hubs", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _IotHubDpsState:
    def __init__(__self__, *,
                 allocation_policy: Optional[pulumi.Input[str]] = None,
                 device_provisioning_host_name: Optional[pulumi.Input[str]] = None,
                 id_scope: Optional[pulumi.Input[str]] = None,
                 linked_hubs: Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_operations_host_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input['IotHubDpsSkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering IotHubDps resources.
        :param pulumi.Input[str] allocation_policy: The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        :param pulumi.Input[str] device_provisioning_host_name: The device endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input[str] id_scope: The unique identifier of the IoT Device Provisioning Service.
        :param pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_operations_host_name: The service endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input['IotHubDpsSkuArgs'] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if allocation_policy is not None:
            pulumi.set(__self__, "allocation_policy", allocation_policy)
        if device_provisioning_host_name is not None:
            pulumi.set(__self__, "device_provisioning_host_name", device_provisioning_host_name)
        if id_scope is not None:
            pulumi.set(__self__, "id_scope", id_scope)
        if linked_hubs is not None:
            pulumi.set(__self__, "linked_hubs", linked_hubs)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if service_operations_host_name is not None:
            pulumi.set(__self__, "service_operations_host_name", service_operations_host_name)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        """
        return pulumi.get(self, "allocation_policy")

    @allocation_policy.setter
    def allocation_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_policy", value)

    @property
    @pulumi.getter(name="deviceProvisioningHostName")
    def device_provisioning_host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The device endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "device_provisioning_host_name")

    @device_provisioning_host_name.setter
    def device_provisioning_host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_provisioning_host_name", value)

    @property
    @pulumi.getter(name="idScope")
    def id_scope(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "id_scope")

    @id_scope.setter
    def id_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id_scope", value)

    @property
    @pulumi.getter(name="linkedHubs")
    def linked_hubs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]]:
        """
        A `linked_hub` block as defined below.
        """
        return pulumi.get(self, "linked_hubs")

    @linked_hubs.setter
    def linked_hubs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IotHubDpsLinkedHubArgs']]]]):
        pulumi.set(self, "linked_hubs", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceOperationsHostName")
    def service_operations_host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The service endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "service_operations_host_name")

    @service_operations_host_name.setter
    def service_operations_host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_operations_host_name", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['IotHubDpsSkuArgs']]:
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['IotHubDpsSkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class IotHubDps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_policy: Optional[pulumi.Input[str]] = None,
                 linked_hubs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an IotHub Device Provisioning Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_iot_hub_dps = azure.iot.IotHubDps("exampleIotHubDps",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            allocation_policy="Hashed",
            sku=azure.iot.IotHubDpsSkuArgs(
                name="S1",
                capacity=1,
            ))
        ```

        ## Import

        IoT Device Provisioning Service can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/iotHubDps:IotHubDps example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_policy: The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IotHubDpsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an IotHub Device Provisioning Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_iot_hub_dps = azure.iot.IotHubDps("exampleIotHubDps",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            allocation_policy="Hashed",
            sku=azure.iot.IotHubDpsSkuArgs(
                name="S1",
                capacity=1,
            ))
        ```

        ## Import

        IoT Device Provisioning Service can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:iot/iotHubDps:IotHubDps example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/provisioningServices/example
        ```

        :param str resource_name: The name of the resource.
        :param IotHubDpsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IotHubDpsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_policy: Optional[pulumi.Input[str]] = None,
                 linked_hubs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = IotHubDpsArgs.__new__(IotHubDpsArgs)

            __props__.__dict__["allocation_policy"] = allocation_policy
            __props__.__dict__["linked_hubs"] = linked_hubs
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["device_provisioning_host_name"] = None
            __props__.__dict__["id_scope"] = None
            __props__.__dict__["service_operations_host_name"] = None
        super(IotHubDps, __self__).__init__(
            'azure:iot/iotHubDps:IotHubDps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_policy: Optional[pulumi.Input[str]] = None,
            device_provisioning_host_name: Optional[pulumi.Input[str]] = None,
            id_scope: Optional[pulumi.Input[str]] = None,
            linked_hubs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_operations_host_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'IotHubDps':
        """
        Get an existing IotHubDps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_policy: The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        :param pulumi.Input[str] device_provisioning_host_name: The device endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input[str] id_scope: The unique identifier of the IoT Device Provisioning Service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IotHubDpsLinkedHubArgs']]]] linked_hubs: A `linked_hub` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_operations_host_name: The service endpoint of the IoT Device Provisioning Service.
        :param pulumi.Input[pulumi.InputType['IotHubDpsSkuArgs']] sku: A `sku` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IotHubDpsState.__new__(_IotHubDpsState)

        __props__.__dict__["allocation_policy"] = allocation_policy
        __props__.__dict__["device_provisioning_host_name"] = device_provisioning_host_name
        __props__.__dict__["id_scope"] = id_scope
        __props__.__dict__["linked_hubs"] = linked_hubs
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["service_operations_host_name"] = service_operations_host_name
        __props__.__dict__["sku"] = sku
        __props__.__dict__["tags"] = tags
        return IotHubDps(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The allocation policy of the IoT Device Provisioning Service (`Hashed`, `GeoLatency` or `Static`). Defaults to `Hashed`.
        """
        return pulumi.get(self, "allocation_policy")

    @property
    @pulumi.getter(name="deviceProvisioningHostName")
    def device_provisioning_host_name(self) -> pulumi.Output[str]:
        """
        The device endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "device_provisioning_host_name")

    @property
    @pulumi.getter(name="idScope")
    def id_scope(self) -> pulumi.Output[str]:
        """
        The unique identifier of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "id_scope")

    @property
    @pulumi.getter(name="linkedHubs")
    def linked_hubs(self) -> pulumi.Output[Optional[Sequence['outputs.IotHubDpsLinkedHub']]]:
        """
        A `linked_hub` block as defined below.
        """
        return pulumi.get(self, "linked_hubs")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceOperationsHostName")
    def service_operations_host_name(self) -> pulumi.Output[str]:
        """
        The service endpoint of the IoT Device Provisioning Service.
        """
        return pulumi.get(self, "service_operations_host_name")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.IotHubDpsSku']:
        """
        A `sku` block as defined below.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

