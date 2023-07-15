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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	topicCmd.Flags().BoolVarP(&hard, "hard", "", false, "Sort by hard")
	topicCmd.Flags().BoolVarP(&medium, "medium", "", false, "Sort by medium")
	topicCmd.Flags().BoolVarP(&easy, "easy", "", false, "Sort by easy")
}
