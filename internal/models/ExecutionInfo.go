package models

type ExecutionInfo struct {
	StartDate    string `bson:"start_date"`
	EndDate      string `bson:"end_date"`
	BudgetFunded string `bson:"budget_funded"`
	OwnFunded    string `bson:"own_funded"`
}
