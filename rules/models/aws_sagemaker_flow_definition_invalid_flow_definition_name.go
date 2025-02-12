// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule checks the pattern is valid
type AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule returns new rule with default attributes
func NewAwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule() *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule {
	return &AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule{
		resourceType:  "aws_sagemaker_flow_definition",
		attributeName: "flow_definition_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-z0-9](-*[a-z0-9]){0,62}`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Name() string {
	return "aws_sagemaker_flow_definition_invalid_flow_definition_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"flow_definition_name must be 63 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"flow_definition_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9](-*[a-z0-9]){0,62}`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
