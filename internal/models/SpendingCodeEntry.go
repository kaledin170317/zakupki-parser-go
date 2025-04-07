package models

type SpendingCodeEntry struct {
	ExpenseCode string  `bson:"expense_code"`
	ReceiptCode string  `bson:"receipt_code"`
	Total       float64 `bson:"total"`
	Year2025    float64 `bson:"year_2025"`
	Year2026    float64 `bson:"year_2026"`
	Year2027    float64 `bson:"year_2027"`
	Later       float64 `bson:"later_years"`
}
