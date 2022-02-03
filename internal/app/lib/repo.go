package lib

import (
	"fmt"
)

type Repo struct {
	s string
}

type RepoCreateArgs struct {
	A, B int
}

func (rp *Repo) Create(args *RepoCreateArgs, reply *int) error {
	fmt.Println("这个方法执行了啊---嘿嘿--- Multiply ")
	*reply = 1
	return nil
}
