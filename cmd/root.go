package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "moldy",
	Short: "Moldy the best project starter of the world",
	Long: `
The best tool for initialize your project
For more information run the info command.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
