package service_test

import (
	"net"
	"testing"

	"github.com/renatospaka/pc-book/pb"
	"github.com/renatospaka/pc-book/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")	// any available port
	require.NoError(t, err)
	
	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}
