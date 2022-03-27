package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KrunchyKernel/telnetClient/UserFuncs"
	"github.com/KrunchyKernel/telnetClient/models"
	"github.com/KrunchyKernel/telnetClient/userprompts"
)

func main() {
	var err error
	var file *models.Config
	fmt.Println(os.Args)
	fmt.Println(models.Welcome)
	c, err := userprompts.Confprompt.Run()
	if strings.ToLower(c) == "y" {
		file, err = UserFuncs.ReadConfigFile()
		if err != nil {
			return
		}
	} else {
		//add prompts to manually get config info
	}
	UserFuncs.InitConnection(*file)
	fmt.Println(err)

}
