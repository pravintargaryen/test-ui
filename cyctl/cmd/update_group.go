package cmd

import (
	"github.com/spf13/cobra"
)

var (
	updateExample = `# Update a module
cyctl update module <module-name> --key=<key> --value=<value>`
)

var updateCMD = &cobra.Command{
	Use:     "update",
	Short:   "Update custom resources like modules",
	Long:    "Update custom resources like modules",
	Example: updateExample,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.AddCommand(updateCMD)
}
