package lib

import (
	"fmt"
	"net/http"
	"net/rpc"
	"testing"

	"github.com/go-facegit/facegit-rpc/internal/app/router"
)

var testPort = ":8068"

func localServer() {
	rpc.HandleHTTP()

	repoLib := new(Repo)

	rpc.Register(repoLib)
	err := http.ListenAndServe(testPort, nil)
	if err != nil {
		fmt.Println("rpc listen err: " + err.Error())
	}
}

func init() {
	// go localServer()

	router.Init("")
	go localServer()
}

func getClient() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", testPort)
	return client, err
}

func TestRepoCreate(t *testing.T) {

	client, err := getClient()
	if err != nil {
		t.Errorf("TestRepoCreate %s", err)
	}

	var reply bool
	err = client.Call("Repo.Create", RepoCreateArgs{"midoks", "facegit"}, &reply)

	fmt.Println("ret", reply, err)
	if err != nil {
		t.Errorf("TestRepoCreate %s", err)
		return
	}
	t.Log("TestRepoCreate ok")
}
