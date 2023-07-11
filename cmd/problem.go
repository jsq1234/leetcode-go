/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// problemCmd represents the problem command
var (
    problemCmd = &cobra.Command{
        Use:   "problem",
        Short: "Download the problem in the current directory",
        Long: ``,
        Args: cobra.MaximumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            if !download && lang == "" && len(args) == 0 {
               fmt.Fprintf(os.Stderr,"Expected 1 args, 0 given.\n")
               cmd.Help()
               os.Exit(1)
            }

        },
    }
    download bool
    lang string
)

func init() {
	rootCmd.AddCommand(problemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// problemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// problemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    problemCmd.Flags().BoolVarP(&download, "download", "d", false, "Download the problem in the current directory")
    problemCmd.Flags().StringVarP(&lang, "lang", "l", "", "Prog. Language to download code")
    problemCmd.MarkFlagsRequiredTogether("download","lang")

}
