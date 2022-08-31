package registry

import (
	"context"
	"time"

	"github.com/openvmi/protobuf_registry_go/pb"
	"github.com/openvmi/vmilog"
	"google.golang.org/grpc"
)

const (
	MODULE_NAME = "utils_healthcheck_registry"
)

func registe(registryUrl string, ip string, port string, serviceName string, tags map[string][]string) bool {
	conn, err := grpc.Dial(registryUrl, grpc.WithInsecure())
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
	return true
}

func AutoRegistry(registryUrl string, ip string, port string, serviceName string, tags map[string][]string, sleepSec int64) {
	vmilog.Info(MODULE_NAME, "AutoRegistry is tarted")
	for {
		vmilog.Info(MODULE_NAME, "Try to registe to:"+registryUrl)
		if !registe(registryUrl, ip, port, serviceName, tags) {
			time.Sleep(10 * time.Second)
			vmilog.Error(MODULE_NAME, "registe fail")
		} else {
			vmilog.Info(MODULE_NAME, "registe success")
		}
		time.Sleep(20 * time.Second)
	}
}
