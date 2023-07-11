/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

// qodCmd represents the qod command
var qodCmd = &cobra.Command{
	Use:   "qod",
	Short: "Get problem of the day",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        core.ProblemOfTheDay()
	},
}

func init() {
	rootCmd.AddCommand(qodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
