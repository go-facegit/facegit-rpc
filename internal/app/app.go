package app

import (
	"fmt"
	"net/http"
	"net/rpc"
	// "os"
	// "path/filepath"
)

func Start(port int) {

	sPortStr := fmt.Sprintf(":%d", port)

	fmt.Println(sPortStr)

	rpc.HandleHTTP()
	err := http.ListenAndServe(sPortStr, nil)
	if err != nil {
		fmt.Println("rpc listen err:%s", err.Error())
	}
}
