// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallFirewallPolicyInvalidNameRule checks the pattern is valid
type AwsNetworkfirewallFirewallPolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallFirewallPolicyInvalidNameRule returns new rule with default attributes
func NewAwsNetworkfirewallFirewallPolicyInvalidNameRule() *AwsNetworkfirewallFirewallPolicyInvalidNameRule {
	return &AwsNetworkfirewallFirewallPolicyInvalidNameRule{
		resourceType:  "aws_networkfirewall_firewall_policy",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9-]+$`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallFirewallPolicyInvalidNameRule) Name() string {
	return "aws_networkfirewall_firewall_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallFirewallPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallFirewallPolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallFirewallPolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallFirewallPolicyInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
