// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupPolicyDiskRetentionRule struct {
	// A `criteria` block as defined below. Changing this forces a new Backup Policy Disk to be created.
	Criteria BackupPolicyDiskRetentionRuleCriteria `pulumi:"criteria"`
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Disk to be created.
	Duration string `pulumi:"duration"`
	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
	Name string `pulumi:"name"`
	// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
	Priority int `pulumi:"priority"`
}

// BackupPolicyDiskRetentionRuleInput is an input type that accepts BackupPolicyDiskRetentionRuleArgs and BackupPolicyDiskRetentionRuleOutput values.
// You can construct a concrete instance of `BackupPolicyDiskRetentionRuleInput` via:
//
//          BackupPolicyDiskRetentionRuleArgs{...}
type BackupPolicyDiskRetentionRuleInput interface {
	pulumi.Input

	ToBackupPolicyDiskRetentionRuleOutput() BackupPolicyDiskRetentionRuleOutput
	ToBackupPolicyDiskRetentionRuleOutputWithContext(context.Context) BackupPolicyDiskRetentionRuleOutput
}

type BackupPolicyDiskRetentionRuleArgs struct {
	// A `criteria` block as defined below. Changing this forces a new Backup Policy Disk to be created.
	Criteria BackupPolicyDiskRetentionRuleCriteriaInput `pulumi:"criteria"`
	// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Disk to be created.
	Duration pulumi.StringInput `pulumi:"duration"`
	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
	Priority pulumi.IntInput `pulumi:"priority"`
}

func (BackupPolicyDiskRetentionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyDiskRetentionRule)(nil)).Elem()
}

func (i BackupPolicyDiskRetentionRuleArgs) ToBackupPolicyDiskRetentionRuleOutput() BackupPolicyDiskRetentionRuleOutput {
	return i.ToBackupPolicyDiskRetentionRuleOutputWithContext(context.Background())
}

func (i BackupPolicyDiskRetentionRuleArgs) ToBackupPolicyDiskRetentionRuleOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyDiskRetentionRuleOutput)
}

// BackupPolicyDiskRetentionRuleArrayInput is an input type that accepts BackupPolicyDiskRetentionRuleArray and BackupPolicyDiskRetentionRuleArrayOutput values.
// You can construct a concrete instance of `BackupPolicyDiskRetentionRuleArrayInput` via:
//
//          BackupPolicyDiskRetentionRuleArray{ BackupPolicyDiskRetentionRuleArgs{...} }
type BackupPolicyDiskRetentionRuleArrayInput interface {
	pulumi.Input

	ToBackupPolicyDiskRetentionRuleArrayOutput() BackupPolicyDiskRetentionRuleArrayOutput
	ToBackupPolicyDiskRetentionRuleArrayOutputWithContext(context.Context) BackupPolicyDiskRetentionRuleArrayOutput
}

type BackupPolicyDiskRetentionRuleArray []BackupPolicyDiskRetentionRuleInput

func (BackupPolicyDiskRetentionRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupPolicyDiskRetentionRule)(nil)).Elem()
}

func (i BackupPolicyDiskRetentionRuleArray) ToBackupPolicyDiskRetentionRuleArrayOutput() BackupPolicyDiskRetentionRuleArrayOutput {
	return i.ToBackupPolicyDiskRetentionRuleArrayOutputWithContext(context.Background())
}

func (i BackupPolicyDiskRetentionRuleArray) ToBackupPolicyDiskRetentionRuleArrayOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyDiskRetentionRuleArrayOutput)
}

type BackupPolicyDiskRetentionRuleOutput struct{ *pulumi.OutputState }

func (BackupPolicyDiskRetentionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyDiskRetentionRule)(nil)).Elem()
}

func (o BackupPolicyDiskRetentionRuleOutput) ToBackupPolicyDiskRetentionRuleOutput() BackupPolicyDiskRetentionRuleOutput {
	return o
}

func (o BackupPolicyDiskRetentionRuleOutput) ToBackupPolicyDiskRetentionRuleOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleOutput {
	return o
}

