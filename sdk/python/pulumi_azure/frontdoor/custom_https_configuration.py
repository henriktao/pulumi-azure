# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CustomHttpsConfiguration']


class CustomHttpsConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_https_configuration: Optional[pulumi.Input[pulumi.InputType['CustomHttpsConfigurationCustomHttpsConfigurationArgs']]] = None,
                 custom_https_provisioning_enabled: Optional[pulumi.Input[bool]] = None,
                 frontend_endpoint_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages the Custom Https Configuration for an Azure Front Door Frontend Endpoint..

        > **NOTE:** Custom https configurations for a Front Door Frontend Endpoint can be defined both within the `frontdoor.Frontdoor` resource via the `custom_https_configuration` block and by using a separate resource, as described in the following sections.

        > **NOTE:** Defining custom https configurations using a separate `frontdoor.CustomHttpsConfiguration` resource allows for parallel creation/update.

        > **NOTE:** UPCOMING BREAKING CHANGE: In order to address the ordering issue we have changed the design on how to retrieve existing sub resources such as frontend endpoints. Existing design will be deprecated and will result in an incorrect configuration. Please refer to the updated documentation below for more information.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        vault = azure.keyvault.get_key_vault(name="example-vault",
            resource_group_name="example-vault-rg")
        example_frontdoor = azure.frontdoor.Frontdoor("exampleFrontdoor",
            resource_group_name=example_resource_group.name,
            enforce_backend_pools_certificate_name_check=False,
            routing_rules=[azure.frontdoor.FrontdoorRoutingRuleArgs(
                name="exampleRoutingRule1",
                accepted_protocols=[
                    "Http",
                    "Https",
                ],
                patterns_to_matches=["/*"],
                frontend_endpoints=["exampleFrontendEndpoint1"],
                forwarding_configuration=azure.frontdoor.FrontdoorRoutingRuleForwardingConfigurationArgs(
                    forwarding_protocol="MatchRequest",
                    backend_pool_name="exampleBackendBing",
                ),
            )],
            backend_pool_load_balancings=[azure.frontdoor.FrontdoorBackendPoolLoadBalancingArgs(
                name="exampleLoadBalancingSettings1",
            )],
            backend_pool_health_probes=[azure.frontdoor.FrontdoorBackendPoolHealthProbeArgs(
                name="exampleHealthProbeSetting1",
            )],
            backend_pools=[azure.frontdoor.FrontdoorBackendPoolArgs(
                name="exampleBackendBing",
                backends=[azure.frontdoor.FrontdoorBackendPoolBackendArgs(
                    host_header="www.bing.com",
                    address="www.bing.com",
                    http_port=80,
                    https_port=443,
                )],
                load_balancing_name="exampleLoadBalancingSettings1",
                health_probe_name="exampleHealthProbeSetting1",
            )],
            frontend_endpoints=[
                azure.frontdoor.FrontdoorFrontendEndpointArgs(
                    name="exampleFrontendEndpoint1",
                    host_name="example-FrontDoor.azurefd.net",
                ),
                azure.frontdoor.FrontdoorFrontendEndpointArgs(
                    name="exampleFrontendEndpoint2",
                    host_name="examplefd1.examplefd.net",
                ),
            ])
        example_custom_https0 = azure.frontdoor.CustomHttpsConfiguration("exampleCustomHttps0",
            frontend_endpoint_id=example_frontdoor.frontend_endpoints_map["exampleFrontendEndpoint1"],
            custom_https_provisioning_enabled=False)
        example_custom_https1 = azure.frontdoor.CustomHttpsConfiguration("exampleCustomHttps1",
            frontend_endpoint_id=example_frontdoor.frontend_endpoints_map["exampleFrontendEndpoint2"],
            custom_https_provisioning_enabled=True,
            custom_https_configuration=azure.frontdoor.CustomHttpsConfigurationCustomHttpsConfigurationArgs(
                certificate_source="AzureKeyVault",
                azure_key_vault_certificate_secret_name="examplefd1",
                azure_key_vault_certificate_secret_version="ec8d0737e0df4f4gb52ecea858e97a73",
                azure_key_vault_certificate_vault_id=vault.id,
            ))
        ```

        ## Import

        Front Door Custom Https Configurations can be imported using the `resource id` of the Frontend Endpoint, e.g.

        ```sh
         $ pulumi import azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration example_custom_https_1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/frontDoors/frontdoor1/frontendEndpoints/endpoint1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CustomHttpsConfigurationCustomHttpsConfigurationArgs']] custom_https_configuration: A `custom_https_configuration` block as defined below.
        :param pulumi.Input[bool] custom_https_provisioning_enabled: Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        :param pulumi.Input[str] frontend_endpoint_id: The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['custom_https_configuration'] = custom_https_configuration
            if custom_https_provisioning_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'custom_https_provisioning_enabled'")
            __props__['custom_https_provisioning_enabled'] = custom_https_provisioning_enabled
            if frontend_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_endpoint_id'")
            __props__['frontend_endpoint_id'] = frontend_endpoint_id
            if resource_group_name is not None and not opts.urn:
                warnings.warn("""This field is no longer used and will be removed in the next major version of the Azure Provider""", DeprecationWarning)
                pulumi.log.warn("resource_group_name is deprecated: This field is no longer used and will be removed in the next major version of the Azure Provider")
            __props__['resource_group_name'] = resource_group_name
        super(CustomHttpsConfiguration, __self__).__init__(
            'azure:frontdoor/customHttpsConfiguration:CustomHttpsConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_https_configuration: Optional[pulumi.Input[pulumi.InputType['CustomHttpsConfigurationCustomHttpsConfigurationArgs']]] = None,
            custom_https_provisioning_enabled: Optional[pulumi.Input[bool]] = None,
            frontend_endpoint_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'CustomHttpsConfiguration':
        """
        Get an existing CustomHttpsConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CustomHttpsConfigurationCustomHttpsConfigurationArgs']] custom_https_configuration: A `custom_https_configuration` block as defined below.
        :param pulumi.Input[bool] custom_https_provisioning_enabled: Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        :param pulumi.Input[str] frontend_endpoint_id: The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_https_configuration"] = custom_https_configuration
        __props__["custom_https_provisioning_enabled"] = custom_https_provisioning_enabled
        __props__["frontend_endpoint_id"] = frontend_endpoint_id
        __props__["resource_group_name"] = resource_group_name
        return CustomHttpsConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customHttpsConfiguration")
    def custom_https_configuration(self) -> pulumi.Output[Optional['outputs.CustomHttpsConfigurationCustomHttpsConfiguration']]:
        """
        A `custom_https_configuration` block as defined below.
        """
        return pulumi.get(self, "custom_https_configuration")

    @property
    @pulumi.getter(name="customHttpsProvisioningEnabled")
    def custom_https_provisioning_enabled(self) -> pulumi.Output[bool]:
        """
        Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
        """
        return pulumi.get(self, "custom_https_provisioning_enabled")

    @property
    @pulumi.getter(name="frontendEndpointId")
    def frontend_endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the FrontDoor Frontend Endpoint which this configuration refers to.
        """
        return pulumi.get(self, "frontend_endpoint_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

