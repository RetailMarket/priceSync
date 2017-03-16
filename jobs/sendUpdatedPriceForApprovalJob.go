package jobs

import (
	"log"
	workflow "github.com/RetailMarket/workFlowClient"
	priceManager "github.com/RetailMarket/priceManagerClient"
	"golang.org/x/net/context"
	"Retail/priceSync/clients"
	"encoding/json"
)

func SendUpdatePriceForApprovalJob() {
	log.Println("Fetching pending update requests...")

	records, err := clients.PriceManagerClient.PendingRecords(context.Background(), &priceManager.Request{})

	if (err != nil) {
		log.Printf("Failed while fetching price update records\nError: %v", err)
		return
	}

	log.Printf("Processing records : %v\n", records.GetEntries())

	notifyService(records)
}
func notifyService(records *priceManager.Records) {
	if (len(records.GetEntries()) != 0) {
		err := notifyWorkflowService(records)

		if err != nil {
			log.Printf("Failed to send pending update requests for approval \n err: %v\n", err)
			return
		}
		notifyPriceManagerService(records)
	}
}

func notifyWorkflowService(records *priceManager.Records) error {
	request := &workflow.Records{}
	recordsInBytes, err := json.Marshal(records);
	if (err != nil) {
		log.Printf("Unable to marshal records: %v", records.GetEntries())
	}
	json.Unmarshal(recordsInBytes, request)

	workflowResponse, err := clients.WorkflowClient.NotifyRecordsPicked(context.Background(), request)
	log.Printf("Workflow Response: %s", workflowResponse.Message)
	return err;
}

func notifyPriceManagerService(records *priceManager.Records) {
	response, err := clients.PriceManagerClient.NotifyRecordsPicked(context.Background(), records)
	if (err != nil) {
		log.Printf("Unable to change status to picked for entries %v\n Error: %v", records, err)
	} else {
		log.Printf("%v\n", response)
	}
}

