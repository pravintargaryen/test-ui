package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cyclops-ui/cycops-cyctl/internal/client"
	"github.com/spf13/cobra"
)

// Mock client to simulate Cyclops API behavior
type MockClient struct {
	client.CyclopsV1Alpha1Client
}

func (m *MockClient) Modules(namespace string) client.ModuleInterface {
	return &MockModuleClient{}
}

type MockModuleClient struct{}

func (m *MockModuleClient) Update(moduleName, key, value string) error {
	// Simulate updating a module
	if moduleName == "fail" {
		return fmt.Errorf("failed to update module")
	}
	return nil
}

func TestUpdateModules(t *testing.T) {
	clientset := &MockClient{}
	moduleNames := []string{"test-module"}
	key := "test-key"
	value := "test-value"

	updateModules(clientset, moduleNames, key, value)
	// Add assertions here as needed
}

func TestUpdateModuleCommand(t *testing.T) {
	cmd := &cobra.Command{}
	b := new(bytes.Buffer)
	cmd.SetOut(b)

	updateModuleCMD.Run(cmd, []string{"test-module"})

	output := b.String()
	expectedOutput := "Module test-module updated successfully with test-key=test-value.\n"
	if output != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output)
	}
}
