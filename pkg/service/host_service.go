package service

import (
	"context"
	"grpc-gateway/pkg/util"
)

type HostService struct {
}

func (s *HostService) SetHostname(ctx context.Context, req *api.SetHostnameRequest) (*api.SetHostnameResponse, error) {
	err := util.SetHostname(req.Hostname)
	if err != nil {
		return &api.SetHostnameResponse{Success: false}, err
	}
	return &api.SetHostnameResponse{Success: true}, nil
}
