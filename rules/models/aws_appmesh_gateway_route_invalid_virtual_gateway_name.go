// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule checks the pattern is valid
type AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule returns new rule with default attributes
func NewAwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule() *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule {
	return &AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule{
		resourceType:  "aws_appmesh_gateway_route",
		attributeName: "virtual_gateway_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule) Name() string {
	return "aws_appmesh_gateway_route_invalid_virtual_gateway_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshGatewayRouteInvalidVirtualGatewayNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"virtual_gateway_name must be 255 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"virtual_gateway_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
