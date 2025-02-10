/*
Copyright © 2025 fhirt
*/
package auth

import (
	"github.com/spf13/cobra"
)

var (
	AuthCmd = &cobra.Command{
		Use:   "auth",
		Short: "login, logout and status commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

const (
	REPOSITORY_USER     = "repository.user"
	REPOSITORY_URL      = "repository.url"
	REPOSITORY_PASSWORD = "repository.password"
)

func init() {
	AuthCmd.AddCommand(statusCmd)
	AuthCmd.AddCommand(loginCmd)
	AuthCmd.AddCommand(logoutCmd)
}
