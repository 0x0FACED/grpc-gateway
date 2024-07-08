package util

import (
	"os/exec"
)

func SetHostname(hostname string) error {
	cmd := exec.Command("hostnamectl", "set-hostname", hostname)
	return cmd.Run()
}

func GetDNSServers() ([]string, error) {
	panic("not impl")
	cmd := exec.Command("systemd-resolve", "--status")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	servers := []string{}
	return servers, nil
}

func AddDNSServer(server string) error {
	panic("not impl")
	return nil
}

func RemoveDNSServer(server string) error {
	panic("not impl")
	return nil
}
