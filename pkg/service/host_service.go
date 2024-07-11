package service

import (
	"context"
	api "grpc-gateway/api/generated"
)

type HostService struct {
	api.UnimplementedHostServiceServer
}

func (s *HostService) SetHostname(ctx context.Context, req *api.SetHostnameRequest) (*api.SetHostnameResponse, error) {
	err := setHostname(req.Hostname)
	if err != nil {
		return &api.SetHostnameResponse{Success: false}, err
	}
	return &api.SetHostnameResponse{Success: true}, nil
}

func (s *HostService) GetDNSServers(ctx context.Context, req *api.GetDNSServersRequest) (*api.GetDNSServersResponse, error) {
	srvs, err := getDNSServers()
	if err != nil {
		return nil, err
	}
	return &api.GetDNSServersResponse{Servers: srvs}, nil
}

func (s *HostService) AddDNSServer(ctx context.Context, req *api.AddDNSServerRequest) (*api.AddDNSServerResponse, error) {
	err := addDNSServer(req.Server)
	if err != nil {
		return &api.AddDNSServerResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}
	return &api.AddDNSServerResponse{Success: true}, err
}

func (s *HostService) RemoveDNSServer(ctx context.Context, req *api.RemoveDNSServerRequest) (*api.RemoveDNSServerResponse, error) {
	err := removeDNSServer(req.Server)
	if err != nil {
		return &api.RemoveDNSServerResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}
	return &api.RemoveDNSServerResponse{Success: true}, nil
}
