package main

import (
	clients "Retail/priceSync/clients"
	"Retail/priceSync/app/jobs"
)

func main() {
	clients.CreateClientConnection()
	defer clients.CloseConnections();

	// running job for sending update price record for approval.
	jobs.SendUpdatePriceForApprovalJob()
}
