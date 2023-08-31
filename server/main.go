package main

import (
	"context"
	"fmt"
	"net"

	controllers "github.com/umamaheswari76/netxd_customer_controller/controllers"
	"github.com/umamaheswari76/netxd_customer_dal/services"
	cst "github.com/umamaheswari76/netxd_customer_proto/customer"

	config "github.com/umamaheswari76/netxd_customer_config/config"
	constants "github.com/umamaheswari76/netxd_customer_config/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// mongodb initialization - we want collection and also the service function
// for performing operations
func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, constants.DatabaseName, "customers")
	controllers.CustomerService = services.InitializeCustomerService(CustomerCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)

	//grpc connection
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to Listen: %v", err)
		return
	}

	//grpc server creation
	s := grpc.NewServer()

	//sending the implemented functions and mapped values in the controller,
	//for registering in the grpc(sending controller with the new server instance)
	cst.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
