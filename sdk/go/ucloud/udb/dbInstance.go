// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package udb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Database instance resource.
//
// > **Note**  Please do confirm if any task pending submission before reset your password, since the password reset will take effect immediately.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/db_instance.html.markdown.
type DbInstance struct {
	s *pulumi.ResourceState
}

// NewDbInstance registers a new resource with the given unique name, arguments, and options.
func NewDbInstance(ctx *pulumi.Context,
	name string, args *DbInstanceArgs, opts ...pulumi.ResourceOpt) (*DbInstance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.Engine == nil {
		return nil, errors.New("missing required argument 'Engine'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.InstanceStorage == nil {
		return nil, errors.New("missing required argument 'InstanceStorage'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["availabilityZone"] = nil
		inputs["backupBeginTime"] = nil
		inputs["backupBlackLists"] = nil
		inputs["backupCount"] = nil
		inputs["backupDate"] = nil
		inputs["chargeType"] = nil
		inputs["duration"] = nil
		inputs["engine"] = nil
		inputs["engineVersion"] = nil
		inputs["instanceStorage"] = nil
		inputs["instanceType"] = nil
		inputs["name"] = nil
		inputs["password"] = nil
		inputs["port"] = nil
		inputs["standbyZone"] = nil
		inputs["subnetId"] = nil
		inputs["tag"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["backupBeginTime"] = args.BackupBeginTime
		inputs["backupBlackLists"] = args.BackupBlackLists
		inputs["backupCount"] = args.BackupCount
		inputs["backupDate"] = args.BackupDate
		inputs["chargeType"] = args.ChargeType
		inputs["duration"] = args.Duration
		inputs["engine"] = args.Engine
		inputs["engineVersion"] = args.EngineVersion
		inputs["instanceStorage"] = args.InstanceStorage
		inputs["instanceType"] = args.InstanceType
		inputs["name"] = args.Name
		inputs["password"] = args.Password
		inputs["port"] = args.Port
		inputs["standbyZone"] = args.StandbyZone
		inputs["subnetId"] = args.SubnetId
		inputs["tag"] = args.Tag
		inputs["vpcId"] = args.VpcId
	}
	inputs["createTime"] = nil
	inputs["expireTime"] = nil
	inputs["modifyTime"] = nil
	inputs["privateIp"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("ucloud:udb/dbInstance:DbInstance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DbInstance{s: s}, nil
}

// GetDbInstance gets an existing DbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DbInstanceState, opts ...pulumi.ResourceOpt) (*DbInstance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["backupBeginTime"] = state.BackupBeginTime
		inputs["backupBlackLists"] = state.BackupBlackLists
		inputs["backupCount"] = state.BackupCount
		inputs["backupDate"] = state.BackupDate
		inputs["chargeType"] = state.ChargeType
		inputs["createTime"] = state.CreateTime
		inputs["duration"] = state.Duration
		inputs["engine"] = state.Engine
		inputs["engineVersion"] = state.EngineVersion
		inputs["expireTime"] = state.ExpireTime
		inputs["instanceStorage"] = state.InstanceStorage
		inputs["instanceType"] = state.InstanceType
		inputs["modifyTime"] = state.ModifyTime
		inputs["name"] = state.Name
		inputs["password"] = state.Password
		inputs["port"] = state.Port
		inputs["privateIp"] = state.PrivateIp
		inputs["standbyZone"] = state.StandbyZone
		inputs["status"] = state.Status
		inputs["subnetId"] = state.SubnetId
		inputs["tag"] = state.Tag
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("ucloud:udb/dbInstance:DbInstance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DbInstance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DbInstance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DbInstance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
func (r *DbInstance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
func (r *DbInstance) BackupBeginTime() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["backupBeginTime"])
}

// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
func (r *DbInstance) BackupBlackLists() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backupBlackLists"])
}

// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
func (r *DbInstance) BackupCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["backupCount"])
}

// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
func (r *DbInstance) BackupDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backupDate"])
}

// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
func (r *DbInstance) ChargeType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["chargeType"])
}

// The creation time of database, formatted by RFC3339 time string.
func (r *DbInstance) CreateTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTime"])
}

// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
func (r *DbInstance) Duration() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["duration"])
}

// The type of database engine, possible values are: "mysql", "percona".
func (r *DbInstance) Engine() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engine"])
}

