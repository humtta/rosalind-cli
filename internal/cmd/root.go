package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{}

	return cmd
}

func Execute() error {
	return newRootCmd().Execute()
}
