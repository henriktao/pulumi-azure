# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetTopicAuthorizationRuleResult',
    'AwaitableGetTopicAuthorizationRuleResult',
    'get_topic_authorization_rule',
    'get_topic_authorization_rule_output',
]

@pulumi.output_type
class GetTopicAuthorizationRuleResult:
    """
    A collection of values returned by getTopicAuthorizationRule.
    """
    def __init__(__self__, id=None, listen=None, manage=None, name=None, namespace_name=None, primary_connection_string=None, primary_connection_string_alias=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_connection_string_alias=None, secondary_key=None, send=None, topic_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listen and not isinstance(listen, bool):
            raise TypeError("Expected argument 'listen' to be a bool")
        pulumi.set(__self__, "listen", listen)
        if manage and not isinstance(manage, bool):
            raise TypeError("Expected argument 'manage' to be a bool")
        pulumi.set(__self__, "manage", manage)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        pulumi.set(__self__, "namespace_name", namespace_name)
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        pulumi.set(__self__, "primary_connection_string", primary_connection_string)
        if primary_connection_string_alias and not isinstance(primary_connection_string_alias, str):
            raise TypeError("Expected argument 'primary_connection_string_alias' to be a str")
        pulumi.set(__self__, "primary_connection_string_alias", primary_connection_string_alias)
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        pulumi.set(__self__, "primary_key", primary_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        pulumi.set(__self__, "secondary_connection_string", secondary_connection_string)
        if secondary_connection_string_alias and not isinstance(secondary_connection_string_alias, str):
            raise TypeError("Expected argument 'secondary_connection_string_alias' to be a str")
        pulumi.set(__self__, "secondary_connection_string_alias", secondary_connection_string_alias)
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        pulumi.set(__self__, "secondary_key", secondary_key)
        if send and not isinstance(send, bool):
            raise TypeError("Expected argument 'send' to be a bool")
        pulumi.set(__self__, "send", send)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def listen(self) -> bool:
        return pulumi.get(self, "listen")

    @property
    @pulumi.getter
    def manage(self) -> bool:
        return pulumi.get(self, "manage")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> str:
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="primaryConnectionString")
    def primary_connection_string(self) -> str:
        """
        The Primary Connection String for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "primary_connection_string")

    @property
    @pulumi.getter(name="primaryConnectionStringAlias")
    def primary_connection_string_alias(self) -> str:
        """
        The alias Primary Connection String for the ServiceBus Namespace, if the namespace is Geo DR paired.
        """
        return pulumi.get(self, "primary_connection_string_alias")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The Primary Key for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryConnectionString")
    def secondary_connection_string(self) -> str:
        """
        The Secondary Connection String for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "secondary_connection_string")

    @property
    @pulumi.getter(name="secondaryConnectionStringAlias")
    def secondary_connection_string_alias(self) -> str:
        """
        The alias Secondary Connection String for the ServiceBus Namespace
        """
        return pulumi.get(self, "secondary_connection_string_alias")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The Secondary Key for the ServiceBus Topic authorization Rule.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter
    def send(self) -> bool:
        return pulumi.get(self, "send")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> str:
        return pulumi.get(self, "topic_name")


class AwaitableGetTopicAuthorizationRuleResult(GetTopicAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicAuthorizationRuleResult(
            id=self.id,
            listen=self.listen,
            manage=self.manage,
            name=self.name,
            namespace_name=self.namespace_name,
            primary_connection_string=self.primary_connection_string,
            primary_connection_string_alias=self.primary_connection_string_alias,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_connection_string=self.secondary_connection_string,
            secondary_connection_string_alias=self.secondary_connection_string_alias,
            secondary_key=self.secondary_key,
            send=self.send,
            topic_name=self.topic_name)


def get_topic_authorization_rule(name: Optional[str] = None,
                                 namespace_name: Optional[str] = None,
                                 resource_group_name: Optional[str] = None,
                                 topic_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicAuthorizationRuleResult:
    """
    Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.servicebus.get_topic_authorization_rule(name="example-tfex_name",
        namespace_name="example-namespace",
        resource_group_name="example-resources",
        topic_name="example-servicebus_topic")
    pulumi.export("servicebusAuthorizationRuleId", data["azurem_servicebus_topic_authorization_rule"]["example"]["id"])
    ```


    :param str name: The name of the ServiceBus Topic Authorization Rule resource.
    :param str namespace_name: The name of the ServiceBus Namespace.
    :param str resource_group_name: The name of the resource group in which the ServiceBus Namespace exists.
    :param str topic_name: The name of the ServiceBus Topic.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:servicebus/getTopicAuthorizationRule:getTopicAuthorizationRule', __args__, opts=opts, typ=GetTopicAuthorizationRuleResult).value

    return AwaitableGetTopicAuthorizationRuleResult(
        id=__ret__.id,
        listen=__ret__.listen,
        manage=__ret__.manage,
        name=__ret__.name,
        namespace_name=__ret__.namespace_name,
        primary_connection_string=__ret__.primary_connection_string,
        primary_connection_string_alias=__ret__.primary_connection_string_alias,
        primary_key=__ret__.primary_key,
        resource_group_name=__ret__.resource_group_name,
        secondary_connection_string=__ret__.secondary_connection_string,
        secondary_connection_string_alias=__ret__.secondary_connection_string_alias,
        secondary_key=__ret__.secondary_key,
        send=__ret__.send,
        topic_name=__ret__.topic_name)


@_utilities.lift_output_func(get_topic_authorization_rule)
def get_topic_authorization_rule_output(name: Optional[pulumi.Input[str]] = None,
                                        namespace_name: Optional[pulumi.Input[str]] = None,
                                        resource_group_name: Optional[pulumi.Input[str]] = None,
                                        topic_name: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicAuthorizationRuleResult]:
    """
    Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.servicebus.get_topic_authorization_rule(name="example-tfex_name",
        namespace_name="example-namespace",
        resource_group_name="example-resources",
        topic_name="example-servicebus_topic")
    pulumi.export("servicebusAuthorizationRuleId", data["azurem_servicebus_topic_authorization_rule"]["example"]["id"])
    ```


    :param str name: The name of the ServiceBus Topic Authorization Rule resource.
    :param str namespace_name: The name of the ServiceBus Namespace.
    :param str resource_group_name: The name of the resource group in which the ServiceBus Namespace exists.
    :param str topic_name: The name of the ServiceBus Topic.
    """
    ...
