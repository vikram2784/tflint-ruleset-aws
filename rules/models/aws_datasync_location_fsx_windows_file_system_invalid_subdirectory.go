// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule checks the pattern is valid
type AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule returns new rule with default attributes
func NewAwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule() *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule {
	return &AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule{
		resourceType:  "aws_datasync_location_fsx_windows_file_system",
		attributeName: "subdirectory",
		max:           4096,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule) Name() string {
	return "aws_datasync_location_fsx_windows_file_system_invalid_subdirectory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidSubdirectoryRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"subdirectory must be 4096 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
