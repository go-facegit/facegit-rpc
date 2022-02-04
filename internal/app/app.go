package app

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/go-facegit/facegit-rpc/internal/app/lib"
)

func Start(port int) {

	rpc.HandleHTTP()
	repoLib := new(lib.Repo)
	rpc.Register(repoLib)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("rpc listen err: " + err.Error())
	}
}
