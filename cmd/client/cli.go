package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "grpc-gateway/api/generated"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	hostname string
	server   string
)

func main() {
	var rootCmd = &cobra.Command{Use: "cli"}
	var setHostnameCmd = &cobra.Command{
		Use:   "set-hostname",
		Short: "Set the hostname",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := api.NewHostServiceClient(conn)

			response, err := c.SetHostname(context.Background(), &api.SetHostnameRequest{Hostname: hostname})
			if err != nil {
				log.Fatalf("could not set hostname: %v", err)
			}
			fmt.Printf("Success: %v\n", response.Success)
		},
	}
	setHostnameCmd.Flags().StringVarP(&hostname, "hostname", "n", "", "New hostname")
	rootCmd.AddCommand(setHostnameCmd)

	var getDNSServersCmd = &cobra.Command{
		Use:   "get-dns-servers",
		Short: "Get DNS servers",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := api.NewHostServiceClient(conn)

			response, err := c.GetDNSServers(context.Background(), &api.GetDNSServersRequest{})
			if err != nil {
				log.Fatalf("could not get DNS servers: %v", err)
			}
			fmt.Printf("DNS Servers: %v\n", response.Servers)
		},
	}
	rootCmd.AddCommand(getDNSServersCmd)

	var addDNSServerCmd = &cobra.Command{
		Use:   "add-dns-server",
		Short: "Add DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := api.NewHostServiceClient(conn)

			response, err := c.AddDNSServer(context.Background(), &api.AddDNSServerRequest{Server: server})
			if err != nil {
				log.Fatalf("could not add DNS server: %v", err)
			}
			fmt.Printf("Success: %v\n", response.Success)
		},
	}
	addDNSServerCmd.Flags().StringVarP(&server, "server", "s", "", "DNS server to add")
	rootCmd.AddCommand(addDNSServerCmd)

	var removeDNSServerCmd = &cobra.Command{
		Use:   "remove-dns-server",
		Short: "Remove DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := api.NewHostServiceClient(conn)

			response, err := c.RemoveDNSServer(context.Background(), &api.RemoveDNSServerRequest{Server: server})
			if err != nil {
				log.Fatalf("could not remove DNS server: %v", err)
			}
			fmt.Printf("Success: %v\n", response.Success)
		},
	}
	removeDNSServerCmd.Flags().StringVarP(&server, "server", "s", "", "DNS server to remove")
	rootCmd.AddCommand(removeDNSServerCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
