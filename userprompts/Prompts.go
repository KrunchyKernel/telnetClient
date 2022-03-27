package userprompts

import (
	"errors"
	prompt "github.com/manifoldco/promptui"
	"strings"
)

func configPromptValid(input string) error {
	if !strings.ContainsAny(input, "ynYN") {
		return errors.New("invalid input")
	}
	return nil
}

var Confprompt = prompt.Prompt{
	Label:       "Use config file?",
	Default:     "yes",
	AllowEdit:   false,
	Validate:    configPromptValid,
	Mask:        0,
	HideEntered: false,
	Templates:   nil,
	IsConfirm:   true,
	IsVimMode:   false,
	Pointer:     nil,
	Stdin:       nil,
	Stdout:      nil,
}
