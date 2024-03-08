package cmd

import (
	"github.com/navneetshukl/create/helpers"
	"github.com/spf13/cobra"
)

// folderCmd command is used to create the folder
var folderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Use to create a folder of given name.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			err := helpers.CreateFolder(name)
			if err != nil {
				continue
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(folderCmd)
}
