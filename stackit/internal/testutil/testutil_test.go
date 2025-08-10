package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-testing/config"
)

func TestProviderConfigBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *ProviderConfigBuilder
		want    string
	}{
		{
			name:    "basic",
			builder: newProviderConfigBuilder(),
			want: `provider "stackit" {
	default_region = "eu01"
}`,
		},
		{
			name:    "set default region",
			builder: newProviderConfigBuilder().DefaultRegion("eu02"),
			want: `provider "stackit" {
	default_region = "eu02"
}`,
		},
		{
			name:    "enable beta resources",
			builder: newProviderConfigBuilder().EnableBetaResources(),
			want: `provider "stackit" {
	default_region = "eu01"
	enable_beta_resources = true
}`,
		},
		{
			name:    "enable single experiment",
			builder: newProviderConfigBuilder().EnableExperiment("iaas"),
			want: `provider "stackit" {
	default_region = "eu01"
	experiments = [ "iaas" ]
}`,
		},
		{
			name:    "enable multiple experiments",
			builder: newProviderConfigBuilder().EnableExperiment("iaas").EnableExperiment("resourcemanager"),
			want: `provider "stackit" {
	default_region = "eu01"
	experiments = [ "iaas", "resourcemanager" ]
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := cmp.Diff(tt.builder.Build(), tt.want)
			if diff != "" {
				t.Fatalf("Provider config does not match: %s", diff)
			}
		})
	}
}

func TestConvertConfigVariable(t *testing.T) {
	tests := []struct {
		name     string
		variable config.Variable
		want     string
	}{
		{
			name:     "string",
			variable: config.StringVariable("test"),
			want:     "test",
		},
		{
			name:     "bool: true",
			variable: config.BoolVariable(true),
			want:     "true",
		},
		{
			name:     "bool: false",
			variable: config.BoolVariable(false),
			want:     "false",
		},
		{
			name:     "integer",
			variable: config.IntegerVariable(10),
			want:     "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertConfigVariable(tt.variable); got != tt.want {
				t.Errorf("ConvertConfigVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}
