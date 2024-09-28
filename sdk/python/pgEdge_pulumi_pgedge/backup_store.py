# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BackupStoreArgs', 'BackupStore']

@pulumi.input_type
class BackupStoreArgs:
    def __init__(__self__, *,
                 cloud_account_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BackupStore resource.
        """
        pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BackupStoreState:
    def __init__(__self__, *,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 cloud_account_type: Optional[pulumi.Input[str]] = None,
                 cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BackupStore resources.
        """
        if cloud_account_id is not None:
            pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if cloud_account_type is not None:
            pulumi.set(__self__, "cloud_account_type", cloud_account_type)
        if cluster_ids is not None:
            pulumi.set(__self__, "cluster_ids", cluster_ids)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter(name="cloudAccountType")
    def cloud_account_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cloud_account_type")

    @cloud_account_type.setter
    def cloud_account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_account_type", value)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "cluster_ids")

    @cluster_ids.setter
    def cluster_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cluster_ids", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class BackupStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BackupStore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BackupStore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BackupStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupStoreArgs.__new__(BackupStoreArgs)

            if cloud_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_account_id'")
            __props__.__dict__["cloud_account_id"] = cloud_account_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["cloud_account_type"] = None
            __props__.__dict__["cluster_ids"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["properties"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(BackupStore, __self__).__init__(
            'pgedge:index/backupStore:BackupStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_account_id: Optional[pulumi.Input[str]] = None,
            cloud_account_type: Optional[pulumi.Input[str]] = None,
            cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'BackupStore':
        """
        Get an existing BackupStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupStoreState.__new__(_BackupStoreState)

        __props__.__dict__["cloud_account_id"] = cloud_account_id
        __props__.__dict__["cloud_account_type"] = cloud_account_type
        __props__.__dict__["cluster_ids"] = cluster_ids
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["name"] = name
        __props__.__dict__["properties"] = properties
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["updated_at"] = updated_at
        return BackupStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="cloudAccountType")
    def cloud_account_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloud_account_type")

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "cluster_ids")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

