package cmd

import (
	"fmt"

	"github.com/cyclops-ui/cycops-cyctl/internal/client"
	"github.com/spf13/cobra"
)

var (
	updateExample = `# Update a single module
cyctl update module <module-name> --key=<key> --value=<value>
`
	updateKey   string
	updateValue string
)

// updateModules updates a specified module from the Cyclops API.
func updateModules(clientset *client.CyclopsV1Alpha1Client, moduleNames []string, key, value string) {
	if len(moduleNames) == 0 {
		fmt.Println("Error: module names cannot be empty")
		return
	}

	for _, moduleName := range moduleNames {
		err := clientset.Modules("cyclops").Update(moduleName, key, value)
		if err != nil {
			fmt.Printf("Error updating module %v: %v\n", moduleName, err)
		} else {
			fmt.Printf("Module %v updated successfully with %s=%s.\n", moduleName, key, value)
		}
	}
}

var updateModuleCMD = &cobra.Command{
	Use:     "module <module-name>",
	Short:   "Update a module",
	Long:    "Update a module in the Cyclops API with the specified key and value.",
	Example: updateExample,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clientset := client.NewCyclopsV1Alpha1Client()
		updateModules(clientset, args, updateKey, updateValue)
	},
}

func init() {
	updateModuleCMD.Flags().StringVar(&updateKey, "key", "", "Key to update in the module")
	updateModuleCMD.Flags().StringVar(&updateValue, "value", "", "Value to set for the specified key in the module")
	updateModuleCMD.MarkFlagRequired("key")
	updateModuleCMD.MarkFlagRequired("value")

	updateCMD.AddCommand(updateModuleCMD)
	RootCmd.AddCommand(updateCMD)
}
