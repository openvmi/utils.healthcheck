package main

import (
	"fmt"
	"net"

	"github.com/openvmi/utils.healthcheck/pkg/health"
	"github.com/openvmi/utils.healthcheck/pkg/registry"
	"github.com/openvmi/vmilog"
	"google.golang.org/grpc"
)

func main() {
	vmilog.SetLevel(vmilog.INFO)
	vmilog.EnableConsole(true)
	vmilog.EnableFile(false)
	h := health.NewHealthCheckService()
	registryUrl := "127.0.0.1:23"
	lis, err := net.Listen("tcp", "127.0.0.1:8854")
	if err != nil {
		fmt.Print(err)
		return
	}
	s := grpc.NewServer()
	health.RegisterService(s, h)
	tags := map[string][]string{
		"caps": {"hello"},
	}
	go registry.AutoRegistry(registryUrl, "127.0.0.1", "8854", "testService", tags, 34)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