// The database engine version, possible values are: "5.5", "5.6", "5.7".
// - 5.5/5.6/5.7 for mysql and percona engine.
func (r *DbInstance) EngineVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engineVersion"])
}

// The expiration time of database, formatted by RFC3339 time string.
func (r *DbInstance) ExpireTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expireTime"])
}

// Specifies the allocated storage size in gigabytes (GB), range from 20 to 3000GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
// - 500GB if the memory chosen is equal or less than 8GB;
// - 1000GB if the memory chosen is from 12 to 24GB;
// - 2000GB if the memory chosen is 32GB;
// - 3000GB if the memory chosen is equal or more than 48GB.
func (r *DbInstance) InstanceStorage() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["instanceStorage"])
}

func (r *DbInstance) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

// The modification time of database, formatted by RFC3339 time string.
func (r *DbInstance) ModifyTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["modifyTime"])
}

func (r *DbInstance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *DbInstance) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
func (r *DbInstance) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The private IP address assigned to the database instance.
func (r *DbInstance) PrivateIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIp"])
}

// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
func (r *DbInstance) StandbyZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["standbyZone"])
}

// Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
func (r *DbInstance) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The ID of subnet.
func (r *DbInstance) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
func (r *DbInstance) Tag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tag"])
}

// The ID of VPC linked to the database instances.
func (r *DbInstance) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering DbInstance resources.
type DbInstanceState struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone interface{}
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime interface{}
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists interface{}
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount interface{}
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate interface{}
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType interface{}
	// The creation time of database, formatted by RFC3339 time string.
	CreateTime interface{}
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration interface{}
	// The type of database engine, possible values are: "mysql", "percona".
	Engine interface{}
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion interface{}
	// The expiration time of database, formatted by RFC3339 time string.
	ExpireTime interface{}
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 3000GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 8GB;
	// - 1000GB if the memory chosen is from 12 to 24GB;
	// - 2000GB if the memory chosen is 32GB;
	// - 3000GB if the memory chosen is equal or more than 48GB.
	InstanceStorage interface{}
	InstanceType    interface{}
	// The modification time of database, formatted by RFC3339 time string.
	ModifyTime interface{}
	Name       interface{}
	Password   interface{}
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port interface{}
	// The private IP address assigned to the database instance.
	PrivateIp interface{}
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone interface{}
	// Specifies the status of database, possible values are: `Init`, `Fail`, `Starting`, `Running`, `Shutdown`, `Shutoff`, `Delete`, `Upgrading`, `Promoting`, `Recovering` and `Recover fail`.
	Status interface{}
	// The ID of subnet.
	SubnetId interface{}
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag interface{}
	// The ID of VPC linked to the database instances.
	VpcId interface{}
}

// The set of arguments for constructing a DbInstance resource.
type DbInstanceArgs struct {
	// Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone interface{}
	// Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.
	BackupBeginTime interface{}
	// The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supported.
	BackupBlackLists interface{}
	// Specifies the number of backup saved per week, it is 7 backups saved per week by default.
	BackupCount interface{}
	// Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.
	BackupDate interface{}
	// The charge type of db instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType interface{}
	// The duration that you will buy the db instance (Default: `1`). The value is `0` when pay by month and the instance will be vaild till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration interface{}
	// The type of database engine, possible values are: "mysql", "percona".
	Engine interface{}
	// The database engine version, possible values are: "5.5", "5.6", "5.7".
	// - 5.5/5.6/5.7 for mysql and percona engine.
	EngineVersion interface{}
	// Specifies the allocated storage size in gigabytes (GB), range from 20 to 3000GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are：
	// - 500GB if the memory chosen is equal or less than 8GB;
	// - 1000GB if the memory chosen is from 12 to 24GB;
	// - 2000GB if the memory chosen is 32GB;
	// - 3000GB if the memory chosen is equal or more than 48GB.
	InstanceStorage interface{}
	InstanceType    interface{}
	Name            interface{}
	Password        interface{}
	// The port on which the database accepts connections, the default port is 3306 for mysql and percona.
	Port interface{}
	// Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.
	StandbyZone interface{}
	// The ID of subnet.
	SubnetId interface{}
	// A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag interface{}
	// The ID of VPC linked to the database instances.
	VpcId interface{}
}