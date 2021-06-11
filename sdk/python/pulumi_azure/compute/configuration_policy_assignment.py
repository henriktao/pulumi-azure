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

__all__ = ['ConfigurationPolicyAssignmentArgs', 'ConfigurationPolicyAssignment']

@pulumi.input_type
class ConfigurationPolicyAssignmentArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs'],
                 virtual_machine_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConfigurationPolicyAssignment resource.
        :param pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs'] configuration: A `configuration` block as defined below.
        :param pulumi.Input[str] virtual_machine_id: The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs']:
        """
        A `configuration` block as defined below.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs']):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_machine_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ConfigurationPolicyAssignmentState:
    def __init__(__self__, *,
                 configuration: Optional[pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigurationPolicyAssignment resources.
        :param pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs'] configuration: A `configuration` block as defined below.
        :param pulumi.Input[str] location: The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_machine_id: The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs']]:
        """
        A `configuration` block as defined below.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['ConfigurationPolicyAssignmentConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_machine_id", value)


class ConfigurationPolicyAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['ConfigurationPolicyAssignmentConfigurationArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Applies a Configuration Policy to a Virtual Machine.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"])
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            ip_configurations=[azure.network.NetworkInterfaceIpConfigurationArgs(
                name="internal",
                subnet_id=example_subnet.id,
                private_ip_address_allocation="Dynamic",
            )])
        example_windows_virtual_machine = azure.compute.WindowsVirtualMachine("exampleWindowsVirtualMachine",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            size="Standard_F2",
            admin_username="adminuser",
            admin_password="P@$$w0rd1234!",
            network_interface_ids=[example_network_interface.id],
            identity=azure.compute.WindowsVirtualMachineIdentityArgs(
                type="SystemAssigned",
            ),
            os_disk=azure.compute.WindowsVirtualMachineOsDiskArgs(
                caching="ReadWrite",
                storage_account_type="Standard_LRS",
            ),
            source_image_reference=azure.compute.WindowsVirtualMachineSourceImageReferenceArgs(
                publisher="MicrosoftWindowsServer",
                offer="WindowsServer",
                sku="2019-Datacenter",
                version="latest",
            ))
        example_extension = azure.compute.Extension("exampleExtension",
            virtual_machine_id=example_windows_virtual_machine.id,
            publisher="Microsoft.GuestConfiguration",
            type="ConfigurationforWindows",
            type_handler_version="1.0",
            auto_upgrade_minor_version=True)
        example_configuration_policy_assignment = azure.compute.ConfigurationPolicyAssignment("exampleConfigurationPolicyAssignment",
            location=example_windows_virtual_machine.location,
            virtual_machine_id=example_windows_virtual_machine.id,
            configuration=azure.compute.ConfigurationPolicyAssignmentConfigurationArgs(
                name="AzureWindowsBaseline",
                version="1.*",
                parameters=[
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Minimum Password Length;ExpectedValue",
                        value="16",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Minimum Password Age;ExpectedValue",
                        value="0",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Maximum Password Age;ExpectedValue",
                        value="30,45",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Enforce Password History;ExpectedValue",
                        value="10",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Password Must Meet Complexity Requirements;ExpectedValue",
                        value="1",
                    ),
                ],
            ))
        ```

        ## Import

        Virtual Machine Configuration Policy Assignments can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConfigurationPolicyAssignmentConfigurationArgs']] configuration: A `configuration` block as defined below.
        :param pulumi.Input[str] location: The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_machine_id: The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationPolicyAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Applies a Configuration Policy to a Virtual Machine.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"])
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"])
        example_network_interface = azure.network.NetworkInterface("exampleNetworkInterface",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            ip_configurations=[azure.network.NetworkInterfaceIpConfigurationArgs(
                name="internal",
                subnet_id=example_subnet.id,
                private_ip_address_allocation="Dynamic",
            )])
        example_windows_virtual_machine = azure.compute.WindowsVirtualMachine("exampleWindowsVirtualMachine",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            size="Standard_F2",
            admin_username="adminuser",
            admin_password="P@$$w0rd1234!",
            network_interface_ids=[example_network_interface.id],
            identity=azure.compute.WindowsVirtualMachineIdentityArgs(
                type="SystemAssigned",
            ),
            os_disk=azure.compute.WindowsVirtualMachineOsDiskArgs(
                caching="ReadWrite",
                storage_account_type="Standard_LRS",
            ),
            source_image_reference=azure.compute.WindowsVirtualMachineSourceImageReferenceArgs(
                publisher="MicrosoftWindowsServer",
                offer="WindowsServer",
                sku="2019-Datacenter",
                version="latest",
            ))
        example_extension = azure.compute.Extension("exampleExtension",
            virtual_machine_id=example_windows_virtual_machine.id,
            publisher="Microsoft.GuestConfiguration",
            type="ConfigurationforWindows",
            type_handler_version="1.0",
            auto_upgrade_minor_version=True)
        example_configuration_policy_assignment = azure.compute.ConfigurationPolicyAssignment("exampleConfigurationPolicyAssignment",
            location=example_windows_virtual_machine.location,
            virtual_machine_id=example_windows_virtual_machine.id,
            configuration=azure.compute.ConfigurationPolicyAssignmentConfigurationArgs(
                name="AzureWindowsBaseline",
                version="1.*",
                parameters=[
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Minimum Password Length;ExpectedValue",
                        value="16",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Minimum Password Age;ExpectedValue",
                        value="0",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Maximum Password Age;ExpectedValue",
                        value="30,45",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Enforce Password History;ExpectedValue",
                        value="10",
                    ),
                    azure.compute.ConfigurationPolicyAssignmentConfigurationParameterArgs(
                        name="Password Must Meet Complexity Requirements;ExpectedValue",
                        value="1",
                    ),
                ],
            ))
        ```

        ## Import

        Virtual Machine Configuration Policy Assignments can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
        ```

        :param str resource_name: The name of the resource.
        :param ConfigurationPolicyAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationPolicyAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['ConfigurationPolicyAssignmentConfigurationArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConfigurationPolicyAssignmentArgs.__new__(ConfigurationPolicyAssignmentArgs)

            if configuration is None and not opts.urn:
                raise TypeError("Missing required property 'configuration'")
            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if virtual_machine_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        super(ConfigurationPolicyAssignment, __self__).__init__(
            'azure:compute/configurationPolicyAssignment:ConfigurationPolicyAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['ConfigurationPolicyAssignmentConfigurationArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'ConfigurationPolicyAssignment':
        """
        Get an existing ConfigurationPolicyAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConfigurationPolicyAssignmentConfigurationArgs']] configuration: A `configuration` block as defined below.
        :param pulumi.Input[str] location: The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_machine_id: The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationPolicyAssignmentState.__new__(_ConfigurationPolicyAssignmentState)

        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        return ConfigurationPolicyAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output['outputs.ConfigurationPolicyAssignmentConfiguration']:
        """
        A `configuration` block as defined below.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where the Virtual Machine Configuration Policy Assignment should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Virtual Machine Configuration Policy Assignment. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the Virtual Machine which this Guest Configuration Assignment should apply to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

