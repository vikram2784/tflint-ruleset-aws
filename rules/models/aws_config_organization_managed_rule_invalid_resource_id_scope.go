// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule() *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule {
	return &AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "resource_id_scope",
		max:           768,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_resource_id_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule) Check(runner tflint.Runner) error {
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
					"resource_id_scope must be 768 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_id_scope must be 1 characters or higher",
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
