package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tinc",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := startFunc(args)
		if err != nil {
			fmt.Println(err)
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}

func startFunc(args []string) error {

	if len(args) < 4 {
		return fmt.Errorf("%s", "Incorrect number of arguments.")
	}
	return nil
}
