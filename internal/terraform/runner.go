package terraform

import (
	"fmt"
	"os"
	"os/exec"
)

func RunPlan() ([]byte, error) {

	tfcmdplan := exec.Command("terraform", "plan", "-out=tfpretty.tfplan")

	planFile := "tfpretty.tfplan"
	defer os.Remove(planFile)

	planOutput, err := tfcmdplan.CombinedOutput()
	if err != nil {
		return planOutput, fmt.Errorf("terraform plan failed: %w", err)
	}

	tfcmdshow := exec.Command("terraform", "show", "-json", "tfpretty.tfplan")

	jsonOutput, err := tfcmdshow.CombinedOutput()
	if err != nil {
		return jsonOutput, fmt.Errorf("terraform show failed: %w", err)
	}

	return jsonOutput, nil
}
