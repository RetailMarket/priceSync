package main

import (
	"Retail/priceSync/clients"
	"Retail/priceSync/jobs"
	"github.com/jasonlvhit/gocron"
)

func main() {
	clients.CreateClientConnection()
	defer clients.CloseConnections();

	// running job for sending update price record for approval.
	gocron.Every(5).Seconds().Do(jobs.SendUpdatePriceForApprovalJob)
	<-gocron.Start()
	defer gocron.Clear()
}
