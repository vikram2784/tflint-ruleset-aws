// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidSessionNameRule checks the pattern is valid
type AwsQuicksightUserInvalidSessionNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightUserInvalidSessionNameRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidSessionNameRule() *AwsQuicksightUserInvalidSessionNameRule {
	return &AwsQuicksightUserInvalidSessionNameRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "session_name",
		max:           64,
		min:           2,
		pattern:       regexp.MustCompile(`^[\w+=.@-]*$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidSessionNameRule) Name() string {
	return "aws_quicksight_user_invalid_session_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidSessionNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidSessionNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidSessionNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidSessionNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"session_name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"session_name must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=.@-]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
