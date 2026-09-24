package terraform

import (
	"strings"
	"testing"
)

func TestRunPlanReturnsErrorWhenTerraformIsUnavailable(t *testing.T) {
	t.Setenv("PATH", "")

	output, err := RunPlan()
	if err == nil {
		t.Fatal("expected RunPlan to return an error when Terraform is unavailable, but got nil")
	}

	if output != nil {
		t.Errorf("expected RunPlan output to be nil when Terraform is unavailable, got %q", output)
	}

	if !strings.Contains(err.Error(), "terraform plan failed") {
		t.Errorf("expected RunPlan error to contain %q, got %q", "terraform plan failed", err)
	}
}
