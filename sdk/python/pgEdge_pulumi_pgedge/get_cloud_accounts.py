# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCloudAccountsResult',
    'AwaitableGetCloudAccountsResult',
    'get_cloud_accounts',
    'get_cloud_accounts_output',
]

@pulumi.output_type
class GetCloudAccountsResult:
    """
    A collection of values returned by getCloudAccounts.
    """
    def __init__(__self__, cloud_accounts=None, id=None):
        if cloud_accounts and not isinstance(cloud_accounts, list):
            raise TypeError("Expected argument 'cloud_accounts' to be a list")
        pulumi.set(__self__, "cloud_accounts", cloud_accounts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="cloudAccounts")
    def cloud_accounts(self) -> Sequence['outputs.GetCloudAccountsCloudAccountResult']:
        return pulumi.get(self, "cloud_accounts")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetCloudAccountsResult(GetCloudAccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudAccountsResult(
            cloud_accounts=self.cloud_accounts,
            id=self.id)


def get_cloud_accounts(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudAccountsResult:
    """
    Data source for pgEdge cloud accounts.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pgedge:index/getCloudAccounts:getCloudAccounts', __args__, opts=opts, typ=GetCloudAccountsResult).value

    return AwaitableGetCloudAccountsResult(
        cloud_accounts=pulumi.get(__ret__, 'cloud_accounts'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_cloud_accounts)
def get_cloud_accounts_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudAccountsResult]:
    """
    Data source for pgEdge cloud accounts.
    """
    ...
