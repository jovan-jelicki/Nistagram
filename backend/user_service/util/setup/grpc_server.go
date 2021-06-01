package setup

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/controllers"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
)

func GRPCServer(db *gorm.DB){
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpc_common.Users_service_address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	server, err := controllers.NewServer(db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Attach the Greeter service to the server
	userspb.RegisterUsersServer(s, server)
	// Serve gRPC server
	log.Println("Serving gRPC on " + grpc_common.Users_service_address)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Users_service_address)
	if err != nil {
		log.Fatalln(err) // TODO: Graceful shutdown
		return
	}

	gatewayMux := runtime.NewServeMux()
	// Register Greeter
	err = userspb.RegisterUsersHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    grpc_common.Users_gateway_address,
		Handler: tracer.TracingWrapper(gatewayMux),
	}

	log.Println("Serving gRPC-Gateway on " + grpc_common.Users_gateway_address)
	log.Fatalln(gwServer.ListenAndServe())
}
