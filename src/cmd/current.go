/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		topk, _ := cmd.Flags().GetInt("topk")
		fmt.Printf("%v\n", topk)
	},
}

func init() {
	publishCmd.AddCommand(currentCmd)

	currentCmd.Flags().IntP("topk", "k", 10, "how many articles to include")
}
