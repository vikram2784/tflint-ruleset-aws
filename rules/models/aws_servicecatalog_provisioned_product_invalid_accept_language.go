// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule checks the pattern is valid
type AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogProvisionedProductInvalidAcceptLanguageRule returns new rule with default attributes
func NewAwsServicecatalogProvisionedProductInvalidAcceptLanguageRule() *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule {
	return &AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule{
		resourceType:  "aws_servicecatalog_provisioned_product",
		attributeName: "accept_language",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule) Name() string {
	return "aws_servicecatalog_provisioned_product_invalid_accept_language"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProvisionedProductInvalidAcceptLanguageRule) Check(runner tflint.Runner) error {
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
					"accept_language must be 100 characters or less",
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
