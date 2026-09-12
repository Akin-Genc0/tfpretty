package terraform

import (
	"encoding/json"
	"fmt"
)

func ParsePlan(data []byte) (Plan, error) {

	var plan Plan

	// Unmarshal the Terraform JSON data into the Plan struct
	err := json.Unmarshal(data, &plan)
	if err != nil {
		return Plan{}, fmt.Errorf("could not parse Terraform plan JSON: %w", err)
	}
	return plan, nil
}

// GetActionGroup determines which Terraform action is being used.

func GetActionGroup(actions []string) (Action, error) {

	if len(actions) == 0 {
		return ActionNoOp, nil
	}

	if len(actions) == 2 {
		if actions[0] == "create" && actions[1] == "delete" ||
			actions[0] == "delete" && actions[1] == "create" {
			return ActionReplace, nil
		}
	}

	if len(actions) == 1 {
		switch actions[0] {
		case "create":
			return ActionCreate, nil
		case "update":
			return ActionUpdate, nil
		case "delete":
			return ActionDelete, nil
		case "no-op":
			return ActionNoOp, nil
		default:
			return ActionNoOp, fmt.Errorf("unknown Terraform action: %s", actions[0])
		}
	}

	return ActionNoOp, fmt.Errorf("unsupported Terraform action group: %v", actions)
}
