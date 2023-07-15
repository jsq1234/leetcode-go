// authCmd represents the auth command
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Manan-Prakash-Singh/leetcode-go/core"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate user using cookies",
	Long: `Authenticate the user using csrftoken and LEETCODE_SESSION.
You must create an environment variable to store cookie.
Do this using, "export LEETCODE_SESSION_KEY="csrftoken=asdasda.....;LEETCODE_SESSION=sadsda...."
and preferably put it inside your ~/.bashrc or ~/.zshrc. To obtain your csrftoken and LEETCODE_SESSION,
open a browser and go to Developer Tools. Click the Network Tab and browse your leetcode account
check the headers. You will find these terms in Cookie field of request headers.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.AuthenticateUser()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
