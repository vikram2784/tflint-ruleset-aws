// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallConfigInvalidResourceIDRule checks the pattern is valid
type AwsRoute53ResolverFirewallConfigInvalidResourceIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverFirewallConfigInvalidResourceIDRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallConfigInvalidResourceIDRule() *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule {
	return &AwsRoute53ResolverFirewallConfigInvalidResourceIDRule{
		resourceType:  "aws_route53_resolver_firewall_config",
		attributeName: "resource_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule) Name() string {
	return "aws_route53_resolver_firewall_config_invalid_resource_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallConfigInvalidResourceIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"resource_id must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"resource_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
