package lib

import (
	"fmt"
	"net/http"
	"net/rpc"
	"testing"
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
	go localServer()
}

func getClient() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", testPort)
	return client, err
}

type ArgsTwo struct {
	UserName    string
	ProjectName string
}

func TestRepoCreate(t *testing.T) {

	client, err := getClient()
	if err != nil {
		t.Errorf("TestRepoCreate %s", err)
	}

	args := ArgsTwo{"midoks", "facegit"}
	var reply int
	err = client.Call("Repo.Create", args, &reply)
	if err != nil {
		t.Errorf("TestRepoCreate %s", err)
	}
	t.Log("TestRepoCreate ok")
}
