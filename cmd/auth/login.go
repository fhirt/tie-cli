/*
Copyright © 2025 fhirt
*/
package auth

import (
	"log"

	"github.com/fhirt/tie-cli/cmd/util"
	"github.com/fhirt/tie-cli/internal/web"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var (
	user       string
	password   string
	repository string
	loginCmd   = &cobra.Command{
		Use:   "login",
		Short: "A brief description of your command",
		Run:   runLogin,
	}
)

func runLogin(cmd *cobra.Command, args []string) {
	if viper.GetString(REPOSITORY_USER) == "" {
		user = util.Prompt(promptui.Prompt{
			Label: "User",
		})
		viper.Set(REPOSITORY_USER, user)
	}
	if viper.GetString(REPOSITORY_PASSWORD) == "" {
		password = util.Prompt(promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		})
		viper.Set(REPOSITORY_PASSWORD, password)
	}
	repositoryClient := web.NewRepositoryClient(viper.GetString(REPOSITORY_URL), viper.GetString(REPOSITORY_USER), viper.GetString(REPOSITORY_PASSWORD))
	err := repositoryClient.Login()
	if err != nil {
		log.Fatal(err)
	}
	viper.WriteConfig()

}

func init() {
	AuthCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "user name")
	AuthCmd.PersistentFlags().StringVarP(&repository, "repository", "r", "dev-repository.tie.ch", "repository url")
	viper.BindPFlag(REPOSITORY_USER, AuthCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag(REPOSITORY_URL, AuthCmd.PersistentFlags().Lookup("repository"))
}
