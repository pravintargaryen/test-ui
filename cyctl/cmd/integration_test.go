package cmd

import (
	"bytes"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = Describe("Update Module Command", func() {
	It("should update a module successfully", func() {
		cmd := exec.Command("cyctl", "update", "module", "test-module", "--key=test-key", "--value=test-value")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		Expect(err).NotTo(HaveOccurred())
		Expect(out.String()).To(ContainSubstring("Module test-module updated successfully"))
	})

	It("should fail to update a nonexistent module", func() {
		cmd := exec.Command("cyctl", "update", "module", "nonexistent-module", "--key=test-key", "--value=test-value")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		Expect(err).To(HaveOccurred())
		Expect(out.String()).To(ContainSubstring("Error updating module"))
	})
})
