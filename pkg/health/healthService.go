package health

import (
	"context"

	"github.com/openvmi/protobuf_registry_go/pb"
	"google.golang.org/grpc"
)

type HealthCheckService struct {
	pb.UnimplementedServiceHealthCheckServer
	Handler IHealthHandler
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (h *HealthCheckService) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if h.Handler == nil {
		return &pb.HealthCheckResponse{Status: "suceess"}, nil
	}
	status := h.Handler.GetStatus()
	return &pb.HealthCheckResponse{Status: status}, nil
}

func RegisterService(server *grpc.Server, service *HealthCheckService) {
	pb.RegisterServiceHealthCheckServer(server, service)
}
