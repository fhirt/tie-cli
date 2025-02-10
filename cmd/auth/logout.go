/*
Copyright © 2025 fhirt
*/
package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from the repository",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(REPOSITORY_PASSWORD, "")
		viper.Set(REPOSITORY_USER, "")
		viper.Set(REPOSITORY_URL, "")
		viper.WriteConfig()
		fmt.Println("logged out")
	},
}