// A `criteria` block as defined below. Changing this forces a new Backup Policy Disk to be created.
func (o BackupPolicyDiskRetentionRuleOutput) Criteria() BackupPolicyDiskRetentionRuleCriteriaOutput {
	return o.ApplyT(func(v BackupPolicyDiskRetentionRule) BackupPolicyDiskRetentionRuleCriteria { return v.Criteria }).(BackupPolicyDiskRetentionRuleCriteriaOutput)
}

// Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Disk to be created.
func (o BackupPolicyDiskRetentionRuleOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyDiskRetentionRule) string { return v.Duration }).(pulumi.StringOutput)
}

// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
func (o BackupPolicyDiskRetentionRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyDiskRetentionRule) string { return v.Name }).(pulumi.StringOutput)
}

// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
func (o BackupPolicyDiskRetentionRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v BackupPolicyDiskRetentionRule) int { return v.Priority }).(pulumi.IntOutput)
}

type BackupPolicyDiskRetentionRuleArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyDiskRetentionRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupPolicyDiskRetentionRule)(nil)).Elem()
}

func (o BackupPolicyDiskRetentionRuleArrayOutput) ToBackupPolicyDiskRetentionRuleArrayOutput() BackupPolicyDiskRetentionRuleArrayOutput {
	return o
}

func (o BackupPolicyDiskRetentionRuleArrayOutput) ToBackupPolicyDiskRetentionRuleArrayOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleArrayOutput {
	return o
}

func (o BackupPolicyDiskRetentionRuleArrayOutput) Index(i pulumi.IntInput) BackupPolicyDiskRetentionRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupPolicyDiskRetentionRule {
		return vs[0].([]BackupPolicyDiskRetentionRule)[vs[1].(int)]
	}).(BackupPolicyDiskRetentionRuleOutput)
}

type BackupPolicyDiskRetentionRuleCriteria struct {
	// Possible values are `FirstOfDay` and `FirstOfWeek`. Changing this forces a new Backup Policy Disk to be created.
	AbsoluteCriteria *string `pulumi:"absoluteCriteria"`
}

// BackupPolicyDiskRetentionRuleCriteriaInput is an input type that accepts BackupPolicyDiskRetentionRuleCriteriaArgs and BackupPolicyDiskRetentionRuleCriteriaOutput values.
// You can construct a concrete instance of `BackupPolicyDiskRetentionRuleCriteriaInput` via:
//
//          BackupPolicyDiskRetentionRuleCriteriaArgs{...}
type BackupPolicyDiskRetentionRuleCriteriaInput interface {
	pulumi.Input

	ToBackupPolicyDiskRetentionRuleCriteriaOutput() BackupPolicyDiskRetentionRuleCriteriaOutput
	ToBackupPolicyDiskRetentionRuleCriteriaOutputWithContext(context.Context) BackupPolicyDiskRetentionRuleCriteriaOutput
}

type BackupPolicyDiskRetentionRuleCriteriaArgs struct {
	// Possible values are `FirstOfDay` and `FirstOfWeek`. Changing this forces a new Backup Policy Disk to be created.
	AbsoluteCriteria pulumi.StringPtrInput `pulumi:"absoluteCriteria"`
}

func (BackupPolicyDiskRetentionRuleCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyDiskRetentionRuleCriteria)(nil)).Elem()
}

func (i BackupPolicyDiskRetentionRuleCriteriaArgs) ToBackupPolicyDiskRetentionRuleCriteriaOutput() BackupPolicyDiskRetentionRuleCriteriaOutput {
	return i.ToBackupPolicyDiskRetentionRuleCriteriaOutputWithContext(context.Background())
}

func (i BackupPolicyDiskRetentionRuleCriteriaArgs) ToBackupPolicyDiskRetentionRuleCriteriaOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyDiskRetentionRuleCriteriaOutput)
}

type BackupPolicyDiskRetentionRuleCriteriaOutput struct{ *pulumi.OutputState }

func (BackupPolicyDiskRetentionRuleCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyDiskRetentionRuleCriteria)(nil)).Elem()
}

