package main

import (
	"testing"

	"github.com/hanks/terraform-variables-generator/version"
)

func Test_createApp(t *testing.T) {
	tests := []struct {
		name        string
		wantName    string
		wantVersion string
	}{
		{
			name:        "Create app instance",
			wantName:    "terraform-variables-genrator",
			wantVersion: version.Version,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createApp()
			if got.Name != tt.wantName {
				t.Errorf("createApp() = %v, want %v", got.Name, tt.wantName)
			}
			if got.Version != tt.wantVersion {
				t.Errorf("createApp() = %v, want %v", got.Version, tt.wantVersion)
			}
		})
	}
}
