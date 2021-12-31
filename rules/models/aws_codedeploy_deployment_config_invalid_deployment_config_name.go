// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule checks the pattern is valid
type AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule returns new rule with default attributes
func NewAwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule() *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule {
	return &AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule{
		resourceType:  "aws_codedeploy_deployment_config",
		attributeName: "deployment_config_name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule) Name() string {
	return "aws_codedeploy_deployment_config_invalid_deployment_config_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule) Check(runner tflint.Runner) error {
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
					"deployment_config_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"deployment_config_name must be 1 characters or higher",
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
