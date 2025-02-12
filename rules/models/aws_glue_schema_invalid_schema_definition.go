// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlueSchemaInvalidSchemaDefinitionRule checks the pattern is valid
type AwsGlueSchemaInvalidSchemaDefinitionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsGlueSchemaInvalidSchemaDefinitionRule returns new rule with default attributes
func NewAwsGlueSchemaInvalidSchemaDefinitionRule() *AwsGlueSchemaInvalidSchemaDefinitionRule {
	return &AwsGlueSchemaInvalidSchemaDefinitionRule{
		resourceType:  "aws_glue_schema",
		attributeName: "schema_definition",
		max:           170000,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Name() string {
	return "aws_glue_schema_invalid_schema_definition"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"schema_definition must be 170000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"schema_definition must be 1 characters or higher",
					attribute.Expr,
				)
			}
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
