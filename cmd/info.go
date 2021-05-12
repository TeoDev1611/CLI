package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var asciiMoldy = `
 _     ____  _    _______  _
/ \__//  _ \/ \  /  _ \  \//
| |\/|| / \|| |  | | \|\  / 
| |  || \_/|| |_/\ |_/|/ /  
\_/  \\____/\____|____/_/   
						   
`
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show all information about Moldy CLI",
	Long:  `Show all information of Moldy CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(asciiMoldy + `
The best tool for initial his project :

Moldy is a Project Starter writed in Go that help for 
start his project and is 100% OPEN SOURCE.

Author: Moldy Community
Contact mail: moldycommunity@gmail.com
Repository: www.github.com/Moldy/Cli
Web Page: www.moldy.github.io

-----------------------------------------------------
Made with love in Ecuador.
`)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
