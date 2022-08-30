package registry

import (
	"context"
	"time"

	"github.com/openvmi/protobuf_registry_go/pb"
	"google.golang.org/grpc"
)

func registe(ip string, port string, serviceName string, tags map[string][]string) bool {
	registryServiceUrl := ip + ":" + port
	conn, err := grpc.Dial(registryServiceUrl, grpc.WithInsecure())
	if err != nil {
		return false
	} else {
		c := pb.NewServiceRegistryClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reqTags := make([]*pb.RegisteRequest_MapFieldEntry, 0)
		for key, value := range tags {
			for _, eachValue := range value {
				_tag := pb.RegisteRequest_MapFieldEntry{
					Key:   key,
					Value: eachValue,
				}
				reqTags = append(reqTags, &_tag)
			}
		}
		_, err := c.Registe(ctx, &pb.RegisteRequest{
			ServiceName: serviceName,
			ServiceIp:   ip,
			ServicePort: port,
			ServiceTag:  reqTags,
		})
		if err != nil {
			return false
		}
	}
}

func AutoRegistry(ip string, port string, serviceName string, tags map[string][]string, sleepSec int64) {
	for {
		if !registe(ip, port, serviceName, tags) {
			time.Sleep(10 * time.Second)
		}
		time.Sleep(20 * time.Second)
	}
}
