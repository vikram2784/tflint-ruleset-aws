// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallFirewallInvalidVpcIDRule checks the pattern is valid
type AwsNetworkfirewallFirewallInvalidVpcIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallFirewallInvalidVpcIDRule returns new rule with default attributes
func NewAwsNetworkfirewallFirewallInvalidVpcIDRule() *AwsNetworkfirewallFirewallInvalidVpcIDRule {
	return &AwsNetworkfirewallFirewallInvalidVpcIDRule{
		resourceType:  "aws_networkfirewall_firewall",
		attributeName: "vpc_id",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^vpc-[0-9a-f]+$`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallFirewallInvalidVpcIDRule) Name() string {
	return "aws_networkfirewall_firewall_invalid_vpc_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallFirewallInvalidVpcIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallFirewallInvalidVpcIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallFirewallInvalidVpcIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallFirewallInvalidVpcIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"vpc_id must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"vpc_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^vpc-[0-9a-f]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
