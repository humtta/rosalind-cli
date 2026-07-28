// Package problem defines types and helpers for Rosalind problems.
package problem

import (
	"fmt"

	"github.com/humtta/rosalind-cli/internal/utils"
)

// Problem represents a Rosalind problem.
type Problem struct {
	ID        string           `json:"id"`
	Title     string           `json:"title"`
	URL       string           `json:"url"`
	Statement ProblemStatement `json:"-"`
}

// ProblemStatement represents the statement of a Rosalind problem.
type ProblemStatement struct {
	Description   string `json:"description"`
	SampleDataset string `json:"sample_dataset"`
	SampleOutput  string `json:"sample_output"`
}

// ValidateID validates the given problem ID.
func ValidateID(id string) error {
	if id == "" {
		return fmt.Errorf("must not be empty")
	}
	if !utils.IsAlphaASCII(id) {
		return fmt.Errorf("must be an ASCII alphabetic string")
	}
	return nil
}
