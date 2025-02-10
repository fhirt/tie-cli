package util

import (
	"log"

	"github.com/manifoldco/promptui"
)

func Prompt(prompt promptui.Prompt) string {
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
