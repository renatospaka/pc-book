package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/renatospaka/pc-book/pb"
	"github.com/renatospaka/pc-book/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	// conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	// create a laptop
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{Laptop: laptop}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)
}

// reading list to fix status and insecure conn of gRPC
// https://stackoverflow.com/questions/70482508/grpc-withinsecure-is-deprecated-use-insecure-newcredentials-instead
// https://stackoverflow.com/questions/63755105/how-can-i-unpack-the-grpc-status-details-error-in-golang
