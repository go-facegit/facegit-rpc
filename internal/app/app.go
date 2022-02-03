package app

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/go-facegit/facegit-rpc/internal/app/lib"
)

func Start(port int) {

	sPortStr := fmt.Sprintf(":%d", port)

	rpc.HandleHTTP()

	repoLib := new(lib.Repo)

	rpc.Register(repoLib)
	err := http.ListenAndServe(sPortStr, nil)
	if err != nil {
		fmt.Println("rpc listen err:%s", err.Error())
	}
}
