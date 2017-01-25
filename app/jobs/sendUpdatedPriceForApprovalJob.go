package jobs

import (
	"log"
	workflow "github.com/RetailMarket/workFlowClient"
	priceManager "github.com/RetailMarket/priceManagerClient"
	"golang.org/x/net/context"
	"time"
	"Retail/priceSync/clients"
)

func SendUpdatePriceForApprovalJob() {
	sendPendingRequestForApproval()
	time.Sleep(time.Second * 10000)
}

func sendPendingRequestForApproval() {
	for range time.Tick(time.Second * 10) {
		log.Println("Fetching pending update requests...")

		priceServiceResponse, err := clients.PriceManagerClient.GetPriceUpdateRecords(context.Background(), &priceManager.FetchRecordsRequest{});
		log.Printf("%v\n", priceServiceResponse.Products)

		if (err != nil) {
			log.Printf("Failed while fetching price update records\nError: %v", err);
			continue;
		}

		if (len(priceServiceResponse.GetProducts()) != 0) {
			workflowRequest := createRequestForWorkflow(priceServiceResponse)

			workflowResponse, err := clients.WorkflowClient.SaveUpdatePriceForApproval(context.Background(), workflowRequest)
			if err != nil {
				log.Printf("Failed to send pending update requests for approval \n err: %v\n", err)
				continue;
			}
			log.Printf("Response: %s", workflowResponse.Message)
		}

		//updateStatus(rows);
	}
}

func createRequestForWorkflow(response *priceManager.FetchRecordsResponse) *workflow.PriceUpdateRequest {
	request := &workflow.PriceUpdateRequest{}
	products := response.GetProducts()
	for i := 0; i < len(products); i++ {
		priceObj := workflow.Product{
			ProductId: products[i].ProductId,
			Version: products[i].Version}
		request.Products = append(request.Products, &priceObj)
	}
	return request;
}

//func updateStatus() {
//	clients.PriceManagerClient.ChangeStatusToPicked(context.Background(), )
//}

