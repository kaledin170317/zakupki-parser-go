package models

type ProcedureInfo struct {
	ApplicationDeadline string `bson:"application_deadline"`
	ProposalDate        string `bson:"proposal_date"`
	ResultsDate         string `bson:"results_date"`
}