func (o BackupPolicyDiskRetentionRuleCriteriaOutput) ToBackupPolicyDiskRetentionRuleCriteriaOutput() BackupPolicyDiskRetentionRuleCriteriaOutput {
	return o
}

func (o BackupPolicyDiskRetentionRuleCriteriaOutput) ToBackupPolicyDiskRetentionRuleCriteriaOutputWithContext(ctx context.Context) BackupPolicyDiskRetentionRuleCriteriaOutput {
	return o
}

// Possible values are `FirstOfDay` and `FirstOfWeek`. Changing this forces a new Backup Policy Disk to be created.
func (o BackupPolicyDiskRetentionRuleCriteriaOutput) AbsoluteCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyDiskRetentionRuleCriteria) *string { return v.AbsoluteCriteria }).(pulumi.StringPtrOutput)
}

type BackupPolicyPostgresqlRetentionRule struct {
	// A `criteria` block as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
	Criteria BackupPolicyPostgresqlRetentionRuleCriteria `pulumi:"criteria"`
	// Duration after which the backup is deleted. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
	Duration string `pulumi:"duration"`
	// The name which should be used for this retention rule. Changing this forces a new Backup Policy PostgreSQL to be created.
	Name string `pulumi:"name"`
	// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Postgre Sql to be created.
	Priority int `pulumi:"priority"`
}

// BackupPolicyPostgresqlRetentionRuleInput is an input type that accepts BackupPolicyPostgresqlRetentionRuleArgs and BackupPolicyPostgresqlRetentionRuleOutput values.
// You can construct a concrete instance of `BackupPolicyPostgresqlRetentionRuleInput` via:
//
//          BackupPolicyPostgresqlRetentionRuleArgs{...}
type BackupPolicyPostgresqlRetentionRuleInput interface {
	pulumi.Input

	ToBackupPolicyPostgresqlRetentionRuleOutput() BackupPolicyPostgresqlRetentionRuleOutput
	ToBackupPolicyPostgresqlRetentionRuleOutputWithContext(context.Context) BackupPolicyPostgresqlRetentionRuleOutput
}

type BackupPolicyPostgresqlRetentionRuleArgs struct {
	// A `criteria` block as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
	Criteria BackupPolicyPostgresqlRetentionRuleCriteriaInput `pulumi:"criteria"`
	// Duration after which the backup is deleted. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
	Duration pulumi.StringInput `pulumi:"duration"`
	// The name which should be used for this retention rule. Changing this forces a new Backup Policy PostgreSQL to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Postgre Sql to be created.
	Priority pulumi.IntInput `pulumi:"priority"`
}

func (BackupPolicyPostgresqlRetentionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyPostgresqlRetentionRule)(nil)).Elem()
}

func (i BackupPolicyPostgresqlRetentionRuleArgs) ToBackupPolicyPostgresqlRetentionRuleOutput() BackupPolicyPostgresqlRetentionRuleOutput {
	return i.ToBackupPolicyPostgresqlRetentionRuleOutputWithContext(context.Background())
}

func (i BackupPolicyPostgresqlRetentionRuleArgs) ToBackupPolicyPostgresqlRetentionRuleOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyPostgresqlRetentionRuleOutput)
}

// BackupPolicyPostgresqlRetentionRuleArrayInput is an input type that accepts BackupPolicyPostgresqlRetentionRuleArray and BackupPolicyPostgresqlRetentionRuleArrayOutput values.
// You can construct a concrete instance of `BackupPolicyPostgresqlRetentionRuleArrayInput` via:
//
//          BackupPolicyPostgresqlRetentionRuleArray{ BackupPolicyPostgresqlRetentionRuleArgs{...} }
type BackupPolicyPostgresqlRetentionRuleArrayInput interface {
	pulumi.Input

	ToBackupPolicyPostgresqlRetentionRuleArrayOutput() BackupPolicyPostgresqlRetentionRuleArrayOutput
	ToBackupPolicyPostgresqlRetentionRuleArrayOutputWithContext(context.Context) BackupPolicyPostgresqlRetentionRuleArrayOutput
}

