package main

import (
	clients "Retail/priceSync/clients"
	"Retail/priceSync/database"
	"Retail/priceSync/app/jobs"
)

func main() {
	database.Init();
	defer database.CloseDb();

	clients.CreateClientConnection()
	clients.CloseConnections();

	// running job for sending update price record for approval.
	jobs.SendUpdatePriceForApprovalJob()

}
