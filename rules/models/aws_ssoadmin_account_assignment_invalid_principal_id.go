// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminAccountAssignmentInvalidPrincipalIDRule checks the pattern is valid
type AwsSsoadminAccountAssignmentInvalidPrincipalIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsoadminAccountAssignmentInvalidPrincipalIDRule returns new rule with default attributes
func NewAwsSsoadminAccountAssignmentInvalidPrincipalIDRule() *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule {
	return &AwsSsoadminAccountAssignmentInvalidPrincipalIDRule{
		resourceType:  "aws_ssoadmin_account_assignment",
		attributeName: "principal_id",
		max:           47,
		min:           1,
		pattern:       regexp.MustCompile(`^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule) Name() string {
	return "aws_ssoadmin_account_assignment_invalid_principal_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"principal_id must be 47 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"principal_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