type BackupPolicyPostgresqlRetentionRuleArray []BackupPolicyPostgresqlRetentionRuleInput

func (BackupPolicyPostgresqlRetentionRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupPolicyPostgresqlRetentionRule)(nil)).Elem()
}

func (i BackupPolicyPostgresqlRetentionRuleArray) ToBackupPolicyPostgresqlRetentionRuleArrayOutput() BackupPolicyPostgresqlRetentionRuleArrayOutput {
	return i.ToBackupPolicyPostgresqlRetentionRuleArrayOutputWithContext(context.Background())
}

func (i BackupPolicyPostgresqlRetentionRuleArray) ToBackupPolicyPostgresqlRetentionRuleArrayOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyPostgresqlRetentionRuleArrayOutput)
}

type BackupPolicyPostgresqlRetentionRuleOutput struct{ *pulumi.OutputState }

func (BackupPolicyPostgresqlRetentionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyPostgresqlRetentionRule)(nil)).Elem()
}

func (o BackupPolicyPostgresqlRetentionRuleOutput) ToBackupPolicyPostgresqlRetentionRuleOutput() BackupPolicyPostgresqlRetentionRuleOutput {
	return o
}

func (o BackupPolicyPostgresqlRetentionRuleOutput) ToBackupPolicyPostgresqlRetentionRuleOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleOutput {
	return o
}

// A `criteria` block as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleOutput) Criteria() BackupPolicyPostgresqlRetentionRuleCriteriaOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRule) BackupPolicyPostgresqlRetentionRuleCriteria {
		return v.Criteria
	}).(BackupPolicyPostgresqlRetentionRuleCriteriaOutput)
}

// Duration after which the backup is deleted. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRule) string { return v.Duration }).(pulumi.StringOutput)
}

// The name which should be used for this retention rule. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRule) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Postgre Sql to be created.
func (o BackupPolicyPostgresqlRetentionRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRule) int { return v.Priority }).(pulumi.IntOutput)
}

type BackupPolicyPostgresqlRetentionRuleArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyPostgresqlRetentionRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupPolicyPostgresqlRetentionRule)(nil)).Elem()
}

func (o BackupPolicyPostgresqlRetentionRuleArrayOutput) ToBackupPolicyPostgresqlRetentionRuleArrayOutput() BackupPolicyPostgresqlRetentionRuleArrayOutput {
	return o
}

func (o BackupPolicyPostgresqlRetentionRuleArrayOutput) ToBackupPolicyPostgresqlRetentionRuleArrayOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleArrayOutput {
	return o
}

func (o BackupPolicyPostgresqlRetentionRuleArrayOutput) Index(i pulumi.IntInput) BackupPolicyPostgresqlRetentionRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupPolicyPostgresqlRetentionRule {
		return vs[0].([]BackupPolicyPostgresqlRetentionRule)[vs[1].(int)]
	}).(BackupPolicyPostgresqlRetentionRuleOutput)
}

type BackupPolicyPostgresqlRetentionRuleCriteria struct {
	// Possible values are `AllBackup`, `FirstOfDay`, `FirstOfWeek`, `FirstOfMonth` and `FirstOfYear`. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy PostgreSQL to be created.
	AbsoluteCriteria *string `pulumi:"absoluteCriteria"`
	// Possible values are `Monday`, `Tuesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`. Changing this forces a new Backup Policy PostgreSQL to be created.
	DaysOfWeeks []string `pulumi:"daysOfWeeks"`
	// Possible values are `January`, `February`, `March`, `April`, `May`, `June`, `July`, `August`, `September`, `October`, `November` and `December`. Changing this forces a new Backup Policy PostgreSQL to be created.
	MonthsOfYears []string `pulumi:"monthsOfYears"`
	// Specifies a list of backup times for backup in the `RFC3339` format. Changing this forces a new Backup Policy Postgre Sql to be created.
	ScheduledBackupTimes []string `pulumi:"scheduledBackupTimes"`
	// Possible values are `First`, `Second`, `Third`, `Fourth` and `Last`. Changing this forces a new Backup Policy PostgreSQL to be created.
	WeeksOfMonths []string `pulumi:"weeksOfMonths"`
}

