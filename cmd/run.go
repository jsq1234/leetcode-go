/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your code with the test cases in your program file.",
	Long: `Run your code against the test cases that are in your program file. Make sure the test cases
    are the very first thing enclosed in /* */ comment block. Multiple test cases are seperated by ";"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		core.RunCode(args[0])
	},
}

var fileName string

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	runCmd.Flags().StringVarP(&fileName, "file", "f", "", "Path of testcase file")
}
