package clients

import (
	workflow "github.com/RetailMarket/workFlowClient"
	priceManager "github.com/RetailMarket/priceManagerClient"
	"log"
	"google.golang.org/grpc"
)

var PriceManagerClient priceManager.PriceManagerClient;

var priceManagerConn *grpc.ClientConn;

var WorkflowClient workflow.WorkFlowClient;

var workflowConn *grpc.ClientConn;

const (
	WORK_FLOW_ADDRESS = "localhost:4000"
	PRICE_MANAGER_ADDRESS = "localhost:3000"
)

func createWorkflowConnection() (workflow.WorkFlowClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(WORK_FLOW_ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return workflow.NewWorkFlowClient(conn), conn
}

func createPriceManagerClientConnection() (priceManager.PriceManagerClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(PRICE_MANAGER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return priceManager.NewPriceManagerClient(conn), conn
}

func CreateClientConnection() {
	WorkflowClient, workflowConn = createWorkflowConnection()
	PriceManagerClient, priceManagerConn = createPriceManagerClientConnection()
}

type clientDetails struct {
	client     interface{}
	connection *grpc.ClientConn
}

func CloseConnections() {
	workflowConn.Close();
	priceManagerConn.Close();
}