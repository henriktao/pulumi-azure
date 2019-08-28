// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Schedule struct {
	s *pulumi.ResourceState
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOpt) (*Schedule, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.NotificationSettings == nil {
		return nil, errors.New("missing required argument 'NotificationSettings'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TaskType == nil {
		return nil, errors.New("missing required argument 'TaskType'")
	}
	if args == nil || args.TimeZoneId == nil {
		return nil, errors.New("missing required argument 'TimeZoneId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dailyRecurrence"] = nil
		inputs["hourlyRecurrence"] = nil
		inputs["labName"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["notificationSettings"] = nil
		inputs["resourceGroupName"] = nil
		inputs["status"] = nil
		inputs["tags"] = nil
		inputs["taskType"] = nil
		inputs["timeZoneId"] = nil
		inputs["weeklyRecurrence"] = nil
	} else {
		inputs["dailyRecurrence"] = args.DailyRecurrence
		inputs["hourlyRecurrence"] = args.HourlyRecurrence
		inputs["labName"] = args.LabName
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["notificationSettings"] = args.NotificationSettings
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["status"] = args.Status
		inputs["tags"] = args.Tags
		inputs["taskType"] = args.TaskType
		inputs["timeZoneId"] = args.TimeZoneId
		inputs["weeklyRecurrence"] = args.WeeklyRecurrence
	}
	s, err := ctx.RegisterResource("azure:devtest/schedule:Schedule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Schedule{s: s}, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ScheduleState, opts ...pulumi.ResourceOpt) (*Schedule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dailyRecurrence"] = state.DailyRecurrence
		inputs["hourlyRecurrence"] = state.HourlyRecurrence
		inputs["labName"] = state.LabName
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["notificationSettings"] = state.NotificationSettings
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["taskType"] = state.TaskType
		inputs["timeZoneId"] = state.TimeZoneId
		inputs["weeklyRecurrence"] = state.WeeklyRecurrence
	}
	s, err := ctx.ReadResource("azure:devtest/schedule:Schedule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Schedule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Schedule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Schedule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Schedule) DailyRecurrence() *pulumi.Output {
	return r.s.State["dailyRecurrence"]
}

func (r *Schedule) HourlyRecurrence() *pulumi.Output {
	return r.s.State["hourlyRecurrence"]
}

func (r *Schedule) LabName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labName"])
}

func (r *Schedule) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

func (r *Schedule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Schedule) NotificationSettings() *pulumi.Output {
	return r.s.State["notificationSettings"]
}

func (r *Schedule) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

func (r *Schedule) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

func (r *Schedule) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

func (r *Schedule) TaskType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["taskType"])
}

func (r *Schedule) TimeZoneId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["timeZoneId"])
}

func (r *Schedule) WeeklyRecurrence() *pulumi.Output {
	return r.s.State["weeklyRecurrence"]
}

// Input properties used for looking up and filtering Schedule resources.
type ScheduleState struct {
	DailyRecurrence interface{}
	HourlyRecurrence interface{}
	LabName interface{}
	Location interface{}
	Name interface{}
	NotificationSettings interface{}
	ResourceGroupName interface{}
	Status interface{}
	Tags interface{}
	TaskType interface{}
	TimeZoneId interface{}
	WeeklyRecurrence interface{}
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	DailyRecurrence interface{}
	HourlyRecurrence interface{}
	LabName interface{}
	Location interface{}
	Name interface{}
	NotificationSettings interface{}
	ResourceGroupName interface{}
	Status interface{}
	Tags interface{}
	TaskType interface{}
	TimeZoneId interface{}
	WeeklyRecurrence interface{}
}
