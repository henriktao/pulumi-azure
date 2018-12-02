# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class ScaleSet(pulumi.CustomResource):
    """
    Manage a virtual machine scale set.
    
    ~> **Note:** All arguments including the administrator login and password will be stored in the raw state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, automatic_os_upgrade=None, boot_diagnostics=None, eviction_policy=None, extensions=None, health_probe_id=None, identity=None, license_type=None, location=None, name=None, network_profiles=None, os_profile=None, os_profile_linux_config=None, os_profile_secrets=None, os_profile_windows_config=None, overprovision=None, plan=None, priority=None, resource_group_name=None, rolling_upgrade_policy=None, single_placement_group=None, sku=None, storage_profile_data_disks=None, storage_profile_image_reference=None, storage_profile_os_disk=None, tags=None, upgrade_policy_mode=None, zones=None):
        """Create a ScaleSet resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['automatic_os_upgrade'] = automatic_os_upgrade

        __props__['boot_diagnostics'] = boot_diagnostics

        __props__['eviction_policy'] = eviction_policy

        __props__['extensions'] = extensions

        __props__['health_probe_id'] = health_probe_id

        __props__['identity'] = identity

        __props__['license_type'] = license_type

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        if not network_profiles:
            raise TypeError('Missing required property network_profiles')
        __props__['network_profiles'] = network_profiles

        if not os_profile:
            raise TypeError('Missing required property os_profile')
        __props__['os_profile'] = os_profile

        __props__['os_profile_linux_config'] = os_profile_linux_config

        __props__['os_profile_secrets'] = os_profile_secrets

        __props__['os_profile_windows_config'] = os_profile_windows_config

        __props__['overprovision'] = overprovision

        __props__['plan'] = plan

        __props__['priority'] = priority

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['rolling_upgrade_policy'] = rolling_upgrade_policy

        __props__['single_placement_group'] = single_placement_group

        if not sku:
            raise TypeError('Missing required property sku')
        __props__['sku'] = sku

        __props__['storage_profile_data_disks'] = storage_profile_data_disks

        __props__['storage_profile_image_reference'] = storage_profile_image_reference

        if not storage_profile_os_disk:
            raise TypeError('Missing required property storage_profile_os_disk')
        __props__['storage_profile_os_disk'] = storage_profile_os_disk

        __props__['tags'] = tags

        if not upgrade_policy_mode:
            raise TypeError('Missing required property upgrade_policy_mode')
        __props__['upgrade_policy_mode'] = upgrade_policy_mode

        __props__['zones'] = zones

        super(ScaleSet, __self__).__init__(
            'azure:compute/scaleSet:ScaleSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

