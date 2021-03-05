# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetMcaAccountScopeResult',
    'AwaitableGetMcaAccountScopeResult',
    'get_mca_account_scope',
]

@pulumi.output_type
class GetMcaAccountScopeResult:
    """
    A collection of values returned by getMcaAccountScope.
    """
    def __init__(__self__, billing_account_name=None, billing_profile_name=None, id=None, invoice_section_name=None):
        if billing_account_name and not isinstance(billing_account_name, str):
            raise TypeError("Expected argument 'billing_account_name' to be a str")
        pulumi.set(__self__, "billing_account_name", billing_account_name)
        if billing_profile_name and not isinstance(billing_profile_name, str):
            raise TypeError("Expected argument 'billing_profile_name' to be a str")
        pulumi.set(__self__, "billing_profile_name", billing_profile_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invoice_section_name and not isinstance(invoice_section_name, str):
            raise TypeError("Expected argument 'invoice_section_name' to be a str")
        pulumi.set(__self__, "invoice_section_name", invoice_section_name)

    @property
    @pulumi.getter(name="billingAccountName")
    def billing_account_name(self) -> str:
        return pulumi.get(self, "billing_account_name")

    @property
    @pulumi.getter(name="billingProfileName")
    def billing_profile_name(self) -> str:
        return pulumi.get(self, "billing_profile_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invoiceSectionName")
    def invoice_section_name(self) -> str:
        return pulumi.get(self, "invoice_section_name")


class AwaitableGetMcaAccountScopeResult(GetMcaAccountScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMcaAccountScopeResult(
            billing_account_name=self.billing_account_name,
            billing_profile_name=self.billing_profile_name,
            id=self.id,
            invoice_section_name=self.invoice_section_name)


def get_mca_account_scope(billing_account_name: Optional[str] = None,
                          billing_profile_name: Optional[str] = None,
                          invoice_section_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMcaAccountScopeResult:
    """
    Use this data source to access an ID for your MCA Account billing scope.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.billing.get_mca_account_scope(billing_account_name="e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
        billing_profile_name="PE2Q-NOIT-BG7-TGB",
        invoice_section_name="MTT4-OBS7-PJA-TGB")
    pulumi.export("id", example.id)
    ```


    :param str billing_account_name: The Billing Account Name of the MCA account.
    :param str billing_profile_name: The Billing Profile Name in the above Billing Account.
    :param str invoice_section_name: The Invoice Section Name in the above Billing Profile.
    """
    __args__ = dict()
    __args__['billingAccountName'] = billing_account_name
    __args__['billingProfileName'] = billing_profile_name
    __args__['invoiceSectionName'] = invoice_section_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:billing/getMcaAccountScope:getMcaAccountScope', __args__, opts=opts, typ=GetMcaAccountScopeResult).value

    return AwaitableGetMcaAccountScopeResult(
        billing_account_name=__ret__.billing_account_name,
        billing_profile_name=__ret__.billing_profile_name,
        id=__ret__.id,
        invoice_section_name=__ret__.invoice_section_name)
