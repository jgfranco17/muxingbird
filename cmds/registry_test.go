package cmds

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCommandRegistry_SetsDefaultFlags(t *testing.T) {
	reg := NewCommandRegistry("testcli", "Test CLI", "0.1.0")
	assert.Equal(t, "testcli", reg.rootCmd.Root().Use)
	assert.Equal(t, "Test CLI", reg.rootCmd.Root().Short)
	assert.Equal(t, "0.1.0", reg.rootCmd.Root().Version)
	flag := reg.rootCmd.Root().PersistentFlags().Lookup("verbose")
	require.NotNil(t, flag)
	assert.Equal(t, "v", flag.Shorthand)
}

func TestRegisterCommands_RegistersSubcommands(t *testing.T) {
	reg := NewCommandRegistry("testcli", "Test CLI", "0.1.0")

	var called bool
	mockCmd := &cobra.Command{
		Use:   "mock",
		Short: "A mock command",
		Run: func(cmd *cobra.Command, args []string) {
			called = true
		},
	}
	reg.RegisterCommands([]*cobra.Command{mockCmd})

	reg.rootCmd.Root().SetArgs([]string{"mock"})
	err := reg.Execute()
	require.NoError(t, err)
	assert.True(t, called)
}
