package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newGetCmd())

	return cmd
}

func Execute() error {
	return newRootCmd().Execute()
}
