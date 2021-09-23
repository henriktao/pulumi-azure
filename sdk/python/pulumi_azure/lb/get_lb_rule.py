# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetLBRuleResult',
    'AwaitableGetLBRuleResult',
    'get_lb_rule',
    'get_lb_rule_output',
]

@pulumi.output_type
class GetLBRuleResult:
    """
    A collection of values returned by getLBRule.
    """
    def __init__(__self__, backend_address_pool_id=None, backend_port=None, disable_outbound_snat=None, enable_floating_ip=None, enable_tcp_reset=None, frontend_ip_configuration_name=None, frontend_port=None, id=None, idle_timeout_in_minutes=None, load_distribution=None, loadbalancer_id=None, name=None, probe_id=None, protocol=None, resource_group_name=None):
        if backend_address_pool_id and not isinstance(backend_address_pool_id, str):
            raise TypeError("Expected argument 'backend_address_pool_id' to be a str")
        pulumi.set(__self__, "backend_address_pool_id", backend_address_pool_id)
        if backend_port and not isinstance(backend_port, int):
            raise TypeError("Expected argument 'backend_port' to be a int")
        pulumi.set(__self__, "backend_port", backend_port)
        if disable_outbound_snat and not isinstance(disable_outbound_snat, bool):
            raise TypeError("Expected argument 'disable_outbound_snat' to be a bool")
        pulumi.set(__self__, "disable_outbound_snat", disable_outbound_snat)
        if enable_floating_ip and not isinstance(enable_floating_ip, bool):
            raise TypeError("Expected argument 'enable_floating_ip' to be a bool")
        pulumi.set(__self__, "enable_floating_ip", enable_floating_ip)
        if enable_tcp_reset and not isinstance(enable_tcp_reset, bool):
            raise TypeError("Expected argument 'enable_tcp_reset' to be a bool")
        pulumi.set(__self__, "enable_tcp_reset", enable_tcp_reset)
        if frontend_ip_configuration_name and not isinstance(frontend_ip_configuration_name, str):
            raise TypeError("Expected argument 'frontend_ip_configuration_name' to be a str")
        pulumi.set(__self__, "frontend_ip_configuration_name", frontend_ip_configuration_name)
        if frontend_port and not isinstance(frontend_port, int):
            raise TypeError("Expected argument 'frontend_port' to be a int")
        pulumi.set(__self__, "frontend_port", frontend_port)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_timeout_in_minutes and not isinstance(idle_timeout_in_minutes, int):
            raise TypeError("Expected argument 'idle_timeout_in_minutes' to be a int")
        pulumi.set(__self__, "idle_timeout_in_minutes", idle_timeout_in_minutes)
        if load_distribution and not isinstance(load_distribution, str):
            raise TypeError("Expected argument 'load_distribution' to be a str")
        pulumi.set(__self__, "load_distribution", load_distribution)
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if probe_id and not isinstance(probe_id, str):
            raise TypeError("Expected argument 'probe_id' to be a str")
        pulumi.set(__self__, "probe_id", probe_id)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)

    @property
    @pulumi.getter(name="backendAddressPoolId")
    def backend_address_pool_id(self) -> str:
        """
        A reference to a Backend Address Pool over which this Load Balancing Rule operates.
        """
        return pulumi.get(self, "backend_address_pool_id")

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> int:
        """
        The port used for internal connections on the endpoint.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter(name="disableOutboundSnat")
    def disable_outbound_snat(self) -> bool:
        """
        If outbound SNAT is enabled for this Load Balancer Rule.
        """
        return pulumi.get(self, "disable_outbound_snat")

    @property
    @pulumi.getter(name="enableFloatingIp")
    def enable_floating_ip(self) -> bool:
        """
        If Floating IPs are enabled for this Load Balancer Rule
        """
        return pulumi.get(self, "enable_floating_ip")

    @property
    @pulumi.getter(name="enableTcpReset")
    def enable_tcp_reset(self) -> bool:
        """
        If TCP Reset is enabled for this Load Balancer Rule.
        """
        return pulumi.get(self, "enable_tcp_reset")

    @property
    @pulumi.getter(name="frontendIpConfigurationName")
    def frontend_ip_configuration_name(self) -> str:
        """
        The name of the frontend IP configuration to which the rule is associated.
        """
        return pulumi.get(self, "frontend_ip_configuration_name")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> int:
        """
        The port for the external endpoint.
        """
        return pulumi.get(self, "frontend_port")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleTimeoutInMinutes")
    def idle_timeout_in_minutes(self) -> int:
        """
        Specifies the idle timeout in minutes for TCP connections.
        """
        return pulumi.get(self, "idle_timeout_in_minutes")

    @property
    @pulumi.getter(name="loadDistribution")
    def load_distribution(self) -> str:
        """
        Specifies the load balancing distribution type used by the Load Balancer.
        """
        return pulumi.get(self, "load_distribution")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> str:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="probeId")
    def probe_id(self) -> str:
        """
        A reference to a Probe used by this Load Balancing Rule.
        """
        return pulumi.get(self, "probe_id")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The transport protocol for the external endpoint.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")


class AwaitableGetLBRuleResult(GetLBRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLBRuleResult(
            backend_address_pool_id=self.backend_address_pool_id,
            backend_port=self.backend_port,
            disable_outbound_snat=self.disable_outbound_snat,
            enable_floating_ip=self.enable_floating_ip,
            enable_tcp_reset=self.enable_tcp_reset,
            frontend_ip_configuration_name=self.frontend_ip_configuration_name,
            frontend_port=self.frontend_port,
            id=self.id,
            idle_timeout_in_minutes=self.idle_timeout_in_minutes,
            load_distribution=self.load_distribution,
            loadbalancer_id=self.loadbalancer_id,
            name=self.name,
            probe_id=self.probe_id,
            protocol=self.protocol,
            resource_group_name=self.resource_group_name)


def get_lb_rule(loadbalancer_id: Optional[str] = None,
                name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLBRuleResult:
    """
    Use this data source to access information about an existing Load Balancer Rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_lb = azure.lb.get_lb(name="example-lb",
        resource_group_name="example-resources")
    example_lb_rule = azure.lb.get_lb_rule(name="first",
        resource_group_name="example-resources",
        loadbalancer_id=example_lb.id)
    pulumi.export("lbRuleId", example_lb_rule.id)
    ```


    :param str loadbalancer_id: The ID of the Load Balancer Rule.
    :param str name: The name of this Load Balancer Rule.
    :param str resource_group_name: The name of the Resource Group where the Load Balancer Rule exists.
    """
    __args__ = dict()
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:lb/getLBRule:getLBRule', __args__, opts=opts, typ=GetLBRuleResult).value

    return AwaitableGetLBRuleResult(
        backend_address_pool_id=__ret__.backend_address_pool_id,
        backend_port=__ret__.backend_port,
        disable_outbound_snat=__ret__.disable_outbound_snat,
        enable_floating_ip=__ret__.enable_floating_ip,
        enable_tcp_reset=__ret__.enable_tcp_reset,
        frontend_ip_configuration_name=__ret__.frontend_ip_configuration_name,
        frontend_port=__ret__.frontend_port,
        id=__ret__.id,
        idle_timeout_in_minutes=__ret__.idle_timeout_in_minutes,
        load_distribution=__ret__.load_distribution,
        loadbalancer_id=__ret__.loadbalancer_id,
        name=__ret__.name,
        probe_id=__ret__.probe_id,
        protocol=__ret__.protocol,
        resource_group_name=__ret__.resource_group_name)


@_utilities.lift_output_func(get_lb_rule)
def get_lb_rule_output(loadbalancer_id: Optional[pulumi.Input[str]] = None,
                       name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLBRuleResult]:
    """
    Use this data source to access information about an existing Load Balancer Rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example_lb = azure.lb.get_lb(name="example-lb",
        resource_group_name="example-resources")
    example_lb_rule = azure.lb.get_lb_rule(name="first",
        resource_group_name="example-resources",
        loadbalancer_id=example_lb.id)
    pulumi.export("lbRuleId", example_lb_rule.id)
    ```


    :param str loadbalancer_id: The ID of the Load Balancer Rule.
    :param str name: The name of this Load Balancer Rule.
    :param str resource_group_name: The name of the Resource Group where the Load Balancer Rule exists.
    """
    ...
