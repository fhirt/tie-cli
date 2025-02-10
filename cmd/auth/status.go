/*
Copyright © 2025 fhirt
*/
package auth

import (
	"fmt"

	"github.com/fhirt/tie-cli/internal/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show login status",
	Run: func(cmd *cobra.Command, args []string) {
		err := web.NewRepositoryClient(viper.GetString(REPOSITORY_URL), viper.GetString(REPOSITORY_USER), viper.GetString(REPOSITORY_PASSWORD)).Login()
		if err != nil {
			fmt.Println("not logged in")
		} else {
			fmt.Println("logged in")
		}
	},
}