// BackupPolicyPostgresqlRetentionRuleCriteriaInput is an input type that accepts BackupPolicyPostgresqlRetentionRuleCriteriaArgs and BackupPolicyPostgresqlRetentionRuleCriteriaOutput values.
// You can construct a concrete instance of `BackupPolicyPostgresqlRetentionRuleCriteriaInput` via:
//
//          BackupPolicyPostgresqlRetentionRuleCriteriaArgs{...}
type BackupPolicyPostgresqlRetentionRuleCriteriaInput interface {
	pulumi.Input

	ToBackupPolicyPostgresqlRetentionRuleCriteriaOutput() BackupPolicyPostgresqlRetentionRuleCriteriaOutput
	ToBackupPolicyPostgresqlRetentionRuleCriteriaOutputWithContext(context.Context) BackupPolicyPostgresqlRetentionRuleCriteriaOutput
}

type BackupPolicyPostgresqlRetentionRuleCriteriaArgs struct {
	// Possible values are `AllBackup`, `FirstOfDay`, `FirstOfWeek`, `FirstOfMonth` and `FirstOfYear`. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy PostgreSQL to be created.
	AbsoluteCriteria pulumi.StringPtrInput `pulumi:"absoluteCriteria"`
	// Possible values are `Monday`, `Tuesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`. Changing this forces a new Backup Policy PostgreSQL to be created.
	DaysOfWeeks pulumi.StringArrayInput `pulumi:"daysOfWeeks"`
	// Possible values are `January`, `February`, `March`, `April`, `May`, `June`, `July`, `August`, `September`, `October`, `November` and `December`. Changing this forces a new Backup Policy PostgreSQL to be created.
	MonthsOfYears pulumi.StringArrayInput `pulumi:"monthsOfYears"`
	// Specifies a list of backup times for backup in the `RFC3339` format. Changing this forces a new Backup Policy Postgre Sql to be created.
	ScheduledBackupTimes pulumi.StringArrayInput `pulumi:"scheduledBackupTimes"`
	// Possible values are `First`, `Second`, `Third`, `Fourth` and `Last`. Changing this forces a new Backup Policy PostgreSQL to be created.
	WeeksOfMonths pulumi.StringArrayInput `pulumi:"weeksOfMonths"`
}

func (BackupPolicyPostgresqlRetentionRuleCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyPostgresqlRetentionRuleCriteria)(nil)).Elem()
}

func (i BackupPolicyPostgresqlRetentionRuleCriteriaArgs) ToBackupPolicyPostgresqlRetentionRuleCriteriaOutput() BackupPolicyPostgresqlRetentionRuleCriteriaOutput {
	return i.ToBackupPolicyPostgresqlRetentionRuleCriteriaOutputWithContext(context.Background())
}

func (i BackupPolicyPostgresqlRetentionRuleCriteriaArgs) ToBackupPolicyPostgresqlRetentionRuleCriteriaOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyPostgresqlRetentionRuleCriteriaOutput)
}

type BackupPolicyPostgresqlRetentionRuleCriteriaOutput struct{ *pulumi.OutputState }

func (BackupPolicyPostgresqlRetentionRuleCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyPostgresqlRetentionRuleCriteria)(nil)).Elem()
}

func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) ToBackupPolicyPostgresqlRetentionRuleCriteriaOutput() BackupPolicyPostgresqlRetentionRuleCriteriaOutput {
	return o
}

func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) ToBackupPolicyPostgresqlRetentionRuleCriteriaOutputWithContext(ctx context.Context) BackupPolicyPostgresqlRetentionRuleCriteriaOutput {
	return o
}

// Possible values are `AllBackup`, `FirstOfDay`, `FirstOfWeek`, `FirstOfMonth` and `FirstOfYear`. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) AbsoluteCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRuleCriteria) *string { return v.AbsoluteCriteria }).(pulumi.StringPtrOutput)
}

