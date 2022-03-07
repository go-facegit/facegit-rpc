package app

import (
	"fmt"
	"net"
	// "net/http"
	// "net/rpc"

	"github.com/go-facegit/facegit-rpc/internal/app/repo"
	"github.com/go-facegit/facegit-rpc/pb"

	"google.golang.org/grpc"
)

func Start(port int) {

	// rpc.HandleHTTP()
	// repoLib := new(repo.Repo)
	// rpc.Register(repoLib)

	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	fmt.Println("rpc listen err: " + err.Error())
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("rpc listen err: " + err.Error())
	}
	rp := repo.Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterRepoServiceServer(grpcServer, &rp)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("grpc listen err: " + err.Error())
	}
}
