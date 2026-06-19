package cmd

import (
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Show the statement of a Rosalind problem",
	}

	return cmd
}
