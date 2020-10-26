// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package builder

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Import              *string           `mapstructure:"import" cty:"import" hcl:"import"`
	Clone               *string           `mapstructure:"clone" cty:"clone" hcl:"clone"`
	Suite               *string           `mapstructure:"suite" cty:"suite" hcl:"suite"`
	Mirror              *string           `mapstructure:"mirror" cty:"mirror" hcl:"mirror"`
	CacheDir            *string           `mapstructure:"cache_dir" cty:"cache_dir" hcl:"cache_dir"`
	MachinesDir         *string           `mapstructure:"machines_dir" cty:"machines_dir" hcl:"machines_dir"`
	Variant             *string           `mapstructure:"variant" cty:"variant" hcl:"variant"`
	Timeout             *string           `mapstructure:"timeout" cty:"timeout" hcl:"timeout"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"import":                     &hcldec.AttrSpec{Name: "import", Type: cty.String, Required: false},
		"clone":                      &hcldec.AttrSpec{Name: "clone", Type: cty.String, Required: false},
		"suite":                      &hcldec.AttrSpec{Name: "suite", Type: cty.String, Required: false},
		"mirror":                     &hcldec.AttrSpec{Name: "mirror", Type: cty.String, Required: false},
		"cache_dir":                  &hcldec.AttrSpec{Name: "cache_dir", Type: cty.String, Required: false},
		"machines_dir":               &hcldec.AttrSpec{Name: "machines_dir", Type: cty.String, Required: false},
		"variant":                    &hcldec.AttrSpec{Name: "variant", Type: cty.String, Required: false},
		"timeout":                    &hcldec.AttrSpec{Name: "timeout", Type: cty.String, Required: false},
	}
	return s
}
