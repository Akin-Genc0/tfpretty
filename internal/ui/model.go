package ui

import (
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

type Screen int

const (
	PlanScreen Screen = iota
	DetailScreen
)

type Model struct {
	Plan   terraform.Plan
	Cursor int
	Screen Screen
}
