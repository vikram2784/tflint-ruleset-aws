// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule checks the pattern is valid
type AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule returns new rule with default attributes
func NewAwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule() *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule {
	return &AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule{
		resourceType:  "aws_launch_template",
		attributeName: "instance_initiated_shutdown_behavior",
		enum: []string{
			"stop",
			"terminate",
		},
	}
}

// Name returns the rule name
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Name() string {
	return "aws_launch_template_invalid_instance_initiated_shutdown_behavior"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as instance_initiated_shutdown_behavior`, truncateLongMessage(val)),
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
