// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferServerInvalidIdentityProviderTypeRule checks the pattern is valid
type AwsTransferServerInvalidIdentityProviderTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsTransferServerInvalidIdentityProviderTypeRule returns new rule with default attributes
func NewAwsTransferServerInvalidIdentityProviderTypeRule() *AwsTransferServerInvalidIdentityProviderTypeRule {
	return &AwsTransferServerInvalidIdentityProviderTypeRule{
		resourceType:  "aws_transfer_server",
		attributeName: "identity_provider_type",
		enum: []string{
			"SERVICE_MANAGED",
			"API_GATEWAY",
			"AWS_DIRECTORY_SERVICE",
			"AWS_LAMBDA",
		},
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Name() string {
	return "aws_transfer_server_invalid_identity_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as identity_provider_type`, truncateLongMessage(val)),
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
