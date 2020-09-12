package cmd

import (
	"fmt"

	"github.com/Alex99y/duplicate-files/pkg/utils"

	"github.com/spf13/cobra"
)

// CobraInterface represents the CMD interface
type CobraInterface struct {
	RootCmd    *cobra.Command
	RootFolder string
}

func (cmd *CobraInterface) setRootCommand() {
	cmd.RootCmd = &cobra.Command{
		Short: "Application to search duplicate files inside a folder",
	}
}

func (cmd *CobraInterface) setVersion() {
	version := &cobra.Command{
		Use:   "version",
		Short: "Print app version",
		Run: func(c *cobra.Command, arg []string) {
			fmt.Print("v0.1.0")
		},
	}
	cmd.RootCmd.AddCommand(version)
}

func (cmd *CobraInterface) setStart() {
	start := &cobra.Command{
		Use:     "start",
		Short:   "Search for duplicated files",
		Example: "start [rootFolder]",
		Long:    "This command receives a folder, find recursively and print all duplicate files inside this folder and his subfolderss",
		Run: func(c *cobra.Command, args []string) {
			cmd.RootFolder = args[0]
		},
		Args: func(c *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("No root folder provided")
			}
			isDir, err := utils.IsDirectory(args[0])
			if err != nil || !isDir {
				return fmt.Errorf("Invalid root folder provided")
			}
			return nil
		},
	}

	cmd.RootCmd.AddCommand(start)
}

// Execute function starts reading arguments from CLI
func (cmd *CobraInterface) Execute() {
	cmd.setRootCommand()
	cmd.setVersion()
	cmd.setStart()
	cmd.RootCmd.Execute()
}
