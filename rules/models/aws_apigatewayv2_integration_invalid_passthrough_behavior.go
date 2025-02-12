// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule checks the pattern is valid
type AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule returns new rule with default attributes
func NewAwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule() *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule {
	return &AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule{
		resourceType:  "aws_apigatewayv2_integration",
		attributeName: "passthrough_behavior",
		enum: []string{
			"WHEN_NO_MATCH",
			"NEVER",
			"WHEN_NO_TEMPLATES",
		},
	}
}

// Name returns the rule name
func (r *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule) Name() string {
	return "aws_apigatewayv2_integration_invalid_passthrough_behavior"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsApigatewayv2IntegrationInvalidPassthroughBehaviorRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as passthrough_behavior`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
