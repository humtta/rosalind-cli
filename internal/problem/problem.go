package problem

import (
	"fmt"

	"github.com/humtta/rosalind-cli/internal/utils"
)

type Problem struct {
	Index     int
	ID        string
	Title     string
	URL       string
	Statement ProblemStatement
}

type ProblemStatement struct {
	Description   string
	SampleDataset string
	SampleOutput  string
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
