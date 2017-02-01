package jobs

import (
	"log"
	workflow "github.com/RetailMarket/workFlowClient"
	priceManager "github.com/RetailMarket/priceManagerClient"
	"golang.org/x/net/context"
	"Retail/priceSync/clients"
)

func SendUpdatePriceForApprovalJob() {
	log.Println("Fetching pending update requests...")

	priceServiceResponse, err := clients.PriceManagerClient.PendingRecords(context.Background(), &priceManager.Request{})
	log.Printf("Processing records : %v\n", priceServiceResponse.Entries)

	if (err != nil) {
		log.Printf("Failed while fetching price update records\nError: %v", err)
		return
	}

	records := priceServiceResponse.GetEntries()
	if (len(records) != 0) {
		err := notifyWorkflowService(records)

		if err != nil {
			log.Printf("Failed to send pending update requests for approval \n err: %v\n", err)
			return
		}
		notifyPriceManagerService(records)
	}
}

func notifyWorkflowService(records []*priceManager.Entry) error {
	request := &workflow.Records{}
	for i := 0; i < len(records); i++ {
		priceObj := workflow.Entry{
			ProductId: records[i].ProductId,
			Version: records[i].Version}
		request.Entries = append(request.Entries, &priceObj)
	}
	workflowResponse, err := clients.WorkflowClient.NotifyRecordsPicked(context.Background(), request)
	log.Printf("Workflow Response: %s", workflowResponse.Message)
	return err;
}

func notifyPriceManagerService(records []*priceManager.Entry) {
	request := &priceManager.Records{Entries:records}
	response, err := clients.PriceManagerClient.NotifyRecordsPicked(context.Background(), request)
	if (err != nil) {
		log.Printf("Unable to change status to picked for entries %v\n Error: %v", records, err)
	} else {
		log.Printf("%v\n", response)
	}
}

