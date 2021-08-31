# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamespaceDisasterRecoveryConfigArgs', 'NamespaceDisasterRecoveryConfig']

@pulumi.input_type
class NamespaceDisasterRecoveryConfigArgs:
    def __init__(__self__, *,
                 partner_namespace_id: pulumi.Input[str],
                 primary_namespace_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NamespaceDisasterRecoveryConfig resource.
        :param pulumi.Input[str] partner_namespace_id: The ID of the Service Bus Namespace to replicate to.
        :param pulumi.Input[str] primary_namespace_id: The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "partner_namespace_id", partner_namespace_id)
        pulumi.set(__self__, "primary_namespace_id", primary_namespace_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="partnerNamespaceId")
    def partner_namespace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Service Bus Namespace to replicate to.
        """
        return pulumi.get(self, "partner_namespace_id")

    @partner_namespace_id.setter
    def partner_namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "partner_namespace_id", value)

    @property
    @pulumi.getter(name="primaryNamespaceId")
    def primary_namespace_id(self) -> pulumi.Input[str]:
        """
        The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "primary_namespace_id")

    @primary_namespace_id.setter
    def primary_namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_namespace_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _NamespaceDisasterRecoveryConfigState:
    def __init__(__self__, *,
                 default_primary_key: Optional[pulumi.Input[str]] = None,
                 default_secondary_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partner_namespace_id: Optional[pulumi.Input[str]] = None,
                 primary_connection_string_alias: Optional[pulumi.Input[str]] = None,
                 primary_namespace_id: Optional[pulumi.Input[str]] = None,
                 secondary_connection_string_alias: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NamespaceDisasterRecoveryConfig resources.
        :param pulumi.Input[str] default_primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] name: Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partner_namespace_id: The ID of the Service Bus Namespace to replicate to.
        :param pulumi.Input[str] primary_connection_string_alias: The alias Primary Connection String for the ServiceBus Namespace.
        :param pulumi.Input[str] primary_namespace_id: The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string_alias: The alias Secondary Connection String for the ServiceBus Namespace
        """
        if default_primary_key is not None:
            pulumi.set(__self__, "default_primary_key", default_primary_key)
        if default_secondary_key is not None:
            pulumi.set(__self__, "default_secondary_key", default_secondary_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partner_namespace_id is not None:
            pulumi.set(__self__, "partner_namespace_id", partner_namespace_id)
        if primary_connection_string_alias is not None:
            pulumi.set(__self__, "primary_connection_string_alias", primary_connection_string_alias)
        if primary_namespace_id is not None:
            pulumi.set(__self__, "primary_namespace_id", primary_namespace_id)
        if secondary_connection_string_alias is not None:
            pulumi.set(__self__, "secondary_connection_string_alias", secondary_connection_string_alias)

    @property
    @pulumi.getter(name="defaultPrimaryKey")
    def default_primary_key(self) -> Optional[pulumi.Input[str]]:
        """
        The primary access key for the authorization rule `RootManageSharedAccessKey`.
        """
        return pulumi.get(self, "default_primary_key")

    @default_primary_key.setter
    def default_primary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_primary_key", value)

    @property
    @pulumi.getter(name="defaultSecondaryKey")
    def default_secondary_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        """
        return pulumi.get(self, "default_secondary_key")

    @default_secondary_key.setter
    def default_secondary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_secondary_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partnerNamespaceId")
    def partner_namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Service Bus Namespace to replicate to.
        """
        return pulumi.get(self, "partner_namespace_id")

    @partner_namespace_id.setter
    def partner_namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partner_namespace_id", value)

    @property
    @pulumi.getter(name="primaryConnectionStringAlias")
    def primary_connection_string_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias Primary Connection String for the ServiceBus Namespace.
        """
        return pulumi.get(self, "primary_connection_string_alias")

    @primary_connection_string_alias.setter
    def primary_connection_string_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_connection_string_alias", value)

    @property
    @pulumi.getter(name="primaryNamespaceId")
    def primary_namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "primary_namespace_id")

    @primary_namespace_id.setter
    def primary_namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_namespace_id", value)

    @property
    @pulumi.getter(name="secondaryConnectionStringAlias")
    def secondary_connection_string_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias Secondary Connection String for the ServiceBus Namespace
        """
        return pulumi.get(self, "secondary_connection_string_alias")

    @secondary_connection_string_alias.setter
    def secondary_connection_string_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_connection_string_alias", value)


class NamespaceDisasterRecoveryConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partner_namespace_id: Optional[pulumi.Input[str]] = None,
                 primary_namespace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Disaster Recovery Config for a Service Bus Namespace.

        > **NOTE:** Disaster Recovery Config is a Premium Sku only capability.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        primary = azure.servicebus.Namespace("primary",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Premium",
            capacity=1)
        secondary = azure.servicebus.Namespace("secondary",
            location="West US",
            resource_group_name=example_resource_group.name,
            sku="Premium",
            capacity=1)
        example_namespace_disaster_recovery_config = azure.servicebus.NamespaceDisasterRecoveryConfig("exampleNamespaceDisasterRecoveryConfig",
            primary_namespace_id=primary.id,
            partner_namespace_id=secondary.id)
        ```

        ## Import

        Service Bus DR configs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/namespaceDisasterRecoveryConfig:NamespaceDisasterRecoveryConfig config1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/disasterRecoveryConfigs/config1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partner_namespace_id: The ID of the Service Bus Namespace to replicate to.
        :param pulumi.Input[str] primary_namespace_id: The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceDisasterRecoveryConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Disaster Recovery Config for a Service Bus Namespace.

        > **NOTE:** Disaster Recovery Config is a Premium Sku only capability.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        primary = azure.servicebus.Namespace("primary",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Premium",
            capacity=1)
        secondary = azure.servicebus.Namespace("secondary",
            location="West US",
            resource_group_name=example_resource_group.name,
            sku="Premium",
            capacity=1)
        example_namespace_disaster_recovery_config = azure.servicebus.NamespaceDisasterRecoveryConfig("exampleNamespaceDisasterRecoveryConfig",
            primary_namespace_id=primary.id,
            partner_namespace_id=secondary.id)
        ```

        ## Import

        Service Bus DR configs can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:servicebus/namespaceDisasterRecoveryConfig:NamespaceDisasterRecoveryConfig config1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/disasterRecoveryConfigs/config1
        ```

        :param str resource_name: The name of the resource.
        :param NamespaceDisasterRecoveryConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceDisasterRecoveryConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partner_namespace_id: Optional[pulumi.Input[str]] = None,
                 primary_namespace_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = NamespaceDisasterRecoveryConfigArgs.__new__(NamespaceDisasterRecoveryConfigArgs)

            __props__.__dict__["name"] = name
            if partner_namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'partner_namespace_id'")
            __props__.__dict__["partner_namespace_id"] = partner_namespace_id
            if primary_namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'primary_namespace_id'")
            __props__.__dict__["primary_namespace_id"] = primary_namespace_id
            __props__.__dict__["default_primary_key"] = None
            __props__.__dict__["default_secondary_key"] = None
            __props__.__dict__["primary_connection_string_alias"] = None
            __props__.__dict__["secondary_connection_string_alias"] = None
        super(NamespaceDisasterRecoveryConfig, __self__).__init__(
            'azure:servicebus/namespaceDisasterRecoveryConfig:NamespaceDisasterRecoveryConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_primary_key: Optional[pulumi.Input[str]] = None,
            default_secondary_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partner_namespace_id: Optional[pulumi.Input[str]] = None,
            primary_connection_string_alias: Optional[pulumi.Input[str]] = None,
            primary_namespace_id: Optional[pulumi.Input[str]] = None,
            secondary_connection_string_alias: Optional[pulumi.Input[str]] = None) -> 'NamespaceDisasterRecoveryConfig':
        """
        Get an existing NamespaceDisasterRecoveryConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_primary_key: The primary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] default_secondary_key: The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        :param pulumi.Input[str] name: Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] partner_namespace_id: The ID of the Service Bus Namespace to replicate to.
        :param pulumi.Input[str] primary_connection_string_alias: The alias Primary Connection String for the ServiceBus Namespace.
        :param pulumi.Input[str] primary_namespace_id: The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_connection_string_alias: The alias Secondary Connection String for the ServiceBus Namespace
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceDisasterRecoveryConfigState.__new__(_NamespaceDisasterRecoveryConfigState)

        __props__.__dict__["default_primary_key"] = default_primary_key
        __props__.__dict__["default_secondary_key"] = default_secondary_key
        __props__.__dict__["name"] = name
        __props__.__dict__["partner_namespace_id"] = partner_namespace_id
        __props__.__dict__["primary_connection_string_alias"] = primary_connection_string_alias
        __props__.__dict__["primary_namespace_id"] = primary_namespace_id
        __props__.__dict__["secondary_connection_string_alias"] = secondary_connection_string_alias
        return NamespaceDisasterRecoveryConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultPrimaryKey")
    def default_primary_key(self) -> pulumi.Output[str]:
        """
        The primary access key for the authorization rule `RootManageSharedAccessKey`.
        """
        return pulumi.get(self, "default_primary_key")

    @property
    @pulumi.getter(name="defaultSecondaryKey")
    def default_secondary_key(self) -> pulumi.Output[str]:
        """
        The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        """
        return pulumi.get(self, "default_secondary_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Disaster Recovery Config. This is the alias DNS name that will be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerNamespaceId")
    def partner_namespace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Service Bus Namespace to replicate to.
        """
        return pulumi.get(self, "partner_namespace_id")

    @property
    @pulumi.getter(name="primaryConnectionStringAlias")
    def primary_connection_string_alias(self) -> pulumi.Output[str]:
        """
        The alias Primary Connection String for the ServiceBus Namespace.
        """
        return pulumi.get(self, "primary_connection_string_alias")

    @property
    @pulumi.getter(name="primaryNamespaceId")
    def primary_namespace_id(self) -> pulumi.Output[str]:
        """
        The ID of the primary Service Bus Namespace to replicate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "primary_namespace_id")

    @property
    @pulumi.getter(name="secondaryConnectionStringAlias")
    def secondary_connection_string_alias(self) -> pulumi.Output[str]:
        """
        The alias Secondary Connection String for the ServiceBus Namespace
        """
        return pulumi.get(self, "secondary_connection_string_alias")

