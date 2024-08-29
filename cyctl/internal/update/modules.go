package update

import (
	"fmt"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1/client"
	"github.com/cyclops-ui/cycops-cyctl/internal/kubeconfig"
	"github.com/spf13/cobra"
) 

var (
	updateModuleExample = `# Update a single module
cyctl update module <module-name> --key=<key> --value=<value>
`
	valuesFile   string
	templateName string
	updateKey    string
	updateValue  string
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
			fmt.Printf("Error from server (NotFound): %v\n", err)
		} else {
			fmt.Printf("Module %v updated successfully.\n", moduleName, key, value)
		}
	}
}

var (
	UpdateModule = &cobra.Command{
		Use:     "modules [module_name]",
		Short:   "update module",
		Long:    "The Update modules command allows you to update already deployed module from the Cyclops API.",
		Example: updateModuleExample,
		Aliases: []string{"module"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			updateModules(kubeconfig.Moduleset, args, updateKey, updateValue )
		},
	}
)

func init() {
	UpdateModule.Flags().StringVar(&updateKey, "key", "", "Key to update in the module")
	UpdateModule.Flags().StringVar(&updateValue, "value", "", "Value to set for the specified key in the module")
	UpdateModule.MarkFlagRequired("key")
	UpdateModule.MarkFlagRequired("value")
}

func main() {
	rootCmd := &cobra.Command{Use: "cyctl"}
	rootCmd.AddCommand(UpdateModule)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}