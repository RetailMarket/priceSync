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
	time.Sleep(time.Second * 100000)
}

func sendPendingRequestForApproval() {
	for range time.Tick(time.Second * 5) {
		log.Println("Fetching pending update requests...")

		priceServiceResponse, err := clients.PriceManagerClient.GetPriceUpdateRecords(context.Background(), &priceManager.FetchRecordsRequest{})
		log.Printf("Processing records : %v\n", priceServiceResponse.Products)

		if (err != nil) {
			log.Printf("Failed while fetching price update records\nError: %v", err)
			continue
		}

		records := priceServiceResponse.GetProducts()
		if (len(records) != 0) {
			workflowRequest := createRequestForWorkflow(records)

			workflowResponse, err := clients.WorkflowClient.SaveUpdatePriceForApproval(context.Background(), workflowRequest)
			if err != nil {
				log.Printf("Failed to send pending update requests for approval \n err: %v\n", err)
				continue
			}
			log.Printf("Workflow Response: %s", workflowResponse.Message)
			updateStatus(records)
		}
	}
}

func createRequestForWorkflow(records []*priceManager.ProductEntry) *workflow.ProductsRequest {
	request := &workflow.ProductsRequest{}
	for i := 0; i < len(records); i++ {
		priceObj := workflow.Product{
			ProductId: records[i].ProductId,
			Version: records[i].Version}
		request.Products = append(request.Products, &priceObj)
	}
	return request
}

func updateStatus(records []*priceManager.ProductEntry) {
	request := &priceManager.ChangeStatusRequest{Products:records}
	response, err := clients.PriceManagerClient.ChangeStatusToPicked(context.Background(), request)
	if (err != nil) {
		log.Printf("Unable to change status to picked for entries %v\n Error: %v", records, err)
	} else {
		log.Printf("%v\n", response)
	}
}

