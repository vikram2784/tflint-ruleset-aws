// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule checks the pattern is valid
type AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule returns new rule with default attributes
func NewAwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule() *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule {
	return &AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule{
		resourceType:  "aws_globalaccelerator_accelerator",
		attributeName: "ip_address_type",
		enum: []string{
			"IPV4",
		},
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule) Name() string {
	return "aws_globalaccelerator_accelerator_invalid_ip_address_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as ip_address_type`, truncateLongMessage(val)),
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
