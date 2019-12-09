package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

//here we set the main command "devoir" to be used with subcommands
var devoir = &cobra.Command{
	Use:   "devoir",
	Short: "commande de test",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

//here we set the subcommand "un" to print un with devoir as the main command
var un = &cobra.Command{
	Use:   "un",
	Short: "Commande pour afficher un",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("un")
	},
}

//here we set the subcommand "deux" to print deux with devoir as the main command
var deux = &cobra.Command{
	Use:   "deux",
	Short: "Commande pour afficher deux",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deux")
	},
}

//Here we add the 2 commands "un" et "deux" before executing the program
func main() {
	devoir.AddCommand(un)
	devoir.AddCommand(deux)
	devoir.Execute()
}
