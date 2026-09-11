package terraform

import (
	"encoding/json"
	"fmt"
)

func ParsePlan(data []byte) (Plan, error) {

	var plan Plan

	err := json.Unmarshal(data, &plan)
	if err != nil {
		return Plan{}, fmt.Errorf("could not parse Terraform plan JSON: %w", err)
	}
	return plan, nil
}
