/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var (
	topicCmd = &cobra.Command{
		Use:   "topic",
		Short: "Get a list of problems of a particular topic",
		Long:  `Get a list of problems of a particular topic`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			core.GetTopics(args[0])
		},
	}
	hard   bool
	medium bool
	easy   bool
)

func init() {

	rootCmd.AddCommand(topicCmd)
	/*
	   TO-DO
	   topicCmd.Flags().BoolVarP(&hard, "hard", "", false, "Sort by hard")
	   topicCmd.Flags().BoolVarP(&medium, "medium", "", false, "Sort by medium")
	   topicCmd.Flags().BoolVarP(&easy, "easy", "", false, "Sort by easy")
	*/
}
