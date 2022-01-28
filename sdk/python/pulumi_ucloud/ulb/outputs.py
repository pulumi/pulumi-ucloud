# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'LBIpSet',
    'GetLBAttachmentLbAttachmentResult',
    'GetLBLbResult',
    'GetLBLbIpSetResult',
    'GetLBListenersLbListenerResult',
    'GetLBRulesLbRuleResult',
    'GetLBSslLbSslResult',
]

@pulumi.output_type
class LBIpSet(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "internetType":
            suggest = "internet_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LBIpSet. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LBIpSet.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LBIpSet.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 internet_type: Optional[str] = None,
                 ip: Optional[str] = None):
        """
        :param str internet_type: Type of Elastic IP routes.
        :param str ip: Elastic IP address.
        """
        if internet_type is not None:
            pulumi.set(__self__, "internet_type", internet_type)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter(name="internetType")
    def internet_type(self) -> Optional[str]:
        """
        Type of Elastic IP routes.
        """
        return pulumi.get(self, "internet_type")

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        Elastic IP address.
        """
        return pulumi.get(self, "ip")


@pulumi.output_type
class GetLBAttachmentLbAttachmentResult(dict):
    def __init__(__self__, *,
                 id: str,
                 port: int,
                 private_ip: str,
                 resource_id: str,
                 status: str):
        """
        :param str id: The ID of LB Attachment.
        :param int port: Port opened on the backend server to receive requests, range: 1-65535.
        :param str private_ip: The private ip address for backend servers.
        :param str resource_id: The ID of a backend server.
        :param str status: The status of backend servers. Possible values are: `normalRunning`, `exceptionRunning`.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of LB Attachment.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Port opened on the backend server to receive requests, range: 1-65535.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The private ip address for backend servers.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The ID of a backend server.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of backend servers. Possible values are: `normalRunning`, `exceptionRunning`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetLBLbResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 id: str,
                 internal: bool,
                 ip_sets: Sequence['outputs.GetLBLbIpSetResult'],
                 name: str,
                 private_ip: str,
                 remark: str,
                 subnet_id: str,
                 tag: str,
                 vpc_id: str):
        """
        :param str create_time: The creation time of Load Balancer, formatted in RFC3339 time string.
        :param str id: The ID of Load Balancer.
        :param bool internal: Indicate whether the load balancer is intranet.
        :param str name: The name of Load Balancer.
        :param str private_ip: The IP address of intranet IP.
        :param str remark: The remarks of Load Balancer.
        :param str subnet_id: The ID of subnet that intrant load balancer belongs to.
        :param str tag: A tag assigned to Load Balancer.
        :param str vpc_id: The ID of the VPC linked to the Load Balancers.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "internal", internal)
        pulumi.set(__self__, "ip_sets", ip_sets)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "tag", tag)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of Load Balancer, formatted in RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of Load Balancer.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def internal(self) -> bool:
        """
        Indicate whether the load balancer is intranet.
        """
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> Sequence['outputs.GetLBLbIpSetResult']:
        return pulumi.get(self, "ip_sets")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of Load Balancer.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The IP address of intranet IP.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        The remarks of Load Balancer.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of subnet that intrant load balancer belongs to.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tag(self) -> str:
        """
        A tag assigned to Load Balancer.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC linked to the Load Balancers.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetLBLbIpSetResult(dict):
    def __init__(__self__, *,
                 internet_type: str,
                 ip: str):
        """
        :param str internet_type: Type of Load Balancer routes.
        :param str ip: Load Balancer address.
        """
        pulumi.set(__self__, "internet_type", internet_type)
        pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter(name="internetType")
    def internet_type(self) -> str:
        """
        Type of Load Balancer routes.
        """
        return pulumi.get(self, "internet_type")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        Load Balancer address.
        """
        return pulumi.get(self, "ip")


@pulumi.output_type
class GetLBListenersLbListenerResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 health_check_type: str,
                 id: str,
                 idle_timeout: int,
                 listen_type: str,
                 method: str,
                 name: str,
                 path: str,
                 persistence: str,
                 persistence_type: str,
                 port: int,
                 protocol: str,
                 status: str):
        """
        :param str domain: Health check domain checking.
        :param str health_check_type: Health check method. Possible values are `port` as port checking and `path` as http checking.
        :param str id: The ID of LB Listener.
        :param int idle_timeout: Amount of time in seconds to wait for the response for in between two sessions if `listen_type` is `request_proxy`, range: 0-86400. Amount of time in seconds to wait for one session if `listen_type` is `packets_transmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
        :param str listen_type: The type of LB Listener. Possible values are `request_proxy` and `packets_transmit`.
        :param str method: The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistent_hash`, `source_port` , `consistent_hash_port`, `weight_roundrobin` and `leastconn`. 
               - The `consistent_hash`, `source_port` , `consistent_hash_port`, `roundrobin`, `source` and `weight_roundrobin` are valid if `listen_type` is `packets_transmit`.
               - The `rundrobin`, `source` and `weight_roundrobin` and `leastconn` are vaild if `listen_type` is `request_proxy`.
        :param str name: The name of LB Listener.
        :param str path: Health check path checking.
        :param str persistence: Indicate whether the persistence session is enabled, it is invaild if `persistence_type` is `none`, an auto-generated string will be exported if `persistence_type` is `server_insert`, a custom string will be exported if `persistence_type` is `user_defined`.
        :param str persistence_type: The type of session persistence of LB Listener. Possible values are: `none` as disabled, `server_insert` as auto-generated string and `user_defined` as cutom string. (Default: `none`).
        :param int port: Port opened on the LB Listener to receive requests, range: 1-65535.
        :param str protocol: LB Listener protocol. Possible values: `http`, `https` if `listen_type` is `request_proxy`, `tcp` and `udp` if `listen_type` is `packets_transmit`.
        :param str status: LB Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "health_check_type", health_check_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "idle_timeout", idle_timeout)
        pulumi.set(__self__, "listen_type", listen_type)
        pulumi.set(__self__, "method", method)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "persistence", persistence)
        pulumi.set(__self__, "persistence_type", persistence_type)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Health check domain checking.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="healthCheckType")
    def health_check_type(self) -> str:
        """
        Health check method. Possible values are `port` as port checking and `path` as http checking.
        """
        return pulumi.get(self, "health_check_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of LB Listener.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> int:
        """
        Amount of time in seconds to wait for the response for in between two sessions if `listen_type` is `request_proxy`, range: 0-86400. Amount of time in seconds to wait for one session if `listen_type` is `packets_transmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
        """
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter(name="listenType")
    def listen_type(self) -> str:
        """
        The type of LB Listener. Possible values are `request_proxy` and `packets_transmit`.
        """
        return pulumi.get(self, "listen_type")

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistent_hash`, `source_port` , `consistent_hash_port`, `weight_roundrobin` and `leastconn`. 
        - The `consistent_hash`, `source_port` , `consistent_hash_port`, `roundrobin`, `source` and `weight_roundrobin` are valid if `listen_type` is `packets_transmit`.
        - The `rundrobin`, `source` and `weight_roundrobin` and `leastconn` are vaild if `listen_type` is `request_proxy`.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of LB Listener.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        Health check path checking.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def persistence(self) -> str:
        """
        Indicate whether the persistence session is enabled, it is invaild if `persistence_type` is `none`, an auto-generated string will be exported if `persistence_type` is `server_insert`, a custom string will be exported if `persistence_type` is `user_defined`.
        """
        return pulumi.get(self, "persistence")

    @property
    @pulumi.getter(name="persistenceType")
    def persistence_type(self) -> str:
        """
        The type of session persistence of LB Listener. Possible values are: `none` as disabled, `server_insert` as auto-generated string and `user_defined` as cutom string. (Default: `none`).
        """
        return pulumi.get(self, "persistence_type")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Port opened on the LB Listener to receive requests, range: 1-65535.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        LB Listener protocol. Possible values: `http`, `https` if `listen_type` is `request_proxy`, `tcp` and `udp` if `listen_type` is `packets_transmit`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        LB Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetLBRulesLbRuleResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 id: str,
                 path: str):
        """
        :param str domain: (Optional) The domain of content forward matching fields. `path` and `domain` cannot coexist.
        :param str id: The ID of LB Rule.
        :param str path: (Optional) The path of Content forward matching fields. `path` and `domain` cannot coexist.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        (Optional) The domain of content forward matching fields. `path` and `domain` cannot coexist.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of LB Rule.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        (Optional) The path of Content forward matching fields. `path` and `domain` cannot coexist.
        """
        return pulumi.get(self, "path")


@pulumi.output_type
class GetLBSslLbSslResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 id: str,
                 name: str):
        """
        :param str create_time: The time of creation for lb ssl, formatted in RFC3339 time string.
        :param str id: The ID of LB SSL certificate resource.
        :param str name: The name of LB SSL certificate resource.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time of creation for lb ssl, formatted in RFC3339 time string.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of LB SSL certificate resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of LB SSL certificate resource.
        """
        return pulumi.get(self, "name")

