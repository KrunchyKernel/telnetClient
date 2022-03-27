package UserFuncs

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/KrunchyKernel/telnetClient/models"
)

func InitConnection(conf models.Config) (conn *net.Conn, err error) {
	var c *net.Conn
	var res *http.Response
	fmt.Println(conf.Target)
	res, err = http.Post(conf.Target, "application/json", conf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("intialization failed")
	}
	*(c), err = net.Dial("tcp", "localhost:8000")
	return c, err
}
