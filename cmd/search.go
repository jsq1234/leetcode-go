/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var (
	count int
	plang string
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search a problem",
	Long: `Search a problem from leetcode. Either enclose the problem in double quotes in case of 
multiple words like "palindrome partionining" or use a dash instead palindrome-partitioning`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		core.SearchProblem(args[0], count)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().IntVarP(&count, "num", "n", 50, "Number of problems to return on search")
}
