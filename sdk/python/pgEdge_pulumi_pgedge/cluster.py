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
from ._inputs import *

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cloud_account_id: pulumi.Input[str],
                 firewalls: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_groups: Optional[pulumi.Input['ClusterNodeGroupsArgs']] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cloud_account_id: Cloud account ID of the cluster
        :param pulumi.Input[str] name: Name of the cluster
        """
        pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if firewalls is not None:
            pulumi.set(__self__, "firewalls", firewalls)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_groups is not None:
            pulumi.set(__self__, "node_groups", node_groups)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Input[str]:
        """
        Cloud account ID of the cluster
        """
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter
    def firewalls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]]:
        return pulumi.get(self, "firewalls")

    @firewalls.setter
    def firewalls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]]):
        pulumi.set(self, "firewalls", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the cluster
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeGroups")
    def node_groups(self) -> Optional[pulumi.Input['ClusterNodeGroupsArgs']]:
        return pulumi.get(self, "node_groups")

    @node_groups.setter
    def node_groups(self, value: Optional[pulumi.Input['ClusterNodeGroupsArgs']]):
        pulumi.set(self, "node_groups", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 firewalls: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_groups: Optional[pulumi.Input['ClusterNodeGroupsArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cloud_account_id: Cloud account ID of the cluster
        :param pulumi.Input[str] created_at: Created at of the cluster
        :param pulumi.Input[str] name: Name of the cluster
        :param pulumi.Input[str] status: Status of the cluster
        """
        if cloud_account_id is not None:
            pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if firewalls is not None:
            pulumi.set(__self__, "firewalls", firewalls)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_groups is not None:
            pulumi.set(__self__, "node_groups", node_groups)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud account ID of the cluster
        """
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Created at of the cluster
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def firewalls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]]:
        return pulumi.get(self, "firewalls")

    @firewalls.setter
    def firewalls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallArgs']]]]):
        pulumi.set(self, "firewalls", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the cluster
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeGroups")
    def node_groups(self) -> Optional[pulumi.Input['ClusterNodeGroupsArgs']]:
        return pulumi.get(self, "node_groups")

    @node_groups.setter
    def node_groups(self, value: Optional[pulumi.Input['ClusterNodeGroupsArgs']]):
        pulumi.set(self, "node_groups", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the cluster
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 firewalls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_groups: Optional[pulumi.Input[pulumi.InputType['ClusterNodeGroupsArgs']]] = None,
                 __props__=None):
        """
        Interface with the pgEdge service API for clusters.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_account_id: Cloud account ID of the cluster
        :param pulumi.Input[str] name: Name of the cluster
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Interface with the pgEdge service API for clusters.

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 firewalls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_groups: Optional[pulumi.Input[pulumi.InputType['ClusterNodeGroupsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            if cloud_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_account_id'")
            __props__.__dict__["cloud_account_id"] = cloud_account_id
            __props__.__dict__["firewalls"] = firewalls
            __props__.__dict__["name"] = name
            __props__.__dict__["node_groups"] = node_groups
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
        super(Cluster, __self__).__init__(
            'pgedge:index/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_account_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            firewalls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_groups: Optional[pulumi.Input[pulumi.InputType['ClusterNodeGroupsArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_account_id: Cloud account ID of the cluster
        :param pulumi.Input[str] created_at: Created at of the cluster
        :param pulumi.Input[str] name: Name of the cluster
        :param pulumi.Input[str] status: Status of the cluster
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cloud_account_id"] = cloud_account_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["firewalls"] = firewalls
        __props__.__dict__["name"] = name
        __props__.__dict__["node_groups"] = node_groups
        __props__.__dict__["status"] = status
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Output[str]:
        """
        Cloud account ID of the cluster
        """
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Created at of the cluster
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def firewalls(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterFirewall']]]:
        return pulumi.get(self, "firewalls")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the cluster
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeGroups")
    def node_groups(self) -> pulumi.Output['outputs.ClusterNodeGroups']:
        return pulumi.get(self, "node_groups")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the cluster
        """
        return pulumi.get(self, "status")

