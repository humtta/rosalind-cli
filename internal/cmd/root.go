package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rosalind",
		Short: "Access Rosalind problems from the terminal",
	}

	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newVersionCmd())

	return cmd
}

func Execute() error {
	return newRootCmd().Execute()
}
