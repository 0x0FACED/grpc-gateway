package service

import (
	"os"
	"os/exec"
	"strings"
)

func setHostname(hostname string) error {
	cmd := exec.Command("hostnamectl", "set-hostname", hostname)
	return cmd.Run()
}

func getDNSServers() ([]string, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	var servers []string
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			servers = append(servers, strings.Fields(line)[1])
		}
	}
	return servers, nil
}

func addDNSServer(server string) error {
	f, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString("nameserver " + server + "\n"); err != nil {
		return err
	}
	return nil
}

func removeDNSServer(server string) error {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	newLines := []string{}
	for _, line := range lines {
		if !strings.Contains(line, server) {
			newLines = append(newLines, line)
		}
	}
	err = os.WriteFile("/etc/resolv.conf", []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return err
	}
	return nil
}
