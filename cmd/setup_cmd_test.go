package cmd

import (
	"bytes"
	"testing"
)

type MockSetup struct{}

func NewMockSetup() Setup {
	return &MockSetup{}
}

func (s *MockSetup) Setup(licensProperties, quickstartPath, setupPath, repoPath, instance string) error {
	return nil
}

func TestSetupCommand(t *testing.T) {
	t.Run("Should throw error if license.properties file is not passed", func(t *testing.T) {
		actual := new(bytes.Buffer)
		setup := NewMockSetup()
		setupCmd := NewSetupCommand(setup)
		setupCmd.SetOut(actual)
		setupCmd.SetErr(actual)
		setupCmd.SetArgs([]string{})
		setupCmd.Execute()

		expected := "Path to license.properties file is required. Exiting...\n"

		AssertOutput(t, actual.String(), expected)
	})

	t.Run("Should throw error if path to quickstart JAR is not passed", func(t *testing.T) {
		actual := new(bytes.Buffer)
		setup := NewMockSetup()
		setupCmd := NewSetupCommand(setup)
		setupCmd.SetOut(actual)
		setupCmd.SetErr(actual)
		setupCmd.SetArgs([]string{"./license.properties"})
		setupCmd.Execute()

		expected := "Path to quickstart JAR is required. Exiting...\n"

		AssertOutput(t, actual.String(), expected)
	})

	t.Run("Setup of local AEM environment is successful", func(t *testing.T) {
		actual := new(bytes.Buffer)
		setup := NewMockSetup()
		setupCmd := NewSetupCommand(setup)
		setupCmd.SetOut(actual)
		setupCmd.SetErr(actual)
		setupCmd.SetArgs([]string{"./license.properties", "./cq-quickstart.jar"})
		setupCmd.Execute()

		expected := "Setup of local AEM environment complete.\n"

		AssertOutput(t, actual.String(), expected)
	})
}

func AssertOutput(t testing.TB, actual, expected string) {
	if actual != expected {
		t.Errorf("Failed: Actual - %q, Expected - %q", actual, expected)
	}
}
