# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApiPolicyArgs', 'ApiPolicy']

@pulumi.input_type
class ApiPolicyArgs:
    def __init__(__self__, *,
                 api_management_name: pulumi.Input[str],
                 api_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiPolicy resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        pulumi.set(__self__, "api_management_name", api_management_name)
        pulumi.set(__self__, "api_name", api_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if xml_content is not None:
            pulumi.set(__self__, "xml_content", xml_content)
        if xml_link is not None:
            pulumi.set(__self__, "xml_link", xml_link)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Input[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> pulumi.Input[str]:
        """
        The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_name")

    @api_name.setter
    def api_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> Optional[pulumi.Input[str]]:
        """
        The XML Content for this Policy as a string.
        """
        return pulumi.get(self, "xml_content")

    @xml_content.setter
    def xml_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_content", value)

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> Optional[pulumi.Input[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    @xml_link.setter
    def xml_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_link", value)


@pulumi.input_type
class _ApiPolicyState:
    def __init__(__self__, *,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiPolicy resources.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        if api_management_name is not None:
            pulumi.set(__self__, "api_management_name", api_management_name)
        if api_name is not None:
            pulumi.set(__self__, "api_name", api_name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if xml_content is not None:
            pulumi.set(__self__, "xml_content", xml_content)
        if xml_link is not None:
            pulumi.set(__self__, "xml_link", xml_link)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @api_management_name.setter
    def api_management_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_management_name", value)

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_name")

    @api_name.setter
    def api_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> Optional[pulumi.Input[str]]:
        """
        The XML Content for this Policy as a string.
        """
        return pulumi.get(self, "xml_content")

    @xml_content.setter
    def xml_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_content", value)

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> Optional[pulumi.Input[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

    @xml_link.setter
    def xml_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "xml_link", value)


class ApiPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an API Management API Policy

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_api = azure.apimanagement.get_api(api_name="my-api",
            api_management_name="example-apim",
            resource_group_name="search-service")
        example_api_policy = azure.apimanagement.ApiPolicy("exampleApiPolicy",
            api_name=example_api.name,
            api_management_name=example_api.api_management_name,
            resource_group_name=example_api.resource_group_name,
            xml_content=\"\"\"<policies>
          <inbound>
            <find-and-replace from="xyz" to="abc" />
          </inbound>
        </policies>
        \"\"\")
        ```

        ## Import

        API Management API Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/apiPolicy:ApiPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/exampleId/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an API Management API Policy

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_api = azure.apimanagement.get_api(api_name="my-api",
            api_management_name="example-apim",
            resource_group_name="search-service")
        example_api_policy = azure.apimanagement.ApiPolicy("exampleApiPolicy",
            api_name=example_api.name,
            api_management_name=example_api.api_management_name,
            resource_group_name=example_api.resource_group_name,
            xml_content=\"\"\"<policies>
          <inbound>
            <find-and-replace from="xyz" to="abc" />
          </inbound>
        </policies>
        \"\"\")
        ```

        ## Import

        API Management API Policy can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:apimanagement/apiPolicy:ApiPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/exampleId/policies/policy
        ```

        :param str resource_name: The name of the resource.
        :param ApiPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_name: Optional[pulumi.Input[str]] = None,
                 api_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 xml_content: Optional[pulumi.Input[str]] = None,
                 xml_link: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApiPolicyArgs.__new__(ApiPolicyArgs)

            if api_management_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_management_name'")
            __props__.__dict__["api_management_name"] = api_management_name
            if api_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_name'")
            __props__.__dict__["api_name"] = api_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["xml_content"] = xml_content
            __props__.__dict__["xml_link"] = xml_link
        super(ApiPolicy, __self__).__init__(
            'azure:apimanagement/apiPolicy:ApiPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_management_name: Optional[pulumi.Input[str]] = None,
            api_name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            xml_content: Optional[pulumi.Input[str]] = None,
            xml_link: Optional[pulumi.Input[str]] = None) -> 'ApiPolicy':
        """
        Get an existing ApiPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] api_name: The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] xml_content: The XML Content for this Policy as a string.
        :param pulumi.Input[str] xml_link: A link to a Policy XML Document, which must be publicly available.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiPolicyState.__new__(_ApiPolicyState)

        __props__.__dict__["api_management_name"] = api_management_name
        __props__.__dict__["api_name"] = api_name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["xml_content"] = xml_content
        __props__.__dict__["xml_link"] = xml_link
        return ApiPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiManagementName")
    def api_management_name(self) -> pulumi.Output[str]:
        """
        The name of the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_management_name")

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> pulumi.Output[str]:
        """
        The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="xmlContent")
    def xml_content(self) -> pulumi.Output[str]:
        """
        The XML Content for this Policy as a string.
        """
        return pulumi.get(self, "xml_content")

    @property
    @pulumi.getter(name="xmlLink")
    def xml_link(self) -> pulumi.Output[Optional[str]]:
        """
        A link to a Policy XML Document, which must be publicly available.
        """
        return pulumi.get(self, "xml_link")

