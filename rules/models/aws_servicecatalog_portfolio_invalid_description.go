// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPortfolioInvalidDescriptionRule checks the pattern is valid
type AwsServicecatalogPortfolioInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogPortfolioInvalidDescriptionRule returns new rule with default attributes
func NewAwsServicecatalogPortfolioInvalidDescriptionRule() *AwsServicecatalogPortfolioInvalidDescriptionRule {
	return &AwsServicecatalogPortfolioInvalidDescriptionRule{
		resourceType:  "aws_servicecatalog_portfolio",
		attributeName: "description",
		max:           2000,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Name() string {
	return "aws_servicecatalog_portfolio_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 2000 characters or less",
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
