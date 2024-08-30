package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/tkmagesh/cisco-advgo-aug-2024/09-grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	appServiceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	// doRequestResponse(ctx, appServiceClient)
	doServerStreaming(ctx, appServiceClient)
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{
		Start: 100,
		End:   200,
	}
	clientStream, err := appServiceClient.GeneratePrimes(ctx, primeReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		resp, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("[doServerStreaming] All the primes numbers have been received!")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[doServerStreaming] Prime No = %d\n", resp.GetPrimeNo())
	}
}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	addResponse, err := appServiceClient.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Result :", addResponse.GetResult())
}
