// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMemoryDBUserInvalidAccessStringRule checks the pattern is valid
type AwsMemoryDBUserInvalidAccessStringRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsMemoryDBUserInvalidAccessStringRule returns new rule with default attributes
func NewAwsMemoryDBUserInvalidAccessStringRule() *AwsMemoryDBUserInvalidAccessStringRule {
	return &AwsMemoryDBUserInvalidAccessStringRule{
		resourceType:  "aws_memorydb_user",
		attributeName: "access_string",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsMemoryDBUserInvalidAccessStringRule) Name() string {
	return "aws_memorydb_user_invalid_access_string"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMemoryDBUserInvalidAccessStringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMemoryDBUserInvalidAccessStringRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMemoryDBUserInvalidAccessStringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMemoryDBUserInvalidAccessStringRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
