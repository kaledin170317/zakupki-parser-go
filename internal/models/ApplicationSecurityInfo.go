package models

type ApplicationSecurityInfo struct {
	Required  bool    `bson:"required"`
	Amount    float64 `bson:"amount"`
	Currency  string  `bson:"currency"`
	Procedure string  `bson:"procedure"`
}