// Possible values are `Monday`, `Tuesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) DaysOfWeeks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRuleCriteria) []string { return v.DaysOfWeeks }).(pulumi.StringArrayOutput)
}

// Possible values are `January`, `February`, `March`, `April`, `May`, `June`, `July`, `August`, `September`, `October`, `November` and `December`. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) MonthsOfYears() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRuleCriteria) []string { return v.MonthsOfYears }).(pulumi.StringArrayOutput)
}

// Specifies a list of backup times for backup in the `RFC3339` format. Changing this forces a new Backup Policy Postgre Sql to be created.
func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) ScheduledBackupTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRuleCriteria) []string { return v.ScheduledBackupTimes }).(pulumi.StringArrayOutput)
}

// Possible values are `First`, `Second`, `Third`, `Fourth` and `Last`. Changing this forces a new Backup Policy PostgreSQL to be created.
func (o BackupPolicyPostgresqlRetentionRuleCriteriaOutput) WeeksOfMonths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackupPolicyPostgresqlRetentionRuleCriteria) []string { return v.WeeksOfMonths }).(pulumi.StringArrayOutput)
}

type BackupVaultIdentity struct {
	// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
	PrincipalId *string `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
	TenantId *string `pulumi:"tenantId"`
	// Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
	Type string `pulumi:"type"`
}

// BackupVaultIdentityInput is an input type that accepts BackupVaultIdentityArgs and BackupVaultIdentityOutput values.
// You can construct a concrete instance of `BackupVaultIdentityInput` via:
//
//          BackupVaultIdentityArgs{...}
type BackupVaultIdentityInput interface {
	pulumi.Input

	ToBackupVaultIdentityOutput() BackupVaultIdentityOutput
	ToBackupVaultIdentityOutputWithContext(context.Context) BackupVaultIdentityOutput
}

type BackupVaultIdentityArgs struct {
	// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (BackupVaultIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVaultIdentity)(nil)).Elem()
}

func (i BackupVaultIdentityArgs) ToBackupVaultIdentityOutput() BackupVaultIdentityOutput {
	return i.ToBackupVaultIdentityOutputWithContext(context.Background())
}

func (i BackupVaultIdentityArgs) ToBackupVaultIdentityOutputWithContext(ctx context.Context) BackupVaultIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultIdentityOutput)
}

func (i BackupVaultIdentityArgs) ToBackupVaultIdentityPtrOutput() BackupVaultIdentityPtrOutput {
	return i.ToBackupVaultIdentityPtrOutputWithContext(context.Background())
}

func (i BackupVaultIdentityArgs) ToBackupVaultIdentityPtrOutputWithContext(ctx context.Context) BackupVaultIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultIdentityOutput).ToBackupVaultIdentityPtrOutputWithContext(ctx)
}

// BackupVaultIdentityPtrInput is an input type that accepts BackupVaultIdentityArgs, BackupVaultIdentityPtr and BackupVaultIdentityPtrOutput values.
// You can construct a concrete instance of `BackupVaultIdentityPtrInput` via:
//
//          BackupVaultIdentityArgs{...}
//
//  or:
//
//          nil
type BackupVaultIdentityPtrInput interface {
	pulumi.Input

	ToBackupVaultIdentityPtrOutput() BackupVaultIdentityPtrOutput
	ToBackupVaultIdentityPtrOutputWithContext(context.Context) BackupVaultIdentityPtrOutput
}

type backupVaultIdentityPtrType BackupVaultIdentityArgs

func BackupVaultIdentityPtr(v *BackupVaultIdentityArgs) BackupVaultIdentityPtrInput {
	return (*backupVaultIdentityPtrType)(v)
}

func (*backupVaultIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVaultIdentity)(nil)).Elem()
}

func (i *backupVaultIdentityPtrType) ToBackupVaultIdentityPtrOutput() BackupVaultIdentityPtrOutput {
	return i.ToBackupVaultIdentityPtrOutputWithContext(context.Background())
}

func (i *backupVaultIdentityPtrType) ToBackupVaultIdentityPtrOutputWithContext(ctx context.Context) BackupVaultIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultIdentityPtrOutput)
}

type BackupVaultIdentityOutput struct{ *pulumi.OutputState }

func (BackupVaultIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVaultIdentity)(nil)).Elem()
}

func (o BackupVaultIdentityOutput) ToBackupVaultIdentityOutput() BackupVaultIdentityOutput {
	return o
}

func (o BackupVaultIdentityOutput) ToBackupVaultIdentityOutputWithContext(ctx context.Context) BackupVaultIdentityOutput {
	return o
}

func (o BackupVaultIdentityOutput) ToBackupVaultIdentityPtrOutput() BackupVaultIdentityPtrOutput {
	return o.ToBackupVaultIdentityPtrOutputWithContext(context.Background())
}

func (o BackupVaultIdentityOutput) ToBackupVaultIdentityPtrOutputWithContext(ctx context.Context) BackupVaultIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupVaultIdentity) *BackupVaultIdentity {
		return &v
	}).(BackupVaultIdentityPtrOutput)
}

// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
func (o BackupVaultIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupVaultIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
func (o BackupVaultIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupVaultIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
func (o BackupVaultIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BackupVaultIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type BackupVaultIdentityPtrOutput struct{ *pulumi.OutputState }

func (BackupVaultIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVaultIdentity)(nil)).Elem()
}

func (o BackupVaultIdentityPtrOutput) ToBackupVaultIdentityPtrOutput() BackupVaultIdentityPtrOutput {
	return o
}

func (o BackupVaultIdentityPtrOutput) ToBackupVaultIdentityPtrOutputWithContext(ctx context.Context) BackupVaultIdentityPtrOutput {
	return o
}

func (o BackupVaultIdentityPtrOutput) Elem() BackupVaultIdentityOutput {
	return o.ApplyT(func(v *BackupVaultIdentity) BackupVaultIdentity {
		if v != nil {
			return *v
		}
		var ret BackupVaultIdentity
		return ret
	}).(BackupVaultIdentityOutput)
}

// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
func (o BackupVaultIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupVaultIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
func (o BackupVaultIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupVaultIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
func (o BackupVaultIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupVaultIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type GetBackupVaultIdentity struct {
	// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
	PrincipalId string `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
	TenantId string `pulumi:"tenantId"`
	// Specifies the identity type of the Backup Vault.
	Type string `pulumi:"type"`
}

