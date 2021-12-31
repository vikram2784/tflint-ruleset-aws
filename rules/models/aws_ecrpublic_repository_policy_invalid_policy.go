// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrpublicRepositoryPolicyInvalidPolicyRule checks the pattern is valid
type AwsEcrpublicRepositoryPolicyInvalidPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsEcrpublicRepositoryPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsEcrpublicRepositoryPolicyInvalidPolicyRule() *AwsEcrpublicRepositoryPolicyInvalidPolicyRule {
	return &AwsEcrpublicRepositoryPolicyInvalidPolicyRule{
		resourceType:  "aws_ecrpublic_repository_policy",
		attributeName: "policy",
		max:           10240,
	}
}

// Name returns the rule name
func (r *AwsEcrpublicRepositoryPolicyInvalidPolicyRule) Name() string {
	return "aws_ecrpublic_repository_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrpublicRepositoryPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrpublicRepositoryPolicyInvalidPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrpublicRepositoryPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrpublicRepositoryPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
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
					"policy must be 10240 characters or less",
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
