package cmd

import (
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Rosalind problems",
	}

	return cmd
}
