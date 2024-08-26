/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// prioritiesCmd represents the priorities command
var prioritiesCmd = &cobra.Command{
	Use:   "priorities",
	Short: "List all your current priorities",
	Long:  `This command allows you to view all the priorities you have added. You can use it to get an overview of your tasks and goals.`,
	Run: func(cmd *cobra.Command, args []string) {
		priorities := []string{
			"Complete Go project",
			"Review Pull Requests",
			"Write the documentation",
		}
		fmt.Println("priorities called")
		for i, priority := range priorities {
			fmt.Printf("%d. %s\n", i+1, priority)
		}

	},
}

func init() {
	rootCmd.AddCommand(prioritiesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//prioritiesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//prioritiesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
