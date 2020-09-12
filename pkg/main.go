package main

import (
	"github.com/Alex99y/duplicate-files/pkg/cmd"
	"github.com/Alex99y/duplicate-files/pkg/core"
)

func main() {
	cobra := new(cmd.CobraInterface)
	cobra.Execute()
	if cobra.RootFolder == "" {
		return
	}
	core.Start(*cobra)
}
