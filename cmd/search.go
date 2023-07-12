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
        val := 10
        if count != 0 {
            val = 5 
        }
        core.SearchProblem(args[0],val)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    searchCmd.Flags().IntVarP(&count,"num","n",10,"Number of problems to return on search")
}