// GetBackupVaultIdentityInput is an input type that accepts GetBackupVaultIdentityArgs and GetBackupVaultIdentityOutput values.
// You can construct a concrete instance of `GetBackupVaultIdentityInput` via:
//
//          GetBackupVaultIdentityArgs{...}
type GetBackupVaultIdentityInput interface {
	pulumi.Input

	ToGetBackupVaultIdentityOutput() GetBackupVaultIdentityOutput
	ToGetBackupVaultIdentityOutputWithContext(context.Context) GetBackupVaultIdentityOutput
}

type GetBackupVaultIdentityArgs struct {
	// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
	// Specifies the identity type of the Backup Vault.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetBackupVaultIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupVaultIdentity)(nil)).Elem()
}

func (i GetBackupVaultIdentityArgs) ToGetBackupVaultIdentityOutput() GetBackupVaultIdentityOutput {
	return i.ToGetBackupVaultIdentityOutputWithContext(context.Background())
}

func (i GetBackupVaultIdentityArgs) ToGetBackupVaultIdentityOutputWithContext(ctx context.Context) GetBackupVaultIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackupVaultIdentityOutput)
}

// GetBackupVaultIdentityArrayInput is an input type that accepts GetBackupVaultIdentityArray and GetBackupVaultIdentityArrayOutput values.
// You can construct a concrete instance of `GetBackupVaultIdentityArrayInput` via:
//
//          GetBackupVaultIdentityArray{ GetBackupVaultIdentityArgs{...} }
type GetBackupVaultIdentityArrayInput interface {
	pulumi.Input

	ToGetBackupVaultIdentityArrayOutput() GetBackupVaultIdentityArrayOutput
	ToGetBackupVaultIdentityArrayOutputWithContext(context.Context) GetBackupVaultIdentityArrayOutput
}

