// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSchemasSchemaInvalidDescriptionRule checks the pattern is valid
type AwsSchemasSchemaInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSchemasSchemaInvalidDescriptionRule returns new rule with default attributes
func NewAwsSchemasSchemaInvalidDescriptionRule() *AwsSchemasSchemaInvalidDescriptionRule {
	return &AwsSchemasSchemaInvalidDescriptionRule{
		resourceType:  "aws_schemas_schema",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsSchemasSchemaInvalidDescriptionRule) Name() string {
	return "aws_schemas_schema_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSchemasSchemaInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSchemasSchemaInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSchemasSchemaInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSchemasSchemaInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
