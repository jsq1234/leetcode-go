/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

// problemCmd represents the problem command
var (
	problemCmd = &cobra.Command{
		Use:   "problem",
		Short: "Download the problem in the current directory",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			core.DownloadProblem(args[0])
		},
	}
	lang string
)

func init() {
	rootCmd.AddCommand(problemCmd)
}
