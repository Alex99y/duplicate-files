package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraInterface represents the CMD interface
type CobraInterface struct {
	RootCmd         *cobra.Command
	NumberOfThreads uint64
	RootFolder      string
}

func (cmd *CobraInterface) setRootCommand() {
	cmd.RootCmd = &cobra.Command{
		Short: "Short",
		Long:  "Long",
	}
}

func (cmd *CobraInterface) setVersion() {
	version := &cobra.Command{
		Use:   "version",
		Short: "Print app version",
		Run: func(c *cobra.Command, arg []string) {
			fmt.Print("v0.0.1")
		},
	}
	cmd.RootCmd.AddCommand(version)
}

func (cmd *CobraInterface) setStart() {
	start := &cobra.Command{
		Use:   "start",
		Short: "Execute duplicate files searcher",
		Long:  "Long description",
		Run: func(c *cobra.Command, arg []string) {
			cmd.RootFolder, _ = c.PersistentFlags().GetString("path")
			cmd.NumberOfThreads, _ = c.PersistentFlags().GetUint64("threads")
		},
	}
	start.PersistentFlags().Uint64P("threads", "t", 4, "--threads 2")
	start.PersistentFlags().StringP("path", "f", "", "--path /home")
	start.MarkPersistentFlagRequired("path")

	cmd.RootCmd.AddCommand(start)
}

// Execute function starts reading arguments from CLI
func (cmd *CobraInterface) Execute() {
	cmd.setRootCommand()
	cmd.setVersion()
	cmd.setStart()
	cmd.RootCmd.Execute()
}