type GetBackupVaultIdentityArray []GetBackupVaultIdentityInput

func (GetBackupVaultIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackupVaultIdentity)(nil)).Elem()
}

func (i GetBackupVaultIdentityArray) ToGetBackupVaultIdentityArrayOutput() GetBackupVaultIdentityArrayOutput {
	return i.ToGetBackupVaultIdentityArrayOutputWithContext(context.Background())
}

func (i GetBackupVaultIdentityArray) ToGetBackupVaultIdentityArrayOutputWithContext(ctx context.Context) GetBackupVaultIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetBackupVaultIdentityArrayOutput)
}

type GetBackupVaultIdentityOutput struct{ *pulumi.OutputState }

func (GetBackupVaultIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupVaultIdentity)(nil)).Elem()
}

func (o GetBackupVaultIdentityOutput) ToGetBackupVaultIdentityOutput() GetBackupVaultIdentityOutput {
	return o
}

func (o GetBackupVaultIdentityOutput) ToGetBackupVaultIdentityOutputWithContext(ctx context.Context) GetBackupVaultIdentityOutput {
	return o
}

// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
func (o GetBackupVaultIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupVaultIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
func (o GetBackupVaultIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupVaultIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

// Specifies the identity type of the Backup Vault.
func (o GetBackupVaultIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupVaultIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type GetBackupVaultIdentityArrayOutput struct{ *pulumi.OutputState }

func (GetBackupVaultIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetBackupVaultIdentity)(nil)).Elem()
}

func (o GetBackupVaultIdentityArrayOutput) ToGetBackupVaultIdentityArrayOutput() GetBackupVaultIdentityArrayOutput {
	return o
}

func (o GetBackupVaultIdentityArrayOutput) ToGetBackupVaultIdentityArrayOutputWithContext(ctx context.Context) GetBackupVaultIdentityArrayOutput {
	return o
}

func (o GetBackupVaultIdentityArrayOutput) Index(i pulumi.IntInput) GetBackupVaultIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetBackupVaultIdentity {
		return vs[0].([]GetBackupVaultIdentity)[vs[1].(int)]
	}).(GetBackupVaultIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyDiskRetentionRuleInput)(nil)).Elem(), BackupPolicyDiskRetentionRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyDiskRetentionRuleArrayInput)(nil)).Elem(), BackupPolicyDiskRetentionRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyDiskRetentionRuleCriteriaInput)(nil)).Elem(), BackupPolicyDiskRetentionRuleCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyPostgresqlRetentionRuleInput)(nil)).Elem(), BackupPolicyPostgresqlRetentionRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyPostgresqlRetentionRuleArrayInput)(nil)).Elem(), BackupPolicyPostgresqlRetentionRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyPostgresqlRetentionRuleCriteriaInput)(nil)).Elem(), BackupPolicyPostgresqlRetentionRuleCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupVaultIdentityInput)(nil)).Elem(), BackupVaultIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupVaultIdentityPtrInput)(nil)).Elem(), BackupVaultIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetBackupVaultIdentityInput)(nil)).Elem(), GetBackupVaultIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetBackupVaultIdentityArrayInput)(nil)).Elem(), GetBackupVaultIdentityArray{})
	pulumi.RegisterOutputType(BackupPolicyDiskRetentionRuleOutput{})
	pulumi.RegisterOutputType(BackupPolicyDiskRetentionRuleArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyDiskRetentionRuleCriteriaOutput{})
	pulumi.RegisterOutputType(BackupPolicyPostgresqlRetentionRuleOutput{})
	pulumi.RegisterOutputType(BackupPolicyPostgresqlRetentionRuleArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyPostgresqlRetentionRuleCriteriaOutput{})
	pulumi.RegisterOutputType(BackupVaultIdentityOutput{})
	pulumi.RegisterOutputType(BackupVaultIdentityPtrOutput{})
	pulumi.RegisterOutputType(GetBackupVaultIdentityOutput{})
	pulumi.RegisterOutputType(GetBackupVaultIdentityArrayOutput{})
}
