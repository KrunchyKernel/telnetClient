package Server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/KrunchyKernel/telnetServer/models"
)

const (
	ConnHost = "localhost"
)

func StartServer(used *[]int, listeners *[]net.Listener) (err error) {
	fmt.Println("initializing server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req here")
		//var res []byte
		var resConf models.Config
		//_, err = r.Body.Read(res)
		//if err != nil {
		//	return
		//}
		err = json.NewDecoder(r.Body).Decode(&resConf)
		fmt.Println(resConf.Target)
		//err = json.Unmarshal(res, resConf)
		if err != nil {
			fmt.Println("some error shit")
			fmt.Println(err.Error())
			return
		}
		fmt.Println(resConf.Port)
		if Contains(*used, resConf.Port) {
			fmt.Println("No")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("Port already in use"))
			if err != nil {
				return
			}
		} else {
			err = StartTCPListen(listeners, resConf.Port)
		}

		*used = append(*(used), resConf.Port)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
	return err
}

func StartTCPListen(listeners *[]net.Listener, port int) error {
	addr := ConnHost + ":" + strconv.Itoa(port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	*listeners = append(*listeners, listen)
	return nil
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
