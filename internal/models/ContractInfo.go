package models

type ContractInfo struct {
	MaxPrice float64 `bson:"max_price"`
	Currency string  `bson:"currency"`
}
