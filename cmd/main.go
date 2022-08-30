package main

import (
	"fmt"
	"net"

	"github.com/openvmi/utils.healthcheck/pkg/health"
	"google.golang.org/grpc"
)

func main() {
	h := health.NewHealthCheckService()
	lis, err := net.Listen("tcp", "127.0.0.1:8854")
	if err != nil {
		fmt.Print(err)
		return
	}
	s := grpc.NewServer()
	health.RegisterService(s, h)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
