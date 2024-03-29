package cmd

import (
	"github.com/navneetshukl/create/helpers"
	"github.com/spf13/cobra"
)

// fileCmd is used to create the file
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Used to create a file with specific name and extension",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Length of argument is ", len(args))
		for _, val := range args {
			err := helpers.CreateFile(val)
			if err != nil {
				continue
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(fileCmd)
}
