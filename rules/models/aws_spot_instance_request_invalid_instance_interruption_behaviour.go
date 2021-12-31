// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule checks the pattern is valid
type AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule returns new rule with default attributes
func NewAwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule() *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule {
	return &AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule{
		resourceType:  "aws_spot_instance_request",
		attributeName: "instance_interruption_behaviour",
		enum: []string{
			"hibernate",
			"stop",
			"terminate",
		},
	}
}

// Name returns the rule name
func (r *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule) Name() string {
	return "aws_spot_instance_request_invalid_instance_interruption_behaviour"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as instance_interruption_behaviour`, truncateLongMessage(val)),
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
