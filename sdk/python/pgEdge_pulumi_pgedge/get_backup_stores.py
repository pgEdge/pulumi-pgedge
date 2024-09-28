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
    'GetBackupStoresResult',
    'AwaitableGetBackupStoresResult',
    'get_backup_stores',
    'get_backup_stores_output',
]

@pulumi.output_type
class GetBackupStoresResult:
    """
    A collection of values returned by getBackupStores.
    """
    def __init__(__self__, backup_stores=None, id=None):
        if backup_stores and not isinstance(backup_stores, list):
            raise TypeError("Expected argument 'backup_stores' to be a list")
        pulumi.set(__self__, "backup_stores", backup_stores)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="backupStores")
    def backup_stores(self) -> Sequence['outputs.GetBackupStoresBackupStoreResult']:
        return pulumi.get(self, "backup_stores")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetBackupStoresResult(GetBackupStoresResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupStoresResult(
            backup_stores=self.backup_stores,
            id=self.id)


def get_backup_stores(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupStoresResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pgedge:index/getBackupStores:getBackupStores', __args__, opts=opts, typ=GetBackupStoresResult).value

    return AwaitableGetBackupStoresResult(
        backup_stores=pulumi.get(__ret__, 'backup_stores'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_backup_stores)
def get_backup_stores_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupStoresResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
