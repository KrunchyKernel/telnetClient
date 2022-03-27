package UserFuncs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/KrunchyKernel/telnetClient/models"
)

func ReadConfigFile() (userConfig *models.Config, err error) {
	var dir string
	var confFile string
	var confBytes []byte

	dir, err = os.Getwd()
	if err != nil {
		return nil, errors.New("error reading config")
	}
	confFile = dir + "/config.json"

	confBytes, err = ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Println("file error: ", err.Error())
	}
	err = json.Unmarshal(confBytes, userConfig)

	if err != nil {
		return nil, err
	}
	return userConfig, err
}
