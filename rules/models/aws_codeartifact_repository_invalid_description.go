// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactRepositoryInvalidDescriptionRule checks the pattern is valid
type AwsCodeartifactRepositoryInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactRepositoryInvalidDescriptionRule returns new rule with default attributes
func NewAwsCodeartifactRepositoryInvalidDescriptionRule() *AwsCodeartifactRepositoryInvalidDescriptionRule {
	return &AwsCodeartifactRepositoryInvalidDescriptionRule{
		resourceType:  "aws_codeartifact_repository",
		attributeName: "description",
		max:           1000,
		pattern:       regexp.MustCompile(`^\P{C}+$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactRepositoryInvalidDescriptionRule) Name() string {
	return "aws_codeartifact_repository_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactRepositoryInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactRepositoryInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactRepositoryInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactRepositoryInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1000 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\P{C}+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
